# TargetTypeDetailsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactoryTargetDetails** | Pointer to [**ArtifactoryTargetDetails**](ArtifactoryTargetDetails.md) |  | [optional] 
**AwsTargetDetails** | Pointer to [**AWSTargetDetails**](AWSTargetDetails.md) |  | [optional] 
**AzureTargetDetails** | Pointer to [**AzureTargetDetails**](AzureTargetDetails.md) |  | [optional] 
**ChefTargetDetails** | Pointer to [**ChefTargetDetails**](ChefTargetDetails.md) |  | [optional] 
**CustomTargetDetails** | Pointer to [**CustomTargetDetails**](CustomTargetDetails.md) |  | [optional] 
**DbTargetDetails** | Pointer to [**DbTargetDetails**](DbTargetDetails.md) |  | [optional] 
**DockerhubTargetDetails** | Pointer to [**DockerhubTargetDetails**](DockerhubTargetDetails.md) |  | [optional] 
**EksTargetDetails** | Pointer to [**EKSTargetDetails**](EKSTargetDetails.md) |  | [optional] 
**GcpTargetDetails** | Pointer to [**GcpTargetDetails**](GcpTargetDetails.md) |  | [optional] 
**GeminiTargetDetails** | Pointer to [**GeminiTargetDetails**](GeminiTargetDetails.md) |  | [optional] 
**GithubTargetDetails** | Pointer to [**GithubTargetDetails**](GithubTargetDetails.md) |  | [optional] 
**GitlabTargetDetails** | Pointer to [**GitlabTargetDetails**](GitlabTargetDetails.md) |  | [optional] 
**GkeTargetDetails** | Pointer to [**GKETargetDetails**](GKETargetDetails.md) |  | [optional] 
**GlobalsignAtlasTargetDetails** | Pointer to [**GlobalSignAtlasTargetDetails**](GlobalSignAtlasTargetDetails.md) |  | [optional] 
**GlobalsignTargetDetails** | Pointer to [**GlobalSignGCCTargetDetails**](GlobalSignGCCTargetDetails.md) |  | [optional] 
**GodaddyTargetDetails** | Pointer to [**GodaddyTargetDetails**](GodaddyTargetDetails.md) |  | [optional] 
**HashiVaultTargetDetails** | Pointer to [**HashiVaultTargetDetails**](HashiVaultTargetDetails.md) |  | [optional] 
**LdapTargetDetails** | Pointer to [**LdapTargetDetails**](LdapTargetDetails.md) |  | [optional] 
**LetsencryptTargetDetails** | Pointer to [**LetsEncryptTargetDetails**](LetsEncryptTargetDetails.md) |  | [optional] 
**LinkedTargetDetails** | Pointer to [**LinkedTargetDetails**](LinkedTargetDetails.md) |  | [optional] 
**MongoDbTargetDetails** | Pointer to [**MongoDBTargetDetails**](MongoDBTargetDetails.md) |  | [optional] 
**NativeK8sTargetDetails** | Pointer to [**NativeK8sTargetDetails**](NativeK8sTargetDetails.md) |  | [optional] 
**OpenaiTargetDetails** | Pointer to [**OpenAITargetDetails**](OpenAITargetDetails.md) |  | [optional] 
**PingTargetDetails** | Pointer to [**PingTargetDetails**](PingTargetDetails.md) |  | [optional] 
**RabbitMqTargetDetails** | Pointer to [**RabbitMQTargetDetails**](RabbitMQTargetDetails.md) |  | [optional] 
**SalesforceTargetDetails** | Pointer to [**SalesforceTargetDetails**](SalesforceTargetDetails.md) |  | [optional] 
**SectigoTargetDetails** | Pointer to [**SectigoTargetDetails**](SectigoTargetDetails.md) |  | [optional] 
**SplunkTargetDetails** | Pointer to [**SplunkTargetDetails**](SplunkTargetDetails.md) |  | [optional] 
**SshTargetDetails** | Pointer to [**SSHTargetDetails**](SSHTargetDetails.md) |  | [optional] 
**VenafiTargetDetails** | Pointer to [**VenafiTargetDetails**](VenafiTargetDetails.md) |  | [optional] 
**WebTargetDetails** | Pointer to [**WebTargetDetails**](WebTargetDetails.md) |  | [optional] 
**WindowsTargetDetails** | Pointer to [**WindowsTargetDetails**](WindowsTargetDetails.md) |  | [optional] 
**ZerosslTargetDetails** | Pointer to [**ZeroSSLTargetDetails**](ZeroSSLTargetDetails.md) |  | [optional] 

