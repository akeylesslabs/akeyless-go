# CreateMcpSecretOAuthClientCreds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. When false, user-defined input/output rules are stored but not enforced; the base security validation still runs.  AraEnabled is tri-state (nil/true/false), not a plain bool: it self-encodes its wire value (see akl.OptionalBool) so an explicit false survives the curl-proxy relay instead of being dropped like a default-false bool flag. | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EnableAgenticRuntimeAuthority** | Pointer to **bool** | EnableAra is the documented spelling of AraEnabled. Both set the same field; --ara-enabled shipped first and stays as an undocumented alias so existing scripts and the Terraform provider keep working. | [optional] 
**EnableAiQuorum** | Pointer to **bool** | Turns on AI Quorum checks for this item. | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | Secret name | 
**OauthClientId** | Pointer to **string** | OAuth client ID | [optional] 
**OauthClientSecret** | Pointer to **string** | OAuth client secret | [optional] 
**OauthScopes** | Pointer to **[]string** | OAuth scopes | [optional] 
**OauthTokenUrl** | Pointer to **string** | OAuth token URL | [optional] 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**ProtectionKey** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | Pointer to **string** | URL of the service | [optional] 

## Methods

### NewCreateMcpSecretOAuthClientCreds

`func NewCreateMcpSecretOAuthClientCreds(name string, ) *CreateMcpSecretOAuthClientCreds`

NewCreateMcpSecretOAuthClientCreds instantiates a new CreateMcpSecretOAuthClientCreds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMcpSecretOAuthClientCredsWithDefaults

`func NewCreateMcpSecretOAuthClientCredsWithDefaults() *CreateMcpSecretOAuthClientCreds`

NewCreateMcpSecretOAuthClientCredsWithDefaults instantiates a new CreateMcpSecretOAuthClientCreds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *CreateMcpSecretOAuthClientCreds) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *CreateMcpSecretOAuthClientCreds) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *CreateMcpSecretOAuthClientCreds) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *CreateMcpSecretOAuthClientCreds) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetAraEnabled

`func (o *CreateMcpSecretOAuthClientCreds) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *CreateMcpSecretOAuthClientCreds) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *CreateMcpSecretOAuthClientCreds) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *CreateMcpSecretOAuthClientCreds) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateMcpSecretOAuthClientCreds) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateMcpSecretOAuthClientCreds) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateMcpSecretOAuthClientCreds) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateMcpSecretOAuthClientCreds) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateMcpSecretOAuthClientCreds) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMcpSecretOAuthClientCreds) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMcpSecretOAuthClientCreds) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMcpSecretOAuthClientCreds) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAgenticRuntimeAuthority

`func (o *CreateMcpSecretOAuthClientCreds) GetEnableAgenticRuntimeAuthority() bool`

GetEnableAgenticRuntimeAuthority returns the EnableAgenticRuntimeAuthority field if non-nil, zero value otherwise.

### GetEnableAgenticRuntimeAuthorityOk

`func (o *CreateMcpSecretOAuthClientCreds) GetEnableAgenticRuntimeAuthorityOk() (*bool, bool)`

GetEnableAgenticRuntimeAuthorityOk returns a tuple with the EnableAgenticRuntimeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAgenticRuntimeAuthority

`func (o *CreateMcpSecretOAuthClientCreds) SetEnableAgenticRuntimeAuthority(v bool)`

SetEnableAgenticRuntimeAuthority sets EnableAgenticRuntimeAuthority field to given value.

### HasEnableAgenticRuntimeAuthority

`func (o *CreateMcpSecretOAuthClientCreds) HasEnableAgenticRuntimeAuthority() bool`

HasEnableAgenticRuntimeAuthority returns a boolean if a field has been set.

### GetEnableAiQuorum

`func (o *CreateMcpSecretOAuthClientCreds) GetEnableAiQuorum() bool`

GetEnableAiQuorum returns the EnableAiQuorum field if non-nil, zero value otherwise.

### GetEnableAiQuorumOk

`func (o *CreateMcpSecretOAuthClientCreds) GetEnableAiQuorumOk() (*bool, bool)`

GetEnableAiQuorumOk returns a tuple with the EnableAiQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiQuorum

`func (o *CreateMcpSecretOAuthClientCreds) SetEnableAiQuorum(v bool)`

SetEnableAiQuorum sets EnableAiQuorum field to given value.

### HasEnableAiQuorum

`func (o *CreateMcpSecretOAuthClientCreds) HasEnableAiQuorum() bool`

HasEnableAiQuorum returns a boolean if a field has been set.

### GetInputRule

`func (o *CreateMcpSecretOAuthClientCreds) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *CreateMcpSecretOAuthClientCreds) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *CreateMcpSecretOAuthClientCreds) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *CreateMcpSecretOAuthClientCreds) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetJson

