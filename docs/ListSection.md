# ListSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | [**[]ListItem**](ListItem.md) |  | 
**Title** | **string** |  | 

## Methods

### NewListSection

`func NewListSection(rows []ListItem, title string, ) *ListSection`

NewListSection instantiates a new ListSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSectionWithDefaults

`func NewListSectionWithDefaults() *ListSection`

NewListSectionWithDefaults instantiates a new ListSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *ListSection) GetRows() []ListItem`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListSection) GetRowsOk() (*[]ListItem, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListSection) SetRows(v []ListItem)`

SetRows sets Rows field to given value.


### GetTitle

`func (o *ListSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListSection) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


