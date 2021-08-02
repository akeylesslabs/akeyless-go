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

// UpdateItem struct for UpdateItem
type UpdateItem struct {
	// List of the new tags that will be attached to this item
	AddTag *[]string `json:"add-tag,omitempty"`
	// Current item name
	Name string `json:"name"`
	// New item metadata
	NewMetadata *string `json:"new-metadata,omitempty"`
	// New item name
	NewName *string `json:"new-name,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// List of the existent tags that will be removed from this item
	RmTag *[]string `json:"rm-tag,omitempty"`
	SecureAccessAllowPortForwading *bool `json:"secure-access-allow-port-forwading,omitempty"`
	SecureAccessAwsAccountId *string `json:"secure-access-aws-account-id,omitempty"`
	SecureAccessAwsNativeCli *bool `json:"secure-access-aws-native-cli,omitempty"`
	SecureAccessAwsRegion *string `json:"secure-access-aws-region,omitempty"`
	SecureAccessBastionApi *string `json:"secure-access-bastion-api,omitempty"`
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	SecureAccessBastionSsh *string `json:"secure-access-bastion-ssh,omitempty"`
	SecureAccessClusterEndpoint *string `json:"secure-access-cluster-endpoint,omitempty"`
	SecureAccessDashboardUrl *string `json:"secure-access-dashboard-url,omitempty"`
	SecureAccessDbName *string `json:"secure-access-db-name,omitempty"`
	SecureAccessDbSchema *string `json:"secure-access-db-schema,omitempty"`
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	SecureAccessRdpDomain *string `json:"secure-access-rdp-domain,omitempty"`
	SecureAccessSshCreds *string `json:"secure-access-ssh-creds,omitempty"`
	SecureAccessSshCredsUser *string `json:"secure-access-ssh-creds-user,omitempty"`
	SecureAccessUrl *string `json:"secure-access-url,omitempty"`
	SecureAccessWebBrowsing *bool `json:"secure-access-web-browsing,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewUpdateItem instantiates a new UpdateItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateItem(name string, ) *UpdateItem {
	this := UpdateItem{}
	this.Name = name
	var newMetadata string = "default_metadata"
	this.NewMetadata = &newMetadata
	return &this
}

// NewUpdateItemWithDefaults instantiates a new UpdateItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateItemWithDefaults() *UpdateItem {
	this := UpdateItem{}
	var newMetadata string = "default_metadata"
	this.NewMetadata = &newMetadata
	return &this
}

// GetAddTag returns the AddTag field value if set, zero value otherwise.
func (o *UpdateItem) GetAddTag() []string {
	if o == nil || o.AddTag == nil {
		var ret []string
		return ret
	}
	return *o.AddTag
}

// GetAddTagOk returns a tuple with the AddTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetAddTagOk() (*[]string, bool) {
	if o == nil || o.AddTag == nil {
		return nil, false
	}
	return o.AddTag, true
}

// HasAddTag returns a boolean if a field has been set.
func (o *UpdateItem) HasAddTag() bool {
	if o != nil && o.AddTag != nil {
		return true
	}

	return false
}

// SetAddTag gets a reference to the given []string and assigns it to the AddTag field.
func (o *UpdateItem) SetAddTag(v []string) {
	o.AddTag = &v
}

// GetName returns the Name field value
func (o *UpdateItem) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateItem) SetName(v string) {
	o.Name = v
}

// GetNewMetadata returns the NewMetadata field value if set, zero value otherwise.
func (o *UpdateItem) GetNewMetadata() string {
	if o == nil || o.NewMetadata == nil {
		var ret string
		return ret
	}
	return *o.NewMetadata
}

// GetNewMetadataOk returns a tuple with the NewMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetNewMetadataOk() (*string, bool) {
	if o == nil || o.NewMetadata == nil {
		return nil, false
	}
	return o.NewMetadata, true
}

