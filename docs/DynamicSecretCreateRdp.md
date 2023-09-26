# DynamicSecretCreateRdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUserExtendSession** | Pointer to **int64** | AllowUserExtendSession | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**FixedUserOnly** | Pointer to **string** | Allow access using externally (IdP) provided username [true/false] | [optional] [default to "false"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
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

### NewDynamicSecretCreateRdp

`func NewDynamicSecretCreateRdp(name string, ) *DynamicSecretCreateRdp`

NewDynamicSecretCreateRdp instantiates a new DynamicSecretCreateRdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateRdpWithDefaults

`func NewDynamicSecretCreateRdpWithDefaults() *DynamicSecretCreateRdp`

NewDynamicSecretCreateRdpWithDefaults instantiates a new DynamicSecretCreateRdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUserExtendSession

`func (o *DynamicSecretCreateRdp) GetAllowUserExtendSession() int64`

GetAllowUserExtendSession returns the AllowUserExtendSession field if non-nil, zero value otherwise.

### GetAllowUserExtendSessionOk

`func (o *DynamicSecretCreateRdp) GetAllowUserExtendSessionOk() (*int64, bool)`

GetAllowUserExtendSessionOk returns a tuple with the AllowUserExtendSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserExtendSession

`func (o *DynamicSecretCreateRdp) SetAllowUserExtendSession(v int64)`

SetAllowUserExtendSession sets AllowUserExtendSession field to given value.

### HasAllowUserExtendSession

`func (o *DynamicSecretCreateRdp) HasAllowUserExtendSession() bool`

HasAllowUserExtendSession returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateRdp) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateRdp) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateRdp) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateRdp) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetFixedUserOnly

`func (o *DynamicSecretCreateRdp) GetFixedUserOnly() string`

GetFixedUserOnly returns the FixedUserOnly field if non-nil, zero value otherwise.

### GetFixedUserOnlyOk

`func (o *DynamicSecretCreateRdp) GetFixedUserOnlyOk() (*string, bool)`

GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserOnly

`func (o *DynamicSecretCreateRdp) SetFixedUserOnly(v string)`

SetFixedUserOnly sets FixedUserOnly field to given value.

### HasFixedUserOnly

`func (o *DynamicSecretCreateRdp) HasFixedUserOnly() bool`

HasFixedUserOnly returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateRdp) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateRdp) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateRdp) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateRdp) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateRdp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateRdp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateRdp) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateRdp) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateRdp) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateRdp) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateRdp) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRdpAdminName

`func (o *DynamicSecretCreateRdp) GetRdpAdminName() string`

GetRdpAdminName returns the RdpAdminName field if non-nil, zero value otherwise.

### GetRdpAdminNameOk

`func (o *DynamicSecretCreateRdp) GetRdpAdminNameOk() (*string, bool)`

GetRdpAdminNameOk returns a tuple with the RdpAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminName

`func (o *DynamicSecretCreateRdp) SetRdpAdminName(v string)`

SetRdpAdminName sets RdpAdminName field to given value.

### HasRdpAdminName

`func (o *DynamicSecretCreateRdp) HasRdpAdminName() bool`

HasRdpAdminName returns a boolean if a field has been set.

### GetRdpAdminPwd

`func (o *DynamicSecretCreateRdp) GetRdpAdminPwd() string`

GetRdpAdminPwd returns the RdpAdminPwd field if non-nil, zero value otherwise.

### GetRdpAdminPwdOk

`func (o *DynamicSecretCreateRdp) GetRdpAdminPwdOk() (*string, bool)`

GetRdpAdminPwdOk returns a tuple with the RdpAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminPwd

`func (o *DynamicSecretCreateRdp) SetRdpAdminPwd(v string)`

SetRdpAdminPwd sets RdpAdminPwd field to given value.

### HasRdpAdminPwd

`func (o *DynamicSecretCreateRdp) HasRdpAdminPwd() bool`

HasRdpAdminPwd returns a boolean if a field has been set.

### GetRdpHostName

`func (o *DynamicSecretCreateRdp) GetRdpHostName() string`

GetRdpHostName returns the RdpHostName field if non-nil, zero value otherwise.

### GetRdpHostNameOk

`func (o *DynamicSecretCreateRdp) GetRdpHostNameOk() (*string, bool)`

GetRdpHostNameOk returns a tuple with the RdpHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostName

`func (o *DynamicSecretCreateRdp) SetRdpHostName(v string)`

SetRdpHostName sets RdpHostName field to given value.

### HasRdpHostName

`func (o *DynamicSecretCreateRdp) HasRdpHostName() bool`

HasRdpHostName returns a boolean if a field has been set.

### GetRdpHostPort

`func (o *DynamicSecretCreateRdp) GetRdpHostPort() string`

GetRdpHostPort returns the RdpHostPort field if non-nil, zero value otherwise.

### GetRdpHostPortOk

`func (o *DynamicSecretCreateRdp) GetRdpHostPortOk() (*string, bool)`

GetRdpHostPortOk returns a tuple with the RdpHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostPort

