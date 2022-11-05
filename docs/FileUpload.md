# FileUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectPath** | **string** |  | 
**FileEncSha256** | **[]int32** |  | 
**FileLength** | **int32** |  | 
**FileSha256** | **[]int32** |  | 
**MediaKey** | **[]int32** |  | 
**MimeType** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewFileUpload

`func NewFileUpload(directPath string, fileEncSha256 []int32, fileLength int32, fileSha256 []int32, mediaKey []int32, mimeType string, url string, ) *FileUpload`

NewFileUpload instantiates a new FileUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUploadWithDefaults

`func NewFileUploadWithDefaults() *FileUpload`

NewFileUploadWithDefaults instantiates a new FileUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectPath

`func (o *FileUpload) GetDirectPath() string`

GetDirectPath returns the DirectPath field if non-nil, zero value otherwise.

### GetDirectPathOk

`func (o *FileUpload) GetDirectPathOk() (*string, bool)`

GetDirectPathOk returns a tuple with the DirectPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPath

`func (o *FileUpload) SetDirectPath(v string)`

SetDirectPath sets DirectPath field to given value.


### GetFileEncSha256

`func (o *FileUpload) GetFileEncSha256() []int32`

GetFileEncSha256 returns the FileEncSha256 field if non-nil, zero value otherwise.

### GetFileEncSha256Ok

`func (o *FileUpload) GetFileEncSha256Ok() (*[]int32, bool)`

GetFileEncSha256Ok returns a tuple with the FileEncSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEncSha256

`func (o *FileUpload) SetFileEncSha256(v []int32)`

SetFileEncSha256 sets FileEncSha256 field to given value.


### GetFileLength

`func (o *FileUpload) GetFileLength() int32`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *FileUpload) GetFileLengthOk() (*int32, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *FileUpload) SetFileLength(v int32)`

SetFileLength sets FileLength field to given value.


### GetFileSha256

`func (o *FileUpload) GetFileSha256() []int32`

GetFileSha256 returns the FileSha256 field if non-nil, zero value otherwise.

### GetFileSha256Ok

`func (o *FileUpload) GetFileSha256Ok() (*[]int32, bool)`

GetFileSha256Ok returns a tuple with the FileSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSha256

`func (o *FileUpload) SetFileSha256(v []int32)`

SetFileSha256 sets FileSha256 field to given value.


### GetMediaKey

`func (o *FileUpload) GetMediaKey() []int32`

GetMediaKey returns the MediaKey field if non-nil, zero value otherwise.

### GetMediaKeyOk

`func (o *FileUpload) GetMediaKeyOk() (*[]int32, bool)`

GetMediaKeyOk returns a tuple with the MediaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaKey

`func (o *FileUpload) SetMediaKey(v []int32)`

SetMediaKey sets MediaKey field to given value.


### GetMimeType

`func (o *FileUpload) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *FileUpload) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *FileUpload) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetUrl

`func (o *FileUpload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileUpload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileUpload) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


