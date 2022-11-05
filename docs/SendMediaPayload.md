# SendMediaPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caption** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**MediaData** | [**FileUpload**](FileUpload.md) |  | 
**MediaType** | **string** |  | 
**To** | **string** |  | 

## Methods

### NewSendMediaPayload

`func NewSendMediaPayload(mediaData FileUpload, mediaType string, to string, ) *SendMediaPayload`

NewSendMediaPayload instantiates a new SendMediaPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMediaPayloadWithDefaults

`func NewSendMediaPayloadWithDefaults() *SendMediaPayload`

NewSendMediaPayloadWithDefaults instantiates a new SendMediaPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaption

`func (o *SendMediaPayload) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *SendMediaPayload) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *SendMediaPayload) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *SendMediaPayload) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetFilename

`func (o *SendMediaPayload) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SendMediaPayload) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SendMediaPayload) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SendMediaPayload) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetMediaData

`func (o *SendMediaPayload) GetMediaData() FileUpload`

GetMediaData returns the MediaData field if non-nil, zero value otherwise.

### GetMediaDataOk

`func (o *SendMediaPayload) GetMediaDataOk() (*FileUpload, bool)`

GetMediaDataOk returns a tuple with the MediaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaData

`func (o *SendMediaPayload) SetMediaData(v FileUpload)`

SetMediaData sets MediaData field to given value.


### GetMediaType

`func (o *SendMediaPayload) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *SendMediaPayload) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *SendMediaPayload) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetTo

`func (o *SendMediaPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SendMediaPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SendMediaPayload) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


