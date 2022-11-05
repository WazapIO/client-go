# ListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**RowId** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewListItem

`func NewListItem(title string, ) *ListItem`

NewListItem instantiates a new ListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemWithDefaults

`func NewListItemWithDefaults() *ListItem`

NewListItemWithDefaults instantiates a new ListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRowId

`func (o *ListItem) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *ListItem) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *ListItem) SetRowId(v string)`

SetRowId sets RowId field to given value.

### HasRowId

`func (o *ListItem) HasRowId() bool`

HasRowId returns a boolean if a field has been set.

### GetTitle

`func (o *ListItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListItem) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


