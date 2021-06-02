# UidRevokeToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodName** | Pointer to **string** | The universal identity auth method name | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**RevokeToken** | **string** | the universal identity token/token-id to revoke | 
**RevokeType** | **string** | revokeSelf/revokeAll (delete only this token/this token and his children) | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewUidRevokeToken

`func NewUidRevokeToken(revokeToken string, revokeType string, ) *UidRevokeToken`

NewUidRevokeToken instantiates a new UidRevokeToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUidRevokeTokenWithDefaults

`func NewUidRevokeTokenWithDefaults() *UidRevokeToken`

NewUidRevokeTokenWithDefaults instantiates a new UidRevokeToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodName

`func (o *UidRevokeToken) GetAuthMethodName() string`

GetAuthMethodName returns the AuthMethodName field if non-nil, zero value otherwise.

### GetAuthMethodNameOk

`func (o *UidRevokeToken) GetAuthMethodNameOk() (*string, bool)`

GetAuthMethodNameOk returns a tuple with the AuthMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodName

`func (o *UidRevokeToken) SetAuthMethodName(v string)`

SetAuthMethodName sets AuthMethodName field to given value.

### HasAuthMethodName

`func (o *UidRevokeToken) HasAuthMethodName() bool`

HasAuthMethodName returns a boolean if a field has been set.

### GetPassword

`func (o *UidRevokeToken) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UidRevokeToken) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UidRevokeToken) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UidRevokeToken) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRevokeToken

`func (o *UidRevokeToken) GetRevokeToken() string`

GetRevokeToken returns the RevokeToken field if non-nil, zero value otherwise.

### GetRevokeTokenOk

`func (o *UidRevokeToken) GetRevokeTokenOk() (*string, bool)`

GetRevokeTokenOk returns a tuple with the RevokeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeToken

`func (o *UidRevokeToken) SetRevokeToken(v string)`

SetRevokeToken sets RevokeToken field to given value.


### GetRevokeType

`func (o *UidRevokeToken) GetRevokeType() string`

GetRevokeType returns the RevokeType field if non-nil, zero value otherwise.

### GetRevokeTypeOk

`func (o *UidRevokeToken) GetRevokeTypeOk() (*string, bool)`

GetRevokeTypeOk returns a tuple with the RevokeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeType

`func (o *UidRevokeToken) SetRevokeType(v string)`

SetRevokeType sets RevokeType field to given value.


### GetToken

`func (o *UidRevokeToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UidRevokeToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UidRevokeToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UidRevokeToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UidRevokeToken) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UidRevokeToken) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UidRevokeToken) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UidRevokeToken) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *UidRevokeToken) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UidRevokeToken) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UidRevokeToken) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UidRevokeToken) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


