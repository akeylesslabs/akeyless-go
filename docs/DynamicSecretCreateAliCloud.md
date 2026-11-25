# DynamicSecretCreateAliCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlicloudAccessKeyId** | Pointer to **string** | AliCloud RAM access key ID | [optional] 
**AlicloudAccessMode** | Pointer to **string** |  | [optional] 
**AlicloudAccessSecretKey** | Pointer to **string** | AliCloud RAM access key secret | [optional] 
**AlicloudExternalId** | Pointer to **string** | The external ID in the trust policy of the AliCloud roles (relevant only for assume_role mode) | [optional] 
**AlicloudRegion** | Pointer to **string** | AliCloud region. Selects the regional STS endpoint. If empty, the global STS endpoint is used | [optional] 
**AlicloudRoleArns** | Pointer to **string** | AliCloud role ARNs allowed in the Assume Role operation (relevant only for assume_role mode) | [optional] 
**AlicloudSessionPolicy** | Pointer to **string** | Session policy document (JSON) that limits the assumed role session (relevant only for assume_role mode) | [optional] 
**AlicloudSetSourceIdentity** | Pointer to **bool** | Send the requester&#39;s identity as the SourceIdentity of the assumed role session, so it shows in ActionTrail (relevant only for assume_role mode). Needs sts:SetSourceIdentity in both the producer identity&#39;s policy and the role trust policy | [optional] [default to false]
**AlicloudSite** | Pointer to **string** | The Alibaba Cloud site of the account: international or china. It selects the console sign-in host of the RAM user | [optional] [default to "international"]
**AlicloudUserConsoleAccess** | Pointer to **bool** | Enable AliCloud RAM user console access (relevant only for ram_user mode) | [optional] [default to false]
**AlicloudUserGroups** | Pointer to **string** | AliCloud user groups (relevant only for ram_user mode) | [optional] 
**AlicloudUserInlinePolicy** | Pointer to **string** | RAM policy document (JSON) for a custom policy created for each RAM user (relevant only for ram_user mode) | [optional] 
**AlicloudUserPolicies** | Pointer to **string** | AliCloud policies attached to the RAM user, in the form &lt;System|Custom&gt;:&lt;PolicyName&gt; (relevant only for ram_user mode) | [optional] 
**AlicloudUserProgrammaticAccess** | Pointer to **bool** | Enable AliCloud RAM user programmatic access (relevant only for ram_user mode) | [optional] [default to true]
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. Mirrors commands.AgenticRulesParams.AraEnabled. | [optional] 
**CustomUsernameTemplate** | Pointer to **string** | Customize how temporary usernames are generated using go template | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EnableAgenticRuntimeAuthority** | Pointer to **bool** | EnableAra is the documented spelling of AraEnabled; --ara-enabled shipped first and stays as an undocumented alias. | [optional] 
**EnableAiQuorum** | Pointer to **bool** | Turns on AI Quorum checks for this item. | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) Mirrors commands.AgenticRulesParams — kept separate because ResourceDS cannot embed it (different package, different struct layout). | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SkipDryRun** | Pointer to **bool** | If set, dry-run will be skipped [true/false] | [optional] 
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

### NewDynamicSecretCreateAliCloud

`func NewDynamicSecretCreateAliCloud(name string, ) *DynamicSecretCreateAliCloud`

NewDynamicSecretCreateAliCloud instantiates a new DynamicSecretCreateAliCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateAliCloudWithDefaults

`func NewDynamicSecretCreateAliCloudWithDefaults() *DynamicSecretCreateAliCloud`

NewDynamicSecretCreateAliCloudWithDefaults instantiates a new DynamicSecretCreateAliCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlicloudAccessKeyId

`func (o *DynamicSecretCreateAliCloud) GetAlicloudAccessKeyId() string`

GetAlicloudAccessKeyId returns the AlicloudAccessKeyId field if non-nil, zero value otherwise.

### GetAlicloudAccessKeyIdOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudAccessKeyIdOk() (*string, bool)`

GetAlicloudAccessKeyIdOk returns a tuple with the AlicloudAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudAccessKeyId

`func (o *DynamicSecretCreateAliCloud) SetAlicloudAccessKeyId(v string)`

SetAlicloudAccessKeyId sets AlicloudAccessKeyId field to given value.

### HasAlicloudAccessKeyId

`func (o *DynamicSecretCreateAliCloud) HasAlicloudAccessKeyId() bool`

HasAlicloudAccessKeyId returns a boolean if a field has been set.

### GetAlicloudAccessMode

`func (o *DynamicSecretCreateAliCloud) GetAlicloudAccessMode() string`