`func (o *DynamicSecretCreateRdp) SetRdpHostPort(v string)`

SetRdpHostPort sets RdpHostPort field to given value.

### HasRdpHostPort

`func (o *DynamicSecretCreateRdp) HasRdpHostPort() bool`

HasRdpHostPort returns a boolean if a field has been set.

### GetRdpUserGroups

`func (o *DynamicSecretCreateRdp) GetRdpUserGroups() string`

GetRdpUserGroups returns the RdpUserGroups field if non-nil, zero value otherwise.

### GetRdpUserGroupsOk

`func (o *DynamicSecretCreateRdp) GetRdpUserGroupsOk() (*string, bool)`

GetRdpUserGroupsOk returns a tuple with the RdpUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpUserGroups

`func (o *DynamicSecretCreateRdp) SetRdpUserGroups(v string)`

SetRdpUserGroups sets RdpUserGroups field to given value.

### HasRdpUserGroups

`func (o *DynamicSecretCreateRdp) HasRdpUserGroups() bool`

HasRdpUserGroups returns a boolean if a field has been set.

### GetSecureAccessAllowExternalUser

`func (o *DynamicSecretCreateRdp) GetSecureAccessAllowExternalUser() bool`

GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field if non-nil, zero value otherwise.

### GetSecureAccessAllowExternalUserOk

`func (o *DynamicSecretCreateRdp) GetSecureAccessAllowExternalUserOk() (*bool, bool)`

GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowExternalUser

`func (o *DynamicSecretCreateRdp) SetSecureAccessAllowExternalUser(v bool)`

SetSecureAccessAllowExternalUser sets SecureAccessAllowExternalUser field to given value.

### HasSecureAccessAllowExternalUser

`func (o *DynamicSecretCreateRdp) HasSecureAccessAllowExternalUser() bool`

HasSecureAccessAllowExternalUser returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateRdp) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateRdp) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateRdp) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateRdp) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretCreateRdp) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretCreateRdp) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretCreateRdp) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretCreateRdp) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdGatewayServer

`func (o *DynamicSecretCreateRdp) GetSecureAccessRdGatewayServer() string`

GetSecureAccessRdGatewayServer returns the SecureAccessRdGatewayServer field if non-nil, zero value otherwise.

### GetSecureAccessRdGatewayServerOk

`func (o *DynamicSecretCreateRdp) GetSecureAccessRdGatewayServerOk() (*string, bool)`

GetSecureAccessRdGatewayServerOk returns a tuple with the SecureAccessRdGatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdGatewayServer

`func (o *DynamicSecretCreateRdp) SetSecureAccessRdGatewayServer(v string)`

SetSecureAccessRdGatewayServer sets SecureAccessRdGatewayServer field to given value.

### HasSecureAccessRdGatewayServer

`func (o *DynamicSecretCreateRdp) HasSecureAccessRdGatewayServer() bool`

HasSecureAccessRdGatewayServer returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *DynamicSecretCreateRdp) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *DynamicSecretCreateRdp) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *DynamicSecretCreateRdp) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *DynamicSecretCreateRdp) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessRdpUser

`func (o *DynamicSecretCreateRdp) GetSecureAccessRdpUser() string`

GetSecureAccessRdpUser returns the SecureAccessRdpUser field if non-nil, zero value otherwise.

### GetSecureAccessRdpUserOk

`func (o *DynamicSecretCreateRdp) GetSecureAccessRdpUserOk() (*string, bool)`

GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpUser

`func (o *DynamicSecretCreateRdp) SetSecureAccessRdpUser(v string)`

SetSecureAccessRdpUser sets SecureAccessRdpUser field to given value.

### HasSecureAccessRdpUser

`func (o *DynamicSecretCreateRdp) HasSecureAccessRdpUser() bool`

HasSecureAccessRdpUser returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateRdp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateRdp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateRdp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateRdp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateRdp) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateRdp) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateRdp) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateRdp) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateRdp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateRdp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateRdp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateRdp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateRdp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateRdp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateRdp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateRdp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateRdp) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateRdp) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateRdp) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateRdp) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetWarnUserBeforeExpiration

`func (o *DynamicSecretCreateRdp) GetWarnUserBeforeExpiration() int64`

GetWarnUserBeforeExpiration returns the WarnUserBeforeExpiration field if non-nil, zero value otherwise.

### GetWarnUserBeforeExpirationOk

`func (o *DynamicSecretCreateRdp) GetWarnUserBeforeExpirationOk() (*int64, bool)`

GetWarnUserBeforeExpirationOk returns a tuple with the WarnUserBeforeExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUserBeforeExpiration

`func (o *DynamicSecretCreateRdp) SetWarnUserBeforeExpiration(v int64)`

SetWarnUserBeforeExpiration sets WarnUserBeforeExpiration field to given value.

### HasWarnUserBeforeExpiration

`func (o *DynamicSecretCreateRdp) HasWarnUserBeforeExpiration() bool`

HasWarnUserBeforeExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


