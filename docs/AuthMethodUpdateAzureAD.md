# AuthMethodUpdateAzureAD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedClientType** | Pointer to **[]string** | limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension] | [optional] 
**Audience** | Pointer to **string** | Deprecated (Deprecated) The audience in the JWT | [optional] [default to "https://management.azure.com/"]
**AuditLogsClaims** | Pointer to **[]string** | Subclaims to include in audit logs, e.g \&quot;--audit-logs-claims email --audit-logs-claims username\&quot; | [optional] 
**BoundGroupId** | Pointer to **[]string** | A list of group ids that the access is restricted to | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**BoundProviders** | Pointer to **[]string** | A list of resource providers that the access is restricted to (e.g, Microsoft.Compute, Microsoft.ManagedIdentity, etc) | [optional] 
**BoundResourceId** | Pointer to **[]string** | A list of full resource ids that the access is restricted to | [optional] 
**BoundResourceNames** | Pointer to **[]string** | A list of resource names that the access is restricted to (e.g, a virtual machine name, scale set name, etc). | [optional] 
**BoundResourceTypes** | Pointer to **[]string** | A list of resource types that the access is restricted to (e.g, virtualMachines, userAssignedIdentities, etc) | [optional] 
**BoundRgId** | Pointer to **[]string** | A list of resource groups that the access is restricted to | [optional] 
**BoundSpid** | Pointer to **[]string** | A list of service principal IDs that the access is restricted to | [optional] 
**BoundSubId** | Pointer to **[]string** | A list of subscription ids that the access is restricted to | [optional] 
**BoundTenantId** | **string** | The Azure tenant id that the access is restricted to | 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the auth method would you like to be notified. | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Issuer** | Pointer to **string** | Issuer URL | [optional] [default to "https://sts.windows.net/---bound_tenant_id---"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwksUri** | Pointer to **string** | The URL to the JSON Web Key Set (JWKS) that containing the public keys that should be used to verify any JSON Web Token (JWT) issued by the authorization server. | [optional] [default to "https://login.microsoftonline.com/common/discovery/keys"]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | Pointer to **string** | A unique identifier (ID) value which is a \&quot;sub claim\&quot; name that contains details uniquely identifying that resource. This \&quot;sub claim\&quot; is used to distinguish between different identities. | [optional] 

## Methods

### NewAuthMethodUpdateAzureAD

`func NewAuthMethodUpdateAzureAD(boundTenantId string, name string, ) *AuthMethodUpdateAzureAD`

NewAuthMethodUpdateAzureAD instantiates a new AuthMethodUpdateAzureAD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodUpdateAzureADWithDefaults

`func NewAuthMethodUpdateAzureADWithDefaults() *AuthMethodUpdateAzureAD`

NewAuthMethodUpdateAzureADWithDefaults instantiates a new AuthMethodUpdateAzureAD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *AuthMethodUpdateAzureAD) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *AuthMethodUpdateAzureAD) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *AuthMethodUpdateAzureAD) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *AuthMethodUpdateAzureAD) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedClientType

`func (o *AuthMethodUpdateAzureAD) GetAllowedClientType() []string`

GetAllowedClientType returns the AllowedClientType field if non-nil, zero value otherwise.

### GetAllowedClientTypeOk

`func (o *AuthMethodUpdateAzureAD) GetAllowedClientTypeOk() (*[]string, bool)`

GetAllowedClientTypeOk returns a tuple with the AllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientType

`func (o *AuthMethodUpdateAzureAD) SetAllowedClientType(v []string)`

SetAllowedClientType sets AllowedClientType field to given value.

### HasAllowedClientType

`func (o *AuthMethodUpdateAzureAD) HasAllowedClientType() bool`

HasAllowedClientType returns a boolean if a field has been set.

### GetAudience

