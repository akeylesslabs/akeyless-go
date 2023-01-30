# ConfigChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigHash** | Pointer to [**ConfigHash**](ConfigHash.md) |  | [optional] 
**LastChange** | Pointer to [**LastConfigChange**](LastConfigChange.md) |  | [optional] 
**LastStatus** | Pointer to [**LastStatusInfo**](LastStatusInfo.md) |  | [optional] 
**RequiredActivity** | Pointer to [**RequiredActivity**](RequiredActivity.md) |  | [optional] 
**UpdateStamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewConfigChange

`func NewConfigChange() *ConfigChange`

NewConfigChange instantiates a new ConfigChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigChangeWithDefaults

`func NewConfigChangeWithDefaults() *ConfigChange`

NewConfigChangeWithDefaults instantiates a new ConfigChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigHash

`func (o *ConfigChange) GetConfigHash() ConfigHash`

GetConfigHash returns the ConfigHash field if non-nil, zero value otherwise.

### GetConfigHashOk

`func (o *ConfigChange) GetConfigHashOk() (*ConfigHash, bool)`

GetConfigHashOk returns a tuple with the ConfigHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigHash

`func (o *ConfigChange) SetConfigHash(v ConfigHash)`

SetConfigHash sets ConfigHash field to given value.

### HasConfigHash

`func (o *ConfigChange) HasConfigHash() bool`

HasConfigHash returns a boolean if a field has been set.

### GetLastChange

`func (o *ConfigChange) GetLastChange() LastConfigChange`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *ConfigChange) GetLastChangeOk() (*LastConfigChange, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *ConfigChange) SetLastChange(v LastConfigChange)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *ConfigChange) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetLastStatus

`func (o *ConfigChange) GetLastStatus() LastStatusInfo`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *ConfigChange) GetLastStatusOk() (*LastStatusInfo, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *ConfigChange) SetLastStatus(v LastStatusInfo)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *ConfigChange) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetRequiredActivity

`func (o *ConfigChange) GetRequiredActivity() RequiredActivity`

GetRequiredActivity returns the RequiredActivity field if non-nil, zero value otherwise.

### GetRequiredActivityOk

`func (o *ConfigChange) GetRequiredActivityOk() (*RequiredActivity, bool)`

GetRequiredActivityOk returns a tuple with the RequiredActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredActivity

`func (o *ConfigChange) SetRequiredActivity(v RequiredActivity)`

SetRequiredActivity sets RequiredActivity field to given value.

### HasRequiredActivity

`func (o *ConfigChange) HasRequiredActivity() bool`

HasRequiredActivity returns a boolean if a field has been set.

### GetUpdateStamp

`func (o *ConfigChange) GetUpdateStamp() int64`

GetUpdateStamp returns the UpdateStamp field if non-nil, zero value otherwise.

### GetUpdateStampOk

`func (o *ConfigChange) GetUpdateStampOk() (*int64, bool)`

GetUpdateStampOk returns a tuple with the UpdateStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStamp

`func (o *ConfigChange) SetUpdateStamp(v int64)`

SetUpdateStamp sets UpdateStamp field to given value.

### HasUpdateStamp

`func (o *ConfigChange) HasUpdateStamp() bool`

HasUpdateStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


