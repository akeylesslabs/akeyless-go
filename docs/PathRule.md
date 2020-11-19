# PathRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | The approved/denied capabilities in the path | [optional] 
**Path** | Pointer to **string** | The path the rule refers to | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPathRule

`func NewPathRule() *PathRule`

NewPathRule instantiates a new PathRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathRuleWithDefaults

`func NewPathRuleWithDefaults() *PathRule`

NewPathRuleWithDefaults instantiates a new PathRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *PathRule) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PathRule) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PathRule) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *PathRule) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetPath

`func (o *PathRule) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PathRule) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PathRule) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PathRule) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *PathRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PathRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PathRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PathRule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


