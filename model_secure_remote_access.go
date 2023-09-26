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

// SecureRemoteAccess struct for SecureRemoteAccess
type SecureRemoteAccess struct {
	AccountId *string `json:"account_id,omitempty"`
	AllowPortForwarding *bool `json:"allow_port_forwarding,omitempty"`
	AllowProvidingExternalUsername *bool `json:"allow_providing_external_username,omitempty"`
	BastionApi *string `json:"bastion_api,omitempty"`
	BastionIssuer *string `json:"bastion_issuer,omitempty"`
	BastionIssuerId *int64 `json:"bastion_issuer_id,omitempty"`
	BastionSsh *string `json:"bastion_ssh,omitempty"`
	Category *string `json:"category,omitempty"`
	DashboardUrl *string `json:"dashboard_url,omitempty"`
	DbName *string `json:"db_name,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Enable *bool `json:"enable,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Host *[]string `json:"host,omitempty"`
	IsCli *bool `json:"is_cli,omitempty"`
	IsWeb *bool `json:"is_web,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Native *bool `json:"native,omitempty"`
	RdGatewayServer *string `json:"rd_gateway_server,omitempty"`
	RdpUser *string `json:"rdp_user,omitempty"`
	Region *string `json:"region,omitempty"`
	RotateAfterDisconnect *bool `json:"rotate_after_disconnect,omitempty"`
	Schema *string `json:"schema,omitempty"`
	SshPassword *bool `json:"ssh_password,omitempty"`
	SshPrivateKey *bool `json:"ssh_private_key,omitempty"`
	SshUser *string `json:"ssh_user,omitempty"`
	Url *string `json:"url,omitempty"`
	UseInternalBastion *bool `json:"use_internal_bastion,omitempty"`
	WebProxy *bool `json:"web_proxy,omitempty"`
}

// NewSecureRemoteAccess instantiates a new SecureRemoteAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecureRemoteAccess() *SecureRemoteAccess {
	this := SecureRemoteAccess{}
	return &this
}

// NewSecureRemoteAccessWithDefaults instantiates a new SecureRemoteAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecureRemoteAccessWithDefaults() *SecureRemoteAccess {
	this := SecureRemoteAccess{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *SecureRemoteAccess) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAllowPortForwarding returns the AllowPortForwarding field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetAllowPortForwarding() bool {
	if o == nil || o.AllowPortForwarding == nil {
		var ret bool
		return ret
	}
	return *o.AllowPortForwarding
}

// GetAllowPortForwardingOk returns a tuple with the AllowPortForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetAllowPortForwardingOk() (*bool, bool) {
	if o == nil || o.AllowPortForwarding == nil {
		return nil, false
	}
	return o.AllowPortForwarding, true
}

// HasAllowPortForwarding returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasAllowPortForwarding() bool {
	if o != nil && o.AllowPortForwarding != nil {
		return true
	}

	return false
}

// SetAllowPortForwarding gets a reference to the given bool and assigns it to the AllowPortForwarding field.
func (o *SecureRemoteAccess) SetAllowPortForwarding(v bool) {
	o.AllowPortForwarding = &v
}

// GetAllowProvidingExternalUsername returns the AllowProvidingExternalUsername field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetAllowProvidingExternalUsername() bool {
	if o == nil || o.AllowProvidingExternalUsername == nil {
		var ret bool
		return ret
	}
	return *o.AllowProvidingExternalUsername
}

// GetAllowProvidingExternalUsernameOk returns a tuple with the AllowProvidingExternalUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetAllowProvidingExternalUsernameOk() (*bool, bool) {
	if o == nil || o.AllowProvidingExternalUsername == nil {
		return nil, false
	}
	return o.AllowProvidingExternalUsername, true
}

// HasAllowProvidingExternalUsername returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasAllowProvidingExternalUsername() bool {
	if o != nil && o.AllowProvidingExternalUsername != nil {
		return true
	}

	return false
}

// SetAllowProvidingExternalUsername gets a reference to the given bool and assigns it to the AllowProvidingExternalUsername field.
func (o *SecureRemoteAccess) SetAllowProvidingExternalUsername(v bool) {
	o.AllowProvidingExternalUsername = &v
}

// GetBastionApi returns the BastionApi field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetBastionApi() string {
	if o == nil || o.BastionApi == nil {
		var ret string
		return ret
	}
	return *o.BastionApi
}

// GetBastionApiOk returns a tuple with the BastionApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetBastionApiOk() (*string, bool) {
	if o == nil || o.BastionApi == nil {
		return nil, false
	}
	return o.BastionApi, true
}

// HasBastionApi returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasBastionApi() bool {
	if o != nil && o.BastionApi != nil {
		return true
	}

	return false
}

// SetBastionApi gets a reference to the given string and assigns it to the BastionApi field.
func (o *SecureRemoteAccess) SetBastionApi(v string) {
	o.BastionApi = &v
}

// GetBastionIssuer returns the BastionIssuer field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetBastionIssuer() string {
	if o == nil || o.BastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.BastionIssuer
}

// GetBastionIssuerOk returns a tuple with the BastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetBastionIssuerOk() (*string, bool) {
	if o == nil || o.BastionIssuer == nil {
		return nil, false
	}
	return o.BastionIssuer, true
}

// HasBastionIssuer returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasBastionIssuer() bool {
	if o != nil && o.BastionIssuer != nil {
		return true
	}

	return false
}

// SetBastionIssuer gets a reference to the given string and assigns it to the BastionIssuer field.
func (o *SecureRemoteAccess) SetBastionIssuer(v string) {
	o.BastionIssuer = &v
}

// GetBastionIssuerId returns the BastionIssuerId field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetBastionIssuerId() int64 {
	if o == nil || o.BastionIssuerId == nil {
		var ret int64
		return ret
	}
	return *o.BastionIssuerId
}

// GetBastionIssuerIdOk returns a tuple with the BastionIssuerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetBastionIssuerIdOk() (*int64, bool) {
	if o == nil || o.BastionIssuerId == nil {
		return nil, false
	}
	return o.BastionIssuerId, true
}

// HasBastionIssuerId returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasBastionIssuerId() bool {
	if o != nil && o.BastionIssuerId != nil {
		return true
	}

	return false
}

// SetBastionIssuerId gets a reference to the given int64 and assigns it to the BastionIssuerId field.
func (o *SecureRemoteAccess) SetBastionIssuerId(v int64) {
	o.BastionIssuerId = &v
}

// GetBastionSsh returns the BastionSsh field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetBastionSsh() string {
	if o == nil || o.BastionSsh == nil {
		var ret string
		return ret
	}
	return *o.BastionSsh
}

// GetBastionSshOk returns a tuple with the BastionSsh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetBastionSshOk() (*string, bool) {
	if o == nil || o.BastionSsh == nil {
		return nil, false
	}
	return o.BastionSsh, true
}

// HasBastionSsh returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasBastionSsh() bool {
	if o != nil && o.BastionSsh != nil {
		return true
	}

	return false
}

// SetBastionSsh gets a reference to the given string and assigns it to the BastionSsh field.
func (o *SecureRemoteAccess) SetBastionSsh(v string) {
	o.BastionSsh = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SecureRemoteAccess) SetCategory(v string) {
	o.Category = &v
}

// GetDashboardUrl returns the DashboardUrl field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetDashboardUrl() string {
	if o == nil || o.DashboardUrl == nil {
		var ret string
		return ret
	}
	return *o.DashboardUrl
}

// GetDashboardUrlOk returns a tuple with the DashboardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetDashboardUrlOk() (*string, bool) {
	if o == nil || o.DashboardUrl == nil {
		return nil, false
	}
	return o.DashboardUrl, true
}

// HasDashboardUrl returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasDashboardUrl() bool {
	if o != nil && o.DashboardUrl != nil {
		return true
	}

	return false
}

// SetDashboardUrl gets a reference to the given string and assigns it to the DashboardUrl field.
func (o *SecureRemoteAccess) SetDashboardUrl(v string) {
	o.DashboardUrl = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasDbName() bool {
	if o != nil && o.DbName != nil {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *SecureRemoteAccess) SetDbName(v string) {
	o.DbName = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *SecureRemoteAccess) SetDomain(v string) {
	o.Domain = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *SecureRemoteAccess) SetEnable(v bool) {
	o.Enable = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *SecureRemoteAccess) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetHost() []string {
	if o == nil || o.Host == nil {
		var ret []string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetHostOk() (*[]string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given []string and assigns it to the Host field.
func (o *SecureRemoteAccess) SetHost(v []string) {
	o.Host = &v
}

// GetIsCli returns the IsCli field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetIsCli() bool {
	if o == nil || o.IsCli == nil {
		var ret bool
		return ret
	}
	return *o.IsCli
}

// GetIsCliOk returns a tuple with the IsCli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetIsCliOk() (*bool, bool) {
	if o == nil || o.IsCli == nil {
		return nil, false
	}
	return o.IsCli, true
}

// HasIsCli returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasIsCli() bool {
	if o != nil && o.IsCli != nil {
		return true
	}

	return false
}

// SetIsCli gets a reference to the given bool and assigns it to the IsCli field.
func (o *SecureRemoteAccess) SetIsCli(v bool) {
	o.IsCli = &v
}

// GetIsWeb returns the IsWeb field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetIsWeb() bool {
	if o == nil || o.IsWeb == nil {
		var ret bool
		return ret
	}
	return *o.IsWeb
}

// GetIsWebOk returns a tuple with the IsWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetIsWebOk() (*bool, bool) {
	if o == nil || o.IsWeb == nil {
		return nil, false
	}
	return o.IsWeb, true
}

// HasIsWeb returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasIsWeb() bool {
	if o != nil && o.IsWeb != nil {
		return true
	}

	return false
}

// SetIsWeb gets a reference to the given bool and assigns it to the IsWeb field.
func (o *SecureRemoteAccess) SetIsWeb(v bool) {
	o.IsWeb = &v
}

// GetIsolated returns the Isolated field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetIsolated() bool {
	if o == nil || o.Isolated == nil {
		var ret bool
		return ret
	}
	return *o.Isolated
}

// GetIsolatedOk returns a tuple with the Isolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetIsolatedOk() (*bool, bool) {
	if o == nil || o.Isolated == nil {
		return nil, false
	}
	return o.Isolated, true
}

// HasIsolated returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasIsolated() bool {
	if o != nil && o.Isolated != nil {
		return true
	}

	return false
}

// SetIsolated gets a reference to the given bool and assigns it to the Isolated field.
func (o *SecureRemoteAccess) SetIsolated(v bool) {
	o.Isolated = &v
}

// GetNative returns the Native field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetNative() bool {
	if o == nil || o.Native == nil {
		var ret bool
		return ret
	}
	return *o.Native
}

// GetNativeOk returns a tuple with the Native field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetNativeOk() (*bool, bool) {
	if o == nil || o.Native == nil {
		return nil, false
	}
	return o.Native, true
}

// HasNative returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasNative() bool {
	if o != nil && o.Native != nil {
		return true
	}

	return false
}

// SetNative gets a reference to the given bool and assigns it to the Native field.
func (o *SecureRemoteAccess) SetNative(v bool) {
	o.Native = &v
}

// GetRdGatewayServer returns the RdGatewayServer field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetRdGatewayServer() string {
	if o == nil || o.RdGatewayServer == nil {
		var ret string
		return ret
	}
	return *o.RdGatewayServer
}

// GetRdGatewayServerOk returns a tuple with the RdGatewayServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetRdGatewayServerOk() (*string, bool) {
	if o == nil || o.RdGatewayServer == nil {
		return nil, false
	}
	return o.RdGatewayServer, true
}

// HasRdGatewayServer returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasRdGatewayServer() bool {
	if o != nil && o.RdGatewayServer != nil {
		return true
	}

	return false
}

// SetRdGatewayServer gets a reference to the given string and assigns it to the RdGatewayServer field.
func (o *SecureRemoteAccess) SetRdGatewayServer(v string) {
	o.RdGatewayServer = &v
}

// GetRdpUser returns the RdpUser field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetRdpUser() string {
	if o == nil || o.RdpUser == nil {
		var ret string
		return ret
	}
	return *o.RdpUser
}

// GetRdpUserOk returns a tuple with the RdpUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetRdpUserOk() (*string, bool) {
	if o == nil || o.RdpUser == nil {
		return nil, false
	}
	return o.RdpUser, true
}

// HasRdpUser returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasRdpUser() bool {
	if o != nil && o.RdpUser != nil {
		return true
	}

	return false
}

// SetRdpUser gets a reference to the given string and assigns it to the RdpUser field.
func (o *SecureRemoteAccess) SetRdpUser(v string) {
	o.RdpUser = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *SecureRemoteAccess) SetRegion(v string) {
	o.Region = &v
}

// GetRotateAfterDisconnect returns the RotateAfterDisconnect field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetRotateAfterDisconnect() bool {
	if o == nil || o.RotateAfterDisconnect == nil {
		var ret bool
		return ret
	}
	return *o.RotateAfterDisconnect
}

// GetRotateAfterDisconnectOk returns a tuple with the RotateAfterDisconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetRotateAfterDisconnectOk() (*bool, bool) {
	if o == nil || o.RotateAfterDisconnect == nil {
		return nil, false
	}
	return o.RotateAfterDisconnect, true
}

// HasRotateAfterDisconnect returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasRotateAfterDisconnect() bool {
	if o != nil && o.RotateAfterDisconnect != nil {
		return true
	}

	return false
}

// SetRotateAfterDisconnect gets a reference to the given bool and assigns it to the RotateAfterDisconnect field.
func (o *SecureRemoteAccess) SetRotateAfterDisconnect(v bool) {
	o.RotateAfterDisconnect = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *SecureRemoteAccess) SetSchema(v string) {
	o.Schema = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetSshPassword() bool {
	if o == nil || o.SshPassword == nil {
		var ret bool
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetSshPasswordOk() (*bool, bool) {
	if o == nil || o.SshPassword == nil {
		return nil, false
	}
	return o.SshPassword, true
}

// HasSshPassword returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasSshPassword() bool {
	if o != nil && o.SshPassword != nil {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given bool and assigns it to the SshPassword field.
func (o *SecureRemoteAccess) SetSshPassword(v bool) {
	o.SshPassword = &v
}

// GetSshPrivateKey returns the SshPrivateKey field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetSshPrivateKey() bool {
	if o == nil || o.SshPrivateKey == nil {
		var ret bool
		return ret
	}
	return *o.SshPrivateKey
}

// GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetSshPrivateKeyOk() (*bool, bool) {
	if o == nil || o.SshPrivateKey == nil {
		return nil, false
	}
	return o.SshPrivateKey, true
}

// HasSshPrivateKey returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasSshPrivateKey() bool {
	if o != nil && o.SshPrivateKey != nil {
		return true
	}

	return false
}

// SetSshPrivateKey gets a reference to the given bool and assigns it to the SshPrivateKey field.
func (o *SecureRemoteAccess) SetSshPrivateKey(v bool) {
	o.SshPrivateKey = &v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetSshUser() string {
	if o == nil || o.SshUser == nil {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetSshUserOk() (*string, bool) {
	if o == nil || o.SshUser == nil {
		return nil, false
	}
	return o.SshUser, true
}

// HasSshUser returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasSshUser() bool {
	if o != nil && o.SshUser != nil {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *SecureRemoteAccess) SetSshUser(v string) {
	o.SshUser = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SecureRemoteAccess) SetUrl(v string) {
	o.Url = &v
}

// GetUseInternalBastion returns the UseInternalBastion field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetUseInternalBastion() bool {
	if o == nil || o.UseInternalBastion == nil {
		var ret bool
		return ret
	}
	return *o.UseInternalBastion
}

// GetUseInternalBastionOk returns a tuple with the UseInternalBastion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetUseInternalBastionOk() (*bool, bool) {
	if o == nil || o.UseInternalBastion == nil {
		return nil, false
	}
	return o.UseInternalBastion, true
}

// HasUseInternalBastion returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasUseInternalBastion() bool {
	if o != nil && o.UseInternalBastion != nil {
		return true
	}

	return false
}

// SetUseInternalBastion gets a reference to the given bool and assigns it to the UseInternalBastion field.
func (o *SecureRemoteAccess) SetUseInternalBastion(v bool) {
	o.UseInternalBastion = &v
}

// GetWebProxy returns the WebProxy field value if set, zero value otherwise.
func (o *SecureRemoteAccess) GetWebProxy() bool {
	if o == nil || o.WebProxy == nil {
		var ret bool
		return ret
	}
	return *o.WebProxy
}

// GetWebProxyOk returns a tuple with the WebProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteAccess) GetWebProxyOk() (*bool, bool) {
	if o == nil || o.WebProxy == nil {
		return nil, false
	}
	return o.WebProxy, true
}

// HasWebProxy returns a boolean if a field has been set.
func (o *SecureRemoteAccess) HasWebProxy() bool {
	if o != nil && o.WebProxy != nil {
		return true
	}

	return false
}

// SetWebProxy gets a reference to the given bool and assigns it to the WebProxy field.
func (o *SecureRemoteAccess) SetWebProxy(v bool) {
	o.WebProxy = &v
}

func (o SecureRemoteAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.AllowPortForwarding != nil {
		toSerialize["allow_port_forwarding"] = o.AllowPortForwarding
	}
	if o.AllowProvidingExternalUsername != nil {
		toSerialize["allow_providing_external_username"] = o.AllowProvidingExternalUsername
	}
	if o.BastionApi != nil {
		toSerialize["bastion_api"] = o.BastionApi
	}
	if o.BastionIssuer != nil {
		toSerialize["bastion_issuer"] = o.BastionIssuer
	}
	if o.BastionIssuerId != nil {
		toSerialize["bastion_issuer_id"] = o.BastionIssuerId
	}
	if o.BastionSsh != nil {
		toSerialize["bastion_ssh"] = o.BastionSsh
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.DashboardUrl != nil {
		toSerialize["dashboard_url"] = o.DashboardUrl
	}
	if o.DbName != nil {
		toSerialize["db_name"] = o.DbName
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.IsCli != nil {
		toSerialize["is_cli"] = o.IsCli
	}
	if o.IsWeb != nil {
		toSerialize["is_web"] = o.IsWeb
	}
	if o.Isolated != nil {
		toSerialize["isolated"] = o.Isolated
	}
	if o.Native != nil {
		toSerialize["native"] = o.Native
	}
	if o.RdGatewayServer != nil {
		toSerialize["rd_gateway_server"] = o.RdGatewayServer
	}
	if o.RdpUser != nil {
		toSerialize["rdp_user"] = o.RdpUser
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.RotateAfterDisconnect != nil {
		toSerialize["rotate_after_disconnect"] = o.RotateAfterDisconnect
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.SshPassword != nil {
		toSerialize["ssh_password"] = o.SshPassword
	}
	if o.SshPrivateKey != nil {
		toSerialize["ssh_private_key"] = o.SshPrivateKey
	}
	if o.SshUser != nil {
		toSerialize["ssh_user"] = o.SshUser
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.UseInternalBastion != nil {
		toSerialize["use_internal_bastion"] = o.UseInternalBastion
	}
	if o.WebProxy != nil {
		toSerialize["web_proxy"] = o.WebProxy
	}
	return json.Marshal(toSerialize)
}

type NullableSecureRemoteAccess struct {
	value *SecureRemoteAccess
	isSet bool
}

func (v NullableSecureRemoteAccess) Get() *SecureRemoteAccess {
	return v.value
}

func (v *NullableSecureRemoteAccess) Set(val *SecureRemoteAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableSecureRemoteAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableSecureRemoteAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecureRemoteAccess(val *SecureRemoteAccess) *NullableSecureRemoteAccess {
	return &NullableSecureRemoteAccess{value: val, isSet: true}
}

func (v NullableSecureRemoteAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecureRemoteAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


