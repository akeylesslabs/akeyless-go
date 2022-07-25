# GetAccountSettingsCommandOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**CustomerFullAddress**](CustomerFullAddress.md) |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**GeneralSettings** | Pointer to [**AccountGeneralSettings**](AccountGeneralSettings.md) |  | [optional] 
**ObjectVersionSettings** | Pointer to [**AccountObjectVersionSettingsOutput**](AccountObjectVersionSettingsOutput.md) |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**SecretManagement** | Pointer to [**SmInfo**](SmInfo.md) |  | [optional] 
**SecureRemoteAccess** | Pointer to [**SraInfo**](SraInfo.md) |  | [optional] 
**SystemAccessCredsSettings** | Pointer to [**SystemAccessCredsSettings**](SystemAccessCredsSettings.md) |  | [optional] 

## Methods

### NewGetAccountSettingsCommandOutput

`func NewGetAccountSettingsCommandOutput() *GetAccountSettingsCommandOutput`

NewGetAccountSettingsCommandOutput instantiates a new GetAccountSettingsCommandOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSettingsCommandOutputWithDefaults

`func NewGetAccountSettingsCommandOutputWithDefaults() *GetAccountSettingsCommandOutput`

NewGetAccountSettingsCommandOutputWithDefaults instantiates a new GetAccountSettingsCommandOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetAccountSettingsCommandOutput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetAccountSettingsCommandOutput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetAccountSettingsCommandOutput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetAccountSettingsCommandOutput) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAddress

`func (o *GetAccountSettingsCommandOutput) GetAddress() CustomerFullAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetAccountSettingsCommandOutput) GetAddressOk() (*CustomerFullAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetAccountSettingsCommandOutput) SetAddress(v CustomerFullAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetAccountSettingsCommandOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCompanyName

`func (o *GetAccountSettingsCommandOutput) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *GetAccountSettingsCommandOutput) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *GetAccountSettingsCommandOutput) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *GetAccountSettingsCommandOutput) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetEmail

`func (o *GetAccountSettingsCommandOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetAccountSettingsCommandOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetAccountSettingsCommandOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetAccountSettingsCommandOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGeneralSettings

`func (o *GetAccountSettingsCommandOutput) GetGeneralSettings() AccountGeneralSettings`

GetGeneralSettings returns the GeneralSettings field if non-nil, zero value otherwise.

### GetGeneralSettingsOk

`func (o *GetAccountSettingsCommandOutput) GetGeneralSettingsOk() (*AccountGeneralSettings, bool)`

GetGeneralSettingsOk returns a tuple with the GeneralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralSettings

`func (o *GetAccountSettingsCommandOutput) SetGeneralSettings(v AccountGeneralSettings)`

SetGeneralSettings sets GeneralSettings field to given value.

### HasGeneralSettings

`func (o *GetAccountSettingsCommandOutput) HasGeneralSettings() bool`

HasGeneralSettings returns a boolean if a field has been set.

### GetObjectVersionSettings

`func (o *GetAccountSettingsCommandOutput) GetObjectVersionSettings() AccountObjectVersionSettingsOutput`

GetObjectVersionSettings returns the ObjectVersionSettings field if non-nil, zero value otherwise.

### GetObjectVersionSettingsOk

`func (o *GetAccountSettingsCommandOutput) GetObjectVersionSettingsOk() (*AccountObjectVersionSettingsOutput, bool)`

GetObjectVersionSettingsOk returns a tuple with the ObjectVersionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectVersionSettings

`func (o *GetAccountSettingsCommandOutput) SetObjectVersionSettings(v AccountObjectVersionSettingsOutput)`

SetObjectVersionSettings sets ObjectVersionSettings field to given value.

### HasObjectVersionSettings

`func (o *GetAccountSettingsCommandOutput) HasObjectVersionSettings() bool`

HasObjectVersionSettings returns a boolean if a field has been set.

### GetPhone

`func (o *GetAccountSettingsCommandOutput) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetAccountSettingsCommandOutput) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetAccountSettingsCommandOutput) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetAccountSettingsCommandOutput) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSecretManagement

`func (o *GetAccountSettingsCommandOutput) GetSecretManagement() SmInfo`

GetSecretManagement returns the SecretManagement field if non-nil, zero value otherwise.

### GetSecretManagementOk

`func (o *GetAccountSettingsCommandOutput) GetSecretManagementOk() (*SmInfo, bool)`

GetSecretManagementOk returns a tuple with the SecretManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagement

`func (o *GetAccountSettingsCommandOutput) SetSecretManagement(v SmInfo)`

SetSecretManagement sets SecretManagement field to given value.

### HasSecretManagement

`func (o *GetAccountSettingsCommandOutput) HasSecretManagement() bool`

HasSecretManagement returns a boolean if a field has been set.

### GetSecureRemoteAccess

`func (o *GetAccountSettingsCommandOutput) GetSecureRemoteAccess() SraInfo`

GetSecureRemoteAccess returns the SecureRemoteAccess field if non-nil, zero value otherwise.

### GetSecureRemoteAccessOk

`func (o *GetAccountSettingsCommandOutput) GetSecureRemoteAccessOk() (*SraInfo, bool)`

GetSecureRemoteAccessOk returns a tuple with the SecureRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureRemoteAccess

`func (o *GetAccountSettingsCommandOutput) SetSecureRemoteAccess(v SraInfo)`

SetSecureRemoteAccess sets SecureRemoteAccess field to given value.

### HasSecureRemoteAccess

`func (o *GetAccountSettingsCommandOutput) HasSecureRemoteAccess() bool`

HasSecureRemoteAccess returns a boolean if a field has been set.

### GetSystemAccessCredsSettings

`func (o *GetAccountSettingsCommandOutput) GetSystemAccessCredsSettings() SystemAccessCredsSettings`

GetSystemAccessCredsSettings returns the SystemAccessCredsSettings field if non-nil, zero value otherwise.

### GetSystemAccessCredsSettingsOk

`func (o *GetAccountSettingsCommandOutput) GetSystemAccessCredsSettingsOk() (*SystemAccessCredsSettings, bool)`

GetSystemAccessCredsSettingsOk returns a tuple with the SystemAccessCredsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccessCredsSettings

`func (o *GetAccountSettingsCommandOutput) SetSystemAccessCredsSettings(v SystemAccessCredsSettings)`

SetSystemAccessCredsSettings sets SystemAccessCredsSettings field to given value.

### HasSystemAccessCredsSettings

`func (o *GetAccountSettingsCommandOutput) HasSystemAccessCredsSettings() bool`

HasSystemAccessCredsSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


