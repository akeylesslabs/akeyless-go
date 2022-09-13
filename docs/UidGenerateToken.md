# UidGenerateToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodName** | **string** | The universal identity auth method name | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUidGenerateToken

`func NewUidGenerateToken(authMethodName string, ) *UidGenerateToken`

NewUidGenerateToken instantiates a new UidGenerateToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUidGenerateTokenWithDefaults

`func NewUidGenerateTokenWithDefaults() *UidGenerateToken`

NewUidGenerateTokenWithDefaults instantiates a new UidGenerateToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodName

`func (o *UidGenerateToken) GetAuthMethodName() string`

GetAuthMethodName returns the AuthMethodName field if non-nil, zero value otherwise.

### GetAuthMethodNameOk

`func (o *UidGenerateToken) GetAuthMethodNameOk() (*string, bool)`

GetAuthMethodNameOk returns a tuple with the AuthMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodName

`func (o *UidGenerateToken) SetAuthMethodName(v string)`

SetAuthMethodName sets AuthMethodName field to given value.


### GetJson

`func (o *UidGenerateToken) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UidGenerateToken) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UidGenerateToken) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UidGenerateToken) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *UidGenerateToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UidGenerateToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UidGenerateToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UidGenerateToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UidGenerateToken) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UidGenerateToken) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UidGenerateToken) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UidGenerateToken) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


