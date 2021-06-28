# UpdateRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsAccess** | Pointer to **string** | Allow this role to view analytics. Currently only &#39;none&#39;, &#39;own&#39;, &#39;all&#39; values are supported, allowing associated auth methods to view reports produced by the same auth methods. | [optional] 
**AuditAccess** | Pointer to **string** | Allow this role to view audit logs. Currently only &#39;none&#39;, &#39;own&#39; and &#39;all&#39; values are supported, allowing associated auth methods to view audit logs produced by the same auth methods. | [optional] 
**GwAnalyticsAccess** | Pointer to **string** | Allow this role to view gw analytics. Currently only &#39;none&#39;, &#39;own&#39;, &#39;all&#39; values are supported, allowing associated auth methods to view reports produced by the same auth methods. | [optional] 
**Name** | **string** | Role name | 
**NewComment** | Pointer to **string** | New comment about the role | [optional] [default to "default_comment"]
**NewName** | Pointer to **string** | New Role name | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewUpdateRole

`func NewUpdateRole(name string, ) *UpdateRole`

NewUpdateRole instantiates a new UpdateRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleWithDefaults

`func NewUpdateRoleWithDefaults() *UpdateRole`

NewUpdateRoleWithDefaults instantiates a new UpdateRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsAccess

`func (o *UpdateRole) GetAnalyticsAccess() string`

GetAnalyticsAccess returns the AnalyticsAccess field if non-nil, zero value otherwise.

### GetAnalyticsAccessOk

`func (o *UpdateRole) GetAnalyticsAccessOk() (*string, bool)`

GetAnalyticsAccessOk returns a tuple with the AnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsAccess

`func (o *UpdateRole) SetAnalyticsAccess(v string)`

SetAnalyticsAccess sets AnalyticsAccess field to given value.

### HasAnalyticsAccess

`func (o *UpdateRole) HasAnalyticsAccess() bool`

HasAnalyticsAccess returns a boolean if a field has been set.

### GetAuditAccess

`func (o *UpdateRole) GetAuditAccess() string`

GetAuditAccess returns the AuditAccess field if non-nil, zero value otherwise.

### GetAuditAccessOk

`func (o *UpdateRole) GetAuditAccessOk() (*string, bool)`

GetAuditAccessOk returns a tuple with the AuditAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditAccess

`func (o *UpdateRole) SetAuditAccess(v string)`

SetAuditAccess sets AuditAccess field to given value.

### HasAuditAccess

`func (o *UpdateRole) HasAuditAccess() bool`

HasAuditAccess returns a boolean if a field has been set.

### GetGwAnalyticsAccess

`func (o *UpdateRole) GetGwAnalyticsAccess() string`

GetGwAnalyticsAccess returns the GwAnalyticsAccess field if non-nil, zero value otherwise.

### GetGwAnalyticsAccessOk

`func (o *UpdateRole) GetGwAnalyticsAccessOk() (*string, bool)`

GetGwAnalyticsAccessOk returns a tuple with the GwAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwAnalyticsAccess

`func (o *UpdateRole) SetGwAnalyticsAccess(v string)`

SetGwAnalyticsAccess sets GwAnalyticsAccess field to given value.

### HasGwAnalyticsAccess

`func (o *UpdateRole) HasGwAnalyticsAccess() bool`

HasGwAnalyticsAccess returns a boolean if a field has been set.

### GetName

`func (o *UpdateRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRole) SetName(v string)`

SetName sets Name field to given value.


### GetNewComment

`func (o *UpdateRole) GetNewComment() string`

GetNewComment returns the NewComment field if non-nil, zero value otherwise.

### GetNewCommentOk

`func (o *UpdateRole) GetNewCommentOk() (*string, bool)`

GetNewCommentOk returns a tuple with the NewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewComment

`func (o *UpdateRole) SetNewComment(v string)`

SetNewComment sets NewComment field to given value.

### HasNewComment

`func (o *UpdateRole) HasNewComment() bool`

HasNewComment returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateRole) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateRole) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateRole) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateRole) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateRole) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateRole) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateRole) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateRole) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *UpdateRole) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateRole) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateRole) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateRole) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateRole) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateRole) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateRole) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateRole) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateRole) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateRole) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateRole) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateRole) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


