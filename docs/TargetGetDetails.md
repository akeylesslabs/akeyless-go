# TargetGetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Target name | 
**ShowVersions** | Pointer to **bool** | Include all target versions in reply | [optional] [default to false]
**TargetVersion** | Pointer to **int32** | Target version | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetGetDetails

`func NewTargetGetDetails(name string, ) *TargetGetDetails`

NewTargetGetDetails instantiates a new TargetGetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGetDetailsWithDefaults

`func NewTargetGetDetailsWithDefaults() *TargetGetDetails`

NewTargetGetDetailsWithDefaults instantiates a new TargetGetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *TargetGetDetails) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetGetDetails) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetGetDetails) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetGetDetails) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *TargetGetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetGetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetGetDetails) SetName(v string)`

SetName sets Name field to given value.


### GetShowVersions

`func (o *TargetGetDetails) GetShowVersions() bool`

GetShowVersions returns the ShowVersions field if non-nil, zero value otherwise.

### GetShowVersionsOk

`func (o *TargetGetDetails) GetShowVersionsOk() (*bool, bool)`

GetShowVersionsOk returns a tuple with the ShowVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowVersions

`func (o *TargetGetDetails) SetShowVersions(v bool)`

SetShowVersions sets ShowVersions field to given value.

### HasShowVersions

`func (o *TargetGetDetails) HasShowVersions() bool`

HasShowVersions returns a boolean if a field has been set.

### GetTargetVersion

`func (o *TargetGetDetails) GetTargetVersion() int32`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *TargetGetDetails) GetTargetVersionOk() (*int32, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *TargetGetDetails) SetTargetVersion(v int32)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *TargetGetDetails) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetToken

`func (o *TargetGetDetails) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetGetDetails) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetGetDetails) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetGetDetails) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetGetDetails) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetGetDetails) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetGetDetails) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetGetDetails) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


