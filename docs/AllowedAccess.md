# AllowedAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**SubClaims** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**SubClaimsCaseInsensitive** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewAllowedAccess

`func NewAllowedAccess() *AllowedAccess`

NewAllowedAccess instantiates a new AllowedAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedAccessWithDefaults

`func NewAllowedAccessWithDefaults() *AllowedAccess`

NewAllowedAccessWithDefaults instantiates a new AllowedAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *AllowedAccess) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *AllowedAccess) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *AllowedAccess) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *AllowedAccess) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetAccessType

`func (o *AllowedAccess) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AllowedAccess) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AllowedAccess) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *AllowedAccess) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetClusterId

`func (o *AllowedAccess) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AllowedAccess) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AllowedAccess) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *AllowedAccess) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AllowedAccess) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AllowedAccess) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AllowedAccess) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AllowedAccess) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AllowedAccess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllowedAccess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllowedAccess) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllowedAccess) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditable

`func (o *AllowedAccess) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AllowedAccess) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AllowedAccess) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *AllowedAccess) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetError

`func (o *AllowedAccess) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AllowedAccess) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AllowedAccess) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AllowedAccess) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *AllowedAccess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedAccess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedAccess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AllowedAccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsValid

`func (o *AllowedAccess) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *AllowedAccess) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *AllowedAccess) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *AllowedAccess) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetName

`func (o *AllowedAccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllowedAccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllowedAccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllowedAccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *AllowedAccess) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AllowedAccess) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AllowedAccess) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AllowedAccess) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSubClaims

`func (o *AllowedAccess) GetSubClaims() map[string][]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *AllowedAccess) GetSubClaimsOk() (*map[string][]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *AllowedAccess) SetSubClaims(v map[string][]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *AllowedAccess) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.

### GetSubClaimsCaseInsensitive

`func (o *AllowedAccess) GetSubClaimsCaseInsensitive() bool`

GetSubClaimsCaseInsensitive returns the SubClaimsCaseInsensitive field if non-nil, zero value otherwise.

### GetSubClaimsCaseInsensitiveOk

`func (o *AllowedAccess) GetSubClaimsCaseInsensitiveOk() (*bool, bool)`

GetSubClaimsCaseInsensitiveOk returns a tuple with the SubClaimsCaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaimsCaseInsensitive

`func (o *AllowedAccess) SetSubClaimsCaseInsensitive(v bool)`

SetSubClaimsCaseInsensitive sets SubClaimsCaseInsensitive field to given value.

### HasSubClaimsCaseInsensitive

`func (o *AllowedAccess) HasSubClaimsCaseInsensitive() bool`

HasSubClaimsCaseInsensitive returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AllowedAccess) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AllowedAccess) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AllowedAccess) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AllowedAccess) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


