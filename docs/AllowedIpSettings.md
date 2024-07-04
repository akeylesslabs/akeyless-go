# AllowedIpSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrWhitelist** | Pointer to **string** |  | [optional] 
**Lock** | Pointer to **bool** |  | [optional] 

## Methods

### NewAllowedIpSettings

`func NewAllowedIpSettings() *AllowedIpSettings`

NewAllowedIpSettings instantiates a new AllowedIpSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedIpSettingsWithDefaults

`func NewAllowedIpSettingsWithDefaults() *AllowedIpSettings`

NewAllowedIpSettingsWithDefaults instantiates a new AllowedIpSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrWhitelist

`func (o *AllowedIpSettings) GetCidrWhitelist() string`

GetCidrWhitelist returns the CidrWhitelist field if non-nil, zero value otherwise.

### GetCidrWhitelistOk

`func (o *AllowedIpSettings) GetCidrWhitelistOk() (*string, bool)`

GetCidrWhitelistOk returns a tuple with the CidrWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrWhitelist

`func (o *AllowedIpSettings) SetCidrWhitelist(v string)`

SetCidrWhitelist sets CidrWhitelist field to given value.

### HasCidrWhitelist

`func (o *AllowedIpSettings) HasCidrWhitelist() bool`

HasCidrWhitelist returns a boolean if a field has been set.

### GetLock

`func (o *AllowedIpSettings) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *AllowedIpSettings) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *AllowedIpSettings) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *AllowedIpSettings) HasLock() bool`

HasLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


