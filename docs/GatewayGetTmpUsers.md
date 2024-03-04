# GatewayGetTmpUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayGetTmpUsers

`func NewGatewayGetTmpUsers(name string, ) *GatewayGetTmpUsers`

NewGatewayGetTmpUsers instantiates a new GatewayGetTmpUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayGetTmpUsersWithDefaults

`func NewGatewayGetTmpUsersWithDefaults() *GatewayGetTmpUsers`

NewGatewayGetTmpUsersWithDefaults instantiates a new GatewayGetTmpUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *GatewayGetTmpUsers) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayGetTmpUsers) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayGetTmpUsers) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayGetTmpUsers) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayGetTmpUsers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayGetTmpUsers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayGetTmpUsers) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *GatewayGetTmpUsers) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayGetTmpUsers) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayGetTmpUsers) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayGetTmpUsers) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayGetTmpUsers) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayGetTmpUsers) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayGetTmpUsers) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayGetTmpUsers) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


