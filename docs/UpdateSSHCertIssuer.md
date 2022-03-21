# UpdateSSHCertIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**AllowedUsers** | **string** | Users allowed to fetch the certificate, e.g root,ubuntu | 
**Extensions** | Pointer to **map[string]string** | Signed certificates with extensions, e.g permit-port-forwarding&#x3D;\\\&quot;\\\&quot; | [optional] 
**Metadata** | Pointer to **string** | A metadata about the issuer | [optional] 
**Name** | **string** | SSH certificate issuer name | 
**NewName** | Pointer to **string** | New item name | [optional] 
**Principals** | Pointer to **string** | Signed certificates with principal, e.g example_role1,example_role2 | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**SecureAccessBastionApi** | Pointer to **string** |  | [optional] 
**SecureAccessBastionSsh** | Pointer to **string** |  | [optional] 
**SecureAccessEnable** | Pointer to **string** |  | [optional] 
**SecureAccessHost** | Pointer to **[]string** |  | [optional] 
**SecureAccessSshCredsUser** | Pointer to **string** |  | [optional] 
**SecureAccessUseInternalBastion** | Pointer to **bool** |  | [optional] 
**SignerKeyName** | **string** | A key to sign the certificate with | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | **int64** | he requested Time To Live for the certificate, in seconds | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateSSHCertIssuer

`func NewUpdateSSHCertIssuer(allowedUsers string, name string, signerKeyName string, ttl int64, ) *UpdateSSHCertIssuer`

NewUpdateSSHCertIssuer instantiates a new UpdateSSHCertIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSSHCertIssuerWithDefaults

`func NewUpdateSSHCertIssuerWithDefaults() *UpdateSSHCertIssuer`

NewUpdateSSHCertIssuerWithDefaults instantiates a new UpdateSSHCertIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *UpdateSSHCertIssuer) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *UpdateSSHCertIssuer) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *UpdateSSHCertIssuer) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *UpdateSSHCertIssuer) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *UpdateSSHCertIssuer) GetAllowedUsers() string`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *UpdateSSHCertIssuer) GetAllowedUsersOk() (*string, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *UpdateSSHCertIssuer) SetAllowedUsers(v string)`

SetAllowedUsers sets AllowedUsers field to given value.


### GetExtensions

`func (o *UpdateSSHCertIssuer) GetExtensions() map[string]string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *UpdateSSHCertIssuer) GetExtensionsOk() (*map[string]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *UpdateSSHCertIssuer) SetExtensions(v map[string]string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *UpdateSSHCertIssuer) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateSSHCertIssuer) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateSSHCertIssuer) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateSSHCertIssuer) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateSSHCertIssuer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UpdateSSHCertIssuer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSSHCertIssuer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSSHCertIssuer) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateSSHCertIssuer) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateSSHCertIssuer) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateSSHCertIssuer) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateSSHCertIssuer) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPrincipals

`func (o *UpdateSSHCertIssuer) GetPrincipals() string`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *UpdateSSHCertIssuer) GetPrincipalsOk() (*string, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *UpdateSSHCertIssuer) SetPrincipals(v string)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *UpdateSSHCertIssuer) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.

### GetRmTag

`func (o *UpdateSSHCertIssuer) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *UpdateSSHCertIssuer) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *UpdateSSHCertIssuer) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *UpdateSSHCertIssuer) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetSecureAccessBastionApi

`func (o *UpdateSSHCertIssuer) GetSecureAccessBastionApi() string`

GetSecureAccessBastionApi returns the SecureAccessBastionApi field if non-nil, zero value otherwise.

### GetSecureAccessBastionApiOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessBastionApiOk() (*string, bool)`

GetSecureAccessBastionApiOk returns a tuple with the SecureAccessBastionApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionApi

`func (o *UpdateSSHCertIssuer) SetSecureAccessBastionApi(v string)`

SetSecureAccessBastionApi sets SecureAccessBastionApi field to given value.

### HasSecureAccessBastionApi

