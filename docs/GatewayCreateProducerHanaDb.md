# GatewayCreateProducerHanaDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. Mirrors commands.AgenticRulesParams.AraEnabled. | [optional] 
**CustomUsernameTemplate** | Pointer to **string** | Customize how temporary usernames are generated using go template | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**HanaDbname** | Pointer to **string** | HanaDb Name | [optional] 
**HanadbCreateStatements** | Pointer to **string** | HanaDb Creation statements | [optional] 
**HanadbHost** | Pointer to **string** | HanaDb Host | [optional] [default to "127.0.0.1"]
**HanadbPassword** | Pointer to **string** | HanaDb Password | [optional] 
**HanadbPort** | Pointer to **string** | HanaDb Port | [optional] [default to "443"]
**HanadbRevocationStatements** | Pointer to **string** | HanaDb Revocation statements | [optional] 
**HanadbUsername** | Pointer to **string** | HanaDb Username | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) Mirrors commands.AgenticRulesParams — kept separate because ResourceDS cannot embed it (different package, different struct layout). | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Deprecated. use secure-access-certificate-issuer | [optional] 
**SecureAccessCertificateIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Secure Access | [optional] 
**SecureAccessDbName** | Pointer to **string** | The DB name (relevant only for DB Dynamic-Secret) | [optional] 
**SecureAccessDbSchema** | Pointer to **string** | The DB schema | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts) | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**SkipDryRun** | Pointer to **string** | If set, dry-run will be skipped | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseCapitalLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one uppercase character from the ISO basic Latin alphabet (A to Z). [true/false] | [optional] 
**UseLowerLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one lowercase character from the ISO basic Latin alphabet (a to z). [true/false] | [optional] 
**UseNumbers** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one numeric character (0 to 9). [true/false] | [optional] 
**UseSpecialCharacters** | Pointer to **string** |  | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerHanaDb

`func NewGatewayCreateProducerHanaDb(name string, ) *GatewayCreateProducerHanaDb`

NewGatewayCreateProducerHanaDb instantiates a new GatewayCreateProducerHanaDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerHanaDbWithDefaults

`func NewGatewayCreateProducerHanaDbWithDefaults() *GatewayCreateProducerHanaDb`

NewGatewayCreateProducerHanaDbWithDefaults instantiates a new GatewayCreateProducerHanaDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAraEnabled

