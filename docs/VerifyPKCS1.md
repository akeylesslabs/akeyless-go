# VerifyPKCS1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | **string** | The name of the RSA key to use in the verification process | 
**Message** | **string** | The message to be verified | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Signature** | **string** | The message&#39;s signature | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewVerifyPKCS1

`func NewVerifyPKCS1(keyName string, message string, signature string, ) *VerifyPKCS1`

NewVerifyPKCS1 instantiates a new VerifyPKCS1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyPKCS1WithDefaults

`func NewVerifyPKCS1WithDefaults() *VerifyPKCS1`

NewVerifyPKCS1WithDefaults instantiates a new VerifyPKCS1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *VerifyPKCS1) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *VerifyPKCS1) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *VerifyPKCS1) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetMessage

`func (o *VerifyPKCS1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VerifyPKCS1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VerifyPKCS1) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPassword

`func (o *VerifyPKCS1) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VerifyPKCS1) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VerifyPKCS1) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VerifyPKCS1) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSignature

`func (o *VerifyPKCS1) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerifyPKCS1) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerifyPKCS1) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetToken

`func (o *VerifyPKCS1) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VerifyPKCS1) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VerifyPKCS1) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VerifyPKCS1) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *VerifyPKCS1) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *VerifyPKCS1) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *VerifyPKCS1) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *VerifyPKCS1) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *VerifyPKCS1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VerifyPKCS1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VerifyPKCS1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VerifyPKCS1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


