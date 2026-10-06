# RotatedSecretCreateOkta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. When false, user-defined input/output rules are stored but not enforced; the base security validation still runs.  AraEnabled is tri-state (nil/true/false), not a plain bool: it self-encodes its wire value (see akl.OptionalBool) so an explicit false survives the curl-proxy relay instead of being dropped like a default-false bool flag. | [optional] 
**AuthenticationCredentials** | Pointer to **string** | Okta uses the SSWS token from the target. | [optional] 
**AutoRotate** | Pointer to **string** |  | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object. | [optional] 
**EnableAgenticRuntimeAuthority** | Pointer to **bool** | EnableAra is the documented spelling of AraEnabled. Both set the same field; --ara-enabled shipped first and stays as an undocumented alias so existing scripts and the Terraform provider keep working. | [optional] 
**EnableAiQuorum** | Pointer to **bool** | Turns on AI Quorum checks for this item. | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** |  | [optional] 
**LockOnRead** | Pointer to **string** | Lock this secret after each successful value read | [optional] 
**LockTtl** | Pointer to **string** | Lock TTL in minutes | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Rotated secret name | 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**RotateOnUnlock** | Pointer to **string** | Rotate this secret after it is unlocked | [optional] 
**RotatedPassword** | Pointer to **string** | rotated-username password (relevant only for rotator-type&#x3D;password) | [optional] 
**RotatedUsername** | Pointer to **string** | username to be rotated, if selected use-self-creds at rotator-creds-type, this username will try to rotate it&#39;s own password, if use-target-creds is selected, target credentials will be use to rotate the rotated-password (relevant only for rotator-type&#x3D;password) | [optional] 
**RotationEventIn** | Pointer to **[]string** | How many days before the rotation of the item would you like to be notified | [optional] 
**RotationHour** | Pointer to **int32** |  | [optional] 
**RotationInterval** | Pointer to **string** |  | [optional] 
**RotatorType** | Pointer to **string** | The Okta rotator supports only password. | [optional] 
**SkipDryRun** | Pointer to **bool** | If set, dry-run will be skipped [true/false] | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | **string** | The Okta target name to associate. | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseCapitalLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one uppercase character from the ISO basic Latin alphabet (A to Z). [true/false] | [optional] 
**UseLowerLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one lowercase character from the ISO basic Latin alphabet (a to z). [true/false] | [optional] 
**UseNumbers** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one numeric character (0 to 9). [true/false] | [optional] 
**UseSpecialCharacters** | Pointer to **string** |  | [optional] 

## Methods

### NewRotatedSecretCreateOkta

`func NewRotatedSecretCreateOkta(name string, targetName string, ) *RotatedSecretCreateOkta`

NewRotatedSecretCreateOkta instantiates a new RotatedSecretCreateOkta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretCreateOktaWithDefaults

`func NewRotatedSecretCreateOktaWithDefaults() *RotatedSecretCreateOkta`

NewRotatedSecretCreateOktaWithDefaults instantiates a new RotatedSecretCreateOkta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAraEnabled

`func (o *RotatedSecretCreateOkta) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *RotatedSecretCreateOkta) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *RotatedSecretCreateOkta) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *RotatedSecretCreateOkta) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetAuthenticationCredentials

`func (o *RotatedSecretCreateOkta) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *RotatedSecretCreateOkta) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *RotatedSecretCreateOkta) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *RotatedSecretCreateOkta) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretCreateOkta) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretCreateOkta) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretCreateOkta) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretCreateOkta) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretCreateOkta) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretCreateOkta) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretCreateOkta) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretCreateOkta) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretCreateOkta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretCreateOkta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretCreateOkta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretCreateOkta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAgenticRuntimeAuthority

`func (o *RotatedSecretCreateOkta) GetEnableAgenticRuntimeAuthority() bool`

GetEnableAgenticRuntimeAuthority returns the EnableAgenticRuntimeAuthority field if non-nil, zero value otherwise.

### GetEnableAgenticRuntimeAuthorityOk

`func (o *RotatedSecretCreateOkta) GetEnableAgenticRuntimeAuthorityOk() (*bool, bool)`

GetEnableAgenticRuntimeAuthorityOk returns a tuple with the EnableAgenticRuntimeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAgenticRuntimeAuthority

`func (o *RotatedSecretCreateOkta) SetEnableAgenticRuntimeAuthority(v bool)`

SetEnableAgenticRuntimeAuthority sets EnableAgenticRuntimeAuthority field to given value.

### HasEnableAgenticRuntimeAuthority

`func (o *RotatedSecretCreateOkta) HasEnableAgenticRuntimeAuthority() bool`

HasEnableAgenticRuntimeAuthority returns a boolean if a field has been set.

### GetEnableAiQuorum

`func (o *RotatedSecretCreateOkta) GetEnableAiQuorum() bool`

GetEnableAiQuorum returns the EnableAiQuorum field if non-nil, zero value otherwise.

### GetEnableAiQuorumOk

`func (o *RotatedSecretCreateOkta) GetEnableAiQuorumOk() (*bool, bool)`

