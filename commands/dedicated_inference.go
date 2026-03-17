/*
Copyright 2018 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
	http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"github.com/digitalocean/doctl"
	"github.com/digitalocean/doctl/commands/displayers"
	"github.com/digitalocean/doctl/do"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// DedicatedInferenceCmd creates the dedicated-inference command and its subcommands.
func DedicatedInferenceCmd() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:     "dedicated-inference",
			Aliases: []string{"di", "dedicated-inferences"},
			Short:   "Display commands for managing dedicated inference endpoints",
			Long:    "The subcommands of `doctl dedicated-inference` manage your dedicated inference endpoints.",
			GroupID: manageResourcesGroup,
		},
	}

	cmdCreate := CmdBuilder(
		cmd,
		RunDedicatedInferenceCreate,
		"create <name>",
		"Create a dedicated inference endpoint",
		"Creates a dedicated inference endpoint on your account. The command requires values for the "+
			"`--name`, `--model-slug`, `--accelerator-slug`, `--region`, `--node-count`, and `--vpc-uuid` flags.",
		Writer,
		aliasOpt("c"),
		displayerType(&displayers.DedicatedInference{}),
	)
	AddStringFlag(cmdCreate, doctl.ArgDedicatedInferenceName, "", "", "Dedicated inference endpoint name", requiredOpt())
	AddStringFlag(cmdCreate, doctl.ArgDedicatedInferenceRegion, "", "", "Region to deploy the inference endpoint in", requiredOpt())
	AddStringFlag(cmdCreate, doctl.ArgDedicatedInferenceModelSlug, "", "", `LLM model slug with provider information (e.g. "hf://mistral/mistral-7b-instruct-v3")`, requiredOpt())
	AddStringFlag(cmdCreate, doctl.ArgDedicatedInferenceAcceleratorSlug, "", "", `GPU accelerator slug (e.g. "gpu-mi300x1-192gb")`, requiredOpt())
	AddIntFlag(cmdCreate, doctl.ArgDedicatedInferenceNodeCount, "", 1, "Number of GPU nodes to allocate", requiredOpt())
	AddStringFlag(cmdCreate, doctl.ArgDedicatedInferenceVPCUUID, "", "", "UUID of the VPC to place the endpoint in", requiredOpt())
	AddStringFlag(cmdCreate, doctl.ArgDedicatedInferenceHuggingFaceToken, "", "", "Hugging Face token for accessing gated models (optional)")
	cmdCreate.Example = `The following example creates a dedicated inference endpoint: doctl dedicated-inference create --name "my-inference" --model-slug "hf://mistral/mistral-7b-instruct-v3" --accelerator-slug "gpu-mi300x1-192gb" --region "nyc2" --node-count 2 --vpc-uuid "997615ce-132d-4bae-9270-9ee21b395e5d"`

	return cmd
}

// RunDedicatedInferenceCreate creates a new dedicated inference endpoint.
func RunDedicatedInferenceCreate(c *CmdConfig) error {
	name, err := c.Doit.GetString(c.NS, doctl.ArgDedicatedInferenceName)
	if err != nil {
		return err
	}
	region, err := c.Doit.GetString(c.NS, doctl.ArgDedicatedInferenceRegion)
	if err != nil {
		return err
	}
	modelSlug, err := c.Doit.GetString(c.NS, doctl.ArgDedicatedInferenceModelSlug)
	if err != nil {
		return err
	}
	acceleratorSlug, err := c.Doit.GetString(c.NS, doctl.ArgDedicatedInferenceAcceleratorSlug)
	if err != nil {
		return err
	}
	nodeCount, err := c.Doit.GetInt(c.NS, doctl.ArgDedicatedInferenceNodeCount)
	if err != nil {
		return err
	}
	vpcUUID, err := c.Doit.GetString(c.NS, doctl.ArgDedicatedInferenceVPCUUID)
	if err != nil {
		return err
	}
	huggingFaceToken, _ := c.Doit.GetString(c.NS, doctl.ArgDedicatedInferenceHuggingFaceToken)

	req := &godo.DedicatedInferenceCreateRequest{
		Name:            name,
		Region:          region,
		ModelSlug:       modelSlug,
		AcceleratorSlug: acceleratorSlug,
		NodeCount:       nodeCount,
		VPCUUID:         vpcUUID,
	}
	if huggingFaceToken != "" {
		req.HuggingFaceToken = huggingFaceToken
	}

	endpoint, err := c.DedicatedInferences().Create(req)
	if err != nil {
		return err
	}
	return c.Display(&displayers.DedicatedInference{DedicatedInferences: do.DedicatedInferences{*endpoint}})
}
