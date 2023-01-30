# UpdateRotatedSecretSC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation | [optional] 
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Secret name | 
**NewMetadata** | Pointer to **string** | New item metadata | [optional] [default to "default_metadata"]
**NewName** | Pointer to **string** | New item name | [optional] 
**NewVersion** | Pointer to **bool** | Whether to create a new version of not | [optional] [default to false]
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (7-365) | [optional] 
**RotationHour** | Pointer to **int32** |  | [optional] 
**RotatorCredsType** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateRotatedSecretSC

`func NewUpdateRotatedSecretSC(name string, ) *UpdateRotatedSecretSC`

NewUpdateRotatedSecretSC instantiates a new UpdateRotatedSecretSC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRotatedSecretSCWithDefaults

`func NewUpdateRotatedSecretSCWithDefaults() *UpdateRotatedSecretSC`

NewUpdateRotatedSecretSCWithDefaults instantiates a new UpdateRotatedSecretSC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *UpdateRotatedSecretSC) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *UpdateRotatedSecretSC) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *UpdateRotatedSecretSC) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *UpdateRotatedSecretSC) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAutoRotate

`func (o *UpdateRotatedSecretSC) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *UpdateRotatedSecretSC) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *UpdateRotatedSecretSC) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *UpdateRotatedSecretSC) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetGatewayUrl

`func (o *UpdateRotatedSecretSC) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *UpdateRotatedSecretSC) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *UpdateRotatedSecretSC) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *UpdateRotatedSecretSC) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetKey

`func (o *UpdateRotatedSecretSC) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateRotatedSecretSC) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateRotatedSecretSC) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateRotatedSecretSC) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateRotatedSecretSC) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRotatedSecretSC) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRotatedSecretSC) SetName(v string)`

SetName sets Name field to given value.


### GetNewMetadata

`func (o *UpdateRotatedSecretSC) GetNewMetadata() string`

GetNewMetadata returns the NewMetadata field if non-nil, zero value otherwise.

### GetNewMetadataOk

`func (o *UpdateRotatedSecretSC) GetNewMetadataOk() (*string, bool)`

GetNewMetadataOk returns a tuple with the NewMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMetadata

`func (o *UpdateRotatedSecretSC) SetNewMetadata(v string)`

SetNewMetadata sets NewMetadata field to given value.

### HasNewMetadata

`func (o *UpdateRotatedSecretSC) HasNewMetadata() bool`

HasNewMetadata returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateRotatedSecretSC) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateRotatedSecretSC) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateRotatedSecretSC) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateRotatedSecretSC) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetNewVersion

`func (o *UpdateRotatedSecretSC) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *UpdateRotatedSecretSC) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *UpdateRotatedSecretSC) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *UpdateRotatedSecretSC) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetRmTag

`func (o *UpdateRotatedSecretSC) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *UpdateRotatedSecretSC) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *UpdateRotatedSecretSC) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *UpdateRotatedSecretSC) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotationInterval

`func (o *UpdateRotatedSecretSC) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *UpdateRotatedSecretSC) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *UpdateRotatedSecretSC) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *UpdateRotatedSecretSC) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotationHour

`func (o *UpdateRotatedSecretSC) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *UpdateRotatedSecretSC) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *UpdateRotatedSecretSC) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *UpdateRotatedSecretSC) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotatorCredsType

`func (o *UpdateRotatedSecretSC) GetRotatorCredsType() string`

GetRotatorCredsType returns the RotatorCredsType field if non-nil, zero value otherwise.

### GetRotatorCredsTypeOk

`func (o *UpdateRotatedSecretSC) GetRotatorCredsTypeOk() (*string, bool)`

GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorCredsType

`func (o *UpdateRotatedSecretSC) SetRotatorCredsType(v string)`

SetRotatorCredsType sets RotatorCredsType field to given value.

### HasRotatorCredsType

`func (o *UpdateRotatedSecretSC) HasRotatorCredsType() bool`

HasRotatorCredsType returns a boolean if a field has been set.

### GetSshPassword

`func (o *UpdateRotatedSecretSC) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateRotatedSecretSC) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateRotatedSecretSC) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateRotatedSecretSC) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshUsername

`func (o *UpdateRotatedSecretSC) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateRotatedSecretSC) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateRotatedSecretSC) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateRotatedSecretSC) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetToken

`func (o *UpdateRotatedSecretSC) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateRotatedSecretSC) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateRotatedSecretSC) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateRotatedSecretSC) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateRotatedSecretSC) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateRotatedSecretSC) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateRotatedSecretSC) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateRotatedSecretSC) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


