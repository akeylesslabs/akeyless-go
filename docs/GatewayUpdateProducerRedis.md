# GatewayUpdateProducerRedis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclRules** | Pointer to **string** | A JSON array list of redis ACL rules to attach to the created user. For available rules see the ACL CAT command https://redis.io/commands/acl-cat By default the user will have permissions to read all keys &#39;[\&quot;~*\&quot;, \&quot;+@read\&quot;]&#39; | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Host** | Pointer to **string** | Redis Host | [optional] [default to "127.0.0.1"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
**Password** | Pointer to **string** | Redis Password | [optional] 
**Port** | Pointer to **string** | Redis Port | [optional] [default to "6379"]
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Ssl** | Pointer to **bool** | Enable/Disable SSL [true/false] | [optional] [default to false]
**SslCertificate** | Pointer to **string** | SSL CA certificate in base64 encoding generated from a trusted Certificate Authority (CA) | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Redis Username | [optional] 

## Methods

### NewGatewayUpdateProducerRedis

`func NewGatewayUpdateProducerRedis(name string, ) *GatewayUpdateProducerRedis`

NewGatewayUpdateProducerRedis instantiates a new GatewayUpdateProducerRedis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerRedisWithDefaults

`func NewGatewayUpdateProducerRedisWithDefaults() *GatewayUpdateProducerRedis`

NewGatewayUpdateProducerRedisWithDefaults instantiates a new GatewayUpdateProducerRedis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclRules

`func (o *GatewayUpdateProducerRedis) GetAclRules() string`

GetAclRules returns the AclRules field if non-nil, zero value otherwise.

### GetAclRulesOk

`func (o *GatewayUpdateProducerRedis) GetAclRulesOk() (*string, bool)`

GetAclRulesOk returns a tuple with the AclRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclRules

`func (o *GatewayUpdateProducerRedis) SetAclRules(v string)`

SetAclRules sets AclRules field to given value.

### HasAclRules

`func (o *GatewayUpdateProducerRedis) HasAclRules() bool`

HasAclRules returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayUpdateProducerRedis) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerRedis) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerRedis) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerRedis) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetHost

`func (o *GatewayUpdateProducerRedis) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GatewayUpdateProducerRedis) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GatewayUpdateProducerRedis) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GatewayUpdateProducerRedis) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateProducerRedis) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateProducerRedis) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateProducerRedis) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateProducerRedis) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerRedis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerRedis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerRedis) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerRedis) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerRedis) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerRedis) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerRedis) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *GatewayUpdateProducerRedis) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayUpdateProducerRedis) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayUpdateProducerRedis) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayUpdateProducerRedis) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *GatewayUpdateProducerRedis) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayUpdateProducerRedis) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayUpdateProducerRedis) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *GatewayUpdateProducerRedis) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerRedis) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerRedis) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerRedis) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerRedis) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSsl

`func (o *GatewayUpdateProducerRedis) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *GatewayUpdateProducerRedis) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *GatewayUpdateProducerRedis) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *GatewayUpdateProducerRedis) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *GatewayUpdateProducerRedis) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *GatewayUpdateProducerRedis) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *GatewayUpdateProducerRedis) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *GatewayUpdateProducerRedis) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerRedis) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerRedis) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerRedis) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerRedis) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerRedis) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerRedis) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerRedis) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerRedis) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerRedis) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerRedis) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerRedis) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerRedis) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerRedis) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerRedis) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerRedis) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerRedis) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerRedis) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerRedis) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerRedis) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerRedis) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayUpdateProducerRedis) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayUpdateProducerRedis) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayUpdateProducerRedis) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayUpdateProducerRedis) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


