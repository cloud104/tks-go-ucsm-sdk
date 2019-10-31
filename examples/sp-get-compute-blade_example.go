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
	client, err := util.AaaLogin(endPoint, username, password)
	//client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if computeBlade, err := util.SpGetComputeBlade(client, spDn); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		if computeBlade != nil {
			fmt.Printf("Location: %s\n", computeBlade.Dn)
			fmt.Printf("Model: %s\n", computeBlade.Model)
			fmt.Printf("Number of CPUs: %d\n", computeBlade.NumOfCpus)
			fmt.Printf("Number of Cores: %d\n", computeBlade.NumOfCores)
			fmt.Printf("Total Memory: %d\n", computeBlade.TotalMemory)
		} else {
			fmt.Println("Error: no compute blade assigned")
		}
	}
}
