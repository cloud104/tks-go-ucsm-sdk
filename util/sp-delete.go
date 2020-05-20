package util

import (
	"github.com/igor-feoktistov/go-ucsm-sdk/api"
	"github.com/igor-feoktistov/go-ucsm-sdk/mo"
)

func SpDelete(c *api.Client, spDn string) (error) {
	var out mo.LsServerPairs
	lsServerPairs := mo.LsServerPairs {
		Pairs: []mo.LsServerPair {
			{
				Key: spDn,
				LsServer: mo.LsServer {
					Dn: spDn,
					Status: "deleted",
				},
			},
		},
	}
	req := api.ConfigConfMosRequest {
		    Cookie: c.Cookie,
		    InHierarchical: "false",
		    InConfigs: lsServerPairs,
	}
	return c.ConfigConfMos(req, &out)
}
