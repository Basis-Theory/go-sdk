// This file was auto-generated by Fern from our API Definition.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Basis-Theory/go-sdk/internal"
)

type BankVerificationRequest struct {
	TokenID       string  `json:"token_id" url:"-"`
	CountryCode   *string `json:"country_code,omitempty" url:"-"`
	RoutingNumber *string `json:"routing_number,omitempty" url:"-"`
}

type BankVerificationResponse struct {
	Status *string `json:"status,omitempty" url:"status,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BankVerificationResponse) GetStatus() *string {
	if b == nil {
		return nil
	}
	return b.Status
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
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BankVerificationResponse) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}
