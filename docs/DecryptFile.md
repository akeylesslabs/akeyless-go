# DecryptFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncContext** | Pointer to **map[string]string** | The encryption context. If this was specified in the encrypt command, it must be specified here or the decryption operation will fail | [optional] 
**InputFile** | **string** | Path to the file to be decrypted. If not provided, the content will be taken from stdin | 
**KeyName** | **string** | The name of the key to use in the decryption process | 
**OutputFilePath** | Pointer to **string** | Path to the output file. If not provided, the output will be sent to stdout | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDecryptFile

`func NewDecryptFile(inputFile string, keyName string, ) *DecryptFile`

NewDecryptFile instantiates a new DecryptFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptFileWithDefaults

`func NewDecryptFileWithDefaults() *DecryptFile`

NewDecryptFileWithDefaults instantiates a new DecryptFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncContext

`func (o *DecryptFile) GetEncContext() map[string]string`

GetEncContext returns the EncContext field if non-nil, zero value otherwise.

### GetEncContextOk

`func (o *DecryptFile) GetEncContextOk() (*map[string]string, bool)`

GetEncContextOk returns a tuple with the EncContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncContext

`func (o *DecryptFile) SetEncContext(v map[string]string)`

SetEncContext sets EncContext field to given value.

### HasEncContext

`func (o *DecryptFile) HasEncContext() bool`

HasEncContext returns a boolean if a field has been set.

### GetInputFile

`func (o *DecryptFile) GetInputFile() string`

GetInputFile returns the InputFile field if non-nil, zero value otherwise.

### GetInputFileOk

`func (o *DecryptFile) GetInputFileOk() (*string, bool)`

GetInputFileOk returns a tuple with the InputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFile

`func (o *DecryptFile) SetInputFile(v string)`

SetInputFile sets InputFile field to given value.


### GetKeyName

`func (o *DecryptFile) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *DecryptFile) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *DecryptFile) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetOutputFilePath

`func (o *DecryptFile) GetOutputFilePath() string`

GetOutputFilePath returns the OutputFilePath field if non-nil, zero value otherwise.

### GetOutputFilePathOk

`func (o *DecryptFile) GetOutputFilePathOk() (*string, bool)`

GetOutputFilePathOk returns a tuple with the OutputFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFilePath

`func (o *DecryptFile) SetOutputFilePath(v string)`

SetOutputFilePath sets OutputFilePath field to given value.

### HasOutputFilePath

`func (o *DecryptFile) HasOutputFilePath() bool`

HasOutputFilePath returns a boolean if a field has been set.

### GetToken

`func (o *DecryptFile) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DecryptFile) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DecryptFile) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DecryptFile) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DecryptFile) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DecryptFile) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DecryptFile) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DecryptFile) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


