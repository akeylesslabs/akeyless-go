# OidcClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPermissionAssignment** | Pointer to [**[]AccessPermissionAssignment**](AccessPermissionAssignment.md) |  | [optional] 
**Audience** | Pointer to **[]string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**LogoutUris** | Pointer to **[]string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**ResponseTypes** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOidcClientInfo

`func NewOidcClientInfo() *OidcClientInfo`

NewOidcClientInfo instantiates a new OidcClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcClientInfoWithDefaults

`func NewOidcClientInfoWithDefaults() *OidcClientInfo`

NewOidcClientInfoWithDefaults instantiates a new OidcClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPermissionAssignment

`func (o *OidcClientInfo) GetAccessPermissionAssignment() []AccessPermissionAssignment`

GetAccessPermissionAssignment returns the AccessPermissionAssignment field if non-nil, zero value otherwise.

### GetAccessPermissionAssignmentOk

`func (o *OidcClientInfo) GetAccessPermissionAssignmentOk() (*[]AccessPermissionAssignment, bool)`

GetAccessPermissionAssignmentOk returns a tuple with the AccessPermissionAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPermissionAssignment

`func (o *OidcClientInfo) SetAccessPermissionAssignment(v []AccessPermissionAssignment)`

SetAccessPermissionAssignment sets AccessPermissionAssignment field to given value.

### HasAccessPermissionAssignment

`func (o *OidcClientInfo) HasAccessPermissionAssignment() bool`

HasAccessPermissionAssignment returns a boolean if a field has been set.

### GetAudience

`func (o *OidcClientInfo) GetAudience() []string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *OidcClientInfo) GetAudienceOk() (*[]string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *OidcClientInfo) SetAudience(v []string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *OidcClientInfo) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetClientId

`func (o *OidcClientInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OidcClientInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OidcClientInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OidcClientInfo) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetGrantTypes

`func (o *OidcClientInfo) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *OidcClientInfo) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *OidcClientInfo) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *OidcClientInfo) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetLogoutUris

`func (o *OidcClientInfo) GetLogoutUris() []string`

GetLogoutUris returns the LogoutUris field if non-nil, zero value otherwise.

### GetLogoutUrisOk

`func (o *OidcClientInfo) GetLogoutUrisOk() (*[]string, bool)`

GetLogoutUrisOk returns a tuple with the LogoutUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUris

`func (o *OidcClientInfo) SetLogoutUris(v []string)`

SetLogoutUris sets LogoutUris field to given value.

### HasLogoutUris

`func (o *OidcClientInfo) HasLogoutUris() bool`

HasLogoutUris returns a boolean if a field has been set.

### GetPublic

`func (o *OidcClientInfo) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *OidcClientInfo) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *OidcClientInfo) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *OidcClientInfo) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRedirectUris

`func (o *OidcClientInfo) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OidcClientInfo) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OidcClientInfo) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *OidcClientInfo) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetResponseTypes

`func (o *OidcClientInfo) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *OidcClientInfo) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *OidcClientInfo) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *OidcClientInfo) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetScopes

`func (o *OidcClientInfo) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OidcClientInfo) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OidcClientInfo) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OidcClientInfo) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


