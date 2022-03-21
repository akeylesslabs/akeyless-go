# UpdateSSHTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**KeepPrevVersion** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**PrivateKeyPassword** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 

## Methods

### NewUpdateSSHTarget

`func NewUpdateSSHTarget(name string, ) *UpdateSSHTarget`

NewUpdateSSHTarget instantiates a new UpdateSSHTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSSHTargetWithDefaults

`func NewUpdateSSHTargetWithDefaults() *UpdateSSHTarget`

NewUpdateSSHTargetWithDefaults instantiates a new UpdateSSHTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateSSHTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateSSHTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateSSHTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateSSHTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetHost

`func (o *UpdateSSHTarget) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateSSHTarget) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateSSHTarget) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateSSHTarget) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateSSHTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateSSHTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateSSHTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateSSHTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateSSHTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateSSHTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateSSHTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateSSHTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateSSHTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSSHTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSSHTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateSSHTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateSSHTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateSSHTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateSSHTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPort

`func (o *UpdateSSHTarget) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateSSHTarget) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateSSHTarget) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateSSHTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateKey

`func (o *UpdateSSHTarget) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *UpdateSSHTarget) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *UpdateSSHTarget) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *UpdateSSHTarget) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyPassword

`func (o *UpdateSSHTarget) GetPrivateKeyPassword() string`

GetPrivateKeyPassword returns the PrivateKeyPassword field if non-nil, zero value otherwise.

### GetPrivateKeyPasswordOk

`func (o *UpdateSSHTarget) GetPrivateKeyPasswordOk() (*string, bool)`

GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassword

`func (o *UpdateSSHTarget) SetPrivateKeyPassword(v string)`

SetPrivateKeyPassword sets PrivateKeyPassword field to given value.

### HasPrivateKeyPassword

`func (o *UpdateSSHTarget) HasPrivateKeyPassword() bool`

HasPrivateKeyPassword returns a boolean if a field has been set.

### GetSshPassword

`func (o *UpdateSSHTarget) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateSSHTarget) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateSSHTarget) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateSSHTarget) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshUsername

`func (o *UpdateSSHTarget) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateSSHTarget) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateSSHTarget) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateSSHTarget) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetToken

`func (o *UpdateSSHTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateSSHTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateSSHTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateSSHTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateSSHTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateSSHTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateSSHTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateSSHTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateSSHTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateSSHTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateSSHTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateSSHTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


