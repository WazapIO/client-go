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

// WebhookPayload struct for WebhookPayload
type WebhookPayload struct {
	WebhookUrl *string `json:"webhook_url,omitempty"`
}

// NewWebhookPayload instantiates a new WebhookPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookPayload() *WebhookPayload {
	this := WebhookPayload{}
	return &this
}

// NewWebhookPayloadWithDefaults instantiates a new WebhookPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookPayloadWithDefaults() *WebhookPayload {
	this := WebhookPayload{}
	return &this
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *WebhookPayload) GetWebhookUrl() string {
	if o == nil || isNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPayload) GetWebhookUrlOk() (*string, bool) {
	if o == nil || isNil(o.WebhookUrl) {
    return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *WebhookPayload) HasWebhookUrl() bool {
	if o != nil && !isNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *WebhookPayload) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o WebhookPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WebhookUrl) {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookPayload struct {
	value *WebhookPayload
	isSet bool
}

func (v NullableWebhookPayload) Get() *WebhookPayload {
	return v.value
}

func (v *NullableWebhookPayload) Set(val *WebhookPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookPayload(val *WebhookPayload) *NullableWebhookPayload {
	return &NullableWebhookPayload{value: val, isSet: true}
}

func (v NullableWebhookPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

