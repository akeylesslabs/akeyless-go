# GatewayUpdateLdapAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | The access ID of the Ldap auth method | [optional] 
**BindDn** | Pointer to **string** | Bind DN | [optional] 
**BindDnPassword** | Pointer to **string** | Bind DN Password | [optional] 
**GroupAttr** | Pointer to **string** | Group Attr | [optional] 
**GroupDn** | Pointer to **string** | Group Dn | [optional] 
**GroupFilter** | Pointer to **string** | Group Filter | [optional] 
**LdapCaCert** | Pointer to **string** | LDAP CA Certificate (base64 encoded) | [optional] 
**LdapEnable** | Pointer to **string** | Enable Ldap | [optional] 
**LdapUrl** | Pointer to **string** | LDAP Server URL, e.g. ldap://planetexpress.com:389 | [optional] 
**LdapAnonymousSearch** | Pointer to **bool** | Ldap Anonymous Search | [optional] 
**SigningKeyData** | Pointer to **string** | The private key (base64 encoded), associated with the public key defined in the Ldap auth | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserAttribute** | Pointer to **string** | User Attribute | [optional] 
**UserDn** | Pointer to **string** | User DN | [optional] 

## Methods

### NewGatewayUpdateLdapAuthConfig

`func NewGatewayUpdateLdapAuthConfig() *GatewayUpdateLdapAuthConfig`

NewGatewayUpdateLdapAuthConfig instantiates a new GatewayUpdateLdapAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLdapAuthConfigWithDefaults

`func NewGatewayUpdateLdapAuthConfigWithDefaults() *GatewayUpdateLdapAuthConfig`

NewGatewayUpdateLdapAuthConfigWithDefaults instantiates a new GatewayUpdateLdapAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *GatewayUpdateLdapAuthConfig) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *GatewayUpdateLdapAuthConfig) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *GatewayUpdateLdapAuthConfig) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *GatewayUpdateLdapAuthConfig) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetBindDn

`func (o *GatewayUpdateLdapAuthConfig) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *GatewayUpdateLdapAuthConfig) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *GatewayUpdateLdapAuthConfig) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *GatewayUpdateLdapAuthConfig) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindDnPassword

`func (o *GatewayUpdateLdapAuthConfig) GetBindDnPassword() string`

GetBindDnPassword returns the BindDnPassword field if non-nil, zero value otherwise.

### GetBindDnPasswordOk

`func (o *GatewayUpdateLdapAuthConfig) GetBindDnPasswordOk() (*string, bool)`

GetBindDnPasswordOk returns a tuple with the BindDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDnPassword

`func (o *GatewayUpdateLdapAuthConfig) SetBindDnPassword(v string)`

SetBindDnPassword sets BindDnPassword field to given value.

### HasBindDnPassword

`func (o *GatewayUpdateLdapAuthConfig) HasBindDnPassword() bool`

HasBindDnPassword returns a boolean if a field has been set.

### GetGroupAttr

`func (o *GatewayUpdateLdapAuthConfig) GetGroupAttr() string`

GetGroupAttr returns the GroupAttr field if non-nil, zero value otherwise.

### GetGroupAttrOk

`func (o *GatewayUpdateLdapAuthConfig) GetGroupAttrOk() (*string, bool)`

GetGroupAttrOk returns a tuple with the GroupAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttr

`func (o *GatewayUpdateLdapAuthConfig) SetGroupAttr(v string)`

SetGroupAttr sets GroupAttr field to given value.

### HasGroupAttr

`func (o *GatewayUpdateLdapAuthConfig) HasGroupAttr() bool`

HasGroupAttr returns a boolean if a field has been set.

### GetGroupDn

`func (o *GatewayUpdateLdapAuthConfig) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *GatewayUpdateLdapAuthConfig) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *GatewayUpdateLdapAuthConfig) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *GatewayUpdateLdapAuthConfig) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetGroupFilter

`func (o *GatewayUpdateLdapAuthConfig) GetGroupFilter() string`

GetGroupFilter returns the GroupFilter field if non-nil, zero value otherwise.

### GetGroupFilterOk

`func (o *GatewayUpdateLdapAuthConfig) GetGroupFilterOk() (*string, bool)`

GetGroupFilterOk returns a tuple with the GroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFilter

`func (o *GatewayUpdateLdapAuthConfig) SetGroupFilter(v string)`

SetGroupFilter sets GroupFilter field to given value.

### HasGroupFilter

`func (o *GatewayUpdateLdapAuthConfig) HasGroupFilter() bool`

