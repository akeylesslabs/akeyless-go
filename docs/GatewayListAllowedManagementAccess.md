# GatewayListAllowedManagementAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayListAllowedManagementAccess

`func NewGatewayListAllowedManagementAccess() *GatewayListAllowedManagementAccess`

NewGatewayListAllowedManagementAccess instantiates a new GatewayListAllowedManagementAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayListAllowedManagementAccessWithDefaults

`func NewGatewayListAllowedManagementAccessWithDefaults() *GatewayListAllowedManagementAccess`

NewGatewayListAllowedManagementAccessWithDefaults instantiates a new GatewayListAllowedManagementAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *GatewayListAllowedManagementAccess) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayListAllowedManagementAccess) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayListAllowedManagementAccess) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayListAllowedManagementAccess) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *GatewayListAllowedManagementAccess) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayListAllowedManagementAccess) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayListAllowedManagementAccess) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayListAllowedManagementAccess) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayListAllowedManagementAccess) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayListAllowedManagementAccess) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayListAllowedManagementAccess) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayListAllowedManagementAccess) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


