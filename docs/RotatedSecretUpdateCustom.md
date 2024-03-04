# RotatedSecretUpdateCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**AuthenticationCredentials** | Pointer to **string** | The credentials to connect with use-user-creds/use-target-creds | [optional] [default to "use-user-creds"]
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation [true/false] | [optional] 
**CustomPayload** | Pointer to **string** | Secret payload to be sent with rotation request | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_metadata"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Rotated secret name | 
**NewName** | Pointer to **string** | New item name | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**RotateAfterDisconnect** | Pointer to **string** | Rotate the value of the secret after SRA session ends [true/false] | [optional] [default to "false"]
**RotationHour** | Pointer to **int32** | The Hour of the rotation in UTC | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (1-365) | [optional] 
**SecureAccessAllowExternalUser** | Pointer to **bool** | Allow providing external user for a domain users | [optional] [default to false]
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessRdpDomain** | Pointer to **string** | Required when the Dynamic Secret is used for a domain user | [optional] 
**SecureAccessRdpUser** | Pointer to **string** | Override the RDP Domain username | [optional] 
**SecureAccessSshUser** | Pointer to **string** | Override the SSH username as indicated in SSH Certificate Issuer | [optional] 
**SecureAccessUrl** | Pointer to **string** | Destination URL to inject secrets | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**SecureAccessWebBrowsing** | Pointer to **bool** | Secure browser via Akeyless Web Access Bastion | [optional] [default to false]
**SecureAccessWebProxy** | Pointer to **bool** | Web-Proxy via Akeyless Web Access Bastion | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotatedSecretUpdateCustom

`func NewRotatedSecretUpdateCustom(name string, ) *RotatedSecretUpdateCustom`

NewRotatedSecretUpdateCustom instantiates a new RotatedSecretUpdateCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretUpdateCustomWithDefaults

`func NewRotatedSecretUpdateCustomWithDefaults() *RotatedSecretUpdateCustom`

NewRotatedSecretUpdateCustomWithDefaults instantiates a new RotatedSecretUpdateCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *RotatedSecretUpdateCustom) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *RotatedSecretUpdateCustom) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *RotatedSecretUpdateCustom) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *RotatedSecretUpdateCustom) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAuthenticationCredentials

`func (o *RotatedSecretUpdateCustom) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *RotatedSecretUpdateCustom) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *RotatedSecretUpdateCustom) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *RotatedSecretUpdateCustom) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretUpdateCustom) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretUpdateCustom) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretUpdateCustom) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretUpdateCustom) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetCustomPayload

`func (o *RotatedSecretUpdateCustom) GetCustomPayload() string`

GetCustomPayload returns the CustomPayload field if non-nil, zero value otherwise.

### GetCustomPayloadOk

`func (o *RotatedSecretUpdateCustom) GetCustomPayloadOk() (*string, bool)`

GetCustomPayloadOk returns a tuple with the CustomPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayload

`func (o *RotatedSecretUpdateCustom) SetCustomPayload(v string)`

SetCustomPayload sets CustomPayload field to given value.

### HasCustomPayload

`func (o *RotatedSecretUpdateCustom) HasCustomPayload() bool`

HasCustomPayload returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretUpdateCustom) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretUpdateCustom) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretUpdateCustom) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretUpdateCustom) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretUpdateCustom) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretUpdateCustom) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretUpdateCustom) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretUpdateCustom) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretUpdateCustom) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretUpdateCustom) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretUpdateCustom) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretUpdateCustom) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *RotatedSecretUpdateCustom) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *RotatedSecretUpdateCustom) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *RotatedSecretUpdateCustom) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *RotatedSecretUpdateCustom) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretUpdateCustom) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretUpdateCustom) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretUpdateCustom) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretUpdateCustom) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretUpdateCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretUpdateCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretUpdateCustom) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *RotatedSecretUpdateCustom) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *RotatedSecretUpdateCustom) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *RotatedSecretUpdateCustom) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *RotatedSecretUpdateCustom) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *RotatedSecretUpdateCustom) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretUpdateCustom) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretUpdateCustom) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretUpdateCustom) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRmTag

`func (o *RotatedSecretUpdateCustom) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *RotatedSecretUpdateCustom) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *RotatedSecretUpdateCustom) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *RotatedSecretUpdateCustom) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotateAfterDisconnect

`func (o *RotatedSecretUpdateCustom) GetRotateAfterDisconnect() string`

GetRotateAfterDisconnect returns the RotateAfterDisconnect field if non-nil, zero value otherwise.

### GetRotateAfterDisconnectOk

`func (o *RotatedSecretUpdateCustom) GetRotateAfterDisconnectOk() (*string, bool)`