`func (o *CreateMcpSecretOAuthClientCreds) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateMcpSecretOAuthClientCreds) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateMcpSecretOAuthClientCreds) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateMcpSecretOAuthClientCreds) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMaxVersions

`func (o *CreateMcpSecretOAuthClientCreds) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *CreateMcpSecretOAuthClientCreds) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *CreateMcpSecretOAuthClientCreds) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *CreateMcpSecretOAuthClientCreds) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateMcpSecretOAuthClientCreds) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateMcpSecretOAuthClientCreds) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateMcpSecretOAuthClientCreds) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateMcpSecretOAuthClientCreds) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateMcpSecretOAuthClientCreds) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMcpSecretOAuthClientCreds) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMcpSecretOAuthClientCreds) SetName(v string)`

SetName sets Name field to given value.


### GetOauthClientId

`func (o *CreateMcpSecretOAuthClientCreds) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *CreateMcpSecretOAuthClientCreds) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *CreateMcpSecretOAuthClientCreds) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *CreateMcpSecretOAuthClientCreds) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### GetOauthClientSecret

`func (o *CreateMcpSecretOAuthClientCreds) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *CreateMcpSecretOAuthClientCreds) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *CreateMcpSecretOAuthClientCreds) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *CreateMcpSecretOAuthClientCreds) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### GetOauthScopes

`func (o *CreateMcpSecretOAuthClientCreds) GetOauthScopes() []string`

GetOauthScopes returns the OauthScopes field if non-nil, zero value otherwise.

### GetOauthScopesOk

`func (o *CreateMcpSecretOAuthClientCreds) GetOauthScopesOk() (*[]string, bool)`

GetOauthScopesOk returns a tuple with the OauthScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthScopes

`func (o *CreateMcpSecretOAuthClientCreds) SetOauthScopes(v []string)`

SetOauthScopes sets OauthScopes field to given value.

### HasOauthScopes

`func (o *CreateMcpSecretOAuthClientCreds) HasOauthScopes() bool`

HasOauthScopes returns a boolean if a field has been set.

### GetOauthTokenUrl

`func (o *CreateMcpSecretOAuthClientCreds) GetOauthTokenUrl() string`

GetOauthTokenUrl returns the OauthTokenUrl field if non-nil, zero value otherwise.

### GetOauthTokenUrlOk

`func (o *CreateMcpSecretOAuthClientCreds) GetOauthTokenUrlOk() (*string, bool)`

GetOauthTokenUrlOk returns a tuple with the OauthTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenUrl

`func (o *CreateMcpSecretOAuthClientCreds) SetOauthTokenUrl(v string)`

SetOauthTokenUrl sets OauthTokenUrl field to given value.

### HasOauthTokenUrl

`func (o *CreateMcpSecretOAuthClientCreds) HasOauthTokenUrl() bool`

HasOauthTokenUrl returns a boolean if a field has been set.

### GetOutputRule

`func (o *CreateMcpSecretOAuthClientCreds) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *CreateMcpSecretOAuthClientCreds) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *CreateMcpSecretOAuthClientCreds) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *CreateMcpSecretOAuthClientCreds) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetProtectionKey

`func (o *CreateMcpSecretOAuthClientCreds) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateMcpSecretOAuthClientCreds) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateMcpSecretOAuthClientCreds) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateMcpSecretOAuthClientCreds) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTags

`func (o *CreateMcpSecretOAuthClientCreds) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateMcpSecretOAuthClientCreds) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateMcpSecretOAuthClientCreds) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateMcpSecretOAuthClientCreds) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *CreateMcpSecretOAuthClientCreds) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateMcpSecretOAuthClientCreds) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateMcpSecretOAuthClientCreds) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateMcpSecretOAuthClientCreds) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateMcpSecretOAuthClientCreds) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateMcpSecretOAuthClientCreds) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateMcpSecretOAuthClientCreds) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateMcpSecretOAuthClientCreds) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *CreateMcpSecretOAuthClientCreds) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateMcpSecretOAuthClientCreds) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateMcpSecretOAuthClientCreds) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateMcpSecretOAuthClientCreds) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


