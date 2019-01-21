package awstypes

import (
	"github.com/hashicorp/terraform/terraform"
)

// Goodnight sweet linter, I'm sorry

type aws_vpc struct {
	Ipv6_cidr_block                  string
	Arn                              string
	Instance_tenancy                 *string
	Enable_classiclink               *bool
	Assign_generated_ipv6_cidr_block *bool
	Main_route_table_id              string
	Default_network_acl_id           string
	Dhcp_options_id                  string
	Default_security_group_id        string
	Ipv6_association_id              string
	Tags                             *map[string]interface{}
	Cidr_block                       string
	Enable_dns_hostnames             *bool
	Enable_dns_support               *bool
	Enable_classiclink_dns_support   *bool
	Default_route_table_id           string
	Owner_id                         string
}

func aws_vpcMapper(r *aws_vpc) *terraform.ResourceConfig {
	return &terraform.ResourceConfig{
		Config: map[string]interface{}{
			"Default_network_acl_id":           r.Default_network_acl_id,
			"Dhcp_options_id":                  r.Dhcp_options_id,
			"Ipv6_cidr_block":                  r.Ipv6_cidr_block,
			"Arn":                              r.Arn,
			"Instance_tenancy":                 r.Instance_tenancy,
			"Enable_classiclink":               r.Enable_classiclink,
			"Assign_generated_ipv6_cidr_block": r.Assign_generated_ipv6_cidr_block,
			"Main_route_table_id":              r.Main_route_table_id,
			"Default_security_group_id":        r.Default_security_group_id,
			"Ipv6_association_id":              r.Ipv6_association_id,
			"Tags":                             r.Tags,
			"Default_route_table_id":           r.Default_route_table_id,
			"Owner_id":                         r.Owner_id,
			"Cidr_block":                       r.Cidr_block,
			"Enable_dns_hostnames":             r.Enable_dns_hostnames,
			"Enable_dns_support":               r.Enable_dns_support,
			"Enable_classiclink_dns_support":   r.Enable_classiclink_dns_support,
		},
	}
}

func aws_vpcUnmapper(r *terraform.ResourceConfig) *aws_vpc {
	return &aws_vpc{
		Arn:                              r.Config["Arn"].(string),
		Instance_tenancy:                 r.Config["Instance_tenancy"].(*string),
		Enable_classiclink:               r.Config["Enable_classiclink"].(*bool),
		Assign_generated_ipv6_cidr_block: r.Config["Assign_generated_ipv6_cidr_block"].(*bool),
		Main_route_table_id:              r.Config["Main_route_table_id"].(string),
		Default_network_acl_id:           r.Config["Default_network_acl_id"].(string),
		Dhcp_options_id:                  r.Config["Dhcp_options_id"].(string),
		Ipv6_cidr_block:                  r.Config["Ipv6_cidr_block"].(string),
		Default_security_group_id:        r.Config["Default_security_group_id"].(string),
		Ipv6_association_id:              r.Config["Ipv6_association_id"].(string),
		Tags:                             r.Config["Tags"].(*map[string]interface{}),
		Cidr_block:                       r.Config["Cidr_block"].(string),
		Enable_dns_hostnames:             r.Config["Enable_dns_hostnames"].(*bool),
		Enable_dns_support:               r.Config["Enable_dns_support"].(*bool),
		Enable_classiclink_dns_support:   r.Config["Enable_classiclink_dns_support"].(*bool),
		Default_route_table_id:           r.Config["Default_route_table_id"].(string),
		Owner_id:                         r.Config["Owner_id"].(string),
	}
}
