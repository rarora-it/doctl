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

// DedicatedInference represents a dedicated inference endpoint response.
type DedicatedInference struct {
	ID                string                               `json:"id,omitempty"`
	OwnerID           int64                                `json:"owner_id,omitempty"`
	State             string                               `json:"state,omitempty"`
	Region            string                               `json:"region,omitempty"`
	VPCUUID           string                               `json:"vpc_uuid,omitempty"`
	Spec              *DedicatedInferenceSpec              `json:"spec,omitempty"`
	PendingDeployment *DedicatedInferencePendingDeployment `json:"pending-deployment,omitempty"`
	CreatedAt         *Timestamp                           `json:"created_at,omitempty"`
	UpdatedAt         *Timestamp                           `json:"updated_at,omitempty"`
}

// DedicatedInferenceSpec represents the deployment specification for a dedicated inference.
type DedicatedInferenceSpec struct {
	Version              int                                  `json:"version" yaml:"version"`
	Name                 string                               `json:"name" yaml:"name"`
	Region               string                               `json:"region" yaml:"region"`
	VPC                  *DedicatedInferenceVPC               `json:"vpc,omitempty" yaml:"vpc,omitempty"`
	EnablePublicEndpoint bool                                 `json:"enable_public_endpoint" yaml:"enable_public_endpoint"`
	ModelDeployments     []*DedicatedInferenceModelDeployment `json:"model_deployments" yaml:"model_deployments"`
}

// DedicatedInferenceVPC represents VPC configuration within a spec.
type DedicatedInferenceVPC struct {
	UUID string `json:"uuid" yaml:"uuid"`
}

// DedicatedInferenceModelDeployment represents a model deployment within a spec.
type DedicatedInferenceModelDeployment struct {
	ModelSlug      string                           `json:"model_slug" yaml:"model_slug"`
	ModelProvider  string                           `json:"model_provider" yaml:"model_provider"`
	WorkloadConfig map[string]interface{}           `json:"workload_config,omitempty" yaml:"workload_config,omitempty"`
	Accelerators   []*DedicatedInferenceAccelerator `json:"accelerators" yaml:"accelerators"`
}

// DedicatedInferenceAccelerator represents an accelerator configuration within a model deployment.
type DedicatedInferenceAccelerator struct {
	Scale           int    `json:"scale" yaml:"scale"`
	Type            string `json:"type" yaml:"type"`
	AcceleratorSlug string `json:"accelerator_slug" yaml:"accelerator_slug"`
}

// DedicatedInferencePendingDeployment represents a pending deployment within the response.
type DedicatedInferencePendingDeployment struct {
	ID        string                  `json:"id,omitempty"`
	Spec      *DedicatedInferenceSpec `json:"spec,omitempty"`
	State     string                  `json:"state,omitempty"`
	CreatedAt *Timestamp              `json:"created_at,omitempty"`
	UpdatedAt *Timestamp              `json:"updated_at,omitempty"`
}

// DedicatedInferenceAccessTokens represents access tokens for model providers.
type DedicatedInferenceAccessTokens struct {
	HuggingFaceToken string `json:"hugging_face_token,omitempty"`
}

// DedicatedInferenceCreateRequest represents the request to create a dedicated inference endpoint.
type DedicatedInferenceCreateRequest struct {
	Spec         *DedicatedInferenceSpec         `json:"spec"`
	AccessTokens *DedicatedInferenceAccessTokens `json:"access_tokens,omitempty"`
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
