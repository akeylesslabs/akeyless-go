# TargetUpdateAtlassian

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
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LockOnRead** | Pointer to **string** | Lock this secret after each successful value read | [optional] 
**LockTtl** | Pointer to **string** | Lock TTL in minutes | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**RotateOnUnlock** | Pointer to **string** | Rotate this secret after it is unlocked | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateAtlassian

`func NewTargetUpdateAtlassian(name string, ) *TargetUpdateAtlassian`

NewTargetUpdateAtlassian instantiates a new TargetUpdateAtlassian object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateAtlassianWithDefaults

`func NewTargetUpdateAtlassianWithDefaults() *TargetUpdateAtlassian`

NewTargetUpdateAtlassianWithDefaults instantiates a new TargetUpdateAtlassian object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlassianAuthType

`func (o *TargetUpdateAtlassian) GetAtlassianAuthType() string`

GetAtlassianAuthType returns the AtlassianAuthType field if non-nil, zero value otherwise.

### GetAtlassianAuthTypeOk

`func (o *TargetUpdateAtlassian) GetAtlassianAuthTypeOk() (*string, bool)`

GetAtlassianAuthTypeOk returns a tuple with the AtlassianAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianAuthType

`func (o *TargetUpdateAtlassian) SetAtlassianAuthType(v string)`

SetAtlassianAuthType sets AtlassianAuthType field to given value.

### HasAtlassianAuthType

`func (o *TargetUpdateAtlassian) HasAtlassianAuthType() bool`

HasAtlassianAuthType returns a boolean if a field has been set.

### GetAtlassianAutomationKey

`func (o *TargetUpdateAtlassian) GetAtlassianAutomationKey() string`

GetAtlassianAutomationKey returns the AtlassianAutomationKey field if non-nil, zero value otherwise.

### GetAtlassianAutomationKeyOk

`func (o *TargetUpdateAtlassian) GetAtlassianAutomationKeyOk() (*string, bool)`

GetAtlassianAutomationKeyOk returns a tuple with the AtlassianAutomationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianAutomationKey

`func (o *TargetUpdateAtlassian) SetAtlassianAutomationKey(v string)`

SetAtlassianAutomationKey sets AtlassianAutomationKey field to given value.

### HasAtlassianAutomationKey

`func (o *TargetUpdateAtlassian) HasAtlassianAutomationKey() bool`

HasAtlassianAutomationKey returns a boolean if a field has been set.

### GetAtlassianCaCertificate

`func (o *TargetUpdateAtlassian) GetAtlassianCaCertificate() string`

GetAtlassianCaCertificate returns the AtlassianCaCertificate field if non-nil, zero value otherwise.

### GetAtlassianCaCertificateOk

`func (o *TargetUpdateAtlassian) GetAtlassianCaCertificateOk() (*string, bool)`

GetAtlassianCaCertificateOk returns a tuple with the AtlassianCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianCaCertificate

`func (o *TargetUpdateAtlassian) SetAtlassianCaCertificate(v string)`

SetAtlassianCaCertificate sets AtlassianCaCertificate field to given value.

### HasAtlassianCaCertificate

`func (o *TargetUpdateAtlassian) HasAtlassianCaCertificate() bool`

HasAtlassianCaCertificate returns a boolean if a field has been set.

### GetAtlassianClientId

`func (o *TargetUpdateAtlassian) GetAtlassianClientId() string`

GetAtlassianClientId returns the AtlassianClientId field if non-nil, zero value otherwise.

### GetAtlassianClientIdOk

`func (o *TargetUpdateAtlassian) GetAtlassianClientIdOk() (*string, bool)`

GetAtlassianClientIdOk returns a tuple with the AtlassianClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianClientId

`func (o *TargetUpdateAtlassian) SetAtlassianClientId(v string)`

SetAtlassianClientId sets AtlassianClientId field to given value.

### HasAtlassianClientId

