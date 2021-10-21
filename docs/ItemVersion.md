# ItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CustomerFragmentId** | Pointer to **string** |  | [optional] 
**DeletionDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ItemVersionState** | Pointer to **string** | ItemState defines the different states an Item can be in | [optional] 
**ProtectionKeyName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WithCustomerFragment** | Pointer to **bool** |  | [optional] 

## Methods

### NewItemVersion

`func NewItemVersion() *ItemVersion`

NewItemVersion instantiates a new ItemVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemVersionWithDefaults

`func NewItemVersionWithDefaults() *ItemVersion`

NewItemVersionWithDefaults instantiates a new ItemVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *ItemVersion) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ItemVersion) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ItemVersion) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ItemVersion) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerFragmentId

`func (o *ItemVersion) GetCustomerFragmentId() string`

GetCustomerFragmentId returns the CustomerFragmentId field if non-nil, zero value otherwise.

### GetCustomerFragmentIdOk

`func (o *ItemVersion) GetCustomerFragmentIdOk() (*string, bool)`

GetCustomerFragmentIdOk returns a tuple with the CustomerFragmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFragmentId

`func (o *ItemVersion) SetCustomerFragmentId(v string)`

SetCustomerFragmentId sets CustomerFragmentId field to given value.

### HasCustomerFragmentId

`func (o *ItemVersion) HasCustomerFragmentId() bool`

HasCustomerFragmentId returns a boolean if a field has been set.

### GetDeletionDate

`func (o *ItemVersion) GetDeletionDate() time.Time`

GetDeletionDate returns the DeletionDate field if non-nil, zero value otherwise.

### GetDeletionDateOk

`func (o *ItemVersion) GetDeletionDateOk() (*time.Time, bool)`

GetDeletionDateOk returns a tuple with the DeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDate

`func (o *ItemVersion) SetDeletionDate(v time.Time)`

SetDeletionDate sets DeletionDate field to given value.

### HasDeletionDate

`func (o *ItemVersion) HasDeletionDate() bool`

HasDeletionDate returns a boolean if a field has been set.

### GetItemVersionState

`func (o *ItemVersion) GetItemVersionState() string`

GetItemVersionState returns the ItemVersionState field if non-nil, zero value otherwise.

### GetItemVersionStateOk

`func (o *ItemVersion) GetItemVersionStateOk() (*string, bool)`

GetItemVersionStateOk returns a tuple with the ItemVersionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersionState

`func (o *ItemVersion) SetItemVersionState(v string)`

SetItemVersionState sets ItemVersionState field to given value.

### HasItemVersionState

`func (o *ItemVersion) HasItemVersionState() bool`

HasItemVersionState returns a boolean if a field has been set.

### GetProtectionKeyName

`func (o *ItemVersion) GetProtectionKeyName() string`

GetProtectionKeyName returns the ProtectionKeyName field if non-nil, zero value otherwise.

### GetProtectionKeyNameOk

`func (o *ItemVersion) GetProtectionKeyNameOk() (*string, bool)`

GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyName

`func (o *ItemVersion) SetProtectionKeyName(v string)`

SetProtectionKeyName sets ProtectionKeyName field to given value.

### HasProtectionKeyName

`func (o *ItemVersion) HasProtectionKeyName() bool`

HasProtectionKeyName returns a boolean if a field has been set.

### GetVersion

`func (o *ItemVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ItemVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ItemVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ItemVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWithCustomerFragment

`func (o *ItemVersion) GetWithCustomerFragment() bool`

GetWithCustomerFragment returns the WithCustomerFragment field if non-nil, zero value otherwise.

### GetWithCustomerFragmentOk

`func (o *ItemVersion) GetWithCustomerFragmentOk() (*bool, bool)`

GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCustomerFragment

`func (o *ItemVersion) SetWithCustomerFragment(v bool)`

SetWithCustomerFragment sets WithCustomerFragment field to given value.

### HasWithCustomerFragment

`func (o *ItemVersion) HasWithCustomerFragment() bool`

HasWithCustomerFragment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


