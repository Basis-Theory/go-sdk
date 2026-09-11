package custom

import (
	"testing"

	basistheory "github.com/Basis-Theory/go-sdk/v7"
)

// Pins the host each environment resolves to. Us and Eu both resolved to the
// compatibility host until the spec grew per-region servers, so the fields
// existed but selected nothing. These assertions fail if a regeneration
// collapses them back.
func TestEnvironmentURLs(t *testing.T) {
	for _, tc := range []struct {
		name string
		got  string
		want string
	}{
		{"Default", basistheory.Environments.Default, "https://api.basistheory.com"},
		{"Us", basistheory.Environments.Us, "https://api.us.basistheory.com"},
		{"Eu", basistheory.Environments.Eu, "https://api.eu.basistheory.com"},
		{"Test", basistheory.Environments.Test, "https://api.test.basistheory.com"},
	} {
		if tc.got != tc.want {
			t.Errorf("Environments.%s = %q, want %q", tc.name, tc.got, tc.want)
		}
	}
}

func TestEnvironmentsAreDistinct(t *testing.T) {
	seen := map[string]string{}
	for name, url := range map[string]string{
		"Default": basistheory.Environments.Default,
		"Us":      basistheory.Environments.Us,
		"Eu":      basistheory.Environments.Eu,
		"Test":    basistheory.Environments.Test,
	} {
		if other, ok := seen[url]; ok {
			t.Errorf("Environments.%s and Environments.%s both resolve to %q", name, other, url)
		}
		seen[url] = name
	}
}
