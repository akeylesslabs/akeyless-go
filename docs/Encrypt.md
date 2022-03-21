# Encrypt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the key to use in the encryption process | [optional] 
**EncryptionContext** | Pointer to **map[string]string** | name-value pair that specifies the encryption context to be used for authenticated encryption. If used here, the same value must be supplied to the decrypt command or decryption will fail | [optional] 
**KeyName** | Pointer to **string** | The name of the key to use in the encryption process | [optional] 
**Plaintext** | Pointer to **string** | Data to be encrypted | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewEncrypt

`func NewEncrypt() *Encrypt`

NewEncrypt instantiates a new Encrypt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptWithDefaults

`func NewEncryptWithDefaults() *Encrypt`

NewEncryptWithDefaults instantiates a new Encrypt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *Encrypt) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Encrypt) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Encrypt) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *Encrypt) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

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

### HasKeyName

`func (o *Encrypt) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

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

### HasPlaintext

`func (o *Encrypt) HasPlaintext() bool`

HasPlaintext returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


