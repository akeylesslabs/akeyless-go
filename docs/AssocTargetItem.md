# AssocTargetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The item to associate | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**TargetName** | **string** | The target to associate | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewAssocTargetItem

`func NewAssocTargetItem(name string, targetName string, ) *AssocTargetItem`

NewAssocTargetItem instantiates a new AssocTargetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssocTargetItemWithDefaults

`func NewAssocTargetItemWithDefaults() *AssocTargetItem`

NewAssocTargetItemWithDefaults instantiates a new AssocTargetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AssocTargetItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssocTargetItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssocTargetItem) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *AssocTargetItem) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AssocTargetItem) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AssocTargetItem) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AssocTargetItem) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

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

### GetUsername

`func (o *AssocTargetItem) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AssocTargetItem) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AssocTargetItem) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AssocTargetItem) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


