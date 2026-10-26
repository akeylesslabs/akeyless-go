# RotatedSecretCreateAliCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AliCloud RAM user AccessKey secret to rotate (relevant only for rotator-type&#x3D;access-key) | [optional] 
**AccessKeyId** | Pointer to **string** | AliCloud RAM user AccessKey ID to rotate (relevant only for rotator-type&#x3D;access-key). On create with authentication-credentials&#x3D;use-target-creds it may be empty, then Akeyless creates the AccessKey on the RAM user of the target and manages it | [optional] 
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. When false, user-defined input/output rules are stored but not enforced; the base security validation still runs.  AraEnabled is tri-state (nil/true/false), not a plain bool: it self-encodes its wire value (see akl.OptionalBool) so an explicit false survives the curl-proxy relay instead of being dropped like a default-false bool flag. | [optional] 
**AuthenticationCredentials** | Pointer to **string** | The credentials to connect with use-user-creds/use-target-creds | [optional] [default to "use-user-creds"]
**AutoRotate** | Pointer to **string** |  | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EnableAgenticRuntimeAuthority** | Pointer to **bool** | EnableAra is the documented spelling of AraEnabled. Both set the same field; --ara-enabled shipped first and stays as an undocumented alias so existing scripts and the Terraform provider keep working. | [optional] 
**EnableAiQuorum** | Pointer to **bool** | Turns on AI Quorum checks for this item. | [optional] 
**GraceRotation** | Pointer to **string** | Enable graceful rotation (keep both versions temporarily). When enabled, a new secret version is created while the previous version is kept for the grace period, so both versions exist for a limited time. [true/false] | [optional] 
**GraceRotationHour** | Pointer to **int32** | The Hour of the grace rotation in UTC | [optional] 
**GraceRotationInterval** | Pointer to **string** | The number of days to wait before deleting the old key (must be bigger than rotation-interval) | [optional] 
**GraceRotationTiming** | Pointer to **string** | When to create the new version relative to the rotation date [after/before] | [optional] 
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
**RotationEventIn** | Pointer to **[]string** | How many days before the rotation of the item would you like to be notified | [optional] 
**RotationHour** | Pointer to **int32** |  | [optional] 
**RotationInterval** | Pointer to **string** |  | [optional] 
**RotatorType** | **string** | The rotator type. options: [target/access-key] | 
**SkipDryRun** | Pointer to **bool** | If set, dry-run will be skipped [true/false] | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | **string** | The target name to associate | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseCapitalLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one uppercase character from the ISO basic Latin alphabet (A to Z). [true/false] | [optional] 
**UseLowerLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one lowercase character from the ISO basic Latin alphabet (a to z). [true/false] | [optional] 
**UseNumbers** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one numeric character (0 to 9). [true/false] | [optional] 
**UseSpecialCharacters** | Pointer to **string** |  | [optional] 

## Methods

### NewRotatedSecretCreateAliCloud

`func NewRotatedSecretCreateAliCloud(name string, rotatorType string, targetName string, ) *RotatedSecretCreateAliCloud`

NewRotatedSecretCreateAliCloud instantiates a new RotatedSecretCreateAliCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretCreateAliCloudWithDefaults

`func NewRotatedSecretCreateAliCloudWithDefaults() *RotatedSecretCreateAliCloud`

NewRotatedSecretCreateAliCloudWithDefaults instantiates a new RotatedSecretCreateAliCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *RotatedSecretCreateAliCloud) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *RotatedSecretCreateAliCloud) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *RotatedSecretCreateAliCloud) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *RotatedSecretCreateAliCloud) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *RotatedSecretCreateAliCloud) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *RotatedSecretCreateAliCloud) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *RotatedSecretCreateAliCloud) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *RotatedSecretCreateAliCloud) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAraEnabled

`func (o *RotatedSecretCreateAliCloud) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *RotatedSecretCreateAliCloud) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *RotatedSecretCreateAliCloud) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *RotatedSecretCreateAliCloud) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetAuthenticationCredentials

`func (o *RotatedSecretCreateAliCloud) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *RotatedSecretCreateAliCloud) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *RotatedSecretCreateAliCloud) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *RotatedSecretCreateAliCloud) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretCreateAliCloud) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretCreateAliCloud) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretCreateAliCloud) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretCreateAliCloud) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretCreateAliCloud) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretCreateAliCloud) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretCreateAliCloud) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretCreateAliCloud) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretCreateAliCloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretCreateAliCloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretCreateAliCloud) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretCreateAliCloud) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAgenticRuntimeAuthority

