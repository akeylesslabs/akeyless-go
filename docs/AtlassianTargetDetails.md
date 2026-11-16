# AtlassianTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlassianAuthType** | Pointer to **string** |  | [optional] 
**AtlassianAutomationKey** | Pointer to **string** | AtlassianAutomationKey is an optional Automation REST API key (Jira settings &gt; System &gt; Automation &gt; API keys). Atlassian&#39;s automation API refuses an account API token, so without this key the scanner reports the automation rules as not read. | [optional] 
**AtlassianCaCertificate** | Pointer to **string** | AtlassianCACertificate is a base64-encoded PEM CA certificate for a Data Center site signed by a private CA. | [optional] 
**AtlassianClientId** | Pointer to **string** |  | [optional] 
**AtlassianClientSecret** | Pointer to **string** |  | [optional] 
**AtlassianCloudId** | Pointer to **string** |  | [optional] 
**AtlassianConfluenceToken** | Pointer to **string** |  | [optional] 
**AtlassianConfluenceUrl** | Pointer to **string** |  | [optional] 
**AtlassianDeploymentType** | Pointer to **string** | When empty, the deployment type, auth type and products are inferred from the fields above, so targets saved before these fields existed keep working. | [optional] 
**AtlassianEmail** | Pointer to **string** |  | [optional] 
**AtlassianJiraToken** | Pointer to **string** |  | [optional] 
**AtlassianJiraUrl** | Pointer to **string** |  | [optional] 
**AtlassianProducts** | Pointer to **[]string** |  | [optional] 
**AtlassianToken** | Pointer to **string** |  | [optional] 
**AtlassianUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewAtlassianTargetDetails

`func NewAtlassianTargetDetails() *AtlassianTargetDetails`

NewAtlassianTargetDetails instantiates a new AtlassianTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlassianTargetDetailsWithDefaults

`func NewAtlassianTargetDetailsWithDefaults() *AtlassianTargetDetails`

NewAtlassianTargetDetailsWithDefaults instantiates a new AtlassianTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlassianAuthType

`func (o *AtlassianTargetDetails) GetAtlassianAuthType() string`

GetAtlassianAuthType returns the AtlassianAuthType field if non-nil, zero value otherwise.

### GetAtlassianAuthTypeOk

`func (o *AtlassianTargetDetails) GetAtlassianAuthTypeOk() (*string, bool)`

GetAtlassianAuthTypeOk returns a tuple with the AtlassianAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianAuthType

`func (o *AtlassianTargetDetails) SetAtlassianAuthType(v string)`

SetAtlassianAuthType sets AtlassianAuthType field to given value.

### HasAtlassianAuthType

`func (o *AtlassianTargetDetails) HasAtlassianAuthType() bool`

HasAtlassianAuthType returns a boolean if a field has been set.

### GetAtlassianAutomationKey

`func (o *AtlassianTargetDetails) GetAtlassianAutomationKey() string`

GetAtlassianAutomationKey returns the AtlassianAutomationKey field if non-nil, zero value otherwise.

### GetAtlassianAutomationKeyOk

`func (o *AtlassianTargetDetails) GetAtlassianAutomationKeyOk() (*string, bool)`

GetAtlassianAutomationKeyOk returns a tuple with the AtlassianAutomationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianAutomationKey

`func (o *AtlassianTargetDetails) SetAtlassianAutomationKey(v string)`

SetAtlassianAutomationKey sets AtlassianAutomationKey field to given value.

### HasAtlassianAutomationKey

`func (o *AtlassianTargetDetails) HasAtlassianAutomationKey() bool`

HasAtlassianAutomationKey returns a boolean if a field has been set.

### GetAtlassianCaCertificate

`func (o *AtlassianTargetDetails) GetAtlassianCaCertificate() string`

GetAtlassianCaCertificate returns the AtlassianCaCertificate field if non-nil, zero value otherwise.

### GetAtlassianCaCertificateOk

`func (o *AtlassianTargetDetails) GetAtlassianCaCertificateOk() (*string, bool)`

GetAtlassianCaCertificateOk returns a tuple with the AtlassianCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianCaCertificate

`func (o *AtlassianTargetDetails) SetAtlassianCaCertificate(v string)`

SetAtlassianCaCertificate sets AtlassianCaCertificate field to given value.

### HasAtlassianCaCertificate

`func (o *AtlassianTargetDetails) HasAtlassianCaCertificate() bool`

HasAtlassianCaCertificate returns a boolean if a field has been set.

### GetAtlassianClientId

`func (o *AtlassianTargetDetails) GetAtlassianClientId() string`

GetAtlassianClientId returns the AtlassianClientId field if non-nil, zero value otherwise.

