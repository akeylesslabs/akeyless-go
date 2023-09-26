# UpdateOidcApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | A comma separated list of allowed audiences | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the OIDC application (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | OIDC application name | 
**PermissionAssignment** | Pointer to **string** | A json string defining the access permission assignment for this app | [optional] 
**Public** | Pointer to **bool** | Set to true if the app is public (cannot keep secrets) | [optional] 
**RedirectUris** | Pointer to **string** | A comma separated list of allowed redirect uris | [optional] 
**Scopes** | Pointer to **string** | A comma separated list of allowed scopes | [optional] [default to "openid"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateOidcApp

`func NewUpdateOidcApp(name string, ) *UpdateOidcApp`

NewUpdateOidcApp instantiates a new UpdateOidcApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOidcAppWithDefaults

`func NewUpdateOidcAppWithDefaults() *UpdateOidcApp`

NewUpdateOidcAppWithDefaults instantiates a new UpdateOidcApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *UpdateOidcApp) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *UpdateOidcApp) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *UpdateOidcApp) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *UpdateOidcApp) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetJson

`func (o *UpdateOidcApp) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateOidcApp) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateOidcApp) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateOidcApp) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *UpdateOidcApp) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateOidcApp) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateOidcApp) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateOidcApp) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateOidcApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOidcApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOidcApp) SetName(v string)`

SetName sets Name field to given value.


### GetPermissionAssignment

`func (o *UpdateOidcApp) GetPermissionAssignment() string`

GetPermissionAssignment returns the PermissionAssignment field if non-nil, zero value otherwise.

### GetPermissionAssignmentOk

`func (o *UpdateOidcApp) GetPermissionAssignmentOk() (*string, bool)`

GetPermissionAssignmentOk returns a tuple with the PermissionAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionAssignment

`func (o *UpdateOidcApp) SetPermissionAssignment(v string)`

SetPermissionAssignment sets PermissionAssignment field to given value.

### HasPermissionAssignment

`func (o *UpdateOidcApp) HasPermissionAssignment() bool`

HasPermissionAssignment returns a boolean if a field has been set.

### GetPublic

`func (o *UpdateOidcApp) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *UpdateOidcApp) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *UpdateOidcApp) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *UpdateOidcApp) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRedirectUris

`func (o *UpdateOidcApp) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *UpdateOidcApp) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *UpdateOidcApp) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *UpdateOidcApp) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *UpdateOidcApp) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateOidcApp) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateOidcApp) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateOidcApp) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetToken

`func (o *UpdateOidcApp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateOidcApp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateOidcApp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateOidcApp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateOidcApp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateOidcApp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateOidcApp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateOidcApp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


