package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

var config *terraform.ResourceConfig

func init() {
	creds, err := (&credentials.SharedCredentialsProvider{}).Retrieve()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Fields derived rom Schema
	config = &terraform.ResourceConfig{
		Config: map[string]interface{}{
			"access_key": creds.AccessKeyID,
			"secret_key": creds.SecretAccessKey,
			"token":      creds.SessionToken,
			"region":     "eu-west-1",
		},
	}
}

func main() {
	p := aws.Provider().(*schema.Provider)
	err := p.Configure(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := p.ResourcesMap["aws_vpc"]
	data := r.Data(&terraform.InstanceState{})
	data.Set("cidr_block", "192.168.0.0/16")
	data.Set("instance_tenancy", "default")
	// data.Set("tags", map[string]string{"Name": "lyra"})
	err = r.Create(data, p.Meta())
	fmt.Println(data.Id())
	if err != nil {
		fmt.Println(err)
		return
	}
}
