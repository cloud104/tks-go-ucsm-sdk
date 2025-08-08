package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func SpDelete(c *api.Client, spDn string) error {
	var out mo.LsServerPairs
	lsServerPairs := mo.LsServerPairs{
		Pairs: []mo.LsServerPair{
			{
				Key: spDn,
				LsServer: mo.LsServer{
					Dn:     spDn,
					Status: "deleted",
				},
			},
		},
	}
	req := api.ConfigConfMosRequest{
		Cookie:         c.Cookie,
		InHierarchical: "false",
		InConfigs:      lsServerPairs,
	}
	if err := c.ConfigConfMos(req, &out); err != nil {
		return fmt.Errorf("failed to delete service profile '%s': %w", spDn, err)
	}
	return nil
}
