# UpdateAccountSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address | [optional] 
**AllowedClientType** | Pointer to **[]string** | A default list of client types that are allowed to authenticate [cli,ui,gateway-admin,sdk,mobile,extension]. | [optional] 
**AllowedEmailDomains** | Pointer to **[]string** | Limits email sharing to the specified domains. Relevant only when item sharing is enabled. By default, all domains are allowed. | [optional] 
**BoundIps** | Pointer to **[]string** | A default list of comma-separated CIDR block that are allowed to authenticate. | [optional] 
**City** | Pointer to **string** | City | [optional] 
**CompanyName** | Pointer to **string** | Company name | [optional] 
**Country** | Pointer to **string** | Country | [optional] 
**DefaultCertificateExpirationNotificationDays** | Pointer to **[]string** | How many days before the expiration of the certificate would you like to be notified. To specify multiple events, use argument multiple times: --default-certificate-expiration-notification-days 1 --default-certificate-expiration-notification-days 5 | [optional] 
**DefaultKeyName** | Pointer to **string** | Set the account default key based on the DFC key name. Use \&quot;set-original-akeyless-default-key\&quot; to revert to using the original default key of the account. | [optional] 
**DefaultShareLinkTtlMinutes** | Pointer to **string** | Set the default ttl in minutes for sharing item number between 60 and 43200 | [optional] 
**DefaultVersioning** | Pointer to **string** | If set to true, new versions is enabled by default | [optional] 
**DpEnableClassicKeyProtection** | Pointer to **string** | Set to update protection with classic keys state [true/false] | [optional] 
**DynamicSecretMaxTtl** | Pointer to **int64** | Set the maximum ttl for dynamic secrets | [optional] 
**DynamicSecretMaxTtlEnable** | Pointer to **string** | Set a maximum ttl for dynamic secrets [true/false] | [optional] 
**EnableAiInsights** | Pointer to **string** | Enable AI insights [true/false] | [optional] 
**EnableDefaultCertificateExpirationEvent** | Pointer to **string** | How many days before the expiration of the certificate would you like to be notified. [true/false] | [optional] 
**EnableItemSharing** | Pointer to **string** | Enable sharing items [true/false] | [optional] 
**EnablePasswordExpiration** | Pointer to **string** | Enable password expiration policy [true/false] | [optional] 
**ForceNewVersions** | Pointer to **string** | If set to true, new version will be created on update | [optional] 
**GwBoundIps** | Pointer to **[]string** | A default list of comma-separated CIDR block that acts as a trusted Gateway entity. | [optional] 
**HidePersonalFolder** | Pointer to **string** | Hide personal folder, if set - users will not be able to use personal folder [true/false] | [optional] 
**HideStaticPassword** | Pointer to **string** | Hide static secret&#39;s password type [true/false] | [optional] 
**InvalidCharacters** | Pointer to **string** | Characters that cannot be used for items/targets/roles/auths/event_forwarder names. Empty string will enforce nothing. | [optional] [default to "notReceivedInvalidCharacter"]
**ItemType** | Pointer to **string** | VersionSettingsObjectType defines object types for account version settings | [optional] 
**ItemsDeletionProtection** | Pointer to **string** | Set or unset the default behaviour of items deletion protection [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtlDefault** | Pointer to **int64** | Default ttl | [optional] 
**JwtTtlMax** | Pointer to **int64** | Maximum ttl | [optional] 
**JwtTtlMin** | Pointer to **int64** | Minimum ttl | [optional] 
**LockAllowedClientType** | Pointer to **string** | Lock allowed-client-type setting in the account [true/false]. | [optional] 
**LockBoundIps** | Pointer to **string** | Lock bound-ips setting globally in the account. | [optional] 
**LockDefaultKey** | Pointer to **string** | Lock the account&#39;s default protection key, if set - users will not be able to use a different protection key, relevant only if default-key-name is configured [true/false] | [optional] 
**LockGwBoundIps** | Pointer to **string** | Lock gw-bound-ips setting in the account. | [optional] 
**MaxRotationInterval** | Pointer to **int32** | Set the maximum rotation interval for rotated secrets auto rotation settings | [optional] 
**MaxRotationIntervalEnable** | Pointer to **string** | Set a maximum rotation interval for rotated secrets auto rotation settings [true/false] | [optional] 
**MaxVersions** | Pointer to **string** | Max versions | [optional] 
**PasswordExpirationDays** | Pointer to **string** | Specifies the number of days that a password is valid before it must be changed. A default value of 90 days is used. | [optional] 
**PasswordExpirationNotificationDays** | Pointer to **string** | Specifies the number of days before a user receives notification that their password will expire. A default value of 14 days is used. | [optional] 
**PasswordLength** | Pointer to **int64** | Password length between 5 - to 50 characters | [optional] 
**Phone** | Pointer to **string** | Phone number | [optional] 
**PostalCode** | Pointer to **string** | Postal code | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UsageEventEnable** | Pointer to **string** | Enable event for objects that have not been used or changed [true/false] | [optional] 
**UsageEventInterval** | Pointer to **int64** | Interval by days for unused objects. Default and minimum interval is 90 days | [optional] 
**UsageEventObjectType** | Pointer to **string** | Usage event is supported for auth method or secrets-and-keys [auth/item] | [optional] 
**UseCapitalLetters** | Pointer to **string** | Password must contain capital letters [true/false] | [optional] 
**UseLowerLetters** | Pointer to **string** | Password must contain lower case letters [true/false] | [optional] 
**UseNumbers** | Pointer to **string** | Password must contain numbers [true/false] | [optional] 
**UseSpecialCharacters** | Pointer to **string** | Password must contain special characters [true/false] | [optional] 

## Methods

### NewUpdateAccountSettings

`func NewUpdateAccountSettings() *UpdateAccountSettings`

NewUpdateAccountSettings instantiates a new UpdateAccountSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountSettingsWithDefaults

`func NewUpdateAccountSettingsWithDefaults() *UpdateAccountSettings`

NewUpdateAccountSettingsWithDefaults instantiates a new UpdateAccountSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdateAccountSettings) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateAccountSettings) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateAccountSettings) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateAccountSettings) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAllowedClientType

