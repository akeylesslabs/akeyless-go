# GatewayUpdateProducerCassandra

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
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayUpdateProducerCassandra

`func NewGatewayUpdateProducerCassandra(name string, ) *GatewayUpdateProducerCassandra`

NewGatewayUpdateProducerCassandra instantiates a new GatewayUpdateProducerCassandra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerCassandraWithDefaults

`func NewGatewayUpdateProducerCassandraWithDefaults() *GatewayUpdateProducerCassandra`

NewGatewayUpdateProducerCassandraWithDefaults instantiates a new GatewayUpdateProducerCassandra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraCreationStatements

`func (o *GatewayUpdateProducerCassandra) GetCassandraCreationStatements() string`

GetCassandraCreationStatements returns the CassandraCreationStatements field if non-nil, zero value otherwise.

### GetCassandraCreationStatementsOk

`func (o *GatewayUpdateProducerCassandra) GetCassandraCreationStatementsOk() (*string, bool)`

GetCassandraCreationStatementsOk returns a tuple with the CassandraCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCreationStatements

`func (o *GatewayUpdateProducerCassandra) SetCassandraCreationStatements(v string)`

SetCassandraCreationStatements sets CassandraCreationStatements field to given value.

### HasCassandraCreationStatements

`func (o *GatewayUpdateProducerCassandra) HasCassandraCreationStatements() bool`

HasCassandraCreationStatements returns a boolean if a field has been set.

### GetCassandraHosts

`func (o *GatewayUpdateProducerCassandra) GetCassandraHosts() string`

GetCassandraHosts returns the CassandraHosts field if non-nil, zero value otherwise.

### GetCassandraHostsOk

`func (o *GatewayUpdateProducerCassandra) GetCassandraHostsOk() (*string, bool)`

GetCassandraHostsOk returns a tuple with the CassandraHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraHosts

`func (o *GatewayUpdateProducerCassandra) SetCassandraHosts(v string)`

SetCassandraHosts sets CassandraHosts field to given value.

### HasCassandraHosts

`func (o *GatewayUpdateProducerCassandra) HasCassandraHosts() bool`

HasCassandraHosts returns a boolean if a field has been set.

### GetCassandraPassword

`func (o *GatewayUpdateProducerCassandra) GetCassandraPassword() string`

GetCassandraPassword returns the CassandraPassword field if non-nil, zero value otherwise.

### GetCassandraPasswordOk

`func (o *GatewayUpdateProducerCassandra) GetCassandraPasswordOk() (*string, bool)`

GetCassandraPasswordOk returns a tuple with the CassandraPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPassword

`func (o *GatewayUpdateProducerCassandra) SetCassandraPassword(v string)`

SetCassandraPassword sets CassandraPassword field to given value.

### HasCassandraPassword

`func (o *GatewayUpdateProducerCassandra) HasCassandraPassword() bool`

HasCassandraPassword returns a boolean if a field has been set.

### GetCassandraPort

`func (o *GatewayUpdateProducerCassandra) GetCassandraPort() string`

GetCassandraPort returns the CassandraPort field if non-nil, zero value otherwise.

### GetCassandraPortOk

`func (o *GatewayUpdateProducerCassandra) GetCassandraPortOk() (*string, bool)`

GetCassandraPortOk returns a tuple with the CassandraPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPort

`func (o *GatewayUpdateProducerCassandra) SetCassandraPort(v string)`

SetCassandraPort sets CassandraPort field to given value.

### HasCassandraPort

`func (o *GatewayUpdateProducerCassandra) HasCassandraPort() bool`

HasCassandraPort returns a boolean if a field has been set.

### GetCassandraUsername

`func (o *GatewayUpdateProducerCassandra) GetCassandraUsername() string`

GetCassandraUsername returns the CassandraUsername field if non-nil, zero value otherwise.

### GetCassandraUsernameOk

`func (o *GatewayUpdateProducerCassandra) GetCassandraUsernameOk() (*string, bool)`

GetCassandraUsernameOk returns a tuple with the CassandraUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraUsername

`func (o *GatewayUpdateProducerCassandra) SetCassandraUsername(v string)`

SetCassandraUsername sets CassandraUsername field to given value.

### HasCassandraUsername

`func (o *GatewayUpdateProducerCassandra) HasCassandraUsername() bool`

HasCassandraUsername returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayUpdateProducerCassandra) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerCassandra) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerCassandra) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerCassandra) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateProducerCassandra) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateProducerCassandra) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateProducerCassandra) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateProducerCassandra) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerCassandra) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerCassandra) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerCassandra) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerCassandra) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerCassandra) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerCassandra) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerCassandra) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerCassandra) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerCassandra) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerCassandra) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerCassandra) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerCassandra) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerCassandra) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerCassandra) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerCassandra) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerCassandra) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerCassandra) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerCassandra) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerCassandra) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerCassandra) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerCassandra) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerCassandra) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerCassandra) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerCassandra) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerCassandra) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerCassandra) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerCassandra) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerCassandra) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerCassandra) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerCassandra) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerCassandra) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


