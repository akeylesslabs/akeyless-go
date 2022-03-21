# VerifyJWTWithClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | **string** | The name of the key to use in the verify JWT process | 
**Jwt** | **string** | JWT | 
**RequiredClaims** | **string** | RequiredClaims | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | **int32** | classic key version | 

## Methods

### NewVerifyJWTWithClassicKey

`func NewVerifyJWTWithClassicKey(displayId string, jwt string, requiredClaims string, version int32, ) *VerifyJWTWithClassicKey`

NewVerifyJWTWithClassicKey instantiates a new VerifyJWTWithClassicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyJWTWithClassicKeyWithDefaults

`func NewVerifyJWTWithClassicKeyWithDefaults() *VerifyJWTWithClassicKey`

NewVerifyJWTWithClassicKeyWithDefaults instantiates a new VerifyJWTWithClassicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *VerifyJWTWithClassicKey) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *VerifyJWTWithClassicKey) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *VerifyJWTWithClassicKey) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetJwt

`func (o *VerifyJWTWithClassicKey) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *VerifyJWTWithClassicKey) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *VerifyJWTWithClassicKey) SetJwt(v string)`

SetJwt sets Jwt field to given value.


### GetRequiredClaims

`func (o *VerifyJWTWithClassicKey) GetRequiredClaims() string`

GetRequiredClaims returns the RequiredClaims field if non-nil, zero value otherwise.

### GetRequiredClaimsOk

`func (o *VerifyJWTWithClassicKey) GetRequiredClaimsOk() (*string, bool)`

GetRequiredClaimsOk returns a tuple with the RequiredClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredClaims

`func (o *VerifyJWTWithClassicKey) SetRequiredClaims(v string)`

SetRequiredClaims sets RequiredClaims field to given value.


### GetToken

`func (o *VerifyJWTWithClassicKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VerifyJWTWithClassicKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VerifyJWTWithClassicKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VerifyJWTWithClassicKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *VerifyJWTWithClassicKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *VerifyJWTWithClassicKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *VerifyJWTWithClassicKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *VerifyJWTWithClassicKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *VerifyJWTWithClassicKey) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VerifyJWTWithClassicKey) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VerifyJWTWithClassicKey) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


