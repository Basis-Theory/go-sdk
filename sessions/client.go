// Code generated by Fern. DO NOT EDIT.

package sessions

import (
	context "context"
	v2 "github.com/Basis-Theory/go-sdk/v2"
	core "github.com/Basis-Theory/go-sdk/v2/core"
	internal "github.com/Basis-Theory/go-sdk/v2/internal"
	option "github.com/Basis-Theory/go-sdk/v2/option"
	http "net/http"
	os "os"
)

type Client struct {
	WithRawResponse *RawClient

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
		WithRawResponse: NewRawClient(options),
		baseURL:         options.BaseURL,
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
	opts ...option.IdempotentRequestOption,
) (*v2.CreateSessionResponse, error) {
	response, err := c.WithRawResponse.Create(
		ctx,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

func (c *Client) Authorize(
	ctx context.Context,
	request *v2.AuthorizeSessionRequest,
	opts ...option.IdempotentRequestOption,
) error {
	_, err := c.WithRawResponse.Authorize(
		ctx,
		request,
		opts...,
	)
	if err != nil {
		return err
	}
	return nil
}
