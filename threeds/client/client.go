// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/Basis-Theory/go-sdk/v2/core"
	internal "github.com/Basis-Theory/go-sdk/v2/internal"
	option "github.com/Basis-Theory/go-sdk/v2/option"
	sessions "github.com/Basis-Theory/go-sdk/v2/threeds/sessions"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	Sessions *sessions.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.APIKey == "" {
		options.APIKey = os.Getenv("BT-API-KEY")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:   options.ToHeader(),
		Sessions: sessions.NewClient(opts...),
	}
}
