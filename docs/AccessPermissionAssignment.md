# AccessPermissionAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to **string** |  | [optional] 
**SubClaims** | Pointer to [**map[string][]string**](array.md) |  | [optional] 

## Methods

### NewAccessPermissionAssignment

`func NewAccessPermissionAssignment() *AccessPermissionAssignment`

NewAccessPermissionAssignment instantiates a new AccessPermissionAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPermissionAssignmentWithDefaults

`func NewAccessPermissionAssignmentWithDefaults() *AccessPermissionAssignment`

NewAccessPermissionAssignmentWithDefaults instantiates a new AccessPermissionAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *AccessPermissionAssignment) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *AccessPermissionAssignment) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *AccessPermissionAssignment) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *AccessPermissionAssignment) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetAccessType

`func (o *AccessPermissionAssignment) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AccessPermissionAssignment) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AccessPermissionAssignment) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *AccessPermissionAssignment) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetSubClaims

`func (o *AccessPermissionAssignment) GetSubClaims() map[string][]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *AccessPermissionAssignment) GetSubClaimsOk() (*map[string][]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *AccessPermissionAssignment) SetSubClaims(v map[string][]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *AccessPermissionAssignment) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


