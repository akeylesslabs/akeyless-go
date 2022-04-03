# GatewayUpdateProducerRdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUserExtendSession** | Pointer to **int64** | AllowUserExtendSession | [optional] 
**FixedUserOnly** | Pointer to **string** | Fixed user | [optional] [default to "false"]
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RdpAdminName** | Pointer to **string** | RDP Admin Name | [optional] 
**RdpAdminPwd** | Pointer to **string** | RDP Admin password | [optional] 
**RdpHostName** | Pointer to **string** | Hostname | [optional] 
**RdpHostPort** | Pointer to **string** | Port | [optional] [default to "22"]
**RdpUserGroups** | Pointer to **string** | Groups | [optional] 
**SecureAccessAllowExternalUser** | Pointer to **bool** |  | [optional] 
**SecureAccessEnable** | Pointer to **string** |  | [optional] 
**SecureAccessHost** | Pointer to **[]string** |  | [optional] 
**SecureAccessRdpDomain** | Pointer to **string** |  | [optional] 
**SecureAccessRdpUser** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**WarnUserBeforeExpiration** | Pointer to **int64** | WarnBeforeUserExpiration | [optional] 

## Methods

### NewGatewayUpdateProducerRdp

`func NewGatewayUpdateProducerRdp(name string, ) *GatewayUpdateProducerRdp`

NewGatewayUpdateProducerRdp instantiates a new GatewayUpdateProducerRdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerRdpWithDefaults

`func NewGatewayUpdateProducerRdpWithDefaults() *GatewayUpdateProducerRdp`

NewGatewayUpdateProducerRdpWithDefaults instantiates a new GatewayUpdateProducerRdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUserExtendSession

`func (o *GatewayUpdateProducerRdp) GetAllowUserExtendSession() int64`

GetAllowUserExtendSession returns the AllowUserExtendSession field if non-nil, zero value otherwise.

### GetAllowUserExtendSessionOk

`func (o *GatewayUpdateProducerRdp) GetAllowUserExtendSessionOk() (*int64, bool)`

GetAllowUserExtendSessionOk returns a tuple with the AllowUserExtendSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserExtendSession

`func (o *GatewayUpdateProducerRdp) SetAllowUserExtendSession(v int64)`

SetAllowUserExtendSession sets AllowUserExtendSession field to given value.

### HasAllowUserExtendSession

`func (o *GatewayUpdateProducerRdp) HasAllowUserExtendSession() bool`

HasAllowUserExtendSession returns a boolean if a field has been set.

### GetFixedUserOnly

`func (o *GatewayUpdateProducerRdp) GetFixedUserOnly() string`

GetFixedUserOnly returns the FixedUserOnly field if non-nil, zero value otherwise.

### GetFixedUserOnlyOk

`func (o *GatewayUpdateProducerRdp) GetFixedUserOnlyOk() (*string, bool)`

GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserOnly

`func (o *GatewayUpdateProducerRdp) SetFixedUserOnly(v string)`

SetFixedUserOnly sets FixedUserOnly field to given value.

### HasFixedUserOnly

`func (o *GatewayUpdateProducerRdp) HasFixedUserOnly() bool`

HasFixedUserOnly returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerRdp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerRdp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerRdp) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerRdp) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerRdp) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerRdp) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerRdp) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerRdp) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerRdp) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerRdp) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerRdp) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRdpAdminName

`func (o *GatewayUpdateProducerRdp) GetRdpAdminName() string`

GetRdpAdminName returns the RdpAdminName field if non-nil, zero value otherwise.

### GetRdpAdminNameOk

`func (o *GatewayUpdateProducerRdp) GetRdpAdminNameOk() (*string, bool)`

GetRdpAdminNameOk returns a tuple with the RdpAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminName

`func (o *GatewayUpdateProducerRdp) SetRdpAdminName(v string)`

SetRdpAdminName sets RdpAdminName field to given value.

### HasRdpAdminName

`func (o *GatewayUpdateProducerRdp) HasRdpAdminName() bool`

HasRdpAdminName returns a boolean if a field has been set.

### GetRdpAdminPwd

`func (o *GatewayUpdateProducerRdp) GetRdpAdminPwd() string`

GetRdpAdminPwd returns the RdpAdminPwd field if non-nil, zero value otherwise.

### GetRdpAdminPwdOk

`func (o *GatewayUpdateProducerRdp) GetRdpAdminPwdOk() (*string, bool)`

GetRdpAdminPwdOk returns a tuple with the RdpAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminPwd

`func (o *GatewayUpdateProducerRdp) SetRdpAdminPwd(v string)`

SetRdpAdminPwd sets RdpAdminPwd field to given value.

### HasRdpAdminPwd

`func (o *GatewayUpdateProducerRdp) HasRdpAdminPwd() bool`

HasRdpAdminPwd returns a boolean if a field has been set.

### GetRdpHostName

`func (o *GatewayUpdateProducerRdp) GetRdpHostName() string`

