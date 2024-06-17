/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// TargetTypeDetailsInput struct for TargetTypeDetailsInput
type TargetTypeDetailsInput struct {
	ArtifactoryTargetDetails *ArtifactoryTargetDetails `json:"artifactory_target_details,omitempty"`
	AwsTargetDetails *AWSTargetDetails `json:"aws_target_details,omitempty"`
	AzureTargetDetails *AzureTargetDetails `json:"azure_target_details,omitempty"`
	ChefTargetDetails *ChefTargetDetails `json:"chef_target_details,omitempty"`
	CustomTargetDetails *CustomTargetDetails `json:"custom_target_details,omitempty"`
	DbTargetDetails *DbTargetDetails `json:"db_target_details,omitempty"`
	DockerhubTargetDetails *DockerhubTargetDetails `json:"dockerhub_target_details,omitempty"`
	EksTargetDetails *EKSTargetDetails `json:"eks_target_details,omitempty"`
	GcpTargetDetails *GcpTargetDetails `json:"gcp_target_details,omitempty"`
	GithubTargetDetails *GithubTargetDetails `json:"github_target_details,omitempty"`
	GkeTargetDetails *GKETargetDetails `json:"gke_target_details,omitempty"`
	GlobalsignAtlasTargetDetails *GlobalSignAtlasTargetDetails `json:"globalsign_atlas_target_details,omitempty"`
	GlobalsignTargetDetails *GlobalSignGCCTargetDetails `json:"globalsign_target_details,omitempty"`
	GodaddyTargetDetails *GodaddyTargetDetails `json:"godaddy_target_details,omitempty"`
	HashiVaultTargetDetails *HashiVaultTargetDetails `json:"hashi_vault_target_details,omitempty"`
	LdapTargetDetails *LdapTargetDetails `json:"ldap_target_details,omitempty"`
	LinkedTargetDetails *LinkedTargetDetails `json:"linked_target_details,omitempty"`
	MongoDbTargetDetails *MongoDBTargetDetails `json:"mongo_db_target_details,omitempty"`
	NativeK8sTargetDetails *NativeK8sTargetDetails `json:"native_k8s_target_details,omitempty"`
	PingTargetDetails *PingTargetDetails `json:"ping_target_details,omitempty"`
	RabbitMqTargetDetails *RabbitMQTargetDetails `json:"rabbit_mq_target_details,omitempty"`
	SalesforceTargetDetails *SalesforceTargetDetails `json:"salesforce_target_details,omitempty"`
	SshTargetDetails *SSHTargetDetails `json:"ssh_target_details,omitempty"`
	VenafiTargetDetails *VenafiTargetDetails `json:"venafi_target_details,omitempty"`
	WebTargetDetails *WebTargetDetails `json:"web_target_details,omitempty"`
	WindowsTargetDetails *WindowsTargetDetails `json:"windows_target_details,omitempty"`
	ZerosslTargetDetails *ZeroSSLTargetDetails `json:"zerossl_target_details,omitempty"`
}

// NewTargetTypeDetailsInput instantiates a new TargetTypeDetailsInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetTypeDetailsInput() *TargetTypeDetailsInput {
	this := TargetTypeDetailsInput{}
	return &this
}

// NewTargetTypeDetailsInputWithDefaults instantiates a new TargetTypeDetailsInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetTypeDetailsInputWithDefaults() *TargetTypeDetailsInput {
	this := TargetTypeDetailsInput{}
	return &this
}

// GetArtifactoryTargetDetails returns the ArtifactoryTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetArtifactoryTargetDetails() ArtifactoryTargetDetails {
	if o == nil || o.ArtifactoryTargetDetails == nil {
		var ret ArtifactoryTargetDetails
		return ret
	}
	return *o.ArtifactoryTargetDetails
}

// GetArtifactoryTargetDetailsOk returns a tuple with the ArtifactoryTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetArtifactoryTargetDetailsOk() (*ArtifactoryTargetDetails, bool) {
	if o == nil || o.ArtifactoryTargetDetails == nil {
		return nil, false
	}
	return o.ArtifactoryTargetDetails, true
}

// HasArtifactoryTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasArtifactoryTargetDetails() bool {
	if o != nil && o.ArtifactoryTargetDetails != nil {
		return true
	}

	return false
}

