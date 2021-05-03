# CreateManagedKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | Managed Key type; options: [AES256GCM, RSA2048] | 
**ManagedKeyValue** | Pointer to **string** | Base64-encoded managed key value | [optional] 
**Metadata** | Pointer to **string** | Metadata about the managed key | [optional] 
**Name** | **string** | ManagedKey name | 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this managed key | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateManagedKey

`func NewCreateManagedKey(alg string, name string, ) *CreateManagedKey`

NewCreateManagedKey instantiates a new CreateManagedKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManagedKeyWithDefaults

`func NewCreateManagedKeyWithDefaults() *CreateManagedKey`

NewCreateManagedKeyWithDefaults instantiates a new CreateManagedKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *CreateManagedKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *CreateManagedKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *CreateManagedKey) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetManagedKeyValue

`func (o *CreateManagedKey) GetManagedKeyValue() string`

GetManagedKeyValue returns the ManagedKeyValue field if non-nil, zero value otherwise.

### GetManagedKeyValueOk

`func (o *CreateManagedKey) GetManagedKeyValueOk() (*string, bool)`

GetManagedKeyValueOk returns a tuple with the ManagedKeyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyValue

`func (o *CreateManagedKey) SetManagedKeyValue(v string)`

SetManagedKeyValue sets ManagedKeyValue field to given value.

### HasManagedKeyValue

`func (o *CreateManagedKey) HasManagedKeyValue() bool`

HasManagedKeyValue returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateManagedKey) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateManagedKey) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateManagedKey) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateManagedKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateManagedKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateManagedKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateManagedKey) SetName(v string)`

SetName sets Name field to given value.


### GetProtectionKey

`func (o *CreateManagedKey) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateManagedKey) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateManagedKey) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateManagedKey) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTags

`func (o *CreateManagedKey) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateManagedKey) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateManagedKey) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateManagedKey) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *CreateManagedKey) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *CreateManagedKey) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *CreateManagedKey) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *CreateManagedKey) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *CreateManagedKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateManagedKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateManagedKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateManagedKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateManagedKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateManagedKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateManagedKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateManagedKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


