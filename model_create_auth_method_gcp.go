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

// CreateAuthMethodGCP createAuthMethodGCP is a command that creates a new auth method that will be able to authenticate using GCP IAM Service Account credentials or GCE instance credentials.
type CreateAuthMethodGCP struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// The audience to verify in the JWT received by the client
	Audience *string `json:"audience,omitempty"`
	// A CIDR whitelist of the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// A comma-separated list of GCP labels formatted as \"key:value\" strings that must be set on authorized GCE instances. TODO: Because GCP labels are not currently ACL'd ....
	BoundLabels *[]string `json:"bound-labels,omitempty"`
	// === Human and Machine authentication section === Array of GCP project IDs. Only entities belonging to any of the provided projects can authenticate.
	BoundProjects *[]string `json:"bound-projects,omitempty"`
	// List of regions that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.
	BoundRegions *[]string `json:"bound-regions,omitempty"`
	// === Human authentication section === List of service accounts the service account must be part of in order to be authenticated.
	BoundServiceAccounts *[]string `json:"bound-service-accounts,omitempty"`
	// === Machine authentication section === List of zones that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.
	BoundZones *[]string `json:"bound-zones,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// ServiceAccount credentials data instead of giving a file path, base64 encoded
	ServiceAccountCredsData *string `json:"service-account-creds-data,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Type of the GCP Access Rules
	Type *string `json:"type,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateAuthMethodGCP instantiates a new CreateAuthMethodGCP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodGCP(name string, ) *CreateAuthMethodGCP {
	this := CreateAuthMethodGCP{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var audience string = "akeyless.io"
	this.Audience = &audience
	this.Name = name
	return &this
}

// NewCreateAuthMethodGCPWithDefaults instantiates a new CreateAuthMethodGCP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodGCPWithDefaults() *CreateAuthMethodGCP {
	this := CreateAuthMethodGCP{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var audience string = "akeyless.io"
	this.Audience = &audience
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *CreateAuthMethodGCP) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *CreateAuthMethodGCP) SetAudience(v string) {
	o.Audience = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *CreateAuthMethodGCP) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetBoundLabels returns the BoundLabels field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetBoundLabels() []string {
	if o == nil || o.BoundLabels == nil {
		var ret []string
		return ret
	}
	return *o.BoundLabels
}

// GetBoundLabelsOk returns a tuple with the BoundLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetBoundLabelsOk() (*[]string, bool) {
	if o == nil || o.BoundLabels == nil {
		return nil, false
	}
	return o.BoundLabels, true
}

// HasBoundLabels returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasBoundLabels() bool {
	if o != nil && o.BoundLabels != nil {
		return true
	}

	return false
}

// SetBoundLabels gets a reference to the given []string and assigns it to the BoundLabels field.
func (o *CreateAuthMethodGCP) SetBoundLabels(v []string) {
	o.BoundLabels = &v
}

// GetBoundProjects returns the BoundProjects field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetBoundProjects() []string {
	if o == nil || o.BoundProjects == nil {
		var ret []string
		return ret
	}
	return *o.BoundProjects
}

// GetBoundProjectsOk returns a tuple with the BoundProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetBoundProjectsOk() (*[]string, bool) {
	if o == nil || o.BoundProjects == nil {
		return nil, false
	}
	return o.BoundProjects, true
}

// HasBoundProjects returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasBoundProjects() bool {
	if o != nil && o.BoundProjects != nil {
		return true
	}

	return false
}

// SetBoundProjects gets a reference to the given []string and assigns it to the BoundProjects field.
func (o *CreateAuthMethodGCP) SetBoundProjects(v []string) {
	o.BoundProjects = &v
}

// GetBoundRegions returns the BoundRegions field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetBoundRegions() []string {
	if o == nil || o.BoundRegions == nil {
		var ret []string
		return ret
	}
	return *o.BoundRegions
}

// GetBoundRegionsOk returns a tuple with the BoundRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetBoundRegionsOk() (*[]string, bool) {
	if o == nil || o.BoundRegions == nil {
		return nil, false
	}
	return o.BoundRegions, true
}

// HasBoundRegions returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasBoundRegions() bool {
	if o != nil && o.BoundRegions != nil {
		return true
	}

	return false
}

// SetBoundRegions gets a reference to the given []string and assigns it to the BoundRegions field.
func (o *CreateAuthMethodGCP) SetBoundRegions(v []string) {
	o.BoundRegions = &v
}

// GetBoundServiceAccounts returns the BoundServiceAccounts field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetBoundServiceAccounts() []string {
	if o == nil || o.BoundServiceAccounts == nil {
		var ret []string
		return ret
	}
	return *o.BoundServiceAccounts
}

// GetBoundServiceAccountsOk returns a tuple with the BoundServiceAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetBoundServiceAccountsOk() (*[]string, bool) {
	if o == nil || o.BoundServiceAccounts == nil {
		return nil, false
	}
	return o.BoundServiceAccounts, true
}

// HasBoundServiceAccounts returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasBoundServiceAccounts() bool {
	if o != nil && o.BoundServiceAccounts != nil {
		return true
	}

	return false
}

// SetBoundServiceAccounts gets a reference to the given []string and assigns it to the BoundServiceAccounts field.
func (o *CreateAuthMethodGCP) SetBoundServiceAccounts(v []string) {
	o.BoundServiceAccounts = &v
}

// GetBoundZones returns the BoundZones field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetBoundZones() []string {
	if o == nil || o.BoundZones == nil {
		var ret []string
		return ret
	}
	return *o.BoundZones
}

// GetBoundZonesOk returns a tuple with the BoundZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetBoundZonesOk() (*[]string, bool) {
	if o == nil || o.BoundZones == nil {
		return nil, false
	}
	return o.BoundZones, true
}

// HasBoundZones returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasBoundZones() bool {
	if o != nil && o.BoundZones != nil {
		return true
	}

	return false
}

// SetBoundZones gets a reference to the given []string and assigns it to the BoundZones field.
func (o *CreateAuthMethodGCP) SetBoundZones(v []string) {
	o.BoundZones = &v
}

// GetName returns the Name field value
func (o *CreateAuthMethodGCP) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAuthMethodGCP) SetName(v string) {
	o.Name = v
}

// GetServiceAccountCredsData returns the ServiceAccountCredsData field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetServiceAccountCredsData() string {
	if o == nil || o.ServiceAccountCredsData == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccountCredsData
}

// GetServiceAccountCredsDataOk returns a tuple with the ServiceAccountCredsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetServiceAccountCredsDataOk() (*string, bool) {
	if o == nil || o.ServiceAccountCredsData == nil {
		return nil, false
	}
	return o.ServiceAccountCredsData, true
}

// HasServiceAccountCredsData returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasServiceAccountCredsData() bool {
	if o != nil && o.ServiceAccountCredsData != nil {
		return true
	}

	return false
}

// SetServiceAccountCredsData gets a reference to the given string and assigns it to the ServiceAccountCredsData field.
func (o *CreateAuthMethodGCP) SetServiceAccountCredsData(v string) {
	o.ServiceAccountCredsData = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateAuthMethodGCP) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateAuthMethodGCP) SetType(v string) {
	o.Type = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateAuthMethodGCP) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodGCP) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateAuthMethodGCP) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateAuthMethodGCP) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateAuthMethodGCP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if o.BoundLabels != nil {
		toSerialize["bound-labels"] = o.BoundLabels
	}
	if o.BoundProjects != nil {
		toSerialize["bound-projects"] = o.BoundProjects
	}
	if o.BoundRegions != nil {
		toSerialize["bound-regions"] = o.BoundRegions
	}
	if o.BoundServiceAccounts != nil {
		toSerialize["bound-service-accounts"] = o.BoundServiceAccounts
	}
	if o.BoundZones != nil {
		toSerialize["bound-zones"] = o.BoundZones
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ServiceAccountCredsData != nil {
		toSerialize["service-account-creds-data"] = o.ServiceAccountCredsData
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAuthMethodGCP struct {
	value *CreateAuthMethodGCP
	isSet bool
}

func (v NullableCreateAuthMethodGCP) Get() *CreateAuthMethodGCP {
	return v.value
}

func (v *NullableCreateAuthMethodGCP) Set(val *CreateAuthMethodGCP) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodGCP) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodGCP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodGCP(val *CreateAuthMethodGCP) *NullableCreateAuthMethodGCP {
	return &NullableCreateAuthMethodGCP{value: val, isSet: true}
}

func (v NullableCreateAuthMethodGCP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodGCP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


