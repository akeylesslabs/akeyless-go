# PolicyCreateKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAlgorithms** | Pointer to **[]string** | Specify allowed key algorithms (e.g., [RSA2048,AES128GCM]) | [optional] 
**AllowedKeyNames** | Pointer to **[]string** | Specify allowed protection key names. To enforce using the account&#39;s default protection key, use &#39;default-account-key&#39; | [optional] 
**AllowedKeyTypes** | Pointer to **[]string** | Specify allowed key protection types (dfc, classic-key) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MaxRotationIntervalDays** | Pointer to **int32** | Set the maximum rotation interval for automatic key rotation. | [optional] 
**ObjectTypes** | Pointer to **[]string** | The object types this policy will apply to (items, targets). If not provided, defaults to [items, targets]. | [optional] 
**Path** | **string** | The path the policy refers to | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewPolicyCreateKeys

`func NewPolicyCreateKeys(path string, ) *PolicyCreateKeys`

NewPolicyCreateKeys instantiates a new PolicyCreateKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCreateKeysWithDefaults

`func NewPolicyCreateKeysWithDefaults() *PolicyCreateKeys`

NewPolicyCreateKeysWithDefaults instantiates a new PolicyCreateKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAlgorithms

`func (o *PolicyCreateKeys) GetAllowedAlgorithms() []string`

GetAllowedAlgorithms returns the AllowedAlgorithms field if non-nil, zero value otherwise.

### GetAllowedAlgorithmsOk

`func (o *PolicyCreateKeys) GetAllowedAlgorithmsOk() (*[]string, bool)`

GetAllowedAlgorithmsOk returns a tuple with the AllowedAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAlgorithms

`func (o *PolicyCreateKeys) SetAllowedAlgorithms(v []string)`

SetAllowedAlgorithms sets AllowedAlgorithms field to given value.

### HasAllowedAlgorithms

`func (o *PolicyCreateKeys) HasAllowedAlgorithms() bool`

HasAllowedAlgorithms returns a boolean if a field has been set.

### GetAllowedKeyNames

`func (o *PolicyCreateKeys) GetAllowedKeyNames() []string`

GetAllowedKeyNames returns the AllowedKeyNames field if non-nil, zero value otherwise.

### GetAllowedKeyNamesOk

`func (o *PolicyCreateKeys) GetAllowedKeyNamesOk() (*[]string, bool)`

GetAllowedKeyNamesOk returns a tuple with the AllowedKeyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKeyNames

`func (o *PolicyCreateKeys) SetAllowedKeyNames(v []string)`

SetAllowedKeyNames sets AllowedKeyNames field to given value.

### HasAllowedKeyNames

`func (o *PolicyCreateKeys) HasAllowedKeyNames() bool`

HasAllowedKeyNames returns a boolean if a field has been set.

### GetAllowedKeyTypes

`func (o *PolicyCreateKeys) GetAllowedKeyTypes() []string`

GetAllowedKeyTypes returns the AllowedKeyTypes field if non-nil, zero value otherwise.

### GetAllowedKeyTypesOk

`func (o *PolicyCreateKeys) GetAllowedKeyTypesOk() (*[]string, bool)`

GetAllowedKeyTypesOk returns a tuple with the AllowedKeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKeyTypes

`func (o *PolicyCreateKeys) SetAllowedKeyTypes(v []string)`

SetAllowedKeyTypes sets AllowedKeyTypes field to given value.

### HasAllowedKeyTypes

`func (o *PolicyCreateKeys) HasAllowedKeyTypes() bool`

HasAllowedKeyTypes returns a boolean if a field has been set.

### GetJson

`func (o *PolicyCreateKeys) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *PolicyCreateKeys) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *PolicyCreateKeys) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *PolicyCreateKeys) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMaxRotationIntervalDays

`func (o *PolicyCreateKeys) GetMaxRotationIntervalDays() int32`

GetMaxRotationIntervalDays returns the MaxRotationIntervalDays field if non-nil, zero value otherwise.

### GetMaxRotationIntervalDaysOk

`func (o *PolicyCreateKeys) GetMaxRotationIntervalDaysOk() (*int32, bool)`

GetMaxRotationIntervalDaysOk returns a tuple with the MaxRotationIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRotationIntervalDays

`func (o *PolicyCreateKeys) SetMaxRotationIntervalDays(v int32)`

SetMaxRotationIntervalDays sets MaxRotationIntervalDays field to given value.

### HasMaxRotationIntervalDays

`func (o *PolicyCreateKeys) HasMaxRotationIntervalDays() bool`

HasMaxRotationIntervalDays returns a boolean if a field has been set.

### GetObjectTypes

`func (o *PolicyCreateKeys) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PolicyCreateKeys) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PolicyCreateKeys) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PolicyCreateKeys) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetPath

`func (o *PolicyCreateKeys) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PolicyCreateKeys) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PolicyCreateKeys) SetPath(v string)`

SetPath sets Path field to given value.


### GetToken

`func (o *PolicyCreateKeys) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyCreateKeys) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyCreateKeys) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyCreateKeys) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *PolicyCreateKeys) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *PolicyCreateKeys) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *PolicyCreateKeys) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *PolicyCreateKeys) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


