package main

import (
	"fmt"

	"go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	sptDn := "org-root/org-Linux/ls-Linux-B200M3-SPT"
	spOrg := "org-root/org-Linux"
	spName := "phygymdev01-lab"
	client, err := util.AaaLogin(endPoint, username, password)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if lsServer, err := util.SptInstantiate(client, sptDn, spOrg, spName); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Dn: %s\n", lsServer.Dn)
                fmt.Printf("\tStatus: %s\n", lsServer.Status)
                fmt.Printf("\tAssign State: %s\n", lsServer.AssignState)
                fmt.Printf("\tAssociation State: %s\n", lsServer.AssocState)
                fmt.Printf("\tConfiguration State: %s\n", lsServer.ConfigState)
	}
}
