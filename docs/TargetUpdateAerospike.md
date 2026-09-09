# TargetUpdateAerospike

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
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_comment"]
**EnableMtls** | Pointer to **bool** | Enable mutual TLS authentication - requires --ssl&#x3D;true (true/false) | [optional] 
**Hostname** | Pointer to **string** | Aerospike host address and port (e.g. url.to.aerospike.db) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LockOnRead** | Pointer to **string** | Lock this secret after each successful value read | [optional] 
**LockTtl** | Pointer to **string** | Lock TTL in minutes | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Namespace** | Pointer to **string** | Namespace name (relevant only for Aerospike db) | [optional] 
**NewComment** | Pointer to **string** | Deprecated - use description | [optional] [default to "default_comment"]
**NewName** | Pointer to **string** | New target name | [optional] 
**Password** | Pointer to **string** | Password for the admin user | [optional] 
**Port** | Pointer to **string** | Database connection port | [optional] 
**RotateOnUnlock** | Pointer to **string** | Rotate this secret after it is unlocked | [optional] 
**SkipServerNameValidation** | Pointer to **string** | Skip server name verification while still validating the certificate chain (true/false). Empty means do not skip. | [optional] 
**Ssl** | Pointer to **bool** | Enable SSL encryption (true/false) | [optional] 
**SslCertificate** | Pointer to **string** | Base64-encoded SSL CA certificate from a trusted Certificate Authority (CA) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateAerospike

`func NewTargetUpdateAerospike(name string, ) *TargetUpdateAerospike`

NewTargetUpdateAerospike instantiates a new TargetUpdateAerospike object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateAerospikeWithDefaults

`func NewTargetUpdateAerospikeWithDefaults() *TargetUpdateAerospike`

NewTargetUpdateAerospikeWithDefaults instantiates a new TargetUpdateAerospike object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminUsername

`func (o *TargetUpdateAerospike) GetAdminUsername() string`

GetAdminUsername returns the AdminUsername field if non-nil, zero value otherwise.

### GetAdminUsernameOk

`func (o *TargetUpdateAerospike) GetAdminUsernameOk() (*string, bool)`

GetAdminUsernameOk returns a tuple with the AdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsername

`func (o *TargetUpdateAerospike) SetAdminUsername(v string)`

SetAdminUsername sets AdminUsername field to given value.

### HasAdminUsername

`func (o *TargetUpdateAerospike) HasAdminUsername() bool`

HasAdminUsername returns a boolean if a field has been set.

### GetAerospikeClientId

`func (o *TargetUpdateAerospike) GetAerospikeClientId() string`

GetAerospikeClientId returns the AerospikeClientId field if non-nil, zero value otherwise.

### GetAerospikeClientIdOk

`func (o *TargetUpdateAerospike) GetAerospikeClientIdOk() (*string, bool)`

GetAerospikeClientIdOk returns a tuple with the AerospikeClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientId

`func (o *TargetUpdateAerospike) SetAerospikeClientId(v string)`

SetAerospikeClientId sets AerospikeClientId field to given value.

### HasAerospikeClientId

`func (o *TargetUpdateAerospike) HasAerospikeClientId() bool`

HasAerospikeClientId returns a boolean if a field has been set.

### GetAerospikeClientSecret

`func (o *TargetUpdateAerospike) GetAerospikeClientSecret() string`

GetAerospikeClientSecret returns the AerospikeClientSecret field if non-nil, zero value otherwise.

### GetAerospikeClientSecretOk

`func (o *TargetUpdateAerospike) GetAerospikeClientSecretOk() (*string, bool)`

GetAerospikeClientSecretOk returns a tuple with the AerospikeClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientSecret

`func (o *TargetUpdateAerospike) SetAerospikeClientSecret(v string)`

SetAerospikeClientSecret sets AerospikeClientSecret field to given value.

### HasAerospikeClientSecret

`func (o *TargetUpdateAerospike) HasAerospikeClientSecret() bool`

HasAerospikeClientSecret returns a boolean if a field has been set.

### GetAerospikeCloud

`func (o *TargetUpdateAerospike) GetAerospikeCloud() bool`

GetAerospikeCloud returns the AerospikeCloud field if non-nil, zero value otherwise.

### GetAerospikeCloudOk

`func (o *TargetUpdateAerospike) GetAerospikeCloudOk() (*bool, bool)`

GetAerospikeCloudOk returns a tuple with the AerospikeCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeCloud

`func (o *TargetUpdateAerospike) SetAerospikeCloud(v bool)`

SetAerospikeCloud sets AerospikeCloud field to given value.

### HasAerospikeCloud

`func (o *TargetUpdateAerospike) HasAerospikeCloud() bool`

HasAerospikeCloud returns a boolean if a field has been set.

### GetAerospikeClusterId

`func (o *TargetUpdateAerospike) GetAerospikeClusterId() string`

GetAerospikeClusterId returns the AerospikeClusterId field if non-nil, zero value otherwise.

### GetAerospikeClusterIdOk

`func (o *TargetUpdateAerospike) GetAerospikeClusterIdOk() (*string, bool)`

GetAerospikeClusterIdOk returns a tuple with the AerospikeClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClusterId

`func (o *TargetUpdateAerospike) SetAerospikeClusterId(v string)`

SetAerospikeClusterId sets AerospikeClusterId field to given value.

### HasAerospikeClusterId

`func (o *TargetUpdateAerospike) HasAerospikeClusterId() bool`

HasAerospikeClusterId returns a boolean if a field has been set.

### GetClientCertificate

`func (o *TargetUpdateAerospike) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *TargetUpdateAerospike) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *TargetUpdateAerospike) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *TargetUpdateAerospike) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetClientPrivateKey

`func (o *TargetUpdateAerospike) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *TargetUpdateAerospike) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *TargetUpdateAerospike) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *TargetUpdateAerospike) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### GetDbServerName

