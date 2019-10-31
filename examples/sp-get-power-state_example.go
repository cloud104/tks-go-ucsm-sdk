package main

import (
	"fmt"

	"go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	spDn := "org-root/org-Linux/ls-phygymdev01-lab1"
	client, err := util.AaaLogin(endPoint, username, password)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if powerState, err := util.SpGetPowerState(client, spDn); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Dn: %s\n", spDn)
                fmt.Printf("\tPower state: %s\n", powerState)
	}
}
