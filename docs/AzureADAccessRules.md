# AzureADAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdEndpoint** | Pointer to **string** | The audience in the JWT. | [optional] 
**BoundGroupIds** | Pointer to **[]string** | The list of group ids that login is restricted to. | [optional] 
**BoundResourceGroups** | Pointer to **[]string** | The list of resource groups that login is restricted to. | [optional] 
**BoundResourceIds** | Pointer to **[]string** | The list of full resource ids that the login is restricted to. | [optional] 
**BoundResourceNames** | Pointer to **[]string** | The list of resource names that the login is restricted to (e.g, a virtual machine name, scale set name, etc). | [optional] 
**BoundResourceProviders** | Pointer to **[]string** | The list of resource providers that login is restricted to (e.g, Microsoft.Compute, Microsoft.ManagedIdentity, etc). | [optional] 
**BoundResourceTypes** | Pointer to **[]string** | The list of resource types that login is restricted to  (e.g, virtualMachines, userAssignedIdentities, etc). | [optional] 
**BoundServicePrincipalIds** | Pointer to **[]string** | The list of service principal IDs that login is restricted to. | [optional] 
**BoundSubscriptionIds** | Pointer to **[]string** | The list of subscription IDs that login is restricted to. | [optional] 
**BoundTenantId** | Pointer to **string** | The tenants id for the Azure Active Directory organization. | [optional] 
**Issuer** | Pointer to **string** | Issuer URL | [optional] 
**JwksUri** | Pointer to **string** | The URL to the JSON Web Key Set (JWKS) that containing the public keys that should be used to verify any JSON Web Token (JWT) issued by the authorization server. | [optional] 

## Methods

### NewAzureADAccessRules

`func NewAzureADAccessRules() *AzureADAccessRules`

NewAzureADAccessRules instantiates a new AzureADAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureADAccessRulesWithDefaults

`func NewAzureADAccessRulesWithDefaults() *AzureADAccessRules`

NewAzureADAccessRulesWithDefaults instantiates a new AzureADAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdEndpoint

`func (o *AzureADAccessRules) GetAdEndpoint() string`

GetAdEndpoint returns the AdEndpoint field if non-nil, zero value otherwise.

### GetAdEndpointOk

`func (o *AzureADAccessRules) GetAdEndpointOk() (*string, bool)`

GetAdEndpointOk returns a tuple with the AdEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdEndpoint

`func (o *AzureADAccessRules) SetAdEndpoint(v string)`

SetAdEndpoint sets AdEndpoint field to given value.

### HasAdEndpoint

`func (o *AzureADAccessRules) HasAdEndpoint() bool`

HasAdEndpoint returns a boolean if a field has been set.

### GetBoundGroupIds

`func (o *AzureADAccessRules) GetBoundGroupIds() []string`

GetBoundGroupIds returns the BoundGroupIds field if non-nil, zero value otherwise.

### GetBoundGroupIdsOk

`func (o *AzureADAccessRules) GetBoundGroupIdsOk() (*[]string, bool)`

GetBoundGroupIdsOk returns a tuple with the BoundGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundGroupIds

`func (o *AzureADAccessRules) SetBoundGroupIds(v []string)`

SetBoundGroupIds sets BoundGroupIds field to given value.

### HasBoundGroupIds

`func (o *AzureADAccessRules) HasBoundGroupIds() bool`

HasBoundGroupIds returns a boolean if a field has been set.

### GetBoundResourceGroups

`func (o *AzureADAccessRules) GetBoundResourceGroups() []string`

GetBoundResourceGroups returns the BoundResourceGroups field if non-nil, zero value otherwise.

### GetBoundResourceGroupsOk

`func (o *AzureADAccessRules) GetBoundResourceGroupsOk() (*[]string, bool)`

GetBoundResourceGroupsOk returns a tuple with the BoundResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceGroups

`func (o *AzureADAccessRules) SetBoundResourceGroups(v []string)`

SetBoundResourceGroups sets BoundResourceGroups field to given value.

### HasBoundResourceGroups

`func (o *AzureADAccessRules) HasBoundResourceGroups() bool`

HasBoundResourceGroups returns a boolean if a field has been set.

### GetBoundResourceIds

`func (o *AzureADAccessRules) GetBoundResourceIds() []string`

GetBoundResourceIds returns the BoundResourceIds field if non-nil, zero value otherwise.

### GetBoundResourceIdsOk

`func (o *AzureADAccessRules) GetBoundResourceIdsOk() (*[]string, bool)`

GetBoundResourceIdsOk returns a tuple with the BoundResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceIds

`func (o *AzureADAccessRules) SetBoundResourceIds(v []string)`

SetBoundResourceIds sets BoundResourceIds field to given value.

### HasBoundResourceIds

`func (o *AzureADAccessRules) HasBoundResourceIds() bool`

HasBoundResourceIds returns a boolean if a field has been set.

### GetBoundResourceNames

