```yaml
package main

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	spDn := "org-root/org-TPSP1"
	client, err := ucsm.AaaLogin(endPoint, username, password)
	//client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if computeBlade, err := ucsm.SpGetComputeBlade(client, spDn); err != nil {
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
	blades, err := ucsm.GetAllComputeBlades(client)
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
		fmt.Printf("Number of CPUs: %d\n", blade.NumOfCpus)
		fmt.Printf("Number of Cores: %d\n", blade.NumOfCores)
		fmt.Printf("Total Memory: %d\n", blade.TotalMemory)
		fmt.Printf("UUID: %s\n", blade.Uuid)
		fmt.Printf("Name: %s\n", blade.Name)
		fmt.Printf("UserLabel: %s\n", blade.UserLabel)
	}
}
```