// SetArtifactoryTargetDetails gets a reference to the given ArtifactoryTargetDetails and assigns it to the ArtifactoryTargetDetails field.
func (o *TargetTypeDetailsInput) SetArtifactoryTargetDetails(v ArtifactoryTargetDetails) {
	o.ArtifactoryTargetDetails = &v
}

// GetAwsTargetDetails returns the AwsTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetAwsTargetDetails() AWSTargetDetails {
	if o == nil || o.AwsTargetDetails == nil {
		var ret AWSTargetDetails
		return ret
	}
	return *o.AwsTargetDetails
}

// GetAwsTargetDetailsOk returns a tuple with the AwsTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetAwsTargetDetailsOk() (*AWSTargetDetails, bool) {
	if o == nil || o.AwsTargetDetails == nil {
		return nil, false
	}
	return o.AwsTargetDetails, true
}

// HasAwsTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasAwsTargetDetails() bool {
	if o != nil && o.AwsTargetDetails != nil {
		return true
	}

	return false
}

// SetAwsTargetDetails gets a reference to the given AWSTargetDetails and assigns it to the AwsTargetDetails field.
func (o *TargetTypeDetailsInput) SetAwsTargetDetails(v AWSTargetDetails) {
	o.AwsTargetDetails = &v
}

// GetAzureTargetDetails returns the AzureTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetAzureTargetDetails() AzureTargetDetails {
	if o == nil || o.AzureTargetDetails == nil {
		var ret AzureTargetDetails
		return ret
	}
	return *o.AzureTargetDetails
}

// GetAzureTargetDetailsOk returns a tuple with the AzureTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetAzureTargetDetailsOk() (*AzureTargetDetails, bool) {
	if o == nil || o.AzureTargetDetails == nil {
		return nil, false
	}
	return o.AzureTargetDetails, true
}

// HasAzureTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasAzureTargetDetails() bool {
	if o != nil && o.AzureTargetDetails != nil {
		return true
	}

	return false
}

// SetAzureTargetDetails gets a reference to the given AzureTargetDetails and assigns it to the AzureTargetDetails field.
func (o *TargetTypeDetailsInput) SetAzureTargetDetails(v AzureTargetDetails) {
	o.AzureTargetDetails = &v
}

// GetChefTargetDetails returns the ChefTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetChefTargetDetails() ChefTargetDetails {
	if o == nil || o.ChefTargetDetails == nil {
		var ret ChefTargetDetails
		return ret
	}
	return *o.ChefTargetDetails
}

// GetChefTargetDetailsOk returns a tuple with the ChefTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetChefTargetDetailsOk() (*ChefTargetDetails, bool) {
	if o == nil || o.ChefTargetDetails == nil {
		return nil, false
	}
	return o.ChefTargetDetails, true
}

// HasChefTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasChefTargetDetails() bool {
	if o != nil && o.ChefTargetDetails != nil {
		return true
	}

	return false
}

// SetChefTargetDetails gets a reference to the given ChefTargetDetails and assigns it to the ChefTargetDetails field.
func (o *TargetTypeDetailsInput) SetChefTargetDetails(v ChefTargetDetails) {
	o.ChefTargetDetails = &v
}

// GetCustomTargetDetails returns the CustomTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetCustomTargetDetails() CustomTargetDetails {
	if o == nil || o.CustomTargetDetails == nil {
		var ret CustomTargetDetails
		return ret
	}
	return *o.CustomTargetDetails
}

// GetCustomTargetDetailsOk returns a tuple with the CustomTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetCustomTargetDetailsOk() (*CustomTargetDetails, bool) {
	if o == nil || o.CustomTargetDetails == nil {
		return nil, false
	}
	return o.CustomTargetDetails, true
}

// HasCustomTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasCustomTargetDetails() bool {
	if o != nil && o.CustomTargetDetails != nil {
		return true
	}

	return false
}

// SetCustomTargetDetails gets a reference to the given CustomTargetDetails and assigns it to the CustomTargetDetails field.
func (o *TargetTypeDetailsInput) SetCustomTargetDetails(v CustomTargetDetails) {
	o.CustomTargetDetails = &v
}

// GetDbTargetDetails returns the DbTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetDbTargetDetails() DbTargetDetails {
	if o == nil || o.DbTargetDetails == nil {
		var ret DbTargetDetails
		return ret
	}
	return *o.DbTargetDetails
}

// GetDbTargetDetailsOk returns a tuple with the DbTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetDbTargetDetailsOk() (*DbTargetDetails, bool) {
	if o == nil || o.DbTargetDetails == nil {
		return nil, false
	}
	return o.DbTargetDetails, true
}

// HasDbTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasDbTargetDetails() bool {
	if o != nil && o.DbTargetDetails != nil {
		return true
	}

	return false
}

// SetDbTargetDetails gets a reference to the given DbTargetDetails and assigns it to the DbTargetDetails field.
func (o *TargetTypeDetailsInput) SetDbTargetDetails(v DbTargetDetails) {
	o.DbTargetDetails = &v
}

// GetDockerhubTargetDetails returns the DockerhubTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetDockerhubTargetDetails() DockerhubTargetDetails {
	if o == nil || o.DockerhubTargetDetails == nil {
		var ret DockerhubTargetDetails
		return ret
	}
	return *o.DockerhubTargetDetails
}

// GetDockerhubTargetDetailsOk returns a tuple with the DockerhubTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetDockerhubTargetDetailsOk() (*DockerhubTargetDetails, bool) {
	if o == nil || o.DockerhubTargetDetails == nil {
		return nil, false
	}
	return o.DockerhubTargetDetails, true
}

// HasDockerhubTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasDockerhubTargetDetails() bool {
	if o != nil && o.DockerhubTargetDetails != nil {
		return true
	}

	return false
}

// SetDockerhubTargetDetails gets a reference to the given DockerhubTargetDetails and assigns it to the DockerhubTargetDetails field.
func (o *TargetTypeDetailsInput) SetDockerhubTargetDetails(v DockerhubTargetDetails) {
	o.DockerhubTargetDetails = &v
}

// GetEksTargetDetails returns the EksTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetEksTargetDetails() EKSTargetDetails {
	if o == nil || o.EksTargetDetails == nil {
		var ret EKSTargetDetails
		return ret
	}
	return *o.EksTargetDetails
}

// GetEksTargetDetailsOk returns a tuple with the EksTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetEksTargetDetailsOk() (*EKSTargetDetails, bool) {
	if o == nil || o.EksTargetDetails == nil {
		return nil, false
	}
	return o.EksTargetDetails, true
}

// HasEksTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasEksTargetDetails() bool {
	if o != nil && o.EksTargetDetails != nil {
		return true
	}

	return false
}

// SetEksTargetDetails gets a reference to the given EKSTargetDetails and assigns it to the EksTargetDetails field.
func (o *TargetTypeDetailsInput) SetEksTargetDetails(v EKSTargetDetails) {
	o.EksTargetDetails = &v
}

// GetGcpTargetDetails returns the GcpTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetGcpTargetDetails() GcpTargetDetails {
	if o == nil || o.GcpTargetDetails == nil {
		var ret GcpTargetDetails
		return ret
	}
	return *o.GcpTargetDetails
}

// GetGcpTargetDetailsOk returns a tuple with the GcpTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetGcpTargetDetailsOk() (*GcpTargetDetails, bool) {
	if o == nil || o.GcpTargetDetails == nil {
		return nil, false
	}
	return o.GcpTargetDetails, true
}

// HasGcpTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasGcpTargetDetails() bool {
	if o != nil && o.GcpTargetDetails != nil {
		return true
	}

	return false
}

// SetGcpTargetDetails gets a reference to the given GcpTargetDetails and assigns it to the GcpTargetDetails field.
func (o *TargetTypeDetailsInput) SetGcpTargetDetails(v GcpTargetDetails) {
	o.GcpTargetDetails = &v
}

// GetGithubTargetDetails returns the GithubTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetGithubTargetDetails() GithubTargetDetails {
	if o == nil || o.GithubTargetDetails == nil {
		var ret GithubTargetDetails
		return ret
	}
	return *o.GithubTargetDetails
}

// GetGithubTargetDetailsOk returns a tuple with the GithubTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetGithubTargetDetailsOk() (*GithubTargetDetails, bool) {
	if o == nil || o.GithubTargetDetails == nil {
		return nil, false
	}
	return o.GithubTargetDetails, true
}

// HasGithubTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasGithubTargetDetails() bool {
	if o != nil && o.GithubTargetDetails != nil {
		return true
	}

	return false
}

// SetGithubTargetDetails gets a reference to the given GithubTargetDetails and assigns it to the GithubTargetDetails field.
func (o *TargetTypeDetailsInput) SetGithubTargetDetails(v GithubTargetDetails) {
	o.GithubTargetDetails = &v
}

