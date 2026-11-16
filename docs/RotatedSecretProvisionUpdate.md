# RotatedSecretProvisionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Rotated secret name | 
**NewTargetName** | Pointer to **string** | New target name when retargeting the provisioning association | [optional] 
**PrivateKeyPassphrase** | Pointer to **string** | Passphrase used to encrypt the private key file (omit to keep current; requires Read) | [optional] 
**PrivateKeyRemotePath** | Pointer to **string** | Remote path for the private key (omit to keep current) | [optional] 
**ProvisionKeys** | Pointer to **string** | Keys to provision: public-key or private-key (omit to keep current) | [optional] 
**ProvisionOnRotation** | Pointer to **bool** | When true, provision keys to this target after a successful key rotation. Only supported for rotator type key-rotator. Omit to keep the current value. Set false to turn it off. | [optional] 
**PublicKeyRemotePath** | Pointer to **string** | Remote path for the public key (omit to keep current) | [optional] 
**TargetName** | **string** | Target name of the provisioning association to update | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotatedSecretProvisionUpdate

`func NewRotatedSecretProvisionUpdate(name string, targetName string, ) *RotatedSecretProvisionUpdate`

NewRotatedSecretProvisionUpdate instantiates a new RotatedSecretProvisionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretProvisionUpdateWithDefaults

`func NewRotatedSecretProvisionUpdateWithDefaults() *RotatedSecretProvisionUpdate`

NewRotatedSecretProvisionUpdateWithDefaults instantiates a new RotatedSecretProvisionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *RotatedSecretProvisionUpdate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretProvisionUpdate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretProvisionUpdate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretProvisionUpdate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretProvisionUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretProvisionUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretProvisionUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetNewTargetName

`func (o *RotatedSecretProvisionUpdate) GetNewTargetName() string`

GetNewTargetName returns the NewTargetName field if non-nil, zero value otherwise.

### GetNewTargetNameOk

`func (o *RotatedSecretProvisionUpdate) GetNewTargetNameOk() (*string, bool)`

GetNewTargetNameOk returns a tuple with the NewTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTargetName

`func (o *RotatedSecretProvisionUpdate) SetNewTargetName(v string)`

SetNewTargetName sets NewTargetName field to given value.

### HasNewTargetName

`func (o *RotatedSecretProvisionUpdate) HasNewTargetName() bool`

HasNewTargetName returns a boolean if a field has been set.

### GetPrivateKeyPassphrase

`func (o *RotatedSecretProvisionUpdate) GetPrivateKeyPassphrase() string`

GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetPrivateKeyPassphraseOk

`func (o *RotatedSecretProvisionUpdate) GetPrivateKeyPassphraseOk() (*string, bool)`

GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassphrase

`func (o *RotatedSecretProvisionUpdate) SetPrivateKeyPassphrase(v string)`

SetPrivateKeyPassphrase sets PrivateKeyPassphrase field to given value.

### HasPrivateKeyPassphrase

`func (o *RotatedSecretProvisionUpdate) HasPrivateKeyPassphrase() bool`

HasPrivateKeyPassphrase returns a boolean if a field has been set.

### GetPrivateKeyRemotePath

`func (o *RotatedSecretProvisionUpdate) GetPrivateKeyRemotePath() string`

GetPrivateKeyRemotePath returns the PrivateKeyRemotePath field if non-nil, zero value otherwise.

### GetPrivateKeyRemotePathOk

`func (o *RotatedSecretProvisionUpdate) GetPrivateKeyRemotePathOk() (*string, bool)`

GetPrivateKeyRemotePathOk returns a tuple with the PrivateKeyRemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyRemotePath

`func (o *RotatedSecretProvisionUpdate) SetPrivateKeyRemotePath(v string)`

SetPrivateKeyRemotePath sets PrivateKeyRemotePath field to given value.

### HasPrivateKeyRemotePath

`func (o *RotatedSecretProvisionUpdate) HasPrivateKeyRemotePath() bool`

HasPrivateKeyRemotePath returns a boolean if a field has been set.

### GetProvisionKeys

`func (o *RotatedSecretProvisionUpdate) GetProvisionKeys() string`

GetProvisionKeys returns the ProvisionKeys field if non-nil, zero value otherwise.

### GetProvisionKeysOk

`func (o *RotatedSecretProvisionUpdate) GetProvisionKeysOk() (*string, bool)`

GetProvisionKeysOk returns a tuple with the ProvisionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionKeys

`func (o *RotatedSecretProvisionUpdate) SetProvisionKeys(v string)`

SetProvisionKeys sets ProvisionKeys field to given value.

### HasProvisionKeys

`func (o *RotatedSecretProvisionUpdate) HasProvisionKeys() bool`

HasProvisionKeys returns a boolean if a field has been set.

### GetProvisionOnRotation

`func (o *RotatedSecretProvisionUpdate) GetProvisionOnRotation() bool`

GetProvisionOnRotation returns the ProvisionOnRotation field if non-nil, zero value otherwise.

### GetProvisionOnRotationOk

`func (o *RotatedSecretProvisionUpdate) GetProvisionOnRotationOk() (*bool, bool)`

GetProvisionOnRotationOk returns a tuple with the ProvisionOnRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionOnRotation

`func (o *RotatedSecretProvisionUpdate) SetProvisionOnRotation(v bool)`

SetProvisionOnRotation sets ProvisionOnRotation field to given value.

### HasProvisionOnRotation

`func (o *RotatedSecretProvisionUpdate) HasProvisionOnRotation() bool`

HasProvisionOnRotation returns a boolean if a field has been set.

### GetPublicKeyRemotePath

`func (o *RotatedSecretProvisionUpdate) GetPublicKeyRemotePath() string`

GetPublicKeyRemotePath returns the PublicKeyRemotePath field if non-nil, zero value otherwise.

### GetPublicKeyRemotePathOk

`func (o *RotatedSecretProvisionUpdate) GetPublicKeyRemotePathOk() (*string, bool)`

GetPublicKeyRemotePathOk returns a tuple with the PublicKeyRemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyRemotePath

`func (o *RotatedSecretProvisionUpdate) SetPublicKeyRemotePath(v string)`

SetPublicKeyRemotePath sets PublicKeyRemotePath field to given value.

### HasPublicKeyRemotePath

`func (o *RotatedSecretProvisionUpdate) HasPublicKeyRemotePath() bool`

HasPublicKeyRemotePath returns a boolean if a field has been set.

### GetTargetName

`func (o *RotatedSecretProvisionUpdate) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *RotatedSecretProvisionUpdate) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *RotatedSecretProvisionUpdate) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetToken

`func (o *RotatedSecretProvisionUpdate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretProvisionUpdate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretProvisionUpdate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretProvisionUpdate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretProvisionUpdate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretProvisionUpdate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretProvisionUpdate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretProvisionUpdate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


