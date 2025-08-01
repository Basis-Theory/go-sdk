// Code generated by Fern. DO NOT EDIT.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Basis-Theory/go-sdk/v2/internal"
	time "time"
)

type ApplePayCreateRequest struct {
	ExpiresAt        *string              `json:"expires_at,omitempty" url:"-"`
	ApplePaymentData *ApplePayMethodToken `json:"apple_payment_data,omitempty" url:"-"`
}

type ApplePayCreateResponse struct {
	ApplePay *ApplePayCreateTokenResponse `json:"apple_pay,omitempty" url:"apple_pay,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApplePayCreateResponse) GetApplePay() *ApplePayCreateTokenResponse {
	if a == nil {
		return nil
	}
	return a.ApplePay
}

func (a *ApplePayCreateResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApplePayCreateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ApplePayCreateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApplePayCreateResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApplePayCreateResponse) String() string {
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

type ApplePayCreateTokenResponse struct {
	ID        *string      `json:"id,omitempty" url:"id,omitempty"`
	Type      *string      `json:"type,omitempty" url:"type,omitempty"`
	TenantID  *string      `json:"tenant_id,omitempty" url:"tenant_id,omitempty"`
	Status    *string      `json:"status,omitempty" url:"status,omitempty"`
	ExpiresAt *time.Time   `json:"expires_at,omitempty" url:"expires_at,omitempty"`
	CreatedBy *string      `json:"created_by,omitempty" url:"created_by,omitempty"`
	CreatedAt *time.Time   `json:"created_at,omitempty" url:"created_at,omitempty"`
	Card      *CardDetails `json:"card,omitempty" url:"card,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApplePayCreateTokenResponse) GetID() *string {
	if a == nil {
		return nil
	}
	return a.ID
}

func (a *ApplePayCreateTokenResponse) GetType() *string {
	if a == nil {
		return nil
	}
	return a.Type
}

func (a *ApplePayCreateTokenResponse) GetTenantID() *string {
	if a == nil {
		return nil
	}
	return a.TenantID
}

func (a *ApplePayCreateTokenResponse) GetStatus() *string {
	if a == nil {
		return nil
	}
	return a.Status
}

func (a *ApplePayCreateTokenResponse) GetExpiresAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.ExpiresAt
}

func (a *ApplePayCreateTokenResponse) GetCreatedBy() *string {
	if a == nil {
		return nil
	}
	return a.CreatedBy
}

func (a *ApplePayCreateTokenResponse) GetCreatedAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.CreatedAt
}

func (a *ApplePayCreateTokenResponse) GetCard() *CardDetails {
	if a == nil {
		return nil
	}
	return a.Card
}

