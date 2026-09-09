# GatewayRevokeTmpUsers

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

### NewGatewayRevokeTmpUsers

`func NewGatewayRevokeTmpUsers(name string, ) *GatewayRevokeTmpUsers`

NewGatewayRevokeTmpUsers instantiates a new GatewayRevokeTmpUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRevokeTmpUsersWithDefaults

`func NewGatewayRevokeTmpUsersWithDefaults() *GatewayRevokeTmpUsers`

NewGatewayRevokeTmpUsersWithDefaults instantiates a new GatewayRevokeTmpUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAraEnabled

`func (o *GatewayRevokeTmpUsers) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *GatewayRevokeTmpUsers) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *GatewayRevokeTmpUsers) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *GatewayRevokeTmpUsers) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetEnableAgenticRuntimeAuthority

`func (o *GatewayRevokeTmpUsers) GetEnableAgenticRuntimeAuthority() bool`

GetEnableAgenticRuntimeAuthority returns the EnableAgenticRuntimeAuthority field if non-nil, zero value otherwise.

### GetEnableAgenticRuntimeAuthorityOk

`func (o *GatewayRevokeTmpUsers) GetEnableAgenticRuntimeAuthorityOk() (*bool, bool)`

GetEnableAgenticRuntimeAuthorityOk returns a tuple with the EnableAgenticRuntimeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAgenticRuntimeAuthority

`func (o *GatewayRevokeTmpUsers) SetEnableAgenticRuntimeAuthority(v bool)`

SetEnableAgenticRuntimeAuthority sets EnableAgenticRuntimeAuthority field to given value.

### HasEnableAgenticRuntimeAuthority

`func (o *GatewayRevokeTmpUsers) HasEnableAgenticRuntimeAuthority() bool`

HasEnableAgenticRuntimeAuthority returns a boolean if a field has been set.

### GetEnableAiQuorum

`func (o *GatewayRevokeTmpUsers) GetEnableAiQuorum() bool`

GetEnableAiQuorum returns the EnableAiQuorum field if non-nil, zero value otherwise.

### GetEnableAiQuorumOk

`func (o *GatewayRevokeTmpUsers) GetEnableAiQuorumOk() (*bool, bool)`

GetEnableAiQuorumOk returns a tuple with the EnableAiQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiQuorum

`func (o *GatewayRevokeTmpUsers) SetEnableAiQuorum(v bool)`

SetEnableAiQuorum sets EnableAiQuorum field to given value.

### HasEnableAiQuorum

`func (o *GatewayRevokeTmpUsers) HasEnableAiQuorum() bool`

HasEnableAiQuorum returns a boolean if a field has been set.

### GetHost

`func (o *GatewayRevokeTmpUsers) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GatewayRevokeTmpUsers) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GatewayRevokeTmpUsers) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GatewayRevokeTmpUsers) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInputRule

`func (o *GatewayRevokeTmpUsers) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *GatewayRevokeTmpUsers) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *GatewayRevokeTmpUsers) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *GatewayRevokeTmpUsers) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetJson

`func (o *GatewayRevokeTmpUsers) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayRevokeTmpUsers) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayRevokeTmpUsers) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayRevokeTmpUsers) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayRevokeTmpUsers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayRevokeTmpUsers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayRevokeTmpUsers) SetName(v string)`

SetName sets Name field to given value.


### GetOutputRule

`func (o *GatewayRevokeTmpUsers) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *GatewayRevokeTmpUsers) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *GatewayRevokeTmpUsers) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *GatewayRevokeTmpUsers) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetRevokeAll

`func (o *GatewayRevokeTmpUsers) GetRevokeAll() bool`

GetRevokeAll returns the RevokeAll field if non-nil, zero value otherwise.

### GetRevokeAllOk

`func (o *GatewayRevokeTmpUsers) GetRevokeAllOk() (*bool, bool)`

GetRevokeAllOk returns a tuple with the RevokeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeAll

`func (o *GatewayRevokeTmpUsers) SetRevokeAll(v bool)`

SetRevokeAll sets RevokeAll field to given value.

### HasRevokeAll

`func (o *GatewayRevokeTmpUsers) HasRevokeAll() bool`

HasRevokeAll returns a boolean if a field has been set.

### GetSkipDryRun

`func (o *GatewayRevokeTmpUsers) GetSkipDryRun() string`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *GatewayRevokeTmpUsers) GetSkipDryRunOk() (*string, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *GatewayRevokeTmpUsers) SetSkipDryRun(v string)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *GatewayRevokeTmpUsers) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetSoftDelete

`func (o *GatewayRevokeTmpUsers) GetSoftDelete() bool`

GetSoftDelete returns the SoftDelete field if non-nil, zero value otherwise.

### GetSoftDeleteOk

`func (o *GatewayRevokeTmpUsers) GetSoftDeleteOk() (*bool, bool)`

GetSoftDeleteOk returns a tuple with the SoftDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDelete

`func (o *GatewayRevokeTmpUsers) SetSoftDelete(v bool)`

SetSoftDelete sets SoftDelete field to given value.

### HasSoftDelete

`func (o *GatewayRevokeTmpUsers) HasSoftDelete() bool`

HasSoftDelete returns a boolean if a field has been set.

### GetTmpCredsId

`func (o *GatewayRevokeTmpUsers) GetTmpCredsId() string`

GetTmpCredsId returns the TmpCredsId field if non-nil, zero value otherwise.

### GetTmpCredsIdOk

`func (o *GatewayRevokeTmpUsers) GetTmpCredsIdOk() (*string, bool)`

GetTmpCredsIdOk returns a tuple with the TmpCredsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpCredsId

`func (o *GatewayRevokeTmpUsers) SetTmpCredsId(v string)`

SetTmpCredsId sets TmpCredsId field to given value.

### HasTmpCredsId

`func (o *GatewayRevokeTmpUsers) HasTmpCredsId() bool`

HasTmpCredsId returns a boolean if a field has been set.

### GetToken

`func (o *GatewayRevokeTmpUsers) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayRevokeTmpUsers) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayRevokeTmpUsers) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayRevokeTmpUsers) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayRevokeTmpUsers) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayRevokeTmpUsers) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayRevokeTmpUsers) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayRevokeTmpUsers) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


