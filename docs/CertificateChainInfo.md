# CertificateChainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateChain** | Pointer to [**[]CertificateInfo**](CertificateInfo.md) |  | [optional] 
**CertificatePem** | Pointer to **string** |  | [optional] 
**ExpirationEvents** | Pointer to [**[]CertificateExpirationEvent**](CertificateExpirationEvent.md) |  | [optional] 

## Methods

### NewCertificateChainInfo

`func NewCertificateChainInfo() *CertificateChainInfo`

NewCertificateChainInfo instantiates a new CertificateChainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateChainInfoWithDefaults

`func NewCertificateChainInfoWithDefaults() *CertificateChainInfo`

NewCertificateChainInfoWithDefaults instantiates a new CertificateChainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateChain

`func (o *CertificateChainInfo) GetCertificateChain() []CertificateInfo`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *CertificateChainInfo) GetCertificateChainOk() (*[]CertificateInfo, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *CertificateChainInfo) SetCertificateChain(v []CertificateInfo)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *CertificateChainInfo) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### GetCertificatePem

`func (o *CertificateChainInfo) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *CertificateChainInfo) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *CertificateChainInfo) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *CertificateChainInfo) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### GetExpirationEvents

`func (o *CertificateChainInfo) GetExpirationEvents() []CertificateExpirationEvent`

GetExpirationEvents returns the ExpirationEvents field if non-nil, zero value otherwise.

### GetExpirationEventsOk

`func (o *CertificateChainInfo) GetExpirationEventsOk() (*[]CertificateExpirationEvent, bool)`

GetExpirationEventsOk returns a tuple with the ExpirationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEvents

`func (o *CertificateChainInfo) SetExpirationEvents(v []CertificateExpirationEvent)`

SetExpirationEvents sets ExpirationEvents field to given value.

### HasExpirationEvents

`func (o *CertificateChainInfo) HasExpirationEvents() bool`

HasExpirationEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


