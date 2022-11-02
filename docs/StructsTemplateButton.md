# StructsTemplateButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ButtonId** | Pointer to **string** | Make sure this is unique | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Type** | **string** | Valid types are replyButton, urlButton &amp; callButton | 

## Methods

### NewStructsTemplateButton

`func NewStructsTemplateButton(title string, type_ string, ) *StructsTemplateButton`

NewStructsTemplateButton instantiates a new StructsTemplateButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructsTemplateButtonWithDefaults

`func NewStructsTemplateButtonWithDefaults() *StructsTemplateButton`

NewStructsTemplateButtonWithDefaults instantiates a new StructsTemplateButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtonId

`func (o *StructsTemplateButton) GetButtonId() string`

GetButtonId returns the ButtonId field if non-nil, zero value otherwise.

### GetButtonIdOk

`func (o *StructsTemplateButton) GetButtonIdOk() (*string, bool)`

GetButtonIdOk returns a tuple with the ButtonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonId

`func (o *StructsTemplateButton) SetButtonId(v string)`

SetButtonId sets ButtonId field to given value.

### HasButtonId

`func (o *StructsTemplateButton) HasButtonId() bool`

HasButtonId returns a boolean if a field has been set.

### GetPayload

`func (o *StructsTemplateButton) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *StructsTemplateButton) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *StructsTemplateButton) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *StructsTemplateButton) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTitle

`func (o *StructsTemplateButton) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StructsTemplateButton) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StructsTemplateButton) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *StructsTemplateButton) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StructsTemplateButton) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StructsTemplateButton) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