`func (o *UpdateAccountSettings) GetAllowedClientType() []string`

GetAllowedClientType returns the AllowedClientType field if non-nil, zero value otherwise.

### GetAllowedClientTypeOk

`func (o *UpdateAccountSettings) GetAllowedClientTypeOk() (*[]string, bool)`

GetAllowedClientTypeOk returns a tuple with the AllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientType

`func (o *UpdateAccountSettings) SetAllowedClientType(v []string)`

SetAllowedClientType sets AllowedClientType field to given value.

### HasAllowedClientType

`func (o *UpdateAccountSettings) HasAllowedClientType() bool`

HasAllowedClientType returns a boolean if a field has been set.

### GetAllowedEmailDomains

`func (o *UpdateAccountSettings) GetAllowedEmailDomains() []string`

GetAllowedEmailDomains returns the AllowedEmailDomains field if non-nil, zero value otherwise.

### GetAllowedEmailDomainsOk

`func (o *UpdateAccountSettings) GetAllowedEmailDomainsOk() (*[]string, bool)`

GetAllowedEmailDomainsOk returns a tuple with the AllowedEmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEmailDomains

`func (o *UpdateAccountSettings) SetAllowedEmailDomains(v []string)`

SetAllowedEmailDomains sets AllowedEmailDomains field to given value.

### HasAllowedEmailDomains

`func (o *UpdateAccountSettings) HasAllowedEmailDomains() bool`

HasAllowedEmailDomains returns a boolean if a field has been set.

### GetBoundIps

`func (o *UpdateAccountSettings) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAccountSettings) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAccountSettings) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAccountSettings) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetCity

`func (o *UpdateAccountSettings) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateAccountSettings) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateAccountSettings) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UpdateAccountSettings) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *UpdateAccountSettings) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UpdateAccountSettings) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UpdateAccountSettings) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UpdateAccountSettings) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *UpdateAccountSettings) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateAccountSettings) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateAccountSettings) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateAccountSettings) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDefaultCertificateExpirationNotificationDays

`func (o *UpdateAccountSettings) GetDefaultCertificateExpirationNotificationDays() []string`

GetDefaultCertificateExpirationNotificationDays returns the DefaultCertificateExpirationNotificationDays field if non-nil, zero value otherwise.

### GetDefaultCertificateExpirationNotificationDaysOk

`func (o *UpdateAccountSettings) GetDefaultCertificateExpirationNotificationDaysOk() (*[]string, bool)`

GetDefaultCertificateExpirationNotificationDaysOk returns a tuple with the DefaultCertificateExpirationNotificationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertificateExpirationNotificationDays

`func (o *UpdateAccountSettings) SetDefaultCertificateExpirationNotificationDays(v []string)`

SetDefaultCertificateExpirationNotificationDays sets DefaultCertificateExpirationNotificationDays field to given value.

