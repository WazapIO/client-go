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

// InstancesInstanceKeySendDocumentPostRequest struct for InstancesInstanceKeySendDocumentPostRequest
type InstancesInstanceKeySendDocumentPostRequest struct {
	// Document file
	File *os.File `json:"file"`
}

// NewInstancesInstanceKeySendDocumentPostRequest instantiates a new InstancesInstanceKeySendDocumentPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstancesInstanceKeySendDocumentPostRequest(file *os.File) *InstancesInstanceKeySendDocumentPostRequest {
	this := InstancesInstanceKeySendDocumentPostRequest{}
	this.File = file
	return &this
}

// NewInstancesInstanceKeySendDocumentPostRequestWithDefaults instantiates a new InstancesInstanceKeySendDocumentPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancesInstanceKeySendDocumentPostRequestWithDefaults() *InstancesInstanceKeySendDocumentPostRequest {
	this := InstancesInstanceKeySendDocumentPostRequest{}
	return &this
}

// GetFile returns the File field value
func (o *InstancesInstanceKeySendDocumentPostRequest) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *InstancesInstanceKeySendDocumentPostRequest) GetFileOk() (**os.File, bool) {
	if o == nil {
    return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *InstancesInstanceKeySendDocumentPostRequest) SetFile(v *os.File) {
	o.File = v
}

func (o InstancesInstanceKeySendDocumentPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableInstancesInstanceKeySendDocumentPostRequest struct {
	value *InstancesInstanceKeySendDocumentPostRequest
	isSet bool
}

func (v NullableInstancesInstanceKeySendDocumentPostRequest) Get() *InstancesInstanceKeySendDocumentPostRequest {
	return v.value
}

func (v *NullableInstancesInstanceKeySendDocumentPostRequest) Set(val *InstancesInstanceKeySendDocumentPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInstancesInstanceKeySendDocumentPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInstancesInstanceKeySendDocumentPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstancesInstanceKeySendDocumentPostRequest(val *InstancesInstanceKeySendDocumentPostRequest) *NullableInstancesInstanceKeySendDocumentPostRequest {
	return &NullableInstancesInstanceKeySendDocumentPostRequest{value: val, isSet: true}
}

func (v NullableInstancesInstanceKeySendDocumentPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstancesInstanceKeySendDocumentPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


