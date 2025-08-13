package ucsm

import (
	"encoding/xml"
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
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
	if err = c.ConfigConfMo(req, &out); err != nil {
		return nil, fmt.Errorf("failed to set attributes for service profile '%s': %w", spDn, err)
	}
	return &out.LsServer, nil
}
