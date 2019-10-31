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
	client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if vnicsEther, err := util.SpGetVnicsEther(client, spDn); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		for _, vnic := range *vnicsEther {
			fmt.Printf("Dn %s\n", vnic.Dn)
			fmt.Printf("\tName %s\n", vnic.Name)
			fmt.Printf("\tAddress %s\n", vnic.Addr)
		}
	}
}
