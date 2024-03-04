# DynamicSecretUpdateRdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUserExtendSession** | Pointer to **int64** | AllowUserExtendSession | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**FixedUserOnly** | Pointer to **string** | Allow access using externally (IdP) provided username [true/false] | [optional] [default to "false"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RdpAdminName** | Pointer to **string** | RDP Admin Name | [optional] 
**RdpAdminPwd** | Pointer to **string** | RDP Admin password | [optional] 
**RdpHostName** | Pointer to **string** | Hostname | [optional] 
**RdpHostPort** | Pointer to **string** | Port | [optional] [default to "22"]
**RdpUserGroups** | Pointer to **string** | Groups | [optional] 
**SecureAccessAllowExternalUser** | Pointer to **bool** | Allow providing external user for a domain users | [optional] [default to false]
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessRdGatewayServer** | Pointer to **string** | RD Gateway server | [optional] 
**SecureAccessRdpDomain** | Pointer to **string** | Required when the Dynamic Secret is used for a domain user | [optional] 
**SecureAccessRdpUser** | Pointer to **string** | Override the RDP Domain username | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**WarnUserBeforeExpiration** | Pointer to **int64** | WarnBeforeUserExpiration | [optional] 

## Methods

### NewDynamicSecretUpdateRdp

`func NewDynamicSecretUpdateRdp(name string, ) *DynamicSecretUpdateRdp`

NewDynamicSecretUpdateRdp instantiates a new DynamicSecretUpdateRdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateRdpWithDefaults

`func NewDynamicSecretUpdateRdpWithDefaults() *DynamicSecretUpdateRdp`

NewDynamicSecretUpdateRdpWithDefaults instantiates a new DynamicSecretUpdateRdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUserExtendSession

`func (o *DynamicSecretUpdateRdp) GetAllowUserExtendSession() int64`

GetAllowUserExtendSession returns the AllowUserExtendSession field if non-nil, zero value otherwise.

### GetAllowUserExtendSessionOk

`func (o *DynamicSecretUpdateRdp) GetAllowUserExtendSessionOk() (*int64, bool)`

GetAllowUserExtendSessionOk returns a tuple with the AllowUserExtendSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserExtendSession

`func (o *DynamicSecretUpdateRdp) SetAllowUserExtendSession(v int64)`

SetAllowUserExtendSession sets AllowUserExtendSession field to given value.

### HasAllowUserExtendSession

`func (o *DynamicSecretUpdateRdp) HasAllowUserExtendSession() bool`

HasAllowUserExtendSession returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretUpdateRdp) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateRdp) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateRdp) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateRdp) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateRdp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateRdp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateRdp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateRdp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixedUserOnly

`func (o *DynamicSecretUpdateRdp) GetFixedUserOnly() string`

GetFixedUserOnly returns the FixedUserOnly field if non-nil, zero value otherwise.

### GetFixedUserOnlyOk

`func (o *DynamicSecretUpdateRdp) GetFixedUserOnlyOk() (*string, bool)`

GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserOnly

`func (o *DynamicSecretUpdateRdp) SetFixedUserOnly(v string)`

SetFixedUserOnly sets FixedUserOnly field to given value.

### HasFixedUserOnly

`func (o *DynamicSecretUpdateRdp) HasFixedUserOnly() bool`

HasFixedUserOnly returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateRdp) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateRdp) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateRdp) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateRdp) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateRdp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateRdp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateRdp) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateRdp) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateRdp) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateRdp) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateRdp) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretUpdateRdp) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretUpdateRdp) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretUpdateRdp) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretUpdateRdp) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateRdp) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateRdp) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateRdp) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateRdp) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRdpAdminName

`func (o *DynamicSecretUpdateRdp) GetRdpAdminName() string`

GetRdpAdminName returns the RdpAdminName field if non-nil, zero value otherwise.

### GetRdpAdminNameOk

`func (o *DynamicSecretUpdateRdp) GetRdpAdminNameOk() (*string, bool)`

GetRdpAdminNameOk returns a tuple with the RdpAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminName

`func (o *DynamicSecretUpdateRdp) SetRdpAdminName(v string)`

SetRdpAdminName sets RdpAdminName field to given value.

### HasRdpAdminName

`func (o *DynamicSecretUpdateRdp) HasRdpAdminName() bool`

HasRdpAdminName returns a boolean if a field has been set.

### GetRdpAdminPwd

`func (o *DynamicSecretUpdateRdp) GetRdpAdminPwd() string`

GetRdpAdminPwd returns the RdpAdminPwd field if non-nil, zero value otherwise.

### GetRdpAdminPwdOk

`func (o *DynamicSecretUpdateRdp) GetRdpAdminPwdOk() (*string, bool)`

GetRdpAdminPwdOk returns a tuple with the RdpAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminPwd

`func (o *DynamicSecretUpdateRdp) SetRdpAdminPwd(v string)`

SetRdpAdminPwd sets RdpAdminPwd field to given value.

### HasRdpAdminPwd

`func (o *DynamicSecretUpdateRdp) HasRdpAdminPwd() bool`

HasRdpAdminPwd returns a boolean if a field has been set.

### GetRdpHostName

`func (o *DynamicSecretUpdateRdp) GetRdpHostName() string`

GetRdpHostName returns the RdpHostName field if non-nil, zero value otherwise.

### GetRdpHostNameOk

`func (o *DynamicSecretUpdateRdp) GetRdpHostNameOk() (*string, bool)`

GetRdpHostNameOk returns a tuple with the RdpHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostName

`func (o *DynamicSecretUpdateRdp) SetRdpHostName(v string)`

SetRdpHostName sets RdpHostName field to given value.