GetAlicloudAccessMode returns the AlicloudAccessMode field if non-nil, zero value otherwise.

### GetAlicloudAccessModeOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudAccessModeOk() (*string, bool)`

GetAlicloudAccessModeOk returns a tuple with the AlicloudAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudAccessMode

`func (o *DynamicSecretCreateAliCloud) SetAlicloudAccessMode(v string)`

SetAlicloudAccessMode sets AlicloudAccessMode field to given value.

### HasAlicloudAccessMode

`func (o *DynamicSecretCreateAliCloud) HasAlicloudAccessMode() bool`

HasAlicloudAccessMode returns a boolean if a field has been set.

### GetAlicloudAccessSecretKey

`func (o *DynamicSecretCreateAliCloud) GetAlicloudAccessSecretKey() string`

GetAlicloudAccessSecretKey returns the AlicloudAccessSecretKey field if non-nil, zero value otherwise.

### GetAlicloudAccessSecretKeyOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudAccessSecretKeyOk() (*string, bool)`

GetAlicloudAccessSecretKeyOk returns a tuple with the AlicloudAccessSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudAccessSecretKey

`func (o *DynamicSecretCreateAliCloud) SetAlicloudAccessSecretKey(v string)`

SetAlicloudAccessSecretKey sets AlicloudAccessSecretKey field to given value.

### HasAlicloudAccessSecretKey

`func (o *DynamicSecretCreateAliCloud) HasAlicloudAccessSecretKey() bool`

HasAlicloudAccessSecretKey returns a boolean if a field has been set.

### GetAlicloudExternalId

`func (o *DynamicSecretCreateAliCloud) GetAlicloudExternalId() string`

GetAlicloudExternalId returns the AlicloudExternalId field if non-nil, zero value otherwise.

### GetAlicloudExternalIdOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudExternalIdOk() (*string, bool)`

GetAlicloudExternalIdOk returns a tuple with the AlicloudExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudExternalId

`func (o *DynamicSecretCreateAliCloud) SetAlicloudExternalId(v string)`

SetAlicloudExternalId sets AlicloudExternalId field to given value.

### HasAlicloudExternalId

`func (o *DynamicSecretCreateAliCloud) HasAlicloudExternalId() bool`

HasAlicloudExternalId returns a boolean if a field has been set.

### GetAlicloudRegion

`func (o *DynamicSecretCreateAliCloud) GetAlicloudRegion() string`

GetAlicloudRegion returns the AlicloudRegion field if non-nil, zero value otherwise.

### GetAlicloudRegionOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudRegionOk() (*string, bool)`

GetAlicloudRegionOk returns a tuple with the AlicloudRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudRegion

`func (o *DynamicSecretCreateAliCloud) SetAlicloudRegion(v string)`

SetAlicloudRegion sets AlicloudRegion field to given value.

### HasAlicloudRegion

`func (o *DynamicSecretCreateAliCloud) HasAlicloudRegion() bool`

HasAlicloudRegion returns a boolean if a field has been set.

### GetAlicloudRoleArns

`func (o *DynamicSecretCreateAliCloud) GetAlicloudRoleArns() string`

GetAlicloudRoleArns returns the AlicloudRoleArns field if non-nil, zero value otherwise.

### GetAlicloudRoleArnsOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudRoleArnsOk() (*string, bool)`

GetAlicloudRoleArnsOk returns a tuple with the AlicloudRoleArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudRoleArns

`func (o *DynamicSecretCreateAliCloud) SetAlicloudRoleArns(v string)`

SetAlicloudRoleArns sets AlicloudRoleArns field to given value.

### HasAlicloudRoleArns

`func (o *DynamicSecretCreateAliCloud) HasAlicloudRoleArns() bool`

HasAlicloudRoleArns returns a boolean if a field has been set.

### GetAlicloudSessionPolicy

`func (o *DynamicSecretCreateAliCloud) GetAlicloudSessionPolicy() string`

GetAlicloudSessionPolicy returns the AlicloudSessionPolicy field if non-nil, zero value otherwise.

### GetAlicloudSessionPolicyOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudSessionPolicyOk() (*string, bool)`

GetAlicloudSessionPolicyOk returns a tuple with the AlicloudSessionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudSessionPolicy

`func (o *DynamicSecretCreateAliCloud) SetAlicloudSessionPolicy(v string)`

SetAlicloudSessionPolicy sets AlicloudSessionPolicy field to given value.

### HasAlicloudSessionPolicy

`func (o *DynamicSecretCreateAliCloud) HasAlicloudSessionPolicy() bool`

