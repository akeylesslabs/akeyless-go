# SignEcDsa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the EC key to use for the signing process | [optional] 
**ItemId** | Pointer to **int64** | The item id of the EC key to use for the signing process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyName** | Pointer to **string** | The name of the EC key to use for the signing process | [optional] 
**Message** | **string** | The input message to sign in a base64 format | 
**Prehashed** | Pointer to **bool** | Markes that the message is already hashed | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | The version of the key to use for signing | [optional] 

## Methods

### NewSignEcDsa

`func NewSignEcDsa(message string, ) *SignEcDsa`

NewSignEcDsa instantiates a new SignEcDsa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignEcDsaWithDefaults

`func NewSignEcDsaWithDefaults() *SignEcDsa`

NewSignEcDsaWithDefaults instantiates a new SignEcDsa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *SignEcDsa) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *SignEcDsa) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *SignEcDsa) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *SignEcDsa) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetItemId

`func (o *SignEcDsa) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SignEcDsa) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SignEcDsa) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SignEcDsa) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *SignEcDsa) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *SignEcDsa) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *SignEcDsa) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *SignEcDsa) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyName

`func (o *SignEcDsa) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *SignEcDsa) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *SignEcDsa) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *SignEcDsa) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetMessage

`func (o *SignEcDsa) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignEcDsa) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignEcDsa) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPrehashed

`func (o *SignEcDsa) GetPrehashed() bool`

GetPrehashed returns the Prehashed field if non-nil, zero value otherwise.

### GetPrehashedOk

`func (o *SignEcDsa) GetPrehashedOk() (*bool, bool)`

GetPrehashedOk returns a tuple with the Prehashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashed

`func (o *SignEcDsa) SetPrehashed(v bool)`

SetPrehashed sets Prehashed field to given value.

### HasPrehashed

`func (o *SignEcDsa) HasPrehashed() bool`

HasPrehashed returns a boolean if a field has been set.

### GetToken

`func (o *SignEcDsa) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SignEcDsa) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SignEcDsa) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SignEcDsa) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *SignEcDsa) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *SignEcDsa) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *SignEcDsa) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *SignEcDsa) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *SignEcDsa) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SignEcDsa) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SignEcDsa) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SignEcDsa) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


