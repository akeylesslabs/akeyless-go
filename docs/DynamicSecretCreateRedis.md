# DynamicSecretCreateRedis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclRules** | Pointer to **string** | A JSON array list of redis ACL rules to attach to the created user. For available rules see the ACL CAT command https://redis.io/commands/acl-cat By default the user will have permissions to read all keys &#39;[\&quot;~*\&quot;, \&quot;+@read\&quot;]&#39; | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Host** | Pointer to **string** | Redis Host | [optional] [default to "127.0.0.1"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**Password** | Pointer to **string** | Redis Password | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
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

### NewDynamicSecretCreateRedis

`func NewDynamicSecretCreateRedis(name string, ) *DynamicSecretCreateRedis`

NewDynamicSecretCreateRedis instantiates a new DynamicSecretCreateRedis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateRedisWithDefaults

`func NewDynamicSecretCreateRedisWithDefaults() *DynamicSecretCreateRedis`

NewDynamicSecretCreateRedisWithDefaults instantiates a new DynamicSecretCreateRedis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclRules

`func (o *DynamicSecretCreateRedis) GetAclRules() string`

GetAclRules returns the AclRules field if non-nil, zero value otherwise.

### GetAclRulesOk

`func (o *DynamicSecretCreateRedis) GetAclRulesOk() (*string, bool)`

GetAclRulesOk returns a tuple with the AclRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclRules

`func (o *DynamicSecretCreateRedis) SetAclRules(v string)`

SetAclRules sets AclRules field to given value.

### HasAclRules

`func (o *DynamicSecretCreateRedis) HasAclRules() bool`

HasAclRules returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateRedis) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateRedis) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateRedis) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateRedis) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateRedis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateRedis) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateRedis) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateRedis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *DynamicSecretCreateRedis) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DynamicSecretCreateRedis) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DynamicSecretCreateRedis) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DynamicSecretCreateRedis) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateRedis) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateRedis) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateRedis) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateRedis) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateRedis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateRedis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateRedis) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *DynamicSecretCreateRedis) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DynamicSecretCreateRedis) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DynamicSecretCreateRedis) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DynamicSecretCreateRedis) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretCreateRedis) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretCreateRedis) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretCreateRedis) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretCreateRedis) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetPort

`func (o *DynamicSecretCreateRedis) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DynamicSecretCreateRedis) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DynamicSecretCreateRedis) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *DynamicSecretCreateRedis) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateRedis) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateRedis) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateRedis) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateRedis) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSsl

`func (o *DynamicSecretCreateRedis) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *DynamicSecretCreateRedis) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *DynamicSecretCreateRedis) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *DynamicSecretCreateRedis) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *DynamicSecretCreateRedis) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *DynamicSecretCreateRedis) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *DynamicSecretCreateRedis) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *DynamicSecretCreateRedis) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateRedis) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateRedis) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateRedis) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateRedis) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateRedis) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateRedis) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateRedis) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateRedis) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateRedis) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateRedis) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateRedis) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateRedis) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateRedis) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateRedis) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateRedis) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateRedis) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateRedis) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateRedis) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateRedis) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateRedis) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *DynamicSecretCreateRedis) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DynamicSecretCreateRedis) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DynamicSecretCreateRedis) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DynamicSecretCreateRedis) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


