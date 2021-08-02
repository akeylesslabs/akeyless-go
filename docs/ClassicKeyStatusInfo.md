# ClassicKeyStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastStatus** | Pointer to **string** | ClassicKeyTargetStatus defines status of classic key target | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewClassicKeyStatusInfo

`func NewClassicKeyStatusInfo() *ClassicKeyStatusInfo`

NewClassicKeyStatusInfo instantiates a new ClassicKeyStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassicKeyStatusInfoWithDefaults

`func NewClassicKeyStatusInfoWithDefaults() *ClassicKeyStatusInfo`

NewClassicKeyStatusInfoWithDefaults instantiates a new ClassicKeyStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorDate

`func (o *ClassicKeyStatusInfo) GetErrorDate() time.Time`

GetErrorDate returns the ErrorDate field if non-nil, zero value otherwise.

### GetErrorDateOk

`func (o *ClassicKeyStatusInfo) GetErrorDateOk() (*time.Time, bool)`

GetErrorDateOk returns a tuple with the ErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDate

`func (o *ClassicKeyStatusInfo) SetErrorDate(v time.Time)`

SetErrorDate sets ErrorDate field to given value.

### HasErrorDate

`func (o *ClassicKeyStatusInfo) HasErrorDate() bool`

HasErrorDate returns a boolean if a field has been set.

### GetLastError

`func (o *ClassicKeyStatusInfo) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ClassicKeyStatusInfo) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ClassicKeyStatusInfo) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ClassicKeyStatusInfo) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastStatus

`func (o *ClassicKeyStatusInfo) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *ClassicKeyStatusInfo) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *ClassicKeyStatusInfo) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *ClassicKeyStatusInfo) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ClassicKeyStatusInfo) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClassicKeyStatusInfo) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClassicKeyStatusInfo) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClassicKeyStatusInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


