# DecryptFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the key to use in the decryption process | [optional] 
**ItemId** | Pointer to **int64** | The item id of the key to use in the decryption process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**KeyName** | **string** | The name of the key to use in the decryption process | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDecryptFile

`func NewDecryptFile(keyName string, ) *DecryptFile`

NewDecryptFile instantiates a new DecryptFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptFileWithDefaults

`func NewDecryptFileWithDefaults() *DecryptFile`

NewDecryptFileWithDefaults instantiates a new DecryptFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


