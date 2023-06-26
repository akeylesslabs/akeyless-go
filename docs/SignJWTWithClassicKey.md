# SignJWTWithClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | **string** | The name of the key to use in the sign JWT process | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtClaims** | **string** | JWTClaims | 
**SigningMethod** | **string** | SigningMethod | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | **int32** | classic key version | 

## Methods

### NewSignJWTWithClassicKey

`func NewSignJWTWithClassicKey(displayId string, jwtClaims string, signingMethod string, version int32, ) *SignJWTWithClassicKey`

NewSignJWTWithClassicKey instantiates a new SignJWTWithClassicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignJWTWithClassicKeyWithDefaults

`func NewSignJWTWithClassicKeyWithDefaults() *SignJWTWithClassicKey`

NewSignJWTWithClassicKeyWithDefaults instantiates a new SignJWTWithClassicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *SignJWTWithClassicKey) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *SignJWTWithClassicKey) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *SignJWTWithClassicKey) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetJson

`func (o *SignJWTWithClassicKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *SignJWTWithClassicKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *SignJWTWithClassicKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *SignJWTWithClassicKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtClaims

`func (o *SignJWTWithClassicKey) GetJwtClaims() string`

GetJwtClaims returns the JwtClaims field if non-nil, zero value otherwise.

### GetJwtClaimsOk

`func (o *SignJWTWithClassicKey) GetJwtClaimsOk() (*string, bool)`

GetJwtClaimsOk returns a tuple with the JwtClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtClaims

`func (o *SignJWTWithClassicKey) SetJwtClaims(v string)`

SetJwtClaims sets JwtClaims field to given value.


### GetSigningMethod

`func (o *SignJWTWithClassicKey) GetSigningMethod() string`

GetSigningMethod returns the SigningMethod field if non-nil, zero value otherwise.

### GetSigningMethodOk

`func (o *SignJWTWithClassicKey) GetSigningMethodOk() (*string, bool)`

GetSigningMethodOk returns a tuple with the SigningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningMethod

`func (o *SignJWTWithClassicKey) SetSigningMethod(v string)`

SetSigningMethod sets SigningMethod field to given value.


### GetToken

`func (o *SignJWTWithClassicKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SignJWTWithClassicKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SignJWTWithClassicKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SignJWTWithClassicKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *SignJWTWithClassicKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *SignJWTWithClassicKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *SignJWTWithClassicKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *SignJWTWithClassicKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *SignJWTWithClassicKey) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SignJWTWithClassicKey) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SignJWTWithClassicKey) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


