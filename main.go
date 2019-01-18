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

func create(p *schema.Provider, resourceType string, resourceData map[string]interface{}) (string, error) {
	r := p.ResourcesMap[resourceType]
	data := r.Data(&terraform.InstanceState{})
	for k, v := range resourceData {
		data.Set(k, v)
	}
	data.MarkNewResource()
	err := r.Create(data, p.Meta())
	if err != nil {
		return "", err
	}
	return data.Id(), nil
}

func read(p *schema.Provider, resourceType string, id string) error {
	r := p.ResourcesMap[resourceType]
	data := r.Data(&terraform.InstanceState{
		ID: id,
	})
	err := r.Read(data, p.Meta())
	if err != nil {
		return err
	}
	fmt.Println("READ:", data)
	return nil
}

func delete(p *schema.Provider, resourceType string, id string) error {
	r := p.ResourcesMap[resourceType]
	data := r.Data(&terraform.InstanceState{
		ID: id,
	})
	return r.Delete(data, p.Meta())
}

func main() {
	p := aws.Provider().(*schema.Provider)
	err := p.Configure(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := create(p, "aws_vpc", map[string]interface{}{
		"cidr_block":       "192.168.0.0/16",
		"instance_tenancy": "default",
		"tags": map[string]interface{}{
			"Name": "lyra",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VPC ID:", id)
	read(p, "aws_vpc", id)
	delete(p, "aws_vpc", id)
}
