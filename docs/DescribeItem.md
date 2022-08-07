# DescribeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | The display id of the item | [optional] 
**ItemId** | Pointer to **int64** | Item id of the item | [optional] 
**Name** | **string** | Item name | 
**ShowVersions** | Pointer to **bool** | Include all item versions in reply | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDescribeItem

`func NewDescribeItem(name string, ) *DescribeItem`

NewDescribeItem instantiates a new DescribeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeItemWithDefaults

`func NewDescribeItemWithDefaults() *DescribeItem`

NewDescribeItemWithDefaults instantiates a new DescribeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *DescribeItem) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *DescribeItem) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *DescribeItem) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *DescribeItem) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetItemId

`func (o *DescribeItem) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DescribeItem) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DescribeItem) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DescribeItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetName

`func (o *DescribeItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeItem) SetName(v string)`

SetName sets Name field to given value.


### GetShowVersions

`func (o *DescribeItem) GetShowVersions() bool`

GetShowVersions returns the ShowVersions field if non-nil, zero value otherwise.

### GetShowVersionsOk

`func (o *DescribeItem) GetShowVersionsOk() (*bool, bool)`

GetShowVersionsOk returns a tuple with the ShowVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowVersions

`func (o *DescribeItem) SetShowVersions(v bool)`

SetShowVersions sets ShowVersions field to given value.

### HasShowVersions

`func (o *DescribeItem) HasShowVersions() bool`

HasShowVersions returns a boolean if a field has been set.

### GetToken

`func (o *DescribeItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DescribeItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DescribeItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DescribeItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DescribeItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DescribeItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


