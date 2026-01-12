# TargetError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 

## Methods

### NewTargetError

`func NewTargetError() *TargetError`

NewTargetError instantiates a new TargetError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetErrorWithDefaults

`func NewTargetErrorWithDefaults() *TargetError`

NewTargetErrorWithDefaults instantiates a new TargetError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *TargetError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TargetError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TargetError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TargetError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHost

`func (o *TargetError) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetError) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetError) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TargetError) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *TargetError) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetError) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetError) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetError) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


