# GatewayAddSubAdmins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowGwApi** | Pointer to **bool** |  | [optional] 
**AllowGwLogin** | Pointer to **bool** |  | [optional] 
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**SubAdminAccessId** | **string** | SubAdmins to add | 
**SubClaims** | Pointer to **map[string]string** | key/val of sub claims, e.g group&#x3D;admins,developers | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayAddSubAdmins

`func NewGatewayAddSubAdmins(subAdminAccessId string, ) *GatewayAddSubAdmins`

NewGatewayAddSubAdmins instantiates a new GatewayAddSubAdmins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayAddSubAdminsWithDefaults

`func NewGatewayAddSubAdminsWithDefaults() *GatewayAddSubAdmins`

NewGatewayAddSubAdminsWithDefaults instantiates a new GatewayAddSubAdmins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowGwApi

`func (o *GatewayAddSubAdmins) GetAllowGwApi() bool`

GetAllowGwApi returns the AllowGwApi field if non-nil, zero value otherwise.

### GetAllowGwApiOk

`func (o *GatewayAddSubAdmins) GetAllowGwApiOk() (*bool, bool)`

GetAllowGwApiOk returns a tuple with the AllowGwApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGwApi

`func (o *GatewayAddSubAdmins) SetAllowGwApi(v bool)`

SetAllowGwApi sets AllowGwApi field to given value.

### HasAllowGwApi

`func (o *GatewayAddSubAdmins) HasAllowGwApi() bool`

HasAllowGwApi returns a boolean if a field has been set.

### GetAllowGwLogin

`func (o *GatewayAddSubAdmins) GetAllowGwLogin() bool`

GetAllowGwLogin returns the AllowGwLogin field if non-nil, zero value otherwise.

### GetAllowGwLoginOk

`func (o *GatewayAddSubAdmins) GetAllowGwLoginOk() (*bool, bool)`

GetAllowGwLoginOk returns a tuple with the AllowGwLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGwLogin

`func (o *GatewayAddSubAdmins) SetAllowGwLogin(v bool)`

SetAllowGwLogin sets AllowGwLogin field to given value.

### HasAllowGwLogin

`func (o *GatewayAddSubAdmins) HasAllowGwLogin() bool`

HasAllowGwLogin returns a boolean if a field has been set.

### GetGatewayUrl

`func (o *GatewayAddSubAdmins) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GatewayAddSubAdmins) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GatewayAddSubAdmins) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *GatewayAddSubAdmins) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetPassword

`func (o *GatewayAddSubAdmins) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayAddSubAdmins) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayAddSubAdmins) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayAddSubAdmins) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSubAdminAccessId

`func (o *GatewayAddSubAdmins) GetSubAdminAccessId() string`

GetSubAdminAccessId returns the SubAdminAccessId field if non-nil, zero value otherwise.

### GetSubAdminAccessIdOk

`func (o *GatewayAddSubAdmins) GetSubAdminAccessIdOk() (*string, bool)`

GetSubAdminAccessIdOk returns a tuple with the SubAdminAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAdminAccessId

`func (o *GatewayAddSubAdmins) SetSubAdminAccessId(v string)`

SetSubAdminAccessId sets SubAdminAccessId field to given value.


### GetSubClaims

`func (o *GatewayAddSubAdmins) GetSubClaims() map[string]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *GatewayAddSubAdmins) GetSubClaimsOk() (*map[string]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *GatewayAddSubAdmins) SetSubClaims(v map[string]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *GatewayAddSubAdmins) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.

### GetToken

`func (o *GatewayAddSubAdmins) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayAddSubAdmins) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayAddSubAdmins) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayAddSubAdmins) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayAddSubAdmins) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayAddSubAdmins) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayAddSubAdmins) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayAddSubAdmins) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayAddSubAdmins) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayAddSubAdmins) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayAddSubAdmins) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayAddSubAdmins) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


