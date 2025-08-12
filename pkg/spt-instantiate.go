package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SptInstantiate(c *api.Client, sptDn string, spOrg string, spName string) (*mo.LsServer, error) {
	var out mo.LsServerMo
	req := api.LsInstantiateNNamedTemplateRequest{
		Cookie:      c.Cookie,
		Dn:          sptDn,
		InTargetOrg: spOrg,
		InNameSet: []api.Dn{
			{
				Value: spName,
			},
		},
		InHierarchical: "false",
	}
	err := c.LsInstantiateNNamedTemplate(req, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate named template: %w", err)
	}
	return &out.LsServer, nil
}
