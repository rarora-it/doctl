package commands

import (
	"testing"

	"github.com/digitalocean/godo"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/doctl/do"
	"github.com/stretchr/testify/assert"
)

// Test data
var (
	testDedicatedInference = do.DedicatedInference{
		DedicatedInference: &godo.DedicatedInference{
			ID:      "00000000-0000-4000-8000-000000000000",
			Name:    "test-dedicated-inference",
			Region:  "nyc2",
			VPCUUID: "00000000-0000-4000-8000-000000000001",
			Status:  "CREATING",
			Endpoints: &godo.DedicatedInferenceEndpoints{
				PublicEndpointFQDN:  "https://test-public.do-infra.ai",
				PrivateEndpointFQDN: "https://test-private.do-infra.ai",
			},
		},
	}
)

func TestDedicatedInferenceCommand(t *testing.T) {
	cmd := DedicatedInferenceCmd()
	assert.NotNil(t, cmd)
	assert.Equal(t, "dedicated-inference", cmd.Name())

	// Verify create is a subcommand
	found := false
	for _, c := range cmd.Commands() {
		if c.Name() == "create" {
			found = true
			break
		}
	}
	assert.True(t, found, "Expected create subcommand")
}

func TestRunDedicatedInferenceCreate(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceName, testDedicatedInference.Name)
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceRegion, testDedicatedInference.Region)
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceModelSlug, "hf://mistral/mistral-7b-instruct-v3")
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceAcceleratorSlug, "gpu-mi300x1-192gb")
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceNodeCount, 2)
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceVPCUUID, testDedicatedInference.VPCUUID)

		expectedReq := &godo.DedicatedInferenceCreateRequest{
			Name:            testDedicatedInference.Name,
			Region:          testDedicatedInference.Region,
			ModelSlug:       "hf://mistral/mistral-7b-instruct-v3",
			AcceleratorSlug: "gpu-mi300x1-192gb",
			NodeCount:       2,
			VPCUUID:         testDedicatedInference.VPCUUID,
		}

		tm.dedicatedInferences.EXPECT().Create(expectedReq).Return(&testDedicatedInference, nil)

		err := RunDedicatedInferenceCreate(config)
		assert.NoError(t, err)
	})
}

func TestRunDedicatedInferenceCreate_WithHuggingFaceToken(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceName, testDedicatedInference.Name)
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceRegion, testDedicatedInference.Region)
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceModelSlug, "hf://mistral/mistral-7b-instruct-v3")
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceAcceleratorSlug, "gpu-mi300x1-192gb")
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceNodeCount, 2)
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceVPCUUID, testDedicatedInference.VPCUUID)
		config.Doit.Set(config.NS, doctl.ArgDedicatedInferenceHuggingFaceToken, "hf_test_token")

		expectedReq := &godo.DedicatedInferenceCreateRequest{
			Name:             testDedicatedInference.Name,
			Region:           testDedicatedInference.Region,
			ModelSlug:        "hf://mistral/mistral-7b-instruct-v3",
			AcceleratorSlug:  "gpu-mi300x1-192gb",
			NodeCount:        2,
			VPCUUID:          testDedicatedInference.VPCUUID,
			HuggingFaceToken: "hf_test_token",
		}

		tm.dedicatedInferences.EXPECT().Create(expectedReq).Return(&testDedicatedInference, nil)

		err := RunDedicatedInferenceCreate(config)
		assert.NoError(t, err)
	})
}
