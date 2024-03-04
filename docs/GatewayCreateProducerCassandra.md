# GatewayCreateProducerCassandra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraCreationStatements** | Pointer to **string** | Cassandra creation statements | [optional] 
**CassandraHosts** | Pointer to **string** | Cassandra hosts IP or addresses, comma separated | [optional] 
**CassandraPassword** | Pointer to **string** | Cassandra superuser password | [optional] 
**CassandraPort** | Pointer to **string** | Cassandra port | [optional] [default to "9042"]
**CassandraUsername** | Pointer to **string** | Cassandra superuser username | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Ssl** | Pointer to **bool** | Enable/Disable SSL [true/false] | [optional] [default to false]
**SslCertificate** | Pointer to **string** | SSL CA certificate in base64 encoding generated from a trusted Certificate Authority (CA) | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerCassandra

`func NewGatewayCreateProducerCassandra(name string, ) *GatewayCreateProducerCassandra`

NewGatewayCreateProducerCassandra instantiates a new GatewayCreateProducerCassandra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerCassandraWithDefaults

`func NewGatewayCreateProducerCassandraWithDefaults() *GatewayCreateProducerCassandra`

NewGatewayCreateProducerCassandraWithDefaults instantiates a new GatewayCreateProducerCassandra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraCreationStatements

`func (o *GatewayCreateProducerCassandra) GetCassandraCreationStatements() string`

GetCassandraCreationStatements returns the CassandraCreationStatements field if non-nil, zero value otherwise.

### GetCassandraCreationStatementsOk

`func (o *GatewayCreateProducerCassandra) GetCassandraCreationStatementsOk() (*string, bool)`

GetCassandraCreationStatementsOk returns a tuple with the CassandraCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCreationStatements

`func (o *GatewayCreateProducerCassandra) SetCassandraCreationStatements(v string)`

SetCassandraCreationStatements sets CassandraCreationStatements field to given value.

### HasCassandraCreationStatements

`func (o *GatewayCreateProducerCassandra) HasCassandraCreationStatements() bool`

HasCassandraCreationStatements returns a boolean if a field has been set.

### GetCassandraHosts

`func (o *GatewayCreateProducerCassandra) GetCassandraHosts() string`

GetCassandraHosts returns the CassandraHosts field if non-nil, zero value otherwise.

### GetCassandraHostsOk

`func (o *GatewayCreateProducerCassandra) GetCassandraHostsOk() (*string, bool)`

GetCassandraHostsOk returns a tuple with the CassandraHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraHosts

`func (o *GatewayCreateProducerCassandra) SetCassandraHosts(v string)`

SetCassandraHosts sets CassandraHosts field to given value.

### HasCassandraHosts

`func (o *GatewayCreateProducerCassandra) HasCassandraHosts() bool`

HasCassandraHosts returns a boolean if a field has been set.

### GetCassandraPassword

`func (o *GatewayCreateProducerCassandra) GetCassandraPassword() string`

GetCassandraPassword returns the CassandraPassword field if non-nil, zero value otherwise.

### GetCassandraPasswordOk

`func (o *GatewayCreateProducerCassandra) GetCassandraPasswordOk() (*string, bool)`

GetCassandraPasswordOk returns a tuple with the CassandraPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPassword

`func (o *GatewayCreateProducerCassandra) SetCassandraPassword(v string)`

SetCassandraPassword sets CassandraPassword field to given value.

### HasCassandraPassword

`func (o *GatewayCreateProducerCassandra) HasCassandraPassword() bool`

HasCassandraPassword returns a boolean if a field has been set.

### GetCassandraPort

`func (o *GatewayCreateProducerCassandra) GetCassandraPort() string`

GetCassandraPort returns the CassandraPort field if non-nil, zero value otherwise.

### GetCassandraPortOk

`func (o *GatewayCreateProducerCassandra) GetCassandraPortOk() (*string, bool)`

GetCassandraPortOk returns a tuple with the CassandraPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPort

`func (o *GatewayCreateProducerCassandra) SetCassandraPort(v string)`

SetCassandraPort sets CassandraPort field to given value.

### HasCassandraPort

`func (o *GatewayCreateProducerCassandra) HasCassandraPort() bool`

HasCassandraPort returns a boolean if a field has been set.

### GetCassandraUsername

`func (o *GatewayCreateProducerCassandra) GetCassandraUsername() string`

GetCassandraUsername returns the CassandraUsername field if non-nil, zero value otherwise.

### GetCassandraUsernameOk

`func (o *GatewayCreateProducerCassandra) GetCassandraUsernameOk() (*string, bool)`

GetCassandraUsernameOk returns a tuple with the CassandraUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraUsername

`func (o *GatewayCreateProducerCassandra) SetCassandraUsername(v string)`

SetCassandraUsername sets CassandraUsername field to given value.

### HasCassandraUsername

`func (o *GatewayCreateProducerCassandra) HasCassandraUsername() bool`

HasCassandraUsername returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayCreateProducerCassandra) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayCreateProducerCassandra) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayCreateProducerCassandra) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayCreateProducerCassandra) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateProducerCassandra) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateProducerCassandra) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateProducerCassandra) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateProducerCassandra) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerCassandra) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerCassandra) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerCassandra) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *GatewayCreateProducerCassandra) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *GatewayCreateProducerCassandra) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *GatewayCreateProducerCassandra) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *GatewayCreateProducerCassandra) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerCassandra) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerCassandra) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerCassandra) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerCassandra) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSsl

`func (o *GatewayCreateProducerCassandra) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *GatewayCreateProducerCassandra) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *GatewayCreateProducerCassandra) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *GatewayCreateProducerCassandra) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *GatewayCreateProducerCassandra) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *GatewayCreateProducerCassandra) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *GatewayCreateProducerCassandra) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *GatewayCreateProducerCassandra) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerCassandra) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerCassandra) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerCassandra) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerCassandra) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerCassandra) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerCassandra) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerCassandra) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerCassandra) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerCassandra) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerCassandra) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerCassandra) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerCassandra) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerCassandra) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerCassandra) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerCassandra) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerCassandra) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerCassandra) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerCassandra) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerCassandra) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerCassandra) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


