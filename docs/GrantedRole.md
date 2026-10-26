# GrantedRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditional** | Pointer to **bool** | Conditional is set when the role association also needs something besides this group: another claim, or membership in another group (\&quot; AND \&quot;). | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewGrantedRole

`func NewGrantedRole() *GrantedRole`

NewGrantedRole instantiates a new GrantedRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantedRoleWithDefaults

`func NewGrantedRoleWithDefaults() *GrantedRole`

NewGrantedRoleWithDefaults instantiates a new GrantedRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditional

`func (o *GrantedRole) GetConditional() bool`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *GrantedRole) GetConditionalOk() (*bool, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *GrantedRole) SetConditional(v bool)`

SetConditional sets Conditional field to given value.

### HasConditional

`func (o *GrantedRole) HasConditional() bool`

HasConditional returns a boolean if a field has been set.

### GetName

`func (o *GrantedRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GrantedRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GrantedRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GrantedRole) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


