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

// StructsListSection struct for StructsListSection
type StructsListSection struct {
	Rows []StructsListItem `json:"rows"`
	Title string `json:"title"`
}

// NewStructsListSection instantiates a new StructsListSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructsListSection(rows []StructsListItem, title string) *StructsListSection {
	this := StructsListSection{}
	this.Rows = rows
	this.Title = title
	return &this
}

// NewStructsListSectionWithDefaults instantiates a new StructsListSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructsListSectionWithDefaults() *StructsListSection {
	this := StructsListSection{}
	return &this
}

// GetRows returns the Rows field value
func (o *StructsListSection) GetRows() []StructsListItem {
	if o == nil {
		var ret []StructsListItem
		return ret
	}

	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *StructsListSection) GetRowsOk() ([]StructsListItem, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rows, true
}

// SetRows sets field value
func (o *StructsListSection) SetRows(v []StructsListItem) {
	o.Rows = v
}

// GetTitle returns the Title field value
func (o *StructsListSection) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *StructsListSection) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *StructsListSection) SetTitle(v string) {
	o.Title = v
}

func (o StructsListSection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rows"] = o.Rows
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableStructsListSection struct {
	value *StructsListSection
	isSet bool
}

func (v NullableStructsListSection) Get() *StructsListSection {
	return v.value
}

func (v *NullableStructsListSection) Set(val *StructsListSection) {
	v.value = val
	v.isSet = true
}

func (v NullableStructsListSection) IsSet() bool {
	return v.isSet
}

func (v *NullableStructsListSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructsListSection(val *StructsListSection) *NullableStructsListSection {
	return &NullableStructsListSection{value: val, isSet: true}
}

func (v NullableStructsListSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructsListSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


