# UpdatePingTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativePort** | Pointer to **string** | Ping Federate administrative port | [optional] [default to "9999"]
**AuthorizationPort** | Pointer to **string** | Ping Federate authorization port | [optional] [default to "9031"]
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**KeepPrevVersion** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Password** | Pointer to **string** | Ping Federate privileged user password | [optional] 
**PingUrl** | Pointer to **string** | Ping URL | [optional] 
**PrivilegedUser** | Pointer to **string** | Ping Federate privileged user | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 

## Methods

### NewUpdatePingTarget

`func NewUpdatePingTarget(name string, ) *UpdatePingTarget`

NewUpdatePingTarget instantiates a new UpdatePingTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePingTargetWithDefaults

`func NewUpdatePingTargetWithDefaults() *UpdatePingTarget`

NewUpdatePingTargetWithDefaults instantiates a new UpdatePingTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativePort

`func (o *UpdatePingTarget) GetAdministrativePort() string`

GetAdministrativePort returns the AdministrativePort field if non-nil, zero value otherwise.

### GetAdministrativePortOk

`func (o *UpdatePingTarget) GetAdministrativePortOk() (*string, bool)`

GetAdministrativePortOk returns a tuple with the AdministrativePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativePort

`func (o *UpdatePingTarget) SetAdministrativePort(v string)`

SetAdministrativePort sets AdministrativePort field to given value.

### HasAdministrativePort

`func (o *UpdatePingTarget) HasAdministrativePort() bool`

HasAdministrativePort returns a boolean if a field has been set.

### GetAuthorizationPort

`func (o *UpdatePingTarget) GetAuthorizationPort() string`

GetAuthorizationPort returns the AuthorizationPort field if non-nil, zero value otherwise.

### GetAuthorizationPortOk

`func (o *UpdatePingTarget) GetAuthorizationPortOk() (*string, bool)`

GetAuthorizationPortOk returns a tuple with the AuthorizationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationPort

`func (o *UpdatePingTarget) SetAuthorizationPort(v string)`

SetAuthorizationPort sets AuthorizationPort field to given value.

### HasAuthorizationPort

`func (o *UpdatePingTarget) HasAuthorizationPort() bool`

HasAuthorizationPort returns a boolean if a field has been set.

### GetComment

`func (o *UpdatePingTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdatePingTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdatePingTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdatePingTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetJson

`func (o *UpdatePingTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdatePingTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdatePingTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdatePingTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdatePingTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdatePingTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdatePingTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdatePingTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdatePingTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdatePingTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdatePingTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdatePingTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdatePingTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePingTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePingTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdatePingTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdatePingTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdatePingTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdatePingTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdatePingTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdatePingTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdatePingTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdatePingTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPingUrl

`func (o *UpdatePingTarget) GetPingUrl() string`

GetPingUrl returns the PingUrl field if non-nil, zero value otherwise.

### GetPingUrlOk

`func (o *UpdatePingTarget) GetPingUrlOk() (*string, bool)`

GetPingUrlOk returns a tuple with the PingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingUrl

`func (o *UpdatePingTarget) SetPingUrl(v string)`

SetPingUrl sets PingUrl field to given value.

### HasPingUrl

`func (o *UpdatePingTarget) HasPingUrl() bool`

HasPingUrl returns a boolean if a field has been set.

### GetPrivilegedUser

`func (o *UpdatePingTarget) GetPrivilegedUser() string`

GetPrivilegedUser returns the PrivilegedUser field if non-nil, zero value otherwise.

### GetPrivilegedUserOk

`func (o *UpdatePingTarget) GetPrivilegedUserOk() (*string, bool)`

GetPrivilegedUserOk returns a tuple with the PrivilegedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedUser

`func (o *UpdatePingTarget) SetPrivilegedUser(v string)`

SetPrivilegedUser sets PrivilegedUser field to given value.

### HasPrivilegedUser

`func (o *UpdatePingTarget) HasPrivilegedUser() bool`

HasPrivilegedUser returns a boolean if a field has been set.

### GetToken

`func (o *UpdatePingTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdatePingTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdatePingTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdatePingTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdatePingTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdatePingTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdatePingTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdatePingTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdatePingTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdatePingTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdatePingTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdatePingTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