GetEnableAiQuorumOk returns a tuple with the EnableAiQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiQuorum

`func (o *RotatedSecretCreateOkta) SetEnableAiQuorum(v bool)`

SetEnableAiQuorum sets EnableAiQuorum field to given value.

### HasEnableAiQuorum

`func (o *RotatedSecretCreateOkta) HasEnableAiQuorum() bool`

HasEnableAiQuorum returns a boolean if a field has been set.

### GetInputRule

`func (o *RotatedSecretCreateOkta) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *RotatedSecretCreateOkta) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *RotatedSecretCreateOkta) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *RotatedSecretCreateOkta) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *RotatedSecretCreateOkta) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *RotatedSecretCreateOkta) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *RotatedSecretCreateOkta) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *RotatedSecretCreateOkta) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretCreateOkta) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretCreateOkta) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretCreateOkta) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretCreateOkta) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretCreateOkta) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretCreateOkta) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretCreateOkta) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretCreateOkta) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLockOnRead

`func (o *RotatedSecretCreateOkta) GetLockOnRead() string`

GetLockOnRead returns the LockOnRead field if non-nil, zero value otherwise.

### GetLockOnReadOk

`func (o *RotatedSecretCreateOkta) GetLockOnReadOk() (*string, bool)`

GetLockOnReadOk returns a tuple with the LockOnRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOnRead

`func (o *RotatedSecretCreateOkta) SetLockOnRead(v string)`

SetLockOnRead sets LockOnRead field to given value.

### HasLockOnRead

`func (o *RotatedSecretCreateOkta) HasLockOnRead() bool`

HasLockOnRead returns a boolean if a field has been set.

### GetLockTtl

`func (o *RotatedSecretCreateOkta) GetLockTtl() string`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *RotatedSecretCreateOkta) GetLockTtlOk() (*string, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *RotatedSecretCreateOkta) SetLockTtl(v string)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *RotatedSecretCreateOkta) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *RotatedSecretCreateOkta) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *RotatedSecretCreateOkta) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *RotatedSecretCreateOkta) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *RotatedSecretCreateOkta) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretCreateOkta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretCreateOkta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretCreateOkta) SetName(v string)`

SetName sets Name field to given value.


### GetOutputRule

`func (o *RotatedSecretCreateOkta) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *RotatedSecretCreateOkta) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *RotatedSecretCreateOkta) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *RotatedSecretCreateOkta) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetPasswordLength

`func (o *RotatedSecretCreateOkta) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretCreateOkta) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretCreateOkta) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretCreateOkta) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRotateOnUnlock

`func (o *RotatedSecretCreateOkta) GetRotateOnUnlock() string`

GetRotateOnUnlock returns the RotateOnUnlock field if non-nil, zero value otherwise.

### GetRotateOnUnlockOk

`func (o *RotatedSecretCreateOkta) GetRotateOnUnlockOk() (*string, bool)`

GetRotateOnUnlockOk returns a tuple with the RotateOnUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateOnUnlock

`func (o *RotatedSecretCreateOkta) SetRotateOnUnlock(v string)`

SetRotateOnUnlock sets RotateOnUnlock field to given value.

### HasRotateOnUnlock

`func (o *RotatedSecretCreateOkta) HasRotateOnUnlock() bool`

HasRotateOnUnlock returns a boolean if a field has been set.

### GetRotatedPassword

`func (o *RotatedSecretCreateOkta) GetRotatedPassword() string`

GetRotatedPassword returns the RotatedPassword field if non-nil, zero value otherwise.

### GetRotatedPasswordOk

`func (o *RotatedSecretCreateOkta) GetRotatedPasswordOk() (*string, bool)`

GetRotatedPasswordOk returns a tuple with the RotatedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedPassword

`func (o *RotatedSecretCreateOkta) SetRotatedPassword(v string)`

SetRotatedPassword sets RotatedPassword field to given value.

### HasRotatedPassword

`func (o *RotatedSecretCreateOkta) HasRotatedPassword() bool`

HasRotatedPassword returns a boolean if a field has been set.

### GetRotatedUsername

`func (o *RotatedSecretCreateOkta) GetRotatedUsername() string`

GetRotatedUsername returns the RotatedUsername field if non-nil, zero value otherwise.

### GetRotatedUsernameOk

`func (o *RotatedSecretCreateOkta) GetRotatedUsernameOk() (*string, bool)`

GetRotatedUsernameOk returns a tuple with the RotatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedUsername

`func (o *RotatedSecretCreateOkta) SetRotatedUsername(v string)`

SetRotatedUsername sets RotatedUsername field to given value.

### HasRotatedUsername

`func (o *RotatedSecretCreateOkta) HasRotatedUsername() bool`

HasRotatedUsername returns a boolean if a field has been set.

### GetRotationEventIn

`func (o *RotatedSecretCreateOkta) GetRotationEventIn() []string`

GetRotationEventIn returns the RotationEventIn field if non-nil, zero value otherwise.

