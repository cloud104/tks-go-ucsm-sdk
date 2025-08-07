package main

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	client, err := util.AaaLogin(endPoint, username, password)
	//client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	blades, err := util.GetAllComputeBlades(client)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	if len(blades) == 0 {
		fmt.Println("Nenhuma blade encontrada.")
		return
	}

	for _, blade := range blades {
		fmt.Println("---------------")
		fmt.Printf("Location: %s\n", blade.Dn)
		fmt.Printf("Model: %s\n", blade.Model)
		fmt.Printf("Serial: %s\n", blade.Serial)
		fmt.Printf("UUID: %s\n", blade.Uuid)
		fmt.Printf("PartNumber: %s\n", blade.PartNumber)
		fmt.Printf("Name: %s\n", blade.Name)
		fmt.Printf("UserLabel: %s\n", blade.UserLabel)
		fmt.Printf("Number of CPUs: %d\n", blade.NumOfCpus)
		fmt.Printf("Number of Cores: %d\n", blade.NumOfCores)
		fmt.Printf("Total Memory: %d\n", blade.TotalMemory)
	}
}
