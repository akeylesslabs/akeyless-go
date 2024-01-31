# AuthMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**AccessDateDisplay** | Pointer to **string** |  | [optional] 
**AccessInfo** | Pointer to [**AuthMethodAccessInfo**](AuthMethodAccessInfo.md) |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AssociatedGwIds** | Pointer to **[]int64** |  | [optional] 
**AuthMethodAccessId** | Pointer to **string** |  | [optional] 
**AuthMethodName** | Pointer to **string** |  | [optional] 
**AuthMethodRolesAssoc** | Pointer to [**[]AuthMethodRoleAssociation**](AuthMethodRoleAssociation.md) |  | [optional] 
**ClientPermissions** | Pointer to **[]string** |  | [optional] 
**CreationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**ModificationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

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

### GetAccessDate

`func (o *AuthMethod) GetAccessDate() time.Time`

GetAccessDate returns the AccessDate field if non-nil, zero value otherwise.

### GetAccessDateOk

`func (o *AuthMethod) GetAccessDateOk() (*time.Time, bool)`

GetAccessDateOk returns a tuple with the AccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDate

`func (o *AuthMethod) SetAccessDate(v time.Time)`

SetAccessDate sets AccessDate field to given value.

### HasAccessDate

`func (o *AuthMethod) HasAccessDate() bool`

HasAccessDate returns a boolean if a field has been set.

### GetAccessDateDisplay

`func (o *AuthMethod) GetAccessDateDisplay() string`

GetAccessDateDisplay returns the AccessDateDisplay field if non-nil, zero value otherwise.

### GetAccessDateDisplayOk

`func (o *AuthMethod) GetAccessDateDisplayOk() (*string, bool)`

GetAccessDateDisplayOk returns a tuple with the AccessDateDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDateDisplay

`func (o *AuthMethod) SetAccessDateDisplay(v string)`

SetAccessDateDisplay sets AccessDateDisplay field to given value.

### HasAccessDateDisplay

`func (o *AuthMethod) HasAccessDateDisplay() bool`

HasAccessDateDisplay returns a boolean if a field has been set.

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

### GetAssociatedGwIds

`func (o *AuthMethod) GetAssociatedGwIds() []int64`

GetAssociatedGwIds returns the AssociatedGwIds field if non-nil, zero value otherwise.

### GetAssociatedGwIdsOk

`func (o *AuthMethod) GetAssociatedGwIdsOk() (*[]int64, bool)`

GetAssociatedGwIdsOk returns a tuple with the AssociatedGwIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedGwIds

`func (o *AuthMethod) SetAssociatedGwIds(v []int64)`

SetAssociatedGwIds sets AssociatedGwIds field to given value.

### HasAssociatedGwIds

`func (o *AuthMethod) HasAssociatedGwIds() bool`

HasAssociatedGwIds returns a boolean if a field has been set.

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

### GetCreationDate

`func (o *AuthMethod) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AuthMethod) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AuthMethod) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AuthMethod) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *AuthMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsApproved

`func (o *AuthMethod) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *AuthMethod) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *AuthMethod) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *AuthMethod) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetModificationDate

`func (o *AuthMethod) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *AuthMethod) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *AuthMethod) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *AuthMethod) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


