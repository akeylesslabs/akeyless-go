# AccountGeneralSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountDefaultKeyItemId** | Pointer to **int64** | AccountDefaultKeyItemID is the item ID of the DFC key item configured as the default protection key | [optional] 
**AccountDefaultKeyName** | Pointer to **string** | AccountDefaultKeyName is the name of the DFC key item configured as the default key This is here simply for the response to include the item name in addition to the display ID so the client can properly show this to the user. It will not be saved to the DB, only the AccountDefaultKeyItemID will. | [optional] 
**AiInsights** | Pointer to [**AiInsightsSetting**](AiInsightsSetting.md) |  | [optional] 
**AllowAutoFill** | Pointer to **bool** |  | [optional] 
**AllowedClientTypes** | Pointer to [**AllowedClientType**](AllowedClientType.md) |  | [optional] 
**AllowedClientsIps** | Pointer to [**AllowedIpSettings**](AllowedIpSettings.md) |  | [optional] 
**AllowedGatewaysIps** | Pointer to [**AllowedIpSettings**](AllowedIpSettings.md) |  | [optional] 
**AuthUsageEvent** | Pointer to [**UsageEventSetting**](UsageEventSetting.md) |  | [optional] 
**CertificateExpirationEvents** | Pointer to [**CertificateExpirationEventsSettings**](CertificateExpirationEventsSettings.md) |  | [optional] 
**DataProtectionSection** | Pointer to [**DataProtectionSection**](DataProtectionSection.md) |  | [optional] 
**DefaultAuthMethod** | Pointer to [**DefaultAuthMethodSettings**](DefaultAuthMethodSettings.md) |  | [optional] 
**DefaultHomePage** | Pointer to [**DefaultHomePage**](DefaultHomePage.md) |  | [optional] 
**DynamicSecretMaxTtl** | Pointer to [**DynamicSecretMaxTtl**](DynamicSecretMaxTtl.md) |  | [optional] 
**EnableRequestForAccess** | Pointer to **bool** |  | [optional] 
**HidePersonalFolder** | Pointer to **bool** |  | [optional] 
**HideStaticPassword** | Pointer to **bool** |  | [optional] 
**InvalidCharacters** | Pointer to **string** | InvalidCharacters is the invalid characters for items/targets/roles/auths/notifier_forwarder naming convention | [optional] 
**ItemUsageEvent** | Pointer to [**UsageEventSetting**](UsageEventSetting.md) |  | [optional] 
**LockDefaultKey** | Pointer to **bool** | LockDefaultKey determines whether the configured default key can be updated by end-users on a per-request basis true - all requests use the configured default key false - every request can determine its protection key (default) nil - change nothing (every request can determine its protection key (default)) This parameter is only relevant if AccountDefaultKeyItemID is not empty | [optional] 
**PasswordExpirationInfo** | Pointer to [**PasswordExpirationInfo**](PasswordExpirationInfo.md) |  | [optional] 
**PasswordPolicy** | Pointer to [**PasswordPolicyInfo**](PasswordPolicyInfo.md) |  | [optional] 
**PasswordScore** | Pointer to [**PasswordScoreSetting**](PasswordScoreSetting.md) |  | [optional] 
**ProtectItemsByDefault** | Pointer to **bool** |  | [optional] 
**RotationSecretMaxInterval** | Pointer to [**RotationSecretMaxInterval**](RotationSecretMaxInterval.md) |  | [optional] 
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

### GetAiInsights

`func (o *AccountGeneralSettings) GetAiInsights() AiInsightsSetting`

GetAiInsights returns the AiInsights field if non-nil, zero value otherwise.

### GetAiInsightsOk

`func (o *AccountGeneralSettings) GetAiInsightsOk() (*AiInsightsSetting, bool)`

GetAiInsightsOk returns a tuple with the AiInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiInsights

`func (o *AccountGeneralSettings) SetAiInsights(v AiInsightsSetting)`

SetAiInsights sets AiInsights field to given value.

### HasAiInsights

`func (o *AccountGeneralSettings) HasAiInsights() bool`

HasAiInsights returns a boolean if a field has been set.

### GetAllowAutoFill

`func (o *AccountGeneralSettings) GetAllowAutoFill() bool`

GetAllowAutoFill returns the AllowAutoFill field if non-nil, zero value otherwise.

### GetAllowAutoFillOk

`func (o *AccountGeneralSettings) GetAllowAutoFillOk() (*bool, bool)`

GetAllowAutoFillOk returns a tuple with the AllowAutoFill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoFill

