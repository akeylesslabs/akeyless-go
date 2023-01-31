# KubernetesAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **string** | Audience is an optional Kubernetes jwt claim to verify | [optional] 
**BoundNamespaces** | Pointer to **[]string** | A list of namespaces that the authentication is restricted to. | [optional] 
**BoundPodNames** | Pointer to **[]string** | A list of pods names that the authentication is restricted to. | [optional] 
**BoundServiceAccountNames** | Pointer to **[]string** | A list of service account names that the authentication is restricted to. | [optional] 
**GenKeyPair** | Pointer to **string** | Generate public/private key (the private key is required for the K8S Auth Config in the Akeyless Gateway) | [optional] 
**PubKey** | Pointer to **string** | The public key value of the Kubernetes auth method configuration in the Akeyless Gateway. | [optional] 

## Methods

### NewKubernetesAccessRules

`func NewKubernetesAccessRules() *KubernetesAccessRules`

NewKubernetesAccessRules instantiates a new KubernetesAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAccessRulesWithDefaults

`func NewKubernetesAccessRulesWithDefaults() *KubernetesAccessRules`

NewKubernetesAccessRulesWithDefaults instantiates a new KubernetesAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *KubernetesAccessRules) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *KubernetesAccessRules) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *KubernetesAccessRules) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *KubernetesAccessRules) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetAudience

`func (o *KubernetesAccessRules) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *KubernetesAccessRules) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *KubernetesAccessRules) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *KubernetesAccessRules) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBoundNamespaces

`func (o *KubernetesAccessRules) GetBoundNamespaces() []string`

GetBoundNamespaces returns the BoundNamespaces field if non-nil, zero value otherwise.

### GetBoundNamespacesOk

`func (o *KubernetesAccessRules) GetBoundNamespacesOk() (*[]string, bool)`

GetBoundNamespacesOk returns a tuple with the BoundNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundNamespaces

`func (o *KubernetesAccessRules) SetBoundNamespaces(v []string)`

SetBoundNamespaces sets BoundNamespaces field to given value.

### HasBoundNamespaces

`func (o *KubernetesAccessRules) HasBoundNamespaces() bool`

HasBoundNamespaces returns a boolean if a field has been set.

### GetBoundPodNames

`func (o *KubernetesAccessRules) GetBoundPodNames() []string`

GetBoundPodNames returns the BoundPodNames field if non-nil, zero value otherwise.

### GetBoundPodNamesOk

`func (o *KubernetesAccessRules) GetBoundPodNamesOk() (*[]string, bool)`

GetBoundPodNamesOk returns a tuple with the BoundPodNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundPodNames

`func (o *KubernetesAccessRules) SetBoundPodNames(v []string)`

SetBoundPodNames sets BoundPodNames field to given value.

### HasBoundPodNames

`func (o *KubernetesAccessRules) HasBoundPodNames() bool`

HasBoundPodNames returns a boolean if a field has been set.

### GetBoundServiceAccountNames

`func (o *KubernetesAccessRules) GetBoundServiceAccountNames() []string`

GetBoundServiceAccountNames returns the BoundServiceAccountNames field if non-nil, zero value otherwise.

### GetBoundServiceAccountNamesOk

`func (o *KubernetesAccessRules) GetBoundServiceAccountNamesOk() (*[]string, bool)`

GetBoundServiceAccountNamesOk returns a tuple with the BoundServiceAccountNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServiceAccountNames

`func (o *KubernetesAccessRules) SetBoundServiceAccountNames(v []string)`

SetBoundServiceAccountNames sets BoundServiceAccountNames field to given value.

### HasBoundServiceAccountNames

`func (o *KubernetesAccessRules) HasBoundServiceAccountNames() bool`

HasBoundServiceAccountNames returns a boolean if a field has been set.

### GetGenKeyPair

`func (o *KubernetesAccessRules) GetGenKeyPair() string`

GetGenKeyPair returns the GenKeyPair field if non-nil, zero value otherwise.

### GetGenKeyPairOk

`func (o *KubernetesAccessRules) GetGenKeyPairOk() (*string, bool)`

GetGenKeyPairOk returns a tuple with the GenKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenKeyPair

`func (o *KubernetesAccessRules) SetGenKeyPair(v string)`

SetGenKeyPair sets GenKeyPair field to given value.

### HasGenKeyPair

`func (o *KubernetesAccessRules) HasGenKeyPair() bool`

HasGenKeyPair returns a boolean if a field has been set.

### GetPubKey

`func (o *KubernetesAccessRules) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *KubernetesAccessRules) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *KubernetesAccessRules) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.

### HasPubKey

`func (o *KubernetesAccessRules) HasPubKey() bool`

HasPubKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


