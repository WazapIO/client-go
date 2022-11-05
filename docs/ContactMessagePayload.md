# ContactMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** |  | 
**Vcard** | [**ContactMessagePayloadVcard**](ContactMessagePayloadVcard.md) |  | 

## Methods

### NewContactMessagePayload

`func NewContactMessagePayload(to string, vcard ContactMessagePayloadVcard, ) *ContactMessagePayload`

NewContactMessagePayload instantiates a new ContactMessagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactMessagePayloadWithDefaults

`func NewContactMessagePayloadWithDefaults() *ContactMessagePayload`

NewContactMessagePayloadWithDefaults instantiates a new ContactMessagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *ContactMessagePayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContactMessagePayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContactMessagePayload) SetTo(v string)`

SetTo sets To field to given value.


### GetVcard

`func (o *ContactMessagePayload) GetVcard() ContactMessagePayloadVcard`

GetVcard returns the Vcard field if non-nil, zero value otherwise.

### GetVcardOk

`func (o *ContactMessagePayload) GetVcardOk() (*ContactMessagePayloadVcard, bool)`

GetVcardOk returns a tuple with the Vcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcard

`func (o *ContactMessagePayload) SetVcard(v ContactMessagePayloadVcard)`

SetVcard sets Vcard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


