# CertificateStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificatePem** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationEvents** | Pointer to [**[]CertificateExpirationEvent**](CertificateExpirationEvent.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewCertificateStore

`func NewCertificateStore() *CertificateStore`

NewCertificateStore instantiates a new CertificateStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoreWithDefaults

`func NewCertificateStoreWithDefaults() *CertificateStore`

NewCertificateStoreWithDefaults instantiates a new CertificateStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificatePem

`func (o *CertificateStore) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *CertificateStore) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *CertificateStore) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *CertificateStore) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### GetCommonName

`func (o *CertificateStore) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateStore) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateStore) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CertificateStore) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CertificateStore) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CertificateStore) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CertificateStore) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CertificateStore) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetExpirationEvents

`func (o *CertificateStore) GetExpirationEvents() []CertificateExpirationEvent`

GetExpirationEvents returns the ExpirationEvents field if non-nil, zero value otherwise.

### GetExpirationEventsOk

`func (o *CertificateStore) GetExpirationEventsOk() (*[]CertificateExpirationEvent, bool)`

GetExpirationEventsOk returns a tuple with the ExpirationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEvents

`func (o *CertificateStore) SetExpirationEvents(v []CertificateExpirationEvent)`

SetExpirationEvents sets ExpirationEvents field to given value.

### HasExpirationEvents

`func (o *CertificateStore) HasExpirationEvents() bool`

HasExpirationEvents returns a boolean if a field has been set.

### GetName

`func (o *CertificateStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateStore) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