`func (o *RotatedSecretCreateAliCloud) GetEnableAgenticRuntimeAuthority() bool`

GetEnableAgenticRuntimeAuthority returns the EnableAgenticRuntimeAuthority field if non-nil, zero value otherwise.

### GetEnableAgenticRuntimeAuthorityOk

`func (o *RotatedSecretCreateAliCloud) GetEnableAgenticRuntimeAuthorityOk() (*bool, bool)`

GetEnableAgenticRuntimeAuthorityOk returns a tuple with the EnableAgenticRuntimeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAgenticRuntimeAuthority

`func (o *RotatedSecretCreateAliCloud) SetEnableAgenticRuntimeAuthority(v bool)`

SetEnableAgenticRuntimeAuthority sets EnableAgenticRuntimeAuthority field to given value.

### HasEnableAgenticRuntimeAuthority

`func (o *RotatedSecretCreateAliCloud) HasEnableAgenticRuntimeAuthority() bool`

HasEnableAgenticRuntimeAuthority returns a boolean if a field has been set.

### GetEnableAiQuorum

`func (o *RotatedSecretCreateAliCloud) GetEnableAiQuorum() bool`

GetEnableAiQuorum returns the EnableAiQuorum field if non-nil, zero value otherwise.

### GetEnableAiQuorumOk

`func (o *RotatedSecretCreateAliCloud) GetEnableAiQuorumOk() (*bool, bool)`

GetEnableAiQuorumOk returns a tuple with the EnableAiQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiQuorum

`func (o *RotatedSecretCreateAliCloud) SetEnableAiQuorum(v bool)`

SetEnableAiQuorum sets EnableAiQuorum field to given value.

### HasEnableAiQuorum

`func (o *RotatedSecretCreateAliCloud) HasEnableAiQuorum() bool`

HasEnableAiQuorum returns a boolean if a field has been set.

### GetGraceRotation

`func (o *RotatedSecretCreateAliCloud) GetGraceRotation() string`

GetGraceRotation returns the GraceRotation field if non-nil, zero value otherwise.

### GetGraceRotationOk

`func (o *RotatedSecretCreateAliCloud) GetGraceRotationOk() (*string, bool)`

GetGraceRotationOk returns a tuple with the GraceRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotation

`func (o *RotatedSecretCreateAliCloud) SetGraceRotation(v string)`

SetGraceRotation sets GraceRotation field to given value.

### HasGraceRotation

`func (o *RotatedSecretCreateAliCloud) HasGraceRotation() bool`

HasGraceRotation returns a boolean if a field has been set.

### GetGraceRotationHour

`func (o *RotatedSecretCreateAliCloud) GetGraceRotationHour() int32`

GetGraceRotationHour returns the GraceRotationHour field if non-nil, zero value otherwise.

### GetGraceRotationHourOk

`func (o *RotatedSecretCreateAliCloud) GetGraceRotationHourOk() (*int32, bool)`

GetGraceRotationHourOk returns a tuple with the GraceRotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotationHour

`func (o *RotatedSecretCreateAliCloud) SetGraceRotationHour(v int32)`

SetGraceRotationHour sets GraceRotationHour field to given value.

### HasGraceRotationHour

`func (o *RotatedSecretCreateAliCloud) HasGraceRotationHour() bool`

HasGraceRotationHour returns a boolean if a field has been set.

### GetGraceRotationInterval

`func (o *RotatedSecretCreateAliCloud) GetGraceRotationInterval() string`

GetGraceRotationInterval returns the GraceRotationInterval field if non-nil, zero value otherwise.

### GetGraceRotationIntervalOk

`func (o *RotatedSecretCreateAliCloud) GetGraceRotationIntervalOk() (*string, bool)`

GetGraceRotationIntervalOk returns a tuple with the GraceRotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotationInterval

`func (o *RotatedSecretCreateAliCloud) SetGraceRotationInterval(v string)`

SetGraceRotationInterval sets GraceRotationInterval field to given value.

### HasGraceRotationInterval

`func (o *RotatedSecretCreateAliCloud) HasGraceRotationInterval() bool`

HasGraceRotationInterval returns a boolean if a field has been set.

### GetGraceRotationTiming

`func (o *RotatedSecretCreateAliCloud) GetGraceRotationTiming() string`

GetGraceRotationTiming returns the GraceRotationTiming field if non-nil, zero value otherwise.

### GetGraceRotationTimingOk

