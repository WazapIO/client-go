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

// SendMediaPayload struct for SendMediaPayload
type SendMediaPayload struct {
	Caption *string `json:"caption,omitempty"`
	Filename *string `json:"filename,omitempty"`
	MediaData FileUpload `json:"media_data"`
	MediaType string `json:"media_type"`
	To string `json:"to"`
}

// NewSendMediaPayload instantiates a new SendMediaPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendMediaPayload(mediaData FileUpload, mediaType string, to string) *SendMediaPayload {
	this := SendMediaPayload{}
	this.MediaData = mediaData
	this.MediaType = mediaType
	this.To = to
	return &this
}

// NewSendMediaPayloadWithDefaults instantiates a new SendMediaPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendMediaPayloadWithDefaults() *SendMediaPayload {
	this := SendMediaPayload{}
	return &this
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *SendMediaPayload) GetCaption() string {
	if o == nil || isNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMediaPayload) GetCaptionOk() (*string, bool) {
	if o == nil || isNil(o.Caption) {
    return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *SendMediaPayload) HasCaption() bool {
	if o != nil && !isNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *SendMediaPayload) SetCaption(v string) {
	o.Caption = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *SendMediaPayload) GetFilename() string {
	if o == nil || isNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMediaPayload) GetFilenameOk() (*string, bool) {
	if o == nil || isNil(o.Filename) {
    return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *SendMediaPayload) HasFilename() bool {
	if o != nil && !isNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *SendMediaPayload) SetFilename(v string) {
	o.Filename = &v
}

// GetMediaData returns the MediaData field value
func (o *SendMediaPayload) GetMediaData() FileUpload {
	if o == nil {
		var ret FileUpload
		return ret
	}

	return o.MediaData
}

// GetMediaDataOk returns a tuple with the MediaData field value
// and a boolean to check if the value has been set.
func (o *SendMediaPayload) GetMediaDataOk() (*FileUpload, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MediaData, true
}

// SetMediaData sets field value
func (o *SendMediaPayload) SetMediaData(v FileUpload) {
	o.MediaData = v
}

// GetMediaType returns the MediaType field value
func (o *SendMediaPayload) GetMediaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value
// and a boolean to check if the value has been set.
func (o *SendMediaPayload) GetMediaTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MediaType, true
}

// SetMediaType sets field value
func (o *SendMediaPayload) SetMediaType(v string) {
	o.MediaType = v
}

// GetTo returns the To field value
func (o *SendMediaPayload) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *SendMediaPayload) GetToOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *SendMediaPayload) SetTo(v string) {
	o.To = v
}

func (o SendMediaPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	if !isNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if true {
		toSerialize["media_data"] = o.MediaData
	}
	if true {
		toSerialize["media_type"] = o.MediaType
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableSendMediaPayload struct {
	value *SendMediaPayload
	isSet bool
}

func (v NullableSendMediaPayload) Get() *SendMediaPayload {
	return v.value
}

func (v *NullableSendMediaPayload) Set(val *SendMediaPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMediaPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMediaPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMediaPayload(val *SendMediaPayload) *NullableSendMediaPayload {
	return &NullableSendMediaPayload{value: val, isSet: true}
}

func (v NullableSendMediaPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMediaPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

