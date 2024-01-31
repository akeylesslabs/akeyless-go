# ProvisionCertificateOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailMessage** | Pointer to **string** |  | [optional] 
**SuccessMessage** | Pointer to **string** |  | [optional] 
**HostNames** | Pointer to **[]string** |  | [optional] 
**ProvisionedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


