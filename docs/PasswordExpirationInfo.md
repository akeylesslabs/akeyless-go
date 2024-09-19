# PasswordExpirationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysBeforeExpire** | Pointer to **int64** |  | [optional] 
**DaysBeforeNotification** | Pointer to **int64** |  | [optional] 
**DaysUntilExpire** | Pointer to **int64** | r/o calculated parameter | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 

## Methods

### NewPasswordExpirationInfo

`func NewPasswordExpirationInfo() *PasswordExpirationInfo`

NewPasswordExpirationInfo instantiates a new PasswordExpirationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordExpirationInfoWithDefaults

`func NewPasswordExpirationInfoWithDefaults() *PasswordExpirationInfo`

NewPasswordExpirationInfoWithDefaults instantiates a new PasswordExpirationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysBeforeExpire

`func (o *PasswordExpirationInfo) GetDaysBeforeExpire() int64`

GetDaysBeforeExpire returns the DaysBeforeExpire field if non-nil, zero value otherwise.

### GetDaysBeforeExpireOk

`func (o *PasswordExpirationInfo) GetDaysBeforeExpireOk() (*int64, bool)`

GetDaysBeforeExpireOk returns a tuple with the DaysBeforeExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBeforeExpire

`func (o *PasswordExpirationInfo) SetDaysBeforeExpire(v int64)`

SetDaysBeforeExpire sets DaysBeforeExpire field to given value.

### HasDaysBeforeExpire

`func (o *PasswordExpirationInfo) HasDaysBeforeExpire() bool`

HasDaysBeforeExpire returns a boolean if a field has been set.

### GetDaysBeforeNotification

`func (o *PasswordExpirationInfo) GetDaysBeforeNotification() int64`

GetDaysBeforeNotification returns the DaysBeforeNotification field if non-nil, zero value otherwise.

### GetDaysBeforeNotificationOk

`func (o *PasswordExpirationInfo) GetDaysBeforeNotificationOk() (*int64, bool)`

GetDaysBeforeNotificationOk returns a tuple with the DaysBeforeNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBeforeNotification

`func (o *PasswordExpirationInfo) SetDaysBeforeNotification(v int64)`

SetDaysBeforeNotification sets DaysBeforeNotification field to given value.

### HasDaysBeforeNotification

`func (o *PasswordExpirationInfo) HasDaysBeforeNotification() bool`

HasDaysBeforeNotification returns a boolean if a field has been set.

### GetDaysUntilExpire

`func (o *PasswordExpirationInfo) GetDaysUntilExpire() int64`

GetDaysUntilExpire returns the DaysUntilExpire field if non-nil, zero value otherwise.

### GetDaysUntilExpireOk

`func (o *PasswordExpirationInfo) GetDaysUntilExpireOk() (*int64, bool)`

GetDaysUntilExpireOk returns a tuple with the DaysUntilExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilExpire

`func (o *PasswordExpirationInfo) SetDaysUntilExpire(v int64)`

SetDaysUntilExpire sets DaysUntilExpire field to given value.

### HasDaysUntilExpire

`func (o *PasswordExpirationInfo) HasDaysUntilExpire() bool`

HasDaysUntilExpire returns a boolean if a field has been set.

### GetEnable

`func (o *PasswordExpirationInfo) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PasswordExpirationInfo) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PasswordExpirationInfo) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PasswordExpirationInfo) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


