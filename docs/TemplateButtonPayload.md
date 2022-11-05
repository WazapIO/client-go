# TemplateButtonPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buttons** | [**[]TemplateButton**](TemplateButton.md) |  | 
**ContentText** | Pointer to **string** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **string** |  | [optional] 
**To** | **string** |  | 

## Methods

### NewTemplateButtonPayload

`func NewTemplateButtonPayload(buttons []TemplateButton, to string, ) *TemplateButtonPayload`

NewTemplateButtonPayload instantiates a new TemplateButtonPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateButtonPayloadWithDefaults

`func NewTemplateButtonPayloadWithDefaults() *TemplateButtonPayload`

NewTemplateButtonPayloadWithDefaults instantiates a new TemplateButtonPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtons

`func (o *TemplateButtonPayload) GetButtons() []TemplateButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *TemplateButtonPayload) GetButtonsOk() (*[]TemplateButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *TemplateButtonPayload) SetButtons(v []TemplateButton)`

SetButtons sets Buttons field to given value.


### GetContentText

`func (o *TemplateButtonPayload) GetContentText() string`

GetContentText returns the ContentText field if non-nil, zero value otherwise.

### GetContentTextOk

`func (o *TemplateButtonPayload) GetContentTextOk() (*string, bool)`

GetContentTextOk returns a tuple with the ContentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentText

`func (o *TemplateButtonPayload) SetContentText(v string)`

SetContentText sets ContentText field to given value.

### HasContentText

`func (o *TemplateButtonPayload) HasContentText() bool`

HasContentText returns a boolean if a field has been set.

### GetFooter

`func (o *TemplateButtonPayload) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *TemplateButtonPayload) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *TemplateButtonPayload) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *TemplateButtonPayload) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHeader

`func (o *TemplateButtonPayload) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *TemplateButtonPayload) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *TemplateButtonPayload) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *TemplateButtonPayload) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetTo

`func (o *TemplateButtonPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TemplateButtonPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TemplateButtonPayload) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


