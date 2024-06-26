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

// EventForwarderCreateWebhook eventForwarderCreateWebhook is a command that creates webhook event forwarder
type EventForwarderCreateWebhook struct {
	// Auth Method Event sources
	AuthMethodsEventSourceLocations *[]string `json:"auth-methods-event-source-locations,omitempty"`
	// Base64 encoded Token string for authentication type Token
	AuthToken *string `json:"auth-token,omitempty"`
	// The Webhook authentication type [user-pass, token, certificate]
	AuthType *string `json:"auth-type,omitempty"`
	// Base64 encoded PEM certificate, relevant for certificate auth-type
	ClientCertData *string `json:"client-cert-data,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated]
	EventTypes *[]string `json:"event-types,omitempty"`
	// Rate of periodic runner repetition in hours
	Every *string `json:"every,omitempty"`
	// Event sources
	GatewaysEventSourceLocations []string `json:"gateways-event-source-locations"`
	// Items Event sources
	ItemsEventSourceLocations *[]string `json:"items-event-source-locations,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// EventForwarder name
	Name string `json:"name"`
	// Password for authentication relevant for user-pass auth-type
	Password *string `json:"password,omitempty"`
	// Base64 encoded PEM RSA Private Key, relevant for certificate auth-type
	PrivateKeyData *string `json:"private-key-data,omitempty"`
	RunnerType string `json:"runner-type"`
	// Base64 encoded PEM certificate of the Webhook
	ServerCertificates *string `json:"server-certificates,omitempty"`
	// Targets Event sources
	TargetsEventSourceLocations *[]string `json:"targets-event-source-locations,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Webhook URL
	Url *string `json:"url,omitempty"`
	// Username for authentication relevant for user-pass auth-type
	Username *string `json:"username,omitempty"`
}

// NewEventForwarderCreateWebhook instantiates a new EventForwarderCreateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventForwarderCreateWebhook(gatewaysEventSourceLocations []string, name string, runnerType string, ) *EventForwarderCreateWebhook {
	this := EventForwarderCreateWebhook{}
	var authType string = "user-pass"
	this.AuthType = &authType
	this.GatewaysEventSourceLocations = gatewaysEventSourceLocations
	var json bool = false
	this.Json = &json
	this.Name = name
	this.RunnerType = runnerType
	return &this
}

// NewEventForwarderCreateWebhookWithDefaults instantiates a new EventForwarderCreateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventForwarderCreateWebhookWithDefaults() *EventForwarderCreateWebhook {
	this := EventForwarderCreateWebhook{}
	var authType string = "user-pass"
	this.AuthType = &authType
	var json bool = false
	this.Json = &json
	return &this
}

// GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetAuthMethodsEventSourceLocations() []string {
	if o == nil || o.AuthMethodsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.AuthMethodsEventSourceLocations
}

// GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.AuthMethodsEventSourceLocations == nil {
		return nil, false
	}
	return o.AuthMethodsEventSourceLocations, true
}

// HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasAuthMethodsEventSourceLocations() bool {
	if o != nil && o.AuthMethodsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetAuthMethodsEventSourceLocations gets a reference to the given []string and assigns it to the AuthMethodsEventSourceLocations field.
func (o *EventForwarderCreateWebhook) SetAuthMethodsEventSourceLocations(v []string) {
	o.AuthMethodsEventSourceLocations = &v
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetAuthToken() string {
	if o == nil || o.AuthToken == nil {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetAuthTokenOk() (*string, bool) {
	if o == nil || o.AuthToken == nil {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasAuthToken() bool {
	if o != nil && o.AuthToken != nil {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *EventForwarderCreateWebhook) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *EventForwarderCreateWebhook) SetAuthType(v string) {
	o.AuthType = &v
}

// GetClientCertData returns the ClientCertData field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetClientCertData() string {
	if o == nil || o.ClientCertData == nil {
		var ret string
		return ret
	}
	return *o.ClientCertData
}

// GetClientCertDataOk returns a tuple with the ClientCertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetClientCertDataOk() (*string, bool) {
	if o == nil || o.ClientCertData == nil {
		return nil, false
	}
	return o.ClientCertData, true
}

// HasClientCertData returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasClientCertData() bool {
	if o != nil && o.ClientCertData != nil {
		return true
	}

	return false
}

// SetClientCertData gets a reference to the given string and assigns it to the ClientCertData field.
func (o *EventForwarderCreateWebhook) SetClientCertData(v string) {
	o.ClientCertData = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventForwarderCreateWebhook) SetDescription(v string) {
	o.Description = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetEventTypesOk() (*[]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *EventForwarderCreateWebhook) SetEventTypes(v []string) {
	o.EventTypes = &v
}

// GetEvery returns the Every field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetEvery() string {
	if o == nil || o.Every == nil {
		var ret string
		return ret
	}
	return *o.Every
}

// GetEveryOk returns a tuple with the Every field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetEveryOk() (*string, bool) {
	if o == nil || o.Every == nil {
		return nil, false
	}
	return o.Every, true
}

// HasEvery returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasEvery() bool {
	if o != nil && o.Every != nil {
		return true
	}

	return false
}

// SetEvery gets a reference to the given string and assigns it to the Every field.
func (o *EventForwarderCreateWebhook) SetEvery(v string) {
	o.Every = &v
}

// GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field value
func (o *EventForwarderCreateWebhook) GetGatewaysEventSourceLocations() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.GatewaysEventSourceLocations
}

// GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetGatewaysEventSourceLocationsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GatewaysEventSourceLocations, true
}

// SetGatewaysEventSourceLocations sets field value
func (o *EventForwarderCreateWebhook) SetGatewaysEventSourceLocations(v []string) {
	o.GatewaysEventSourceLocations = v
}

// GetItemsEventSourceLocations returns the ItemsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetItemsEventSourceLocations() []string {
	if o == nil || o.ItemsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.ItemsEventSourceLocations
}

// GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetItemsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.ItemsEventSourceLocations == nil {
		return nil, false
	}
	return o.ItemsEventSourceLocations, true
}

// HasItemsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasItemsEventSourceLocations() bool {
	if o != nil && o.ItemsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetItemsEventSourceLocations gets a reference to the given []string and assigns it to the ItemsEventSourceLocations field.
func (o *EventForwarderCreateWebhook) SetItemsEventSourceLocations(v []string) {
	o.ItemsEventSourceLocations = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *EventForwarderCreateWebhook) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EventForwarderCreateWebhook) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *EventForwarderCreateWebhook) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventForwarderCreateWebhook) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *EventForwarderCreateWebhook) SetPassword(v string) {
	o.Password = &v
}

// GetPrivateKeyData returns the PrivateKeyData field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetPrivateKeyData() string {
	if o == nil || o.PrivateKeyData == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyData
}

// GetPrivateKeyDataOk returns a tuple with the PrivateKeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetPrivateKeyDataOk() (*string, bool) {
	if o == nil || o.PrivateKeyData == nil {
		return nil, false
	}
	return o.PrivateKeyData, true
}

// HasPrivateKeyData returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasPrivateKeyData() bool {
	if o != nil && o.PrivateKeyData != nil {
		return true
	}

	return false
}

// SetPrivateKeyData gets a reference to the given string and assigns it to the PrivateKeyData field.
func (o *EventForwarderCreateWebhook) SetPrivateKeyData(v string) {
	o.PrivateKeyData = &v
}

// GetRunnerType returns the RunnerType field value
func (o *EventForwarderCreateWebhook) GetRunnerType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RunnerType
}

// GetRunnerTypeOk returns a tuple with the RunnerType field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetRunnerTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RunnerType, true
}

// SetRunnerType sets field value
func (o *EventForwarderCreateWebhook) SetRunnerType(v string) {
	o.RunnerType = v
}

// GetServerCertificates returns the ServerCertificates field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetServerCertificates() string {
	if o == nil || o.ServerCertificates == nil {
		var ret string
		return ret
	}
	return *o.ServerCertificates
}

// GetServerCertificatesOk returns a tuple with the ServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetServerCertificatesOk() (*string, bool) {
	if o == nil || o.ServerCertificates == nil {
		return nil, false
	}
	return o.ServerCertificates, true
}

// HasServerCertificates returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasServerCertificates() bool {
	if o != nil && o.ServerCertificates != nil {
		return true
	}

	return false
}

// SetServerCertificates gets a reference to the given string and assigns it to the ServerCertificates field.
func (o *EventForwarderCreateWebhook) SetServerCertificates(v string) {
	o.ServerCertificates = &v
}

// GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetTargetsEventSourceLocations() []string {
	if o == nil || o.TargetsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.TargetsEventSourceLocations
}

// GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetTargetsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.TargetsEventSourceLocations == nil {
		return nil, false
	}
	return o.TargetsEventSourceLocations, true
}

// HasTargetsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasTargetsEventSourceLocations() bool {
	if o != nil && o.TargetsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetTargetsEventSourceLocations gets a reference to the given []string and assigns it to the TargetsEventSourceLocations field.
func (o *EventForwarderCreateWebhook) SetTargetsEventSourceLocations(v []string) {
	o.TargetsEventSourceLocations = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EventForwarderCreateWebhook) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EventForwarderCreateWebhook) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EventForwarderCreateWebhook) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EventForwarderCreateWebhook) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateWebhook) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EventForwarderCreateWebhook) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EventForwarderCreateWebhook) SetUsername(v string) {
	o.Username = &v
}

func (o EventForwarderCreateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthMethodsEventSourceLocations != nil {
		toSerialize["auth-methods-event-source-locations"] = o.AuthMethodsEventSourceLocations
	}
	if o.AuthToken != nil {
		toSerialize["auth-token"] = o.AuthToken
	}
	if o.AuthType != nil {
		toSerialize["auth-type"] = o.AuthType
	}
	if o.ClientCertData != nil {
		toSerialize["client-cert-data"] = o.ClientCertData
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EventTypes != nil {
		toSerialize["event-types"] = o.EventTypes
	}
	if o.Every != nil {
		toSerialize["every"] = o.Every
	}
	if true {
		toSerialize["gateways-event-source-locations"] = o.GatewaysEventSourceLocations
	}
	if o.ItemsEventSourceLocations != nil {
		toSerialize["items-event-source-locations"] = o.ItemsEventSourceLocations
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PrivateKeyData != nil {
		toSerialize["private-key-data"] = o.PrivateKeyData
	}
	if true {
		toSerialize["runner-type"] = o.RunnerType
	}
	if o.ServerCertificates != nil {
		toSerialize["server-certificates"] = o.ServerCertificates
	}
	if o.TargetsEventSourceLocations != nil {
		toSerialize["targets-event-source-locations"] = o.TargetsEventSourceLocations
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableEventForwarderCreateWebhook struct {
	value *EventForwarderCreateWebhook
	isSet bool
}

func (v NullableEventForwarderCreateWebhook) Get() *EventForwarderCreateWebhook {
	return v.value
}

func (v *NullableEventForwarderCreateWebhook) Set(val *EventForwarderCreateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableEventForwarderCreateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableEventForwarderCreateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventForwarderCreateWebhook(val *EventForwarderCreateWebhook) *NullableEventForwarderCreateWebhook {
	return &NullableEventForwarderCreateWebhook{value: val, isSet: true}
}

func (v NullableEventForwarderCreateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventForwarderCreateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


