package ucsm

import (
	"fmt"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

func GetAllComputeBlades(c *api.Client) ([]mo.ComputeBlade, error) {
	var out mo.Blades

	req := api.ConfigResolveClassRequest{
		Cookie:         c.Cookie,
		ClassID:        "computeBlade",
		InHierarchical: "false",
		// Sem InFilter para buscar todas
	}

	if err := c.ConfigResolveClass(req, &out); err != nil {
		return nil, fmt.Errorf("failed to resolve compute blades: %w", err)
	}

	return out.ComputeBlades, nil
}
