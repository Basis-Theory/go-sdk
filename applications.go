// Code generated by Fern. DO NOT EDIT.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Basis-Theory/go-sdk/v2/internal"
)

type CreateApplicationRequest struct {
	Name        string        `json:"name" url:"-"`
	Type        string        `json:"type" url:"-"`
	Permissions []string      `json:"permissions,omitempty" url:"-"`
	Rules       []*AccessRule `json:"rules,omitempty" url:"-"`
	CreateKey   *bool         `json:"create_key,omitempty" url:"-"`
}

type ApplicationsListRequest struct {
	ID    []*string `json:"-" url:"id,omitempty"`
	Type  []*string `json:"-" url:"type,omitempty"`
	Page  *int      `json:"-" url:"page,omitempty"`
	Start *string   `json:"-" url:"start,omitempty"`
	Size  *int      `json:"-" url:"size,omitempty"`
}

type ApplicationPaginatedList struct {
	Pagination *Pagination    `json:"pagination,omitempty" url:"pagination,omitempty"`
	Data       []*Application `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApplicationPaginatedList) GetPagination() *Pagination {
	if a == nil {
		return nil
	}
	return a.Pagination
}

func (a *ApplicationPaginatedList) GetData() []*Application {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ApplicationPaginatedList) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApplicationPaginatedList) UnmarshalJSON(data []byte) error {
	type unmarshaler ApplicationPaginatedList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApplicationPaginatedList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApplicationPaginatedList) String() string {
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

type UpdateApplicationRequest struct {
	Name        string        `json:"name" url:"-"`
	Permissions []string      `json:"permissions,omitempty" url:"-"`
	Rules       []*AccessRule `json:"rules,omitempty" url:"-"`
}
