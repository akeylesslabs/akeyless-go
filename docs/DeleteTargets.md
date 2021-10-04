# DeleteTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceDeletion** | Pointer to **bool** | Enforce deletion | [optional] [default to false]
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Path** | **string** | Path to delete the targets from | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewDeleteTargets

`func NewDeleteTargets(path string, ) *DeleteTargets`

NewDeleteTargets instantiates a new DeleteTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTargetsWithDefaults

`func NewDeleteTargetsWithDefaults() *DeleteTargets`

NewDeleteTargetsWithDefaults instantiates a new DeleteTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceDeletion

`func (o *DeleteTargets) GetForceDeletion() bool`

GetForceDeletion returns the ForceDeletion field if non-nil, zero value otherwise.

### GetForceDeletionOk

`func (o *DeleteTargets) GetForceDeletionOk() (*bool, bool)`

GetForceDeletionOk returns a tuple with the ForceDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDeletion

`func (o *DeleteTargets) SetForceDeletion(v bool)`

SetForceDeletion sets ForceDeletion field to given value.

### HasForceDeletion

`func (o *DeleteTargets) HasForceDeletion() bool`

HasForceDeletion returns a boolean if a field has been set.

### GetPassword

`func (o *DeleteTargets) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeleteTargets) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeleteTargets) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DeleteTargets) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *DeleteTargets) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeleteTargets) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeleteTargets) SetPath(v string)`

SetPath sets Path field to given value.


### GetToken

`func (o *DeleteTargets) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteTargets) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteTargets) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeleteTargets) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeleteTargets) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeleteTargets) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeleteTargets) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeleteTargets) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *DeleteTargets) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeleteTargets) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeleteTargets) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DeleteTargets) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


