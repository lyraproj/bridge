package main

import (
	"fmt"

	"github.com/scottyw/lyra-bridge/pkg/generated"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

var config *terraform.ResourceConfig

func ptr(s string) *string {
	return &s
}

func init() {
	// Fields derived from Schema
	config = &terraform.ResourceConfig{
		Config: map[string]interface{}{
			"region": "eu-west-1",
		},
	}
}

func create(p *schema.Provider, resourceType string, resourceConfig *terraform.ResourceConfig) (string, error) {
	// To get Terraform to create a new resource, the ID must be blank and existing state must be empty (since the
	// resource does not exist yet), and the diff object should have no old state and all of the new state.
	info := &terraform.InstanceInfo{Type: resourceType}
	state := &terraform.InstanceState{}
	diff, err := p.Diff(info, state, resourceConfig)
	if err != nil {
		return "", err
	}
	newstate, err := p.Apply(info, state, diff)
	if newstate == nil {
		// contract.Assertf(err != nil, "expected non-nil error with nil state during Create")
		return "", err
	}
	// contract.Assertf(newstate.ID != "", "Expected non-empty ID for new state during Create")
	return newstate.ID, nil
}

func read(p *schema.Provider, resourceType string, id string) (*terraform.InstanceState, error) {
	info := &terraform.InstanceInfo{Type: resourceType}
	state := &terraform.InstanceState{ID: id}
	newstate, err := p.Refresh(info, state)
	if err != nil {
		return nil, err
	}
	return newstate, nil
}

func delete(p *schema.Provider, resourceType string, id string) error {
	info := &terraform.InstanceInfo{Type: resourceType}
	state := &terraform.InstanceState{ID: id}
	diff := &terraform.InstanceDiff{Destroy: true}
	_, err := p.Apply(info, state, diff)
	return err
}

func main() {

	// Configure the provider
	p := aws.Provider().(*schema.Provider)
	err := p.Configure(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create VPC
	vpc := &generated.Aws_vpc{
		Cidr_block:       "192.168.0.0/16",
		Instance_tenancy: ptr("default"),
		Tags: &map[string]interface{}{
			"Name": "lyra-test",
		},
	}
	vid, err := create(p, "aws_vpc", generated.Aws_vpcMapper(vpc))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VPC ID:", vid)

	// Read VPC
	state, err := read(p, "aws_vpc", vid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(state)

	// Create subnet
	subnet := &generated.Aws_subnet{
		Vpc_id:     vid,
		Cidr_block: "192.168.1.0/24",
		Tags: &map[string]interface{}{
			"Name": "lyra-test",
		},
	}
	sid, err := create(p, "aws_subnet", generated.Aws_subnetMapper(subnet))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("SUBNET ID:", sid)

	// Read subnet
	state, err = read(p, "aws_subnet", sid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(state)

	// Delete subnet
	delete(p, "aws_subnet", sid)

	// Delete VPC
	delete(p, "aws_vpc", vid)
}
