# EsmListSecretsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretsList** | Pointer to [**[]SecretInfo**](SecretInfo.md) |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEsmListSecretsOutput

`func NewEsmListSecretsOutput() *EsmListSecretsOutput`

NewEsmListSecretsOutput instantiates a new EsmListSecretsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsmListSecretsOutputWithDefaults

`func NewEsmListSecretsOutputWithDefaults() *EsmListSecretsOutput`

NewEsmListSecretsOutputWithDefaults instantiates a new EsmListSecretsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretsList

`func (o *EsmListSecretsOutput) GetSecretsList() []SecretInfo`

GetSecretsList returns the SecretsList field if non-nil, zero value otherwise.

### GetSecretsListOk

`func (o *EsmListSecretsOutput) GetSecretsListOk() (*[]SecretInfo, bool)`

GetSecretsListOk returns a tuple with the SecretsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsList

`func (o *EsmListSecretsOutput) SetSecretsList(v []SecretInfo)`

SetSecretsList sets SecretsList field to given value.

### HasSecretsList

`func (o *EsmListSecretsOutput) HasSecretsList() bool`

HasSecretsList returns a boolean if a field has been set.

### GetWarnings

`func (o *EsmListSecretsOutput) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *EsmListSecretsOutput) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *EsmListSecretsOutput) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *EsmListSecretsOutput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


