# APIKeyAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The public key value of the API-key. | [optional] 

## Methods

### NewAPIKeyAccessRules

`func NewAPIKeyAccessRules() *APIKeyAccessRules`

NewAPIKeyAccessRules instantiates a new APIKeyAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyAccessRulesWithDefaults

`func NewAPIKeyAccessRulesWithDefaults() *APIKeyAccessRules`

NewAPIKeyAccessRulesWithDefaults instantiates a new APIKeyAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *APIKeyAccessRules) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *APIKeyAccessRules) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *APIKeyAccessRules) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *APIKeyAccessRules) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetKey

`func (o *APIKeyAccessRules) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *APIKeyAccessRules) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *APIKeyAccessRules) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *APIKeyAccessRules) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


