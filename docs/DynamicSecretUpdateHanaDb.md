# DynamicSecretUpdateHanaDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**HanaDbname** | Pointer to **string** | HanaDb Name | [optional] 
**HanadbCreateStatements** | Pointer to **string** | HanaDb Creation statements | [optional] 
**HanadbHost** | Pointer to **string** | HanaDb Host | [optional] [default to "127.0.0.1"]
**HanadbPassword** | Pointer to **string** | HanaDb Password | [optional] 
**HanadbPort** | Pointer to **string** | HanaDb Port | [optional] [default to "443"]
**HanadbRevocationStatements** | Pointer to **string** | HanaDb Revocation statements | [optional] 
**HanadbUsername** | Pointer to **string** | HanaDb Username | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessDbSchema** | Pointer to **string** | The DB schema | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts) | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretUpdateHanaDb

`func NewDynamicSecretUpdateHanaDb(name string, ) *DynamicSecretUpdateHanaDb`

NewDynamicSecretUpdateHanaDb instantiates a new DynamicSecretUpdateHanaDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateHanaDbWithDefaults

`func NewDynamicSecretUpdateHanaDbWithDefaults() *DynamicSecretUpdateHanaDb`

NewDynamicSecretUpdateHanaDbWithDefaults instantiates a new DynamicSecretUpdateHanaDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretUpdateHanaDb) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateHanaDb) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateHanaDb) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateHanaDb) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetHanaDbname

`func (o *DynamicSecretUpdateHanaDb) GetHanaDbname() string`

GetHanaDbname returns the HanaDbname field if non-nil, zero value otherwise.

### GetHanaDbnameOk

`func (o *DynamicSecretUpdateHanaDb) GetHanaDbnameOk() (*string, bool)`

GetHanaDbnameOk returns a tuple with the HanaDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanaDbname

`func (o *DynamicSecretUpdateHanaDb) SetHanaDbname(v string)`

SetHanaDbname sets HanaDbname field to given value.

### HasHanaDbname

`func (o *DynamicSecretUpdateHanaDb) HasHanaDbname() bool`

HasHanaDbname returns a boolean if a field has been set.

### GetHanadbCreateStatements

`func (o *DynamicSecretUpdateHanaDb) GetHanadbCreateStatements() string`

GetHanadbCreateStatements returns the HanadbCreateStatements field if non-nil, zero value otherwise.

### GetHanadbCreateStatementsOk

`func (o *DynamicSecretUpdateHanaDb) GetHanadbCreateStatementsOk() (*string, bool)`

GetHanadbCreateStatementsOk returns a tuple with the HanadbCreateStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbCreateStatements

`func (o *DynamicSecretUpdateHanaDb) SetHanadbCreateStatements(v string)`

SetHanadbCreateStatements sets HanadbCreateStatements field to given value.

### HasHanadbCreateStatements

`func (o *DynamicSecretUpdateHanaDb) HasHanadbCreateStatements() bool`

HasHanadbCreateStatements returns a boolean if a field has been set.

### GetHanadbHost

`func (o *DynamicSecretUpdateHanaDb) GetHanadbHost() string`

GetHanadbHost returns the HanadbHost field if non-nil, zero value otherwise.

### GetHanadbHostOk

`func (o *DynamicSecretUpdateHanaDb) GetHanadbHostOk() (*string, bool)`

GetHanadbHostOk returns a tuple with the HanadbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbHost

`func (o *DynamicSecretUpdateHanaDb) SetHanadbHost(v string)`

SetHanadbHost sets HanadbHost field to given value.

### HasHanadbHost

`func (o *DynamicSecretUpdateHanaDb) HasHanadbHost() bool`

HasHanadbHost returns a boolean if a field has been set.

### GetHanadbPassword

`func (o *DynamicSecretUpdateHanaDb) GetHanadbPassword() string`

GetHanadbPassword returns the HanadbPassword field if non-nil, zero value otherwise.

### GetHanadbPasswordOk

`func (o *DynamicSecretUpdateHanaDb) GetHanadbPasswordOk() (*string, bool)`

GetHanadbPasswordOk returns a tuple with the HanadbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbPassword

`func (o *DynamicSecretUpdateHanaDb) SetHanadbPassword(v string)`

