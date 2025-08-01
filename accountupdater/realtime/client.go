// Code generated by Fern. DO NOT EDIT.

package realtime

import (
	context "context"
	v2 "github.com/Basis-Theory/go-sdk/v2"
	accountupdater "github.com/Basis-Theory/go-sdk/v2/accountupdater"
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

// Returns the update result
func (c *Client) Invoke(
	ctx context.Context,
	request *accountupdater.AccountUpdaterRealTimeRequest,
	opts ...option.RequestOption,
) (*v2.AccountUpdaterRealTimeResponse, error) {
	response, err := c.WithRawResponse.Invoke(
		ctx,
		request,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
