# UsageEventSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**EnableTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**IntervalByDays** | Pointer to **int64** |  | [optional] 

## Methods

### NewUsageEventSetting

`func NewUsageEventSetting() *UsageEventSetting`

NewUsageEventSetting instantiates a new UsageEventSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageEventSettingWithDefaults

`func NewUsageEventSettingWithDefaults() *UsageEventSetting`

NewUsageEventSettingWithDefaults instantiates a new UsageEventSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *UsageEventSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UsageEventSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UsageEventSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *UsageEventSetting) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTime

`func (o *UsageEventSetting) GetEnableTime() time.Time`

GetEnableTime returns the EnableTime field if non-nil, zero value otherwise.

### GetEnableTimeOk

`func (o *UsageEventSetting) GetEnableTimeOk() (*time.Time, bool)`

GetEnableTimeOk returns a tuple with the EnableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTime

`func (o *UsageEventSetting) SetEnableTime(v time.Time)`

SetEnableTime sets EnableTime field to given value.

### HasEnableTime

`func (o *UsageEventSetting) HasEnableTime() bool`

HasEnableTime returns a boolean if a field has been set.

### GetIntervalByDays

`func (o *UsageEventSetting) GetIntervalByDays() int64`

GetIntervalByDays returns the IntervalByDays field if non-nil, zero value otherwise.

### GetIntervalByDaysOk

`func (o *UsageEventSetting) GetIntervalByDaysOk() (*int64, bool)`

GetIntervalByDaysOk returns a tuple with the IntervalByDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalByDays

`func (o *UsageEventSetting) SetIntervalByDays(v int64)`

SetIntervalByDays sets IntervalByDays field to given value.

### HasIntervalByDays

`func (o *UsageEventSetting) HasIntervalByDays() bool`

HasIntervalByDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


