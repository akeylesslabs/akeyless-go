# EncryptGPG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the key to use in the encryption process | [optional] 
**InputFormat** | Pointer to **string** | If specified, the plaintext input is assumed to be formatted accordingly. Current supported options: [base64] | [optional] 
**ItemId** | Pointer to **int64** | The item id of the key to use in the encryption process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyName** | **string** | The name of the key to use in the encryption process | 
**Plaintext** | **string** | Data to be encrypted | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewEncryptGPG

`func NewEncryptGPG(keyName string, plaintext string, ) *EncryptGPG`

NewEncryptGPG instantiates a new EncryptGPG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptGPGWithDefaults

`func NewEncryptGPGWithDefaults() *EncryptGPG`

NewEncryptGPGWithDefaults instantiates a new EncryptGPG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *EncryptGPG) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *EncryptGPG) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *EncryptGPG) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *EncryptGPG) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetInputFormat

`func (o *EncryptGPG) GetInputFormat() string`

GetInputFormat returns the InputFormat field if non-nil, zero value otherwise.

### GetInputFormatOk

`func (o *EncryptGPG) GetInputFormatOk() (*string, bool)`

GetInputFormatOk returns a tuple with the InputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFormat

`func (o *EncryptGPG) SetInputFormat(v string)`

SetInputFormat sets InputFormat field to given value.

### HasInputFormat

`func (o *EncryptGPG) HasInputFormat() bool`

HasInputFormat returns a boolean if a field has been set.

### GetItemId

`func (o *EncryptGPG) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *EncryptGPG) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *EncryptGPG) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *EncryptGPG) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *EncryptGPG) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EncryptGPG) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EncryptGPG) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EncryptGPG) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyName

`func (o *EncryptGPG) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *EncryptGPG) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *EncryptGPG) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetPlaintext

`func (o *EncryptGPG) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *EncryptGPG) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *EncryptGPG) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.


### GetToken

`func (o *EncryptGPG) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EncryptGPG) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EncryptGPG) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EncryptGPG) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EncryptGPG) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EncryptGPG) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EncryptGPG) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EncryptGPG) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


