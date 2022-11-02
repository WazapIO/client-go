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

// StructsFileUpload struct for StructsFileUpload
type StructsFileUpload struct {
	DirectPath string `json:"direct_path"`
	FileEncSha256 []int32 `json:"file_enc_sha256"`
	FileLength int32 `json:"file_length"`
	FileSha256 []int32 `json:"file_sha256"`
	MediaKey []int32 `json:"media_key"`
	MimeType string `json:"mime_type"`
	Url string `json:"url"`
}

// NewStructsFileUpload instantiates a new StructsFileUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructsFileUpload(directPath string, fileEncSha256 []int32, fileLength int32, fileSha256 []int32, mediaKey []int32, mimeType string, url string) *StructsFileUpload {
	this := StructsFileUpload{}
	this.DirectPath = directPath
	this.FileEncSha256 = fileEncSha256
	this.FileLength = fileLength
	this.FileSha256 = fileSha256
	this.MediaKey = mediaKey
	this.MimeType = mimeType
	this.Url = url
	return &this
}

// NewStructsFileUploadWithDefaults instantiates a new StructsFileUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructsFileUploadWithDefaults() *StructsFileUpload {
	this := StructsFileUpload{}
	return &this
}

// GetDirectPath returns the DirectPath field value
func (o *StructsFileUpload) GetDirectPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DirectPath
}

// GetDirectPathOk returns a tuple with the DirectPath field value
// and a boolean to check if the value has been set.
func (o *StructsFileUpload) GetDirectPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DirectPath, true
}

// SetDirectPath sets field value
func (o *StructsFileUpload) SetDirectPath(v string) {
	o.DirectPath = v
}

// GetFileEncSha256 returns the FileEncSha256 field value
func (o *StructsFileUpload) GetFileEncSha256() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.FileEncSha256
}

// GetFileEncSha256Ok returns a tuple with the FileEncSha256 field value
// and a boolean to check if the value has been set.
func (o *StructsFileUpload) GetFileEncSha256Ok() ([]int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.FileEncSha256, true
}

// SetFileEncSha256 sets field value
func (o *StructsFileUpload) SetFileEncSha256(v []int32) {
	o.FileEncSha256 = v
}

// GetFileLength returns the FileLength field value
func (o *StructsFileUpload) GetFileLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileLength
}

// GetFileLengthOk returns a tuple with the FileLength field value
// and a boolean to check if the value has been set.
func (o *StructsFileUpload) GetFileLengthOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FileLength, true
}

// SetFileLength sets field value
func (o *StructsFileUpload) SetFileLength(v int32) {
	o.FileLength = v
}

// GetFileSha256 returns the FileSha256 field value
func (o *StructsFileUpload) GetFileSha256() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.FileSha256
}

// GetFileSha256Ok returns a tuple with the FileSha256 field value
// and a boolean to check if the value has been set.
func (o *StructsFileUpload) GetFileSha256Ok() ([]int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.FileSha256, true
}

// SetFileSha256 sets field value
func (o *StructsFileUpload) SetFileSha256(v []int32) {
	o.FileSha256 = v
}

// GetMediaKey returns the MediaKey field value
func (o *StructsFileUpload) GetMediaKey() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.MediaKey
}

// GetMediaKeyOk returns a tuple with the MediaKey field value
// and a boolean to check if the value has been set.
func (o *StructsFileUpload) GetMediaKeyOk() ([]int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MediaKey, true
}

// SetMediaKey sets field value
func (o *StructsFileUpload) SetMediaKey(v []int32) {
	o.MediaKey = v
}

// GetMimeType returns the MimeType field value
func (o *StructsFileUpload) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *StructsFileUpload) GetMimeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *StructsFileUpload) SetMimeType(v string) {
	o.MimeType = v
}

// GetUrl returns the Url field value
func (o *StructsFileUpload) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *StructsFileUpload) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *StructsFileUpload) SetUrl(v string) {
	o.Url = v
}

func (o StructsFileUpload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["direct_path"] = o.DirectPath
	}
	if true {
		toSerialize["file_enc_sha256"] = o.FileEncSha256
	}
	if true {
		toSerialize["file_length"] = o.FileLength
	}
	if true {
		toSerialize["file_sha256"] = o.FileSha256
	}
	if true {
		toSerialize["media_key"] = o.MediaKey
	}
	if true {
		toSerialize["mime_type"] = o.MimeType
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableStructsFileUpload struct {
	value *StructsFileUpload
	isSet bool
}

func (v NullableStructsFileUpload) Get() *StructsFileUpload {
	return v.value
}

func (v *NullableStructsFileUpload) Set(val *StructsFileUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableStructsFileUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableStructsFileUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructsFileUpload(val *StructsFileUpload) *NullableStructsFileUpload {
	return &NullableStructsFileUpload{value: val, isSet: true}
}

func (v NullableStructsFileUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructsFileUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


