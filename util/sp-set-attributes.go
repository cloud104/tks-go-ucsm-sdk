package util

import (
	"encoding/xml"

	"github.com/gfalves87/tks-go-ucsm-sdk/api"
	"github.com/gfalves87/tks-go-ucsm-sdk/mo"
)

type LsServerAttributes struct {
	Descr  string `xml:"descr,attr"`
	UsrLbl string `xml:"usrLbl,attr"`
}

type LsServerAttrMo struct {
	XMLName  xml.Name
	LsServer LsServerAttributes `xml:"lsServer"`
}

func SpSetAttributes(c *api.Client, spDn string, spDescription string, spUserLabel string) (lsServer *mo.LsServer, err error) {
	var out mo.LsServerMo
	lsServerMo := LsServerAttrMo{
		LsServer: LsServerAttributes{
			Descr:  spDescription,
			UsrLbl: spUserLabel,
		},
	}
	req := api.ConfigConfMoRequest{
		Cookie:         c.Cookie,
		Dn:             spDn,
		InHierarchical: "false",
		InConfig:       lsServerMo,
	}
	if err = c.ConfigConfMo(req, &out); err == nil {
		lsServer = &out.LsServer
	}
	return
}
