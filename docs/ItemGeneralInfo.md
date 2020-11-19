# ItemGeneralInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIssueDetails** | Pointer to [**CertificateIssueInfo**](CertificateIssueInfo.md) |  | [optional] 
**DynamicSecretProducerDetails** | Pointer to [**DynamicSecretProducerInfo**](DynamicSecretProducerInfo.md) |  | [optional] 

## Methods

### NewItemGeneralInfo

`func NewItemGeneralInfo() *ItemGeneralInfo`

NewItemGeneralInfo instantiates a new ItemGeneralInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemGeneralInfoWithDefaults

`func NewItemGeneralInfoWithDefaults() *ItemGeneralInfo`

NewItemGeneralInfoWithDefaults instantiates a new ItemGeneralInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIssueDetails

`func (o *ItemGeneralInfo) GetCertIssueDetails() CertificateIssueInfo`

GetCertIssueDetails returns the CertIssueDetails field if non-nil, zero value otherwise.

### GetCertIssueDetailsOk

`func (o *ItemGeneralInfo) GetCertIssueDetailsOk() (*CertificateIssueInfo, bool)`

GetCertIssueDetailsOk returns a tuple with the CertIssueDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssueDetails

`func (o *ItemGeneralInfo) SetCertIssueDetails(v CertificateIssueInfo)`

SetCertIssueDetails sets CertIssueDetails field to given value.

### HasCertIssueDetails

`func (o *ItemGeneralInfo) HasCertIssueDetails() bool`

HasCertIssueDetails returns a boolean if a field has been set.

### GetDynamicSecretProducerDetails

`func (o *ItemGeneralInfo) GetDynamicSecretProducerDetails() DynamicSecretProducerInfo`

GetDynamicSecretProducerDetails returns the DynamicSecretProducerDetails field if non-nil, zero value otherwise.

### GetDynamicSecretProducerDetailsOk

`func (o *ItemGeneralInfo) GetDynamicSecretProducerDetailsOk() (*DynamicSecretProducerInfo, bool)`

GetDynamicSecretProducerDetailsOk returns a tuple with the DynamicSecretProducerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretProducerDetails

`func (o *ItemGeneralInfo) SetDynamicSecretProducerDetails(v DynamicSecretProducerInfo)`

SetDynamicSecretProducerDetails sets DynamicSecretProducerDetails field to given value.

### HasDynamicSecretProducerDetails

`func (o *ItemGeneralInfo) HasDynamicSecretProducerDetails() bool`

HasDynamicSecretProducerDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


