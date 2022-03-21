# GatewayCreateProducerLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindDn** | Pointer to **string** | Bind DN | [optional] 
**BindDnPassword** | Pointer to **string** | Bind DN Password | [optional] 
**ExternalUsername** | Pointer to **string** | Fixed user | [optional] [default to "false"]
**LdapCaCert** | Pointer to **string** | CA Certificate File Content | [optional] 
**LdapUrl** | Pointer to **string** | LDAP Server URL | [optional] 
**Name** | **string** | Producer name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExpiration** | Pointer to **string** | Token expiration | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserAttribute** | Pointer to **string** | User Attribute | [optional] 
**UserDn** | Pointer to **string** | User DN | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerLdap

`func NewGatewayCreateProducerLdap(name string, ) *GatewayCreateProducerLdap`

NewGatewayCreateProducerLdap instantiates a new GatewayCreateProducerLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerLdapWithDefaults

`func NewGatewayCreateProducerLdapWithDefaults() *GatewayCreateProducerLdap`

NewGatewayCreateProducerLdapWithDefaults instantiates a new GatewayCreateProducerLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindDn

`func (o *GatewayCreateProducerLdap) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *GatewayCreateProducerLdap) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *GatewayCreateProducerLdap) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *GatewayCreateProducerLdap) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindDnPassword

`func (o *GatewayCreateProducerLdap) GetBindDnPassword() string`

GetBindDnPassword returns the BindDnPassword field if non-nil, zero value otherwise.

### GetBindDnPasswordOk

`func (o *GatewayCreateProducerLdap) GetBindDnPasswordOk() (*string, bool)`

GetBindDnPasswordOk returns a tuple with the BindDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDnPassword

`func (o *GatewayCreateProducerLdap) SetBindDnPassword(v string)`

SetBindDnPassword sets BindDnPassword field to given value.

### HasBindDnPassword

`func (o *GatewayCreateProducerLdap) HasBindDnPassword() bool`

HasBindDnPassword returns a boolean if a field has been set.

### GetExternalUsername

`func (o *GatewayCreateProducerLdap) GetExternalUsername() string`

GetExternalUsername returns the ExternalUsername field if non-nil, zero value otherwise.

### GetExternalUsernameOk

`func (o *GatewayCreateProducerLdap) GetExternalUsernameOk() (*string, bool)`

GetExternalUsernameOk returns a tuple with the ExternalUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUsername

`func (o *GatewayCreateProducerLdap) SetExternalUsername(v string)`

SetExternalUsername sets ExternalUsername field to given value.

### HasExternalUsername

`func (o *GatewayCreateProducerLdap) HasExternalUsername() bool`

HasExternalUsername returns a boolean if a field has been set.

### GetLdapCaCert

`func (o *GatewayCreateProducerLdap) GetLdapCaCert() string`

GetLdapCaCert returns the LdapCaCert field if non-nil, zero value otherwise.

### GetLdapCaCertOk

`func (o *GatewayCreateProducerLdap) GetLdapCaCertOk() (*string, bool)`

GetLdapCaCertOk returns a tuple with the LdapCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCaCert

`func (o *GatewayCreateProducerLdap) SetLdapCaCert(v string)`

SetLdapCaCert sets LdapCaCert field to given value.

### HasLdapCaCert

`func (o *GatewayCreateProducerLdap) HasLdapCaCert() bool`

HasLdapCaCert returns a boolean if a field has been set.

### GetLdapUrl

`func (o *GatewayCreateProducerLdap) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *GatewayCreateProducerLdap) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *GatewayCreateProducerLdap) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *GatewayCreateProducerLdap) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerLdap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerLdap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerLdap) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerLdap) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerLdap) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerLdap) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerLdap) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerLdap) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerLdap) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerLdap) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerLdap) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerLdap) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerLdap) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerLdap) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerLdap) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerLdap) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerLdap) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerLdap) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerLdap) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenExpiration

`func (o *GatewayCreateProducerLdap) GetTokenExpiration() string`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *GatewayCreateProducerLdap) GetTokenExpirationOk() (*string, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *GatewayCreateProducerLdap) SetTokenExpiration(v string)`

SetTokenExpiration sets TokenExpiration field to given value.

### HasTokenExpiration

`func (o *GatewayCreateProducerLdap) HasTokenExpiration() bool`

HasTokenExpiration returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerLdap) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerLdap) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerLdap) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerLdap) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserAttribute

`func (o *GatewayCreateProducerLdap) GetUserAttribute() string`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *GatewayCreateProducerLdap) GetUserAttributeOk() (*string, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *GatewayCreateProducerLdap) SetUserAttribute(v string)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *GatewayCreateProducerLdap) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### GetUserDn

`func (o *GatewayCreateProducerLdap) GetUserDn() string`

GetUserDn returns the UserDn field if non-nil, zero value otherwise.

### GetUserDnOk

`func (o *GatewayCreateProducerLdap) GetUserDnOk() (*string, bool)`

GetUserDnOk returns a tuple with the UserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDn

`func (o *GatewayCreateProducerLdap) SetUserDn(v string)`

SetUserDn sets UserDn field to given value.

### HasUserDn

`func (o *GatewayCreateProducerLdap) HasUserDn() bool`

HasUserDn returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerLdap) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerLdap) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerLdap) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerLdap) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


