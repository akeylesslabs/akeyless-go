# CertificateExpirationEventsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**ExpirationsList** | Pointer to [**[]CertificateExpirationEvent**](CertificateExpirationEvent.md) | ExpirationEventsList is the default expiration events for the account | [optional] 

## Methods

### NewCertificateExpirationEventsSettings

`func NewCertificateExpirationEventsSettings() *CertificateExpirationEventsSettings`

NewCertificateExpirationEventsSettings instantiates a new CertificateExpirationEventsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateExpirationEventsSettingsWithDefaults

`func NewCertificateExpirationEventsSettingsWithDefaults() *CertificateExpirationEventsSettings`

NewCertificateExpirationEventsSettingsWithDefaults instantiates a new CertificateExpirationEventsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *CertificateExpirationEventsSettings) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CertificateExpirationEventsSettings) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CertificateExpirationEventsSettings) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CertificateExpirationEventsSettings) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetExpirationsList

`func (o *CertificateExpirationEventsSettings) GetExpirationsList() []CertificateExpirationEvent`

GetExpirationsList returns the ExpirationsList field if non-nil, zero value otherwise.

### GetExpirationsListOk

`func (o *CertificateExpirationEventsSettings) GetExpirationsListOk() (*[]CertificateExpirationEvent, bool)`

GetExpirationsListOk returns a tuple with the ExpirationsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationsList

`func (o *CertificateExpirationEventsSettings) SetExpirationsList(v []CertificateExpirationEvent)`

SetExpirationsList sets ExpirationsList field to given value.

### HasExpirationsList

`func (o *CertificateExpirationEventsSettings) HasExpirationsList() bool`

HasExpirationsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