## Methods

### NewTargetTypeDetailsInput

`func NewTargetTypeDetailsInput() *TargetTypeDetailsInput`

NewTargetTypeDetailsInput instantiates a new TargetTypeDetailsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetTypeDetailsInputWithDefaults

`func NewTargetTypeDetailsInputWithDefaults() *TargetTypeDetailsInput`

NewTargetTypeDetailsInputWithDefaults instantiates a new TargetTypeDetailsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactoryTargetDetails

`func (o *TargetTypeDetailsInput) GetArtifactoryTargetDetails() ArtifactoryTargetDetails`

GetArtifactoryTargetDetails returns the ArtifactoryTargetDetails field if non-nil, zero value otherwise.

### GetArtifactoryTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetArtifactoryTargetDetailsOk() (*ArtifactoryTargetDetails, bool)`

GetArtifactoryTargetDetailsOk returns a tuple with the ArtifactoryTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryTargetDetails

`func (o *TargetTypeDetailsInput) SetArtifactoryTargetDetails(v ArtifactoryTargetDetails)`

SetArtifactoryTargetDetails sets ArtifactoryTargetDetails field to given value.

### HasArtifactoryTargetDetails

`func (o *TargetTypeDetailsInput) HasArtifactoryTargetDetails() bool`

HasArtifactoryTargetDetails returns a boolean if a field has been set.

### GetAwsTargetDetails

`func (o *TargetTypeDetailsInput) GetAwsTargetDetails() AWSTargetDetails`

GetAwsTargetDetails returns the AwsTargetDetails field if non-nil, zero value otherwise.

### GetAwsTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetAwsTargetDetailsOk() (*AWSTargetDetails, bool)`

GetAwsTargetDetailsOk returns a tuple with the AwsTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetDetails

`func (o *TargetTypeDetailsInput) SetAwsTargetDetails(v AWSTargetDetails)`

SetAwsTargetDetails sets AwsTargetDetails field to given value.

### HasAwsTargetDetails

`func (o *TargetTypeDetailsInput) HasAwsTargetDetails() bool`

HasAwsTargetDetails returns a boolean if a field has been set.

### GetAzureTargetDetails

`func (o *TargetTypeDetailsInput) GetAzureTargetDetails() AzureTargetDetails`

GetAzureTargetDetails returns the AzureTargetDetails field if non-nil, zero value otherwise.

### GetAzureTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetAzureTargetDetailsOk() (*AzureTargetDetails, bool)`

GetAzureTargetDetailsOk returns a tuple with the AzureTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTargetDetails

`func (o *TargetTypeDetailsInput) SetAzureTargetDetails(v AzureTargetDetails)`

SetAzureTargetDetails sets AzureTargetDetails field to given value.

### HasAzureTargetDetails

`func (o *TargetTypeDetailsInput) HasAzureTargetDetails() bool`

HasAzureTargetDetails returns a boolean if a field has been set.

### GetChefTargetDetails

`func (o *TargetTypeDetailsInput) GetChefTargetDetails() ChefTargetDetails`

GetChefTargetDetails returns the ChefTargetDetails field if non-nil, zero value otherwise.

### GetChefTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetChefTargetDetailsOk() (*ChefTargetDetails, bool)`

GetChefTargetDetailsOk returns a tuple with the ChefTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefTargetDetails

`func (o *TargetTypeDetailsInput) SetChefTargetDetails(v ChefTargetDetails)`

SetChefTargetDetails sets ChefTargetDetails field to given value.

### HasChefTargetDetails

`func (o *TargetTypeDetailsInput) HasChefTargetDetails() bool`

