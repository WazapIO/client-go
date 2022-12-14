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

// GroupUpdateDescriptionPayload struct for GroupUpdateDescriptionPayload
type GroupUpdateDescriptionPayload struct {
	Description *string `json:"description,omitempty"`
}

// NewGroupUpdateDescriptionPayload instantiates a new GroupUpdateDescriptionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUpdateDescriptionPayload() *GroupUpdateDescriptionPayload {
	this := GroupUpdateDescriptionPayload{}
	return &this
}

// NewGroupUpdateDescriptionPayloadWithDefaults instantiates a new GroupUpdateDescriptionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUpdateDescriptionPayloadWithDefaults() *GroupUpdateDescriptionPayload {
	this := GroupUpdateDescriptionPayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupUpdateDescriptionPayload) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdateDescriptionPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupUpdateDescriptionPayload) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupUpdateDescriptionPayload) SetDescription(v string) {
	o.Description = &v
}

func (o GroupUpdateDescriptionPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableGroupUpdateDescriptionPayload struct {
	value *GroupUpdateDescriptionPayload
	isSet bool
}

func (v NullableGroupUpdateDescriptionPayload) Get() *GroupUpdateDescriptionPayload {
	return v.value
}

func (v *NullableGroupUpdateDescriptionPayload) Set(val *GroupUpdateDescriptionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdateDescriptionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdateDescriptionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdateDescriptionPayload(val *GroupUpdateDescriptionPayload) *NullableGroupUpdateDescriptionPayload {
	return &NullableGroupUpdateDescriptionPayload{value: val, isSet: true}
}

func (v NullableGroupUpdateDescriptionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdateDescriptionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


