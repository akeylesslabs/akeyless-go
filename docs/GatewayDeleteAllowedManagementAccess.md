# GatewayDeleteAllowedManagementAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**SubAdminId** | **string** | SubAdminID to be removed | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayDeleteAllowedManagementAccess

`func NewGatewayDeleteAllowedManagementAccess(subAdminId string, ) *GatewayDeleteAllowedManagementAccess`

NewGatewayDeleteAllowedManagementAccess instantiates a new GatewayDeleteAllowedManagementAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayDeleteAllowedManagementAccessWithDefaults

`func NewGatewayDeleteAllowedManagementAccessWithDefaults() *GatewayDeleteAllowedManagementAccess`

NewGatewayDeleteAllowedManagementAccessWithDefaults instantiates a new GatewayDeleteAllowedManagementAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *GatewayDeleteAllowedManagementAccess) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayDeleteAllowedManagementAccess) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayDeleteAllowedManagementAccess) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayDeleteAllowedManagementAccess) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSubAdminId

`func (o *GatewayDeleteAllowedManagementAccess) GetSubAdminId() string`

GetSubAdminId returns the SubAdminId field if non-nil, zero value otherwise.

### GetSubAdminIdOk

`func (o *GatewayDeleteAllowedManagementAccess) GetSubAdminIdOk() (*string, bool)`

GetSubAdminIdOk returns a tuple with the SubAdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAdminId

`func (o *GatewayDeleteAllowedManagementAccess) SetSubAdminId(v string)`

SetSubAdminId sets SubAdminId field to given value.


### GetToken

`func (o *GatewayDeleteAllowedManagementAccess) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayDeleteAllowedManagementAccess) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayDeleteAllowedManagementAccess) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayDeleteAllowedManagementAccess) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayDeleteAllowedManagementAccess) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayDeleteAllowedManagementAccess) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayDeleteAllowedManagementAccess) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayDeleteAllowedManagementAccess) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayDeleteAllowedManagementAccess) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayDeleteAllowedManagementAccess) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayDeleteAllowedManagementAccess) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayDeleteAllowedManagementAccess) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


