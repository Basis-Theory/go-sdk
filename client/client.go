// This file was auto-generated by Fern from our API Definition.

package client

import (
	applicationkeys "github.com/fern-demo/basis-theory-go/applicationkeys"
	applications "github.com/fern-demo/basis-theory-go/applications"
	applicationtemplates "github.com/fern-demo/basis-theory-go/applicationtemplates"
	core "github.com/fern-demo/basis-theory-go/core"
	logs "github.com/fern-demo/basis-theory-go/logs"
	option "github.com/fern-demo/basis-theory-go/option"
	permissions "github.com/fern-demo/basis-theory-go/permissions"
	proxies "github.com/fern-demo/basis-theory-go/proxies"
	reactorformulas "github.com/fern-demo/basis-theory-go/reactorformulas"
	reactors "github.com/fern-demo/basis-theory-go/reactors"
	roles "github.com/fern-demo/basis-theory-go/roles"
	sessions "github.com/fern-demo/basis-theory-go/sessions"
	tenantsclient "github.com/fern-demo/basis-theory-go/tenants/client"
	threeds "github.com/fern-demo/basis-theory-go/threeds"
	tokenize "github.com/fern-demo/basis-theory-go/tokenize"
	tokens "github.com/fern-demo/basis-theory-go/tokens"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Applications         *applications.Client
	ApplicationKeys      *applicationkeys.Client
	ApplicationTemplates *applicationtemplates.Client
	Logs                 *logs.Client
	Permissions          *permissions.Client
	Proxies              *proxies.Client
	Reactorformulas      *reactorformulas.Client
	Reactors             *reactors.Client
	Roles                *roles.Client
	Sessions             *sessions.Client
	Tenants              *tenantsclient.Client
	Threeds              *threeds.Client
	Tokenize             *tokenize.Client
	Tokens               *tokens.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.APIKey == "" {
		options.APIKey = os.Getenv("BT-API-KEY")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:               options.ToHeader(),
		Applications:         applications.NewClient(opts...),
		ApplicationKeys:      applicationkeys.NewClient(opts...),
		ApplicationTemplates: applicationtemplates.NewClient(opts...),
		Logs:                 logs.NewClient(opts...),
		Permissions:          permissions.NewClient(opts...),
		Proxies:              proxies.NewClient(opts...),
		Reactorformulas:      reactorformulas.NewClient(opts...),
		Reactors:             reactors.NewClient(opts...),
		Roles:                roles.NewClient(opts...),
		Sessions:             sessions.NewClient(opts...),
		Tenants:              tenantsclient.NewClient(opts...),
		Threeds:              threeds.NewClient(opts...),
		Tokenize:             tokenize.NewClient(opts...),
		Tokens:               tokens.NewClient(opts...),
	}
}
