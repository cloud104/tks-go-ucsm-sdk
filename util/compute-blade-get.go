package util

import (
	"go-ucsm-sdk/api"
	"go-ucsm-sdk/mo"
)

func ComputeBladeGetAvailable(c *api.Client, bladeModel string) (computeBlades *[]mo.ComputeBlade, err error) {
	var out mo.Blades
	filter := api.FilterAnd {
		Filters: []api.FilterAny {
			api.FilterEq {
				FilterProperty: api.FilterProperty {
					Class: "computeBlade",
					Property: "adminState",
					Value: "in-service",
				},
			},
			api.FilterEq {
				FilterProperty: api.FilterProperty {
					Class: "computeBlade",
					Property: "operState",
					Value: "unassociated",
				},
			},
			api.FilterEq {
				FilterProperty: api.FilterProperty {
					Class: "computeBlade",
					Property: "operability",
					Value: "operable",
				},
			},
			api.FilterEq {
				FilterProperty: api.FilterProperty {
					Class: "computeBlade",
					Property: "availability",
					Value: "available",
				},
			},
		},
	}
	if bladeModel != "" {
		modelFilter := api.FilterEq {
			FilterProperty: api.FilterProperty {
				Class: "computeBlade",
				Property: "model",
				Value: bladeModel,
			},
		}
		filter.Filters = append(filter.Filters, modelFilter)
	}
	req := api.ConfigResolveClassRequest {
		    Cookie: c.Cookie,
		    ClassId: "computeBlade",
		    InHierarchical: "false",
		    InFilter: filter,
	}
	if err = c.ConfigResolveClass(req, &out); err == nil {
		computeBlades = &out.ComputeBlades
	}
	return
}
