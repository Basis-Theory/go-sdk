// This file was auto-generated by Fern from our API Definition.

package basistheory

type CreateThreeDsSessionRequest struct {
	Pan           *string            `json:"pan,omitempty" url:"-"`
	TokenID       *string            `json:"token_id,omitempty" url:"-"`
	TokenIntentID *string            `json:"token_intent_id,omitempty" url:"-"`
	Type          *string            `json:"type,omitempty" url:"-"`
	Device        *string            `json:"device,omitempty" url:"-"`
	DeviceInfo    *ThreeDsDeviceInfo `json:"device_info,omitempty" url:"-"`
}
