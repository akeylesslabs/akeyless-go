# CreateSSHTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Host** | Pointer to **string** | SSH host name | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Port** | Pointer to **string** | SSH port | [optional] [default to "22"]
**PrivateKey** | Pointer to **string** | SSH private key | [optional] 
**PrivateKeyPassword** | Pointer to **string** | SSH private key password | [optional] 
**SshPassword** | Pointer to **string** | SSH password to rotate | [optional] 
**SshUsername** | Pointer to **string** | SSH username | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateSSHTarget

`func NewCreateSSHTarget(name string, ) *CreateSSHTarget`

NewCreateSSHTarget instantiates a new CreateSSHTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSSHTargetWithDefaults

`func NewCreateSSHTargetWithDefaults() *CreateSSHTarget`

NewCreateSSHTargetWithDefaults instantiates a new CreateSSHTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateSSHTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateSSHTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateSSHTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateSSHTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSSHTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSSHTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSSHTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSSHTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *CreateSSHTarget) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateSSHTarget) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateSSHTarget) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateSSHTarget) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *CreateSSHTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateSSHTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateSSHTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateSSHTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *CreateSSHTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateSSHTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateSSHTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateSSHTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateSSHTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSSHTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSSHTarget) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *CreateSSHTarget) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateSSHTarget) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateSSHTarget) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateSSHTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CreateSSHTarget) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateSSHTarget) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateSSHTarget) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CreateSSHTarget) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyPassword

`func (o *CreateSSHTarget) GetPrivateKeyPassword() string`

GetPrivateKeyPassword returns the PrivateKeyPassword field if non-nil, zero value otherwise.

### GetPrivateKeyPasswordOk

`func (o *CreateSSHTarget) GetPrivateKeyPasswordOk() (*string, bool)`

GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassword

`func (o *CreateSSHTarget) SetPrivateKeyPassword(v string)`

SetPrivateKeyPassword sets PrivateKeyPassword field to given value.

### HasPrivateKeyPassword

`func (o *CreateSSHTarget) HasPrivateKeyPassword() bool`

HasPrivateKeyPassword returns a boolean if a field has been set.

### GetSshPassword

`func (o *CreateSSHTarget) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *CreateSSHTarget) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *CreateSSHTarget) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *CreateSSHTarget) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshUsername

`func (o *CreateSSHTarget) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *CreateSSHTarget) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *CreateSSHTarget) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *CreateSSHTarget) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetToken

`func (o *CreateSSHTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateSSHTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateSSHTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateSSHTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateSSHTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateSSHTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateSSHTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateSSHTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


