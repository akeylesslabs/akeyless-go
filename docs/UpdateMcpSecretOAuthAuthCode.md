# UpdateMcpSecretOAuthAuthCode

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
**OauthRedirectUri** | Pointer to **string** | OAuth redirect URI | [optional] 
**OauthRefreshToken** | Pointer to **string** | OAuth refresh token | [optional] 
**OauthScopes** | Pointer to **[]string** | OAuth scopes | [optional] 
**OauthTokenUrl** | Pointer to **string** | OAuth token URL | [optional] 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | Pointer to **string** | URL of the service | [optional] 

## Methods

### NewUpdateMcpSecretOAuthAuthCode

`func NewUpdateMcpSecretOAuthAuthCode(name string, ) *UpdateMcpSecretOAuthAuthCode`

NewUpdateMcpSecretOAuthAuthCode instantiates a new UpdateMcpSecretOAuthAuthCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMcpSecretOAuthAuthCodeWithDefaults

`func NewUpdateMcpSecretOAuthAuthCodeWithDefaults() *UpdateMcpSecretOAuthAuthCode`

NewUpdateMcpSecretOAuthAuthCodeWithDefaults instantiates a new UpdateMcpSecretOAuthAuthCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *UpdateMcpSecretOAuthAuthCode) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *UpdateMcpSecretOAuthAuthCode) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *UpdateMcpSecretOAuthAuthCode) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetAraEnabled

`func (o *UpdateMcpSecretOAuthAuthCode) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *UpdateMcpSecretOAuthAuthCode) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *UpdateMcpSecretOAuthAuthCode) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetEnableAgenticRuntimeAuthority

`func (o *UpdateMcpSecretOAuthAuthCode) GetEnableAgenticRuntimeAuthority() bool`

GetEnableAgenticRuntimeAuthority returns the EnableAgenticRuntimeAuthority field if non-nil, zero value otherwise.

### GetEnableAgenticRuntimeAuthorityOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetEnableAgenticRuntimeAuthorityOk() (*bool, bool)`

GetEnableAgenticRuntimeAuthorityOk returns a tuple with the EnableAgenticRuntimeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAgenticRuntimeAuthority

`func (o *UpdateMcpSecretOAuthAuthCode) SetEnableAgenticRuntimeAuthority(v bool)`

SetEnableAgenticRuntimeAuthority sets EnableAgenticRuntimeAuthority field to given value.

### HasEnableAgenticRuntimeAuthority

`func (o *UpdateMcpSecretOAuthAuthCode) HasEnableAgenticRuntimeAuthority() bool`

HasEnableAgenticRuntimeAuthority returns a boolean if a field has been set.

### GetEnableAiQuorum

`func (o *UpdateMcpSecretOAuthAuthCode) GetEnableAiQuorum() bool`

GetEnableAiQuorum returns the EnableAiQuorum field if non-nil, zero value otherwise.

### GetEnableAiQuorumOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetEnableAiQuorumOk() (*bool, bool)`

GetEnableAiQuorumOk returns a tuple with the EnableAiQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiQuorum

`func (o *UpdateMcpSecretOAuthAuthCode) SetEnableAiQuorum(v bool)`

SetEnableAiQuorum sets EnableAiQuorum field to given value.

### HasEnableAiQuorum

`func (o *UpdateMcpSecretOAuthAuthCode) HasEnableAiQuorum() bool`

HasEnableAiQuorum returns a boolean if a field has been set.

### GetInputRule

`func (o *UpdateMcpSecretOAuthAuthCode) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *UpdateMcpSecretOAuthAuthCode) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *UpdateMcpSecretOAuthAuthCode) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetJson

`func (o *UpdateMcpSecretOAuthAuthCode) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateMcpSecretOAuthAuthCode) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateMcpSecretOAuthAuthCode) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateMcpSecretOAuthAuthCode) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateMcpSecretOAuthAuthCode) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateMcpSecretOAuthAuthCode) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateMcpSecretOAuthAuthCode) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateMcpSecretOAuthAuthCode) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateMcpSecretOAuthAuthCode) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastVersion

`func (o *UpdateMcpSecretOAuthAuthCode) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *UpdateMcpSecretOAuthAuthCode) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *UpdateMcpSecretOAuthAuthCode) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### GetName

`func (o *UpdateMcpSecretOAuthAuthCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMcpSecretOAuthAuthCode) SetName(v string)`

SetName sets Name field to given value.


### GetOauthClientId

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *UpdateMcpSecretOAuthAuthCode) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *UpdateMcpSecretOAuthAuthCode) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### GetOauthClientSecret

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *UpdateMcpSecretOAuthAuthCode) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *UpdateMcpSecretOAuthAuthCode) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### GetOauthRedirectUri

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthRedirectUri() string`

GetOauthRedirectUri returns the OauthRedirectUri field if non-nil, zero value otherwise.

### GetOauthRedirectUriOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthRedirectUriOk() (*string, bool)`

GetOauthRedirectUriOk returns a tuple with the OauthRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthRedirectUri

`func (o *UpdateMcpSecretOAuthAuthCode) SetOauthRedirectUri(v string)`

SetOauthRedirectUri sets OauthRedirectUri field to given value.

### HasOauthRedirectUri

`func (o *UpdateMcpSecretOAuthAuthCode) HasOauthRedirectUri() bool`

HasOauthRedirectUri returns a boolean if a field has been set.

### GetOauthRefreshToken

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthRefreshToken() string`

GetOauthRefreshToken returns the OauthRefreshToken field if non-nil, zero value otherwise.

### GetOauthRefreshTokenOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthRefreshTokenOk() (*string, bool)`

GetOauthRefreshTokenOk returns a tuple with the OauthRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthRefreshToken

`func (o *UpdateMcpSecretOAuthAuthCode) SetOauthRefreshToken(v string)`

SetOauthRefreshToken sets OauthRefreshToken field to given value.

### HasOauthRefreshToken

`func (o *UpdateMcpSecretOAuthAuthCode) HasOauthRefreshToken() bool`

HasOauthRefreshToken returns a boolean if a field has been set.

### GetOauthScopes

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthScopes() []string`

GetOauthScopes returns the OauthScopes field if non-nil, zero value otherwise.

### GetOauthScopesOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthScopesOk() (*[]string, bool)`

GetOauthScopesOk returns a tuple with the OauthScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthScopes

`func (o *UpdateMcpSecretOAuthAuthCode) SetOauthScopes(v []string)`

SetOauthScopes sets OauthScopes field to given value.

### HasOauthScopes

`func (o *UpdateMcpSecretOAuthAuthCode) HasOauthScopes() bool`

HasOauthScopes returns a boolean if a field has been set.

### GetOauthTokenUrl

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthTokenUrl() string`

GetOauthTokenUrl returns the OauthTokenUrl field if non-nil, zero value otherwise.

### GetOauthTokenUrlOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetOauthTokenUrlOk() (*string, bool)`

GetOauthTokenUrlOk returns a tuple with the OauthTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenUrl

`func (o *UpdateMcpSecretOAuthAuthCode) SetOauthTokenUrl(v string)`

SetOauthTokenUrl sets OauthTokenUrl field to given value.

### HasOauthTokenUrl

`func (o *UpdateMcpSecretOAuthAuthCode) HasOauthTokenUrl() bool`

HasOauthTokenUrl returns a boolean if a field has been set.

### GetOutputRule

`func (o *UpdateMcpSecretOAuthAuthCode) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *UpdateMcpSecretOAuthAuthCode) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *UpdateMcpSecretOAuthAuthCode) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetToken

`func (o *UpdateMcpSecretOAuthAuthCode) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateMcpSecretOAuthAuthCode) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateMcpSecretOAuthAuthCode) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateMcpSecretOAuthAuthCode) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateMcpSecretOAuthAuthCode) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateMcpSecretOAuthAuthCode) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateMcpSecretOAuthAuthCode) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateMcpSecretOAuthAuthCode) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateMcpSecretOAuthAuthCode) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateMcpSecretOAuthAuthCode) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