`func (o *AccountGeneralSettings) SetAllowAutoFill(v bool)`

SetAllowAutoFill sets AllowAutoFill field to given value.

### HasAllowAutoFill

`func (o *AccountGeneralSettings) HasAllowAutoFill() bool`

HasAllowAutoFill returns a boolean if a field has been set.

### GetAllowedClientTypes

`func (o *AccountGeneralSettings) GetAllowedClientTypes() AllowedClientType`

GetAllowedClientTypes returns the AllowedClientTypes field if non-nil, zero value otherwise.

### GetAllowedClientTypesOk

`func (o *AccountGeneralSettings) GetAllowedClientTypesOk() (*AllowedClientType, bool)`

GetAllowedClientTypesOk returns a tuple with the AllowedClientTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientTypes

`func (o *AccountGeneralSettings) SetAllowedClientTypes(v AllowedClientType)`

SetAllowedClientTypes sets AllowedClientTypes field to given value.

### HasAllowedClientTypes

`func (o *AccountGeneralSettings) HasAllowedClientTypes() bool`

HasAllowedClientTypes returns a boolean if a field has been set.

### GetAllowedClientsIps

`func (o *AccountGeneralSettings) GetAllowedClientsIps() AllowedIpSettings`

GetAllowedClientsIps returns the AllowedClientsIps field if non-nil, zero value otherwise.

### GetAllowedClientsIpsOk

`func (o *AccountGeneralSettings) GetAllowedClientsIpsOk() (*AllowedIpSettings, bool)`

GetAllowedClientsIpsOk returns a tuple with the AllowedClientsIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientsIps

`func (o *AccountGeneralSettings) SetAllowedClientsIps(v AllowedIpSettings)`

SetAllowedClientsIps sets AllowedClientsIps field to given value.

### HasAllowedClientsIps

`func (o *AccountGeneralSettings) HasAllowedClientsIps() bool`

HasAllowedClientsIps returns a boolean if a field has been set.

### GetAllowedGatewaysIps

`func (o *AccountGeneralSettings) GetAllowedGatewaysIps() AllowedIpSettings`

GetAllowedGatewaysIps returns the AllowedGatewaysIps field if non-nil, zero value otherwise.

### GetAllowedGatewaysIpsOk

`func (o *AccountGeneralSettings) GetAllowedGatewaysIpsOk() (*AllowedIpSettings, bool)`

GetAllowedGatewaysIpsOk returns a tuple with the AllowedGatewaysIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGatewaysIps

`func (o *AccountGeneralSettings) SetAllowedGatewaysIps(v AllowedIpSettings)`

SetAllowedGatewaysIps sets AllowedGatewaysIps field to given value.

### HasAllowedGatewaysIps

`func (o *AccountGeneralSettings) HasAllowedGatewaysIps() bool`

HasAllowedGatewaysIps returns a boolean if a field has been set.

### GetAuthUsageEvent

`func (o *AccountGeneralSettings) GetAuthUsageEvent() UsageEventSetting`

GetAuthUsageEvent returns the AuthUsageEvent field if non-nil, zero value otherwise.

### GetAuthUsageEventOk

`func (o *AccountGeneralSettings) GetAuthUsageEventOk() (*UsageEventSetting, bool)`

GetAuthUsageEventOk returns a tuple with the AuthUsageEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsageEvent

`func (o *AccountGeneralSettings) SetAuthUsageEvent(v UsageEventSetting)`

SetAuthUsageEvent sets AuthUsageEvent field to given value.

### HasAuthUsageEvent

`func (o *AccountGeneralSettings) HasAuthUsageEvent() bool`

HasAuthUsageEvent returns a boolean if a field has been set.

### GetCertificateExpirationEvents

`func (o *AccountGeneralSettings) GetCertificateExpirationEvents() CertificateExpirationEventsSettings`

GetCertificateExpirationEvents returns the CertificateExpirationEvents field if non-nil, zero value otherwise.

### GetCertificateExpirationEventsOk

`func (o *AccountGeneralSettings) GetCertificateExpirationEventsOk() (*CertificateExpirationEventsSettings, bool)`

GetCertificateExpirationEventsOk returns a tuple with the CertificateExpirationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExpirationEvents

`func (o *AccountGeneralSettings) SetCertificateExpirationEvents(v CertificateExpirationEventsSettings)`

SetCertificateExpirationEvents sets CertificateExpirationEvents field to given value.

### HasCertificateExpirationEvents

`func (o *AccountGeneralSettings) HasCertificateExpirationEvents() bool`

