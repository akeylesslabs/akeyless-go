# AccountGeneralSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataProtectionSection** | Pointer to [**DataProtectionSection**](DataProtectionSection.md) |  | [optional] 
**EnableRequestForAccess** | Pointer to **bool** |  | [optional] 
**PasswordPolicy** | Pointer to [**PasswordPolicyInfo**](PasswordPolicyInfo.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