### GetAtlassianClientIdOk

`func (o *AtlassianTargetDetails) GetAtlassianClientIdOk() (*string, bool)`

GetAtlassianClientIdOk returns a tuple with the AtlassianClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianClientId

`func (o *AtlassianTargetDetails) SetAtlassianClientId(v string)`

SetAtlassianClientId sets AtlassianClientId field to given value.

### HasAtlassianClientId

`func (o *AtlassianTargetDetails) HasAtlassianClientId() bool`

HasAtlassianClientId returns a boolean if a field has been set.

### GetAtlassianClientSecret

`func (o *AtlassianTargetDetails) GetAtlassianClientSecret() string`

GetAtlassianClientSecret returns the AtlassianClientSecret field if non-nil, zero value otherwise.

### GetAtlassianClientSecretOk

`func (o *AtlassianTargetDetails) GetAtlassianClientSecretOk() (*string, bool)`

GetAtlassianClientSecretOk returns a tuple with the AtlassianClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianClientSecret

`func (o *AtlassianTargetDetails) SetAtlassianClientSecret(v string)`

SetAtlassianClientSecret sets AtlassianClientSecret field to given value.

### HasAtlassianClientSecret

`func (o *AtlassianTargetDetails) HasAtlassianClientSecret() bool`

HasAtlassianClientSecret returns a boolean if a field has been set.

### GetAtlassianCloudId

`func (o *AtlassianTargetDetails) GetAtlassianCloudId() string`

GetAtlassianCloudId returns the AtlassianCloudId field if non-nil, zero value otherwise.

### GetAtlassianCloudIdOk

`func (o *AtlassianTargetDetails) GetAtlassianCloudIdOk() (*string, bool)`

GetAtlassianCloudIdOk returns a tuple with the AtlassianCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianCloudId

`func (o *AtlassianTargetDetails) SetAtlassianCloudId(v string)`

SetAtlassianCloudId sets AtlassianCloudId field to given value.

### HasAtlassianCloudId

`func (o *AtlassianTargetDetails) HasAtlassianCloudId() bool`

HasAtlassianCloudId returns a boolean if a field has been set.

### GetAtlassianConfluenceToken

`func (o *AtlassianTargetDetails) GetAtlassianConfluenceToken() string`

GetAtlassianConfluenceToken returns the AtlassianConfluenceToken field if non-nil, zero value otherwise.

### GetAtlassianConfluenceTokenOk

`func (o *AtlassianTargetDetails) GetAtlassianConfluenceTokenOk() (*string, bool)`

GetAtlassianConfluenceTokenOk returns a tuple with the AtlassianConfluenceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianConfluenceToken

`func (o *AtlassianTargetDetails) SetAtlassianConfluenceToken(v string)`

SetAtlassianConfluenceToken sets AtlassianConfluenceToken field to given value.

### HasAtlassianConfluenceToken

`func (o *AtlassianTargetDetails) HasAtlassianConfluenceToken() bool`

HasAtlassianConfluenceToken returns a boolean if a field has been set.

### GetAtlassianConfluenceUrl

`func (o *AtlassianTargetDetails) GetAtlassianConfluenceUrl() string`

GetAtlassianConfluenceUrl returns the AtlassianConfluenceUrl field if non-nil, zero value otherwise.

### GetAtlassianConfluenceUrlOk

`func (o *AtlassianTargetDetails) GetAtlassianConfluenceUrlOk() (*string, bool)`

GetAtlassianConfluenceUrlOk returns a tuple with the AtlassianConfluenceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianConfluenceUrl

`func (o *AtlassianTargetDetails) SetAtlassianConfluenceUrl(v string)`

SetAtlassianConfluenceUrl sets AtlassianConfluenceUrl field to given value.

### HasAtlassianConfluenceUrl

`func (o *AtlassianTargetDetails) HasAtlassianConfluenceUrl() bool`

HasAtlassianConfluenceUrl returns a boolean if a field has been set.

### GetAtlassianDeploymentType

`func (o *AtlassianTargetDetails) GetAtlassianDeploymentType() string`

GetAtlassianDeploymentType returns the AtlassianDeploymentType field if non-nil, zero value otherwise.

### GetAtlassianDeploymentTypeOk

`func (o *AtlassianTargetDetails) GetAtlassianDeploymentTypeOk() (*string, bool)`

GetAtlassianDeploymentTypeOk returns a tuple with the AtlassianDeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianDeploymentType

`func (o *AtlassianTargetDetails) SetAtlassianDeploymentType(v string)`

SetAtlassianDeploymentType sets AtlassianDeploymentType field to given value.

### HasAtlassianDeploymentType

