package ucsm

import (
	"encoding/xml"
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

// Re-defining models.LsServer to allow empty SrcTemplName value
type LsServerUnbind struct {
	SrcTemplName string `xml:"srcTemplName,attr"`
}

type LsServerUnbindMo struct {
	XMLName  xml.Name
	LsServer LsServerUnbind `xml:"lsServer"`
}

func SpUnbindFromSpt(c *api.Client, spDn string) (lsServer *mo.LsServer, err error) {
	var out mo.LsServerMo
	lsServerMo := LsServerUnbindMo{
		LsServer: LsServerUnbind{
			SrcTemplName: "",
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
	return lsServer, fmt.Errorf("failed to unbind server profile from template %s: %w", spDn, err)
}