`func (o *GatewayCreateProducerHanaDb) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *GatewayCreateProducerHanaDb) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *GatewayCreateProducerHanaDb) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *GatewayCreateProducerHanaDb) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetCustomUsernameTemplate

`func (o *GatewayCreateProducerHanaDb) GetCustomUsernameTemplate() string`

GetCustomUsernameTemplate returns the CustomUsernameTemplate field if non-nil, zero value otherwise.

### GetCustomUsernameTemplateOk

`func (o *GatewayCreateProducerHanaDb) GetCustomUsernameTemplateOk() (*string, bool)`

GetCustomUsernameTemplateOk returns a tuple with the CustomUsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUsernameTemplate

`func (o *GatewayCreateProducerHanaDb) SetCustomUsernameTemplate(v string)`

SetCustomUsernameTemplate sets CustomUsernameTemplate field to given value.

### HasCustomUsernameTemplate

`func (o *GatewayCreateProducerHanaDb) HasCustomUsernameTemplate() bool`

HasCustomUsernameTemplate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayCreateProducerHanaDb) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayCreateProducerHanaDb) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayCreateProducerHanaDb) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayCreateProducerHanaDb) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetHanaDbname

`func (o *GatewayCreateProducerHanaDb) GetHanaDbname() string`

GetHanaDbname returns the HanaDbname field if non-nil, zero value otherwise.

### GetHanaDbnameOk

`func (o *GatewayCreateProducerHanaDb) GetHanaDbnameOk() (*string, bool)`

GetHanaDbnameOk returns a tuple with the HanaDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanaDbname

`func (o *GatewayCreateProducerHanaDb) SetHanaDbname(v string)`

SetHanaDbname sets HanaDbname field to given value.

### HasHanaDbname

`func (o *GatewayCreateProducerHanaDb) HasHanaDbname() bool`

HasHanaDbname returns a boolean if a field has been set.

### GetHanadbCreateStatements

`func (o *GatewayCreateProducerHanaDb) GetHanadbCreateStatements() string`

GetHanadbCreateStatements returns the HanadbCreateStatements field if non-nil, zero value otherwise.

### GetHanadbCreateStatementsOk

`func (o *GatewayCreateProducerHanaDb) GetHanadbCreateStatementsOk() (*string, bool)`

GetHanadbCreateStatementsOk returns a tuple with the HanadbCreateStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbCreateStatements

`func (o *GatewayCreateProducerHanaDb) SetHanadbCreateStatements(v string)`

SetHanadbCreateStatements sets HanadbCreateStatements field to given value.

### HasHanadbCreateStatements

`func (o *GatewayCreateProducerHanaDb) HasHanadbCreateStatements() bool`

HasHanadbCreateStatements returns a boolean if a field has been set.

### GetHanadbHost

`func (o *GatewayCreateProducerHanaDb) GetHanadbHost() string`

GetHanadbHost returns the HanadbHost field if non-nil, zero value otherwise.

### GetHanadbHostOk

`func (o *GatewayCreateProducerHanaDb) GetHanadbHostOk() (*string, bool)`

GetHanadbHostOk returns a tuple with the HanadbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbHost

`func (o *GatewayCreateProducerHanaDb) SetHanadbHost(v string)`

SetHanadbHost sets HanadbHost field to given value.

### HasHanadbHost

`func (o *GatewayCreateProducerHanaDb) HasHanadbHost() bool`

HasHanadbHost returns a boolean if a field has been set.

### GetHanadbPassword

`func (o *GatewayCreateProducerHanaDb) GetHanadbPassword() string`

GetHanadbPassword returns the HanadbPassword field if non-nil, zero value otherwise.

### GetHanadbPasswordOk

`func (o *GatewayCreateProducerHanaDb) GetHanadbPasswordOk() (*string, bool)`

GetHanadbPasswordOk returns a tuple with the HanadbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbPassword

`func (o *GatewayCreateProducerHanaDb) SetHanadbPassword(v string)`

SetHanadbPassword sets HanadbPassword field to given value.

### HasHanadbPassword

`func (o *GatewayCreateProducerHanaDb) HasHanadbPassword() bool`

HasHanadbPassword returns a boolean if a field has been set.

### GetHanadbPort

`func (o *GatewayCreateProducerHanaDb) GetHanadbPort() string`

GetHanadbPort returns the HanadbPort field if non-nil, zero value otherwise.

### GetHanadbPortOk

`func (o *GatewayCreateProducerHanaDb) GetHanadbPortOk() (*string, bool)`

GetHanadbPortOk returns a tuple with the HanadbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbPort

`func (o *GatewayCreateProducerHanaDb) SetHanadbPort(v string)`

SetHanadbPort sets HanadbPort field to given value.

### HasHanadbPort

`func (o *GatewayCreateProducerHanaDb) HasHanadbPort() bool`

HasHanadbPort returns a boolean if a field has been set.

### GetHanadbRevocationStatements

`func (o *GatewayCreateProducerHanaDb) GetHanadbRevocationStatements() string`

GetHanadbRevocationStatements returns the HanadbRevocationStatements field if non-nil, zero value otherwise.

### GetHanadbRevocationStatementsOk

`func (o *GatewayCreateProducerHanaDb) GetHanadbRevocationStatementsOk() (*string, bool)`

GetHanadbRevocationStatementsOk returns a tuple with the HanadbRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbRevocationStatements

`func (o *GatewayCreateProducerHanaDb) SetHanadbRevocationStatements(v string)`

SetHanadbRevocationStatements sets HanadbRevocationStatements field to given value.

### HasHanadbRevocationStatements

`func (o *GatewayCreateProducerHanaDb) HasHanadbRevocationStatements() bool`

HasHanadbRevocationStatements returns a boolean if a field has been set.

### GetHanadbUsername

`func (o *GatewayCreateProducerHanaDb) GetHanadbUsername() string`

GetHanadbUsername returns the HanadbUsername field if non-nil, zero value otherwise.

### GetHanadbUsernameOk

`func (o *GatewayCreateProducerHanaDb) GetHanadbUsernameOk() (*string, bool)`

GetHanadbUsernameOk returns a tuple with the HanadbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbUsername

`func (o *GatewayCreateProducerHanaDb) SetHanadbUsername(v string)`

SetHanadbUsername sets HanadbUsername field to given value.

### HasHanadbUsername

`func (o *GatewayCreateProducerHanaDb) HasHanadbUsername() bool`

HasHanadbUsername returns a boolean if a field has been set.

### GetInputRule

`func (o *GatewayCreateProducerHanaDb) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *GatewayCreateProducerHanaDb) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *GatewayCreateProducerHanaDb) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *GatewayCreateProducerHanaDb) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *GatewayCreateProducerHanaDb) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *GatewayCreateProducerHanaDb) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *GatewayCreateProducerHanaDb) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *GatewayCreateProducerHanaDb) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateProducerHanaDb) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateProducerHanaDb) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateProducerHanaDb) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateProducerHanaDb) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerHanaDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerHanaDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerHanaDb) SetName(v string)`

SetName sets Name field to given value.


### GetOutputRule

`func (o *GatewayCreateProducerHanaDb) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *GatewayCreateProducerHanaDb) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *GatewayCreateProducerHanaDb) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *GatewayCreateProducerHanaDb) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetPasswordLength

