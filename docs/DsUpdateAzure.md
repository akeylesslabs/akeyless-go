# DsUpdateAzure

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
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
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

### NewDsUpdateAzure

`func NewDsUpdateAzure(name string, ) *DsUpdateAzure`

NewDsUpdateAzure instantiates a new DsUpdateAzure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsUpdateAzureWithDefaults

`func NewDsUpdateAzureWithDefaults() *DsUpdateAzure`

NewDsUpdateAzureWithDefaults instantiates a new DsUpdateAzure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppObjId

`func (o *DsUpdateAzure) GetAppObjId() string`

GetAppObjId returns the AppObjId field if non-nil, zero value otherwise.

### GetAppObjIdOk

`func (o *DsUpdateAzure) GetAppObjIdOk() (*string, bool)`

GetAppObjIdOk returns a tuple with the AppObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppObjId

`func (o *DsUpdateAzure) SetAppObjId(v string)`

SetAppObjId sets AppObjId field to given value.

### HasAppObjId

`func (o *DsUpdateAzure) HasAppObjId() bool`

HasAppObjId returns a boolean if a field has been set.

### GetAzureClientId

`func (o *DsUpdateAzure) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *DsUpdateAzure) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *DsUpdateAzure) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *DsUpdateAzure) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureClientSecret

`func (o *DsUpdateAzure) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *DsUpdateAzure) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *DsUpdateAzure) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### HasAzureClientSecret

`func (o *DsUpdateAzure) HasAzureClientSecret() bool`

HasAzureClientSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *DsUpdateAzure) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *DsUpdateAzure) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *DsUpdateAzure) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *DsUpdateAzure) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DsUpdateAzure) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DsUpdateAzure) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DsUpdateAzure) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DsUpdateAzure) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetFixedUserClaimKeyname

`func (o *DsUpdateAzure) GetFixedUserClaimKeyname() string`

GetFixedUserClaimKeyname returns the FixedUserClaimKeyname field if non-nil, zero value otherwise.

### GetFixedUserClaimKeynameOk

`func (o *DsUpdateAzure) GetFixedUserClaimKeynameOk() (*string, bool)`

GetFixedUserClaimKeynameOk returns a tuple with the FixedUserClaimKeyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserClaimKeyname

`func (o *DsUpdateAzure) SetFixedUserClaimKeyname(v string)`

SetFixedUserClaimKeyname sets FixedUserClaimKeyname field to given value.

### HasFixedUserClaimKeyname

`func (o *DsUpdateAzure) HasFixedUserClaimKeyname() bool`

HasFixedUserClaimKeyname returns a boolean if a field has been set.

### GetFixedUserOnly

`func (o *DsUpdateAzure) GetFixedUserOnly() bool`

GetFixedUserOnly returns the FixedUserOnly field if non-nil, zero value otherwise.

### GetFixedUserOnlyOk

`func (o *DsUpdateAzure) GetFixedUserOnlyOk() (*bool, bool)`

GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserOnly

`func (o *DsUpdateAzure) SetFixedUserOnly(v bool)`

SetFixedUserOnly sets FixedUserOnly field to given value.

### HasFixedUserOnly

`func (o *DsUpdateAzure) HasFixedUserOnly() bool`

HasFixedUserOnly returns a boolean if a field has been set.

### GetJson

