// Code generated by Fern. DO NOT EDIT.

package roles

import (
	context "context"
	v2 "github.com/Basis-Theory/go-sdk/v2"
	core "github.com/Basis-Theory/go-sdk/v2/core"
	internal "github.com/Basis-Theory/go-sdk/v2/internal"
	option "github.com/Basis-Theory/go-sdk/v2/option"
	http "net/http"
)

type RawClient struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewRawClient(options *core.RequestOptions) *RawClient {
	return &RawClient{
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

func (r *RawClient) List(
	ctx context.Context,
	opts ...option.RequestOption,
) (*core.Response[[]*v2.Role], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.basistheory.com",
	)
	endpointURL := baseURL + "/roles"
	headers := internal.MergeHeaders(
		r.header.Clone(),
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
	}
	var response []*v2.Role
	raw, err := r.caller.Call(
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
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[[]*v2.Role]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}