HasChefTargetDetails returns a boolean if a field has been set.

### GetCustomTargetDetails

`func (o *TargetTypeDetailsInput) GetCustomTargetDetails() CustomTargetDetails`

GetCustomTargetDetails returns the CustomTargetDetails field if non-nil, zero value otherwise.

### GetCustomTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetCustomTargetDetailsOk() (*CustomTargetDetails, bool)`

GetCustomTargetDetailsOk returns a tuple with the CustomTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTargetDetails

`func (o *TargetTypeDetailsInput) SetCustomTargetDetails(v CustomTargetDetails)`

SetCustomTargetDetails sets CustomTargetDetails field to given value.

### HasCustomTargetDetails

`func (o *TargetTypeDetailsInput) HasCustomTargetDetails() bool`

HasCustomTargetDetails returns a boolean if a field has been set.

### GetDbTargetDetails

`func (o *TargetTypeDetailsInput) GetDbTargetDetails() DbTargetDetails`

GetDbTargetDetails returns the DbTargetDetails field if non-nil, zero value otherwise.

### GetDbTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetDbTargetDetailsOk() (*DbTargetDetails, bool)`

GetDbTargetDetailsOk returns a tuple with the DbTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTargetDetails

`func (o *TargetTypeDetailsInput) SetDbTargetDetails(v DbTargetDetails)`

SetDbTargetDetails sets DbTargetDetails field to given value.

### HasDbTargetDetails

`func (o *TargetTypeDetailsInput) HasDbTargetDetails() bool`

HasDbTargetDetails returns a boolean if a field has been set.

### GetDockerhubTargetDetails

`func (o *TargetTypeDetailsInput) GetDockerhubTargetDetails() DockerhubTargetDetails`

GetDockerhubTargetDetails returns the DockerhubTargetDetails field if non-nil, zero value otherwise.

### GetDockerhubTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetDockerhubTargetDetailsOk() (*DockerhubTargetDetails, bool)`

GetDockerhubTargetDetailsOk returns a tuple with the DockerhubTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubTargetDetails

`func (o *TargetTypeDetailsInput) SetDockerhubTargetDetails(v DockerhubTargetDetails)`

SetDockerhubTargetDetails sets DockerhubTargetDetails field to given value.

### HasDockerhubTargetDetails

`func (o *TargetTypeDetailsInput) HasDockerhubTargetDetails() bool`

HasDockerhubTargetDetails returns a boolean if a field has been set.

### GetEksTargetDetails

`func (o *TargetTypeDetailsInput) GetEksTargetDetails() EKSTargetDetails`

GetEksTargetDetails returns the EksTargetDetails field if non-nil, zero value otherwise.

### GetEksTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetEksTargetDetailsOk() (*EKSTargetDetails, bool)`

GetEksTargetDetailsOk returns a tuple with the EksTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksTargetDetails

`func (o *TargetTypeDetailsInput) SetEksTargetDetails(v EKSTargetDetails)`

SetEksTargetDetails sets EksTargetDetails field to given value.

### HasEksTargetDetails

`func (o *TargetTypeDetailsInput) HasEksTargetDetails() bool`

HasEksTargetDetails returns a boolean if a field has been set.

### GetGcpTargetDetails

`func (o *TargetTypeDetailsInput) GetGcpTargetDetails() GcpTargetDetails`

GetGcpTargetDetails returns the GcpTargetDetails field if non-nil, zero value otherwise.

### GetGcpTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetGcpTargetDetailsOk() (*GcpTargetDetails, bool)`

GetGcpTargetDetailsOk returns a tuple with the GcpTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTargetDetails

`func (o *TargetTypeDetailsInput) SetGcpTargetDetails(v GcpTargetDetails)`

SetGcpTargetDetails sets GcpTargetDetails field to given value.

### HasGcpTargetDetails

`func (o *TargetTypeDetailsInput) HasGcpTargetDetails() bool`

HasGcpTargetDetails returns a boolean if a field has been set.

### GetGeminiTargetDetails

