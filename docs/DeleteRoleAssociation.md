# DeleteRoleAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssocId** | **string** | The association id to be deleted | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewDeleteRoleAssociation

`func NewDeleteRoleAssociation(assocId string, ) *DeleteRoleAssociation`

NewDeleteRoleAssociation instantiates a new DeleteRoleAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRoleAssociationWithDefaults

`func NewDeleteRoleAssociationWithDefaults() *DeleteRoleAssociation`

NewDeleteRoleAssociationWithDefaults instantiates a new DeleteRoleAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssocId

`func (o *DeleteRoleAssociation) GetAssocId() string`

GetAssocId returns the AssocId field if non-nil, zero value otherwise.

### GetAssocIdOk

`func (o *DeleteRoleAssociation) GetAssocIdOk() (*string, bool)`

GetAssocIdOk returns a tuple with the AssocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocId

`func (o *DeleteRoleAssociation) SetAssocId(v string)`

SetAssocId sets AssocId field to given value.


### GetPassword

`func (o *DeleteRoleAssociation) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeleteRoleAssociation) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeleteRoleAssociation) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DeleteRoleAssociation) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *DeleteRoleAssociation) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteRoleAssociation) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteRoleAssociation) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeleteRoleAssociation) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeleteRoleAssociation) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeleteRoleAssociation) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeleteRoleAssociation) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeleteRoleAssociation) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *DeleteRoleAssociation) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeleteRoleAssociation) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeleteRoleAssociation) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DeleteRoleAssociation) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


