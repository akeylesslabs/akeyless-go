# ReverseRBACClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assocs** | Pointer to [**[]AuthMethodRoleAssociation**](AuthMethodRoleAssociation.md) |  | [optional] 
**AuthMethodName** | Pointer to **string** |  | [optional] 

## Methods

### NewReverseRBACClient

`func NewReverseRBACClient() *ReverseRBACClient`

NewReverseRBACClient instantiates a new ReverseRBACClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseRBACClientWithDefaults

`func NewReverseRBACClientWithDefaults() *ReverseRBACClient`

NewReverseRBACClientWithDefaults instantiates a new ReverseRBACClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssocs

`func (o *ReverseRBACClient) GetAssocs() []AuthMethodRoleAssociation`

GetAssocs returns the Assocs field if non-nil, zero value otherwise.

### GetAssocsOk

`func (o *ReverseRBACClient) GetAssocsOk() (*[]AuthMethodRoleAssociation, bool)`

GetAssocsOk returns a tuple with the Assocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocs

`func (o *ReverseRBACClient) SetAssocs(v []AuthMethodRoleAssociation)`

SetAssocs sets Assocs field to given value.

### HasAssocs

`func (o *ReverseRBACClient) HasAssocs() bool`

HasAssocs returns a boolean if a field has been set.

### GetAuthMethodName

`func (o *ReverseRBACClient) GetAuthMethodName() string`

GetAuthMethodName returns the AuthMethodName field if non-nil, zero value otherwise.

### GetAuthMethodNameOk

`func (o *ReverseRBACClient) GetAuthMethodNameOk() (*string, bool)`

GetAuthMethodNameOk returns a tuple with the AuthMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodName

`func (o *ReverseRBACClient) SetAuthMethodName(v string)`

SetAuthMethodName sets AuthMethodName field to given value.

### HasAuthMethodName

`func (o *ReverseRBACClient) HasAuthMethodName() bool`

HasAuthMethodName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


