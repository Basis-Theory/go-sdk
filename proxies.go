// Code generated by Fern. DO NOT EDIT.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Basis-Theory/go-sdk/v2/internal"
	time "time"
)

type CreateProxyRequest struct {
	Name              string             `json:"name" url:"-"`
	DestinationURL    string             `json:"destination_url" url:"-"`
	RequestReactorID  *string            `json:"request_reactor_id,omitempty" url:"-"`
	ResponseReactorID *string            `json:"response_reactor_id,omitempty" url:"-"`
	RequestTransform  *ProxyTransform    `json:"request_transform,omitempty" url:"-"`
	ResponseTransform *ProxyTransform    `json:"response_transform,omitempty" url:"-"`
	Application       *Application       `json:"application,omitempty" url:"-"`
	Configuration     map[string]*string `json:"configuration,omitempty" url:"-"`
	RequireAuth       *bool              `json:"require_auth,omitempty" url:"-"`
}

type ProxiesListRequest struct {
	ID    []*string `json:"-" url:"id,omitempty"`
	Name  *string   `json:"-" url:"name,omitempty"`
	Page  *int      `json:"-" url:"page,omitempty"`
	Start *string   `json:"-" url:"start,omitempty"`
	Size  *int      `json:"-" url:"size,omitempty"`
}

type PatchProxyRequest struct {
	Name              *string            `json:"name,omitempty" url:"-"`
	DestinationURL    *string            `json:"destination_url,omitempty" url:"-"`
	RequestTransform  *ProxyTransform    `json:"request_transform,omitempty" url:"-"`
	ResponseTransform *ProxyTransform    `json:"response_transform,omitempty" url:"-"`
	Application       *Application       `json:"application,omitempty" url:"-"`
	Configuration     map[string]*string `json:"configuration,omitempty" url:"-"`
	RequireAuth       *bool              `json:"require_auth,omitempty" url:"-"`
}

type Proxy struct {
	ID                *string            `json:"id,omitempty" url:"id,omitempty"`
	Key               *string            `json:"key,omitempty" url:"key,omitempty"`
	TenantID          *string            `json:"tenant_id,omitempty" url:"tenant_id,omitempty"`
	Name              *string            `json:"name,omitempty" url:"name,omitempty"`
	DestinationURL    *string            `json:"destination_url,omitempty" url:"destination_url,omitempty"`
	RequestReactorID  *string            `json:"request_reactor_id,omitempty" url:"request_reactor_id,omitempty"`
	ResponseReactorID *string            `json:"response_reactor_id,omitempty" url:"response_reactor_id,omitempty"`
	RequireAuth       *bool              `json:"require_auth,omitempty" url:"require_auth,omitempty"`
	RequestTransform  *ProxyTransform    `json:"request_transform,omitempty" url:"request_transform,omitempty"`
	ResponseTransform *ProxyTransform    `json:"response_transform,omitempty" url:"response_transform,omitempty"`
	ApplicationID     *string            `json:"application_id,omitempty" url:"application_id,omitempty"`
	Configuration     map[string]*string `json:"configuration,omitempty" url:"configuration,omitempty"`
	ProxyHost         *string            `json:"proxy_host,omitempty" url:"proxy_host,omitempty"`
	Timeout           *int               `json:"timeout,omitempty" url:"timeout,omitempty"`
	ClientCertificate *string            `json:"client_certificate,omitempty" url:"client_certificate,omitempty"`
	CreatedBy         *string            `json:"created_by,omitempty" url:"created_by,omitempty"`
	CreatedAt         *time.Time         `json:"created_at,omitempty" url:"created_at,omitempty"`
	ModifiedBy        *string            `json:"modified_by,omitempty" url:"modified_by,omitempty"`
	ModifiedAt        *time.Time         `json:"modified_at,omitempty" url:"modified_at,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *Proxy) GetID() *string {
	if p == nil {
		return nil
	}
	return p.ID
}

func (p *Proxy) GetKey() *string {
	if p == nil {
		return nil
	}
	return p.Key
}

func (p *Proxy) GetTenantID() *string {
	if p == nil {
		return nil
	}
	return p.TenantID
}

func (p *Proxy) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *Proxy) GetDestinationURL() *string {
	if p == nil {
		return nil
	}
	return p.DestinationURL
}

func (p *Proxy) GetRequestReactorID() *string {
	if p == nil {
		return nil
	}
	return p.RequestReactorID
}

func (p *Proxy) GetResponseReactorID() *string {
	if p == nil {
		return nil
	}
	return p.ResponseReactorID
}

func (p *Proxy) GetRequireAuth() *bool {
	if p == nil {
		return nil
	}
	return p.RequireAuth
}

func (p *Proxy) GetRequestTransform() *ProxyTransform {
	if p == nil {
		return nil
	}
	return p.RequestTransform
}

func (p *Proxy) GetResponseTransform() *ProxyTransform {
	if p == nil {
		return nil
	}
	return p.ResponseTransform
}

func (p *Proxy) GetApplicationID() *string {
	if p == nil {
		return nil
	}
	return p.ApplicationID
}

func (p *Proxy) GetConfiguration() map[string]*string {
	if p == nil {
		return nil
	}
	return p.Configuration
}

func (p *Proxy) GetProxyHost() *string {
	if p == nil {
		return nil
	}
	return p.ProxyHost
}

func (p *Proxy) GetTimeout() *int {
	if p == nil {
		return nil
	}
	return p.Timeout
}

func (p *Proxy) GetClientCertificate() *string {
	if p == nil {
		return nil
	}
	return p.ClientCertificate
}

func (p *Proxy) GetCreatedBy() *string {
	if p == nil {
		return nil
	}
	return p.CreatedBy
}

func (p *Proxy) GetCreatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *Proxy) GetModifiedBy() *string {
	if p == nil {
		return nil
	}
	return p.ModifiedBy
}

func (p *Proxy) GetModifiedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.ModifiedAt
}

func (p *Proxy) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *Proxy) UnmarshalJSON(data []byte) error {
	type embed Proxy
	var unmarshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = Proxy(unmarshaler.embed)
	p.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	p.ModifiedAt = unmarshaler.ModifiedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *Proxy) MarshalJSON() ([]byte, error) {
	type embed Proxy
	var marshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed:      embed(*p),
		CreatedAt:  internal.NewOptionalDateTime(p.CreatedAt),
		ModifiedAt: internal.NewOptionalDateTime(p.ModifiedAt),
	}
	return json.Marshal(marshaler)
}

func (p *Proxy) String() string {
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

type ProxyPaginatedList struct {
	Pagination *Pagination `json:"pagination,omitempty" url:"pagination,omitempty"`
	Data       []*Proxy    `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *ProxyPaginatedList) GetPagination() *Pagination {
	if p == nil {
		return nil
	}
	return p.Pagination
}

