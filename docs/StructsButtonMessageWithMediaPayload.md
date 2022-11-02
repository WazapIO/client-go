# StructsButtonMessageWithMediaPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buttons** | Pointer to [**[]StructsReplyButton**](StructsReplyButton.md) |  | [optional] 
**FooterText** | Pointer to **string** |  | [optional] 
**MediaData** | Pointer to [**StructsFileUpload**](StructsFileUpload.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewStructsButtonMessageWithMediaPayload

`func NewStructsButtonMessageWithMediaPayload() *StructsButtonMessageWithMediaPayload`

NewStructsButtonMessageWithMediaPayload instantiates a new StructsButtonMessageWithMediaPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructsButtonMessageWithMediaPayloadWithDefaults

`func NewStructsButtonMessageWithMediaPayloadWithDefaults() *StructsButtonMessageWithMediaPayload`

NewStructsButtonMessageWithMediaPayloadWithDefaults instantiates a new StructsButtonMessageWithMediaPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtons

`func (o *StructsButtonMessageWithMediaPayload) GetButtons() []StructsReplyButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *StructsButtonMessageWithMediaPayload) GetButtonsOk() (*[]StructsReplyButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *StructsButtonMessageWithMediaPayload) SetButtons(v []StructsReplyButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *StructsButtonMessageWithMediaPayload) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetFooterText

`func (o *StructsButtonMessageWithMediaPayload) GetFooterText() string`

GetFooterText returns the FooterText field if non-nil, zero value otherwise.

### GetFooterTextOk

`func (o *StructsButtonMessageWithMediaPayload) GetFooterTextOk() (*string, bool)`

GetFooterTextOk returns a tuple with the FooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterText

`func (o *StructsButtonMessageWithMediaPayload) SetFooterText(v string)`

SetFooterText sets FooterText field to given value.

### HasFooterText

`func (o *StructsButtonMessageWithMediaPayload) HasFooterText() bool`

HasFooterText returns a boolean if a field has been set.

### GetMediaData

`func (o *StructsButtonMessageWithMediaPayload) GetMediaData() StructsFileUpload`

GetMediaData returns the MediaData field if non-nil, zero value otherwise.

### GetMediaDataOk

`func (o *StructsButtonMessageWithMediaPayload) GetMediaDataOk() (*StructsFileUpload, bool)`

GetMediaDataOk returns a tuple with the MediaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaData

`func (o *StructsButtonMessageWithMediaPayload) SetMediaData(v StructsFileUpload)`

SetMediaData sets MediaData field to given value.

### HasMediaData

`func (o *StructsButtonMessageWithMediaPayload) HasMediaData() bool`

HasMediaData returns a boolean if a field has been set.

### GetText

`func (o *StructsButtonMessageWithMediaPayload) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StructsButtonMessageWithMediaPayload) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StructsButtonMessageWithMediaPayload) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StructsButtonMessageWithMediaPayload) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTo

`func (o *StructsButtonMessageWithMediaPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *StructsButtonMessageWithMediaPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *StructsButtonMessageWithMediaPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *StructsButtonMessageWithMediaPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


