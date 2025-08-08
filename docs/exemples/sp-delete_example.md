package main

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	spDn := "org-root/org-Linux/ls-phygymdev01-lab"
	client, err := ucsm.AaaLogin(endPoint, username, password)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if err := ucsm.SpDelete(client, spDn); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Deleted SP %s\n", spDn)
	}
}
