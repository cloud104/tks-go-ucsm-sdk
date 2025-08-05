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
	client, err := util.AaaLogin(endPoint, username, password)
	//client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if lsBinding, err := util.SpGetBinding(client, spDn); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Dn: %s\n", lsBinding.Dn)
		fmt.Printf("\tAssigned Blade: %s\n", lsBinding.AssignedToDn)
	}
}
