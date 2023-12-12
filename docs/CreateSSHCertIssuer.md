# CreateSSHCertIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshCertIssuerHostProvider** | Pointer to **string** |  | [optional] 
**AllowedUsers** | **string** | Users allowed to fetch the certificate, e.g root,ubuntu | 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Extensions** | Pointer to **map[string]string** | Signed certificates with extensions, e.g permit-port-forwarding&#x3D;\\\&quot;\\\&quot; | [optional] 
**HostProvider** | Pointer to **string** | Host provider type [explicit/target] | [optional] [default to "explicit"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | SSH certificate issuer name | 
**Principals** | Pointer to **string** | Signed certificates with principal, e.g example_role1,example_role2 | [optional] 
**SecureAccessBastionApi** | Pointer to **string** | Bastion&#39;s SSH control API endpoint. E.g. https://my.bastion:9900 | [optional] 
**SecureAccessBastionSsh** | Pointer to **string** | Bastion&#39;s SSH server. E.g. my.bastion:22 | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessSshCredsUser** | Pointer to **string** | SSH username to connect to target server, must be in &#39;Allowed Users&#39; list | [optional] 
**SecureAccessUseInternalBastion** | Pointer to **bool** | Use internal SSH Bastion | [optional] 
**SignerKeyName** | **string** | A key to sign the certificate with | 
**Tag** | Pointer to **[]string** | List of the tags attached to this key | [optional] 
**Target** | Pointer to **[]string** | A list of existing targets to be associated, Relevant only for Secure Remote Access, To specify multiple targets use argument multiple times | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | **int64** | The requested Time To Live for the certificate, in seconds | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

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

### GetSshCertIssuerHostProvider

`func (o *CreateSSHCertIssuer) GetSshCertIssuerHostProvider() string`

GetSshCertIssuerHostProvider returns the SshCertIssuerHostProvider field if non-nil, zero value otherwise.

### GetSshCertIssuerHostProviderOk

`func (o *CreateSSHCertIssuer) GetSshCertIssuerHostProviderOk() (*string, bool)`

GetSshCertIssuerHostProviderOk returns a tuple with the SshCertIssuerHostProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCertIssuerHostProvider

`func (o *CreateSSHCertIssuer) SetSshCertIssuerHostProvider(v string)`

SetSshCertIssuerHostProvider sets SshCertIssuerHostProvider field to given value.

### HasSshCertIssuerHostProvider

`func (o *CreateSSHCertIssuer) HasSshCertIssuerHostProvider() bool`

HasSshCertIssuerHostProvider returns a boolean if a field has been set.

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


### GetDeleteProtection

`func (o *CreateSSHCertIssuer) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateSSHCertIssuer) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateSSHCertIssuer) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateSSHCertIssuer) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSSHCertIssuer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSSHCertIssuer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSSHCertIssuer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSSHCertIssuer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### GetHostProvider

`func (o *CreateSSHCertIssuer) GetHostProvider() string`

GetHostProvider returns the HostProvider field if non-nil, zero value otherwise.

### GetHostProviderOk

`func (o *CreateSSHCertIssuer) GetHostProviderOk() (*string, bool)`

GetHostProviderOk returns a tuple with the HostProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostProvider

`func (o *CreateSSHCertIssuer) SetHostProvider(v string)`

SetHostProvider sets HostProvider field to given value.

### HasHostProvider

`func (o *CreateSSHCertIssuer) HasHostProvider() bool`

HasHostProvider returns a boolean if a field has been set.

### GetJson

`func (o *CreateSSHCertIssuer) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateSSHCertIssuer) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateSSHCertIssuer) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateSSHCertIssuer) HasJson() bool`

HasJson returns a boolean if a field has been set.

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

### GetSecureAccessBastionApi

`func (o *CreateSSHCertIssuer) GetSecureAccessBastionApi() string`

GetSecureAccessBastionApi returns the SecureAccessBastionApi field if non-nil, zero value otherwise.

### GetSecureAccessBastionApiOk

`func (o *CreateSSHCertIssuer) GetSecureAccessBastionApiOk() (*string, bool)`

GetSecureAccessBastionApiOk returns a tuple with the SecureAccessBastionApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionApi

`func (o *CreateSSHCertIssuer) SetSecureAccessBastionApi(v string)`

SetSecureAccessBastionApi sets SecureAccessBastionApi field to given value.

### HasSecureAccessBastionApi

`func (o *CreateSSHCertIssuer) HasSecureAccessBastionApi() bool`

HasSecureAccessBastionApi returns a boolean if a field has been set.

### GetSecureAccessBastionSsh

`func (o *CreateSSHCertIssuer) GetSecureAccessBastionSsh() string`

GetSecureAccessBastionSsh returns the SecureAccessBastionSsh field if non-nil, zero value otherwise.

### GetSecureAccessBastionSshOk

`func (o *CreateSSHCertIssuer) GetSecureAccessBastionSshOk() (*string, bool)`

GetSecureAccessBastionSshOk returns a tuple with the SecureAccessBastionSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionSsh

`func (o *CreateSSHCertIssuer) SetSecureAccessBastionSsh(v string)`

SetSecureAccessBastionSsh sets SecureAccessBastionSsh field to given value.

### HasSecureAccessBastionSsh

`func (o *CreateSSHCertIssuer) HasSecureAccessBastionSsh() bool`

HasSecureAccessBastionSsh returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *CreateSSHCertIssuer) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *CreateSSHCertIssuer) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *CreateSSHCertIssuer) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *CreateSSHCertIssuer) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *CreateSSHCertIssuer) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *CreateSSHCertIssuer) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *CreateSSHCertIssuer) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *CreateSSHCertIssuer) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessSshCredsUser

`func (o *CreateSSHCertIssuer) GetSecureAccessSshCredsUser() string`

GetSecureAccessSshCredsUser returns the SecureAccessSshCredsUser field if non-nil, zero value otherwise.

### GetSecureAccessSshCredsUserOk

`func (o *CreateSSHCertIssuer) GetSecureAccessSshCredsUserOk() (*string, bool)`

GetSecureAccessSshCredsUserOk returns a tuple with the SecureAccessSshCredsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessSshCredsUser

`func (o *CreateSSHCertIssuer) SetSecureAccessSshCredsUser(v string)`

SetSecureAccessSshCredsUser sets SecureAccessSshCredsUser field to given value.

### HasSecureAccessSshCredsUser

`func (o *CreateSSHCertIssuer) HasSecureAccessSshCredsUser() bool`

HasSecureAccessSshCredsUser returns a boolean if a field has been set.

### GetSecureAccessUseInternalBastion

`func (o *CreateSSHCertIssuer) GetSecureAccessUseInternalBastion() bool`

GetSecureAccessUseInternalBastion returns the SecureAccessUseInternalBastion field if non-nil, zero value otherwise.

### GetSecureAccessUseInternalBastionOk

`func (o *CreateSSHCertIssuer) GetSecureAccessUseInternalBastionOk() (*bool, bool)`

GetSecureAccessUseInternalBastionOk returns a tuple with the SecureAccessUseInternalBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUseInternalBastion

`func (o *CreateSSHCertIssuer) SetSecureAccessUseInternalBastion(v bool)`

SetSecureAccessUseInternalBastion sets SecureAccessUseInternalBastion field to given value.

### HasSecureAccessUseInternalBastion

`func (o *CreateSSHCertIssuer) HasSecureAccessUseInternalBastion() bool`

HasSecureAccessUseInternalBastion returns a boolean if a field has been set.

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


### GetTag

`func (o *CreateSSHCertIssuer) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateSSHCertIssuer) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateSSHCertIssuer) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateSSHCertIssuer) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTarget

`func (o *CreateSSHCertIssuer) GetTarget() []string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateSSHCertIssuer) GetTargetOk() (*[]string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateSSHCertIssuer) SetTarget(v []string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CreateSSHCertIssuer) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


