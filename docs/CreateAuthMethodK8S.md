# CreateAuthMethodK8S

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**Audience** | Pointer to **string** | The audience in the Kubernetes JWT that the access is restricted to | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**BoundNamespaces** | Pointer to **[]string** | A list of namespaces that the access is restricted to | [optional] 
**BoundPodNames** | Pointer to **[]string** | A list of pod names that the access is restricted to | [optional] 
**BoundSaNames** | Pointer to **[]string** | A list of service account names that the access is restricted to | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GenKey** | Pointer to **string** | Automatically generate key-pair for K8S configuration. If set to false, a public key needs to be provided [true/false] | [optional] [default to "true"]
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**PublicKey** | Pointer to **string** | Base64-encoded or PEM formatted public key data for K8S authentication method is required [RSA2048] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateAuthMethodK8S

`func NewCreateAuthMethodK8S(name string, ) *CreateAuthMethodK8S`

NewCreateAuthMethodK8S instantiates a new CreateAuthMethodK8S object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodK8SWithDefaults

`func NewCreateAuthMethodK8SWithDefaults() *CreateAuthMethodK8S`

NewCreateAuthMethodK8SWithDefaults instantiates a new CreateAuthMethodK8S object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodK8S) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodK8S) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodK8S) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodK8S) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAudience

`func (o *CreateAuthMethodK8S) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CreateAuthMethodK8S) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CreateAuthMethodK8S) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CreateAuthMethodK8S) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBoundIps

`func (o *CreateAuthMethodK8S) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodK8S) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodK8S) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodK8S) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundNamespaces

`func (o *CreateAuthMethodK8S) GetBoundNamespaces() []string`

GetBoundNamespaces returns the BoundNamespaces field if non-nil, zero value otherwise.

### GetBoundNamespacesOk

`func (o *CreateAuthMethodK8S) GetBoundNamespacesOk() (*[]string, bool)`

GetBoundNamespacesOk returns a tuple with the BoundNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundNamespaces

`func (o *CreateAuthMethodK8S) SetBoundNamespaces(v []string)`

SetBoundNamespaces sets BoundNamespaces field to given value.

### HasBoundNamespaces

`func (o *CreateAuthMethodK8S) HasBoundNamespaces() bool`

HasBoundNamespaces returns a boolean if a field has been set.

### GetBoundPodNames

`func (o *CreateAuthMethodK8S) GetBoundPodNames() []string`

GetBoundPodNames returns the BoundPodNames field if non-nil, zero value otherwise.

### GetBoundPodNamesOk

`func (o *CreateAuthMethodK8S) GetBoundPodNamesOk() (*[]string, bool)`

GetBoundPodNamesOk returns a tuple with the BoundPodNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundPodNames

`func (o *CreateAuthMethodK8S) SetBoundPodNames(v []string)`

SetBoundPodNames sets BoundPodNames field to given value.

### HasBoundPodNames

`func (o *CreateAuthMethodK8S) HasBoundPodNames() bool`

HasBoundPodNames returns a boolean if a field has been set.

### GetBoundSaNames

`func (o *CreateAuthMethodK8S) GetBoundSaNames() []string`

GetBoundSaNames returns the BoundSaNames field if non-nil, zero value otherwise.

### GetBoundSaNamesOk

`func (o *CreateAuthMethodK8S) GetBoundSaNamesOk() (*[]string, bool)`

GetBoundSaNamesOk returns a tuple with the BoundSaNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSaNames

`func (o *CreateAuthMethodK8S) SetBoundSaNames(v []string)`

SetBoundSaNames sets BoundSaNames field to given value.

### HasBoundSaNames

`func (o *CreateAuthMethodK8S) HasBoundSaNames() bool`

HasBoundSaNames returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAuthMethodK8S) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAuthMethodK8S) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAuthMethodK8S) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAuthMethodK8S) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *CreateAuthMethodK8S) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodK8S) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodK8S) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodK8S) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGenKey

`func (o *CreateAuthMethodK8S) GetGenKey() string`

GetGenKey returns the GenKey field if non-nil, zero value otherwise.

### GetGenKeyOk

`func (o *CreateAuthMethodK8S) GetGenKeyOk() (*string, bool)`

GetGenKeyOk returns a tuple with the GenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenKey

`func (o *CreateAuthMethodK8S) SetGenKey(v string)`

SetGenKey sets GenKey field to given value.

### HasGenKey

`func (o *CreateAuthMethodK8S) HasGenKey() bool`

HasGenKey returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *CreateAuthMethodK8S) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *CreateAuthMethodK8S) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *CreateAuthMethodK8S) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *CreateAuthMethodK8S) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *CreateAuthMethodK8S) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateAuthMethodK8S) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateAuthMethodK8S) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateAuthMethodK8S) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *CreateAuthMethodK8S) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *CreateAuthMethodK8S) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *CreateAuthMethodK8S) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *CreateAuthMethodK8S) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodK8S) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodK8S) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodK8S) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *CreateAuthMethodK8S) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CreateAuthMethodK8S) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CreateAuthMethodK8S) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CreateAuthMethodK8S) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetToken

`func (o *CreateAuthMethodK8S) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodK8S) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodK8S) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodK8S) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodK8S) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodK8S) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodK8S) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodK8S) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


