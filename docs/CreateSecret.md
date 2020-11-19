# CreateSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **string** | Metadata about the secret | [optional] 
**MultilineValue** | Pointer to **bool** | The provided value is a multiline value (separated by &#39;\\n&#39;) | [optional] 
**Name** | **string** | Secret name | 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Value** | **string** | The secret value | 

## Methods

### NewCreateSecret

`func NewCreateSecret(name string, value string, ) *CreateSecret`

NewCreateSecret instantiates a new CreateSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecretWithDefaults

`func NewCreateSecretWithDefaults() *CreateSecret`

NewCreateSecretWithDefaults instantiates a new CreateSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CreateSecret) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSecret) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSecret) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSecret) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMultilineValue

`func (o *CreateSecret) GetMultilineValue() bool`

GetMultilineValue returns the MultilineValue field if non-nil, zero value otherwise.

### GetMultilineValueOk

`func (o *CreateSecret) GetMultilineValueOk() (*bool, bool)`

GetMultilineValueOk returns a tuple with the MultilineValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultilineValue

`func (o *CreateSecret) SetMultilineValue(v bool)`

SetMultilineValue sets MultilineValue field to given value.

### HasMultilineValue

`func (o *CreateSecret) HasMultilineValue() bool`

HasMultilineValue returns a boolean if a field has been set.

### GetName

`func (o *CreateSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSecret) SetName(v string)`

SetName sets Name field to given value.


### GetProtectionKey

`func (o *CreateSecret) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateSecret) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateSecret) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateSecret) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTags

`func (o *CreateSecret) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSecret) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSecret) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateSecret) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *CreateSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateSecret) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateSecret) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateSecret) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateSecret) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetValue

`func (o *CreateSecret) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateSecret) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateSecret) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


