# SetGroupPictureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** | Group picture in JPEG | 

## Methods

### NewSetGroupPictureRequest

`func NewSetGroupPictureRequest(file *os.File, ) *SetGroupPictureRequest`

NewSetGroupPictureRequest instantiates a new SetGroupPictureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetGroupPictureRequestWithDefaults

`func NewSetGroupPictureRequestWithDefaults() *SetGroupPictureRequest`

NewSetGroupPictureRequestWithDefaults instantiates a new SetGroupPictureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *SetGroupPictureRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SetGroupPictureRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SetGroupPictureRequest) SetFile(v *os.File)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


