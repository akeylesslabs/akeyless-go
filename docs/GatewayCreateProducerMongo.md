# GatewayCreateProducerMongo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MongodbAtlasApiPrivateKey** | Pointer to **string** | MongoDB Atlas private key | [optional] 
**MongodbAtlasApiPublicKey** | Pointer to **string** | MongoDB Atlas public key | [optional] 
**MongodbAtlasProjectId** | Pointer to **string** | MongoDB Atlas project ID | [optional] 
**MongodbDefaultAuthDb** | Pointer to **string** | MongoDB server default authentication database | [optional] 
**MongodbHostPort** | Pointer to **string** | MongoDB server host and port | [optional] 
**MongodbName** | **string** | MongoDB Name | 
**MongodbPassword** | Pointer to **string** | MongoDB server password. You will prompted to provide a password if it will not appear in CLI parameters | [optional] 
**MongodbRoles** | Pointer to **string** | MongoDB Roles | [optional] [default to "[]"]
**MongodbServerUri** | Pointer to **string** | MongoDB server URI | [optional] 
**MongodbUriOptions** | Pointer to **string** | MongoDB server URI options | [optional] 
**MongodbUsername** | Pointer to **string** | MongoDB server username | [optional] 
**Name** | **string** | Producer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Encrypt producer with following key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** |  | [optional] 
**SecureAccessEnable** | Pointer to **string** |  | [optional] 
**SecureAccessHost** | Pointer to **[]string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayCreateProducerMongo

`func NewGatewayCreateProducerMongo(mongodbName string, name string, ) *GatewayCreateProducerMongo`

NewGatewayCreateProducerMongo instantiates a new GatewayCreateProducerMongo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerMongoWithDefaults

`func NewGatewayCreateProducerMongoWithDefaults() *GatewayCreateProducerMongo`

NewGatewayCreateProducerMongoWithDefaults instantiates a new GatewayCreateProducerMongo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMongodbAtlasApiPrivateKey

`func (o *GatewayCreateProducerMongo) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *GatewayCreateProducerMongo) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *GatewayCreateProducerMongo) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *GatewayCreateProducerMongo) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *GatewayCreateProducerMongo) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *GatewayCreateProducerMongo) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *GatewayCreateProducerMongo) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *GatewayCreateProducerMongo) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *GatewayCreateProducerMongo) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *GatewayCreateProducerMongo) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *GatewayCreateProducerMongo) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *GatewayCreateProducerMongo) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *GatewayCreateProducerMongo) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *GatewayCreateProducerMongo) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *GatewayCreateProducerMongo) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *GatewayCreateProducerMongo) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbHostPort

`func (o *GatewayCreateProducerMongo) GetMongodbHostPort() string`

GetMongodbHostPort returns the MongodbHostPort field if non-nil, zero value otherwise.

### GetMongodbHostPortOk

`func (o *GatewayCreateProducerMongo) GetMongodbHostPortOk() (*string, bool)`

GetMongodbHostPortOk returns a tuple with the MongodbHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbHostPort

`func (o *GatewayCreateProducerMongo) SetMongodbHostPort(v string)`

SetMongodbHostPort sets MongodbHostPort field to given value.

### HasMongodbHostPort

`func (o *GatewayCreateProducerMongo) HasMongodbHostPort() bool`

HasMongodbHostPort returns a boolean if a field has been set.

### GetMongodbName

`func (o *GatewayCreateProducerMongo) GetMongodbName() string`

GetMongodbName returns the MongodbName field if non-nil, zero value otherwise.

### GetMongodbNameOk

`func (o *GatewayCreateProducerMongo) GetMongodbNameOk() (*string, bool)`

GetMongodbNameOk returns a tuple with the MongodbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbName

`func (o *GatewayCreateProducerMongo) SetMongodbName(v string)`

SetMongodbName sets MongodbName field to given value.


### GetMongodbPassword

`func (o *GatewayCreateProducerMongo) GetMongodbPassword() string`

GetMongodbPassword returns the MongodbPassword field if non-nil, zero value otherwise.

### GetMongodbPasswordOk

`func (o *GatewayCreateProducerMongo) GetMongodbPasswordOk() (*string, bool)`

GetMongodbPasswordOk returns a tuple with the MongodbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbPassword

`func (o *GatewayCreateProducerMongo) SetMongodbPassword(v string)`

SetMongodbPassword sets MongodbPassword field to given value.

### HasMongodbPassword

`func (o *GatewayCreateProducerMongo) HasMongodbPassword() bool`

HasMongodbPassword returns a boolean if a field has been set.

### GetMongodbRoles

`func (o *GatewayCreateProducerMongo) GetMongodbRoles() string`

GetMongodbRoles returns the MongodbRoles field if non-nil, zero value otherwise.

### GetMongodbRolesOk

`func (o *GatewayCreateProducerMongo) GetMongodbRolesOk() (*string, bool)`

GetMongodbRolesOk returns a tuple with the MongodbRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbRoles

`func (o *GatewayCreateProducerMongo) SetMongodbRoles(v string)`

SetMongodbRoles sets MongodbRoles field to given value.

### HasMongodbRoles

`func (o *GatewayCreateProducerMongo) HasMongodbRoles() bool`

HasMongodbRoles returns a boolean if a field has been set.

### GetMongodbServerUri

`func (o *GatewayCreateProducerMongo) GetMongodbServerUri() string`

GetMongodbServerUri returns the MongodbServerUri field if non-nil, zero value otherwise.

### GetMongodbServerUriOk

`func (o *GatewayCreateProducerMongo) GetMongodbServerUriOk() (*string, bool)`

GetMongodbServerUriOk returns a tuple with the MongodbServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbServerUri

`func (o *GatewayCreateProducerMongo) SetMongodbServerUri(v string)`

SetMongodbServerUri sets MongodbServerUri field to given value.

### HasMongodbServerUri

`func (o *GatewayCreateProducerMongo) HasMongodbServerUri() bool`

HasMongodbServerUri returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *GatewayCreateProducerMongo) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *GatewayCreateProducerMongo) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *GatewayCreateProducerMongo) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *GatewayCreateProducerMongo) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetMongodbUsername

`func (o *GatewayCreateProducerMongo) GetMongodbUsername() string`

GetMongodbUsername returns the MongodbUsername field if non-nil, zero value otherwise.

### GetMongodbUsernameOk

`func (o *GatewayCreateProducerMongo) GetMongodbUsernameOk() (*string, bool)`

GetMongodbUsernameOk returns a tuple with the MongodbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUsername

`func (o *GatewayCreateProducerMongo) SetMongodbUsername(v string)`

SetMongodbUsername sets MongodbUsername field to given value.

### HasMongodbUsername

`func (o *GatewayCreateProducerMongo) HasMongodbUsername() bool`

HasMongodbUsername returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerMongo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerMongo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerMongo) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *GatewayCreateProducerMongo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateProducerMongo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateProducerMongo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateProducerMongo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerMongo) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerMongo) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerMongo) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerMongo) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerMongo) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayCreateProducerMongo) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerMongo) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayCreateProducerMongo) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayCreateProducerMongo) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayCreateProducerMongo) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayCreateProducerMongo) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayCreateProducerMongo) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayCreateProducerMongo) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayCreateProducerMongo) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayCreateProducerMongo) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayCreateProducerMongo) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerMongo) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerMongo) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerMongo) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerMongo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerMongo) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerMongo) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerMongo) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerMongo) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerMongo) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerMongo) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerMongo) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerMongo) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayCreateProducerMongo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateProducerMongo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateProducerMongo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateProducerMongo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


