# SignPKCS1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the key to use in the signing process | [optional] 
**ItemId** | Pointer to **int64** | The item id of the key to use in the signing process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**KeyName** | Pointer to **string** | The name of the RSA key to use in the signing process | [optional] 
**Message** | **string** | The message to be signed | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewSignPKCS1

`func NewSignPKCS1(message string, ) *SignPKCS1`

NewSignPKCS1 instantiates a new SignPKCS1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignPKCS1WithDefaults

`func NewSignPKCS1WithDefaults() *SignPKCS1`

NewSignPKCS1WithDefaults instantiates a new SignPKCS1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *SignPKCS1) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *SignPKCS1) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *SignPKCS1) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *SignPKCS1) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetItemId

`func (o *SignPKCS1) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SignPKCS1) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SignPKCS1) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SignPKCS1) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *SignPKCS1) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *SignPKCS1) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *SignPKCS1) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *SignPKCS1) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyName

`func (o *SignPKCS1) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *SignPKCS1) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *SignPKCS1) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *SignPKCS1) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetMessage

`func (o *SignPKCS1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignPKCS1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignPKCS1) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetToken

`func (o *SignPKCS1) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SignPKCS1) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SignPKCS1) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SignPKCS1) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *SignPKCS1) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *SignPKCS1) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *SignPKCS1) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *SignPKCS1) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


