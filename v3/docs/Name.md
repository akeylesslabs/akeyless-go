# Name

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **[]string** |  | [optional] 
**ExtraNames** | Pointer to [**[]AttributeTypeAndValue**](AttributeTypeAndValue.md) | ExtraNames contains attributes to be copied, raw, into any marshaled distinguished names. Values override any attributes with the same OID. The ExtraNames field is not populated when parsing, see Names. | [optional] 
**Locality** | Pointer to **[]string** |  | [optional] 
**Names** | Pointer to [**[]AttributeTypeAndValue**](AttributeTypeAndValue.md) | Names contains all parsed attributes. When parsing distinguished names, this can be used to extract non-standard attributes that are not parsed by this package. When marshaling to RDNSequences, the Names field is ignored, see ExtraNames. | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**StreetAddress** | Pointer to **[]string** |  | [optional] 

## Methods

### NewName

`func NewName() *Name`

NewName instantiates a new Name object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameWithDefaults

`func NewNameWithDefaults() *Name`

NewNameWithDefaults instantiates a new Name object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *Name) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Name) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Name) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Name) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExtraNames

`func (o *Name) GetExtraNames() []AttributeTypeAndValue`

GetExtraNames returns the ExtraNames field if non-nil, zero value otherwise.

### GetExtraNamesOk

`func (o *Name) GetExtraNamesOk() (*[]AttributeTypeAndValue, bool)`

GetExtraNamesOk returns a tuple with the ExtraNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraNames

`func (o *Name) SetExtraNames(v []AttributeTypeAndValue)`

SetExtraNames sets ExtraNames field to given value.

### HasExtraNames

`func (o *Name) HasExtraNames() bool`

HasExtraNames returns a boolean if a field has been set.

### GetLocality

`func (o *Name) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *Name) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *Name) SetLocality(v []string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *Name) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetNames

`func (o *Name) GetNames() []AttributeTypeAndValue`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Name) GetNamesOk() (*[]AttributeTypeAndValue, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Name) SetNames(v []AttributeTypeAndValue)`

SetNames sets Names field to given value.

### HasNames

`func (o *Name) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Name) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Name) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Name) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Name) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStreetAddress

`func (o *Name) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *Name) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *Name) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *Name) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


