# GatewayAddAllowedManagementAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowGwApi** | Pointer to **bool** |  | [optional] 
**AllowGwLogin** | Pointer to **bool** |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**SubAdminAccessId** | **string** | SubAdmins to add | 
**SubClaims** | Pointer to **map[string]string** | key/val of sub claims, e.g group&#x3D;admins,developers | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayAddAllowedManagementAccess

`func NewGatewayAddAllowedManagementAccess(subAdminAccessId string, ) *GatewayAddAllowedManagementAccess`

NewGatewayAddAllowedManagementAccess instantiates a new GatewayAddAllowedManagementAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayAddAllowedManagementAccessWithDefaults

`func NewGatewayAddAllowedManagementAccessWithDefaults() *GatewayAddAllowedManagementAccess`

NewGatewayAddAllowedManagementAccessWithDefaults instantiates a new GatewayAddAllowedManagementAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowGwApi

`func (o *GatewayAddAllowedManagementAccess) GetAllowGwApi() bool`

GetAllowGwApi returns the AllowGwApi field if non-nil, zero value otherwise.

### GetAllowGwApiOk

`func (o *GatewayAddAllowedManagementAccess) GetAllowGwApiOk() (*bool, bool)`

GetAllowGwApiOk returns a tuple with the AllowGwApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGwApi

`func (o *GatewayAddAllowedManagementAccess) SetAllowGwApi(v bool)`

SetAllowGwApi sets AllowGwApi field to given value.

### HasAllowGwApi

`func (o *GatewayAddAllowedManagementAccess) HasAllowGwApi() bool`

HasAllowGwApi returns a boolean if a field has been set.

### GetAllowGwLogin

`func (o *GatewayAddAllowedManagementAccess) GetAllowGwLogin() bool`

GetAllowGwLogin returns the AllowGwLogin field if non-nil, zero value otherwise.

### GetAllowGwLoginOk

`func (o *GatewayAddAllowedManagementAccess) GetAllowGwLoginOk() (*bool, bool)`

GetAllowGwLoginOk returns a tuple with the AllowGwLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGwLogin

`func (o *GatewayAddAllowedManagementAccess) SetAllowGwLogin(v bool)`

SetAllowGwLogin sets AllowGwLogin field to given value.

### HasAllowGwLogin

`func (o *GatewayAddAllowedManagementAccess) HasAllowGwLogin() bool`

HasAllowGwLogin returns a boolean if a field has been set.

### GetJson

`func (o *GatewayAddAllowedManagementAccess) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayAddAllowedManagementAccess) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayAddAllowedManagementAccess) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayAddAllowedManagementAccess) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSubAdminAccessId

`func (o *GatewayAddAllowedManagementAccess) GetSubAdminAccessId() string`

GetSubAdminAccessId returns the SubAdminAccessId field if non-nil, zero value otherwise.

### GetSubAdminAccessIdOk

`func (o *GatewayAddAllowedManagementAccess) GetSubAdminAccessIdOk() (*string, bool)`

GetSubAdminAccessIdOk returns a tuple with the SubAdminAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAdminAccessId

`func (o *GatewayAddAllowedManagementAccess) SetSubAdminAccessId(v string)`

SetSubAdminAccessId sets SubAdminAccessId field to given value.


### GetSubClaims

`func (o *GatewayAddAllowedManagementAccess) GetSubClaims() map[string]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *GatewayAddAllowedManagementAccess) GetSubClaimsOk() (*map[string]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *GatewayAddAllowedManagementAccess) SetSubClaims(v map[string]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *GatewayAddAllowedManagementAccess) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.

### GetToken

`func (o *GatewayAddAllowedManagementAccess) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayAddAllowedManagementAccess) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayAddAllowedManagementAccess) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayAddAllowedManagementAccess) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayAddAllowedManagementAccess) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayAddAllowedManagementAccess) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayAddAllowedManagementAccess) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayAddAllowedManagementAccess) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


