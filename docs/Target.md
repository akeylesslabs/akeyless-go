# Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientPermissions** | Pointer to **[]string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**LastVersion** | Pointer to **int32** |  | [optional] 
**ProtectionKeyName** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **int64** |  | [optional] 
**TargetItemsAssoc** | Pointer to [**[]TargetItemAssociation**](TargetItemAssociation.md) |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**TargetVersions** | Pointer to [**[]ItemVersion**](ItemVersion.md) |  | [optional] 
**WithCustomerFragment** | Pointer to **bool** |  | [optional] 

## Methods

### NewTarget

`func NewTarget() *Target`

NewTarget instantiates a new Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetWithDefaults

`func NewTargetWithDefaults() *Target`

NewTargetWithDefaults instantiates a new Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientPermissions

`func (o *Target) GetClientPermissions() []string`

GetClientPermissions returns the ClientPermissions field if non-nil, zero value otherwise.

### GetClientPermissionsOk

`func (o *Target) GetClientPermissionsOk() (*[]string, bool)`

GetClientPermissionsOk returns a tuple with the ClientPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPermissions

`func (o *Target) SetClientPermissions(v []string)`

SetClientPermissions sets ClientPermissions field to given value.

### HasClientPermissions

`func (o *Target) HasClientPermissions() bool`

HasClientPermissions returns a boolean if a field has been set.

### GetComment

`func (o *Target) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Target) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Target) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Target) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLastVersion

`func (o *Target) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *Target) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *Target) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *Target) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### GetProtectionKeyName

`func (o *Target) GetProtectionKeyName() string`

GetProtectionKeyName returns the ProtectionKeyName field if non-nil, zero value otherwise.

### GetProtectionKeyNameOk

`func (o *Target) GetProtectionKeyNameOk() (*string, bool)`

GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyName

`func (o *Target) SetProtectionKeyName(v string)`

SetProtectionKeyName sets ProtectionKeyName field to given value.

### HasProtectionKeyName

`func (o *Target) HasProtectionKeyName() bool`

HasProtectionKeyName returns a boolean if a field has been set.

### GetTargetId

`func (o *Target) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Target) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Target) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Target) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetItemsAssoc

`func (o *Target) GetTargetItemsAssoc() []TargetItemAssociation`

GetTargetItemsAssoc returns the TargetItemsAssoc field if non-nil, zero value otherwise.

### GetTargetItemsAssocOk

`func (o *Target) GetTargetItemsAssocOk() (*[]TargetItemAssociation, bool)`

GetTargetItemsAssocOk returns a tuple with the TargetItemsAssoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetItemsAssoc

`func (o *Target) SetTargetItemsAssoc(v []TargetItemAssociation)`

SetTargetItemsAssoc sets TargetItemsAssoc field to given value.

### HasTargetItemsAssoc

`func (o *Target) HasTargetItemsAssoc() bool`

HasTargetItemsAssoc returns a boolean if a field has been set.

### GetTargetName

`func (o *Target) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *Target) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *Target) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *Target) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetType

`func (o *Target) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *Target) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *Target) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *Target) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargetVersions

`func (o *Target) GetTargetVersions() []ItemVersion`

GetTargetVersions returns the TargetVersions field if non-nil, zero value otherwise.

### GetTargetVersionsOk

`func (o *Target) GetTargetVersionsOk() (*[]ItemVersion, bool)`

GetTargetVersionsOk returns a tuple with the TargetVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersions

`func (o *Target) SetTargetVersions(v []ItemVersion)`

SetTargetVersions sets TargetVersions field to given value.

### HasTargetVersions

`func (o *Target) HasTargetVersions() bool`

HasTargetVersions returns a boolean if a field has been set.

### GetWithCustomerFragment

`func (o *Target) GetWithCustomerFragment() bool`

GetWithCustomerFragment returns the WithCustomerFragment field if non-nil, zero value otherwise.

### GetWithCustomerFragmentOk

`func (o *Target) GetWithCustomerFragmentOk() (*bool, bool)`

GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCustomerFragment

`func (o *Target) SetWithCustomerFragment(v bool)`

SetWithCustomerFragment sets WithCustomerFragment field to given value.

### HasWithCustomerFragment

`func (o *Target) HasWithCustomerFragment() bool`

HasWithCustomerFragment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


