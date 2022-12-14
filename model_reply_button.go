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

// ReplyButton struct for ReplyButton
type ReplyButton struct {
	ButtonId *string `json:"button_id,omitempty"`
	ButtonText *string `json:"button_text,omitempty"`
}

// NewReplyButton instantiates a new ReplyButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplyButton() *ReplyButton {
	this := ReplyButton{}
	return &this
}

// NewReplyButtonWithDefaults instantiates a new ReplyButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplyButtonWithDefaults() *ReplyButton {
	this := ReplyButton{}
	return &this
}

// GetButtonId returns the ButtonId field value if set, zero value otherwise.
func (o *ReplyButton) GetButtonId() string {
	if o == nil || isNil(o.ButtonId) {
		var ret string
		return ret
	}
	return *o.ButtonId
}

// GetButtonIdOk returns a tuple with the ButtonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyButton) GetButtonIdOk() (*string, bool) {
	if o == nil || isNil(o.ButtonId) {
    return nil, false
	}
	return o.ButtonId, true
}

// HasButtonId returns a boolean if a field has been set.
func (o *ReplyButton) HasButtonId() bool {
	if o != nil && !isNil(o.ButtonId) {
		return true
	}

	return false
}

// SetButtonId gets a reference to the given string and assigns it to the ButtonId field.
func (o *ReplyButton) SetButtonId(v string) {
	o.ButtonId = &v
}

// GetButtonText returns the ButtonText field value if set, zero value otherwise.
func (o *ReplyButton) GetButtonText() string {
	if o == nil || isNil(o.ButtonText) {
		var ret string
		return ret
	}
	return *o.ButtonText
}

// GetButtonTextOk returns a tuple with the ButtonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyButton) GetButtonTextOk() (*string, bool) {
	if o == nil || isNil(o.ButtonText) {
    return nil, false
	}
	return o.ButtonText, true
}

// HasButtonText returns a boolean if a field has been set.
func (o *ReplyButton) HasButtonText() bool {
	if o != nil && !isNil(o.ButtonText) {
		return true
	}

	return false
}

// SetButtonText gets a reference to the given string and assigns it to the ButtonText field.
func (o *ReplyButton) SetButtonText(v string) {
	o.ButtonText = &v
}

func (o ReplyButton) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ButtonId) {
		toSerialize["button_id"] = o.ButtonId
	}
	if !isNil(o.ButtonText) {
		toSerialize["button_text"] = o.ButtonText
	}
	return json.Marshal(toSerialize)
}

type NullableReplyButton struct {
	value *ReplyButton
	isSet bool
}

func (v NullableReplyButton) Get() *ReplyButton {
	return v.value
}

func (v *NullableReplyButton) Set(val *ReplyButton) {
	v.value = val
	v.isSet = true
}

func (v NullableReplyButton) IsSet() bool {
	return v.isSet
}

func (v *NullableReplyButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplyButton(val *ReplyButton) *NullableReplyButton {
	return &NullableReplyButton{value: val, isSet: true}
}

func (v NullableReplyButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplyButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


