package main

import (
	"fmt"

	"go-ucsm-sdk/util"
	"go-ucsm-sdk/mo"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	spDn := "org-root/org-Linux/ls-phygymdev01-lab"
	iscsiVnicName := "iscsi0"
	iscsiInitiatorName := "iqn.2005-02.com.open-iscsi:phygymdev01-lab.1"
	iscsiVnicAddr := mo.VnicIPv4IscsiAddr {
				Addr: "192.168.1.25",
				Subnet: "255.255.255.0", 
				DefGw: "192.168.1.1",
				PrimDns: "192.168.4.10",
				SecDns: "0.0.0.0",
	}
	iscsiTargets := []mo.VnicIScsiStaticTargetIf {
				{
					IpAddress: "192.168.1.10",
					Name: "iqn.1992-08.com.netapp:sn.f104401e43da11e760d600a098ade5c8:vs.23",
					Port: "3260",
					Priority: "2",
					VnicLuns: []mo.VnicLun {{Bootable: "yes",Id: "0",},},
				},
				{
					IpAddress: "192.168.2.10",
					Name: "iqn.1992-08.com.netapp:sn.f104401e43da11e760d600a098ade5c8:vs.23",
					Port: "3260",
					Priority: "1",
					VnicLuns: []mo.VnicLun {{Bootable: "yes",Id: "0",},},
				},
	}
	client, err := util.AaaLogin(endPoint, username, password)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if vnicIscsi, err := util.SpSetIscsiBoot(client, spDn, iscsiVnicName, iscsiInitiatorName, iscsiVnicAddr, iscsiTargets); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Dn: %s\n", vnicIscsi.Dn)
		fmt.Printf("\t\tConfig State: %s\n", vnicIscsi.ConfigState)
	}
}
