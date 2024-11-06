# KubeconfigUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**KubeconfigUserExec**](KubeconfigUserExec.md) |  | [optional] 

## Methods

### NewKubeconfigUser

`func NewKubeconfigUser() *KubeconfigUser`

NewKubeconfigUser instantiates a new KubeconfigUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeconfigUserWithDefaults

`func NewKubeconfigUserWithDefaults() *KubeconfigUser`

NewKubeconfigUserWithDefaults instantiates a new KubeconfigUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubeconfigUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubeconfigUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubeconfigUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubeconfigUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUser

`func (o *KubeconfigUser) GetUser() KubeconfigUserExec`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *KubeconfigUser) GetUserOk() (*KubeconfigUserExec, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *KubeconfigUser) SetUser(v KubeconfigUserExec)`

SetUser sets User field to given value.

### HasUser

`func (o *KubeconfigUser) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


