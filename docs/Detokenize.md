# Detokenize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | **string** | Data to be decrypted | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenizerName** | **string** | The name of the tokenizer to use in the decryption process | 
**Tweak** | Pointer to **string** | Base64 encoded tweak for vaultless encryption | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDetokenize

`func NewDetokenize(ciphertext string, tokenizerName string, ) *Detokenize`

NewDetokenize instantiates a new Detokenize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetokenizeWithDefaults

`func NewDetokenizeWithDefaults() *Detokenize`

NewDetokenizeWithDefaults instantiates a new Detokenize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *Detokenize) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *Detokenize) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *Detokenize) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### GetJson

`func (o *Detokenize) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *Detokenize) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *Detokenize) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *Detokenize) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *Detokenize) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Detokenize) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Detokenize) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Detokenize) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenizerName

`func (o *Detokenize) GetTokenizerName() string`

GetTokenizerName returns the TokenizerName field if non-nil, zero value otherwise.

### GetTokenizerNameOk

`func (o *Detokenize) GetTokenizerNameOk() (*string, bool)`

GetTokenizerNameOk returns a tuple with the TokenizerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizerName

`func (o *Detokenize) SetTokenizerName(v string)`

SetTokenizerName sets TokenizerName field to given value.


### GetTweak

`func (o *Detokenize) GetTweak() string`

GetTweak returns the Tweak field if non-nil, zero value otherwise.

### GetTweakOk

`func (o *Detokenize) GetTweakOk() (*string, bool)`

GetTweakOk returns a tuple with the Tweak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweak

`func (o *Detokenize) SetTweak(v string)`

SetTweak sets Tweak field to given value.

### HasTweak

`func (o *Detokenize) HasTweak() bool`

HasTweak returns a boolean if a field has been set.

### GetUidToken

`func (o *Detokenize) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *Detokenize) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *Detokenize) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *Detokenize) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


