# RegexpTokenizerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alphabet** | Pointer to **string** | The Alphabet used for the tokenization | [optional] 
**DecryptionTemplate** | Pointer to **string** | Transformation to perform on the decrypted data | [optional] 
**EncryptionTemplate** | Pointer to **string** | Transformation to perform on the encrypted data, if the required output template doesn&#39;t match the input string The output Should still be valid for the Pattern, otherwise the secret would be able to be decrypted. | [optional] 
**Pattern** | Pointer to **string** | Regexp pattern to extract and deposit the text/encdata | [optional] 

## Methods

### NewRegexpTokenizerInfo

`func NewRegexpTokenizerInfo() *RegexpTokenizerInfo`

NewRegexpTokenizerInfo instantiates a new RegexpTokenizerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegexpTokenizerInfoWithDefaults

`func NewRegexpTokenizerInfoWithDefaults() *RegexpTokenizerInfo`

NewRegexpTokenizerInfoWithDefaults instantiates a new RegexpTokenizerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlphabet

`func (o *RegexpTokenizerInfo) GetAlphabet() string`

GetAlphabet returns the Alphabet field if non-nil, zero value otherwise.

### GetAlphabetOk

`func (o *RegexpTokenizerInfo) GetAlphabetOk() (*string, bool)`

GetAlphabetOk returns a tuple with the Alphabet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphabet

`func (o *RegexpTokenizerInfo) SetAlphabet(v string)`

SetAlphabet sets Alphabet field to given value.

### HasAlphabet

`func (o *RegexpTokenizerInfo) HasAlphabet() bool`

HasAlphabet returns a boolean if a field has been set.

### GetDecryptionTemplate

`func (o *RegexpTokenizerInfo) GetDecryptionTemplate() string`

GetDecryptionTemplate returns the DecryptionTemplate field if non-nil, zero value otherwise.

### GetDecryptionTemplateOk

`func (o *RegexpTokenizerInfo) GetDecryptionTemplateOk() (*string, bool)`

GetDecryptionTemplateOk returns a tuple with the DecryptionTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptionTemplate

`func (o *RegexpTokenizerInfo) SetDecryptionTemplate(v string)`

SetDecryptionTemplate sets DecryptionTemplate field to given value.

### HasDecryptionTemplate

`func (o *RegexpTokenizerInfo) HasDecryptionTemplate() bool`

HasDecryptionTemplate returns a boolean if a field has been set.

### GetEncryptionTemplate

`func (o *RegexpTokenizerInfo) GetEncryptionTemplate() string`

GetEncryptionTemplate returns the EncryptionTemplate field if non-nil, zero value otherwise.

### GetEncryptionTemplateOk

`func (o *RegexpTokenizerInfo) GetEncryptionTemplateOk() (*string, bool)`

GetEncryptionTemplateOk returns a tuple with the EncryptionTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionTemplate

`func (o *RegexpTokenizerInfo) SetEncryptionTemplate(v string)`

SetEncryptionTemplate sets EncryptionTemplate field to given value.

### HasEncryptionTemplate

`func (o *RegexpTokenizerInfo) HasEncryptionTemplate() bool`

HasEncryptionTemplate returns a boolean if a field has been set.

### GetPattern

`func (o *RegexpTokenizerInfo) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *RegexpTokenizerInfo) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *RegexpTokenizerInfo) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *RegexpTokenizerInfo) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


