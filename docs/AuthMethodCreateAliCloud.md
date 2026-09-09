# AuthMethodCreateAliCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedClientType** | Pointer to **[]string** | limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension] | [optional] 
**AuditLogsClaims** | Pointer to **[]string** | Subclaims to include in audit logs, e.g \&quot;--audit-logs-claims email --audit-logs-claims username\&quot; | [optional] 
**BoundAccountId** | Pointer to **[]string** | A list of AliCloud account-IDs that the access is restricted to | [optional] 
**BoundArn** | Pointer to **[]string** | A list of full arns that the access is restricted to | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**BoundResourceId** | Pointer to **[]string** | A list of full resource ids that the access is restricted to | [optional] 
**BoundRoleId** | Pointer to **[]string** | A list of full role ids that the access is restricted to | [optional] 
**BoundRoleName** | Pointer to **[]string** | A list of full role-name that the access is restricted to | [optional] 
**BoundUserId** | Pointer to **[]string** | A list of full user ids that the access is restricted to | [optional] 
**BoundUserName** | Pointer to **[]string** | A list of full user-name that the access is restricted to | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the auth method would you like to be notified. | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca, ara] | [optional] 
**StsUrl** | Pointer to **string** | sts URL | [optional] [default to "https://sts.aliyuncs.com"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | Pointer to **string** | A unique identifier (ID) value which is a \&quot;sub claim\&quot; name that contains details uniquely identifying that resource. This \&quot;sub claim\&quot; is used to distinguish between different identities. | [optional] 

## Methods

### NewAuthMethodCreateAliCloud

`func NewAuthMethodCreateAliCloud(name string, ) *AuthMethodCreateAliCloud`

NewAuthMethodCreateAliCloud instantiates a new AuthMethodCreateAliCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodCreateAliCloudWithDefaults

`func NewAuthMethodCreateAliCloudWithDefaults() *AuthMethodCreateAliCloud`

NewAuthMethodCreateAliCloudWithDefaults instantiates a new AuthMethodCreateAliCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *AuthMethodCreateAliCloud) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *AuthMethodCreateAliCloud) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *AuthMethodCreateAliCloud) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *AuthMethodCreateAliCloud) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedClientType

`func (o *AuthMethodCreateAliCloud) GetAllowedClientType() []string`

GetAllowedClientType returns the AllowedClientType field if non-nil, zero value otherwise.

### GetAllowedClientTypeOk

`func (o *AuthMethodCreateAliCloud) GetAllowedClientTypeOk() (*[]string, bool)`

GetAllowedClientTypeOk returns a tuple with the AllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientType

`func (o *AuthMethodCreateAliCloud) SetAllowedClientType(v []string)`

SetAllowedClientType sets AllowedClientType field to given value.

### HasAllowedClientType

`func (o *AuthMethodCreateAliCloud) HasAllowedClientType() bool`

HasAllowedClientType returns a boolean if a field has been set.

### GetAuditLogsClaims

`func (o *AuthMethodCreateAliCloud) GetAuditLogsClaims() []string`

GetAuditLogsClaims returns the AuditLogsClaims field if non-nil, zero value otherwise.

### GetAuditLogsClaimsOk

`func (o *AuthMethodCreateAliCloud) GetAuditLogsClaimsOk() (*[]string, bool)`

GetAuditLogsClaimsOk returns a tuple with the AuditLogsClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsClaims

`func (o *AuthMethodCreateAliCloud) SetAuditLogsClaims(v []string)`

SetAuditLogsClaims sets AuditLogsClaims field to given value.

### HasAuditLogsClaims

`func (o *AuthMethodCreateAliCloud) HasAuditLogsClaims() bool`

HasAuditLogsClaims returns a boolean if a field has been set.

### GetBoundAccountId

`func (o *AuthMethodCreateAliCloud) GetBoundAccountId() []string`

GetBoundAccountId returns the BoundAccountId field if non-nil, zero value otherwise.

### GetBoundAccountIdOk

`func (o *AuthMethodCreateAliCloud) GetBoundAccountIdOk() (*[]string, bool)`

GetBoundAccountIdOk returns a tuple with the BoundAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAccountId

`func (o *AuthMethodCreateAliCloud) SetBoundAccountId(v []string)`

SetBoundAccountId sets BoundAccountId field to given value.

### HasBoundAccountId

`func (o *AuthMethodCreateAliCloud) HasBoundAccountId() bool`

HasBoundAccountId returns a boolean if a field has been set.

### GetBoundArn

`func (o *AuthMethodCreateAliCloud) GetBoundArn() []string`

GetBoundArn returns the BoundArn field if non-nil, zero value otherwise.

### GetBoundArnOk

`func (o *AuthMethodCreateAliCloud) GetBoundArnOk() (*[]string, bool)`

GetBoundArnOk returns a tuple with the BoundArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundArn

`func (o *AuthMethodCreateAliCloud) SetBoundArn(v []string)`

SetBoundArn sets BoundArn field to given value.

### HasBoundArn

`func (o *AuthMethodCreateAliCloud) HasBoundArn() bool`

HasBoundArn returns a boolean if a field has been set.

### GetBoundIps

