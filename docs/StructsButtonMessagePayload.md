# StructsButtonMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buttons** | Pointer to [**[]StructsReplyButton**](StructsReplyButton.md) |  | [optional] 
**FooterText** | Pointer to **string** |  | [optional] 
**HeaderText** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewStructsButtonMessagePayload

`func NewStructsButtonMessagePayload() *StructsButtonMessagePayload`

NewStructsButtonMessagePayload instantiates a new StructsButtonMessagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructsButtonMessagePayloadWithDefaults

`func NewStructsButtonMessagePayloadWithDefaults() *StructsButtonMessagePayload`

NewStructsButtonMessagePayloadWithDefaults instantiates a new StructsButtonMessagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtons

`func (o *StructsButtonMessagePayload) GetButtons() []StructsReplyButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *StructsButtonMessagePayload) GetButtonsOk() (*[]StructsReplyButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *StructsButtonMessagePayload) SetButtons(v []StructsReplyButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *StructsButtonMessagePayload) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetFooterText

`func (o *StructsButtonMessagePayload) GetFooterText() string`

GetFooterText returns the FooterText field if non-nil, zero value otherwise.

### GetFooterTextOk

`func (o *StructsButtonMessagePayload) GetFooterTextOk() (*string, bool)`

GetFooterTextOk returns a tuple with the FooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterText

`func (o *StructsButtonMessagePayload) SetFooterText(v string)`

SetFooterText sets FooterText field to given value.

### HasFooterText

`func (o *StructsButtonMessagePayload) HasFooterText() bool`

HasFooterText returns a boolean if a field has been set.

### GetHeaderText

`func (o *StructsButtonMessagePayload) GetHeaderText() string`

GetHeaderText returns the HeaderText field if non-nil, zero value otherwise.

### GetHeaderTextOk

`func (o *StructsButtonMessagePayload) GetHeaderTextOk() (*string, bool)`

GetHeaderTextOk returns a tuple with the HeaderText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderText

`func (o *StructsButtonMessagePayload) SetHeaderText(v string)`

SetHeaderText sets HeaderText field to given value.

### HasHeaderText

`func (o *StructsButtonMessagePayload) HasHeaderText() bool`

HasHeaderText returns a boolean if a field has been set.

### GetText

`func (o *StructsButtonMessagePayload) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StructsButtonMessagePayload) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StructsButtonMessagePayload) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StructsButtonMessagePayload) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTo

`func (o *StructsButtonMessagePayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *StructsButtonMessagePayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *StructsButtonMessagePayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *StructsButtonMessagePayload) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


