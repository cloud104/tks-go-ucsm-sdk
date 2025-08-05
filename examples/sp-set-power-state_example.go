package main

import (
	"fmt"

	"github.com/gfalves87/tks-go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	spDn := "org-root/org-Linux/ls-phygymdev01-lab"
	powerState := "down"
	client, err := util.AaaLogin(endPoint, username, password)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if lsPower, err := util.SpSetPowerState(client, spDn, powerState); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Dn: %s\n", lsPower.Dn)
		fmt.Printf("\tPower state: %s\n", lsPower.State)
	}
}
