# RotatedSecretUpdateAliCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AliCloud RAM user AccessKey secret to rotate (relevant only for rotator-type&#x3D;access-key) | [optional] 
**AccessKeyId** | Pointer to **string** | AliCloud RAM user AccessKey ID to rotate (relevant only for rotator-type&#x3D;access-key). On create with authentication-credentials&#x3D;use-target-creds it may be empty, then Akeyless creates the AccessKey on the RAM user of the target and manages it | [optional] 
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. When false, user-defined input/output rules are stored but not enforced; the base security validation still runs.  AraEnabled is tri-state (nil/true/false), not a plain bool: it self-encodes its wire value (see akl.OptionalBool) so an explicit false survives the curl-proxy relay instead of being dropped like a default-false bool flag. | [optional] 
**AuthenticationCredentials** | Pointer to **string** | The credentials to connect with use-user-creds/use-target-creds | [optional] [default to "use-user-creds"]
**AutoRotate** | Pointer to **string** |  | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_metadata"]
**EnableAgenticRuntimeAuthority** | Pointer to **bool** | EnableAra is the documented spelling of AraEnabled. Both set the same field; --ara-enabled shipped first and stays as an undocumented alias so existing scripts and the Terraform provider keep working. | [optional] 
**EnableAiQuorum** | Pointer to **bool** | Turns on AI Quorum checks for this item. | [optional] 
**GraceRotation** | Pointer to **string** | Enable graceful rotation (keep both versions temporarily). When enabled, a new secret version is created while the previous version is kept for the grace period, so both versions exist for a limited time. [true/false] | [optional] 
**GraceRotationHour** | Pointer to **int32** | The Hour of the grace rotation in UTC | [optional] 
**GraceRotationInterval** | Pointer to **string** | The number of days to wait before deleting the old key (must be bigger than rotation-interval) | [optional] 
**GraceRotationTiming** | Pointer to **string** | When to create the new version relative to the rotation date [after/before] | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**LockOnRead** | Pointer to **string** | Lock this secret after each successful value read | [optional] 
**LockTtl** | Pointer to **string** | Lock TTL in minutes | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Rotated secret name | 
**NewName** | Pointer to **string** | New item name | [optional] 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**RotateOnUnlock** | Pointer to **string** | Rotate this secret after it is unlocked | [optional] 
**RotationEventIn** | Pointer to **[]string** | How many days before the rotation of the item would you like to be notified | [optional] 
**RotationHour** | Pointer to **int32** |  | [optional] 
**RotationInterval** | Pointer to **string** |  | [optional] 
**SkipDryRun** | Pointer to **bool** | If set, dry-run will be skipped [true/false] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseCapitalLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one uppercase character from the ISO basic Latin alphabet (A to Z). [true/false] | [optional] 
**UseLowerLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one lowercase character from the ISO basic Latin alphabet (a to z). [true/false] | [optional] 
**UseNumbers** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one numeric character (0 to 9). [true/false] | [optional] 
**UseSpecialCharacters** | Pointer to **string** |  | [optional] 

## Methods

### NewRotatedSecretUpdateAliCloud

`func NewRotatedSecretUpdateAliCloud(name string, ) *RotatedSecretUpdateAliCloud`

NewRotatedSecretUpdateAliCloud instantiates a new RotatedSecretUpdateAliCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretUpdateAliCloudWithDefaults

`func NewRotatedSecretUpdateAliCloudWithDefaults() *RotatedSecretUpdateAliCloud`

NewRotatedSecretUpdateAliCloudWithDefaults instantiates a new RotatedSecretUpdateAliCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *RotatedSecretUpdateAliCloud) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *RotatedSecretUpdateAliCloud) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *RotatedSecretUpdateAliCloud) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *RotatedSecretUpdateAliCloud) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *RotatedSecretUpdateAliCloud) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *RotatedSecretUpdateAliCloud) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *RotatedSecretUpdateAliCloud) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *RotatedSecretUpdateAliCloud) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAddTag

`func (o *RotatedSecretUpdateAliCloud) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *RotatedSecretUpdateAliCloud) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *RotatedSecretUpdateAliCloud) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *RotatedSecretUpdateAliCloud) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAraEnabled

`func (o *RotatedSecretUpdateAliCloud) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *RotatedSecretUpdateAliCloud) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *RotatedSecretUpdateAliCloud) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *RotatedSecretUpdateAliCloud) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetAuthenticationCredentials

