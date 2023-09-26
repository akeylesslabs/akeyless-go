# EncryptFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the key to use in the encryption process | [optional] 
**In** | **string** | Path to the file to be encrypted. If not provided, the content will be taken from stdin | 
**ItemId** | Pointer to **int64** | The item id of the key to use in the encryption process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyName** | **string** | The name of the key to use in the encryption process | 
**Out** | Pointer to **string** | Path to the output file. If not provided, the output will be sent to stdout | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewEncryptFile

`func NewEncryptFile(in string, keyName string, ) *EncryptFile`

NewEncryptFile instantiates a new EncryptFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptFileWithDefaults

`func NewEncryptFileWithDefaults() *EncryptFile`

NewEncryptFileWithDefaults instantiates a new EncryptFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *EncryptFile) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *EncryptFile) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *EncryptFile) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *EncryptFile) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetIn

`func (o *EncryptFile) GetIn() string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *EncryptFile) GetInOk() (*string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *EncryptFile) SetIn(v string)`

SetIn sets In field to given value.


### GetItemId

`func (o *EncryptFile) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *EncryptFile) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *EncryptFile) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *EncryptFile) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *EncryptFile) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EncryptFile) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EncryptFile) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EncryptFile) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyName

`func (o *EncryptFile) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *EncryptFile) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *EncryptFile) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetOut

`func (o *EncryptFile) GetOut() string`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *EncryptFile) GetOutOk() (*string, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *EncryptFile) SetOut(v string)`

SetOut sets Out field to given value.

### HasOut

`func (o *EncryptFile) HasOut() bool`

HasOut returns a boolean if a field has been set.

### GetToken

`func (o *EncryptFile) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EncryptFile) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EncryptFile) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EncryptFile) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EncryptFile) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EncryptFile) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EncryptFile) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EncryptFile) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


