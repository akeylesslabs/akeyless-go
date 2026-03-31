# ItemLockingSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**MaxLockTtlMins** | Pointer to **int64** |  | [optional] 

## Methods

### NewItemLockingSetting

`func NewItemLockingSetting() *ItemLockingSetting`

NewItemLockingSetting instantiates a new ItemLockingSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemLockingSettingWithDefaults

`func NewItemLockingSettingWithDefaults() *ItemLockingSetting`

NewItemLockingSettingWithDefaults instantiates a new ItemLockingSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ItemLockingSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ItemLockingSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ItemLockingSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ItemLockingSetting) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMaxLockTtlMins

`func (o *ItemLockingSetting) GetMaxLockTtlMins() int64`

GetMaxLockTtlMins returns the MaxLockTtlMins field if non-nil, zero value otherwise.

### GetMaxLockTtlMinsOk

`func (o *ItemLockingSetting) GetMaxLockTtlMinsOk() (*int64, bool)`

GetMaxLockTtlMinsOk returns a tuple with the MaxLockTtlMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLockTtlMins

`func (o *ItemLockingSetting) SetMaxLockTtlMins(v int64)`

SetMaxLockTtlMins sets MaxLockTtlMins field to given value.

### HasMaxLockTtlMins

`func (o *ItemLockingSetting) HasMaxLockTtlMins() bool`

HasMaxLockTtlMins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


