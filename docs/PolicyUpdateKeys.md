# PolicyUpdateKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAlgorithms** | Pointer to **[]string** | Specify allowed key algorithms (e.g., [RSA2048,AES128GCM]) | [optional] 
**AllowedKeyNames** | Pointer to **[]string** | Specify allowed protection key names. To enforce using the account&#39;s default protection key, use &#39;default-account-key&#39; | [optional] 
**AllowedKeyTypes** | Pointer to **[]string** | Specify allowed key protection types (dfc, classic-key) | [optional] 
**Id** | **string** | Policy id | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MaxRotationIntervalDays** | Pointer to **int32** | Set the maximum rotation interval for automatic key rotation. | [optional] 
**ObjectTypes** | Pointer to **[]string** | The object type this policy will apply to (items, targets) | [optional] 
**Path** | Pointer to **string** | The path the policy refers to | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewPolicyUpdateKeys

`func NewPolicyUpdateKeys(id string, ) *PolicyUpdateKeys`

NewPolicyUpdateKeys instantiates a new PolicyUpdateKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUpdateKeysWithDefaults

`func NewPolicyUpdateKeysWithDefaults() *PolicyUpdateKeys`

NewPolicyUpdateKeysWithDefaults instantiates a new PolicyUpdateKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAlgorithms

`func (o *PolicyUpdateKeys) GetAllowedAlgorithms() []string`

GetAllowedAlgorithms returns the AllowedAlgorithms field if non-nil, zero value otherwise.

### GetAllowedAlgorithmsOk

`func (o *PolicyUpdateKeys) GetAllowedAlgorithmsOk() (*[]string, bool)`

GetAllowedAlgorithmsOk returns a tuple with the AllowedAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAlgorithms

`func (o *PolicyUpdateKeys) SetAllowedAlgorithms(v []string)`

SetAllowedAlgorithms sets AllowedAlgorithms field to given value.

### HasAllowedAlgorithms

`func (o *PolicyUpdateKeys) HasAllowedAlgorithms() bool`

HasAllowedAlgorithms returns a boolean if a field has been set.

### GetAllowedKeyNames

`func (o *PolicyUpdateKeys) GetAllowedKeyNames() []string`

GetAllowedKeyNames returns the AllowedKeyNames field if non-nil, zero value otherwise.

### GetAllowedKeyNamesOk

`func (o *PolicyUpdateKeys) GetAllowedKeyNamesOk() (*[]string, bool)`

GetAllowedKeyNamesOk returns a tuple with the AllowedKeyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKeyNames

`func (o *PolicyUpdateKeys) SetAllowedKeyNames(v []string)`

SetAllowedKeyNames sets AllowedKeyNames field to given value.

### HasAllowedKeyNames

`func (o *PolicyUpdateKeys) HasAllowedKeyNames() bool`

HasAllowedKeyNames returns a boolean if a field has been set.

### GetAllowedKeyTypes

`func (o *PolicyUpdateKeys) GetAllowedKeyTypes() []string`

GetAllowedKeyTypes returns the AllowedKeyTypes field if non-nil, zero value otherwise.

### GetAllowedKeyTypesOk

`func (o *PolicyUpdateKeys) GetAllowedKeyTypesOk() (*[]string, bool)`

GetAllowedKeyTypesOk returns a tuple with the AllowedKeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKeyTypes

`func (o *PolicyUpdateKeys) SetAllowedKeyTypes(v []string)`

SetAllowedKeyTypes sets AllowedKeyTypes field to given value.

### HasAllowedKeyTypes

`func (o *PolicyUpdateKeys) HasAllowedKeyTypes() bool`

HasAllowedKeyTypes returns a boolean if a field has been set.

### GetId

`func (o *PolicyUpdateKeys) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyUpdateKeys) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyUpdateKeys) SetId(v string)`

SetId sets Id field to given value.


### GetJson

`func (o *PolicyUpdateKeys) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *PolicyUpdateKeys) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *PolicyUpdateKeys) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *PolicyUpdateKeys) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMaxRotationIntervalDays

`func (o *PolicyUpdateKeys) GetMaxRotationIntervalDays() int32`

GetMaxRotationIntervalDays returns the MaxRotationIntervalDays field if non-nil, zero value otherwise.

### GetMaxRotationIntervalDaysOk

`func (o *PolicyUpdateKeys) GetMaxRotationIntervalDaysOk() (*int32, bool)`

GetMaxRotationIntervalDaysOk returns a tuple with the MaxRotationIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRotationIntervalDays

`func (o *PolicyUpdateKeys) SetMaxRotationIntervalDays(v int32)`

SetMaxRotationIntervalDays sets MaxRotationIntervalDays field to given value.

### HasMaxRotationIntervalDays

`func (o *PolicyUpdateKeys) HasMaxRotationIntervalDays() bool`

HasMaxRotationIntervalDays returns a boolean if a field has been set.

### GetObjectTypes

`func (o *PolicyUpdateKeys) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PolicyUpdateKeys) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PolicyUpdateKeys) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PolicyUpdateKeys) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetPath

`func (o *PolicyUpdateKeys) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PolicyUpdateKeys) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PolicyUpdateKeys) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PolicyUpdateKeys) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetToken

`func (o *PolicyUpdateKeys) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyUpdateKeys) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyUpdateKeys) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyUpdateKeys) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *PolicyUpdateKeys) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *PolicyUpdateKeys) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *PolicyUpdateKeys) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *PolicyUpdateKeys) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


