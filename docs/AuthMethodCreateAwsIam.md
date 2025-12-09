# AuthMethodCreateAwsIam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedClientType** | Pointer to **[]string** | limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension] | [optional] 
**AuditLogsClaims** | Pointer to **[]string** | Subclaims to include in audit logs, e.g \&quot;--audit-logs-claims email --audit-logs-claims username\&quot; | [optional] 
**BoundArn** | Pointer to **[]string** | A list of full arns that the access is restricted to | [optional] 
**BoundAwsAccountId** | **[]string** | A list of AWS account-IDs that the access is restricted to | 
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
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca] | [optional] 
**StsUrl** | Pointer to **string** | sts URL | [optional] [default to "https://sts.amazonaws.com"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | Pointer to **string** | A unique identifier (ID) value which is a \&quot;sub claim\&quot; name that contains details uniquely identifying that resource. This \&quot;sub claim\&quot; is used to distinguish between different identities. | [optional] 

## Methods

### NewAuthMethodCreateAwsIam

`func NewAuthMethodCreateAwsIam(boundAwsAccountId []string, name string, ) *AuthMethodCreateAwsIam`

NewAuthMethodCreateAwsIam instantiates a new AuthMethodCreateAwsIam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodCreateAwsIamWithDefaults

`func NewAuthMethodCreateAwsIamWithDefaults() *AuthMethodCreateAwsIam`

NewAuthMethodCreateAwsIamWithDefaults instantiates a new AuthMethodCreateAwsIam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *AuthMethodCreateAwsIam) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *AuthMethodCreateAwsIam) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *AuthMethodCreateAwsIam) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *AuthMethodCreateAwsIam) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedClientType

`func (o *AuthMethodCreateAwsIam) GetAllowedClientType() []string`

GetAllowedClientType returns the AllowedClientType field if non-nil, zero value otherwise.

### GetAllowedClientTypeOk

`func (o *AuthMethodCreateAwsIam) GetAllowedClientTypeOk() (*[]string, bool)`

GetAllowedClientTypeOk returns a tuple with the AllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientType

`func (o *AuthMethodCreateAwsIam) SetAllowedClientType(v []string)`

SetAllowedClientType sets AllowedClientType field to given value.

### HasAllowedClientType

`func (o *AuthMethodCreateAwsIam) HasAllowedClientType() bool`

HasAllowedClientType returns a boolean if a field has been set.

### GetAuditLogsClaims

`func (o *AuthMethodCreateAwsIam) GetAuditLogsClaims() []string`

GetAuditLogsClaims returns the AuditLogsClaims field if non-nil, zero value otherwise.

### GetAuditLogsClaimsOk

`func (o *AuthMethodCreateAwsIam) GetAuditLogsClaimsOk() (*[]string, bool)`

GetAuditLogsClaimsOk returns a tuple with the AuditLogsClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsClaims

`func (o *AuthMethodCreateAwsIam) SetAuditLogsClaims(v []string)`

SetAuditLogsClaims sets AuditLogsClaims field to given value.

### HasAuditLogsClaims

`func (o *AuthMethodCreateAwsIam) HasAuditLogsClaims() bool`

HasAuditLogsClaims returns a boolean if a field has been set.

### GetBoundArn

`func (o *AuthMethodCreateAwsIam) GetBoundArn() []string`

GetBoundArn returns the BoundArn field if non-nil, zero value otherwise.

### GetBoundArnOk

`func (o *AuthMethodCreateAwsIam) GetBoundArnOk() (*[]string, bool)`

GetBoundArnOk returns a tuple with the BoundArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundArn

`func (o *AuthMethodCreateAwsIam) SetBoundArn(v []string)`

SetBoundArn sets BoundArn field to given value.

### HasBoundArn

`func (o *AuthMethodCreateAwsIam) HasBoundArn() bool`

HasBoundArn returns a boolean if a field has been set.

### GetBoundAwsAccountId

`func (o *AuthMethodCreateAwsIam) GetBoundAwsAccountId() []string`

GetBoundAwsAccountId returns the BoundAwsAccountId field if non-nil, zero value otherwise.

### GetBoundAwsAccountIdOk

`func (o *AuthMethodCreateAwsIam) GetBoundAwsAccountIdOk() (*[]string, bool)`

GetBoundAwsAccountIdOk returns a tuple with the BoundAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAwsAccountId

`func (o *AuthMethodCreateAwsIam) SetBoundAwsAccountId(v []string)`

SetBoundAwsAccountId sets BoundAwsAccountId field to given value.


### GetBoundIps

