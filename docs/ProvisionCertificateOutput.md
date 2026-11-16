# ProvisionCertificateOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailMessage** | Pointer to **string** |  | [optional] 
**SuccessMessage** | Pointer to **string** |  | [optional] 
**AssociationId** | Pointer to **string** | AssociationID and TargetName identify the row of a Results entry. | [optional] 
**HostNames** | Pointer to **[]string** |  | [optional] 
**ProvisionedAt** | Pointer to **time.Time** |  | [optional] 
**Results** | Pointer to [**[]ProvisionCertificateOutput**](ProvisionCertificateOutput.md) | Results holds one entry per row. It is set only when several rows were handled; the top-level fields then summarize them. | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 

## Methods

### NewProvisionCertificateOutput

`func NewProvisionCertificateOutput() *ProvisionCertificateOutput`

NewProvisionCertificateOutput instantiates a new ProvisionCertificateOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionCertificateOutputWithDefaults

`func NewProvisionCertificateOutputWithDefaults() *ProvisionCertificateOutput`

NewProvisionCertificateOutputWithDefaults instantiates a new ProvisionCertificateOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailMessage

`func (o *ProvisionCertificateOutput) GetFailMessage() string`

GetFailMessage returns the FailMessage field if non-nil, zero value otherwise.

### GetFailMessageOk

`func (o *ProvisionCertificateOutput) GetFailMessageOk() (*string, bool)`

GetFailMessageOk returns a tuple with the FailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailMessage

`func (o *ProvisionCertificateOutput) SetFailMessage(v string)`

SetFailMessage sets FailMessage field to given value.

### HasFailMessage

`func (o *ProvisionCertificateOutput) HasFailMessage() bool`

HasFailMessage returns a boolean if a field has been set.

### GetSuccessMessage

`func (o *ProvisionCertificateOutput) GetSuccessMessage() string`

GetSuccessMessage returns the SuccessMessage field if non-nil, zero value otherwise.

### GetSuccessMessageOk

`func (o *ProvisionCertificateOutput) GetSuccessMessageOk() (*string, bool)`

GetSuccessMessageOk returns a tuple with the SuccessMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessMessage

`func (o *ProvisionCertificateOutput) SetSuccessMessage(v string)`

SetSuccessMessage sets SuccessMessage field to given value.

### HasSuccessMessage

`func (o *ProvisionCertificateOutput) HasSuccessMessage() bool`

HasSuccessMessage returns a boolean if a field has been set.

### GetAssociationId

`func (o *ProvisionCertificateOutput) GetAssociationId() string`

GetAssociationId returns the AssociationId field if non-nil, zero value otherwise.

### GetAssociationIdOk

`func (o *ProvisionCertificateOutput) GetAssociationIdOk() (*string, bool)`

GetAssociationIdOk returns a tuple with the AssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationId

`func (o *ProvisionCertificateOutput) SetAssociationId(v string)`

SetAssociationId sets AssociationId field to given value.

### HasAssociationId

`func (o *ProvisionCertificateOutput) HasAssociationId() bool`

HasAssociationId returns a boolean if a field has been set.

### GetHostNames

`func (o *ProvisionCertificateOutput) GetHostNames() []string`

GetHostNames returns the HostNames field if non-nil, zero value otherwise.

### GetHostNamesOk

`func (o *ProvisionCertificateOutput) GetHostNamesOk() (*[]string, bool)`

GetHostNamesOk returns a tuple with the HostNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNames

`func (o *ProvisionCertificateOutput) SetHostNames(v []string)`

SetHostNames sets HostNames field to given value.

### HasHostNames

`func (o *ProvisionCertificateOutput) HasHostNames() bool`

HasHostNames returns a boolean if a field has been set.

### GetProvisionedAt

`func (o *ProvisionCertificateOutput) GetProvisionedAt() time.Time`

GetProvisionedAt returns the ProvisionedAt field if non-nil, zero value otherwise.

### GetProvisionedAtOk

`func (o *ProvisionCertificateOutput) GetProvisionedAtOk() (*time.Time, bool)`

GetProvisionedAtOk returns a tuple with the ProvisionedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedAt

`func (o *ProvisionCertificateOutput) SetProvisionedAt(v time.Time)`

SetProvisionedAt sets ProvisionedAt field to given value.

### HasProvisionedAt

`func (o *ProvisionCertificateOutput) HasProvisionedAt() bool`

HasProvisionedAt returns a boolean if a field has been set.

### GetResults

`func (o *ProvisionCertificateOutput) GetResults() []ProvisionCertificateOutput`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ProvisionCertificateOutput) GetResultsOk() (*[]ProvisionCertificateOutput, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ProvisionCertificateOutput) SetResults(v []ProvisionCertificateOutput)`

SetResults sets Results field to given value.

### HasResults

`func (o *ProvisionCertificateOutput) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTargetName

`func (o *ProvisionCertificateOutput) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ProvisionCertificateOutput) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ProvisionCertificateOutput) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ProvisionCertificateOutput) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


