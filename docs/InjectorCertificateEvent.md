# InjectorCertificateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FingerprintSha256** | Pointer to **string** |  | [optional] 
**InjectorName** | Pointer to **string** |  | [optional] 
**InjectorNamespace** | Pointer to **string** |  | [optional] 
**InstallationId** | Pointer to **string** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**Threshold** | Pointer to **string** |  | [optional] 

## Methods

### NewInjectorCertificateEvent

`func NewInjectorCertificateEvent() *InjectorCertificateEvent`

NewInjectorCertificateEvent instantiates a new InjectorCertificateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInjectorCertificateEventWithDefaults

`func NewInjectorCertificateEventWithDefaults() *InjectorCertificateEvent`

NewInjectorCertificateEventWithDefaults instantiates a new InjectorCertificateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprintSha256

`func (o *InjectorCertificateEvent) GetFingerprintSha256() string`

GetFingerprintSha256 returns the FingerprintSha256 field if non-nil, zero value otherwise.

### GetFingerprintSha256Ok

`func (o *InjectorCertificateEvent) GetFingerprintSha256Ok() (*string, bool)`

GetFingerprintSha256Ok returns a tuple with the FingerprintSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintSha256

`func (o *InjectorCertificateEvent) SetFingerprintSha256(v string)`

SetFingerprintSha256 sets FingerprintSha256 field to given value.

### HasFingerprintSha256

`func (o *InjectorCertificateEvent) HasFingerprintSha256() bool`

HasFingerprintSha256 returns a boolean if a field has been set.

### GetInjectorName

`func (o *InjectorCertificateEvent) GetInjectorName() string`

GetInjectorName returns the InjectorName field if non-nil, zero value otherwise.

### GetInjectorNameOk

`func (o *InjectorCertificateEvent) GetInjectorNameOk() (*string, bool)`

GetInjectorNameOk returns a tuple with the InjectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectorName

`func (o *InjectorCertificateEvent) SetInjectorName(v string)`

SetInjectorName sets InjectorName field to given value.

### HasInjectorName

`func (o *InjectorCertificateEvent) HasInjectorName() bool`

HasInjectorName returns a boolean if a field has been set.

### GetInjectorNamespace

`func (o *InjectorCertificateEvent) GetInjectorNamespace() string`

GetInjectorNamespace returns the InjectorNamespace field if non-nil, zero value otherwise.

### GetInjectorNamespaceOk

`func (o *InjectorCertificateEvent) GetInjectorNamespaceOk() (*string, bool)`

GetInjectorNamespaceOk returns a tuple with the InjectorNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectorNamespace

`func (o *InjectorCertificateEvent) SetInjectorNamespace(v string)`

SetInjectorNamespace sets InjectorNamespace field to given value.

### HasInjectorNamespace

`func (o *InjectorCertificateEvent) HasInjectorNamespace() bool`

HasInjectorNamespace returns a boolean if a field has been set.

### GetInstallationId

`func (o *InjectorCertificateEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *InjectorCertificateEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *InjectorCertificateEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *InjectorCertificateEvent) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetNotAfter

`func (o *InjectorCertificateEvent) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *InjectorCertificateEvent) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *InjectorCertificateEvent) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *InjectorCertificateEvent) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetThreshold

`func (o *InjectorCertificateEvent) GetThreshold() string`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *InjectorCertificateEvent) GetThresholdOk() (*string, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *InjectorCertificateEvent) SetThreshold(v string)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *InjectorCertificateEvent) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


