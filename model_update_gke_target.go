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

// UpdateGKETarget struct for UpdateGKETarget
type UpdateGKETarget struct {
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	// GKE Service Account key file path
	GkeAccountKey *string `json:"gke-account-key,omitempty"`
	// GKE cluster CA certificate
	GkeClusterCert string `json:"gke-cluster-cert"`
	// GKE cluster URL endpoint
	GkeClusterEndpoint string `json:"gke-cluster-endpoint"`
	// GKE cluster name
	GkeClusterName string `json:"gke-cluster-name"`
	// GKE service account email
	GkeServiceAccountEmail string `json:"gke-service-account-email"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// New target name
	NewName *string `json:"new-name,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Create new version for the target
	UpdateVersion *bool `json:"update-version,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewUpdateGKETarget instantiates a new UpdateGKETarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGKETarget(gkeClusterCert string, gkeClusterEndpoint string, gkeClusterName string, gkeServiceAccountEmail string, name string, ) *UpdateGKETarget {
	this := UpdateGKETarget{}
	this.GkeClusterCert = gkeClusterCert
	this.GkeClusterEndpoint = gkeClusterEndpoint
	this.GkeClusterName = gkeClusterName
	this.GkeServiceAccountEmail = gkeServiceAccountEmail
	this.Name = name
	var updateVersion bool = false
	this.UpdateVersion = &updateVersion
	return &this
}

// NewUpdateGKETargetWithDefaults instantiates a new UpdateGKETarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGKETargetWithDefaults() *UpdateGKETarget {
	this := UpdateGKETarget{}
	var updateVersion bool = false
	this.UpdateVersion = &updateVersion
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateGKETarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateGKETarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateGKETarget) SetComment(v string) {
	o.Comment = &v
}

// GetGkeAccountKey returns the GkeAccountKey field value if set, zero value otherwise.
func (o *UpdateGKETarget) GetGkeAccountKey() string {
	if o == nil || o.GkeAccountKey == nil {
		var ret string
		return ret
	}
	return *o.GkeAccountKey
}

// GetGkeAccountKeyOk returns a tuple with the GkeAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetGkeAccountKeyOk() (*string, bool) {
	if o == nil || o.GkeAccountKey == nil {
		return nil, false
	}
	return o.GkeAccountKey, true
}

// HasGkeAccountKey returns a boolean if a field has been set.
func (o *UpdateGKETarget) HasGkeAccountKey() bool {
	if o != nil && o.GkeAccountKey != nil {
		return true
	}

	return false
}

// SetGkeAccountKey gets a reference to the given string and assigns it to the GkeAccountKey field.
func (o *UpdateGKETarget) SetGkeAccountKey(v string) {
	o.GkeAccountKey = &v
}

// GetGkeClusterCert returns the GkeClusterCert field value
func (o *UpdateGKETarget) GetGkeClusterCert() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.GkeClusterCert
}

// GetGkeClusterCertOk returns a tuple with the GkeClusterCert field value
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetGkeClusterCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GkeClusterCert, true
}

// SetGkeClusterCert sets field value
func (o *UpdateGKETarget) SetGkeClusterCert(v string) {
	o.GkeClusterCert = v
}

// GetGkeClusterEndpoint returns the GkeClusterEndpoint field value
func (o *UpdateGKETarget) GetGkeClusterEndpoint() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.GkeClusterEndpoint
}

// GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field value
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetGkeClusterEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GkeClusterEndpoint, true
}

// SetGkeClusterEndpoint sets field value
func (o *UpdateGKETarget) SetGkeClusterEndpoint(v string) {
	o.GkeClusterEndpoint = v
}

// GetGkeClusterName returns the GkeClusterName field value
func (o *UpdateGKETarget) GetGkeClusterName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.GkeClusterName
}

// GetGkeClusterNameOk returns a tuple with the GkeClusterName field value
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetGkeClusterNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GkeClusterName, true
}

// SetGkeClusterName sets field value
func (o *UpdateGKETarget) SetGkeClusterName(v string) {
	o.GkeClusterName = v
}

// GetGkeServiceAccountEmail returns the GkeServiceAccountEmail field value
func (o *UpdateGKETarget) GetGkeServiceAccountEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.GkeServiceAccountEmail
}

// GetGkeServiceAccountEmailOk returns a tuple with the GkeServiceAccountEmail field value
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetGkeServiceAccountEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GkeServiceAccountEmail, true
}

// SetGkeServiceAccountEmail sets field value
func (o *UpdateGKETarget) SetGkeServiceAccountEmail(v string) {
	o.GkeServiceAccountEmail = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateGKETarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateGKETarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateGKETarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *UpdateGKETarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateGKETarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateGKETarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateGKETarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateGKETarget) SetNewName(v string) {
	o.NewName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateGKETarget) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateGKETarget) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateGKETarget) SetPassword(v string) {
	o.Password = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateGKETarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateGKETarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateGKETarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateGKETarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateGKETarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateGKETarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateGKETarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateGKETarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateGKETarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateGKETarget) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGKETarget) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateGKETarget) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateGKETarget) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateGKETarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.GkeAccountKey != nil {
		toSerialize["gke-account-key"] = o.GkeAccountKey
	}
	if true {
		toSerialize["gke-cluster-cert"] = o.GkeClusterCert
	}
	if true {
		toSerialize["gke-cluster-endpoint"] = o.GkeClusterEndpoint
	}
	if true {
		toSerialize["gke-cluster-name"] = o.GkeClusterName
	}
	if true {
		toSerialize["gke-service-account-email"] = o.GkeServiceAccountEmail
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UpdateVersion != nil {
		toSerialize["update-version"] = o.UpdateVersion
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateGKETarget struct {
	value *UpdateGKETarget
	isSet bool
}

func (v NullableUpdateGKETarget) Get() *UpdateGKETarget {
	return v.value
}

func (v *NullableUpdateGKETarget) Set(val *UpdateGKETarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGKETarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGKETarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGKETarget(val *UpdateGKETarget) *NullableUpdateGKETarget {
	return &NullableUpdateGKETarget{value: val, isSet: true}
}

func (v NullableUpdateGKETarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGKETarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