### HasDefaultCertificateExpirationNotificationDays

`func (o *UpdateAccountSettings) HasDefaultCertificateExpirationNotificationDays() bool`

HasDefaultCertificateExpirationNotificationDays returns a boolean if a field has been set.

### GetDefaultKeyName

`func (o *UpdateAccountSettings) GetDefaultKeyName() string`

GetDefaultKeyName returns the DefaultKeyName field if non-nil, zero value otherwise.

### GetDefaultKeyNameOk

`func (o *UpdateAccountSettings) GetDefaultKeyNameOk() (*string, bool)`

GetDefaultKeyNameOk returns a tuple with the DefaultKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeyName

`func (o *UpdateAccountSettings) SetDefaultKeyName(v string)`

SetDefaultKeyName sets DefaultKeyName field to given value.

### HasDefaultKeyName

`func (o *UpdateAccountSettings) HasDefaultKeyName() bool`

HasDefaultKeyName returns a boolean if a field has been set.

### GetDefaultShareLinkTtlMinutes

`func (o *UpdateAccountSettings) GetDefaultShareLinkTtlMinutes() string`

GetDefaultShareLinkTtlMinutes returns the DefaultShareLinkTtlMinutes field if non-nil, zero value otherwise.

### GetDefaultShareLinkTtlMinutesOk

`func (o *UpdateAccountSettings) GetDefaultShareLinkTtlMinutesOk() (*string, bool)`

GetDefaultShareLinkTtlMinutesOk returns a tuple with the DefaultShareLinkTtlMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShareLinkTtlMinutes

`func (o *UpdateAccountSettings) SetDefaultShareLinkTtlMinutes(v string)`

SetDefaultShareLinkTtlMinutes sets DefaultShareLinkTtlMinutes field to given value.

### HasDefaultShareLinkTtlMinutes

`func (o *UpdateAccountSettings) HasDefaultShareLinkTtlMinutes() bool`

HasDefaultShareLinkTtlMinutes returns a boolean if a field has been set.

### GetDefaultVersioning

`func (o *UpdateAccountSettings) GetDefaultVersioning() string`

GetDefaultVersioning returns the DefaultVersioning field if non-nil, zero value otherwise.

### GetDefaultVersioningOk

`func (o *UpdateAccountSettings) GetDefaultVersioningOk() (*string, bool)`

GetDefaultVersioningOk returns a tuple with the DefaultVersioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersioning

`func (o *UpdateAccountSettings) SetDefaultVersioning(v string)`

SetDefaultVersioning sets DefaultVersioning field to given value.

### HasDefaultVersioning

`func (o *UpdateAccountSettings) HasDefaultVersioning() bool`

HasDefaultVersioning returns a boolean if a field has been set.

### GetDpEnableClassicKeyProtection

`func (o *UpdateAccountSettings) GetDpEnableClassicKeyProtection() string`

GetDpEnableClassicKeyProtection returns the DpEnableClassicKeyProtection field if non-nil, zero value otherwise.

### GetDpEnableClassicKeyProtectionOk

`func (o *UpdateAccountSettings) GetDpEnableClassicKeyProtectionOk() (*string, bool)`

GetDpEnableClassicKeyProtectionOk returns a tuple with the DpEnableClassicKeyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpEnableClassicKeyProtection

`func (o *UpdateAccountSettings) SetDpEnableClassicKeyProtection(v string)`

SetDpEnableClassicKeyProtection sets DpEnableClassicKeyProtection field to given value.

### HasDpEnableClassicKeyProtection

`func (o *UpdateAccountSettings) HasDpEnableClassicKeyProtection() bool`

HasDpEnableClassicKeyProtection returns a boolean if a field has been set.

### GetDynamicSecretMaxTtl

`func (o *UpdateAccountSettings) GetDynamicSecretMaxTtl() int64`

GetDynamicSecretMaxTtl returns the DynamicSecretMaxTtl field if non-nil, zero value otherwise.

### GetDynamicSecretMaxTtlOk

`func (o *UpdateAccountSettings) GetDynamicSecretMaxTtlOk() (*int64, bool)`

GetDynamicSecretMaxTtlOk returns a tuple with the DynamicSecretMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretMaxTtl

`func (o *UpdateAccountSettings) SetDynamicSecretMaxTtl(v int64)`

SetDynamicSecretMaxTtl sets DynamicSecretMaxTtl field to given value.

### HasDynamicSecretMaxTtl

`func (o *UpdateAccountSettings) HasDynamicSecretMaxTtl() bool`

