# RotatedSecretCreateWindows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationCredentials** | Pointer to **string** | The credentials to connect with use-user-creds/use-target-creds | [optional] [default to "use-user-creds"]
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation [true/false] | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Rotated secret name | 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**RotateAfterDisconnect** | Pointer to **string** | Rotate the value of the secret after SRA session ends [true/false] | [optional] [default to "false"]
**RotatedPassword** | Pointer to **string** | rotated-username password (relevant only for rotator-type&#x3D;password) | [optional] 
**RotatedUsername** | Pointer to **string** | username to be rotated, if selected use-self-creds at rotator-creds-type, this username will try to rotate it&#39;s own password, if use-target-creds is selected, target credentials will be use to rotate the rotated-password (relevant only for rotator-type&#x3D;password) | [optional] 
**RotationHour** | Pointer to **int32** | The Hour of the rotation in UTC | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (1-365) | [optional] 
**RotatorType** | **string** | The rotator type. options: [target/password] | 
**SamePassword** | Pointer to **string** | Rotate same password for each host from the Linked Target (relevant only for Linked Target) | [optional] 
**SecureAccessAllowExternalUser** | Pointer to **bool** | Allow providing external user for a domain users | [optional] [default to false]
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessRdpDomain** | Pointer to **string** | Required when the Dynamic Secret is used for a domain user | [optional] 
**SecureAccessRdpUser** | Pointer to **string** | Override the RDP Domain username | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotatedSecretCreateWindows

`func NewRotatedSecretCreateWindows(name string, rotatorType string, targetName string, ) *RotatedSecretCreateWindows`

NewRotatedSecretCreateWindows instantiates a new RotatedSecretCreateWindows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretCreateWindowsWithDefaults

`func NewRotatedSecretCreateWindowsWithDefaults() *RotatedSecretCreateWindows`

NewRotatedSecretCreateWindowsWithDefaults instantiates a new RotatedSecretCreateWindows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationCredentials

`func (o *RotatedSecretCreateWindows) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *RotatedSecretCreateWindows) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *RotatedSecretCreateWindows) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *RotatedSecretCreateWindows) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretCreateWindows) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretCreateWindows) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretCreateWindows) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretCreateWindows) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretCreateWindows) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretCreateWindows) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretCreateWindows) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretCreateWindows) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretCreateWindows) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretCreateWindows) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretCreateWindows) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretCreateWindows) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretCreateWindows) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretCreateWindows) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretCreateWindows) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretCreateWindows) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretCreateWindows) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretCreateWindows) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretCreateWindows) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretCreateWindows) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretCreateWindows) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretCreateWindows) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretCreateWindows) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *RotatedSecretCreateWindows) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretCreateWindows) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretCreateWindows) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretCreateWindows) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRotateAfterDisconnect

`func (o *RotatedSecretCreateWindows) GetRotateAfterDisconnect() string`

GetRotateAfterDisconnect returns the RotateAfterDisconnect field if non-nil, zero value otherwise.

### GetRotateAfterDisconnectOk

`func (o *RotatedSecretCreateWindows) GetRotateAfterDisconnectOk() (*string, bool)`

GetRotateAfterDisconnectOk returns a tuple with the RotateAfterDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAfterDisconnect

`func (o *RotatedSecretCreateWindows) SetRotateAfterDisconnect(v string)`

SetRotateAfterDisconnect sets RotateAfterDisconnect field to given value.

### HasRotateAfterDisconnect

`func (o *RotatedSecretCreateWindows) HasRotateAfterDisconnect() bool`

HasRotateAfterDisconnect returns a boolean if a field has been set.

### GetRotatedPassword

`func (o *RotatedSecretCreateWindows) GetRotatedPassword() string`

GetRotatedPassword returns the RotatedPassword field if non-nil, zero value otherwise.

### GetRotatedPasswordOk

`func (o *RotatedSecretCreateWindows) GetRotatedPasswordOk() (*string, bool)`

GetRotatedPasswordOk returns a tuple with the RotatedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedPassword

`func (o *RotatedSecretCreateWindows) SetRotatedPassword(v string)`

SetRotatedPassword sets RotatedPassword field to given value.

### HasRotatedPassword

`func (o *RotatedSecretCreateWindows) HasRotatedPassword() bool`

HasRotatedPassword returns a boolean if a field has been set.

### GetRotatedUsername

`func (o *RotatedSecretCreateWindows) GetRotatedUsername() string`

GetRotatedUsername returns the RotatedUsername field if non-nil, zero value otherwise.

### GetRotatedUsernameOk

`func (o *RotatedSecretCreateWindows) GetRotatedUsernameOk() (*string, bool)`

GetRotatedUsernameOk returns a tuple with the RotatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedUsername

`func (o *RotatedSecretCreateWindows) SetRotatedUsername(v string)`

SetRotatedUsername sets RotatedUsername field to given value.

### HasRotatedUsername

`func (o *RotatedSecretCreateWindows) HasRotatedUsername() bool`

HasRotatedUsername returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretCreateWindows) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretCreateWindows) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretCreateWindows) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretCreateWindows) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretCreateWindows) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretCreateWindows) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretCreateWindows) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretCreateWindows) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotatorType

`func (o *RotatedSecretCreateWindows) GetRotatorType() string`

GetRotatorType returns the RotatorType field if non-nil, zero value otherwise.

### GetRotatorTypeOk

`func (o *RotatedSecretCreateWindows) GetRotatorTypeOk() (*string, bool)`

GetRotatorTypeOk returns a tuple with the RotatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorType

`func (o *RotatedSecretCreateWindows) SetRotatorType(v string)`

SetRotatorType sets RotatorType field to given value.


### GetSamePassword

`func (o *RotatedSecretCreateWindows) GetSamePassword() string`

GetSamePassword returns the SamePassword field if non-nil, zero value otherwise.

### GetSamePasswordOk

`func (o *RotatedSecretCreateWindows) GetSamePasswordOk() (*string, bool)`

GetSamePasswordOk returns a tuple with the SamePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePassword

`func (o *RotatedSecretCreateWindows) SetSamePassword(v string)`

SetSamePassword sets SamePassword field to given value.

### HasSamePassword

`func (o *RotatedSecretCreateWindows) HasSamePassword() bool`

HasSamePassword returns a boolean if a field has been set.

### GetSecureAccessAllowExternalUser

`func (o *RotatedSecretCreateWindows) GetSecureAccessAllowExternalUser() bool`

GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field if non-nil, zero value otherwise.

### GetSecureAccessAllowExternalUserOk

`func (o *RotatedSecretCreateWindows) GetSecureAccessAllowExternalUserOk() (*bool, bool)`

GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowExternalUser

`func (o *RotatedSecretCreateWindows) SetSecureAccessAllowExternalUser(v bool)`

SetSecureAccessAllowExternalUser sets SecureAccessAllowExternalUser field to given value.

### HasSecureAccessAllowExternalUser

`func (o *RotatedSecretCreateWindows) HasSecureAccessAllowExternalUser() bool`

HasSecureAccessAllowExternalUser returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *RotatedSecretCreateWindows) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *RotatedSecretCreateWindows) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *RotatedSecretCreateWindows) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *RotatedSecretCreateWindows) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *RotatedSecretCreateWindows) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *RotatedSecretCreateWindows) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *RotatedSecretCreateWindows) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *RotatedSecretCreateWindows) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *RotatedSecretCreateWindows) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *RotatedSecretCreateWindows) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *RotatedSecretCreateWindows) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *RotatedSecretCreateWindows) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessRdpUser

`func (o *RotatedSecretCreateWindows) GetSecureAccessRdpUser() string`

GetSecureAccessRdpUser returns the SecureAccessRdpUser field if non-nil, zero value otherwise.

### GetSecureAccessRdpUserOk

`func (o *RotatedSecretCreateWindows) GetSecureAccessRdpUserOk() (*string, bool)`

GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpUser

`func (o *RotatedSecretCreateWindows) SetSecureAccessRdpUser(v string)`

SetSecureAccessRdpUser sets SecureAccessRdpUser field to given value.

### HasSecureAccessRdpUser

`func (o *RotatedSecretCreateWindows) HasSecureAccessRdpUser() bool`

HasSecureAccessRdpUser returns a boolean if a field has been set.

### GetTags

`func (o *RotatedSecretCreateWindows) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RotatedSecretCreateWindows) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RotatedSecretCreateWindows) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RotatedSecretCreateWindows) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *RotatedSecretCreateWindows) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *RotatedSecretCreateWindows) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *RotatedSecretCreateWindows) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetToken

`func (o *RotatedSecretCreateWindows) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretCreateWindows) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretCreateWindows) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretCreateWindows) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretCreateWindows) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretCreateWindows) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretCreateWindows) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretCreateWindows) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