// GetGkeTargetDetails returns the GkeTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetGkeTargetDetails() GKETargetDetails {
	if o == nil || o.GkeTargetDetails == nil {
		var ret GKETargetDetails
		return ret
	}
	return *o.GkeTargetDetails
}

// GetGkeTargetDetailsOk returns a tuple with the GkeTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetGkeTargetDetailsOk() (*GKETargetDetails, bool) {
	if o == nil || o.GkeTargetDetails == nil {
		return nil, false
	}
	return o.GkeTargetDetails, true
}

// HasGkeTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasGkeTargetDetails() bool {
	if o != nil && o.GkeTargetDetails != nil {
		return true
	}

	return false
}

// SetGkeTargetDetails gets a reference to the given GKETargetDetails and assigns it to the GkeTargetDetails field.
func (o *TargetTypeDetailsInput) SetGkeTargetDetails(v GKETargetDetails) {
	o.GkeTargetDetails = &v
}

// GetGlobalsignAtlasTargetDetails returns the GlobalsignAtlasTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetGlobalsignAtlasTargetDetails() GlobalSignAtlasTargetDetails {
	if o == nil || o.GlobalsignAtlasTargetDetails == nil {
		var ret GlobalSignAtlasTargetDetails
		return ret
	}
	return *o.GlobalsignAtlasTargetDetails
}

// GetGlobalsignAtlasTargetDetailsOk returns a tuple with the GlobalsignAtlasTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetGlobalsignAtlasTargetDetailsOk() (*GlobalSignAtlasTargetDetails, bool) {
	if o == nil || o.GlobalsignAtlasTargetDetails == nil {
		return nil, false
	}
	return o.GlobalsignAtlasTargetDetails, true
}

// HasGlobalsignAtlasTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasGlobalsignAtlasTargetDetails() bool {
	if o != nil && o.GlobalsignAtlasTargetDetails != nil {
		return true
	}

	return false
}

// SetGlobalsignAtlasTargetDetails gets a reference to the given GlobalSignAtlasTargetDetails and assigns it to the GlobalsignAtlasTargetDetails field.
func (o *TargetTypeDetailsInput) SetGlobalsignAtlasTargetDetails(v GlobalSignAtlasTargetDetails) {
	o.GlobalsignAtlasTargetDetails = &v
}

// GetGlobalsignTargetDetails returns the GlobalsignTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetGlobalsignTargetDetails() GlobalSignGCCTargetDetails {
	if o == nil || o.GlobalsignTargetDetails == nil {
		var ret GlobalSignGCCTargetDetails
		return ret
	}
	return *o.GlobalsignTargetDetails
}

// GetGlobalsignTargetDetailsOk returns a tuple with the GlobalsignTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetGlobalsignTargetDetailsOk() (*GlobalSignGCCTargetDetails, bool) {
	if o == nil || o.GlobalsignTargetDetails == nil {
		return nil, false
	}
	return o.GlobalsignTargetDetails, true
}

// HasGlobalsignTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasGlobalsignTargetDetails() bool {
	if o != nil && o.GlobalsignTargetDetails != nil {
		return true
	}

	return false
}

// SetGlobalsignTargetDetails gets a reference to the given GlobalSignGCCTargetDetails and assigns it to the GlobalsignTargetDetails field.
func (o *TargetTypeDetailsInput) SetGlobalsignTargetDetails(v GlobalSignGCCTargetDetails) {
	o.GlobalsignTargetDetails = &v
}

// GetGodaddyTargetDetails returns the GodaddyTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetGodaddyTargetDetails() GodaddyTargetDetails {
	if o == nil || o.GodaddyTargetDetails == nil {
		var ret GodaddyTargetDetails
		return ret
	}
	return *o.GodaddyTargetDetails
}

// GetGodaddyTargetDetailsOk returns a tuple with the GodaddyTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetGodaddyTargetDetailsOk() (*GodaddyTargetDetails, bool) {
	if o == nil || o.GodaddyTargetDetails == nil {
		return nil, false
	}
	return o.GodaddyTargetDetails, true
}

// HasGodaddyTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasGodaddyTargetDetails() bool {
	if o != nil && o.GodaddyTargetDetails != nil {
		return true
	}

	return false
}

// SetGodaddyTargetDetails gets a reference to the given GodaddyTargetDetails and assigns it to the GodaddyTargetDetails field.
func (o *TargetTypeDetailsInput) SetGodaddyTargetDetails(v GodaddyTargetDetails) {
	o.GodaddyTargetDetails = &v
}

