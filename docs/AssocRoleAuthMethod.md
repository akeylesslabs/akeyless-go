# AssocRoleAuthMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmName** | **string** | The auth method to associate | 
**RoleName** | **string** | The role to associate | 
**SubClaims** | Pointer to **map[string]string** | key/val of sub claims, e.g group&#x3D;admins,developers | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewAssocRoleAuthMethod

`func NewAssocRoleAuthMethod(amName string, roleName string, ) *AssocRoleAuthMethod`

NewAssocRoleAuthMethod instantiates a new AssocRoleAuthMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssocRoleAuthMethodWithDefaults

`func NewAssocRoleAuthMethodWithDefaults() *AssocRoleAuthMethod`

NewAssocRoleAuthMethodWithDefaults instantiates a new AssocRoleAuthMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmName

`func (o *AssocRoleAuthMethod) GetAmName() string`

GetAmName returns the AmName field if non-nil, zero value otherwise.

### GetAmNameOk

`func (o *AssocRoleAuthMethod) GetAmNameOk() (*string, bool)`

GetAmNameOk returns a tuple with the AmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmName

`func (o *AssocRoleAuthMethod) SetAmName(v string)`

SetAmName sets AmName field to given value.


### GetRoleName

`func (o *AssocRoleAuthMethod) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AssocRoleAuthMethod) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AssocRoleAuthMethod) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetSubClaims

`func (o *AssocRoleAuthMethod) GetSubClaims() map[string]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *AssocRoleAuthMethod) GetSubClaimsOk() (*map[string]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *AssocRoleAuthMethod) SetSubClaims(v map[string]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *AssocRoleAuthMethod) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.

### GetToken

`func (o *AssocRoleAuthMethod) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AssocRoleAuthMethod) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AssocRoleAuthMethod) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AssocRoleAuthMethod) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AssocRoleAuthMethod) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AssocRoleAuthMethod) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AssocRoleAuthMethod) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AssocRoleAuthMethod) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


