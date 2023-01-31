# RegexpTokenizerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alphabet** | Pointer to **string** | The Alphabet used for the tokenization | [optional] 
**DecodingTemplate** | Pointer to **string** | Transformation to perform on the decrypted data | [optional] 
**EncodingTemplate** | Pointer to **string** | Transformation to perform on the encrypted data, if the required output template doesn&#39;t match the input string The output Should still be valid for the Pattern, otherwise the secret would be able to be decrypted. | [optional] 
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

### GetDecodingTemplate

`func (o *RegexpTokenizerInfo) GetDecodingTemplate() string`

GetDecodingTemplate returns the DecodingTemplate field if non-nil, zero value otherwise.

### GetDecodingTemplateOk

`func (o *RegexpTokenizerInfo) GetDecodingTemplateOk() (*string, bool)`

GetDecodingTemplateOk returns a tuple with the DecodingTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodingTemplate

`func (o *RegexpTokenizerInfo) SetDecodingTemplate(v string)`

SetDecodingTemplate sets DecodingTemplate field to given value.

### HasDecodingTemplate

`func (o *RegexpTokenizerInfo) HasDecodingTemplate() bool`

HasDecodingTemplate returns a boolean if a field has been set.

### GetEncodingTemplate

`func (o *RegexpTokenizerInfo) GetEncodingTemplate() string`

GetEncodingTemplate returns the EncodingTemplate field if non-nil, zero value otherwise.

### GetEncodingTemplateOk

`func (o *RegexpTokenizerInfo) GetEncodingTemplateOk() (*string, bool)`

GetEncodingTemplateOk returns a tuple with the EncodingTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingTemplate

`func (o *RegexpTokenizerInfo) SetEncodingTemplate(v string)`

SetEncodingTemplate sets EncodingTemplate field to given value.

### HasEncodingTemplate

`func (o *RegexpTokenizerInfo) HasEncodingTemplate() bool`

HasEncodingTemplate returns a boolean if a field has been set.

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


