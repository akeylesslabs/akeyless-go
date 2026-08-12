# AerospikeTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AerospikeAdminUsername** | Pointer to **string** |  | [optional] 
**AerospikeClientCertificate** | Pointer to **string** |  | [optional] 
**AerospikeClientId** | Pointer to **string** |  | [optional] 
**AerospikeClientPrivateKey** | Pointer to **string** |  | [optional] 
**AerospikeClientSecret** | Pointer to **string** |  | [optional] 
**AerospikeCloud** | Pointer to **bool** |  | [optional] 
**AerospikeClusterId** | Pointer to **string** |  | [optional] 
**AerospikeDbServerName** | Pointer to **string** |  | [optional] 
**AerospikeEnableMtls** | Pointer to **bool** |  | [optional] 
**AerospikeHostname** | Pointer to **string** |  | [optional] 
**AerospikeNamespace** | Pointer to **string** |  | [optional] 
**AerospikePassword** | Pointer to **string** |  | [optional] 
**AerospikePort** | Pointer to **string** |  | [optional] 
**AerospikeSkipServerNameValidation** | Pointer to **string** |  | [optional] 
**AerospikeSslConnectionCertificate** | Pointer to **string** |  | [optional] 
**AerospikeSslConnectionMode** | Pointer to **bool** |  | [optional] 

## Methods

### NewAerospikeTargetDetails

`func NewAerospikeTargetDetails() *AerospikeTargetDetails`

NewAerospikeTargetDetails instantiates a new AerospikeTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAerospikeTargetDetailsWithDefaults

`func NewAerospikeTargetDetailsWithDefaults() *AerospikeTargetDetails`

NewAerospikeTargetDetailsWithDefaults instantiates a new AerospikeTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAerospikeAdminUsername

`func (o *AerospikeTargetDetails) GetAerospikeAdminUsername() string`

GetAerospikeAdminUsername returns the AerospikeAdminUsername field if non-nil, zero value otherwise.

### GetAerospikeAdminUsernameOk

`func (o *AerospikeTargetDetails) GetAerospikeAdminUsernameOk() (*string, bool)`

GetAerospikeAdminUsernameOk returns a tuple with the AerospikeAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeAdminUsername

`func (o *AerospikeTargetDetails) SetAerospikeAdminUsername(v string)`

SetAerospikeAdminUsername sets AerospikeAdminUsername field to given value.

### HasAerospikeAdminUsername

`func (o *AerospikeTargetDetails) HasAerospikeAdminUsername() bool`

HasAerospikeAdminUsername returns a boolean if a field has been set.

### GetAerospikeClientCertificate

`func (o *AerospikeTargetDetails) GetAerospikeClientCertificate() string`

GetAerospikeClientCertificate returns the AerospikeClientCertificate field if non-nil, zero value otherwise.

### GetAerospikeClientCertificateOk

`func (o *AerospikeTargetDetails) GetAerospikeClientCertificateOk() (*string, bool)`

GetAerospikeClientCertificateOk returns a tuple with the AerospikeClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientCertificate

`func (o *AerospikeTargetDetails) SetAerospikeClientCertificate(v string)`

SetAerospikeClientCertificate sets AerospikeClientCertificate field to given value.

### HasAerospikeClientCertificate

`func (o *AerospikeTargetDetails) HasAerospikeClientCertificate() bool`

HasAerospikeClientCertificate returns a boolean if a field has been set.

### GetAerospikeClientId

`func (o *AerospikeTargetDetails) GetAerospikeClientId() string`

GetAerospikeClientId returns the AerospikeClientId field if non-nil, zero value otherwise.

### GetAerospikeClientIdOk

`func (o *AerospikeTargetDetails) GetAerospikeClientIdOk() (*string, bool)`

GetAerospikeClientIdOk returns a tuple with the AerospikeClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientId

`func (o *AerospikeTargetDetails) SetAerospikeClientId(v string)`

SetAerospikeClientId sets AerospikeClientId field to given value.

### HasAerospikeClientId

`func (o *AerospikeTargetDetails) HasAerospikeClientId() bool`

HasAerospikeClientId returns a boolean if a field has been set.

### GetAerospikeClientPrivateKey

`func (o *AerospikeTargetDetails) GetAerospikeClientPrivateKey() string`

GetAerospikeClientPrivateKey returns the AerospikeClientPrivateKey field if non-nil, zero value otherwise.

### GetAerospikeClientPrivateKeyOk

`func (o *AerospikeTargetDetails) GetAerospikeClientPrivateKeyOk() (*string, bool)`

GetAerospikeClientPrivateKeyOk returns a tuple with the AerospikeClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientPrivateKey

`func (o *AerospikeTargetDetails) SetAerospikeClientPrivateKey(v string)`

SetAerospikeClientPrivateKey sets AerospikeClientPrivateKey field to given value.

### HasAerospikeClientPrivateKey

`func (o *AerospikeTargetDetails) HasAerospikeClientPrivateKey() bool`

