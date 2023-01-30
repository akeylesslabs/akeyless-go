# VerifyPKICertWithClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | **string** | The name of the key to use in the verify PKICert process | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**PkiCert** | **string** | PkiCert | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | **int32** | classic key version | 

## Methods

### NewVerifyPKICertWithClassicKey

`func NewVerifyPKICertWithClassicKey(displayId string, pkiCert string, version int32, ) *VerifyPKICertWithClassicKey`

NewVerifyPKICertWithClassicKey instantiates a new VerifyPKICertWithClassicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyPKICertWithClassicKeyWithDefaults

`func NewVerifyPKICertWithClassicKeyWithDefaults() *VerifyPKICertWithClassicKey`

NewVerifyPKICertWithClassicKeyWithDefaults instantiates a new VerifyPKICertWithClassicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *VerifyPKICertWithClassicKey) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *VerifyPKICertWithClassicKey) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *VerifyPKICertWithClassicKey) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetJson

`func (o *VerifyPKICertWithClassicKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VerifyPKICertWithClassicKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VerifyPKICertWithClassicKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *VerifyPKICertWithClassicKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetPkiCert

`func (o *VerifyPKICertWithClassicKey) GetPkiCert() string`

GetPkiCert returns the PkiCert field if non-nil, zero value otherwise.

### GetPkiCertOk

`func (o *VerifyPKICertWithClassicKey) GetPkiCertOk() (*string, bool)`

GetPkiCertOk returns a tuple with the PkiCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCert

`func (o *VerifyPKICertWithClassicKey) SetPkiCert(v string)`

SetPkiCert sets PkiCert field to given value.


### GetToken

`func (o *VerifyPKICertWithClassicKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VerifyPKICertWithClassicKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VerifyPKICertWithClassicKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VerifyPKICertWithClassicKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *VerifyPKICertWithClassicKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *VerifyPKICertWithClassicKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *VerifyPKICertWithClassicKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *VerifyPKICertWithClassicKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *VerifyPKICertWithClassicKey) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VerifyPKICertWithClassicKey) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VerifyPKICertWithClassicKey) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


