# DynamicSecretTmpCredsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AraEnabled** | Pointer to **bool** | Enable or disable Agentic Runtime Authority rule enforcement for this item. Mirrors commands.AgenticRulesParams.AraEnabled. | [optional] 
**Host** | **string** | Host | 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) Mirrors commands.AgenticRulesParams — kept separate because ResourceDS cannot embed it (different package, different struct layout). | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewTtlMin** | **int64** | New TTL in Minutes | 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**SkipDryRun** | Pointer to **string** | If set, dry-run will be skipped | [optional] 
**TmpCredsId** | **string** | Tmp Creds ID | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDynamicSecretTmpCredsUpdate

`func NewDynamicSecretTmpCredsUpdate(host string, name string, newTtlMin int64, tmpCredsId string, ) *DynamicSecretTmpCredsUpdate`

NewDynamicSecretTmpCredsUpdate instantiates a new DynamicSecretTmpCredsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretTmpCredsUpdateWithDefaults

`func NewDynamicSecretTmpCredsUpdateWithDefaults() *DynamicSecretTmpCredsUpdate`

NewDynamicSecretTmpCredsUpdateWithDefaults instantiates a new DynamicSecretTmpCredsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAraEnabled

`func (o *DynamicSecretTmpCredsUpdate) GetAraEnabled() bool`

GetAraEnabled returns the AraEnabled field if non-nil, zero value otherwise.

### GetAraEnabledOk

`func (o *DynamicSecretTmpCredsUpdate) GetAraEnabledOk() (*bool, bool)`

GetAraEnabledOk returns a tuple with the AraEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAraEnabled

`func (o *DynamicSecretTmpCredsUpdate) SetAraEnabled(v bool)`

SetAraEnabled sets AraEnabled field to given value.

### HasAraEnabled

`func (o *DynamicSecretTmpCredsUpdate) HasAraEnabled() bool`

HasAraEnabled returns a boolean if a field has been set.

### GetHost

`func (o *DynamicSecretTmpCredsUpdate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DynamicSecretTmpCredsUpdate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DynamicSecretTmpCredsUpdate) SetHost(v string)`

SetHost sets Host field to given value.


### GetInputRule

`func (o *DynamicSecretTmpCredsUpdate) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *DynamicSecretTmpCredsUpdate) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *DynamicSecretTmpCredsUpdate) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *DynamicSecretTmpCredsUpdate) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretTmpCredsUpdate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretTmpCredsUpdate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretTmpCredsUpdate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretTmpCredsUpdate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretTmpCredsUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretTmpCredsUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretTmpCredsUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetNewTtlMin

`func (o *DynamicSecretTmpCredsUpdate) GetNewTtlMin() int64`

GetNewTtlMin returns the NewTtlMin field if non-nil, zero value otherwise.

### GetNewTtlMinOk

`func (o *DynamicSecretTmpCredsUpdate) GetNewTtlMinOk() (*int64, bool)`

GetNewTtlMinOk returns a tuple with the NewTtlMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTtlMin

`func (o *DynamicSecretTmpCredsUpdate) SetNewTtlMin(v int64)`

SetNewTtlMin sets NewTtlMin field to given value.


### GetOutputRule

`func (o *DynamicSecretTmpCredsUpdate) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *DynamicSecretTmpCredsUpdate) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *DynamicSecretTmpCredsUpdate) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *DynamicSecretTmpCredsUpdate) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetSkipDryRun

`func (o *DynamicSecretTmpCredsUpdate) GetSkipDryRun() string`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *DynamicSecretTmpCredsUpdate) GetSkipDryRunOk() (*string, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *DynamicSecretTmpCredsUpdate) SetSkipDryRun(v string)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *DynamicSecretTmpCredsUpdate) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetTmpCredsId

`func (o *DynamicSecretTmpCredsUpdate) GetTmpCredsId() string`

GetTmpCredsId returns the TmpCredsId field if non-nil, zero value otherwise.

### GetTmpCredsIdOk

`func (o *DynamicSecretTmpCredsUpdate) GetTmpCredsIdOk() (*string, bool)`

GetTmpCredsIdOk returns a tuple with the TmpCredsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpCredsId

`func (o *DynamicSecretTmpCredsUpdate) SetTmpCredsId(v string)`

SetTmpCredsId sets TmpCredsId field to given value.


### GetToken

`func (o *DynamicSecretTmpCredsUpdate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretTmpCredsUpdate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretTmpCredsUpdate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretTmpCredsUpdate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretTmpCredsUpdate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretTmpCredsUpdate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretTmpCredsUpdate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretTmpCredsUpdate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