### GetRotationEventInOk

`func (o *RotatedSecretCreateOkta) GetRotationEventInOk() (*[]string, bool)`

GetRotationEventInOk returns a tuple with the RotationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationEventIn

`func (o *RotatedSecretCreateOkta) SetRotationEventIn(v []string)`

SetRotationEventIn sets RotationEventIn field to given value.

### HasRotationEventIn

`func (o *RotatedSecretCreateOkta) HasRotationEventIn() bool`

HasRotationEventIn returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretCreateOkta) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretCreateOkta) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretCreateOkta) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretCreateOkta) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretCreateOkta) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretCreateOkta) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretCreateOkta) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretCreateOkta) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotatorType

`func (o *RotatedSecretCreateOkta) GetRotatorType() string`

GetRotatorType returns the RotatorType field if non-nil, zero value otherwise.

### GetRotatorTypeOk

`func (o *RotatedSecretCreateOkta) GetRotatorTypeOk() (*string, bool)`

GetRotatorTypeOk returns a tuple with the RotatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorType

`func (o *RotatedSecretCreateOkta) SetRotatorType(v string)`

SetRotatorType sets RotatorType field to given value.

### HasRotatorType

`func (o *RotatedSecretCreateOkta) HasRotatorType() bool`

HasRotatorType returns a boolean if a field has been set.

### GetSkipDryRun

`func (o *RotatedSecretCreateOkta) GetSkipDryRun() bool`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *RotatedSecretCreateOkta) GetSkipDryRunOk() (*bool, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *RotatedSecretCreateOkta) SetSkipDryRun(v bool)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *RotatedSecretCreateOkta) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetTags

`func (o *RotatedSecretCreateOkta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RotatedSecretCreateOkta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RotatedSecretCreateOkta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RotatedSecretCreateOkta) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *RotatedSecretCreateOkta) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *RotatedSecretCreateOkta) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *RotatedSecretCreateOkta) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetToken

`func (o *RotatedSecretCreateOkta) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretCreateOkta) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretCreateOkta) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretCreateOkta) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretCreateOkta) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretCreateOkta) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretCreateOkta) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretCreateOkta) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseCapitalLetters

`func (o *RotatedSecretCreateOkta) GetUseCapitalLetters() string`

GetUseCapitalLetters returns the UseCapitalLetters field if non-nil, zero value otherwise.

### GetUseCapitalLettersOk

`func (o *RotatedSecretCreateOkta) GetUseCapitalLettersOk() (*string, bool)`

GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCapitalLetters

`func (o *RotatedSecretCreateOkta) SetUseCapitalLetters(v string)`

SetUseCapitalLetters sets UseCapitalLetters field to given value.

### HasUseCapitalLetters

`func (o *RotatedSecretCreateOkta) HasUseCapitalLetters() bool`

HasUseCapitalLetters returns a boolean if a field has been set.

### GetUseLowerLetters

`func (o *RotatedSecretCreateOkta) GetUseLowerLetters() string`

GetUseLowerLetters returns the UseLowerLetters field if non-nil, zero value otherwise.

### GetUseLowerLettersOk

`func (o *RotatedSecretCreateOkta) GetUseLowerLettersOk() (*string, bool)`

GetUseLowerLettersOk returns a tuple with the UseLowerLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLowerLetters

`func (o *RotatedSecretCreateOkta) SetUseLowerLetters(v string)`

SetUseLowerLetters sets UseLowerLetters field to given value.

### HasUseLowerLetters

`func (o *RotatedSecretCreateOkta) HasUseLowerLetters() bool`

HasUseLowerLetters returns a boolean if a field has been set.

### GetUseNumbers

`func (o *RotatedSecretCreateOkta) GetUseNumbers() string`

GetUseNumbers returns the UseNumbers field if non-nil, zero value otherwise.

### GetUseNumbersOk

`func (o *RotatedSecretCreateOkta) GetUseNumbersOk() (*string, bool)`

GetUseNumbersOk returns a tuple with the UseNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumbers

`func (o *RotatedSecretCreateOkta) SetUseNumbers(v string)`

SetUseNumbers sets UseNumbers field to given value.

### HasUseNumbers

`func (o *RotatedSecretCreateOkta) HasUseNumbers() bool`

HasUseNumbers returns a boolean if a field has been set.

### GetUseSpecialCharacters

`func (o *RotatedSecretCreateOkta) GetUseSpecialCharacters() string`

GetUseSpecialCharacters returns the UseSpecialCharacters field if non-nil, zero value otherwise.

### GetUseSpecialCharactersOk

`func (o *RotatedSecretCreateOkta) GetUseSpecialCharactersOk() (*string, bool)`

GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpecialCharacters

`func (o *RotatedSecretCreateOkta) SetUseSpecialCharacters(v string)`

SetUseSpecialCharacters sets UseSpecialCharacters field to given value.

### HasUseSpecialCharacters

`func (o *RotatedSecretCreateOkta) HasUseSpecialCharacters() bool`

HasUseSpecialCharacters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


