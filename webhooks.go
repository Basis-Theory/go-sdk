// Code generated by Fern. DO NOT EDIT.

package basistheory

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Basis-Theory/go-sdk/v2/internal"
	time "time"
)

type CreateWebhookRequest struct {
	// The name of the webhook
	Name string `json:"name" url:"-"`
	// The URL to which the webhook will send events
	URL string `json:"url" url:"-"`
	// The email address to use for management notification events. Ie: webhook disabled
	NotifyEmail *string `json:"notify_email,omitempty" url:"-"`
	// An array of event types that the webhook will listen for
	Events []string `json:"events,omitempty" url:"-"`
}

type Webhook struct {
	ID       string        `json:"id" url:"id"`
	TenantID string        `json:"tenant_id" url:"tenant_id"`
	Status   WebhookStatus `json:"status" url:"status"`
	Name     string        `json:"name" url:"name"`
	URL      string        `json:"url" url:"url"`
	// The email address to use for management notification events. Ie: webhook disabled
	NotifyEmail *string    `json:"notify_email,omitempty" url:"notify_email,omitempty"`
	Events      []string   `json:"events" url:"events"`
	CreatedBy   string     `json:"created_by" url:"created_by"`
	CreatedAt   time.Time  `json:"created_at" url:"created_at"`
	ModifiedBy  *string    `json:"modified_by,omitempty" url:"modified_by,omitempty"`
	ModifiedAt  *time.Time `json:"modified_at,omitempty" url:"modified_at,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (w *Webhook) GetID() string {
	if w == nil {
		return ""
	}
	return w.ID
}

func (w *Webhook) GetTenantID() string {
	if w == nil {
		return ""
	}
	return w.TenantID
}

func (w *Webhook) GetStatus() WebhookStatus {
	if w == nil {
		return ""
	}
	return w.Status
}

func (w *Webhook) GetName() string {
	if w == nil {
		return ""
	}
	return w.Name
}

func (w *Webhook) GetURL() string {
	if w == nil {
		return ""
	}
	return w.URL
}

func (w *Webhook) GetNotifyEmail() *string {
	if w == nil {
		return nil
	}
	return w.NotifyEmail
}

func (w *Webhook) GetEvents() []string {
	if w == nil {
		return nil
	}
	return w.Events
}

func (w *Webhook) GetCreatedBy() string {
	if w == nil {
		return ""
	}
	return w.CreatedBy
}

func (w *Webhook) GetCreatedAt() time.Time {
	if w == nil {
		return time.Time{}
	}
	return w.CreatedAt
}

func (w *Webhook) GetModifiedBy() *string {
	if w == nil {
		return nil
	}
	return w.ModifiedBy
}

func (w *Webhook) GetModifiedAt() *time.Time {
	if w == nil {
		return nil
	}
	return w.ModifiedAt
}

func (w *Webhook) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *Webhook) UnmarshalJSON(data []byte) error {
	type embed Webhook
	var unmarshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed: embed(*w),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*w = Webhook(unmarshaler.embed)
	w.CreatedAt = unmarshaler.CreatedAt.Time()
	w.ModifiedAt = unmarshaler.ModifiedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	w.rawJSON = json.RawMessage(data)
	return nil
}

func (w *Webhook) MarshalJSON() ([]byte, error) {
	type embed Webhook
	var marshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed:      embed(*w),
		CreatedAt:  internal.NewDateTime(w.CreatedAt),
		ModifiedAt: internal.NewOptionalDateTime(w.ModifiedAt),
	}
	return json.Marshal(marshaler)
}

func (w *Webhook) String() string {
	if len(w.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(w.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebhookList struct {
	Pagination *WebhookListPagination `json:"pagination" url:"pagination"`
	Data       []*Webhook             `json:"data" url:"data"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (w *WebhookList) GetPagination() *WebhookListPagination {
	if w == nil {
		return nil
	}
	return w.Pagination
}

func (w *WebhookList) GetData() []*Webhook {
	if w == nil {
		return nil
	}
	return w.Data
}

func (w *WebhookList) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *WebhookList) UnmarshalJSON(data []byte) error {
	type unmarshaler WebhookList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WebhookList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	w.rawJSON = json.RawMessage(data)
	return nil
}

func (w *WebhookList) String() string {
	if len(w.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(w.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebhookListPagination struct {
	PageSize *int    `json:"page_size,omitempty" url:"page_size,omitempty"`
	Next     *string `json:"next,omitempty" url:"next,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (w *WebhookListPagination) GetPageSize() *int {
	if w == nil {
		return nil
	}
	return w.PageSize
}

func (w *WebhookListPagination) GetNext() *string {
	if w == nil {
		return nil
	}
	return w.Next
}

func (w *WebhookListPagination) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *WebhookListPagination) UnmarshalJSON(data []byte) error {
	type unmarshaler WebhookListPagination
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WebhookListPagination(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	w.rawJSON = json.RawMessage(data)
	return nil
}

func (w *WebhookListPagination) String() string {
	if len(w.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(w.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebhookStatus string

const (
	WebhookStatusEnabled  WebhookStatus = "enabled"
	WebhookStatusDisabled WebhookStatus = "disabled"
)

func NewWebhookStatusFromString(s string) (WebhookStatus, error) {
	switch s {
	case "enabled":
		return WebhookStatusEnabled, nil
	case "disabled":
		return WebhookStatusDisabled, nil
	}
	var t WebhookStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (w WebhookStatus) Ptr() *WebhookStatus {
	return &w
}

type UpdateWebhookRequest struct {
	// The name of the webhook
	Name string `json:"name" url:"-"`
	// The URL to which the webhook will send events
	URL string `json:"url" url:"-"`
	// The email address to use for management notification events. Ie: webhook disabled
	NotifyEmail *string `json:"notify_email,omitempty" url:"-"`
	// An array of event types that the webhook will listen for
	Events []string `json:"events,omitempty" url:"-"`
}
