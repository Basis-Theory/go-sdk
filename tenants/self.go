// This file was auto-generated by Fern from our API Definition.

package tenants

type UpdateTenantRequest struct {
	Name     string             `json:"name" url:"-"`
	Settings map[string]*string `json:"settings,omitempty" url:"-"`
}