// GetHashiVaultTargetDetails returns the HashiVaultTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetHashiVaultTargetDetails() HashiVaultTargetDetails {
	if o == nil || o.HashiVaultTargetDetails == nil {
		var ret HashiVaultTargetDetails
		return ret
	}
	return *o.HashiVaultTargetDetails
}

// GetHashiVaultTargetDetailsOk returns a tuple with the HashiVaultTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetHashiVaultTargetDetailsOk() (*HashiVaultTargetDetails, bool) {
	if o == nil || o.HashiVaultTargetDetails == nil {
		return nil, false
	}
	return o.HashiVaultTargetDetails, true
}

// HasHashiVaultTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasHashiVaultTargetDetails() bool {
	if o != nil && o.HashiVaultTargetDetails != nil {
		return true
	}

	return false
}

// SetHashiVaultTargetDetails gets a reference to the given HashiVaultTargetDetails and assigns it to the HashiVaultTargetDetails field.
func (o *TargetTypeDetailsInput) SetHashiVaultTargetDetails(v HashiVaultTargetDetails) {
	o.HashiVaultTargetDetails = &v
}

// GetLdapTargetDetails returns the LdapTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetLdapTargetDetails() LdapTargetDetails {
	if o == nil || o.LdapTargetDetails == nil {
		var ret LdapTargetDetails
		return ret
	}
	return *o.LdapTargetDetails
}

// GetLdapTargetDetailsOk returns a tuple with the LdapTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetLdapTargetDetailsOk() (*LdapTargetDetails, bool) {
	if o == nil || o.LdapTargetDetails == nil {
		return nil, false
	}
	return o.LdapTargetDetails, true
}

// HasLdapTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasLdapTargetDetails() bool {
	if o != nil && o.LdapTargetDetails != nil {
		return true
	}

	return false
}

// SetLdapTargetDetails gets a reference to the given LdapTargetDetails and assigns it to the LdapTargetDetails field.
func (o *TargetTypeDetailsInput) SetLdapTargetDetails(v LdapTargetDetails) {
	o.LdapTargetDetails = &v
}

// GetLinkedTargetDetails returns the LinkedTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetLinkedTargetDetails() LinkedTargetDetails {
	if o == nil || o.LinkedTargetDetails == nil {
		var ret LinkedTargetDetails
		return ret
	}
	return *o.LinkedTargetDetails
}

// GetLinkedTargetDetailsOk returns a tuple with the LinkedTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetLinkedTargetDetailsOk() (*LinkedTargetDetails, bool) {
	if o == nil || o.LinkedTargetDetails == nil {
		return nil, false
	}
	return o.LinkedTargetDetails, true
}

// HasLinkedTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasLinkedTargetDetails() bool {
	if o != nil && o.LinkedTargetDetails != nil {
		return true
	}

	return false
}

// SetLinkedTargetDetails gets a reference to the given LinkedTargetDetails and assigns it to the LinkedTargetDetails field.
func (o *TargetTypeDetailsInput) SetLinkedTargetDetails(v LinkedTargetDetails) {
	o.LinkedTargetDetails = &v
}

// GetMongoDbTargetDetails returns the MongoDbTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetMongoDbTargetDetails() MongoDBTargetDetails {
	if o == nil || o.MongoDbTargetDetails == nil {
		var ret MongoDBTargetDetails
		return ret
	}
	return *o.MongoDbTargetDetails
}

// GetMongoDbTargetDetailsOk returns a tuple with the MongoDbTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetMongoDbTargetDetailsOk() (*MongoDBTargetDetails, bool) {
	if o == nil || o.MongoDbTargetDetails == nil {
		return nil, false
	}
	return o.MongoDbTargetDetails, true
}

// HasMongoDbTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasMongoDbTargetDetails() bool {
	if o != nil && o.MongoDbTargetDetails != nil {
		return true
	}

	return false
}

// SetMongoDbTargetDetails gets a reference to the given MongoDBTargetDetails and assigns it to the MongoDbTargetDetails field.
func (o *TargetTypeDetailsInput) SetMongoDbTargetDetails(v MongoDBTargetDetails) {
	o.MongoDbTargetDetails = &v
}

// GetNativeK8sTargetDetails returns the NativeK8sTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetNativeK8sTargetDetails() NativeK8sTargetDetails {
	if o == nil || o.NativeK8sTargetDetails == nil {
		var ret NativeK8sTargetDetails
		return ret
	}
	return *o.NativeK8sTargetDetails
}

