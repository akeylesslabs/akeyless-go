# VerifyRsaSsaPss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the RSA key to use in the verification process | [optional] 
**HashFunction** | Pointer to **string** | HashFunction defines the hash function (e.g. sha-256) | [optional] 
**ItemId** | Pointer to **int64** | The item id of the RSA key to use in the verification process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyName** | Pointer to **string** | The name of the RSA key to use in the verification process | [optional] 
**Message** | **string** | The input message to verify in a base64 format | 
**Prehashed** | Pointer to **bool** | Markes that the message is already hashed | [optional] 
**Signature** | **string** | The message&#39;s signature | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | The version of the key to use for verification | [optional] 

## Methods

### NewVerifyRsaSsaPss

`func NewVerifyRsaSsaPss(message string, signature string, ) *VerifyRsaSsaPss`

NewVerifyRsaSsaPss instantiates a new VerifyRsaSsaPss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyRsaSsaPssWithDefaults

`func NewVerifyRsaSsaPssWithDefaults() *VerifyRsaSsaPss`

NewVerifyRsaSsaPssWithDefaults instantiates a new VerifyRsaSsaPss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *VerifyRsaSsaPss) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *VerifyRsaSsaPss) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *VerifyRsaSsaPss) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *VerifyRsaSsaPss) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetHashFunction

`func (o *VerifyRsaSsaPss) GetHashFunction() string`

GetHashFunction returns the HashFunction field if non-nil, zero value otherwise.

### GetHashFunctionOk

`func (o *VerifyRsaSsaPss) GetHashFunctionOk() (*string, bool)`

GetHashFunctionOk returns a tuple with the HashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashFunction

`func (o *VerifyRsaSsaPss) SetHashFunction(v string)`

SetHashFunction sets HashFunction field to given value.

### HasHashFunction

`func (o *VerifyRsaSsaPss) HasHashFunction() bool`

HasHashFunction returns a boolean if a field has been set.

### GetItemId

`func (o *VerifyRsaSsaPss) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *VerifyRsaSsaPss) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *VerifyRsaSsaPss) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *VerifyRsaSsaPss) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *VerifyRsaSsaPss) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VerifyRsaSsaPss) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VerifyRsaSsaPss) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *VerifyRsaSsaPss) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyName

`func (o *VerifyRsaSsaPss) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *VerifyRsaSsaPss) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *VerifyRsaSsaPss) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *VerifyRsaSsaPss) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetMessage

`func (o *VerifyRsaSsaPss) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VerifyRsaSsaPss) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VerifyRsaSsaPss) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPrehashed

`func (o *VerifyRsaSsaPss) GetPrehashed() bool`

GetPrehashed returns the Prehashed field if non-nil, zero value otherwise.

### GetPrehashedOk

`func (o *VerifyRsaSsaPss) GetPrehashedOk() (*bool, bool)`

GetPrehashedOk returns a tuple with the Prehashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashed

`func (o *VerifyRsaSsaPss) SetPrehashed(v bool)`

SetPrehashed sets Prehashed field to given value.

### HasPrehashed

`func (o *VerifyRsaSsaPss) HasPrehashed() bool`

HasPrehashed returns a boolean if a field has been set.

### GetSignature

`func (o *VerifyRsaSsaPss) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerifyRsaSsaPss) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerifyRsaSsaPss) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetToken

`func (o *VerifyRsaSsaPss) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VerifyRsaSsaPss) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VerifyRsaSsaPss) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VerifyRsaSsaPss) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *VerifyRsaSsaPss) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *VerifyRsaSsaPss) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *VerifyRsaSsaPss) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *VerifyRsaSsaPss) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *VerifyRsaSsaPss) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VerifyRsaSsaPss) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VerifyRsaSsaPss) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VerifyRsaSsaPss) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


