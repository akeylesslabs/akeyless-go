# CertificateChainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRenewCertificate** | Pointer to **bool** |  | [optional] 
**CertificateChain** | Pointer to [**[]CertificateInfo**](CertificateInfo.md) |  | [optional] 
**CertificateFormat** | Pointer to **string** |  | [optional] 
**CertificateHasPrivateKey** | Pointer to **bool** |  | [optional] 
**CertificateIssuerGwClusterId** | Pointer to **int64** |  | [optional] 
**CertificateIssuerGwClusterUrl** | Pointer to **string** |  | [optional] 
**CertificateIssuerItemId** | Pointer to **int64** |  | [optional] 
**CertificateIssuerName** | Pointer to **string** |  | [optional] 
**CertificatePem** | Pointer to **string** |  | [optional] 
**CertificateStatus** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**CsrPem** | Pointer to **string** | CSRPEM contains the PEM-encoded CSR for pending certificates (HTTP-01 challenge) | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationEvents** | Pointer to [**[]CertificateExpirationEvent**](CertificateExpirationEvent.md) |  | [optional] 
**ExternalCaId** | Pointer to [**NullString**](NullString.md) |  | [optional] 
**IssuanceStatus** | Pointer to **string** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**RenewBeforeExpirationInDays** | Pointer to **int64** |  | [optional] 

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

### GetAutoRenewCertificate

`func (o *CertificateChainInfo) GetAutoRenewCertificate() bool`

GetAutoRenewCertificate returns the AutoRenewCertificate field if non-nil, zero value otherwise.

### GetAutoRenewCertificateOk

`func (o *CertificateChainInfo) GetAutoRenewCertificateOk() (*bool, bool)`

GetAutoRenewCertificateOk returns a tuple with the AutoRenewCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewCertificate

`func (o *CertificateChainInfo) SetAutoRenewCertificate(v bool)`

SetAutoRenewCertificate sets AutoRenewCertificate field to given value.

### HasAutoRenewCertificate

`func (o *CertificateChainInfo) HasAutoRenewCertificate() bool`

HasAutoRenewCertificate returns a boolean if a field has been set.

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

### GetCertificateFormat

`func (o *CertificateChainInfo) GetCertificateFormat() string`

GetCertificateFormat returns the CertificateFormat field if non-nil, zero value otherwise.

### GetCertificateFormatOk

`func (o *CertificateChainInfo) GetCertificateFormatOk() (*string, bool)`

GetCertificateFormatOk returns a tuple with the CertificateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFormat

`func (o *CertificateChainInfo) SetCertificateFormat(v string)`

SetCertificateFormat sets CertificateFormat field to given value.

### HasCertificateFormat

`func (o *CertificateChainInfo) HasCertificateFormat() bool`

HasCertificateFormat returns a boolean if a field has been set.

### GetCertificateHasPrivateKey

`func (o *CertificateChainInfo) GetCertificateHasPrivateKey() bool`

GetCertificateHasPrivateKey returns the CertificateHasPrivateKey field if non-nil, zero value otherwise.

### GetCertificateHasPrivateKeyOk

`func (o *CertificateChainInfo) GetCertificateHasPrivateKeyOk() (*bool, bool)`

GetCertificateHasPrivateKeyOk returns a tuple with the CertificateHasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateHasPrivateKey

`func (o *CertificateChainInfo) SetCertificateHasPrivateKey(v bool)`

SetCertificateHasPrivateKey sets CertificateHasPrivateKey field to given value.

### HasCertificateHasPrivateKey

`func (o *CertificateChainInfo) HasCertificateHasPrivateKey() bool`

HasCertificateHasPrivateKey returns a boolean if a field has been set.

### GetCertificateIssuerGwClusterId

`func (o *CertificateChainInfo) GetCertificateIssuerGwClusterId() int64`

GetCertificateIssuerGwClusterId returns the CertificateIssuerGwClusterId field if non-nil, zero value otherwise.

### GetCertificateIssuerGwClusterIdOk

`func (o *CertificateChainInfo) GetCertificateIssuerGwClusterIdOk() (*int64, bool)`

GetCertificateIssuerGwClusterIdOk returns a tuple with the CertificateIssuerGwClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssuerGwClusterId

`func (o *CertificateChainInfo) SetCertificateIssuerGwClusterId(v int64)`

SetCertificateIssuerGwClusterId sets CertificateIssuerGwClusterId field to given value.

### HasCertificateIssuerGwClusterId

`func (o *CertificateChainInfo) HasCertificateIssuerGwClusterId() bool`

HasCertificateIssuerGwClusterId returns a boolean if a field has been set.

### GetCertificateIssuerGwClusterUrl

`func (o *CertificateChainInfo) GetCertificateIssuerGwClusterUrl() string`

GetCertificateIssuerGwClusterUrl returns the CertificateIssuerGwClusterUrl field if non-nil, zero value otherwise.

### GetCertificateIssuerGwClusterUrlOk

`func (o *CertificateChainInfo) GetCertificateIssuerGwClusterUrlOk() (*string, bool)`

GetCertificateIssuerGwClusterUrlOk returns a tuple with the CertificateIssuerGwClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssuerGwClusterUrl

`func (o *CertificateChainInfo) SetCertificateIssuerGwClusterUrl(v string)`

SetCertificateIssuerGwClusterUrl sets CertificateIssuerGwClusterUrl field to given value.

### HasCertificateIssuerGwClusterUrl

`func (o *CertificateChainInfo) HasCertificateIssuerGwClusterUrl() bool`

