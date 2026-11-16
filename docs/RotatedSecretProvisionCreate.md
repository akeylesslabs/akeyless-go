# RotatedSecretProvisionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Rotated secret name | 
**PrivateKeyPassphrase** | Pointer to **string** | Optional passphrase used to encrypt the private key file on the remote host | [optional] 
**PrivateKeyRemotePath** | Pointer to **string** | Remote path for the private key (required when provision-keys is private-key on host targets) | [optional] 
**ProvisionKeys** | Pointer to **string** | Keys to provision: public-key (default) or private-key (also pushes the public key) | [optional] 
**ProvisionOnRotation** | Pointer to **bool** | When true, provision keys to this target after a successful key rotation. Only supported for rotator type key-rotator. Defaults to true. Set false to turn it off. | [optional] 
**PublicKeyRemotePath** | Pointer to **string** | Remote path for the public key (defaults from the rotated secret) | [optional] 
**TargetName** | **string** | Target name to provision keys to (must differ from the rotation target) | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotatedSecretProvisionCreate

`func NewRotatedSecretProvisionCreate(name string, targetName string, ) *RotatedSecretProvisionCreate`

NewRotatedSecretProvisionCreate instantiates a new RotatedSecretProvisionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretProvisionCreateWithDefaults

`func NewRotatedSecretProvisionCreateWithDefaults() *RotatedSecretProvisionCreate`

NewRotatedSecretProvisionCreateWithDefaults instantiates a new RotatedSecretProvisionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *RotatedSecretProvisionCreate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretProvisionCreate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretProvisionCreate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretProvisionCreate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretProvisionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretProvisionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretProvisionCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPrivateKeyPassphrase

`func (o *RotatedSecretProvisionCreate) GetPrivateKeyPassphrase() string`

GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetPrivateKeyPassphraseOk

`func (o *RotatedSecretProvisionCreate) GetPrivateKeyPassphraseOk() (*string, bool)`

GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassphrase

`func (o *RotatedSecretProvisionCreate) SetPrivateKeyPassphrase(v string)`

SetPrivateKeyPassphrase sets PrivateKeyPassphrase field to given value.

### HasPrivateKeyPassphrase

`func (o *RotatedSecretProvisionCreate) HasPrivateKeyPassphrase() bool`

HasPrivateKeyPassphrase returns a boolean if a field has been set.

### GetPrivateKeyRemotePath

`func (o *RotatedSecretProvisionCreate) GetPrivateKeyRemotePath() string`

GetPrivateKeyRemotePath returns the PrivateKeyRemotePath field if non-nil, zero value otherwise.

### GetPrivateKeyRemotePathOk

`func (o *RotatedSecretProvisionCreate) GetPrivateKeyRemotePathOk() (*string, bool)`

GetPrivateKeyRemotePathOk returns a tuple with the PrivateKeyRemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyRemotePath

`func (o *RotatedSecretProvisionCreate) SetPrivateKeyRemotePath(v string)`

SetPrivateKeyRemotePath sets PrivateKeyRemotePath field to given value.

### HasPrivateKeyRemotePath

`func (o *RotatedSecretProvisionCreate) HasPrivateKeyRemotePath() bool`

HasPrivateKeyRemotePath returns a boolean if a field has been set.

### GetProvisionKeys

`func (o *RotatedSecretProvisionCreate) GetProvisionKeys() string`

GetProvisionKeys returns the ProvisionKeys field if non-nil, zero value otherwise.

### GetProvisionKeysOk

`func (o *RotatedSecretProvisionCreate) GetProvisionKeysOk() (*string, bool)`

GetProvisionKeysOk returns a tuple with the ProvisionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionKeys

`func (o *RotatedSecretProvisionCreate) SetProvisionKeys(v string)`

SetProvisionKeys sets ProvisionKeys field to given value.

### HasProvisionKeys

`func (o *RotatedSecretProvisionCreate) HasProvisionKeys() bool`

HasProvisionKeys returns a boolean if a field has been set.

### GetProvisionOnRotation

`func (o *RotatedSecretProvisionCreate) GetProvisionOnRotation() bool`

GetProvisionOnRotation returns the ProvisionOnRotation field if non-nil, zero value otherwise.

### GetProvisionOnRotationOk

`func (o *RotatedSecretProvisionCreate) GetProvisionOnRotationOk() (*bool, bool)`

GetProvisionOnRotationOk returns a tuple with the ProvisionOnRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionOnRotation

`func (o *RotatedSecretProvisionCreate) SetProvisionOnRotation(v bool)`

SetProvisionOnRotation sets ProvisionOnRotation field to given value.

### HasProvisionOnRotation

`func (o *RotatedSecretProvisionCreate) HasProvisionOnRotation() bool`

HasProvisionOnRotation returns a boolean if a field has been set.

### GetPublicKeyRemotePath

`func (o *RotatedSecretProvisionCreate) GetPublicKeyRemotePath() string`

GetPublicKeyRemotePath returns the PublicKeyRemotePath field if non-nil, zero value otherwise.

### GetPublicKeyRemotePathOk

`func (o *RotatedSecretProvisionCreate) GetPublicKeyRemotePathOk() (*string, bool)`

GetPublicKeyRemotePathOk returns a tuple with the PublicKeyRemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyRemotePath

`func (o *RotatedSecretProvisionCreate) SetPublicKeyRemotePath(v string)`

SetPublicKeyRemotePath sets PublicKeyRemotePath field to given value.

### HasPublicKeyRemotePath

`func (o *RotatedSecretProvisionCreate) HasPublicKeyRemotePath() bool`

HasPublicKeyRemotePath returns a boolean if a field has been set.

### GetTargetName

`func (o *RotatedSecretProvisionCreate) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *RotatedSecretProvisionCreate) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *RotatedSecretProvisionCreate) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetToken

`func (o *RotatedSecretProvisionCreate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretProvisionCreate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretProvisionCreate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretProvisionCreate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretProvisionCreate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretProvisionCreate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretProvisionCreate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretProvisionCreate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


