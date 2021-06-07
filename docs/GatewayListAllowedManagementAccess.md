# GatewayListAllowedManagementAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

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

### GetGatewayUrl

`func (o *GatewayListAllowedManagementAccess) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GatewayListAllowedManagementAccess) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GatewayListAllowedManagementAccess) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *GatewayListAllowedManagementAccess) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetPassword

`func (o *GatewayListAllowedManagementAccess) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayListAllowedManagementAccess) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayListAllowedManagementAccess) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayListAllowedManagementAccess) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

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

### GetUsername

`func (o *GatewayListAllowedManagementAccess) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayListAllowedManagementAccess) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayListAllowedManagementAccess) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayListAllowedManagementAccess) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


