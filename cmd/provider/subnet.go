package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scottyw/lyra-bridge/pkg/bridge"
	"github.com/scottyw/lyra-bridge/pkg/generated"
)

// Aws_subnetHandler ...
type Aws_subnetHandler struct {
	Provider *schema.Provider
}

// Create a subnet
func (h *Aws_subnetHandler) Create(desired *generated.Aws_subnet) (*generated.Aws_subnet, string, error) {
	rState := generated.Aws_subnetMapper(desired)
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

// Read a subnet
func (h *Aws_subnetHandler) Read(externalID string) (*generated.Aws_subnet, error) {
	actual, err := bridge.Read(h.Provider, "aws_subnet", externalID)
	if err != nil {
		return nil, err
	}
	return generated.Aws_subnetUnmapper(actual), nil
}

// Delete a subnet
func (h *Aws_subnetHandler) Delete(externalID string) error {
	return bridge.Delete(h.Provider, "aws_subnet", externalID)
}
