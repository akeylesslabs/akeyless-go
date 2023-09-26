# GatewayUpdateTlsCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertData** | Pointer to **string** | TLS Certificate (base64 encoded) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyData** | Pointer to **string** | TLS Private Key (base64 encoded) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateTlsCert

`func NewGatewayUpdateTlsCert() *GatewayUpdateTlsCert`

NewGatewayUpdateTlsCert instantiates a new GatewayUpdateTlsCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateTlsCertWithDefaults

`func NewGatewayUpdateTlsCertWithDefaults() *GatewayUpdateTlsCert`

NewGatewayUpdateTlsCertWithDefaults instantiates a new GatewayUpdateTlsCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertData

`func (o *GatewayUpdateTlsCert) GetCertData() string`

GetCertData returns the CertData field if non-nil, zero value otherwise.

### GetCertDataOk

`func (o *GatewayUpdateTlsCert) GetCertDataOk() (*string, bool)`

GetCertDataOk returns a tuple with the CertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertData

`func (o *GatewayUpdateTlsCert) SetCertData(v string)`

SetCertData sets CertData field to given value.

### HasCertData

`func (o *GatewayUpdateTlsCert) HasCertData() bool`

HasCertData returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateTlsCert) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateTlsCert) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateTlsCert) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateTlsCert) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyData

`func (o *GatewayUpdateTlsCert) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *GatewayUpdateTlsCert) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *GatewayUpdateTlsCert) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *GatewayUpdateTlsCert) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateTlsCert) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateTlsCert) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateTlsCert) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateTlsCert) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateTlsCert) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateTlsCert) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateTlsCert) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateTlsCert) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


