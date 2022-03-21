# DeleteTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceDeletion** | Pointer to **bool** | Enforce deletion | [optional] [default to false]
**Name** | **string** | Target name | 
**TargetVersion** | Pointer to **int32** | Target version | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDeleteTarget

`func NewDeleteTarget(name string, ) *DeleteTarget`

NewDeleteTarget instantiates a new DeleteTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTargetWithDefaults

`func NewDeleteTargetWithDefaults() *DeleteTarget`

NewDeleteTargetWithDefaults instantiates a new DeleteTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceDeletion

`func (o *DeleteTarget) GetForceDeletion() bool`

GetForceDeletion returns the ForceDeletion field if non-nil, zero value otherwise.

### GetForceDeletionOk

`func (o *DeleteTarget) GetForceDeletionOk() (*bool, bool)`

GetForceDeletionOk returns a tuple with the ForceDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDeletion

`func (o *DeleteTarget) SetForceDeletion(v bool)`

SetForceDeletion sets ForceDeletion field to given value.

### HasForceDeletion

`func (o *DeleteTarget) HasForceDeletion() bool`

HasForceDeletion returns a boolean if a field has been set.

### GetName

`func (o *DeleteTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeleteTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeleteTarget) SetName(v string)`

SetName sets Name field to given value.


### GetTargetVersion

`func (o *DeleteTarget) GetTargetVersion() int32`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *DeleteTarget) GetTargetVersionOk() (*int32, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *DeleteTarget) SetTargetVersion(v int32)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *DeleteTarget) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetToken

`func (o *DeleteTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeleteTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeleteTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeleteTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeleteTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeleteTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


