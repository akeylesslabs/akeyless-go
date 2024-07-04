# TargetCreatePing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativePort** | Pointer to **string** | Ping Federate administrative port | [optional] [default to "9999"]
**AuthorizationPort** | Pointer to **string** | Ping Federate authorization port | [optional] [default to "9031"]
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Ping Federate privileged user password | [optional] 
**PingUrl** | Pointer to **string** | Ping URL | [optional] 
**PrivilegedUser** | Pointer to **string** | Ping Federate privileged user | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreatePing

`func NewTargetCreatePing(name string, ) *TargetCreatePing`

NewTargetCreatePing instantiates a new TargetCreatePing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreatePingWithDefaults

`func NewTargetCreatePingWithDefaults() *TargetCreatePing`

NewTargetCreatePingWithDefaults instantiates a new TargetCreatePing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativePort

`func (o *TargetCreatePing) GetAdministrativePort() string`

GetAdministrativePort returns the AdministrativePort field if non-nil, zero value otherwise.

### GetAdministrativePortOk

`func (o *TargetCreatePing) GetAdministrativePortOk() (*string, bool)`

GetAdministrativePortOk returns a tuple with the AdministrativePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativePort

`func (o *TargetCreatePing) SetAdministrativePort(v string)`

SetAdministrativePort sets AdministrativePort field to given value.

### HasAdministrativePort

`func (o *TargetCreatePing) HasAdministrativePort() bool`

HasAdministrativePort returns a boolean if a field has been set.

### GetAuthorizationPort

`func (o *TargetCreatePing) GetAuthorizationPort() string`

GetAuthorizationPort returns the AuthorizationPort field if non-nil, zero value otherwise.

### GetAuthorizationPortOk

`func (o *TargetCreatePing) GetAuthorizationPortOk() (*string, bool)`

GetAuthorizationPortOk returns a tuple with the AuthorizationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationPort

`func (o *TargetCreatePing) SetAuthorizationPort(v string)`

SetAuthorizationPort sets AuthorizationPort field to given value.

### HasAuthorizationPort

`func (o *TargetCreatePing) HasAuthorizationPort() bool`

HasAuthorizationPort returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreatePing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreatePing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreatePing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreatePing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreatePing) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreatePing) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreatePing) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreatePing) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreatePing) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreatePing) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreatePing) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreatePing) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreatePing) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreatePing) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreatePing) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreatePing) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreatePing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreatePing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreatePing) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *TargetCreatePing) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetCreatePing) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetCreatePing) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TargetCreatePing) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPingUrl

`func (o *TargetCreatePing) GetPingUrl() string`

GetPingUrl returns the PingUrl field if non-nil, zero value otherwise.

### GetPingUrlOk

`func (o *TargetCreatePing) GetPingUrlOk() (*string, bool)`

GetPingUrlOk returns a tuple with the PingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingUrl

`func (o *TargetCreatePing) SetPingUrl(v string)`

SetPingUrl sets PingUrl field to given value.

### HasPingUrl

`func (o *TargetCreatePing) HasPingUrl() bool`

HasPingUrl returns a boolean if a field has been set.

### GetPrivilegedUser

`func (o *TargetCreatePing) GetPrivilegedUser() string`

GetPrivilegedUser returns the PrivilegedUser field if non-nil, zero value otherwise.

### GetPrivilegedUserOk

`func (o *TargetCreatePing) GetPrivilegedUserOk() (*string, bool)`

GetPrivilegedUserOk returns a tuple with the PrivilegedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedUser

`func (o *TargetCreatePing) SetPrivilegedUser(v string)`

SetPrivilegedUser sets PrivilegedUser field to given value.

### HasPrivilegedUser

`func (o *TargetCreatePing) HasPrivilegedUser() bool`

HasPrivilegedUser returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreatePing) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreatePing) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreatePing) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreatePing) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreatePing) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreatePing) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreatePing) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreatePing) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


