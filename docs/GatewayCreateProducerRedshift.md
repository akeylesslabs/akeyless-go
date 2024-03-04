# GatewayCreateProducerRedshift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationStatements** | Pointer to **string** | Redshift Creation statements | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKey** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RedshiftDbName** | Pointer to **string** | Redshift DB Name | [optional] 
**RedshiftHost** | Pointer to **string** | Redshift Host | [optional] [default to "127.0.0.1"]
**RedshiftPassword** | Pointer to **string** | Redshift Password | [optional] 
**RedshiftPort** | Pointer to **string** | Redshift Port | [optional] [default to "5439"]
**RedshiftUsername** | Pointer to **string** | Redshift Username | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts) | [optional] 
**Ssl** | Pointer to **bool** | Enable/Disable SSL [true/false] | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerRedshift

`func NewGatewayCreateProducerRedshift(name string, ) *GatewayCreateProducerRedshift`

NewGatewayCreateProducerRedshift instantiates a new GatewayCreateProducerRedshift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerRedshiftWithDefaults

`func NewGatewayCreateProducerRedshiftWithDefaults() *GatewayCreateProducerRedshift`

NewGatewayCreateProducerRedshiftWithDefaults instantiates a new GatewayCreateProducerRedshift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationStatements

`func (o *GatewayCreateProducerRedshift) GetCreationStatements() string`

GetCreationStatements returns the CreationStatements field if non-nil, zero value otherwise.

### GetCreationStatementsOk

`func (o *GatewayCreateProducerRedshift) GetCreationStatementsOk() (*string, bool)`

GetCreationStatementsOk returns a tuple with the CreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStatements

`func (o *GatewayCreateProducerRedshift) SetCreationStatements(v string)`

SetCreationStatements sets CreationStatements field to given value.

### HasCreationStatements

`func (o *GatewayCreateProducerRedshift) HasCreationStatements() bool`

HasCreationStatements returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayCreateProducerRedshift) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayCreateProducerRedshift) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayCreateProducerRedshift) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayCreateProducerRedshift) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateProducerRedshift) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateProducerRedshift) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateProducerRedshift) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateProducerRedshift) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerRedshift) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerRedshift) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerRedshift) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *GatewayCreateProducerRedshift) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *GatewayCreateProducerRedshift) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *GatewayCreateProducerRedshift) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *GatewayCreateProducerRedshift) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKey

`func (o *GatewayCreateProducerRedshift) GetProducerEncryptionKey() string`

GetProducerEncryptionKey returns the ProducerEncryptionKey field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyOk

`func (o *GatewayCreateProducerRedshift) GetProducerEncryptionKeyOk() (*string, bool)`

GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKey

`func (o *GatewayCreateProducerRedshift) SetProducerEncryptionKey(v string)`

SetProducerEncryptionKey sets ProducerEncryptionKey field to given value.

### HasProducerEncryptionKey

`func (o *GatewayCreateProducerRedshift) HasProducerEncryptionKey() bool`

HasProducerEncryptionKey returns a boolean if a field has been set.

### GetRedshiftDbName

`func (o *GatewayCreateProducerRedshift) GetRedshiftDbName() string`

GetRedshiftDbName returns the RedshiftDbName field if non-nil, zero value otherwise.

### GetRedshiftDbNameOk

`func (o *GatewayCreateProducerRedshift) GetRedshiftDbNameOk() (*string, bool)`

GetRedshiftDbNameOk returns a tuple with the RedshiftDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftDbName

`func (o *GatewayCreateProducerRedshift) SetRedshiftDbName(v string)`

SetRedshiftDbName sets RedshiftDbName field to given value.

### HasRedshiftDbName

`func (o *GatewayCreateProducerRedshift) HasRedshiftDbName() bool`

HasRedshiftDbName returns a boolean if a field has been set.

### GetRedshiftHost

`func (o *GatewayCreateProducerRedshift) GetRedshiftHost() string`

GetRedshiftHost returns the RedshiftHost field if non-nil, zero value otherwise.

### GetRedshiftHostOk

`func (o *GatewayCreateProducerRedshift) GetRedshiftHostOk() (*string, bool)`

GetRedshiftHostOk returns a tuple with the RedshiftHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftHost

`func (o *GatewayCreateProducerRedshift) SetRedshiftHost(v string)`

SetRedshiftHost sets RedshiftHost field to given value.

### HasRedshiftHost

`func (o *GatewayCreateProducerRedshift) HasRedshiftHost() bool`

HasRedshiftHost returns a boolean if a field has been set.

### GetRedshiftPassword

`func (o *GatewayCreateProducerRedshift) GetRedshiftPassword() string`

GetRedshiftPassword returns the RedshiftPassword field if non-nil, zero value otherwise.

### GetRedshiftPasswordOk

`func (o *GatewayCreateProducerRedshift) GetRedshiftPasswordOk() (*string, bool)`

GetRedshiftPasswordOk returns a tuple with the RedshiftPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftPassword

`func (o *GatewayCreateProducerRedshift) SetRedshiftPassword(v string)`

SetRedshiftPassword sets RedshiftPassword field to given value.

### HasRedshiftPassword

`func (o *GatewayCreateProducerRedshift) HasRedshiftPassword() bool`

HasRedshiftPassword returns a boolean if a field has been set.

### GetRedshiftPort

`func (o *GatewayCreateProducerRedshift) GetRedshiftPort() string`

GetRedshiftPort returns the RedshiftPort field if non-nil, zero value otherwise.

### GetRedshiftPortOk

`func (o *GatewayCreateProducerRedshift) GetRedshiftPortOk() (*string, bool)`

GetRedshiftPortOk returns a tuple with the RedshiftPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftPort

`func (o *GatewayCreateProducerRedshift) SetRedshiftPort(v string)`

SetRedshiftPort sets RedshiftPort field to given value.

### HasRedshiftPort

`func (o *GatewayCreateProducerRedshift) HasRedshiftPort() bool`

HasRedshiftPort returns a boolean if a field has been set.

### GetRedshiftUsername

`func (o *GatewayCreateProducerRedshift) GetRedshiftUsername() string`

GetRedshiftUsername returns the RedshiftUsername field if non-nil, zero value otherwise.

### GetRedshiftUsernameOk

`func (o *GatewayCreateProducerRedshift) GetRedshiftUsernameOk() (*string, bool)`

GetRedshiftUsernameOk returns a tuple with the RedshiftUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftUsername

`func (o *GatewayCreateProducerRedshift) SetRedshiftUsername(v string)`

SetRedshiftUsername sets RedshiftUsername field to given value.

### HasRedshiftUsername

`func (o *GatewayCreateProducerRedshift) HasRedshiftUsername() bool`

HasRedshiftUsername returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayCreateProducerRedshift) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayCreateProducerRedshift) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayCreateProducerRedshift) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayCreateProducerRedshift) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayCreateProducerRedshift) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayCreateProducerRedshift) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayCreateProducerRedshift) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayCreateProducerRedshift) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSsl

`func (o *GatewayCreateProducerRedshift) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *GatewayCreateProducerRedshift) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *GatewayCreateProducerRedshift) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *GatewayCreateProducerRedshift) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerRedshift) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerRedshift) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerRedshift) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerRedshift) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerRedshift) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerRedshift) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerRedshift) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerRedshift) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerRedshift) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerRedshift) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerRedshift) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerRedshift) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerRedshift) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerRedshift) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerRedshift) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerRedshift) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerRedshift) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerRedshift) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerRedshift) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerRedshift) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


