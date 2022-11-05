# ButtonMessageWithMediaPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buttons** | Pointer to [**[]ReplyButton**](ReplyButton.md) |  | [optional] 
**FooterText** | Pointer to **string** |  | [optional] 
**MediaData** | Pointer to [**FileUpload**](FileUpload.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewButtonMessageWithMediaPayload

`func NewButtonMessageWithMediaPayload() *ButtonMessageWithMediaPayload`

NewButtonMessageWithMediaPayload instantiates a new ButtonMessageWithMediaPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewButtonMessageWithMediaPayloadWithDefaults

`func NewButtonMessageWithMediaPayloadWithDefaults() *ButtonMessageWithMediaPayload`

NewButtonMessageWithMediaPayloadWithDefaults instantiates a new ButtonMessageWithMediaPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtons

`func (o *ButtonMessageWithMediaPayload) GetButtons() []ReplyButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *ButtonMessageWithMediaPayload) GetButtonsOk() (*[]ReplyButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *ButtonMessageWithMediaPayload) SetButtons(v []ReplyButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *ButtonMessageWithMediaPayload) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetFooterText

`func (o *ButtonMessageWithMediaPayload) GetFooterText() string`

GetFooterText returns the FooterText field if non-nil, zero value otherwise.

### GetFooterTextOk

`func (o *ButtonMessageWithMediaPayload) GetFooterTextOk() (*string, bool)`

GetFooterTextOk returns a tuple with the FooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterText

`func (o *ButtonMessageWithMediaPayload) SetFooterText(v string)`

SetFooterText sets FooterText field to given value.

### HasFooterText

`func (o *ButtonMessageWithMediaPayload) HasFooterText() bool`

HasFooterText returns a boolean if a field has been set.

### GetMediaData

`func (o *ButtonMessageWithMediaPayload) GetMediaData() FileUpload`

GetMediaData returns the MediaData field if non-nil, zero value otherwise.

### GetMediaDataOk

`func (o *ButtonMessageWithMediaPayload) GetMediaDataOk() (*FileUpload, bool)`

GetMediaDataOk returns a tuple with the MediaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaData

`func (o *ButtonMessageWithMediaPayload) SetMediaData(v FileUpload)`

SetMediaData sets MediaData field to given value.

### HasMediaData

`func (o *ButtonMessageWithMediaPayload) HasMediaData() bool`

HasMediaData returns a boolean if a field has been set.

### GetText

`func (o *ButtonMessageWithMediaPayload) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ButtonMessageWithMediaPayload) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ButtonMessageWithMediaPayload) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ButtonMessageWithMediaPayload) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTo

`func (o *ButtonMessageWithMediaPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ButtonMessageWithMediaPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ButtonMessageWithMediaPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ButtonMessageWithMediaPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


