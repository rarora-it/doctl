package godo

import (
	"context"
	"net/http"
)

const (
	dedicatedInferenceBasePath = "/v2/dedicated-inferences"
)

// DedicatedInferenceService is an interface for managing dedicated inference endpoints
// of the DigitalOcean API.
type DedicatedInferenceService interface {
	Create(context.Context, *DedicatedInferenceCreateRequest) (*DedicatedInference, *Response, error)
}

var _ DedicatedInferenceService = &DedicatedInferenceServiceOp{}

// DedicatedInferenceServiceOp handles communication with the dedicated inference
// related methods of the DigitalOcean API.
type DedicatedInferenceServiceOp struct {
	client *Client
}

// DedicatedInference represents a dedicated inference endpoint.
type DedicatedInference struct {
	ID        string                       `json:"id,omitempty"`
	Name      string                       `json:"name,omitempty"`
	Region    string                       `json:"region,omitempty"`
	VPCUUID   string                       `json:"vpc_uuid,omitempty"`
	Status    string                       `json:"status,omitempty"`
	Endpoints *DedicatedInferenceEndpoints `json:"endpoints,omitempty"`
	CreatedAt *Timestamp                   `json:"created_at,omitempty"`
	UpdatedAt *Timestamp                   `json:"updated_at,omitempty"`
}

// DedicatedInferenceEndpoints represents the endpoints for a dedicated inference.
type DedicatedInferenceEndpoints struct {
	PublicEndpointFQDN  string `json:"public_endpoint_fqdn,omitempty"`
	PrivateEndpointFQDN string `json:"private_endpoint_fqdn,omitempty"`
}

// DedicatedInferenceCreateRequest represents the request to create a dedicated inference endpoint.
type DedicatedInferenceCreateRequest struct {
	Name             string `json:"name"`
	Region           string `json:"region"`
	ModelSlug        string `json:"model_slug"`
	AcceleratorSlug  string `json:"accelerator_slug"`
	NodeCount        int    `json:"node_count"`
	VPCUUID          string `json:"vpc_uuid"`
	HuggingFaceToken string `json:"hugging_face_token,omitempty"`
}

type dedicatedInferenceRoot struct {
	DedicatedInference *DedicatedInference `json:"dedicated_inference"`
}

// Create creates a new dedicated inference endpoint.
func (s *DedicatedInferenceServiceOp) Create(ctx context.Context, create *DedicatedInferenceCreateRequest) (*DedicatedInference, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, dedicatedInferenceBasePath, create)
	if err != nil {
		return nil, nil, err
	}

	root := new(dedicatedInferenceRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	return root.DedicatedInference, resp, nil
}
