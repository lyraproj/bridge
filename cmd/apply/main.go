package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/scottyw/lyra-bridge/pkg/generated"
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

func main() {

	// Configure the provider
	p := aws.Provider().(*schema.Provider)
	err := p.Configure(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create VPC
	vpcHandler := generated.Aws_vpcHandler{Provider: p}
	vpc := &generated.Aws_vpc{
		Cidr_block:       "192.168.0.0/16",
		Instance_tenancy: ptr("default"),
		Tags: &map[string]interface{}{
			"Name": "lyra-test",
		},
	}
	vpc, vid, err := vpcHandler.Create(vpc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CREATED VPC:", spew.Sdump(vpc))
	fmt.Println("CREATED VPC ID:", vid)

	// Read VPC
	vpc, err = vpcHandler.Read(vid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("READ VPC:", spew.Sdump(vpc))

	// Create SUBNET
	subnetHandler := generated.Aws_subnetHandler{Provider: p}
	subnet := &generated.Aws_subnet{
		Vpc_id:     vid,
		Cidr_block: "192.168.1.0/24",
		Tags: &map[string]interface{}{
			"Name": "lyra-test",
		},
	}
	subnet, sid, err := subnetHandler.Create(subnet)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CREATED SUBNET:", spew.Sdump(subnet))
	fmt.Println("CREATED SUBNET ID:", sid)

	// Read SUBNET
	subnet, err = subnetHandler.Read(sid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("READ SUBNET:", spew.Sdump(subnet))

	// Delete SUBNET
	subnetHandler.Delete(sid)

	// Delete VPC
	vpcHandler.Delete(vid)

}
