# BasisTheory Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com?utm_source=github&utm_medium=github&utm_campaign=readme&utm_source=https%3A%2F%2Fgithub.com%2FBasis-Theory%2Fgo-sdk)

The BasisTheory Go library provides convenient access to the BasisTheory APIs from Go.

## Table of Contents

- [Requirements](#requirements)
- [Usage](#usage)
- [Optional Parameters](#optional-parameters)
- [Automatic Pagination](#automatic-pagination)
- [Timeouts](#timeouts)
- [Request Options](#request-options)
- [Automatic Retries](#automatic-retries)
- [Errors](#errors)
- [Contributing](#contributing)
- [Documentation](#documentation)
- [Environments](#environments)
- [Advanced](#advanced)
  - [Response Headers](#response-headers)
  - [Retries](#retries)
  - [Timeouts](#timeouts)

## Requirements

This module requires Go version >= 1.18.

# Installation

Run the following command to use the basistheory Go library in your module:

```sh
go get github.com/basis-theory/go-sdk
```

## Usage

Instantiate and use the client with the following:

```go
package example

import (
    client "github.com/Basis-Theory/go-sdk/v5/client"
    option "github.com/Basis-Theory/go-sdk/v5/option"
    context "context"
)

func do() {
    client := client.NewClient(
        option.WithAPIKey(
            "<value>",
        ),
    )
    client.Tenants.Self.Get(
        context.TODO(),
    )
}
```

## Optional Parameters

This library models optional primitives and enum types as pointers. This is primarily meant to distinguish
default zero values from explicit values (e.g. `false` for `bool` and `""` for `string`). A collection of
helper functions are provided to easily map a primitive or enum to its pointer-equivalent (e.g. `basistheory.String`).

For example, consider the `client.Applications.Create` endpoint usage below:

```go
response, err := client.Applications.Create(
  context.TODO(),
  &basistheory.CreateApplicationRequest{
    Name:      "name",
    Type:      "type",
    CreateKey: basistheory.Bool(true),
  },
)
```

## Automatic Pagination

List endpoints are paginated. The SDK provides an iterator so that you can simply loop over the items:

```go
page, err := client.Applications.List(
  context.TODO(),
  &basistheory.ApplicationsListRequest{
    Page: basistheory.Int(1),
  },
)
if err != nil {
  return nil, err
}
iter := page.Iterator()
for iter.Next() {
  application := iter.Current()
  fmt.Printf("Got application: %v\n", application.Name)
}
if err := iter.Err(); err != nil {
  // Handle the error!
}
```

You can also iterate page-by-page:

```go
for page != nil {
  for _, event := range page.Results {
    fmt.Printf("Got event: %v\n", event.ID)
  }
  page, err = page.GetNextPage()
  if errors.Is(err, core.ErrNoPages) {
    break
  }
  if err != nil {
    // Handle the error!
  }
}
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard `context` library. Setting
a one second timeout for an individual API call looks like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := client.Applications.Create(
  ctx,
  &basistheory.CreateApplicationRequest{
    Name:      "name",
    Type:      "type",
  },
)
```

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes configuring
authorization tokens, or providing your own instrumented `*http.Client`.

These request options can either be
specified on the client so that they're applied on every request, or for an individual request, like so:

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

```go
// Specify default options applied on every request.
client := client.NewClient(
    option.WithToken("<YOUR_API_KEY>"),
    option.WithHTTPClient(
        &http.Client{
            Timeout: 5 * time.Second,
        },
    ),
)

// Specify options for an individual request.
response, err := client.Tenants.Self.Get(
    ...,
    option.WithToken("<YOUR_API_KEY>"),
)
```

## Automatic Retries

The basistheory Go client is instrumented with automatic retries with exponential backoff. A request will be
retried as long as the request is deemed retriable and the number of retry attempts has not grown larger
than the configured retry limit (default: 2).

A request is deemed retriable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

You can use the `option.WithMaxAttempts` option to configure the maximum retry limit to
your liking. For example, if you want to disable retries for the client entirely, you can
set this value to 1 like so:

```go
client := basistheoryclient.NewClient(
  option.WithMaxAttempts(1),
)
```

This can be done for an individual request, too:

```go
response, err := client.Applications.Create(
  ctx,
  &basistheory.CreateApplicationRequest{
    Name:      "name",
    Type:      "type",
  },
  option.WithMaxAttempts(1),
)
```

## Errors

Structured error types are returned from API calls that return non-success status codes. These errors are compatible
with the `errors.Is` and `errors.As` APIs, so you can access the error like so:

```go
response, err := client.Tenants.Self.Get(...)
if err != nil {
    var apiError *core.APIError
    if errors.As(err, apiError) {
        // Do something with the API error ...
    }
    return err
}
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
## Documentation

API reference documentation is available [here](https://api.basistheory.com).

## Environments

You can choose between different environments by using the `option.WithBaseURL` option. You can configure any arbitrary base
URL, which is particularly useful in test environments.

```go
client := client.NewClient(
    option.WithBaseURL(basistheory.Environments.Default),
)
```

## Advanced

### Response Headers

You can access the raw HTTP response data by using the `WithRawResponse` field on the client. This is useful
when you need to examine the response headers received from the API call.

```go
response, err := client.Tenants.Self.WithRawResponse.Get(...)
if err != nil {
    return err
}
fmt.Printf("Got response headers: %v", response.Header)
```

### Retries

The SDK is instrumented with automatic retries with exponential backoff. A request will be retried as long
as the request is deemed retryable and the number of retry attempts has not grown larger than the configured
retry limit (default: 2).

A request is deemed retryable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

Use the `option.WithMaxAttempts` option to configure this behavior for the entire client or an individual request:

```go
client := client.NewClient(
    option.WithMaxAttempts(1),
)

response, err := client.Tenants.Self.Get(
    ...,
    option.WithMaxAttempts(1),
)
```

### Timeouts

Setting a timeout for each individual request is as simple as using the standard context library. Setting a one second timeout for an individual API call looks like the following:

```go
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()

response, err := client.Tenants.Self.Get(ctx, ...)
```
