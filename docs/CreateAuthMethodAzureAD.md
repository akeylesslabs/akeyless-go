# CreateAuthMethodAzureAD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**Audience** | Pointer to **string** | The audience in the JWT | [optional] [default to "https://management.azure.com/"]
**BoundGroupId** | Pointer to **[]string** | A list of group ids that the access is restricted to | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist of the IPs that the access is restricted to | [optional] 
**BoundProviders** | Pointer to **[]string** | A list of resource providers that the access is restricted to (e.g, Microsoft.Compute, Microsoft.ManagedIdentity, etc) | [optional] 
**BoundResourceId** | Pointer to **[]string** | A list of full resource ids that the access is restricted to | [optional] 
**BoundResourceNames** | Pointer to **[]string** | A list of resource names that the access is restricted to (e.g, a virtual machine name, scale set name, etc). | [optional] 
**BoundResourceTypes** | Pointer to **[]string** | A list of resource types that the access is restricted to (e.g, virtualMachines, userAssignedIdentities, etc) | [optional] 
**BoundRgId** | Pointer to **[]string** | A list of resource groups that the access is restricted to | [optional] 
**BoundSpid** | Pointer to **[]string** | A list of service principal IDs that the access is restricted to | [optional] 
**BoundSubId** | Pointer to **[]string** | A list of subscription ids that the access is restricted to | [optional] 
**BoundTenantId** | **string** | The Azure tenant id that the access is restricted to | 
**Issuer** | Pointer to **string** | Issuer URL | [optional] [default to "https://sts.windows.net/---bound_tenant_id---"]
**JwksUri** | Pointer to **string** | The URL to the JSON Web Key Set (JWKS) that containing the public keys that should be used to verify any JSON Web Token (JWT) issued by the authorization server. | [optional] [default to "https://login.microsoftonline.com/common/discovery/keys"]
**Name** | **string** | Auth Method name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateAuthMethodAzureAD

`func NewCreateAuthMethodAzureAD(boundTenantId string, name string, ) *CreateAuthMethodAzureAD`

NewCreateAuthMethodAzureAD instantiates a new CreateAuthMethodAzureAD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodAzureADWithDefaults

`func NewCreateAuthMethodAzureADWithDefaults() *CreateAuthMethodAzureAD`

NewCreateAuthMethodAzureADWithDefaults instantiates a new CreateAuthMethodAzureAD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodAzureAD) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodAzureAD) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodAzureAD) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodAzureAD) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAudience

`func (o *CreateAuthMethodAzureAD) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CreateAuthMethodAzureAD) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CreateAuthMethodAzureAD) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CreateAuthMethodAzureAD) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBoundGroupId

`func (o *CreateAuthMethodAzureAD) GetBoundGroupId() []string`

GetBoundGroupId returns the BoundGroupId field if non-nil, zero value otherwise.

### GetBoundGroupIdOk

`func (o *CreateAuthMethodAzureAD) GetBoundGroupIdOk() (*[]string, bool)`

GetBoundGroupIdOk returns a tuple with the BoundGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundGroupId

`func (o *CreateAuthMethodAzureAD) SetBoundGroupId(v []string)`

SetBoundGroupId sets BoundGroupId field to given value.

### HasBoundGroupId

`func (o *CreateAuthMethodAzureAD) HasBoundGroupId() bool`

HasBoundGroupId returns a boolean if a field has been set.

### GetBoundIps

`func (o *CreateAuthMethodAzureAD) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodAzureAD) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodAzureAD) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodAzureAD) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundProviders

`func (o *CreateAuthMethodAzureAD) GetBoundProviders() []string`

GetBoundProviders returns the BoundProviders field if non-nil, zero value otherwise.

### GetBoundProvidersOk

`func (o *CreateAuthMethodAzureAD) GetBoundProvidersOk() (*[]string, bool)`

GetBoundProvidersOk returns a tuple with the BoundProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProviders

`func (o *CreateAuthMethodAzureAD) SetBoundProviders(v []string)`

SetBoundProviders sets BoundProviders field to given value.

### HasBoundProviders

`func (o *CreateAuthMethodAzureAD) HasBoundProviders() bool`

HasBoundProviders returns a boolean if a field has been set.

### GetBoundResourceId

`func (o *CreateAuthMethodAzureAD) GetBoundResourceId() []string`

GetBoundResourceId returns the BoundResourceId field if non-nil, zero value otherwise.

### GetBoundResourceIdOk

`func (o *CreateAuthMethodAzureAD) GetBoundResourceIdOk() (*[]string, bool)`

GetBoundResourceIdOk returns a tuple with the BoundResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceId

`func (o *CreateAuthMethodAzureAD) SetBoundResourceId(v []string)`

SetBoundResourceId sets BoundResourceId field to given value.

### HasBoundResourceId

`func (o *CreateAuthMethodAzureAD) HasBoundResourceId() bool`

HasBoundResourceId returns a boolean if a field has been set.

