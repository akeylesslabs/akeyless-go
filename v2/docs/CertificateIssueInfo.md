# CertificateIssueInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIssuerType** | Pointer to **string** |  | [optional] 
**MaxTtl** | Pointer to **int64** |  | [optional] 
**PkiCertIssuerDetails** | Pointer to [**PKICertificateIssueDetails**](PKICertificateIssueDetails.md) |  | [optional] 
**SshCertIssuerDetails** | Pointer to [**SSHCertificateIssueDetails**](SSHCertificateIssueDetails.md) |  | [optional] 

## Methods

### NewCertificateIssueInfo

`func NewCertificateIssueInfo() *CertificateIssueInfo`

NewCertificateIssueInfo instantiates a new CertificateIssueInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateIssueInfoWithDefaults

`func NewCertificateIssueInfoWithDefaults() *CertificateIssueInfo`

NewCertificateIssueInfoWithDefaults instantiates a new CertificateIssueInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIssuerType

`func (o *CertificateIssueInfo) GetCertIssuerType() string`

GetCertIssuerType returns the CertIssuerType field if non-nil, zero value otherwise.

### GetCertIssuerTypeOk

`func (o *CertificateIssueInfo) GetCertIssuerTypeOk() (*string, bool)`

GetCertIssuerTypeOk returns a tuple with the CertIssuerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerType

`func (o *CertificateIssueInfo) SetCertIssuerType(v string)`

SetCertIssuerType sets CertIssuerType field to given value.

### HasCertIssuerType

`func (o *CertificateIssueInfo) HasCertIssuerType() bool`

HasCertIssuerType returns a boolean if a field has been set.

### GetMaxTtl

`func (o *CertificateIssueInfo) GetMaxTtl() int64`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *CertificateIssueInfo) GetMaxTtlOk() (*int64, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *CertificateIssueInfo) SetMaxTtl(v int64)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *CertificateIssueInfo) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPkiCertIssuerDetails

`func (o *CertificateIssueInfo) GetPkiCertIssuerDetails() PKICertificateIssueDetails`

GetPkiCertIssuerDetails returns the PkiCertIssuerDetails field if non-nil, zero value otherwise.

### GetPkiCertIssuerDetailsOk

`func (o *CertificateIssueInfo) GetPkiCertIssuerDetailsOk() (*PKICertificateIssueDetails, bool)`

GetPkiCertIssuerDetailsOk returns a tuple with the PkiCertIssuerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCertIssuerDetails

`func (o *CertificateIssueInfo) SetPkiCertIssuerDetails(v PKICertificateIssueDetails)`

SetPkiCertIssuerDetails sets PkiCertIssuerDetails field to given value.

### HasPkiCertIssuerDetails

`func (o *CertificateIssueInfo) HasPkiCertIssuerDetails() bool`

HasPkiCertIssuerDetails returns a boolean if a field has been set.

### GetSshCertIssuerDetails

`func (o *CertificateIssueInfo) GetSshCertIssuerDetails() SSHCertificateIssueDetails`

GetSshCertIssuerDetails returns the SshCertIssuerDetails field if non-nil, zero value otherwise.

### GetSshCertIssuerDetailsOk

`func (o *CertificateIssueInfo) GetSshCertIssuerDetailsOk() (*SSHCertificateIssueDetails, bool)`

GetSshCertIssuerDetailsOk returns a tuple with the SshCertIssuerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCertIssuerDetails

`func (o *CertificateIssueInfo) SetSshCertIssuerDetails(v SSHCertificateIssueDetails)`

SetSshCertIssuerDetails sets SshCertIssuerDetails field to given value.

### HasSshCertIssuerDetails

`func (o *CertificateIssueInfo) HasSshCertIssuerDetails() bool`

HasSshCertIssuerDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