`func (o *TargetUpdateAtlassian) HasAtlassianClientId() bool`

HasAtlassianClientId returns a boolean if a field has been set.

### GetAtlassianClientSecret

`func (o *TargetUpdateAtlassian) GetAtlassianClientSecret() string`

GetAtlassianClientSecret returns the AtlassianClientSecret field if non-nil, zero value otherwise.

### GetAtlassianClientSecretOk

`func (o *TargetUpdateAtlassian) GetAtlassianClientSecretOk() (*string, bool)`

GetAtlassianClientSecretOk returns a tuple with the AtlassianClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianClientSecret

`func (o *TargetUpdateAtlassian) SetAtlassianClientSecret(v string)`

SetAtlassianClientSecret sets AtlassianClientSecret field to given value.

### HasAtlassianClientSecret

`func (o *TargetUpdateAtlassian) HasAtlassianClientSecret() bool`

HasAtlassianClientSecret returns a boolean if a field has been set.

### GetAtlassianCloudId

`func (o *TargetUpdateAtlassian) GetAtlassianCloudId() string`

GetAtlassianCloudId returns the AtlassianCloudId field if non-nil, zero value otherwise.

### GetAtlassianCloudIdOk

`func (o *TargetUpdateAtlassian) GetAtlassianCloudIdOk() (*string, bool)`

GetAtlassianCloudIdOk returns a tuple with the AtlassianCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianCloudId

`func (o *TargetUpdateAtlassian) SetAtlassianCloudId(v string)`

SetAtlassianCloudId sets AtlassianCloudId field to given value.

### HasAtlassianCloudId

`func (o *TargetUpdateAtlassian) HasAtlassianCloudId() bool`

HasAtlassianCloudId returns a boolean if a field has been set.

### GetAtlassianConfluenceToken

`func (o *TargetUpdateAtlassian) GetAtlassianConfluenceToken() string`

GetAtlassianConfluenceToken returns the AtlassianConfluenceToken field if non-nil, zero value otherwise.

### GetAtlassianConfluenceTokenOk

`func (o *TargetUpdateAtlassian) GetAtlassianConfluenceTokenOk() (*string, bool)`

GetAtlassianConfluenceTokenOk returns a tuple with the AtlassianConfluenceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianConfluenceToken

`func (o *TargetUpdateAtlassian) SetAtlassianConfluenceToken(v string)`

SetAtlassianConfluenceToken sets AtlassianConfluenceToken field to given value.

### HasAtlassianConfluenceToken

`func (o *TargetUpdateAtlassian) HasAtlassianConfluenceToken() bool`

HasAtlassianConfluenceToken returns a boolean if a field has been set.

### GetAtlassianConfluenceUrl

`func (o *TargetUpdateAtlassian) GetAtlassianConfluenceUrl() string`

GetAtlassianConfluenceUrl returns the AtlassianConfluenceUrl field if non-nil, zero value otherwise.

### GetAtlassianConfluenceUrlOk

`func (o *TargetUpdateAtlassian) GetAtlassianConfluenceUrlOk() (*string, bool)`

GetAtlassianConfluenceUrlOk returns a tuple with the AtlassianConfluenceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianConfluenceUrl

`func (o *TargetUpdateAtlassian) SetAtlassianConfluenceUrl(v string)`

SetAtlassianConfluenceUrl sets AtlassianConfluenceUrl field to given value.

### HasAtlassianConfluenceUrl

`func (o *TargetUpdateAtlassian) HasAtlassianConfluenceUrl() bool`

HasAtlassianConfluenceUrl returns a boolean if a field has been set.

### GetAtlassianDeploymentType

`func (o *TargetUpdateAtlassian) GetAtlassianDeploymentType() string`

GetAtlassianDeploymentType returns the AtlassianDeploymentType field if non-nil, zero value otherwise.

### GetAtlassianDeploymentTypeOk

