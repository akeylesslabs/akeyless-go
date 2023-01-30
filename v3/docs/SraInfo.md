# SraInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sla** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **string** | Tier represents a level of extensibility the account will have, defined by various limits for different resources of Akeyless e.g - A StarterTier may have a limit of 3 Client resources and 50 Secret resources | [optional] 
**UserType** | Pointer to **string** |  | [optional] 

## Methods

### NewSraInfo

`func NewSraInfo() *SraInfo`

NewSraInfo instantiates a new SraInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSraInfoWithDefaults

`func NewSraInfoWithDefaults() *SraInfo`

NewSraInfoWithDefaults instantiates a new SraInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSla

`func (o *SraInfo) GetSla() string`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *SraInfo) GetSlaOk() (*string, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *SraInfo) SetSla(v string)`

SetSla sets Sla field to given value.

### HasSla

`func (o *SraInfo) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetTier

`func (o *SraInfo) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SraInfo) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SraInfo) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *SraInfo) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUserType

`func (o *SraInfo) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *SraInfo) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *SraInfo) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *SraInfo) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


