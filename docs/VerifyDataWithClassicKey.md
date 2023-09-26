# VerifyDataWithClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Data | 
**DisplayId** | Pointer to **string** | The display id of the key to use in the verification process | [optional] 
**Hashed** | Pointer to **bool** | Defines whether the data should be hashed as part of the signing. If true, the data will not be hashed | [optional] [default to false]
**HashingMethod** | Pointer to **string** | HashingMethod | [optional] [default to "SHA256"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | The name of the key to use in the verification process | 
**Signature** | **string** | The data&#39;s signature in a Base64 format. | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | **int32** | classic key version | 

## Methods

### NewVerifyDataWithClassicKey

`func NewVerifyDataWithClassicKey(data string, name string, signature string, version int32, ) *VerifyDataWithClassicKey`

NewVerifyDataWithClassicKey instantiates a new VerifyDataWithClassicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyDataWithClassicKeyWithDefaults

`func NewVerifyDataWithClassicKeyWithDefaults() *VerifyDataWithClassicKey`

NewVerifyDataWithClassicKeyWithDefaults instantiates a new VerifyDataWithClassicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VerifyDataWithClassicKey) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VerifyDataWithClassicKey) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VerifyDataWithClassicKey) SetData(v string)`

SetData sets Data field to given value.


### GetDisplayId

`func (o *VerifyDataWithClassicKey) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *VerifyDataWithClassicKey) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *VerifyDataWithClassicKey) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *VerifyDataWithClassicKey) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetHashed

`func (o *VerifyDataWithClassicKey) GetHashed() bool`

GetHashed returns the Hashed field if non-nil, zero value otherwise.

### GetHashedOk

`func (o *VerifyDataWithClassicKey) GetHashedOk() (*bool, bool)`

GetHashedOk returns a tuple with the Hashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashed

`func (o *VerifyDataWithClassicKey) SetHashed(v bool)`

SetHashed sets Hashed field to given value.

### HasHashed

`func (o *VerifyDataWithClassicKey) HasHashed() bool`

HasHashed returns a boolean if a field has been set.

### GetHashingMethod

`func (o *VerifyDataWithClassicKey) GetHashingMethod() string`

GetHashingMethod returns the HashingMethod field if non-nil, zero value otherwise.

### GetHashingMethodOk

`func (o *VerifyDataWithClassicKey) GetHashingMethodOk() (*string, bool)`

GetHashingMethodOk returns a tuple with the HashingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashingMethod

`func (o *VerifyDataWithClassicKey) SetHashingMethod(v string)`

SetHashingMethod sets HashingMethod field to given value.

### HasHashingMethod

`func (o *VerifyDataWithClassicKey) HasHashingMethod() bool`

HasHashingMethod returns a boolean if a field has been set.

### GetJson

`func (o *VerifyDataWithClassicKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VerifyDataWithClassicKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VerifyDataWithClassicKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *VerifyDataWithClassicKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *VerifyDataWithClassicKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerifyDataWithClassicKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerifyDataWithClassicKey) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *VerifyDataWithClassicKey) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerifyDataWithClassicKey) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerifyDataWithClassicKey) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetToken

`func (o *VerifyDataWithClassicKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VerifyDataWithClassicKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VerifyDataWithClassicKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VerifyDataWithClassicKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *VerifyDataWithClassicKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *VerifyDataWithClassicKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *VerifyDataWithClassicKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *VerifyDataWithClassicKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *VerifyDataWithClassicKey) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VerifyDataWithClassicKey) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VerifyDataWithClassicKey) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


