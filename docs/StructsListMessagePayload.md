# StructsListMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ButtonText** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MultiSelect** | Pointer to **bool** |  | [optional] 
**Sections** | [**[]StructsListSection**](StructsListSection.md) |  | 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**To** | **string** |  | 

## Methods

### NewStructsListMessagePayload

`func NewStructsListMessagePayload(sections []StructsListSection, to string, ) *StructsListMessagePayload`

NewStructsListMessagePayload instantiates a new StructsListMessagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructsListMessagePayloadWithDefaults

`func NewStructsListMessagePayloadWithDefaults() *StructsListMessagePayload`

NewStructsListMessagePayloadWithDefaults instantiates a new StructsListMessagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtonText

`func (o *StructsListMessagePayload) GetButtonText() string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *StructsListMessagePayload) GetButtonTextOk() (*string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *StructsListMessagePayload) SetButtonText(v string)`

SetButtonText sets ButtonText field to given value.

### HasButtonText

`func (o *StructsListMessagePayload) HasButtonText() bool`

HasButtonText returns a boolean if a field has been set.

### GetDescription

`func (o *StructsListMessagePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StructsListMessagePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StructsListMessagePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StructsListMessagePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiSelect

`func (o *StructsListMessagePayload) GetMultiSelect() bool`

GetMultiSelect returns the MultiSelect field if non-nil, zero value otherwise.

### GetMultiSelectOk

`func (o *StructsListMessagePayload) GetMultiSelectOk() (*bool, bool)`

GetMultiSelectOk returns a tuple with the MultiSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSelect

`func (o *StructsListMessagePayload) SetMultiSelect(v bool)`

SetMultiSelect sets MultiSelect field to given value.

### HasMultiSelect

`func (o *StructsListMessagePayload) HasMultiSelect() bool`

HasMultiSelect returns a boolean if a field has been set.

### GetSections

`func (o *StructsListMessagePayload) GetSections() []StructsListSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *StructsListMessagePayload) GetSectionsOk() (*[]StructsListSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *StructsListMessagePayload) SetSections(v []StructsListSection)`

SetSections sets Sections field to given value.


### GetText

`func (o *StructsListMessagePayload) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StructsListMessagePayload) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StructsListMessagePayload) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StructsListMessagePayload) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *StructsListMessagePayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StructsListMessagePayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StructsListMessagePayload) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StructsListMessagePayload) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTo

`func (o *StructsListMessagePayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *StructsListMessagePayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *StructsListMessagePayload) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


