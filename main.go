package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

var config *terraform.ResourceConfig

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
	r := p.ResourcesMap[resourceType]
	data := r.Data(&terraform.InstanceState{ID: id})
	return r.Delete(data, p.Meta())
}

func main() {
	p := aws.Provider().(*schema.Provider)
	err := p.Configure(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	resourceConfig := &terraform.ResourceConfig{
		Config: map[string]interface{}{
			"cidr_block":       "192.168.0.0/16",
			"instance_tenancy": "default",
			"tags": map[string]interface{}{
				"Name": "lyra",
			},
		},
	}
	id, err := create(p, "aws_vpc", resourceConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VPC ID:", id)
	state, err := read(p, "aws_vpc", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(state)
	delete(p, "aws_vpc", id)
}
