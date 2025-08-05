package util

import (
	"github.com/gfalves87/tks-go-ucsm-sdk/api"
	"github.com/gfalves87/tks-go-ucsm-sdk/mo"
)

func SpGetComputeBlade(c *api.Client, spDn string) (computeBlade *mo.ComputeBlade, err error) {
	var out mo.Blades
	spFilter := api.FilterEq{
		FilterProperty: api.FilterProperty{
			Class:    "computeBlade",
			Property: "assignedToDn",
			Value:    spDn,
		},
	}
	req := api.ConfigResolveClassRequest{
		Cookie:         c.Cookie,
		ClassId:        "computeBlade",
		InHierarchical: "false",
		InFilter:       spFilter,
	}
	if err = c.ConfigResolveClass(req, &out); err == nil {
		for _, blade := range out.ComputeBlades {
			if blade.AssignedToDn == spDn {
				computeBlade = &blade
			}
		}
	}
	return
}

func GetAllComputeBlades(c *api.Client) ([]mo.ComputeBlade, error) {
	var out mo.Blades

	req := api.ConfigResolveClassRequest{
		Cookie:         c.Cookie,
		ClassId:        "computeBlade",
		InHierarchical: "false",
		// Sem InFilter para buscar todas
	}

	if err := c.ConfigResolveClass(req, &out); err != nil {
		return nil, err
	}

	return out.ComputeBlades, nil
}
