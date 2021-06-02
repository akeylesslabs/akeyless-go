# CreateSSHCertIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUsers** | **string** | Users allowed to fetch the certificate, e.g root,ubuntu | 
**Extensions** | Pointer to **map[string]string** | Signed certificates with extensions, e.g permit-port-forwarding&#x3D;\\\&quot;\\\&quot; | [optional] 
**Metadata** | Pointer to **string** | A metadata about the issuer | [optional] 
**Name** | **string** | SSH certificate issuer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Principals** | Pointer to **string** | Signed certificates with principal, e.g example_role1,example_role2 | [optional] 
**SignerKeyName** | **string** | A key to sign the certificate with | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | **int64** | The requested Time To Live for the certificate, use second units | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreateSSHCertIssuer

`func NewCreateSSHCertIssuer(allowedUsers string, name string, signerKeyName string, ttl int64, ) *CreateSSHCertIssuer`

NewCreateSSHCertIssuer instantiates a new CreateSSHCertIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSSHCertIssuerWithDefaults

`func NewCreateSSHCertIssuerWithDefaults() *CreateSSHCertIssuer`

NewCreateSSHCertIssuerWithDefaults instantiates a new CreateSSHCertIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedUsers

`func (o *CreateSSHCertIssuer) GetAllowedUsers() string`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *CreateSSHCertIssuer) GetAllowedUsersOk() (*string, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *CreateSSHCertIssuer) SetAllowedUsers(v string)`

SetAllowedUsers sets AllowedUsers field to given value.


### GetExtensions

`func (o *CreateSSHCertIssuer) GetExtensions() map[string]string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CreateSSHCertIssuer) GetExtensionsOk() (*map[string]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CreateSSHCertIssuer) SetExtensions(v map[string]string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CreateSSHCertIssuer) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateSSHCertIssuer) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSSHCertIssuer) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSSHCertIssuer) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSSHCertIssuer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateSSHCertIssuer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSSHCertIssuer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSSHCertIssuer) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateSSHCertIssuer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateSSHCertIssuer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateSSHCertIssuer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateSSHCertIssuer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrincipals

`func (o *CreateSSHCertIssuer) GetPrincipals() string`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *CreateSSHCertIssuer) GetPrincipalsOk() (*string, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *CreateSSHCertIssuer) SetPrincipals(v string)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *CreateSSHCertIssuer) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.

### GetSignerKeyName

`func (o *CreateSSHCertIssuer) GetSignerKeyName() string`

GetSignerKeyName returns the SignerKeyName field if non-nil, zero value otherwise.

### GetSignerKeyNameOk

`func (o *CreateSSHCertIssuer) GetSignerKeyNameOk() (*string, bool)`

GetSignerKeyNameOk returns a tuple with the SignerKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerKeyName

`func (o *CreateSSHCertIssuer) SetSignerKeyName(v string)`

SetSignerKeyName sets SignerKeyName field to given value.


### GetToken

`func (o *CreateSSHCertIssuer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateSSHCertIssuer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateSSHCertIssuer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateSSHCertIssuer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *CreateSSHCertIssuer) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateSSHCertIssuer) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateSSHCertIssuer) SetTtl(v int64)`

SetTtl sets Ttl field to given value.


### GetUidToken

`func (o *CreateSSHCertIssuer) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateSSHCertIssuer) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateSSHCertIssuer) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateSSHCertIssuer) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateSSHCertIssuer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateSSHCertIssuer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateSSHCertIssuer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateSSHCertIssuer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