SetHanadbPassword sets HanadbPassword field to given value.

### HasHanadbPassword

`func (o *DynamicSecretUpdateHanaDb) HasHanadbPassword() bool`

HasHanadbPassword returns a boolean if a field has been set.

### GetHanadbPort

`func (o *DynamicSecretUpdateHanaDb) GetHanadbPort() string`

GetHanadbPort returns the HanadbPort field if non-nil, zero value otherwise.

### GetHanadbPortOk

`func (o *DynamicSecretUpdateHanaDb) GetHanadbPortOk() (*string, bool)`

GetHanadbPortOk returns a tuple with the HanadbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbPort

`func (o *DynamicSecretUpdateHanaDb) SetHanadbPort(v string)`

SetHanadbPort sets HanadbPort field to given value.

### HasHanadbPort

`func (o *DynamicSecretUpdateHanaDb) HasHanadbPort() bool`

HasHanadbPort returns a boolean if a field has been set.

### GetHanadbRevocationStatements

`func (o *DynamicSecretUpdateHanaDb) GetHanadbRevocationStatements() string`

GetHanadbRevocationStatements returns the HanadbRevocationStatements field if non-nil, zero value otherwise.

### GetHanadbRevocationStatementsOk

`func (o *DynamicSecretUpdateHanaDb) GetHanadbRevocationStatementsOk() (*string, bool)`

GetHanadbRevocationStatementsOk returns a tuple with the HanadbRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbRevocationStatements

`func (o *DynamicSecretUpdateHanaDb) SetHanadbRevocationStatements(v string)`

SetHanadbRevocationStatements sets HanadbRevocationStatements field to given value.

### HasHanadbRevocationStatements

`func (o *DynamicSecretUpdateHanaDb) HasHanadbRevocationStatements() bool`

HasHanadbRevocationStatements returns a boolean if a field has been set.

### GetHanadbUsername

`func (o *DynamicSecretUpdateHanaDb) GetHanadbUsername() string`

GetHanadbUsername returns the HanadbUsername field if non-nil, zero value otherwise.

### GetHanadbUsernameOk

`func (o *DynamicSecretUpdateHanaDb) GetHanadbUsernameOk() (*string, bool)`

GetHanadbUsernameOk returns a tuple with the HanadbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbUsername

`func (o *DynamicSecretUpdateHanaDb) SetHanadbUsername(v string)`

SetHanadbUsername sets HanadbUsername field to given value.

### HasHanadbUsername

`func (o *DynamicSecretUpdateHanaDb) HasHanadbUsername() bool`

HasHanadbUsername returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateHanaDb) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateHanaDb) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateHanaDb) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateHanaDb) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateHanaDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateHanaDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateHanaDb) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateHanaDb) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateHanaDb) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateHanaDb) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateHanaDb) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateHanaDb) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateHanaDb) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateHanaDb) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateHanaDb) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateHanaDb) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateHanaDb) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessDbSchema

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessDbSchema() string`

GetSecureAccessDbSchema returns the SecureAccessDbSchema field if non-nil, zero value otherwise.

### GetSecureAccessDbSchemaOk

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessDbSchemaOk() (*string, bool)`

GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbSchema

`func (o *DynamicSecretUpdateHanaDb) SetSecureAccessDbSchema(v string)`

SetSecureAccessDbSchema sets SecureAccessDbSchema field to given value.

### HasSecureAccessDbSchema

`func (o *DynamicSecretUpdateHanaDb) HasSecureAccessDbSchema() bool`

HasSecureAccessDbSchema returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateHanaDb) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateHanaDb) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretUpdateHanaDb) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretUpdateHanaDb) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretUpdateHanaDb) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretUpdateHanaDb) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretUpdateHanaDb) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateHanaDb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateHanaDb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateHanaDb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateHanaDb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateHanaDb) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateHanaDb) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateHanaDb) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateHanaDb) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateHanaDb) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateHanaDb) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateHanaDb) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateHanaDb) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateHanaDb) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateHanaDb) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateHanaDb) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateHanaDb) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateHanaDb) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateHanaDb) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateHanaDb) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateHanaDb) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


