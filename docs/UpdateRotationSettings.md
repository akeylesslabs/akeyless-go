# UpdateRotationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRotate** | **bool** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | **string** | Key name | 
**RotationInterval** | Pointer to **int64** | The number of days to wait between every automatic key rotation (7-365) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateRotationSettings

`func NewUpdateRotationSettings(autoRotate bool, name string, ) *UpdateRotationSettings`

NewUpdateRotationSettings instantiates a new UpdateRotationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRotationSettingsWithDefaults

`func NewUpdateRotationSettingsWithDefaults() *UpdateRotationSettings`

NewUpdateRotationSettingsWithDefaults instantiates a new UpdateRotationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRotate

`func (o *UpdateRotationSettings) GetAutoRotate() bool`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *UpdateRotationSettings) GetAutoRotateOk() (*bool, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *UpdateRotationSettings) SetAutoRotate(v bool)`

SetAutoRotate sets AutoRotate field to given value.


### GetJson

`func (o *UpdateRotationSettings) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateRotationSettings) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateRotationSettings) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateRotationSettings) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *UpdateRotationSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRotationSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRotationSettings) SetName(v string)`

SetName sets Name field to given value.


### GetRotationInterval

`func (o *UpdateRotationSettings) GetRotationInterval() int64`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *UpdateRotationSettings) GetRotationIntervalOk() (*int64, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *UpdateRotationSettings) SetRotationInterval(v int64)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *UpdateRotationSettings) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetToken

`func (o *UpdateRotationSettings) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateRotationSettings) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateRotationSettings) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateRotationSettings) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateRotationSettings) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateRotationSettings) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateRotationSettings) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateRotationSettings) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


