# ManagedKeyStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | Pointer to **int64** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastStatus** | Pointer to **string** | ManagedKeyTargetStatus defines status of managed key target | [optional] 

## Methods

### NewManagedKeyStatusInfo

`func NewManagedKeyStatusInfo() *ManagedKeyStatusInfo`

NewManagedKeyStatusInfo instantiates a new ManagedKeyStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedKeyStatusInfoWithDefaults

`func NewManagedKeyStatusInfoWithDefaults() *ManagedKeyStatusInfo`

NewManagedKeyStatusInfoWithDefaults instantiates a new ManagedKeyStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *ManagedKeyStatusInfo) GetKeyId() int64`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ManagedKeyStatusInfo) GetKeyIdOk() (*int64, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ManagedKeyStatusInfo) SetKeyId(v int64)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ManagedKeyStatusInfo) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetLastError

`func (o *ManagedKeyStatusInfo) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ManagedKeyStatusInfo) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ManagedKeyStatusInfo) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ManagedKeyStatusInfo) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastStatus

`func (o *ManagedKeyStatusInfo) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *ManagedKeyStatusInfo) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *ManagedKeyStatusInfo) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *ManagedKeyStatusInfo) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


