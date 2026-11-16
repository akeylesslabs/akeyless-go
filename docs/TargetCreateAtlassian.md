# TargetCreateAtlassian

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlassianAuthType** | Pointer to **string** | Atlassian auth type: oauth, service_account_token or user_api_token on cloud, pat on data center | [optional] 
**AtlassianAutomationKey** | Pointer to **string** | Optional Automation REST API key (Jira settings &gt; System &gt; Automation &gt; API keys); the automation API refuses an account token, so without it the automation rules are reported as not read | [optional] 
**AtlassianCaCertificate** | Pointer to **string** | CA certificate (base64-encoded PEM) of a Data Center site signed by a private CA | [optional] 
**AtlassianClientId** | Pointer to **string** | Atlassian OAuth 2.0 service account client ID | [optional] 
**AtlassianClientSecret** | Pointer to **string** | Atlassian OAuth 2.0 service account client secret | [optional] 
**AtlassianCloudId** | Pointer to **string** | Atlassian Cloud ID of the site (resolved from the site URL when empty) | [optional] 
**AtlassianConfluenceToken** | Pointer to **string** | Confluence Data Center personal access token | [optional] 
**AtlassianConfluenceUrl** | Pointer to **string** | Confluence Data Center base URL | [optional] 
**AtlassianDeploymentType** | Pointer to **string** | Atlassian deployment type: cloud or data_center (inferred from the URL when empty) | [optional] 
**AtlassianEmail** | Pointer to **string** | Atlassian account email paired with a user API token | [optional] 
**AtlassianJiraToken** | Pointer to **string** | Jira Data Center personal access token | [optional] 
**AtlassianJiraUrl** | Pointer to **string** | Jira Data Center base URL | [optional] 
**AtlassianProducts** | Pointer to **[]string** | Atlassian products to connect to on cloud: jira, confluence or both (default both) | [optional] 
**AtlassianToken** | Pointer to **string** | Atlassian Cloud service account API token or user API token | [optional] 
**AtlassianUrl** | Pointer to **string** | Atlassian Cloud site base URL (https://&lt;site&gt;.atlassian.net) | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LockOnRead** | Pointer to **string** | Lock this secret after each successful value read | [optional] 
**LockTtl** | Pointer to **string** | Lock TTL in minutes | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**RotateOnUnlock** | Pointer to **string** | Rotate this secret after it is unlocked | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateAtlassian

`func NewTargetCreateAtlassian(name string, ) *TargetCreateAtlassian`

NewTargetCreateAtlassian instantiates a new TargetCreateAtlassian object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateAtlassianWithDefaults

`func NewTargetCreateAtlassianWithDefaults() *TargetCreateAtlassian`

NewTargetCreateAtlassianWithDefaults instantiates a new TargetCreateAtlassian object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlassianAuthType

`func (o *TargetCreateAtlassian) GetAtlassianAuthType() string`

GetAtlassianAuthType returns the AtlassianAuthType field if non-nil, zero value otherwise.

### GetAtlassianAuthTypeOk

`func (o *TargetCreateAtlassian) GetAtlassianAuthTypeOk() (*string, bool)`

GetAtlassianAuthTypeOk returns a tuple with the AtlassianAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianAuthType

`func (o *TargetCreateAtlassian) SetAtlassianAuthType(v string)`

SetAtlassianAuthType sets AtlassianAuthType field to given value.

### HasAtlassianAuthType

`func (o *TargetCreateAtlassian) HasAtlassianAuthType() bool`

HasAtlassianAuthType returns a boolean if a field has been set.

### GetAtlassianAutomationKey

`func (o *TargetCreateAtlassian) GetAtlassianAutomationKey() string`

GetAtlassianAutomationKey returns the AtlassianAutomationKey field if non-nil, zero value otherwise.

### GetAtlassianAutomationKeyOk

`func (o *TargetCreateAtlassian) GetAtlassianAutomationKeyOk() (*string, bool)`

GetAtlassianAutomationKeyOk returns a tuple with the AtlassianAutomationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianAutomationKey

`func (o *TargetCreateAtlassian) SetAtlassianAutomationKey(v string)`

SetAtlassianAutomationKey sets AtlassianAutomationKey field to given value.

### HasAtlassianAutomationKey

`func (o *TargetCreateAtlassian) HasAtlassianAutomationKey() bool`

HasAtlassianAutomationKey returns a boolean if a field has been set.

### GetAtlassianCaCertificate

`func (o *TargetCreateAtlassian) GetAtlassianCaCertificate() string`

GetAtlassianCaCertificate returns the AtlassianCaCertificate field if non-nil, zero value otherwise.

### GetAtlassianCaCertificateOk

`func (o *TargetCreateAtlassian) GetAtlassianCaCertificateOk() (*string, bool)`

GetAtlassianCaCertificateOk returns a tuple with the AtlassianCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianCaCertificate

`func (o *TargetCreateAtlassian) SetAtlassianCaCertificate(v string)`

SetAtlassianCaCertificate sets AtlassianCaCertificate field to given value.

### HasAtlassianCaCertificate

`func (o *TargetCreateAtlassian) HasAtlassianCaCertificate() bool`

HasAtlassianCaCertificate returns a boolean if a field has been set.

### GetAtlassianClientId

`func (o *TargetCreateAtlassian) GetAtlassianClientId() string`

GetAtlassianClientId returns the AtlassianClientId field if non-nil, zero value otherwise.

### GetAtlassianClientIdOk

`func (o *TargetCreateAtlassian) GetAtlassianClientIdOk() (*string, bool)`

GetAtlassianClientIdOk returns a tuple with the AtlassianClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianClientId

`func (o *TargetCreateAtlassian) SetAtlassianClientId(v string)`

SetAtlassianClientId sets AtlassianClientId field to given value.

### HasAtlassianClientId

`func (o *TargetCreateAtlassian) HasAtlassianClientId() bool`

HasAtlassianClientId returns a boolean if a field has been set.

### GetAtlassianClientSecret

`func (o *TargetCreateAtlassian) GetAtlassianClientSecret() string`

GetAtlassianClientSecret returns the AtlassianClientSecret field if non-nil, zero value otherwise.

### GetAtlassianClientSecretOk

`func (o *TargetCreateAtlassian) GetAtlassianClientSecretOk() (*string, bool)`

GetAtlassianClientSecretOk returns a tuple with the AtlassianClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianClientSecret

`func (o *TargetCreateAtlassian) SetAtlassianClientSecret(v string)`

SetAtlassianClientSecret sets AtlassianClientSecret field to given value.

### HasAtlassianClientSecret

`func (o *TargetCreateAtlassian) HasAtlassianClientSecret() bool`

HasAtlassianClientSecret returns a boolean if a field has been set.

### GetAtlassianCloudId

`func (o *TargetCreateAtlassian) GetAtlassianCloudId() string`

GetAtlassianCloudId returns the AtlassianCloudId field if non-nil, zero value otherwise.

### GetAtlassianCloudIdOk

`func (o *TargetCreateAtlassian) GetAtlassianCloudIdOk() (*string, bool)`

GetAtlassianCloudIdOk returns a tuple with the AtlassianCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianCloudId

`func (o *TargetCreateAtlassian) SetAtlassianCloudId(v string)`

SetAtlassianCloudId sets AtlassianCloudId field to given value.

### HasAtlassianCloudId

`func (o *TargetCreateAtlassian) HasAtlassianCloudId() bool`

HasAtlassianCloudId returns a boolean if a field has been set.

### GetAtlassianConfluenceToken

`func (o *TargetCreateAtlassian) GetAtlassianConfluenceToken() string`

GetAtlassianConfluenceToken returns the AtlassianConfluenceToken field if non-nil, zero value otherwise.

### GetAtlassianConfluenceTokenOk

`func (o *TargetCreateAtlassian) GetAtlassianConfluenceTokenOk() (*string, bool)`

GetAtlassianConfluenceTokenOk returns a tuple with the AtlassianConfluenceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianConfluenceToken

`func (o *TargetCreateAtlassian) SetAtlassianConfluenceToken(v string)`

SetAtlassianConfluenceToken sets AtlassianConfluenceToken field to given value.

### HasAtlassianConfluenceToken

`func (o *TargetCreateAtlassian) HasAtlassianConfluenceToken() bool`

HasAtlassianConfluenceToken returns a boolean if a field has been set.

### GetAtlassianConfluenceUrl

`func (o *TargetCreateAtlassian) GetAtlassianConfluenceUrl() string`

GetAtlassianConfluenceUrl returns the AtlassianConfluenceUrl field if non-nil, zero value otherwise.

### GetAtlassianConfluenceUrlOk

`func (o *TargetCreateAtlassian) GetAtlassianConfluenceUrlOk() (*string, bool)`

GetAtlassianConfluenceUrlOk returns a tuple with the AtlassianConfluenceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianConfluenceUrl

`func (o *TargetCreateAtlassian) SetAtlassianConfluenceUrl(v string)`

SetAtlassianConfluenceUrl sets AtlassianConfluenceUrl field to given value.

### HasAtlassianConfluenceUrl

`func (o *TargetCreateAtlassian) HasAtlassianConfluenceUrl() bool`

HasAtlassianConfluenceUrl returns a boolean if a field has been set.

### GetAtlassianDeploymentType

`func (o *TargetCreateAtlassian) GetAtlassianDeploymentType() string`

GetAtlassianDeploymentType returns the AtlassianDeploymentType field if non-nil, zero value otherwise.

### GetAtlassianDeploymentTypeOk

`func (o *TargetCreateAtlassian) GetAtlassianDeploymentTypeOk() (*string, bool)`

GetAtlassianDeploymentTypeOk returns a tuple with the AtlassianDeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianDeploymentType

`func (o *TargetCreateAtlassian) SetAtlassianDeploymentType(v string)`

SetAtlassianDeploymentType sets AtlassianDeploymentType field to given value.

### HasAtlassianDeploymentType

`func (o *TargetCreateAtlassian) HasAtlassianDeploymentType() bool`

HasAtlassianDeploymentType returns a boolean if a field has been set.

### GetAtlassianEmail

`func (o *TargetCreateAtlassian) GetAtlassianEmail() string`

GetAtlassianEmail returns the AtlassianEmail field if non-nil, zero value otherwise.

### GetAtlassianEmailOk

`func (o *TargetCreateAtlassian) GetAtlassianEmailOk() (*string, bool)`

GetAtlassianEmailOk returns a tuple with the AtlassianEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianEmail

`func (o *TargetCreateAtlassian) SetAtlassianEmail(v string)`

SetAtlassianEmail sets AtlassianEmail field to given value.

### HasAtlassianEmail

`func (o *TargetCreateAtlassian) HasAtlassianEmail() bool`

HasAtlassianEmail returns a boolean if a field has been set.

### GetAtlassianJiraToken

`func (o *TargetCreateAtlassian) GetAtlassianJiraToken() string`

GetAtlassianJiraToken returns the AtlassianJiraToken field if non-nil, zero value otherwise.

### GetAtlassianJiraTokenOk

`func (o *TargetCreateAtlassian) GetAtlassianJiraTokenOk() (*string, bool)`

GetAtlassianJiraTokenOk returns a tuple with the AtlassianJiraToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianJiraToken

`func (o *TargetCreateAtlassian) SetAtlassianJiraToken(v string)`

SetAtlassianJiraToken sets AtlassianJiraToken field to given value.

### HasAtlassianJiraToken

`func (o *TargetCreateAtlassian) HasAtlassianJiraToken() bool`

HasAtlassianJiraToken returns a boolean if a field has been set.

### GetAtlassianJiraUrl

`func (o *TargetCreateAtlassian) GetAtlassianJiraUrl() string`

GetAtlassianJiraUrl returns the AtlassianJiraUrl field if non-nil, zero value otherwise.

### GetAtlassianJiraUrlOk

`func (o *TargetCreateAtlassian) GetAtlassianJiraUrlOk() (*string, bool)`

GetAtlassianJiraUrlOk returns a tuple with the AtlassianJiraUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianJiraUrl

`func (o *TargetCreateAtlassian) SetAtlassianJiraUrl(v string)`

SetAtlassianJiraUrl sets AtlassianJiraUrl field to given value.

### HasAtlassianJiraUrl

`func (o *TargetCreateAtlassian) HasAtlassianJiraUrl() bool`

HasAtlassianJiraUrl returns a boolean if a field has been set.

### GetAtlassianProducts

`func (o *TargetCreateAtlassian) GetAtlassianProducts() []string`

GetAtlassianProducts returns the AtlassianProducts field if non-nil, zero value otherwise.

### GetAtlassianProductsOk

`func (o *TargetCreateAtlassian) GetAtlassianProductsOk() (*[]string, bool)`

GetAtlassianProductsOk returns a tuple with the AtlassianProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianProducts

`func (o *TargetCreateAtlassian) SetAtlassianProducts(v []string)`

SetAtlassianProducts sets AtlassianProducts field to given value.

### HasAtlassianProducts

`func (o *TargetCreateAtlassian) HasAtlassianProducts() bool`

HasAtlassianProducts returns a boolean if a field has been set.

### GetAtlassianToken

`func (o *TargetCreateAtlassian) GetAtlassianToken() string`

GetAtlassianToken returns the AtlassianToken field if non-nil, zero value otherwise.

### GetAtlassianTokenOk

`func (o *TargetCreateAtlassian) GetAtlassianTokenOk() (*string, bool)`

GetAtlassianTokenOk returns a tuple with the AtlassianToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianToken

`func (o *TargetCreateAtlassian) SetAtlassianToken(v string)`

SetAtlassianToken sets AtlassianToken field to given value.

### HasAtlassianToken

`func (o *TargetCreateAtlassian) HasAtlassianToken() bool`

HasAtlassianToken returns a boolean if a field has been set.

### GetAtlassianUrl

`func (o *TargetCreateAtlassian) GetAtlassianUrl() string`

GetAtlassianUrl returns the AtlassianUrl field if non-nil, zero value otherwise.

### GetAtlassianUrlOk

`func (o *TargetCreateAtlassian) GetAtlassianUrlOk() (*string, bool)`

GetAtlassianUrlOk returns a tuple with the AtlassianUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianUrl

`func (o *TargetCreateAtlassian) SetAtlassianUrl(v string)`

SetAtlassianUrl sets AtlassianUrl field to given value.

### HasAtlassianUrl

`func (o *TargetCreateAtlassian) HasAtlassianUrl() bool`

HasAtlassianUrl returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *TargetCreateAtlassian) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *TargetCreateAtlassian) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *TargetCreateAtlassian) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *TargetCreateAtlassian) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateAtlassian) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateAtlassian) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateAtlassian) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateAtlassian) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateAtlassian) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateAtlassian) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateAtlassian) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateAtlassian) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateAtlassian) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateAtlassian) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateAtlassian) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateAtlassian) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLockOnRead

`func (o *TargetCreateAtlassian) GetLockOnRead() string`

GetLockOnRead returns the LockOnRead field if non-nil, zero value otherwise.

### GetLockOnReadOk

`func (o *TargetCreateAtlassian) GetLockOnReadOk() (*string, bool)`

GetLockOnReadOk returns a tuple with the LockOnRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOnRead

`func (o *TargetCreateAtlassian) SetLockOnRead(v string)`

SetLockOnRead sets LockOnRead field to given value.

### HasLockOnRead

`func (o *TargetCreateAtlassian) HasLockOnRead() bool`

HasLockOnRead returns a boolean if a field has been set.

### GetLockTtl

`func (o *TargetCreateAtlassian) GetLockTtl() string`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *TargetCreateAtlassian) GetLockTtlOk() (*string, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *TargetCreateAtlassian) SetLockTtl(v string)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *TargetCreateAtlassian) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateAtlassian) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateAtlassian) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateAtlassian) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateAtlassian) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateAtlassian) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateAtlassian) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateAtlassian) SetName(v string)`

SetName sets Name field to given value.


### GetRotateOnUnlock

`func (o *TargetCreateAtlassian) GetRotateOnUnlock() string`

GetRotateOnUnlock returns the RotateOnUnlock field if non-nil, zero value otherwise.

### GetRotateOnUnlockOk

`func (o *TargetCreateAtlassian) GetRotateOnUnlockOk() (*string, bool)`

GetRotateOnUnlockOk returns a tuple with the RotateOnUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateOnUnlock

`func (o *TargetCreateAtlassian) SetRotateOnUnlock(v string)`

SetRotateOnUnlock sets RotateOnUnlock field to given value.

### HasRotateOnUnlock

`func (o *TargetCreateAtlassian) HasRotateOnUnlock() bool`

HasRotateOnUnlock returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateAtlassian) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateAtlassian) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateAtlassian) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateAtlassian) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateAtlassian) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateAtlassian) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateAtlassian) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateAtlassian) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


