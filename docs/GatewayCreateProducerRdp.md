# GatewayCreateProducerRdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUserExtendSession** | Pointer to **int64** | AllowUserExtendSession | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item | [optional] 
**FixedUserOnly** | Pointer to **string** | Fixed user | [optional] [default to "false"]
**Name** | **string** | Producer name | 
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

### NewGatewayCreateProducerRdp

`func NewGatewayCreateProducerRdp(name string, ) *GatewayCreateProducerRdp`

NewGatewayCreateProducerRdp instantiates a new GatewayCreateProducerRdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerRdpWithDefaults

`func NewGatewayCreateProducerRdpWithDefaults() *GatewayCreateProducerRdp`

NewGatewayCreateProducerRdpWithDefaults instantiates a new GatewayCreateProducerRdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUserExtendSession

`func (o *GatewayCreateProducerRdp) GetAllowUserExtendSession() int64`

GetAllowUserExtendSession returns the AllowUserExtendSession field if non-nil, zero value otherwise.

### GetAllowUserExtendSessionOk

`func (o *GatewayCreateProducerRdp) GetAllowUserExtendSessionOk() (*int64, bool)`

GetAllowUserExtendSessionOk returns a tuple with the AllowUserExtendSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserExtendSession

`func (o *GatewayCreateProducerRdp) SetAllowUserExtendSession(v int64)`

SetAllowUserExtendSession sets AllowUserExtendSession field to given value.

### HasAllowUserExtendSession

`func (o *GatewayCreateProducerRdp) HasAllowUserExtendSession() bool`

HasAllowUserExtendSession returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayCreateProducerRdp) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayCreateProducerRdp) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayCreateProducerRdp) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayCreateProducerRdp) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetFixedUserOnly

`func (o *GatewayCreateProducerRdp) GetFixedUserOnly() string`

GetFixedUserOnly returns the FixedUserOnly field if non-nil, zero value otherwise.

### GetFixedUserOnlyOk

`func (o *GatewayCreateProducerRdp) GetFixedUserOnlyOk() (*string, bool)`

GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserOnly

`func (o *GatewayCreateProducerRdp) SetFixedUserOnly(v string)`

SetFixedUserOnly sets FixedUserOnly field to given value.

### HasFixedUserOnly

`func (o *GatewayCreateProducerRdp) HasFixedUserOnly() bool`

HasFixedUserOnly returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerRdp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerRdp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerRdp) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerRdp) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerRdp) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerRdp) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerRdp) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRdpAdminName

`func (o *GatewayCreateProducerRdp) GetRdpAdminName() string`

GetRdpAdminName returns the RdpAdminName field if non-nil, zero value otherwise.

### GetRdpAdminNameOk

`func (o *GatewayCreateProducerRdp) GetRdpAdminNameOk() (*string, bool)`

GetRdpAdminNameOk returns a tuple with the RdpAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminName

`func (o *GatewayCreateProducerRdp) SetRdpAdminName(v string)`

SetRdpAdminName sets RdpAdminName field to given value.

### HasRdpAdminName

`func (o *GatewayCreateProducerRdp) HasRdpAdminName() bool`

HasRdpAdminName returns a boolean if a field has been set.

### GetRdpAdminPwd

`func (o *GatewayCreateProducerRdp) GetRdpAdminPwd() string`

GetRdpAdminPwd returns the RdpAdminPwd field if non-nil, zero value otherwise.

### GetRdpAdminPwdOk

`func (o *GatewayCreateProducerRdp) GetRdpAdminPwdOk() (*string, bool)`

GetRdpAdminPwdOk returns a tuple with the RdpAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminPwd

`func (o *GatewayCreateProducerRdp) SetRdpAdminPwd(v string)`

SetRdpAdminPwd sets RdpAdminPwd field to given value.

### HasRdpAdminPwd

`func (o *GatewayCreateProducerRdp) HasRdpAdminPwd() bool`

HasRdpAdminPwd returns a boolean if a field has been set.

### GetRdpHostName

`func (o *GatewayCreateProducerRdp) GetRdpHostName() string`

GetRdpHostName returns the RdpHostName field if non-nil, zero value otherwise.

### GetRdpHostNameOk

`func (o *GatewayCreateProducerRdp) GetRdpHostNameOk() (*string, bool)`

GetRdpHostNameOk returns a tuple with the RdpHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostName

`func (o *GatewayCreateProducerRdp) SetRdpHostName(v string)`

SetRdpHostName sets RdpHostName field to given value.

### HasRdpHostName

`func (o *GatewayCreateProducerRdp) HasRdpHostName() bool`

