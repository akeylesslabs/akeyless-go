# PoliciesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | Pointer to **bool** | Aggregate missing configurations from parent policies (requires --paths) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**ObjectType** | Pointer to **[]string** | Optional object types filter (items or targets) | [optional] 
**Paths** | Pointer to **[]string** | Filter by exact policy paths | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Types** | Pointer to **[]string** | Filter by policy types | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewPoliciesList

`func NewPoliciesList() *PoliciesList`

NewPoliciesList instantiates a new PoliciesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesListWithDefaults

`func NewPoliciesListWithDefaults() *PoliciesList`

NewPoliciesListWithDefaults instantiates a new PoliciesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *PoliciesList) GetAggregate() bool`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *PoliciesList) GetAggregateOk() (*bool, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *PoliciesList) SetAggregate(v bool)`

SetAggregate sets Aggregate field to given value.

### HasAggregate

`func (o *PoliciesList) HasAggregate() bool`

HasAggregate returns a boolean if a field has been set.

### GetJson

`func (o *PoliciesList) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *PoliciesList) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *PoliciesList) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *PoliciesList) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetObjectType

`func (o *PoliciesList) GetObjectType() []string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoliciesList) GetObjectTypeOk() (*[]string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoliciesList) SetObjectType(v []string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PoliciesList) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPaths

`func (o *PoliciesList) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *PoliciesList) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *PoliciesList) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *PoliciesList) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetToken

`func (o *PoliciesList) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PoliciesList) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PoliciesList) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PoliciesList) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTypes

`func (o *PoliciesList) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *PoliciesList) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *PoliciesList) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *PoliciesList) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUidToken

`func (o *PoliciesList) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *PoliciesList) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *PoliciesList) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *PoliciesList) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


