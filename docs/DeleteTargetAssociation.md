# DeleteTargetAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssocId** | Pointer to **string** | The association id to be deleted | [optional] 
**Name** | **string** | Item name | 
**TargetName** | Pointer to **string** | The target to associate | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDeleteTargetAssociation

`func NewDeleteTargetAssociation(name string, ) *DeleteTargetAssociation`

NewDeleteTargetAssociation instantiates a new DeleteTargetAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTargetAssociationWithDefaults

`func NewDeleteTargetAssociationWithDefaults() *DeleteTargetAssociation`

NewDeleteTargetAssociationWithDefaults instantiates a new DeleteTargetAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssocId

`func (o *DeleteTargetAssociation) GetAssocId() string`

GetAssocId returns the AssocId field if non-nil, zero value otherwise.

### GetAssocIdOk

`func (o *DeleteTargetAssociation) GetAssocIdOk() (*string, bool)`

GetAssocIdOk returns a tuple with the AssocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocId

`func (o *DeleteTargetAssociation) SetAssocId(v string)`

SetAssocId sets AssocId field to given value.

### HasAssocId

`func (o *DeleteTargetAssociation) HasAssocId() bool`

HasAssocId returns a boolean if a field has been set.

### GetName

`func (o *DeleteTargetAssociation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeleteTargetAssociation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeleteTargetAssociation) SetName(v string)`

SetName sets Name field to given value.


### GetTargetName

`func (o *DeleteTargetAssociation) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DeleteTargetAssociation) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DeleteTargetAssociation) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DeleteTargetAssociation) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DeleteTargetAssociation) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteTargetAssociation) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteTargetAssociation) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeleteTargetAssociation) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeleteTargetAssociation) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeleteTargetAssociation) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeleteTargetAssociation) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeleteTargetAssociation) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


