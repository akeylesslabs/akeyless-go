# AiInsightsConfigPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Models** | Pointer to [**[]AiModelEntry**](AiModelEntry.md) | Models holds every configured model, in whatever order and with whatever Default flag was stored - it is NOT canonicalized on write, so nothing may assume the Default sits at index 0. Empty on configs written before multi-model support. Never read it directly: use EffectiveModels for the list as stored (which also handles the legacy case), or PolicyModels for exactly one Default in row 1 followed by the Quorum models. | [optional] 
**TargetId** | Pointer to **int64** |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** | Version is an optimistic-concurrency token, bumped by gator on every accepted write.  Every mutation of this part is a read-modify-write across the network (the gateway reads the whole part, edits one entry, writes it back), and the write replaces the part wholesale. Without a token, two admins adding a quorum model at the same time silently lose one of the two - which, since the list must always carry exactly one Default, can also change which model serves every other AI feature.  Zero means \&quot;unversioned\&quot;: a client that predates this field, whose write gator accepts rather than rejecting outright. See updateGatewayAiInsightsConfig. | [optional] 

## Methods

### NewAiInsightsConfigPart

`func NewAiInsightsConfigPart() *AiInsightsConfigPart`

NewAiInsightsConfigPart instantiates a new AiInsightsConfigPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiInsightsConfigPartWithDefaults

`func NewAiInsightsConfigPartWithDefaults() *AiInsightsConfigPart`

NewAiInsightsConfigPartWithDefaults instantiates a new AiInsightsConfigPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AiInsightsConfigPart) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AiInsightsConfigPart) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AiInsightsConfigPart) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AiInsightsConfigPart) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetModel

`func (o *AiInsightsConfigPart) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiInsightsConfigPart) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiInsightsConfigPart) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiInsightsConfigPart) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModels

`func (o *AiInsightsConfigPart) GetModels() []AiModelEntry`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AiInsightsConfigPart) GetModelsOk() (*[]AiModelEntry, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AiInsightsConfigPart) SetModels(v []AiModelEntry)`

SetModels sets Models field to given value.

### HasModels

`func (o *AiInsightsConfigPart) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetTargetId

`func (o *AiInsightsConfigPart) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AiInsightsConfigPart) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AiInsightsConfigPart) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AiInsightsConfigPart) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetName

`func (o *AiInsightsConfigPart) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *AiInsightsConfigPart) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *AiInsightsConfigPart) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *AiInsightsConfigPart) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetVersion

`func (o *AiInsightsConfigPart) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AiInsightsConfigPart) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AiInsightsConfigPart) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AiInsightsConfigPart) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