GetRotateAfterDisconnectOk returns a tuple with the RotateAfterDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAfterDisconnect

`func (o *RotatedSecretUpdateCustom) SetRotateAfterDisconnect(v string)`

SetRotateAfterDisconnect sets RotateAfterDisconnect field to given value.

### HasRotateAfterDisconnect

`func (o *RotatedSecretUpdateCustom) HasRotateAfterDisconnect() bool`

HasRotateAfterDisconnect returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretUpdateCustom) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretUpdateCustom) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretUpdateCustom) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretUpdateCustom) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretUpdateCustom) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretUpdateCustom) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretUpdateCustom) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretUpdateCustom) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetSecureAccessAllowExternalUser

`func (o *RotatedSecretUpdateCustom) GetSecureAccessAllowExternalUser() bool`

GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field if non-nil, zero value otherwise.

### GetSecureAccessAllowExternalUserOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessAllowExternalUserOk() (*bool, bool)`

GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowExternalUser

`func (o *RotatedSecretUpdateCustom) SetSecureAccessAllowExternalUser(v bool)`

SetSecureAccessAllowExternalUser sets SecureAccessAllowExternalUser field to given value.

### HasSecureAccessAllowExternalUser

`func (o *RotatedSecretUpdateCustom) HasSecureAccessAllowExternalUser() bool`

HasSecureAccessAllowExternalUser returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *RotatedSecretUpdateCustom) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *RotatedSecretUpdateCustom) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *RotatedSecretUpdateCustom) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *RotatedSecretUpdateCustom) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *RotatedSecretUpdateCustom) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *RotatedSecretUpdateCustom) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *RotatedSecretUpdateCustom) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *RotatedSecretUpdateCustom) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *RotatedSecretUpdateCustom) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *RotatedSecretUpdateCustom) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *RotatedSecretUpdateCustom) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *RotatedSecretUpdateCustom) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessRdpUser

`func (o *RotatedSecretUpdateCustom) GetSecureAccessRdpUser() string`

GetSecureAccessRdpUser returns the SecureAccessRdpUser field if non-nil, zero value otherwise.

### GetSecureAccessRdpUserOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessRdpUserOk() (*string, bool)`

GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpUser

`func (o *RotatedSecretUpdateCustom) SetSecureAccessRdpUser(v string)`

SetSecureAccessRdpUser sets SecureAccessRdpUser field to given value.

### HasSecureAccessRdpUser

`func (o *RotatedSecretUpdateCustom) HasSecureAccessRdpUser() bool`

HasSecureAccessRdpUser returns a boolean if a field has been set.

### GetSecureAccessSshUser

`func (o *RotatedSecretUpdateCustom) GetSecureAccessSshUser() string`

GetSecureAccessSshUser returns the SecureAccessSshUser field if non-nil, zero value otherwise.

### GetSecureAccessSshUserOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessSshUserOk() (*string, bool)`

GetSecureAccessSshUserOk returns a tuple with the SecureAccessSshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessSshUser

`func (o *RotatedSecretUpdateCustom) SetSecureAccessSshUser(v string)`

SetSecureAccessSshUser sets SecureAccessSshUser field to given value.

### HasSecureAccessSshUser

`func (o *RotatedSecretUpdateCustom) HasSecureAccessSshUser() bool`

HasSecureAccessSshUser returns a boolean if a field has been set.

### GetSecureAccessUrl

`func (o *RotatedSecretUpdateCustom) GetSecureAccessUrl() string`

GetSecureAccessUrl returns the SecureAccessUrl field if non-nil, zero value otherwise.

### GetSecureAccessUrlOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessUrlOk() (*string, bool)`

GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUrl

`func (o *RotatedSecretUpdateCustom) SetSecureAccessUrl(v string)`

SetSecureAccessUrl sets SecureAccessUrl field to given value.

### HasSecureAccessUrl

`func (o *RotatedSecretUpdateCustom) HasSecureAccessUrl() bool`

HasSecureAccessUrl returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *RotatedSecretUpdateCustom) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *RotatedSecretUpdateCustom) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *RotatedSecretUpdateCustom) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *RotatedSecretUpdateCustom) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *RotatedSecretUpdateCustom) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *RotatedSecretUpdateCustom) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *RotatedSecretUpdateCustom) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *RotatedSecretUpdateCustom) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *RotatedSecretUpdateCustom) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *RotatedSecretUpdateCustom) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetToken

`func (o *RotatedSecretUpdateCustom) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretUpdateCustom) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretUpdateCustom) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretUpdateCustom) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretUpdateCustom) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretUpdateCustom) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretUpdateCustom) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretUpdateCustom) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


