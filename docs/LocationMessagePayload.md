# LocationMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**LocationMessagePayloadLocation**](LocationMessagePayloadLocation.md) |  | 
**To** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewLocationMessagePayload

`func NewLocationMessagePayload(location LocationMessagePayloadLocation, to string, ) *LocationMessagePayload`

NewLocationMessagePayload instantiates a new LocationMessagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationMessagePayloadWithDefaults

`func NewLocationMessagePayloadWithDefaults() *LocationMessagePayload`

NewLocationMessagePayloadWithDefaults instantiates a new LocationMessagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *LocationMessagePayload) GetLocation() LocationMessagePayloadLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LocationMessagePayload) GetLocationOk() (*LocationMessagePayloadLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LocationMessagePayload) SetLocation(v LocationMessagePayloadLocation)`

SetLocation sets Location field to given value.


### GetTo

`func (o *LocationMessagePayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LocationMessagePayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LocationMessagePayload) SetTo(v string)`

SetTo sets To field to given value.


### GetUrl

`func (o *LocationMessagePayload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LocationMessagePayload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LocationMessagePayload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LocationMessagePayload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


