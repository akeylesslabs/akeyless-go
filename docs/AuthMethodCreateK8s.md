# AuthMethodCreateK8s

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedClientType** | Pointer to **[]string** | limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension] | [optional] 
**Audience** | Pointer to **string** | The audience in the Kubernetes JWT that the access is restricted to | [optional] 
**AuditLogsClaims** | Pointer to **[]string** | Subclaims to include in audit logs, e.g \&quot;--audit-logs-claims email --audit-logs-claims username\&quot; | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**BoundNamespaces** | Pointer to **[]string** | A list of namespaces that the access is restricted to | [optional] 
**BoundPodNames** | Pointer to **[]string** | A list of pod names that the access is restricted to | [optional] 
**BoundSaNames** | Pointer to **[]string** | A list of service account names that the access is restricted to | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the auth method would you like to be notified. | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GenKey** | Pointer to **string** | Automatically generate key-pair for K8S configuration. If set to false, a public key needs to be provided [true/false] | [optional] [default to "true"]
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca] | [optional] 
**PublicKey** | Pointer to **string** | Base64-encoded or PEM formatted public key data for K8S authentication method is required [RSA2048] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewAuthMethodCreateK8s

`func NewAuthMethodCreateK8s(name string, ) *AuthMethodCreateK8s`

NewAuthMethodCreateK8s instantiates a new AuthMethodCreateK8s object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodCreateK8sWithDefaults

`func NewAuthMethodCreateK8sWithDefaults() *AuthMethodCreateK8s`

NewAuthMethodCreateK8sWithDefaults instantiates a new AuthMethodCreateK8s object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *AuthMethodCreateK8s) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *AuthMethodCreateK8s) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *AuthMethodCreateK8s) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *AuthMethodCreateK8s) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedClientType

`func (o *AuthMethodCreateK8s) GetAllowedClientType() []string`

GetAllowedClientType returns the AllowedClientType field if non-nil, zero value otherwise.

### GetAllowedClientTypeOk

`func (o *AuthMethodCreateK8s) GetAllowedClientTypeOk() (*[]string, bool)`

GetAllowedClientTypeOk returns a tuple with the AllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientType

`func (o *AuthMethodCreateK8s) SetAllowedClientType(v []string)`

SetAllowedClientType sets AllowedClientType field to given value.

### HasAllowedClientType

`func (o *AuthMethodCreateK8s) HasAllowedClientType() bool`

HasAllowedClientType returns a boolean if a field has been set.

### GetAudience

`func (o *AuthMethodCreateK8s) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *AuthMethodCreateK8s) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *AuthMethodCreateK8s) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *AuthMethodCreateK8s) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetAuditLogsClaims

`func (o *AuthMethodCreateK8s) GetAuditLogsClaims() []string`

GetAuditLogsClaims returns the AuditLogsClaims field if non-nil, zero value otherwise.

### GetAuditLogsClaimsOk

`func (o *AuthMethodCreateK8s) GetAuditLogsClaimsOk() (*[]string, bool)`

GetAuditLogsClaimsOk returns a tuple with the AuditLogsClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsClaims

`func (o *AuthMethodCreateK8s) SetAuditLogsClaims(v []string)`

SetAuditLogsClaims sets AuditLogsClaims field to given value.

### HasAuditLogsClaims

`func (o *AuthMethodCreateK8s) HasAuditLogsClaims() bool`

HasAuditLogsClaims returns a boolean if a field has been set.

### GetBoundIps

`func (o *AuthMethodCreateK8s) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *AuthMethodCreateK8s) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *AuthMethodCreateK8s) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *AuthMethodCreateK8s) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundNamespaces

`func (o *AuthMethodCreateK8s) GetBoundNamespaces() []string`

GetBoundNamespaces returns the BoundNamespaces field if non-nil, zero value otherwise.

### GetBoundNamespacesOk

`func (o *AuthMethodCreateK8s) GetBoundNamespacesOk() (*[]string, bool)`

GetBoundNamespacesOk returns a tuple with the BoundNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundNamespaces

