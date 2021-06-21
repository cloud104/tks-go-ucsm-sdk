package util

import (
	"encoding/xml"

	"github.com/igor-feoktistov/go-ucsm-sdk/api"
	"github.com/igor-feoktistov/go-ucsm-sdk/mo"
)

type LsServerDescr struct {
	Descr string `xml:"descr,attr,omitempty"`
}

type LsServerDescrMo struct {
        XMLName xml.Name
        LsServer LsServerDescr `xml:"lsServer"`
}

func SpSetDescr(c *api.Client, spDn string, spDescr string) (lsServer *mo.LsServer, err error) {
	var out mo.LsServerMo
	lsServerMo := LsServerDescrMo {
			LsServer: LsServerDescr {
				    Descr: spDescr,
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
