# StructsTemplateButtonPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buttons** | [**[]StructsTemplateButton**](StructsTemplateButton.md) |  | 
**ContentText** | Pointer to **string** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **string** |  | [optional] 
**To** | **string** |  | 

## Methods

### NewStructsTemplateButtonPayload

`func NewStructsTemplateButtonPayload(buttons []StructsTemplateButton, to string, ) *StructsTemplateButtonPayload`

NewStructsTemplateButtonPayload instantiates a new StructsTemplateButtonPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructsTemplateButtonPayloadWithDefaults

`func NewStructsTemplateButtonPayloadWithDefaults() *StructsTemplateButtonPayload`

NewStructsTemplateButtonPayloadWithDefaults instantiates a new StructsTemplateButtonPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtons

`func (o *StructsTemplateButtonPayload) GetButtons() []StructsTemplateButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *StructsTemplateButtonPayload) GetButtonsOk() (*[]StructsTemplateButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *StructsTemplateButtonPayload) SetButtons(v []StructsTemplateButton)`

SetButtons sets Buttons field to given value.


### GetContentText

`func (o *StructsTemplateButtonPayload) GetContentText() string`

GetContentText returns the ContentText field if non-nil, zero value otherwise.

### GetContentTextOk

`func (o *StructsTemplateButtonPayload) GetContentTextOk() (*string, bool)`

GetContentTextOk returns a tuple with the ContentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentText

`func (o *StructsTemplateButtonPayload) SetContentText(v string)`

SetContentText sets ContentText field to given value.

### HasContentText

`func (o *StructsTemplateButtonPayload) HasContentText() bool`

HasContentText returns a boolean if a field has been set.

### GetFooter

`func (o *StructsTemplateButtonPayload) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *StructsTemplateButtonPayload) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *StructsTemplateButtonPayload) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *StructsTemplateButtonPayload) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHeader

`func (o *StructsTemplateButtonPayload) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *StructsTemplateButtonPayload) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *StructsTemplateButtonPayload) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *StructsTemplateButtonPayload) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetTo

`func (o *StructsTemplateButtonPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *StructsTemplateButtonPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *StructsTemplateButtonPayload) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


