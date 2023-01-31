# RollbackSecretOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewRollbackSecretOutput

`func NewRollbackSecretOutput() *RollbackSecretOutput`

NewRollbackSecretOutput instantiates a new RollbackSecretOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackSecretOutputWithDefaults

`func NewRollbackSecretOutputWithDefaults() *RollbackSecretOutput`

NewRollbackSecretOutputWithDefaults instantiates a new RollbackSecretOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RollbackSecretOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RollbackSecretOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RollbackSecretOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RollbackSecretOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *RollbackSecretOutput) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RollbackSecretOutput) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RollbackSecretOutput) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RollbackSecretOutput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


