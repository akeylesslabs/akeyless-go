# RotatedSecretCreateLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | Pointer to **string** |  | [optional] 
**AuthenticationCredentials** | Pointer to **string** | The credentials to connect with use-user-creds/use-target-creds | [optional] [default to "use-user-creds"]
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation [true/false] | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**HostProvider** | Pointer to **string** | Host provider type [explicit/target], Relevant only for Secure Remote Access of ssh cert issuer and ldap rotated secret | [optional] [default to "explicit"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Rotated secret name | 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**RotateAfterDisconnect** | Pointer to **string** | Rotate the value of the secret after SRA session ends [true/false] | [optional] [default to "false"]
**RotatedPassword** | Pointer to **string** | rotated-username password (relevant only for rotator-type&#x3D;ldap) | [optional] 
**RotatedUsername** | Pointer to **string** | username to be rotated, if selected use-self-creds at rotator-creds-type, this username will try to rotate it&#39;s own password, if use-target-creds is selected, target credentials will be use to rotate the rotated-password (relevant only for rotator-type&#x3D;ldap) | [optional] 
**RotationHour** | Pointer to **int32** | The Hour of the rotation in UTC | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (1-365) | [optional] 
**RotatorType** | **string** | The rotator type. options: [target/ldap] | 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessRdpDomain** | Pointer to **string** | Required when the Dynamic Secret is used for a domain user | [optional] 
**SecureAccessUrl** | Pointer to **string** | Destination URL to inject secrets | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**Target** | Pointer to **[]string** | A list of linked targets to be associated, Relevant only for Secure Remote Access for ssh cert issuer and ldap rotated secret, To specify multiple targets use argument multiple times | [optional] 
**TargetName** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserAttribute** | Pointer to **string** | LDAP User Attribute, Default value \&quot;cn\&quot; | [optional] [default to "cn"]
**UserDn** | Pointer to **string** | Base DN to Perform User Search | [optional] 

## Methods

### NewRotatedSecretCreateLdap

`func NewRotatedSecretCreateLdap(name string, rotatorType string, targetName string, ) *RotatedSecretCreateLdap`

NewRotatedSecretCreateLdap instantiates a new RotatedSecretCreateLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretCreateLdapWithDefaults

`func NewRotatedSecretCreateLdapWithDefaults() *RotatedSecretCreateLdap`

NewRotatedSecretCreateLdapWithDefaults instantiates a new RotatedSecretCreateLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *RotatedSecretCreateLdap) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *RotatedSecretCreateLdap) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *RotatedSecretCreateLdap) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *RotatedSecretCreateLdap) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetAuthenticationCredentials

