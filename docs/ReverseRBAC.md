# ReverseRBAC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Path to an object | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | **string** | Type of object (item, am, role) | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewReverseRBAC

`func NewReverseRBAC(path string, type_ string, ) *ReverseRBAC`

NewReverseRBAC instantiates a new ReverseRBAC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseRBACWithDefaults

`func NewReverseRBACWithDefaults() *ReverseRBAC`

NewReverseRBACWithDefaults instantiates a new ReverseRBAC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ReverseRBAC) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReverseRBAC) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReverseRBAC) SetPath(v string)`

SetPath sets Path field to given value.


### GetToken

`func (o *ReverseRBAC) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ReverseRBAC) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ReverseRBAC) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ReverseRBAC) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *ReverseRBAC) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReverseRBAC) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReverseRBAC) SetType(v string)`

SetType sets Type field to given value.


### GetUidToken

`func (o *ReverseRBAC) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ReverseRBAC) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ReverseRBAC) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ReverseRBAC) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