`func (o *GatewayCreateProducerHanaDb) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *GatewayCreateProducerHanaDb) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *GatewayCreateProducerHanaDb) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *GatewayCreateProducerHanaDb) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerHanaDb) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerHanaDb) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerHanaDb) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerHanaDb) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerHanaDb) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayCreateProducerHanaDb) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessCertificateIssuer

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessCertificateIssuer() string`

GetSecureAccessCertificateIssuer returns the SecureAccessCertificateIssuer field if non-nil, zero value otherwise.

### GetSecureAccessCertificateIssuerOk

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessCertificateIssuerOk() (*string, bool)`

GetSecureAccessCertificateIssuerOk returns a tuple with the SecureAccessCertificateIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessCertificateIssuer

`func (o *GatewayCreateProducerHanaDb) SetSecureAccessCertificateIssuer(v string)`

SetSecureAccessCertificateIssuer sets SecureAccessCertificateIssuer field to given value.

### HasSecureAccessCertificateIssuer

`func (o *GatewayCreateProducerHanaDb) HasSecureAccessCertificateIssuer() bool`

HasSecureAccessCertificateIssuer returns a boolean if a field has been set.

### GetSecureAccessDbName

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessDbName() string`

GetSecureAccessDbName returns the SecureAccessDbName field if non-nil, zero value otherwise.

### GetSecureAccessDbNameOk

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessDbNameOk() (*string, bool)`

GetSecureAccessDbNameOk returns a tuple with the SecureAccessDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbName

`func (o *GatewayCreateProducerHanaDb) SetSecureAccessDbName(v string)`

SetSecureAccessDbName sets SecureAccessDbName field to given value.

### HasSecureAccessDbName

`func (o *GatewayCreateProducerHanaDb) HasSecureAccessDbName() bool`

HasSecureAccessDbName returns a boolean if a field has been set.

### GetSecureAccessDbSchema

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessDbSchema() string`

GetSecureAccessDbSchema returns the SecureAccessDbSchema field if non-nil, zero value otherwise.