HasDynamicSecretMaxTtl returns a boolean if a field has been set.

### GetDynamicSecretMaxTtlEnable

`func (o *UpdateAccountSettings) GetDynamicSecretMaxTtlEnable() string`

GetDynamicSecretMaxTtlEnable returns the DynamicSecretMaxTtlEnable field if non-nil, zero value otherwise.

### GetDynamicSecretMaxTtlEnableOk

`func (o *UpdateAccountSettings) GetDynamicSecretMaxTtlEnableOk() (*string, bool)`

GetDynamicSecretMaxTtlEnableOk returns a tuple with the DynamicSecretMaxTtlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretMaxTtlEnable

`func (o *UpdateAccountSettings) SetDynamicSecretMaxTtlEnable(v string)`

SetDynamicSecretMaxTtlEnable sets DynamicSecretMaxTtlEnable field to given value.

### HasDynamicSecretMaxTtlEnable

`func (o *UpdateAccountSettings) HasDynamicSecretMaxTtlEnable() bool`

HasDynamicSecretMaxTtlEnable returns a boolean if a field has been set.

### GetEnableAiInsights

`func (o *UpdateAccountSettings) GetEnableAiInsights() string`

GetEnableAiInsights returns the EnableAiInsights field if non-nil, zero value otherwise.

### GetEnableAiInsightsOk

`func (o *UpdateAccountSettings) GetEnableAiInsightsOk() (*string, bool)`

GetEnableAiInsightsOk returns a tuple with the EnableAiInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiInsights

`func (o *UpdateAccountSettings) SetEnableAiInsights(v string)`

SetEnableAiInsights sets EnableAiInsights field to given value.

### HasEnableAiInsights

`func (o *UpdateAccountSettings) HasEnableAiInsights() bool`

HasEnableAiInsights returns a boolean if a field has been set.

### GetEnableDefaultCertificateExpirationEvent

`func (o *UpdateAccountSettings) GetEnableDefaultCertificateExpirationEvent() string`

GetEnableDefaultCertificateExpirationEvent returns the EnableDefaultCertificateExpirationEvent field if non-nil, zero value otherwise.

### GetEnableDefaultCertificateExpirationEventOk

`func (o *UpdateAccountSettings) GetEnableDefaultCertificateExpirationEventOk() (*string, bool)`

GetEnableDefaultCertificateExpirationEventOk returns a tuple with the EnableDefaultCertificateExpirationEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDefaultCertificateExpirationEvent

`func (o *UpdateAccountSettings) SetEnableDefaultCertificateExpirationEvent(v string)`

SetEnableDefaultCertificateExpirationEvent sets EnableDefaultCertificateExpirationEvent field to given value.

### HasEnableDefaultCertificateExpirationEvent

`func (o *UpdateAccountSettings) HasEnableDefaultCertificateExpirationEvent() bool`

HasEnableDefaultCertificateExpirationEvent returns a boolean if a field has been set.

### GetEnableItemSharing

`func (o *UpdateAccountSettings) GetEnableItemSharing() string`

GetEnableItemSharing returns the EnableItemSharing field if non-nil, zero value otherwise.

### GetEnableItemSharingOk

`func (o *UpdateAccountSettings) GetEnableItemSharingOk() (*string, bool)`

GetEnableItemSharingOk returns a tuple with the EnableItemSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableItemSharing

`func (o *UpdateAccountSettings) SetEnableItemSharing(v string)`

SetEnableItemSharing sets EnableItemSharing field to given value.

### HasEnableItemSharing

`func (o *UpdateAccountSettings) HasEnableItemSharing() bool`

HasEnableItemSharing returns a boolean if a field has been set.

### GetEnablePasswordExpiration

`func (o *UpdateAccountSettings) GetEnablePasswordExpiration() string`

GetEnablePasswordExpiration returns the EnablePasswordExpiration field if non-nil, zero value otherwise.

### GetEnablePasswordExpirationOk

`func (o *UpdateAccountSettings) GetEnablePasswordExpirationOk() (*string, bool)`

GetEnablePasswordExpirationOk returns a tuple with the EnablePasswordExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordExpiration

`func (o *UpdateAccountSettings) SetEnablePasswordExpiration(v string)`

SetEnablePasswordExpiration sets EnablePasswordExpiration field to given value.

### HasEnablePasswordExpiration

`func (o *UpdateAccountSettings) HasEnablePasswordExpiration() bool`

HasEnablePasswordExpiration returns a boolean if a field has been set.

### GetForceNewVersions

`func (o *UpdateAccountSettings) GetForceNewVersions() string`

