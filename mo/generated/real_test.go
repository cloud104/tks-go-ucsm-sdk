package generated

import (
	"encoding/xml"
	"os"
	"testing"
)

// Respostas reais do UCSM 4.2(3p), capturadas com o debug do client
// (cookie removido). Ver testdata/.

func loadFixture(t *testing.T, name string, out any) {
	t.Helper()
	raw, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}
	if err := xml.Unmarshal(raw, out); err != nil {
		t.Fatalf("%s: %v", name, err)
	}
}

func TestRealServiceProfile(t *testing.T) {
	var resp struct {
		OutConfig LsServerList `xml:"outConfig"`
	}
	loadFixture(t, "lsServer_real.xml", &resp)

	if len(resp.OutConfig.Items) != 1 {
		t.Fatalf("lsServer: got %d", len(resp.OutConfig.Items))
	}
	sp := resp.OutConfig.Items[0]
	if sp.PnDn != "sys/chassis-1/blade-1" || sp.OperSrcTemplName == "" || sp.AssocState != "associated" {
		t.Errorf("lsServer: pnDn=%q template=%q assoc=%q", sp.PnDn, sp.OperSrcTemplName, sp.AssocState)
	}

	vnics := map[string]VnicEther{}
	for _, v := range sp.VnicEther {
		vnics[v.Name] = v
	}
	if len(vnics) != 5 {
		t.Fatalf("vnicEther: got %d", len(vnics))
	}

	// regiond-a: 2672 nativa e 2606 (L3Out) tagged.
	r, ok := vnics["regiond-a"]
	if !ok {
		t.Fatal("vNIC regiond-a ausente")
	}
	if r.Addr != "00:25:B5:00:00:07" || r.SwitchID != "A-B" || r.IdentPoolName == "" {
		t.Errorf("regiond-a: addr=%q switch=%q pool=%q", r.Addr, r.SwitchID, r.IdentPoolName)
	}
	native := map[string]string{}
	for _, vl := range r.VnicEtherIf {
		native[vl.Vnet] = vl.DefaultNet
	}
	if native["2672"] != "yes" || native["2606"] != "no" {
		t.Errorf("regiond-a VLANs: %v", native)
	}
}

func TestRealComputeBlade(t *testing.T) {
	var resp struct {
		OutConfig ComputeBladeList `xml:"outConfig"`
	}
	loadFixture(t, "computeBlade_real.xml", &resp)

	if len(resp.OutConfig.Items) != 1 {
		t.Fatalf("computeBlade: got %d", len(resp.OutConfig.Items))
	}
	b := resp.OutConfig.Items[0]
	if b.Serial == "" || b.Model != "UCSB-B200-M5" || b.AssignedToDn == "" {
		t.Errorf("blade: serial=%q model=%q assignedTo=%q", b.Serial, b.Model, b.AssignedToDn)
	}
}

func TestRealVnicByMAC(t *testing.T) {
	var resp struct {
		OutConfigs VnicEtherList `xml:"outConfigs"`
	}
	loadFixture(t, "vnicEther_by_mac.xml", &resp)

	if len(resp.OutConfigs.Items) != 1 || resp.OutConfigs.Items[0].Addr != "00:25:B5:00:00:07" {
		t.Fatalf("vnicEther: %+v", resp.OutConfigs.Items)
	}
	if len(resp.OutConfigs.Items[0].VnicEtherIf) == 0 {
		t.Error("vnicEtherIf ausente na consulta hierárquica")
	}
}
