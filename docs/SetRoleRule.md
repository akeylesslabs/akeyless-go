# SetRoleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **[]string** | List of the approved/denied capabilities in the path options: [read, create, update, delete, list, deny] | 
**Path** | **string** | The path the rule refers to | 
**RoleName** | **string** | The role name to be updated | 
**RuleType** | Pointer to **string** | item-rule, target-rule, role-rule, auth-method-rule, search-rule, reports-rule, gw-reports-rule or sra-reports-rule | [optional] [default to "item-rule"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewSetRoleRule

`func NewSetRoleRule(capability []string, path string, roleName string, ) *SetRoleRule`

NewSetRoleRule instantiates a new SetRoleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRoleRuleWithDefaults

`func NewSetRoleRuleWithDefaults() *SetRoleRule`

NewSetRoleRuleWithDefaults instantiates a new SetRoleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *SetRoleRule) GetCapability() []string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *SetRoleRule) GetCapabilityOk() (*[]string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *SetRoleRule) SetCapability(v []string)`

SetCapability sets Capability field to given value.


### GetPath

`func (o *SetRoleRule) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SetRoleRule) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SetRoleRule) SetPath(v string)`

SetPath sets Path field to given value.


### GetRoleName

`func (o *SetRoleRule) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *SetRoleRule) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *SetRoleRule) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetRuleType

`func (o *SetRoleRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *SetRoleRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *SetRoleRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *SetRoleRule) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetToken

`func (o *SetRoleRule) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SetRoleRule) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SetRoleRule) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SetRoleRule) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *SetRoleRule) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *SetRoleRule) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *SetRoleRule) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *SetRoleRule) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


