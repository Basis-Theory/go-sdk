// This file was auto-generated by Fern from our API Definition.

package applepay

type ApplePayDomainDeregistrationRequest struct {
	Domain string `json:"domain" url:"-"`
}

type ApplePayDomainRegistrationRequest struct {
	Domain string `json:"domain" url:"-"`
}

type ApplePayDomainRegistrationListRequest struct {
	Domains []string `json:"domains,omitempty" url:"-"`
}
