package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

func mungeFieldName(k string) string {
	if k == "type" {
		k = "resource_type"
	}
	return strings.Title(k)
}

func getType(s *schema.Schema) string {
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
	if s.Optional {
		return "*" + t
	}
	return t
}

func generateResource(rType string, r *schema.Resource) {
	fmt.Printf("type %s struct {\n", rType)
	for k, v := range r.Schema {
		k = mungeFieldName(k)
		goType := getType(v)
		if goType != "" {
			fmt.Println("    ", k, goType)
		} else {
			log.Printf("Ignoring unsupported schema: %s -> %s -> %v", rType, k, v.Type)
		}
	}
	fmt.Printf("}\n\n")
}

func generateMapper(rType string, r *schema.Resource) {
	fmt.Printf("func %sMapper(r *%s) *terraform.ResourceConfig {\n", rType, rType)
	fmt.Printf("    return &terraform.ResourceConfig{\n")
	fmt.Printf("    Config: map[string]interface{}{\n")
	for k, v := range r.Schema {
		k = mungeFieldName(k)
		goType := getType(v)
		if goType != "" {
			fmt.Printf("    \"%s\":    r.%s,\n", k, k)
		} else {
			log.Printf("Ignoring unsupported schema: %s -> %s -> %v", rType, k, v.Type)
		}
	}
	fmt.Printf("},\n}\n}\n\n")
}

func generateUnmapper(rType string, r *schema.Resource) {
	fmt.Printf("func %sUnmapper(r *terraform.ResourceConfig) *%s {\n", rType, rType)
	fmt.Printf("    return &%s{\n", rType)
	for k, v := range r.Schema {
		k = mungeFieldName(k)
		goType := getType(v)
		if goType != "" {
			fmt.Printf("    %s:    r.Config[\"%s\"].(%s),\n", k, k, goType)
		} else {
			log.Printf("Ignoring unsupported schema: %s -> %s -> %v", rType, k, v.Type)
		}
	}
	fmt.Printf("}\n}\n\n")
}

var prefix = `
package awstypes

import (
	"github.com/hashicorp/terraform/terraform"
)

// Goodnight sweet linter, I'm sorry

`

func main() {
	fmt.Printf(prefix)
	p := aws.Provider().(*schema.Provider)
	for rType, r := range p.ResourcesMap {
		if rType != "aws_vpc" {
			continue
		}
		generateResource(rType, r)
		generateMapper(rType, r)
		generateUnmapper(rType, r)
	}
}
