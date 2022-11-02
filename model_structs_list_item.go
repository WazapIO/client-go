/*
WhatsAPI Go

The V2 of WhatsAPI Go

API version: 2.0
Contact: manjit@sponsorbook.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package com.whatsapi

import (
	"encoding/json"
)

// StructsListItem struct for StructsListItem
type StructsListItem struct {
	Description *string `json:"description,omitempty"`
	RowId *string `json:"row_id,omitempty"`
	Title string `json:"title"`
}

// NewStructsListItem instantiates a new StructsListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructsListItem(title string) *StructsListItem {
	this := StructsListItem{}
	this.Title = title
	return &this
}

// NewStructsListItemWithDefaults instantiates a new StructsListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructsListItemWithDefaults() *StructsListItem {
	this := StructsListItem{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StructsListItem) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsListItem) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StructsListItem) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StructsListItem) SetDescription(v string) {
	o.Description = &v
}

// GetRowId returns the RowId field value if set, zero value otherwise.
func (o *StructsListItem) GetRowId() string {
	if o == nil || isNil(o.RowId) {
		var ret string
		return ret
	}
	return *o.RowId
}

// GetRowIdOk returns a tuple with the RowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsListItem) GetRowIdOk() (*string, bool) {
	if o == nil || isNil(o.RowId) {
    return nil, false
	}
	return o.RowId, true
}

// HasRowId returns a boolean if a field has been set.
func (o *StructsListItem) HasRowId() bool {
	if o != nil && !isNil(o.RowId) {
		return true
	}

	return false
}

// SetRowId gets a reference to the given string and assigns it to the RowId field.
func (o *StructsListItem) SetRowId(v string) {
	o.RowId = &v
}

// GetTitle returns the Title field value
func (o *StructsListItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *StructsListItem) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *StructsListItem) SetTitle(v string) {
	o.Title = v
}

func (o StructsListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.RowId) {
		toSerialize["row_id"] = o.RowId
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableStructsListItem struct {
	value *StructsListItem
	isSet bool
}

func (v NullableStructsListItem) Get() *StructsListItem {
	return v.value
}

func (v *NullableStructsListItem) Set(val *StructsListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStructsListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStructsListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructsListItem(val *StructsListItem) *NullableStructsListItem {
	return &NullableStructsListItem{value: val, isSet: true}
}

func (v NullableStructsListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructsListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


