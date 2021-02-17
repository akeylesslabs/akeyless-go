# GatewayDeleteProducer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**Name** | **string** | Producer name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayDeleteProducer

`func NewGatewayDeleteProducer(name string, ) *GatewayDeleteProducer`

NewGatewayDeleteProducer instantiates a new GatewayDeleteProducer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayDeleteProducerWithDefaults

`func NewGatewayDeleteProducerWithDefaults() *GatewayDeleteProducer`

NewGatewayDeleteProducerWithDefaults instantiates a new GatewayDeleteProducer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayUrl

`func (o *GatewayDeleteProducer) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GatewayDeleteProducer) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GatewayDeleteProducer) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *GatewayDeleteProducer) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetName

`func (o *GatewayDeleteProducer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayDeleteProducer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayDeleteProducer) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *GatewayDeleteProducer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayDeleteProducer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayDeleteProducer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayDeleteProducer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayDeleteProducer) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayDeleteProducer) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayDeleteProducer) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayDeleteProducer) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

