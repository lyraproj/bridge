package bridge

import (
	"github.com/hashicorp/terraform/flatmap"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Create a resource using the Terrform provider
func Create(p *schema.Provider, resourceType string, resourceConfig *terraform.ResourceConfig) (string, error) {
	// To get Terraform to create a new resource, the ID must be blank and existing state must be empty (since the
	// resource does not exist yet), and the diff object should have no old state and all of the new state.
	info := &terraform.InstanceInfo{Type: resourceType}
	state := &terraform.InstanceState{}
	diff, err := p.Diff(info, state, resourceConfig)
	if err != nil {
		return "", err
	}
	newstate, err := p.Apply(info, state, diff)
	if newstate == nil {
		// contract.Assertf(err != nil, "expected non-nil error with nil state during Create")
		return "", err
	}
	// contract.Assertf(newstate.ID != "", "Expected non-empty ID for new state during Create")
	return newstate.ID, nil
}

// Read a resource using the Terrform provider
func Read(p *schema.Provider, resourceType string, id string) (map[string]interface{}, error) {
	info := &terraform.InstanceInfo{Type: resourceType}
	state := &terraform.InstanceState{ID: id}
	newstate, err := p.Refresh(info, state)
	if err != nil {
		return nil, err
	}
	return expand(id, newstate), nil
}

// Delete a resource using the Terrform provider
func Delete(p *schema.Provider, resourceType string, id string) error {
	info := &terraform.InstanceInfo{Type: resourceType}
	state := &terraform.InstanceState{ID: id}
	diff := &terraform.InstanceDiff{Destroy: true}
	_, err := p.Apply(info, state, diff)
	return err
}

func expand(id string, state *terraform.InstanceState) map[string]interface{} {
	var outs map[string]interface{}
	if state != nil {
		outs = make(map[string]interface{})
		outs["external_id"] = id
		attrs := state.Attributes
		for _, key := range flatmap.Map(attrs).Keys() {
			outs[key] = flatmap.Expand(attrs, key)
		}
	}
	return outs
}
