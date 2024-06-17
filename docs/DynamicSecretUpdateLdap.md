# DynamicSecretUpdateLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | Pointer to **string** |  | [optional] 
**BindDn** | Pointer to **string** | Bind DN | [optional] 
**BindDnPassword** | Pointer to **string** | Bind DN Password | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**ExternalUsername** | Pointer to **string** | Externally provided username [true/false] | [optional] [default to "false"]
**GroupDn** | Pointer to **string** | Group DN which the temporary user should be added | [optional] 
**HostProvider** | Pointer to **string** | Host provider type [explicit/target], Default Host provider is explicit, Relevant only for Secure Remote Access of ssh cert issuer, ldap rotated secret and ldap dynamic secret | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**LdapCaCert** | Pointer to **string** | CA Certificate File Content | [optional] 
**LdapUrl** | Pointer to **string** | LDAP Server URL | [optional] 
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessRdGatewayServer** | Pointer to **string** | RD Gateway server | [optional] 
**SecureAccessRdpDomain** | Pointer to **string** | Required when the Dynamic Secret is used for a domain user | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**Target** | Pointer to **[]string** | A list of linked targets to be associated, Relevant only for Secure Remote Access for ssh cert issuer, ldap rotated secret and ldap dynamic secret, To specify multiple targets use argument multiple times | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExpiration** | Pointer to **string** | Token expiration | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserAttribute** | Pointer to **string** | User Attribute | [optional] 
**UserDn** | Pointer to **string** | User DN | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretUpdateLdap

`func NewDynamicSecretUpdateLdap(name string, ) *DynamicSecretUpdateLdap`

NewDynamicSecretUpdateLdap instantiates a new DynamicSecretUpdateLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateLdapWithDefaults

`func NewDynamicSecretUpdateLdapWithDefaults() *DynamicSecretUpdateLdap`

NewDynamicSecretUpdateLdapWithDefaults instantiates a new DynamicSecretUpdateLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *DynamicSecretUpdateLdap) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DynamicSecretUpdateLdap) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DynamicSecretUpdateLdap) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *DynamicSecretUpdateLdap) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetBindDn

`func (o *DynamicSecretUpdateLdap) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *DynamicSecretUpdateLdap) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *DynamicSecretUpdateLdap) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *DynamicSecretUpdateLdap) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindDnPassword

`func (o *DynamicSecretUpdateLdap) GetBindDnPassword() string`

GetBindDnPassword returns the BindDnPassword field if non-nil, zero value otherwise.

### GetBindDnPasswordOk

`func (o *DynamicSecretUpdateLdap) GetBindDnPasswordOk() (*string, bool)`

GetBindDnPasswordOk returns a tuple with the BindDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDnPassword

`func (o *DynamicSecretUpdateLdap) SetBindDnPassword(v string)`

SetBindDnPassword sets BindDnPassword field to given value.

### HasBindDnPassword

`func (o *DynamicSecretUpdateLdap) HasBindDnPassword() bool`

HasBindDnPassword returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretUpdateLdap) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateLdap) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateLdap) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateLdap) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateLdap) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateLdap) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateLdap) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateLdap) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalUsername

`func (o *DynamicSecretUpdateLdap) GetExternalUsername() string`

GetExternalUsername returns the ExternalUsername field if non-nil, zero value otherwise.

### GetExternalUsernameOk

`func (o *DynamicSecretUpdateLdap) GetExternalUsernameOk() (*string, bool)`

GetExternalUsernameOk returns a tuple with the ExternalUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUsername

`func (o *DynamicSecretUpdateLdap) SetExternalUsername(v string)`

SetExternalUsername sets ExternalUsername field to given value.

### HasExternalUsername

`func (o *DynamicSecretUpdateLdap) HasExternalUsername() bool`

HasExternalUsername returns a boolean if a field has been set.

### GetGroupDn

`func (o *DynamicSecretUpdateLdap) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *DynamicSecretUpdateLdap) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *DynamicSecretUpdateLdap) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *DynamicSecretUpdateLdap) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetHostProvider

`func (o *DynamicSecretUpdateLdap) GetHostProvider() string`

GetHostProvider returns the HostProvider field if non-nil, zero value otherwise.

### GetHostProviderOk

`func (o *DynamicSecretUpdateLdap) GetHostProviderOk() (*string, bool)`

GetHostProviderOk returns a tuple with the HostProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostProvider

`func (o *DynamicSecretUpdateLdap) SetHostProvider(v string)`

SetHostProvider sets HostProvider field to given value.

### HasHostProvider

