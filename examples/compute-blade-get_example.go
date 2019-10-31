package main

import (
	"fmt"

	"go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	bladeModel := "UCSB-B200-M4"
	client, err := util.AaaLogin(endPoint, username, password)
	//client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if computeBlades, err := util.ComputeBladeGetAvailable(client, bladeModel); err != nil {
		fmt.Print(err)
	} else {
		for _, blade := range *computeBlades {
			fmt.Printf("%s:\n", blade.Dn)
			fmt.Printf("\tModel: %s\n", blade.Model)
			fmt.Printf("\tNumber of CPUs: %d\n", blade.NumOfCpus)
			fmt.Printf("\tNumber of Cores: %d\n", blade.NumOfCores)
			fmt.Printf("\tTotal Memory: %d\n", blade.TotalMemory)
			fmt.Printf("\tAvailability: %s\n", blade.Availability)

		}
	}
}
