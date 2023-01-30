# AWSIAMAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **[]string** | The list of account ids that the login is restricted to. | [optional] 
**Arn** | Pointer to **[]string** | The list of ARNs that the login is restricted to. | [optional] 
**ResourceId** | Pointer to **[]string** | The list of resource ids that the login is restricted to. | [optional] 
**RoleId** | Pointer to **[]string** | The list of role ids that the login is restricted to. | [optional] 
**RoleName** | Pointer to **[]string** | The list of role names that the login is restricted to. | [optional] 
**StsEndpoint** | Pointer to **string** | The sts URL. | [optional] 
**UserId** | Pointer to **[]string** | The list of user ids that the login is restricted to. | [optional] 
**UserName** | Pointer to **[]string** | The list of user names that the login is restricted to. | [optional] 

## Methods

### NewAWSIAMAccessRules

`func NewAWSIAMAccessRules() *AWSIAMAccessRules`

NewAWSIAMAccessRules instantiates a new AWSIAMAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSIAMAccessRulesWithDefaults

`func NewAWSIAMAccessRulesWithDefaults() *AWSIAMAccessRules`

NewAWSIAMAccessRulesWithDefaults instantiates a new AWSIAMAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AWSIAMAccessRules) GetAccountId() []string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSIAMAccessRules) GetAccountIdOk() (*[]string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSIAMAccessRules) SetAccountId(v []string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AWSIAMAccessRules) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetArn

`func (o *AWSIAMAccessRules) GetArn() []string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *AWSIAMAccessRules) GetArnOk() (*[]string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *AWSIAMAccessRules) SetArn(v []string)`

SetArn sets Arn field to given value.

### HasArn

`func (o *AWSIAMAccessRules) HasArn() bool`

HasArn returns a boolean if a field has been set.

### GetResourceId

`func (o *AWSIAMAccessRules) GetResourceId() []string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AWSIAMAccessRules) GetResourceIdOk() (*[]string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AWSIAMAccessRules) SetResourceId(v []string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AWSIAMAccessRules) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetRoleId

`func (o *AWSIAMAccessRules) GetRoleId() []string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AWSIAMAccessRules) GetRoleIdOk() (*[]string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AWSIAMAccessRules) SetRoleId(v []string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *AWSIAMAccessRules) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *AWSIAMAccessRules) GetRoleName() []string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AWSIAMAccessRules) GetRoleNameOk() (*[]string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AWSIAMAccessRules) SetRoleName(v []string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AWSIAMAccessRules) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetStsEndpoint

`func (o *AWSIAMAccessRules) GetStsEndpoint() string`

GetStsEndpoint returns the StsEndpoint field if non-nil, zero value otherwise.

### GetStsEndpointOk

`func (o *AWSIAMAccessRules) GetStsEndpointOk() (*string, bool)`

GetStsEndpointOk returns a tuple with the StsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsEndpoint

`func (o *AWSIAMAccessRules) SetStsEndpoint(v string)`

SetStsEndpoint sets StsEndpoint field to given value.

### HasStsEndpoint

`func (o *AWSIAMAccessRules) HasStsEndpoint() bool`

HasStsEndpoint returns a boolean if a field has been set.

### GetUserId

`func (o *AWSIAMAccessRules) GetUserId() []string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AWSIAMAccessRules) GetUserIdOk() (*[]string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AWSIAMAccessRules) SetUserId(v []string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AWSIAMAccessRules) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *AWSIAMAccessRules) GetUserName() []string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AWSIAMAccessRules) GetUserNameOk() (*[]string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AWSIAMAccessRules) SetUserName(v []string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AWSIAMAccessRules) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


