# KmipMoveServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**NewRoot** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewKmipMoveServer

`func NewKmipMoveServer() *KmipMoveServer`

NewKmipMoveServer instantiates a new KmipMoveServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipMoveServerWithDefaults

`func NewKmipMoveServerWithDefaults() *KmipMoveServer`

NewKmipMoveServerWithDefaults instantiates a new KmipMoveServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *KmipMoveServer) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *KmipMoveServer) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *KmipMoveServer) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *KmipMoveServer) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetNewRoot

`func (o *KmipMoveServer) GetNewRoot() string`

GetNewRoot returns the NewRoot field if non-nil, zero value otherwise.

### GetNewRootOk

`func (o *KmipMoveServer) GetNewRootOk() (*string, bool)`

GetNewRootOk returns a tuple with the NewRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRoot

`func (o *KmipMoveServer) SetNewRoot(v string)`

SetNewRoot sets NewRoot field to given value.

### HasNewRoot

`func (o *KmipMoveServer) HasNewRoot() bool`

HasNewRoot returns a boolean if a field has been set.

### GetToken

`func (o *KmipMoveServer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipMoveServer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipMoveServer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipMoveServer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipMoveServer) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipMoveServer) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipMoveServer) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipMoveServer) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