`func (o *AuthMethodUpdateAzureAD) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *AuthMethodUpdateAzureAD) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *AuthMethodUpdateAzureAD) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *AuthMethodUpdateAzureAD) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetAuditLogsClaims

`func (o *AuthMethodUpdateAzureAD) GetAuditLogsClaims() []string`

GetAuditLogsClaims returns the AuditLogsClaims field if non-nil, zero value otherwise.

### GetAuditLogsClaimsOk

`func (o *AuthMethodUpdateAzureAD) GetAuditLogsClaimsOk() (*[]string, bool)`

GetAuditLogsClaimsOk returns a tuple with the AuditLogsClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsClaims

`func (o *AuthMethodUpdateAzureAD) SetAuditLogsClaims(v []string)`

SetAuditLogsClaims sets AuditLogsClaims field to given value.

### HasAuditLogsClaims

`func (o *AuthMethodUpdateAzureAD) HasAuditLogsClaims() bool`

HasAuditLogsClaims returns a boolean if a field has been set.

### GetBoundGroupId

`func (o *AuthMethodUpdateAzureAD) GetBoundGroupId() []string`

GetBoundGroupId returns the BoundGroupId field if non-nil, zero value otherwise.

### GetBoundGroupIdOk

`func (o *AuthMethodUpdateAzureAD) GetBoundGroupIdOk() (*[]string, bool)`

GetBoundGroupIdOk returns a tuple with the BoundGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundGroupId

`func (o *AuthMethodUpdateAzureAD) SetBoundGroupId(v []string)`

SetBoundGroupId sets BoundGroupId field to given value.

### HasBoundGroupId

`func (o *AuthMethodUpdateAzureAD) HasBoundGroupId() bool`

HasBoundGroupId returns a boolean if a field has been set.

### GetBoundIps

`func (o *AuthMethodUpdateAzureAD) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *AuthMethodUpdateAzureAD) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *AuthMethodUpdateAzureAD) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *AuthMethodUpdateAzureAD) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundProviders

`func (o *AuthMethodUpdateAzureAD) GetBoundProviders() []string`

GetBoundProviders returns the BoundProviders field if non-nil, zero value otherwise.

### GetBoundProvidersOk

`func (o *AuthMethodUpdateAzureAD) GetBoundProvidersOk() (*[]string, bool)`

GetBoundProvidersOk returns a tuple with the BoundProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProviders

`func (o *AuthMethodUpdateAzureAD) SetBoundProviders(v []string)`

SetBoundProviders sets BoundProviders field to given value.

### HasBoundProviders

`func (o *AuthMethodUpdateAzureAD) HasBoundProviders() bool`

HasBoundProviders returns a boolean if a field has been set.

### GetBoundResourceId

`func (o *AuthMethodUpdateAzureAD) GetBoundResourceId() []string`

GetBoundResourceId returns the BoundResourceId field if non-nil, zero value otherwise.

### GetBoundResourceIdOk

`func (o *AuthMethodUpdateAzureAD) GetBoundResourceIdOk() (*[]string, bool)`

GetBoundResourceIdOk returns a tuple with the BoundResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceId

`func (o *AuthMethodUpdateAzureAD) SetBoundResourceId(v []string)`

SetBoundResourceId sets BoundResourceId field to given value.

### HasBoundResourceId

`func (o *AuthMethodUpdateAzureAD) HasBoundResourceId() bool`

HasBoundResourceId returns a boolean if a field has been set.

### GetBoundResourceNames

`func (o *AuthMethodUpdateAzureAD) GetBoundResourceNames() []string`

GetBoundResourceNames returns the BoundResourceNames field if non-nil, zero value otherwise.

### GetBoundResourceNamesOk

`func (o *AuthMethodUpdateAzureAD) GetBoundResourceNamesOk() (*[]string, bool)`

GetBoundResourceNamesOk returns a tuple with the BoundResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceNames

`func (o *AuthMethodUpdateAzureAD) SetBoundResourceNames(v []string)`

SetBoundResourceNames sets BoundResourceNames field to given value.

### HasBoundResourceNames

`func (o *AuthMethodUpdateAzureAD) HasBoundResourceNames() bool`

HasBoundResourceNames returns a boolean if a field has been set.

### GetBoundResourceTypes

`func (o *AuthMethodUpdateAzureAD) GetBoundResourceTypes() []string`

GetBoundResourceTypes returns the BoundResourceTypes field if non-nil, zero value otherwise.

### GetBoundResourceTypesOk

`func (o *AuthMethodUpdateAzureAD) GetBoundResourceTypesOk() (*[]string, bool)`

GetBoundResourceTypesOk returns a tuple with the BoundResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceTypes

`func (o *AuthMethodUpdateAzureAD) SetBoundResourceTypes(v []string)`

SetBoundResourceTypes sets BoundResourceTypes field to given value.

### HasBoundResourceTypes

`func (o *AuthMethodUpdateAzureAD) HasBoundResourceTypes() bool`

HasBoundResourceTypes returns a boolean if a field has been set.

### GetBoundRgId

`func (o *AuthMethodUpdateAzureAD) GetBoundRgId() []string`

GetBoundRgId returns the BoundRgId field if non-nil, zero value otherwise.

### GetBoundRgIdOk

`func (o *AuthMethodUpdateAzureAD) GetBoundRgIdOk() (*[]string, bool)`

GetBoundRgIdOk returns a tuple with the BoundRgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRgId

`func (o *AuthMethodUpdateAzureAD) SetBoundRgId(v []string)`

SetBoundRgId sets BoundRgId field to given value.

### HasBoundRgId

`func (o *AuthMethodUpdateAzureAD) HasBoundRgId() bool`

HasBoundRgId returns a boolean if a field has been set.