HasAerospikeClientPrivateKey returns a boolean if a field has been set.

### GetAerospikeClientSecret

`func (o *AerospikeTargetDetails) GetAerospikeClientSecret() string`

GetAerospikeClientSecret returns the AerospikeClientSecret field if non-nil, zero value otherwise.

### GetAerospikeClientSecretOk

`func (o *AerospikeTargetDetails) GetAerospikeClientSecretOk() (*string, bool)`

GetAerospikeClientSecretOk returns a tuple with the AerospikeClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientSecret

`func (o *AerospikeTargetDetails) SetAerospikeClientSecret(v string)`

SetAerospikeClientSecret sets AerospikeClientSecret field to given value.

### HasAerospikeClientSecret

`func (o *AerospikeTargetDetails) HasAerospikeClientSecret() bool`

HasAerospikeClientSecret returns a boolean if a field has been set.

### GetAerospikeCloud

`func (o *AerospikeTargetDetails) GetAerospikeCloud() bool`

GetAerospikeCloud returns the AerospikeCloud field if non-nil, zero value otherwise.

### GetAerospikeCloudOk

`func (o *AerospikeTargetDetails) GetAerospikeCloudOk() (*bool, bool)`

GetAerospikeCloudOk returns a tuple with the AerospikeCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeCloud

`func (o *AerospikeTargetDetails) SetAerospikeCloud(v bool)`

SetAerospikeCloud sets AerospikeCloud field to given value.

### HasAerospikeCloud

`func (o *AerospikeTargetDetails) HasAerospikeCloud() bool`

HasAerospikeCloud returns a boolean if a field has been set.

### GetAerospikeClusterId

`func (o *AerospikeTargetDetails) GetAerospikeClusterId() string`

GetAerospikeClusterId returns the AerospikeClusterId field if non-nil, zero value otherwise.

### GetAerospikeClusterIdOk

`func (o *AerospikeTargetDetails) GetAerospikeClusterIdOk() (*string, bool)`

GetAerospikeClusterIdOk returns a tuple with the AerospikeClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClusterId

`func (o *AerospikeTargetDetails) SetAerospikeClusterId(v string)`

SetAerospikeClusterId sets AerospikeClusterId field to given value.

### HasAerospikeClusterId

`func (o *AerospikeTargetDetails) HasAerospikeClusterId() bool`

HasAerospikeClusterId returns a boolean if a field has been set.

### GetAerospikeDbServerName

`func (o *AerospikeTargetDetails) GetAerospikeDbServerName() string`

GetAerospikeDbServerName returns the AerospikeDbServerName field if non-nil, zero value otherwise.

### GetAerospikeDbServerNameOk

`func (o *AerospikeTargetDetails) GetAerospikeDbServerNameOk() (*string, bool)`

GetAerospikeDbServerNameOk returns a tuple with the AerospikeDbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeDbServerName

`func (o *AerospikeTargetDetails) SetAerospikeDbServerName(v string)`

SetAerospikeDbServerName sets AerospikeDbServerName field to given value.

### HasAerospikeDbServerName

`func (o *AerospikeTargetDetails) HasAerospikeDbServerName() bool`

HasAerospikeDbServerName returns a boolean if a field has been set.

### GetAerospikeEnableMtls

`func (o *AerospikeTargetDetails) GetAerospikeEnableMtls() bool`

GetAerospikeEnableMtls returns the AerospikeEnableMtls field if non-nil, zero value otherwise.

### GetAerospikeEnableMtlsOk

`func (o *AerospikeTargetDetails) GetAerospikeEnableMtlsOk() (*bool, bool)`

GetAerospikeEnableMtlsOk returns a tuple with the AerospikeEnableMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeEnableMtls

`func (o *AerospikeTargetDetails) SetAerospikeEnableMtls(v bool)`

SetAerospikeEnableMtls sets AerospikeEnableMtls field to given value.

### HasAerospikeEnableMtls

`func (o *AerospikeTargetDetails) HasAerospikeEnableMtls() bool`

HasAerospikeEnableMtls returns a boolean if a field has been set.

### GetAerospikeHostname

`func (o *AerospikeTargetDetails) GetAerospikeHostname() string`

GetAerospikeHostname returns the AerospikeHostname field if non-nil, zero value otherwise.

### GetAerospikeHostnameOk

`func (o *AerospikeTargetDetails) GetAerospikeHostnameOk() (*string, bool)`

GetAerospikeHostnameOk returns a tuple with the AerospikeHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeHostname

`func (o *AerospikeTargetDetails) SetAerospikeHostname(v string)`

SetAerospikeHostname sets AerospikeHostname field to given value.

### HasAerospikeHostname

`func (o *AerospikeTargetDetails) HasAerospikeHostname() bool`

HasAerospikeHostname returns a boolean if a field has been set.

### GetAerospikeNamespace

`func (o *AerospikeTargetDetails) GetAerospikeNamespace() string`

