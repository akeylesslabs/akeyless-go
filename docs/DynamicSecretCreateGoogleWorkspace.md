# DynamicSecretCreateGoogleWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | **string** |  | 
**AdminEmail** | **string** | Admin user email | 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**FixedUserClaimKeyname** | Pointer to **string** | For externally provided users, denotes the key-name of IdP claim to extract the username from | [optional] [default to "ext_email"]
**GcpKey** | Pointer to **string** | Base64-encoded service account private key text | [optional] 
**GroupEmail** | Pointer to **string** | A group email, relevant only for group access-mode | [optional] 
**GroupRole** | Pointer to **string** |  | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RoleName** | Pointer to **string** | Name of the admin role to assign to the user, relevant only for role access-mode | [optional] 
**RoleScope** | Pointer to **string** |  | [optional] 
**SecureAccessDelay** | Pointer to **int64** | The delay duration, in seconds, to wait after generating just-in-time credentials. Accepted range: 0-120 seconds | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessUrl** | Pointer to **string** | Destination URL to inject secrets | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to true]
**SecureAccessWebBrowsing** | Pointer to **bool** | Secure browser via Akeyless&#39;s Secure Remote Access (SRA) | [optional] [default to false]
**SecureAccessWebProxy** | Pointer to **bool** | Web-Proxy via Akeyless&#39;s Secure Remote Access (SRA) | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Name of existing target to use in dynamic secret creation | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateGoogleWorkspace

`func NewDynamicSecretCreateGoogleWorkspace(accessMode string, adminEmail string, name string, ) *DynamicSecretCreateGoogleWorkspace`

NewDynamicSecretCreateGoogleWorkspace instantiates a new DynamicSecretCreateGoogleWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateGoogleWorkspaceWithDefaults

`func NewDynamicSecretCreateGoogleWorkspaceWithDefaults() *DynamicSecretCreateGoogleWorkspace`

NewDynamicSecretCreateGoogleWorkspaceWithDefaults instantiates a new DynamicSecretCreateGoogleWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *DynamicSecretCreateGoogleWorkspace) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *DynamicSecretCreateGoogleWorkspace) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.


### GetAdminEmail

`func (o *DynamicSecretCreateGoogleWorkspace) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *DynamicSecretCreateGoogleWorkspace) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.


### GetDeleteProtection

`func (o *DynamicSecretCreateGoogleWorkspace) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateGoogleWorkspace) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateGoogleWorkspace) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateGoogleWorkspace) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateGoogleWorkspace) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateGoogleWorkspace) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixedUserClaimKeyname

`func (o *DynamicSecretCreateGoogleWorkspace) GetFixedUserClaimKeyname() string`

GetFixedUserClaimKeyname returns the FixedUserClaimKeyname field if non-nil, zero value otherwise.

### GetFixedUserClaimKeynameOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetFixedUserClaimKeynameOk() (*string, bool)`

GetFixedUserClaimKeynameOk returns a tuple with the FixedUserClaimKeyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserClaimKeyname

`func (o *DynamicSecretCreateGoogleWorkspace) SetFixedUserClaimKeyname(v string)`

SetFixedUserClaimKeyname sets FixedUserClaimKeyname field to given value.

### HasFixedUserClaimKeyname

`func (o *DynamicSecretCreateGoogleWorkspace) HasFixedUserClaimKeyname() bool`

HasFixedUserClaimKeyname returns a boolean if a field has been set.

### GetGcpKey

`func (o *DynamicSecretCreateGoogleWorkspace) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *DynamicSecretCreateGoogleWorkspace) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *DynamicSecretCreateGoogleWorkspace) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetGroupEmail

`func (o *DynamicSecretCreateGoogleWorkspace) GetGroupEmail() string`

GetGroupEmail returns the GroupEmail field if non-nil, zero value otherwise.

### GetGroupEmailOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetGroupEmailOk() (*string, bool)`

GetGroupEmailOk returns a tuple with the GroupEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEmail

`func (o *DynamicSecretCreateGoogleWorkspace) SetGroupEmail(v string)`

SetGroupEmail sets GroupEmail field to given value.

### HasGroupEmail

