// This file was auto-generated by Fern from our API Definition.

package signingkey

import (
	bytes "bytes"
	context "context"
	core "github.com/Basis-Theory/go-sdk/core"
	option "github.com/Basis-Theory/go-sdk/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
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
		header: options.ToHeader(),
	}
}

// Returns the signing key
func (c *Client) Get(
	ctx context.Context,
	opts ...option.RequestOption,
) (string, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.basistheory.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/webhooks/signing-key"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	response := bytes.NewBuffer(nil)
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        response,
		},
	); err != nil {
		return "", err
	}
	return response.String(), nil
}
