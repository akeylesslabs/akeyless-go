# CreateTokenizer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alphabet** | Pointer to **string** | Alphabet to use in regexp vaultless tokenization | [optional] 
**DecodingTemplate** | Pointer to **string** | The Decoding output template to use in regexp vaultless tokenization | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item | [optional] 
**EncodingTemplate** | Pointer to **string** | The Encoding output template to use in regexp vaultless tokenization | [optional] 
**EncryptionKeyName** | Pointer to **string** | AES key name to use in vaultless tokenization | [optional] 
**Metadata** | Pointer to **string** | A metadata about the tokenizer | [optional] 
**Name** | **string** | Tokenizer name | 
**Pattern** | Pointer to **string** | Pattern to use in regexp vaultless tokenization | [optional] 
**Tag** | Pointer to **[]string** | List of the tags attached to this key | [optional] 
**TemplateType** | **string** | Which template type this tokenizer is used for [SSN,CreditCard,USPhoneNumber,Email,Regexp] | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenizerType** | **string** | Tokenizer type | 
**TweakType** | Pointer to **string** | The tweak type to use in vaultless tokenization [Supplied, Generated, Internal, Masking] | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateTokenizer

`func NewCreateTokenizer(name string, templateType string, tokenizerType string, ) *CreateTokenizer`

NewCreateTokenizer instantiates a new CreateTokenizer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenizerWithDefaults

`func NewCreateTokenizerWithDefaults() *CreateTokenizer`

NewCreateTokenizerWithDefaults instantiates a new CreateTokenizer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlphabet

`func (o *CreateTokenizer) GetAlphabet() string`

GetAlphabet returns the Alphabet field if non-nil, zero value otherwise.

### GetAlphabetOk

`func (o *CreateTokenizer) GetAlphabetOk() (*string, bool)`

GetAlphabetOk returns a tuple with the Alphabet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphabet

`func (o *CreateTokenizer) SetAlphabet(v string)`

SetAlphabet sets Alphabet field to given value.

### HasAlphabet

`func (o *CreateTokenizer) HasAlphabet() bool`

HasAlphabet returns a boolean if a field has been set.

### GetDecodingTemplate

`func (o *CreateTokenizer) GetDecodingTemplate() string`

GetDecodingTemplate returns the DecodingTemplate field if non-nil, zero value otherwise.

### GetDecodingTemplateOk

`func (o *CreateTokenizer) GetDecodingTemplateOk() (*string, bool)`

GetDecodingTemplateOk returns a tuple with the DecodingTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodingTemplate

`func (o *CreateTokenizer) SetDecodingTemplate(v string)`

SetDecodingTemplate sets DecodingTemplate field to given value.

### HasDecodingTemplate

`func (o *CreateTokenizer) HasDecodingTemplate() bool`

HasDecodingTemplate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateTokenizer) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateTokenizer) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateTokenizer) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateTokenizer) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetEncodingTemplate

`func (o *CreateTokenizer) GetEncodingTemplate() string`

GetEncodingTemplate returns the EncodingTemplate field if non-nil, zero value otherwise.

### GetEncodingTemplateOk

`func (o *CreateTokenizer) GetEncodingTemplateOk() (*string, bool)`

GetEncodingTemplateOk returns a tuple with the EncodingTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingTemplate

`func (o *CreateTokenizer) SetEncodingTemplate(v string)`

SetEncodingTemplate sets EncodingTemplate field to given value.

### HasEncodingTemplate

`func (o *CreateTokenizer) HasEncodingTemplate() bool`

HasEncodingTemplate returns a boolean if a field has been set.

### GetEncryptionKeyName

`func (o *CreateTokenizer) GetEncryptionKeyName() string`

GetEncryptionKeyName returns the EncryptionKeyName field if non-nil, zero value otherwise.

### GetEncryptionKeyNameOk

`func (o *CreateTokenizer) GetEncryptionKeyNameOk() (*string, bool)`

GetEncryptionKeyNameOk returns a tuple with the EncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyName

`func (o *CreateTokenizer) SetEncryptionKeyName(v string)`

SetEncryptionKeyName sets EncryptionKeyName field to given value.

### HasEncryptionKeyName

`func (o *CreateTokenizer) HasEncryptionKeyName() bool`

HasEncryptionKeyName returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateTokenizer) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateTokenizer) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateTokenizer) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateTokenizer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateTokenizer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTokenizer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTokenizer) SetName(v string)`

SetName sets Name field to given value.


### GetPattern

`func (o *CreateTokenizer) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateTokenizer) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateTokenizer) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CreateTokenizer) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetTag

`func (o *CreateTokenizer) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTokenizer) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTokenizer) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTokenizer) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTemplateType

`func (o *CreateTokenizer) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *CreateTokenizer) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *CreateTokenizer) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.


### GetToken

`func (o *CreateTokenizer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateTokenizer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateTokenizer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateTokenizer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenizerType

`func (o *CreateTokenizer) GetTokenizerType() string`

GetTokenizerType returns the TokenizerType field if non-nil, zero value otherwise.

### GetTokenizerTypeOk

`func (o *CreateTokenizer) GetTokenizerTypeOk() (*string, bool)`

GetTokenizerTypeOk returns a tuple with the TokenizerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizerType

`func (o *CreateTokenizer) SetTokenizerType(v string)`

SetTokenizerType sets TokenizerType field to given value.


### GetTweakType

`func (o *CreateTokenizer) GetTweakType() string`

GetTweakType returns the TweakType field if non-nil, zero value otherwise.

### GetTweakTypeOk

`func (o *CreateTokenizer) GetTweakTypeOk() (*string, bool)`

GetTweakTypeOk returns a tuple with the TweakType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweakType

`func (o *CreateTokenizer) SetTweakType(v string)`

SetTweakType sets TweakType field to given value.

### HasTweakType

`func (o *CreateTokenizer) HasTweakType() bool`

HasTweakType returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateTokenizer) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateTokenizer) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateTokenizer) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateTokenizer) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