`func (o *RotatedSecretCreateAliCloud) GetGraceRotationTimingOk() (*string, bool)`

GetGraceRotationTimingOk returns a tuple with the GraceRotationTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotationTiming

`func (o *RotatedSecretCreateAliCloud) SetGraceRotationTiming(v string)`

SetGraceRotationTiming sets GraceRotationTiming field to given value.

### HasGraceRotationTiming

`func (o *RotatedSecretCreateAliCloud) HasGraceRotationTiming() bool`

HasGraceRotationTiming returns a boolean if a field has been set.

### GetInputRule

`func (o *RotatedSecretCreateAliCloud) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *RotatedSecretCreateAliCloud) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *RotatedSecretCreateAliCloud) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *RotatedSecretCreateAliCloud) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *RotatedSecretCreateAliCloud) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *RotatedSecretCreateAliCloud) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *RotatedSecretCreateAliCloud) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *RotatedSecretCreateAliCloud) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretCreateAliCloud) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretCreateAliCloud) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretCreateAliCloud) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretCreateAliCloud) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretCreateAliCloud) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretCreateAliCloud) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretCreateAliCloud) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretCreateAliCloud) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLockOnRead

`func (o *RotatedSecretCreateAliCloud) GetLockOnRead() string`

GetLockOnRead returns the LockOnRead field if non-nil, zero value otherwise.

### GetLockOnReadOk

`func (o *RotatedSecretCreateAliCloud) GetLockOnReadOk() (*string, bool)`

GetLockOnReadOk returns a tuple with the LockOnRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOnRead

`func (o *RotatedSecretCreateAliCloud) SetLockOnRead(v string)`

SetLockOnRead sets LockOnRead field to given value.

### HasLockOnRead

`func (o *RotatedSecretCreateAliCloud) HasLockOnRead() bool`

HasLockOnRead returns a boolean if a field has been set.

### GetLockTtl

`func (o *RotatedSecretCreateAliCloud) GetLockTtl() string`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *RotatedSecretCreateAliCloud) GetLockTtlOk() (*string, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *RotatedSecretCreateAliCloud) SetLockTtl(v string)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *RotatedSecretCreateAliCloud) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *RotatedSecretCreateAliCloud) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *RotatedSecretCreateAliCloud) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *RotatedSecretCreateAliCloud) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *RotatedSecretCreateAliCloud) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretCreateAliCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretCreateAliCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretCreateAliCloud) SetName(v string)`

SetName sets Name field to given value.


### GetOutputRule

`func (o *RotatedSecretCreateAliCloud) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *RotatedSecretCreateAliCloud) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *RotatedSecretCreateAliCloud) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *RotatedSecretCreateAliCloud) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetPasswordLength

`func (o *RotatedSecretCreateAliCloud) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretCreateAliCloud) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretCreateAliCloud) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretCreateAliCloud) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRotateOnUnlock

`func (o *RotatedSecretCreateAliCloud) GetRotateOnUnlock() string`

GetRotateOnUnlock returns the RotateOnUnlock field if non-nil, zero value otherwise.

### GetRotateOnUnlockOk

`func (o *RotatedSecretCreateAliCloud) GetRotateOnUnlockOk() (*string, bool)`

GetRotateOnUnlockOk returns a tuple with the RotateOnUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateOnUnlock

`func (o *RotatedSecretCreateAliCloud) SetRotateOnUnlock(v string)`

SetRotateOnUnlock sets RotateOnUnlock field to given value.

### HasRotateOnUnlock

`func (o *RotatedSecretCreateAliCloud) HasRotateOnUnlock() bool`

HasRotateOnUnlock returns a boolean if a field has been set.

### GetRotationEventIn

`func (o *RotatedSecretCreateAliCloud) GetRotationEventIn() []string`

GetRotationEventIn returns the RotationEventIn field if non-nil, zero value otherwise.

### GetRotationEventInOk

`func (o *RotatedSecretCreateAliCloud) GetRotationEventInOk() (*[]string, bool)`

GetRotationEventInOk returns a tuple with the RotationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationEventIn

`func (o *RotatedSecretCreateAliCloud) SetRotationEventIn(v []string)`

SetRotationEventIn sets RotationEventIn field to given value.

### HasRotationEventIn

`func (o *RotatedSecretCreateAliCloud) HasRotationEventIn() bool`

HasRotationEventIn returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretCreateAliCloud) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretCreateAliCloud) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretCreateAliCloud) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretCreateAliCloud) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretCreateAliCloud) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretCreateAliCloud) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretCreateAliCloud) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretCreateAliCloud) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotatorType

`func (o *RotatedSecretCreateAliCloud) GetRotatorType() string`

GetRotatorType returns the RotatorType field if non-nil, zero value otherwise.

### GetRotatorTypeOk

`func (o *RotatedSecretCreateAliCloud) GetRotatorTypeOk() (*string, bool)`

GetRotatorTypeOk returns a tuple with the RotatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorType

`func (o *RotatedSecretCreateAliCloud) SetRotatorType(v string)`

SetRotatorType sets RotatorType field to given value.


### GetSkipDryRun

`func (o *RotatedSecretCreateAliCloud) GetSkipDryRun() bool`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *RotatedSecretCreateAliCloud) GetSkipDryRunOk() (*bool, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *RotatedSecretCreateAliCloud) SetSkipDryRun(v bool)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *RotatedSecretCreateAliCloud) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetTags

`func (o *RotatedSecretCreateAliCloud) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RotatedSecretCreateAliCloud) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RotatedSecretCreateAliCloud) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RotatedSecretCreateAliCloud) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *RotatedSecretCreateAliCloud) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *RotatedSecretCreateAliCloud) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *RotatedSecretCreateAliCloud) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetToken

`func (o *RotatedSecretCreateAliCloud) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretCreateAliCloud) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretCreateAliCloud) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretCreateAliCloud) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretCreateAliCloud) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretCreateAliCloud) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretCreateAliCloud) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretCreateAliCloud) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseCapitalLetters

`func (o *RotatedSecretCreateAliCloud) GetUseCapitalLetters() string`

GetUseCapitalLetters returns the UseCapitalLetters field if non-nil, zero value otherwise.

### GetUseCapitalLettersOk

`func (o *RotatedSecretCreateAliCloud) GetUseCapitalLettersOk() (*string, bool)`

GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCapitalLetters

`func (o *RotatedSecretCreateAliCloud) SetUseCapitalLetters(v string)`

SetUseCapitalLetters sets UseCapitalLetters field to given value.

### HasUseCapitalLetters

`func (o *RotatedSecretCreateAliCloud) HasUseCapitalLetters() bool`

HasUseCapitalLetters returns a boolean if a field has been set.

### GetUseLowerLetters

`func (o *RotatedSecretCreateAliCloud) GetUseLowerLetters() string`

GetUseLowerLetters returns the UseLowerLetters field if non-nil, zero value otherwise.

### GetUseLowerLettersOk

`func (o *RotatedSecretCreateAliCloud) GetUseLowerLettersOk() (*string, bool)`

GetUseLowerLettersOk returns a tuple with the UseLowerLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLowerLetters

`func (o *RotatedSecretCreateAliCloud) SetUseLowerLetters(v string)`

SetUseLowerLetters sets UseLowerLetters field to given value.

### HasUseLowerLetters

`func (o *RotatedSecretCreateAliCloud) HasUseLowerLetters() bool`

HasUseLowerLetters returns a boolean if a field has been set.

### GetUseNumbers

`func (o *RotatedSecretCreateAliCloud) GetUseNumbers() string`

GetUseNumbers returns the UseNumbers field if non-nil, zero value otherwise.

### GetUseNumbersOk

`func (o *RotatedSecretCreateAliCloud) GetUseNumbersOk() (*string, bool)`

GetUseNumbersOk returns a tuple with the UseNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumbers

`func (o *RotatedSecretCreateAliCloud) SetUseNumbers(v string)`

SetUseNumbers sets UseNumbers field to given value.

### HasUseNumbers

`func (o *RotatedSecretCreateAliCloud) HasUseNumbers() bool`

HasUseNumbers returns a boolean if a field has been set.

### GetUseSpecialCharacters

`func (o *RotatedSecretCreateAliCloud) GetUseSpecialCharacters() string`

GetUseSpecialCharacters returns the UseSpecialCharacters field if non-nil, zero value otherwise.

### GetUseSpecialCharactersOk

`func (o *RotatedSecretCreateAliCloud) GetUseSpecialCharactersOk() (*string, bool)`

GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpecialCharacters

`func (o *RotatedSecretCreateAliCloud) SetUseSpecialCharacters(v string)`

SetUseSpecialCharacters sets UseSpecialCharacters field to given value.

### HasUseSpecialCharacters

`func (o *RotatedSecretCreateAliCloud) HasUseSpecialCharacters() bool`

HasUseSpecialCharacters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


