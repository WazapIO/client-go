# PaymentRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ExpiryMinutes** | Pointer to **int32** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentRequestPayload

`func NewPaymentRequestPayload() *PaymentRequestPayload`

NewPaymentRequestPayload instantiates a new PaymentRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestPayloadWithDefaults

`func NewPaymentRequestPayloadWithDefaults() *PaymentRequestPayload`

NewPaymentRequestPayloadWithDefaults instantiates a new PaymentRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentRequestPayload) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequestPayload) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequestPayload) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentRequestPayload) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentRequestPayload) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentRequestPayload) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentRequestPayload) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentRequestPayload) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExpiryMinutes

`func (o *PaymentRequestPayload) GetExpiryMinutes() int32`

GetExpiryMinutes returns the ExpiryMinutes field if non-nil, zero value otherwise.

### GetExpiryMinutesOk

`func (o *PaymentRequestPayload) GetExpiryMinutesOk() (*int32, bool)`

GetExpiryMinutesOk returns a tuple with the ExpiryMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMinutes

`func (o *PaymentRequestPayload) SetExpiryMinutes(v int32)`

SetExpiryMinutes sets ExpiryMinutes field to given value.

### HasExpiryMinutes

`func (o *PaymentRequestPayload) HasExpiryMinutes() bool`

HasExpiryMinutes returns a boolean if a field has been set.

### GetTo

`func (o *PaymentRequestPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PaymentRequestPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PaymentRequestPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *PaymentRequestPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