`func (o *AtlassianTargetDetails) HasAtlassianDeploymentType() bool`

HasAtlassianDeploymentType returns a boolean if a field has been set.

### GetAtlassianEmail

`func (o *AtlassianTargetDetails) GetAtlassianEmail() string`

GetAtlassianEmail returns the AtlassianEmail field if non-nil, zero value otherwise.

### GetAtlassianEmailOk

`func (o *AtlassianTargetDetails) GetAtlassianEmailOk() (*string, bool)`

GetAtlassianEmailOk returns a tuple with the AtlassianEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianEmail

`func (o *AtlassianTargetDetails) SetAtlassianEmail(v string)`

SetAtlassianEmail sets AtlassianEmail field to given value.

### HasAtlassianEmail

`func (o *AtlassianTargetDetails) HasAtlassianEmail() bool`

HasAtlassianEmail returns a boolean if a field has been set.

### GetAtlassianJiraToken

`func (o *AtlassianTargetDetails) GetAtlassianJiraToken() string`

GetAtlassianJiraToken returns the AtlassianJiraToken field if non-nil, zero value otherwise.

### GetAtlassianJiraTokenOk

`func (o *AtlassianTargetDetails) GetAtlassianJiraTokenOk() (*string, bool)`

GetAtlassianJiraTokenOk returns a tuple with the AtlassianJiraToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianJiraToken

`func (o *AtlassianTargetDetails) SetAtlassianJiraToken(v string)`

SetAtlassianJiraToken sets AtlassianJiraToken field to given value.

### HasAtlassianJiraToken

`func (o *AtlassianTargetDetails) HasAtlassianJiraToken() bool`

HasAtlassianJiraToken returns a boolean if a field has been set.

### GetAtlassianJiraUrl

`func (o *AtlassianTargetDetails) GetAtlassianJiraUrl() string`

GetAtlassianJiraUrl returns the AtlassianJiraUrl field if non-nil, zero value otherwise.

### GetAtlassianJiraUrlOk

`func (o *AtlassianTargetDetails) GetAtlassianJiraUrlOk() (*string, bool)`

GetAtlassianJiraUrlOk returns a tuple with the AtlassianJiraUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianJiraUrl

`func (o *AtlassianTargetDetails) SetAtlassianJiraUrl(v string)`

SetAtlassianJiraUrl sets AtlassianJiraUrl field to given value.

### HasAtlassianJiraUrl

`func (o *AtlassianTargetDetails) HasAtlassianJiraUrl() bool`

HasAtlassianJiraUrl returns a boolean if a field has been set.

### GetAtlassianProducts

`func (o *AtlassianTargetDetails) GetAtlassianProducts() []string`

GetAtlassianProducts returns the AtlassianProducts field if non-nil, zero value otherwise.

### GetAtlassianProductsOk

`func (o *AtlassianTargetDetails) GetAtlassianProductsOk() (*[]string, bool)`

GetAtlassianProductsOk returns a tuple with the AtlassianProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianProducts

`func (o *AtlassianTargetDetails) SetAtlassianProducts(v []string)`

SetAtlassianProducts sets AtlassianProducts field to given value.

### HasAtlassianProducts

`func (o *AtlassianTargetDetails) HasAtlassianProducts() bool`

HasAtlassianProducts returns a boolean if a field has been set.

### GetAtlassianToken

`func (o *AtlassianTargetDetails) GetAtlassianToken() string`

GetAtlassianToken returns the AtlassianToken field if non-nil, zero value otherwise.

### GetAtlassianTokenOk

`func (o *AtlassianTargetDetails) GetAtlassianTokenOk() (*string, bool)`

GetAtlassianTokenOk returns a tuple with the AtlassianToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianToken

`func (o *AtlassianTargetDetails) SetAtlassianToken(v string)`

SetAtlassianToken sets AtlassianToken field to given value.

### HasAtlassianToken

`func (o *AtlassianTargetDetails) HasAtlassianToken() bool`

HasAtlassianToken returns a boolean if a field has been set.

### GetAtlassianUrl

`func (o *AtlassianTargetDetails) GetAtlassianUrl() string`

GetAtlassianUrl returns the AtlassianUrl field if non-nil, zero value otherwise.

### GetAtlassianUrlOk

`func (o *AtlassianTargetDetails) GetAtlassianUrlOk() (*string, bool)`

GetAtlassianUrlOk returns a tuple with the AtlassianUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianUrl

`func (o *AtlassianTargetDetails) SetAtlassianUrl(v string)`

SetAtlassianUrl sets AtlassianUrl field to given value.

### HasAtlassianUrl

`func (o *AtlassianTargetDetails) HasAtlassianUrl() bool`

HasAtlassianUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