// GetNativeK8sTargetDetailsOk returns a tuple with the NativeK8sTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetNativeK8sTargetDetailsOk() (*NativeK8sTargetDetails, bool) {
	if o == nil || o.NativeK8sTargetDetails == nil {
		return nil, false
	}
	return o.NativeK8sTargetDetails, true
}

// HasNativeK8sTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasNativeK8sTargetDetails() bool {
	if o != nil && o.NativeK8sTargetDetails != nil {
		return true
	}

	return false
}

// SetNativeK8sTargetDetails gets a reference to the given NativeK8sTargetDetails and assigns it to the NativeK8sTargetDetails field.
func (o *TargetTypeDetailsInput) SetNativeK8sTargetDetails(v NativeK8sTargetDetails) {
	o.NativeK8sTargetDetails = &v
}

// GetPingTargetDetails returns the PingTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetPingTargetDetails() PingTargetDetails {
	if o == nil || o.PingTargetDetails == nil {
		var ret PingTargetDetails
		return ret
	}
	return *o.PingTargetDetails
}

// GetPingTargetDetailsOk returns a tuple with the PingTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetPingTargetDetailsOk() (*PingTargetDetails, bool) {
	if o == nil || o.PingTargetDetails == nil {
		return nil, false
	}
	return o.PingTargetDetails, true
}

// HasPingTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasPingTargetDetails() bool {
	if o != nil && o.PingTargetDetails != nil {
		return true
	}

	return false
}

// SetPingTargetDetails gets a reference to the given PingTargetDetails and assigns it to the PingTargetDetails field.
func (o *TargetTypeDetailsInput) SetPingTargetDetails(v PingTargetDetails) {
	o.PingTargetDetails = &v
}

// GetRabbitMqTargetDetails returns the RabbitMqTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetRabbitMqTargetDetails() RabbitMQTargetDetails {
	if o == nil || o.RabbitMqTargetDetails == nil {
		var ret RabbitMQTargetDetails
		return ret
	}
	return *o.RabbitMqTargetDetails
}

// GetRabbitMqTargetDetailsOk returns a tuple with the RabbitMqTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetRabbitMqTargetDetailsOk() (*RabbitMQTargetDetails, bool) {
	if o == nil || o.RabbitMqTargetDetails == nil {
		return nil, false
	}
	return o.RabbitMqTargetDetails, true
}

// HasRabbitMqTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasRabbitMqTargetDetails() bool {
	if o != nil && o.RabbitMqTargetDetails != nil {
		return true
	}

	return false
}

// SetRabbitMqTargetDetails gets a reference to the given RabbitMQTargetDetails and assigns it to the RabbitMqTargetDetails field.
func (o *TargetTypeDetailsInput) SetRabbitMqTargetDetails(v RabbitMQTargetDetails) {
	o.RabbitMqTargetDetails = &v
}

// GetSalesforceTargetDetails returns the SalesforceTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetSalesforceTargetDetails() SalesforceTargetDetails {
	if o == nil || o.SalesforceTargetDetails == nil {
		var ret SalesforceTargetDetails
		return ret
	}
	return *o.SalesforceTargetDetails
}

// GetSalesforceTargetDetailsOk returns a tuple with the SalesforceTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetSalesforceTargetDetailsOk() (*SalesforceTargetDetails, bool) {
	if o == nil || o.SalesforceTargetDetails == nil {
		return nil, false
	}
	return o.SalesforceTargetDetails, true
}

// HasSalesforceTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasSalesforceTargetDetails() bool {
	if o != nil && o.SalesforceTargetDetails != nil {
		return true
	}

	return false
}

// SetSalesforceTargetDetails gets a reference to the given SalesforceTargetDetails and assigns it to the SalesforceTargetDetails field.
func (o *TargetTypeDetailsInput) SetSalesforceTargetDetails(v SalesforceTargetDetails) {
	o.SalesforceTargetDetails = &v
}

// GetSshTargetDetails returns the SshTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetSshTargetDetails() SSHTargetDetails {
	if o == nil || o.SshTargetDetails == nil {
		var ret SSHTargetDetails
		return ret
	}
	return *o.SshTargetDetails
}

// GetSshTargetDetailsOk returns a tuple with the SshTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetSshTargetDetailsOk() (*SSHTargetDetails, bool) {
	if o == nil || o.SshTargetDetails == nil {
		return nil, false
	}
	return o.SshTargetDetails, true
}

// HasSshTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasSshTargetDetails() bool {
	if o != nil && o.SshTargetDetails != nil {
		return true
	}

	return false
}

// SetSshTargetDetails gets a reference to the given SSHTargetDetails and assigns it to the SshTargetDetails field.
func (o *TargetTypeDetailsInput) SetSshTargetDetails(v SSHTargetDetails) {
	o.SshTargetDetails = &v
}

// GetVenafiTargetDetails returns the VenafiTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetVenafiTargetDetails() VenafiTargetDetails {
	if o == nil || o.VenafiTargetDetails == nil {
		var ret VenafiTargetDetails
		return ret
	}
	return *o.VenafiTargetDetails
}

// GetVenafiTargetDetailsOk returns a tuple with the VenafiTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetVenafiTargetDetailsOk() (*VenafiTargetDetails, bool) {
	if o == nil || o.VenafiTargetDetails == nil {
		return nil, false
	}
	return o.VenafiTargetDetails, true
}

// HasVenafiTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasVenafiTargetDetails() bool {
	if o != nil && o.VenafiTargetDetails != nil {
		return true
	}

	return false
}

// SetVenafiTargetDetails gets a reference to the given VenafiTargetDetails and assigns it to the VenafiTargetDetails field.
func (o *TargetTypeDetailsInput) SetVenafiTargetDetails(v VenafiTargetDetails) {
	o.VenafiTargetDetails = &v
}

// GetWebTargetDetails returns the WebTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetWebTargetDetails() WebTargetDetails {
	if o == nil || o.WebTargetDetails == nil {
		var ret WebTargetDetails
		return ret
	}
	return *o.WebTargetDetails
}

// GetWebTargetDetailsOk returns a tuple with the WebTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetWebTargetDetailsOk() (*WebTargetDetails, bool) {
	if o == nil || o.WebTargetDetails == nil {
		return nil, false
	}
	return o.WebTargetDetails, true
}

// HasWebTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasWebTargetDetails() bool {
	if o != nil && o.WebTargetDetails != nil {
		return true
	}

	return false
}

// SetWebTargetDetails gets a reference to the given WebTargetDetails and assigns it to the WebTargetDetails field.
func (o *TargetTypeDetailsInput) SetWebTargetDetails(v WebTargetDetails) {
	o.WebTargetDetails = &v
}

// GetWindowsTargetDetails returns the WindowsTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetWindowsTargetDetails() WindowsTargetDetails {
	if o == nil || o.WindowsTargetDetails == nil {
		var ret WindowsTargetDetails
		return ret
	}
	return *o.WindowsTargetDetails
}

// GetWindowsTargetDetailsOk returns a tuple with the WindowsTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetWindowsTargetDetailsOk() (*WindowsTargetDetails, bool) {
	if o == nil || o.WindowsTargetDetails == nil {
		return nil, false
	}
	return o.WindowsTargetDetails, true
}

// HasWindowsTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasWindowsTargetDetails() bool {
	if o != nil && o.WindowsTargetDetails != nil {
		return true
	}

	return false
}

// SetWindowsTargetDetails gets a reference to the given WindowsTargetDetails and assigns it to the WindowsTargetDetails field.
func (o *TargetTypeDetailsInput) SetWindowsTargetDetails(v WindowsTargetDetails) {
	o.WindowsTargetDetails = &v
}

// GetZerosslTargetDetails returns the ZerosslTargetDetails field value if set, zero value otherwise.
func (o *TargetTypeDetailsInput) GetZerosslTargetDetails() ZeroSSLTargetDetails {
	if o == nil || o.ZerosslTargetDetails == nil {
		var ret ZeroSSLTargetDetails
		return ret
	}
	return *o.ZerosslTargetDetails
}

// GetZerosslTargetDetailsOk returns a tuple with the ZerosslTargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTypeDetailsInput) GetZerosslTargetDetailsOk() (*ZeroSSLTargetDetails, bool) {
	if o == nil || o.ZerosslTargetDetails == nil {
		return nil, false
	}
	return o.ZerosslTargetDetails, true
}

// HasZerosslTargetDetails returns a boolean if a field has been set.
func (o *TargetTypeDetailsInput) HasZerosslTargetDetails() bool {
	if o != nil && o.ZerosslTargetDetails != nil {
		return true
	}

	return false
}

