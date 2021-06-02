# Encrypt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionContext** | Pointer to **map[string]string** | name-value pair that specifies the encryption context to be used for authenticated encryption. If used here, the same value must be supplied to the decrypt command or decryption will fail | [optional] 
**KeyName** | **string** | The name of the key to use in the encryption process | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Plaintext** | **string** | Data to be encrypted | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewEncrypt

`func NewEncrypt(keyName string, plaintext string, ) *Encrypt`

NewEncrypt instantiates a new Encrypt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptWithDefaults

`func NewEncryptWithDefaults() *Encrypt`

NewEncryptWithDefaults instantiates a new Encrypt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionContext

`func (o *Encrypt) GetEncryptionContext() map[string]string`

GetEncryptionContext returns the EncryptionContext field if non-nil, zero value otherwise.

### GetEncryptionContextOk

`func (o *Encrypt) GetEncryptionContextOk() (*map[string]string, bool)`

GetEncryptionContextOk returns a tuple with the EncryptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionContext

`func (o *Encrypt) SetEncryptionContext(v map[string]string)`

SetEncryptionContext sets EncryptionContext field to given value.

### HasEncryptionContext

`func (o *Encrypt) HasEncryptionContext() bool`

HasEncryptionContext returns a boolean if a field has been set.

### GetKeyName

`func (o *Encrypt) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *Encrypt) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *Encrypt) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetPassword

`func (o *Encrypt) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Encrypt) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Encrypt) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Encrypt) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPlaintext

`func (o *Encrypt) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *Encrypt) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *Encrypt) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.


### GetToken

`func (o *Encrypt) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Encrypt) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Encrypt) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Encrypt) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *Encrypt) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *Encrypt) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *Encrypt) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *Encrypt) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *Encrypt) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Encrypt) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Encrypt) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Encrypt) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