GetForceNewVersions returns the ForceNewVersions field if non-nil, zero value otherwise.

### GetForceNewVersionsOk

`func (o *UpdateAccountSettings) GetForceNewVersionsOk() (*string, bool)`

GetForceNewVersionsOk returns a tuple with the ForceNewVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNewVersions

`func (o *UpdateAccountSettings) SetForceNewVersions(v string)`

SetForceNewVersions sets ForceNewVersions field to given value.

### HasForceNewVersions

`func (o *UpdateAccountSettings) HasForceNewVersions() bool`

HasForceNewVersions returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *UpdateAccountSettings) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *UpdateAccountSettings) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *UpdateAccountSettings) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *UpdateAccountSettings) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetHidePersonalFolder

`func (o *UpdateAccountSettings) GetHidePersonalFolder() string`

GetHidePersonalFolder returns the HidePersonalFolder field if non-nil, zero value otherwise.

### GetHidePersonalFolderOk

`func (o *UpdateAccountSettings) GetHidePersonalFolderOk() (*string, bool)`

GetHidePersonalFolderOk returns a tuple with the HidePersonalFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePersonalFolder

`func (o *UpdateAccountSettings) SetHidePersonalFolder(v string)`

SetHidePersonalFolder sets HidePersonalFolder field to given value.

### HasHidePersonalFolder

`func (o *UpdateAccountSettings) HasHidePersonalFolder() bool`

HasHidePersonalFolder returns a boolean if a field has been set.

### GetHideStaticPassword

`func (o *UpdateAccountSettings) GetHideStaticPassword() string`

GetHideStaticPassword returns the HideStaticPassword field if non-nil, zero value otherwise.

### GetHideStaticPasswordOk

`func (o *UpdateAccountSettings) GetHideStaticPasswordOk() (*string, bool)`

GetHideStaticPasswordOk returns a tuple with the HideStaticPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideStaticPassword

`func (o *UpdateAccountSettings) SetHideStaticPassword(v string)`

SetHideStaticPassword sets HideStaticPassword field to given value.

### HasHideStaticPassword

`func (o *UpdateAccountSettings) HasHideStaticPassword() bool`

HasHideStaticPassword returns a boolean if a field has been set.

### GetInvalidCharacters

`func (o *UpdateAccountSettings) GetInvalidCharacters() string`

GetInvalidCharacters returns the InvalidCharacters field if non-nil, zero value otherwise.

### GetInvalidCharactersOk

`func (o *UpdateAccountSettings) GetInvalidCharactersOk() (*string, bool)`

GetInvalidCharactersOk returns a tuple with the InvalidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCharacters

`func (o *UpdateAccountSettings) SetInvalidCharacters(v string)`

SetInvalidCharacters sets InvalidCharacters field to given value.

### HasInvalidCharacters

`func (o *UpdateAccountSettings) HasInvalidCharacters() bool`

HasInvalidCharacters returns a boolean if a field has been set.

### GetItemType

`func (o *UpdateAccountSettings) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *UpdateAccountSettings) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *UpdateAccountSettings) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *UpdateAccountSettings) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetItemsDeletionProtection

`func (o *UpdateAccountSettings) GetItemsDeletionProtection() string`

GetItemsDeletionProtection returns the ItemsDeletionProtection field if non-nil, zero value otherwise.

### GetItemsDeletionProtectionOk

`func (o *UpdateAccountSettings) GetItemsDeletionProtectionOk() (*string, bool)`

GetItemsDeletionProtectionOk returns a tuple with the ItemsDeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsDeletionProtection

`func (o *UpdateAccountSettings) SetItemsDeletionProtection(v string)`

SetItemsDeletionProtection sets ItemsDeletionProtection field to given value.

### HasItemsDeletionProtection

`func (o *UpdateAccountSettings) HasItemsDeletionProtection() bool`

HasItemsDeletionProtection returns a boolean if a field has been set.

### GetJson

`func (o *UpdateAccountSettings) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateAccountSettings) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateAccountSettings) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateAccountSettings) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtlDefault

`func (o *UpdateAccountSettings) GetJwtTtlDefault() int64`

GetJwtTtlDefault returns the JwtTtlDefault field if non-nil, zero value otherwise.

### GetJwtTtlDefaultOk

`func (o *UpdateAccountSettings) GetJwtTtlDefaultOk() (*int64, bool)`

GetJwtTtlDefaultOk returns a tuple with the JwtTtlDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtlDefault

`func (o *UpdateAccountSettings) SetJwtTtlDefault(v int64)`

