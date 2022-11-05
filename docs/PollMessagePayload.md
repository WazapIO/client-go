# PollMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | **[]string** |  | 
**SelectableOptionsCount** | **int32** |  | 
**Title** | **string** |  | 
**To** | **string** |  | 

## Methods

### NewPollMessagePayload

`func NewPollMessagePayload(options []string, selectableOptionsCount int32, title string, to string, ) *PollMessagePayload`

NewPollMessagePayload instantiates a new PollMessagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollMessagePayloadWithDefaults

`func NewPollMessagePayloadWithDefaults() *PollMessagePayload`

NewPollMessagePayloadWithDefaults instantiates a new PollMessagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *PollMessagePayload) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PollMessagePayload) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PollMessagePayload) SetOptions(v []string)`

SetOptions sets Options field to given value.


### GetSelectableOptionsCount

`func (o *PollMessagePayload) GetSelectableOptionsCount() int32`

GetSelectableOptionsCount returns the SelectableOptionsCount field if non-nil, zero value otherwise.

### GetSelectableOptionsCountOk

`func (o *PollMessagePayload) GetSelectableOptionsCountOk() (*int32, bool)`

GetSelectableOptionsCountOk returns a tuple with the SelectableOptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectableOptionsCount

`func (o *PollMessagePayload) SetSelectableOptionsCount(v int32)`

SetSelectableOptionsCount sets SelectableOptionsCount field to given value.


### GetTitle

`func (o *PollMessagePayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PollMessagePayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PollMessagePayload) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTo

`func (o *PollMessagePayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PollMessagePayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PollMessagePayload) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


