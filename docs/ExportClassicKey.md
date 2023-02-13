# ExportClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreCache** | Pointer to **string** | Retrieve the Secret value without checking the Gateway&#39;s cache [true/false]. This flag is only relevant when using the RestAPI | [optional] [default to "false"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | ClassicKey name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | Classic key version | [optional] 

## Methods

### NewExportClassicKey

`func NewExportClassicKey(name string, ) *ExportClassicKey`

NewExportClassicKey instantiates a new ExportClassicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportClassicKeyWithDefaults

`func NewExportClassicKeyWithDefaults() *ExportClassicKey`

NewExportClassicKeyWithDefaults instantiates a new ExportClassicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreCache

`func (o *ExportClassicKey) GetIgnoreCache() string`

GetIgnoreCache returns the IgnoreCache field if non-nil, zero value otherwise.

### GetIgnoreCacheOk

`func (o *ExportClassicKey) GetIgnoreCacheOk() (*string, bool)`

GetIgnoreCacheOk returns a tuple with the IgnoreCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCache

`func (o *ExportClassicKey) SetIgnoreCache(v string)`

SetIgnoreCache sets IgnoreCache field to given value.

### HasIgnoreCache

`func (o *ExportClassicKey) HasIgnoreCache() bool`

HasIgnoreCache returns a boolean if a field has been set.

### GetJson

`func (o *ExportClassicKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ExportClassicKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ExportClassicKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ExportClassicKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *ExportClassicKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportClassicKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportClassicKey) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *ExportClassicKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ExportClassicKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ExportClassicKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ExportClassicKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ExportClassicKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ExportClassicKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ExportClassicKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ExportClassicKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *ExportClassicKey) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExportClassicKey) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExportClassicKey) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExportClassicKey) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


