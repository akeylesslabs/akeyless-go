# SignDataWithClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Data | 
**DisplayId** | **string** | The name of the key to use in the sign data process | 
**Hashed** | Pointer to **bool** | Defines whether the data should be hashed as part of the signing. If true, the data will not be hashed | [optional] [default to false]
**HashingMethod** | Pointer to **string** | HashingMethod | [optional] [default to "SHA256"]
**IgnoreCache** | Pointer to **string** | Retrieve the Secret value without checking the Gateway&#39;s cache [true/false]. This flag is only relevant when using the RestAPI | [optional] [default to "false"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | ClassicKey name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | **int32** | classic key version | 

## Methods

### NewSignDataWithClassicKey

`func NewSignDataWithClassicKey(data string, displayId string, name string, version int32, ) *SignDataWithClassicKey`

NewSignDataWithClassicKey instantiates a new SignDataWithClassicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignDataWithClassicKeyWithDefaults

`func NewSignDataWithClassicKeyWithDefaults() *SignDataWithClassicKey`

NewSignDataWithClassicKeyWithDefaults instantiates a new SignDataWithClassicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SignDataWithClassicKey) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SignDataWithClassicKey) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SignDataWithClassicKey) SetData(v string)`

SetData sets Data field to given value.


### GetDisplayId

`func (o *SignDataWithClassicKey) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *SignDataWithClassicKey) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *SignDataWithClassicKey) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetHashed

`func (o *SignDataWithClassicKey) GetHashed() bool`

GetHashed returns the Hashed field if non-nil, zero value otherwise.

### GetHashedOk

`func (o *SignDataWithClassicKey) GetHashedOk() (*bool, bool)`

GetHashedOk returns a tuple with the Hashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashed

`func (o *SignDataWithClassicKey) SetHashed(v bool)`

SetHashed sets Hashed field to given value.

### HasHashed

`func (o *SignDataWithClassicKey) HasHashed() bool`

HasHashed returns a boolean if a field has been set.

### GetHashingMethod

`func (o *SignDataWithClassicKey) GetHashingMethod() string`

GetHashingMethod returns the HashingMethod field if non-nil, zero value otherwise.

### GetHashingMethodOk

`func (o *SignDataWithClassicKey) GetHashingMethodOk() (*string, bool)`

GetHashingMethodOk returns a tuple with the HashingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashingMethod

`func (o *SignDataWithClassicKey) SetHashingMethod(v string)`

SetHashingMethod sets HashingMethod field to given value.

### HasHashingMethod

`func (o *SignDataWithClassicKey) HasHashingMethod() bool`

HasHashingMethod returns a boolean if a field has been set.

### GetIgnoreCache

`func (o *SignDataWithClassicKey) GetIgnoreCache() string`

GetIgnoreCache returns the IgnoreCache field if non-nil, zero value otherwise.

### GetIgnoreCacheOk

`func (o *SignDataWithClassicKey) GetIgnoreCacheOk() (*string, bool)`

GetIgnoreCacheOk returns a tuple with the IgnoreCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCache

`func (o *SignDataWithClassicKey) SetIgnoreCache(v string)`

SetIgnoreCache sets IgnoreCache field to given value.

### HasIgnoreCache

`func (o *SignDataWithClassicKey) HasIgnoreCache() bool`

HasIgnoreCache returns a boolean if a field has been set.

### GetJson

`func (o *SignDataWithClassicKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *SignDataWithClassicKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *SignDataWithClassicKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *SignDataWithClassicKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *SignDataWithClassicKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignDataWithClassicKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignDataWithClassicKey) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *SignDataWithClassicKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SignDataWithClassicKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SignDataWithClassicKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SignDataWithClassicKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *SignDataWithClassicKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *SignDataWithClassicKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *SignDataWithClassicKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *SignDataWithClassicKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *SignDataWithClassicKey) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SignDataWithClassicKey) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SignDataWithClassicKey) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


