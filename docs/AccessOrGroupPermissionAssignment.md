# AccessOrGroupPermissionAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** |  | [optional] 
**AssignmentType** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**SubClaims** | Pointer to [**map[string][]string**](array.md) |  | [optional] 

## Methods

### NewAccessOrGroupPermissionAssignment

`func NewAccessOrGroupPermissionAssignment() *AccessOrGroupPermissionAssignment`

NewAccessOrGroupPermissionAssignment instantiates a new AccessOrGroupPermissionAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessOrGroupPermissionAssignmentWithDefaults

`func NewAccessOrGroupPermissionAssignmentWithDefaults() *AccessOrGroupPermissionAssignment`

NewAccessOrGroupPermissionAssignmentWithDefaults instantiates a new AccessOrGroupPermissionAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *AccessOrGroupPermissionAssignment) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *AccessOrGroupPermissionAssignment) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *AccessOrGroupPermissionAssignment) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *AccessOrGroupPermissionAssignment) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetAssignmentType

`func (o *AccessOrGroupPermissionAssignment) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *AccessOrGroupPermissionAssignment) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *AccessOrGroupPermissionAssignment) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *AccessOrGroupPermissionAssignment) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetGroupId

`func (o *AccessOrGroupPermissionAssignment) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AccessOrGroupPermissionAssignment) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AccessOrGroupPermissionAssignment) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AccessOrGroupPermissionAssignment) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSubClaims

`func (o *AccessOrGroupPermissionAssignment) GetSubClaims() map[string][]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *AccessOrGroupPermissionAssignment) GetSubClaimsOk() (*map[string][]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *AccessOrGroupPermissionAssignment) SetSubClaims(v map[string][]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *AccessOrGroupPermissionAssignment) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


