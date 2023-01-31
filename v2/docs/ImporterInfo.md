# ImporterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalItemId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewImporterInfo

`func NewImporterInfo() *ImporterInfo`

NewImporterInfo instantiates a new ImporterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImporterInfoWithDefaults

`func NewImporterInfoWithDefaults() *ImporterInfo`

NewImporterInfoWithDefaults instantiates a new ImporterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalItemId

`func (o *ImporterInfo) GetExternalItemId() string`

GetExternalItemId returns the ExternalItemId field if non-nil, zero value otherwise.

### GetExternalItemIdOk

`func (o *ImporterInfo) GetExternalItemIdOk() (*string, bool)`

GetExternalItemIdOk returns a tuple with the ExternalItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalItemId

`func (o *ImporterInfo) SetExternalItemId(v string)`

SetExternalItemId sets ExternalItemId field to given value.

### HasExternalItemId

`func (o *ImporterInfo) HasExternalItemId() bool`

HasExternalItemId returns a boolean if a field has been set.

### GetVersion

`func (o *ImporterInfo) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImporterInfo) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImporterInfo) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ImporterInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


