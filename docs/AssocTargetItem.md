# AssocTargetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemName** | **string** | The item to associate | 
**TargetName** | **string** | The target to associate | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewAssocTargetItem

`func NewAssocTargetItem(itemName string, targetName string, ) *AssocTargetItem`

NewAssocTargetItem instantiates a new AssocTargetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssocTargetItemWithDefaults

`func NewAssocTargetItemWithDefaults() *AssocTargetItem`

NewAssocTargetItemWithDefaults instantiates a new AssocTargetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemName

`func (o *AssocTargetItem) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *AssocTargetItem) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *AssocTargetItem) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetTargetName

`func (o *AssocTargetItem) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *AssocTargetItem) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *AssocTargetItem) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetToken

`func (o *AssocTargetItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AssocTargetItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AssocTargetItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AssocTargetItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AssocTargetItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AssocTargetItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AssocTargetItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AssocTargetItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


