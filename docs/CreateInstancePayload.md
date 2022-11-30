# CreateInstancePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceKey** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateInstancePayload

`func NewCreateInstancePayload() *CreateInstancePayload`

NewCreateInstancePayload instantiates a new CreateInstancePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstancePayloadWithDefaults

`func NewCreateInstancePayloadWithDefaults() *CreateInstancePayload`

NewCreateInstancePayloadWithDefaults instantiates a new CreateInstancePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceKey

`func (o *CreateInstancePayload) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *CreateInstancePayload) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *CreateInstancePayload) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.

### HasInstanceKey

`func (o *CreateInstancePayload) HasInstanceKey() bool`

HasInstanceKey returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateInstancePayload) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateInstancePayload) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateInstancePayload) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateInstancePayload) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


