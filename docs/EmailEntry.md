# EmailEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToEmail** | Pointer to **string** |  | [optional] 
**ToName** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailEntry

`func NewEmailEntry() *EmailEntry`

NewEmailEntry instantiates a new EmailEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailEntryWithDefaults

`func NewEmailEntryWithDefaults() *EmailEntry`

NewEmailEntryWithDefaults instantiates a new EmailEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToEmail

`func (o *EmailEntry) GetToEmail() string`

GetToEmail returns the ToEmail field if non-nil, zero value otherwise.

### GetToEmailOk

`func (o *EmailEntry) GetToEmailOk() (*string, bool)`

GetToEmailOk returns a tuple with the ToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEmail

`func (o *EmailEntry) SetToEmail(v string)`

SetToEmail sets ToEmail field to given value.

### HasToEmail

`func (o *EmailEntry) HasToEmail() bool`

HasToEmail returns a boolean if a field has been set.

### GetToName

`func (o *EmailEntry) GetToName() string`

GetToName returns the ToName field if non-nil, zero value otherwise.

### GetToNameOk

`func (o *EmailEntry) GetToNameOk() (*string, bool)`

GetToNameOk returns a tuple with the ToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToName

`func (o *EmailEntry) SetToName(v string)`

SetToName sets ToName field to given value.

### HasToName

`func (o *EmailEntry) HasToName() bool`

HasToName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


