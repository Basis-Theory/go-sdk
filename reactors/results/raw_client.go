// Code generated by Fern. DO NOT EDIT.

package results

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

func (r *RawClient) Get(
	ctx context.Context,
	id string,
	requestID string,
	opts ...option.RequestOption,
) (*core.Response[any], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.basistheory.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/reactors/%v/results/%v",
		id,
		requestID,
	)
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
		404: func(apiError *core.APIError) error {
			return &v2.NotFoundError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &v2.UnprocessableEntityError{
				APIError: apiError,
			}
		},
	}
	var response any
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
	return &core.Response[any]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}
