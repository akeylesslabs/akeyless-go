# RollbackSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Secret name | 
**OldVersion** | **int32** | Old secret version to rollback to | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRollbackSecret

`func NewRollbackSecret(name string, oldVersion int32, ) *RollbackSecret`

NewRollbackSecret instantiates a new RollbackSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackSecretWithDefaults

`func NewRollbackSecretWithDefaults() *RollbackSecret`

NewRollbackSecretWithDefaults instantiates a new RollbackSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RollbackSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RollbackSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RollbackSecret) SetName(v string)`

SetName sets Name field to given value.


### GetOldVersion

`func (o *RollbackSecret) GetOldVersion() int32`

GetOldVersion returns the OldVersion field if non-nil, zero value otherwise.

### GetOldVersionOk

`func (o *RollbackSecret) GetOldVersionOk() (*int32, bool)`

GetOldVersionOk returns a tuple with the OldVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldVersion

`func (o *RollbackSecret) SetOldVersion(v int32)`

SetOldVersion sets OldVersion field to given value.


### GetToken

`func (o *RollbackSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RollbackSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RollbackSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RollbackSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RollbackSecret) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RollbackSecret) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RollbackSecret) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RollbackSecret) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


