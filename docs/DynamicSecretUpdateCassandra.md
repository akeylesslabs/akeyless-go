# DynamicSecretUpdateCassandra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraCreationStatements** | Pointer to **string** | Cassandra creation statements | [optional] 
**CassandraHosts** | Pointer to **string** | Cassandra hosts IP or addresses, comma separated | [optional] 
**CassandraPassword** | Pointer to **string** | Cassandra superuser password | [optional] 
**CassandraPort** | Pointer to **string** | Cassandra port | [optional] [default to "9042"]
**CassandraUsername** | Pointer to **string** | Cassandra superuser username | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
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

### NewDynamicSecretUpdateCassandra

`func NewDynamicSecretUpdateCassandra(name string, ) *DynamicSecretUpdateCassandra`

NewDynamicSecretUpdateCassandra instantiates a new DynamicSecretUpdateCassandra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateCassandraWithDefaults

`func NewDynamicSecretUpdateCassandraWithDefaults() *DynamicSecretUpdateCassandra`

NewDynamicSecretUpdateCassandraWithDefaults instantiates a new DynamicSecretUpdateCassandra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraCreationStatements

`func (o *DynamicSecretUpdateCassandra) GetCassandraCreationStatements() string`

GetCassandraCreationStatements returns the CassandraCreationStatements field if non-nil, zero value otherwise.

### GetCassandraCreationStatementsOk

`func (o *DynamicSecretUpdateCassandra) GetCassandraCreationStatementsOk() (*string, bool)`

GetCassandraCreationStatementsOk returns a tuple with the CassandraCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCreationStatements

`func (o *DynamicSecretUpdateCassandra) SetCassandraCreationStatements(v string)`

SetCassandraCreationStatements sets CassandraCreationStatements field to given value.

### HasCassandraCreationStatements

`func (o *DynamicSecretUpdateCassandra) HasCassandraCreationStatements() bool`

HasCassandraCreationStatements returns a boolean if a field has been set.

### GetCassandraHosts

`func (o *DynamicSecretUpdateCassandra) GetCassandraHosts() string`

GetCassandraHosts returns the CassandraHosts field if non-nil, zero value otherwise.

### GetCassandraHostsOk

`func (o *DynamicSecretUpdateCassandra) GetCassandraHostsOk() (*string, bool)`

GetCassandraHostsOk returns a tuple with the CassandraHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraHosts

`func (o *DynamicSecretUpdateCassandra) SetCassandraHosts(v string)`

SetCassandraHosts sets CassandraHosts field to given value.

### HasCassandraHosts

`func (o *DynamicSecretUpdateCassandra) HasCassandraHosts() bool`

HasCassandraHosts returns a boolean if a field has been set.

### GetCassandraPassword

`func (o *DynamicSecretUpdateCassandra) GetCassandraPassword() string`

GetCassandraPassword returns the CassandraPassword field if non-nil, zero value otherwise.

### GetCassandraPasswordOk

`func (o *DynamicSecretUpdateCassandra) GetCassandraPasswordOk() (*string, bool)`

GetCassandraPasswordOk returns a tuple with the CassandraPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPassword

`func (o *DynamicSecretUpdateCassandra) SetCassandraPassword(v string)`

SetCassandraPassword sets CassandraPassword field to given value.

### HasCassandraPassword

`func (o *DynamicSecretUpdateCassandra) HasCassandraPassword() bool`

HasCassandraPassword returns a boolean if a field has been set.

### GetCassandraPort

`func (o *DynamicSecretUpdateCassandra) GetCassandraPort() string`

GetCassandraPort returns the CassandraPort field if non-nil, zero value otherwise.

### GetCassandraPortOk

`func (o *DynamicSecretUpdateCassandra) GetCassandraPortOk() (*string, bool)`

GetCassandraPortOk returns a tuple with the CassandraPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPort

`func (o *DynamicSecretUpdateCassandra) SetCassandraPort(v string)`

SetCassandraPort sets CassandraPort field to given value.

### HasCassandraPort

`func (o *DynamicSecretUpdateCassandra) HasCassandraPort() bool`

HasCassandraPort returns a boolean if a field has been set.

### GetCassandraUsername

`func (o *DynamicSecretUpdateCassandra) GetCassandraUsername() string`

GetCassandraUsername returns the CassandraUsername field if non-nil, zero value otherwise.

### GetCassandraUsernameOk

`func (o *DynamicSecretUpdateCassandra) GetCassandraUsernameOk() (*string, bool)`

GetCassandraUsernameOk returns a tuple with the CassandraUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraUsername

`func (o *DynamicSecretUpdateCassandra) SetCassandraUsername(v string)`

SetCassandraUsername sets CassandraUsername field to given value.

### HasCassandraUsername

`func (o *DynamicSecretUpdateCassandra) HasCassandraUsername() bool`

HasCassandraUsername returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretUpdateCassandra) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateCassandra) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateCassandra) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateCassandra) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateCassandra) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateCassandra) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateCassandra) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateCassandra) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateCassandra) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateCassandra) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateCassandra) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateCassandra) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateCassandra) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateCassandra) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateCassandra) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateCassandra) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateCassandra) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateCassandra) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateCassandra) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretUpdateCassandra) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretUpdateCassandra) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretUpdateCassandra) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretUpdateCassandra) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateCassandra) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateCassandra) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateCassandra) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateCassandra) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSsl

`func (o *DynamicSecretUpdateCassandra) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *DynamicSecretUpdateCassandra) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *DynamicSecretUpdateCassandra) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *DynamicSecretUpdateCassandra) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *DynamicSecretUpdateCassandra) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *DynamicSecretUpdateCassandra) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *DynamicSecretUpdateCassandra) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *DynamicSecretUpdateCassandra) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateCassandra) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateCassandra) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateCassandra) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateCassandra) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateCassandra) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateCassandra) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateCassandra) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateCassandra) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateCassandra) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateCassandra) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateCassandra) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateCassandra) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateCassandra) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateCassandra) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateCassandra) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateCassandra) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateCassandra) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateCassandra) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateCassandra) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateCassandra) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


