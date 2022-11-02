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

// StructsListMessagePayload struct for StructsListMessagePayload
type StructsListMessagePayload struct {
	ButtonText *string `json:"button_text,omitempty"`
	Description *string `json:"description,omitempty"`
	MultiSelect *bool `json:"multi_select,omitempty"`
	Sections []StructsListSection `json:"sections"`
	Text *string `json:"text,omitempty"`
	Title *string `json:"title,omitempty"`
	To string `json:"to"`
}

// NewStructsListMessagePayload instantiates a new StructsListMessagePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructsListMessagePayload(sections []StructsListSection, to string) *StructsListMessagePayload {
	this := StructsListMessagePayload{}
	this.Sections = sections
	this.To = to
	return &this
}

// NewStructsListMessagePayloadWithDefaults instantiates a new StructsListMessagePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructsListMessagePayloadWithDefaults() *StructsListMessagePayload {
	this := StructsListMessagePayload{}
	return &this
}

// GetButtonText returns the ButtonText field value if set, zero value otherwise.
func (o *StructsListMessagePayload) GetButtonText() string {
	if o == nil || isNil(o.ButtonText) {
		var ret string
		return ret
	}
	return *o.ButtonText
}

// GetButtonTextOk returns a tuple with the ButtonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsListMessagePayload) GetButtonTextOk() (*string, bool) {
	if o == nil || isNil(o.ButtonText) {
    return nil, false
	}
	return o.ButtonText, true
}

// HasButtonText returns a boolean if a field has been set.
func (o *StructsListMessagePayload) HasButtonText() bool {
	if o != nil && !isNil(o.ButtonText) {
		return true
	}

	return false
}

// SetButtonText gets a reference to the given string and assigns it to the ButtonText field.
func (o *StructsListMessagePayload) SetButtonText(v string) {
	o.ButtonText = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StructsListMessagePayload) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsListMessagePayload) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StructsListMessagePayload) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StructsListMessagePayload) SetDescription(v string) {
	o.Description = &v
}

// GetMultiSelect returns the MultiSelect field value if set, zero value otherwise.
func (o *StructsListMessagePayload) GetMultiSelect() bool {
	if o == nil || isNil(o.MultiSelect) {
		var ret bool
		return ret
	}
	return *o.MultiSelect
}

// GetMultiSelectOk returns a tuple with the MultiSelect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsListMessagePayload) GetMultiSelectOk() (*bool, bool) {
	if o == nil || isNil(o.MultiSelect) {
    return nil, false
	}
	return o.MultiSelect, true
}

// HasMultiSelect returns a boolean if a field has been set.
func (o *StructsListMessagePayload) HasMultiSelect() bool {
	if o != nil && !isNil(o.MultiSelect) {
		return true
	}

	return false
}

// SetMultiSelect gets a reference to the given bool and assigns it to the MultiSelect field.
func (o *StructsListMessagePayload) SetMultiSelect(v bool) {
	o.MultiSelect = &v
}

// GetSections returns the Sections field value
func (o *StructsListMessagePayload) GetSections() []StructsListSection {
	if o == nil {
		var ret []StructsListSection
		return ret
	}

	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value
// and a boolean to check if the value has been set.
func (o *StructsListMessagePayload) GetSectionsOk() ([]StructsListSection, bool) {
	if o == nil {
    return nil, false
	}
	return o.Sections, true
}

// SetSections sets field value
func (o *StructsListMessagePayload) SetSections(v []StructsListSection) {
	o.Sections = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *StructsListMessagePayload) GetText() string {
	if o == nil || isNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsListMessagePayload) GetTextOk() (*string, bool) {
	if o == nil || isNil(o.Text) {
    return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *StructsListMessagePayload) HasText() bool {
	if o != nil && !isNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *StructsListMessagePayload) SetText(v string) {
	o.Text = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *StructsListMessagePayload) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructsListMessagePayload) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *StructsListMessagePayload) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *StructsListMessagePayload) SetTitle(v string) {
	o.Title = &v
}

// GetTo returns the To field value
func (o *StructsListMessagePayload) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *StructsListMessagePayload) GetToOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *StructsListMessagePayload) SetTo(v string) {
	o.To = v
}

func (o StructsListMessagePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ButtonText) {
		toSerialize["button_text"] = o.ButtonText
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.MultiSelect) {
		toSerialize["multi_select"] = o.MultiSelect
	}
	if true {
		toSerialize["sections"] = o.Sections
	}
	if !isNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableStructsListMessagePayload struct {
	value *StructsListMessagePayload
	isSet bool
}

func (v NullableStructsListMessagePayload) Get() *StructsListMessagePayload {
	return v.value
}

func (v *NullableStructsListMessagePayload) Set(val *StructsListMessagePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableStructsListMessagePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableStructsListMessagePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructsListMessagePayload(val *StructsListMessagePayload) *NullableStructsListMessagePayload {
	return &NullableStructsListMessagePayload{value: val, isSet: true}
}

func (v NullableStructsListMessagePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructsListMessagePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