SetJwtTtlDefault sets JwtTtlDefault field to given value.

### HasJwtTtlDefault

`func (o *UpdateAccountSettings) HasJwtTtlDefault() bool`

HasJwtTtlDefault returns a boolean if a field has been set.

### GetJwtTtlMax

`func (o *UpdateAccountSettings) GetJwtTtlMax() int64`

GetJwtTtlMax returns the JwtTtlMax field if non-nil, zero value otherwise.

### GetJwtTtlMaxOk

`func (o *UpdateAccountSettings) GetJwtTtlMaxOk() (*int64, bool)`

GetJwtTtlMaxOk returns a tuple with the JwtTtlMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtlMax

`func (o *UpdateAccountSettings) SetJwtTtlMax(v int64)`

SetJwtTtlMax sets JwtTtlMax field to given value.

### HasJwtTtlMax

`func (o *UpdateAccountSettings) HasJwtTtlMax() bool`

HasJwtTtlMax returns a boolean if a field has been set.

### GetJwtTtlMin

`func (o *UpdateAccountSettings) GetJwtTtlMin() int64`

GetJwtTtlMin returns the JwtTtlMin field if non-nil, zero value otherwise.

### GetJwtTtlMinOk

`func (o *UpdateAccountSettings) GetJwtTtlMinOk() (*int64, bool)`

GetJwtTtlMinOk returns a tuple with the JwtTtlMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtlMin

`func (o *UpdateAccountSettings) SetJwtTtlMin(v int64)`

SetJwtTtlMin sets JwtTtlMin field to given value.

### HasJwtTtlMin

`func (o *UpdateAccountSettings) HasJwtTtlMin() bool`

HasJwtTtlMin returns a boolean if a field has been set.

### GetLockAllowedClientType

`func (o *UpdateAccountSettings) GetLockAllowedClientType() string`

GetLockAllowedClientType returns the LockAllowedClientType field if non-nil, zero value otherwise.

### GetLockAllowedClientTypeOk

`func (o *UpdateAccountSettings) GetLockAllowedClientTypeOk() (*string, bool)`

GetLockAllowedClientTypeOk returns a tuple with the LockAllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockAllowedClientType

`func (o *UpdateAccountSettings) SetLockAllowedClientType(v string)`

SetLockAllowedClientType sets LockAllowedClientType field to given value.

### HasLockAllowedClientType

`func (o *UpdateAccountSettings) HasLockAllowedClientType() bool`

HasLockAllowedClientType returns a boolean if a field has been set.

### GetLockBoundIps

`func (o *UpdateAccountSettings) GetLockBoundIps() string`

GetLockBoundIps returns the LockBoundIps field if non-nil, zero value otherwise.

### GetLockBoundIpsOk

`func (o *UpdateAccountSettings) GetLockBoundIpsOk() (*string, bool)`

GetLockBoundIpsOk returns a tuple with the LockBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockBoundIps

`func (o *UpdateAccountSettings) SetLockBoundIps(v string)`

SetLockBoundIps sets LockBoundIps field to given value.

### HasLockBoundIps

`func (o *UpdateAccountSettings) HasLockBoundIps() bool`

HasLockBoundIps returns a boolean if a field has been set.

### GetLockDefaultKey

`func (o *UpdateAccountSettings) GetLockDefaultKey() string`

GetLockDefaultKey returns the LockDefaultKey field if non-nil, zero value otherwise.

### GetLockDefaultKeyOk

`func (o *UpdateAccountSettings) GetLockDefaultKeyOk() (*string, bool)`

GetLockDefaultKeyOk returns a tuple with the LockDefaultKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockDefaultKey

`func (o *UpdateAccountSettings) SetLockDefaultKey(v string)`

SetLockDefaultKey sets LockDefaultKey field to given value.

### HasLockDefaultKey

`func (o *UpdateAccountSettings) HasLockDefaultKey() bool`

HasLockDefaultKey returns a boolean if a field has been set.

### GetLockGwBoundIps

`func (o *UpdateAccountSettings) GetLockGwBoundIps() string`

GetLockGwBoundIps returns the LockGwBoundIps field if non-nil, zero value otherwise.

### GetLockGwBoundIpsOk

`func (o *UpdateAccountSettings) GetLockGwBoundIpsOk() (*string, bool)`

GetLockGwBoundIpsOk returns a tuple with the LockGwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockGwBoundIps

`func (o *UpdateAccountSettings) SetLockGwBoundIps(v string)`

