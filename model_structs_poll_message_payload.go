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

// StructsPollMessagePayload struct for StructsPollMessagePayload
type StructsPollMessagePayload struct {
	Options []string `json:"options"`
	SelectableOptionsCount int32 `json:"selectable_options_count"`
	Title string `json:"title"`
	To string `json:"to"`
}

// NewStructsPollMessagePayload instantiates a new StructsPollMessagePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructsPollMessagePayload(options []string, selectableOptionsCount int32, title string, to string) *StructsPollMessagePayload {
	this := StructsPollMessagePayload{}
	this.Options = options
	this.SelectableOptionsCount = selectableOptionsCount
	this.Title = title
	this.To = to
	return &this
}

// NewStructsPollMessagePayloadWithDefaults instantiates a new StructsPollMessagePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructsPollMessagePayloadWithDefaults() *StructsPollMessagePayload {
	this := StructsPollMessagePayload{}
	return &this
}

// GetOptions returns the Options field value
func (o *StructsPollMessagePayload) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *StructsPollMessagePayload) GetOptionsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *StructsPollMessagePayload) SetOptions(v []string) {
	o.Options = v
}

// GetSelectableOptionsCount returns the SelectableOptionsCount field value
func (o *StructsPollMessagePayload) GetSelectableOptionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SelectableOptionsCount
}

// GetSelectableOptionsCountOk returns a tuple with the SelectableOptionsCount field value
// and a boolean to check if the value has been set.
func (o *StructsPollMessagePayload) GetSelectableOptionsCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SelectableOptionsCount, true
}

// SetSelectableOptionsCount sets field value
func (o *StructsPollMessagePayload) SetSelectableOptionsCount(v int32) {
	o.SelectableOptionsCount = v
}

// GetTitle returns the Title field value
func (o *StructsPollMessagePayload) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *StructsPollMessagePayload) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *StructsPollMessagePayload) SetTitle(v string) {
	o.Title = v
}

// GetTo returns the To field value
func (o *StructsPollMessagePayload) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *StructsPollMessagePayload) GetToOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *StructsPollMessagePayload) SetTo(v string) {
	o.To = v
}

func (o StructsPollMessagePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["options"] = o.Options
	}
	if true {
		toSerialize["selectable_options_count"] = o.SelectableOptionsCount
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableStructsPollMessagePayload struct {
	value *StructsPollMessagePayload
	isSet bool
}

func (v NullableStructsPollMessagePayload) Get() *StructsPollMessagePayload {
	return v.value
}

func (v *NullableStructsPollMessagePayload) Set(val *StructsPollMessagePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableStructsPollMessagePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableStructsPollMessagePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructsPollMessagePayload(val *StructsPollMessagePayload) *NullableStructsPollMessagePayload {
	return &NullableStructsPollMessagePayload{value: val, isSet: true}
}

func (v NullableStructsPollMessagePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructsPollMessagePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


