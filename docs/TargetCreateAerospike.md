# TargetCreateAerospike

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUsername** | Pointer to **string** | Username of an account with the user-admin role | [optional] 
**AerospikeClientId** | Pointer to **string** | Client ID for Aerospike Cloud authentication (relevant only for Aerospike Cloud) | [optional] 
**AerospikeClientSecret** | Pointer to **string** | Client secret for Aerospike Cloud authentication (relevant only for Aerospike Cloud) | [optional] 
**AerospikeCloud** | Pointer to **bool** | Set to &#39;true&#39; for Aerospike Cloud deployments | [optional] 
**AerospikeClusterId** | Pointer to **string** | Cloud cluster ID (relevant only for Aerospike Cloud) | [optional] 
**ClientCertificate** | Pointer to **string** | Client certificate for mTLS (mTLS only) | [optional] 
**ClientPrivateKey** | Pointer to **string** | Client private key for mTLS (mTLS only) | [optional] 
**DbServerName** | Pointer to **string** | TLS server name used to verify the certificate hostname. If empty, the Aerospike hostname is used. | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EnableMtls** | Pointer to **bool** | Enable mutual TLS authentication - requires --ssl&#x3D;true (true/false) | [optional] 
**Hostname** | Pointer to **string** | Aerospike host address and port (e.g. url.to.aerospike.db) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Namespace** | Pointer to **string** | Namespace name (relevant only for Aerospike db) | [optional] 
**Password** | Pointer to **string** | Password for the admin user | [optional] 
**Port** | Pointer to **string** | Database connection port | [optional] 
**SkipServerNameValidation** | Pointer to **string** | Skip server name verification while still validating the certificate chain (true/false). Empty means do not skip. | [optional] 
**Ssl** | Pointer to **bool** | Enable SSL encryption (true/false) | [optional] 
**SslCertificate** | Pointer to **string** | Base64-encoded SSL CA certificate from a trusted Certificate Authority (CA) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateAerospike

`func NewTargetCreateAerospike(name string, ) *TargetCreateAerospike`

NewTargetCreateAerospike instantiates a new TargetCreateAerospike object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateAerospikeWithDefaults

`func NewTargetCreateAerospikeWithDefaults() *TargetCreateAerospike`

NewTargetCreateAerospikeWithDefaults instantiates a new TargetCreateAerospike object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminUsername

`func (o *TargetCreateAerospike) GetAdminUsername() string`

GetAdminUsername returns the AdminUsername field if non-nil, zero value otherwise.

### GetAdminUsernameOk

`func (o *TargetCreateAerospike) GetAdminUsernameOk() (*string, bool)`

GetAdminUsernameOk returns a tuple with the AdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsername

`func (o *TargetCreateAerospike) SetAdminUsername(v string)`

SetAdminUsername sets AdminUsername field to given value.

### HasAdminUsername

`func (o *TargetCreateAerospike) HasAdminUsername() bool`

HasAdminUsername returns a boolean if a field has been set.

### GetAerospikeClientId

`func (o *TargetCreateAerospike) GetAerospikeClientId() string`

GetAerospikeClientId returns the AerospikeClientId field if non-nil, zero value otherwise.

### GetAerospikeClientIdOk

`func (o *TargetCreateAerospike) GetAerospikeClientIdOk() (*string, bool)`

GetAerospikeClientIdOk returns a tuple with the AerospikeClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientId

`func (o *TargetCreateAerospike) SetAerospikeClientId(v string)`

SetAerospikeClientId sets AerospikeClientId field to given value.

### HasAerospikeClientId

`func (o *TargetCreateAerospike) HasAerospikeClientId() bool`

HasAerospikeClientId returns a boolean if a field has been set.

### GetAerospikeClientSecret

`func (o *TargetCreateAerospike) GetAerospikeClientSecret() string`

GetAerospikeClientSecret returns the AerospikeClientSecret field if non-nil, zero value otherwise.

### GetAerospikeClientSecretOk

`func (o *TargetCreateAerospike) GetAerospikeClientSecretOk() (*string, bool)`

GetAerospikeClientSecretOk returns a tuple with the AerospikeClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientSecret

`func (o *TargetCreateAerospike) SetAerospikeClientSecret(v string)`

SetAerospikeClientSecret sets AerospikeClientSecret field to given value.

### HasAerospikeClientSecret

`func (o *TargetCreateAerospike) HasAerospikeClientSecret() bool`

HasAerospikeClientSecret returns a boolean if a field has been set.

### GetAerospikeCloud

`func (o *TargetCreateAerospike) GetAerospikeCloud() bool`

GetAerospikeCloud returns the AerospikeCloud field if non-nil, zero value otherwise.

### GetAerospikeCloudOk

