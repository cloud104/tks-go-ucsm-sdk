package util

import (
	"encoding/xml"

	"github.com/igor-feoktistov/go-ucsm-sdk/api"
	"github.com/igor-feoktistov/go-ucsm-sdk/mo"
)

type LsServerAttributes struct {
	AssetTag string `xml:"assetTag,attr,omitempty"`
	Descr string    `xml:"descr,attr,omitempty"`
	UsrLbl string   `xml:"usrLbl,attr,omitempty"`
}

type LsServerAttrMo struct {
        XMLName xml.Name
        LsServer LsServerAttributes `xml:"lsServer"`
}

func SpSetAttributes(c *api.Client, spDn string, spAssetTag string, spDescription string, spUserLabel string) (lsServer *mo.LsServer, err error) {
	var out mo.LsServerMo
	lsServerMo := LsServerAttrMo {
			LsServer: LsServerAttributes {
				    AssetTag: spAssetTag,
				    Descr: spDescription,
				    UsrLbl: spUserLabel,
			},
	}
	req := api.ConfigConfMoRequest {
		    Cookie: c.Cookie,
		    Dn: spDn,
		    InHierarchical: "false",
		    InConfig: lsServerMo,
	}
	if err = c.ConfigConfMo(req, &out); err == nil {
		lsServer = &out.LsServer
	}
	return
}
