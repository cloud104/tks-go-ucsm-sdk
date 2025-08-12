package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SpGetVnicsIScsi(c *api.Client, spDn string) (*[]mo.VnicIScsi, error) {
	var out mo.VnicsIScsi
	req := api.ConfigResolveChildrenRequest{
		Cookie:         c.Cookie,
		InDn:           spDn,
		ClassID:        "vnicIScsi",
		InHierarchical: "true",
	}
	if err := c.ConfigResolveChildren(req, &out); err != nil {
		return nil, fmt.Errorf("failed to resolve children: %w", err)
	}
	return &out.Vnics, nil
}