### GetSecureAccessDbSchemaOk

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessDbSchemaOk() (*string, bool)`

GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbSchema

`func (o *GatewayCreateProducerHanaDb) SetSecureAccessDbSchema(v string)`

SetSecureAccessDbSchema sets SecureAccessDbSchema field to given value.

### HasSecureAccessDbSchema

`func (o *GatewayCreateProducerHanaDb) HasSecureAccessDbSchema() bool`

HasSecureAccessDbSchema returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayCreateProducerHanaDb) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayCreateProducerHanaDb) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayCreateProducerHanaDb) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayCreateProducerHanaDb) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayCreateProducerHanaDb) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayCreateProducerHanaDb) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayCreateProducerHanaDb) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSkipDryRun

`func (o *GatewayCreateProducerHanaDb) GetSkipDryRun() string`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *GatewayCreateProducerHanaDb) GetSkipDryRunOk() (*string, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *GatewayCreateProducerHanaDb) SetSkipDryRun(v string)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *GatewayCreateProducerHanaDb) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerHanaDb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerHanaDb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerHanaDb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerHanaDb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerHanaDb) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerHanaDb) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerHanaDb) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerHanaDb) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerHanaDb) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerHanaDb) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerHanaDb) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerHanaDb) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerHanaDb) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerHanaDb) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerHanaDb) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerHanaDb) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseCapitalLetters

`func (o *GatewayCreateProducerHanaDb) GetUseCapitalLetters() string`

GetUseCapitalLetters returns the UseCapitalLetters field if non-nil, zero value otherwise.

### GetUseCapitalLettersOk

`func (o *GatewayCreateProducerHanaDb) GetUseCapitalLettersOk() (*string, bool)`

GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCapitalLetters

`func (o *GatewayCreateProducerHanaDb) SetUseCapitalLetters(v string)`

SetUseCapitalLetters sets UseCapitalLetters field to given value.

### HasUseCapitalLetters

`func (o *GatewayCreateProducerHanaDb) HasUseCapitalLetters() bool`

HasUseCapitalLetters returns a boolean if a field has been set.

### GetUseLowerLetters

`func (o *GatewayCreateProducerHanaDb) GetUseLowerLetters() string`

GetUseLowerLetters returns the UseLowerLetters field if non-nil, zero value otherwise.

### GetUseLowerLettersOk

`func (o *GatewayCreateProducerHanaDb) GetUseLowerLettersOk() (*string, bool)`

GetUseLowerLettersOk returns a tuple with the UseLowerLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLowerLetters

`func (o *GatewayCreateProducerHanaDb) SetUseLowerLetters(v string)`

SetUseLowerLetters sets UseLowerLetters field to given value.

### HasUseLowerLetters

`func (o *GatewayCreateProducerHanaDb) HasUseLowerLetters() bool`

HasUseLowerLetters returns a boolean if a field has been set.

### GetUseNumbers

`func (o *GatewayCreateProducerHanaDb) GetUseNumbers() string`

GetUseNumbers returns the UseNumbers field if non-nil, zero value otherwise.

### GetUseNumbersOk

`func (o *GatewayCreateProducerHanaDb) GetUseNumbersOk() (*string, bool)`

GetUseNumbersOk returns a tuple with the UseNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumbers

`func (o *GatewayCreateProducerHanaDb) SetUseNumbers(v string)`

SetUseNumbers sets UseNumbers field to given value.

### HasUseNumbers

`func (o *GatewayCreateProducerHanaDb) HasUseNumbers() bool`

HasUseNumbers returns a boolean if a field has been set.

### GetUseSpecialCharacters

`func (o *GatewayCreateProducerHanaDb) GetUseSpecialCharacters() string`

GetUseSpecialCharacters returns the UseSpecialCharacters field if non-nil, zero value otherwise.

### GetUseSpecialCharactersOk

`func (o *GatewayCreateProducerHanaDb) GetUseSpecialCharactersOk() (*string, bool)`

GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpecialCharacters

`func (o *GatewayCreateProducerHanaDb) SetUseSpecialCharacters(v string)`

SetUseSpecialCharacters sets UseSpecialCharacters field to given value.

### HasUseSpecialCharacters

`func (o *GatewayCreateProducerHanaDb) HasUseSpecialCharacters() bool`

HasUseSpecialCharacters returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerHanaDb) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerHanaDb) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerHanaDb) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerHanaDb) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


