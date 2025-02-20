# PasswordBreachInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreachCheckDate** | Pointer to **time.Time** |  | [optional] 
**BreachCount** | Pointer to **int64** |  | [optional] 
**BreachSuggestions** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewPasswordBreachInfo

`func NewPasswordBreachInfo() *PasswordBreachInfo`

NewPasswordBreachInfo instantiates a new PasswordBreachInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordBreachInfoWithDefaults

`func NewPasswordBreachInfoWithDefaults() *PasswordBreachInfo`

NewPasswordBreachInfoWithDefaults instantiates a new PasswordBreachInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreachCheckDate

`func (o *PasswordBreachInfo) GetBreachCheckDate() time.Time`

GetBreachCheckDate returns the BreachCheckDate field if non-nil, zero value otherwise.

### GetBreachCheckDateOk

`func (o *PasswordBreachInfo) GetBreachCheckDateOk() (*time.Time, bool)`

GetBreachCheckDateOk returns a tuple with the BreachCheckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreachCheckDate

`func (o *PasswordBreachInfo) SetBreachCheckDate(v time.Time)`

SetBreachCheckDate sets BreachCheckDate field to given value.

### HasBreachCheckDate

`func (o *PasswordBreachInfo) HasBreachCheckDate() bool`

HasBreachCheckDate returns a boolean if a field has been set.

### GetBreachCount

`func (o *PasswordBreachInfo) GetBreachCount() int64`

GetBreachCount returns the BreachCount field if non-nil, zero value otherwise.

### GetBreachCountOk

`func (o *PasswordBreachInfo) GetBreachCountOk() (*int64, bool)`

GetBreachCountOk returns a tuple with the BreachCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreachCount

`func (o *PasswordBreachInfo) SetBreachCount(v int64)`

SetBreachCount sets BreachCount field to given value.

### HasBreachCount

`func (o *PasswordBreachInfo) HasBreachCount() bool`

HasBreachCount returns a boolean if a field has been set.

### GetBreachSuggestions

`func (o *PasswordBreachInfo) GetBreachSuggestions() []string`

GetBreachSuggestions returns the BreachSuggestions field if non-nil, zero value otherwise.

### GetBreachSuggestionsOk

`func (o *PasswordBreachInfo) GetBreachSuggestionsOk() (*[]string, bool)`

GetBreachSuggestionsOk returns a tuple with the BreachSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreachSuggestions

`func (o *PasswordBreachInfo) SetBreachSuggestions(v []string)`

SetBreachSuggestions sets BreachSuggestions field to given value.

### HasBreachSuggestions

`func (o *PasswordBreachInfo) HasBreachSuggestions() bool`

HasBreachSuggestions returns a boolean if a field has been set.

### GetStatus

`func (o *PasswordBreachInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PasswordBreachInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PasswordBreachInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PasswordBreachInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