// HasNewMetadata returns a boolean if a field has been set.
func (o *UpdateItem) HasNewMetadata() bool {
	if o != nil && o.NewMetadata != nil {
		return true
	}

	return false
}

// SetNewMetadata gets a reference to the given string and assigns it to the NewMetadata field.
func (o *UpdateItem) SetNewMetadata(v string) {
	o.NewMetadata = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateItem) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateItem) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateItem) SetNewName(v string) {
	o.NewName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateItem) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateItem) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateItem) SetPassword(v string) {
	o.Password = &v
}

// GetRmTag returns the RmTag field value if set, zero value otherwise.
func (o *UpdateItem) GetRmTag() []string {
	if o == nil || o.RmTag == nil {
		var ret []string
		return ret
	}
	return *o.RmTag
}

// GetRmTagOk returns a tuple with the RmTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetRmTagOk() (*[]string, bool) {
	if o == nil || o.RmTag == nil {
		return nil, false
	}
	return o.RmTag, true
}

// HasRmTag returns a boolean if a field has been set.
func (o *UpdateItem) HasRmTag() bool {
	if o != nil && o.RmTag != nil {
		return true
	}

	return false
}

// SetRmTag gets a reference to the given []string and assigns it to the RmTag field.
func (o *UpdateItem) SetRmTag(v []string) {
	o.RmTag = &v
}

// GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessAllowPortForwading() bool {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAllowPortForwading
}

// GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessAllowPortForwadingOk() (*bool, bool) {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		return nil, false
	}
	return o.SecureAccessAllowPortForwading, true
}

// HasSecureAccessAllowPortForwading returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessAllowPortForwading() bool {
	if o != nil && o.SecureAccessAllowPortForwading != nil {
		return true
	}

	return false
}

// SetSecureAccessAllowPortForwading gets a reference to the given bool and assigns it to the SecureAccessAllowPortForwading field.
func (o *UpdateItem) SetSecureAccessAllowPortForwading(v bool) {
	o.SecureAccessAllowPortForwading = &v
}

// GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessAwsAccountId() string {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessAwsAccountId
}

// GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessAwsAccountIdOk() (*string, bool) {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		return nil, false
	}
	return o.SecureAccessAwsAccountId, true
}

// HasSecureAccessAwsAccountId returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessAwsAccountId() bool {
	if o != nil && o.SecureAccessAwsAccountId != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsAccountId gets a reference to the given string and assigns it to the SecureAccessAwsAccountId field.
func (o *UpdateItem) SetSecureAccessAwsAccountId(v string) {
	o.SecureAccessAwsAccountId = &v
}

// GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessAwsNativeCli() bool {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAwsNativeCli
}

// GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessAwsNativeCliOk() (*bool, bool) {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		return nil, false
	}
	return o.SecureAccessAwsNativeCli, true
}

// HasSecureAccessAwsNativeCli returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessAwsNativeCli() bool {
	if o != nil && o.SecureAccessAwsNativeCli != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsNativeCli gets a reference to the given bool and assigns it to the SecureAccessAwsNativeCli field.
func (o *UpdateItem) SetSecureAccessAwsNativeCli(v bool) {
	o.SecureAccessAwsNativeCli = &v
}

// GetSecureAccessAwsRegion returns the SecureAccessAwsRegion field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessAwsRegion() string {
	if o == nil || o.SecureAccessAwsRegion == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessAwsRegion
}

// GetSecureAccessAwsRegionOk returns a tuple with the SecureAccessAwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessAwsRegionOk() (*string, bool) {
	if o == nil || o.SecureAccessAwsRegion == nil {
		return nil, false
	}
	return o.SecureAccessAwsRegion, true
}

// HasSecureAccessAwsRegion returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessAwsRegion() bool {
	if o != nil && o.SecureAccessAwsRegion != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsRegion gets a reference to the given string and assigns it to the SecureAccessAwsRegion field.
func (o *UpdateItem) SetSecureAccessAwsRegion(v string) {
	o.SecureAccessAwsRegion = &v
}

