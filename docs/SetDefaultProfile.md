# SetDefaultProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Profile** | **string** | Profile name to use as the default | 

## Methods

### NewSetDefaultProfile

`func NewSetDefaultProfile(profile string, ) *SetDefaultProfile`

NewSetDefaultProfile instantiates a new SetDefaultProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDefaultProfileWithDefaults

`func NewSetDefaultProfileWithDefaults() *SetDefaultProfile`

NewSetDefaultProfileWithDefaults instantiates a new SetDefaultProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *SetDefaultProfile) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *SetDefaultProfile) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *SetDefaultProfile) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *SetDefaultProfile) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetProfile

`func (o *SetDefaultProfile) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SetDefaultProfile) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SetDefaultProfile) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


