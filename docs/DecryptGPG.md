# DecryptGPG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | **string** | Ciphertext to be decrypted in base64 encoded format | 
**DisplayId** | Pointer to **string** | The display id of the key to use in the decryption process | [optional] 
**ItemId** | Pointer to **int64** | The item id of the key to use in the decryption process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyName** | **string** | The name of the key to use in the decryption process | 
**OutputFormat** | Pointer to **string** | If specified, the output will be formatted accordingly. options: [base64] | [optional] 
**Passphrase** | Pointer to **string** | Passphrase that was used to generate the key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDecryptGPG

`func NewDecryptGPG(ciphertext string, keyName string, ) *DecryptGPG`

NewDecryptGPG instantiates a new DecryptGPG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptGPGWithDefaults

`func NewDecryptGPGWithDefaults() *DecryptGPG`

NewDecryptGPGWithDefaults instantiates a new DecryptGPG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *DecryptGPG) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *DecryptGPG) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *DecryptGPG) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### GetDisplayId

`func (o *DecryptGPG) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *DecryptGPG) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *DecryptGPG) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *DecryptGPG) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetItemId

`func (o *DecryptGPG) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DecryptGPG) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DecryptGPG) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DecryptGPG) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *DecryptGPG) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DecryptGPG) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DecryptGPG) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DecryptGPG) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyName

`func (o *DecryptGPG) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *DecryptGPG) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *DecryptGPG) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetOutputFormat

`func (o *DecryptGPG) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *DecryptGPG) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *DecryptGPG) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *DecryptGPG) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPassphrase

`func (o *DecryptGPG) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *DecryptGPG) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *DecryptGPG) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *DecryptGPG) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetToken

`func (o *DecryptGPG) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DecryptGPG) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DecryptGPG) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DecryptGPG) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DecryptGPG) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DecryptGPG) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DecryptGPG) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DecryptGPG) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


