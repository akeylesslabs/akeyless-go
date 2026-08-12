# GatewayUpdateProducerRdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | Pointer to **string** |  | [optional] 
**AllowUserExtendSession** | Pointer to **int64** | AllowUserExtendSession | [optional] 
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. Mirrors commands.AgenticRulesParams.AraEnabled. | [optional] 
**CustomUsernameTemplate** | Pointer to **string** | Customize how temporary usernames are generated using go template | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**FixedUserClaimKeyname** | Pointer to **string** | For externally provided users, denotes the key-name of IdP claim to extract the username from (relevant only for fixed-user-only&#x3D;true) | [optional] [default to "ext_username"]
**FixedUserOnly** | Pointer to **string** | Allow access using externally (IdP) provided username [true/false] | [optional] [default to "false"]
**HostProvider** | Pointer to **string** | Host provider type [explicit/target], Default Host provider is explicit, Relevant only for SRA items. | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) Mirrors commands.AgenticRulesParams — kept separate because ResourceDS cannot embed it (different package, different struct layout). | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RdpAdminName** | Pointer to **string** | RDP Admin Name | [optional] 
**RdpAdminPwd** | Pointer to **string** | RDP Admin password | [optional] 
**RdpHostName** | Pointer to **string** | Hostname | [optional] 
**RdpHostPort** | Pointer to **string** | Port | [optional] [default to "22"]
**RdpUserGroups** | Pointer to **string** | Groups | [optional] 
**SecureAccessAllowExternalUser** | Pointer to **bool** | Allow providing external user for a domain users | [optional] [default to false]
**SecureAccessBastionIssuer** | Pointer to **string** | Deprecated. use secure-access-certificate-issuer | [optional] 
**SecureAccessCertificateIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Secure Access | [optional] 
**SecureAccessDelay** | Pointer to **int64** | The delay duration, in seconds, to wait after generating just-in-time credentials. Accepted range: 0-120 seconds | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessEnforceHostsRestriction** | Pointer to **bool** | Enforce connections only to allowed SRA hosts | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessRdGatewayServer** | Pointer to **string** | RD Gateway server | [optional] 
**SecureAccessRdpDomain** | Pointer to **string** | Required when the Dynamic Secret is used for a domain user | [optional] 
**SecureAccessRdpUser** | Pointer to **string** | Override the RDP Domain username | [optional] 
**SkipDryRun** | Pointer to **string** | If set, dry-run will be skipped | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**Target** | Pointer to **[]string** | A list of targets to be associated with an SRA item, To specify multiple targets use argument multiple times | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseCapitalLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one uppercase character from the ISO basic Latin alphabet (A to Z). [true/false] | [optional] 
**UseLowerLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one lowercase character from the ISO basic Latin alphabet (a to z). [true/false] | [optional] 
**UseNumbers** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one numeric character (0 to 9). [true/false] | [optional] 
**UseSpecialCharacters** | Pointer to **string** |  | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**WarnUserBeforeExpiration** | Pointer to **int64** | WarnBeforeUserExpiration | [optional] 

## Methods

### NewGatewayUpdateProducerRdp

`func NewGatewayUpdateProducerRdp(name string, ) *GatewayUpdateProducerRdp`

NewGatewayUpdateProducerRdp instantiates a new GatewayUpdateProducerRdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerRdpWithDefaults

`func NewGatewayUpdateProducerRdpWithDefaults() *GatewayUpdateProducerRdp`

NewGatewayUpdateProducerRdpWithDefaults instantiates a new GatewayUpdateProducerRdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *GatewayUpdateProducerRdp) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *GatewayUpdateProducerRdp) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *GatewayUpdateProducerRdp) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *GatewayUpdateProducerRdp) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetAllowUserExtendSession

`func (o *GatewayUpdateProducerRdp) GetAllowUserExtendSession() int64`

GetAllowUserExtendSession returns the AllowUserExtendSession field if non-nil, zero value otherwise.

### GetAllowUserExtendSessionOk

`func (o *GatewayUpdateProducerRdp) GetAllowUserExtendSessionOk() (*int64, bool)`

GetAllowUserExtendSessionOk returns a tuple with the AllowUserExtendSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserExtendSession

`func (o *GatewayUpdateProducerRdp) SetAllowUserExtendSession(v int64)`

SetAllowUserExtendSession sets AllowUserExtendSession field to given value.

### HasAllowUserExtendSession

`func (o *GatewayUpdateProducerRdp) HasAllowUserExtendSession() bool`