`func (o *TargetCreateAerospike) GetAerospikeCloudOk() (*bool, bool)`

GetAerospikeCloudOk returns a tuple with the AerospikeCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeCloud

`func (o *TargetCreateAerospike) SetAerospikeCloud(v bool)`

SetAerospikeCloud sets AerospikeCloud field to given value.

### HasAerospikeCloud

`func (o *TargetCreateAerospike) HasAerospikeCloud() bool`

HasAerospikeCloud returns a boolean if a field has been set.

### GetAerospikeClusterId

`func (o *TargetCreateAerospike) GetAerospikeClusterId() string`

GetAerospikeClusterId returns the AerospikeClusterId field if non-nil, zero value otherwise.

### GetAerospikeClusterIdOk

`func (o *TargetCreateAerospike) GetAerospikeClusterIdOk() (*string, bool)`

GetAerospikeClusterIdOk returns a tuple with the AerospikeClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClusterId

`func (o *TargetCreateAerospike) SetAerospikeClusterId(v string)`

SetAerospikeClusterId sets AerospikeClusterId field to given value.

### HasAerospikeClusterId

`func (o *TargetCreateAerospike) HasAerospikeClusterId() bool`

HasAerospikeClusterId returns a boolean if a field has been set.

### GetClientCertificate

`func (o *TargetCreateAerospike) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *TargetCreateAerospike) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *TargetCreateAerospike) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *TargetCreateAerospike) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetClientPrivateKey

`func (o *TargetCreateAerospike) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *TargetCreateAerospike) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *TargetCreateAerospike) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *TargetCreateAerospike) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### GetDbServerName

`func (o *TargetCreateAerospike) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *TargetCreateAerospike) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *TargetCreateAerospike) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *TargetCreateAerospike) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *TargetCreateAerospike) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *TargetCreateAerospike) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *TargetCreateAerospike) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *TargetCreateAerospike) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateAerospike) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateAerospike) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateAerospike) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateAerospike) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableMtls

`func (o *TargetCreateAerospike) GetEnableMtls() bool`

GetEnableMtls returns the EnableMtls field if non-nil, zero value otherwise.

### GetEnableMtlsOk

`func (o *TargetCreateAerospike) GetEnableMtlsOk() (*bool, bool)`

GetEnableMtlsOk returns a tuple with the EnableMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMtls

`func (o *TargetCreateAerospike) SetEnableMtls(v bool)`

SetEnableMtls sets EnableMtls field to given value.

### HasEnableMtls

`func (o *TargetCreateAerospike) HasEnableMtls() bool`

HasEnableMtls returns a boolean if a field has been set.

### GetHostname

`func (o *TargetCreateAerospike) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *TargetCreateAerospike) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *TargetCreateAerospike) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *TargetCreateAerospike) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateAerospike) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateAerospike) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateAerospike) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateAerospike) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateAerospike) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateAerospike) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateAerospike) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateAerospike) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateAerospike) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateAerospike) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateAerospike) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateAerospike) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateAerospike) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateAerospike) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateAerospike) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *TargetCreateAerospike) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TargetCreateAerospike) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TargetCreateAerospike) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *TargetCreateAerospike) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPassword

`func (o *TargetCreateAerospike) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetCreateAerospike) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetCreateAerospike) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TargetCreateAerospike) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *TargetCreateAerospike) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetCreateAerospike) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetCreateAerospike) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetCreateAerospike) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSkipServerNameValidation

`func (o *TargetCreateAerospike) GetSkipServerNameValidation() string`

GetSkipServerNameValidation returns the SkipServerNameValidation field if non-nil, zero value otherwise.

### GetSkipServerNameValidationOk

`func (o *TargetCreateAerospike) GetSkipServerNameValidationOk() (*string, bool)`

GetSkipServerNameValidationOk returns a tuple with the SkipServerNameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipServerNameValidation

`func (o *TargetCreateAerospike) SetSkipServerNameValidation(v string)`

SetSkipServerNameValidation sets SkipServerNameValidation field to given value.

### HasSkipServerNameValidation

`func (o *TargetCreateAerospike) HasSkipServerNameValidation() bool`

HasSkipServerNameValidation returns a boolean if a field has been set.

### GetSsl

`func (o *TargetCreateAerospike) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *TargetCreateAerospike) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *TargetCreateAerospike) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *TargetCreateAerospike) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *TargetCreateAerospike) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *TargetCreateAerospike) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *TargetCreateAerospike) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *TargetCreateAerospike) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateAerospike) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateAerospike) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateAerospike) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateAerospike) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateAerospike) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateAerospike) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateAerospike) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateAerospike) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


