# CreatePasskey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**Alg** | **string** | Passkey type; options: [EC256, EC384, EC512] | [default to "EC256"]
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | ClassicKey name | 
**OriginUrl** | Pointer to **[]string** | Originating websites for this passkey | [optional] 
**ProtectionKeyName** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | For Password Management use | [optional] 

## Methods

### NewCreatePasskey

`func NewCreatePasskey(alg string, name string, ) *CreatePasskey`

NewCreatePasskey instantiates a new CreatePasskey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePasskeyWithDefaults

`func NewCreatePasskeyWithDefaults() *CreatePasskey`

NewCreatePasskeyWithDefaults instantiates a new CreatePasskey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *CreatePasskey) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *CreatePasskey) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *CreatePasskey) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *CreatePasskey) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetAlg

`func (o *CreatePasskey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *CreatePasskey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *CreatePasskey) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetDeleteProtection

`func (o *CreatePasskey) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreatePasskey) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreatePasskey) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreatePasskey) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePasskey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePasskey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePasskey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePasskey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *CreatePasskey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreatePasskey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreatePasskey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreatePasskey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *CreatePasskey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePasskey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePasskey) SetName(v string)`

SetName sets Name field to given value.


### GetOriginUrl

`func (o *CreatePasskey) GetOriginUrl() []string`

GetOriginUrl returns the OriginUrl field if non-nil, zero value otherwise.

### GetOriginUrlOk

`func (o *CreatePasskey) GetOriginUrlOk() (*[]string, bool)`

GetOriginUrlOk returns a tuple with the OriginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginUrl

`func (o *CreatePasskey) SetOriginUrl(v []string)`

SetOriginUrl sets OriginUrl field to given value.

### HasOriginUrl

`func (o *CreatePasskey) HasOriginUrl() bool`

HasOriginUrl returns a boolean if a field has been set.

### GetProtectionKeyName

`func (o *CreatePasskey) GetProtectionKeyName() string`

GetProtectionKeyName returns the ProtectionKeyName field if non-nil, zero value otherwise.

### GetProtectionKeyNameOk

`func (o *CreatePasskey) GetProtectionKeyNameOk() (*string, bool)`

GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyName

`func (o *CreatePasskey) SetProtectionKeyName(v string)`

SetProtectionKeyName sets ProtectionKeyName field to given value.

### HasProtectionKeyName

`func (o *CreatePasskey) HasProtectionKeyName() bool`

HasProtectionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *CreatePasskey) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreatePasskey) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreatePasskey) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreatePasskey) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *CreatePasskey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreatePasskey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreatePasskey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreatePasskey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreatePasskey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreatePasskey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreatePasskey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreatePasskey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreatePasskey) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreatePasskey) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreatePasskey) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreatePasskey) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


