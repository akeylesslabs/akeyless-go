# UpdateAuthMethodAzureAD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**Audience** | Pointer to **string** | Deprecated (Deprecated) The audience in the JWT | [optional] [default to "https://management.azure.com/"]
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
**Description** | Pointer to **string** | Auth Method description | [optional] 
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

## Methods

### NewUpdateAuthMethodAzureAD

`func NewUpdateAuthMethodAzureAD(boundTenantId string, name string, ) *UpdateAuthMethodAzureAD`

NewUpdateAuthMethodAzureAD instantiates a new UpdateAuthMethodAzureAD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthMethodAzureADWithDefaults

`func NewUpdateAuthMethodAzureADWithDefaults() *UpdateAuthMethodAzureAD`

NewUpdateAuthMethodAzureADWithDefaults instantiates a new UpdateAuthMethodAzureAD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *UpdateAuthMethodAzureAD) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *UpdateAuthMethodAzureAD) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *UpdateAuthMethodAzureAD) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *UpdateAuthMethodAzureAD) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAudience

`func (o *UpdateAuthMethodAzureAD) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *UpdateAuthMethodAzureAD) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *UpdateAuthMethodAzureAD) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *UpdateAuthMethodAzureAD) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBoundGroupId

`func (o *UpdateAuthMethodAzureAD) GetBoundGroupId() []string`

GetBoundGroupId returns the BoundGroupId field if non-nil, zero value otherwise.

### GetBoundGroupIdOk

`func (o *UpdateAuthMethodAzureAD) GetBoundGroupIdOk() (*[]string, bool)`

GetBoundGroupIdOk returns a tuple with the BoundGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundGroupId

`func (o *UpdateAuthMethodAzureAD) SetBoundGroupId(v []string)`

SetBoundGroupId sets BoundGroupId field to given value.

### HasBoundGroupId

`func (o *UpdateAuthMethodAzureAD) HasBoundGroupId() bool`

HasBoundGroupId returns a boolean if a field has been set.

### GetBoundIps

`func (o *UpdateAuthMethodAzureAD) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAuthMethodAzureAD) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAuthMethodAzureAD) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAuthMethodAzureAD) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundProviders

`func (o *UpdateAuthMethodAzureAD) GetBoundProviders() []string`

GetBoundProviders returns the BoundProviders field if non-nil, zero value otherwise.

### GetBoundProvidersOk

`func (o *UpdateAuthMethodAzureAD) GetBoundProvidersOk() (*[]string, bool)`

GetBoundProvidersOk returns a tuple with the BoundProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProviders

`func (o *UpdateAuthMethodAzureAD) SetBoundProviders(v []string)`

SetBoundProviders sets BoundProviders field to given value.

### HasBoundProviders

`func (o *UpdateAuthMethodAzureAD) HasBoundProviders() bool`

HasBoundProviders returns a boolean if a field has been set.

### GetBoundResourceId

`func (o *UpdateAuthMethodAzureAD) GetBoundResourceId() []string`

GetBoundResourceId returns the BoundResourceId field if non-nil, zero value otherwise.

### GetBoundResourceIdOk

`func (o *UpdateAuthMethodAzureAD) GetBoundResourceIdOk() (*[]string, bool)`

GetBoundResourceIdOk returns a tuple with the BoundResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceId

`func (o *UpdateAuthMethodAzureAD) SetBoundResourceId(v []string)`

SetBoundResourceId sets BoundResourceId field to given value.

### HasBoundResourceId

`func (o *UpdateAuthMethodAzureAD) HasBoundResourceId() bool`

HasBoundResourceId returns a boolean if a field has been set.

### GetBoundResourceNames

`func (o *UpdateAuthMethodAzureAD) GetBoundResourceNames() []string`

GetBoundResourceNames returns the BoundResourceNames field if non-nil, zero value otherwise.

### GetBoundResourceNamesOk

`func (o *UpdateAuthMethodAzureAD) GetBoundResourceNamesOk() (*[]string, bool)`

GetBoundResourceNamesOk returns a tuple with the BoundResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceNames

`func (o *UpdateAuthMethodAzureAD) SetBoundResourceNames(v []string)`

SetBoundResourceNames sets BoundResourceNames field to given value.

### HasBoundResourceNames

`func (o *UpdateAuthMethodAzureAD) HasBoundResourceNames() bool`

HasBoundResourceNames returns a boolean if a field has been set.

### GetBoundResourceTypes

`func (o *UpdateAuthMethodAzureAD) GetBoundResourceTypes() []string`

GetBoundResourceTypes returns the BoundResourceTypes field if non-nil, zero value otherwise.

### GetBoundResourceTypesOk

`func (o *UpdateAuthMethodAzureAD) GetBoundResourceTypesOk() (*[]string, bool)`

GetBoundResourceTypesOk returns a tuple with the BoundResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceTypes

`func (o *UpdateAuthMethodAzureAD) SetBoundResourceTypes(v []string)`

