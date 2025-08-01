// Code generated by Fern. DO NOT EDIT.

package client

import (
	core "github.com/Basis-Theory/go-sdk/v2/core"
	internal "github.com/Basis-Theory/go-sdk/v2/internal"
	option "github.com/Basis-Theory/go-sdk/v2/option"
	connections "github.com/Basis-Theory/go-sdk/v2/tenants/connections"
	invitations "github.com/Basis-Theory/go-sdk/v2/tenants/invitations"
	members "github.com/Basis-Theory/go-sdk/v2/tenants/members"
	owner "github.com/Basis-Theory/go-sdk/v2/tenants/owner"
	self "github.com/Basis-Theory/go-sdk/v2/tenants/self"
	http "net/http"
	os "os"
)

type Client struct {
	Connections *connections.Client
	Invitations *invitations.Client
	Members     *members.Client
	Owner       *owner.Client
	Self        *self.Client

	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.APIKey == "" {
		options.APIKey = os.Getenv("BT-API-KEY")
	}
	return &Client{
		Connections: connections.NewClient(opts...),
		Invitations: invitations.NewClient(opts...),
		Members:     members.NewClient(opts...),
		Owner:       owner.NewClient(opts...),
		Self:        self.NewClient(opts...),
		baseURL:     options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}
