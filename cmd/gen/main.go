package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

var providerTemplate = `

// {{.TitleType}}Handler ...
type {{.TitleType}}Handler struct {
	Provider *schema.Provider
}

// Create ...
func (h *{{.TitleType}}Handler) Create(desired *{{.TitleType}}) (*{{.TitleType}}, string, error) {
	rState := {{.TitleType}}Mapper(desired)
	id, err := bridge.Create(h.Provider, "{{.TFType}}", rState)
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
func (h *{{.TitleType}}Handler) Read(externalID string) (*{{.TitleType}}, error) {
	actual, err := bridge.Read(h.Provider, "{{.TFType}}", externalID)
	if err != nil {
		return nil, err
	}
	return {{.TitleType}}Unmapper(actual), nil
}

// Delete ...
func (h *{{.TitleType}}Handler) Delete(externalID string) error {
	return bridge.Delete(h.Provider, "{{.TFType}}", externalID)
}

`

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
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scottyw/lyra-bridge/pkg/bridge"
	"github.com/hashicorp/terraform/terraform"
)

// Goodnight sweet linter, I'm sorry

var Resources = map[string]interface{} {
`

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

type providerType struct {
	TitleType string
	TFType    string
}

func generateProvider(rType string) {
	tmpl := template.Must(template.New("provider").Parse(providerTemplate))
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, providerType{strings.Title(rType), rType})
	if err != nil {
		panic(err)
	}
	fmt.Printf(buf.String())
}

func main() {
	fmt.Printf(prefix)

	p := aws.Provider().(*schema.Provider)

	for rType, _ := range p.ResourcesMap {
		// if rType != "aws_vpc" && rType != "aws_subnet" {
		// 	continue
		// }
		fmt.Printf("    \"%s\": %sHandler{},\n", rType, strings.Title(rType))
	}
	fmt.Printf("}\n\n")

	for rType, r := range p.ResourcesMap {
		// if rType != "aws_vpc" && rType != "aws_subnet" {
		// 	continue
		// }
		generateResource(strings.Title(rType), r)
		generateMapper(strings.Title(rType), r)
		generateUnmapper(strings.Title(rType), r)
		generateProvider(rType)
	}

}