### GetBoundResourceNames

`func (o *CreateAuthMethodAzureAD) GetBoundResourceNames() []string`

GetBoundResourceNames returns the BoundResourceNames field if non-nil, zero value otherwise.

### GetBoundResourceNamesOk

`func (o *CreateAuthMethodAzureAD) GetBoundResourceNamesOk() (*[]string, bool)`

GetBoundResourceNamesOk returns a tuple with the BoundResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceNames

`func (o *CreateAuthMethodAzureAD) SetBoundResourceNames(v []string)`

SetBoundResourceNames sets BoundResourceNames field to given value.

### HasBoundResourceNames

`func (o *CreateAuthMethodAzureAD) HasBoundResourceNames() bool`

HasBoundResourceNames returns a boolean if a field has been set.

### GetBoundResourceTypes

`func (o *CreateAuthMethodAzureAD) GetBoundResourceTypes() []string`

GetBoundResourceTypes returns the BoundResourceTypes field if non-nil, zero value otherwise.

### GetBoundResourceTypesOk

`func (o *CreateAuthMethodAzureAD) GetBoundResourceTypesOk() (*[]string, bool)`

GetBoundResourceTypesOk returns a tuple with the BoundResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceTypes

`func (o *CreateAuthMethodAzureAD) SetBoundResourceTypes(v []string)`

SetBoundResourceTypes sets BoundResourceTypes field to given value.

### HasBoundResourceTypes

`func (o *CreateAuthMethodAzureAD) HasBoundResourceTypes() bool`

HasBoundResourceTypes returns a boolean if a field has been set.

### GetBoundRgId

`func (o *CreateAuthMethodAzureAD) GetBoundRgId() []string`

GetBoundRgId returns the BoundRgId field if non-nil, zero value otherwise.

### GetBoundRgIdOk

`func (o *CreateAuthMethodAzureAD) GetBoundRgIdOk() (*[]string, bool)`

GetBoundRgIdOk returns a tuple with the BoundRgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRgId

`func (o *CreateAuthMethodAzureAD) SetBoundRgId(v []string)`

SetBoundRgId sets BoundRgId field to given value.

### HasBoundRgId

`func (o *CreateAuthMethodAzureAD) HasBoundRgId() bool`

HasBoundRgId returns a boolean if a field has been set.

### GetBoundSpid

`func (o *CreateAuthMethodAzureAD) GetBoundSpid() []string`

GetBoundSpid returns the BoundSpid field if non-nil, zero value otherwise.

### GetBoundSpidOk

`func (o *CreateAuthMethodAzureAD) GetBoundSpidOk() (*[]string, bool)`

GetBoundSpidOk returns a tuple with the BoundSpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSpid

`func (o *CreateAuthMethodAzureAD) SetBoundSpid(v []string)`

SetBoundSpid sets BoundSpid field to given value.

### HasBoundSpid

`func (o *CreateAuthMethodAzureAD) HasBoundSpid() bool`

HasBoundSpid returns a boolean if a field has been set.

### GetBoundSubId

`func (o *CreateAuthMethodAzureAD) GetBoundSubId() []string`

GetBoundSubId returns the BoundSubId field if non-nil, zero value otherwise.

### GetBoundSubIdOk

`func (o *CreateAuthMethodAzureAD) GetBoundSubIdOk() (*[]string, bool)`

GetBoundSubIdOk returns a tuple with the BoundSubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSubId

`func (o *CreateAuthMethodAzureAD) SetBoundSubId(v []string)`

SetBoundSubId sets BoundSubId field to given value.

### HasBoundSubId

`func (o *CreateAuthMethodAzureAD) HasBoundSubId() bool`

HasBoundSubId returns a boolean if a field has been set.

### GetBoundTenantId

`func (o *CreateAuthMethodAzureAD) GetBoundTenantId() string`

GetBoundTenantId returns the BoundTenantId field if non-nil, zero value otherwise.

### GetBoundTenantIdOk

`func (o *CreateAuthMethodAzureAD) GetBoundTenantIdOk() (*string, bool)`

GetBoundTenantIdOk returns a tuple with the BoundTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTenantId

`func (o *CreateAuthMethodAzureAD) SetBoundTenantId(v string)`

SetBoundTenantId sets BoundTenantId field to given value.


### GetIssuer

`func (o *CreateAuthMethodAzureAD) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CreateAuthMethodAzureAD) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CreateAuthMethodAzureAD) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CreateAuthMethodAzureAD) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJwksUri

`func (o *CreateAuthMethodAzureAD) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *CreateAuthMethodAzureAD) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *CreateAuthMethodAzureAD) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *CreateAuthMethodAzureAD) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodAzureAD) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodAzureAD) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodAzureAD) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *CreateAuthMethodAzureAD) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodAzureAD) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodAzureAD) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodAzureAD) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodAzureAD) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodAzureAD) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodAzureAD) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodAzureAD) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