`func (o *RotatedSecretUpdateAliCloud) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *RotatedSecretUpdateAliCloud) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *RotatedSecretUpdateAliCloud) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *RotatedSecretUpdateAliCloud) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretUpdateAliCloud) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretUpdateAliCloud) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretUpdateAliCloud) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretUpdateAliCloud) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretUpdateAliCloud) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretUpdateAliCloud) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretUpdateAliCloud) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretUpdateAliCloud) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretUpdateAliCloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretUpdateAliCloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretUpdateAliCloud) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretUpdateAliCloud) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAgenticRuntimeAuthority

`func (o *RotatedSecretUpdateAliCloud) GetEnableAgenticRuntimeAuthority() bool`

GetEnableAgenticRuntimeAuthority returns the EnableAgenticRuntimeAuthority field if non-nil, zero value otherwise.

### GetEnableAgenticRuntimeAuthorityOk

`func (o *RotatedSecretUpdateAliCloud) GetEnableAgenticRuntimeAuthorityOk() (*bool, bool)`

GetEnableAgenticRuntimeAuthorityOk returns a tuple with the EnableAgenticRuntimeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAgenticRuntimeAuthority

`func (o *RotatedSecretUpdateAliCloud) SetEnableAgenticRuntimeAuthority(v bool)`

SetEnableAgenticRuntimeAuthority sets EnableAgenticRuntimeAuthority field to given value.

### HasEnableAgenticRuntimeAuthority

`func (o *RotatedSecretUpdateAliCloud) HasEnableAgenticRuntimeAuthority() bool`

HasEnableAgenticRuntimeAuthority returns a boolean if a field has been set.

### GetEnableAiQuorum

`func (o *RotatedSecretUpdateAliCloud) GetEnableAiQuorum() bool`

GetEnableAiQuorum returns the EnableAiQuorum field if non-nil, zero value otherwise.

### GetEnableAiQuorumOk

`func (o *RotatedSecretUpdateAliCloud) GetEnableAiQuorumOk() (*bool, bool)`

GetEnableAiQuorumOk returns a tuple with the EnableAiQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiQuorum

`func (o *RotatedSecretUpdateAliCloud) SetEnableAiQuorum(v bool)`

SetEnableAiQuorum sets EnableAiQuorum field to given value.

### HasEnableAiQuorum

`func (o *RotatedSecretUpdateAliCloud) HasEnableAiQuorum() bool`

HasEnableAiQuorum returns a boolean if a field has been set.

### GetGraceRotation

`func (o *RotatedSecretUpdateAliCloud) GetGraceRotation() string`

GetGraceRotation returns the GraceRotation field if non-nil, zero value otherwise.

### GetGraceRotationOk

`func (o *RotatedSecretUpdateAliCloud) GetGraceRotationOk() (*string, bool)`

GetGraceRotationOk returns a tuple with the GraceRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotation

`func (o *RotatedSecretUpdateAliCloud) SetGraceRotation(v string)`

SetGraceRotation sets GraceRotation field to given value.

### HasGraceRotation

`func (o *RotatedSecretUpdateAliCloud) HasGraceRotation() bool`

HasGraceRotation returns a boolean if a field has been set.

### GetGraceRotationHour

`func (o *RotatedSecretUpdateAliCloud) GetGraceRotationHour() int32`

GetGraceRotationHour returns the GraceRotationHour field if non-nil, zero value otherwise.

### GetGraceRotationHourOk

`func (o *RotatedSecretUpdateAliCloud) GetGraceRotationHourOk() (*int32, bool)`

GetGraceRotationHourOk returns a tuple with the GraceRotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotationHour

`func (o *RotatedSecretUpdateAliCloud) SetGraceRotationHour(v int32)`

SetGraceRotationHour sets GraceRotationHour field to given value.

### HasGraceRotationHour

`func (o *RotatedSecretUpdateAliCloud) HasGraceRotationHour() bool`

HasGraceRotationHour returns a boolean if a field has been set.

### GetGraceRotationInterval

`func (o *RotatedSecretUpdateAliCloud) GetGraceRotationInterval() string`

GetGraceRotationInterval returns the GraceRotationInterval field if non-nil, zero value otherwise.

### GetGraceRotationIntervalOk

`func (o *RotatedSecretUpdateAliCloud) GetGraceRotationIntervalOk() (*string, bool)`

GetGraceRotationIntervalOk returns a tuple with the GraceRotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotationInterval

`func (o *RotatedSecretUpdateAliCloud) SetGraceRotationInterval(v string)`

SetGraceRotationInterval sets GraceRotationInterval field to given value.

### HasGraceRotationInterval

`func (o *RotatedSecretUpdateAliCloud) HasGraceRotationInterval() bool`

HasGraceRotationInterval returns a boolean if a field has been set.

### GetGraceRotationTiming

`func (o *RotatedSecretUpdateAliCloud) GetGraceRotationTiming() string`

GetGraceRotationTiming returns the GraceRotationTiming field if non-nil, zero value otherwise.

### GetGraceRotationTimingOk

`func (o *RotatedSecretUpdateAliCloud) GetGraceRotationTimingOk() (*string, bool)`

GetGraceRotationTimingOk returns a tuple with the GraceRotationTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotationTiming

`func (o *RotatedSecretUpdateAliCloud) SetGraceRotationTiming(v string)`

SetGraceRotationTiming sets GraceRotationTiming field to given value.

### HasGraceRotationTiming

`func (o *RotatedSecretUpdateAliCloud) HasGraceRotationTiming() bool`

HasGraceRotationTiming returns a boolean if a field has been set.

### GetInputRule

`func (o *RotatedSecretUpdateAliCloud) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *RotatedSecretUpdateAliCloud) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *RotatedSecretUpdateAliCloud) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *RotatedSecretUpdateAliCloud) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *RotatedSecretUpdateAliCloud) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *RotatedSecretUpdateAliCloud) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *RotatedSecretUpdateAliCloud) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *RotatedSecretUpdateAliCloud) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretUpdateAliCloud) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretUpdateAliCloud) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretUpdateAliCloud) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretUpdateAliCloud) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *RotatedSecretUpdateAliCloud) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *RotatedSecretUpdateAliCloud) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *RotatedSecretUpdateAliCloud) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *RotatedSecretUpdateAliCloud) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretUpdateAliCloud) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretUpdateAliCloud) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretUpdateAliCloud) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretUpdateAliCloud) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLockOnRead

