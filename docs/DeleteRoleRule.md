# DeleteRoleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Path** | **string** | The path the rule refers to | 
**RoleName** | **string** | The role name to be updated | 
**RuleType** | Pointer to **string** | item-rule, role-rule, auth-method-rule, search-rule or reports-rule | [optional] [default to "item-rule"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewDeleteRoleRule

`func NewDeleteRoleRule(path string, roleName string, ) *DeleteRoleRule`

NewDeleteRoleRule instantiates a new DeleteRoleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRoleRuleWithDefaults

`func NewDeleteRoleRuleWithDefaults() *DeleteRoleRule`

NewDeleteRoleRuleWithDefaults instantiates a new DeleteRoleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *DeleteRoleRule) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeleteRoleRule) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeleteRoleRule) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DeleteRoleRule) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *DeleteRoleRule) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeleteRoleRule) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeleteRoleRule) SetPath(v string)`

SetPath sets Path field to given value.


### GetRoleName

`func (o *DeleteRoleRule) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *DeleteRoleRule) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *DeleteRoleRule) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetRuleType

`func (o *DeleteRoleRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *DeleteRoleRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *DeleteRoleRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *DeleteRoleRule) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetToken

`func (o *DeleteRoleRule) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteRoleRule) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteRoleRule) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeleteRoleRule) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeleteRoleRule) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeleteRoleRule) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeleteRoleRule) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeleteRoleRule) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *DeleteRoleRule) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeleteRoleRule) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeleteRoleRule) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DeleteRoleRule) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


