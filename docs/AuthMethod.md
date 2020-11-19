# AuthMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessInfo** | Pointer to [**AuthMethodAccessInfo**](AuthMethodAccessInfo.md) |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AuthMethodAccessId** | Pointer to **string** |  | [optional] 
**AuthMethodName** | Pointer to **string** |  | [optional] 
**AuthMethodRolesAssoc** | Pointer to [**[]AuthMethodRoleAssociation**](AuthMethodRoleAssociation.md) |  | [optional] 
**ClientPermissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAuthMethod

`func NewAuthMethod() *AuthMethod`

NewAuthMethod instantiates a new AuthMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodWithDefaults

`func NewAuthMethodWithDefaults() *AuthMethod`

NewAuthMethodWithDefaults instantiates a new AuthMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessInfo

`func (o *AuthMethod) GetAccessInfo() AuthMethodAccessInfo`

GetAccessInfo returns the AccessInfo field if non-nil, zero value otherwise.

### GetAccessInfoOk

`func (o *AuthMethod) GetAccessInfoOk() (*AuthMethodAccessInfo, bool)`

GetAccessInfoOk returns a tuple with the AccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessInfo

`func (o *AuthMethod) SetAccessInfo(v AuthMethodAccessInfo)`

SetAccessInfo sets AccessInfo field to given value.

### HasAccessInfo

`func (o *AuthMethod) HasAccessInfo() bool`

HasAccessInfo returns a boolean if a field has been set.

### GetAccountId

`func (o *AuthMethod) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AuthMethod) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AuthMethod) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AuthMethod) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAuthMethodAccessId

`func (o *AuthMethod) GetAuthMethodAccessId() string`

GetAuthMethodAccessId returns the AuthMethodAccessId field if non-nil, zero value otherwise.

### GetAuthMethodAccessIdOk

`func (o *AuthMethod) GetAuthMethodAccessIdOk() (*string, bool)`

GetAuthMethodAccessIdOk returns a tuple with the AuthMethodAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodAccessId

`func (o *AuthMethod) SetAuthMethodAccessId(v string)`

SetAuthMethodAccessId sets AuthMethodAccessId field to given value.

### HasAuthMethodAccessId

`func (o *AuthMethod) HasAuthMethodAccessId() bool`

HasAuthMethodAccessId returns a boolean if a field has been set.

### GetAuthMethodName

`func (o *AuthMethod) GetAuthMethodName() string`

GetAuthMethodName returns the AuthMethodName field if non-nil, zero value otherwise.

### GetAuthMethodNameOk

`func (o *AuthMethod) GetAuthMethodNameOk() (*string, bool)`

GetAuthMethodNameOk returns a tuple with the AuthMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodName

`func (o *AuthMethod) SetAuthMethodName(v string)`

SetAuthMethodName sets AuthMethodName field to given value.

### HasAuthMethodName

`func (o *AuthMethod) HasAuthMethodName() bool`

HasAuthMethodName returns a boolean if a field has been set.

### GetAuthMethodRolesAssoc

`func (o *AuthMethod) GetAuthMethodRolesAssoc() []AuthMethodRoleAssociation`

GetAuthMethodRolesAssoc returns the AuthMethodRolesAssoc field if non-nil, zero value otherwise.

### GetAuthMethodRolesAssocOk

`func (o *AuthMethod) GetAuthMethodRolesAssocOk() (*[]AuthMethodRoleAssociation, bool)`

GetAuthMethodRolesAssocOk returns a tuple with the AuthMethodRolesAssoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodRolesAssoc

`func (o *AuthMethod) SetAuthMethodRolesAssoc(v []AuthMethodRoleAssociation)`

SetAuthMethodRolesAssoc sets AuthMethodRolesAssoc field to given value.

### HasAuthMethodRolesAssoc

`func (o *AuthMethod) HasAuthMethodRolesAssoc() bool`

HasAuthMethodRolesAssoc returns a boolean if a field has been set.

### GetClientPermissions

`func (o *AuthMethod) GetClientPermissions() []string`

GetClientPermissions returns the ClientPermissions field if non-nil, zero value otherwise.

### GetClientPermissionsOk

`func (o *AuthMethod) GetClientPermissionsOk() (*[]string, bool)`

GetClientPermissionsOk returns a tuple with the ClientPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPermissions

`func (o *AuthMethod) SetClientPermissions(v []string)`

SetClientPermissions sets ClientPermissions field to given value.

### HasClientPermissions

`func (o *AuthMethod) HasClientPermissions() bool`

HasClientPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


