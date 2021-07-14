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

// UpdateEKSTarget struct for UpdateEKSTarget
type UpdateEKSTarget struct {
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	// Access Key ID
	EksAccessKeyId string `json:"eks-access-key-id"`
	// EKS cluster CA certificate
	EksClusterCaCert string `json:"eks-cluster-ca-cert"`
	// EKS cluster URL endpoint
	EksClusterEndpoint string `json:"eks-cluster-endpoint"`
	// EKS cluster name
	EksClusterName string `json:"eks-cluster-name"`
	// Region
	EksRegion *string `json:"eks-region,omitempty"`
	// Secret Access Key
	EksSecretAccessKey string `json:"eks-secret-access-key"`
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

// NewUpdateEKSTarget instantiates a new UpdateEKSTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEKSTarget(eksAccessKeyId string, eksClusterCaCert string, eksClusterEndpoint string, eksClusterName string, eksSecretAccessKey string, name string, ) *UpdateEKSTarget {
	this := UpdateEKSTarget{}
	this.EksAccessKeyId = eksAccessKeyId
	this.EksClusterCaCert = eksClusterCaCert
	this.EksClusterEndpoint = eksClusterEndpoint
	this.EksClusterName = eksClusterName
	var eksRegion string = "us-east-2"
	this.EksRegion = &eksRegion
	this.EksSecretAccessKey = eksSecretAccessKey
	this.Name = name
	var updateVersion bool = false
	this.UpdateVersion = &updateVersion
	return &this
}

// NewUpdateEKSTargetWithDefaults instantiates a new UpdateEKSTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEKSTargetWithDefaults() *UpdateEKSTarget {
	this := UpdateEKSTarget{}
	var eksRegion string = "us-east-2"
	this.EksRegion = &eksRegion
	var updateVersion bool = false
	this.UpdateVersion = &updateVersion
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateEKSTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateEKSTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateEKSTarget) SetComment(v string) {
	o.Comment = &v
}

// GetEksAccessKeyId returns the EksAccessKeyId field value
func (o *UpdateEKSTarget) GetEksAccessKeyId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EksAccessKeyId
}

// GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field value
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetEksAccessKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EksAccessKeyId, true
}

// SetEksAccessKeyId sets field value
func (o *UpdateEKSTarget) SetEksAccessKeyId(v string) {
	o.EksAccessKeyId = v
}

// GetEksClusterCaCert returns the EksClusterCaCert field value
func (o *UpdateEKSTarget) GetEksClusterCaCert() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EksClusterCaCert
}

// GetEksClusterCaCertOk returns a tuple with the EksClusterCaCert field value
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetEksClusterCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EksClusterCaCert, true
}

// SetEksClusterCaCert sets field value
func (o *UpdateEKSTarget) SetEksClusterCaCert(v string) {
	o.EksClusterCaCert = v
}

// GetEksClusterEndpoint returns the EksClusterEndpoint field value
func (o *UpdateEKSTarget) GetEksClusterEndpoint() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EksClusterEndpoint
}

// GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field value
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetEksClusterEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EksClusterEndpoint, true
}

// SetEksClusterEndpoint sets field value
func (o *UpdateEKSTarget) SetEksClusterEndpoint(v string) {
	o.EksClusterEndpoint = v
}

// GetEksClusterName returns the EksClusterName field value
func (o *UpdateEKSTarget) GetEksClusterName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EksClusterName
}

// GetEksClusterNameOk returns a tuple with the EksClusterName field value
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetEksClusterNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EksClusterName, true
}

// SetEksClusterName sets field value
func (o *UpdateEKSTarget) SetEksClusterName(v string) {
	o.EksClusterName = v
}

// GetEksRegion returns the EksRegion field value if set, zero value otherwise.
func (o *UpdateEKSTarget) GetEksRegion() string {
	if o == nil || o.EksRegion == nil {
		var ret string
		return ret
	}
	return *o.EksRegion
}

// GetEksRegionOk returns a tuple with the EksRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetEksRegionOk() (*string, bool) {
	if o == nil || o.EksRegion == nil {
		return nil, false
	}
	return o.EksRegion, true
}

// HasEksRegion returns a boolean if a field has been set.
func (o *UpdateEKSTarget) HasEksRegion() bool {
	if o != nil && o.EksRegion != nil {
		return true
	}

	return false
}

// SetEksRegion gets a reference to the given string and assigns it to the EksRegion field.
func (o *UpdateEKSTarget) SetEksRegion(v string) {
	o.EksRegion = &v
}

// GetEksSecretAccessKey returns the EksSecretAccessKey field value
func (o *UpdateEKSTarget) GetEksSecretAccessKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EksSecretAccessKey
}

// GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetEksSecretAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EksSecretAccessKey, true
}

// SetEksSecretAccessKey sets field value
func (o *UpdateEKSTarget) SetEksSecretAccessKey(v string) {
	o.EksSecretAccessKey = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateEKSTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateEKSTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateEKSTarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *UpdateEKSTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateEKSTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateEKSTarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateEKSTarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateEKSTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateEKSTarget) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateEKSTarget) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateEKSTarget) SetPassword(v string) {
	o.Password = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateEKSTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateEKSTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateEKSTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateEKSTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateEKSTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateEKSTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateEKSTarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateEKSTarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateEKSTarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateEKSTarget) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTarget) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateEKSTarget) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateEKSTarget) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateEKSTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["eks-access-key-id"] = o.EksAccessKeyId
	}
	if true {
		toSerialize["eks-cluster-ca-cert"] = o.EksClusterCaCert
	}
	if true {
		toSerialize["eks-cluster-endpoint"] = o.EksClusterEndpoint
	}
	if true {
		toSerialize["eks-cluster-name"] = o.EksClusterName
	}
	if o.EksRegion != nil {
		toSerialize["eks-region"] = o.EksRegion
	}
	if true {
		toSerialize["eks-secret-access-key"] = o.EksSecretAccessKey
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

type NullableUpdateEKSTarget struct {
	value *UpdateEKSTarget
	isSet bool
}

func (v NullableUpdateEKSTarget) Get() *UpdateEKSTarget {
	return v.value
}

func (v *NullableUpdateEKSTarget) Set(val *UpdateEKSTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEKSTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEKSTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEKSTarget(val *UpdateEKSTarget) *NullableUpdateEKSTarget {
	return &NullableUpdateEKSTarget{value: val, isSet: true}
}

func (v NullableUpdateEKSTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEKSTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

