# ObjectVersionSettingsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemType** | Pointer to **string** | VersionSettingsObjectType defines object types for account version settings | [optional] 
**MaxVersions** | Pointer to **string** |  | [optional] 

## Methods

### NewObjectVersionSettingsOutput

`func NewObjectVersionSettingsOutput() *ObjectVersionSettingsOutput`

NewObjectVersionSettingsOutput instantiates a new ObjectVersionSettingsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectVersionSettingsOutputWithDefaults

`func NewObjectVersionSettingsOutputWithDefaults() *ObjectVersionSettingsOutput`

NewObjectVersionSettingsOutputWithDefaults instantiates a new ObjectVersionSettingsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemType

`func (o *ObjectVersionSettingsOutput) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ObjectVersionSettingsOutput) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ObjectVersionSettingsOutput) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ObjectVersionSettingsOutput) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetMaxVersions

`func (o *ObjectVersionSettingsOutput) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *ObjectVersionSettingsOutput) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *ObjectVersionSettingsOutput) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *ObjectVersionSettingsOutput) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


