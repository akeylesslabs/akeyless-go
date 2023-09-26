# DsCreateHanaDb

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

### NewDsCreateHanaDb

`func NewDsCreateHanaDb(name string, ) *DsCreateHanaDb`

NewDsCreateHanaDb instantiates a new DsCreateHanaDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsCreateHanaDbWithDefaults

`func NewDsCreateHanaDbWithDefaults() *DsCreateHanaDb`

NewDsCreateHanaDbWithDefaults instantiates a new DsCreateHanaDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DsCreateHanaDb) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DsCreateHanaDb) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DsCreateHanaDb) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DsCreateHanaDb) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetHanaDbname

`func (o *DsCreateHanaDb) GetHanaDbname() string`

GetHanaDbname returns the HanaDbname field if non-nil, zero value otherwise.

### GetHanaDbnameOk

`func (o *DsCreateHanaDb) GetHanaDbnameOk() (*string, bool)`

GetHanaDbnameOk returns a tuple with the HanaDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanaDbname

`func (o *DsCreateHanaDb) SetHanaDbname(v string)`

SetHanaDbname sets HanaDbname field to given value.

### HasHanaDbname

`func (o *DsCreateHanaDb) HasHanaDbname() bool`

HasHanaDbname returns a boolean if a field has been set.

### GetHanadbCreateStatements

`func (o *DsCreateHanaDb) GetHanadbCreateStatements() string`

GetHanadbCreateStatements returns the HanadbCreateStatements field if non-nil, zero value otherwise.

### GetHanadbCreateStatementsOk

`func (o *DsCreateHanaDb) GetHanadbCreateStatementsOk() (*string, bool)`

GetHanadbCreateStatementsOk returns a tuple with the HanadbCreateStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbCreateStatements

`func (o *DsCreateHanaDb) SetHanadbCreateStatements(v string)`

SetHanadbCreateStatements sets HanadbCreateStatements field to given value.

### HasHanadbCreateStatements

`func (o *DsCreateHanaDb) HasHanadbCreateStatements() bool`

HasHanadbCreateStatements returns a boolean if a field has been set.

### GetHanadbHost

`func (o *DsCreateHanaDb) GetHanadbHost() string`

GetHanadbHost returns the HanadbHost field if non-nil, zero value otherwise.

### GetHanadbHostOk

`func (o *DsCreateHanaDb) GetHanadbHostOk() (*string, bool)`

GetHanadbHostOk returns a tuple with the HanadbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbHost

`func (o *DsCreateHanaDb) SetHanadbHost(v string)`

SetHanadbHost sets HanadbHost field to given value.

### HasHanadbHost

`func (o *DsCreateHanaDb) HasHanadbHost() bool`

HasHanadbHost returns a boolean if a field has been set.

### GetHanadbPassword

`func (o *DsCreateHanaDb) GetHanadbPassword() string`

GetHanadbPassword returns the HanadbPassword field if non-nil, zero value otherwise.

### GetHanadbPasswordOk

`func (o *DsCreateHanaDb) GetHanadbPasswordOk() (*string, bool)`

GetHanadbPasswordOk returns a tuple with the HanadbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbPassword

`func (o *DsCreateHanaDb) SetHanadbPassword(v string)`

SetHanadbPassword sets HanadbPassword field to given value.

### HasHanadbPassword

`func (o *DsCreateHanaDb) HasHanadbPassword() bool`

HasHanadbPassword returns a boolean if a field has been set.

### GetHanadbPort

`func (o *DsCreateHanaDb) GetHanadbPort() string`

GetHanadbPort returns the HanadbPort field if non-nil, zero value otherwise.

### GetHanadbPortOk

`func (o *DsCreateHanaDb) GetHanadbPortOk() (*string, bool)`

GetHanadbPortOk returns a tuple with the HanadbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbPort

`func (o *DsCreateHanaDb) SetHanadbPort(v string)`

