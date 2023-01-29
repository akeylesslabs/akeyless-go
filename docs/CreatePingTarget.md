# CreatePingTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativePort** | Pointer to **string** | Ping Federate administrative port | [optional] [default to "9999"]
**AuthorizationPort** | Pointer to **string** | Ping Federate authorization port | [optional] [default to "9031"]
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Ping Federate privileged user password | [optional] 
**PingUrl** | Pointer to **string** | Ping URL | [optional] 
**PrivilegedUser** | Pointer to **string** | Ping Federate privileged user | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreatePingTarget

`func NewCreatePingTarget(name string, ) *CreatePingTarget`

NewCreatePingTarget instantiates a new CreatePingTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePingTargetWithDefaults

`func NewCreatePingTargetWithDefaults() *CreatePingTarget`

NewCreatePingTargetWithDefaults instantiates a new CreatePingTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativePort

`func (o *CreatePingTarget) GetAdministrativePort() string`

GetAdministrativePort returns the AdministrativePort field if non-nil, zero value otherwise.

### GetAdministrativePortOk

`func (o *CreatePingTarget) GetAdministrativePortOk() (*string, bool)`

GetAdministrativePortOk returns a tuple with the AdministrativePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativePort

`func (o *CreatePingTarget) SetAdministrativePort(v string)`

SetAdministrativePort sets AdministrativePort field to given value.

### HasAdministrativePort

`func (o *CreatePingTarget) HasAdministrativePort() bool`

HasAdministrativePort returns a boolean if a field has been set.

### GetAuthorizationPort

`func (o *CreatePingTarget) GetAuthorizationPort() string`

GetAuthorizationPort returns the AuthorizationPort field if non-nil, zero value otherwise.

### GetAuthorizationPortOk

`func (o *CreatePingTarget) GetAuthorizationPortOk() (*string, bool)`

GetAuthorizationPortOk returns a tuple with the AuthorizationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationPort

`func (o *CreatePingTarget) SetAuthorizationPort(v string)`

SetAuthorizationPort sets AuthorizationPort field to given value.

### HasAuthorizationPort

`func (o *CreatePingTarget) HasAuthorizationPort() bool`

HasAuthorizationPort returns a boolean if a field has been set.

### GetComment

`func (o *CreatePingTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreatePingTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreatePingTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreatePingTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePingTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePingTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePingTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePingTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *CreatePingTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreatePingTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreatePingTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreatePingTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *CreatePingTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreatePingTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreatePingTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreatePingTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreatePingTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePingTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePingTarget) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreatePingTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreatePingTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreatePingTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreatePingTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPingUrl

`func (o *CreatePingTarget) GetPingUrl() string`

GetPingUrl returns the PingUrl field if non-nil, zero value otherwise.

### GetPingUrlOk

`func (o *CreatePingTarget) GetPingUrlOk() (*string, bool)`

GetPingUrlOk returns a tuple with the PingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingUrl

`func (o *CreatePingTarget) SetPingUrl(v string)`

SetPingUrl sets PingUrl field to given value.

### HasPingUrl

`func (o *CreatePingTarget) HasPingUrl() bool`

HasPingUrl returns a boolean if a field has been set.

### GetPrivilegedUser

`func (o *CreatePingTarget) GetPrivilegedUser() string`

GetPrivilegedUser returns the PrivilegedUser field if non-nil, zero value otherwise.

### GetPrivilegedUserOk

`func (o *CreatePingTarget) GetPrivilegedUserOk() (*string, bool)`

GetPrivilegedUserOk returns a tuple with the PrivilegedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedUser

`func (o *CreatePingTarget) SetPrivilegedUser(v string)`

SetPrivilegedUser sets PrivilegedUser field to given value.

### HasPrivilegedUser

`func (o *CreatePingTarget) HasPrivilegedUser() bool`

HasPrivilegedUser returns a boolean if a field has been set.

### GetToken

`func (o *CreatePingTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreatePingTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreatePingTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreatePingTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreatePingTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreatePingTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreatePingTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreatePingTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


