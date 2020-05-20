package util

import (
	"fmt"
	"time"

	"github.com/igor-feoktistov/go-ucsm-sdk/api"
	"github.com/igor-feoktistov/go-ucsm-sdk/mo"
)

func SpWaitForAssociation(c *api.Client, spDn string, waitMax int) (assocState string, err error) {
	waitUntil := time.Now().Add(time.Second * time.Duration(waitMax))
	var lsServers []*mo.LsServer
	for time.Now().Before(waitUntil) {
		if lsServers, err = ServerGet(c, spDn, "instance"); err == nil {
			if len(lsServers) == 0 {
				err = fmt.Errorf("ServerGet: no server %s found", spDn)
				return
			}
			assocState = lsServers[0].AssocState
			if assocState == "associating" {
				time.Sleep(2 * time.Second)
				continue
			}
			// Let to settle down on the state different from "associating"
			// to make sure it's not turning to "associating" again from
			// "failure" or "associated" (known issue)
			for n := 0; n < 5 && assocState != "associating" && err == nil; n++ {
				time.Sleep(1 * time.Second)
				if lsServers, err = ServerGet(c, spDn, "instance"); err == nil {
					assocState = lsServers[0].AssocState
				}
			}
			if lsServers[0].AssocState == "associating" {
				continue
			}
		}
		return
	}
	return
}