SetHanadbPort sets HanadbPort field to given value.

### HasHanadbPort

`func (o *DsCreateHanaDb) HasHanadbPort() bool`

HasHanadbPort returns a boolean if a field has been set.

### GetHanadbRevocationStatements

`func (o *DsCreateHanaDb) GetHanadbRevocationStatements() string`

GetHanadbRevocationStatements returns the HanadbRevocationStatements field if non-nil, zero value otherwise.

### GetHanadbRevocationStatementsOk

`func (o *DsCreateHanaDb) GetHanadbRevocationStatementsOk() (*string, bool)`

GetHanadbRevocationStatementsOk returns a tuple with the HanadbRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbRevocationStatements

`func (o *DsCreateHanaDb) SetHanadbRevocationStatements(v string)`

SetHanadbRevocationStatements sets HanadbRevocationStatements field to given value.

### HasHanadbRevocationStatements

`func (o *DsCreateHanaDb) HasHanadbRevocationStatements() bool`

HasHanadbRevocationStatements returns a boolean if a field has been set.

### GetHanadbUsername

`func (o *DsCreateHanaDb) GetHanadbUsername() string`

GetHanadbUsername returns the HanadbUsername field if non-nil, zero value otherwise.

### GetHanadbUsernameOk

`func (o *DsCreateHanaDb) GetHanadbUsernameOk() (*string, bool)`

GetHanadbUsernameOk returns a tuple with the HanadbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbUsername

`func (o *DsCreateHanaDb) SetHanadbUsername(v string)`

SetHanadbUsername sets HanadbUsername field to given value.

### HasHanadbUsername

`func (o *DsCreateHanaDb) HasHanadbUsername() bool`

HasHanadbUsername returns a boolean if a field has been set.

### GetJson

`func (o *DsCreateHanaDb) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DsCreateHanaDb) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DsCreateHanaDb) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DsCreateHanaDb) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DsCreateHanaDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DsCreateHanaDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DsCreateHanaDb) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DsCreateHanaDb) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DsCreateHanaDb) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DsCreateHanaDb) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DsCreateHanaDb) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DsCreateHanaDb) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DsCreateHanaDb) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DsCreateHanaDb) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DsCreateHanaDb) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessDbSchema

`func (o *DsCreateHanaDb) GetSecureAccessDbSchema() string`

GetSecureAccessDbSchema returns the SecureAccessDbSchema field if non-nil, zero value otherwise.

### GetSecureAccessDbSchemaOk

`func (o *DsCreateHanaDb) GetSecureAccessDbSchemaOk() (*string, bool)`

GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbSchema

`func (o *DsCreateHanaDb) SetSecureAccessDbSchema(v string)`

SetSecureAccessDbSchema sets SecureAccessDbSchema field to given value.

### HasSecureAccessDbSchema

`func (o *DsCreateHanaDb) HasSecureAccessDbSchema() bool`

HasSecureAccessDbSchema returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DsCreateHanaDb) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DsCreateHanaDb) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DsCreateHanaDb) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DsCreateHanaDb) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DsCreateHanaDb) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DsCreateHanaDb) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DsCreateHanaDb) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DsCreateHanaDb) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DsCreateHanaDb) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DsCreateHanaDb) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DsCreateHanaDb) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DsCreateHanaDb) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DsCreateHanaDb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DsCreateHanaDb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DsCreateHanaDb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DsCreateHanaDb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DsCreateHanaDb) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DsCreateHanaDb) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DsCreateHanaDb) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DsCreateHanaDb) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DsCreateHanaDb) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DsCreateHanaDb) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DsCreateHanaDb) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DsCreateHanaDb) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DsCreateHanaDb) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DsCreateHanaDb) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DsCreateHanaDb) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DsCreateHanaDb) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DsCreateHanaDb) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DsCreateHanaDb) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DsCreateHanaDb) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DsCreateHanaDb) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


