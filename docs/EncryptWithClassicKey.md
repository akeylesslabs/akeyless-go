# EncryptWithClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | **string** | The name of the key to use in the encryption process | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Plaintext** | **string** | Data to be encrypted | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | **int32** | classic key version | 

## Methods

### NewEncryptWithClassicKey

`func NewEncryptWithClassicKey(displayId string, plaintext string, version int32, ) *EncryptWithClassicKey`

NewEncryptWithClassicKey instantiates a new EncryptWithClassicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptWithClassicKeyWithDefaults

`func NewEncryptWithClassicKeyWithDefaults() *EncryptWithClassicKey`

NewEncryptWithClassicKeyWithDefaults instantiates a new EncryptWithClassicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *EncryptWithClassicKey) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *EncryptWithClassicKey) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *EncryptWithClassicKey) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetJson

`func (o *EncryptWithClassicKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EncryptWithClassicKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EncryptWithClassicKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EncryptWithClassicKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetPlaintext

`func (o *EncryptWithClassicKey) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *EncryptWithClassicKey) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *EncryptWithClassicKey) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.


### GetToken

`func (o *EncryptWithClassicKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EncryptWithClassicKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EncryptWithClassicKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EncryptWithClassicKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EncryptWithClassicKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EncryptWithClassicKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EncryptWithClassicKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EncryptWithClassicKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *EncryptWithClassicKey) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EncryptWithClassicKey) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EncryptWithClassicKey) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


