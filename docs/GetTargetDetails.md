# GetTargetDetails

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

### NewGetTargetDetails

`func NewGetTargetDetails(name string, ) *GetTargetDetails`

NewGetTargetDetails instantiates a new GetTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTargetDetailsWithDefaults

`func NewGetTargetDetailsWithDefaults() *GetTargetDetails`

NewGetTargetDetailsWithDefaults instantiates a new GetTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *GetTargetDetails) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GetTargetDetails) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GetTargetDetails) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GetTargetDetails) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GetTargetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTargetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTargetDetails) SetName(v string)`

SetName sets Name field to given value.


### GetShowVersions

`func (o *GetTargetDetails) GetShowVersions() bool`

GetShowVersions returns the ShowVersions field if non-nil, zero value otherwise.

### GetShowVersionsOk

`func (o *GetTargetDetails) GetShowVersionsOk() (*bool, bool)`

GetShowVersionsOk returns a tuple with the ShowVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowVersions

`func (o *GetTargetDetails) SetShowVersions(v bool)`

SetShowVersions sets ShowVersions field to given value.

### HasShowVersions

`func (o *GetTargetDetails) HasShowVersions() bool`

HasShowVersions returns a boolean if a field has been set.

### GetTargetVersion

`func (o *GetTargetDetails) GetTargetVersion() int32`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *GetTargetDetails) GetTargetVersionOk() (*int32, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *GetTargetDetails) SetTargetVersion(v int32)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *GetTargetDetails) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetToken

`func (o *GetTargetDetails) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetTargetDetails) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetTargetDetails) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetTargetDetails) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GetTargetDetails) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GetTargetDetails) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GetTargetDetails) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GetTargetDetails) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