`func (o *TargetTypeDetailsInput) GetGeminiTargetDetails() GeminiTargetDetails`

GetGeminiTargetDetails returns the GeminiTargetDetails field if non-nil, zero value otherwise.

### GetGeminiTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetGeminiTargetDetailsOk() (*GeminiTargetDetails, bool)`

GetGeminiTargetDetailsOk returns a tuple with the GeminiTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeminiTargetDetails

`func (o *TargetTypeDetailsInput) SetGeminiTargetDetails(v GeminiTargetDetails)`

SetGeminiTargetDetails sets GeminiTargetDetails field to given value.

### HasGeminiTargetDetails

`func (o *TargetTypeDetailsInput) HasGeminiTargetDetails() bool`

HasGeminiTargetDetails returns a boolean if a field has been set.

### GetGithubTargetDetails

`func (o *TargetTypeDetailsInput) GetGithubTargetDetails() GithubTargetDetails`

GetGithubTargetDetails returns the GithubTargetDetails field if non-nil, zero value otherwise.

### GetGithubTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetGithubTargetDetailsOk() (*GithubTargetDetails, bool)`

GetGithubTargetDetailsOk returns a tuple with the GithubTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubTargetDetails

`func (o *TargetTypeDetailsInput) SetGithubTargetDetails(v GithubTargetDetails)`

SetGithubTargetDetails sets GithubTargetDetails field to given value.

### HasGithubTargetDetails

`func (o *TargetTypeDetailsInput) HasGithubTargetDetails() bool`

HasGithubTargetDetails returns a boolean if a field has been set.

### GetGitlabTargetDetails

`func (o *TargetTypeDetailsInput) GetGitlabTargetDetails() GitlabTargetDetails`

GetGitlabTargetDetails returns the GitlabTargetDetails field if non-nil, zero value otherwise.

### GetGitlabTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetGitlabTargetDetailsOk() (*GitlabTargetDetails, bool)`

GetGitlabTargetDetailsOk returns a tuple with the GitlabTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabTargetDetails

`func (o *TargetTypeDetailsInput) SetGitlabTargetDetails(v GitlabTargetDetails)`

SetGitlabTargetDetails sets GitlabTargetDetails field to given value.

### HasGitlabTargetDetails

`func (o *TargetTypeDetailsInput) HasGitlabTargetDetails() bool`

HasGitlabTargetDetails returns a boolean if a field has been set.

### GetGkeTargetDetails

`func (o *TargetTypeDetailsInput) GetGkeTargetDetails() GKETargetDetails`

GetGkeTargetDetails returns the GkeTargetDetails field if non-nil, zero value otherwise.

### GetGkeTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetGkeTargetDetailsOk() (*GKETargetDetails, bool)`

GetGkeTargetDetailsOk returns a tuple with the GkeTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeTargetDetails

`func (o *TargetTypeDetailsInput) SetGkeTargetDetails(v GKETargetDetails)`

SetGkeTargetDetails sets GkeTargetDetails field to given value.

### HasGkeTargetDetails

`func (o *TargetTypeDetailsInput) HasGkeTargetDetails() bool`

HasGkeTargetDetails returns a boolean if a field has been set.

### GetGlobalsignAtlasTargetDetails

`func (o *TargetTypeDetailsInput) GetGlobalsignAtlasTargetDetails() GlobalSignAtlasTargetDetails`

GetGlobalsignAtlasTargetDetails returns the GlobalsignAtlasTargetDetails field if non-nil, zero value otherwise.

### GetGlobalsignAtlasTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetGlobalsignAtlasTargetDetailsOk() (*GlobalSignAtlasTargetDetails, bool)`

GetGlobalsignAtlasTargetDetailsOk returns a tuple with the GlobalsignAtlasTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalsignAtlasTargetDetails

`func (o *TargetTypeDetailsInput) SetGlobalsignAtlasTargetDetails(v GlobalSignAtlasTargetDetails)`

SetGlobalsignAtlasTargetDetails sets GlobalsignAtlasTargetDetails field to given value.

### HasGlobalsignAtlasTargetDetails