`func (o *AuthMethodCreateAliCloud) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *AuthMethodCreateAliCloud) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *AuthMethodCreateAliCloud) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *AuthMethodCreateAliCloud) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundResourceId

`func (o *AuthMethodCreateAliCloud) GetBoundResourceId() []string`

GetBoundResourceId returns the BoundResourceId field if non-nil, zero value otherwise.

### GetBoundResourceIdOk

`func (o *AuthMethodCreateAliCloud) GetBoundResourceIdOk() (*[]string, bool)`

GetBoundResourceIdOk returns a tuple with the BoundResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceId

`func (o *AuthMethodCreateAliCloud) SetBoundResourceId(v []string)`

SetBoundResourceId sets BoundResourceId field to given value.

### HasBoundResourceId

`func (o *AuthMethodCreateAliCloud) HasBoundResourceId() bool`

HasBoundResourceId returns a boolean if a field has been set.

### GetBoundRoleId

`func (o *AuthMethodCreateAliCloud) GetBoundRoleId() []string`

GetBoundRoleId returns the BoundRoleId field if non-nil, zero value otherwise.

### GetBoundRoleIdOk

`func (o *AuthMethodCreateAliCloud) GetBoundRoleIdOk() (*[]string, bool)`

GetBoundRoleIdOk returns a tuple with the BoundRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoleId

`func (o *AuthMethodCreateAliCloud) SetBoundRoleId(v []string)`

SetBoundRoleId sets BoundRoleId field to given value.

### HasBoundRoleId

`func (o *AuthMethodCreateAliCloud) HasBoundRoleId() bool`

HasBoundRoleId returns a boolean if a field has been set.

### GetBoundRoleName

`func (o *AuthMethodCreateAliCloud) GetBoundRoleName() []string`

GetBoundRoleName returns the BoundRoleName field if non-nil, zero value otherwise.

### GetBoundRoleNameOk

`func (o *AuthMethodCreateAliCloud) GetBoundRoleNameOk() (*[]string, bool)`

GetBoundRoleNameOk returns a tuple with the BoundRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoleName

`func (o *AuthMethodCreateAliCloud) SetBoundRoleName(v []string)`

SetBoundRoleName sets BoundRoleName field to given value.

### HasBoundRoleName

`func (o *AuthMethodCreateAliCloud) HasBoundRoleName() bool`

HasBoundRoleName returns a boolean if a field has been set.

### GetBoundUserId

`func (o *AuthMethodCreateAliCloud) GetBoundUserId() []string`

GetBoundUserId returns the BoundUserId field if non-nil, zero value otherwise.

### GetBoundUserIdOk

`func (o *AuthMethodCreateAliCloud) GetBoundUserIdOk() (*[]string, bool)`

GetBoundUserIdOk returns a tuple with the BoundUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserId

`func (o *AuthMethodCreateAliCloud) SetBoundUserId(v []string)`

SetBoundUserId sets BoundUserId field to given value.

### HasBoundUserId

`func (o *AuthMethodCreateAliCloud) HasBoundUserId() bool`

HasBoundUserId returns a boolean if a field has been set.

### GetBoundUserName

`func (o *AuthMethodCreateAliCloud) GetBoundUserName() []string`

GetBoundUserName returns the BoundUserName field if non-nil, zero value otherwise.

### GetBoundUserNameOk

`func (o *AuthMethodCreateAliCloud) GetBoundUserNameOk() (*[]string, bool)`

GetBoundUserNameOk returns a tuple with the BoundUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserName

`func (o *AuthMethodCreateAliCloud) SetBoundUserName(v []string)`

SetBoundUserName sets BoundUserName field to given value.

### HasBoundUserName

`func (o *AuthMethodCreateAliCloud) HasBoundUserName() bool`

HasBoundUserName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *AuthMethodCreateAliCloud) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *AuthMethodCreateAliCloud) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *AuthMethodCreateAliCloud) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *AuthMethodCreateAliCloud) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *AuthMethodCreateAliCloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthMethodCreateAliCloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthMethodCreateAliCloud) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthMethodCreateAliCloud) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *AuthMethodCreateAliCloud) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *AuthMethodCreateAliCloud) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *AuthMethodCreateAliCloud) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *AuthMethodCreateAliCloud) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *AuthMethodCreateAliCloud) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *AuthMethodCreateAliCloud) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *AuthMethodCreateAliCloud) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *AuthMethodCreateAliCloud) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *AuthMethodCreateAliCloud) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *AuthMethodCreateAliCloud) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *AuthMethodCreateAliCloud) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *AuthMethodCreateAliCloud) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *AuthMethodCreateAliCloud) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AuthMethodCreateAliCloud) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AuthMethodCreateAliCloud) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AuthMethodCreateAliCloud) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *AuthMethodCreateAliCloud) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *AuthMethodCreateAliCloud) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *AuthMethodCreateAliCloud) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *AuthMethodCreateAliCloud) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *AuthMethodCreateAliCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthMethodCreateAliCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthMethodCreateAliCloud) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *AuthMethodCreateAliCloud) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuthMethodCreateAliCloud) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuthMethodCreateAliCloud) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuthMethodCreateAliCloud) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetStsUrl

`func (o *AuthMethodCreateAliCloud) GetStsUrl() string`

GetStsUrl returns the StsUrl field if non-nil, zero value otherwise.

### GetStsUrlOk

`func (o *AuthMethodCreateAliCloud) GetStsUrlOk() (*string, bool)`

GetStsUrlOk returns a tuple with the StsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsUrl

`func (o *AuthMethodCreateAliCloud) SetStsUrl(v string)`

SetStsUrl sets StsUrl field to given value.

### HasStsUrl

`func (o *AuthMethodCreateAliCloud) HasStsUrl() bool`

HasStsUrl returns a boolean if a field has been set.

### GetToken

`func (o *AuthMethodCreateAliCloud) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthMethodCreateAliCloud) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthMethodCreateAliCloud) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthMethodCreateAliCloud) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AuthMethodCreateAliCloud) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AuthMethodCreateAliCloud) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AuthMethodCreateAliCloud) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AuthMethodCreateAliCloud) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *AuthMethodCreateAliCloud) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *AuthMethodCreateAliCloud) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *AuthMethodCreateAliCloud) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *AuthMethodCreateAliCloud) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


