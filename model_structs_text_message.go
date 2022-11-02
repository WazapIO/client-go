/*
WhatsAPI Go

The V2 of WhatsAPI Go

API version: 2.0
Contact: manjit@sponsorbook.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StructsTextMessage struct for StructsTextMessage
type StructsTextMessage struct {
	Text string `json:"text"`
	To string `json:"to"`
}

// NewStructsTextMessage instantiates a new StructsTextMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructsTextMessage(text string, to string) *StructsTextMessage {
	this := StructsTextMessage{}
	this.Text = text
	this.To = to
	return &this
}

// NewStructsTextMessageWithDefaults instantiates a new StructsTextMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructsTextMessageWithDefaults() *StructsTextMessage {
	this := StructsTextMessage{}
	return &this
}

// GetText returns the Text field value
func (o *StructsTextMessage) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *StructsTextMessage) GetTextOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *StructsTextMessage) SetText(v string) {
	o.Text = v
}

// GetTo returns the To field value
func (o *StructsTextMessage) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *StructsTextMessage) GetToOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *StructsTextMessage) SetTo(v string) {
	o.To = v
}

func (o StructsTextMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["text"] = o.Text
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableStructsTextMessage struct {
	value *StructsTextMessage
	isSet bool
}

func (v NullableStructsTextMessage) Get() *StructsTextMessage {
	return v.value
}

func (v *NullableStructsTextMessage) Set(val *StructsTextMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableStructsTextMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableStructsTextMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructsTextMessage(val *StructsTextMessage) *NullableStructsTextMessage {
	return &NullableStructsTextMessage{value: val, isSet: true}
}

func (v NullableStructsTextMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructsTextMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