`func (o *TargetTypeDetailsInput) HasGlobalsignAtlasTargetDetails() bool`

HasGlobalsignAtlasTargetDetails returns a boolean if a field has been set.

### GetGlobalsignTargetDetails

`func (o *TargetTypeDetailsInput) GetGlobalsignTargetDetails() GlobalSignGCCTargetDetails`

GetGlobalsignTargetDetails returns the GlobalsignTargetDetails field if non-nil, zero value otherwise.

### GetGlobalsignTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetGlobalsignTargetDetailsOk() (*GlobalSignGCCTargetDetails, bool)`

GetGlobalsignTargetDetailsOk returns a tuple with the GlobalsignTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalsignTargetDetails

`func (o *TargetTypeDetailsInput) SetGlobalsignTargetDetails(v GlobalSignGCCTargetDetails)`

SetGlobalsignTargetDetails sets GlobalsignTargetDetails field to given value.

### HasGlobalsignTargetDetails

`func (o *TargetTypeDetailsInput) HasGlobalsignTargetDetails() bool`

HasGlobalsignTargetDetails returns a boolean if a field has been set.

### GetGodaddyTargetDetails

`func (o *TargetTypeDetailsInput) GetGodaddyTargetDetails() GodaddyTargetDetails`

GetGodaddyTargetDetails returns the GodaddyTargetDetails field if non-nil, zero value otherwise.

### GetGodaddyTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetGodaddyTargetDetailsOk() (*GodaddyTargetDetails, bool)`

GetGodaddyTargetDetailsOk returns a tuple with the GodaddyTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGodaddyTargetDetails

`func (o *TargetTypeDetailsInput) SetGodaddyTargetDetails(v GodaddyTargetDetails)`

SetGodaddyTargetDetails sets GodaddyTargetDetails field to given value.

### HasGodaddyTargetDetails

`func (o *TargetTypeDetailsInput) HasGodaddyTargetDetails() bool`

HasGodaddyTargetDetails returns a boolean if a field has been set.

### GetHashiVaultTargetDetails

`func (o *TargetTypeDetailsInput) GetHashiVaultTargetDetails() HashiVaultTargetDetails`

GetHashiVaultTargetDetails returns the HashiVaultTargetDetails field if non-nil, zero value otherwise.

### GetHashiVaultTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetHashiVaultTargetDetailsOk() (*HashiVaultTargetDetails, bool)`

GetHashiVaultTargetDetailsOk returns a tuple with the HashiVaultTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiVaultTargetDetails

`func (o *TargetTypeDetailsInput) SetHashiVaultTargetDetails(v HashiVaultTargetDetails)`

SetHashiVaultTargetDetails sets HashiVaultTargetDetails field to given value.

### HasHashiVaultTargetDetails

`func (o *TargetTypeDetailsInput) HasHashiVaultTargetDetails() bool`

HasHashiVaultTargetDetails returns a boolean if a field has been set.

### GetLdapTargetDetails

`func (o *TargetTypeDetailsInput) GetLdapTargetDetails() LdapTargetDetails`

GetLdapTargetDetails returns the LdapTargetDetails field if non-nil, zero value otherwise.

### GetLdapTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetLdapTargetDetailsOk() (*LdapTargetDetails, bool)`

GetLdapTargetDetailsOk returns a tuple with the LdapTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapTargetDetails

`func (o *TargetTypeDetailsInput) SetLdapTargetDetails(v LdapTargetDetails)`

SetLdapTargetDetails sets LdapTargetDetails field to given value.

### HasLdapTargetDetails

`func (o *TargetTypeDetailsInput) HasLdapTargetDetails() bool`

HasLdapTargetDetails returns a boolean if a field has been set.

### GetLetsencryptTargetDetails

`func (o *TargetTypeDetailsInput) GetLetsencryptTargetDetails() LetsEncryptTargetDetails`

GetLetsencryptTargetDetails returns the LetsencryptTargetDetails field if non-nil, zero value otherwise.

### GetLetsencryptTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetLetsencryptTargetDetailsOk() (*LetsEncryptTargetDetails, bool)`

