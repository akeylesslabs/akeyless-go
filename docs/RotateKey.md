# RotateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | **string** | Key name | 
**NewKeyData** | Pointer to **string** | The new base64 encoded value for the classic key. relevant only for keys provided by user (&#39;bring-your-own-key&#39;) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotateKey

`func NewRotateKey(name string, ) *RotateKey`

NewRotateKey instantiates a new RotateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateKeyWithDefaults

`func NewRotateKeyWithDefaults() *RotateKey`

NewRotateKeyWithDefaults instantiates a new RotateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *RotateKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotateKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotateKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotateKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RotateKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotateKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotateKey) SetName(v string)`

SetName sets Name field to given value.


### GetNewKeyData

`func (o *RotateKey) GetNewKeyData() string`

GetNewKeyData returns the NewKeyData field if non-nil, zero value otherwise.

### GetNewKeyDataOk

`func (o *RotateKey) GetNewKeyDataOk() (*string, bool)`

GetNewKeyDataOk returns a tuple with the NewKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewKeyData

`func (o *RotateKey) SetNewKeyData(v string)`

SetNewKeyData sets NewKeyData field to given value.

### HasNewKeyData

`func (o *RotateKey) HasNewKeyData() bool`

HasNewKeyData returns a boolean if a field has been set.

### GetToken

`func (o *RotateKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotateKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotateKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotateKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotateKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotateKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotateKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotateKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