HasAllowUserExtendSession returns a boolean if a field has been set.

### GetAraEnabled

`func (o *GatewayUpdateProducerRdp) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *GatewayUpdateProducerRdp) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *GatewayUpdateProducerRdp) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *GatewayUpdateProducerRdp) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetCustomUsernameTemplate

`func (o *GatewayUpdateProducerRdp) GetCustomUsernameTemplate() string`

GetCustomUsernameTemplate returns the CustomUsernameTemplate field if non-nil, zero value otherwise.

### GetCustomUsernameTemplateOk

`func (o *GatewayUpdateProducerRdp) GetCustomUsernameTemplateOk() (*string, bool)`

GetCustomUsernameTemplateOk returns a tuple with the CustomUsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUsernameTemplate

`func (o *GatewayUpdateProducerRdp) SetCustomUsernameTemplate(v string)`

SetCustomUsernameTemplate sets CustomUsernameTemplate field to given value.

### HasCustomUsernameTemplate

`func (o *GatewayUpdateProducerRdp) HasCustomUsernameTemplate() bool`

HasCustomUsernameTemplate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayUpdateProducerRdp) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerRdp) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerRdp) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerRdp) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetFixedUserClaimKeyname

`func (o *GatewayUpdateProducerRdp) GetFixedUserClaimKeyname() string`

GetFixedUserClaimKeyname returns the FixedUserClaimKeyname field if non-nil, zero value otherwise.

### GetFixedUserClaimKeynameOk

`func (o *GatewayUpdateProducerRdp) GetFixedUserClaimKeynameOk() (*string, bool)`

GetFixedUserClaimKeynameOk returns a tuple with the FixedUserClaimKeyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserClaimKeyname

`func (o *GatewayUpdateProducerRdp) SetFixedUserClaimKeyname(v string)`

SetFixedUserClaimKeyname sets FixedUserClaimKeyname field to given value.

### HasFixedUserClaimKeyname

`func (o *GatewayUpdateProducerRdp) HasFixedUserClaimKeyname() bool`

HasFixedUserClaimKeyname returns a boolean if a field has been set.

### GetFixedUserOnly

`func (o *GatewayUpdateProducerRdp) GetFixedUserOnly() string`

GetFixedUserOnly returns the FixedUserOnly field if non-nil, zero value otherwise.

### GetFixedUserOnlyOk

`func (o *GatewayUpdateProducerRdp) GetFixedUserOnlyOk() (*string, bool)`

GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserOnly

`func (o *GatewayUpdateProducerRdp) SetFixedUserOnly(v string)`

SetFixedUserOnly sets FixedUserOnly field to given value.

### HasFixedUserOnly

`func (o *GatewayUpdateProducerRdp) HasFixedUserOnly() bool`

HasFixedUserOnly returns a boolean if a field has been set.

### GetHostProvider

`func (o *GatewayUpdateProducerRdp) GetHostProvider() string`

GetHostProvider returns the HostProvider field if non-nil, zero value otherwise.

### GetHostProviderOk

`func (o *GatewayUpdateProducerRdp) GetHostProviderOk() (*string, bool)`

GetHostProviderOk returns a tuple with the HostProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostProvider

`func (o *GatewayUpdateProducerRdp) SetHostProvider(v string)`

SetHostProvider sets HostProvider field to given value.

### HasHostProvider

`func (o *GatewayUpdateProducerRdp) HasHostProvider() bool`

HasHostProvider returns a boolean if a field has been set.

### GetInputRule

`func (o *GatewayUpdateProducerRdp) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *GatewayUpdateProducerRdp) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *GatewayUpdateProducerRdp) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *GatewayUpdateProducerRdp) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *GatewayUpdateProducerRdp) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *GatewayUpdateProducerRdp) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *GatewayUpdateProducerRdp) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *GatewayUpdateProducerRdp) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateProducerRdp) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateProducerRdp) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateProducerRdp) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateProducerRdp) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerRdp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerRdp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerRdp) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerRdp) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerRdp) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerRdp) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerRdp) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetOutputRule

`func (o *GatewayUpdateProducerRdp) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *GatewayUpdateProducerRdp) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *GatewayUpdateProducerRdp) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *GatewayUpdateProducerRdp) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetPasswordLength

`func (o *GatewayUpdateProducerRdp) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *GatewayUpdateProducerRdp) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *GatewayUpdateProducerRdp) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *GatewayUpdateProducerRdp) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerRdp) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerRdp) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerRdp) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerRdp) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRdpAdminName

`func (o *GatewayUpdateProducerRdp) GetRdpAdminName() string`

GetRdpAdminName returns the RdpAdminName field if non-nil, zero value otherwise.

### GetRdpAdminNameOk

`func (o *GatewayUpdateProducerRdp) GetRdpAdminNameOk() (*string, bool)`

GetRdpAdminNameOk returns a tuple with the RdpAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminName

`func (o *GatewayUpdateProducerRdp) SetRdpAdminName(v string)`

SetRdpAdminName sets RdpAdminName field to given value.

### HasRdpAdminName

`func (o *GatewayUpdateProducerRdp) HasRdpAdminName() bool`

HasRdpAdminName returns a boolean if a field has been set.

### GetRdpAdminPwd

`func (o *GatewayUpdateProducerRdp) GetRdpAdminPwd() string`

GetRdpAdminPwd returns the RdpAdminPwd field if non-nil, zero value otherwise.

### GetRdpAdminPwdOk

`func (o *GatewayUpdateProducerRdp) GetRdpAdminPwdOk() (*string, bool)`

GetRdpAdminPwdOk returns a tuple with the RdpAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminPwd

`func (o *GatewayUpdateProducerRdp) SetRdpAdminPwd(v string)`

SetRdpAdminPwd sets RdpAdminPwd field to given value.

### HasRdpAdminPwd

`func (o *GatewayUpdateProducerRdp) HasRdpAdminPwd() bool`

HasRdpAdminPwd returns a boolean if a field has been set.

### GetRdpHostName

`func (o *GatewayUpdateProducerRdp) GetRdpHostName() string`

GetRdpHostName returns the RdpHostName field if non-nil, zero value otherwise.

### GetRdpHostNameOk

`func (o *GatewayUpdateProducerRdp) GetRdpHostNameOk() (*string, bool)`

GetRdpHostNameOk returns a tuple with the RdpHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostName

`func (o *GatewayUpdateProducerRdp) SetRdpHostName(v string)`

SetRdpHostName sets RdpHostName field to given value.

### HasRdpHostName

`func (o *GatewayUpdateProducerRdp) HasRdpHostName() bool`

HasRdpHostName returns a boolean if a field has been set.

### GetRdpHostPort

`func (o *GatewayUpdateProducerRdp) GetRdpHostPort() string`

GetRdpHostPort returns the RdpHostPort field if non-nil, zero value otherwise.

### GetRdpHostPortOk

`func (o *GatewayUpdateProducerRdp) GetRdpHostPortOk() (*string, bool)`

GetRdpHostPortOk returns a tuple with the RdpHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostPort

`func (o *GatewayUpdateProducerRdp) SetRdpHostPort(v string)`

SetRdpHostPort sets RdpHostPort field to given value.

### HasRdpHostPort

`func (o *GatewayUpdateProducerRdp) HasRdpHostPort() bool`

HasRdpHostPort returns a boolean if a field has been set.

### GetRdpUserGroups

`func (o *GatewayUpdateProducerRdp) GetRdpUserGroups() string`

GetRdpUserGroups returns the RdpUserGroups field if non-nil, zero value otherwise.

### GetRdpUserGroupsOk

`func (o *GatewayUpdateProducerRdp) GetRdpUserGroupsOk() (*string, bool)`

GetRdpUserGroupsOk returns a tuple with the RdpUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpUserGroups

`func (o *GatewayUpdateProducerRdp) SetRdpUserGroups(v string)`

SetRdpUserGroups sets RdpUserGroups field to given value.

### HasRdpUserGroups

`func (o *GatewayUpdateProducerRdp) HasRdpUserGroups() bool`

HasRdpUserGroups returns a boolean if a field has been set.

### GetSecureAccessAllowExternalUser

`func (o *GatewayUpdateProducerRdp) GetSecureAccessAllowExternalUser() bool`

GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field if non-nil, zero value otherwise.

### GetSecureAccessAllowExternalUserOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessAllowExternalUserOk() (*bool, bool)`

GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowExternalUser

`func (o *GatewayUpdateProducerRdp) SetSecureAccessAllowExternalUser(v bool)`

SetSecureAccessAllowExternalUser sets SecureAccessAllowExternalUser field to given value.

### HasSecureAccessAllowExternalUser

`func (o *GatewayUpdateProducerRdp) HasSecureAccessAllowExternalUser() bool`

HasSecureAccessAllowExternalUser returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerRdp) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerRdp) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerRdp) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessCertificateIssuer

