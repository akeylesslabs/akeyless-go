# DynamicSecretCreateLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindDn** | Pointer to **string** | Bind DN | [optional] 
**BindDnPassword** | Pointer to **string** | Bind DN Password | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**ExternalUsername** | Pointer to **string** | Externally provided username [true/false] | [optional] [default to "false"]
**GroupDn** | Pointer to **string** | Group DN which the temporary user should be added | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**LdapCaCert** | Pointer to **string** | CA Certificate File Content | [optional] 
**LdapUrl** | Pointer to **string** | LDAP Server URL | [optional] 
**Name** | **string** | Dynamic secret name | 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExpiration** | Pointer to **string** | Token expiration | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserAttribute** | Pointer to **string** | User Attribute | [optional] 
**UserDn** | Pointer to **string** | User DN | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateLdap

`func NewDynamicSecretCreateLdap(name string, ) *DynamicSecretCreateLdap`

NewDynamicSecretCreateLdap instantiates a new DynamicSecretCreateLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateLdapWithDefaults

`func NewDynamicSecretCreateLdapWithDefaults() *DynamicSecretCreateLdap`

NewDynamicSecretCreateLdapWithDefaults instantiates a new DynamicSecretCreateLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindDn

`func (o *DynamicSecretCreateLdap) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *DynamicSecretCreateLdap) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *DynamicSecretCreateLdap) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *DynamicSecretCreateLdap) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindDnPassword

`func (o *DynamicSecretCreateLdap) GetBindDnPassword() string`

GetBindDnPassword returns the BindDnPassword field if non-nil, zero value otherwise.

### GetBindDnPasswordOk

`func (o *DynamicSecretCreateLdap) GetBindDnPasswordOk() (*string, bool)`

GetBindDnPasswordOk returns a tuple with the BindDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDnPassword

`func (o *DynamicSecretCreateLdap) SetBindDnPassword(v string)`

SetBindDnPassword sets BindDnPassword field to given value.

### HasBindDnPassword

`func (o *DynamicSecretCreateLdap) HasBindDnPassword() bool`

HasBindDnPassword returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateLdap) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateLdap) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateLdap) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateLdap) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateLdap) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateLdap) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateLdap) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateLdap) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalUsername

`func (o *DynamicSecretCreateLdap) GetExternalUsername() string`

GetExternalUsername returns the ExternalUsername field if non-nil, zero value otherwise.

### GetExternalUsernameOk

`func (o *DynamicSecretCreateLdap) GetExternalUsernameOk() (*string, bool)`

GetExternalUsernameOk returns a tuple with the ExternalUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUsername

`func (o *DynamicSecretCreateLdap) SetExternalUsername(v string)`

SetExternalUsername sets ExternalUsername field to given value.

### HasExternalUsername

`func (o *DynamicSecretCreateLdap) HasExternalUsername() bool`

HasExternalUsername returns a boolean if a field has been set.

### GetGroupDn

`func (o *DynamicSecretCreateLdap) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *DynamicSecretCreateLdap) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *DynamicSecretCreateLdap) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *DynamicSecretCreateLdap) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateLdap) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateLdap) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateLdap) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateLdap) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLdapCaCert

`func (o *DynamicSecretCreateLdap) GetLdapCaCert() string`

GetLdapCaCert returns the LdapCaCert field if non-nil, zero value otherwise.

### GetLdapCaCertOk

`func (o *DynamicSecretCreateLdap) GetLdapCaCertOk() (*string, bool)`

GetLdapCaCertOk returns a tuple with the LdapCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCaCert

`func (o *DynamicSecretCreateLdap) SetLdapCaCert(v string)`

SetLdapCaCert sets LdapCaCert field to given value.

### HasLdapCaCert

`func (o *DynamicSecretCreateLdap) HasLdapCaCert() bool`

HasLdapCaCert returns a boolean if a field has been set.

### GetLdapUrl

`func (o *DynamicSecretCreateLdap) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *DynamicSecretCreateLdap) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *DynamicSecretCreateLdap) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *DynamicSecretCreateLdap) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateLdap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateLdap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateLdap) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *DynamicSecretCreateLdap) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretCreateLdap) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretCreateLdap) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretCreateLdap) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateLdap) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateLdap) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateLdap) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateLdap) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateLdap) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateLdap) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateLdap) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateLdap) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateLdap) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateLdap) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateLdap) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateLdap) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateLdap) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateLdap) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateLdap) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateLdap) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenExpiration

`func (o *DynamicSecretCreateLdap) GetTokenExpiration() string`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *DynamicSecretCreateLdap) GetTokenExpirationOk() (*string, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *DynamicSecretCreateLdap) SetTokenExpiration(v string)`

SetTokenExpiration sets TokenExpiration field to given value.

### HasTokenExpiration

`func (o *DynamicSecretCreateLdap) HasTokenExpiration() bool`

HasTokenExpiration returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateLdap) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateLdap) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateLdap) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateLdap) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserAttribute

`func (o *DynamicSecretCreateLdap) GetUserAttribute() string`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *DynamicSecretCreateLdap) GetUserAttributeOk() (*string, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *DynamicSecretCreateLdap) SetUserAttribute(v string)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *DynamicSecretCreateLdap) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### GetUserDn

`func (o *DynamicSecretCreateLdap) GetUserDn() string`

GetUserDn returns the UserDn field if non-nil, zero value otherwise.

### GetUserDnOk

`func (o *DynamicSecretCreateLdap) GetUserDnOk() (*string, bool)`

GetUserDnOk returns a tuple with the UserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDn

`func (o *DynamicSecretCreateLdap) SetUserDn(v string)`

SetUserDn sets UserDn field to given value.

### HasUserDn

`func (o *DynamicSecretCreateLdap) HasUserDn() bool`

HasUserDn returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateLdap) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateLdap) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateLdap) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateLdap) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


