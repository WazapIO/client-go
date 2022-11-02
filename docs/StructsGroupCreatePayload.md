# StructsGroupCreatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStructsGroupCreatePayload

`func NewStructsGroupCreatePayload() *StructsGroupCreatePayload`

NewStructsGroupCreatePayload instantiates a new StructsGroupCreatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructsGroupCreatePayloadWithDefaults

`func NewStructsGroupCreatePayloadWithDefaults() *StructsGroupCreatePayload`

NewStructsGroupCreatePayloadWithDefaults instantiates a new StructsGroupCreatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *StructsGroupCreatePayload) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *StructsGroupCreatePayload) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *StructsGroupCreatePayload) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *StructsGroupCreatePayload) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetParticipants

`func (o *StructsGroupCreatePayload) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *StructsGroupCreatePayload) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *StructsGroupCreatePayload) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *StructsGroupCreatePayload) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


