# Decrypt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | **string** | Ciphertext to be decrypted in base64 encoded format | 
**EncryptionContext** | Pointer to **map[string]string** | The encryption context. If this was specified in the encrypt command, it must be specified here or the decryption operation will fail | [optional] 
**KeyName** | **string** | The name of the key to use in the decryption process | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewDecrypt

`func NewDecrypt(ciphertext string, keyName string, ) *Decrypt`

NewDecrypt instantiates a new Decrypt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptWithDefaults

`func NewDecryptWithDefaults() *Decrypt`

NewDecryptWithDefaults instantiates a new Decrypt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *Decrypt) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *Decrypt) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *Decrypt) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### GetEncryptionContext

`func (o *Decrypt) GetEncryptionContext() map[string]string`

GetEncryptionContext returns the EncryptionContext field if non-nil, zero value otherwise.

### GetEncryptionContextOk

`func (o *Decrypt) GetEncryptionContextOk() (*map[string]string, bool)`

GetEncryptionContextOk returns a tuple with the EncryptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionContext

`func (o *Decrypt) SetEncryptionContext(v map[string]string)`

SetEncryptionContext sets EncryptionContext field to given value.

### HasEncryptionContext

`func (o *Decrypt) HasEncryptionContext() bool`

HasEncryptionContext returns a boolean if a field has been set.

### GetKeyName

`func (o *Decrypt) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *Decrypt) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *Decrypt) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetPassword

`func (o *Decrypt) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Decrypt) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Decrypt) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Decrypt) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *Decrypt) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Decrypt) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Decrypt) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Decrypt) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *Decrypt) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *Decrypt) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *Decrypt) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *Decrypt) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *Decrypt) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Decrypt) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Decrypt) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Decrypt) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


