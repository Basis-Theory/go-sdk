// This file was auto-generated by Fern from our API Definition.

package tenants

import (
	gosdk "github.com/Basis-Theory/go-sdk"
)

type CreateTenantConnectionRequest struct {
	Strategy string                         `json:"strategy" url:"-"`
	Options  *gosdk.TenantConnectionOptions `json:"options,omitempty" url:"-"`
}
