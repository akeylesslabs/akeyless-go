# VerifyJWTWithClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | **string** | The name of the key to use in the verify JWT process | 
**JwtClaims** | **string** | JWTClaims | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Signature** | **string** | Signature | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Version** | **int32** | classic key version | 

## Methods

### NewVerifyJWTWithClassicKey

`func NewVerifyJWTWithClassicKey(displayId string, jwtClaims string, signature string, version int32, ) *VerifyJWTWithClassicKey`

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


### GetJwtClaims

`func (o *VerifyJWTWithClassicKey) GetJwtClaims() string`

GetJwtClaims returns the JwtClaims field if non-nil, zero value otherwise.

### GetJwtClaimsOk

`func (o *VerifyJWTWithClassicKey) GetJwtClaimsOk() (*string, bool)`

GetJwtClaimsOk returns a tuple with the JwtClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtClaims

`func (o *VerifyJWTWithClassicKey) SetJwtClaims(v string)`

SetJwtClaims sets JwtClaims field to given value.


### GetPassword

`func (o *VerifyJWTWithClassicKey) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VerifyJWTWithClassicKey) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VerifyJWTWithClassicKey) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VerifyJWTWithClassicKey) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSignature

`func (o *VerifyJWTWithClassicKey) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerifyJWTWithClassicKey) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerifyJWTWithClassicKey) SetSignature(v string)`

SetSignature sets Signature field to given value.


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

### GetUsername

`func (o *VerifyJWTWithClassicKey) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VerifyJWTWithClassicKey) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VerifyJWTWithClassicKey) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VerifyJWTWithClassicKey) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

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