GetLetsencryptTargetDetailsOk returns a tuple with the LetsencryptTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsencryptTargetDetails

`func (o *TargetTypeDetailsInput) SetLetsencryptTargetDetails(v LetsEncryptTargetDetails)`

SetLetsencryptTargetDetails sets LetsencryptTargetDetails field to given value.

### HasLetsencryptTargetDetails

`func (o *TargetTypeDetailsInput) HasLetsencryptTargetDetails() bool`

HasLetsencryptTargetDetails returns a boolean if a field has been set.

### GetLinkedTargetDetails

`func (o *TargetTypeDetailsInput) GetLinkedTargetDetails() LinkedTargetDetails`

GetLinkedTargetDetails returns the LinkedTargetDetails field if non-nil, zero value otherwise.

### GetLinkedTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetLinkedTargetDetailsOk() (*LinkedTargetDetails, bool)`

GetLinkedTargetDetailsOk returns a tuple with the LinkedTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTargetDetails

`func (o *TargetTypeDetailsInput) SetLinkedTargetDetails(v LinkedTargetDetails)`

SetLinkedTargetDetails sets LinkedTargetDetails field to given value.

### HasLinkedTargetDetails

`func (o *TargetTypeDetailsInput) HasLinkedTargetDetails() bool`

HasLinkedTargetDetails returns a boolean if a field has been set.

### GetMongoDbTargetDetails

`func (o *TargetTypeDetailsInput) GetMongoDbTargetDetails() MongoDBTargetDetails`

GetMongoDbTargetDetails returns the MongoDbTargetDetails field if non-nil, zero value otherwise.

### GetMongoDbTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetMongoDbTargetDetailsOk() (*MongoDBTargetDetails, bool)`

GetMongoDbTargetDetailsOk returns a tuple with the MongoDbTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDbTargetDetails

`func (o *TargetTypeDetailsInput) SetMongoDbTargetDetails(v MongoDBTargetDetails)`

SetMongoDbTargetDetails sets MongoDbTargetDetails field to given value.

### HasMongoDbTargetDetails

`func (o *TargetTypeDetailsInput) HasMongoDbTargetDetails() bool`

HasMongoDbTargetDetails returns a boolean if a field has been set.

### GetNativeK8sTargetDetails

`func (o *TargetTypeDetailsInput) GetNativeK8sTargetDetails() NativeK8sTargetDetails`

GetNativeK8sTargetDetails returns the NativeK8sTargetDetails field if non-nil, zero value otherwise.

### GetNativeK8sTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetNativeK8sTargetDetailsOk() (*NativeK8sTargetDetails, bool)`

GetNativeK8sTargetDetailsOk returns a tuple with the NativeK8sTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeK8sTargetDetails

`func (o *TargetTypeDetailsInput) SetNativeK8sTargetDetails(v NativeK8sTargetDetails)`

SetNativeK8sTargetDetails sets NativeK8sTargetDetails field to given value.

### HasNativeK8sTargetDetails

`func (o *TargetTypeDetailsInput) HasNativeK8sTargetDetails() bool`

HasNativeK8sTargetDetails returns a boolean if a field has been set.

### GetOpenaiTargetDetails

`func (o *TargetTypeDetailsInput) GetOpenaiTargetDetails() OpenAITargetDetails`

GetOpenaiTargetDetails returns the OpenaiTargetDetails field if non-nil, zero value otherwise.

### GetOpenaiTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetOpenaiTargetDetailsOk() (*OpenAITargetDetails, bool)`

GetOpenaiTargetDetailsOk returns a tuple with the OpenaiTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenaiTargetDetails

`func (o *TargetTypeDetailsInput) SetOpenaiTargetDetails(v OpenAITargetDetails)`

SetOpenaiTargetDetails sets OpenaiTargetDetails field to given value.

### HasOpenaiTargetDetails

`func (o *TargetTypeDetailsInput) HasOpenaiTargetDetails() bool`

HasOpenaiTargetDetails returns a boolean if a field has been set.

### GetPingTargetDetails

