# GatewayUpdateK8SAuthConfigOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**PartsChange** | Pointer to [**ConfigChange**](ConfigChange.md) |  | [optional] 
**TotalHash** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewayUpdateK8SAuthConfigOutput

`func NewGatewayUpdateK8SAuthConfigOutput() *GatewayUpdateK8SAuthConfigOutput`

NewGatewayUpdateK8SAuthConfigOutput instantiates a new GatewayUpdateK8SAuthConfigOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateK8SAuthConfigOutputWithDefaults

`func NewGatewayUpdateK8SAuthConfigOutputWithDefaults() *GatewayUpdateK8SAuthConfigOutput`

NewGatewayUpdateK8SAuthConfigOutputWithDefaults instantiates a new GatewayUpdateK8SAuthConfigOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *GatewayUpdateK8SAuthConfigOutput) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *GatewayUpdateK8SAuthConfigOutput) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *GatewayUpdateK8SAuthConfigOutput) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *GatewayUpdateK8SAuthConfigOutput) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetPartsChange

`func (o *GatewayUpdateK8SAuthConfigOutput) GetPartsChange() ConfigChange`

GetPartsChange returns the PartsChange field if non-nil, zero value otherwise.

### GetPartsChangeOk

`func (o *GatewayUpdateK8SAuthConfigOutput) GetPartsChangeOk() (*ConfigChange, bool)`

GetPartsChangeOk returns a tuple with the PartsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsChange

`func (o *GatewayUpdateK8SAuthConfigOutput) SetPartsChange(v ConfigChange)`

SetPartsChange sets PartsChange field to given value.

### HasPartsChange

`func (o *GatewayUpdateK8SAuthConfigOutput) HasPartsChange() bool`

HasPartsChange returns a boolean if a field has been set.

### GetTotalHash

`func (o *GatewayUpdateK8SAuthConfigOutput) GetTotalHash() string`

GetTotalHash returns the TotalHash field if non-nil, zero value otherwise.

### GetTotalHashOk

`func (o *GatewayUpdateK8SAuthConfigOutput) GetTotalHashOk() (*string, bool)`

GetTotalHashOk returns a tuple with the TotalHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHash

`func (o *GatewayUpdateK8SAuthConfigOutput) SetTotalHash(v string)`

SetTotalHash sets TotalHash field to given value.

### HasTotalHash

`func (o *GatewayUpdateK8SAuthConfigOutput) HasTotalHash() bool`

HasTotalHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