// GetSecureAccessBastionApi returns the SecureAccessBastionApi field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessBastionApi() string {
	if o == nil || o.SecureAccessBastionApi == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionApi
}

// GetSecureAccessBastionApiOk returns a tuple with the SecureAccessBastionApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessBastionApiOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionApi == nil {
		return nil, false
	}
	return o.SecureAccessBastionApi, true
}

// HasSecureAccessBastionApi returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessBastionApi() bool {
	if o != nil && o.SecureAccessBastionApi != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionApi gets a reference to the given string and assigns it to the SecureAccessBastionApi field.
func (o *UpdateItem) SetSecureAccessBastionApi(v string) {
	o.SecureAccessBastionApi = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *UpdateItem) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessBastionSsh returns the SecureAccessBastionSsh field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessBastionSsh() string {
	if o == nil || o.SecureAccessBastionSsh == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionSsh
}

// GetSecureAccessBastionSshOk returns a tuple with the SecureAccessBastionSsh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessBastionSshOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionSsh == nil {
		return nil, false
	}
	return o.SecureAccessBastionSsh, true
}

// HasSecureAccessBastionSsh returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessBastionSsh() bool {
	if o != nil && o.SecureAccessBastionSsh != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionSsh gets a reference to the given string and assigns it to the SecureAccessBastionSsh field.
func (o *UpdateItem) SetSecureAccessBastionSsh(v string) {
	o.SecureAccessBastionSsh = &v
}

// GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessClusterEndpoint() string {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessClusterEndpoint
}

// GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessClusterEndpointOk() (*string, bool) {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		return nil, false
	}
	return o.SecureAccessClusterEndpoint, true
}

// HasSecureAccessClusterEndpoint returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessClusterEndpoint() bool {
	if o != nil && o.SecureAccessClusterEndpoint != nil {
		return true
	}

	return false
}

// SetSecureAccessClusterEndpoint gets a reference to the given string and assigns it to the SecureAccessClusterEndpoint field.
func (o *UpdateItem) SetSecureAccessClusterEndpoint(v string) {
	o.SecureAccessClusterEndpoint = &v
}

// GetSecureAccessDashboardUrl returns the SecureAccessDashboardUrl field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessDashboardUrl() string {
	if o == nil || o.SecureAccessDashboardUrl == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDashboardUrl
}

// GetSecureAccessDashboardUrlOk returns a tuple with the SecureAccessDashboardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessDashboardUrlOk() (*string, bool) {
	if o == nil || o.SecureAccessDashboardUrl == nil {
		return nil, false
	}
	return o.SecureAccessDashboardUrl, true
}

// HasSecureAccessDashboardUrl returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessDashboardUrl() bool {
	if o != nil && o.SecureAccessDashboardUrl != nil {
		return true
	}

	return false
}

// SetSecureAccessDashboardUrl gets a reference to the given string and assigns it to the SecureAccessDashboardUrl field.
func (o *UpdateItem) SetSecureAccessDashboardUrl(v string) {
	o.SecureAccessDashboardUrl = &v
}

// GetSecureAccessDbName returns the SecureAccessDbName field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessDbName() string {
	if o == nil || o.SecureAccessDbName == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDbName
}

// GetSecureAccessDbNameOk returns a tuple with the SecureAccessDbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessDbNameOk() (*string, bool) {
	if o == nil || o.SecureAccessDbName == nil {
		return nil, false
	}
	return o.SecureAccessDbName, true
}

// HasSecureAccessDbName returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessDbName() bool {
	if o != nil && o.SecureAccessDbName != nil {
		return true
	}

	return false
}

// SetSecureAccessDbName gets a reference to the given string and assigns it to the SecureAccessDbName field.
func (o *UpdateItem) SetSecureAccessDbName(v string) {
	o.SecureAccessDbName = &v
}

