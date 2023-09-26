# DynamicSecretCreateCassandra

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
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateCassandra

`func NewDynamicSecretCreateCassandra(name string, ) *DynamicSecretCreateCassandra`

NewDynamicSecretCreateCassandra instantiates a new DynamicSecretCreateCassandra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateCassandraWithDefaults

`func NewDynamicSecretCreateCassandraWithDefaults() *DynamicSecretCreateCassandra`

NewDynamicSecretCreateCassandraWithDefaults instantiates a new DynamicSecretCreateCassandra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraCreationStatements

`func (o *DynamicSecretCreateCassandra) GetCassandraCreationStatements() string`

GetCassandraCreationStatements returns the CassandraCreationStatements field if non-nil, zero value otherwise.

### GetCassandraCreationStatementsOk

`func (o *DynamicSecretCreateCassandra) GetCassandraCreationStatementsOk() (*string, bool)`

GetCassandraCreationStatementsOk returns a tuple with the CassandraCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCreationStatements

`func (o *DynamicSecretCreateCassandra) SetCassandraCreationStatements(v string)`

SetCassandraCreationStatements sets CassandraCreationStatements field to given value.

### HasCassandraCreationStatements

`func (o *DynamicSecretCreateCassandra) HasCassandraCreationStatements() bool`

HasCassandraCreationStatements returns a boolean if a field has been set.

### GetCassandraHosts

`func (o *DynamicSecretCreateCassandra) GetCassandraHosts() string`

GetCassandraHosts returns the CassandraHosts field if non-nil, zero value otherwise.

### GetCassandraHostsOk

`func (o *DynamicSecretCreateCassandra) GetCassandraHostsOk() (*string, bool)`

GetCassandraHostsOk returns a tuple with the CassandraHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraHosts

`func (o *DynamicSecretCreateCassandra) SetCassandraHosts(v string)`

SetCassandraHosts sets CassandraHosts field to given value.

### HasCassandraHosts

`func (o *DynamicSecretCreateCassandra) HasCassandraHosts() bool`

HasCassandraHosts returns a boolean if a field has been set.

### GetCassandraPassword

`func (o *DynamicSecretCreateCassandra) GetCassandraPassword() string`

GetCassandraPassword returns the CassandraPassword field if non-nil, zero value otherwise.

### GetCassandraPasswordOk

`func (o *DynamicSecretCreateCassandra) GetCassandraPasswordOk() (*string, bool)`

GetCassandraPasswordOk returns a tuple with the CassandraPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPassword

`func (o *DynamicSecretCreateCassandra) SetCassandraPassword(v string)`

SetCassandraPassword sets CassandraPassword field to given value.

### HasCassandraPassword

`func (o *DynamicSecretCreateCassandra) HasCassandraPassword() bool`

HasCassandraPassword returns a boolean if a field has been set.

### GetCassandraPort

`func (o *DynamicSecretCreateCassandra) GetCassandraPort() string`

GetCassandraPort returns the CassandraPort field if non-nil, zero value otherwise.

### GetCassandraPortOk

`func (o *DynamicSecretCreateCassandra) GetCassandraPortOk() (*string, bool)`

GetCassandraPortOk returns a tuple with the CassandraPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPort

`func (o *DynamicSecretCreateCassandra) SetCassandraPort(v string)`

SetCassandraPort sets CassandraPort field to given value.

### HasCassandraPort

`func (o *DynamicSecretCreateCassandra) HasCassandraPort() bool`

HasCassandraPort returns a boolean if a field has been set.

### GetCassandraUsername

`func (o *DynamicSecretCreateCassandra) GetCassandraUsername() string`

GetCassandraUsername returns the CassandraUsername field if non-nil, zero value otherwise.

### GetCassandraUsernameOk

`func (o *DynamicSecretCreateCassandra) GetCassandraUsernameOk() (*string, bool)`

GetCassandraUsernameOk returns a tuple with the CassandraUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraUsername

`func (o *DynamicSecretCreateCassandra) SetCassandraUsername(v string)`

SetCassandraUsername sets CassandraUsername field to given value.

### HasCassandraUsername

`func (o *DynamicSecretCreateCassandra) HasCassandraUsername() bool`

HasCassandraUsername returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateCassandra) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateCassandra) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateCassandra) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateCassandra) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateCassandra) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateCassandra) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateCassandra) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateCassandra) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateCassandra) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateCassandra) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateCassandra) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateCassandra) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateCassandra) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateCassandra) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateCassandra) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateCassandra) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateCassandra) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateCassandra) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateCassandra) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateCassandra) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateCassandra) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateCassandra) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateCassandra) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateCassandra) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateCassandra) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateCassandra) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateCassandra) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateCassandra) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateCassandra) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateCassandra) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateCassandra) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateCassandra) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateCassandra) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateCassandra) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateCassandra) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