`func (o *AuthMethodCreateK8s) SetBoundNamespaces(v []string)`

SetBoundNamespaces sets BoundNamespaces field to given value.

### HasBoundNamespaces

`func (o *AuthMethodCreateK8s) HasBoundNamespaces() bool`

HasBoundNamespaces returns a boolean if a field has been set.

### GetBoundPodNames

`func (o *AuthMethodCreateK8s) GetBoundPodNames() []string`

GetBoundPodNames returns the BoundPodNames field if non-nil, zero value otherwise.

### GetBoundPodNamesOk

`func (o *AuthMethodCreateK8s) GetBoundPodNamesOk() (*[]string, bool)`

GetBoundPodNamesOk returns a tuple with the BoundPodNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundPodNames

`func (o *AuthMethodCreateK8s) SetBoundPodNames(v []string)`

SetBoundPodNames sets BoundPodNames field to given value.

### HasBoundPodNames

`func (o *AuthMethodCreateK8s) HasBoundPodNames() bool`

HasBoundPodNames returns a boolean if a field has been set.

### GetBoundSaNames

`func (o *AuthMethodCreateK8s) GetBoundSaNames() []string`

GetBoundSaNames returns the BoundSaNames field if non-nil, zero value otherwise.

### GetBoundSaNamesOk

`func (o *AuthMethodCreateK8s) GetBoundSaNamesOk() (*[]string, bool)`

GetBoundSaNamesOk returns a tuple with the BoundSaNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSaNames

`func (o *AuthMethodCreateK8s) SetBoundSaNames(v []string)`

SetBoundSaNames sets BoundSaNames field to given value.

### HasBoundSaNames

`func (o *AuthMethodCreateK8s) HasBoundSaNames() bool`

HasBoundSaNames returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *AuthMethodCreateK8s) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *AuthMethodCreateK8s) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *AuthMethodCreateK8s) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *AuthMethodCreateK8s) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *AuthMethodCreateK8s) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthMethodCreateK8s) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthMethodCreateK8s) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthMethodCreateK8s) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *AuthMethodCreateK8s) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *AuthMethodCreateK8s) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *AuthMethodCreateK8s) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *AuthMethodCreateK8s) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *AuthMethodCreateK8s) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *AuthMethodCreateK8s) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *AuthMethodCreateK8s) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *AuthMethodCreateK8s) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGenKey

`func (o *AuthMethodCreateK8s) GetGenKey() string`

GetGenKey returns the GenKey field if non-nil, zero value otherwise.

### GetGenKeyOk

`func (o *AuthMethodCreateK8s) GetGenKeyOk() (*string, bool)`

GetGenKeyOk returns a tuple with the GenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenKey

`func (o *AuthMethodCreateK8s) SetGenKey(v string)`

SetGenKey sets GenKey field to given value.

### HasGenKey

`func (o *AuthMethodCreateK8s) HasGenKey() bool`

HasGenKey returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *AuthMethodCreateK8s) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *AuthMethodCreateK8s) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *AuthMethodCreateK8s) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *AuthMethodCreateK8s) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *AuthMethodCreateK8s) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AuthMethodCreateK8s) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AuthMethodCreateK8s) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AuthMethodCreateK8s) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *AuthMethodCreateK8s) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *AuthMethodCreateK8s) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *AuthMethodCreateK8s) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *AuthMethodCreateK8s) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *AuthMethodCreateK8s) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthMethodCreateK8s) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthMethodCreateK8s) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *AuthMethodCreateK8s) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuthMethodCreateK8s) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuthMethodCreateK8s) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuthMethodCreateK8s) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetPublicKey

`func (o *AuthMethodCreateK8s) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AuthMethodCreateK8s) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AuthMethodCreateK8s) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *AuthMethodCreateK8s) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetToken

`func (o *AuthMethodCreateK8s) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthMethodCreateK8s) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthMethodCreateK8s) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthMethodCreateK8s) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AuthMethodCreateK8s) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AuthMethodCreateK8s) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AuthMethodCreateK8s) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AuthMethodCreateK8s) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