`func (o *TargetUpdateAtlassian) GetAtlassianDeploymentTypeOk() (*string, bool)`

GetAtlassianDeploymentTypeOk returns a tuple with the AtlassianDeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianDeploymentType

`func (o *TargetUpdateAtlassian) SetAtlassianDeploymentType(v string)`

SetAtlassianDeploymentType sets AtlassianDeploymentType field to given value.

### HasAtlassianDeploymentType

`func (o *TargetUpdateAtlassian) HasAtlassianDeploymentType() bool`

HasAtlassianDeploymentType returns a boolean if a field has been set.

### GetAtlassianEmail

`func (o *TargetUpdateAtlassian) GetAtlassianEmail() string`

GetAtlassianEmail returns the AtlassianEmail field if non-nil, zero value otherwise.

### GetAtlassianEmailOk

`func (o *TargetUpdateAtlassian) GetAtlassianEmailOk() (*string, bool)`

GetAtlassianEmailOk returns a tuple with the AtlassianEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianEmail

`func (o *TargetUpdateAtlassian) SetAtlassianEmail(v string)`

SetAtlassianEmail sets AtlassianEmail field to given value.

### HasAtlassianEmail

`func (o *TargetUpdateAtlassian) HasAtlassianEmail() bool`

HasAtlassianEmail returns a boolean if a field has been set.

### GetAtlassianJiraToken

`func (o *TargetUpdateAtlassian) GetAtlassianJiraToken() string`

GetAtlassianJiraToken returns the AtlassianJiraToken field if non-nil, zero value otherwise.

### GetAtlassianJiraTokenOk

`func (o *TargetUpdateAtlassian) GetAtlassianJiraTokenOk() (*string, bool)`

GetAtlassianJiraTokenOk returns a tuple with the AtlassianJiraToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianJiraToken

`func (o *TargetUpdateAtlassian) SetAtlassianJiraToken(v string)`

SetAtlassianJiraToken sets AtlassianJiraToken field to given value.

### HasAtlassianJiraToken

`func (o *TargetUpdateAtlassian) HasAtlassianJiraToken() bool`

HasAtlassianJiraToken returns a boolean if a field has been set.

### GetAtlassianJiraUrl

`func (o *TargetUpdateAtlassian) GetAtlassianJiraUrl() string`

GetAtlassianJiraUrl returns the AtlassianJiraUrl field if non-nil, zero value otherwise.

### GetAtlassianJiraUrlOk

`func (o *TargetUpdateAtlassian) GetAtlassianJiraUrlOk() (*string, bool)`

GetAtlassianJiraUrlOk returns a tuple with the AtlassianJiraUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianJiraUrl

`func (o *TargetUpdateAtlassian) SetAtlassianJiraUrl(v string)`

SetAtlassianJiraUrl sets AtlassianJiraUrl field to given value.

### HasAtlassianJiraUrl

`func (o *TargetUpdateAtlassian) HasAtlassianJiraUrl() bool`

HasAtlassianJiraUrl returns a boolean if a field has been set.

### GetAtlassianProducts

`func (o *TargetUpdateAtlassian) GetAtlassianProducts() []string`

GetAtlassianProducts returns the AtlassianProducts field if non-nil, zero value otherwise.

### GetAtlassianProductsOk

`func (o *TargetUpdateAtlassian) GetAtlassianProductsOk() (*[]string, bool)`

GetAtlassianProductsOk returns a tuple with the AtlassianProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianProducts

`func (o *TargetUpdateAtlassian) SetAtlassianProducts(v []string)`

SetAtlassianProducts sets AtlassianProducts field to given value.

### HasAtlassianProducts

`func (o *TargetUpdateAtlassian) HasAtlassianProducts() bool`

HasAtlassianProducts returns a boolean if a field has been set.

### GetAtlassianToken

`func (o *TargetUpdateAtlassian) GetAtlassianToken() string`

GetAtlassianToken returns the AtlassianToken field if non-nil, zero value otherwise.

