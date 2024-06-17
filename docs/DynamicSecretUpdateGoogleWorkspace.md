# DynamicSecretUpdateGoogleWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | **string** |  | 
**AdminEmail** | **string** | Admin user email | 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GcpKey** | Pointer to **string** | Base64-encoded service account private key text | [optional] 
**GroupEmail** | Pointer to **string** | A group email, relevant only for group access-mode | [optional] 
**GroupRole** | Pointer to **string** |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RoleName** | Pointer to **string** | Name of the admin role to assign to the user, relevant only for role access-mode | [optional] 
**RoleScope** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Name of existing target to use in dynamic secret creation | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretUpdateGoogleWorkspace

`func NewDynamicSecretUpdateGoogleWorkspace(accessMode string, adminEmail string, name string, ) *DynamicSecretUpdateGoogleWorkspace`

NewDynamicSecretUpdateGoogleWorkspace instantiates a new DynamicSecretUpdateGoogleWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateGoogleWorkspaceWithDefaults

`func NewDynamicSecretUpdateGoogleWorkspaceWithDefaults() *DynamicSecretUpdateGoogleWorkspace`

NewDynamicSecretUpdateGoogleWorkspaceWithDefaults instantiates a new DynamicSecretUpdateGoogleWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *DynamicSecretUpdateGoogleWorkspace) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *DynamicSecretUpdateGoogleWorkspace) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.


### GetAdminEmail

`func (o *DynamicSecretUpdateGoogleWorkspace) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *DynamicSecretUpdateGoogleWorkspace) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.


### GetDeleteProtection

`func (o *DynamicSecretUpdateGoogleWorkspace) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateGoogleWorkspace) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateGoogleWorkspace) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateGoogleWorkspace) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateGoogleWorkspace) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateGoogleWorkspace) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGcpKey

`func (o *DynamicSecretUpdateGoogleWorkspace) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *DynamicSecretUpdateGoogleWorkspace) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *DynamicSecretUpdateGoogleWorkspace) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetGroupEmail

`func (o *DynamicSecretUpdateGoogleWorkspace) GetGroupEmail() string`

GetGroupEmail returns the GroupEmail field if non-nil, zero value otherwise.

### GetGroupEmailOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetGroupEmailOk() (*string, bool)`

GetGroupEmailOk returns a tuple with the GroupEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEmail

`func (o *DynamicSecretUpdateGoogleWorkspace) SetGroupEmail(v string)`

SetGroupEmail sets GroupEmail field to given value.

### HasGroupEmail

`func (o *DynamicSecretUpdateGoogleWorkspace) HasGroupEmail() bool`

HasGroupEmail returns a boolean if a field has been set.

### GetGroupRole

`func (o *DynamicSecretUpdateGoogleWorkspace) GetGroupRole() string`

GetGroupRole returns the GroupRole field if non-nil, zero value otherwise.

### GetGroupRoleOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetGroupRoleOk() (*string, bool)`

GetGroupRoleOk returns a tuple with the GroupRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRole

`func (o *DynamicSecretUpdateGoogleWorkspace) SetGroupRole(v string)`

SetGroupRole sets GroupRole field to given value.

### HasGroupRole

`func (o *DynamicSecretUpdateGoogleWorkspace) HasGroupRole() bool`

HasGroupRole returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateGoogleWorkspace) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateGoogleWorkspace) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateGoogleWorkspace) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateGoogleWorkspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateGoogleWorkspace) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateGoogleWorkspace) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateGoogleWorkspace) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateGoogleWorkspace) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateGoogleWorkspace) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateGoogleWorkspace) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateGoogleWorkspace) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRoleName

`func (o *DynamicSecretUpdateGoogleWorkspace) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *DynamicSecretUpdateGoogleWorkspace) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *DynamicSecretUpdateGoogleWorkspace) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRoleScope

`func (o *DynamicSecretUpdateGoogleWorkspace) GetRoleScope() string`

GetRoleScope returns the RoleScope field if non-nil, zero value otherwise.

### GetRoleScopeOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetRoleScopeOk() (*string, bool)`

GetRoleScopeOk returns a tuple with the RoleScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleScope

`func (o *DynamicSecretUpdateGoogleWorkspace) SetRoleScope(v string)`

SetRoleScope sets RoleScope field to given value.

### HasRoleScope

`func (o *DynamicSecretUpdateGoogleWorkspace) HasRoleScope() bool`

HasRoleScope returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateGoogleWorkspace) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateGoogleWorkspace) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateGoogleWorkspace) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateGoogleWorkspace) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateGoogleWorkspace) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateGoogleWorkspace) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateGoogleWorkspace) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateGoogleWorkspace) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateGoogleWorkspace) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateGoogleWorkspace) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateGoogleWorkspace) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateGoogleWorkspace) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateGoogleWorkspace) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateGoogleWorkspace) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateGoogleWorkspace) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateGoogleWorkspace) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


