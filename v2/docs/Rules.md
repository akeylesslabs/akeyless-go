# Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Is admin | [optional] 
**PathRules** | Pointer to [**[]PathRule**](PathRule.md) | The path the rules refers to | [optional] 

## Methods

### NewRules

`func NewRules() *Rules`

NewRules instantiates a new Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesWithDefaults

`func NewRulesWithDefaults() *Rules`

NewRulesWithDefaults instantiates a new Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *Rules) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *Rules) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *Rules) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *Rules) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetPathRules

`func (o *Rules) GetPathRules() []PathRule`

GetPathRules returns the PathRules field if non-nil, zero value otherwise.

### GetPathRulesOk

`func (o *Rules) GetPathRulesOk() (*[]PathRule, bool)`

GetPathRulesOk returns a tuple with the PathRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathRules

`func (o *Rules) SetPathRules(v []PathRule)`

SetPathRules sets PathRules field to given value.

### HasPathRules

`func (o *Rules) HasPathRules() bool`

HasPathRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


