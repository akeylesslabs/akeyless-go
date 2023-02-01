# SmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sla** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **string** | Tier represents a level of extensibility the account will have, defined by various limits for different resources of Akeyless e.g - A StarterTier may have a limit of 3 Client resources and 50 Secret resources | [optional] 

## Methods

### NewSmInfo

`func NewSmInfo() *SmInfo`

NewSmInfo instantiates a new SmInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmInfoWithDefaults

`func NewSmInfoWithDefaults() *SmInfo`

NewSmInfoWithDefaults instantiates a new SmInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSla

`func (o *SmInfo) GetSla() string`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *SmInfo) GetSlaOk() (*string, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *SmInfo) SetSla(v string)`

SetSla sets Sla field to given value.

### HasSla

`func (o *SmInfo) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetTier

`func (o *SmInfo) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SmInfo) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SmInfo) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *SmInfo) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


