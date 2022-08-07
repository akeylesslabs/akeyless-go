# DecryptPKCS1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | **string** | Ciphertext to be decrypted in base64 encoded format | 
**DisplayId** | Pointer to **string** | The display id of the key to use in the decryption process | [optional] 
**ItemId** | Pointer to **int64** | The item id of the key to use in the decryption process | [optional] 
**KeyName** | **string** | The name of the key to use in the decryption process | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDecryptPKCS1

`func NewDecryptPKCS1(ciphertext string, keyName string, ) *DecryptPKCS1`

NewDecryptPKCS1 instantiates a new DecryptPKCS1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptPKCS1WithDefaults

`func NewDecryptPKCS1WithDefaults() *DecryptPKCS1`

NewDecryptPKCS1WithDefaults instantiates a new DecryptPKCS1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *DecryptPKCS1) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *DecryptPKCS1) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *DecryptPKCS1) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### GetDisplayId

`func (o *DecryptPKCS1) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *DecryptPKCS1) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *DecryptPKCS1) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *DecryptPKCS1) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetItemId

`func (o *DecryptPKCS1) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DecryptPKCS1) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DecryptPKCS1) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DecryptPKCS1) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetKeyName

`func (o *DecryptPKCS1) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *DecryptPKCS1) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *DecryptPKCS1) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetToken

`func (o *DecryptPKCS1) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DecryptPKCS1) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DecryptPKCS1) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DecryptPKCS1) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DecryptPKCS1) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DecryptPKCS1) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DecryptPKCS1) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DecryptPKCS1) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


