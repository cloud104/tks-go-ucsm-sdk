package util

import (
	"time"

	"go-ucsm-sdk/api"
	"go-ucsm-sdk/mo"
)

func SpWaitForAssociation(c *api.Client, spDn string, waitMax int) (assocState string, err error) {
	waitUntil := time.Now().Add(time.Second * time.Duration(waitMax))
	var lsServers []*mo.LsServer
	for time.Now().Before(waitUntil) {
		if lsServers, err = ServerGet(c, spDn, "instance"); err == nil {
			assocState = lsServers[0].AssocState
			if assocState == "associating" {
				time.Sleep(2 * time.Second)
				continue
			}
			if assocState == "associated" {
				// Let to settle down on this state
				// to make sure it's not turning to "associating" again (known issue)
				for n := 0; n < 5 && assocState == "associated" && err == nil; n++ {
					time.Sleep(1 * time.Second)
					if lsServers, err = ServerGet(c, spDn, "instance"); err == nil {
						assocState = lsServers[0].AssocState
					}
				}
				if lsServers[0].AssocState == "associating" {
					continue
				}
			}
		}
		return
	}
	return
}
