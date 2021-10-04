# CreateLdapTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | Access ID | 
**BindDn** | Pointer to **string** | Bind DN | [optional] 
**BindDnPassword** | Pointer to **string** | Bind DN Password | [optional] 
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**EnableAnonymSearch** | Pointer to **bool** | EnableAnonymousSearch | [optional] 
**GroupAttribute** | Pointer to **string** | Group attribute | [optional] 
**GroupDn** | Pointer to **string** | Group DN | [optional] 
**GroupFilter** | Pointer to **string** | Group attribute | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LdapCaCert** | Pointer to **string** | CA Certificate File Content | [optional] 
**LdapUrl** | **string** | LDAP Server URL | 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**PrivateKey** | Pointer to **string** | Base64-encoded ldap private key text | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExpiration** | Pointer to **string** | Token expiration | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserAttribute** | Pointer to **string** | User Attribute | [optional] 
**UserDn** | **string** | User DN | 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreateLdapTarget

`func NewCreateLdapTarget(accessId string, ldapUrl string, name string, userDn string, ) *CreateLdapTarget`

NewCreateLdapTarget instantiates a new CreateLdapTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLdapTargetWithDefaults

`func NewCreateLdapTargetWithDefaults() *CreateLdapTarget`

NewCreateLdapTargetWithDefaults instantiates a new CreateLdapTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *CreateLdapTarget) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *CreateLdapTarget) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *CreateLdapTarget) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


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

### HasBindDn

`func (o *CreateLdapTarget) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

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

### HasBindDnPassword

`func (o *CreateLdapTarget) HasBindDnPassword() bool`

HasBindDnPassword returns a boolean if a field has been set.

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

### GetEnableAnonymSearch

`func (o *CreateLdapTarget) GetEnableAnonymSearch() bool`

GetEnableAnonymSearch returns the EnableAnonymSearch field if non-nil, zero value otherwise.

### GetEnableAnonymSearchOk

`func (o *CreateLdapTarget) GetEnableAnonymSearchOk() (*bool, bool)`

GetEnableAnonymSearchOk returns a tuple with the EnableAnonymSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAnonymSearch

`func (o *CreateLdapTarget) SetEnableAnonymSearch(v bool)`

SetEnableAnonymSearch sets EnableAnonymSearch field to given value.

### HasEnableAnonymSearch

`func (o *CreateLdapTarget) HasEnableAnonymSearch() bool`

HasEnableAnonymSearch returns a boolean if a field has been set.

### GetGroupAttribute

`func (o *CreateLdapTarget) GetGroupAttribute() string`

GetGroupAttribute returns the GroupAttribute field if non-nil, zero value otherwise.

### GetGroupAttributeOk

`func (o *CreateLdapTarget) GetGroupAttributeOk() (*string, bool)`

GetGroupAttributeOk returns a tuple with the GroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttribute

`func (o *CreateLdapTarget) SetGroupAttribute(v string)`

SetGroupAttribute sets GroupAttribute field to given value.

### HasGroupAttribute

`func (o *CreateLdapTarget) HasGroupAttribute() bool`

HasGroupAttribute returns a boolean if a field has been set.

### GetGroupDn

`func (o *CreateLdapTarget) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *CreateLdapTarget) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *CreateLdapTarget) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *CreateLdapTarget) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetGroupFilter

`func (o *CreateLdapTarget) GetGroupFilter() string`

GetGroupFilter returns the GroupFilter field if non-nil, zero value otherwise.

### GetGroupFilterOk

`func (o *CreateLdapTarget) GetGroupFilterOk() (*string, bool)`

GetGroupFilterOk returns a tuple with the GroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFilter

`func (o *CreateLdapTarget) SetGroupFilter(v string)`

SetGroupFilter sets GroupFilter field to given value.

### HasGroupFilter

`func (o *CreateLdapTarget) HasGroupFilter() bool`

HasGroupFilter returns a boolean if a field has been set.

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


### GetPassword

`func (o *CreateLdapTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateLdapTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateLdapTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateLdapTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CreateLdapTarget) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateLdapTarget) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateLdapTarget) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CreateLdapTarget) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

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

### GetUserAttribute

`func (o *CreateLdapTarget) GetUserAttribute() string`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *CreateLdapTarget) GetUserAttributeOk() (*string, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *CreateLdapTarget) SetUserAttribute(v string)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *CreateLdapTarget) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### GetUserDn

`func (o *CreateLdapTarget) GetUserDn() string`

GetUserDn returns the UserDn field if non-nil, zero value otherwise.

### GetUserDnOk

`func (o *CreateLdapTarget) GetUserDnOk() (*string, bool)`

GetUserDnOk returns a tuple with the UserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDn

`func (o *CreateLdapTarget) SetUserDn(v string)`

SetUserDn sets UserDn field to given value.


### GetUsername

`func (o *CreateLdapTarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateLdapTarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateLdapTarget) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateLdapTarget) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


