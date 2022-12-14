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
	"os"
)

// UpdateProfilePicRequest struct for UpdateProfilePicRequest
type UpdateProfilePicRequest struct {
	// Image file
	File *os.File `json:"file"`
}

// NewUpdateProfilePicRequest instantiates a new UpdateProfilePicRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfilePicRequest(file *os.File) *UpdateProfilePicRequest {
	this := UpdateProfilePicRequest{}
	this.File = file
	return &this
}

// NewUpdateProfilePicRequestWithDefaults instantiates a new UpdateProfilePicRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfilePicRequestWithDefaults() *UpdateProfilePicRequest {
	this := UpdateProfilePicRequest{}
	return &this
}

// GetFile returns the File field value
func (o *UpdateProfilePicRequest) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *UpdateProfilePicRequest) GetFileOk() (**os.File, bool) {
	if o == nil {
    return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *UpdateProfilePicRequest) SetFile(v *os.File) {
	o.File = v
}

func (o UpdateProfilePicRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProfilePicRequest struct {
	value *UpdateProfilePicRequest
	isSet bool
}

func (v NullableUpdateProfilePicRequest) Get() *UpdateProfilePicRequest {
	return v.value
}

func (v *NullableUpdateProfilePicRequest) Set(val *UpdateProfilePicRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfilePicRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfilePicRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfilePicRequest(val *UpdateProfilePicRequest) *NullableUpdateProfilePicRequest {
	return &NullableUpdateProfilePicRequest{value: val, isSet: true}
}

func (v NullableUpdateProfilePicRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfilePicRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


