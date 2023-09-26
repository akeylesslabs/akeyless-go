# DsCreateCassandra

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

### NewDsCreateCassandra

`func NewDsCreateCassandra(name string, ) *DsCreateCassandra`

NewDsCreateCassandra instantiates a new DsCreateCassandra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsCreateCassandraWithDefaults

`func NewDsCreateCassandraWithDefaults() *DsCreateCassandra`

NewDsCreateCassandraWithDefaults instantiates a new DsCreateCassandra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraCreationStatements

`func (o *DsCreateCassandra) GetCassandraCreationStatements() string`

GetCassandraCreationStatements returns the CassandraCreationStatements field if non-nil, zero value otherwise.

### GetCassandraCreationStatementsOk

`func (o *DsCreateCassandra) GetCassandraCreationStatementsOk() (*string, bool)`

GetCassandraCreationStatementsOk returns a tuple with the CassandraCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCreationStatements

`func (o *DsCreateCassandra) SetCassandraCreationStatements(v string)`

SetCassandraCreationStatements sets CassandraCreationStatements field to given value.

### HasCassandraCreationStatements

`func (o *DsCreateCassandra) HasCassandraCreationStatements() bool`

HasCassandraCreationStatements returns a boolean if a field has been set.

### GetCassandraHosts

`func (o *DsCreateCassandra) GetCassandraHosts() string`

GetCassandraHosts returns the CassandraHosts field if non-nil, zero value otherwise.

### GetCassandraHostsOk

`func (o *DsCreateCassandra) GetCassandraHostsOk() (*string, bool)`

GetCassandraHostsOk returns a tuple with the CassandraHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraHosts

`func (o *DsCreateCassandra) SetCassandraHosts(v string)`

SetCassandraHosts sets CassandraHosts field to given value.

### HasCassandraHosts

`func (o *DsCreateCassandra) HasCassandraHosts() bool`

HasCassandraHosts returns a boolean if a field has been set.

### GetCassandraPassword

`func (o *DsCreateCassandra) GetCassandraPassword() string`

GetCassandraPassword returns the CassandraPassword field if non-nil, zero value otherwise.

### GetCassandraPasswordOk

`func (o *DsCreateCassandra) GetCassandraPasswordOk() (*string, bool)`

GetCassandraPasswordOk returns a tuple with the CassandraPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPassword

`func (o *DsCreateCassandra) SetCassandraPassword(v string)`

SetCassandraPassword sets CassandraPassword field to given value.

### HasCassandraPassword

`func (o *DsCreateCassandra) HasCassandraPassword() bool`

HasCassandraPassword returns a boolean if a field has been set.

### GetCassandraPort

`func (o *DsCreateCassandra) GetCassandraPort() string`

GetCassandraPort returns the CassandraPort field if non-nil, zero value otherwise.

### GetCassandraPortOk

`func (o *DsCreateCassandra) GetCassandraPortOk() (*string, bool)`

GetCassandraPortOk returns a tuple with the CassandraPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPort

`func (o *DsCreateCassandra) SetCassandraPort(v string)`

SetCassandraPort sets CassandraPort field to given value.

### HasCassandraPort

`func (o *DsCreateCassandra) HasCassandraPort() bool`

HasCassandraPort returns a boolean if a field has been set.

### GetCassandraUsername

`func (o *DsCreateCassandra) GetCassandraUsername() string`

GetCassandraUsername returns the CassandraUsername field if non-nil, zero value otherwise.

### GetCassandraUsernameOk

`func (o *DsCreateCassandra) GetCassandraUsernameOk() (*string, bool)`

GetCassandraUsernameOk returns a tuple with the CassandraUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraUsername

`func (o *DsCreateCassandra) SetCassandraUsername(v string)`

SetCassandraUsername sets CassandraUsername field to given value.

### HasCassandraUsername

`func (o *DsCreateCassandra) HasCassandraUsername() bool`

HasCassandraUsername returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DsCreateCassandra) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DsCreateCassandra) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DsCreateCassandra) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DsCreateCassandra) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DsCreateCassandra) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DsCreateCassandra) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DsCreateCassandra) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DsCreateCassandra) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DsCreateCassandra) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DsCreateCassandra) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DsCreateCassandra) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DsCreateCassandra) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DsCreateCassandra) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DsCreateCassandra) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DsCreateCassandra) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *DsCreateCassandra) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DsCreateCassandra) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DsCreateCassandra) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DsCreateCassandra) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DsCreateCassandra) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DsCreateCassandra) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DsCreateCassandra) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DsCreateCassandra) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DsCreateCassandra) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DsCreateCassandra) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DsCreateCassandra) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DsCreateCassandra) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DsCreateCassandra) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DsCreateCassandra) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DsCreateCassandra) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DsCreateCassandra) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DsCreateCassandra) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DsCreateCassandra) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DsCreateCassandra) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DsCreateCassandra) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


