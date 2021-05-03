# UpdateManagedKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation | [optional] 
**Name** | **string** | ManagedKey name | 
**NewMetadata** | Pointer to **string** | New item metadata | [optional] [default to "default_metadata"]
**NewName** | Pointer to **string** | New managed key name | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (7-365) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]

## Methods

### NewUpdateManagedKey

`func NewUpdateManagedKey(name string, ) *UpdateManagedKey`

NewUpdateManagedKey instantiates a new UpdateManagedKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateManagedKeyWithDefaults

`func NewUpdateManagedKeyWithDefaults() *UpdateManagedKey`

NewUpdateManagedKeyWithDefaults instantiates a new UpdateManagedKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *UpdateManagedKey) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *UpdateManagedKey) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *UpdateManagedKey) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *UpdateManagedKey) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAutoRotate

`func (o *UpdateManagedKey) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *UpdateManagedKey) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *UpdateManagedKey) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *UpdateManagedKey) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetName

`func (o *UpdateManagedKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateManagedKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateManagedKey) SetName(v string)`

SetName sets Name field to given value.


### GetNewMetadata

`func (o *UpdateManagedKey) GetNewMetadata() string`

GetNewMetadata returns the NewMetadata field if non-nil, zero value otherwise.

### GetNewMetadataOk

`func (o *UpdateManagedKey) GetNewMetadataOk() (*string, bool)`

GetNewMetadataOk returns a tuple with the NewMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMetadata

`func (o *UpdateManagedKey) SetNewMetadata(v string)`

SetNewMetadata sets NewMetadata field to given value.

### HasNewMetadata

`func (o *UpdateManagedKey) HasNewMetadata() bool`

HasNewMetadata returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateManagedKey) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateManagedKey) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateManagedKey) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateManagedKey) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRmTag

`func (o *UpdateManagedKey) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *UpdateManagedKey) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *UpdateManagedKey) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *UpdateManagedKey) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotationInterval

`func (o *UpdateManagedKey) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *UpdateManagedKey) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *UpdateManagedKey) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *UpdateManagedKey) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetToken

`func (o *UpdateManagedKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateManagedKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateManagedKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateManagedKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateManagedKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateManagedKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateManagedKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateManagedKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateUrl

`func (o *UpdateManagedKey) GetUpdateUrl() string`

GetUpdateUrl returns the UpdateUrl field if non-nil, zero value otherwise.

### GetUpdateUrlOk

`func (o *UpdateManagedKey) GetUpdateUrlOk() (*string, bool)`

GetUpdateUrlOk returns a tuple with the UpdateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUrl

`func (o *UpdateManagedKey) SetUpdateUrl(v string)`

SetUpdateUrl sets UpdateUrl field to given value.

### HasUpdateUrl

`func (o *UpdateManagedKey) HasUpdateUrl() bool`

HasUpdateUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