HasRdpHostName returns a boolean if a field has been set.

### GetRdpHostPort

`func (o *GatewayCreateProducerRdp) GetRdpHostPort() string`

GetRdpHostPort returns the RdpHostPort field if non-nil, zero value otherwise.

### GetRdpHostPortOk

`func (o *GatewayCreateProducerRdp) GetRdpHostPortOk() (*string, bool)`

GetRdpHostPortOk returns a tuple with the RdpHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostPort

`func (o *GatewayCreateProducerRdp) SetRdpHostPort(v string)`

SetRdpHostPort sets RdpHostPort field to given value.

### HasRdpHostPort

`func (o *GatewayCreateProducerRdp) HasRdpHostPort() bool`

HasRdpHostPort returns a boolean if a field has been set.

### GetRdpUserGroups

`func (o *GatewayCreateProducerRdp) GetRdpUserGroups() string`

GetRdpUserGroups returns the RdpUserGroups field if non-nil, zero value otherwise.

### GetRdpUserGroupsOk

`func (o *GatewayCreateProducerRdp) GetRdpUserGroupsOk() (*string, bool)`

GetRdpUserGroupsOk returns a tuple with the RdpUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpUserGroups

`func (o *GatewayCreateProducerRdp) SetRdpUserGroups(v string)`

SetRdpUserGroups sets RdpUserGroups field to given value.

### HasRdpUserGroups

`func (o *GatewayCreateProducerRdp) HasRdpUserGroups() bool`

HasRdpUserGroups returns a boolean if a field has been set.

### GetSecureAccessAllowExternalUser

`func (o *GatewayCreateProducerRdp) GetSecureAccessAllowExternalUser() bool`

GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field if non-nil, zero value otherwise.

### GetSecureAccessAllowExternalUserOk

`func (o *GatewayCreateProducerRdp) GetSecureAccessAllowExternalUserOk() (*bool, bool)`

GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowExternalUser

`func (o *GatewayCreateProducerRdp) SetSecureAccessAllowExternalUser(v bool)`

SetSecureAccessAllowExternalUser sets SecureAccessAllowExternalUser field to given value.

### HasSecureAccessAllowExternalUser

`func (o *GatewayCreateProducerRdp) HasSecureAccessAllowExternalUser() bool`

HasSecureAccessAllowExternalUser returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayCreateProducerRdp) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayCreateProducerRdp) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayCreateProducerRdp) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayCreateProducerRdp) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayCreateProducerRdp) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayCreateProducerRdp) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayCreateProducerRdp) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayCreateProducerRdp) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *GatewayCreateProducerRdp) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *GatewayCreateProducerRdp) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *GatewayCreateProducerRdp) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *GatewayCreateProducerRdp) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessRdpUser

`func (o *GatewayCreateProducerRdp) GetSecureAccessRdpUser() string`

GetSecureAccessRdpUser returns the SecureAccessRdpUser field if non-nil, zero value otherwise.

### GetSecureAccessRdpUserOk

`func (o *GatewayCreateProducerRdp) GetSecureAccessRdpUserOk() (*string, bool)`

GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpUser

`func (o *GatewayCreateProducerRdp) SetSecureAccessRdpUser(v string)`

SetSecureAccessRdpUser sets SecureAccessRdpUser field to given value.

### HasSecureAccessRdpUser

`func (o *GatewayCreateProducerRdp) HasSecureAccessRdpUser() bool`

HasSecureAccessRdpUser returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerRdp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerRdp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerRdp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerRdp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerRdp) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerRdp) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerRdp) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerRdp) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerRdp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerRdp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerRdp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerRdp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerRdp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerRdp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerRdp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerRdp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerRdp) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerRdp) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerRdp) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerRdp) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetWarnUserBeforeExpiration

`func (o *GatewayCreateProducerRdp) GetWarnUserBeforeExpiration() int64`

GetWarnUserBeforeExpiration returns the WarnUserBeforeExpiration field if non-nil, zero value otherwise.

### GetWarnUserBeforeExpirationOk

`func (o *GatewayCreateProducerRdp) GetWarnUserBeforeExpirationOk() (*int64, bool)`

GetWarnUserBeforeExpirationOk returns a tuple with the WarnUserBeforeExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUserBeforeExpiration

`func (o *GatewayCreateProducerRdp) SetWarnUserBeforeExpiration(v int64)`

SetWarnUserBeforeExpiration sets WarnUserBeforeExpiration field to given value.

### HasWarnUserBeforeExpiration

`func (o *GatewayCreateProducerRdp) HasWarnUserBeforeExpiration() bool`

HasWarnUserBeforeExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


