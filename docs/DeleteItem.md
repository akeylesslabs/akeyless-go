# DeleteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteImmediately** | Pointer to **bool** | When delete-in-days&#x3D;-1, must be set | [optional] [default to false]
**DeleteInDays** | Pointer to **int64** | The number of days to wait before deleting the item (relevant for keys only) | [optional] [default to 7]
**Name** | **string** | Item name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | The specific version you want to delete - 0&#x3D;last version, -1&#x3D;entire item with all versions (default) | [optional] [default to -1]

## Methods

### NewDeleteItem

`func NewDeleteItem(name string, ) *DeleteItem`

NewDeleteItem instantiates a new DeleteItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteItemWithDefaults

`func NewDeleteItemWithDefaults() *DeleteItem`

NewDeleteItemWithDefaults instantiates a new DeleteItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteImmediately

`func (o *DeleteItem) GetDeleteImmediately() bool`

GetDeleteImmediately returns the DeleteImmediately field if non-nil, zero value otherwise.

### GetDeleteImmediatelyOk

`func (o *DeleteItem) GetDeleteImmediatelyOk() (*bool, bool)`

GetDeleteImmediatelyOk returns a tuple with the DeleteImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteImmediately

`func (o *DeleteItem) SetDeleteImmediately(v bool)`

SetDeleteImmediately sets DeleteImmediately field to given value.

### HasDeleteImmediately

`func (o *DeleteItem) HasDeleteImmediately() bool`

HasDeleteImmediately returns a boolean if a field has been set.

### GetDeleteInDays

`func (o *DeleteItem) GetDeleteInDays() int64`

GetDeleteInDays returns the DeleteInDays field if non-nil, zero value otherwise.

### GetDeleteInDaysOk

`func (o *DeleteItem) GetDeleteInDaysOk() (*int64, bool)`

GetDeleteInDaysOk returns a tuple with the DeleteInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInDays

`func (o *DeleteItem) SetDeleteInDays(v int64)`

SetDeleteInDays sets DeleteInDays field to given value.

### HasDeleteInDays

`func (o *DeleteItem) HasDeleteInDays() bool`

HasDeleteInDays returns a boolean if a field has been set.

### GetName

`func (o *DeleteItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeleteItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeleteItem) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *DeleteItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeleteItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeleteItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeleteItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeleteItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeleteItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *DeleteItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeleteItem) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeleteItem) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeleteItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


