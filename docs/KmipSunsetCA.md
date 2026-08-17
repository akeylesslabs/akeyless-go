# KmipSunsetCA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaId** | **string** | CA ID to sunset | 
**Force** | Pointer to **bool** | Force sunset even if issued clients or recent usage are detected | [optional] [default to false]
**GracePeriod** | Pointer to **int64** | Grace period in seconds for recent CA usage checks | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewKmipSunsetCA

`func NewKmipSunsetCA(caId string, ) *KmipSunsetCA`

NewKmipSunsetCA instantiates a new KmipSunsetCA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipSunsetCAWithDefaults

`func NewKmipSunsetCAWithDefaults() *KmipSunsetCA`

NewKmipSunsetCAWithDefaults instantiates a new KmipSunsetCA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaId

`func (o *KmipSunsetCA) GetCaId() string`

GetCaId returns the CaId field if non-nil, zero value otherwise.

### GetCaIdOk

`func (o *KmipSunsetCA) GetCaIdOk() (*string, bool)`

GetCaIdOk returns a tuple with the CaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaId

`func (o *KmipSunsetCA) SetCaId(v string)`

SetCaId sets CaId field to given value.


### GetForce

`func (o *KmipSunsetCA) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *KmipSunsetCA) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *KmipSunsetCA) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *KmipSunsetCA) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetGracePeriod

`func (o *KmipSunsetCA) GetGracePeriod() int64`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *KmipSunsetCA) GetGracePeriodOk() (*int64, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *KmipSunsetCA) SetGracePeriod(v int64)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *KmipSunsetCA) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetJson

`func (o *KmipSunsetCA) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *KmipSunsetCA) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *KmipSunsetCA) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *KmipSunsetCA) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *KmipSunsetCA) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipSunsetCA) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipSunsetCA) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipSunsetCA) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipSunsetCA) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipSunsetCA) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipSunsetCA) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipSunsetCA) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


