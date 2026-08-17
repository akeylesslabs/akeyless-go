# KmipGetCABundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewKmipGetCABundle

`func NewKmipGetCABundle() *KmipGetCABundle`

NewKmipGetCABundle instantiates a new KmipGetCABundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipGetCABundleWithDefaults

`func NewKmipGetCABundleWithDefaults() *KmipGetCABundle`

NewKmipGetCABundleWithDefaults instantiates a new KmipGetCABundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *KmipGetCABundle) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *KmipGetCABundle) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *KmipGetCABundle) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *KmipGetCABundle) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *KmipGetCABundle) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipGetCABundle) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipGetCABundle) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipGetCABundle) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipGetCABundle) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipGetCABundle) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipGetCABundle) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipGetCABundle) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


