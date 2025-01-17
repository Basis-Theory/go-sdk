// This file was auto-generated by Fern from our API Definition.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/Basis-Theory/go-sdk/core"
)

type BankVerificationRequest struct {
	TokenID       string  `json:"token_id" url:"-"`
	CountryCode   *string `json:"country_code,omitempty" url:"-"`
	RoutingNumber *string `json:"routing_number,omitempty" url:"-"`
}

type BankVerificationResponse struct {
	Status *string `json:"status,omitempty" url:"status,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BankVerificationResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BankVerificationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BankVerificationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BankVerificationResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BankVerificationResponse) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}
