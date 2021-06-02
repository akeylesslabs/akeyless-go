# SignPKCS1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | **string** | The name of the RSA key to use in the signing process | 
**Message** | **string** | The message to be signed | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewSignPKCS1

`func NewSignPKCS1(keyName string, message string, ) *SignPKCS1`

NewSignPKCS1 instantiates a new SignPKCS1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignPKCS1WithDefaults

`func NewSignPKCS1WithDefaults() *SignPKCS1`

NewSignPKCS1WithDefaults instantiates a new SignPKCS1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *SignPKCS1) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *SignPKCS1) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *SignPKCS1) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetMessage

`func (o *SignPKCS1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignPKCS1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignPKCS1) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPassword

`func (o *SignPKCS1) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SignPKCS1) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SignPKCS1) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SignPKCS1) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *SignPKCS1) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SignPKCS1) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SignPKCS1) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SignPKCS1) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *SignPKCS1) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *SignPKCS1) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *SignPKCS1) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *SignPKCS1) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *SignPKCS1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SignPKCS1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SignPKCS1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SignPKCS1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


