# AiModelEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** | Provider is display-only: it carries the provider label the model catalog reported for Model (meaningful for Bedrock, which fronts several providers) so the console can render its Provider column without re-fetching the whole catalog per row. Never used for credential resolution or validation - the target type is the authority there. | [optional] 
**Role** | Pointer to **string** | A gateway has exactly one AiModelRoleDefault entry whenever any entry exists: it serves AI Insight, ISI, and the standard policy decisions in Secretless AI / ARA. AiModelRoleQuorum entries are consulted alongside the default to double-check policy decisions before a risky action is allowed. | [optional] 
**TargetId** | Pointer to **int64** | TargetId uses the repo-wide spelling (see types.Target.TargetId and AiInsightsConfigPart.TargetId) rather than the Go-idiomatic TargetID. | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**TargetType** | Pointer to **string** | TargetType is display-only, captured at write time like Provider. The console shows it as a read-only column; resolving it per row from the targets list instead would mean an extra lookup that can disagree with what the entry was created against. | [optional] 

## Methods

### NewAiModelEntry

`func NewAiModelEntry() *AiModelEntry`

NewAiModelEntry instantiates a new AiModelEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelEntryWithDefaults

`func NewAiModelEntryWithDefaults() *AiModelEntry`

NewAiModelEntryWithDefaults instantiates a new AiModelEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiModelEntry) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiModelEntry) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiModelEntry) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiModelEntry) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProvider

`func (o *AiModelEntry) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiModelEntry) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiModelEntry) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AiModelEntry) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRole

`func (o *AiModelEntry) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AiModelEntry) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AiModelEntry) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AiModelEntry) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTargetId

`func (o *AiModelEntry) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AiModelEntry) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AiModelEntry) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AiModelEntry) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetName

`func (o *AiModelEntry) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *AiModelEntry) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *AiModelEntry) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *AiModelEntry) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetType

`func (o *AiModelEntry) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AiModelEntry) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AiModelEntry) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *AiModelEntry) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


