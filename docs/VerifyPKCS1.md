# VerifyPKCS1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the key to use in the verification process | [optional] 
**HashFunction** | Pointer to **string** | HashFunction defines the hash function (e.g. sha-256) | [optional] 
**InputFormat** | Pointer to **string** | Select default assumed format for the plaintext message. Currently supported options: [base64] | [optional] 
**ItemId** | Pointer to **int64** | The item id of the key to use in the verification process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyName** | **string** | The name of the RSA key to use in the verification process | 
**Message** | **string** | The message to be verified | 
**Prehashed** | Pointer to **bool** | Markes that the message is already hashed | [optional] 
**Signature** | **string** | The message&#39;s signature | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | The version of the key to use for verification | [optional] 

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

### GetDisplayId

`func (o *VerifyPKCS1) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *VerifyPKCS1) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *VerifyPKCS1) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *VerifyPKCS1) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetHashFunction

`func (o *VerifyPKCS1) GetHashFunction() string`

GetHashFunction returns the HashFunction field if non-nil, zero value otherwise.

### GetHashFunctionOk

`func (o *VerifyPKCS1) GetHashFunctionOk() (*string, bool)`

GetHashFunctionOk returns a tuple with the HashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashFunction

`func (o *VerifyPKCS1) SetHashFunction(v string)`

SetHashFunction sets HashFunction field to given value.

### HasHashFunction

`func (o *VerifyPKCS1) HasHashFunction() bool`

HasHashFunction returns a boolean if a field has been set.

### GetInputFormat

`func (o *VerifyPKCS1) GetInputFormat() string`

GetInputFormat returns the InputFormat field if non-nil, zero value otherwise.

### GetInputFormatOk

`func (o *VerifyPKCS1) GetInputFormatOk() (*string, bool)`

GetInputFormatOk returns a tuple with the InputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFormat

`func (o *VerifyPKCS1) SetInputFormat(v string)`

SetInputFormat sets InputFormat field to given value.

### HasInputFormat

`func (o *VerifyPKCS1) HasInputFormat() bool`

HasInputFormat returns a boolean if a field has been set.

### GetItemId

`func (o *VerifyPKCS1) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *VerifyPKCS1) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *VerifyPKCS1) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *VerifyPKCS1) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *VerifyPKCS1) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VerifyPKCS1) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VerifyPKCS1) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *VerifyPKCS1) HasJson() bool`

HasJson returns a boolean if a field has been set.

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


### GetPrehashed

`func (o *VerifyPKCS1) GetPrehashed() bool`

GetPrehashed returns the Prehashed field if non-nil, zero value otherwise.

### GetPrehashedOk

`func (o *VerifyPKCS1) GetPrehashedOk() (*bool, bool)`

GetPrehashedOk returns a tuple with the Prehashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashed

`func (o *VerifyPKCS1) SetPrehashed(v bool)`

SetPrehashed sets Prehashed field to given value.

### HasPrehashed

`func (o *VerifyPKCS1) HasPrehashed() bool`

HasPrehashed returns a boolean if a field has been set.

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

### GetVersion

`func (o *VerifyPKCS1) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VerifyPKCS1) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VerifyPKCS1) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VerifyPKCS1) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


