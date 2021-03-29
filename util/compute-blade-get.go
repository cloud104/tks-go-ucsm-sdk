package util

import (
	"regexp"

	"github.com/igor-feoktistov/go-ucsm-sdk/api"
	"github.com/igor-feoktistov/go-ucsm-sdk/mo"
)

const (
	wcardChars = `[\^\*\+\.\(\)\[\]\|\?\$]`
)

type BladeSpec struct {
	Dn           string `yaml:"dn,omitempty" json:"dn,omitempty" xml:"dn,omitempty"`
	Model        string `yaml:"model,omitempty" json:"model,omitempty" xml:"model,omitempty"`
	Serial       string `yaml:"serial,omitempty" json:"serial,omitempty" xml:"serial,omitempty"`
	NumOfCpus    string `yaml:"numOfCpus,omitempty" json:"numOfCpus,omitempty" xml:"numOfCpus,omitempty"`
	NumOfCores   string `yaml:"numOfCores,omitempty" json:"numOfCores,omitempty" xml:"numOfCores,omitempty"`
	NumOfThreads string `yaml:"numOfThreads,omitempty" json:"numOfThreads,omitempty" xml:"numOfThreads,omitempty"`
	TotalMemory  string `yaml:"totalMemory,omitempty" json:"totalMemory,omitempty" xml:"totalMemory,omitempty"`
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
		var matched bool
		if bladeSpec.Dn != "" {
			if matched, err = regexp.MatchString(wcardChars, bladeSpec.Dn); err != nil {
				return
			}
			if matched {
				filter.Filters = append(filter.Filters, api.FilterWildcard {
									    FilterProperty: api.FilterProperty {
										    Class: "computeBlade",
										    Property: "dn",
										    Value: bladeSpec.Dn,
										},
									})
			} else {
				filter.Filters = append(filter.Filters, api.FilterEq {
									    FilterProperty: api.FilterProperty {
										    Class: "computeBlade",
										    Property: "dn",
										    Value: bladeSpec.Dn,
									    },
									})
			}
		}
		if bladeSpec.Model != "" {
			if matched, err = regexp.MatchString(wcardChars, bladeSpec.Model); err != nil {
				return
			}
			if matched {
				filter.Filters = append(filter.Filters, api.FilterWildcard {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "model",
											Value: bladeSpec.Model,
										},
									})
			} else {
				filter.Filters = append(filter.Filters, api.FilterEq {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "model",
											Value: bladeSpec.Model,
										},
									})
			}
		}
		rangeRegexp, _ := regexp.Compile(`([0-9]+)\s*-\s*([0-9]+)`)
		if bladeSpec.NumOfCpus != "" {
			cpusRange := rangeRegexp.FindStringSubmatch(bladeSpec.NumOfCpus)
			if len(cpusRange) == 3 {
				filter.Filters = append(filter.Filters, api.FilterGe {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "numOfCpus",
											Value: cpusRange[1],
										},
									})
				filter.Filters = append(filter.Filters, api.FilterLe {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "numOfCpus",
											Value: cpusRange[2],
										},
									})
			} else {
				filter.Filters = append(filter.Filters, api.FilterEq {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "numOfCpus",
											Value: bladeSpec.NumOfCpus,
										},
									})
			}
		}
		if bladeSpec.NumOfCores != "" {
			coresRange := rangeRegexp.FindStringSubmatch(bladeSpec.NumOfCores)
			if len(coresRange) == 3 {
				filter.Filters = append(filter.Filters, api.FilterGe {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "numOfCores",
											Value: coresRange[1],
										},
									})
				filter.Filters = append(filter.Filters, api.FilterLe {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "numOfCores",
											Value: coresRange[2],
										},
									})
			} else {
				filter.Filters = append(filter.Filters, api.FilterEq {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "numOfCores",
											Value: bladeSpec.NumOfCores,
										},
									})
			}
		}
		if bladeSpec.NumOfThreads != "" {
			threadsRange := rangeRegexp.FindStringSubmatch(bladeSpec.NumOfThreads)
			if len(threadsRange) == 3 {
				filter.Filters = append(filter.Filters, api.FilterGe {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "numOfThreads",
											Value: threadsRange[1],
										},
									})
				filter.Filters = append(filter.Filters, api.FilterLe {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "numOfThreads",
											Value: threadsRange[2],
										},
									})
			} else {
				filter.Filters = append(filter.Filters, api.FilterEq {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "numOfThreads",
											Value: bladeSpec.NumOfThreads,
										},
									})
			}
		}
		if bladeSpec.TotalMemory != "" {
			memRange := rangeRegexp.FindStringSubmatch(bladeSpec.TotalMemory)
			if len(memRange) == 3 {
				filter.Filters = append(filter.Filters, api.FilterGe {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "totalMemory",
											Value: memRange[1],
										},
									})
				filter.Filters = append(filter.Filters, api.FilterLe {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "totalMemory",
											Value: memRange[2],
										},
									})
			} else {
				filter.Filters = append(filter.Filters, api.FilterEq {
										FilterProperty: api.FilterProperty {
											Class: "computeBlade",
											Property: "totalMemory",
											Value: bladeSpec.TotalMemory,
										},
									})
			}
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