SetBoundResourceTypes sets BoundResourceTypes field to given value.

### HasBoundResourceTypes

`func (o *UpdateAuthMethodAzureAD) HasBoundResourceTypes() bool`

HasBoundResourceTypes returns a boolean if a field has been set.

### GetBoundRgId

`func (o *UpdateAuthMethodAzureAD) GetBoundRgId() []string`

GetBoundRgId returns the BoundRgId field if non-nil, zero value otherwise.

### GetBoundRgIdOk

`func (o *UpdateAuthMethodAzureAD) GetBoundRgIdOk() (*[]string, bool)`

GetBoundRgIdOk returns a tuple with the BoundRgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRgId

`func (o *UpdateAuthMethodAzureAD) SetBoundRgId(v []string)`

SetBoundRgId sets BoundRgId field to given value.

### HasBoundRgId

`func (o *UpdateAuthMethodAzureAD) HasBoundRgId() bool`

HasBoundRgId returns a boolean if a field has been set.

### GetBoundSpid

`func (o *UpdateAuthMethodAzureAD) GetBoundSpid() []string`

GetBoundSpid returns the BoundSpid field if non-nil, zero value otherwise.

### GetBoundSpidOk

`func (o *UpdateAuthMethodAzureAD) GetBoundSpidOk() (*[]string, bool)`

GetBoundSpidOk returns a tuple with the BoundSpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSpid

`func (o *UpdateAuthMethodAzureAD) SetBoundSpid(v []string)`

SetBoundSpid sets BoundSpid field to given value.

### HasBoundSpid

`func (o *UpdateAuthMethodAzureAD) HasBoundSpid() bool`

HasBoundSpid returns a boolean if a field has been set.

### GetBoundSubId

`func (o *UpdateAuthMethodAzureAD) GetBoundSubId() []string`

GetBoundSubId returns the BoundSubId field if non-nil, zero value otherwise.

### GetBoundSubIdOk

`func (o *UpdateAuthMethodAzureAD) GetBoundSubIdOk() (*[]string, bool)`

GetBoundSubIdOk returns a tuple with the BoundSubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSubId

`func (o *UpdateAuthMethodAzureAD) SetBoundSubId(v []string)`

SetBoundSubId sets BoundSubId field to given value.

### HasBoundSubId

`func (o *UpdateAuthMethodAzureAD) HasBoundSubId() bool`

HasBoundSubId returns a boolean if a field has been set.

### GetBoundTenantId

`func (o *UpdateAuthMethodAzureAD) GetBoundTenantId() string`

GetBoundTenantId returns the BoundTenantId field if non-nil, zero value otherwise.

### GetBoundTenantIdOk

`func (o *UpdateAuthMethodAzureAD) GetBoundTenantIdOk() (*string, bool)`

GetBoundTenantIdOk returns a tuple with the BoundTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTenantId

`func (o *UpdateAuthMethodAzureAD) SetBoundTenantId(v string)`

SetBoundTenantId sets BoundTenantId field to given value.


### GetDescription

`func (o *UpdateAuthMethodAzureAD) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAuthMethodAzureAD) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAuthMethodAzureAD) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAuthMethodAzureAD) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *UpdateAuthMethodAzureAD) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *UpdateAuthMethodAzureAD) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *UpdateAuthMethodAzureAD) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *UpdateAuthMethodAzureAD) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *UpdateAuthMethodAzureAD) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *UpdateAuthMethodAzureAD) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *UpdateAuthMethodAzureAD) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *UpdateAuthMethodAzureAD) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetIssuer

`func (o *UpdateAuthMethodAzureAD) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *UpdateAuthMethodAzureAD) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *UpdateAuthMethodAzureAD) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *UpdateAuthMethodAzureAD) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJson

`func (o *UpdateAuthMethodAzureAD) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateAuthMethodAzureAD) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateAuthMethodAzureAD) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateAuthMethodAzureAD) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwksUri

`func (o *UpdateAuthMethodAzureAD) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *UpdateAuthMethodAzureAD) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *UpdateAuthMethodAzureAD) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *UpdateAuthMethodAzureAD) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetJwtTtl

`func (o *UpdateAuthMethodAzureAD) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *UpdateAuthMethodAzureAD) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *UpdateAuthMethodAzureAD) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *UpdateAuthMethodAzureAD) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *UpdateAuthMethodAzureAD) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthMethodAzureAD) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthMethodAzureAD) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateAuthMethodAzureAD) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateAuthMethodAzureAD) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateAuthMethodAzureAD) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateAuthMethodAzureAD) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProductType

`func (o *UpdateAuthMethodAzureAD) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *UpdateAuthMethodAzureAD) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *UpdateAuthMethodAzureAD) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *UpdateAuthMethodAzureAD) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAuthMethodAzureAD) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAuthMethodAzureAD) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAuthMethodAzureAD) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAuthMethodAzureAD) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAuthMethodAzureAD) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAuthMethodAzureAD) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAuthMethodAzureAD) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAuthMethodAzureAD) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


