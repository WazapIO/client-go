# StructsTemplateButtonWithMediaPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buttons** | [**[]StructsTemplateButton**](StructsTemplateButton.md) |  | 
**ContentText** | Pointer to **string** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**Media** | [**StructsFileUpload**](StructsFileUpload.md) |  | 
**To** | **string** |  | 

## Methods

### NewStructsTemplateButtonWithMediaPayload

`func NewStructsTemplateButtonWithMediaPayload(buttons []StructsTemplateButton, media StructsFileUpload, to string, ) *StructsTemplateButtonWithMediaPayload`

NewStructsTemplateButtonWithMediaPayload instantiates a new StructsTemplateButtonWithMediaPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructsTemplateButtonWithMediaPayloadWithDefaults

`func NewStructsTemplateButtonWithMediaPayloadWithDefaults() *StructsTemplateButtonWithMediaPayload`

NewStructsTemplateButtonWithMediaPayloadWithDefaults instantiates a new StructsTemplateButtonWithMediaPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtons

`func (o *StructsTemplateButtonWithMediaPayload) GetButtons() []StructsTemplateButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *StructsTemplateButtonWithMediaPayload) GetButtonsOk() (*[]StructsTemplateButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *StructsTemplateButtonWithMediaPayload) SetButtons(v []StructsTemplateButton)`

SetButtons sets Buttons field to given value.


### GetContentText

`func (o *StructsTemplateButtonWithMediaPayload) GetContentText() string`

GetContentText returns the ContentText field if non-nil, zero value otherwise.

### GetContentTextOk

`func (o *StructsTemplateButtonWithMediaPayload) GetContentTextOk() (*string, bool)`

GetContentTextOk returns a tuple with the ContentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentText

`func (o *StructsTemplateButtonWithMediaPayload) SetContentText(v string)`

SetContentText sets ContentText field to given value.

### HasContentText

`func (o *StructsTemplateButtonWithMediaPayload) HasContentText() bool`

HasContentText returns a boolean if a field has been set.

### GetFooter

`func (o *StructsTemplateButtonWithMediaPayload) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *StructsTemplateButtonWithMediaPayload) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *StructsTemplateButtonWithMediaPayload) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *StructsTemplateButtonWithMediaPayload) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetMedia

`func (o *StructsTemplateButtonWithMediaPayload) GetMedia() StructsFileUpload`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *StructsTemplateButtonWithMediaPayload) GetMediaOk() (*StructsFileUpload, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *StructsTemplateButtonWithMediaPayload) SetMedia(v StructsFileUpload)`

SetMedia sets Media field to given value.


### GetTo

`func (o *StructsTemplateButtonWithMediaPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *StructsTemplateButtonWithMediaPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *StructsTemplateButtonWithMediaPayload) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


