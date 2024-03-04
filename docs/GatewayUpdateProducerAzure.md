# GatewayUpdateProducerAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppObjId** | Pointer to **string** | Azure App Object Id | [optional] 
**AzureClientId** | Pointer to **string** | Azure Client ID | [optional] 
**AzureClientSecret** | Pointer to **string** | Azure Client Secret | [optional] 
**AzureTenantId** | Pointer to **string** | Azure Tenant ID | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**FixedUserClaimKeyname** | Pointer to **string** | FixedUserClaimKeyname | [optional] [default to "false"]
**FixedUserOnly** | Pointer to **bool** | Fixed user | [optional] [default to false]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to true]
**SecureAccessWebBrowsing** | Pointer to **bool** | Secure browser via Akeyless Web Access Bastion | [optional] [default to false]
**SecureAccessWebProxy** | Pointer to **bool** | Web-Proxy via Akeyless Web Access Bastion | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserGroupObjId** | Pointer to **string** | User Group Object Id | [optional] 
**UserPortalAccess** | Pointer to **bool** | Azure User portal access | [optional] [default to false]
**UserPrincipalName** | Pointer to **string** | User Principal Name | [optional] 
**UserProgrammaticAccess** | Pointer to **bool** | Azure User programmatic access | [optional] [default to false]
**UserRoleTemplateId** | Pointer to **string** | User Role Template Id | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayUpdateProducerAzure

`func NewGatewayUpdateProducerAzure(name string, ) *GatewayUpdateProducerAzure`

NewGatewayUpdateProducerAzure instantiates a new GatewayUpdateProducerAzure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerAzureWithDefaults

`func NewGatewayUpdateProducerAzureWithDefaults() *GatewayUpdateProducerAzure`

NewGatewayUpdateProducerAzureWithDefaults instantiates a new GatewayUpdateProducerAzure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppObjId

`func (o *GatewayUpdateProducerAzure) GetAppObjId() string`

GetAppObjId returns the AppObjId field if non-nil, zero value otherwise.

### GetAppObjIdOk

`func (o *GatewayUpdateProducerAzure) GetAppObjIdOk() (*string, bool)`

GetAppObjIdOk returns a tuple with the AppObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppObjId

`func (o *GatewayUpdateProducerAzure) SetAppObjId(v string)`

SetAppObjId sets AppObjId field to given value.

### HasAppObjId

`func (o *GatewayUpdateProducerAzure) HasAppObjId() bool`

HasAppObjId returns a boolean if a field has been set.

### GetAzureClientId

`func (o *GatewayUpdateProducerAzure) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *GatewayUpdateProducerAzure) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *GatewayUpdateProducerAzure) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *GatewayUpdateProducerAzure) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureClientSecret

`func (o *GatewayUpdateProducerAzure) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *GatewayUpdateProducerAzure) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *GatewayUpdateProducerAzure) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### HasAzureClientSecret

`func (o *GatewayUpdateProducerAzure) HasAzureClientSecret() bool`

HasAzureClientSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *GatewayUpdateProducerAzure) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *GatewayUpdateProducerAzure) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *GatewayUpdateProducerAzure) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *GatewayUpdateProducerAzure) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayUpdateProducerAzure) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerAzure) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerAzure) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerAzure) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetFixedUserClaimKeyname

`func (o *GatewayUpdateProducerAzure) GetFixedUserClaimKeyname() string`

GetFixedUserClaimKeyname returns the FixedUserClaimKeyname field if non-nil, zero value otherwise.

### GetFixedUserClaimKeynameOk

`func (o *GatewayUpdateProducerAzure) GetFixedUserClaimKeynameOk() (*string, bool)`

GetFixedUserClaimKeynameOk returns a tuple with the FixedUserClaimKeyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserClaimKeyname

`func (o *GatewayUpdateProducerAzure) SetFixedUserClaimKeyname(v string)`

SetFixedUserClaimKeyname sets FixedUserClaimKeyname field to given value.

### HasFixedUserClaimKeyname

`func (o *GatewayUpdateProducerAzure) HasFixedUserClaimKeyname() bool`

HasFixedUserClaimKeyname returns a boolean if a field has been set.

### GetFixedUserOnly

`func (o *GatewayUpdateProducerAzure) GetFixedUserOnly() bool`

GetFixedUserOnly returns the FixedUserOnly field if non-nil, zero value otherwise.

### GetFixedUserOnlyOk

`func (o *GatewayUpdateProducerAzure) GetFixedUserOnlyOk() (*bool, bool)`

GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserOnly

`func (o *GatewayUpdateProducerAzure) SetFixedUserOnly(v bool)`

SetFixedUserOnly sets FixedUserOnly field to given value.

### HasFixedUserOnly

`func (o *GatewayUpdateProducerAzure) HasFixedUserOnly() bool`