HasAlicloudSessionPolicy returns a boolean if a field has been set.

### GetAlicloudSetSourceIdentity

`func (o *DynamicSecretCreateAliCloud) GetAlicloudSetSourceIdentity() bool`

GetAlicloudSetSourceIdentity returns the AlicloudSetSourceIdentity field if non-nil, zero value otherwise.

### GetAlicloudSetSourceIdentityOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudSetSourceIdentityOk() (*bool, bool)`

GetAlicloudSetSourceIdentityOk returns a tuple with the AlicloudSetSourceIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudSetSourceIdentity

`func (o *DynamicSecretCreateAliCloud) SetAlicloudSetSourceIdentity(v bool)`

SetAlicloudSetSourceIdentity sets AlicloudSetSourceIdentity field to given value.

### HasAlicloudSetSourceIdentity

`func (o *DynamicSecretCreateAliCloud) HasAlicloudSetSourceIdentity() bool`

HasAlicloudSetSourceIdentity returns a boolean if a field has been set.

### GetAlicloudSite

`func (o *DynamicSecretCreateAliCloud) GetAlicloudSite() string`

GetAlicloudSite returns the AlicloudSite field if non-nil, zero value otherwise.

### GetAlicloudSiteOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudSiteOk() (*string, bool)`

GetAlicloudSiteOk returns a tuple with the AlicloudSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudSite

`func (o *DynamicSecretCreateAliCloud) SetAlicloudSite(v string)`

SetAlicloudSite sets AlicloudSite field to given value.

### HasAlicloudSite

`func (o *DynamicSecretCreateAliCloud) HasAlicloudSite() bool`

HasAlicloudSite returns a boolean if a field has been set.

### GetAlicloudUserConsoleAccess

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserConsoleAccess() bool`

GetAlicloudUserConsoleAccess returns the AlicloudUserConsoleAccess field if non-nil, zero value otherwise.

### GetAlicloudUserConsoleAccessOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserConsoleAccessOk() (*bool, bool)`

GetAlicloudUserConsoleAccessOk returns a tuple with the AlicloudUserConsoleAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudUserConsoleAccess

`func (o *DynamicSecretCreateAliCloud) SetAlicloudUserConsoleAccess(v bool)`

SetAlicloudUserConsoleAccess sets AlicloudUserConsoleAccess field to given value.

### HasAlicloudUserConsoleAccess

`func (o *DynamicSecretCreateAliCloud) HasAlicloudUserConsoleAccess() bool`

HasAlicloudUserConsoleAccess returns a boolean if a field has been set.

### GetAlicloudUserGroups

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserGroups() string`

GetAlicloudUserGroups returns the AlicloudUserGroups field if non-nil, zero value otherwise.

### GetAlicloudUserGroupsOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserGroupsOk() (*string, bool)`

GetAlicloudUserGroupsOk returns a tuple with the AlicloudUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudUserGroups

`func (o *DynamicSecretCreateAliCloud) SetAlicloudUserGroups(v string)`

SetAlicloudUserGroups sets AlicloudUserGroups field to given value.

### HasAlicloudUserGroups

`func (o *DynamicSecretCreateAliCloud) HasAlicloudUserGroups() bool`

HasAlicloudUserGroups returns a boolean if a field has been set.

### GetAlicloudUserInlinePolicy

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserInlinePolicy() string`

GetAlicloudUserInlinePolicy returns the AlicloudUserInlinePolicy field if non-nil, zero value otherwise.

### GetAlicloudUserInlinePolicyOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserInlinePolicyOk() (*string, bool)`

GetAlicloudUserInlinePolicyOk returns a tuple with the AlicloudUserInlinePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudUserInlinePolicy

`func (o *DynamicSecretCreateAliCloud) SetAlicloudUserInlinePolicy(v string)`

SetAlicloudUserInlinePolicy sets AlicloudUserInlinePolicy field to given value.

### HasAlicloudUserInlinePolicy

`func (o *DynamicSecretCreateAliCloud) HasAlicloudUserInlinePolicy() bool`

HasAlicloudUserInlinePolicy returns a boolean if a field has been set.

### GetAlicloudUserPolicies

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserPolicies() string`

GetAlicloudUserPolicies returns the AlicloudUserPolicies field if non-nil, zero value otherwise.

### GetAlicloudUserPoliciesOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserPoliciesOk() (*string, bool)`

GetAlicloudUserPoliciesOk returns a tuple with the AlicloudUserPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudUserPolicies

`func (o *DynamicSecretCreateAliCloud) SetAlicloudUserPolicies(v string)`

