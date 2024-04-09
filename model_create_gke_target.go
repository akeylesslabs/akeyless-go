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

// CreateGKETarget struct for CreateGKETarget
type CreateGKETarget struct {
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// GKE Service Account key file path
	GkeAccountKey *string `json:"gke-account-key,omitempty"`
	// GKE cluster CA certificate
	GkeClusterCert *string `json:"gke-cluster-cert,omitempty"`
	// GKE cluster URL endpoint
	GkeClusterEndpoint *string `json:"gke-cluster-endpoint,omitempty"`
	// GKE cluster name
	GkeClusterName *string `json:"gke-cluster-name,omitempty"`
	// GKE service account email
	GkeServiceAccountEmail *string `json:"gke-service-account-email,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Set the maximum number of versions, limited by the account settings defaults.
	MaxVersions *string `json:"max-versions,omitempty"`
	// Target name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	UseGwCloudIdentity *bool `json:"use-gw-cloud-identity,omitempty"`
}

// NewCreateGKETarget instantiates a new CreateGKETarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGKETarget(name string, ) *CreateGKETarget {
	this := CreateGKETarget{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewCreateGKETargetWithDefaults instantiates a new CreateGKETarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGKETargetWithDefaults() *CreateGKETarget {
	this := CreateGKETarget{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateGKETarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateGKETarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateGKETarget) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateGKETarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGKETarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateGKETarget) SetDescription(v string) {
	o.Description = &v
}

// GetGkeAccountKey returns the GkeAccountKey field value if set, zero value otherwise.
func (o *CreateGKETarget) GetGkeAccountKey() string {
	if o == nil || o.GkeAccountKey == nil {
		var ret string
		return ret
	}
	return *o.GkeAccountKey
}

// GetGkeAccountKeyOk returns a tuple with the GkeAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetGkeAccountKeyOk() (*string, bool) {
	if o == nil || o.GkeAccountKey == nil {
		return nil, false
	}
	return o.GkeAccountKey, true
}

// HasGkeAccountKey returns a boolean if a field has been set.
func (o *CreateGKETarget) HasGkeAccountKey() bool {
	if o != nil && o.GkeAccountKey != nil {
		return true
	}

	return false
}

// SetGkeAccountKey gets a reference to the given string and assigns it to the GkeAccountKey field.
func (o *CreateGKETarget) SetGkeAccountKey(v string) {
	o.GkeAccountKey = &v
}

// GetGkeClusterCert returns the GkeClusterCert field value if set, zero value otherwise.
func (o *CreateGKETarget) GetGkeClusterCert() string {
	if o == nil || o.GkeClusterCert == nil {
		var ret string
		return ret
	}
	return *o.GkeClusterCert
}

// GetGkeClusterCertOk returns a tuple with the GkeClusterCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetGkeClusterCertOk() (*string, bool) {
	if o == nil || o.GkeClusterCert == nil {
		return nil, false
	}
	return o.GkeClusterCert, true
}

// HasGkeClusterCert returns a boolean if a field has been set.
func (o *CreateGKETarget) HasGkeClusterCert() bool {
	if o != nil && o.GkeClusterCert != nil {
		return true
	}

	return false
}

// SetGkeClusterCert gets a reference to the given string and assigns it to the GkeClusterCert field.
func (o *CreateGKETarget) SetGkeClusterCert(v string) {
	o.GkeClusterCert = &v
}

// GetGkeClusterEndpoint returns the GkeClusterEndpoint field value if set, zero value otherwise.
func (o *CreateGKETarget) GetGkeClusterEndpoint() string {
	if o == nil || o.GkeClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.GkeClusterEndpoint
}

// GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetGkeClusterEndpointOk() (*string, bool) {
	if o == nil || o.GkeClusterEndpoint == nil {
		return nil, false
	}
	return o.GkeClusterEndpoint, true
}

// HasGkeClusterEndpoint returns a boolean if a field has been set.
func (o *CreateGKETarget) HasGkeClusterEndpoint() bool {
	if o != nil && o.GkeClusterEndpoint != nil {
		return true
	}

	return false
}

// SetGkeClusterEndpoint gets a reference to the given string and assigns it to the GkeClusterEndpoint field.
func (o *CreateGKETarget) SetGkeClusterEndpoint(v string) {
	o.GkeClusterEndpoint = &v
}

// GetGkeClusterName returns the GkeClusterName field value if set, zero value otherwise.
func (o *CreateGKETarget) GetGkeClusterName() string {
	if o == nil || o.GkeClusterName == nil {
		var ret string
		return ret
	}
	return *o.GkeClusterName
}

// GetGkeClusterNameOk returns a tuple with the GkeClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetGkeClusterNameOk() (*string, bool) {
	if o == nil || o.GkeClusterName == nil {
		return nil, false
	}
	return o.GkeClusterName, true
}

// HasGkeClusterName returns a boolean if a field has been set.
func (o *CreateGKETarget) HasGkeClusterName() bool {
	if o != nil && o.GkeClusterName != nil {
		return true
	}

	return false
}

// SetGkeClusterName gets a reference to the given string and assigns it to the GkeClusterName field.
func (o *CreateGKETarget) SetGkeClusterName(v string) {
	o.GkeClusterName = &v
}

// GetGkeServiceAccountEmail returns the GkeServiceAccountEmail field value if set, zero value otherwise.
func (o *CreateGKETarget) GetGkeServiceAccountEmail() string {
	if o == nil || o.GkeServiceAccountEmail == nil {
		var ret string
		return ret
	}
	return *o.GkeServiceAccountEmail
}

// GetGkeServiceAccountEmailOk returns a tuple with the GkeServiceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetGkeServiceAccountEmailOk() (*string, bool) {
	if o == nil || o.GkeServiceAccountEmail == nil {
		return nil, false
	}
	return o.GkeServiceAccountEmail, true
}

// HasGkeServiceAccountEmail returns a boolean if a field has been set.
func (o *CreateGKETarget) HasGkeServiceAccountEmail() bool {
	if o != nil && o.GkeServiceAccountEmail != nil {
		return true
	}

	return false
}

// SetGkeServiceAccountEmail gets a reference to the given string and assigns it to the GkeServiceAccountEmail field.
func (o *CreateGKETarget) SetGkeServiceAccountEmail(v string) {
	o.GkeServiceAccountEmail = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateGKETarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateGKETarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateGKETarget) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateGKETarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateGKETarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateGKETarget) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *CreateGKETarget) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *CreateGKETarget) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *CreateGKETarget) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetName returns the Name field value
func (o *CreateGKETarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGKETarget) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateGKETarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateGKETarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateGKETarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateGKETarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateGKETarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateGKETarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUseGwCloudIdentity returns the UseGwCloudIdentity field value if set, zero value otherwise.
func (o *CreateGKETarget) GetUseGwCloudIdentity() bool {
	if o == nil || o.UseGwCloudIdentity == nil {
		var ret bool
		return ret
	}
	return *o.UseGwCloudIdentity
}

// GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETarget) GetUseGwCloudIdentityOk() (*bool, bool) {
	if o == nil || o.UseGwCloudIdentity == nil {
		return nil, false
	}
	return o.UseGwCloudIdentity, true
}

// HasUseGwCloudIdentity returns a boolean if a field has been set.
func (o *CreateGKETarget) HasUseGwCloudIdentity() bool {
	if o != nil && o.UseGwCloudIdentity != nil {
		return true
	}

	return false
}

// SetUseGwCloudIdentity gets a reference to the given bool and assigns it to the UseGwCloudIdentity field.
func (o *CreateGKETarget) SetUseGwCloudIdentity(v bool) {
	o.UseGwCloudIdentity = &v
}

func (o CreateGKETarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GkeAccountKey != nil {
		toSerialize["gke-account-key"] = o.GkeAccountKey
	}
	if o.GkeClusterCert != nil {
		toSerialize["gke-cluster-cert"] = o.GkeClusterCert
	}
	if o.GkeClusterEndpoint != nil {
		toSerialize["gke-cluster-endpoint"] = o.GkeClusterEndpoint
	}
	if o.GkeClusterName != nil {
		toSerialize["gke-cluster-name"] = o.GkeClusterName
	}
	if o.GkeServiceAccountEmail != nil {
		toSerialize["gke-service-account-email"] = o.GkeServiceAccountEmail
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.MaxVersions != nil {
		toSerialize["max-versions"] = o.MaxVersions
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UseGwCloudIdentity != nil {
		toSerialize["use-gw-cloud-identity"] = o.UseGwCloudIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGKETarget struct {
	value *CreateGKETarget
	isSet bool
}

func (v NullableCreateGKETarget) Get() *CreateGKETarget {
	return v.value
}

func (v *NullableCreateGKETarget) Set(val *CreateGKETarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGKETarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGKETarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGKETarget(val *CreateGKETarget) *NullableCreateGKETarget {
	return &NullableCreateGKETarget{value: val, isSet: true}
}

func (v NullableCreateGKETarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGKETarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


