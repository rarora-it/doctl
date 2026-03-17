package displayers

import (
	"io"

	"github.com/digitalocean/doctl/do"
)

// DedicatedInference wraps a slice of DedicatedInference for display.
type DedicatedInference struct {
	DedicatedInferences do.DedicatedInferences
}

var _ Displayable = &DedicatedInference{}

func (d *DedicatedInference) JSON(out io.Writer) error {
	return writeJSON(d.DedicatedInferences, out)
}

func (d *DedicatedInference) Cols() []string {
	return []string{
		"ID",
		"Name",
		"Region",
		"VPCUUID",
		"Status",
		"PublicEndpoint",
		"PrivateEndpoint",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (d *DedicatedInference) ColMap() map[string]string {
	return map[string]string{
		"ID":              "ID",
		"Name":            "Name",
		"Region":          "Region",
		"VPCUUID":         "VPC UUID",
		"Status":          "Status",
		"PublicEndpoint":  "Public Endpoint",
		"PrivateEndpoint": "Private Endpoint",
		"CreatedAt":       "Created At",
		"UpdatedAt":       "Updated At",
	}
}

func (d *DedicatedInference) KV() []map[string]any {
	if d == nil || d.DedicatedInferences == nil {
		return []map[string]any{}
	}
	out := make([]map[string]any, 0, len(d.DedicatedInferences))
	for _, di := range d.DedicatedInferences {
		publicEndpoint := ""
		privateEndpoint := ""
		if di.Endpoints != nil {
			publicEndpoint = di.Endpoints.PublicEndpointFQDN
			privateEndpoint = di.Endpoints.PrivateEndpointFQDN
		}
		out = append(out, map[string]any{
			"ID":              di.ID,
			"Name":            di.Name,
			"Region":          di.Region,
			"VPCUUID":         di.VPCUUID,
			"Status":          di.Status,
			"PublicEndpoint":  publicEndpoint,
			"PrivateEndpoint": privateEndpoint,
			"CreatedAt":       di.CreatedAt,
			"UpdatedAt":       di.UpdatedAt,
		})
	}
	return out
}
