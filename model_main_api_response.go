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

// MainAPIResponse struct for MainAPIResponse
type MainAPIResponse struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Error *bool `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
	Status *int32 `json:"status,omitempty"`
}

// NewMainAPIResponse instantiates a new MainAPIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMainAPIResponse() *MainAPIResponse {
	this := MainAPIResponse{}
	return &this
}

// NewMainAPIResponseWithDefaults instantiates a new MainAPIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMainAPIResponseWithDefaults() *MainAPIResponse {
	this := MainAPIResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MainAPIResponse) GetData() map[string]interface{} {
	if o == nil || isNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainAPIResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Data) {
    return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MainAPIResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *MainAPIResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MainAPIResponse) GetError() bool {
	if o == nil || isNil(o.Error) {
		var ret bool
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainAPIResponse) GetErrorOk() (*bool, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MainAPIResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given bool and assigns it to the Error field.
func (o *MainAPIResponse) SetError(v bool) {
	o.Error = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MainAPIResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainAPIResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MainAPIResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MainAPIResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MainAPIResponse) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainAPIResponse) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MainAPIResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *MainAPIResponse) SetStatus(v int32) {
	o.Status = &v
}

func (o MainAPIResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMainAPIResponse struct {
	value *MainAPIResponse
	isSet bool
}

func (v NullableMainAPIResponse) Get() *MainAPIResponse {
	return v.value
}

func (v *NullableMainAPIResponse) Set(val *MainAPIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMainAPIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMainAPIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMainAPIResponse(val *MainAPIResponse) *NullableMainAPIResponse {
	return &NullableMainAPIResponse{value: val, isSet: true}
}

func (v NullableMainAPIResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMainAPIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


