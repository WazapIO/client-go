# GroupInviteMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caption** | Pointer to **string** |  | [optional] 
**ExpiryMinutes** | Pointer to **int32** |  | [optional] 
**InviteCode** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupInviteMessagePayload

`func NewGroupInviteMessagePayload() *GroupInviteMessagePayload`

NewGroupInviteMessagePayload instantiates a new GroupInviteMessagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupInviteMessagePayloadWithDefaults

`func NewGroupInviteMessagePayloadWithDefaults() *GroupInviteMessagePayload`

NewGroupInviteMessagePayloadWithDefaults instantiates a new GroupInviteMessagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaption

`func (o *GroupInviteMessagePayload) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *GroupInviteMessagePayload) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *GroupInviteMessagePayload) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *GroupInviteMessagePayload) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetExpiryMinutes

`func (o *GroupInviteMessagePayload) GetExpiryMinutes() int32`

GetExpiryMinutes returns the ExpiryMinutes field if non-nil, zero value otherwise.

### GetExpiryMinutesOk

`func (o *GroupInviteMessagePayload) GetExpiryMinutesOk() (*int32, bool)`

GetExpiryMinutesOk returns a tuple with the ExpiryMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMinutes

`func (o *GroupInviteMessagePayload) SetExpiryMinutes(v int32)`

SetExpiryMinutes sets ExpiryMinutes field to given value.

### HasExpiryMinutes

`func (o *GroupInviteMessagePayload) HasExpiryMinutes() bool`

HasExpiryMinutes returns a boolean if a field has been set.

### GetInviteCode

`func (o *GroupInviteMessagePayload) GetInviteCode() string`

GetInviteCode returns the InviteCode field if non-nil, zero value otherwise.

### GetInviteCodeOk

`func (o *GroupInviteMessagePayload) GetInviteCodeOk() (*string, bool)`

GetInviteCodeOk returns a tuple with the InviteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteCode

`func (o *GroupInviteMessagePayload) SetInviteCode(v string)`

SetInviteCode sets InviteCode field to given value.

### HasInviteCode

`func (o *GroupInviteMessagePayload) HasInviteCode() bool`

HasInviteCode returns a boolean if a field has been set.

### GetTo

`func (o *GroupInviteMessagePayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GroupInviteMessagePayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GroupInviteMessagePayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GroupInviteMessagePayload) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


