# CertificateVersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotAfter** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**NotBefore** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCertificateVersionInfo

`func NewCertificateVersionInfo() *CertificateVersionInfo`

NewCertificateVersionInfo instantiates a new CertificateVersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateVersionInfoWithDefaults

`func NewCertificateVersionInfoWithDefaults() *CertificateVersionInfo`

NewCertificateVersionInfoWithDefaults instantiates a new CertificateVersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotAfter

`func (o *CertificateVersionInfo) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateVersionInfo) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateVersionInfo) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateVersionInfo) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificateVersionInfo) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateVersionInfo) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateVersionInfo) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateVersionInfo) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateVersionInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateVersionInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateVersionInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateVersionInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


