package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

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
		if k == "type" {
			k = "resource_type"
		}
		goType := getType(v)
		if goType != "" {
			fmt.Println("    ", k, goType)
		} else {
			log.Printf("Ignoring unsupported schema: %s -> %s -> %v", rType, k, v.Type)
		}
	}
	fmt.Printf("}\n\n")
}

func main() {
	fmt.Printf("package awstypes\n\n// Goodnight sweet linter, I'm sorry\n\n")
	p := aws.Provider().(*schema.Provider)
	for rType, r := range p.ResourcesMap {
		// if rType != "aws_vpc" {
		// 	continue
		// }
		generateResource(rType, r)
	}
}
