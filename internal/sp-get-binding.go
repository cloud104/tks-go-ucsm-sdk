package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SpGetBinding(c *api.Client, spDn string) (lsBinding *mo.LsBinding, err error) {
	var out mo.LsBindingMo
	req := api.ConfigResolveChildrenRequest{
		Cookie:         c.Cookie,
		InDn:           spDn,
		ClassID:        "lsBinding",
		InHierarchical: "false",
	}
	if err = c.ConfigResolveChildren(req, &out); err == nil {
		lsBinding = &out.LsBinding
	}
	return lsBinding, fmt.Errorf("failed to resolve children: %w", err)
}
