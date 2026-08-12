# UscCreateSecretOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartialFailure** | Pointer to **string** | PartialFailure aggregates per-target create failures when some targets still succeeded. | [optional] 
**SecretId** | Pointer to **string** |  | [optional] 
**SelectedEnvironments** | Pointer to **string** | SelectedEnvironments is the subset of GitHub environments where create succeeded (comma-separated). | [optional] 
**SelectedRepositories** | Pointer to **string** | SelectedRepositories is the subset of GitHub repositories where create succeeded (comma-separated). | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewUscCreateSecretOutput

`func NewUscCreateSecretOutput() *UscCreateSecretOutput`

NewUscCreateSecretOutput instantiates a new UscCreateSecretOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUscCreateSecretOutputWithDefaults

`func NewUscCreateSecretOutputWithDefaults() *UscCreateSecretOutput`

NewUscCreateSecretOutputWithDefaults instantiates a new UscCreateSecretOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartialFailure

`func (o *UscCreateSecretOutput) GetPartialFailure() string`

GetPartialFailure returns the PartialFailure field if non-nil, zero value otherwise.

### GetPartialFailureOk

`func (o *UscCreateSecretOutput) GetPartialFailureOk() (*string, bool)`

GetPartialFailureOk returns a tuple with the PartialFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialFailure

`func (o *UscCreateSecretOutput) SetPartialFailure(v string)`

SetPartialFailure sets PartialFailure field to given value.

### HasPartialFailure

`func (o *UscCreateSecretOutput) HasPartialFailure() bool`

HasPartialFailure returns a boolean if a field has been set.

### GetSecretId

`func (o *UscCreateSecretOutput) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *UscCreateSecretOutput) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *UscCreateSecretOutput) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.

### HasSecretId

`func (o *UscCreateSecretOutput) HasSecretId() bool`

HasSecretId returns a boolean if a field has been set.

### GetSelectedEnvironments

`func (o *UscCreateSecretOutput) GetSelectedEnvironments() string`

GetSelectedEnvironments returns the SelectedEnvironments field if non-nil, zero value otherwise.

### GetSelectedEnvironmentsOk

`func (o *UscCreateSecretOutput) GetSelectedEnvironmentsOk() (*string, bool)`

GetSelectedEnvironmentsOk returns a tuple with the SelectedEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedEnvironments

`func (o *UscCreateSecretOutput) SetSelectedEnvironments(v string)`

SetSelectedEnvironments sets SelectedEnvironments field to given value.

### HasSelectedEnvironments

`func (o *UscCreateSecretOutput) HasSelectedEnvironments() bool`

HasSelectedEnvironments returns a boolean if a field has been set.

### GetSelectedRepositories

`func (o *UscCreateSecretOutput) GetSelectedRepositories() string`

GetSelectedRepositories returns the SelectedRepositories field if non-nil, zero value otherwise.

### GetSelectedRepositoriesOk

`func (o *UscCreateSecretOutput) GetSelectedRepositoriesOk() (*string, bool)`

GetSelectedRepositoriesOk returns a tuple with the SelectedRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositories

`func (o *UscCreateSecretOutput) SetSelectedRepositories(v string)`

SetSelectedRepositories sets SelectedRepositories field to given value.

### HasSelectedRepositories

`func (o *UscCreateSecretOutput) HasSelectedRepositories() bool`

HasSelectedRepositories returns a boolean if a field has been set.

### GetVersionId

`func (o *UscCreateSecretOutput) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UscCreateSecretOutput) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UscCreateSecretOutput) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *UscCreateSecretOutput) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


