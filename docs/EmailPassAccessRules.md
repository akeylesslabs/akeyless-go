# EmailPassAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** | The Email value | [optional] 
**HashPass** | Pointer to **string** | The password value | [optional] 

## Methods

### NewEmailPassAccessRules

`func NewEmailPassAccessRules() *EmailPassAccessRules`

NewEmailPassAccessRules instantiates a new EmailPassAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailPassAccessRulesWithDefaults

`func NewEmailPassAccessRulesWithDefaults() *EmailPassAccessRules`

NewEmailPassAccessRulesWithDefaults instantiates a new EmailPassAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *EmailPassAccessRules) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *EmailPassAccessRules) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *EmailPassAccessRules) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *EmailPassAccessRules) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetEmail

`func (o *EmailPassAccessRules) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailPassAccessRules) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailPassAccessRules) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailPassAccessRules) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHashPass

`func (o *EmailPassAccessRules) GetHashPass() string`

GetHashPass returns the HashPass field if non-nil, zero value otherwise.

### GetHashPassOk

`func (o *EmailPassAccessRules) GetHashPassOk() (*string, bool)`

GetHashPassOk returns a tuple with the HashPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashPass

`func (o *EmailPassAccessRules) SetHashPass(v string)`

SetHashPass sets HashPass field to given value.

### HasHashPass

`func (o *EmailPassAccessRules) HasHashPass() bool`

HasHashPass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


