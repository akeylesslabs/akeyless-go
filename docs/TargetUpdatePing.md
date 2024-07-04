# TargetUpdatePing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativePort** | Pointer to **string** | Ping Federate administrative port | [optional] [default to "9999"]
**AuthorizationPort** | Pointer to **string** | Ping Federate authorization port | [optional] [default to "9031"]
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Password** | Pointer to **string** | Ping Federate privileged user password | [optional] 
**PingUrl** | Pointer to **string** | Ping URL | [optional] 
**PrivilegedUser** | Pointer to **string** | Ping Federate privileged user | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdatePing

`func NewTargetUpdatePing(name string, ) *TargetUpdatePing`

NewTargetUpdatePing instantiates a new TargetUpdatePing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdatePingWithDefaults

`func NewTargetUpdatePingWithDefaults() *TargetUpdatePing`

NewTargetUpdatePingWithDefaults instantiates a new TargetUpdatePing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativePort

`func (o *TargetUpdatePing) GetAdministrativePort() string`

GetAdministrativePort returns the AdministrativePort field if non-nil, zero value otherwise.

### GetAdministrativePortOk

`func (o *TargetUpdatePing) GetAdministrativePortOk() (*string, bool)`

GetAdministrativePortOk returns a tuple with the AdministrativePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativePort

`func (o *TargetUpdatePing) SetAdministrativePort(v string)`

SetAdministrativePort sets AdministrativePort field to given value.

### HasAdministrativePort

`func (o *TargetUpdatePing) HasAdministrativePort() bool`

HasAdministrativePort returns a boolean if a field has been set.

### GetAuthorizationPort

`func (o *TargetUpdatePing) GetAuthorizationPort() string`

GetAuthorizationPort returns the AuthorizationPort field if non-nil, zero value otherwise.

### GetAuthorizationPortOk

`func (o *TargetUpdatePing) GetAuthorizationPortOk() (*string, bool)`

GetAuthorizationPortOk returns a tuple with the AuthorizationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationPort

`func (o *TargetUpdatePing) SetAuthorizationPort(v string)`

SetAuthorizationPort sets AuthorizationPort field to given value.

### HasAuthorizationPort

`func (o *TargetUpdatePing) HasAuthorizationPort() bool`

HasAuthorizationPort returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdatePing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdatePing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdatePing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdatePing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdatePing) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdatePing) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdatePing) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdatePing) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdatePing) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdatePing) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdatePing) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdatePing) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdatePing) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdatePing) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdatePing) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdatePing) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdatePing) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdatePing) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdatePing) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdatePing) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdatePing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdatePing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdatePing) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdatePing) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdatePing) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdatePing) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdatePing) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *TargetUpdatePing) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetUpdatePing) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetUpdatePing) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TargetUpdatePing) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPingUrl

`func (o *TargetUpdatePing) GetPingUrl() string`

GetPingUrl returns the PingUrl field if non-nil, zero value otherwise.

### GetPingUrlOk

`func (o *TargetUpdatePing) GetPingUrlOk() (*string, bool)`

GetPingUrlOk returns a tuple with the PingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingUrl

`func (o *TargetUpdatePing) SetPingUrl(v string)`

SetPingUrl sets PingUrl field to given value.

### HasPingUrl

`func (o *TargetUpdatePing) HasPingUrl() bool`

HasPingUrl returns a boolean if a field has been set.

### GetPrivilegedUser

`func (o *TargetUpdatePing) GetPrivilegedUser() string`

GetPrivilegedUser returns the PrivilegedUser field if non-nil, zero value otherwise.

### GetPrivilegedUserOk

`func (o *TargetUpdatePing) GetPrivilegedUserOk() (*string, bool)`

GetPrivilegedUserOk returns a tuple with the PrivilegedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedUser

`func (o *TargetUpdatePing) SetPrivilegedUser(v string)`

SetPrivilegedUser sets PrivilegedUser field to given value.

### HasPrivilegedUser

`func (o *TargetUpdatePing) HasPrivilegedUser() bool`

HasPrivilegedUser returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdatePing) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdatePing) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdatePing) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdatePing) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdatePing) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdatePing) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdatePing) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdatePing) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


