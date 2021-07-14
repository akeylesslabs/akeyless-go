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

// CreateNativeK8STarget struct for CreateNativeK8STarget
type CreateNativeK8STarget struct {
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	// K8S cluster CA certificate
	K8sClusterCaCert string `json:"k8s-cluster-ca-cert"`
	// K8S cluster URL endpoint
	K8sClusterEndpoint string `json:"k8s-cluster-endpoint"`
	// K8S cluster Bearer token
	K8sClusterToken string `json:"k8s-cluster-token"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewCreateNativeK8STarget instantiates a new CreateNativeK8STarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNativeK8STarget(k8sClusterCaCert string, k8sClusterEndpoint string, k8sClusterToken string, name string, ) *CreateNativeK8STarget {
	this := CreateNativeK8STarget{}
	this.K8sClusterCaCert = k8sClusterCaCert
	this.K8sClusterEndpoint = k8sClusterEndpoint
	this.K8sClusterToken = k8sClusterToken
	this.Name = name
	return &this
}

// NewCreateNativeK8STargetWithDefaults instantiates a new CreateNativeK8STarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNativeK8STargetWithDefaults() *CreateNativeK8STarget {
	this := CreateNativeK8STarget{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateNativeK8STarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateNativeK8STarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateNativeK8STarget) SetComment(v string) {
	o.Comment = &v
}

// GetK8sClusterCaCert returns the K8sClusterCaCert field value
func (o *CreateNativeK8STarget) GetK8sClusterCaCert() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sClusterCaCert
}

// GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field value
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetK8sClusterCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sClusterCaCert, true
}

// SetK8sClusterCaCert sets field value
func (o *CreateNativeK8STarget) SetK8sClusterCaCert(v string) {
	o.K8sClusterCaCert = v
}

// GetK8sClusterEndpoint returns the K8sClusterEndpoint field value
func (o *CreateNativeK8STarget) GetK8sClusterEndpoint() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sClusterEndpoint
}

// GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field value
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetK8sClusterEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sClusterEndpoint, true
}

// SetK8sClusterEndpoint sets field value
func (o *CreateNativeK8STarget) SetK8sClusterEndpoint(v string) {
	o.K8sClusterEndpoint = v
}

// GetK8sClusterToken returns the K8sClusterToken field value
func (o *CreateNativeK8STarget) GetK8sClusterToken() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sClusterToken
}

// GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field value
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetK8sClusterTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sClusterToken, true
}

// SetK8sClusterToken sets field value
func (o *CreateNativeK8STarget) SetK8sClusterToken(v string) {
	o.K8sClusterToken = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateNativeK8STarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateNativeK8STarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateNativeK8STarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *CreateNativeK8STarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNativeK8STarget) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateNativeK8STarget) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateNativeK8STarget) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateNativeK8STarget) SetPassword(v string) {
	o.Password = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateNativeK8STarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateNativeK8STarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateNativeK8STarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateNativeK8STarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateNativeK8STarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateNativeK8STarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateNativeK8STarget) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STarget) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateNativeK8STarget) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateNativeK8STarget) SetUsername(v string) {
	o.Username = &v
}

func (o CreateNativeK8STarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["k8s-cluster-ca-cert"] = o.K8sClusterCaCert
	}
	if true {
		toSerialize["k8s-cluster-endpoint"] = o.K8sClusterEndpoint
	}
	if true {
		toSerialize["k8s-cluster-token"] = o.K8sClusterToken
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

type NullableCreateNativeK8STarget struct {
	value *CreateNativeK8STarget
	isSet bool
}

func (v NullableCreateNativeK8STarget) Get() *CreateNativeK8STarget {
	return v.value
}

func (v *NullableCreateNativeK8STarget) Set(val *CreateNativeK8STarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNativeK8STarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNativeK8STarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNativeK8STarget(val *CreateNativeK8STarget) *NullableCreateNativeK8STarget {
	return &NullableCreateNativeK8STarget{value: val, isSet: true}
}

func (v NullableCreateNativeK8STarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNativeK8STarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

