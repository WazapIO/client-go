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

// GroupCreatePayload struct for GroupCreatePayload
type GroupCreatePayload struct {
	GroupName *string `json:"group_name,omitempty"`
	Participants []string `json:"participants,omitempty"`
}

// NewGroupCreatePayload instantiates a new GroupCreatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCreatePayload() *GroupCreatePayload {
	this := GroupCreatePayload{}
	return &this
}

// NewGroupCreatePayloadWithDefaults instantiates a new GroupCreatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCreatePayloadWithDefaults() *GroupCreatePayload {
	this := GroupCreatePayload{}
	return &this
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *GroupCreatePayload) GetGroupName() string {
	if o == nil || isNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreatePayload) GetGroupNameOk() (*string, bool) {
	if o == nil || isNil(o.GroupName) {
    return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *GroupCreatePayload) HasGroupName() bool {
	if o != nil && !isNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *GroupCreatePayload) SetGroupName(v string) {
	o.GroupName = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *GroupCreatePayload) GetParticipants() []string {
	if o == nil || isNil(o.Participants) {
		var ret []string
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreatePayload) GetParticipantsOk() ([]string, bool) {
	if o == nil || isNil(o.Participants) {
    return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *GroupCreatePayload) HasParticipants() bool {
	if o != nil && !isNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []string and assigns it to the Participants field.
func (o *GroupCreatePayload) SetParticipants(v []string) {
	o.Participants = v
}

func (o GroupCreatePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !isNil(o.Participants) {
		toSerialize["participants"] = o.Participants
	}
	return json.Marshal(toSerialize)
}

type NullableGroupCreatePayload struct {
	value *GroupCreatePayload
	isSet bool
}

func (v NullableGroupCreatePayload) Get() *GroupCreatePayload {
	return v.value
}

func (v *NullableGroupCreatePayload) Set(val *GroupCreatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCreatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCreatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCreatePayload(val *GroupCreatePayload) *NullableGroupCreatePayload {
	return &NullableGroupCreatePayload{value: val, isSet: true}
}

func (v NullableGroupCreatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCreatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


