# DynamicSecretUpdateRedshift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationStatements** | Pointer to **string** | Redshift Creation statements | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
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

### NewDynamicSecretUpdateRedshift

`func NewDynamicSecretUpdateRedshift(name string, ) *DynamicSecretUpdateRedshift`

NewDynamicSecretUpdateRedshift instantiates a new DynamicSecretUpdateRedshift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateRedshiftWithDefaults

`func NewDynamicSecretUpdateRedshiftWithDefaults() *DynamicSecretUpdateRedshift`

NewDynamicSecretUpdateRedshiftWithDefaults instantiates a new DynamicSecretUpdateRedshift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationStatements

`func (o *DynamicSecretUpdateRedshift) GetCreationStatements() string`

GetCreationStatements returns the CreationStatements field if non-nil, zero value otherwise.

### GetCreationStatementsOk

`func (o *DynamicSecretUpdateRedshift) GetCreationStatementsOk() (*string, bool)`

GetCreationStatementsOk returns a tuple with the CreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStatements

`func (o *DynamicSecretUpdateRedshift) SetCreationStatements(v string)`

SetCreationStatements sets CreationStatements field to given value.

### HasCreationStatements

`func (o *DynamicSecretUpdateRedshift) HasCreationStatements() bool`

HasCreationStatements returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretUpdateRedshift) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateRedshift) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateRedshift) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateRedshift) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateRedshift) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateRedshift) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateRedshift) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateRedshift) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateRedshift) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateRedshift) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateRedshift) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateRedshift) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateRedshift) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateRedshift) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateRedshift) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateRedshift) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateRedshift) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateRedshift) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateRedshift) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretUpdateRedshift) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretUpdateRedshift) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretUpdateRedshift) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretUpdateRedshift) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKey

`func (o *DynamicSecretUpdateRedshift) GetProducerEncryptionKey() string`

GetProducerEncryptionKey returns the ProducerEncryptionKey field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyOk

`func (o *DynamicSecretUpdateRedshift) GetProducerEncryptionKeyOk() (*string, bool)`

GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKey

`func (o *DynamicSecretUpdateRedshift) SetProducerEncryptionKey(v string)`

SetProducerEncryptionKey sets ProducerEncryptionKey field to given value.

### HasProducerEncryptionKey

`func (o *DynamicSecretUpdateRedshift) HasProducerEncryptionKey() bool`

HasProducerEncryptionKey returns a boolean if a field has been set.

### GetRedshiftDbName

`func (o *DynamicSecretUpdateRedshift) GetRedshiftDbName() string`

GetRedshiftDbName returns the RedshiftDbName field if non-nil, zero value otherwise.

### GetRedshiftDbNameOk

`func (o *DynamicSecretUpdateRedshift) GetRedshiftDbNameOk() (*string, bool)`

GetRedshiftDbNameOk returns a tuple with the RedshiftDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftDbName

`func (o *DynamicSecretUpdateRedshift) SetRedshiftDbName(v string)`

SetRedshiftDbName sets RedshiftDbName field to given value.

### HasRedshiftDbName

`func (o *DynamicSecretUpdateRedshift) HasRedshiftDbName() bool`

HasRedshiftDbName returns a boolean if a field has been set.

### GetRedshiftHost

`func (o *DynamicSecretUpdateRedshift) GetRedshiftHost() string`

GetRedshiftHost returns the RedshiftHost field if non-nil, zero value otherwise.

### GetRedshiftHostOk

`func (o *DynamicSecretUpdateRedshift) GetRedshiftHostOk() (*string, bool)`

GetRedshiftHostOk returns a tuple with the RedshiftHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftHost

`func (o *DynamicSecretUpdateRedshift) SetRedshiftHost(v string)`

SetRedshiftHost sets RedshiftHost field to given value.

### HasRedshiftHost

`func (o *DynamicSecretUpdateRedshift) HasRedshiftHost() bool`

HasRedshiftHost returns a boolean if a field has been set.

### GetRedshiftPassword

`func (o *DynamicSecretUpdateRedshift) GetRedshiftPassword() string`

GetRedshiftPassword returns the RedshiftPassword field if non-nil, zero value otherwise.

### GetRedshiftPasswordOk

`func (o *DynamicSecretUpdateRedshift) GetRedshiftPasswordOk() (*string, bool)`

GetRedshiftPasswordOk returns a tuple with the RedshiftPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftPassword

`func (o *DynamicSecretUpdateRedshift) SetRedshiftPassword(v string)`

SetRedshiftPassword sets RedshiftPassword field to given value.

### HasRedshiftPassword

`func (o *DynamicSecretUpdateRedshift) HasRedshiftPassword() bool`

HasRedshiftPassword returns a boolean if a field has been set.

### GetRedshiftPort

`func (o *DynamicSecretUpdateRedshift) GetRedshiftPort() string`

GetRedshiftPort returns the RedshiftPort field if non-nil, zero value otherwise.

### GetRedshiftPortOk

`func (o *DynamicSecretUpdateRedshift) GetRedshiftPortOk() (*string, bool)`

GetRedshiftPortOk returns a tuple with the RedshiftPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftPort

`func (o *DynamicSecretUpdateRedshift) SetRedshiftPort(v string)`

SetRedshiftPort sets RedshiftPort field to given value.

### HasRedshiftPort

`func (o *DynamicSecretUpdateRedshift) HasRedshiftPort() bool`

HasRedshiftPort returns a boolean if a field has been set.

### GetRedshiftUsername

`func (o *DynamicSecretUpdateRedshift) GetRedshiftUsername() string`

GetRedshiftUsername returns the RedshiftUsername field if non-nil, zero value otherwise.

### GetRedshiftUsernameOk

`func (o *DynamicSecretUpdateRedshift) GetRedshiftUsernameOk() (*string, bool)`

GetRedshiftUsernameOk returns a tuple with the RedshiftUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftUsername

`func (o *DynamicSecretUpdateRedshift) SetRedshiftUsername(v string)`

SetRedshiftUsername sets RedshiftUsername field to given value.

### HasRedshiftUsername

`func (o *DynamicSecretUpdateRedshift) HasRedshiftUsername() bool`

HasRedshiftUsername returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateRedshift) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateRedshift) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateRedshift) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateRedshift) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretUpdateRedshift) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretUpdateRedshift) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretUpdateRedshift) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretUpdateRedshift) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSsl

`func (o *DynamicSecretUpdateRedshift) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *DynamicSecretUpdateRedshift) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *DynamicSecretUpdateRedshift) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *DynamicSecretUpdateRedshift) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateRedshift) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateRedshift) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateRedshift) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateRedshift) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateRedshift) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateRedshift) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateRedshift) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateRedshift) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateRedshift) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateRedshift) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateRedshift) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateRedshift) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateRedshift) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateRedshift) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateRedshift) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateRedshift) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateRedshift) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateRedshift) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateRedshift) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateRedshift) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


