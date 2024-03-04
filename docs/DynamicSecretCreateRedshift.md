# DynamicSecretCreateRedshift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationStatements** | Pointer to **string** | Redshift Creation statements | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
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

### NewDynamicSecretCreateRedshift

`func NewDynamicSecretCreateRedshift(name string, ) *DynamicSecretCreateRedshift`

NewDynamicSecretCreateRedshift instantiates a new DynamicSecretCreateRedshift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateRedshiftWithDefaults

`func NewDynamicSecretCreateRedshiftWithDefaults() *DynamicSecretCreateRedshift`

NewDynamicSecretCreateRedshiftWithDefaults instantiates a new DynamicSecretCreateRedshift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationStatements

`func (o *DynamicSecretCreateRedshift) GetCreationStatements() string`

GetCreationStatements returns the CreationStatements field if non-nil, zero value otherwise.

### GetCreationStatementsOk

`func (o *DynamicSecretCreateRedshift) GetCreationStatementsOk() (*string, bool)`

GetCreationStatementsOk returns a tuple with the CreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStatements

`func (o *DynamicSecretCreateRedshift) SetCreationStatements(v string)`

SetCreationStatements sets CreationStatements field to given value.

### HasCreationStatements

`func (o *DynamicSecretCreateRedshift) HasCreationStatements() bool`

HasCreationStatements returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateRedshift) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateRedshift) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateRedshift) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateRedshift) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateRedshift) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateRedshift) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateRedshift) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateRedshift) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateRedshift) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateRedshift) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateRedshift) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateRedshift) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateRedshift) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateRedshift) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateRedshift) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *DynamicSecretCreateRedshift) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretCreateRedshift) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretCreateRedshift) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretCreateRedshift) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKey

`func (o *DynamicSecretCreateRedshift) GetProducerEncryptionKey() string`

GetProducerEncryptionKey returns the ProducerEncryptionKey field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyOk

`func (o *DynamicSecretCreateRedshift) GetProducerEncryptionKeyOk() (*string, bool)`

GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKey

`func (o *DynamicSecretCreateRedshift) SetProducerEncryptionKey(v string)`

SetProducerEncryptionKey sets ProducerEncryptionKey field to given value.

### HasProducerEncryptionKey

`func (o *DynamicSecretCreateRedshift) HasProducerEncryptionKey() bool`

HasProducerEncryptionKey returns a boolean if a field has been set.

### GetRedshiftDbName

`func (o *DynamicSecretCreateRedshift) GetRedshiftDbName() string`

GetRedshiftDbName returns the RedshiftDbName field if non-nil, zero value otherwise.

### GetRedshiftDbNameOk

`func (o *DynamicSecretCreateRedshift) GetRedshiftDbNameOk() (*string, bool)`

GetRedshiftDbNameOk returns a tuple with the RedshiftDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftDbName

`func (o *DynamicSecretCreateRedshift) SetRedshiftDbName(v string)`

SetRedshiftDbName sets RedshiftDbName field to given value.

### HasRedshiftDbName

`func (o *DynamicSecretCreateRedshift) HasRedshiftDbName() bool`

HasRedshiftDbName returns a boolean if a field has been set.

### GetRedshiftHost

`func (o *DynamicSecretCreateRedshift) GetRedshiftHost() string`

GetRedshiftHost returns the RedshiftHost field if non-nil, zero value otherwise.

### GetRedshiftHostOk

`func (o *DynamicSecretCreateRedshift) GetRedshiftHostOk() (*string, bool)`

GetRedshiftHostOk returns a tuple with the RedshiftHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftHost

`func (o *DynamicSecretCreateRedshift) SetRedshiftHost(v string)`

SetRedshiftHost sets RedshiftHost field to given value.

### HasRedshiftHost

`func (o *DynamicSecretCreateRedshift) HasRedshiftHost() bool`

HasRedshiftHost returns a boolean if a field has been set.

### GetRedshiftPassword

`func (o *DynamicSecretCreateRedshift) GetRedshiftPassword() string`

GetRedshiftPassword returns the RedshiftPassword field if non-nil, zero value otherwise.

### GetRedshiftPasswordOk

`func (o *DynamicSecretCreateRedshift) GetRedshiftPasswordOk() (*string, bool)`

GetRedshiftPasswordOk returns a tuple with the RedshiftPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftPassword

`func (o *DynamicSecretCreateRedshift) SetRedshiftPassword(v string)`

SetRedshiftPassword sets RedshiftPassword field to given value.

### HasRedshiftPassword

`func (o *DynamicSecretCreateRedshift) HasRedshiftPassword() bool`

HasRedshiftPassword returns a boolean if a field has been set.

### GetRedshiftPort

`func (o *DynamicSecretCreateRedshift) GetRedshiftPort() string`

GetRedshiftPort returns the RedshiftPort field if non-nil, zero value otherwise.

### GetRedshiftPortOk

`func (o *DynamicSecretCreateRedshift) GetRedshiftPortOk() (*string, bool)`

GetRedshiftPortOk returns a tuple with the RedshiftPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftPort

`func (o *DynamicSecretCreateRedshift) SetRedshiftPort(v string)`

SetRedshiftPort sets RedshiftPort field to given value.

### HasRedshiftPort

`func (o *DynamicSecretCreateRedshift) HasRedshiftPort() bool`

HasRedshiftPort returns a boolean if a field has been set.

### GetRedshiftUsername

`func (o *DynamicSecretCreateRedshift) GetRedshiftUsername() string`

GetRedshiftUsername returns the RedshiftUsername field if non-nil, zero value otherwise.

### GetRedshiftUsernameOk

`func (o *DynamicSecretCreateRedshift) GetRedshiftUsernameOk() (*string, bool)`

GetRedshiftUsernameOk returns a tuple with the RedshiftUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftUsername

`func (o *DynamicSecretCreateRedshift) SetRedshiftUsername(v string)`

SetRedshiftUsername sets RedshiftUsername field to given value.

### HasRedshiftUsername

`func (o *DynamicSecretCreateRedshift) HasRedshiftUsername() bool`

HasRedshiftUsername returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateRedshift) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateRedshift) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateRedshift) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateRedshift) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretCreateRedshift) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretCreateRedshift) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretCreateRedshift) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretCreateRedshift) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSsl

`func (o *DynamicSecretCreateRedshift) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *DynamicSecretCreateRedshift) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *DynamicSecretCreateRedshift) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *DynamicSecretCreateRedshift) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateRedshift) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateRedshift) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateRedshift) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateRedshift) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateRedshift) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateRedshift) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateRedshift) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateRedshift) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateRedshift) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateRedshift) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateRedshift) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateRedshift) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateRedshift) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateRedshift) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateRedshift) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateRedshift) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateRedshift) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateRedshift) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateRedshift) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateRedshift) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


