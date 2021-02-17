# GatewayAddSubAdmins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**SubAdmin** | **[]string** | SubAdmins to add | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayAddSubAdmins

`func NewGatewayAddSubAdmins(subAdmin []string, ) *GatewayAddSubAdmins`

NewGatewayAddSubAdmins instantiates a new GatewayAddSubAdmins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayAddSubAdminsWithDefaults

`func NewGatewayAddSubAdminsWithDefaults() *GatewayAddSubAdmins`

NewGatewayAddSubAdminsWithDefaults instantiates a new GatewayAddSubAdmins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayUrl

`func (o *GatewayAddSubAdmins) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GatewayAddSubAdmins) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GatewayAddSubAdmins) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *GatewayAddSubAdmins) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetSubAdmin

`func (o *GatewayAddSubAdmins) GetSubAdmin() []string`

GetSubAdmin returns the SubAdmin field if non-nil, zero value otherwise.

### GetSubAdminOk

`func (o *GatewayAddSubAdmins) GetSubAdminOk() (*[]string, bool)`

GetSubAdminOk returns a tuple with the SubAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAdmin

`func (o *GatewayAddSubAdmins) SetSubAdmin(v []string)`

SetSubAdmin sets SubAdmin field to given value.


### GetToken

`func (o *GatewayAddSubAdmins) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayAddSubAdmins) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayAddSubAdmins) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayAddSubAdmins) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayAddSubAdmins) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayAddSubAdmins) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayAddSubAdmins) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayAddSubAdmins) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


