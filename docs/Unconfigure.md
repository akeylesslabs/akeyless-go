# Unconfigure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **string** | The profile name to be removed | [optional] [default to "default"]

## Methods

### NewUnconfigure

`func NewUnconfigure() *Unconfigure`

NewUnconfigure instantiates a new Unconfigure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnconfigureWithDefaults

`func NewUnconfigureWithDefaults() *Unconfigure`

NewUnconfigureWithDefaults instantiates a new Unconfigure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *Unconfigure) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Unconfigure) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Unconfigure) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Unconfigure) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


