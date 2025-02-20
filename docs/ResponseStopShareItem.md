# ResponseStopShareItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]EmailError**](EmailError.md) |  | [optional] 
**ItemName** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseStopShareItem

`func NewResponseStopShareItem() *ResponseStopShareItem`

NewResponseStopShareItem instantiates a new ResponseStopShareItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStopShareItemWithDefaults

`func NewResponseStopShareItemWithDefaults() *ResponseStopShareItem`

NewResponseStopShareItemWithDefaults instantiates a new ResponseStopShareItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ResponseStopShareItem) GetErrors() []EmailError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ResponseStopShareItem) GetErrorsOk() (*[]EmailError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ResponseStopShareItem) SetErrors(v []EmailError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ResponseStopShareItem) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetItemName

`func (o *ResponseStopShareItem) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *ResponseStopShareItem) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *ResponseStopShareItem) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *ResponseStopShareItem) HasItemName() bool`

HasItemName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


