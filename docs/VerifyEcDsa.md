# VerifyEcDsa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the EC key to use for the verification process | [optional] 
**ItemId** | Pointer to **int64** | The item id of the EC key to use for the verification process | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyName** | Pointer to **string** | The name of the EC key to use for the verification process | [optional] 
**Message** | **string** | The message to be verified in a base64 format | 
**Signature** | **string** | The message&#39;s signature | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewVerifyEcDsa

`func NewVerifyEcDsa(message string, signature string, ) *VerifyEcDsa`

NewVerifyEcDsa instantiates a new VerifyEcDsa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyEcDsaWithDefaults

`func NewVerifyEcDsaWithDefaults() *VerifyEcDsa`

NewVerifyEcDsaWithDefaults instantiates a new VerifyEcDsa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *VerifyEcDsa) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *VerifyEcDsa) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *VerifyEcDsa) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *VerifyEcDsa) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetItemId

`func (o *VerifyEcDsa) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *VerifyEcDsa) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *VerifyEcDsa) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *VerifyEcDsa) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *VerifyEcDsa) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VerifyEcDsa) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VerifyEcDsa) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *VerifyEcDsa) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyName

`func (o *VerifyEcDsa) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *VerifyEcDsa) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *VerifyEcDsa) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *VerifyEcDsa) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetMessage

`func (o *VerifyEcDsa) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VerifyEcDsa) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VerifyEcDsa) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSignature

`func (o *VerifyEcDsa) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerifyEcDsa) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerifyEcDsa) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetToken

`func (o *VerifyEcDsa) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VerifyEcDsa) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VerifyEcDsa) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VerifyEcDsa) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *VerifyEcDsa) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *VerifyEcDsa) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *VerifyEcDsa) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *VerifyEcDsa) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


