# ListMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ButtonText** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MultiSelect** | Pointer to **bool** |  | [optional] 
**Sections** | [**[]ListSection**](ListSection.md) |  | 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**To** | **string** |  | 

## Methods

### NewListMessagePayload

`func NewListMessagePayload(sections []ListSection, to string, ) *ListMessagePayload`

NewListMessagePayload instantiates a new ListMessagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessagePayloadWithDefaults

`func NewListMessagePayloadWithDefaults() *ListMessagePayload`

NewListMessagePayloadWithDefaults instantiates a new ListMessagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtonText

`func (o *ListMessagePayload) GetButtonText() string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *ListMessagePayload) GetButtonTextOk() (*string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *ListMessagePayload) SetButtonText(v string)`

SetButtonText sets ButtonText field to given value.

### HasButtonText

`func (o *ListMessagePayload) HasButtonText() bool`

HasButtonText returns a boolean if a field has been set.

### GetDescription

`func (o *ListMessagePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListMessagePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListMessagePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListMessagePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiSelect

`func (o *ListMessagePayload) GetMultiSelect() bool`

GetMultiSelect returns the MultiSelect field if non-nil, zero value otherwise.

### GetMultiSelectOk

`func (o *ListMessagePayload) GetMultiSelectOk() (*bool, bool)`

GetMultiSelectOk returns a tuple with the MultiSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSelect

`func (o *ListMessagePayload) SetMultiSelect(v bool)`

SetMultiSelect sets MultiSelect field to given value.

### HasMultiSelect

`func (o *ListMessagePayload) HasMultiSelect() bool`

HasMultiSelect returns a boolean if a field has been set.

### GetSections

`func (o *ListMessagePayload) GetSections() []ListSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ListMessagePayload) GetSectionsOk() (*[]ListSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ListMessagePayload) SetSections(v []ListSection)`

SetSections sets Sections field to given value.


### GetText

`func (o *ListMessagePayload) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ListMessagePayload) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ListMessagePayload) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ListMessagePayload) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *ListMessagePayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListMessagePayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListMessagePayload) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListMessagePayload) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTo

`func (o *ListMessagePayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ListMessagePayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ListMessagePayload) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


