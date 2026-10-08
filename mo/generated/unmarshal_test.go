package generated

import (
	"encoding/xml"
	"testing"
)

// Resposta de configResolveClass com hierarquia lsServer > vnicEther > vnicEtherIf.
const sample = `<configResolveClass cookie="x" response="yes" classId="lsServer">
  <outConfigs>
    <lsServer dn="org-root/ls-worker-04" name="worker-04" pnDn="sys/chassis-1/blade-4"
              srcTemplName="" operSrcTemplName="" assocState="associated" childAction="deleteNonPresent">
      <vnicEther dn="org-root/ls-worker-04/ether-eth0" name="eth0" addr="00:25:B5:B8:00:1A"
                 switchId="A" mtu="9000" identPoolName="mac-pool-a" cdnPropInSync="yes">
        <vnicEtherIf dn="org-root/ls-worker-04/ether-eth0/if-dados" name="dados" vnet="106" defaultNet="yes"/>
        <vnicEtherIf dn="org-root/ls-worker-04/ether-eth0/if-outra" name="outra" vnet="200" defaultNet="no"/>
      </vnicEther>
      <faultInst code="F0999"/>
    </lsServer>
  </outConfigs>
</configResolveClass>`

type resolveClass struct {
	OutConfigs struct {
		LsServers []LsServer `xml:"lsServer"`
	} `xml:"outConfigs"`
}

func TestUnmarshalHierarchy(t *testing.T) {
	var r resolveClass
	if err := xml.Unmarshal([]byte(sample), &r); err != nil {
		t.Fatal(err)
	}
	if len(r.OutConfigs.LsServers) != 1 {
		t.Fatalf("lsServer: got %d", len(r.OutConfigs.LsServers))
	}
	sp := r.OutConfigs.LsServers[0]
	if sp.PnDn != "sys/chassis-1/blade-4" || sp.AssocState != "associated" {
		t.Errorf("lsServer attrs: %+v", sp)
	}
	if len(sp.VnicEther) != 1 {
		t.Fatalf("vnicEther: got %d", len(sp.VnicEther))
	}
	v := sp.VnicEther[0]
	if v.Addr != "00:25:B5:B8:00:1A" || v.SwitchID != "A" || v.IdentPoolName != "mac-pool-a" || v.CdnPropInSync != "yes" {
		t.Errorf("vnicEther attrs: %+v", v)
	}
	if len(v.VnicEtherIf) != 2 {
		t.Fatalf("vnicEtherIf: got %d", len(v.VnicEtherIf))
	}
	if v.VnicEtherIf[0].Vnet != "106" || v.VnicEtherIf[0].DefaultNet != "yes" {
		t.Errorf("vnicEtherIf attrs: %+v", v.VnicEtherIf[0])
	}
}

func TestUnmarshalList(t *testing.T) {
	// Conteúdo de outConfigs de um configResolveClass de vnicEther.
	const out = `<outConfigs>
  <vnicEther dn="org-root/ls-a/ether-eth0" addr="00:25:B5:00:00:07">
    <vnicEtherIf name="dados" vnet="106" defaultNet="yes"/>
  </vnicEther>
  <vnicEther dn="org-root/ls-b/ether-eth0" addr="00:25:B5:00:00:08"/>
</outConfigs>`
	var l VnicEtherList
	if err := xml.Unmarshal([]byte(out), &l); err != nil {
		t.Fatal(err)
	}
	if len(l.Items) != 2 || l.Items[0].VnicEtherIf[0].Vnet != "106" {
		t.Fatalf("lista inesperada: %+v", l.Items)
	}
}