### GetBoundSpid

`func (o *AuthMethodUpdateAzureAD) GetBoundSpid() []string`

GetBoundSpid returns the BoundSpid field if non-nil, zero value otherwise.

### GetBoundSpidOk

`func (o *AuthMethodUpdateAzureAD) GetBoundSpidOk() (*[]string, bool)`

GetBoundSpidOk returns a tuple with the BoundSpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSpid

`func (o *AuthMethodUpdateAzureAD) SetBoundSpid(v []string)`

SetBoundSpid sets BoundSpid field to given value.

### HasBoundSpid

`func (o *AuthMethodUpdateAzureAD) HasBoundSpid() bool`

HasBoundSpid returns a boolean if a field has been set.

### GetBoundSubId

`func (o *AuthMethodUpdateAzureAD) GetBoundSubId() []string`

GetBoundSubId returns the BoundSubId field if non-nil, zero value otherwise.

### GetBoundSubIdOk

`func (o *AuthMethodUpdateAzureAD) GetBoundSubIdOk() (*[]string, bool)`

GetBoundSubIdOk returns a tuple with the BoundSubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSubId

`func (o *AuthMethodUpdateAzureAD) SetBoundSubId(v []string)`

SetBoundSubId sets BoundSubId field to given value.

### HasBoundSubId

`func (o *AuthMethodUpdateAzureAD) HasBoundSubId() bool`

HasBoundSubId returns a boolean if a field has been set.

### GetBoundTenantId

`func (o *AuthMethodUpdateAzureAD) GetBoundTenantId() string`

GetBoundTenantId returns the BoundTenantId field if non-nil, zero value otherwise.

### GetBoundTenantIdOk

`func (o *AuthMethodUpdateAzureAD) GetBoundTenantIdOk() (*string, bool)`

GetBoundTenantIdOk returns a tuple with the BoundTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTenantId

`func (o *AuthMethodUpdateAzureAD) SetBoundTenantId(v string)`

SetBoundTenantId sets BoundTenantId field to given value.


### GetDeleteProtection

`func (o *AuthMethodUpdateAzureAD) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *AuthMethodUpdateAzureAD) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *AuthMethodUpdateAzureAD) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *AuthMethodUpdateAzureAD) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *AuthMethodUpdateAzureAD) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthMethodUpdateAzureAD) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthMethodUpdateAzureAD) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthMethodUpdateAzureAD) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *AuthMethodUpdateAzureAD) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *AuthMethodUpdateAzureAD) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *AuthMethodUpdateAzureAD) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *AuthMethodUpdateAzureAD) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *AuthMethodUpdateAzureAD) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *AuthMethodUpdateAzureAD) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *AuthMethodUpdateAzureAD) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *AuthMethodUpdateAzureAD) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *AuthMethodUpdateAzureAD) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *AuthMethodUpdateAzureAD) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *AuthMethodUpdateAzureAD) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *AuthMethodUpdateAzureAD) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetIssuer

`func (o *AuthMethodUpdateAzureAD) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AuthMethodUpdateAzureAD) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AuthMethodUpdateAzureAD) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AuthMethodUpdateAzureAD) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJson

`func (o *AuthMethodUpdateAzureAD) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AuthMethodUpdateAzureAD) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AuthMethodUpdateAzureAD) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AuthMethodUpdateAzureAD) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwksUri

`func (o *AuthMethodUpdateAzureAD) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *AuthMethodUpdateAzureAD) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *AuthMethodUpdateAzureAD) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *AuthMethodUpdateAzureAD) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetJwtTtl

`func (o *AuthMethodUpdateAzureAD) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *AuthMethodUpdateAzureAD) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *AuthMethodUpdateAzureAD) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *AuthMethodUpdateAzureAD) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *AuthMethodUpdateAzureAD) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthMethodUpdateAzureAD) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthMethodUpdateAzureAD) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *AuthMethodUpdateAzureAD) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *AuthMethodUpdateAzureAD) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *AuthMethodUpdateAzureAD) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *AuthMethodUpdateAzureAD) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProductType

`func (o *AuthMethodUpdateAzureAD) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuthMethodUpdateAzureAD) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuthMethodUpdateAzureAD) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuthMethodUpdateAzureAD) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetToken

`func (o *AuthMethodUpdateAzureAD) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthMethodUpdateAzureAD) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthMethodUpdateAzureAD) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthMethodUpdateAzureAD) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AuthMethodUpdateAzureAD) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AuthMethodUpdateAzureAD) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AuthMethodUpdateAzureAD) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AuthMethodUpdateAzureAD) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *AuthMethodUpdateAzureAD) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *AuthMethodUpdateAzureAD) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *AuthMethodUpdateAzureAD) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *AuthMethodUpdateAzureAD) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


