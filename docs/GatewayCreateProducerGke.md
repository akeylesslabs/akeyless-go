# GatewayCreateProducerGke

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**GkeClusterCert** | **string** | GKE cluster CA certificate | 
**GkeClusterEndpoint** | **string** | GKE cluster URL endpoint | 
**GkeClusterName** | **string** | GKE cluster name | 
**GkeServiceAccountEmail** | **string** | GKE service account email | 
**Name** | **string** | Producer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayCreateProducerGke

`func NewGatewayCreateProducerGke(gkeClusterCert string, gkeClusterEndpoint string, gkeClusterName string, gkeServiceAccountEmail string, name string, ) *GatewayCreateProducerGke`

NewGatewayCreateProducerGke instantiates a new GatewayCreateProducerGke object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerGkeWithDefaults

`func NewGatewayCreateProducerGkeWithDefaults() *GatewayCreateProducerGke`

NewGatewayCreateProducerGkeWithDefaults instantiates a new GatewayCreateProducerGke object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayUrl

`func (o *GatewayCreateProducerGke) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GatewayCreateProducerGke) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GatewayCreateProducerGke) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *GatewayCreateProducerGke) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetGkeClusterCert

`func (o *GatewayCreateProducerGke) GetGkeClusterCert() string`

GetGkeClusterCert returns the GkeClusterCert field if non-nil, zero value otherwise.

### GetGkeClusterCertOk

`func (o *GatewayCreateProducerGke) GetGkeClusterCertOk() (*string, bool)`

GetGkeClusterCertOk returns a tuple with the GkeClusterCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterCert

`func (o *GatewayCreateProducerGke) SetGkeClusterCert(v string)`

SetGkeClusterCert sets GkeClusterCert field to given value.


### GetGkeClusterEndpoint

`func (o *GatewayCreateProducerGke) GetGkeClusterEndpoint() string`

GetGkeClusterEndpoint returns the GkeClusterEndpoint field if non-nil, zero value otherwise.

### GetGkeClusterEndpointOk

`func (o *GatewayCreateProducerGke) GetGkeClusterEndpointOk() (*string, bool)`

GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterEndpoint

`func (o *GatewayCreateProducerGke) SetGkeClusterEndpoint(v string)`

SetGkeClusterEndpoint sets GkeClusterEndpoint field to given value.


### GetGkeClusterName

`func (o *GatewayCreateProducerGke) GetGkeClusterName() string`

GetGkeClusterName returns the GkeClusterName field if non-nil, zero value otherwise.

### GetGkeClusterNameOk

`func (o *GatewayCreateProducerGke) GetGkeClusterNameOk() (*string, bool)`

GetGkeClusterNameOk returns a tuple with the GkeClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterName

`func (o *GatewayCreateProducerGke) SetGkeClusterName(v string)`

SetGkeClusterName sets GkeClusterName field to given value.


### GetGkeServiceAccountEmail

`func (o *GatewayCreateProducerGke) GetGkeServiceAccountEmail() string`

GetGkeServiceAccountEmail returns the GkeServiceAccountEmail field if non-nil, zero value otherwise.

### GetGkeServiceAccountEmailOk

`func (o *GatewayCreateProducerGke) GetGkeServiceAccountEmailOk() (*string, bool)`

GetGkeServiceAccountEmailOk returns a tuple with the GkeServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeServiceAccountEmail

`func (o *GatewayCreateProducerGke) SetGkeServiceAccountEmail(v string)`

SetGkeServiceAccountEmail sets GkeServiceAccountEmail field to given value.


### GetName

`func (o *GatewayCreateProducerGke) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerGke) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerGke) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *GatewayCreateProducerGke) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateProducerGke) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateProducerGke) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateProducerGke) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerGke) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerGke) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerGke) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerGke) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerGke) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerGke) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerGke) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerGke) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerGke) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerGke) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerGke) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerGke) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerGke) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerGke) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerGke) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerGke) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayCreateProducerGke) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateProducerGke) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateProducerGke) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateProducerGke) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


