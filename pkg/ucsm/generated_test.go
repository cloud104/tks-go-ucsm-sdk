package ucsm

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
)

// fakeUCSM responde às chamadas da XML API com respostas fixas,
// passando pelo api.Client real (doPost + getInnerXML).
func fakeUCSM(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var root struct {
			XMLName xml.Name
			Dn      string `xml:"dn,attr"`
		}
		_ = xml.Unmarshal(body, &root)
		switch root.XMLName.Local {
		case "aaaLogin":
			io.WriteString(w, `<aaaLogin cookie="" response="yes" outCookie="c00k1e"/>`)
		case "configResolveClass":
			if !strings.Contains(string(body), `property="addr"`) {
				t.Errorf("filtro por addr ausente: %s", body)
			}
			io.WriteString(w, `<configResolveClass cookie="c00k1e" response="yes" classId="vnicEther"><outConfigs>
<vnicEther dn="org-root/ls-worker-04/ether-eth0" name="eth0" addr="00:25:B5:B8:00:1A" switchId="A" identPoolName="mac-a">
<vnicEtherIf dn="org-root/ls-worker-04/ether-eth0/if-dados" name="dados" vnet="106" defaultNet="yes"/>
</vnicEther></outConfigs></configResolveClass>`)
		case "configResolveDn":
			switch root.Dn {
			case "org-root/ls-worker-04":
				io.WriteString(w, `<configResolveDn dn="org-root/ls-worker-04" cookie="c00k1e" response="yes"><outConfig>
<lsServer dn="org-root/ls-worker-04" name="worker-04" pnDn="sys/chassis-1/blade-4" operSrcTemplName="org-root/ls-tpl-worker">
<vnicEther dn="org-root/ls-worker-04/ether-eth0" name="eth0" addr="00:25:B5:B8:00:1A" switchId="A" mtu="9000">
<vnicEtherIf name="dados" vnet="106" defaultNet="yes"/><vnicEtherIf name="pxe" vnet="124" defaultNet="no"/>
</vnicEther><lsPower state="up"/></lsServer></outConfig></configResolveDn>`)
			case "sys/chassis-1/blade-4":
				io.WriteString(w, `<configResolveDn dn="sys/chassis-1/blade-4" cookie="c00k1e" response="yes"><outConfig>
<computeBlade dn="sys/chassis-1/blade-4" serial="FCH1234ABCD" model="UCSB-B200-M5" numOfCpus="2" assignedToDn="org-root/ls-worker-04"/>
</outConfig></configResolveDn>`)
			default: // DN inexistente: UCSM devolve outConfig vazio
				io.WriteString(w, `<configResolveDn dn="`+root.Dn+`" cookie="c00k1e" response="yes"><outConfig/></configResolveDn>`)
			}
		default:
			t.Errorf("chamada inesperada: %s", root.XMLName.Local)
		}
	}))
}

func TestGeneratedHelpers(t *testing.T) {
	srv := fakeUCSM(t)
	defer srv.Close()

	c, err := api.NewClient(api.Config{Endpoint: srv.URL + "/"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.AaaLogin(); err != nil {
		t.Fatal(err)
	}

	vnics, err := FindVnicsEtherByMAC(c, "00:25:b5:b8:00:1a")
	if err != nil {
		t.Fatal(err)
	}
	if len(vnics) != 1 || vnics[0].VnicEtherIf[0].Vnet != "106" {
		t.Fatalf("vnics: %+v", vnics)
	}

	sp, err := GetServiceProfile(c, "org-root/ls-worker-04")
	if err != nil {
		t.Fatal(err)
	}
	if sp.PnDn != "sys/chassis-1/blade-4" || sp.OperSrcTemplName == "" || len(sp.VnicEther[0].VnicEtherIf) != 2 {
		t.Fatalf("sp: %+v", sp)
	}

	b, err := GetComputeBlade(c, sp.PnDn)
	if err != nil {
		t.Fatal(err)
	}
	if b.Serial != "FCH1234ABCD" || b.AssignedToDn != sp.Dn {
		t.Fatalf("blade: %+v", b)
	}

	if _, err := GetComputeRackUnit(c, "sys/rack-unit-9"); err == nil {
		t.Fatal("esperava erro para DN inexistente")
	}
}