SetLockGwBoundIps sets LockGwBoundIps field to given value.

### HasLockGwBoundIps

`func (o *UpdateAccountSettings) HasLockGwBoundIps() bool`

HasLockGwBoundIps returns a boolean if a field has been set.

### GetMaxRotationInterval

`func (o *UpdateAccountSettings) GetMaxRotationInterval() int32`

GetMaxRotationInterval returns the MaxRotationInterval field if non-nil, zero value otherwise.

### GetMaxRotationIntervalOk

`func (o *UpdateAccountSettings) GetMaxRotationIntervalOk() (*int32, bool)`

GetMaxRotationIntervalOk returns a tuple with the MaxRotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRotationInterval

`func (o *UpdateAccountSettings) SetMaxRotationInterval(v int32)`

SetMaxRotationInterval sets MaxRotationInterval field to given value.

### HasMaxRotationInterval

`func (o *UpdateAccountSettings) HasMaxRotationInterval() bool`

HasMaxRotationInterval returns a boolean if a field has been set.

### GetMaxRotationIntervalEnable

`func (o *UpdateAccountSettings) GetMaxRotationIntervalEnable() string`

GetMaxRotationIntervalEnable returns the MaxRotationIntervalEnable field if non-nil, zero value otherwise.

### GetMaxRotationIntervalEnableOk

`func (o *UpdateAccountSettings) GetMaxRotationIntervalEnableOk() (*string, bool)`

GetMaxRotationIntervalEnableOk returns a tuple with the MaxRotationIntervalEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRotationIntervalEnable

`func (o *UpdateAccountSettings) SetMaxRotationIntervalEnable(v string)`

SetMaxRotationIntervalEnable sets MaxRotationIntervalEnable field to given value.

### HasMaxRotationIntervalEnable

`func (o *UpdateAccountSettings) HasMaxRotationIntervalEnable() bool`

HasMaxRotationIntervalEnable returns a boolean if a field has been set.

### GetMaxVersions

`func (o *UpdateAccountSettings) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *UpdateAccountSettings) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *UpdateAccountSettings) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *UpdateAccountSettings) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetPasswordExpirationDays

`func (o *UpdateAccountSettings) GetPasswordExpirationDays() string`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *UpdateAccountSettings) GetPasswordExpirationDaysOk() (*string, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationDays

`func (o *UpdateAccountSettings) SetPasswordExpirationDays(v string)`

SetPasswordExpirationDays sets PasswordExpirationDays field to given value.

### HasPasswordExpirationDays

`func (o *UpdateAccountSettings) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### GetPasswordExpirationNotificationDays

`func (o *UpdateAccountSettings) GetPasswordExpirationNotificationDays() string`

GetPasswordExpirationNotificationDays returns the PasswordExpirationNotificationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationNotificationDaysOk

`func (o *UpdateAccountSettings) GetPasswordExpirationNotificationDaysOk() (*string, bool)`

GetPasswordExpirationNotificationDaysOk returns a tuple with the PasswordExpirationNotificationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationNotificationDays

`func (o *UpdateAccountSettings) SetPasswordExpirationNotificationDays(v string)`

SetPasswordExpirationNotificationDays sets PasswordExpirationNotificationDays field to given value.

### HasPasswordExpirationNotificationDays

`func (o *UpdateAccountSettings) HasPasswordExpirationNotificationDays() bool`

HasPasswordExpirationNotificationDays returns a boolean if a field has been set.

### GetPasswordLength

`func (o *UpdateAccountSettings) GetPasswordLength() int64`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *UpdateAccountSettings) GetPasswordLengthOk() (*int64, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *UpdateAccountSettings) SetPasswordLength(v int64)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *UpdateAccountSettings) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateAccountSettings) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateAccountSettings) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateAccountSettings) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateAccountSettings) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *UpdateAccountSettings) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UpdateAccountSettings) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UpdateAccountSettings) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UpdateAccountSettings) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAccountSettings) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAccountSettings) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAccountSettings) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAccountSettings) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAccountSettings) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAccountSettings) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAccountSettings) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAccountSettings) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsageEventEnable

`func (o *UpdateAccountSettings) GetUsageEventEnable() string`

GetUsageEventEnable returns the UsageEventEnable field if non-nil, zero value otherwise.

### GetUsageEventEnableOk

`func (o *UpdateAccountSettings) GetUsageEventEnableOk() (*string, bool)`

GetUsageEventEnableOk returns a tuple with the UsageEventEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventEnable

`func (o *UpdateAccountSettings) SetUsageEventEnable(v string)`

