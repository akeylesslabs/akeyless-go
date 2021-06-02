# DeleteTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforceDeletion** | Pointer to **bool** | Enforce deletion | [optional] [default to false]
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**TargetVersion** | Pointer to **int32** | Target version | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

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

### GetEnforceDeletion

`func (o *DeleteTarget) GetEnforceDeletion() bool`

GetEnforceDeletion returns the EnforceDeletion field if non-nil, zero value otherwise.

### GetEnforceDeletionOk

`func (o *DeleteTarget) GetEnforceDeletionOk() (*bool, bool)`

GetEnforceDeletionOk returns a tuple with the EnforceDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceDeletion

`func (o *DeleteTarget) SetEnforceDeletion(v bool)`

SetEnforceDeletion sets EnforceDeletion field to given value.

### HasEnforceDeletion

`func (o *DeleteTarget) HasEnforceDeletion() bool`

HasEnforceDeletion returns a boolean if a field has been set.

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


### GetPassword

`func (o *DeleteTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeleteTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeleteTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DeleteTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

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

### GetUsername

`func (o *DeleteTarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeleteTarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeleteTarget) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DeleteTarget) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