GetRdpHostName returns the RdpHostName field if non-nil, zero value otherwise.

### GetRdpHostNameOk

`func (o *GatewayUpdateProducerRdp) GetRdpHostNameOk() (*string, bool)`

GetRdpHostNameOk returns a tuple with the RdpHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostName

`func (o *GatewayUpdateProducerRdp) SetRdpHostName(v string)`

SetRdpHostName sets RdpHostName field to given value.

### HasRdpHostName

`func (o *GatewayUpdateProducerRdp) HasRdpHostName() bool`

HasRdpHostName returns a boolean if a field has been set.

### GetRdpHostPort

`func (o *GatewayUpdateProducerRdp) GetRdpHostPort() string`

GetRdpHostPort returns the RdpHostPort field if non-nil, zero value otherwise.

### GetRdpHostPortOk

`func (o *GatewayUpdateProducerRdp) GetRdpHostPortOk() (*string, bool)`

GetRdpHostPortOk returns a tuple with the RdpHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostPort

`func (o *GatewayUpdateProducerRdp) SetRdpHostPort(v string)`

SetRdpHostPort sets RdpHostPort field to given value.

### HasRdpHostPort

`func (o *GatewayUpdateProducerRdp) HasRdpHostPort() bool`

HasRdpHostPort returns a boolean if a field has been set.

### GetRdpUserGroups

`func (o *GatewayUpdateProducerRdp) GetRdpUserGroups() string`

GetRdpUserGroups returns the RdpUserGroups field if non-nil, zero value otherwise.

### GetRdpUserGroupsOk

`func (o *GatewayUpdateProducerRdp) GetRdpUserGroupsOk() (*string, bool)`

GetRdpUserGroupsOk returns a tuple with the RdpUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpUserGroups

`func (o *GatewayUpdateProducerRdp) SetRdpUserGroups(v string)`

SetRdpUserGroups sets RdpUserGroups field to given value.

### HasRdpUserGroups

`func (o *GatewayUpdateProducerRdp) HasRdpUserGroups() bool`

HasRdpUserGroups returns a boolean if a field has been set.

### GetSecureAccessAllowExternalUser

`func (o *GatewayUpdateProducerRdp) GetSecureAccessAllowExternalUser() bool`

GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field if non-nil, zero value otherwise.

### GetSecureAccessAllowExternalUserOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessAllowExternalUserOk() (*bool, bool)`

GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowExternalUser

`func (o *GatewayUpdateProducerRdp) SetSecureAccessAllowExternalUser(v bool)`

SetSecureAccessAllowExternalUser sets SecureAccessAllowExternalUser field to given value.

### HasSecureAccessAllowExternalUser

`func (o *GatewayUpdateProducerRdp) HasSecureAccessAllowExternalUser() bool`

HasSecureAccessAllowExternalUser returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayUpdateProducerRdp) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayUpdateProducerRdp) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayUpdateProducerRdp) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayUpdateProducerRdp) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayUpdateProducerRdp) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayUpdateProducerRdp) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *GatewayUpdateProducerRdp) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *GatewayUpdateProducerRdp) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessRdpUser

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdpUser() string`

GetSecureAccessRdpUser returns the SecureAccessRdpUser field if non-nil, zero value otherwise.

### GetSecureAccessRdpUserOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdpUserOk() (*string, bool)`

GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpUser

`func (o *GatewayUpdateProducerRdp) SetSecureAccessRdpUser(v string)`

SetSecureAccessRdpUser sets SecureAccessRdpUser field to given value.

### HasSecureAccessRdpUser

`func (o *GatewayUpdateProducerRdp) HasSecureAccessRdpUser() bool`

HasSecureAccessRdpUser returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerRdp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerRdp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerRdp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerRdp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerRdp) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerRdp) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerRdp) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerRdp) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerRdp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerRdp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerRdp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerRdp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerRdp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerRdp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerRdp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerRdp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerRdp) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerRdp) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerRdp) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerRdp) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetWarnUserBeforeExpiration

`func (o *GatewayUpdateProducerRdp) GetWarnUserBeforeExpiration() int64`

GetWarnUserBeforeExpiration returns the WarnUserBeforeExpiration field if non-nil, zero value otherwise.

### GetWarnUserBeforeExpirationOk

`func (o *GatewayUpdateProducerRdp) GetWarnUserBeforeExpirationOk() (*int64, bool)`

GetWarnUserBeforeExpirationOk returns a tuple with the WarnUserBeforeExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUserBeforeExpiration

`func (o *GatewayUpdateProducerRdp) SetWarnUserBeforeExpiration(v int64)`

SetWarnUserBeforeExpiration sets WarnUserBeforeExpiration field to given value.

### HasWarnUserBeforeExpiration

`func (o *GatewayUpdateProducerRdp) HasWarnUserBeforeExpiration() bool`

HasWarnUserBeforeExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


