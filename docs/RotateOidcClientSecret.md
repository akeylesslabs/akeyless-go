# RotateOidcClientSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | OIDC application name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotateOidcClientSecret

`func NewRotateOidcClientSecret(name string, ) *RotateOidcClientSecret`

NewRotateOidcClientSecret instantiates a new RotateOidcClientSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateOidcClientSecretWithDefaults

`func NewRotateOidcClientSecretWithDefaults() *RotateOidcClientSecret`

NewRotateOidcClientSecretWithDefaults instantiates a new RotateOidcClientSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *RotateOidcClientSecret) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotateOidcClientSecret) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotateOidcClientSecret) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotateOidcClientSecret) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RotateOidcClientSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotateOidcClientSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotateOidcClientSecret) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *RotateOidcClientSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotateOidcClientSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotateOidcClientSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotateOidcClientSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotateOidcClientSecret) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotateOidcClientSecret) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotateOidcClientSecret) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotateOidcClientSecret) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