HasCertificateExpirationEvents returns a boolean if a field has been set.

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

### GetDefaultAuthMethod

`func (o *AccountGeneralSettings) GetDefaultAuthMethod() DefaultAuthMethodSettings`

GetDefaultAuthMethod returns the DefaultAuthMethod field if non-nil, zero value otherwise.

### GetDefaultAuthMethodOk

`func (o *AccountGeneralSettings) GetDefaultAuthMethodOk() (*DefaultAuthMethodSettings, bool)`

GetDefaultAuthMethodOk returns a tuple with the DefaultAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAuthMethod

`func (o *AccountGeneralSettings) SetDefaultAuthMethod(v DefaultAuthMethodSettings)`

SetDefaultAuthMethod sets DefaultAuthMethod field to given value.

### HasDefaultAuthMethod

`func (o *AccountGeneralSettings) HasDefaultAuthMethod() bool`

HasDefaultAuthMethod returns a boolean if a field has been set.

### GetDefaultHomePage

`func (o *AccountGeneralSettings) GetDefaultHomePage() DefaultHomePage`

GetDefaultHomePage returns the DefaultHomePage field if non-nil, zero value otherwise.

### GetDefaultHomePageOk

`func (o *AccountGeneralSettings) GetDefaultHomePageOk() (*DefaultHomePage, bool)`

GetDefaultHomePageOk returns a tuple with the DefaultHomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHomePage

`func (o *AccountGeneralSettings) SetDefaultHomePage(v DefaultHomePage)`

SetDefaultHomePage sets DefaultHomePage field to given value.

### HasDefaultHomePage

`func (o *AccountGeneralSettings) HasDefaultHomePage() bool`

HasDefaultHomePage returns a boolean if a field has been set.

### GetDynamicSecretMaxTtl

`func (o *AccountGeneralSettings) GetDynamicSecretMaxTtl() DynamicSecretMaxTtl`

GetDynamicSecretMaxTtl returns the DynamicSecretMaxTtl field if non-nil, zero value otherwise.

### GetDynamicSecretMaxTtlOk

`func (o *AccountGeneralSettings) GetDynamicSecretMaxTtlOk() (*DynamicSecretMaxTtl, bool)`

GetDynamicSecretMaxTtlOk returns a tuple with the DynamicSecretMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretMaxTtl

`func (o *AccountGeneralSettings) SetDynamicSecretMaxTtl(v DynamicSecretMaxTtl)`

SetDynamicSecretMaxTtl sets DynamicSecretMaxTtl field to given value.

### HasDynamicSecretMaxTtl

`func (o *AccountGeneralSettings) HasDynamicSecretMaxTtl() bool`

HasDynamicSecretMaxTtl returns a boolean if a field has been set.

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

### GetHidePersonalFolder

`func (o *AccountGeneralSettings) GetHidePersonalFolder() bool`

GetHidePersonalFolder returns the HidePersonalFolder field if non-nil, zero value otherwise.

### GetHidePersonalFolderOk

`func (o *AccountGeneralSettings) GetHidePersonalFolderOk() (*bool, bool)`

GetHidePersonalFolderOk returns a tuple with the HidePersonalFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePersonalFolder

`func (o *AccountGeneralSettings) SetHidePersonalFolder(v bool)`

SetHidePersonalFolder sets HidePersonalFolder field to given value.

### HasHidePersonalFolder

`func (o *AccountGeneralSettings) HasHidePersonalFolder() bool`

HasHidePersonalFolder returns a boolean if a field has been set.

### GetHideStaticPassword

`func (o *AccountGeneralSettings) GetHideStaticPassword() bool`

GetHideStaticPassword returns the HideStaticPassword field if non-nil, zero value otherwise.

### GetHideStaticPasswordOk

`func (o *AccountGeneralSettings) GetHideStaticPasswordOk() (*bool, bool)`

GetHideStaticPasswordOk returns a tuple with the HideStaticPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideStaticPassword

`func (o *AccountGeneralSettings) SetHideStaticPassword(v bool)`

SetHideStaticPassword sets HideStaticPassword field to given value.

### HasHideStaticPassword

`func (o *AccountGeneralSettings) HasHideStaticPassword() bool`

HasHideStaticPassword returns a boolean if a field has been set.

### GetInvalidCharacters

`func (o *AccountGeneralSettings) GetInvalidCharacters() string`

GetInvalidCharacters returns the InvalidCharacters field if non-nil, zero value otherwise.

### GetInvalidCharactersOk

`func (o *AccountGeneralSettings) GetInvalidCharactersOk() (*string, bool)`

