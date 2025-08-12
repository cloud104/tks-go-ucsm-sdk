package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
)

func SpGetAssociationState(c *api.Client, spDn string) (string, error) {
	var assocState string
	if lsServers, err := ServerGet(c, spDn, "instance"); err == nil {
		if len(lsServers) == 0 {
			err := fmt.Errorf("ServerGet: no server %s found", spDn)
			return "", fmt.Errorf("failed to get association state: %w", err)
		}
		assocState = lsServers[0].AssocState
	}
	return assocState, nil
}
