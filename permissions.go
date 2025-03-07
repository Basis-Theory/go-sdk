// This file was auto-generated by Fern from our API Definition.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/Basis-Theory/go-sdk/core"
)

type PermissionsListRequest struct {
	ApplicationType *string `json:"-" url:"application_type,omitempty"`
}

type Permission struct {
	Type             *string  `json:"type,omitempty" url:"type,omitempty"`
	Description      *string  `json:"description,omitempty" url:"description,omitempty"`
	ApplicationTypes []string `json:"application_types,omitempty" url:"application_types,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *Permission) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *Permission) UnmarshalJSON(data []byte) error {
	type unmarshaler Permission
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = Permission(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *Permission) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}
