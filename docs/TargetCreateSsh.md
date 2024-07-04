# TargetCreateSsh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**Host** | Pointer to **string** | SSH host name | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Port** | Pointer to **string** | SSH port | [optional] [default to "22"]
**PrivateKey** | Pointer to **string** | SSH private key | [optional] 
**PrivateKeyPassword** | Pointer to **string** | SSH private key password | [optional] 
**SshPassword** | Pointer to **string** | SSH password to rotate | [optional] 
**SshUsername** | Pointer to **string** | SSH username | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateSsh

`func NewTargetCreateSsh(name string, ) *TargetCreateSsh`

NewTargetCreateSsh instantiates a new TargetCreateSsh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateSshWithDefaults

`func NewTargetCreateSshWithDefaults() *TargetCreateSsh`

NewTargetCreateSshWithDefaults instantiates a new TargetCreateSsh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TargetCreateSsh) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateSsh) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateSsh) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateSsh) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *TargetCreateSsh) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetCreateSsh) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetCreateSsh) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TargetCreateSsh) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateSsh) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateSsh) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateSsh) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateSsh) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateSsh) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateSsh) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateSsh) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateSsh) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateSsh) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateSsh) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateSsh) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateSsh) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateSsh) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateSsh) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateSsh) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *TargetCreateSsh) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetCreateSsh) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetCreateSsh) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetCreateSsh) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateKey

`func (o *TargetCreateSsh) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *TargetCreateSsh) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *TargetCreateSsh) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *TargetCreateSsh) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyPassword

`func (o *TargetCreateSsh) GetPrivateKeyPassword() string`

GetPrivateKeyPassword returns the PrivateKeyPassword field if non-nil, zero value otherwise.

### GetPrivateKeyPasswordOk

`func (o *TargetCreateSsh) GetPrivateKeyPasswordOk() (*string, bool)`

GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassword

`func (o *TargetCreateSsh) SetPrivateKeyPassword(v string)`

SetPrivateKeyPassword sets PrivateKeyPassword field to given value.

### HasPrivateKeyPassword

`func (o *TargetCreateSsh) HasPrivateKeyPassword() bool`

HasPrivateKeyPassword returns a boolean if a field has been set.

### GetSshPassword

`func (o *TargetCreateSsh) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *TargetCreateSsh) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *TargetCreateSsh) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *TargetCreateSsh) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshUsername

`func (o *TargetCreateSsh) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *TargetCreateSsh) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *TargetCreateSsh) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *TargetCreateSsh) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateSsh) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateSsh) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateSsh) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateSsh) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateSsh) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateSsh) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateSsh) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateSsh) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


