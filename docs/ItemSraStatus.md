# ItemSraStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountByHostInfo** | Pointer to **map[string]int64** |  | [optional] 
**CountInfo** | Pointer to [**map[string]map[string]int64**](map.md) |  | [optional] 
**HostsInUse** | Pointer to **[]string** |  | [optional] 
**IsInUse** | Pointer to **bool** |  | [optional] 
**LastUsedItem** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewItemSraStatus

`func NewItemSraStatus() *ItemSraStatus`

NewItemSraStatus instantiates a new ItemSraStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSraStatusWithDefaults

`func NewItemSraStatusWithDefaults() *ItemSraStatus`

NewItemSraStatusWithDefaults instantiates a new ItemSraStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountByHostInfo

`func (o *ItemSraStatus) GetCountByHostInfo() map[string]int64`

GetCountByHostInfo returns the CountByHostInfo field if non-nil, zero value otherwise.

### GetCountByHostInfoOk

`func (o *ItemSraStatus) GetCountByHostInfoOk() (*map[string]int64, bool)`

GetCountByHostInfoOk returns a tuple with the CountByHostInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountByHostInfo

`func (o *ItemSraStatus) SetCountByHostInfo(v map[string]int64)`

SetCountByHostInfo sets CountByHostInfo field to given value.

### HasCountByHostInfo

`func (o *ItemSraStatus) HasCountByHostInfo() bool`

HasCountByHostInfo returns a boolean if a field has been set.

### GetCountInfo

`func (o *ItemSraStatus) GetCountInfo() map[string]map[string]int64`

GetCountInfo returns the CountInfo field if non-nil, zero value otherwise.

### GetCountInfoOk

`func (o *ItemSraStatus) GetCountInfoOk() (*map[string]map[string]int64, bool)`

GetCountInfoOk returns a tuple with the CountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountInfo

`func (o *ItemSraStatus) SetCountInfo(v map[string]map[string]int64)`

SetCountInfo sets CountInfo field to given value.

### HasCountInfo

`func (o *ItemSraStatus) HasCountInfo() bool`

HasCountInfo returns a boolean if a field has been set.

### GetHostsInUse

`func (o *ItemSraStatus) GetHostsInUse() []string`

GetHostsInUse returns the HostsInUse field if non-nil, zero value otherwise.

### GetHostsInUseOk

`func (o *ItemSraStatus) GetHostsInUseOk() (*[]string, bool)`

GetHostsInUseOk returns a tuple with the HostsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsInUse

`func (o *ItemSraStatus) SetHostsInUse(v []string)`

SetHostsInUse sets HostsInUse field to given value.

### HasHostsInUse

`func (o *ItemSraStatus) HasHostsInUse() bool`

HasHostsInUse returns a boolean if a field has been set.

### GetIsInUse

`func (o *ItemSraStatus) GetIsInUse() bool`

GetIsInUse returns the IsInUse field if non-nil, zero value otherwise.

### GetIsInUseOk

`func (o *ItemSraStatus) GetIsInUseOk() (*bool, bool)`

GetIsInUseOk returns a tuple with the IsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInUse

`func (o *ItemSraStatus) SetIsInUse(v bool)`

SetIsInUse sets IsInUse field to given value.

### HasIsInUse

`func (o *ItemSraStatus) HasIsInUse() bool`

HasIsInUse returns a boolean if a field has been set.

### GetLastUsedItem

`func (o *ItemSraStatus) GetLastUsedItem() time.Time`

GetLastUsedItem returns the LastUsedItem field if non-nil, zero value otherwise.

### GetLastUsedItemOk

`func (o *ItemSraStatus) GetLastUsedItemOk() (*time.Time, bool)`

GetLastUsedItemOk returns a tuple with the LastUsedItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedItem

`func (o *ItemSraStatus) SetLastUsedItem(v time.Time)`

SetLastUsedItem sets LastUsedItem field to given value.

### HasLastUsedItem

`func (o *ItemSraStatus) HasLastUsedItem() bool`

HasLastUsedItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


