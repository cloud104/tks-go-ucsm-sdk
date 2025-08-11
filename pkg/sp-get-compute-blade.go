package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
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
		ClassID:        "computeBlade",
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
	return computeBlade, fmt.Errorf("failed to resolve compute blade: %w", err)
}

func GetAllComputeBlades(c *api.Client) ([]mo.ComputeBlade, error) {
	var out mo.Blades

	req := api.ConfigResolveClassRequest{
		Cookie:         c.Cookie,
		ClassID:        "computeBlade",
		InHierarchical: "false",
		// Sem InFilter para buscar todas
	}

	if err := c.ConfigResolveClass(req, &out); err != nil {
		return nil, fmt.Errorf("failed to resolve compute blades: %w", err)
	}

	return out.ComputeBlades, nil
}
