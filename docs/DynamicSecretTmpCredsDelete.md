# DynamicSecretTmpCredsDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. Mirrors commands.AgenticRulesParams.AraEnabled. | [optional] 
**EnableAgenticRuntimeAuthority** | Pointer to **bool** | EnableAra is the documented spelling of AraEnabled; --ara-enabled shipped first and stays as an undocumented alias. | [optional] 
**EnableAiQuorum** | Pointer to **bool** | Turns on AI Quorum checks for this item. | [optional] 
**Host** | Pointer to **string** | Host | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) Mirrors commands.AgenticRulesParams — kept separate because ResourceDS cannot embed it (different package, different struct layout). | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**RevokeAll** | Pointer to **bool** | Revoke All Temp Creds | [optional] 
**SkipDryRun** | Pointer to **string** | If set, dry-run will be skipped | [optional] 
**SoftDelete** | Pointer to **bool** | Soft Delete | [optional] 
**TmpCredsId** | Pointer to **string** | Tmp Creds ID | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDynamicSecretTmpCredsDelete

`func NewDynamicSecretTmpCredsDelete(name string, ) *DynamicSecretTmpCredsDelete`

NewDynamicSecretTmpCredsDelete instantiates a new DynamicSecretTmpCredsDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretTmpCredsDeleteWithDefaults

`func NewDynamicSecretTmpCredsDeleteWithDefaults() *DynamicSecretTmpCredsDelete`

NewDynamicSecretTmpCredsDeleteWithDefaults instantiates a new DynamicSecretTmpCredsDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAraEnabled

`func (o *DynamicSecretTmpCredsDelete) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *DynamicSecretTmpCredsDelete) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *DynamicSecretTmpCredsDelete) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *DynamicSecretTmpCredsDelete) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetEnableAgenticRuntimeAuthority

`func (o *DynamicSecretTmpCredsDelete) GetEnableAgenticRuntimeAuthority() bool`

GetEnableAgenticRuntimeAuthority returns the EnableAgenticRuntimeAuthority field if non-nil, zero value otherwise.

### GetEnableAgenticRuntimeAuthorityOk

`func (o *DynamicSecretTmpCredsDelete) GetEnableAgenticRuntimeAuthorityOk() (*bool, bool)`

GetEnableAgenticRuntimeAuthorityOk returns a tuple with the EnableAgenticRuntimeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAgenticRuntimeAuthority

`func (o *DynamicSecretTmpCredsDelete) SetEnableAgenticRuntimeAuthority(v bool)`

SetEnableAgenticRuntimeAuthority sets EnableAgenticRuntimeAuthority field to given value.

### HasEnableAgenticRuntimeAuthority

`func (o *DynamicSecretTmpCredsDelete) HasEnableAgenticRuntimeAuthority() bool`

HasEnableAgenticRuntimeAuthority returns a boolean if a field has been set.

### GetEnableAiQuorum

`func (o *DynamicSecretTmpCredsDelete) GetEnableAiQuorum() bool`

GetEnableAiQuorum returns the EnableAiQuorum field if non-nil, zero value otherwise.

### GetEnableAiQuorumOk

`func (o *DynamicSecretTmpCredsDelete) GetEnableAiQuorumOk() (*bool, bool)`

GetEnableAiQuorumOk returns a tuple with the EnableAiQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiQuorum

`func (o *DynamicSecretTmpCredsDelete) SetEnableAiQuorum(v bool)`

SetEnableAiQuorum sets EnableAiQuorum field to given value.

### HasEnableAiQuorum

`func (o *DynamicSecretTmpCredsDelete) HasEnableAiQuorum() bool`

HasEnableAiQuorum returns a boolean if a field has been set.

### GetHost

`func (o *DynamicSecretTmpCredsDelete) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DynamicSecretTmpCredsDelete) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DynamicSecretTmpCredsDelete) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DynamicSecretTmpCredsDelete) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInputRule

`func (o *DynamicSecretTmpCredsDelete) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *DynamicSecretTmpCredsDelete) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *DynamicSecretTmpCredsDelete) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *DynamicSecretTmpCredsDelete) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretTmpCredsDelete) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretTmpCredsDelete) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretTmpCredsDelete) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretTmpCredsDelete) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretTmpCredsDelete) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretTmpCredsDelete) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretTmpCredsDelete) SetName(v string)`

SetName sets Name field to given value.


### GetOutputRule

`func (o *DynamicSecretTmpCredsDelete) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *DynamicSecretTmpCredsDelete) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *DynamicSecretTmpCredsDelete) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *DynamicSecretTmpCredsDelete) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetRevokeAll

`func (o *DynamicSecretTmpCredsDelete) GetRevokeAll() bool`

GetRevokeAll returns the RevokeAll field if non-nil, zero value otherwise.

### GetRevokeAllOk

`func (o *DynamicSecretTmpCredsDelete) GetRevokeAllOk() (*bool, bool)`

GetRevokeAllOk returns a tuple with the RevokeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeAll

`func (o *DynamicSecretTmpCredsDelete) SetRevokeAll(v bool)`

SetRevokeAll sets RevokeAll field to given value.

### HasRevokeAll

`func (o *DynamicSecretTmpCredsDelete) HasRevokeAll() bool`

HasRevokeAll returns a boolean if a field has been set.

### GetSkipDryRun

`func (o *DynamicSecretTmpCredsDelete) GetSkipDryRun() string`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *DynamicSecretTmpCredsDelete) GetSkipDryRunOk() (*string, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *DynamicSecretTmpCredsDelete) SetSkipDryRun(v string)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *DynamicSecretTmpCredsDelete) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetSoftDelete

`func (o *DynamicSecretTmpCredsDelete) GetSoftDelete() bool`

GetSoftDelete returns the SoftDelete field if non-nil, zero value otherwise.

### GetSoftDeleteOk

`func (o *DynamicSecretTmpCredsDelete) GetSoftDeleteOk() (*bool, bool)`

GetSoftDeleteOk returns a tuple with the SoftDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDelete

`func (o *DynamicSecretTmpCredsDelete) SetSoftDelete(v bool)`

SetSoftDelete sets SoftDelete field to given value.

### HasSoftDelete

`func (o *DynamicSecretTmpCredsDelete) HasSoftDelete() bool`

HasSoftDelete returns a boolean if a field has been set.

### GetTmpCredsId

`func (o *DynamicSecretTmpCredsDelete) GetTmpCredsId() string`

GetTmpCredsId returns the TmpCredsId field if non-nil, zero value otherwise.

### GetTmpCredsIdOk

`func (o *DynamicSecretTmpCredsDelete) GetTmpCredsIdOk() (*string, bool)`

GetTmpCredsIdOk returns a tuple with the TmpCredsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpCredsId

`func (o *DynamicSecretTmpCredsDelete) SetTmpCredsId(v string)`

SetTmpCredsId sets TmpCredsId field to given value.

### HasTmpCredsId

`func (o *DynamicSecretTmpCredsDelete) HasTmpCredsId() bool`

HasTmpCredsId returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretTmpCredsDelete) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretTmpCredsDelete) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretTmpCredsDelete) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretTmpCredsDelete) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretTmpCredsDelete) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretTmpCredsDelete) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretTmpCredsDelete) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretTmpCredsDelete) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


