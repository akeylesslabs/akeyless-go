# SmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **string** |  | [optional] 

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

### GetPackage

`func (o *SmInfo) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *SmInfo) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *SmInfo) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *SmInfo) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

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


