# RotateKeyOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassicKeyGwUrl** | Pointer to **string** |  | [optional] 
**ItemType** | Pointer to **string** |  | [optional] 
**NewItemVersion** | Pointer to **int32** |  | [optional] 
**NextRotationDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRotateKeyOutput

`func NewRotateKeyOutput() *RotateKeyOutput`

NewRotateKeyOutput instantiates a new RotateKeyOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateKeyOutputWithDefaults

`func NewRotateKeyOutputWithDefaults() *RotateKeyOutput`

NewRotateKeyOutputWithDefaults instantiates a new RotateKeyOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassicKeyGwUrl

`func (o *RotateKeyOutput) GetClassicKeyGwUrl() string`

GetClassicKeyGwUrl returns the ClassicKeyGwUrl field if non-nil, zero value otherwise.

### GetClassicKeyGwUrlOk

`func (o *RotateKeyOutput) GetClassicKeyGwUrlOk() (*string, bool)`

GetClassicKeyGwUrlOk returns a tuple with the ClassicKeyGwUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicKeyGwUrl

`func (o *RotateKeyOutput) SetClassicKeyGwUrl(v string)`

SetClassicKeyGwUrl sets ClassicKeyGwUrl field to given value.

### HasClassicKeyGwUrl

`func (o *RotateKeyOutput) HasClassicKeyGwUrl() bool`

HasClassicKeyGwUrl returns a boolean if a field has been set.

### GetItemType

`func (o *RotateKeyOutput) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *RotateKeyOutput) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *RotateKeyOutput) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *RotateKeyOutput) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetNewItemVersion

`func (o *RotateKeyOutput) GetNewItemVersion() int32`

GetNewItemVersion returns the NewItemVersion field if non-nil, zero value otherwise.

### GetNewItemVersionOk

`func (o *RotateKeyOutput) GetNewItemVersionOk() (*int32, bool)`

GetNewItemVersionOk returns a tuple with the NewItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewItemVersion

`func (o *RotateKeyOutput) SetNewItemVersion(v int32)`

SetNewItemVersion sets NewItemVersion field to given value.

### HasNewItemVersion

`func (o *RotateKeyOutput) HasNewItemVersion() bool`

HasNewItemVersion returns a boolean if a field has been set.

### GetNextRotationDate

`func (o *RotateKeyOutput) GetNextRotationDate() time.Time`

GetNextRotationDate returns the NextRotationDate field if non-nil, zero value otherwise.

### GetNextRotationDateOk

`func (o *RotateKeyOutput) GetNextRotationDateOk() (*time.Time, bool)`

GetNextRotationDateOk returns a tuple with the NextRotationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRotationDate

`func (o *RotateKeyOutput) SetNextRotationDate(v time.Time)`

SetNextRotationDate sets NextRotationDate field to given value.

### HasNextRotationDate

`func (o *RotateKeyOutput) HasNextRotationDate() bool`

HasNextRotationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


