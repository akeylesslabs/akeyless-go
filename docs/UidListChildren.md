# UidListChildren

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodName** | Pointer to **string** | The universal identity auth method name, required only when uid-token is not provided | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewUidListChildren

`func NewUidListChildren() *UidListChildren`

NewUidListChildren instantiates a new UidListChildren object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUidListChildrenWithDefaults

`func NewUidListChildrenWithDefaults() *UidListChildren`

NewUidListChildrenWithDefaults instantiates a new UidListChildren object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodName

`func (o *UidListChildren) GetAuthMethodName() string`

GetAuthMethodName returns the AuthMethodName field if non-nil, zero value otherwise.

### GetAuthMethodNameOk

`func (o *UidListChildren) GetAuthMethodNameOk() (*string, bool)`

GetAuthMethodNameOk returns a tuple with the AuthMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodName

`func (o *UidListChildren) SetAuthMethodName(v string)`

SetAuthMethodName sets AuthMethodName field to given value.

### HasAuthMethodName

`func (o *UidListChildren) HasAuthMethodName() bool`

HasAuthMethodName returns a boolean if a field has been set.

### GetPassword

`func (o *UidListChildren) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UidListChildren) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UidListChildren) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UidListChildren) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *UidListChildren) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UidListChildren) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UidListChildren) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UidListChildren) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UidListChildren) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UidListChildren) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UidListChildren) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UidListChildren) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *UidListChildren) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UidListChildren) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UidListChildren) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UidListChildren) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


