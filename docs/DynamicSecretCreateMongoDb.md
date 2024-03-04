# DynamicSecretCreateMongoDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MongodbAtlasApiPrivateKey** | Pointer to **string** | MongoDB Atlas private key | [optional] 
**MongodbAtlasApiPublicKey** | Pointer to **string** | MongoDB Atlas public key | [optional] 
**MongodbAtlasProjectId** | Pointer to **string** | MongoDB Atlas project ID | [optional] 
**MongodbCustomData** | Pointer to **string** | MongoDB custom data | [optional] 
**MongodbDefaultAuthDb** | Pointer to **string** | MongoDB server default authentication database | [optional] 
**MongodbHostPort** | Pointer to **string** | MongoDB server host and port | [optional] 
**MongodbName** | Pointer to **string** | MongoDB Name | [optional] 
**MongodbPassword** | Pointer to **string** | MongoDB server password. You will prompted to provide a password if it will not appear in CLI parameters | [optional] 
**MongodbRoles** | Pointer to **string** | MongoDB Roles | [optional] [default to "[]"]
**MongodbServerUri** | Pointer to **string** | MongoDB server URI | [optional] 
**MongodbUriOptions** | Pointer to **string** | MongoDB server URI options | [optional] 
**MongodbUsername** | Pointer to **string** | MongoDB server username | [optional] 
**Name** | **string** | Dynamic secret name | 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Encrypt producer with following key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts) | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateMongoDb

`func NewDynamicSecretCreateMongoDb(name string, ) *DynamicSecretCreateMongoDb`

NewDynamicSecretCreateMongoDb instantiates a new DynamicSecretCreateMongoDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateMongoDbWithDefaults

`func NewDynamicSecretCreateMongoDbWithDefaults() *DynamicSecretCreateMongoDb`

NewDynamicSecretCreateMongoDbWithDefaults instantiates a new DynamicSecretCreateMongoDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretCreateMongoDb) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateMongoDb) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateMongoDb) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateMongoDb) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateMongoDb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateMongoDb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateMongoDb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateMongoDb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateMongoDb) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateMongoDb) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateMongoDb) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateMongoDb) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMongodbAtlasApiPrivateKey

`func (o *DynamicSecretCreateMongoDb) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *DynamicSecretCreateMongoDb) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *DynamicSecretCreateMongoDb) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *DynamicSecretCreateMongoDb) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *DynamicSecretCreateMongoDb) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *DynamicSecretCreateMongoDb) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *DynamicSecretCreateMongoDb) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *DynamicSecretCreateMongoDb) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *DynamicSecretCreateMongoDb) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbCustomData

`func (o *DynamicSecretCreateMongoDb) GetMongodbCustomData() string`

GetMongodbCustomData returns the MongodbCustomData field if non-nil, zero value otherwise.

### GetMongodbCustomDataOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbCustomDataOk() (*string, bool)`

GetMongodbCustomDataOk returns a tuple with the MongodbCustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbCustomData

`func (o *DynamicSecretCreateMongoDb) SetMongodbCustomData(v string)`

SetMongodbCustomData sets MongodbCustomData field to given value.

### HasMongodbCustomData

`func (o *DynamicSecretCreateMongoDb) HasMongodbCustomData() bool`

HasMongodbCustomData returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *DynamicSecretCreateMongoDb) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *DynamicSecretCreateMongoDb) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *DynamicSecretCreateMongoDb) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbHostPort

`func (o *DynamicSecretCreateMongoDb) GetMongodbHostPort() string`

GetMongodbHostPort returns the MongodbHostPort field if non-nil, zero value otherwise.

### GetMongodbHostPortOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbHostPortOk() (*string, bool)`

GetMongodbHostPortOk returns a tuple with the MongodbHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbHostPort

`func (o *DynamicSecretCreateMongoDb) SetMongodbHostPort(v string)`

SetMongodbHostPort sets MongodbHostPort field to given value.

### HasMongodbHostPort

`func (o *DynamicSecretCreateMongoDb) HasMongodbHostPort() bool`

HasMongodbHostPort returns a boolean if a field has been set.

### GetMongodbName

`func (o *DynamicSecretCreateMongoDb) GetMongodbName() string`

GetMongodbName returns the MongodbName field if non-nil, zero value otherwise.

### GetMongodbNameOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbNameOk() (*string, bool)`

