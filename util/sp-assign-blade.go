package util

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gfalves87/tks-go-ucsm-sdk/api"
	"github.com/gfalves87/tks-go-ucsm-sdk/mo"
)

func SpAssignBlade(c *api.Client, spDn string, pnDn string) (lsBinding *mo.LsBinding, err error) {
	var out mo.LsBindingPairs
	var lsServerPairs mo.LsServerPairs
	var lsIssuesPairs mo.LsIssuesPairs
	inMoPairs := mo.MoPairs{
		Pairs: []mo.MoPair{
			mo.MoPair{
				Key: spDn + "/pn",
				Mo: mo.LsBinding{
					Dn:   spDn + "/pn",
					PnDn: pnDn,
				},
			},
			mo.MoPair{
				Key: spDn + "/pn-req",
				Mo: mo.LsRequirement{
					Dn:     spDn + "/pn-req",
					Status: "deleted",
				},
			},
		},
	}
	req := api.ConfigEstimateImpactRequest{
		Cookie:    c.Cookie,
		InConfigs: inMoPairs,
	}
	if err = c.ConfigEstimateImpact(req, &lsServerPairs, &lsIssuesPairs); err == nil {
		var lsServer *mo.LsServer
		var lsIssues *mo.LsIssues
		for _, pair := range lsServerPairs.Pairs {
			if pair.LsServer.XMLName.Local == "lsServer" && pair.Key == spDn {
				lsServer = &pair.LsServer
				break
			}
		}
		for _, pair := range lsIssuesPairs.Pairs {
			if pair.LsIssues.XMLName.Local == "lsIssues" && pair.Key == spDn+"/config-issue" {
				lsIssues = &pair.LsIssues
				break
			}
		}
		if lsServer != nil && lsIssues != nil {
			var issues []string
			if lsIssues.IscsiConfigIssues != "" {
				issues = append(issues, "IscsiConfigIssues: "+lsIssues.IscsiConfigIssues)
			}
			if lsIssues.NetworkConfigIssues != "" {
				issues = append(issues, "NetworkConfigIssues: "+lsIssues.NetworkConfigIssues)
			}
			if lsIssues.ServerConfigIssues != "" {
				issues = append(issues, "ServerConfigIssues: "+lsIssues.ServerConfigIssues)
			}
			if lsIssues.ServerExtdConfigIssues != "" {
				issues = append(issues, "ServerExtdConfigIssues: "+lsIssues.ServerExtdConfigIssues)
			}
			if lsIssues.StorageConfigIssues != "" {
				issues = append(issues, "StorageConfigIssues: "+lsIssues.StorageConfigIssues)
			}
			if lsIssues.VnicConfigIssues != "" {
				issues = append(issues, "VnicConfigIssues: "+lsIssues.VnicConfigIssues)
			}
			if lsServer.ConfigState != "applied" {
				err = errors.New(fmt.Sprintf("configState=\"%s\", configQualifier=\"%s\", issues=\"%s\"", lsServer.ConfigState, lsServer.ConfigQualifier, strings.Join(issues, ", ")))
			}
		}
	}
	if err == nil {
		req := api.ConfigConfMosRequest{
			Cookie:         c.Cookie,
			InHierarchical: "false",
			InConfigs:      inMoPairs,
		}
		if err = c.ConfigConfMos(req, &out); err == nil {
			if lsBinding, err = SpGetBinding(c, spDn); err == nil {
				if lsBinding.OperState != "used" {
					err = errors.New(fmt.Sprintf("operState=\"%s\", issues=\"%s\"", lsBinding.OperState, lsBinding.Issues))
				}
			}
		}
	}
	return
}
