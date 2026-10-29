package config

import (
	"encoding/json"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// TestSchemaResourcesHaveExternalNameConfig guards against a Terraform
// provider bump adding a resource that is silently left unwired.
func TestSchemaResourcesHaveExternalNameConfig(t *testing.T) {
	var schema struct {
		ProviderSchemas map[string]struct {
			ResourceSchemas map[string]json.RawMessage `json:"resource_schemas"`
		} `json:"provider_schemas"`
	}
	if err := json.Unmarshal([]byte(providerSchema), &schema); err != nil {
		t.Fatalf("cannot parse embedded schema.json: %v", err)
	}
	for _, p := range schema.ProviderSchemas {
		for name := range p.ResourceSchemas {
			if _, ok := ExternalNameConfigs[name]; !ok {
				t.Errorf("resource %q exists in schema.json but has no ExternalNameConfigs entry", name)
			}
		}
		if len(p.ResourceSchemas) != len(ExternalNameConfigs) {
			t.Errorf("schema.json has %d resources, ExternalNameConfigs has %d", len(p.ResourceSchemas), len(ExternalNameConfigs))
		}
	}
}

// TestNumericStateIDAsExternalName reproduces the bunny.net numeric id
// problem: 16 of 24 resources declare `id` as a number, so Terraform writes
// e.g. {"id": 577654} into the state. Upstream upjet only accepts string ids
// ("cannot work with a non-string id"); the wildbitca/upjet fork accepts them.
func TestNumericStateIDAsExternalName(t *testing.T) {
	cases := map[string]struct {
		state map[string]any
		want  string
	}{
		"Numeric": {state: map[string]any{"id": float64(577654)}, want: "577654"},
		"String":  {state: map[string]any{"id": "abc-123"}, want: "abc-123"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := ujconfig.IdentifierFromProvider.GetExternalNameFn(tc.state)
			if err != nil {
				t.Fatalf("GetExternalNameFn(%v): %v", tc.state, err)
			}
			if got != tc.want {
				t.Errorf("GetExternalNameFn(%v) = %q, want %q", tc.state, got, tc.want)
			}
		})
	}
}

func TestGetProvider(t *testing.T) {
	if got := len(GetProvider().Resources); got != len(ExternalNameConfigs) {
		t.Errorf("GetProvider() configured %d resources, want %d", got, len(ExternalNameConfigs))
	}
	if got := len(GetProviderNamespaced().Resources); got != len(ExternalNameConfigs) {
		t.Errorf("GetProviderNamespaced() configured %d resources, want %d", got, len(ExternalNameConfigs))
	}
}
