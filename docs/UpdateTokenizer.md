# UpdateTokenizer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**Alphabet** | Pointer to **string** | Alphabet to use in regexp vaultless tokenization | [optional] 
**DecodingTemplate** | Pointer to **string** | The Decoding output template to use in regexp vaultless tokenization | [optional] 
**DeleteProtection** | Pointer to **string** |  | [optional] 
**EncodingTemplate** | Pointer to **string** | The Encoding output template to use in regexp vaultless tokenization | [optional] 
**EncryptionKeyName** | Pointer to **string** | AES key name to use in vaultless tokenization | [optional] 
**Name** | **string** | Current item name | 
**NewMetadata** | Pointer to **string** | New item metadata | [optional] [default to "default_metadata"]
**NewName** | Pointer to **string** | New item name | [optional] 
**Pattern** | Pointer to **string** | Pattern to use in regexp vaultless tokenization | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**TemplateType** | **string** | Which template type this tokenizer is used for [SSN,CreditCard,USPhoneNumber,Email,Regexp] | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenizerType** | **string** | Tokenizer type | 
**TweakType** | Pointer to **string** | The tweak type to use in vaultless tokenization [Supplied, Generated, Internal, Masking] | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateTokenizer

`func NewUpdateTokenizer(name string, templateType string, tokenizerType string, ) *UpdateTokenizer`

NewUpdateTokenizer instantiates a new UpdateTokenizer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTokenizerWithDefaults

`func NewUpdateTokenizerWithDefaults() *UpdateTokenizer`

NewUpdateTokenizerWithDefaults instantiates a new UpdateTokenizer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *UpdateTokenizer) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *UpdateTokenizer) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *UpdateTokenizer) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *UpdateTokenizer) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAlphabet

`func (o *UpdateTokenizer) GetAlphabet() string`

GetAlphabet returns the Alphabet field if non-nil, zero value otherwise.

### GetAlphabetOk

`func (o *UpdateTokenizer) GetAlphabetOk() (*string, bool)`

GetAlphabetOk returns a tuple with the Alphabet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphabet

`func (o *UpdateTokenizer) SetAlphabet(v string)`

SetAlphabet sets Alphabet field to given value.

### HasAlphabet

`func (o *UpdateTokenizer) HasAlphabet() bool`

HasAlphabet returns a boolean if a field has been set.

### GetDecodingTemplate

`func (o *UpdateTokenizer) GetDecodingTemplate() string`

GetDecodingTemplate returns the DecodingTemplate field if non-nil, zero value otherwise.

### GetDecodingTemplateOk

`func (o *UpdateTokenizer) GetDecodingTemplateOk() (*string, bool)`

GetDecodingTemplateOk returns a tuple with the DecodingTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodingTemplate

`func (o *UpdateTokenizer) SetDecodingTemplate(v string)`

SetDecodingTemplate sets DecodingTemplate field to given value.

### HasDecodingTemplate

`func (o *UpdateTokenizer) HasDecodingTemplate() bool`

HasDecodingTemplate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *UpdateTokenizer) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *UpdateTokenizer) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *UpdateTokenizer) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *UpdateTokenizer) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetEncodingTemplate

`func (o *UpdateTokenizer) GetEncodingTemplate() string`

GetEncodingTemplate returns the EncodingTemplate field if non-nil, zero value otherwise.

### GetEncodingTemplateOk

`func (o *UpdateTokenizer) GetEncodingTemplateOk() (*string, bool)`

GetEncodingTemplateOk returns a tuple with the EncodingTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingTemplate

`func (o *UpdateTokenizer) SetEncodingTemplate(v string)`

SetEncodingTemplate sets EncodingTemplate field to given value.

### HasEncodingTemplate

`func (o *UpdateTokenizer) HasEncodingTemplate() bool`

HasEncodingTemplate returns a boolean if a field has been set.

### GetEncryptionKeyName

`func (o *UpdateTokenizer) GetEncryptionKeyName() string`

GetEncryptionKeyName returns the EncryptionKeyName field if non-nil, zero value otherwise.

### GetEncryptionKeyNameOk

`func (o *UpdateTokenizer) GetEncryptionKeyNameOk() (*string, bool)`

GetEncryptionKeyNameOk returns a tuple with the EncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyName

`func (o *UpdateTokenizer) SetEncryptionKeyName(v string)`

SetEncryptionKeyName sets EncryptionKeyName field to given value.

### HasEncryptionKeyName

`func (o *UpdateTokenizer) HasEncryptionKeyName() bool`

HasEncryptionKeyName returns a boolean if a field has been set.

### GetName

`func (o *UpdateTokenizer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTokenizer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTokenizer) SetName(v string)`

SetName sets Name field to given value.


### GetNewMetadata

`func (o *UpdateTokenizer) GetNewMetadata() string`

GetNewMetadata returns the NewMetadata field if non-nil, zero value otherwise.

### GetNewMetadataOk

`func (o *UpdateTokenizer) GetNewMetadataOk() (*string, bool)`

GetNewMetadataOk returns a tuple with the NewMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMetadata

`func (o *UpdateTokenizer) SetNewMetadata(v string)`

SetNewMetadata sets NewMetadata field to given value.

### HasNewMetadata

`func (o *UpdateTokenizer) HasNewMetadata() bool`

HasNewMetadata returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateTokenizer) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateTokenizer) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateTokenizer) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateTokenizer) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPattern

`func (o *UpdateTokenizer) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *UpdateTokenizer) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *UpdateTokenizer) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *UpdateTokenizer) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRmTag

`func (o *UpdateTokenizer) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *UpdateTokenizer) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *UpdateTokenizer) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *UpdateTokenizer) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetTemplateType

`func (o *UpdateTokenizer) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *UpdateTokenizer) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *UpdateTokenizer) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.


### GetToken

`func (o *UpdateTokenizer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateTokenizer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateTokenizer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateTokenizer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenizerType

`func (o *UpdateTokenizer) GetTokenizerType() string`

GetTokenizerType returns the TokenizerType field if non-nil, zero value otherwise.

### GetTokenizerTypeOk

`func (o *UpdateTokenizer) GetTokenizerTypeOk() (*string, bool)`

GetTokenizerTypeOk returns a tuple with the TokenizerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizerType

`func (o *UpdateTokenizer) SetTokenizerType(v string)`

SetTokenizerType sets TokenizerType field to given value.


### GetTweakType

`func (o *UpdateTokenizer) GetTweakType() string`

GetTweakType returns the TweakType field if non-nil, zero value otherwise.

### GetTweakTypeOk

`func (o *UpdateTokenizer) GetTweakTypeOk() (*string, bool)`

GetTweakTypeOk returns a tuple with the TweakType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweakType

`func (o *UpdateTokenizer) SetTweakType(v string)`

SetTweakType sets TweakType field to given value.

### HasTweakType

`func (o *UpdateTokenizer) HasTweakType() bool`

HasTweakType returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateTokenizer) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateTokenizer) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateTokenizer) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateTokenizer) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


