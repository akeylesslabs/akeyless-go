# GatewayDeleteSubAdmins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**SubAdmin** | **[]string** | SubAdmins to be removed | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayDeleteSubAdmins

`func NewGatewayDeleteSubAdmins(subAdmin []string, ) *GatewayDeleteSubAdmins`

NewGatewayDeleteSubAdmins instantiates a new GatewayDeleteSubAdmins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayDeleteSubAdminsWithDefaults

`func NewGatewayDeleteSubAdminsWithDefaults() *GatewayDeleteSubAdmins`

NewGatewayDeleteSubAdminsWithDefaults instantiates a new GatewayDeleteSubAdmins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayUrl

`func (o *GatewayDeleteSubAdmins) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GatewayDeleteSubAdmins) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GatewayDeleteSubAdmins) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *GatewayDeleteSubAdmins) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetSubAdmin

`func (o *GatewayDeleteSubAdmins) GetSubAdmin() []string`

GetSubAdmin returns the SubAdmin field if non-nil, zero value otherwise.

### GetSubAdminOk

`func (o *GatewayDeleteSubAdmins) GetSubAdminOk() (*[]string, bool)`

GetSubAdminOk returns a tuple with the SubAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAdmin

`func (o *GatewayDeleteSubAdmins) SetSubAdmin(v []string)`

SetSubAdmin sets SubAdmin field to given value.


### GetToken

`func (o *GatewayDeleteSubAdmins) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayDeleteSubAdmins) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayDeleteSubAdmins) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayDeleteSubAdmins) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayDeleteSubAdmins) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayDeleteSubAdmins) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayDeleteSubAdmins) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayDeleteSubAdmins) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


