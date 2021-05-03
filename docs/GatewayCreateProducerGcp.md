# GatewayCreateProducerGcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**GcpCredType** | **string** |  | 
**GcpKey** | Pointer to **string** | Base64-encoded service account private key text | [optional] 
**GcpKeyAlgo** | Pointer to **string** | Service account key algorithm, e.g. KEY_ALG_RSA_1024 | [optional] 
**GcpSaEmail** | **string** | GCP service account email | 
**GcpTokenScopes** | Pointer to **string** | Access token scopes list, e.g. scope1,scope2 | [optional] 
**Name** | **string** | Producer name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerGcp

`func NewGatewayCreateProducerGcp(gcpCredType string, gcpSaEmail string, name string, ) *GatewayCreateProducerGcp`

NewGatewayCreateProducerGcp instantiates a new GatewayCreateProducerGcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerGcpWithDefaults

`func NewGatewayCreateProducerGcpWithDefaults() *GatewayCreateProducerGcp`

NewGatewayCreateProducerGcpWithDefaults instantiates a new GatewayCreateProducerGcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayUrl

`func (o *GatewayCreateProducerGcp) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GatewayCreateProducerGcp) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GatewayCreateProducerGcp) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *GatewayCreateProducerGcp) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetGcpCredType

`func (o *GatewayCreateProducerGcp) GetGcpCredType() string`

GetGcpCredType returns the GcpCredType field if non-nil, zero value otherwise.

### GetGcpCredTypeOk

`func (o *GatewayCreateProducerGcp) GetGcpCredTypeOk() (*string, bool)`

GetGcpCredTypeOk returns a tuple with the GcpCredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredType

`func (o *GatewayCreateProducerGcp) SetGcpCredType(v string)`

SetGcpCredType sets GcpCredType field to given value.


### GetGcpKey

`func (o *GatewayCreateProducerGcp) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *GatewayCreateProducerGcp) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *GatewayCreateProducerGcp) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *GatewayCreateProducerGcp) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetGcpKeyAlgo

`func (o *GatewayCreateProducerGcp) GetGcpKeyAlgo() string`

GetGcpKeyAlgo returns the GcpKeyAlgo field if non-nil, zero value otherwise.

### GetGcpKeyAlgoOk

`func (o *GatewayCreateProducerGcp) GetGcpKeyAlgoOk() (*string, bool)`

GetGcpKeyAlgoOk returns a tuple with the GcpKeyAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKeyAlgo

`func (o *GatewayCreateProducerGcp) SetGcpKeyAlgo(v string)`

SetGcpKeyAlgo sets GcpKeyAlgo field to given value.

### HasGcpKeyAlgo

`func (o *GatewayCreateProducerGcp) HasGcpKeyAlgo() bool`

HasGcpKeyAlgo returns a boolean if a field has been set.

### GetGcpSaEmail

`func (o *GatewayCreateProducerGcp) GetGcpSaEmail() string`

GetGcpSaEmail returns the GcpSaEmail field if non-nil, zero value otherwise.

### GetGcpSaEmailOk

`func (o *GatewayCreateProducerGcp) GetGcpSaEmailOk() (*string, bool)`

GetGcpSaEmailOk returns a tuple with the GcpSaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSaEmail

`func (o *GatewayCreateProducerGcp) SetGcpSaEmail(v string)`

SetGcpSaEmail sets GcpSaEmail field to given value.


### GetGcpTokenScopes

`func (o *GatewayCreateProducerGcp) GetGcpTokenScopes() string`

GetGcpTokenScopes returns the GcpTokenScopes field if non-nil, zero value otherwise.

### GetGcpTokenScopesOk

`func (o *GatewayCreateProducerGcp) GetGcpTokenScopesOk() (*string, bool)`

GetGcpTokenScopesOk returns a tuple with the GcpTokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTokenScopes

`func (o *GatewayCreateProducerGcp) SetGcpTokenScopes(v string)`

SetGcpTokenScopes sets GcpTokenScopes field to given value.

### HasGcpTokenScopes

`func (o *GatewayCreateProducerGcp) HasGcpTokenScopes() bool`

HasGcpTokenScopes returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerGcp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerGcp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerGcp) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerGcp) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerGcp) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerGcp) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerGcp) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerGcp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerGcp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerGcp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerGcp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerGcp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerGcp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerGcp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerGcp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerGcp) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerGcp) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerGcp) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerGcp) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


