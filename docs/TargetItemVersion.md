# TargetItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CustomerFragmentId** | Pointer to **string** |  | [optional] 
**DeletionDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ItemVersionState** | Pointer to **string** | ItemState defines the different states an Item can be in | [optional] 
**LatestVersion** | Pointer to **bool** |  | [optional] 
**ProtectionKeyName** | Pointer to **string** |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WithCustomerFragment** | Pointer to **bool** |  | [optional] 

## Methods

### NewTargetItemVersion

`func NewTargetItemVersion() *TargetItemVersion`

NewTargetItemVersion instantiates a new TargetItemVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetItemVersionWithDefaults

`func NewTargetItemVersionWithDefaults() *TargetItemVersion`

NewTargetItemVersionWithDefaults instantiates a new TargetItemVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *TargetItemVersion) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *TargetItemVersion) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *TargetItemVersion) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *TargetItemVersion) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerFragmentId

`func (o *TargetItemVersion) GetCustomerFragmentId() string`

GetCustomerFragmentId returns the CustomerFragmentId field if non-nil, zero value otherwise.

### GetCustomerFragmentIdOk

`func (o *TargetItemVersion) GetCustomerFragmentIdOk() (*string, bool)`

GetCustomerFragmentIdOk returns a tuple with the CustomerFragmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFragmentId

`func (o *TargetItemVersion) SetCustomerFragmentId(v string)`

SetCustomerFragmentId sets CustomerFragmentId field to given value.

### HasCustomerFragmentId

`func (o *TargetItemVersion) HasCustomerFragmentId() bool`

HasCustomerFragmentId returns a boolean if a field has been set.

### GetDeletionDate

`func (o *TargetItemVersion) GetDeletionDate() time.Time`

GetDeletionDate returns the DeletionDate field if non-nil, zero value otherwise.

### GetDeletionDateOk

`func (o *TargetItemVersion) GetDeletionDateOk() (*time.Time, bool)`

GetDeletionDateOk returns a tuple with the DeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDate

`func (o *TargetItemVersion) SetDeletionDate(v time.Time)`

SetDeletionDate sets DeletionDate field to given value.

### HasDeletionDate

`func (o *TargetItemVersion) HasDeletionDate() bool`

HasDeletionDate returns a boolean if a field has been set.

### GetItemVersionState

`func (o *TargetItemVersion) GetItemVersionState() string`

GetItemVersionState returns the ItemVersionState field if non-nil, zero value otherwise.

### GetItemVersionStateOk

`func (o *TargetItemVersion) GetItemVersionStateOk() (*string, bool)`

GetItemVersionStateOk returns a tuple with the ItemVersionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersionState

`func (o *TargetItemVersion) SetItemVersionState(v string)`

SetItemVersionState sets ItemVersionState field to given value.

### HasItemVersionState

`func (o *TargetItemVersion) HasItemVersionState() bool`

HasItemVersionState returns a boolean if a field has been set.

### GetLatestVersion

`func (o *TargetItemVersion) GetLatestVersion() bool`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *TargetItemVersion) GetLatestVersionOk() (*bool, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *TargetItemVersion) SetLatestVersion(v bool)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *TargetItemVersion) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetProtectionKeyName

`func (o *TargetItemVersion) GetProtectionKeyName() string`

GetProtectionKeyName returns the ProtectionKeyName field if non-nil, zero value otherwise.

### GetProtectionKeyNameOk

`func (o *TargetItemVersion) GetProtectionKeyNameOk() (*string, bool)`

GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyName

`func (o *TargetItemVersion) SetProtectionKeyName(v string)`

SetProtectionKeyName sets ProtectionKeyName field to given value.

### HasProtectionKeyName

`func (o *TargetItemVersion) HasProtectionKeyName() bool`

HasProtectionKeyName returns a boolean if a field has been set.

### GetTargetName

`func (o *TargetItemVersion) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *TargetItemVersion) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *TargetItemVersion) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *TargetItemVersion) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetVersion

`func (o *TargetItemVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TargetItemVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TargetItemVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TargetItemVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWithCustomerFragment

`func (o *TargetItemVersion) GetWithCustomerFragment() bool`

GetWithCustomerFragment returns the WithCustomerFragment field if non-nil, zero value otherwise.

### GetWithCustomerFragmentOk

`func (o *TargetItemVersion) GetWithCustomerFragmentOk() (*bool, bool)`

GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCustomerFragment

`func (o *TargetItemVersion) SetWithCustomerFragment(v bool)`

SetWithCustomerFragment sets WithCustomerFragment field to given value.

### HasWithCustomerFragment

`func (o *TargetItemVersion) HasWithCustomerFragment() bool`

HasWithCustomerFragment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


