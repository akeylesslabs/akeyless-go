# KmipDescribeServerOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**CaCert** | Pointer to **[]int32** |  | [optional] 
**CertificateIssueDate** | Pointer to **time.Time** |  | [optional] 
**CertificateTtlInSeconds** | Pointer to **int64** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Root** | Pointer to **string** |  | [optional] 

## Methods

### NewKmipDescribeServerOutput

`func NewKmipDescribeServerOutput() *KmipDescribeServerOutput`

NewKmipDescribeServerOutput instantiates a new KmipDescribeServerOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipDescribeServerOutputWithDefaults

`func NewKmipDescribeServerOutputWithDefaults() *KmipDescribeServerOutput`

NewKmipDescribeServerOutputWithDefaults instantiates a new KmipDescribeServerOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *KmipDescribeServerOutput) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *KmipDescribeServerOutput) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *KmipDescribeServerOutput) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *KmipDescribeServerOutput) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCaCert

`func (o *KmipDescribeServerOutput) GetCaCert() []int32`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *KmipDescribeServerOutput) GetCaCertOk() (*[]int32, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *KmipDescribeServerOutput) SetCaCert(v []int32)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *KmipDescribeServerOutput) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetCertificateIssueDate

`func (o *KmipDescribeServerOutput) GetCertificateIssueDate() time.Time`

GetCertificateIssueDate returns the CertificateIssueDate field if non-nil, zero value otherwise.

### GetCertificateIssueDateOk

`func (o *KmipDescribeServerOutput) GetCertificateIssueDateOk() (*time.Time, bool)`

GetCertificateIssueDateOk returns a tuple with the CertificateIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssueDate

`func (o *KmipDescribeServerOutput) SetCertificateIssueDate(v time.Time)`

SetCertificateIssueDate sets CertificateIssueDate field to given value.

### HasCertificateIssueDate

`func (o *KmipDescribeServerOutput) HasCertificateIssueDate() bool`

HasCertificateIssueDate returns a boolean if a field has been set.

### GetCertificateTtlInSeconds

`func (o *KmipDescribeServerOutput) GetCertificateTtlInSeconds() int64`

GetCertificateTtlInSeconds returns the CertificateTtlInSeconds field if non-nil, zero value otherwise.

### GetCertificateTtlInSecondsOk

`func (o *KmipDescribeServerOutput) GetCertificateTtlInSecondsOk() (*int64, bool)`

GetCertificateTtlInSecondsOk returns a tuple with the CertificateTtlInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTtlInSeconds

`func (o *KmipDescribeServerOutput) SetCertificateTtlInSeconds(v int64)`

SetCertificateTtlInSeconds sets CertificateTtlInSeconds field to given value.

### HasCertificateTtlInSeconds

`func (o *KmipDescribeServerOutput) HasCertificateTtlInSeconds() bool`

HasCertificateTtlInSeconds returns a boolean if a field has been set.

### GetHostname

`func (o *KmipDescribeServerOutput) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *KmipDescribeServerOutput) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *KmipDescribeServerOutput) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *KmipDescribeServerOutput) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetRoot

`func (o *KmipDescribeServerOutput) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *KmipDescribeServerOutput) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *KmipDescribeServerOutput) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *KmipDescribeServerOutput) HasRoot() bool`

HasRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