`func (o *TargetTypeDetailsInput) GetPingTargetDetails() PingTargetDetails`

GetPingTargetDetails returns the PingTargetDetails field if non-nil, zero value otherwise.

### GetPingTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetPingTargetDetailsOk() (*PingTargetDetails, bool)`

GetPingTargetDetailsOk returns a tuple with the PingTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTargetDetails

`func (o *TargetTypeDetailsInput) SetPingTargetDetails(v PingTargetDetails)`

SetPingTargetDetails sets PingTargetDetails field to given value.

### HasPingTargetDetails

`func (o *TargetTypeDetailsInput) HasPingTargetDetails() bool`

HasPingTargetDetails returns a boolean if a field has been set.

### GetRabbitMqTargetDetails

`func (o *TargetTypeDetailsInput) GetRabbitMqTargetDetails() RabbitMQTargetDetails`

GetRabbitMqTargetDetails returns the RabbitMqTargetDetails field if non-nil, zero value otherwise.

### GetRabbitMqTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetRabbitMqTargetDetailsOk() (*RabbitMQTargetDetails, bool)`

GetRabbitMqTargetDetailsOk returns a tuple with the RabbitMqTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitMqTargetDetails

`func (o *TargetTypeDetailsInput) SetRabbitMqTargetDetails(v RabbitMQTargetDetails)`

SetRabbitMqTargetDetails sets RabbitMqTargetDetails field to given value.

### HasRabbitMqTargetDetails

`func (o *TargetTypeDetailsInput) HasRabbitMqTargetDetails() bool`

HasRabbitMqTargetDetails returns a boolean if a field has been set.

### GetSalesforceTargetDetails

`func (o *TargetTypeDetailsInput) GetSalesforceTargetDetails() SalesforceTargetDetails`

GetSalesforceTargetDetails returns the SalesforceTargetDetails field if non-nil, zero value otherwise.

### GetSalesforceTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetSalesforceTargetDetailsOk() (*SalesforceTargetDetails, bool)`

GetSalesforceTargetDetailsOk returns a tuple with the SalesforceTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesforceTargetDetails

`func (o *TargetTypeDetailsInput) SetSalesforceTargetDetails(v SalesforceTargetDetails)`

SetSalesforceTargetDetails sets SalesforceTargetDetails field to given value.

### HasSalesforceTargetDetails

`func (o *TargetTypeDetailsInput) HasSalesforceTargetDetails() bool`

HasSalesforceTargetDetails returns a boolean if a field has been set.

### GetSectigoTargetDetails

`func (o *TargetTypeDetailsInput) GetSectigoTargetDetails() SectigoTargetDetails`

GetSectigoTargetDetails returns the SectigoTargetDetails field if non-nil, zero value otherwise.

### GetSectigoTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetSectigoTargetDetailsOk() (*SectigoTargetDetails, bool)`

GetSectigoTargetDetailsOk returns a tuple with the SectigoTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectigoTargetDetails

`func (o *TargetTypeDetailsInput) SetSectigoTargetDetails(v SectigoTargetDetails)`

SetSectigoTargetDetails sets SectigoTargetDetails field to given value.

### HasSectigoTargetDetails

`func (o *TargetTypeDetailsInput) HasSectigoTargetDetails() bool`

HasSectigoTargetDetails returns a boolean if a field has been set.

### GetSplunkTargetDetails

`func (o *TargetTypeDetailsInput) GetSplunkTargetDetails() SplunkTargetDetails`

GetSplunkTargetDetails returns the SplunkTargetDetails field if non-nil, zero value otherwise.

### GetSplunkTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetSplunkTargetDetailsOk() (*SplunkTargetDetails, bool)`

GetSplunkTargetDetailsOk returns a tuple with the SplunkTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkTargetDetails

`func (o *TargetTypeDetailsInput) SetSplunkTargetDetails(v SplunkTargetDetails)`

SetSplunkTargetDetails sets SplunkTargetDetails field to given value.

### HasSplunkTargetDetails

`func (o *TargetTypeDetailsInput) HasSplunkTargetDetails() bool`