`func (o *DynamicSecretCreateGoogleWorkspace) HasGroupEmail() bool`

HasGroupEmail returns a boolean if a field has been set.

### GetGroupRole

`func (o *DynamicSecretCreateGoogleWorkspace) GetGroupRole() string`

GetGroupRole returns the GroupRole field if non-nil, zero value otherwise.

### GetGroupRoleOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetGroupRoleOk() (*string, bool)`

GetGroupRoleOk returns a tuple with the GroupRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRole

`func (o *DynamicSecretCreateGoogleWorkspace) SetGroupRole(v string)`

SetGroupRole sets GroupRole field to given value.

### HasGroupRole

`func (o *DynamicSecretCreateGoogleWorkspace) HasGroupRole() bool`

HasGroupRole returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *DynamicSecretCreateGoogleWorkspace) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *DynamicSecretCreateGoogleWorkspace) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *DynamicSecretCreateGoogleWorkspace) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateGoogleWorkspace) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateGoogleWorkspace) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateGoogleWorkspace) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateGoogleWorkspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateGoogleWorkspace) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateGoogleWorkspace) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateGoogleWorkspace) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateGoogleWorkspace) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRoleName

`func (o *DynamicSecretCreateGoogleWorkspace) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *DynamicSecretCreateGoogleWorkspace) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *DynamicSecretCreateGoogleWorkspace) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRoleScope

`func (o *DynamicSecretCreateGoogleWorkspace) GetRoleScope() string`

GetRoleScope returns the RoleScope field if non-nil, zero value otherwise.

### GetRoleScopeOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetRoleScopeOk() (*string, bool)`

GetRoleScopeOk returns a tuple with the RoleScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleScope

`func (o *DynamicSecretCreateGoogleWorkspace) SetRoleScope(v string)`

SetRoleScope sets RoleScope field to given value.

### HasRoleScope

`func (o *DynamicSecretCreateGoogleWorkspace) HasRoleScope() bool`

HasRoleScope returns a boolean if a field has been set.

### GetSecureAccessDelay

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessDelay() int64`

GetSecureAccessDelay returns the SecureAccessDelay field if non-nil, zero value otherwise.

### GetSecureAccessDelayOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessDelayOk() (*int64, bool)`

GetSecureAccessDelayOk returns a tuple with the SecureAccessDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDelay

`func (o *DynamicSecretCreateGoogleWorkspace) SetSecureAccessDelay(v int64)`

SetSecureAccessDelay sets SecureAccessDelay field to given value.

### HasSecureAccessDelay

`func (o *DynamicSecretCreateGoogleWorkspace) HasSecureAccessDelay() bool`

HasSecureAccessDelay returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateGoogleWorkspace) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateGoogleWorkspace) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessUrl

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessUrl() string`

GetSecureAccessUrl returns the SecureAccessUrl field if non-nil, zero value otherwise.

### GetSecureAccessUrlOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessUrlOk() (*string, bool)`

GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUrl

`func (o *DynamicSecretCreateGoogleWorkspace) SetSecureAccessUrl(v string)`

SetSecureAccessUrl sets SecureAccessUrl field to given value.

### HasSecureAccessUrl

`func (o *DynamicSecretCreateGoogleWorkspace) HasSecureAccessUrl() bool`

HasSecureAccessUrl returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretCreateGoogleWorkspace) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretCreateGoogleWorkspace) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *DynamicSecretCreateGoogleWorkspace) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *DynamicSecretCreateGoogleWorkspace) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *DynamicSecretCreateGoogleWorkspace) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *DynamicSecretCreateGoogleWorkspace) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateGoogleWorkspace) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateGoogleWorkspace) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateGoogleWorkspace) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateGoogleWorkspace) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateGoogleWorkspace) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateGoogleWorkspace) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateGoogleWorkspace) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateGoogleWorkspace) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateGoogleWorkspace) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateGoogleWorkspace) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateGoogleWorkspace) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateGoogleWorkspace) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateGoogleWorkspace) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateGoogleWorkspace) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateGoogleWorkspace) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateGoogleWorkspace) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


