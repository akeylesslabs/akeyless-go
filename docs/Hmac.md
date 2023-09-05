# Hmac

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the key to use in the encryption process | [optional] 
**HashFunction** | Pointer to **string** | Hash function [sha-256,sha-512] | [optional] [default to "sha-256"]
**InputFormat** | Pointer to **string** | Select default assumed format for any plaintext input. Currently supported options: [base64] | [optional] 
**ItemId** | Pointer to **int64** | The item id of the key to use in the encryption process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyName** | **string** | The name of the key to use in the encryption process | 
**Plaintext** | Pointer to **string** | Data to perform hmac on | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewHmac

`func NewHmac(keyName string, ) *Hmac`

NewHmac instantiates a new Hmac object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHmacWithDefaults

`func NewHmacWithDefaults() *Hmac`

NewHmacWithDefaults instantiates a new Hmac object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *Hmac) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Hmac) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Hmac) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *Hmac) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetHashFunction

`func (o *Hmac) GetHashFunction() string`

GetHashFunction returns the HashFunction field if non-nil, zero value otherwise.

### GetHashFunctionOk

`func (o *Hmac) GetHashFunctionOk() (*string, bool)`

GetHashFunctionOk returns a tuple with the HashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashFunction

`func (o *Hmac) SetHashFunction(v string)`

SetHashFunction sets HashFunction field to given value.

### HasHashFunction

`func (o *Hmac) HasHashFunction() bool`

HasHashFunction returns a boolean if a field has been set.

### GetInputFormat

`func (o *Hmac) GetInputFormat() string`

GetInputFormat returns the InputFormat field if non-nil, zero value otherwise.

### GetInputFormatOk

`func (o *Hmac) GetInputFormatOk() (*string, bool)`

GetInputFormatOk returns a tuple with the InputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFormat

`func (o *Hmac) SetInputFormat(v string)`

SetInputFormat sets InputFormat field to given value.

### HasInputFormat

`func (o *Hmac) HasInputFormat() bool`

HasInputFormat returns a boolean if a field has been set.

### GetItemId

`func (o *Hmac) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *Hmac) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *Hmac) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *Hmac) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *Hmac) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *Hmac) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *Hmac) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *Hmac) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyName

`func (o *Hmac) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *Hmac) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *Hmac) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetPlaintext

`func (o *Hmac) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *Hmac) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *Hmac) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.

### HasPlaintext

`func (o *Hmac) HasPlaintext() bool`

HasPlaintext returns a boolean if a field has been set.

### GetToken

`func (o *Hmac) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Hmac) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Hmac) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Hmac) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *Hmac) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *Hmac) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *Hmac) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *Hmac) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


