# CreateRotatedSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation | [optional] 
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**Metadata** | Pointer to **string** | Metadata about the secret | [optional] 
**Name** | **string** | Secret name | 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (7-365) | [optional] 
**RotationHour** | Pointer to **int32** |  | [optional] 
**RotatorCredsType** | Pointer to **string** |  | [optional] 
**RotatorType** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateRotatedSecret

`func NewCreateRotatedSecret(name string, ) *CreateRotatedSecret`

NewCreateRotatedSecret instantiates a new CreateRotatedSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRotatedSecretWithDefaults

`func NewCreateRotatedSecretWithDefaults() *CreateRotatedSecret`

NewCreateRotatedSecretWithDefaults instantiates a new CreateRotatedSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRotate

`func (o *CreateRotatedSecret) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *CreateRotatedSecret) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *CreateRotatedSecret) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *CreateRotatedSecret) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetGatewayUrl

`func (o *CreateRotatedSecret) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *CreateRotatedSecret) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *CreateRotatedSecret) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *CreateRotatedSecret) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateRotatedSecret) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateRotatedSecret) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateRotatedSecret) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateRotatedSecret) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateRotatedSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRotatedSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRotatedSecret) SetName(v string)`

SetName sets Name field to given value.


### GetProtectionKey

`func (o *CreateRotatedSecret) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateRotatedSecret) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateRotatedSecret) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateRotatedSecret) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetRotationInterval

`func (o *CreateRotatedSecret) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *CreateRotatedSecret) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *CreateRotatedSecret) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *CreateRotatedSecret) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotationHour

`func (o *CreateRotatedSecret) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *CreateRotatedSecret) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *CreateRotatedSecret) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *CreateRotatedSecret) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotatorCredsType

`func (o *CreateRotatedSecret) GetRotatorCredsType() string`

GetRotatorCredsType returns the RotatorCredsType field if non-nil, zero value otherwise.

### GetRotatorCredsTypeOk

`func (o *CreateRotatedSecret) GetRotatorCredsTypeOk() (*string, bool)`

GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorCredsType

`func (o *CreateRotatedSecret) SetRotatorCredsType(v string)`

SetRotatorCredsType sets RotatorCredsType field to given value.

### HasRotatorCredsType

`func (o *CreateRotatedSecret) HasRotatorCredsType() bool`

HasRotatorCredsType returns a boolean if a field has been set.

### GetRotatorType

`func (o *CreateRotatedSecret) GetRotatorType() string`

GetRotatorType returns the RotatorType field if non-nil, zero value otherwise.

### GetRotatorTypeOk

`func (o *CreateRotatedSecret) GetRotatorTypeOk() (*string, bool)`

GetRotatorTypeOk returns a tuple with the RotatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorType

`func (o *CreateRotatedSecret) SetRotatorType(v string)`

SetRotatorType sets RotatorType field to given value.

### HasRotatorType

`func (o *CreateRotatedSecret) HasRotatorType() bool`

HasRotatorType returns a boolean if a field has been set.

### GetSshPassword

`func (o *CreateRotatedSecret) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *CreateRotatedSecret) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *CreateRotatedSecret) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *CreateRotatedSecret) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshUsername

`func (o *CreateRotatedSecret) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *CreateRotatedSecret) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *CreateRotatedSecret) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *CreateRotatedSecret) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetTags

`func (o *CreateRotatedSecret) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRotatedSecret) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRotatedSecret) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRotatedSecret) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *CreateRotatedSecret) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *CreateRotatedSecret) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *CreateRotatedSecret) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *CreateRotatedSecret) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *CreateRotatedSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateRotatedSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateRotatedSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateRotatedSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateRotatedSecret) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateRotatedSecret) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateRotatedSecret) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateRotatedSecret) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