HasCertificateIssuerGwClusterUrl returns a boolean if a field has been set.

### GetCertificateIssuerItemId

`func (o *CertificateChainInfo) GetCertificateIssuerItemId() int64`

GetCertificateIssuerItemId returns the CertificateIssuerItemId field if non-nil, zero value otherwise.

### GetCertificateIssuerItemIdOk

`func (o *CertificateChainInfo) GetCertificateIssuerItemIdOk() (*int64, bool)`

GetCertificateIssuerItemIdOk returns a tuple with the CertificateIssuerItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssuerItemId

`func (o *CertificateChainInfo) SetCertificateIssuerItemId(v int64)`

SetCertificateIssuerItemId sets CertificateIssuerItemId field to given value.

### HasCertificateIssuerItemId

`func (o *CertificateChainInfo) HasCertificateIssuerItemId() bool`

HasCertificateIssuerItemId returns a boolean if a field has been set.

### GetCertificateIssuerName

`func (o *CertificateChainInfo) GetCertificateIssuerName() string`

GetCertificateIssuerName returns the CertificateIssuerName field if non-nil, zero value otherwise.

### GetCertificateIssuerNameOk

`func (o *CertificateChainInfo) GetCertificateIssuerNameOk() (*string, bool)`

GetCertificateIssuerNameOk returns a tuple with the CertificateIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssuerName

`func (o *CertificateChainInfo) SetCertificateIssuerName(v string)`

SetCertificateIssuerName sets CertificateIssuerName field to given value.

### HasCertificateIssuerName

`func (o *CertificateChainInfo) HasCertificateIssuerName() bool`

HasCertificateIssuerName returns a boolean if a field has been set.

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

### GetCertificateStatus

`func (o *CertificateChainInfo) GetCertificateStatus() string`

GetCertificateStatus returns the CertificateStatus field if non-nil, zero value otherwise.

### GetCertificateStatusOk

`func (o *CertificateChainInfo) GetCertificateStatusOk() (*string, bool)`

GetCertificateStatusOk returns a tuple with the CertificateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatus

`func (o *CertificateChainInfo) SetCertificateStatus(v string)`

SetCertificateStatus sets CertificateStatus field to given value.

### HasCertificateStatus

`func (o *CertificateChainInfo) HasCertificateStatus() bool`

HasCertificateStatus returns a boolean if a field has been set.

### GetCommonName

`func (o *CertificateChainInfo) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateChainInfo) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateChainInfo) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CertificateChainInfo) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCsrPem

`func (o *CertificateChainInfo) GetCsrPem() string`

GetCsrPem returns the CsrPem field if non-nil, zero value otherwise.

### GetCsrPemOk

`func (o *CertificateChainInfo) GetCsrPemOk() (*string, bool)`

GetCsrPemOk returns a tuple with the CsrPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrPem

`func (o *CertificateChainInfo) SetCsrPem(v string)`

SetCsrPem sets CsrPem field to given value.

### HasCsrPem

`func (o *CertificateChainInfo) HasCsrPem() bool`

HasCsrPem returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CertificateChainInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CertificateChainInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CertificateChainInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CertificateChainInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CertificateChainInfo) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CertificateChainInfo) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CertificateChainInfo) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CertificateChainInfo) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

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

### GetExternalCaId

`func (o *CertificateChainInfo) GetExternalCaId() NullString`

GetExternalCaId returns the ExternalCaId field if non-nil, zero value otherwise.

### GetExternalCaIdOk

`func (o *CertificateChainInfo) GetExternalCaIdOk() (*NullString, bool)`

GetExternalCaIdOk returns a tuple with the ExternalCaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCaId

`func (o *CertificateChainInfo) SetExternalCaId(v NullString)`

SetExternalCaId sets ExternalCaId field to given value.

### HasExternalCaId

`func (o *CertificateChainInfo) HasExternalCaId() bool`

HasExternalCaId returns a boolean if a field has been set.

### GetIssuanceStatus

`func (o *CertificateChainInfo) GetIssuanceStatus() string`

GetIssuanceStatus returns the IssuanceStatus field if non-nil, zero value otherwise.

### GetIssuanceStatusOk

`func (o *CertificateChainInfo) GetIssuanceStatusOk() (*string, bool)`

GetIssuanceStatusOk returns a tuple with the IssuanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceStatus

`func (o *CertificateChainInfo) SetIssuanceStatus(v string)`

SetIssuanceStatus sets IssuanceStatus field to given value.

### HasIssuanceStatus

`func (o *CertificateChainInfo) HasIssuanceStatus() bool`

HasIssuanceStatus returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificateChainInfo) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateChainInfo) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateChainInfo) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateChainInfo) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetRenewBeforeExpirationInDays

`func (o *CertificateChainInfo) GetRenewBeforeExpirationInDays() int64`

GetRenewBeforeExpirationInDays returns the RenewBeforeExpirationInDays field if non-nil, zero value otherwise.

### GetRenewBeforeExpirationInDaysOk

`func (o *CertificateChainInfo) GetRenewBeforeExpirationInDaysOk() (*int64, bool)`

GetRenewBeforeExpirationInDaysOk returns a tuple with the RenewBeforeExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewBeforeExpirationInDays

`func (o *CertificateChainInfo) SetRenewBeforeExpirationInDays(v int64)`

SetRenewBeforeExpirationInDays sets RenewBeforeExpirationInDays field to given value.

### HasRenewBeforeExpirationInDays

`func (o *CertificateChainInfo) HasRenewBeforeExpirationInDays() bool`

HasRenewBeforeExpirationInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


