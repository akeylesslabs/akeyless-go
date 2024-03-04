# GenCustomerFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**HsmKeyLabel** | Pointer to **string** | The label of the hsm key to use for customer fragment operations (relevant for hsm_wrapped/hsm_protected customer fragments) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | Pointer to **string** | Customer fragment name | [optional] 
**Type** | Pointer to **string** | Customer fragment type [standard/hsm_wrapped/hsm_secured] | [optional] [default to "standard"]

## Methods

### NewGenCustomerFragment

`func NewGenCustomerFragment() *GenCustomerFragment`

NewGenCustomerFragment instantiates a new GenCustomerFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenCustomerFragmentWithDefaults

`func NewGenCustomerFragmentWithDefaults() *GenCustomerFragment`

NewGenCustomerFragmentWithDefaults instantiates a new GenCustomerFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GenCustomerFragment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenCustomerFragment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenCustomerFragment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenCustomerFragment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHsmKeyLabel

`func (o *GenCustomerFragment) GetHsmKeyLabel() string`

GetHsmKeyLabel returns the HsmKeyLabel field if non-nil, zero value otherwise.

### GetHsmKeyLabelOk

`func (o *GenCustomerFragment) GetHsmKeyLabelOk() (*string, bool)`

GetHsmKeyLabelOk returns a tuple with the HsmKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmKeyLabel

`func (o *GenCustomerFragment) SetHsmKeyLabel(v string)`

SetHsmKeyLabel sets HsmKeyLabel field to given value.

### HasHsmKeyLabel

`func (o *GenCustomerFragment) HasHsmKeyLabel() bool`

HasHsmKeyLabel returns a boolean if a field has been set.

### GetJson

`func (o *GenCustomerFragment) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GenCustomerFragment) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GenCustomerFragment) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GenCustomerFragment) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMetadata

`func (o *GenCustomerFragment) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenCustomerFragment) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenCustomerFragment) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GenCustomerFragment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *GenCustomerFragment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenCustomerFragment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenCustomerFragment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenCustomerFragment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GenCustomerFragment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenCustomerFragment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenCustomerFragment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GenCustomerFragment) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


