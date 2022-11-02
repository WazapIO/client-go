/*
WhatsAPI Go

The V2 of WhatsAPI Go

API version: 2.0
Contact: manjit@sponsorbook.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package whatsapi

import (
	"encoding/json"
)

// StructsReplyButton struct for StructsReplyButton
type StructsReplyButton struct {
	ButtonId *string `json:"button_id,omitempty"`
	ButtonText *string `json:"button_text,omitempty"`
}

// NewStructsReplyButton instantiates a new StructsReplyButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructsReplyButton() *StructsReplyButton {
	this := StructsReplyButton{}
	return &this
}

// NewStructsReplyButtonWithDefaults instantiates a new StructsReplyButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructsReplyButtonWithDefaults() *StructsReplyButton {
	this := StructsReplyButton{}
	return &this
}

// GetButtonId returns the ButtonId field value if set, zero value otherwise.
func (o *StructsReplyButton) GetButtonId() string {
	if o == nil || isNil(o.ButtonId) {
		var ret string
		return ret
	}
	return *o.ButtonId
}

// GetButtonIdOk returns a tuple with the ButtonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsReplyButton) GetButtonIdOk() (*string, bool) {
	if o == nil || isNil(o.ButtonId) {
    return nil, false
	}
	return o.ButtonId, true
}

// HasButtonId returns a boolean if a field has been set.
func (o *StructsReplyButton) HasButtonId() bool {
	if o != nil && !isNil(o.ButtonId) {
		return true
	}

	return false
}

// SetButtonId gets a reference to the given string and assigns it to the ButtonId field.
func (o *StructsReplyButton) SetButtonId(v string) {
	o.ButtonId = &v
}

// GetButtonText returns the ButtonText field value if set, zero value otherwise.
func (o *StructsReplyButton) GetButtonText() string {
	if o == nil || isNil(o.ButtonText) {
		var ret string
		return ret
	}
	return *o.ButtonText
}

// GetButtonTextOk returns a tuple with the ButtonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsReplyButton) GetButtonTextOk() (*string, bool) {
	if o == nil || isNil(o.ButtonText) {
    return nil, false
	}
	return o.ButtonText, true
}

// HasButtonText returns a boolean if a field has been set.
func (o *StructsReplyButton) HasButtonText() bool {
	if o != nil && !isNil(o.ButtonText) {
		return true
	}

	return false
}

// SetButtonText gets a reference to the given string and assigns it to the ButtonText field.
func (o *StructsReplyButton) SetButtonText(v string) {
	o.ButtonText = &v
}

func (o StructsReplyButton) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ButtonId) {
		toSerialize["button_id"] = o.ButtonId
	}
	if !isNil(o.ButtonText) {
		toSerialize["button_text"] = o.ButtonText
	}
	return json.Marshal(toSerialize)
}

type NullableStructsReplyButton struct {
	value *StructsReplyButton
	isSet bool
}

func (v NullableStructsReplyButton) Get() *StructsReplyButton {
	return v.value
}

func (v *NullableStructsReplyButton) Set(val *StructsReplyButton) {
	v.value = val
	v.isSet = true
}

func (v NullableStructsReplyButton) IsSet() bool {
	return v.isSet
}

func (v *NullableStructsReplyButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructsReplyButton(val *StructsReplyButton) *NullableStructsReplyButton {
	return &NullableStructsReplyButton{value: val, isSet: true}
}

func (v NullableStructsReplyButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructsReplyButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


