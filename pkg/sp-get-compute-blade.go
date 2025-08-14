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
	err = c.ConfigResolveClass(req, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve classID <computeBlade>: %w", err)
	}
	for _, blade := range out.ComputeBlades {
		if blade.AssignedToDn == spDn {
			computeBlade = &blade
			return computeBlade, nil
		}
	}
	return nil, fmt.Errorf("no compute blade assigned with this filter: %w", err)
}
