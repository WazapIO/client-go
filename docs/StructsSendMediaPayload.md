# StructsSendMediaPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caption** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**MediaData** | [**StructsFileUpload**](StructsFileUpload.md) |  | 
**MediaType** | **string** |  | 
**To** | **string** |  | 

## Methods

### NewStructsSendMediaPayload

`func NewStructsSendMediaPayload(mediaData StructsFileUpload, mediaType string, to string, ) *StructsSendMediaPayload`

NewStructsSendMediaPayload instantiates a new StructsSendMediaPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructsSendMediaPayloadWithDefaults

`func NewStructsSendMediaPayloadWithDefaults() *StructsSendMediaPayload`

NewStructsSendMediaPayloadWithDefaults instantiates a new StructsSendMediaPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaption

`func (o *StructsSendMediaPayload) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *StructsSendMediaPayload) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *StructsSendMediaPayload) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *StructsSendMediaPayload) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetFilename

`func (o *StructsSendMediaPayload) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *StructsSendMediaPayload) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *StructsSendMediaPayload) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *StructsSendMediaPayload) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetMediaData

`func (o *StructsSendMediaPayload) GetMediaData() StructsFileUpload`

GetMediaData returns the MediaData field if non-nil, zero value otherwise.

### GetMediaDataOk

`func (o *StructsSendMediaPayload) GetMediaDataOk() (*StructsFileUpload, bool)`

GetMediaDataOk returns a tuple with the MediaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaData

`func (o *StructsSendMediaPayload) SetMediaData(v StructsFileUpload)`

SetMediaData sets MediaData field to given value.


### GetMediaType

`func (o *StructsSendMediaPayload) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *StructsSendMediaPayload) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *StructsSendMediaPayload) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetTo

`func (o *StructsSendMediaPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *StructsSendMediaPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *StructsSendMediaPayload) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


