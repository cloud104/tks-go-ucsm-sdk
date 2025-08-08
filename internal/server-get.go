package ucsm

import (
	"fmt"
	"strings"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func ServerGet(c *api.Client, dnFilter string, lsType string) (lsServers []*mo.LsServer, err error) {
	var startOrg string
	if ndx := strings.LastIndex(dnFilter, "/"); ndx != -1 {
		startOrg = dnFilter[:ndx]
	} else {
		startOrg = "org-root"
	}
	var out mo.LsServerPairs
	filter := api.FilterAnd{
		Filters: []api.FilterAny{
			api.FilterWildcard{
				FilterProperty: api.FilterProperty{
					Class:    "lsServer",
					Property: "dn",
					Value:    dnFilter,
				},
			},
			api.FilterEq{
				FilterProperty: api.FilterProperty{
					Class:    "lsServer",
					Property: "type",
					Value:    lsType,
				},
			},
		},
	}
	req := api.OrgResolveElementsRequest{
		Cookie:         c.Cookie,
		Dn:             startOrg,
		InHierarchical: "false",
		InClass:        "lsServer",
		InSingleLevel:  "true",
		InFilter:       filter,
	}
	if err = c.OrgResolveElements(req, &out); err == nil {
		for _, pair := range out.Pairs {
			lsServers = append(lsServers, &pair.LsServer)
		}
	}
	return lsServers, fmt.Errorf("failed to resolve elements: %w", err)
}
