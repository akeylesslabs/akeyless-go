# SAMLAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSAMLAttribute

`func NewSAMLAttribute() *SAMLAttribute`

NewSAMLAttribute instantiates a new SAMLAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLAttributeWithDefaults

`func NewSAMLAttributeWithDefaults() *SAMLAttribute`

NewSAMLAttributeWithDefaults instantiates a new SAMLAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SAMLAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SAMLAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SAMLAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SAMLAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValues

`func (o *SAMLAttribute) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SAMLAttribute) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SAMLAttribute) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *SAMLAttribute) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


