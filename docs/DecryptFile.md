# DecryptFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CyphertextHeader** | Pointer to **string** |  | [optional] 
**DisplayId** | Pointer to **string** | The display id of the key to use in the decryption process | [optional] 
**In** | **string** | Path to the file to be decrypted. If not provided, the content will be taken from stdin | 
**ItemId** | Pointer to **int64** | The item id of the key to use in the decryption process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyName** | **string** | The name of the key to use in the decryption process | 
**Out** | Pointer to **string** | Path to the output file. If not provided, the output will be sent to stdout | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | key version (relevant only for classic key) | [optional] 

## Methods

### NewDecryptFile

`func NewDecryptFile(in string, keyName string, ) *DecryptFile`

NewDecryptFile instantiates a new DecryptFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptFileWithDefaults

`func NewDecryptFileWithDefaults() *DecryptFile`

NewDecryptFileWithDefaults instantiates a new DecryptFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCyphertextHeader

`func (o *DecryptFile) GetCyphertextHeader() string`

GetCyphertextHeader returns the CyphertextHeader field if non-nil, zero value otherwise.

### GetCyphertextHeaderOk

`func (o *DecryptFile) GetCyphertextHeaderOk() (*string, bool)`

GetCyphertextHeaderOk returns a tuple with the CyphertextHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyphertextHeader

`func (o *DecryptFile) SetCyphertextHeader(v string)`

SetCyphertextHeader sets CyphertextHeader field to given value.

### HasCyphertextHeader

`func (o *DecryptFile) HasCyphertextHeader() bool`

HasCyphertextHeader returns a boolean if a field has been set.

### GetDisplayId

`func (o *DecryptFile) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *DecryptFile) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *DecryptFile) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *DecryptFile) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetIn

`func (o *DecryptFile) GetIn() string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *DecryptFile) GetInOk() (*string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *DecryptFile) SetIn(v string)`

SetIn sets In field to given value.


### GetItemId

`func (o *DecryptFile) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DecryptFile) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DecryptFile) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DecryptFile) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *DecryptFile) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DecryptFile) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DecryptFile) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DecryptFile) HasJson() bool`

HasJson returns a boolean if a field has been set.

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


### GetOut

`func (o *DecryptFile) GetOut() string`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *DecryptFile) GetOutOk() (*string, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *DecryptFile) SetOut(v string)`

SetOut sets Out field to given value.

### HasOut

`func (o *DecryptFile) HasOut() bool`

HasOut returns a boolean if a field has been set.

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

### GetVersion

`func (o *DecryptFile) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DecryptFile) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DecryptFile) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DecryptFile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


