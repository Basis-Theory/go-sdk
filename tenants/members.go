// Code generated by Fern. DO NOT EDIT.

package tenants

type MembersListRequest struct {
	UserID []*string `json:"-" url:"user_id,omitempty"`
	Page   *int      `json:"-" url:"page,omitempty"`
	Start  *string   `json:"-" url:"start,omitempty"`
	Size   *int      `json:"-" url:"size,omitempty"`
}

type UpdateTenantMemberRequest struct {
	Role string `json:"role" url:"-"`
}
