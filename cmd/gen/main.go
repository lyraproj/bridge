package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

func getGoType(s *schema.Schema) string {
	var t string
	switch s.Type {
	case schema.TypeString:
		t = "string"
	case schema.TypeMap:
		// FIXME won't always be string to interface{}?
		t = "map[string]interface{}"
	case schema.TypeBool:
		t = "bool"
	default:
		return ""
	}
	return t
}

func generateResource(rType string, r *schema.Resource) {
	fmt.Printf("type %s struct {\n", rType)
	for k, v := range r.Schema {
		if k == "type" {
			k = "resource_type"
		}
		goType := getGoType(v)
		if goType == "" {
			log.Printf("Ignoring unsupported schema: %s -> %s -> %v", rType, k, v.Type)
			continue
		}
		if v.Optional {
			goType = "*" + goType
		}
		fmt.Println("    ", strings.Title(k), goType)
	}
	fmt.Printf("}\n\n")
}

func generateMapper(rType string, r *schema.Resource) {
	fmt.Printf(mapperPrefix, rType, rType)
	for k, v := range r.Schema {
		if k == "type" {
			k = "resource_type"
		}
		if getGoType(v) == "" {
			log.Printf("Ignoring unsupported schema: %s -> %s -> %v", rType, k, v.Type)
			continue
		}
		if v.Optional {
			fmt.Printf("if r.%s != nil {\n", strings.Title(k))
			fmt.Printf("    config[\"%s\"] = *r.%s\n", k, strings.Title(k))
			fmt.Printf("}\n")
		} else {
			fmt.Printf("    config[\"%s\"] = r.%s\n", k, strings.Title(k))
		}
	}
	fmt.Printf(mapperSuffix)
}

func generateUnmapper(rType string, r *schema.Resource) {
	fmt.Printf(unmapperPrefix, rType, rType, rType)
	for k, v := range r.Schema {
		if k == "type" {
			k = "resource_type"
		}
		goType := getGoType(v)
		if goType == "" {
			log.Printf("Ignoring unsupported schema: %s -> %s -> %v", rType, k, v.Type)
			continue
		}
		if v.Optional {
			fmt.Printf(unmapperWithPointerDeref, k, goType, strings.Title(k))
		} else {
			fmt.Printf(unmapper, k, strings.Title(k), goType)
		}
	}
	fmt.Printf(unmapperSuffix)
}

var mapperPrefix = `
func %sMapper(r *%s) *terraform.ResourceConfig {
	config := map[string]interface{}{}
 	`

var mapperSuffix = `return &terraform.ResourceConfig{
		Config: config,
	}
}
`

var unmapperPrefix = `
func %sUnmapper(state map[string]interface{}) *%s {
	r := &%s{}
`

var unmapper = `
if x, ok := state["%s"]; ok {
	r.%s = x.(%s)
}
`

var unmapperWithPointerDeref = `
if x, ok := state["%s"]; ok {
	x := x.(%s)
	r.%s = &x
}
`

var unmapperSuffix = `	return r
}
`

var prefix = `
package generated

import (
	"github.com/hashicorp/terraform/terraform"
)

// Goodnight sweet linter, I'm sorry

`

func main() {
	fmt.Printf(prefix)
	p := aws.Provider().(*schema.Provider)
	for rType, r := range p.ResourcesMap {
		if rType != "aws_vpc" && rType != "aws_subnet" {
			continue
		}
		rType = strings.Title(rType)
		generateResource(rType, r)
		generateMapper(rType, r)
		generateUnmapper(rType, r)
	}
}
