# ValidateTokenOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to **string** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewValidateTokenOutput

`func NewValidateTokenOutput() *ValidateTokenOutput`

NewValidateTokenOutput instantiates a new ValidateTokenOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateTokenOutputWithDefaults

`func NewValidateTokenOutputWithDefaults() *ValidateTokenOutput`

NewValidateTokenOutputWithDefaults instantiates a new ValidateTokenOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *ValidateTokenOutput) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ValidateTokenOutput) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ValidateTokenOutput) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ValidateTokenOutput) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetIsValid

`func (o *ValidateTokenOutput) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ValidateTokenOutput) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ValidateTokenOutput) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *ValidateTokenOutput) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetReason

`func (o *ValidateTokenOutput) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ValidateTokenOutput) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ValidateTokenOutput) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ValidateTokenOutput) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