### GetAtlassianTokenOk

`func (o *TargetUpdateAtlassian) GetAtlassianTokenOk() (*string, bool)`

GetAtlassianTokenOk returns a tuple with the AtlassianToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianToken

`func (o *TargetUpdateAtlassian) SetAtlassianToken(v string)`

SetAtlassianToken sets AtlassianToken field to given value.

### HasAtlassianToken

`func (o *TargetUpdateAtlassian) HasAtlassianToken() bool`

HasAtlassianToken returns a boolean if a field has been set.

### GetAtlassianUrl

`func (o *TargetUpdateAtlassian) GetAtlassianUrl() string`

GetAtlassianUrl returns the AtlassianUrl field if non-nil, zero value otherwise.

### GetAtlassianUrlOk

`func (o *TargetUpdateAtlassian) GetAtlassianUrlOk() (*string, bool)`

GetAtlassianUrlOk returns a tuple with the AtlassianUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianUrl

`func (o *TargetUpdateAtlassian) SetAtlassianUrl(v string)`

SetAtlassianUrl sets AtlassianUrl field to given value.

### HasAtlassianUrl

`func (o *TargetUpdateAtlassian) HasAtlassianUrl() bool`

HasAtlassianUrl returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *TargetUpdateAtlassian) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *TargetUpdateAtlassian) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *TargetUpdateAtlassian) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *TargetUpdateAtlassian) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateAtlassian) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateAtlassian) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateAtlassian) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateAtlassian) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateAtlassian) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateAtlassian) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateAtlassian) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateAtlassian) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateAtlassian) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateAtlassian) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateAtlassian) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateAtlassian) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateAtlassian) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateAtlassian) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateAtlassian) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateAtlassian) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLockOnRead

`func (o *TargetUpdateAtlassian) GetLockOnRead() string`

GetLockOnRead returns the LockOnRead field if non-nil, zero value otherwise.

### GetLockOnReadOk

`func (o *TargetUpdateAtlassian) GetLockOnReadOk() (*string, bool)`

GetLockOnReadOk returns a tuple with the LockOnRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOnRead

`func (o *TargetUpdateAtlassian) SetLockOnRead(v string)`

SetLockOnRead sets LockOnRead field to given value.

### HasLockOnRead

`func (o *TargetUpdateAtlassian) HasLockOnRead() bool`

HasLockOnRead returns a boolean if a field has been set.

### GetLockTtl

`func (o *TargetUpdateAtlassian) GetLockTtl() string`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *TargetUpdateAtlassian) GetLockTtlOk() (*string, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *TargetUpdateAtlassian) SetLockTtl(v string)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *TargetUpdateAtlassian) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateAtlassian) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateAtlassian) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateAtlassian) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateAtlassian) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateAtlassian) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateAtlassian) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateAtlassian) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateAtlassian) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateAtlassian) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateAtlassian) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateAtlassian) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRotateOnUnlock

`func (o *TargetUpdateAtlassian) GetRotateOnUnlock() string`

GetRotateOnUnlock returns the RotateOnUnlock field if non-nil, zero value otherwise.

### GetRotateOnUnlockOk

`func (o *TargetUpdateAtlassian) GetRotateOnUnlockOk() (*string, bool)`

GetRotateOnUnlockOk returns a tuple with the RotateOnUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateOnUnlock

`func (o *TargetUpdateAtlassian) SetRotateOnUnlock(v string)`

SetRotateOnUnlock sets RotateOnUnlock field to given value.

### HasRotateOnUnlock

`func (o *TargetUpdateAtlassian) HasRotateOnUnlock() bool`

HasRotateOnUnlock returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateAtlassian) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateAtlassian) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateAtlassian) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateAtlassian) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateAtlassian) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateAtlassian) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateAtlassian) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateAtlassian) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


