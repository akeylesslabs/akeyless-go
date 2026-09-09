# CreateUSC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureKvName** | Pointer to **string** | Azure Key Vault name (Relevant only for Azure targets) | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the Universal Secrets Connector | [optional] 
**EnvironmentNames** | Pointer to **string** | The environments in repo-name/environment-name format, comma-separated (only relevant for: github-scope&#x3D;repository-environment) | [optional] 
**GcpFolderId** | Pointer to **string** | GCP Folder ID (Relevant only for GCP targets with folder scope) | [optional] 
**GcpOrganizationId** | Pointer to **string** | GCP Organization ID (Relevant only for GCP targets with folder or organization scope) | [optional] 
**GcpProjectId** | Pointer to **string** | GCP Project ID to manage secrets in (Relevant only for GCP targets with project scope). If empty, the project is taken from the target. | [optional] 
**GcpScope** | Pointer to **string** | The GCP USC scope [project/folder/organization]. With folder/organization, all projects under the folder/org are manageable. Relevant only for GCP targets. Defaults to project when empty. | [optional] [default to "project"]
**GcpSmRegions** | Pointer to **string** | GCP Secret Manager regions to query for regional secrets (comma-separated, e.g., us-east1,us-west1). Max 12 regions. Required when listing with object-type&#x3D;regional-secrets. | [optional] 
**GithubScope** | Pointer to **string** | The scope where secrets will be created, available options: [repository, organization, repository-environment] | [optional] [default to "repository"]
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**K8sNamespace** | Pointer to **string** | K8s namespace (Relevant to Kubernetes targets) | [optional] 
**Name** | **string** | Universal Secrets Connector name | 
**OrganizationName** | Pointer to **string** | The organization name to create the secret in | [optional] 
**RepositoryAccess** | Pointer to **string** |  | [optional] [default to "public"]
**RepositoryNames** | Pointer to **string** | The repository names, comma-separated (only relevant for: github-scope&#x3D;repository) | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this Universal Secrets Connector | [optional] 
**TargetToAssociate** | **string** | Target Universal Secrets Connector to connect | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscPrefix** | Pointer to **string** | Prefix for all secrets created in AWS Secrets Manager | [optional] 
**UscTags** | Pointer to **string** | Comma-separated list of tags to apply to all secrets created on the remote USC | [optional] 
**UsePrefixAsFilter** | Pointer to **string** | Whether to filter the USC secret list using the specified usc-prefix [true/false] | [optional] [default to "false"]
**UseTagsAsFilter** | Pointer to **bool** | Filter the USC secret list by the value(s) of --usc-tags. [true|false] | [optional] 

## Methods

### NewCreateUSC

`func NewCreateUSC(name string, targetToAssociate string, ) *CreateUSC`

NewCreateUSC instantiates a new CreateUSC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUSCWithDefaults

`func NewCreateUSCWithDefaults() *CreateUSC`

NewCreateUSCWithDefaults instantiates a new CreateUSC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureKvName

`func (o *CreateUSC) GetAzureKvName() string`

GetAzureKvName returns the AzureKvName field if non-nil, zero value otherwise.

### GetAzureKvNameOk

`func (o *CreateUSC) GetAzureKvNameOk() (*string, bool)`

GetAzureKvNameOk returns a tuple with the AzureKvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKvName

`func (o *CreateUSC) SetAzureKvName(v string)`

SetAzureKvName sets AzureKvName field to given value.

### HasAzureKvName

`func (o *CreateUSC) HasAzureKvName() bool`

HasAzureKvName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateUSC) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateUSC) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateUSC) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateUSC) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateUSC) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUSC) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUSC) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUSC) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentNames

`func (o *CreateUSC) GetEnvironmentNames() string`

GetEnvironmentNames returns the EnvironmentNames field if non-nil, zero value otherwise.

### GetEnvironmentNamesOk

`func (o *CreateUSC) GetEnvironmentNamesOk() (*string, bool)`

GetEnvironmentNamesOk returns a tuple with the EnvironmentNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentNames

`func (o *CreateUSC) SetEnvironmentNames(v string)`

SetEnvironmentNames sets EnvironmentNames field to given value.

### HasEnvironmentNames

`func (o *CreateUSC) HasEnvironmentNames() bool`

HasEnvironmentNames returns a boolean if a field has been set.

