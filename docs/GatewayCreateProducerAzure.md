# GatewayCreateProducerAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppObjId** | Pointer to **string** | Azure App Object Id | [optional] 
**ClientId** | **string** | Azure Client ID | 
**ClientSecret** | **string** | Azure Client Secret | 
**Name** | **string** | Producer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**TenantId** | **string** | Azure Tenant ID | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserGroupObjId** | Pointer to **string** | User Group Object Id | [optional] 
**UserPortalAccess** | Pointer to **bool** | Azure User portal access | [optional] [default to false]
**UserPrincipalName** | Pointer to **string** | User Principal Name | [optional] 
**UserProgrammaticAccess** | Pointer to **bool** | Azure User programmatic access | [optional] [default to true]
**UserRoleTemplateId** | Pointer to **string** | User Role Template Id | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayCreateProducerAzure

`func NewGatewayCreateProducerAzure(clientId string, clientSecret string, name string, tenantId string, ) *GatewayCreateProducerAzure`

NewGatewayCreateProducerAzure instantiates a new GatewayCreateProducerAzure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerAzureWithDefaults

`func NewGatewayCreateProducerAzureWithDefaults() *GatewayCreateProducerAzure`

NewGatewayCreateProducerAzureWithDefaults instantiates a new GatewayCreateProducerAzure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppObjId

`func (o *GatewayCreateProducerAzure) GetAppObjId() string`

GetAppObjId returns the AppObjId field if non-nil, zero value otherwise.

### GetAppObjIdOk

`func (o *GatewayCreateProducerAzure) GetAppObjIdOk() (*string, bool)`

GetAppObjIdOk returns a tuple with the AppObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppObjId

`func (o *GatewayCreateProducerAzure) SetAppObjId(v string)`

SetAppObjId sets AppObjId field to given value.

### HasAppObjId

`func (o *GatewayCreateProducerAzure) HasAppObjId() bool`

HasAppObjId returns a boolean if a field has been set.

### GetClientId

`func (o *GatewayCreateProducerAzure) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GatewayCreateProducerAzure) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GatewayCreateProducerAzure) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *GatewayCreateProducerAzure) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GatewayCreateProducerAzure) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GatewayCreateProducerAzure) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetName

`func (o *GatewayCreateProducerAzure) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerAzure) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerAzure) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *GatewayCreateProducerAzure) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateProducerAzure) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateProducerAzure) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateProducerAzure) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerAzure) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerAzure) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerAzure) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerAzure) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTenantId

`func (o *GatewayCreateProducerAzure) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GatewayCreateProducerAzure) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GatewayCreateProducerAzure) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetToken

`func (o *GatewayCreateProducerAzure) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerAzure) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerAzure) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerAzure) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerAzure) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerAzure) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerAzure) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerAzure) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserGroupObjId

`func (o *GatewayCreateProducerAzure) GetUserGroupObjId() string`

GetUserGroupObjId returns the UserGroupObjId field if non-nil, zero value otherwise.

### GetUserGroupObjIdOk

`func (o *GatewayCreateProducerAzure) GetUserGroupObjIdOk() (*string, bool)`

GetUserGroupObjIdOk returns a tuple with the UserGroupObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupObjId

`func (o *GatewayCreateProducerAzure) SetUserGroupObjId(v string)`

SetUserGroupObjId sets UserGroupObjId field to given value.

### HasUserGroupObjId

`func (o *GatewayCreateProducerAzure) HasUserGroupObjId() bool`

HasUserGroupObjId returns a boolean if a field has been set.

### GetUserPortalAccess

`func (o *GatewayCreateProducerAzure) GetUserPortalAccess() bool`

GetUserPortalAccess returns the UserPortalAccess field if non-nil, zero value otherwise.

### GetUserPortalAccessOk

`func (o *GatewayCreateProducerAzure) GetUserPortalAccessOk() (*bool, bool)`

GetUserPortalAccessOk returns a tuple with the UserPortalAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPortalAccess

`func (o *GatewayCreateProducerAzure) SetUserPortalAccess(v bool)`

SetUserPortalAccess sets UserPortalAccess field to given value.

### HasUserPortalAccess

`func (o *GatewayCreateProducerAzure) HasUserPortalAccess() bool`

HasUserPortalAccess returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *GatewayCreateProducerAzure) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *GatewayCreateProducerAzure) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *GatewayCreateProducerAzure) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *GatewayCreateProducerAzure) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetUserProgrammaticAccess

`func (o *GatewayCreateProducerAzure) GetUserProgrammaticAccess() bool`

GetUserProgrammaticAccess returns the UserProgrammaticAccess field if non-nil, zero value otherwise.

### GetUserProgrammaticAccessOk

`func (o *GatewayCreateProducerAzure) GetUserProgrammaticAccessOk() (*bool, bool)`

GetUserProgrammaticAccessOk returns a tuple with the UserProgrammaticAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProgrammaticAccess

`func (o *GatewayCreateProducerAzure) SetUserProgrammaticAccess(v bool)`

SetUserProgrammaticAccess sets UserProgrammaticAccess field to given value.

### HasUserProgrammaticAccess

`func (o *GatewayCreateProducerAzure) HasUserProgrammaticAccess() bool`

HasUserProgrammaticAccess returns a boolean if a field has been set.

### GetUserRoleTemplateId

`func (o *GatewayCreateProducerAzure) GetUserRoleTemplateId() string`

GetUserRoleTemplateId returns the UserRoleTemplateId field if non-nil, zero value otherwise.

### GetUserRoleTemplateIdOk

`func (o *GatewayCreateProducerAzure) GetUserRoleTemplateIdOk() (*string, bool)`

GetUserRoleTemplateIdOk returns a tuple with the UserRoleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoleTemplateId

`func (o *GatewayCreateProducerAzure) SetUserRoleTemplateId(v string)`

SetUserRoleTemplateId sets UserRoleTemplateId field to given value.

### HasUserRoleTemplateId

`func (o *GatewayCreateProducerAzure) HasUserRoleTemplateId() bool`

HasUserRoleTemplateId returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerAzure) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerAzure) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerAzure) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerAzure) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayCreateProducerAzure) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateProducerAzure) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateProducerAzure) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateProducerAzure) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