GetInvalidCharactersOk returns a tuple with the InvalidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCharacters

`func (o *AccountGeneralSettings) SetInvalidCharacters(v string)`

SetInvalidCharacters sets InvalidCharacters field to given value.

### HasInvalidCharacters

`func (o *AccountGeneralSettings) HasInvalidCharacters() bool`

HasInvalidCharacters returns a boolean if a field has been set.

### GetItemUsageEvent

`func (o *AccountGeneralSettings) GetItemUsageEvent() UsageEventSetting`

GetItemUsageEvent returns the ItemUsageEvent field if non-nil, zero value otherwise.

### GetItemUsageEventOk

`func (o *AccountGeneralSettings) GetItemUsageEventOk() (*UsageEventSetting, bool)`

GetItemUsageEventOk returns a tuple with the ItemUsageEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUsageEvent

`func (o *AccountGeneralSettings) SetItemUsageEvent(v UsageEventSetting)`

SetItemUsageEvent sets ItemUsageEvent field to given value.

### HasItemUsageEvent

`func (o *AccountGeneralSettings) HasItemUsageEvent() bool`

HasItemUsageEvent returns a boolean if a field has been set.

### GetLockDefaultKey

`func (o *AccountGeneralSettings) GetLockDefaultKey() bool`

GetLockDefaultKey returns the LockDefaultKey field if non-nil, zero value otherwise.

### GetLockDefaultKeyOk

`func (o *AccountGeneralSettings) GetLockDefaultKeyOk() (*bool, bool)`

GetLockDefaultKeyOk returns a tuple with the LockDefaultKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockDefaultKey

`func (o *AccountGeneralSettings) SetLockDefaultKey(v bool)`

SetLockDefaultKey sets LockDefaultKey field to given value.

### HasLockDefaultKey

`func (o *AccountGeneralSettings) HasLockDefaultKey() bool`

HasLockDefaultKey returns a boolean if a field has been set.

### GetPasswordExpirationInfo

`func (o *AccountGeneralSettings) GetPasswordExpirationInfo() PasswordExpirationInfo`

GetPasswordExpirationInfo returns the PasswordExpirationInfo field if non-nil, zero value otherwise.

### GetPasswordExpirationInfoOk

`func (o *AccountGeneralSettings) GetPasswordExpirationInfoOk() (*PasswordExpirationInfo, bool)`

GetPasswordExpirationInfoOk returns a tuple with the PasswordExpirationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationInfo

`func (o *AccountGeneralSettings) SetPasswordExpirationInfo(v PasswordExpirationInfo)`

SetPasswordExpirationInfo sets PasswordExpirationInfo field to given value.

### HasPasswordExpirationInfo

`func (o *AccountGeneralSettings) HasPasswordExpirationInfo() bool`

HasPasswordExpirationInfo returns a boolean if a field has been set.

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

### GetPasswordScore

`func (o *AccountGeneralSettings) GetPasswordScore() PasswordScoreSetting`

GetPasswordScore returns the PasswordScore field if non-nil, zero value otherwise.

### GetPasswordScoreOk

`func (o *AccountGeneralSettings) GetPasswordScoreOk() (*PasswordScoreSetting, bool)`

GetPasswordScoreOk returns a tuple with the PasswordScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordScore

`func (o *AccountGeneralSettings) SetPasswordScore(v PasswordScoreSetting)`

SetPasswordScore sets PasswordScore field to given value.

### HasPasswordScore

`func (o *AccountGeneralSettings) HasPasswordScore() bool`

HasPasswordScore returns a boolean if a field has been set.

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

### GetRotationSecretMaxInterval

`func (o *AccountGeneralSettings) GetRotationSecretMaxInterval() RotationSecretMaxInterval`

GetRotationSecretMaxInterval returns the RotationSecretMaxInterval field if non-nil, zero value otherwise.

### GetRotationSecretMaxIntervalOk

`func (o *AccountGeneralSettings) GetRotationSecretMaxIntervalOk() (*RotationSecretMaxInterval, bool)`

GetRotationSecretMaxIntervalOk returns a tuple with the RotationSecretMaxInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationSecretMaxInterval

`func (o *AccountGeneralSettings) SetRotationSecretMaxInterval(v RotationSecretMaxInterval)`

SetRotationSecretMaxInterval sets RotationSecretMaxInterval field to given value.

### HasRotationSecretMaxInterval

`func (o *AccountGeneralSettings) HasRotationSecretMaxInterval() bool`

HasRotationSecretMaxInterval returns a boolean if a field has been set.

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


