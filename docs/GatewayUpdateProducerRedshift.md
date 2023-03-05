# GatewayUpdateProducerRedshift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationStatements** | Pointer to **string** | Redshift Creation statements | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**ProducerEncryptionKey** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RedshiftDbName** | Pointer to **string** | Redshift DB Name | [optional] 
**RedshiftHost** | Pointer to **string** | Redshift Host | [optional] [default to "127.0.0.1"]
**RedshiftPassword** | Pointer to **string** | Redshift Password | [optional] 
**RedshiftPort** | Pointer to **string** | Redshift Port | [optional] [default to "5439"]
**RedshiftUsername** | Pointer to **string** | Redshift Username | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections | [optional] 
**Ssl** | Pointer to **bool** | Enable/Disable SSL [true/false] | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayUpdateProducerRedshift

`func NewGatewayUpdateProducerRedshift(name string, ) *GatewayUpdateProducerRedshift`

NewGatewayUpdateProducerRedshift instantiates a new GatewayUpdateProducerRedshift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerRedshiftWithDefaults

`func NewGatewayUpdateProducerRedshiftWithDefaults() *GatewayUpdateProducerRedshift`

NewGatewayUpdateProducerRedshiftWithDefaults instantiates a new GatewayUpdateProducerRedshift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationStatements

`func (o *GatewayUpdateProducerRedshift) GetCreationStatements() string`

GetCreationStatements returns the CreationStatements field if non-nil, zero value otherwise.

### GetCreationStatementsOk

`func (o *GatewayUpdateProducerRedshift) GetCreationStatementsOk() (*string, bool)`

GetCreationStatementsOk returns a tuple with the CreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStatements

`func (o *GatewayUpdateProducerRedshift) SetCreationStatements(v string)`

SetCreationStatements sets CreationStatements field to given value.

### HasCreationStatements

`func (o *GatewayUpdateProducerRedshift) HasCreationStatements() bool`

HasCreationStatements returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayUpdateProducerRedshift) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerRedshift) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerRedshift) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerRedshift) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateProducerRedshift) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateProducerRedshift) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateProducerRedshift) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateProducerRedshift) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerRedshift) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerRedshift) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerRedshift) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerRedshift) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerRedshift) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerRedshift) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerRedshift) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKey

`func (o *GatewayUpdateProducerRedshift) GetProducerEncryptionKey() string`

GetProducerEncryptionKey returns the ProducerEncryptionKey field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyOk

`func (o *GatewayUpdateProducerRedshift) GetProducerEncryptionKeyOk() (*string, bool)`

GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKey

`func (o *GatewayUpdateProducerRedshift) SetProducerEncryptionKey(v string)`

SetProducerEncryptionKey sets ProducerEncryptionKey field to given value.

### HasProducerEncryptionKey

`func (o *GatewayUpdateProducerRedshift) HasProducerEncryptionKey() bool`

HasProducerEncryptionKey returns a boolean if a field has been set.

### GetRedshiftDbName

`func (o *GatewayUpdateProducerRedshift) GetRedshiftDbName() string`

GetRedshiftDbName returns the RedshiftDbName field if non-nil, zero value otherwise.

### GetRedshiftDbNameOk

`func (o *GatewayUpdateProducerRedshift) GetRedshiftDbNameOk() (*string, bool)`

GetRedshiftDbNameOk returns a tuple with the RedshiftDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftDbName

`func (o *GatewayUpdateProducerRedshift) SetRedshiftDbName(v string)`

SetRedshiftDbName sets RedshiftDbName field to given value.

### HasRedshiftDbName

`func (o *GatewayUpdateProducerRedshift) HasRedshiftDbName() bool`

HasRedshiftDbName returns a boolean if a field has been set.

### GetRedshiftHost

`func (o *GatewayUpdateProducerRedshift) GetRedshiftHost() string`

GetRedshiftHost returns the RedshiftHost field if non-nil, zero value otherwise.

### GetRedshiftHostOk

`func (o *GatewayUpdateProducerRedshift) GetRedshiftHostOk() (*string, bool)`

GetRedshiftHostOk returns a tuple with the RedshiftHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftHost

`func (o *GatewayUpdateProducerRedshift) SetRedshiftHost(v string)`

SetRedshiftHost sets RedshiftHost field to given value.

### HasRedshiftHost

`func (o *GatewayUpdateProducerRedshift) HasRedshiftHost() bool`

HasRedshiftHost returns a boolean if a field has been set.

### GetRedshiftPassword

`func (o *GatewayUpdateProducerRedshift) GetRedshiftPassword() string`

GetRedshiftPassword returns the RedshiftPassword field if non-nil, zero value otherwise.

### GetRedshiftPasswordOk

`func (o *GatewayUpdateProducerRedshift) GetRedshiftPasswordOk() (*string, bool)`

GetRedshiftPasswordOk returns a tuple with the RedshiftPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftPassword

`func (o *GatewayUpdateProducerRedshift) SetRedshiftPassword(v string)`

SetRedshiftPassword sets RedshiftPassword field to given value.

### HasRedshiftPassword

`func (o *GatewayUpdateProducerRedshift) HasRedshiftPassword() bool`

HasRedshiftPassword returns a boolean if a field has been set.

### GetRedshiftPort

`func (o *GatewayUpdateProducerRedshift) GetRedshiftPort() string`

GetRedshiftPort returns the RedshiftPort field if non-nil, zero value otherwise.

### GetRedshiftPortOk

`func (o *GatewayUpdateProducerRedshift) GetRedshiftPortOk() (*string, bool)`

GetRedshiftPortOk returns a tuple with the RedshiftPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftPort

`func (o *GatewayUpdateProducerRedshift) SetRedshiftPort(v string)`

SetRedshiftPort sets RedshiftPort field to given value.

### HasRedshiftPort

`func (o *GatewayUpdateProducerRedshift) HasRedshiftPort() bool`

HasRedshiftPort returns a boolean if a field has been set.

### GetRedshiftUsername

`func (o *GatewayUpdateProducerRedshift) GetRedshiftUsername() string`

GetRedshiftUsername returns the RedshiftUsername field if non-nil, zero value otherwise.

### GetRedshiftUsernameOk

`func (o *GatewayUpdateProducerRedshift) GetRedshiftUsernameOk() (*string, bool)`

GetRedshiftUsernameOk returns a tuple with the RedshiftUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftUsername

`func (o *GatewayUpdateProducerRedshift) SetRedshiftUsername(v string)`

SetRedshiftUsername sets RedshiftUsername field to given value.

### HasRedshiftUsername

`func (o *GatewayUpdateProducerRedshift) HasRedshiftUsername() bool`

HasRedshiftUsername returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayUpdateProducerRedshift) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayUpdateProducerRedshift) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayUpdateProducerRedshift) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayUpdateProducerRedshift) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayUpdateProducerRedshift) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayUpdateProducerRedshift) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayUpdateProducerRedshift) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayUpdateProducerRedshift) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSsl

`func (o *GatewayUpdateProducerRedshift) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *GatewayUpdateProducerRedshift) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *GatewayUpdateProducerRedshift) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *GatewayUpdateProducerRedshift) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerRedshift) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerRedshift) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerRedshift) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerRedshift) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerRedshift) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerRedshift) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerRedshift) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerRedshift) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerRedshift) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerRedshift) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerRedshift) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerRedshift) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerRedshift) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerRedshift) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerRedshift) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerRedshift) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerRedshift) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerRedshift) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerRedshift) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerRedshift) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