`func (o *DynamicSecretUpdateLdap) HasHostProvider() bool`

HasHostProvider returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateLdap) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateLdap) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateLdap) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateLdap) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLdapCaCert

`func (o *DynamicSecretUpdateLdap) GetLdapCaCert() string`

GetLdapCaCert returns the LdapCaCert field if non-nil, zero value otherwise.

### GetLdapCaCertOk

`func (o *DynamicSecretUpdateLdap) GetLdapCaCertOk() (*string, bool)`

GetLdapCaCertOk returns a tuple with the LdapCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCaCert

`func (o *DynamicSecretUpdateLdap) SetLdapCaCert(v string)`

SetLdapCaCert sets LdapCaCert field to given value.

### HasLdapCaCert

`func (o *DynamicSecretUpdateLdap) HasLdapCaCert() bool`

HasLdapCaCert returns a boolean if a field has been set.

### GetLdapUrl

`func (o *DynamicSecretUpdateLdap) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *DynamicSecretUpdateLdap) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *DynamicSecretUpdateLdap) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *DynamicSecretUpdateLdap) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateLdap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateLdap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateLdap) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateLdap) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateLdap) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateLdap) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateLdap) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretUpdateLdap) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretUpdateLdap) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretUpdateLdap) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretUpdateLdap) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateLdap) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateLdap) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateLdap) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateLdap) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateLdap) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateLdap) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateLdap) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateLdap) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretUpdateLdap) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretUpdateLdap) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretUpdateLdap) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretUpdateLdap) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdGatewayServer

`func (o *DynamicSecretUpdateLdap) GetSecureAccessRdGatewayServer() string`

GetSecureAccessRdGatewayServer returns the SecureAccessRdGatewayServer field if non-nil, zero value otherwise.

### GetSecureAccessRdGatewayServerOk

`func (o *DynamicSecretUpdateLdap) GetSecureAccessRdGatewayServerOk() (*string, bool)`

GetSecureAccessRdGatewayServerOk returns a tuple with the SecureAccessRdGatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdGatewayServer

`func (o *DynamicSecretUpdateLdap) SetSecureAccessRdGatewayServer(v string)`

SetSecureAccessRdGatewayServer sets SecureAccessRdGatewayServer field to given value.

### HasSecureAccessRdGatewayServer

`func (o *DynamicSecretUpdateLdap) HasSecureAccessRdGatewayServer() bool`

HasSecureAccessRdGatewayServer returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *DynamicSecretUpdateLdap) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *DynamicSecretUpdateLdap) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *DynamicSecretUpdateLdap) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *DynamicSecretUpdateLdap) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateLdap) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateLdap) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateLdap) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateLdap) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTarget

`func (o *DynamicSecretUpdateLdap) GetTarget() []string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DynamicSecretUpdateLdap) GetTargetOk() (*[]string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DynamicSecretUpdateLdap) SetTarget(v []string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DynamicSecretUpdateLdap) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateLdap) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateLdap) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateLdap) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateLdap) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateLdap) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateLdap) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateLdap) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateLdap) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenExpiration

`func (o *DynamicSecretUpdateLdap) GetTokenExpiration() string`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *DynamicSecretUpdateLdap) GetTokenExpirationOk() (*string, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *DynamicSecretUpdateLdap) SetTokenExpiration(v string)`

SetTokenExpiration sets TokenExpiration field to given value.

### HasTokenExpiration

`func (o *DynamicSecretUpdateLdap) HasTokenExpiration() bool`

HasTokenExpiration returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateLdap) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateLdap) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateLdap) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateLdap) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserAttribute

`func (o *DynamicSecretUpdateLdap) GetUserAttribute() string`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *DynamicSecretUpdateLdap) GetUserAttributeOk() (*string, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *DynamicSecretUpdateLdap) SetUserAttribute(v string)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *DynamicSecretUpdateLdap) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### GetUserDn

`func (o *DynamicSecretUpdateLdap) GetUserDn() string`

GetUserDn returns the UserDn field if non-nil, zero value otherwise.

### GetUserDnOk

`func (o *DynamicSecretUpdateLdap) GetUserDnOk() (*string, bool)`

GetUserDnOk returns a tuple with the UserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDn

`func (o *DynamicSecretUpdateLdap) SetUserDn(v string)`

SetUserDn sets UserDn field to given value.

### HasUserDn

`func (o *DynamicSecretUpdateLdap) HasUserDn() bool`

HasUserDn returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateLdap) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateLdap) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateLdap) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateLdap) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


