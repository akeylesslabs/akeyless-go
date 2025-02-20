# UnwrapToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**SharedToken** | **string** | The shared token | 

## Methods

### NewUnwrapToken

`func NewUnwrapToken(sharedToken string, ) *UnwrapToken`

NewUnwrapToken instantiates a new UnwrapToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnwrapTokenWithDefaults

`func NewUnwrapTokenWithDefaults() *UnwrapToken`

NewUnwrapTokenWithDefaults instantiates a new UnwrapToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *UnwrapToken) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UnwrapToken) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UnwrapToken) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UnwrapToken) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSharedToken

`func (o *UnwrapToken) GetSharedToken() string`

GetSharedToken returns the SharedToken field if non-nil, zero value otherwise.

### GetSharedTokenOk

`func (o *UnwrapToken) GetSharedTokenOk() (*string, bool)`

GetSharedTokenOk returns a tuple with the SharedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedToken

`func (o *UnwrapToken) SetSharedToken(v string)`

SetSharedToken sets SharedToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


