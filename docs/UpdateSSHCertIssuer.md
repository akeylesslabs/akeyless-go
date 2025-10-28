# UpdateSSHCertIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | Pointer to **string** |  | [optional] 
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**AllowedUsers** | **string** | Users allowed to fetch the certificate, e.g root,ubuntu | [default to "-"]
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Extensions** | Pointer to **map[string]string** | Signed certificates with extensions, e.g permit-port-forwarding&#x3D;\\\&quot;\\\&quot; | [optional] 
**ExternalUsername** | Pointer to **string** | Externally provided username [true/false] | [optional] [default to "false"]
**FixedUserClaimKeyname** | Pointer to **string** | For externally provided users, denotes the key-name of IdP claim to extract the username from (relevant only for external-username&#x3D;true) | [optional] 
**HostProvider** | Pointer to **string** | Host provider type [explicit/target], Default Host provider is explicit, Relevant only for Secure Remote Access of ssh cert issuer, ldap rotated secret and ldap dynamic secret | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | SSH certificate issuer name | 
**NewName** | Pointer to **string** | New item name | [optional] 
**Principals** | Pointer to **string** | Signed certificates with principal, e.g example_role1,example_role2 | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**SecureAccessApi** | Pointer to **string** | Secure Access SSH control API endpoint. E.g. https://my.sra-server:9900 | [optional] 
**SecureAccessBastionApi** | Pointer to **string** | Deprecated. use secure-access-api | [optional] 
**SecureAccessBastionSsh** | Pointer to **string** | Deprecated. use secure-access-ssh | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessEnforceHostsRestriction** | Pointer to **bool** | Enable this flag to enforce connections only to the hosts listed in --secure-access-host | [optional] 
**SecureAccessGateway** | Pointer to **string** |  | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessSsh** | Pointer to **string** | Bastion&#39;s SSH server. E.g. my.sra-server:22 | [optional] 
**SecureAccessSshCredsUser** | Pointer to **string** | SSH username to connect to target server, must be in &#39;Allowed Users&#39; list | [optional] 
**SecureAccessUseInternalBastion** | Pointer to **bool** | Deprecated. Use secure-access-use-internal-ssh-access | [optional] 
**SecureAccessUseInternalSshAccess** | Pointer to **bool** | Use internal SSH Access | [optional] 
**SignerKeyName** | **string** | A key to sign the certificate with | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | **int64** | The requested Time To Live for the certificate, in seconds | 
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

### GetProviderType

`func (o *UpdateSSHCertIssuer) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *UpdateSSHCertIssuer) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *UpdateSSHCertIssuer) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *UpdateSSHCertIssuer) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

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


### GetDeleteProtection

`func (o *UpdateSSHCertIssuer) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *UpdateSSHCertIssuer) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *UpdateSSHCertIssuer) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *UpdateSSHCertIssuer) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSSHCertIssuer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSSHCertIssuer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSSHCertIssuer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSSHCertIssuer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### GetExternalUsername

`func (o *UpdateSSHCertIssuer) GetExternalUsername() string`

GetExternalUsername returns the ExternalUsername field if non-nil, zero value otherwise.

### GetExternalUsernameOk

`func (o *UpdateSSHCertIssuer) GetExternalUsernameOk() (*string, bool)`

GetExternalUsernameOk returns a tuple with the ExternalUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUsername

`func (o *UpdateSSHCertIssuer) SetExternalUsername(v string)`

SetExternalUsername sets ExternalUsername field to given value.

### HasExternalUsername

`func (o *UpdateSSHCertIssuer) HasExternalUsername() bool`

HasExternalUsername returns a boolean if a field has been set.

### GetFixedUserClaimKeyname

`func (o *UpdateSSHCertIssuer) GetFixedUserClaimKeyname() string`

GetFixedUserClaimKeyname returns the FixedUserClaimKeyname field if non-nil, zero value otherwise.

### GetFixedUserClaimKeynameOk

`func (o *UpdateSSHCertIssuer) GetFixedUserClaimKeynameOk() (*string, bool)`

GetFixedUserClaimKeynameOk returns a tuple with the FixedUserClaimKeyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserClaimKeyname

`func (o *UpdateSSHCertIssuer) SetFixedUserClaimKeyname(v string)`

SetFixedUserClaimKeyname sets FixedUserClaimKeyname field to given value.

### HasFixedUserClaimKeyname

`func (o *UpdateSSHCertIssuer) HasFixedUserClaimKeyname() bool`

HasFixedUserClaimKeyname returns a boolean if a field has been set.

### GetHostProvider

`func (o *UpdateSSHCertIssuer) GetHostProvider() string`

GetHostProvider returns the HostProvider field if non-nil, zero value otherwise.

### GetHostProviderOk

`func (o *UpdateSSHCertIssuer) GetHostProviderOk() (*string, bool)`

GetHostProviderOk returns a tuple with the HostProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostProvider

