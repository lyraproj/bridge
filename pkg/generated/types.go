package generated

import (
	"github.com/hashicorp/terraform/terraform"
)

// Goodnight sweet linter, I'm sorry

type Aws_subnet struct {
	Availability_zone               *string
	Availability_zone_id            *string
	Map_public_ip_on_launch         *bool
	Vpc_id                          string
	Cidr_block                      string
	Ipv6_cidr_block_association_id  string
	Arn                             string
	Tags                            *map[string]interface{}
	Owner_id                        string
	Ipv6_cidr_block                 *string
	Assign_ipv6_address_on_creation *bool
}

func Aws_subnetMapper(r *Aws_subnet) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	if r.Availability_zone_id != nil {
		config["availability_zone_id"] = *r.Availability_zone_id
	}
	if r.Map_public_ip_on_launch != nil {
		config["map_public_ip_on_launch"] = *r.Map_public_ip_on_launch
	}
	config["vpc_id"] = r.Vpc_id
	config["cidr_block"] = r.Cidr_block
	config["ipv6_cidr_block_association_id"] = r.Ipv6_cidr_block_association_id
	config["arn"] = r.Arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	if r.Ipv6_cidr_block != nil {
		config["ipv6_cidr_block"] = *r.Ipv6_cidr_block
	}
	if r.Assign_ipv6_address_on_creation != nil {
		config["assign_ipv6_address_on_creation"] = *r.Assign_ipv6_address_on_creation
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

func Aws_subnetUnmapper(state map[string]interface{}) *Aws_subnet {
	r := &Aws_subnet{}

	if x, ok := state["arn"]; ok {
		r.Arn = x.(string)
	}

	if x, ok := state["tags"]; ok {
		x := x.(map[string]interface{})
		r.Tags = &x
	}

	if x, ok := state["owner_id"]; ok {
		r.Owner_id = x.(string)
	}

	if x, ok := state["ipv6_cidr_block"]; ok {
		x := x.(string)
		r.Ipv6_cidr_block = &x
	}

	if x, ok := state["assign_ipv6_address_on_creation"]; ok {
		x := x.(bool)
		r.Assign_ipv6_address_on_creation = &x
	}

	if x, ok := state["ipv6_cidr_block_association_id"]; ok {
		r.Ipv6_cidr_block_association_id = x.(string)
	}

	if x, ok := state["availability_zone_id"]; ok {
		x := x.(string)
		r.Availability_zone_id = &x
	}

	if x, ok := state["map_public_ip_on_launch"]; ok {
		x := x.(bool)
		r.Map_public_ip_on_launch = &x
	}

	if x, ok := state["vpc_id"]; ok {
		r.Vpc_id = x.(string)
	}

	if x, ok := state["cidr_block"]; ok {
		r.Cidr_block = x.(string)
	}

	if x, ok := state["availability_zone"]; ok {
		x := x.(string)
		r.Availability_zone = &x
	}
	return r
}

type Aws_vpc struct {
	Default_security_group_id        string
	Default_route_table_id           string
	Cidr_block                       string
	Instance_tenancy                 *string
	Ipv6_association_id              string
	Ipv6_cidr_block                  string
	Arn                              string
	Tags                             *map[string]interface{}
	Enable_dns_hostnames             *bool
	Enable_dns_support               *bool
	Main_route_table_id              string
	Dhcp_options_id                  string
	Owner_id                         string
	Enable_classiclink               *bool
	Enable_classiclink_dns_support   *bool
	Assign_generated_ipv6_cidr_block *bool
	Default_network_acl_id           string
}

func Aws_vpcMapper(r *Aws_vpc) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["default_security_group_id"] = r.Default_security_group_id
	config["default_route_table_id"] = r.Default_route_table_id
	config["cidr_block"] = r.Cidr_block
	if r.Instance_tenancy != nil {
		config["instance_tenancy"] = *r.Instance_tenancy
	}
	config["main_route_table_id"] = r.Main_route_table_id
	config["dhcp_options_id"] = r.Dhcp_options_id
	config["ipv6_association_id"] = r.Ipv6_association_id
	config["ipv6_cidr_block"] = r.Ipv6_cidr_block
	config["arn"] = r.Arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Enable_dns_hostnames != nil {
		config["enable_dns_hostnames"] = *r.Enable_dns_hostnames
	}
	if r.Enable_dns_support != nil {
		config["enable_dns_support"] = *r.Enable_dns_support
	}
	config["owner_id"] = r.Owner_id
	if r.Assign_generated_ipv6_cidr_block != nil {
		config["assign_generated_ipv6_cidr_block"] = *r.Assign_generated_ipv6_cidr_block
	}
	config["default_network_acl_id"] = r.Default_network_acl_id
	if r.Enable_classiclink != nil {
		config["enable_classiclink"] = *r.Enable_classiclink
	}
	if r.Enable_classiclink_dns_support != nil {
		config["enable_classiclink_dns_support"] = *r.Enable_classiclink_dns_support
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

func Aws_vpcUnmapper(state map[string]interface{}) *Aws_vpc {
	r := &Aws_vpc{}

	if x, ok := state["default_security_group_id"]; ok {
		r.Default_security_group_id = x.(string)
	}

	if x, ok := state["default_route_table_id"]; ok {
		r.Default_route_table_id = x.(string)
	}

	if x, ok := state["cidr_block"]; ok {
		r.Cidr_block = x.(string)
	}

	if x, ok := state["instance_tenancy"]; ok {
		x := x.(string)
		r.Instance_tenancy = &x
	}

	if x, ok := state["enable_dns_hostnames"]; ok {
		x := x.(bool)
		r.Enable_dns_hostnames = &x
	}

	if x, ok := state["enable_dns_support"]; ok {
		x := x.(bool)
		r.Enable_dns_support = &x
	}

	if x, ok := state["main_route_table_id"]; ok {
		r.Main_route_table_id = x.(string)
	}

	if x, ok := state["dhcp_options_id"]; ok {
		r.Dhcp_options_id = x.(string)
	}

	if x, ok := state["ipv6_association_id"]; ok {
		r.Ipv6_association_id = x.(string)
	}

	if x, ok := state["ipv6_cidr_block"]; ok {
		r.Ipv6_cidr_block = x.(string)
	}

	if x, ok := state["arn"]; ok {
		r.Arn = x.(string)
	}

	if x, ok := state["tags"]; ok {
		x := x.(map[string]interface{})
		r.Tags = &x
	}

	if x, ok := state["owner_id"]; ok {
		r.Owner_id = x.(string)
	}

	if x, ok := state["enable_classiclink"]; ok {
		x := x.(bool)
		r.Enable_classiclink = &x
	}

	if x, ok := state["enable_classiclink_dns_support"]; ok {
		x := x.(bool)
		r.Enable_classiclink_dns_support = &x
	}

	if x, ok := state["assign_generated_ipv6_cidr_block"]; ok {
		x := x.(bool)
		r.Assign_generated_ipv6_cidr_block = &x
	}

	if x, ok := state["default_network_acl_id"]; ok {
		r.Default_network_acl_id = x.(string)
	}
	return r
}
