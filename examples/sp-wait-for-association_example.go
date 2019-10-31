package main

import (
	"fmt"

	"go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	spDn := "org-root/org-Linux/ls-phygymdev01-lab"
	waitMax := 1800
	client, err := util.AaaLogin(endPoint, username, password)
	//client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if assocState, err := util.SpWaitForAssociation(client, spDn, waitMax); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("SP association state: %s\n", assocState)
	}
}
