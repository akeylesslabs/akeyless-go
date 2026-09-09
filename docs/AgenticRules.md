# AgenticRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled is a pointer so rules persisted before this field existed (nil) keep enforcing, rather than silently switching off. | [optional] 
**InputRules** | Pointer to [**[]AgenticRule**](AgenticRule.md) |  | [optional] 
**OutputRules** | Pointer to [**[]AgenticRule**](AgenticRule.md) |  | [optional] 
**QuorumEnabled** | Pointer to **bool** | QuorumEnabled asks for this item&#39;s policy decisions to be evaluated by every model configured on the gateway rather than the Default alone.  Also a pointer, but with the opposite nil meaning to Enabled above: nil is OFF. Enabled defaults on because it governs rules that were already being enforced before the field existed, whereas quorum is new behavior that multiplies latency and denies fail-closed - an item that never asked for it must not acquire it by upgrade. | [optional] 

## Methods

### NewAgenticRules

`func NewAgenticRules() *AgenticRules`

NewAgenticRules instantiates a new AgenticRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticRulesWithDefaults

`func NewAgenticRulesWithDefaults() *AgenticRules`

NewAgenticRulesWithDefaults instantiates a new AgenticRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AgenticRules) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgenticRules) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgenticRules) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgenticRules) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInputRules

`func (o *AgenticRules) GetInputRules() []AgenticRule`

GetInputRules returns the InputRules field if non-nil, zero value otherwise.

### GetInputRulesOk

`func (o *AgenticRules) GetInputRulesOk() (*[]AgenticRule, bool)`

GetInputRulesOk returns a tuple with the InputRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRules

`func (o *AgenticRules) SetInputRules(v []AgenticRule)`

SetInputRules sets InputRules field to given value.

### HasInputRules

`func (o *AgenticRules) HasInputRules() bool`

HasInputRules returns a boolean if a field has been set.

### GetOutputRules

`func (o *AgenticRules) GetOutputRules() []AgenticRule`

GetOutputRules returns the OutputRules field if non-nil, zero value otherwise.

### GetOutputRulesOk

`func (o *AgenticRules) GetOutputRulesOk() (*[]AgenticRule, bool)`

GetOutputRulesOk returns a tuple with the OutputRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRules

`func (o *AgenticRules) SetOutputRules(v []AgenticRule)`

SetOutputRules sets OutputRules field to given value.

### HasOutputRules

`func (o *AgenticRules) HasOutputRules() bool`

HasOutputRules returns a boolean if a field has been set.

### GetQuorumEnabled

`func (o *AgenticRules) GetQuorumEnabled() bool`

GetQuorumEnabled returns the QuorumEnabled field if non-nil, zero value otherwise.

### GetQuorumEnabledOk

`func (o *AgenticRules) GetQuorumEnabledOk() (*bool, bool)`

GetQuorumEnabledOk returns a tuple with the QuorumEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorumEnabled

`func (o *AgenticRules) SetQuorumEnabled(v bool)`

SetQuorumEnabled sets QuorumEnabled field to given value.

### HasQuorumEnabled

`func (o *AgenticRules) HasQuorumEnabled() bool`

HasQuorumEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


