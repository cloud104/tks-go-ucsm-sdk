package util

import (
	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SpGetVnicsIScsi(c *api.Client, spDn string) (vnicsIScsi *[]mo.VnicIScsi, err error) {
	var out mo.VnicsIScsi
	req := api.ConfigResolveChildrenRequest{
		Cookie:         c.Cookie,
		InDn:           spDn,
		ClassId:        "vnicIScsi",
		InHierarchical: "true",
	}
	if err = c.ConfigResolveChildren(req, &out); err == nil {
		vnicsIScsi = &out.Vnics
	}
	return
}
