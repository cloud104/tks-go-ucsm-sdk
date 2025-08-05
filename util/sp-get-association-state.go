package util

import (
	"fmt"

	"github.com/gfalves87/tks-go-ucsm-sdk/api"
	"github.com/gfalves87/tks-go-ucsm-sdk/mo"
)

func SpGetAssociationState(c *api.Client, spDn string) (assocState string, err error) {
	var lsServers []*mo.LsServer
	if lsServers, err = ServerGet(c, spDn, "instance"); err == nil {
		if len(lsServers) > 0 {
			assocState = lsServers[0].AssocState
		} else {
			err = fmt.Errorf("ServerGet: no server %s found", spDn)
		}
	}
	return
}