GetAerospikeNamespace returns the AerospikeNamespace field if non-nil, zero value otherwise.

### GetAerospikeNamespaceOk

`func (o *AerospikeTargetDetails) GetAerospikeNamespaceOk() (*string, bool)`

GetAerospikeNamespaceOk returns a tuple with the AerospikeNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeNamespace

`func (o *AerospikeTargetDetails) SetAerospikeNamespace(v string)`

SetAerospikeNamespace sets AerospikeNamespace field to given value.

### HasAerospikeNamespace

`func (o *AerospikeTargetDetails) HasAerospikeNamespace() bool`

HasAerospikeNamespace returns a boolean if a field has been set.

### GetAerospikePassword

`func (o *AerospikeTargetDetails) GetAerospikePassword() string`

GetAerospikePassword returns the AerospikePassword field if non-nil, zero value otherwise.

### GetAerospikePasswordOk

`func (o *AerospikeTargetDetails) GetAerospikePasswordOk() (*string, bool)`

GetAerospikePasswordOk returns a tuple with the AerospikePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikePassword

`func (o *AerospikeTargetDetails) SetAerospikePassword(v string)`

SetAerospikePassword sets AerospikePassword field to given value.

### HasAerospikePassword

`func (o *AerospikeTargetDetails) HasAerospikePassword() bool`

HasAerospikePassword returns a boolean if a field has been set.

### GetAerospikePort

`func (o *AerospikeTargetDetails) GetAerospikePort() string`

GetAerospikePort returns the AerospikePort field if non-nil, zero value otherwise.

### GetAerospikePortOk

`func (o *AerospikeTargetDetails) GetAerospikePortOk() (*string, bool)`

GetAerospikePortOk returns a tuple with the AerospikePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikePort

`func (o *AerospikeTargetDetails) SetAerospikePort(v string)`

SetAerospikePort sets AerospikePort field to given value.

### HasAerospikePort

`func (o *AerospikeTargetDetails) HasAerospikePort() bool`

HasAerospikePort returns a boolean if a field has been set.

### GetAerospikeSkipServerNameValidation

`func (o *AerospikeTargetDetails) GetAerospikeSkipServerNameValidation() string`

GetAerospikeSkipServerNameValidation returns the AerospikeSkipServerNameValidation field if non-nil, zero value otherwise.

### GetAerospikeSkipServerNameValidationOk

`func (o *AerospikeTargetDetails) GetAerospikeSkipServerNameValidationOk() (*string, bool)`

GetAerospikeSkipServerNameValidationOk returns a tuple with the AerospikeSkipServerNameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeSkipServerNameValidation

`func (o *AerospikeTargetDetails) SetAerospikeSkipServerNameValidation(v string)`

SetAerospikeSkipServerNameValidation sets AerospikeSkipServerNameValidation field to given value.

### HasAerospikeSkipServerNameValidation

`func (o *AerospikeTargetDetails) HasAerospikeSkipServerNameValidation() bool`

HasAerospikeSkipServerNameValidation returns a boolean if a field has been set.

### GetAerospikeSslConnectionCertificate

`func (o *AerospikeTargetDetails) GetAerospikeSslConnectionCertificate() string`

GetAerospikeSslConnectionCertificate returns the AerospikeSslConnectionCertificate field if non-nil, zero value otherwise.

### GetAerospikeSslConnectionCertificateOk

`func (o *AerospikeTargetDetails) GetAerospikeSslConnectionCertificateOk() (*string, bool)`

GetAerospikeSslConnectionCertificateOk returns a tuple with the AerospikeSslConnectionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeSslConnectionCertificate

`func (o *AerospikeTargetDetails) SetAerospikeSslConnectionCertificate(v string)`

SetAerospikeSslConnectionCertificate sets AerospikeSslConnectionCertificate field to given value.

### HasAerospikeSslConnectionCertificate

`func (o *AerospikeTargetDetails) HasAerospikeSslConnectionCertificate() bool`

HasAerospikeSslConnectionCertificate returns a boolean if a field has been set.

### GetAerospikeSslConnectionMode

`func (o *AerospikeTargetDetails) GetAerospikeSslConnectionMode() bool`

GetAerospikeSslConnectionMode returns the AerospikeSslConnectionMode field if non-nil, zero value otherwise.

### GetAerospikeSslConnectionModeOk

`func (o *AerospikeTargetDetails) GetAerospikeSslConnectionModeOk() (*bool, bool)`

GetAerospikeSslConnectionModeOk returns a tuple with the AerospikeSslConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeSslConnectionMode

`func (o *AerospikeTargetDetails) SetAerospikeSslConnectionMode(v bool)`

SetAerospikeSslConnectionMode sets AerospikeSslConnectionMode field to given value.

### HasAerospikeSslConnectionMode

`func (o *AerospikeTargetDetails) HasAerospikeSslConnectionMode() bool`

HasAerospikeSslConnectionMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


