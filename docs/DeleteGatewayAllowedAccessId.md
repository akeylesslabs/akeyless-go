# DeleteGatewayAllowedAccessId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | The access id to be stripped of access to gateway | 
**ClusterName** | **string** | The name of the updated cluster, e.g. acc-abcd12345678/p-123456789012/defaultCluster | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDeleteGatewayAllowedAccessId

`func NewDeleteGatewayAllowedAccessId(accessId string, clusterName string, ) *DeleteGatewayAllowedAccessId`

NewDeleteGatewayAllowedAccessId instantiates a new DeleteGatewayAllowedAccessId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteGatewayAllowedAccessIdWithDefaults

`func NewDeleteGatewayAllowedAccessIdWithDefaults() *DeleteGatewayAllowedAccessId`

NewDeleteGatewayAllowedAccessIdWithDefaults instantiates a new DeleteGatewayAllowedAccessId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *DeleteGatewayAllowedAccessId) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *DeleteGatewayAllowedAccessId) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *DeleteGatewayAllowedAccessId) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetClusterName

`func (o *DeleteGatewayAllowedAccessId) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DeleteGatewayAllowedAccessId) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DeleteGatewayAllowedAccessId) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetJson

`func (o *DeleteGatewayAllowedAccessId) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DeleteGatewayAllowedAccessId) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DeleteGatewayAllowedAccessId) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DeleteGatewayAllowedAccessId) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *DeleteGatewayAllowedAccessId) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteGatewayAllowedAccessId) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteGatewayAllowedAccessId) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeleteGatewayAllowedAccessId) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeleteGatewayAllowedAccessId) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeleteGatewayAllowedAccessId) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeleteGatewayAllowedAccessId) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeleteGatewayAllowedAccessId) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


