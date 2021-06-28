# KMIPClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]PathRule**](PathRule.md) |  | [optional] 

## Methods

### NewKMIPClient

`func NewKMIPClient() *KMIPClient`

NewKMIPClient instantiates a new KMIPClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMIPClientWithDefaults

`func NewKMIPClientWithDefaults() *KMIPClient`

NewKMIPClientWithDefaults instantiates a new KMIPClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KMIPClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KMIPClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KMIPClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KMIPClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *KMIPClient) GetRules() []PathRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *KMIPClient) GetRulesOk() (*[]PathRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *KMIPClient) SetRules(v []PathRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *KMIPClient) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


