// This file was auto-generated by Fern from our API Definition.

package invitations

import (
	context "context"
	fmt "fmt"
	gosdk "github.com/Basis-Theory/go-sdk"
	core "github.com/Basis-Theory/go-sdk/core"
	internal "github.com/Basis-Theory/go-sdk/internal"
	option "github.com/Basis-Theory/go-sdk/option"
	tenants "github.com/Basis-Theory/go-sdk/tenants"
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

func (c *Client) List(
	ctx context.Context,
	request *tenants.InvitationsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*gosdk.TenantInvitationResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.basistheory.com",
	)
	endpointURL := baseURL + "/tenants/self/invitations"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		401: func(apiError *core.APIError) error {
			return &gosdk.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &gosdk.ForbiddenError{
				APIError: apiError,
			}
		},
	}

	prepareCall := func(pageRequest *internal.PageRequest[*int]) *internal.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("page", fmt.Sprintf("%v", *pageRequest.Cursor))
		}
		nextURL := endpointURL
		if len(queryParams) > 0 {
			nextURL += "?" + queryParams.Encode()
		}
		return &internal.CallParams{
			URL:             nextURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        pageRequest.Response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		}
	}
	next := 1
	if request.Page != nil {
		next = *request.Page
	}
	readPageResponse := func(response *gosdk.TenantInvitationResponsePaginatedList) *internal.PageResponse[*int, *gosdk.TenantInvitationResponse] {
		next += 1
		results := response.Data
		return &internal.PageResponse[*int, *gosdk.TenantInvitationResponse]{
			Next:    &next,
			Results: results,
		}
	}
	pager := internal.NewOffsetPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, &next)
}

func (c *Client) Create(
	ctx context.Context,
	request *tenants.CreateTenantInvitationRequest,
	opts ...option.IdempotentRequestOption,
) (*gosdk.TenantInvitationResponse, error) {
	options := core.NewIdempotentRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.basistheory.com",
	)
	endpointURL := baseURL + "/tenants/self/invitations"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &gosdk.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &gosdk.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &gosdk.ForbiddenError{
				APIError: apiError,
			}
		},
	}

	var response *gosdk.TenantInvitationResponse
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

func (c *Client) Resend(
	ctx context.Context,
	invitationID string,
	opts ...option.IdempotentRequestOption,
) (*gosdk.TenantInvitationResponse, error) {
	options := core.NewIdempotentRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.basistheory.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/tenants/self/invitations/%v/resend",
		invitationID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &gosdk.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &gosdk.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &gosdk.ForbiddenError{
				APIError: apiError,
			}
		},
	}

	var response *gosdk.TenantInvitationResponse
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
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Get(
	ctx context.Context,
	invitationID string,
	opts ...option.RequestOption,
) (*gosdk.TenantInvitationResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.basistheory.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/tenants/self/invitations/%v",
		invitationID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		401: func(apiError *core.APIError) error {
			return &gosdk.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &gosdk.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &gosdk.NotFoundError{
				APIError: apiError,
			}
		},
	}

	var response *gosdk.TenantInvitationResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
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

func (c *Client) Delete(
	ctx context.Context,
	invitationID string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.basistheory.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/tenants/self/invitations/%v",
		invitationID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		401: func(apiError *core.APIError) error {
			return &gosdk.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &gosdk.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &gosdk.NotFoundError{
				APIError: apiError,
			}
		},
	}

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
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return err
	}
	return nil
}