`func (o *TargetUpdateAerospike) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *TargetUpdateAerospike) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *TargetUpdateAerospike) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *TargetUpdateAerospike) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *TargetUpdateAerospike) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *TargetUpdateAerospike) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *TargetUpdateAerospike) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *TargetUpdateAerospike) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateAerospike) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateAerospike) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateAerospike) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateAerospike) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableMtls

`func (o *TargetUpdateAerospike) GetEnableMtls() bool`

GetEnableMtls returns the EnableMtls field if non-nil, zero value otherwise.

### GetEnableMtlsOk

`func (o *TargetUpdateAerospike) GetEnableMtlsOk() (*bool, bool)`

GetEnableMtlsOk returns a tuple with the EnableMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMtls

`func (o *TargetUpdateAerospike) SetEnableMtls(v bool)`

SetEnableMtls sets EnableMtls field to given value.

### HasEnableMtls

`func (o *TargetUpdateAerospike) HasEnableMtls() bool`

HasEnableMtls returns a boolean if a field has been set.

### GetHostname

`func (o *TargetUpdateAerospike) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *TargetUpdateAerospike) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *TargetUpdateAerospike) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *TargetUpdateAerospike) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateAerospike) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateAerospike) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateAerospike) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateAerospike) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateAerospike) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateAerospike) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateAerospike) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateAerospike) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateAerospike) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateAerospike) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateAerospike) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateAerospike) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLockOnRead

`func (o *TargetUpdateAerospike) GetLockOnRead() string`

GetLockOnRead returns the LockOnRead field if non-nil, zero value otherwise.

### GetLockOnReadOk

`func (o *TargetUpdateAerospike) GetLockOnReadOk() (*string, bool)`

GetLockOnReadOk returns a tuple with the LockOnRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOnRead

`func (o *TargetUpdateAerospike) SetLockOnRead(v string)`

SetLockOnRead sets LockOnRead field to given value.

### HasLockOnRead

`func (o *TargetUpdateAerospike) HasLockOnRead() bool`

HasLockOnRead returns a boolean if a field has been set.

### GetLockTtl

`func (o *TargetUpdateAerospike) GetLockTtl() string`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *TargetUpdateAerospike) GetLockTtlOk() (*string, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *TargetUpdateAerospike) SetLockTtl(v string)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *TargetUpdateAerospike) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateAerospike) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateAerospike) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateAerospike) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateAerospike) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateAerospike) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateAerospike) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateAerospike) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *TargetUpdateAerospike) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TargetUpdateAerospike) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TargetUpdateAerospike) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *TargetUpdateAerospike) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNewComment

`func (o *TargetUpdateAerospike) GetNewComment() string`

GetNewComment returns the NewComment field if non-nil, zero value otherwise.

### GetNewCommentOk

`func (o *TargetUpdateAerospike) GetNewCommentOk() (*string, bool)`

GetNewCommentOk returns a tuple with the NewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewComment

`func (o *TargetUpdateAerospike) SetNewComment(v string)`

SetNewComment sets NewComment field to given value.

### HasNewComment

`func (o *TargetUpdateAerospike) HasNewComment() bool`

HasNewComment returns a boolean if a field has been set.

### GetNewName

`func (o *TargetUpdateAerospike) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateAerospike) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateAerospike) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateAerospike) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *TargetUpdateAerospike) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetUpdateAerospike) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetUpdateAerospike) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TargetUpdateAerospike) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *TargetUpdateAerospike) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetUpdateAerospike) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetUpdateAerospike) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetUpdateAerospike) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRotateOnUnlock

`func (o *TargetUpdateAerospike) GetRotateOnUnlock() string`

GetRotateOnUnlock returns the RotateOnUnlock field if non-nil, zero value otherwise.

### GetRotateOnUnlockOk

`func (o *TargetUpdateAerospike) GetRotateOnUnlockOk() (*string, bool)`

GetRotateOnUnlockOk returns a tuple with the RotateOnUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateOnUnlock

`func (o *TargetUpdateAerospike) SetRotateOnUnlock(v string)`

SetRotateOnUnlock sets RotateOnUnlock field to given value.

### HasRotateOnUnlock

`func (o *TargetUpdateAerospike) HasRotateOnUnlock() bool`

HasRotateOnUnlock returns a boolean if a field has been set.

### GetSkipServerNameValidation

`func (o *TargetUpdateAerospike) GetSkipServerNameValidation() string`

GetSkipServerNameValidation returns the SkipServerNameValidation field if non-nil, zero value otherwise.

### GetSkipServerNameValidationOk

`func (o *TargetUpdateAerospike) GetSkipServerNameValidationOk() (*string, bool)`

GetSkipServerNameValidationOk returns a tuple with the SkipServerNameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipServerNameValidation

`func (o *TargetUpdateAerospike) SetSkipServerNameValidation(v string)`

SetSkipServerNameValidation sets SkipServerNameValidation field to given value.

### HasSkipServerNameValidation

`func (o *TargetUpdateAerospike) HasSkipServerNameValidation() bool`

HasSkipServerNameValidation returns a boolean if a field has been set.

### GetSsl

`func (o *TargetUpdateAerospike) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *TargetUpdateAerospike) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *TargetUpdateAerospike) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *TargetUpdateAerospike) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *TargetUpdateAerospike) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *TargetUpdateAerospike) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *TargetUpdateAerospike) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *TargetUpdateAerospike) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateAerospike) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateAerospike) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateAerospike) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateAerospike) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateAerospike) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateAerospike) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateAerospike) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateAerospike) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


