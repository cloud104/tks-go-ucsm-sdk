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
	spDn := "org-root"
	pnDn := "sys/chassis-1/blade-2"
	client, err := ucsm.AaaLogin(endPoint, username, password)
	//client.SetDebug(true)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if lsBinding, err := ucsm.SpAssignBlade(client, spDn, pnDn); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Dn: %s\n", lsBinding.Dn)
		fmt.Printf("\tAssigned Blade: %s\n", lsBinding.AssignedToDn)
		if lsBinding.FaultInst != nil {
			fmt.Printf("\tFault: %s\n", lsBinding.FaultInst.Descr)
		}
	}
}
```