SetAlicloudUserPolicies sets AlicloudUserPolicies field to given value.

### HasAlicloudUserPolicies

`func (o *DynamicSecretCreateAliCloud) HasAlicloudUserPolicies() bool`

HasAlicloudUserPolicies returns a boolean if a field has been set.

### GetAlicloudUserProgrammaticAccess

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserProgrammaticAccess() bool`

GetAlicloudUserProgrammaticAccess returns the AlicloudUserProgrammaticAccess field if non-nil, zero value otherwise.

### GetAlicloudUserProgrammaticAccessOk

`func (o *DynamicSecretCreateAliCloud) GetAlicloudUserProgrammaticAccessOk() (*bool, bool)`

GetAlicloudUserProgrammaticAccessOk returns a tuple with the AlicloudUserProgrammaticAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlicloudUserProgrammaticAccess

`func (o *DynamicSecretCreateAliCloud) SetAlicloudUserProgrammaticAccess(v bool)`

SetAlicloudUserProgrammaticAccess sets AlicloudUserProgrammaticAccess field to given value.

### HasAlicloudUserProgrammaticAccess

`func (o *DynamicSecretCreateAliCloud) HasAlicloudUserProgrammaticAccess() bool`

HasAlicloudUserProgrammaticAccess returns a boolean if a field has been set.

### GetAraEnabled

`func (o *DynamicSecretCreateAliCloud) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *DynamicSecretCreateAliCloud) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *DynamicSecretCreateAliCloud) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *DynamicSecretCreateAliCloud) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetCustomUsernameTemplate

`func (o *DynamicSecretCreateAliCloud) GetCustomUsernameTemplate() string`

GetCustomUsernameTemplate returns the CustomUsernameTemplate field if non-nil, zero value otherwise.

### GetCustomUsernameTemplateOk

`func (o *DynamicSecretCreateAliCloud) GetCustomUsernameTemplateOk() (*string, bool)`

GetCustomUsernameTemplateOk returns a tuple with the CustomUsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUsernameTemplate

`func (o *DynamicSecretCreateAliCloud) SetCustomUsernameTemplate(v string)`

SetCustomUsernameTemplate sets CustomUsernameTemplate field to given value.

### HasCustomUsernameTemplate

`func (o *DynamicSecretCreateAliCloud) HasCustomUsernameTemplate() bool`

HasCustomUsernameTemplate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateAliCloud) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateAliCloud) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateAliCloud) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateAliCloud) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateAliCloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateAliCloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateAliCloud) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateAliCloud) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAgenticRuntimeAuthority

`func (o *DynamicSecretCreateAliCloud) GetEnableAgenticRuntimeAuthority() bool`

GetEnableAgenticRuntimeAuthority returns the EnableAgenticRuntimeAuthority field if non-nil, zero value otherwise.

### GetEnableAgenticRuntimeAuthorityOk

`func (o *DynamicSecretCreateAliCloud) GetEnableAgenticRuntimeAuthorityOk() (*bool, bool)`

GetEnableAgenticRuntimeAuthorityOk returns a tuple with the EnableAgenticRuntimeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAgenticRuntimeAuthority

`func (o *DynamicSecretCreateAliCloud) SetEnableAgenticRuntimeAuthority(v bool)`

SetEnableAgenticRuntimeAuthority sets EnableAgenticRuntimeAuthority field to given value.

### HasEnableAgenticRuntimeAuthority

`func (o *DynamicSecretCreateAliCloud) HasEnableAgenticRuntimeAuthority() bool`

HasEnableAgenticRuntimeAuthority returns a boolean if a field has been set.

### GetEnableAiQuorum

`func (o *DynamicSecretCreateAliCloud) GetEnableAiQuorum() bool`

GetEnableAiQuorum returns the EnableAiQuorum field if non-nil, zero value otherwise.

### GetEnableAiQuorumOk

`func (o *DynamicSecretCreateAliCloud) GetEnableAiQuorumOk() (*bool, bool)`

GetEnableAiQuorumOk returns a tuple with the EnableAiQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiQuorum

`func (o *DynamicSecretCreateAliCloud) SetEnableAiQuorum(v bool)`

SetEnableAiQuorum sets EnableAiQuorum field to given value.

### HasEnableAiQuorum

`func (o *DynamicSecretCreateAliCloud) HasEnableAiQuorum() bool`

HasEnableAiQuorum returns a boolean if a field has been set.

### GetInputRule

