package ucsm

import (
	"fmt"
	"strings"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func ListAllVnicsEther(c *api.Client) (*[]mo.VnicEther, error) {
	var out mo.VnicsEther
	req := api.ConfigResolveClassRequest{
		Cookie:         c.Cookie,
		ClassID:        "vnicEther",
		InHierarchical: "false",
	}
	if err := c.ConfigResolveClass(req, &out); err != nil {
		return nil, fmt.Errorf("failed to resolve children classID <vnicEther>:  %w", err)
	}
	return &out.Vnics, nil
}

func GetVnicsEtherbyMAC(c *api.Client, macaddress string) (*[]mo.VnicEther, error) {
	var out mo.VnicsEther
	// remove ":" from macaddress break filter
	macaddress = strings.ToUpper(strings.ReplaceAll(macaddress, ":", ""))

	macFilter := api.FilterEq{
		FilterProperty: api.FilterProperty{
			Class:    "vnicEther",
			Property: "addr",
			Value:    macaddress,
		},
	}
	req := api.ConfigResolveClassRequest{
		Cookie:         c.Cookie,
		ClassID:        "vnicEther",
		InHierarchical: "true",
		InFilter:       macFilter,
	}

	if err := c.ConfigResolveClass(req, &out); err != nil {
		return nil, fmt.Errorf("failed to resolve classID <vnicEther>:  %w", err)
	}
	return &out.Vnics, nil
}