GetMongodbNameOk returns a tuple with the MongodbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbName

`func (o *DynamicSecretCreateMongoDb) SetMongodbName(v string)`

SetMongodbName sets MongodbName field to given value.

### HasMongodbName

`func (o *DynamicSecretCreateMongoDb) HasMongodbName() bool`

HasMongodbName returns a boolean if a field has been set.

### GetMongodbPassword

`func (o *DynamicSecretCreateMongoDb) GetMongodbPassword() string`

GetMongodbPassword returns the MongodbPassword field if non-nil, zero value otherwise.

### GetMongodbPasswordOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbPasswordOk() (*string, bool)`

GetMongodbPasswordOk returns a tuple with the MongodbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbPassword

`func (o *DynamicSecretCreateMongoDb) SetMongodbPassword(v string)`

SetMongodbPassword sets MongodbPassword field to given value.

### HasMongodbPassword

`func (o *DynamicSecretCreateMongoDb) HasMongodbPassword() bool`

HasMongodbPassword returns a boolean if a field has been set.

### GetMongodbRoles

`func (o *DynamicSecretCreateMongoDb) GetMongodbRoles() string`

GetMongodbRoles returns the MongodbRoles field if non-nil, zero value otherwise.

### GetMongodbRolesOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbRolesOk() (*string, bool)`

GetMongodbRolesOk returns a tuple with the MongodbRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbRoles

`func (o *DynamicSecretCreateMongoDb) SetMongodbRoles(v string)`

SetMongodbRoles sets MongodbRoles field to given value.

### HasMongodbRoles

`func (o *DynamicSecretCreateMongoDb) HasMongodbRoles() bool`

HasMongodbRoles returns a boolean if a field has been set.

### GetMongodbServerUri

`func (o *DynamicSecretCreateMongoDb) GetMongodbServerUri() string`

GetMongodbServerUri returns the MongodbServerUri field if non-nil, zero value otherwise.

### GetMongodbServerUriOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbServerUriOk() (*string, bool)`

GetMongodbServerUriOk returns a tuple with the MongodbServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbServerUri

`func (o *DynamicSecretCreateMongoDb) SetMongodbServerUri(v string)`

SetMongodbServerUri sets MongodbServerUri field to given value.

### HasMongodbServerUri

`func (o *DynamicSecretCreateMongoDb) HasMongodbServerUri() bool`

HasMongodbServerUri returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *DynamicSecretCreateMongoDb) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *DynamicSecretCreateMongoDb) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *DynamicSecretCreateMongoDb) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetMongodbUsername

`func (o *DynamicSecretCreateMongoDb) GetMongodbUsername() string`

GetMongodbUsername returns the MongodbUsername field if non-nil, zero value otherwise.

### GetMongodbUsernameOk

`func (o *DynamicSecretCreateMongoDb) GetMongodbUsernameOk() (*string, bool)`

GetMongodbUsernameOk returns a tuple with the MongodbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUsername

`func (o *DynamicSecretCreateMongoDb) SetMongodbUsername(v string)`

SetMongodbUsername sets MongodbUsername field to given value.

### HasMongodbUsername

`func (o *DynamicSecretCreateMongoDb) HasMongodbUsername() bool`

HasMongodbUsername returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateMongoDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateMongoDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateMongoDb) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *DynamicSecretCreateMongoDb) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretCreateMongoDb) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretCreateMongoDb) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretCreateMongoDb) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateMongoDb) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateMongoDb) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateMongoDb) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateMongoDb) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateMongoDb) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretCreateMongoDb) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateMongoDb) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretCreateMongoDb) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateMongoDb) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateMongoDb) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateMongoDb) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateMongoDb) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretCreateMongoDb) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretCreateMongoDb) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretCreateMongoDb) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretCreateMongoDb) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretCreateMongoDb) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretCreateMongoDb) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretCreateMongoDb) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretCreateMongoDb) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateMongoDb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateMongoDb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateMongoDb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateMongoDb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateMongoDb) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateMongoDb) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateMongoDb) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateMongoDb) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateMongoDb) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateMongoDb) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateMongoDb) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateMongoDb) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateMongoDb) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateMongoDb) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateMongoDb) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateMongoDb) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateMongoDb) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateMongoDb) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateMongoDb) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateMongoDb) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


