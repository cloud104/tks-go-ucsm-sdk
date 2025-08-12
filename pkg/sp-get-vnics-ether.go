package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SpGetVnicsEther(c *api.Client, spDn string) (*[]mo.VnicEther, error) {
	var out mo.VnicsEther
	req := api.ConfigResolveChildrenRequest{
		Cookie:         c.Cookie,
		InDn:           spDn,
		ClassID:        "vnicEther",
		InHierarchical: "true",
	}
	if err := c.ConfigResolveChildren(req, &out); err != nil {
		return nil, fmt.Errorf("failed to resolve children classID <vnicEther>:  %w", err)
	}
	return &out.Vnics, nil
}
