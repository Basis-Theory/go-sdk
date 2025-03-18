// This file was auto-generated by Fern from our API Definition.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Basis-Theory/go-sdk/internal"
)

type ApplePayTokenizeRequest struct {
	ApplePaymentMethodToken *ApplePayMethodToken `json:"apple_payment_method_token,omitempty" url:"-"`
}

type ApplePayMethodToken struct {
	PaymentData           *PaymentData `json:"paymentData,omitempty" url:"paymentData,omitempty"`
	TransactionIdentifier *string      `json:"transactionIdentifier,omitempty" url:"transactionIdentifier,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApplePayMethodToken) GetPaymentData() *PaymentData {
	if a == nil {
		return nil
	}
	return a.PaymentData
}

func (a *ApplePayMethodToken) GetTransactionIdentifier() *string {
	if a == nil {
		return nil
	}
	return a.TransactionIdentifier
}

func (a *ApplePayMethodToken) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApplePayMethodToken) UnmarshalJSON(data []byte) error {
	type unmarshaler ApplePayMethodToken
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApplePayMethodToken(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApplePayMethodToken) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ApplePayTokenizeResponse struct {
	TokenIntent *CreateTokenIntentResponse `json:"token_intent,omitempty" url:"token_intent,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApplePayTokenizeResponse) GetTokenIntent() *CreateTokenIntentResponse {
	if a == nil {
		return nil
	}
	return a.TokenIntent
}

func (a *ApplePayTokenizeResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApplePayTokenizeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ApplePayTokenizeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApplePayTokenizeResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApplePayTokenizeResponse) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type Header struct {
	PublicKeyHash      *string `json:"publicKeyHash,omitempty" url:"publicKeyHash,omitempty"`
	EphemeralPublicKey *string `json:"ephemeralPublicKey,omitempty" url:"ephemeralPublicKey,omitempty"`
	TransactionID      *string `json:"transactionId,omitempty" url:"transactionId,omitempty"`
	ApplicationData    *string `json:"applicationData,omitempty" url:"applicationData,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (h *Header) GetPublicKeyHash() *string {
	if h == nil {
		return nil
	}
	return h.PublicKeyHash
}

func (h *Header) GetEphemeralPublicKey() *string {
	if h == nil {
		return nil
	}
	return h.EphemeralPublicKey
}

func (h *Header) GetTransactionID() *string {
	if h == nil {
		return nil
	}
	return h.TransactionID
}

func (h *Header) GetApplicationData() *string {
	if h == nil {
		return nil
	}
	return h.ApplicationData
}

func (h *Header) GetExtraProperties() map[string]interface{} {
	return h.extraProperties
}

func (h *Header) UnmarshalJSON(data []byte) error {
	type unmarshaler Header
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*h = Header(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *h)
	if err != nil {
		return err
	}
	h.extraProperties = extraProperties
	h.rawJSON = json.RawMessage(data)
	return nil
}

func (h *Header) String() string {
	if len(h.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(h.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(h); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", h)
}

type PaymentData struct {
	Data      *string `json:"data,omitempty" url:"data,omitempty"`
	Signature *string `json:"signature,omitempty" url:"signature,omitempty"`
	Header    *Header `json:"header,omitempty" url:"header,omitempty"`
	Version   *string `json:"version,omitempty" url:"version,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaymentData) GetData() *string {
	if p == nil {
		return nil
	}
	return p.Data
}

func (p *PaymentData) GetSignature() *string {
	if p == nil {
		return nil
	}
	return p.Signature
}

func (p *PaymentData) GetHeader() *Header {
	if p == nil {
		return nil
	}
	return p.Header
}

func (p *PaymentData) GetVersion() *string {
	if p == nil {
		return nil
	}
	return p.Version
}

func (p *PaymentData) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaymentData) UnmarshalJSON(data []byte) error {
	type unmarshaler PaymentData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaymentData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaymentData) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}
