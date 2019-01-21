package generated

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/scottyw/lyra-bridge/pkg/bridge"
)

// Goodnight sweet linter, I'm sorry

type Aws_subnet struct {
	Vpc_id                          string
	Availability_zone_id            *string
	Assign_ipv6_address_on_creation *bool
	Ipv6_cidr_block_association_id  string
	Tags                            *map[string]interface{}
	Cidr_block                      string
	Ipv6_cidr_block                 *string
	Availability_zone               *string
	Map_public_ip_on_launch         *bool
	Arn                             string
	Owner_id                        string
}

func Aws_subnetMapper(r *Aws_subnet) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["cidr_block"] = r.Cidr_block
	if r.Ipv6_cidr_block != nil {
		config["ipv6_cidr_block"] = *r.Ipv6_cidr_block
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	if r.Map_public_ip_on_launch != nil {
		config["map_public_ip_on_launch"] = *r.Map_public_ip_on_launch
	}
	config["arn"] = r.Arn
	config["owner_id"] = r.Owner_id
	config["vpc_id"] = r.Vpc_id
	if r.Availability_zone_id != nil {
		config["availability_zone_id"] = *r.Availability_zone_id
	}
	if r.Assign_ipv6_address_on_creation != nil {
		config["assign_ipv6_address_on_creation"] = *r.Assign_ipv6_address_on_creation
	}
	config["ipv6_cidr_block_association_id"] = r.Ipv6_cidr_block_association_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

func Aws_subnetUnmapper(state map[string]interface{}) *Aws_subnet {
	r := &Aws_subnet{}

	if x, ok := state["tags"]; ok {
		x := x.(map[string]interface{})
		r.Tags = &x
	}

	if x, ok := state["vpc_id"]; ok {
		r.Vpc_id = x.(string)
	}

	if x, ok := state["availability_zone_id"]; ok {
		x := x.(string)
		r.Availability_zone_id = &x
	}

	if x, ok := state["assign_ipv6_address_on_creation"]; ok {
		x := x.(bool)
		r.Assign_ipv6_address_on_creation = &x
	}

	if x, ok := state["ipv6_cidr_block_association_id"]; ok {
		r.Ipv6_cidr_block_association_id = x.(string)
	}

	if x, ok := state["arn"]; ok {
		r.Arn = x.(string)
	}

	if x, ok := state["owner_id"]; ok {
		r.Owner_id = x.(string)
	}

	if x, ok := state["cidr_block"]; ok {
		r.Cidr_block = x.(string)
	}

	if x, ok := state["ipv6_cidr_block"]; ok {
		x := x.(string)
		r.Ipv6_cidr_block = &x
	}

	if x, ok := state["availability_zone"]; ok {
		x := x.(string)
		r.Availability_zone = &x
	}

	if x, ok := state["map_public_ip_on_launch"]; ok {
		x := x.(bool)
		r.Map_public_ip_on_launch = &x
	}
	return r
}

// Aws_subnetHandler ...
type Aws_subnetHandler struct {
	Provider *schema.Provider
}

// Create ...
func (h *Aws_subnetHandler) Create(desired *Aws_subnet) (*Aws_subnet, string, error) {
	rState := Aws_subnetMapper(desired)
	id, err := bridge.Create(h.Provider, "aws_subnet", rState)
	if err != nil {
		return nil, "", err
	}
	actual, err := h.Read(id)
	if err != nil {
		return nil, "", err
	}
	return actual, id, nil
}

// Read ...
func (h *Aws_subnetHandler) Read(externalID string) (*Aws_subnet, error) {
	actual, err := bridge.Read(h.Provider, "aws_subnet", externalID)
	if err != nil {
		return nil, err
	}
	return Aws_subnetUnmapper(actual), nil
}

// Delete ...
func (h *Aws_subnetHandler) Delete(externalID string) error {
	return bridge.Delete(h.Provider, "aws_subnet", externalID)
}

type Aws_vpc struct {
	Enable_dns_support               *bool
	Main_route_table_id              string
	Tags                             *map[string]interface{}
	Owner_id                         string
	Enable_classiclink               *bool
	Assign_generated_ipv6_cidr_block *bool
	Default_network_acl_id           string
	Default_route_table_id           string
	Ipv6_cidr_block                  string
	Cidr_block                       string
	Instance_tenancy                 *string
	Enable_dns_hostnames             *bool
	Enable_classiclink_dns_support   *bool
	Dhcp_options_id                  string
	Ipv6_association_id              string
	Default_security_group_id        string
	Arn                              string
}

func Aws_vpcMapper(r *Aws_vpc) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Enable_dns_hostnames != nil {
		config["enable_dns_hostnames"] = *r.Enable_dns_hostnames
	}
	if r.Enable_classiclink_dns_support != nil {
		config["enable_classiclink_dns_support"] = *r.Enable_classiclink_dns_support
	}
	config["dhcp_options_id"] = r.Dhcp_options_id
	config["ipv6_association_id"] = r.Ipv6_association_id
	config["cidr_block"] = r.Cidr_block
	if r.Instance_tenancy != nil {
		config["instance_tenancy"] = *r.Instance_tenancy
	}
	config["default_security_group_id"] = r.Default_security_group_id
	config["arn"] = r.Arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	if r.Enable_dns_support != nil {
		config["enable_dns_support"] = *r.Enable_dns_support
	}
	config["main_route_table_id"] = r.Main_route_table_id
	config["default_network_acl_id"] = r.Default_network_acl_id
	config["default_route_table_id"] = r.Default_route_table_id
	config["ipv6_cidr_block"] = r.Ipv6_cidr_block
	if r.Enable_classiclink != nil {
		config["enable_classiclink"] = *r.Enable_classiclink
	}
	if r.Assign_generated_ipv6_cidr_block != nil {
		config["assign_generated_ipv6_cidr_block"] = *r.Assign_generated_ipv6_cidr_block
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

	if x, ok := state["arn"]; ok {
		r.Arn = x.(string)
	}

	if x, ok := state["enable_dns_support"]; ok {
		x := x.(bool)
		r.Enable_dns_support = &x
	}

	if x, ok := state["main_route_table_id"]; ok {
		r.Main_route_table_id = x.(string)
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

	if x, ok := state["assign_generated_ipv6_cidr_block"]; ok {
		x := x.(bool)
		r.Assign_generated_ipv6_cidr_block = &x
	}

	if x, ok := state["default_network_acl_id"]; ok {
		r.Default_network_acl_id = x.(string)
	}

	if x, ok := state["default_route_table_id"]; ok {
		r.Default_route_table_id = x.(string)
	}

	if x, ok := state["ipv6_cidr_block"]; ok {
		r.Ipv6_cidr_block = x.(string)
	}

	if x, ok := state["ipv6_association_id"]; ok {
		r.Ipv6_association_id = x.(string)
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

	if x, ok := state["enable_classiclink_dns_support"]; ok {
		x := x.(bool)
		r.Enable_classiclink_dns_support = &x
	}

	if x, ok := state["dhcp_options_id"]; ok {
		r.Dhcp_options_id = x.(string)
	}
	return r
}

// Aws_vpcHandler ...
type Aws_vpcHandler struct {
	Provider *schema.Provider
}

// Create ...
func (h *Aws_vpcHandler) Create(desired *Aws_vpc) (*Aws_vpc, string, error) {
	rState := Aws_vpcMapper(desired)
	id, err := bridge.Create(h.Provider, "aws_vpc", rState)
	if err != nil {
		return nil, "", err
	}
	actual, err := h.Read(id)
	if err != nil {
		return nil, "", err
	}
	return actual, id, nil
}

// Read ...
func (h *Aws_vpcHandler) Read(externalID string) (*Aws_vpc, error) {
	actual, err := bridge.Read(h.Provider, "aws_vpc", externalID)
	if err != nil {
		return nil, err
	}
	return Aws_vpcUnmapper(actual), nil
}

// Delete ...
func (h *Aws_vpcHandler) Delete(externalID string) error {
	return bridge.Delete(h.Provider, "aws_vpc", externalID)
}
