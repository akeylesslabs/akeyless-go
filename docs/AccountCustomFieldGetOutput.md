# AccountCustomFieldGetOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**DeletionDate** | Pointer to [**NullTime**](NullTime.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ModificationDate** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountCustomFieldGetOutput

`func NewAccountCustomFieldGetOutput() *AccountCustomFieldGetOutput`

NewAccountCustomFieldGetOutput instantiates a new AccountCustomFieldGetOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCustomFieldGetOutputWithDefaults

`func NewAccountCustomFieldGetOutputWithDefaults() *AccountCustomFieldGetOutput`

NewAccountCustomFieldGetOutputWithDefaults instantiates a new AccountCustomFieldGetOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountCustomFieldGetOutput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountCustomFieldGetOutput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountCustomFieldGetOutput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountCustomFieldGetOutput) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreationDate

`func (o *AccountCustomFieldGetOutput) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AccountCustomFieldGetOutput) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AccountCustomFieldGetOutput) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AccountCustomFieldGetOutput) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDeletionDate

`func (o *AccountCustomFieldGetOutput) GetDeletionDate() NullTime`

GetDeletionDate returns the DeletionDate field if non-nil, zero value otherwise.

### GetDeletionDateOk

`func (o *AccountCustomFieldGetOutput) GetDeletionDateOk() (*NullTime, bool)`

GetDeletionDateOk returns a tuple with the DeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDate

`func (o *AccountCustomFieldGetOutput) SetDeletionDate(v NullTime)`

SetDeletionDate sets DeletionDate field to given value.

### HasDeletionDate

`func (o *AccountCustomFieldGetOutput) HasDeletionDate() bool`

HasDeletionDate returns a boolean if a field has been set.

### GetId

`func (o *AccountCustomFieldGetOutput) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountCustomFieldGetOutput) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountCustomFieldGetOutput) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccountCustomFieldGetOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModificationDate

`func (o *AccountCustomFieldGetOutput) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *AccountCustomFieldGetOutput) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *AccountCustomFieldGetOutput) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *AccountCustomFieldGetOutput) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetName

`func (o *AccountCustomFieldGetOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCustomFieldGetOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCustomFieldGetOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountCustomFieldGetOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *AccountCustomFieldGetOutput) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AccountCustomFieldGetOutput) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AccountCustomFieldGetOutput) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AccountCustomFieldGetOutput) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectType

`func (o *AccountCustomFieldGetOutput) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccountCustomFieldGetOutput) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccountCustomFieldGetOutput) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AccountCustomFieldGetOutput) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetRequired

`func (o *AccountCustomFieldGetOutput) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AccountCustomFieldGetOutput) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AccountCustomFieldGetOutput) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AccountCustomFieldGetOutput) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