`func (o *DsUpdateAzure) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DsUpdateAzure) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DsUpdateAzure) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DsUpdateAzure) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DsUpdateAzure) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DsUpdateAzure) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DsUpdateAzure) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DsUpdateAzure) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DsUpdateAzure) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DsUpdateAzure) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DsUpdateAzure) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DsUpdateAzure) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DsUpdateAzure) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DsUpdateAzure) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DsUpdateAzure) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DsUpdateAzure) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DsUpdateAzure) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DsUpdateAzure) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DsUpdateAzure) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DsUpdateAzure) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DsUpdateAzure) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DsUpdateAzure) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DsUpdateAzure) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *DsUpdateAzure) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *DsUpdateAzure) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *DsUpdateAzure) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *DsUpdateAzure) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *DsUpdateAzure) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *DsUpdateAzure) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *DsUpdateAzure) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *DsUpdateAzure) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *DsUpdateAzure) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DsUpdateAzure) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DsUpdateAzure) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DsUpdateAzure) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DsUpdateAzure) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DsUpdateAzure) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DsUpdateAzure) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DsUpdateAzure) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DsUpdateAzure) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DsUpdateAzure) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DsUpdateAzure) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DsUpdateAzure) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DsUpdateAzure) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DsUpdateAzure) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DsUpdateAzure) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DsUpdateAzure) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserGroupObjId

`func (o *DsUpdateAzure) GetUserGroupObjId() string`

GetUserGroupObjId returns the UserGroupObjId field if non-nil, zero value otherwise.

### GetUserGroupObjIdOk

`func (o *DsUpdateAzure) GetUserGroupObjIdOk() (*string, bool)`

GetUserGroupObjIdOk returns a tuple with the UserGroupObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupObjId

`func (o *DsUpdateAzure) SetUserGroupObjId(v string)`

SetUserGroupObjId sets UserGroupObjId field to given value.

### HasUserGroupObjId

`func (o *DsUpdateAzure) HasUserGroupObjId() bool`

HasUserGroupObjId returns a boolean if a field has been set.

### GetUserPortalAccess

`func (o *DsUpdateAzure) GetUserPortalAccess() bool`

GetUserPortalAccess returns the UserPortalAccess field if non-nil, zero value otherwise.

### GetUserPortalAccessOk

`func (o *DsUpdateAzure) GetUserPortalAccessOk() (*bool, bool)`

GetUserPortalAccessOk returns a tuple with the UserPortalAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPortalAccess

`func (o *DsUpdateAzure) SetUserPortalAccess(v bool)`

SetUserPortalAccess sets UserPortalAccess field to given value.

### HasUserPortalAccess

`func (o *DsUpdateAzure) HasUserPortalAccess() bool`

HasUserPortalAccess returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *DsUpdateAzure) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *DsUpdateAzure) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *DsUpdateAzure) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *DsUpdateAzure) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetUserProgrammaticAccess

`func (o *DsUpdateAzure) GetUserProgrammaticAccess() bool`

GetUserProgrammaticAccess returns the UserProgrammaticAccess field if non-nil, zero value otherwise.

### GetUserProgrammaticAccessOk

`func (o *DsUpdateAzure) GetUserProgrammaticAccessOk() (*bool, bool)`

GetUserProgrammaticAccessOk returns a tuple with the UserProgrammaticAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProgrammaticAccess

`func (o *DsUpdateAzure) SetUserProgrammaticAccess(v bool)`

SetUserProgrammaticAccess sets UserProgrammaticAccess field to given value.

### HasUserProgrammaticAccess

`func (o *DsUpdateAzure) HasUserProgrammaticAccess() bool`

HasUserProgrammaticAccess returns a boolean if a field has been set.

### GetUserRoleTemplateId

`func (o *DsUpdateAzure) GetUserRoleTemplateId() string`

GetUserRoleTemplateId returns the UserRoleTemplateId field if non-nil, zero value otherwise.

### GetUserRoleTemplateIdOk

`func (o *DsUpdateAzure) GetUserRoleTemplateIdOk() (*string, bool)`

GetUserRoleTemplateIdOk returns a tuple with the UserRoleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoleTemplateId

`func (o *DsUpdateAzure) SetUserRoleTemplateId(v string)`

SetUserRoleTemplateId sets UserRoleTemplateId field to given value.

### HasUserRoleTemplateId

`func (o *DsUpdateAzure) HasUserRoleTemplateId() bool`

HasUserRoleTemplateId returns a boolean if a field has been set.

### GetUserTtl

`func (o *DsUpdateAzure) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DsUpdateAzure) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DsUpdateAzure) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DsUpdateAzure) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


