# PasswordScoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Suggestions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPasswordScoreInfo

`func NewPasswordScoreInfo() *PasswordScoreInfo`

NewPasswordScoreInfo instantiates a new PasswordScoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordScoreInfoWithDefaults

`func NewPasswordScoreInfoWithDefaults() *PasswordScoreInfo`

NewPasswordScoreInfoWithDefaults instantiates a new PasswordScoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *PasswordScoreInfo) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *PasswordScoreInfo) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *PasswordScoreInfo) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *PasswordScoreInfo) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetStatus

`func (o *PasswordScoreInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PasswordScoreInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PasswordScoreInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PasswordScoreInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuggestions

`func (o *PasswordScoreInfo) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *PasswordScoreInfo) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *PasswordScoreInfo) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *PasswordScoreInfo) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