// SetZerosslTargetDetails gets a reference to the given ZeroSSLTargetDetails and assigns it to the ZerosslTargetDetails field.
func (o *TargetTypeDetailsInput) SetZerosslTargetDetails(v ZeroSSLTargetDetails) {
	o.ZerosslTargetDetails = &v
}

func (o TargetTypeDetailsInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactoryTargetDetails != nil {
		toSerialize["artifactory_target_details"] = o.ArtifactoryTargetDetails
	}
	if o.AwsTargetDetails != nil {
		toSerialize["aws_target_details"] = o.AwsTargetDetails
	}
	if o.AzureTargetDetails != nil {
		toSerialize["azure_target_details"] = o.AzureTargetDetails
	}
	if o.ChefTargetDetails != nil {
		toSerialize["chef_target_details"] = o.ChefTargetDetails
	}
	if o.CustomTargetDetails != nil {
		toSerialize["custom_target_details"] = o.CustomTargetDetails
	}
	if o.DbTargetDetails != nil {
		toSerialize["db_target_details"] = o.DbTargetDetails
	}
	if o.DockerhubTargetDetails != nil {
		toSerialize["dockerhub_target_details"] = o.DockerhubTargetDetails
	}
	if o.EksTargetDetails != nil {
		toSerialize["eks_target_details"] = o.EksTargetDetails
	}
	if o.GcpTargetDetails != nil {
		toSerialize["gcp_target_details"] = o.GcpTargetDetails
	}
	if o.GithubTargetDetails != nil {
		toSerialize["github_target_details"] = o.GithubTargetDetails
	}
	if o.GkeTargetDetails != nil {
		toSerialize["gke_target_details"] = o.GkeTargetDetails
	}
	if o.GlobalsignAtlasTargetDetails != nil {
		toSerialize["globalsign_atlas_target_details"] = o.GlobalsignAtlasTargetDetails
	}
	if o.GlobalsignTargetDetails != nil {
		toSerialize["globalsign_target_details"] = o.GlobalsignTargetDetails
	}
	if o.GodaddyTargetDetails != nil {
		toSerialize["godaddy_target_details"] = o.GodaddyTargetDetails
	}
	if o.HashiVaultTargetDetails != nil {
		toSerialize["hashi_vault_target_details"] = o.HashiVaultTargetDetails
	}
	if o.LdapTargetDetails != nil {
		toSerialize["ldap_target_details"] = o.LdapTargetDetails
	}
	if o.LinkedTargetDetails != nil {
		toSerialize["linked_target_details"] = o.LinkedTargetDetails
	}
	if o.MongoDbTargetDetails != nil {
		toSerialize["mongo_db_target_details"] = o.MongoDbTargetDetails
	}
	if o.NativeK8sTargetDetails != nil {
		toSerialize["native_k8s_target_details"] = o.NativeK8sTargetDetails
	}
	if o.PingTargetDetails != nil {
		toSerialize["ping_target_details"] = o.PingTargetDetails
	}
	if o.RabbitMqTargetDetails != nil {
		toSerialize["rabbit_mq_target_details"] = o.RabbitMqTargetDetails
	}
	if o.SalesforceTargetDetails != nil {
		toSerialize["salesforce_target_details"] = o.SalesforceTargetDetails
	}
	if o.SshTargetDetails != nil {
		toSerialize["ssh_target_details"] = o.SshTargetDetails
	}
	if o.VenafiTargetDetails != nil {
		toSerialize["venafi_target_details"] = o.VenafiTargetDetails
	}
	if o.WebTargetDetails != nil {
		toSerialize["web_target_details"] = o.WebTargetDetails
	}
	if o.WindowsTargetDetails != nil {
		toSerialize["windows_target_details"] = o.WindowsTargetDetails
	}
	if o.ZerosslTargetDetails != nil {
		toSerialize["zerossl_target_details"] = o.ZerosslTargetDetails
	}
	return json.Marshal(toSerialize)
}

type NullableTargetTypeDetailsInput struct {
	value *TargetTypeDetailsInput
	isSet bool
}

func (v NullableTargetTypeDetailsInput) Get() *TargetTypeDetailsInput {
	return v.value
}

func (v *NullableTargetTypeDetailsInput) Set(val *TargetTypeDetailsInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetTypeDetailsInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetTypeDetailsInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetTypeDetailsInput(val *TargetTypeDetailsInput) *NullableTargetTypeDetailsInput {
	return &NullableTargetTypeDetailsInput{value: val, isSet: true}
}

func (v NullableTargetTypeDetailsInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetTypeDetailsInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