HasGroupFilter returns a boolean if a field has been set.

### GetLdapCaCert

`func (o *GatewayUpdateLdapAuthConfig) GetLdapCaCert() string`

GetLdapCaCert returns the LdapCaCert field if non-nil, zero value otherwise.

### GetLdapCaCertOk

`func (o *GatewayUpdateLdapAuthConfig) GetLdapCaCertOk() (*string, bool)`

GetLdapCaCertOk returns a tuple with the LdapCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCaCert

`func (o *GatewayUpdateLdapAuthConfig) SetLdapCaCert(v string)`

SetLdapCaCert sets LdapCaCert field to given value.

### HasLdapCaCert

`func (o *GatewayUpdateLdapAuthConfig) HasLdapCaCert() bool`

HasLdapCaCert returns a boolean if a field has been set.

### GetLdapEnable

`func (o *GatewayUpdateLdapAuthConfig) GetLdapEnable() string`

GetLdapEnable returns the LdapEnable field if non-nil, zero value otherwise.

### GetLdapEnableOk

`func (o *GatewayUpdateLdapAuthConfig) GetLdapEnableOk() (*string, bool)`

GetLdapEnableOk returns a tuple with the LdapEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapEnable

`func (o *GatewayUpdateLdapAuthConfig) SetLdapEnable(v string)`

SetLdapEnable sets LdapEnable field to given value.

### HasLdapEnable

`func (o *GatewayUpdateLdapAuthConfig) HasLdapEnable() bool`

HasLdapEnable returns a boolean if a field has been set.

### GetLdapUrl

`func (o *GatewayUpdateLdapAuthConfig) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *GatewayUpdateLdapAuthConfig) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *GatewayUpdateLdapAuthConfig) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *GatewayUpdateLdapAuthConfig) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetLdapAnonymousSearch

`func (o *GatewayUpdateLdapAuthConfig) GetLdapAnonymousSearch() bool`

GetLdapAnonymousSearch returns the LdapAnonymousSearch field if non-nil, zero value otherwise.

### GetLdapAnonymousSearchOk

`func (o *GatewayUpdateLdapAuthConfig) GetLdapAnonymousSearchOk() (*bool, bool)`

GetLdapAnonymousSearchOk returns a tuple with the LdapAnonymousSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAnonymousSearch

`func (o *GatewayUpdateLdapAuthConfig) SetLdapAnonymousSearch(v bool)`

SetLdapAnonymousSearch sets LdapAnonymousSearch field to given value.

### HasLdapAnonymousSearch

`func (o *GatewayUpdateLdapAuthConfig) HasLdapAnonymousSearch() bool`

HasLdapAnonymousSearch returns a boolean if a field has been set.

### GetSigningKeyData

`func (o *GatewayUpdateLdapAuthConfig) GetSigningKeyData() string`

GetSigningKeyData returns the SigningKeyData field if non-nil, zero value otherwise.

### GetSigningKeyDataOk

`func (o *GatewayUpdateLdapAuthConfig) GetSigningKeyDataOk() (*string, bool)`

GetSigningKeyDataOk returns a tuple with the SigningKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyData

`func (o *GatewayUpdateLdapAuthConfig) SetSigningKeyData(v string)`

SetSigningKeyData sets SigningKeyData field to given value.

### HasSigningKeyData

`func (o *GatewayUpdateLdapAuthConfig) HasSigningKeyData() bool`

HasSigningKeyData returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLdapAuthConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLdapAuthConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLdapAuthConfig) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLdapAuthConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLdapAuthConfig) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLdapAuthConfig) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLdapAuthConfig) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLdapAuthConfig) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserAttribute

`func (o *GatewayUpdateLdapAuthConfig) GetUserAttribute() string`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *GatewayUpdateLdapAuthConfig) GetUserAttributeOk() (*string, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *GatewayUpdateLdapAuthConfig) SetUserAttribute(v string)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *GatewayUpdateLdapAuthConfig) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### GetUserDn

`func (o *GatewayUpdateLdapAuthConfig) GetUserDn() string`

GetUserDn returns the UserDn field if non-nil, zero value otherwise.

### GetUserDnOk

`func (o *GatewayUpdateLdapAuthConfig) GetUserDnOk() (*string, bool)`

GetUserDnOk returns a tuple with the UserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDn

`func (o *GatewayUpdateLdapAuthConfig) SetUserDn(v string)`

SetUserDn sets UserDn field to given value.

### HasUserDn

`func (o *GatewayUpdateLdapAuthConfig) HasUserDn() bool`

HasUserDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


