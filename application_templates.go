// This file was auto-generated by Fern from our API Definition.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Basis-Theory/go-sdk/internal"
)

type ApplicationTemplate struct {
	ID              *string       `json:"id,omitempty" url:"id,omitempty"`
	Name            *string       `json:"name,omitempty" url:"name,omitempty"`
	Description     *string       `json:"description,omitempty" url:"description,omitempty"`
	ApplicationType *string       `json:"application_type,omitempty" url:"application_type,omitempty"`
	TemplateType    *string       `json:"template_type,omitempty" url:"template_type,omitempty"`
	IsStarter       *bool         `json:"is_starter,omitempty" url:"is_starter,omitempty"`
	Rules           []*AccessRule `json:"rules,omitempty" url:"rules,omitempty"`
	Permissions     []string      `json:"permissions,omitempty" url:"permissions,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApplicationTemplate) GetID() *string {
	if a == nil {
		return nil
	}
	return a.ID
}

func (a *ApplicationTemplate) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *ApplicationTemplate) GetDescription() *string {
	if a == nil {
		return nil
	}
	return a.Description
}

func (a *ApplicationTemplate) GetApplicationType() *string {
	if a == nil {
		return nil
	}
	return a.ApplicationType
}

func (a *ApplicationTemplate) GetTemplateType() *string {
	if a == nil {
		return nil
	}
	return a.TemplateType
}

func (a *ApplicationTemplate) GetIsStarter() *bool {
	if a == nil {
		return nil
	}
	return a.IsStarter
}

func (a *ApplicationTemplate) GetRules() []*AccessRule {
	if a == nil {
		return nil
	}
	return a.Rules
}

func (a *ApplicationTemplate) GetPermissions() []string {
	if a == nil {
		return nil
	}
	return a.Permissions
}

func (a *ApplicationTemplate) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApplicationTemplate) UnmarshalJSON(data []byte) error {
	type unmarshaler ApplicationTemplate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApplicationTemplate(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApplicationTemplate) String() string {
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
