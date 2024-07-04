# AuthMethodList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Filter by auth method name or part of it | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**PaginationToken** | Pointer to **string** | Next page reference | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **[]string** | The Auth method types list of the requested method. In case it is empty, all types of auth methods will be returned. options: [api_key, azure_ad, oauth2/jwt, saml2, ldap, aws_iam, oidc, universal_identity, gcp, k8s, cert] | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewAuthMethodList

`func NewAuthMethodList() *AuthMethodList`

NewAuthMethodList instantiates a new AuthMethodList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodListWithDefaults

`func NewAuthMethodListWithDefaults() *AuthMethodList`

NewAuthMethodListWithDefaults instantiates a new AuthMethodList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AuthMethodList) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AuthMethodList) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AuthMethodList) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AuthMethodList) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetJson

`func (o *AuthMethodList) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AuthMethodList) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AuthMethodList) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AuthMethodList) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetPaginationToken

`func (o *AuthMethodList) GetPaginationToken() string`

GetPaginationToken returns the PaginationToken field if non-nil, zero value otherwise.

### GetPaginationTokenOk

`func (o *AuthMethodList) GetPaginationTokenOk() (*string, bool)`

GetPaginationTokenOk returns a tuple with the PaginationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationToken

`func (o *AuthMethodList) SetPaginationToken(v string)`

SetPaginationToken sets PaginationToken field to given value.

### HasPaginationToken

`func (o *AuthMethodList) HasPaginationToken() bool`

HasPaginationToken returns a boolean if a field has been set.

### GetToken

`func (o *AuthMethodList) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthMethodList) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthMethodList) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthMethodList) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *AuthMethodList) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthMethodList) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthMethodList) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthMethodList) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *AuthMethodList) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AuthMethodList) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AuthMethodList) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AuthMethodList) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


