# AnthropicTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnthropicUrl** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** | UseGatewayCloudIdentity authenticates with Anthropic Workload Identity Federation instead of ApiKey: the Gateway exchanges its own identity token for a short-lived Anthropic access token. ApiKey must be empty. See go/src/targets/llm/anthropic_wif_client.go. | [optional] 
**WifAzureAudience** | Pointer to **string** | WIFAzureAudience is the identifier URI (api://&lt;app-id&gt;) of the Entra app registration that represents the Claude API audience. Required with the azure source, empty with every other source. | [optional] 
**WifFederationRuleId** | Pointer to **string** | WIFFederationRuleID is the federation rule ID (fdrl_...). | [optional] 
**WifIdentityTokenSource** | Pointer to **string** | WIFIdentityTokenSource selects where the Gateway takes its identity token: file (default, projected Kubernetes or SPIFFE token), gcp (Google Cloud metadata server), aws (STS web identity token) or azure (Microsoft Entra ID). | [optional] 
**WifOrganizationId** | Pointer to **string** | WIFOrganizationID is the UUID of the Anthropic organization. | [optional] 
**WifServiceAccountId** | Pointer to **string** | WIFServiceAccountID is the Claude service account ID (svac_...). | [optional] 
**WifWorkspaceId** | Pointer to **string** | WIFWorkspaceID is the optional workspace ID (wrkspc_...). Required by Anthropic only when the rule is enabled for more than one workspace. | [optional] 

## Methods

### NewAnthropicTargetDetails

`func NewAnthropicTargetDetails() *AnthropicTargetDetails`

NewAnthropicTargetDetails instantiates a new AnthropicTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicTargetDetailsWithDefaults

`func NewAnthropicTargetDetailsWithDefaults() *AnthropicTargetDetails`

NewAnthropicTargetDetailsWithDefaults instantiates a new AnthropicTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnthropicUrl

`func (o *AnthropicTargetDetails) GetAnthropicUrl() string`

GetAnthropicUrl returns the AnthropicUrl field if non-nil, zero value otherwise.

### GetAnthropicUrlOk

`func (o *AnthropicTargetDetails) GetAnthropicUrlOk() (*string, bool)`

GetAnthropicUrlOk returns a tuple with the AnthropicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnthropicUrl

`func (o *AnthropicTargetDetails) SetAnthropicUrl(v string)`

SetAnthropicUrl sets AnthropicUrl field to given value.

### HasAnthropicUrl

`func (o *AnthropicTargetDetails) HasAnthropicUrl() bool`

HasAnthropicUrl returns a boolean if a field has been set.

### GetApiKey

`func (o *AnthropicTargetDetails) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AnthropicTargetDetails) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AnthropicTargetDetails) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AnthropicTargetDetails) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *AnthropicTargetDetails) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *AnthropicTargetDetails) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *AnthropicTargetDetails) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *AnthropicTargetDetails) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.

### GetWifAzureAudience

`func (o *AnthropicTargetDetails) GetWifAzureAudience() string`

GetWifAzureAudience returns the WifAzureAudience field if non-nil, zero value otherwise.

### GetWifAzureAudienceOk

`func (o *AnthropicTargetDetails) GetWifAzureAudienceOk() (*string, bool)`

GetWifAzureAudienceOk returns a tuple with the WifAzureAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifAzureAudience

`func (o *AnthropicTargetDetails) SetWifAzureAudience(v string)`

SetWifAzureAudience sets WifAzureAudience field to given value.

### HasWifAzureAudience

`func (o *AnthropicTargetDetails) HasWifAzureAudience() bool`

HasWifAzureAudience returns a boolean if a field has been set.

### GetWifFederationRuleId

`func (o *AnthropicTargetDetails) GetWifFederationRuleId() string`

GetWifFederationRuleId returns the WifFederationRuleId field if non-nil, zero value otherwise.

### GetWifFederationRuleIdOk

`func (o *AnthropicTargetDetails) GetWifFederationRuleIdOk() (*string, bool)`

GetWifFederationRuleIdOk returns a tuple with the WifFederationRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifFederationRuleId

`func (o *AnthropicTargetDetails) SetWifFederationRuleId(v string)`

SetWifFederationRuleId sets WifFederationRuleId field to given value.

### HasWifFederationRuleId

`func (o *AnthropicTargetDetails) HasWifFederationRuleId() bool`

HasWifFederationRuleId returns a boolean if a field has been set.

### GetWifIdentityTokenSource

`func (o *AnthropicTargetDetails) GetWifIdentityTokenSource() string`

GetWifIdentityTokenSource returns the WifIdentityTokenSource field if non-nil, zero value otherwise.

### GetWifIdentityTokenSourceOk

`func (o *AnthropicTargetDetails) GetWifIdentityTokenSourceOk() (*string, bool)`

GetWifIdentityTokenSourceOk returns a tuple with the WifIdentityTokenSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifIdentityTokenSource

`func (o *AnthropicTargetDetails) SetWifIdentityTokenSource(v string)`

SetWifIdentityTokenSource sets WifIdentityTokenSource field to given value.

### HasWifIdentityTokenSource

`func (o *AnthropicTargetDetails) HasWifIdentityTokenSource() bool`

HasWifIdentityTokenSource returns a boolean if a field has been set.

### GetWifOrganizationId

`func (o *AnthropicTargetDetails) GetWifOrganizationId() string`

GetWifOrganizationId returns the WifOrganizationId field if non-nil, zero value otherwise.

### GetWifOrganizationIdOk

`func (o *AnthropicTargetDetails) GetWifOrganizationIdOk() (*string, bool)`

GetWifOrganizationIdOk returns a tuple with the WifOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifOrganizationId

`func (o *AnthropicTargetDetails) SetWifOrganizationId(v string)`

SetWifOrganizationId sets WifOrganizationId field to given value.

### HasWifOrganizationId

`func (o *AnthropicTargetDetails) HasWifOrganizationId() bool`

HasWifOrganizationId returns a boolean if a field has been set.

### GetWifServiceAccountId

`func (o *AnthropicTargetDetails) GetWifServiceAccountId() string`

GetWifServiceAccountId returns the WifServiceAccountId field if non-nil, zero value otherwise.

### GetWifServiceAccountIdOk

`func (o *AnthropicTargetDetails) GetWifServiceAccountIdOk() (*string, bool)`

GetWifServiceAccountIdOk returns a tuple with the WifServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifServiceAccountId

`func (o *AnthropicTargetDetails) SetWifServiceAccountId(v string)`

SetWifServiceAccountId sets WifServiceAccountId field to given value.

### HasWifServiceAccountId

`func (o *AnthropicTargetDetails) HasWifServiceAccountId() bool`

HasWifServiceAccountId returns a boolean if a field has been set.

### GetWifWorkspaceId

`func (o *AnthropicTargetDetails) GetWifWorkspaceId() string`

GetWifWorkspaceId returns the WifWorkspaceId field if non-nil, zero value otherwise.

### GetWifWorkspaceIdOk

`func (o *AnthropicTargetDetails) GetWifWorkspaceIdOk() (*string, bool)`

GetWifWorkspaceIdOk returns a tuple with the WifWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifWorkspaceId

`func (o *AnthropicTargetDetails) SetWifWorkspaceId(v string)`

SetWifWorkspaceId sets WifWorkspaceId field to given value.

### HasWifWorkspaceId

`func (o *AnthropicTargetDetails) HasWifWorkspaceId() bool`

HasWifWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


