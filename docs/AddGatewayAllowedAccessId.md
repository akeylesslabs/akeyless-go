# AddGatewayAllowedAccessId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | The access id that will be able to access to gateway | 
**ClusterName** | **string** | The name of the updated cluster, e.g. acc-abcd12345678/p-123456789012/defaultCluster | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**SubClaims** | Pointer to **map[string]string** | key/val of sub claims, e.g group&#x3D;admins,developers | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewAddGatewayAllowedAccessId

`func NewAddGatewayAllowedAccessId(accessId string, clusterName string, ) *AddGatewayAllowedAccessId`

NewAddGatewayAllowedAccessId instantiates a new AddGatewayAllowedAccessId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGatewayAllowedAccessIdWithDefaults

`func NewAddGatewayAllowedAccessIdWithDefaults() *AddGatewayAllowedAccessId`

NewAddGatewayAllowedAccessIdWithDefaults instantiates a new AddGatewayAllowedAccessId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *AddGatewayAllowedAccessId) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *AddGatewayAllowedAccessId) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *AddGatewayAllowedAccessId) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetClusterName

`func (o *AddGatewayAllowedAccessId) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *AddGatewayAllowedAccessId) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *AddGatewayAllowedAccessId) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetJson

`func (o *AddGatewayAllowedAccessId) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AddGatewayAllowedAccessId) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AddGatewayAllowedAccessId) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AddGatewayAllowedAccessId) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSubClaims

`func (o *AddGatewayAllowedAccessId) GetSubClaims() map[string]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *AddGatewayAllowedAccessId) GetSubClaimsOk() (*map[string]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *AddGatewayAllowedAccessId) SetSubClaims(v map[string]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *AddGatewayAllowedAccessId) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.

### GetToken

`func (o *AddGatewayAllowedAccessId) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddGatewayAllowedAccessId) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddGatewayAllowedAccessId) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AddGatewayAllowedAccessId) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AddGatewayAllowedAccessId) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AddGatewayAllowedAccessId) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AddGatewayAllowedAccessId) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AddGatewayAllowedAccessId) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


