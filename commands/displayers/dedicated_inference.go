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
		"State",
		"Region",
		"VPCUUID",
		"PendingDeploymentState",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (d *DedicatedInference) ColMap() map[string]string {
	return map[string]string{
		"ID":                     "ID",
		"Name":                   "Name",
		"State":                  "State",
		"Region":                 "Region",
		"VPCUUID":                "VPC UUID",
		"PendingDeploymentState": "Pending Deployment State",
		"CreatedAt":              "Created At",
		"UpdatedAt":              "Updated At",
	}
}

func (d *DedicatedInference) KV() []map[string]any {
	if d == nil || d.DedicatedInferences == nil {
		return []map[string]any{}
	}
	out := make([]map[string]any, 0, len(d.DedicatedInferences))
	for _, di := range d.DedicatedInferences {
		name := ""
		if di.Spec != nil {
			name = di.Spec.Name
		}
		pendingState := ""
		if di.PendingDeployment != nil {
			pendingState = di.PendingDeployment.State
		}
		out = append(out, map[string]any{
			"ID":                     di.ID,
			"Name":                   name,
			"State":                  di.State,
			"Region":                 di.Region,
			"VPCUUID":                di.VPCUUID,
			"PendingDeploymentState": pendingState,
			"CreatedAt":              di.CreatedAt,
			"UpdatedAt":              di.UpdatedAt,
		})
	}
	return out
}
