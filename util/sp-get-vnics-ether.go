package util

import (
	"go-ucsm-sdk/api"
	"go-ucsm-sdk/mo"
)

func SpGetVnicsEther(c *api.Client, spDn string) (vnicsEther *[]mo.VnicEther, err error) {
	var out mo.VnicsEther
	req := api.ConfigResolveChildrenRequest {
		    Cookie: c.Cookie,
		    InDn: spDn,
		    ClassId: "vnicEther",
		    InHierarchical: "true",
	}
	if err = c.ConfigResolveChildren(req, &out); err == nil {
		vnicsEther = &out.Vnics
	}
	return
}
