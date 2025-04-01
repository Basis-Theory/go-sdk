// This file was auto-generated by Fern from our API Definition.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Basis-Theory/go-sdk/internal"
)

type CreateNetworkTokenRequest struct {
	Data           *Card           `json:"data,omitempty" url:"-"`
	MerchantID     *string         `json:"merchant_id,omitempty" url:"-"`
	CardholderInfo *CardholderInfo `json:"cardholder_info,omitempty" url:"-"`
	Containers     []string        `json:"containers,omitempty" url:"-"`
}

type Address struct {
	Line1       *string `json:"line1,omitempty" url:"line1,omitempty"`
	Line2       *string `json:"line2,omitempty" url:"line2,omitempty"`
	Line3       *string `json:"line3,omitempty" url:"line3,omitempty"`
	PostalCode  *string `json:"postal_code,omitempty" url:"postal_code,omitempty"`
	City        *string `json:"city,omitempty" url:"city,omitempty"`
	StateCode   *string `json:"state_code,omitempty" url:"state_code,omitempty"`
	CountryCode *string `json:"country_code,omitempty" url:"country_code,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *Address) GetLine1() *string {
	if a == nil {
		return nil
	}
	return a.Line1
}

func (a *Address) GetLine2() *string {
	if a == nil {
		return nil
	}
	return a.Line2
}

func (a *Address) GetLine3() *string {
	if a == nil {
		return nil
	}
	return a.Line3
}

func (a *Address) GetPostalCode() *string {
	if a == nil {
		return nil
	}
	return a.PostalCode
}

func (a *Address) GetCity() *string {
	if a == nil {
		return nil
	}
	return a.City
}

func (a *Address) GetStateCode() *string {
	if a == nil {
		return nil
	}
	return a.StateCode
}

func (a *Address) GetCountryCode() *string {
	if a == nil {
		return nil
	}
	return a.CountryCode
}

func (a *Address) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *Address) UnmarshalJSON(data []byte) error {
	type unmarshaler Address
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = Address(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *Address) String() string {
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

type Card struct {
	Number          *string `json:"number,omitempty" url:"number,omitempty"`
	ExpirationMonth *int    `json:"expiration_month,omitempty" url:"expiration_month,omitempty"`
	ExpirationYear  *int    `json:"expiration_year,omitempty" url:"expiration_year,omitempty"`
	Cvc             *string `json:"cvc,omitempty" url:"cvc,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *Card) GetNumber() *string {
	if c == nil {
		return nil
	}
	return c.Number
}

func (c *Card) GetExpirationMonth() *int {
	if c == nil {
		return nil
	}
	return c.ExpirationMonth
}

func (c *Card) GetExpirationYear() *int {
	if c == nil {
		return nil
	}
	return c.ExpirationYear
}

func (c *Card) GetCvc() *string {
	if c == nil {
		return nil
	}
	return c.Cvc
}

func (c *Card) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *Card) UnmarshalJSON(data []byte) error {
	type unmarshaler Card
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = Card(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *Card) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CardholderInfo struct {
	Name    *string  `json:"name,omitempty" url:"name,omitempty"`
	Address *Address `json:"address,omitempty" url:"address,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CardholderInfo) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CardholderInfo) GetAddress() *Address {
	if c == nil {
		return nil
	}
	return c.Address
}

func (c *CardholderInfo) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CardholderInfo) UnmarshalJSON(data []byte) error {
	type unmarshaler CardholderInfo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CardholderInfo(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CardholderInfo) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}
