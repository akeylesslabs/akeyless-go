# CreateLdapTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindDn** | **string** | Bind DN | 
**BindDnPassword** | **string** | Bind DN Password | 
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LdapCaCert** | Pointer to **string** | CA Certificate File Content | [optional] 
**LdapUrl** | **string** | LDAP Server URL | 
**Name** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExpiration** | Pointer to **string** | Token expiration | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateLdapTarget

`func NewCreateLdapTarget(bindDn string, bindDnPassword string, ldapUrl string, name string, ) *CreateLdapTarget`

NewCreateLdapTarget instantiates a new CreateLdapTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLdapTargetWithDefaults

`func NewCreateLdapTargetWithDefaults() *CreateLdapTarget`

NewCreateLdapTargetWithDefaults instantiates a new CreateLdapTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindDn

`func (o *CreateLdapTarget) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *CreateLdapTarget) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *CreateLdapTarget) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.


### GetBindDnPassword

`func (o *CreateLdapTarget) GetBindDnPassword() string`

GetBindDnPassword returns the BindDnPassword field if non-nil, zero value otherwise.

### GetBindDnPasswordOk

`func (o *CreateLdapTarget) GetBindDnPasswordOk() (*string, bool)`

GetBindDnPasswordOk returns a tuple with the BindDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDnPassword

`func (o *CreateLdapTarget) SetBindDnPassword(v string)`

SetBindDnPassword sets BindDnPassword field to given value.


### GetComment

`func (o *CreateLdapTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateLdapTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateLdapTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateLdapTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetKey

`func (o *CreateLdapTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateLdapTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateLdapTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateLdapTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLdapCaCert

`func (o *CreateLdapTarget) GetLdapCaCert() string`

GetLdapCaCert returns the LdapCaCert field if non-nil, zero value otherwise.

### GetLdapCaCertOk

`func (o *CreateLdapTarget) GetLdapCaCertOk() (*string, bool)`

GetLdapCaCertOk returns a tuple with the LdapCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCaCert

`func (o *CreateLdapTarget) SetLdapCaCert(v string)`

SetLdapCaCert sets LdapCaCert field to given value.

### HasLdapCaCert

`func (o *CreateLdapTarget) HasLdapCaCert() bool`

HasLdapCaCert returns a boolean if a field has been set.

### GetLdapUrl

`func (o *CreateLdapTarget) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *CreateLdapTarget) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *CreateLdapTarget) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.


### GetName

`func (o *CreateLdapTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLdapTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLdapTarget) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *CreateLdapTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateLdapTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateLdapTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateLdapTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenExpiration

`func (o *CreateLdapTarget) GetTokenExpiration() string`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *CreateLdapTarget) GetTokenExpirationOk() (*string, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *CreateLdapTarget) SetTokenExpiration(v string)`

SetTokenExpiration sets TokenExpiration field to given value.

### HasTokenExpiration

`func (o *CreateLdapTarget) HasTokenExpiration() bool`

HasTokenExpiration returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateLdapTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateLdapTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateLdapTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateLdapTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