HasFixedUserOnly returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateProducerAzure) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateProducerAzure) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateProducerAzure) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateProducerAzure) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerAzure) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerAzure) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerAzure) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerAzure) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerAzure) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerAzure) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerAzure) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *GatewayUpdateProducerAzure) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *GatewayUpdateProducerAzure) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *GatewayUpdateProducerAzure) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *GatewayUpdateProducerAzure) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerAzure) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerAzure) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerAzure) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerAzure) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayUpdateProducerAzure) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayUpdateProducerAzure) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayUpdateProducerAzure) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayUpdateProducerAzure) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayUpdateProducerAzure) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayUpdateProducerAzure) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayUpdateProducerAzure) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayUpdateProducerAzure) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *GatewayUpdateProducerAzure) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *GatewayUpdateProducerAzure) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *GatewayUpdateProducerAzure) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *GatewayUpdateProducerAzure) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *GatewayUpdateProducerAzure) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *GatewayUpdateProducerAzure) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *GatewayUpdateProducerAzure) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *GatewayUpdateProducerAzure) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerAzure) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerAzure) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerAzure) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerAzure) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerAzure) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerAzure) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerAzure) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerAzure) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerAzure) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerAzure) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerAzure) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerAzure) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerAzure) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerAzure) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerAzure) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerAzure) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserGroupObjId

`func (o *GatewayUpdateProducerAzure) GetUserGroupObjId() string`

GetUserGroupObjId returns the UserGroupObjId field if non-nil, zero value otherwise.

### GetUserGroupObjIdOk

`func (o *GatewayUpdateProducerAzure) GetUserGroupObjIdOk() (*string, bool)`

GetUserGroupObjIdOk returns a tuple with the UserGroupObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupObjId

`func (o *GatewayUpdateProducerAzure) SetUserGroupObjId(v string)`

SetUserGroupObjId sets UserGroupObjId field to given value.

### HasUserGroupObjId

`func (o *GatewayUpdateProducerAzure) HasUserGroupObjId() bool`

HasUserGroupObjId returns a boolean if a field has been set.

### GetUserPortalAccess

`func (o *GatewayUpdateProducerAzure) GetUserPortalAccess() bool`

GetUserPortalAccess returns the UserPortalAccess field if non-nil, zero value otherwise.

### GetUserPortalAccessOk

`func (o *GatewayUpdateProducerAzure) GetUserPortalAccessOk() (*bool, bool)`

GetUserPortalAccessOk returns a tuple with the UserPortalAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPortalAccess

`func (o *GatewayUpdateProducerAzure) SetUserPortalAccess(v bool)`

SetUserPortalAccess sets UserPortalAccess field to given value.

### HasUserPortalAccess

`func (o *GatewayUpdateProducerAzure) HasUserPortalAccess() bool`

HasUserPortalAccess returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *GatewayUpdateProducerAzure) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *GatewayUpdateProducerAzure) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *GatewayUpdateProducerAzure) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *GatewayUpdateProducerAzure) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetUserProgrammaticAccess

`func (o *GatewayUpdateProducerAzure) GetUserProgrammaticAccess() bool`

GetUserProgrammaticAccess returns the UserProgrammaticAccess field if non-nil, zero value otherwise.

### GetUserProgrammaticAccessOk

`func (o *GatewayUpdateProducerAzure) GetUserProgrammaticAccessOk() (*bool, bool)`

GetUserProgrammaticAccessOk returns a tuple with the UserProgrammaticAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProgrammaticAccess

`func (o *GatewayUpdateProducerAzure) SetUserProgrammaticAccess(v bool)`

SetUserProgrammaticAccess sets UserProgrammaticAccess field to given value.

### HasUserProgrammaticAccess

`func (o *GatewayUpdateProducerAzure) HasUserProgrammaticAccess() bool`

HasUserProgrammaticAccess returns a boolean if a field has been set.

### GetUserRoleTemplateId

`func (o *GatewayUpdateProducerAzure) GetUserRoleTemplateId() string`

GetUserRoleTemplateId returns the UserRoleTemplateId field if non-nil, zero value otherwise.

### GetUserRoleTemplateIdOk

`func (o *GatewayUpdateProducerAzure) GetUserRoleTemplateIdOk() (*string, bool)`

GetUserRoleTemplateIdOk returns a tuple with the UserRoleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoleTemplateId

`func (o *GatewayUpdateProducerAzure) SetUserRoleTemplateId(v string)`

SetUserRoleTemplateId sets UserRoleTemplateId field to given value.

### HasUserRoleTemplateId

`func (o *GatewayUpdateProducerAzure) HasUserRoleTemplateId() bool`

HasUserRoleTemplateId returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerAzure) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerAzure) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerAzure) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerAzure) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