`func (o *AzureADAccessRules) GetBoundResourceNames() []string`

GetBoundResourceNames returns the BoundResourceNames field if non-nil, zero value otherwise.

### GetBoundResourceNamesOk

`func (o *AzureADAccessRules) GetBoundResourceNamesOk() (*[]string, bool)`

GetBoundResourceNamesOk returns a tuple with the BoundResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceNames

`func (o *AzureADAccessRules) SetBoundResourceNames(v []string)`

SetBoundResourceNames sets BoundResourceNames field to given value.

### HasBoundResourceNames

`func (o *AzureADAccessRules) HasBoundResourceNames() bool`

HasBoundResourceNames returns a boolean if a field has been set.

### GetBoundResourceProviders

`func (o *AzureADAccessRules) GetBoundResourceProviders() []string`

GetBoundResourceProviders returns the BoundResourceProviders field if non-nil, zero value otherwise.

### GetBoundResourceProvidersOk

`func (o *AzureADAccessRules) GetBoundResourceProvidersOk() (*[]string, bool)`

GetBoundResourceProvidersOk returns a tuple with the BoundResourceProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceProviders

`func (o *AzureADAccessRules) SetBoundResourceProviders(v []string)`

SetBoundResourceProviders sets BoundResourceProviders field to given value.

### HasBoundResourceProviders

`func (o *AzureADAccessRules) HasBoundResourceProviders() bool`

HasBoundResourceProviders returns a boolean if a field has been set.

### GetBoundResourceTypes

`func (o *AzureADAccessRules) GetBoundResourceTypes() []string`

GetBoundResourceTypes returns the BoundResourceTypes field if non-nil, zero value otherwise.

### GetBoundResourceTypesOk

`func (o *AzureADAccessRules) GetBoundResourceTypesOk() (*[]string, bool)`

GetBoundResourceTypesOk returns a tuple with the BoundResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceTypes

`func (o *AzureADAccessRules) SetBoundResourceTypes(v []string)`

SetBoundResourceTypes sets BoundResourceTypes field to given value.

### HasBoundResourceTypes

`func (o *AzureADAccessRules) HasBoundResourceTypes() bool`

HasBoundResourceTypes returns a boolean if a field has been set.

### GetBoundServicePrincipalIds

`func (o *AzureADAccessRules) GetBoundServicePrincipalIds() []string`

GetBoundServicePrincipalIds returns the BoundServicePrincipalIds field if non-nil, zero value otherwise.

### GetBoundServicePrincipalIdsOk

`func (o *AzureADAccessRules) GetBoundServicePrincipalIdsOk() (*[]string, bool)`

GetBoundServicePrincipalIdsOk returns a tuple with the BoundServicePrincipalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServicePrincipalIds

`func (o *AzureADAccessRules) SetBoundServicePrincipalIds(v []string)`

SetBoundServicePrincipalIds sets BoundServicePrincipalIds field to given value.

### HasBoundServicePrincipalIds

`func (o *AzureADAccessRules) HasBoundServicePrincipalIds() bool`

HasBoundServicePrincipalIds returns a boolean if a field has been set.

### GetBoundSubscriptionIds

`func (o *AzureADAccessRules) GetBoundSubscriptionIds() []string`

GetBoundSubscriptionIds returns the BoundSubscriptionIds field if non-nil, zero value otherwise.

### GetBoundSubscriptionIdsOk

`func (o *AzureADAccessRules) GetBoundSubscriptionIdsOk() (*[]string, bool)`

GetBoundSubscriptionIdsOk returns a tuple with the BoundSubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSubscriptionIds

`func (o *AzureADAccessRules) SetBoundSubscriptionIds(v []string)`

SetBoundSubscriptionIds sets BoundSubscriptionIds field to given value.

### HasBoundSubscriptionIds

`func (o *AzureADAccessRules) HasBoundSubscriptionIds() bool`

HasBoundSubscriptionIds returns a boolean if a field has been set.

### GetBoundTenantId

`func (o *AzureADAccessRules) GetBoundTenantId() string`

GetBoundTenantId returns the BoundTenantId field if non-nil, zero value otherwise.

### GetBoundTenantIdOk

`func (o *AzureADAccessRules) GetBoundTenantIdOk() (*string, bool)`

GetBoundTenantIdOk returns a tuple with the BoundTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTenantId

`func (o *AzureADAccessRules) SetBoundTenantId(v string)`

SetBoundTenantId sets BoundTenantId field to given value.

### HasBoundTenantId

`func (o *AzureADAccessRules) HasBoundTenantId() bool`

HasBoundTenantId returns a boolean if a field has been set.

### GetIssuer

`func (o *AzureADAccessRules) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AzureADAccessRules) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AzureADAccessRules) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AzureADAccessRules) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJwksUri

`func (o *AzureADAccessRules) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *AzureADAccessRules) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *AzureADAccessRules) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *AzureADAccessRules) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


