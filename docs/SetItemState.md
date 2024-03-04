# SetItemState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredState** | **string** | Desired item state (Enabled, Disabled) | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Current item name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | The specific version you want to update: 0&#x3D;item level state (default) (relevant only for keys) | [optional] [default to 0]

## Methods

### NewSetItemState

`func NewSetItemState(desiredState string, name string, ) *SetItemState`

NewSetItemState instantiates a new SetItemState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetItemStateWithDefaults

`func NewSetItemStateWithDefaults() *SetItemState`

NewSetItemStateWithDefaults instantiates a new SetItemState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredState

`func (o *SetItemState) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *SetItemState) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *SetItemState) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.


### GetJson

`func (o *SetItemState) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *SetItemState) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *SetItemState) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *SetItemState) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *SetItemState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetItemState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetItemState) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *SetItemState) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SetItemState) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SetItemState) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SetItemState) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *SetItemState) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *SetItemState) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *SetItemState) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *SetItemState) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *SetItemState) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SetItemState) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SetItemState) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SetItemState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


