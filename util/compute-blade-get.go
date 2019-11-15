package util

import (
	"strconv"

	"go-ucsm-sdk/api"
	"go-ucsm-sdk/mo"
)

type BladeSpec struct {
	Dn          string `yaml:"dn,omitempty"`
	Model       string `yaml:"model,omitempty"`
	NumOfCpus   int    `yaml:"numOfCpus,omitempty"`
	NumOfCores  int    `yaml:"numOfCores,omitempty"`
	TotalMemory int    `yaml:"totalMemory,omitempty"`
}

func ComputeBladeGetAvailable(c *api.Client, bladeSpec *BladeSpec) (computeBlades *[]mo.ComputeBlade, err error) {
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
	if bladeSpec != nil {
		if bladeSpec.Dn != "" {
			dnFilter := api.FilterEq {
				FilterProperty: api.FilterProperty {
					Class: "computeBlade",
					Property: "dn",
					Value: bladeSpec.Dn,
				},
			}
			filter.Filters = append(filter.Filters, dnFilter)
		}
		if bladeSpec.Model != "" {
			modelFilter := api.FilterEq {
				FilterProperty: api.FilterProperty {
					Class: "computeBlade",
					Property: "model",
					Value: bladeSpec.Model,
				},
			}
			filter.Filters = append(filter.Filters, modelFilter)
		}
		if bladeSpec.NumOfCpus > 0 {
			numOfCpusFilter := api.FilterEq {
				FilterProperty: api.FilterProperty {
					Class: "computeBlade",
					Property: "numOfCpus",
					Value: strconv.Itoa(bladeSpec.NumOfCpus),
				},
			}
			filter.Filters = append(filter.Filters, numOfCpusFilter)
		}
		if bladeSpec.NumOfCores > 0 {
			numOfCoresFilter := api.FilterEq {
				FilterProperty: api.FilterProperty {
					Class: "computeBlade",
					Property: "numOfCores",
					Value: strconv.Itoa(bladeSpec.NumOfCores),
				},
			}
			filter.Filters = append(filter.Filters, numOfCoresFilter)
		}
		if bladeSpec.TotalMemory > 0 {
			totalMemoryFilter := api.FilterEq {
				FilterProperty: api.FilterProperty {
					Class: "computeBlade",
					Property: "totalMemory",
					Value: strconv.Itoa(bladeSpec.TotalMemory),
				},
			}
			filter.Filters = append(filter.Filters, totalMemoryFilter)
		}
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
