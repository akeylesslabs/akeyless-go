# MoveObjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectsType** | Pointer to **string** | The objects type to move (item/auth_method/role) | [optional] [default to "item"]
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Source** | **string** | Source path to move the objects from | 
**Target** | **string** | Target path to move the objects to | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewMoveObjects

`func NewMoveObjects(source string, target string, ) *MoveObjects`

NewMoveObjects instantiates a new MoveObjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveObjectsWithDefaults

`func NewMoveObjectsWithDefaults() *MoveObjects`

NewMoveObjectsWithDefaults instantiates a new MoveObjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectsType

`func (o *MoveObjects) GetObjectsType() string`

GetObjectsType returns the ObjectsType field if non-nil, zero value otherwise.

### GetObjectsTypeOk

`func (o *MoveObjects) GetObjectsTypeOk() (*string, bool)`

GetObjectsTypeOk returns a tuple with the ObjectsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsType

`func (o *MoveObjects) SetObjectsType(v string)`

SetObjectsType sets ObjectsType field to given value.

### HasObjectsType

`func (o *MoveObjects) HasObjectsType() bool`

HasObjectsType returns a boolean if a field has been set.

### GetPassword

`func (o *MoveObjects) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MoveObjects) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MoveObjects) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MoveObjects) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSource

`func (o *MoveObjects) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MoveObjects) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MoveObjects) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *MoveObjects) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MoveObjects) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MoveObjects) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetToken

`func (o *MoveObjects) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MoveObjects) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MoveObjects) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MoveObjects) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *MoveObjects) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *MoveObjects) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *MoveObjects) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *MoveObjects) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *MoveObjects) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MoveObjects) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MoveObjects) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MoveObjects) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


