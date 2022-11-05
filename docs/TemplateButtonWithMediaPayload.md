# TemplateButtonWithMediaPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buttons** | [**[]TemplateButton**](TemplateButton.md) |  | 
**ContentText** | Pointer to **string** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**Media** | [**FileUpload**](FileUpload.md) |  | 
**To** | **string** |  | 

## Methods

### NewTemplateButtonWithMediaPayload

`func NewTemplateButtonWithMediaPayload(buttons []TemplateButton, media FileUpload, to string, ) *TemplateButtonWithMediaPayload`

NewTemplateButtonWithMediaPayload instantiates a new TemplateButtonWithMediaPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateButtonWithMediaPayloadWithDefaults

`func NewTemplateButtonWithMediaPayloadWithDefaults() *TemplateButtonWithMediaPayload`

NewTemplateButtonWithMediaPayloadWithDefaults instantiates a new TemplateButtonWithMediaPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtons

`func (o *TemplateButtonWithMediaPayload) GetButtons() []TemplateButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *TemplateButtonWithMediaPayload) GetButtonsOk() (*[]TemplateButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *TemplateButtonWithMediaPayload) SetButtons(v []TemplateButton)`

SetButtons sets Buttons field to given value.


### GetContentText

`func (o *TemplateButtonWithMediaPayload) GetContentText() string`

GetContentText returns the ContentText field if non-nil, zero value otherwise.

### GetContentTextOk

`func (o *TemplateButtonWithMediaPayload) GetContentTextOk() (*string, bool)`

GetContentTextOk returns a tuple with the ContentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentText

`func (o *TemplateButtonWithMediaPayload) SetContentText(v string)`

SetContentText sets ContentText field to given value.

### HasContentText

`func (o *TemplateButtonWithMediaPayload) HasContentText() bool`

HasContentText returns a boolean if a field has been set.

### GetFooter

`func (o *TemplateButtonWithMediaPayload) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *TemplateButtonWithMediaPayload) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *TemplateButtonWithMediaPayload) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *TemplateButtonWithMediaPayload) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetMedia

`func (o *TemplateButtonWithMediaPayload) GetMedia() FileUpload`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *TemplateButtonWithMediaPayload) GetMediaOk() (*FileUpload, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *TemplateButtonWithMediaPayload) SetMedia(v FileUpload)`

SetMedia sets Media field to given value.


### GetTo

`func (o *TemplateButtonWithMediaPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TemplateButtonWithMediaPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TemplateButtonWithMediaPayload) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