HasSplunkTargetDetails returns a boolean if a field has been set.

### GetSshTargetDetails

`func (o *TargetTypeDetailsInput) GetSshTargetDetails() SSHTargetDetails`

GetSshTargetDetails returns the SshTargetDetails field if non-nil, zero value otherwise.

### GetSshTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetSshTargetDetailsOk() (*SSHTargetDetails, bool)`

GetSshTargetDetailsOk returns a tuple with the SshTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshTargetDetails

`func (o *TargetTypeDetailsInput) SetSshTargetDetails(v SSHTargetDetails)`

SetSshTargetDetails sets SshTargetDetails field to given value.

### HasSshTargetDetails

`func (o *TargetTypeDetailsInput) HasSshTargetDetails() bool`

HasSshTargetDetails returns a boolean if a field has been set.

### GetVenafiTargetDetails

`func (o *TargetTypeDetailsInput) GetVenafiTargetDetails() VenafiTargetDetails`

GetVenafiTargetDetails returns the VenafiTargetDetails field if non-nil, zero value otherwise.

### GetVenafiTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetVenafiTargetDetailsOk() (*VenafiTargetDetails, bool)`

GetVenafiTargetDetailsOk returns a tuple with the VenafiTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiTargetDetails

`func (o *TargetTypeDetailsInput) SetVenafiTargetDetails(v VenafiTargetDetails)`

SetVenafiTargetDetails sets VenafiTargetDetails field to given value.

### HasVenafiTargetDetails

`func (o *TargetTypeDetailsInput) HasVenafiTargetDetails() bool`

HasVenafiTargetDetails returns a boolean if a field has been set.

### GetWebTargetDetails

`func (o *TargetTypeDetailsInput) GetWebTargetDetails() WebTargetDetails`

GetWebTargetDetails returns the WebTargetDetails field if non-nil, zero value otherwise.

### GetWebTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetWebTargetDetailsOk() (*WebTargetDetails, bool)`

GetWebTargetDetailsOk returns a tuple with the WebTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTargetDetails

`func (o *TargetTypeDetailsInput) SetWebTargetDetails(v WebTargetDetails)`

SetWebTargetDetails sets WebTargetDetails field to given value.

### HasWebTargetDetails

`func (o *TargetTypeDetailsInput) HasWebTargetDetails() bool`

HasWebTargetDetails returns a boolean if a field has been set.

### GetWindowsTargetDetails

`func (o *TargetTypeDetailsInput) GetWindowsTargetDetails() WindowsTargetDetails`

GetWindowsTargetDetails returns the WindowsTargetDetails field if non-nil, zero value otherwise.

### GetWindowsTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetWindowsTargetDetailsOk() (*WindowsTargetDetails, bool)`

GetWindowsTargetDetailsOk returns a tuple with the WindowsTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsTargetDetails

`func (o *TargetTypeDetailsInput) SetWindowsTargetDetails(v WindowsTargetDetails)`

SetWindowsTargetDetails sets WindowsTargetDetails field to given value.

### HasWindowsTargetDetails

`func (o *TargetTypeDetailsInput) HasWindowsTargetDetails() bool`

HasWindowsTargetDetails returns a boolean if a field has been set.

### GetZerosslTargetDetails

`func (o *TargetTypeDetailsInput) GetZerosslTargetDetails() ZeroSSLTargetDetails`

GetZerosslTargetDetails returns the ZerosslTargetDetails field if non-nil, zero value otherwise.

### GetZerosslTargetDetailsOk

`func (o *TargetTypeDetailsInput) GetZerosslTargetDetailsOk() (*ZeroSSLTargetDetails, bool)`

GetZerosslTargetDetailsOk returns a tuple with the ZerosslTargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZerosslTargetDetails

`func (o *TargetTypeDetailsInput) SetZerosslTargetDetails(v ZeroSSLTargetDetails)`

SetZerosslTargetDetails sets ZerosslTargetDetails field to given value.

### HasZerosslTargetDetails

`func (o *TargetTypeDetailsInput) HasZerosslTargetDetails() bool`

HasZerosslTargetDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