func (p *ProxyPaginatedList) GetData() []*Proxy {
	if p == nil {
		return nil
	}
	return p.Data
}

func (p *ProxyPaginatedList) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *ProxyPaginatedList) UnmarshalJSON(data []byte) error {
	type unmarshaler ProxyPaginatedList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProxyPaginatedList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *ProxyPaginatedList) String() string {
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

type ProxyTransform struct {
	Type        *string `json:"type,omitempty" url:"type,omitempty"`
	Code        *string `json:"code,omitempty" url:"code,omitempty"`
	Matcher     *string `json:"matcher,omitempty" url:"matcher,omitempty"`
	Expression  *string `json:"expression,omitempty" url:"expression,omitempty"`
	Replacement *string `json:"replacement,omitempty" url:"replacement,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *ProxyTransform) GetType() *string {
	if p == nil {
		return nil
	}
	return p.Type
}

func (p *ProxyTransform) GetCode() *string {
	if p == nil {
		return nil
	}
	return p.Code
}

func (p *ProxyTransform) GetMatcher() *string {
	if p == nil {
		return nil
	}
	return p.Matcher
}

func (p *ProxyTransform) GetExpression() *string {
	if p == nil {
		return nil
	}
	return p.Expression
}

func (p *ProxyTransform) GetReplacement() *string {
	if p == nil {
		return nil
	}
	return p.Replacement
}

func (p *ProxyTransform) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *ProxyTransform) UnmarshalJSON(data []byte) error {
	type unmarshaler ProxyTransform
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProxyTransform(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *ProxyTransform) String() string {
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

type UpdateProxyRequest struct {
	Name              string             `json:"name" url:"-"`
	DestinationURL    string             `json:"destination_url" url:"-"`
	RequestReactorID  *string            `json:"request_reactor_id,omitempty" url:"-"`
	ResponseReactorID *string            `json:"response_reactor_id,omitempty" url:"-"`
	RequestTransform  *ProxyTransform    `json:"request_transform,omitempty" url:"-"`
	ResponseTransform *ProxyTransform    `json:"response_transform,omitempty" url:"-"`
	Application       *Application       `json:"application,omitempty" url:"-"`
	Configuration     map[string]*string `json:"configuration,omitempty" url:"-"`
	RequireAuth       *bool              `json:"require_auth,omitempty" url:"-"`
}