func (a *ApplePayCreateTokenResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApplePayCreateTokenResponse) UnmarshalJSON(data []byte) error {
	type embed ApplePayCreateTokenResponse
	var unmarshaler = struct {
		embed
		ExpiresAt *internal.DateTime `json:"expires_at,omitempty"`
		CreatedAt *internal.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = ApplePayCreateTokenResponse(unmarshaler.embed)
	a.ExpiresAt = unmarshaler.ExpiresAt.TimePtr()
	a.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApplePayCreateTokenResponse) MarshalJSON() ([]byte, error) {
	type embed ApplePayCreateTokenResponse
	var marshaler = struct {
		embed
		ExpiresAt *internal.DateTime `json:"expires_at,omitempty"`
		CreatedAt *internal.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*a),
		ExpiresAt: internal.NewOptionalDateTime(a.ExpiresAt),
		CreatedAt: internal.NewOptionalDateTime(a.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (a *ApplePayCreateTokenResponse) String() string {
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

type ApplePayToken struct {
	ID             *string         `json:"id,omitempty" url:"id,omitempty"`
	Type           *string         `json:"type,omitempty" url:"type,omitempty"`
	TenantID       *string         `json:"tenant_id,omitempty" url:"tenant_id,omitempty"`
	Status         *string         `json:"status,omitempty" url:"status,omitempty"`
	ExpiresAt      *time.Time      `json:"expires_at,omitempty" url:"expires_at,omitempty"`
	CreatedBy      *string         `json:"created_by,omitempty" url:"created_by,omitempty"`
	CreatedAt      *time.Time      `json:"created_at,omitempty" url:"created_at,omitempty"`
	ModifiedBy     *string         `json:"modified_by,omitempty" url:"modified_by,omitempty"`
	ModifiedAt     *time.Time      `json:"modified_at,omitempty" url:"modified_at,omitempty"`
	Card           *CardDetails    `json:"card,omitempty" url:"card,omitempty"`
	Data           interface{}     `json:"data,omitempty" url:"data,omitempty"`
	Authentication *Authentication `json:"authentication,omitempty" url:"authentication,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApplePayToken) GetID() *string {
	if a == nil {
		return nil
	}
	return a.ID
}

func (a *ApplePayToken) GetType() *string {
	if a == nil {
		return nil
	}
	return a.Type
}

func (a *ApplePayToken) GetTenantID() *string {
	if a == nil {
		return nil
	}
	return a.TenantID
}

func (a *ApplePayToken) GetStatus() *string {
	if a == nil {
		return nil
	}
	return a.Status
}

func (a *ApplePayToken) GetExpiresAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.ExpiresAt
}

func (a *ApplePayToken) GetCreatedBy() *string {
	if a == nil {
		return nil
	}
	return a.CreatedBy
}

func (a *ApplePayToken) GetCreatedAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.CreatedAt
}

func (a *ApplePayToken) GetModifiedBy() *string {
	if a == nil {
		return nil
	}
	return a.ModifiedBy
}

func (a *ApplePayToken) GetModifiedAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.ModifiedAt
}

func (a *ApplePayToken) GetCard() *CardDetails {
	if a == nil {
		return nil
	}
	return a.Card
}

func (a *ApplePayToken) GetData() interface{} {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ApplePayToken) GetAuthentication() *Authentication {
	if a == nil {
		return nil
	}
	return a.Authentication
}

func (a *ApplePayToken) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApplePayToken) UnmarshalJSON(data []byte) error {
	type embed ApplePayToken
	var unmarshaler = struct {
		embed
		ExpiresAt  *internal.DateTime `json:"expires_at,omitempty"`
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = ApplePayToken(unmarshaler.embed)
	a.ExpiresAt = unmarshaler.ExpiresAt.TimePtr()
	a.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	a.ModifiedAt = unmarshaler.ModifiedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApplePayToken) MarshalJSON() ([]byte, error) {
	type embed ApplePayToken
	var marshaler = struct {
		embed
		ExpiresAt  *internal.DateTime `json:"expires_at,omitempty"`
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed:      embed(*a),
		ExpiresAt:  internal.NewOptionalDateTime(a.ExpiresAt),
		CreatedAt:  internal.NewOptionalDateTime(a.CreatedAt),
		ModifiedAt: internal.NewOptionalDateTime(a.ModifiedAt),
	}
	return json.Marshal(marshaler)
}

func (a *ApplePayToken) String() string {
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

type Authentication struct {
	ThreedsCryptogram       *string                              `json:"threeds_cryptogram,omitempty" url:"threeds_cryptogram,omitempty"`
	EciIndicator            *string                              `json:"eci_indicator,omitempty" url:"eci_indicator,omitempty"`
	AuthenticationResponses []*SubmerchantAuthenticationResponse `json:"authentication_responses,omitempty" url:"authentication_responses,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *Authentication) GetThreedsCryptogram() *string {
	if a == nil {
		return nil
	}
	return a.ThreedsCryptogram
}

func (a *Authentication) GetEciIndicator() *string {
	if a == nil {
		return nil
	}
	return a.EciIndicator
}

func (a *Authentication) GetAuthenticationResponses() []*SubmerchantAuthenticationResponse {
	if a == nil {
		return nil
	}
	return a.AuthenticationResponses
}

func (a *Authentication) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *Authentication) UnmarshalJSON(data []byte) error {
	type unmarshaler Authentication
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = Authentication(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *Authentication) String() string {
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

type SubmerchantAuthenticationResponse struct {
	MerchantIdentifier *string `json:"merchant_identifier,omitempty" url:"merchant_identifier,omitempty"`
	AuthenticationData *string `json:"authentication_data,omitempty" url:"authentication_data,omitempty"`
	TransactionAmount  *string `json:"transaction_amount,omitempty" url:"transaction_amount,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SubmerchantAuthenticationResponse) GetMerchantIdentifier() *string {
	if s == nil {
		return nil
	}
	return s.MerchantIdentifier
}

func (s *SubmerchantAuthenticationResponse) GetAuthenticationData() *string {
	if s == nil {
		return nil
	}
	return s.AuthenticationData
}

func (s *SubmerchantAuthenticationResponse) GetTransactionAmount() *string {
	if s == nil {
		return nil
	}
	return s.TransactionAmount
}

func (s *SubmerchantAuthenticationResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SubmerchantAuthenticationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SubmerchantAuthenticationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SubmerchantAuthenticationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SubmerchantAuthenticationResponse) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
