# CreateOidcApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**Audience** | Pointer to **string** | A comma separated list of allowed audiences | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the OIDC application (if empty, the account default protectionKey key will be used) | [optional] 
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | OIDC application name | 
**PermissionAssignment** | Pointer to **string** | A json string defining the access permission assignment for this app | [optional] 
**Public** | Pointer to **bool** | Set to true if the app is public (cannot keep secrets) | [optional] 
**RedirectUris** | Pointer to **string** | A comma separated list of allowed redirect uris | [optional] 
**Scopes** | Pointer to **string** | A comma separated list of allowed scopes | [optional] [default to "openid"]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateOidcApp

`func NewCreateOidcApp(name string, ) *CreateOidcApp`

NewCreateOidcApp instantiates a new CreateOidcApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOidcAppWithDefaults

`func NewCreateOidcAppWithDefaults() *CreateOidcApp`

NewCreateOidcAppWithDefaults instantiates a new CreateOidcApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *CreateOidcApp) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *CreateOidcApp) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *CreateOidcApp) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *CreateOidcApp) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetAudience

`func (o *CreateOidcApp) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CreateOidcApp) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CreateOidcApp) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CreateOidcApp) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateOidcApp) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateOidcApp) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateOidcApp) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateOidcApp) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOidcApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOidcApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOidcApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOidcApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *CreateOidcApp) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateOidcApp) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateOidcApp) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateOidcApp) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *CreateOidcApp) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateOidcApp) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateOidcApp) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateOidcApp) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateOidcApp) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateOidcApp) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateOidcApp) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateOidcApp) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateOidcApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOidcApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOidcApp) SetName(v string)`

SetName sets Name field to given value.


### GetPermissionAssignment

`func (o *CreateOidcApp) GetPermissionAssignment() string`

GetPermissionAssignment returns the PermissionAssignment field if non-nil, zero value otherwise.

### GetPermissionAssignmentOk

`func (o *CreateOidcApp) GetPermissionAssignmentOk() (*string, bool)`

GetPermissionAssignmentOk returns a tuple with the PermissionAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionAssignment

`func (o *CreateOidcApp) SetPermissionAssignment(v string)`

SetPermissionAssignment sets PermissionAssignment field to given value.

### HasPermissionAssignment

`func (o *CreateOidcApp) HasPermissionAssignment() bool`

HasPermissionAssignment returns a boolean if a field has been set.

### GetPublic

`func (o *CreateOidcApp) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CreateOidcApp) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CreateOidcApp) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CreateOidcApp) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRedirectUris

`func (o *CreateOidcApp) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *CreateOidcApp) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *CreateOidcApp) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *CreateOidcApp) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *CreateOidcApp) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateOidcApp) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateOidcApp) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateOidcApp) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTags

`func (o *CreateOidcApp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateOidcApp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateOidcApp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateOidcApp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *CreateOidcApp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateOidcApp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateOidcApp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateOidcApp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateOidcApp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateOidcApp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateOidcApp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateOidcApp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


