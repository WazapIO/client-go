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

// StructsLocationMessagePayload struct for StructsLocationMessagePayload
type StructsLocationMessagePayload struct {
	Location StructsLocationMessagePayloadLocation `json:"location"`
	To string `json:"to"`
	Url *string `json:"url,omitempty"`
}

// NewStructsLocationMessagePayload instantiates a new StructsLocationMessagePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructsLocationMessagePayload(location StructsLocationMessagePayloadLocation, to string) *StructsLocationMessagePayload {
	this := StructsLocationMessagePayload{}
	this.Location = location
	this.To = to
	return &this
}

// NewStructsLocationMessagePayloadWithDefaults instantiates a new StructsLocationMessagePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructsLocationMessagePayloadWithDefaults() *StructsLocationMessagePayload {
	this := StructsLocationMessagePayload{}
	return &this
}

// GetLocation returns the Location field value
func (o *StructsLocationMessagePayload) GetLocation() StructsLocationMessagePayloadLocation {
	if o == nil {
		var ret StructsLocationMessagePayloadLocation
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *StructsLocationMessagePayload) GetLocationOk() (*StructsLocationMessagePayloadLocation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *StructsLocationMessagePayload) SetLocation(v StructsLocationMessagePayloadLocation) {
	o.Location = v
}

// GetTo returns the To field value
func (o *StructsLocationMessagePayload) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *StructsLocationMessagePayload) GetToOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *StructsLocationMessagePayload) SetTo(v string) {
	o.To = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *StructsLocationMessagePayload) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsLocationMessagePayload) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *StructsLocationMessagePayload) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *StructsLocationMessagePayload) SetUrl(v string) {
	o.Url = &v
}

func (o StructsLocationMessagePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["to"] = o.To
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableStructsLocationMessagePayload struct {
	value *StructsLocationMessagePayload
	isSet bool
}

func (v NullableStructsLocationMessagePayload) Get() *StructsLocationMessagePayload {
	return v.value
}

func (v *NullableStructsLocationMessagePayload) Set(val *StructsLocationMessagePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableStructsLocationMessagePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableStructsLocationMessagePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructsLocationMessagePayload(val *StructsLocationMessagePayload) *NullableStructsLocationMessagePayload {
	return &NullableStructsLocationMessagePayload{value: val, isSet: true}
}

func (v NullableStructsLocationMessagePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructsLocationMessagePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


