package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SpGetVnicsEther(c *api.Client, spDn string) (vnicsEther *[]mo.VnicEther, err error) {
	var out mo.VnicsEther
	req := api.ConfigResolveChildrenRequest{
		Cookie:         c.Cookie,
		InDn:           spDn,
		ClassID:        "vnicEther",
		InHierarchical: "true",
	}
	if err = c.ConfigResolveChildren(req, &out); err == nil {
		vnicsEther = &out.Vnics
	}
	return vnicsEther, fmt.Errorf("failed to resolve children: %w", err)
}