`func (o *AuthMethodCreateAwsIam) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *AuthMethodCreateAwsIam) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *AuthMethodCreateAwsIam) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *AuthMethodCreateAwsIam) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundResourceId

`func (o *AuthMethodCreateAwsIam) GetBoundResourceId() []string`

GetBoundResourceId returns the BoundResourceId field if non-nil, zero value otherwise.

### GetBoundResourceIdOk

`func (o *AuthMethodCreateAwsIam) GetBoundResourceIdOk() (*[]string, bool)`

GetBoundResourceIdOk returns a tuple with the BoundResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceId

`func (o *AuthMethodCreateAwsIam) SetBoundResourceId(v []string)`

SetBoundResourceId sets BoundResourceId field to given value.

### HasBoundResourceId

`func (o *AuthMethodCreateAwsIam) HasBoundResourceId() bool`

HasBoundResourceId returns a boolean if a field has been set.

### GetBoundRoleId

`func (o *AuthMethodCreateAwsIam) GetBoundRoleId() []string`

GetBoundRoleId returns the BoundRoleId field if non-nil, zero value otherwise.

### GetBoundRoleIdOk

`func (o *AuthMethodCreateAwsIam) GetBoundRoleIdOk() (*[]string, bool)`

GetBoundRoleIdOk returns a tuple with the BoundRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoleId

`func (o *AuthMethodCreateAwsIam) SetBoundRoleId(v []string)`

SetBoundRoleId sets BoundRoleId field to given value.

### HasBoundRoleId

`func (o *AuthMethodCreateAwsIam) HasBoundRoleId() bool`

HasBoundRoleId returns a boolean if a field has been set.

### GetBoundRoleName

`func (o *AuthMethodCreateAwsIam) GetBoundRoleName() []string`

GetBoundRoleName returns the BoundRoleName field if non-nil, zero value otherwise.

### GetBoundRoleNameOk

`func (o *AuthMethodCreateAwsIam) GetBoundRoleNameOk() (*[]string, bool)`

GetBoundRoleNameOk returns a tuple with the BoundRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoleName

`func (o *AuthMethodCreateAwsIam) SetBoundRoleName(v []string)`

SetBoundRoleName sets BoundRoleName field to given value.

### HasBoundRoleName

`func (o *AuthMethodCreateAwsIam) HasBoundRoleName() bool`

HasBoundRoleName returns a boolean if a field has been set.

### GetBoundUserId

`func (o *AuthMethodCreateAwsIam) GetBoundUserId() []string`

GetBoundUserId returns the BoundUserId field if non-nil, zero value otherwise.

### GetBoundUserIdOk

`func (o *AuthMethodCreateAwsIam) GetBoundUserIdOk() (*[]string, bool)`

GetBoundUserIdOk returns a tuple with the BoundUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserId

`func (o *AuthMethodCreateAwsIam) SetBoundUserId(v []string)`

SetBoundUserId sets BoundUserId field to given value.

### HasBoundUserId

`func (o *AuthMethodCreateAwsIam) HasBoundUserId() bool`

HasBoundUserId returns a boolean if a field has been set.

### GetBoundUserName

`func (o *AuthMethodCreateAwsIam) GetBoundUserName() []string`

GetBoundUserName returns the BoundUserName field if non-nil, zero value otherwise.

### GetBoundUserNameOk

`func (o *AuthMethodCreateAwsIam) GetBoundUserNameOk() (*[]string, bool)`

GetBoundUserNameOk returns a tuple with the BoundUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserName

`func (o *AuthMethodCreateAwsIam) SetBoundUserName(v []string)`

SetBoundUserName sets BoundUserName field to given value.

### HasBoundUserName

`func (o *AuthMethodCreateAwsIam) HasBoundUserName() bool`

HasBoundUserName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *AuthMethodCreateAwsIam) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *AuthMethodCreateAwsIam) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *AuthMethodCreateAwsIam) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *AuthMethodCreateAwsIam) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *AuthMethodCreateAwsIam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthMethodCreateAwsIam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthMethodCreateAwsIam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthMethodCreateAwsIam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *AuthMethodCreateAwsIam) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *AuthMethodCreateAwsIam) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *AuthMethodCreateAwsIam) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *AuthMethodCreateAwsIam) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *AuthMethodCreateAwsIam) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *AuthMethodCreateAwsIam) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *AuthMethodCreateAwsIam) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *AuthMethodCreateAwsIam) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *AuthMethodCreateAwsIam) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *AuthMethodCreateAwsIam) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *AuthMethodCreateAwsIam) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *AuthMethodCreateAwsIam) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *AuthMethodCreateAwsIam) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AuthMethodCreateAwsIam) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AuthMethodCreateAwsIam) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AuthMethodCreateAwsIam) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *AuthMethodCreateAwsIam) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *AuthMethodCreateAwsIam) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *AuthMethodCreateAwsIam) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *AuthMethodCreateAwsIam) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *AuthMethodCreateAwsIam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthMethodCreateAwsIam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthMethodCreateAwsIam) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *AuthMethodCreateAwsIam) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuthMethodCreateAwsIam) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuthMethodCreateAwsIam) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuthMethodCreateAwsIam) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetStsUrl

`func (o *AuthMethodCreateAwsIam) GetStsUrl() string`

GetStsUrl returns the StsUrl field if non-nil, zero value otherwise.

### GetStsUrlOk

`func (o *AuthMethodCreateAwsIam) GetStsUrlOk() (*string, bool)`

GetStsUrlOk returns a tuple with the StsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsUrl

`func (o *AuthMethodCreateAwsIam) SetStsUrl(v string)`

SetStsUrl sets StsUrl field to given value.

### HasStsUrl

`func (o *AuthMethodCreateAwsIam) HasStsUrl() bool`

HasStsUrl returns a boolean if a field has been set.

### GetToken

`func (o *AuthMethodCreateAwsIam) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthMethodCreateAwsIam) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthMethodCreateAwsIam) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthMethodCreateAwsIam) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AuthMethodCreateAwsIam) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AuthMethodCreateAwsIam) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AuthMethodCreateAwsIam) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AuthMethodCreateAwsIam) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *AuthMethodCreateAwsIam) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *AuthMethodCreateAwsIam) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *AuthMethodCreateAwsIam) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *AuthMethodCreateAwsIam) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


