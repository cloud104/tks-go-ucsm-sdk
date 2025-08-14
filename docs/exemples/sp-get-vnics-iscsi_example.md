```go
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
	client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if vnicsIScsi, err := ucsm.SpGetVnicsIScsi(client, spDn); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		for _, vnic := range *vnicsIScsi {
			fmt.Printf("Dn %s\n", vnic.Dn)
			fmt.Printf("\tInitiator Name: %s\n", vnic.InitiatorName)
			fmt.Printf("\tIP Address: %s\n", vnic.VnicVlan.VnicIPv4If.VnicIPv4IscsiAddr.Addr)
			fmt.Printf("\tSubnet: %s\n", vnic.VnicVlan.VnicIPv4If.VnicIPv4IscsiAddr.Subnet)
			fmt.Printf("\tGateway: %s\n", vnic.VnicVlan.VnicIPv4If.VnicIPv4IscsiAddr.DefGw)
			fmt.Println("\tiSCSI Targets:")
			for _, target := range vnic.VnicVlan.VnicIScsiStaticTargets {
				fmt.Printf("\t\tName: %s\n", target.Name)
				fmt.Printf("\t\t\tIP Address: %s\n", target.IpAddress)
				fmt.Printf("\t\t\tPort: %s\n", target.Port)
				fmt.Printf("\t\t\tPriority: %s\n", target.Priority)
				for _, lun := range target.VnicLuns {
					fmt.Println("\t\t\tLuns:")
					fmt.Printf("\t\t\t\tId: %s\n", lun.Id)
					fmt.Printf("\t\t\t\tBootable: %s\n", lun.Bootable)

				}
			}

		}
	}
}
```