`func (o *RotatedSecretUpdateAliCloud) GetLockOnRead() string`

GetLockOnRead returns the LockOnRead field if non-nil, zero value otherwise.

### GetLockOnReadOk

`func (o *RotatedSecretUpdateAliCloud) GetLockOnReadOk() (*string, bool)`

GetLockOnReadOk returns a tuple with the LockOnRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOnRead

`func (o *RotatedSecretUpdateAliCloud) SetLockOnRead(v string)`

SetLockOnRead sets LockOnRead field to given value.

### HasLockOnRead

`func (o *RotatedSecretUpdateAliCloud) HasLockOnRead() bool`

HasLockOnRead returns a boolean if a field has been set.

### GetLockTtl

`func (o *RotatedSecretUpdateAliCloud) GetLockTtl() string`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *RotatedSecretUpdateAliCloud) GetLockTtlOk() (*string, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *RotatedSecretUpdateAliCloud) SetLockTtl(v string)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *RotatedSecretUpdateAliCloud) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *RotatedSecretUpdateAliCloud) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *RotatedSecretUpdateAliCloud) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *RotatedSecretUpdateAliCloud) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *RotatedSecretUpdateAliCloud) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretUpdateAliCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretUpdateAliCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretUpdateAliCloud) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *RotatedSecretUpdateAliCloud) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *RotatedSecretUpdateAliCloud) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *RotatedSecretUpdateAliCloud) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *RotatedSecretUpdateAliCloud) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetOutputRule

`func (o *RotatedSecretUpdateAliCloud) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *RotatedSecretUpdateAliCloud) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *RotatedSecretUpdateAliCloud) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *RotatedSecretUpdateAliCloud) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetPasswordLength

`func (o *RotatedSecretUpdateAliCloud) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretUpdateAliCloud) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretUpdateAliCloud) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretUpdateAliCloud) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRmTag

`func (o *RotatedSecretUpdateAliCloud) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *RotatedSecretUpdateAliCloud) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *RotatedSecretUpdateAliCloud) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *RotatedSecretUpdateAliCloud) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotateOnUnlock

`func (o *RotatedSecretUpdateAliCloud) GetRotateOnUnlock() string`

GetRotateOnUnlock returns the RotateOnUnlock field if non-nil, zero value otherwise.

### GetRotateOnUnlockOk

`func (o *RotatedSecretUpdateAliCloud) GetRotateOnUnlockOk() (*string, bool)`

GetRotateOnUnlockOk returns a tuple with the RotateOnUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateOnUnlock

`func (o *RotatedSecretUpdateAliCloud) SetRotateOnUnlock(v string)`

SetRotateOnUnlock sets RotateOnUnlock field to given value.

### HasRotateOnUnlock

`func (o *RotatedSecretUpdateAliCloud) HasRotateOnUnlock() bool`

HasRotateOnUnlock returns a boolean if a field has been set.

### GetRotationEventIn

`func (o *RotatedSecretUpdateAliCloud) GetRotationEventIn() []string`

GetRotationEventIn returns the RotationEventIn field if non-nil, zero value otherwise.

### GetRotationEventInOk

`func (o *RotatedSecretUpdateAliCloud) GetRotationEventInOk() (*[]string, bool)`

GetRotationEventInOk returns a tuple with the RotationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationEventIn

`func (o *RotatedSecretUpdateAliCloud) SetRotationEventIn(v []string)`

SetRotationEventIn sets RotationEventIn field to given value.

### HasRotationEventIn

`func (o *RotatedSecretUpdateAliCloud) HasRotationEventIn() bool`

HasRotationEventIn returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretUpdateAliCloud) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretUpdateAliCloud) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretUpdateAliCloud) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretUpdateAliCloud) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretUpdateAliCloud) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretUpdateAliCloud) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretUpdateAliCloud) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretUpdateAliCloud) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetSkipDryRun

`func (o *RotatedSecretUpdateAliCloud) GetSkipDryRun() bool`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *RotatedSecretUpdateAliCloud) GetSkipDryRunOk() (*bool, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *RotatedSecretUpdateAliCloud) SetSkipDryRun(v bool)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *RotatedSecretUpdateAliCloud) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetToken

`func (o *RotatedSecretUpdateAliCloud) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretUpdateAliCloud) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretUpdateAliCloud) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretUpdateAliCloud) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretUpdateAliCloud) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretUpdateAliCloud) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretUpdateAliCloud) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretUpdateAliCloud) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseCapitalLetters

`func (o *RotatedSecretUpdateAliCloud) GetUseCapitalLetters() string`

GetUseCapitalLetters returns the UseCapitalLetters field if non-nil, zero value otherwise.

### GetUseCapitalLettersOk

`func (o *RotatedSecretUpdateAliCloud) GetUseCapitalLettersOk() (*string, bool)`

GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCapitalLetters

`func (o *RotatedSecretUpdateAliCloud) SetUseCapitalLetters(v string)`

SetUseCapitalLetters sets UseCapitalLetters field to given value.

### HasUseCapitalLetters

`func (o *RotatedSecretUpdateAliCloud) HasUseCapitalLetters() bool`

HasUseCapitalLetters returns a boolean if a field has been set.

### GetUseLowerLetters

`func (o *RotatedSecretUpdateAliCloud) GetUseLowerLetters() string`

GetUseLowerLetters returns the UseLowerLetters field if non-nil, zero value otherwise.

### GetUseLowerLettersOk

`func (o *RotatedSecretUpdateAliCloud) GetUseLowerLettersOk() (*string, bool)`

GetUseLowerLettersOk returns a tuple with the UseLowerLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLowerLetters

`func (o *RotatedSecretUpdateAliCloud) SetUseLowerLetters(v string)`

SetUseLowerLetters sets UseLowerLetters field to given value.

### HasUseLowerLetters

`func (o *RotatedSecretUpdateAliCloud) HasUseLowerLetters() bool`

HasUseLowerLetters returns a boolean if a field has been set.

### GetUseNumbers

`func (o *RotatedSecretUpdateAliCloud) GetUseNumbers() string`

GetUseNumbers returns the UseNumbers field if non-nil, zero value otherwise.

### GetUseNumbersOk

`func (o *RotatedSecretUpdateAliCloud) GetUseNumbersOk() (*string, bool)`

GetUseNumbersOk returns a tuple with the UseNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumbers

`func (o *RotatedSecretUpdateAliCloud) SetUseNumbers(v string)`

SetUseNumbers sets UseNumbers field to given value.

### HasUseNumbers

`func (o *RotatedSecretUpdateAliCloud) HasUseNumbers() bool`

HasUseNumbers returns a boolean if a field has been set.

### GetUseSpecialCharacters

`func (o *RotatedSecretUpdateAliCloud) GetUseSpecialCharacters() string`

GetUseSpecialCharacters returns the UseSpecialCharacters field if non-nil, zero value otherwise.

### GetUseSpecialCharactersOk

`func (o *RotatedSecretUpdateAliCloud) GetUseSpecialCharactersOk() (*string, bool)`

GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpecialCharacters

`func (o *RotatedSecretUpdateAliCloud) SetUseSpecialCharacters(v string)`

SetUseSpecialCharacters sets UseSpecialCharacters field to given value.

### HasUseSpecialCharacters

`func (o *RotatedSecretUpdateAliCloud) HasUseSpecialCharacters() bool`

HasUseSpecialCharacters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