### HasRdpHostName

`func (o *DynamicSecretUpdateRdp) HasRdpHostName() bool`

HasRdpHostName returns a boolean if a field has been set.

### GetRdpHostPort

`func (o *DynamicSecretUpdateRdp) GetRdpHostPort() string`

GetRdpHostPort returns the RdpHostPort field if non-nil, zero value otherwise.

### GetRdpHostPortOk

`func (o *DynamicSecretUpdateRdp) GetRdpHostPortOk() (*string, bool)`

GetRdpHostPortOk returns a tuple with the RdpHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostPort

`func (o *DynamicSecretUpdateRdp) SetRdpHostPort(v string)`

SetRdpHostPort sets RdpHostPort field to given value.

### HasRdpHostPort

`func (o *DynamicSecretUpdateRdp) HasRdpHostPort() bool`

HasRdpHostPort returns a boolean if a field has been set.

### GetRdpUserGroups

`func (o *DynamicSecretUpdateRdp) GetRdpUserGroups() string`

GetRdpUserGroups returns the RdpUserGroups field if non-nil, zero value otherwise.

### GetRdpUserGroupsOk

`func (o *DynamicSecretUpdateRdp) GetRdpUserGroupsOk() (*string, bool)`

GetRdpUserGroupsOk returns a tuple with the RdpUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpUserGroups

`func (o *DynamicSecretUpdateRdp) SetRdpUserGroups(v string)`

SetRdpUserGroups sets RdpUserGroups field to given value.

### HasRdpUserGroups

`func (o *DynamicSecretUpdateRdp) HasRdpUserGroups() bool`

HasRdpUserGroups returns a boolean if a field has been set.

### GetSecureAccessAllowExternalUser

`func (o *DynamicSecretUpdateRdp) GetSecureAccessAllowExternalUser() bool`

GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field if non-nil, zero value otherwise.

### GetSecureAccessAllowExternalUserOk

`func (o *DynamicSecretUpdateRdp) GetSecureAccessAllowExternalUserOk() (*bool, bool)`

GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowExternalUser

`func (o *DynamicSecretUpdateRdp) SetSecureAccessAllowExternalUser(v bool)`

SetSecureAccessAllowExternalUser sets SecureAccessAllowExternalUser field to given value.

### HasSecureAccessAllowExternalUser

`func (o *DynamicSecretUpdateRdp) HasSecureAccessAllowExternalUser() bool`

HasSecureAccessAllowExternalUser returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateRdp) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateRdp) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateRdp) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateRdp) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretUpdateRdp) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretUpdateRdp) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretUpdateRdp) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretUpdateRdp) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdGatewayServer

`func (o *DynamicSecretUpdateRdp) GetSecureAccessRdGatewayServer() string`

GetSecureAccessRdGatewayServer returns the SecureAccessRdGatewayServer field if non-nil, zero value otherwise.

### GetSecureAccessRdGatewayServerOk

`func (o *DynamicSecretUpdateRdp) GetSecureAccessRdGatewayServerOk() (*string, bool)`

GetSecureAccessRdGatewayServerOk returns a tuple with the SecureAccessRdGatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdGatewayServer

`func (o *DynamicSecretUpdateRdp) SetSecureAccessRdGatewayServer(v string)`

SetSecureAccessRdGatewayServer sets SecureAccessRdGatewayServer field to given value.

### HasSecureAccessRdGatewayServer

`func (o *DynamicSecretUpdateRdp) HasSecureAccessRdGatewayServer() bool`

HasSecureAccessRdGatewayServer returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *DynamicSecretUpdateRdp) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *DynamicSecretUpdateRdp) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *DynamicSecretUpdateRdp) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *DynamicSecretUpdateRdp) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessRdpUser

`func (o *DynamicSecretUpdateRdp) GetSecureAccessRdpUser() string`

GetSecureAccessRdpUser returns the SecureAccessRdpUser field if non-nil, zero value otherwise.

### GetSecureAccessRdpUserOk

`func (o *DynamicSecretUpdateRdp) GetSecureAccessRdpUserOk() (*string, bool)`

GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpUser

`func (o *DynamicSecretUpdateRdp) SetSecureAccessRdpUser(v string)`

SetSecureAccessRdpUser sets SecureAccessRdpUser field to given value.

### HasSecureAccessRdpUser

`func (o *DynamicSecretUpdateRdp) HasSecureAccessRdpUser() bool`

HasSecureAccessRdpUser returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateRdp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateRdp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateRdp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateRdp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateRdp) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateRdp) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateRdp) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateRdp) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateRdp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateRdp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateRdp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateRdp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateRdp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateRdp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateRdp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateRdp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateRdp) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateRdp) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateRdp) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateRdp) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetWarnUserBeforeExpiration

`func (o *DynamicSecretUpdateRdp) GetWarnUserBeforeExpiration() int64`

GetWarnUserBeforeExpiration returns the WarnUserBeforeExpiration field if non-nil, zero value otherwise.

### GetWarnUserBeforeExpirationOk

`func (o *DynamicSecretUpdateRdp) GetWarnUserBeforeExpirationOk() (*int64, bool)`

GetWarnUserBeforeExpirationOk returns a tuple with the WarnUserBeforeExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUserBeforeExpiration

`func (o *DynamicSecretUpdateRdp) SetWarnUserBeforeExpiration(v int64)`

SetWarnUserBeforeExpiration sets WarnUserBeforeExpiration field to given value.

### HasWarnUserBeforeExpiration

`func (o *DynamicSecretUpdateRdp) HasWarnUserBeforeExpiration() bool`

HasWarnUserBeforeExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