`func (o *UpdateSSHCertIssuer) HasSecureAccessBastionApi() bool`

HasSecureAccessBastionApi returns a boolean if a field has been set.

### GetSecureAccessBastionSsh

`func (o *UpdateSSHCertIssuer) GetSecureAccessBastionSsh() string`

GetSecureAccessBastionSsh returns the SecureAccessBastionSsh field if non-nil, zero value otherwise.

### GetSecureAccessBastionSshOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessBastionSshOk() (*string, bool)`

GetSecureAccessBastionSshOk returns a tuple with the SecureAccessBastionSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionSsh

`func (o *UpdateSSHCertIssuer) SetSecureAccessBastionSsh(v string)`

SetSecureAccessBastionSsh sets SecureAccessBastionSsh field to given value.

### HasSecureAccessBastionSsh

`func (o *UpdateSSHCertIssuer) HasSecureAccessBastionSsh() bool`

HasSecureAccessBastionSsh returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *UpdateSSHCertIssuer) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *UpdateSSHCertIssuer) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *UpdateSSHCertIssuer) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *UpdateSSHCertIssuer) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *UpdateSSHCertIssuer) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *UpdateSSHCertIssuer) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessSshCredsUser

`func (o *UpdateSSHCertIssuer) GetSecureAccessSshCredsUser() string`

GetSecureAccessSshCredsUser returns the SecureAccessSshCredsUser field if non-nil, zero value otherwise.

### GetSecureAccessSshCredsUserOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessSshCredsUserOk() (*string, bool)`

GetSecureAccessSshCredsUserOk returns a tuple with the SecureAccessSshCredsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessSshCredsUser

`func (o *UpdateSSHCertIssuer) SetSecureAccessSshCredsUser(v string)`

SetSecureAccessSshCredsUser sets SecureAccessSshCredsUser field to given value.

### HasSecureAccessSshCredsUser

`func (o *UpdateSSHCertIssuer) HasSecureAccessSshCredsUser() bool`

HasSecureAccessSshCredsUser returns a boolean if a field has been set.

### GetSecureAccessUseInternalBastion

`func (o *UpdateSSHCertIssuer) GetSecureAccessUseInternalBastion() bool`

GetSecureAccessUseInternalBastion returns the SecureAccessUseInternalBastion field if non-nil, zero value otherwise.

### GetSecureAccessUseInternalBastionOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessUseInternalBastionOk() (*bool, bool)`

GetSecureAccessUseInternalBastionOk returns a tuple with the SecureAccessUseInternalBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUseInternalBastion

`func (o *UpdateSSHCertIssuer) SetSecureAccessUseInternalBastion(v bool)`

SetSecureAccessUseInternalBastion sets SecureAccessUseInternalBastion field to given value.

### HasSecureAccessUseInternalBastion

`func (o *UpdateSSHCertIssuer) HasSecureAccessUseInternalBastion() bool`

HasSecureAccessUseInternalBastion returns a boolean if a field has been set.

### GetSignerKeyName

`func (o *UpdateSSHCertIssuer) GetSignerKeyName() string`

GetSignerKeyName returns the SignerKeyName field if non-nil, zero value otherwise.

### GetSignerKeyNameOk

`func (o *UpdateSSHCertIssuer) GetSignerKeyNameOk() (*string, bool)`

GetSignerKeyNameOk returns a tuple with the SignerKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerKeyName

`func (o *UpdateSSHCertIssuer) SetSignerKeyName(v string)`

SetSignerKeyName sets SignerKeyName field to given value.


### GetToken

`func (o *UpdateSSHCertIssuer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateSSHCertIssuer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateSSHCertIssuer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateSSHCertIssuer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *UpdateSSHCertIssuer) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *UpdateSSHCertIssuer) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *UpdateSSHCertIssuer) SetTtl(v int64)`

SetTtl sets Ttl field to given value.


### GetUidToken

`func (o *UpdateSSHCertIssuer) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateSSHCertIssuer) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateSSHCertIssuer) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateSSHCertIssuer) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