`func (o *UpdateSSHCertIssuer) SetHostProvider(v string)`

SetHostProvider sets HostProvider field to given value.

### HasHostProvider

`func (o *UpdateSSHCertIssuer) HasHostProvider() bool`

HasHostProvider returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *UpdateSSHCertIssuer) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *UpdateSSHCertIssuer) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *UpdateSSHCertIssuer) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *UpdateSSHCertIssuer) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *UpdateSSHCertIssuer) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateSSHCertIssuer) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateSSHCertIssuer) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateSSHCertIssuer) HasJson() bool`

HasJson returns a boolean if a field has been set.

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

### GetSecureAccessApi

`func (o *UpdateSSHCertIssuer) GetSecureAccessApi() string`

GetSecureAccessApi returns the SecureAccessApi field if non-nil, zero value otherwise.

### GetSecureAccessApiOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessApiOk() (*string, bool)`

GetSecureAccessApiOk returns a tuple with the SecureAccessApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessApi

`func (o *UpdateSSHCertIssuer) SetSecureAccessApi(v string)`

SetSecureAccessApi sets SecureAccessApi field to given value.

### HasSecureAccessApi

`func (o *UpdateSSHCertIssuer) HasSecureAccessApi() bool`

HasSecureAccessApi returns a boolean if a field has been set.

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

### GetSecureAccessEnforceHostsRestriction

`func (o *UpdateSSHCertIssuer) GetSecureAccessEnforceHostsRestriction() bool`

GetSecureAccessEnforceHostsRestriction returns the SecureAccessEnforceHostsRestriction field if non-nil, zero value otherwise.

### GetSecureAccessEnforceHostsRestrictionOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessEnforceHostsRestrictionOk() (*bool, bool)`

GetSecureAccessEnforceHostsRestrictionOk returns a tuple with the SecureAccessEnforceHostsRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnforceHostsRestriction

`func (o *UpdateSSHCertIssuer) SetSecureAccessEnforceHostsRestriction(v bool)`

SetSecureAccessEnforceHostsRestriction sets SecureAccessEnforceHostsRestriction field to given value.

### HasSecureAccessEnforceHostsRestriction

`func (o *UpdateSSHCertIssuer) HasSecureAccessEnforceHostsRestriction() bool`

HasSecureAccessEnforceHostsRestriction returns a boolean if a field has been set.

### GetSecureAccessGateway

`func (o *UpdateSSHCertIssuer) GetSecureAccessGateway() string`

GetSecureAccessGateway returns the SecureAccessGateway field if non-nil, zero value otherwise.

### GetSecureAccessGatewayOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessGatewayOk() (*string, bool)`

GetSecureAccessGatewayOk returns a tuple with the SecureAccessGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessGateway

`func (o *UpdateSSHCertIssuer) SetSecureAccessGateway(v string)`

SetSecureAccessGateway sets SecureAccessGateway field to given value.

### HasSecureAccessGateway

`func (o *UpdateSSHCertIssuer) HasSecureAccessGateway() bool`

HasSecureAccessGateway returns a boolean if a field has been set.

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

### GetSecureAccessSsh

`func (o *UpdateSSHCertIssuer) GetSecureAccessSsh() string`

GetSecureAccessSsh returns the SecureAccessSsh field if non-nil, zero value otherwise.

### GetSecureAccessSshOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessSshOk() (*string, bool)`

GetSecureAccessSshOk returns a tuple with the SecureAccessSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessSsh

`func (o *UpdateSSHCertIssuer) SetSecureAccessSsh(v string)`

SetSecureAccessSsh sets SecureAccessSsh field to given value.

### HasSecureAccessSsh

`func (o *UpdateSSHCertIssuer) HasSecureAccessSsh() bool`

HasSecureAccessSsh returns a boolean if a field has been set.

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

### GetSecureAccessUseInternalSshAccess

`func (o *UpdateSSHCertIssuer) GetSecureAccessUseInternalSshAccess() bool`

GetSecureAccessUseInternalSshAccess returns the SecureAccessUseInternalSshAccess field if non-nil, zero value otherwise.

### GetSecureAccessUseInternalSshAccessOk

`func (o *UpdateSSHCertIssuer) GetSecureAccessUseInternalSshAccessOk() (*bool, bool)`

GetSecureAccessUseInternalSshAccessOk returns a tuple with the SecureAccessUseInternalSshAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUseInternalSshAccess

`func (o *UpdateSSHCertIssuer) SetSecureAccessUseInternalSshAccess(v bool)`

SetSecureAccessUseInternalSshAccess sets SecureAccessUseInternalSshAccess field to given value.

### HasSecureAccessUseInternalSshAccess

`func (o *UpdateSSHCertIssuer) HasSecureAccessUseInternalSshAccess() bool`

HasSecureAccessUseInternalSshAccess returns a boolean if a field has been set.

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