`func (o *DynamicSecretCreateAliCloud) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *DynamicSecretCreateAliCloud) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *DynamicSecretCreateAliCloud) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *DynamicSecretCreateAliCloud) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *DynamicSecretCreateAliCloud) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *DynamicSecretCreateAliCloud) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *DynamicSecretCreateAliCloud) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *DynamicSecretCreateAliCloud) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateAliCloud) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateAliCloud) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateAliCloud) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateAliCloud) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateAliCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateAliCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateAliCloud) SetName(v string)`

SetName sets Name field to given value.


### GetOutputRule

`func (o *DynamicSecretCreateAliCloud) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *DynamicSecretCreateAliCloud) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *DynamicSecretCreateAliCloud) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *DynamicSecretCreateAliCloud) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretCreateAliCloud) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretCreateAliCloud) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretCreateAliCloud) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretCreateAliCloud) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateAliCloud) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateAliCloud) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateAliCloud) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateAliCloud) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSkipDryRun

`func (o *DynamicSecretCreateAliCloud) GetSkipDryRun() bool`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *DynamicSecretCreateAliCloud) GetSkipDryRunOk() (*bool, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *DynamicSecretCreateAliCloud) SetSkipDryRun(v bool)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *DynamicSecretCreateAliCloud) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateAliCloud) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateAliCloud) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateAliCloud) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateAliCloud) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateAliCloud) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateAliCloud) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateAliCloud) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateAliCloud) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateAliCloud) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateAliCloud) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateAliCloud) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateAliCloud) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateAliCloud) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateAliCloud) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateAliCloud) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateAliCloud) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseCapitalLetters

`func (o *DynamicSecretCreateAliCloud) GetUseCapitalLetters() string`

GetUseCapitalLetters returns the UseCapitalLetters field if non-nil, zero value otherwise.

### GetUseCapitalLettersOk

`func (o *DynamicSecretCreateAliCloud) GetUseCapitalLettersOk() (*string, bool)`

GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCapitalLetters

`func (o *DynamicSecretCreateAliCloud) SetUseCapitalLetters(v string)`

SetUseCapitalLetters sets UseCapitalLetters field to given value.

### HasUseCapitalLetters

`func (o *DynamicSecretCreateAliCloud) HasUseCapitalLetters() bool`

HasUseCapitalLetters returns a boolean if a field has been set.

### GetUseLowerLetters

`func (o *DynamicSecretCreateAliCloud) GetUseLowerLetters() string`

GetUseLowerLetters returns the UseLowerLetters field if non-nil, zero value otherwise.

### GetUseLowerLettersOk

`func (o *DynamicSecretCreateAliCloud) GetUseLowerLettersOk() (*string, bool)`

GetUseLowerLettersOk returns a tuple with the UseLowerLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLowerLetters

`func (o *DynamicSecretCreateAliCloud) SetUseLowerLetters(v string)`

SetUseLowerLetters sets UseLowerLetters field to given value.

### HasUseLowerLetters

`func (o *DynamicSecretCreateAliCloud) HasUseLowerLetters() bool`

HasUseLowerLetters returns a boolean if a field has been set.

### GetUseNumbers

`func (o *DynamicSecretCreateAliCloud) GetUseNumbers() string`

GetUseNumbers returns the UseNumbers field if non-nil, zero value otherwise.

### GetUseNumbersOk

`func (o *DynamicSecretCreateAliCloud) GetUseNumbersOk() (*string, bool)`

GetUseNumbersOk returns a tuple with the UseNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumbers

`func (o *DynamicSecretCreateAliCloud) SetUseNumbers(v string)`

SetUseNumbers sets UseNumbers field to given value.

### HasUseNumbers

`func (o *DynamicSecretCreateAliCloud) HasUseNumbers() bool`

HasUseNumbers returns a boolean if a field has been set.

### GetUseSpecialCharacters

`func (o *DynamicSecretCreateAliCloud) GetUseSpecialCharacters() string`

GetUseSpecialCharacters returns the UseSpecialCharacters field if non-nil, zero value otherwise.

### GetUseSpecialCharactersOk

`func (o *DynamicSecretCreateAliCloud) GetUseSpecialCharactersOk() (*string, bool)`

GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpecialCharacters

`func (o *DynamicSecretCreateAliCloud) SetUseSpecialCharacters(v string)`

SetUseSpecialCharacters sets UseSpecialCharacters field to given value.

### HasUseSpecialCharacters

`func (o *DynamicSecretCreateAliCloud) HasUseSpecialCharacters() bool`

HasUseSpecialCharacters returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateAliCloud) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateAliCloud) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateAliCloud) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateAliCloud) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


