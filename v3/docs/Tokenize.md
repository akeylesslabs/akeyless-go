# Tokenize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Plaintext** | **string** | Data to be encrypted | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenizerName** | **string** | The name of the tokenizer to use in the encryption process | 
**Tweak** | Pointer to **string** | Base64 encoded tweak for vaultless encryption | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTokenize

`func NewTokenize(plaintext string, tokenizerName string, ) *Tokenize`

NewTokenize instantiates a new Tokenize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizeWithDefaults

`func NewTokenizeWithDefaults() *Tokenize`

NewTokenizeWithDefaults instantiates a new Tokenize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *Tokenize) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *Tokenize) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *Tokenize) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *Tokenize) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetPlaintext

`func (o *Tokenize) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *Tokenize) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *Tokenize) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.


### GetToken

`func (o *Tokenize) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Tokenize) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Tokenize) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Tokenize) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenizerName

`func (o *Tokenize) GetTokenizerName() string`

GetTokenizerName returns the TokenizerName field if non-nil, zero value otherwise.

### GetTokenizerNameOk

`func (o *Tokenize) GetTokenizerNameOk() (*string, bool)`

GetTokenizerNameOk returns a tuple with the TokenizerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizerName

`func (o *Tokenize) SetTokenizerName(v string)`

SetTokenizerName sets TokenizerName field to given value.


### GetTweak

`func (o *Tokenize) GetTweak() string`

GetTweak returns the Tweak field if non-nil, zero value otherwise.

### GetTweakOk

`func (o *Tokenize) GetTweakOk() (*string, bool)`

GetTweakOk returns a tuple with the Tweak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweak

`func (o *Tokenize) SetTweak(v string)`

SetTweak sets Tweak field to given value.

### HasTweak

`func (o *Tokenize) HasTweak() bool`

HasTweak returns a boolean if a field has been set.

### GetUidToken

`func (o *Tokenize) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *Tokenize) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *Tokenize) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *Tokenize) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