SetUsageEventEnable sets UsageEventEnable field to given value.

### HasUsageEventEnable

`func (o *UpdateAccountSettings) HasUsageEventEnable() bool`

HasUsageEventEnable returns a boolean if a field has been set.

### GetUsageEventInterval

`func (o *UpdateAccountSettings) GetUsageEventInterval() int64`

GetUsageEventInterval returns the UsageEventInterval field if non-nil, zero value otherwise.

### GetUsageEventIntervalOk

`func (o *UpdateAccountSettings) GetUsageEventIntervalOk() (*int64, bool)`

GetUsageEventIntervalOk returns a tuple with the UsageEventInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventInterval

`func (o *UpdateAccountSettings) SetUsageEventInterval(v int64)`

SetUsageEventInterval sets UsageEventInterval field to given value.

### HasUsageEventInterval

`func (o *UpdateAccountSettings) HasUsageEventInterval() bool`

HasUsageEventInterval returns a boolean if a field has been set.

### GetUsageEventObjectType

`func (o *UpdateAccountSettings) GetUsageEventObjectType() string`

GetUsageEventObjectType returns the UsageEventObjectType field if non-nil, zero value otherwise.

### GetUsageEventObjectTypeOk

`func (o *UpdateAccountSettings) GetUsageEventObjectTypeOk() (*string, bool)`

GetUsageEventObjectTypeOk returns a tuple with the UsageEventObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventObjectType

`func (o *UpdateAccountSettings) SetUsageEventObjectType(v string)`

SetUsageEventObjectType sets UsageEventObjectType field to given value.

### HasUsageEventObjectType

`func (o *UpdateAccountSettings) HasUsageEventObjectType() bool`

HasUsageEventObjectType returns a boolean if a field has been set.

### GetUseCapitalLetters

`func (o *UpdateAccountSettings) GetUseCapitalLetters() string`

GetUseCapitalLetters returns the UseCapitalLetters field if non-nil, zero value otherwise.

### GetUseCapitalLettersOk

`func (o *UpdateAccountSettings) GetUseCapitalLettersOk() (*string, bool)`

GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCapitalLetters

`func (o *UpdateAccountSettings) SetUseCapitalLetters(v string)`

SetUseCapitalLetters sets UseCapitalLetters field to given value.

### HasUseCapitalLetters

`func (o *UpdateAccountSettings) HasUseCapitalLetters() bool`

HasUseCapitalLetters returns a boolean if a field has been set.

### GetUseLowerLetters

`func (o *UpdateAccountSettings) GetUseLowerLetters() string`

GetUseLowerLetters returns the UseLowerLetters field if non-nil, zero value otherwise.

### GetUseLowerLettersOk

`func (o *UpdateAccountSettings) GetUseLowerLettersOk() (*string, bool)`

GetUseLowerLettersOk returns a tuple with the UseLowerLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLowerLetters

`func (o *UpdateAccountSettings) SetUseLowerLetters(v string)`

SetUseLowerLetters sets UseLowerLetters field to given value.

### HasUseLowerLetters

`func (o *UpdateAccountSettings) HasUseLowerLetters() bool`

HasUseLowerLetters returns a boolean if a field has been set.

### GetUseNumbers

`func (o *UpdateAccountSettings) GetUseNumbers() string`

GetUseNumbers returns the UseNumbers field if non-nil, zero value otherwise.

### GetUseNumbersOk

`func (o *UpdateAccountSettings) GetUseNumbersOk() (*string, bool)`

GetUseNumbersOk returns a tuple with the UseNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumbers

`func (o *UpdateAccountSettings) SetUseNumbers(v string)`

SetUseNumbers sets UseNumbers field to given value.

### HasUseNumbers

`func (o *UpdateAccountSettings) HasUseNumbers() bool`

HasUseNumbers returns a boolean if a field has been set.

### GetUseSpecialCharacters

`func (o *UpdateAccountSettings) GetUseSpecialCharacters() string`

GetUseSpecialCharacters returns the UseSpecialCharacters field if non-nil, zero value otherwise.

### GetUseSpecialCharactersOk

`func (o *UpdateAccountSettings) GetUseSpecialCharactersOk() (*string, bool)`

GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpecialCharacters

`func (o *UpdateAccountSettings) SetUseSpecialCharacters(v string)`

SetUseSpecialCharacters sets UseSpecialCharacters field to given value.

### HasUseSpecialCharacters

`func (o *UpdateAccountSettings) HasUseSpecialCharacters() bool`

HasUseSpecialCharacters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


