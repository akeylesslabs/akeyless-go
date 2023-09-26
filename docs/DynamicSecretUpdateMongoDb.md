# DynamicSecretUpdateMongoDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
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
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
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

### NewDynamicSecretUpdateMongoDb

`func NewDynamicSecretUpdateMongoDb(name string, ) *DynamicSecretUpdateMongoDb`

NewDynamicSecretUpdateMongoDb instantiates a new DynamicSecretUpdateMongoDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateMongoDbWithDefaults

`func NewDynamicSecretUpdateMongoDbWithDefaults() *DynamicSecretUpdateMongoDb`

NewDynamicSecretUpdateMongoDbWithDefaults instantiates a new DynamicSecretUpdateMongoDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretUpdateMongoDb) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateMongoDb) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateMongoDb) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateMongoDb) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateMongoDb) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateMongoDb) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateMongoDb) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateMongoDb) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMongodbAtlasApiPrivateKey

`func (o *DynamicSecretUpdateMongoDb) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *DynamicSecretUpdateMongoDb) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *DynamicSecretUpdateMongoDb) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *DynamicSecretUpdateMongoDb) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *DynamicSecretUpdateMongoDb) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *DynamicSecretUpdateMongoDb) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *DynamicSecretUpdateMongoDb) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *DynamicSecretUpdateMongoDb) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *DynamicSecretUpdateMongoDb) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbCustomData

`func (o *DynamicSecretUpdateMongoDb) GetMongodbCustomData() string`

GetMongodbCustomData returns the MongodbCustomData field if non-nil, zero value otherwise.

### GetMongodbCustomDataOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbCustomDataOk() (*string, bool)`

GetMongodbCustomDataOk returns a tuple with the MongodbCustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbCustomData

`func (o *DynamicSecretUpdateMongoDb) SetMongodbCustomData(v string)`

SetMongodbCustomData sets MongodbCustomData field to given value.

### HasMongodbCustomData

`func (o *DynamicSecretUpdateMongoDb) HasMongodbCustomData() bool`

HasMongodbCustomData returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *DynamicSecretUpdateMongoDb) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *DynamicSecretUpdateMongoDb) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *DynamicSecretUpdateMongoDb) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbHostPort

`func (o *DynamicSecretUpdateMongoDb) GetMongodbHostPort() string`

GetMongodbHostPort returns the MongodbHostPort field if non-nil, zero value otherwise.

### GetMongodbHostPortOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbHostPortOk() (*string, bool)`

GetMongodbHostPortOk returns a tuple with the MongodbHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbHostPort

`func (o *DynamicSecretUpdateMongoDb) SetMongodbHostPort(v string)`

SetMongodbHostPort sets MongodbHostPort field to given value.

### HasMongodbHostPort

`func (o *DynamicSecretUpdateMongoDb) HasMongodbHostPort() bool`

HasMongodbHostPort returns a boolean if a field has been set.

### GetMongodbName

`func (o *DynamicSecretUpdateMongoDb) GetMongodbName() string`

GetMongodbName returns the MongodbName field if non-nil, zero value otherwise.

### GetMongodbNameOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbNameOk() (*string, bool)`

GetMongodbNameOk returns a tuple with the MongodbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbName

`func (o *DynamicSecretUpdateMongoDb) SetMongodbName(v string)`

SetMongodbName sets MongodbName field to given value.

### HasMongodbName

`func (o *DynamicSecretUpdateMongoDb) HasMongodbName() bool`

HasMongodbName returns a boolean if a field has been set.

### GetMongodbPassword

`func (o *DynamicSecretUpdateMongoDb) GetMongodbPassword() string`

GetMongodbPassword returns the MongodbPassword field if non-nil, zero value otherwise.

### GetMongodbPasswordOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbPasswordOk() (*string, bool)`

GetMongodbPasswordOk returns a tuple with the MongodbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbPassword

`func (o *DynamicSecretUpdateMongoDb) SetMongodbPassword(v string)`

SetMongodbPassword sets MongodbPassword field to given value.

### HasMongodbPassword

`func (o *DynamicSecretUpdateMongoDb) HasMongodbPassword() bool`

HasMongodbPassword returns a boolean if a field has been set.

### GetMongodbRoles

`func (o *DynamicSecretUpdateMongoDb) GetMongodbRoles() string`

GetMongodbRoles returns the MongodbRoles field if non-nil, zero value otherwise.

### GetMongodbRolesOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbRolesOk() (*string, bool)`

GetMongodbRolesOk returns a tuple with the MongodbRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbRoles

`func (o *DynamicSecretUpdateMongoDb) SetMongodbRoles(v string)`

SetMongodbRoles sets MongodbRoles field to given value.

### HasMongodbRoles

`func (o *DynamicSecretUpdateMongoDb) HasMongodbRoles() bool`

HasMongodbRoles returns a boolean if a field has been set.

### GetMongodbServerUri

`func (o *DynamicSecretUpdateMongoDb) GetMongodbServerUri() string`

GetMongodbServerUri returns the MongodbServerUri field if non-nil, zero value otherwise.

### GetMongodbServerUriOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbServerUriOk() (*string, bool)`

GetMongodbServerUriOk returns a tuple with the MongodbServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbServerUri

`func (o *DynamicSecretUpdateMongoDb) SetMongodbServerUri(v string)`

SetMongodbServerUri sets MongodbServerUri field to given value.

### HasMongodbServerUri

`func (o *DynamicSecretUpdateMongoDb) HasMongodbServerUri() bool`

HasMongodbServerUri returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *DynamicSecretUpdateMongoDb) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *DynamicSecretUpdateMongoDb) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *DynamicSecretUpdateMongoDb) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetMongodbUsername

`func (o *DynamicSecretUpdateMongoDb) GetMongodbUsername() string`

GetMongodbUsername returns the MongodbUsername field if non-nil, zero value otherwise.

### GetMongodbUsernameOk

`func (o *DynamicSecretUpdateMongoDb) GetMongodbUsernameOk() (*string, bool)`

GetMongodbUsernameOk returns a tuple with the MongodbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUsername

`func (o *DynamicSecretUpdateMongoDb) SetMongodbUsername(v string)`

SetMongodbUsername sets MongodbUsername field to given value.

### HasMongodbUsername

`func (o *DynamicSecretUpdateMongoDb) HasMongodbUsername() bool`

HasMongodbUsername returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateMongoDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateMongoDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateMongoDb) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateMongoDb) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateMongoDb) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateMongoDb) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateMongoDb) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateMongoDb) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateMongoDb) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateMongoDb) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateMongoDb) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateMongoDb) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretUpdateMongoDb) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateMongoDb) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateMongoDb) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateMongoDb) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateMongoDb) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateMongoDb) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateMongoDb) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretUpdateMongoDb) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretUpdateMongoDb) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretUpdateMongoDb) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretUpdateMongoDb) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretUpdateMongoDb) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretUpdateMongoDb) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretUpdateMongoDb) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretUpdateMongoDb) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateMongoDb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateMongoDb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateMongoDb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateMongoDb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateMongoDb) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateMongoDb) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateMongoDb) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateMongoDb) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateMongoDb) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateMongoDb) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateMongoDb) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateMongoDb) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateMongoDb) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateMongoDb) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateMongoDb) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateMongoDb) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateMongoDb) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateMongoDb) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateMongoDb) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateMongoDb) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


