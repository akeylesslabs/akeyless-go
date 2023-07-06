# AccountGeneralSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountDefaultKeyItemId** | Pointer to **int64** | AccountDefaultKeyItemID is the item ID of the DFC key item configured as the default protection key | [optional] 
**AccountDefaultKeyName** | Pointer to **string** | AccountDefaultKeyName is the name of the DFC key item configured as the default key This is here simply for the response to include the item name in addition to the display ID so the client can properly show this to the user. It will not be saved to the DB, only the AccountDefaultKeyItemID will. | [optional] 
**DataProtectionSection** | Pointer to [**DataProtectionSection**](DataProtectionSection.md) |  | [optional] 
**EnableRequestForAccess** | Pointer to **bool** |  | [optional] 
**PasswordPolicy** | Pointer to [**PasswordPolicyInfo**](PasswordPolicyInfo.md) |  | [optional] 
**ProtectItemsByDefault** | Pointer to **bool** |  | [optional] 
**SharingPolicy** | Pointer to [**SharingPolicyInfo**](SharingPolicyInfo.md) |  | [optional] 

## Methods

### NewAccountGeneralSettings

`func NewAccountGeneralSettings() *AccountGeneralSettings`

NewAccountGeneralSettings instantiates a new AccountGeneralSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGeneralSettingsWithDefaults

`func NewAccountGeneralSettingsWithDefaults() *AccountGeneralSettings`

NewAccountGeneralSettingsWithDefaults instantiates a new AccountGeneralSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountDefaultKeyItemId

`func (o *AccountGeneralSettings) GetAccountDefaultKeyItemId() int64`

GetAccountDefaultKeyItemId returns the AccountDefaultKeyItemId field if non-nil, zero value otherwise.

### GetAccountDefaultKeyItemIdOk

`func (o *AccountGeneralSettings) GetAccountDefaultKeyItemIdOk() (*int64, bool)`

GetAccountDefaultKeyItemIdOk returns a tuple with the AccountDefaultKeyItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDefaultKeyItemId

`func (o *AccountGeneralSettings) SetAccountDefaultKeyItemId(v int64)`

SetAccountDefaultKeyItemId sets AccountDefaultKeyItemId field to given value.

### HasAccountDefaultKeyItemId

`func (o *AccountGeneralSettings) HasAccountDefaultKeyItemId() bool`

HasAccountDefaultKeyItemId returns a boolean if a field has been set.

### GetAccountDefaultKeyName

`func (o *AccountGeneralSettings) GetAccountDefaultKeyName() string`

GetAccountDefaultKeyName returns the AccountDefaultKeyName field if non-nil, zero value otherwise.

### GetAccountDefaultKeyNameOk

`func (o *AccountGeneralSettings) GetAccountDefaultKeyNameOk() (*string, bool)`

GetAccountDefaultKeyNameOk returns a tuple with the AccountDefaultKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDefaultKeyName

`func (o *AccountGeneralSettings) SetAccountDefaultKeyName(v string)`

SetAccountDefaultKeyName sets AccountDefaultKeyName field to given value.

### HasAccountDefaultKeyName

`func (o *AccountGeneralSettings) HasAccountDefaultKeyName() bool`

HasAccountDefaultKeyName returns a boolean if a field has been set.

### GetDataProtectionSection

`func (o *AccountGeneralSettings) GetDataProtectionSection() DataProtectionSection`

GetDataProtectionSection returns the DataProtectionSection field if non-nil, zero value otherwise.

### GetDataProtectionSectionOk

`func (o *AccountGeneralSettings) GetDataProtectionSectionOk() (*DataProtectionSection, bool)`

GetDataProtectionSectionOk returns a tuple with the DataProtectionSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionSection

`func (o *AccountGeneralSettings) SetDataProtectionSection(v DataProtectionSection)`

SetDataProtectionSection sets DataProtectionSection field to given value.

### HasDataProtectionSection

`func (o *AccountGeneralSettings) HasDataProtectionSection() bool`

HasDataProtectionSection returns a boolean if a field has been set.

### GetEnableRequestForAccess

`func (o *AccountGeneralSettings) GetEnableRequestForAccess() bool`

GetEnableRequestForAccess returns the EnableRequestForAccess field if non-nil, zero value otherwise.

### GetEnableRequestForAccessOk

`func (o *AccountGeneralSettings) GetEnableRequestForAccessOk() (*bool, bool)`

GetEnableRequestForAccessOk returns a tuple with the EnableRequestForAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRequestForAccess

`func (o *AccountGeneralSettings) SetEnableRequestForAccess(v bool)`

SetEnableRequestForAccess sets EnableRequestForAccess field to given value.

### HasEnableRequestForAccess

`func (o *AccountGeneralSettings) HasEnableRequestForAccess() bool`

HasEnableRequestForAccess returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *AccountGeneralSettings) GetPasswordPolicy() PasswordPolicyInfo`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *AccountGeneralSettings) GetPasswordPolicyOk() (*PasswordPolicyInfo, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *AccountGeneralSettings) SetPasswordPolicy(v PasswordPolicyInfo)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *AccountGeneralSettings) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetProtectItemsByDefault

`func (o *AccountGeneralSettings) GetProtectItemsByDefault() bool`

GetProtectItemsByDefault returns the ProtectItemsByDefault field if non-nil, zero value otherwise.

### GetProtectItemsByDefaultOk

`func (o *AccountGeneralSettings) GetProtectItemsByDefaultOk() (*bool, bool)`

GetProtectItemsByDefaultOk returns a tuple with the ProtectItemsByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectItemsByDefault

`func (o *AccountGeneralSettings) SetProtectItemsByDefault(v bool)`

SetProtectItemsByDefault sets ProtectItemsByDefault field to given value.

### HasProtectItemsByDefault

`func (o *AccountGeneralSettings) HasProtectItemsByDefault() bool`

HasProtectItemsByDefault returns a boolean if a field has been set.

### GetSharingPolicy

`func (o *AccountGeneralSettings) GetSharingPolicy() SharingPolicyInfo`

GetSharingPolicy returns the SharingPolicy field if non-nil, zero value otherwise.

### GetSharingPolicyOk

`func (o *AccountGeneralSettings) GetSharingPolicyOk() (*SharingPolicyInfo, bool)`

GetSharingPolicyOk returns a tuple with the SharingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingPolicy

`func (o *AccountGeneralSettings) SetSharingPolicy(v SharingPolicyInfo)`

SetSharingPolicy sets SharingPolicy field to given value.

### HasSharingPolicy

`func (o *AccountGeneralSettings) HasSharingPolicy() bool`

HasSharingPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


