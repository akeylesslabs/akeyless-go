# CustomDnsTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**ProviderType** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomDnsTargetDetails

`func NewCustomDnsTargetDetails() *CustomDnsTargetDetails`

NewCustomDnsTargetDetails instantiates a new CustomDnsTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDnsTargetDetailsWithDefaults

`func NewCustomDnsTargetDetailsWithDefaults() *CustomDnsTargetDetails`

NewCustomDnsTargetDetailsWithDefaults instantiates a new CustomDnsTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *CustomDnsTargetDetails) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CustomDnsTargetDetails) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CustomDnsTargetDetails) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CustomDnsTargetDetails) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetProviderType

`func (o *CustomDnsTargetDetails) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CustomDnsTargetDetails) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CustomDnsTargetDetails) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *CustomDnsTargetDetails) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


