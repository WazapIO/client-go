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
	"os"
)

// InstancesInstanceKeySendImagePostRequest struct for InstancesInstanceKeySendImagePostRequest
type InstancesInstanceKeySendImagePostRequest struct {
	// Image file
	File *os.File `json:"file"`
}

// NewInstancesInstanceKeySendImagePostRequest instantiates a new InstancesInstanceKeySendImagePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstancesInstanceKeySendImagePostRequest(file *os.File) *InstancesInstanceKeySendImagePostRequest {
	this := InstancesInstanceKeySendImagePostRequest{}
	this.File = file
	return &this
}

// NewInstancesInstanceKeySendImagePostRequestWithDefaults instantiates a new InstancesInstanceKeySendImagePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancesInstanceKeySendImagePostRequestWithDefaults() *InstancesInstanceKeySendImagePostRequest {
	this := InstancesInstanceKeySendImagePostRequest{}
	return &this
}

// GetFile returns the File field value
func (o *InstancesInstanceKeySendImagePostRequest) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *InstancesInstanceKeySendImagePostRequest) GetFileOk() (**os.File, bool) {
	if o == nil {
    return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *InstancesInstanceKeySendImagePostRequest) SetFile(v *os.File) {
	o.File = v
}

func (o InstancesInstanceKeySendImagePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableInstancesInstanceKeySendImagePostRequest struct {
	value *InstancesInstanceKeySendImagePostRequest
	isSet bool
}

func (v NullableInstancesInstanceKeySendImagePostRequest) Get() *InstancesInstanceKeySendImagePostRequest {
	return v.value
}

func (v *NullableInstancesInstanceKeySendImagePostRequest) Set(val *InstancesInstanceKeySendImagePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInstancesInstanceKeySendImagePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInstancesInstanceKeySendImagePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstancesInstanceKeySendImagePostRequest(val *InstancesInstanceKeySendImagePostRequest) *NullableInstancesInstanceKeySendImagePostRequest {
	return &NullableInstancesInstanceKeySendImagePostRequest{value: val, isSet: true}
}

func (v NullableInstancesInstanceKeySendImagePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstancesInstanceKeySendImagePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


