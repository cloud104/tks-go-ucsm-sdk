package ucsm

// Helpers de leitura que devolvem as structs geradas (mo/generated).
// Convivem com os helpers baseados em mo; migre os chamadores aos poucos.

import (
	"fmt"
	"strings"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo/generated"
)

// resolveDn busca um único objeto pelo DN e decodifica o outConfig em out
// (um *generated.XList). Se o DN não existir, out.Items fica vazio.
func resolveDn(c *api.Client, dn string, hierarchical bool, out any) error {
	req := api.ConfigResolveDnRequest{
		Cookie:         c.Cookie,
		Dn:             dn,
		InHierarchical: fmt.Sprintf("%t", hierarchical),
	}
	if err := c.ConfigResolveDn(req, out); err != nil {
		return fmt.Errorf("configResolveDn %q: %w", dn, err)
	}
	return nil
}

// FindVnicsEtherByMAC devolve as vNICs com o MAC informado, já com as
// VLANs (vnicEtherIf) de cada uma.
func FindVnicsEtherByMAC(c *api.Client, mac string) ([]generated.VnicEther, error) {
	// Mesma normalização usada em GetVnicsEtherbyMAC.
	mac = strings.ToUpper(strings.ReplaceAll(mac, ":", ""))

	var out generated.VnicEtherList
	req := api.ConfigResolveClassRequest{
		Cookie:         c.Cookie,
		ClassID:        "vnicEther",
		InHierarchical: "true",
		InFilter: api.FilterEq{
			FilterProperty: api.FilterProperty{Class: "vnicEther", Property: "addr", Value: mac},
		},
	}
	if err := c.ConfigResolveClass(req, &out); err != nil {
		return nil, fmt.Errorf("configResolveClass <vnicEther> addr=%q: %w", mac, err)
	}
	return out.Items, nil
}

// GetServiceProfile devolve o service profile (lsServer) do DN informado,
// com as vNICs e as VLANs.
func GetServiceProfile(c *api.Client, dn string) (*generated.LsServer, error) {
	var out generated.LsServerList
	if err := resolveDn(c, dn, true, &out); err != nil {
		return nil, err
	}
	if len(out.Items) == 0 {
		return nil, fmt.Errorf("service profile %q não encontrado", dn)
	}
	return &out.Items[0], nil
}

// GetComputeBlade devolve o blade no DN físico informado (lsServer.PnDn).
func GetComputeBlade(c *api.Client, dn string) (*generated.ComputeBlade, error) {
	var out generated.ComputeBladeList
	if err := resolveDn(c, dn, false, &out); err != nil {
		return nil, err
	}
	if len(out.Items) == 0 {
		return nil, fmt.Errorf("blade %q não encontrado", dn)
	}
	return &out.Items[0], nil
}

// GetComputeRackUnit devolve o servidor rack no DN físico informado.
func GetComputeRackUnit(c *api.Client, dn string) (*generated.ComputeRackUnit, error) {
	var out generated.ComputeRackUnitList
	if err := resolveDn(c, dn, false, &out); err != nil {
		return nil, err
	}
	if len(out.Items) == 0 {
		return nil, fmt.Errorf("rack unit %q não encontrado", dn)
	}
	return &out.Items[0], nil
}
