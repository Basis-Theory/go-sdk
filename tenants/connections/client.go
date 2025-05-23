// This file was auto-generated by Fern from our API Definition.

package connections

import (
	context "context"
	v2 "github.com/Basis-Theory/go-sdk/v2"
	core "github.com/Basis-Theory/go-sdk/v2/core"
	internal "github.com/Basis-Theory/go-sdk/v2/internal"
	option "github.com/Basis-Theory/go-sdk/v2/option"
	tenants "github.com/Basis-Theory/go-sdk/v2/tenants"
	http "net/http"
	os "os"
)

type Client struct {
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
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

func (c *Client) Create(
	ctx context.Context,
	request *tenants.CreateTenantConnectionRequest,
	opts ...option.IdempotentRequestOption,
) (*v2.CreateTenantConnectionResponse, error) {
	options := core.NewIdempotentRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.basistheory.com",
	)
	endpointURL := baseURL + "/tenants/self/connections"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &v2.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &v2.ForbiddenError{
				APIError: apiError,
			}
		},
	}

	var response *v2.CreateTenantConnectionResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Delete(
	ctx context.Context,
	opts ...option.RequestOption,
) (*v2.CreateTenantConnectionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.basistheory.com",
	)
	endpointURL := baseURL + "/tenants/self/connections"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		401: func(apiError *core.APIError) error {
			return &v2.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &v2.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &v2.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *v2.CreateTenantConnectionResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
