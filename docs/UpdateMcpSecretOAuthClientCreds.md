# UpdateMcpSecretOAuthClientCreds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. When false, user-defined input/output rules are stored but not enforced; the base security validation still runs.  AraEnabled is tri-state (nil/true/false), not a plain bool: it self-encodes its wire value (see akl.OptionalBool) so an explicit false survives the curl-proxy relay instead of being dropped like a default-false bool flag. | [optional] 
**EnableAgenticRuntimeAuthority** | Pointer to **bool** | EnableAra is the documented spelling of AraEnabled. Both set the same field; --ara-enabled shipped first and stays as an undocumented alias so existing scripts and the Terraform provider keep working. | [optional] 
**EnableAiQuorum** | Pointer to **bool** | Turns on AI Quorum checks for this item. | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**LastVersion** | Pointer to **int32** |  | [optional] 
**Name** | **string** | Secret name | 
**OauthClientId** | Pointer to **string** | OAuth client ID | [optional] 
**OauthClientSecret** | Pointer to **string** | OAuth client secret | [optional] 
**OauthScopes** | Pointer to **[]string** | OAuth scopes | [optional] 
**OauthTokenUrl** | Pointer to **string** | OAuth token URL | [optional] 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | Pointer to **string** | URL of the service | [optional] 

## Methods

### NewUpdateMcpSecretOAuthClientCreds

`func NewUpdateMcpSecretOAuthClientCreds(name string, ) *UpdateMcpSecretOAuthClientCreds`

NewUpdateMcpSecretOAuthClientCreds instantiates a new UpdateMcpSecretOAuthClientCreds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMcpSecretOAuthClientCredsWithDefaults

`func NewUpdateMcpSecretOAuthClientCredsWithDefaults() *UpdateMcpSecretOAuthClientCreds`

NewUpdateMcpSecretOAuthClientCredsWithDefaults instantiates a new UpdateMcpSecretOAuthClientCreds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *UpdateMcpSecretOAuthClientCreds) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *UpdateMcpSecretOAuthClientCreds) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *UpdateMcpSecretOAuthClientCreds) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetAraEnabled

`func (o *UpdateMcpSecretOAuthClientCreds) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *UpdateMcpSecretOAuthClientCreds) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *UpdateMcpSecretOAuthClientCreds) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetEnableAgenticRuntimeAuthority

`func (o *UpdateMcpSecretOAuthClientCreds) GetEnableAgenticRuntimeAuthority() bool`

GetEnableAgenticRuntimeAuthority returns the EnableAgenticRuntimeAuthority field if non-nil, zero value otherwise.

### GetEnableAgenticRuntimeAuthorityOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetEnableAgenticRuntimeAuthorityOk() (*bool, bool)`

GetEnableAgenticRuntimeAuthorityOk returns a tuple with the EnableAgenticRuntimeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAgenticRuntimeAuthority

`func (o *UpdateMcpSecretOAuthClientCreds) SetEnableAgenticRuntimeAuthority(v bool)`

SetEnableAgenticRuntimeAuthority sets EnableAgenticRuntimeAuthority field to given value.

### HasEnableAgenticRuntimeAuthority

`func (o *UpdateMcpSecretOAuthClientCreds) HasEnableAgenticRuntimeAuthority() bool`

HasEnableAgenticRuntimeAuthority returns a boolean if a field has been set.

### GetEnableAiQuorum

`func (o *UpdateMcpSecretOAuthClientCreds) GetEnableAiQuorum() bool`

GetEnableAiQuorum returns the EnableAiQuorum field if non-nil, zero value otherwise.

### GetEnableAiQuorumOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetEnableAiQuorumOk() (*bool, bool)`

GetEnableAiQuorumOk returns a tuple with the EnableAiQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiQuorum

`func (o *UpdateMcpSecretOAuthClientCreds) SetEnableAiQuorum(v bool)`

SetEnableAiQuorum sets EnableAiQuorum field to given value.

### HasEnableAiQuorum

`func (o *UpdateMcpSecretOAuthClientCreds) HasEnableAiQuorum() bool`

HasEnableAiQuorum returns a boolean if a field has been set.

### GetInputRule

`func (o *UpdateMcpSecretOAuthClientCreds) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *UpdateMcpSecretOAuthClientCreds) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *UpdateMcpSecretOAuthClientCreds) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetJson

`func (o *UpdateMcpSecretOAuthClientCreds) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateMcpSecretOAuthClientCreds) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateMcpSecretOAuthClientCreds) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateMcpSecretOAuthClientCreds) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateMcpSecretOAuthClientCreds) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateMcpSecretOAuthClientCreds) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateMcpSecretOAuthClientCreds) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateMcpSecretOAuthClientCreds) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateMcpSecretOAuthClientCreds) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastVersion

`func (o *UpdateMcpSecretOAuthClientCreds) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *UpdateMcpSecretOAuthClientCreds) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *UpdateMcpSecretOAuthClientCreds) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### GetName

`func (o *UpdateMcpSecretOAuthClientCreds) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMcpSecretOAuthClientCreds) SetName(v string)`

SetName sets Name field to given value.


### GetOauthClientId

`func (o *UpdateMcpSecretOAuthClientCreds) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *UpdateMcpSecretOAuthClientCreds) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *UpdateMcpSecretOAuthClientCreds) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### GetOauthClientSecret

`func (o *UpdateMcpSecretOAuthClientCreds) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *UpdateMcpSecretOAuthClientCreds) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *UpdateMcpSecretOAuthClientCreds) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### GetOauthScopes

`func (o *UpdateMcpSecretOAuthClientCreds) GetOauthScopes() []string`

GetOauthScopes returns the OauthScopes field if non-nil, zero value otherwise.

### GetOauthScopesOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetOauthScopesOk() (*[]string, bool)`

GetOauthScopesOk returns a tuple with the OauthScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthScopes

`func (o *UpdateMcpSecretOAuthClientCreds) SetOauthScopes(v []string)`

SetOauthScopes sets OauthScopes field to given value.

### HasOauthScopes

`func (o *UpdateMcpSecretOAuthClientCreds) HasOauthScopes() bool`

HasOauthScopes returns a boolean if a field has been set.

### GetOauthTokenUrl

`func (o *UpdateMcpSecretOAuthClientCreds) GetOauthTokenUrl() string`

GetOauthTokenUrl returns the OauthTokenUrl field if non-nil, zero value otherwise.

### GetOauthTokenUrlOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetOauthTokenUrlOk() (*string, bool)`

GetOauthTokenUrlOk returns a tuple with the OauthTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenUrl

`func (o *UpdateMcpSecretOAuthClientCreds) SetOauthTokenUrl(v string)`

SetOauthTokenUrl sets OauthTokenUrl field to given value.

### HasOauthTokenUrl

`func (o *UpdateMcpSecretOAuthClientCreds) HasOauthTokenUrl() bool`

HasOauthTokenUrl returns a boolean if a field has been set.

### GetOutputRule

`func (o *UpdateMcpSecretOAuthClientCreds) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *UpdateMcpSecretOAuthClientCreds) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *UpdateMcpSecretOAuthClientCreds) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetToken

`func (o *UpdateMcpSecretOAuthClientCreds) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateMcpSecretOAuthClientCreds) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateMcpSecretOAuthClientCreds) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateMcpSecretOAuthClientCreds) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateMcpSecretOAuthClientCreds) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateMcpSecretOAuthClientCreds) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateMcpSecretOAuthClientCreds) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateMcpSecretOAuthClientCreds) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateMcpSecretOAuthClientCreds) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateMcpSecretOAuthClientCreds) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


