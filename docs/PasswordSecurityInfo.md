# PasswordSecurityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreachInfo** | Pointer to [**PasswordBreachInfo**](PasswordBreachInfo.md) |  | [optional] 
**ScoreInfo** | Pointer to [**PasswordScoreInfo**](PasswordScoreInfo.md) |  | [optional] 

## Methods

### NewPasswordSecurityInfo

`func NewPasswordSecurityInfo() *PasswordSecurityInfo`

NewPasswordSecurityInfo instantiates a new PasswordSecurityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordSecurityInfoWithDefaults

`func NewPasswordSecurityInfoWithDefaults() *PasswordSecurityInfo`

NewPasswordSecurityInfoWithDefaults instantiates a new PasswordSecurityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreachInfo

`func (o *PasswordSecurityInfo) GetBreachInfo() PasswordBreachInfo`

GetBreachInfo returns the BreachInfo field if non-nil, zero value otherwise.

### GetBreachInfoOk

`func (o *PasswordSecurityInfo) GetBreachInfoOk() (*PasswordBreachInfo, bool)`

GetBreachInfoOk returns a tuple with the BreachInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreachInfo

`func (o *PasswordSecurityInfo) SetBreachInfo(v PasswordBreachInfo)`

SetBreachInfo sets BreachInfo field to given value.

### HasBreachInfo

`func (o *PasswordSecurityInfo) HasBreachInfo() bool`

HasBreachInfo returns a boolean if a field has been set.

### GetScoreInfo

`func (o *PasswordSecurityInfo) GetScoreInfo() PasswordScoreInfo`

GetScoreInfo returns the ScoreInfo field if non-nil, zero value otherwise.

### GetScoreInfoOk

`func (o *PasswordSecurityInfo) GetScoreInfoOk() (*PasswordScoreInfo, bool)`

GetScoreInfoOk returns a tuple with the ScoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreInfo

`func (o *PasswordSecurityInfo) SetScoreInfo(v PasswordScoreInfo)`

SetScoreInfo sets ScoreInfo field to given value.

### HasScoreInfo

`func (o *PasswordSecurityInfo) HasScoreInfo() bool`

HasScoreInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