// GetSecureAccessDbSchema returns the SecureAccessDbSchema field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessDbSchema() string {
	if o == nil || o.SecureAccessDbSchema == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDbSchema
}

// GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessDbSchemaOk() (*string, bool) {
	if o == nil || o.SecureAccessDbSchema == nil {
		return nil, false
	}
	return o.SecureAccessDbSchema, true
}

// HasSecureAccessDbSchema returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessDbSchema() bool {
	if o != nil && o.SecureAccessDbSchema != nil {
		return true
	}

	return false
}

// SetSecureAccessDbSchema gets a reference to the given string and assigns it to the SecureAccessDbSchema field.
func (o *UpdateItem) SetSecureAccessDbSchema(v string) {
	o.SecureAccessDbSchema = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *UpdateItem) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *UpdateItem) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessRdpDomain() string {
	if o == nil || o.SecureAccessRdpDomain == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessRdpDomain
}

// GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessRdpDomainOk() (*string, bool) {
	if o == nil || o.SecureAccessRdpDomain == nil {
		return nil, false
	}
	return o.SecureAccessRdpDomain, true
}

// HasSecureAccessRdpDomain returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessRdpDomain() bool {
	if o != nil && o.SecureAccessRdpDomain != nil {
		return true
	}

	return false
}

// SetSecureAccessRdpDomain gets a reference to the given string and assigns it to the SecureAccessRdpDomain field.
func (o *UpdateItem) SetSecureAccessRdpDomain(v string) {
	o.SecureAccessRdpDomain = &v
}

// GetSecureAccessSshCreds returns the SecureAccessSshCreds field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessSshCreds() string {
	if o == nil || o.SecureAccessSshCreds == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessSshCreds
}

// GetSecureAccessSshCredsOk returns a tuple with the SecureAccessSshCreds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessSshCredsOk() (*string, bool) {
	if o == nil || o.SecureAccessSshCreds == nil {
		return nil, false
	}
	return o.SecureAccessSshCreds, true
}

// HasSecureAccessSshCreds returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessSshCreds() bool {
	if o != nil && o.SecureAccessSshCreds != nil {
		return true
	}

	return false
}

// SetSecureAccessSshCreds gets a reference to the given string and assigns it to the SecureAccessSshCreds field.
func (o *UpdateItem) SetSecureAccessSshCreds(v string) {
	o.SecureAccessSshCreds = &v
}

// GetSecureAccessSshCredsUser returns the SecureAccessSshCredsUser field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessSshCredsUser() string {
	if o == nil || o.SecureAccessSshCredsUser == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessSshCredsUser
}

// GetSecureAccessSshCredsUserOk returns a tuple with the SecureAccessSshCredsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessSshCredsUserOk() (*string, bool) {
	if o == nil || o.SecureAccessSshCredsUser == nil {
		return nil, false
	}
	return o.SecureAccessSshCredsUser, true
}

// HasSecureAccessSshCredsUser returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessSshCredsUser() bool {
	if o != nil && o.SecureAccessSshCredsUser != nil {
		return true
	}

	return false
}

// SetSecureAccessSshCredsUser gets a reference to the given string and assigns it to the SecureAccessSshCredsUser field.
func (o *UpdateItem) SetSecureAccessSshCredsUser(v string) {
	o.SecureAccessSshCredsUser = &v
}

// GetSecureAccessUrl returns the SecureAccessUrl field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessUrl() string {
	if o == nil || o.SecureAccessUrl == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessUrl
}

// GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessUrlOk() (*string, bool) {
	if o == nil || o.SecureAccessUrl == nil {
		return nil, false
	}
	return o.SecureAccessUrl, true
}

// HasSecureAccessUrl returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessUrl() bool {
	if o != nil && o.SecureAccessUrl != nil {
		return true
	}

	return false
}

// SetSecureAccessUrl gets a reference to the given string and assigns it to the SecureAccessUrl field.
func (o *UpdateItem) SetSecureAccessUrl(v string) {
	o.SecureAccessUrl = &v
}

// GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field value if set, zero value otherwise.
func (o *UpdateItem) GetSecureAccessWebBrowsing() bool {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebBrowsing
}

// GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetSecureAccessWebBrowsingOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		return nil, false
	}
	return o.SecureAccessWebBrowsing, true
}

// HasSecureAccessWebBrowsing returns a boolean if a field has been set.
func (o *UpdateItem) HasSecureAccessWebBrowsing() bool {
	if o != nil && o.SecureAccessWebBrowsing != nil {
		return true
	}

	return false
}

// SetSecureAccessWebBrowsing gets a reference to the given bool and assigns it to the SecureAccessWebBrowsing field.
func (o *UpdateItem) SetSecureAccessWebBrowsing(v bool) {
	o.SecureAccessWebBrowsing = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateItem) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateItem) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateItem) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateItem) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateItem) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateItem) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateItem) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItem) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateItem) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateItem) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddTag != nil {
		toSerialize["add-tag"] = o.AddTag
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewMetadata != nil {
		toSerialize["new-metadata"] = o.NewMetadata
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.RmTag != nil {
		toSerialize["rm-tag"] = o.RmTag
	}
	if o.SecureAccessAllowPortForwading != nil {
		toSerialize["secure-access-allow-port-forwading"] = o.SecureAccessAllowPortForwading
	}
	if o.SecureAccessAwsAccountId != nil {
		toSerialize["secure-access-aws-account-id"] = o.SecureAccessAwsAccountId
	}
	if o.SecureAccessAwsNativeCli != nil {
		toSerialize["secure-access-aws-native-cli"] = o.SecureAccessAwsNativeCli
	}
	if o.SecureAccessAwsRegion != nil {
		toSerialize["secure-access-aws-region"] = o.SecureAccessAwsRegion
	}
	if o.SecureAccessBastionApi != nil {
		toSerialize["secure-access-bastion-api"] = o.SecureAccessBastionApi
	}
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessBastionSsh != nil {
		toSerialize["secure-access-bastion-ssh"] = o.SecureAccessBastionSsh
	}
	if o.SecureAccessClusterEndpoint != nil {
		toSerialize["secure-access-cluster-endpoint"] = o.SecureAccessClusterEndpoint
	}
	if o.SecureAccessDashboardUrl != nil {
		toSerialize["secure-access-dashboard-url"] = o.SecureAccessDashboardUrl
	}
	if o.SecureAccessDbName != nil {
		toSerialize["secure-access-db-name"] = o.SecureAccessDbName
	}
	if o.SecureAccessDbSchema != nil {
		toSerialize["secure-access-db-schema"] = o.SecureAccessDbSchema
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessHost != nil {
		toSerialize["secure-access-host"] = o.SecureAccessHost
	}
	if o.SecureAccessRdpDomain != nil {
		toSerialize["secure-access-rdp-domain"] = o.SecureAccessRdpDomain
	}
	if o.SecureAccessSshCreds != nil {
		toSerialize["secure-access-ssh-creds"] = o.SecureAccessSshCreds
	}
	if o.SecureAccessSshCredsUser != nil {
		toSerialize["secure-access-ssh-creds-user"] = o.SecureAccessSshCredsUser
	}
	if o.SecureAccessUrl != nil {
		toSerialize["secure-access-url"] = o.SecureAccessUrl
	}
	if o.SecureAccessWebBrowsing != nil {
		toSerialize["secure-access-web-browsing"] = o.SecureAccessWebBrowsing
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateItem struct {
	value *UpdateItem
	isSet bool
}

func (v NullableUpdateItem) Get() *UpdateItem {
	return v.value
}

func (v *NullableUpdateItem) Set(val *UpdateItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateItem(val *UpdateItem) *NullableUpdateItem {
	return &NullableUpdateItem{value: val, isSet: true}
}

func (v NullableUpdateItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