`func (o *GatewayUpdateProducerRdp) GetSecureAccessCertificateIssuer() string`

GetSecureAccessCertificateIssuer returns the SecureAccessCertificateIssuer field if non-nil, zero value otherwise.

### GetSecureAccessCertificateIssuerOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessCertificateIssuerOk() (*string, bool)`

GetSecureAccessCertificateIssuerOk returns a tuple with the SecureAccessCertificateIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessCertificateIssuer

`func (o *GatewayUpdateProducerRdp) SetSecureAccessCertificateIssuer(v string)`

SetSecureAccessCertificateIssuer sets SecureAccessCertificateIssuer field to given value.

### HasSecureAccessCertificateIssuer

`func (o *GatewayUpdateProducerRdp) HasSecureAccessCertificateIssuer() bool`

HasSecureAccessCertificateIssuer returns a boolean if a field has been set.

### GetSecureAccessDelay

`func (o *GatewayUpdateProducerRdp) GetSecureAccessDelay() int64`

GetSecureAccessDelay returns the SecureAccessDelay field if non-nil, zero value otherwise.

### GetSecureAccessDelayOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessDelayOk() (*int64, bool)`

GetSecureAccessDelayOk returns a tuple with the SecureAccessDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDelay

`func (o *GatewayUpdateProducerRdp) SetSecureAccessDelay(v int64)`

SetSecureAccessDelay sets SecureAccessDelay field to given value.

### HasSecureAccessDelay

`func (o *GatewayUpdateProducerRdp) HasSecureAccessDelay() bool`

HasSecureAccessDelay returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayUpdateProducerRdp) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayUpdateProducerRdp) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayUpdateProducerRdp) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessEnforceHostsRestriction

`func (o *GatewayUpdateProducerRdp) GetSecureAccessEnforceHostsRestriction() bool`

GetSecureAccessEnforceHostsRestriction returns the SecureAccessEnforceHostsRestriction field if non-nil, zero value otherwise.

### GetSecureAccessEnforceHostsRestrictionOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessEnforceHostsRestrictionOk() (*bool, bool)`

GetSecureAccessEnforceHostsRestrictionOk returns a tuple with the SecureAccessEnforceHostsRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnforceHostsRestriction

`func (o *GatewayUpdateProducerRdp) SetSecureAccessEnforceHostsRestriction(v bool)`

SetSecureAccessEnforceHostsRestriction sets SecureAccessEnforceHostsRestriction field to given value.

### HasSecureAccessEnforceHostsRestriction

`func (o *GatewayUpdateProducerRdp) HasSecureAccessEnforceHostsRestriction() bool`

HasSecureAccessEnforceHostsRestriction returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayUpdateProducerRdp) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayUpdateProducerRdp) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayUpdateProducerRdp) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdGatewayServer

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdGatewayServer() string`

GetSecureAccessRdGatewayServer returns the SecureAccessRdGatewayServer field if non-nil, zero value otherwise.

### GetSecureAccessRdGatewayServerOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdGatewayServerOk() (*string, bool)`

GetSecureAccessRdGatewayServerOk returns a tuple with the SecureAccessRdGatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdGatewayServer

`func (o *GatewayUpdateProducerRdp) SetSecureAccessRdGatewayServer(v string)`

SetSecureAccessRdGatewayServer sets SecureAccessRdGatewayServer field to given value.

### HasSecureAccessRdGatewayServer

`func (o *GatewayUpdateProducerRdp) HasSecureAccessRdGatewayServer() bool`

HasSecureAccessRdGatewayServer returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *GatewayUpdateProducerRdp) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *GatewayUpdateProducerRdp) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessRdpUser

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdpUser() string`

GetSecureAccessRdpUser returns the SecureAccessRdpUser field if non-nil, zero value otherwise.

### GetSecureAccessRdpUserOk

`func (o *GatewayUpdateProducerRdp) GetSecureAccessRdpUserOk() (*string, bool)`

GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpUser

`func (o *GatewayUpdateProducerRdp) SetSecureAccessRdpUser(v string)`

SetSecureAccessRdpUser sets SecureAccessRdpUser field to given value.

### HasSecureAccessRdpUser

`func (o *GatewayUpdateProducerRdp) HasSecureAccessRdpUser() bool`

HasSecureAccessRdpUser returns a boolean if a field has been set.

### GetSkipDryRun

`func (o *GatewayUpdateProducerRdp) GetSkipDryRun() string`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *GatewayUpdateProducerRdp) GetSkipDryRunOk() (*string, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *GatewayUpdateProducerRdp) SetSkipDryRun(v string)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *GatewayUpdateProducerRdp) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerRdp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerRdp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerRdp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerRdp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTarget

`func (o *GatewayUpdateProducerRdp) GetTarget() []string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GatewayUpdateProducerRdp) GetTargetOk() (*[]string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GatewayUpdateProducerRdp) SetTarget(v []string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GatewayUpdateProducerRdp) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerRdp) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerRdp) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerRdp) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerRdp) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerRdp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerRdp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerRdp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerRdp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerRdp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerRdp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerRdp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerRdp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseCapitalLetters

`func (o *GatewayUpdateProducerRdp) GetUseCapitalLetters() string`

GetUseCapitalLetters returns the UseCapitalLetters field if non-nil, zero value otherwise.

### GetUseCapitalLettersOk

`func (o *GatewayUpdateProducerRdp) GetUseCapitalLettersOk() (*string, bool)`

GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCapitalLetters

`func (o *GatewayUpdateProducerRdp) SetUseCapitalLetters(v string)`

SetUseCapitalLetters sets UseCapitalLetters field to given value.

### HasUseCapitalLetters

`func (o *GatewayUpdateProducerRdp) HasUseCapitalLetters() bool`

HasUseCapitalLetters returns a boolean if a field has been set.

### GetUseLowerLetters

`func (o *GatewayUpdateProducerRdp) GetUseLowerLetters() string`

GetUseLowerLetters returns the UseLowerLetters field if non-nil, zero value otherwise.

### GetUseLowerLettersOk

`func (o *GatewayUpdateProducerRdp) GetUseLowerLettersOk() (*string, bool)`

GetUseLowerLettersOk returns a tuple with the UseLowerLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLowerLetters

`func (o *GatewayUpdateProducerRdp) SetUseLowerLetters(v string)`

SetUseLowerLetters sets UseLowerLetters field to given value.

### HasUseLowerLetters

`func (o *GatewayUpdateProducerRdp) HasUseLowerLetters() bool`

HasUseLowerLetters returns a boolean if a field has been set.

### GetUseNumbers

`func (o *GatewayUpdateProducerRdp) GetUseNumbers() string`

GetUseNumbers returns the UseNumbers field if non-nil, zero value otherwise.

### GetUseNumbersOk

`func (o *GatewayUpdateProducerRdp) GetUseNumbersOk() (*string, bool)`

GetUseNumbersOk returns a tuple with the UseNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumbers

`func (o *GatewayUpdateProducerRdp) SetUseNumbers(v string)`

SetUseNumbers sets UseNumbers field to given value.

### HasUseNumbers

`func (o *GatewayUpdateProducerRdp) HasUseNumbers() bool`

HasUseNumbers returns a boolean if a field has been set.

### GetUseSpecialCharacters

`func (o *GatewayUpdateProducerRdp) GetUseSpecialCharacters() string`

GetUseSpecialCharacters returns the UseSpecialCharacters field if non-nil, zero value otherwise.

### GetUseSpecialCharactersOk

`func (o *GatewayUpdateProducerRdp) GetUseSpecialCharactersOk() (*string, bool)`

GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpecialCharacters

`func (o *GatewayUpdateProducerRdp) SetUseSpecialCharacters(v string)`

SetUseSpecialCharacters sets UseSpecialCharacters field to given value.

### HasUseSpecialCharacters

`func (o *GatewayUpdateProducerRdp) HasUseSpecialCharacters() bool`

HasUseSpecialCharacters returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerRdp) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerRdp) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerRdp) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerRdp) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetWarnUserBeforeExpiration

`func (o *GatewayUpdateProducerRdp) GetWarnUserBeforeExpiration() int64`

GetWarnUserBeforeExpiration returns the WarnUserBeforeExpiration field if non-nil, zero value otherwise.

### GetWarnUserBeforeExpirationOk

`func (o *GatewayUpdateProducerRdp) GetWarnUserBeforeExpirationOk() (*int64, bool)`

GetWarnUserBeforeExpirationOk returns a tuple with the WarnUserBeforeExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUserBeforeExpiration

`func (o *GatewayUpdateProducerRdp) SetWarnUserBeforeExpiration(v int64)`

SetWarnUserBeforeExpiration sets WarnUserBeforeExpiration field to given value.

### HasWarnUserBeforeExpiration

`func (o *GatewayUpdateProducerRdp) HasWarnUserBeforeExpiration() bool`

HasWarnUserBeforeExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


