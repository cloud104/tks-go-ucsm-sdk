package main

import (
	"fmt"

	"github.com/gfalves87/tks-go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	//srvDnFilter := "org-root/org-Linux/ls-phygymdev01-lab"
	//srvType := "instance"
	srvDnFilter := "org-root/org-Linux/ls-Linux-B200M3-SPT"
	srvType := "initial-template"
	client, err := util.AaaLogin(endPoint, username, password)
	//client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if lsServers, err := util.ServerGet(client, srvDnFilter, srvType); err != nil {
		fmt.Print(err)
	} else {
		for _, lsServer := range lsServers {
			if lsServer.Type == "instance" {
				fmt.Printf("SP: %s\n", lsServer.Dn)
				fmt.Printf("\tType: %s\n", lsServer.Type)
				fmt.Printf("\tAssociation State: %s\n", lsServer.AssocState)
				fmt.Printf("\tConfiguration State: %s\n", lsServer.ConfigState)
				fmt.Printf("\tAssign State: %s\n", lsServer.AssignState)
				if lsServer.AssignState == "assigned" {
					fmt.Printf("\tAssigned blade: %s\n", lsServer.PnDn)
				}
			}
			if lsServer.Type == "initial-template" {
				fmt.Printf("SPT: %s\n", lsServer.Dn)
				fmt.Printf("\tType: %s\n", lsServer.Type)
			}
		}
	}
}
