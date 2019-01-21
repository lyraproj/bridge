package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scottyw/lyra-bridge/pkg/bridge"
	"github.com/scottyw/lyra-bridge/pkg/generated"
)

// Aws_vpcHandler ...
type Aws_vpcHandler struct {
	Provider *schema.Provider
}

// Create a VPC
func (h *Aws_vpcHandler) Create(desired *generated.Aws_vpc) (*generated.Aws_vpc, string, error) {
	rState := generated.Aws_vpcMapper(desired)
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

// Read a VPC
func (h *Aws_vpcHandler) Read(externalID string) (*generated.Aws_vpc, error) {
	actual, err := bridge.Read(h.Provider, "aws_vpc", externalID)
	if err != nil {
		return nil, err
	}
	return generated.Aws_vpcUnmapper(actual), nil
}

// Delete a VPC
func (h *Aws_vpcHandler) Delete(externalID string) error {
	return bridge.Delete(h.Provider, "aws_vpc", externalID)
}