### GetGcpFolderId

`func (o *CreateUSC) GetGcpFolderId() string`

GetGcpFolderId returns the GcpFolderId field if non-nil, zero value otherwise.

### GetGcpFolderIdOk

`func (o *CreateUSC) GetGcpFolderIdOk() (*string, bool)`

GetGcpFolderIdOk returns a tuple with the GcpFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpFolderId

`func (o *CreateUSC) SetGcpFolderId(v string)`

SetGcpFolderId sets GcpFolderId field to given value.

### HasGcpFolderId

`func (o *CreateUSC) HasGcpFolderId() bool`

HasGcpFolderId returns a boolean if a field has been set.

### GetGcpOrganizationId

`func (o *CreateUSC) GetGcpOrganizationId() string`

GetGcpOrganizationId returns the GcpOrganizationId field if non-nil, zero value otherwise.

### GetGcpOrganizationIdOk

`func (o *CreateUSC) GetGcpOrganizationIdOk() (*string, bool)`

GetGcpOrganizationIdOk returns a tuple with the GcpOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpOrganizationId

`func (o *CreateUSC) SetGcpOrganizationId(v string)`

SetGcpOrganizationId sets GcpOrganizationId field to given value.

### HasGcpOrganizationId

`func (o *CreateUSC) HasGcpOrganizationId() bool`

HasGcpOrganizationId returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *CreateUSC) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *CreateUSC) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *CreateUSC) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *CreateUSC) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.

### GetGcpScope

`func (o *CreateUSC) GetGcpScope() string`

GetGcpScope returns the GcpScope field if non-nil, zero value otherwise.

### GetGcpScopeOk

`func (o *CreateUSC) GetGcpScopeOk() (*string, bool)`

GetGcpScopeOk returns a tuple with the GcpScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpScope

`func (o *CreateUSC) SetGcpScope(v string)`

SetGcpScope sets GcpScope field to given value.

### HasGcpScope

`func (o *CreateUSC) HasGcpScope() bool`

HasGcpScope returns a boolean if a field has been set.

### GetGcpSmRegions

`func (o *CreateUSC) GetGcpSmRegions() string`

GetGcpSmRegions returns the GcpSmRegions field if non-nil, zero value otherwise.

### GetGcpSmRegionsOk

`func (o *CreateUSC) GetGcpSmRegionsOk() (*string, bool)`

GetGcpSmRegionsOk returns a tuple with the GcpSmRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSmRegions

`func (o *CreateUSC) SetGcpSmRegions(v string)`

SetGcpSmRegions sets GcpSmRegions field to given value.

### HasGcpSmRegions

`func (o *CreateUSC) HasGcpSmRegions() bool`

HasGcpSmRegions returns a boolean if a field has been set.

### GetGithubScope

`func (o *CreateUSC) GetGithubScope() string`

GetGithubScope returns the GithubScope field if non-nil, zero value otherwise.

### GetGithubScopeOk

`func (o *CreateUSC) GetGithubScopeOk() (*string, bool)`

GetGithubScopeOk returns a tuple with the GithubScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubScope

`func (o *CreateUSC) SetGithubScope(v string)`

SetGithubScope sets GithubScope field to given value.

### HasGithubScope

`func (o *CreateUSC) HasGithubScope() bool`

HasGithubScope returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *CreateUSC) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *CreateUSC) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *CreateUSC) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *CreateUSC) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *CreateUSC) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateUSC) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateUSC) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateUSC) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sNamespace

`func (o *CreateUSC) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *CreateUSC) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *CreateUSC) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *CreateUSC) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### GetName

`func (o *CreateUSC) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUSC) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUSC) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationName

`func (o *CreateUSC) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CreateUSC) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CreateUSC) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *CreateUSC) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetRepositoryAccess

`func (o *CreateUSC) GetRepositoryAccess() string`

GetRepositoryAccess returns the RepositoryAccess field if non-nil, zero value otherwise.

### GetRepositoryAccessOk

`func (o *CreateUSC) GetRepositoryAccessOk() (*string, bool)`

GetRepositoryAccessOk returns a tuple with the RepositoryAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryAccess

`func (o *CreateUSC) SetRepositoryAccess(v string)`

SetRepositoryAccess sets RepositoryAccess field to given value.

### HasRepositoryAccess

`func (o *CreateUSC) HasRepositoryAccess() bool`

HasRepositoryAccess returns a boolean if a field has been set.

### GetRepositoryNames

`func (o *CreateUSC) GetRepositoryNames() string`

GetRepositoryNames returns the RepositoryNames field if non-nil, zero value otherwise.

### GetRepositoryNamesOk

`func (o *CreateUSC) GetRepositoryNamesOk() (*string, bool)`

GetRepositoryNamesOk returns a tuple with the RepositoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryNames

`func (o *CreateUSC) SetRepositoryNames(v string)`

SetRepositoryNames sets RepositoryNames field to given value.

### HasRepositoryNames

`func (o *CreateUSC) HasRepositoryNames() bool`

HasRepositoryNames returns a boolean if a field has been set.

### GetTags

`func (o *CreateUSC) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateUSC) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateUSC) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateUSC) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetToAssociate

`func (o *CreateUSC) GetTargetToAssociate() string`

GetTargetToAssociate returns the TargetToAssociate field if non-nil, zero value otherwise.

### GetTargetToAssociateOk

`func (o *CreateUSC) GetTargetToAssociateOk() (*string, bool)`

GetTargetToAssociateOk returns a tuple with the TargetToAssociate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetToAssociate

`func (o *CreateUSC) SetTargetToAssociate(v string)`

SetTargetToAssociate sets TargetToAssociate field to given value.


### GetToken

`func (o *CreateUSC) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateUSC) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateUSC) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateUSC) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateUSC) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateUSC) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateUSC) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateUSC) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscPrefix

`func (o *CreateUSC) GetUscPrefix() string`

GetUscPrefix returns the UscPrefix field if non-nil, zero value otherwise.

### GetUscPrefixOk

`func (o *CreateUSC) GetUscPrefixOk() (*string, bool)`

GetUscPrefixOk returns a tuple with the UscPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscPrefix

`func (o *CreateUSC) SetUscPrefix(v string)`

SetUscPrefix sets UscPrefix field to given value.

### HasUscPrefix

`func (o *CreateUSC) HasUscPrefix() bool`

HasUscPrefix returns a boolean if a field has been set.

### GetUscTags

`func (o *CreateUSC) GetUscTags() string`

GetUscTags returns the UscTags field if non-nil, zero value otherwise.

### GetUscTagsOk

`func (o *CreateUSC) GetUscTagsOk() (*string, bool)`

GetUscTagsOk returns a tuple with the UscTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscTags

`func (o *CreateUSC) SetUscTags(v string)`

SetUscTags sets UscTags field to given value.

### HasUscTags

`func (o *CreateUSC) HasUscTags() bool`

HasUscTags returns a boolean if a field has been set.

### GetUsePrefixAsFilter

`func (o *CreateUSC) GetUsePrefixAsFilter() string`

GetUsePrefixAsFilter returns the UsePrefixAsFilter field if non-nil, zero value otherwise.

### GetUsePrefixAsFilterOk

`func (o *CreateUSC) GetUsePrefixAsFilterOk() (*string, bool)`

GetUsePrefixAsFilterOk returns a tuple with the UsePrefixAsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePrefixAsFilter

`func (o *CreateUSC) SetUsePrefixAsFilter(v string)`

SetUsePrefixAsFilter sets UsePrefixAsFilter field to given value.

### HasUsePrefixAsFilter

`func (o *CreateUSC) HasUsePrefixAsFilter() bool`

HasUsePrefixAsFilter returns a boolean if a field has been set.

### GetUseTagsAsFilter

`func (o *CreateUSC) GetUseTagsAsFilter() bool`

GetUseTagsAsFilter returns the UseTagsAsFilter field if non-nil, zero value otherwise.

### GetUseTagsAsFilterOk

`func (o *CreateUSC) GetUseTagsAsFilterOk() (*bool, bool)`

GetUseTagsAsFilterOk returns a tuple with the UseTagsAsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTagsAsFilter

`func (o *CreateUSC) SetUseTagsAsFilter(v bool)`

SetUseTagsAsFilter sets UseTagsAsFilter field to given value.

### HasUseTagsAsFilter

`func (o *CreateUSC) HasUseTagsAsFilter() bool`

HasUseTagsAsFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