`func (o *RotatedSecretCreateLdap) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *RotatedSecretCreateLdap) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *RotatedSecretCreateLdap) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *RotatedSecretCreateLdap) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretCreateLdap) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretCreateLdap) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretCreateLdap) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretCreateLdap) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretCreateLdap) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretCreateLdap) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretCreateLdap) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretCreateLdap) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretCreateLdap) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretCreateLdap) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretCreateLdap) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretCreateLdap) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostProvider

`func (o *RotatedSecretCreateLdap) GetHostProvider() string`

GetHostProvider returns the HostProvider field if non-nil, zero value otherwise.

### GetHostProviderOk

`func (o *RotatedSecretCreateLdap) GetHostProviderOk() (*string, bool)`

GetHostProviderOk returns a tuple with the HostProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostProvider

`func (o *RotatedSecretCreateLdap) SetHostProvider(v string)`

SetHostProvider sets HostProvider field to given value.

### HasHostProvider

`func (o *RotatedSecretCreateLdap) HasHostProvider() bool`

HasHostProvider returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretCreateLdap) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretCreateLdap) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretCreateLdap) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretCreateLdap) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretCreateLdap) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretCreateLdap) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretCreateLdap) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretCreateLdap) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretCreateLdap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretCreateLdap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretCreateLdap) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *RotatedSecretCreateLdap) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretCreateLdap) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretCreateLdap) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretCreateLdap) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRotateAfterDisconnect

`func (o *RotatedSecretCreateLdap) GetRotateAfterDisconnect() string`

GetRotateAfterDisconnect returns the RotateAfterDisconnect field if non-nil, zero value otherwise.

### GetRotateAfterDisconnectOk

`func (o *RotatedSecretCreateLdap) GetRotateAfterDisconnectOk() (*string, bool)`

GetRotateAfterDisconnectOk returns a tuple with the RotateAfterDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAfterDisconnect

`func (o *RotatedSecretCreateLdap) SetRotateAfterDisconnect(v string)`

SetRotateAfterDisconnect sets RotateAfterDisconnect field to given value.

### HasRotateAfterDisconnect

`func (o *RotatedSecretCreateLdap) HasRotateAfterDisconnect() bool`

HasRotateAfterDisconnect returns a boolean if a field has been set.

### GetRotatedPassword

`func (o *RotatedSecretCreateLdap) GetRotatedPassword() string`

GetRotatedPassword returns the RotatedPassword field if non-nil, zero value otherwise.

### GetRotatedPasswordOk

`func (o *RotatedSecretCreateLdap) GetRotatedPasswordOk() (*string, bool)`

GetRotatedPasswordOk returns a tuple with the RotatedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedPassword

`func (o *RotatedSecretCreateLdap) SetRotatedPassword(v string)`

SetRotatedPassword sets RotatedPassword field to given value.

### HasRotatedPassword

`func (o *RotatedSecretCreateLdap) HasRotatedPassword() bool`

HasRotatedPassword returns a boolean if a field has been set.

### GetRotatedUsername

`func (o *RotatedSecretCreateLdap) GetRotatedUsername() string`

GetRotatedUsername returns the RotatedUsername field if non-nil, zero value otherwise.

### GetRotatedUsernameOk

`func (o *RotatedSecretCreateLdap) GetRotatedUsernameOk() (*string, bool)`

GetRotatedUsernameOk returns a tuple with the RotatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedUsername

`func (o *RotatedSecretCreateLdap) SetRotatedUsername(v string)`

SetRotatedUsername sets RotatedUsername field to given value.

### HasRotatedUsername

`func (o *RotatedSecretCreateLdap) HasRotatedUsername() bool`

HasRotatedUsername returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretCreateLdap) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretCreateLdap) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretCreateLdap) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretCreateLdap) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretCreateLdap) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretCreateLdap) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretCreateLdap) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretCreateLdap) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotatorType

`func (o *RotatedSecretCreateLdap) GetRotatorType() string`

GetRotatorType returns the RotatorType field if non-nil, zero value otherwise.

### GetRotatorTypeOk

`func (o *RotatedSecretCreateLdap) GetRotatorTypeOk() (*string, bool)`

GetRotatorTypeOk returns a tuple with the RotatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorType

`func (o *RotatedSecretCreateLdap) SetRotatorType(v string)`

SetRotatorType sets RotatorType field to given value.


### GetSecureAccessEnable

`func (o *RotatedSecretCreateLdap) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *RotatedSecretCreateLdap) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *RotatedSecretCreateLdap) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *RotatedSecretCreateLdap) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *RotatedSecretCreateLdap) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *RotatedSecretCreateLdap) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *RotatedSecretCreateLdap) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *RotatedSecretCreateLdap) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *RotatedSecretCreateLdap) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *RotatedSecretCreateLdap) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *RotatedSecretCreateLdap) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *RotatedSecretCreateLdap) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessUrl

`func (o *RotatedSecretCreateLdap) GetSecureAccessUrl() string`

GetSecureAccessUrl returns the SecureAccessUrl field if non-nil, zero value otherwise.

### GetSecureAccessUrlOk

`func (o *RotatedSecretCreateLdap) GetSecureAccessUrlOk() (*string, bool)`

GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUrl

`func (o *RotatedSecretCreateLdap) SetSecureAccessUrl(v string)`

SetSecureAccessUrl sets SecureAccessUrl field to given value.

### HasSecureAccessUrl

`func (o *RotatedSecretCreateLdap) HasSecureAccessUrl() bool`

HasSecureAccessUrl returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *RotatedSecretCreateLdap) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *RotatedSecretCreateLdap) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *RotatedSecretCreateLdap) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *RotatedSecretCreateLdap) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *RotatedSecretCreateLdap) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RotatedSecretCreateLdap) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RotatedSecretCreateLdap) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RotatedSecretCreateLdap) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTarget

`func (o *RotatedSecretCreateLdap) GetTarget() []string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RotatedSecretCreateLdap) GetTargetOk() (*[]string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RotatedSecretCreateLdap) SetTarget(v []string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RotatedSecretCreateLdap) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetName

`func (o *RotatedSecretCreateLdap) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *RotatedSecretCreateLdap) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *RotatedSecretCreateLdap) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetToken

`func (o *RotatedSecretCreateLdap) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretCreateLdap) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretCreateLdap) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretCreateLdap) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretCreateLdap) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretCreateLdap) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretCreateLdap) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretCreateLdap) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserAttribute

`func (o *RotatedSecretCreateLdap) GetUserAttribute() string`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *RotatedSecretCreateLdap) GetUserAttributeOk() (*string, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *RotatedSecretCreateLdap) SetUserAttribute(v string)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *RotatedSecretCreateLdap) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### GetUserDn

`func (o *RotatedSecretCreateLdap) GetUserDn() string`

GetUserDn returns the UserDn field if non-nil, zero value otherwise.

### GetUserDnOk

`func (o *RotatedSecretCreateLdap) GetUserDnOk() (*string, bool)`

GetUserDnOk returns a tuple with the UserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDn

`func (o *RotatedSecretCreateLdap) SetUserDn(v string)`

SetUserDn sets UserDn field to given value.

### HasUserDn

`func (o *RotatedSecretCreateLdap) HasUserDn() bool`

HasUserDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


