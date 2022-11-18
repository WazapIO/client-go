/*
WhatsAPI Go

The V2 of WhatsAPI Go

API version: 2.0
Contact: manjit@sponsorbook.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package WhatsAPI

import (
	"encoding/json"
)

// GroupInviteMessagePayload struct for GroupInviteMessagePayload
type GroupInviteMessagePayload struct {
	Caption *string `json:"caption,omitempty"`
	ExpiryMinutes *int32 `json:"expiry_minutes,omitempty"`
	InviteCode *string `json:"invite_code,omitempty"`
	To *string `json:"to,omitempty"`
}

// NewGroupInviteMessagePayload instantiates a new GroupInviteMessagePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupInviteMessagePayload() *GroupInviteMessagePayload {
	this := GroupInviteMessagePayload{}
	return &this
}

// NewGroupInviteMessagePayloadWithDefaults instantiates a new GroupInviteMessagePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupInviteMessagePayloadWithDefaults() *GroupInviteMessagePayload {
	this := GroupInviteMessagePayload{}
	return &this
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *GroupInviteMessagePayload) GetCaption() string {
	if o == nil || isNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInviteMessagePayload) GetCaptionOk() (*string, bool) {
	if o == nil || isNil(o.Caption) {
    return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *GroupInviteMessagePayload) HasCaption() bool {
	if o != nil && !isNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *GroupInviteMessagePayload) SetCaption(v string) {
	o.Caption = &v
}

// GetExpiryMinutes returns the ExpiryMinutes field value if set, zero value otherwise.
func (o *GroupInviteMessagePayload) GetExpiryMinutes() int32 {
	if o == nil || isNil(o.ExpiryMinutes) {
		var ret int32
		return ret
	}
	return *o.ExpiryMinutes
}

// GetExpiryMinutesOk returns a tuple with the ExpiryMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInviteMessagePayload) GetExpiryMinutesOk() (*int32, bool) {
	if o == nil || isNil(o.ExpiryMinutes) {
    return nil, false
	}
	return o.ExpiryMinutes, true
}

// HasExpiryMinutes returns a boolean if a field has been set.
func (o *GroupInviteMessagePayload) HasExpiryMinutes() bool {
	if o != nil && !isNil(o.ExpiryMinutes) {
		return true
	}

	return false
}

// SetExpiryMinutes gets a reference to the given int32 and assigns it to the ExpiryMinutes field.
func (o *GroupInviteMessagePayload) SetExpiryMinutes(v int32) {
	o.ExpiryMinutes = &v
}

// GetInviteCode returns the InviteCode field value if set, zero value otherwise.
func (o *GroupInviteMessagePayload) GetInviteCode() string {
	if o == nil || isNil(o.InviteCode) {
		var ret string
		return ret
	}
	return *o.InviteCode
}

// GetInviteCodeOk returns a tuple with the InviteCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInviteMessagePayload) GetInviteCodeOk() (*string, bool) {
	if o == nil || isNil(o.InviteCode) {
    return nil, false
	}
	return o.InviteCode, true
}

// HasInviteCode returns a boolean if a field has been set.
func (o *GroupInviteMessagePayload) HasInviteCode() bool {
	if o != nil && !isNil(o.InviteCode) {
		return true
	}

	return false
}

// SetInviteCode gets a reference to the given string and assigns it to the InviteCode field.
func (o *GroupInviteMessagePayload) SetInviteCode(v string) {
	o.InviteCode = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *GroupInviteMessagePayload) GetTo() string {
	if o == nil || isNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInviteMessagePayload) GetToOk() (*string, bool) {
	if o == nil || isNil(o.To) {
    return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *GroupInviteMessagePayload) HasTo() bool {
	if o != nil && !isNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *GroupInviteMessagePayload) SetTo(v string) {
	o.To = &v
}

func (o GroupInviteMessagePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	if !isNil(o.ExpiryMinutes) {
		toSerialize["expiry_minutes"] = o.ExpiryMinutes
	}
	if !isNil(o.InviteCode) {
		toSerialize["invite_code"] = o.InviteCode
	}
	if !isNil(o.To) {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableGroupInviteMessagePayload struct {
	value *GroupInviteMessagePayload
	isSet bool
}

func (v NullableGroupInviteMessagePayload) Get() *GroupInviteMessagePayload {
	return v.value
}

func (v *NullableGroupInviteMessagePayload) Set(val *GroupInviteMessagePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupInviteMessagePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupInviteMessagePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupInviteMessagePayload(val *GroupInviteMessagePayload) *NullableGroupInviteMessagePayload {
	return &NullableGroupInviteMessagePayload{value: val, isSet: true}
}

func (v NullableGroupInviteMessagePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupInviteMessagePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

