/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GatewayGetK8SAuthConfigOutput struct for GatewayGetK8SAuthConfigOutput
type GatewayGetK8SAuthConfigOutput struct {
	// AuthMethodTokenExpiration is time in seconds of expiration of the Akeyless Kube Auth Method token
	AmTokenExpiration *int64 `json:"am_token_expiration,omitempty"`
	// AuthMethodAccessId of the Kubernetes auth method
	AuthMethodAccessId *string `json:"auth_method_access_id,omitempty"`
	// AuthMethodSigningKey is the private key (in base64 of the PEM format) associated with the public key defined in the Kubernetes auth method, that used to sign the internal token for the Akeyless Kubernetes Auth Method
	AuthMethodPrvKeyPem *string `json:"auth_method_prv_key_pem,omitempty"`
	// ClusterApiType defines types of API access to cluster
	ClusterApiType *string `json:"cluster_api_type,omitempty"`
	// DisableISSValidation is optional parameter to disable ISS validation
	DisableIssValidation *bool `json:"disable_iss_validation,omitempty"`
	Id *string `json:"id,omitempty"`
	// K8SCACert is the CA Cert to use to call into the kubernetes API
	K8sCaCert *string `json:"k8s_ca_cert,omitempty"`
	// K8SHost is the url string for the kubernetes API
	K8sHost *string `json:"k8s_host,omitempty"`
	// K8SIssuer is the claim that specifies who issued the Kubernetes token
	K8sIssuer *string `json:"k8s_issuer,omitempty"`
	// K8SPublicKeysPEM is the list of public key in PEM format
	K8sPubKeysPem []string `json:"k8s_pub_keys_pem,omitempty"`
	// K8STokenReviewerJWT is the bearer for clusterApiTypeK8s, used during TokenReview API call
	K8sTokenReviewerJwt *string `json:"k8s_token_reviewer_jwt,omitempty"`
	Name *string `json:"name,omitempty"`
	ProtectionKey *string `json:"protection_key,omitempty"`
	// RancherApiKey the bear token for clusterApiTypeRancher
	RancherApiKey *string `json:"rancher_api_key,omitempty"`
	// RancherClusterId cluster id as define in rancher (in case of clusterApiTypeRancher)
	RancherClusterId *string `json:"rancher_cluster_id,omitempty"`
	// UseLocalCAJwt is an optional parameter to set defaulting to using the local service account when running in a Kubernetes pod
	UseLocalCaJwt *bool `json:"use_local_ca_jwt,omitempty"`
}

// NewGatewayGetK8SAuthConfigOutput instantiates a new GatewayGetK8SAuthConfigOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayGetK8SAuthConfigOutput() *GatewayGetK8SAuthConfigOutput {
	this := GatewayGetK8SAuthConfigOutput{}
	return &this
}

// NewGatewayGetK8SAuthConfigOutputWithDefaults instantiates a new GatewayGetK8SAuthConfigOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayGetK8SAuthConfigOutputWithDefaults() *GatewayGetK8SAuthConfigOutput {
	this := GatewayGetK8SAuthConfigOutput{}
	return &this
}

// GetAmTokenExpiration returns the AmTokenExpiration field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetAmTokenExpiration() int64 {
	if o == nil || o.AmTokenExpiration == nil {
		var ret int64
		return ret
	}
	return *o.AmTokenExpiration
}

// GetAmTokenExpirationOk returns a tuple with the AmTokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetAmTokenExpirationOk() (*int64, bool) {
	if o == nil || o.AmTokenExpiration == nil {
		return nil, false
	}
	return o.AmTokenExpiration, true
}

// HasAmTokenExpiration returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasAmTokenExpiration() bool {
	if o != nil && o.AmTokenExpiration != nil {
		return true
	}

	return false
}

// SetAmTokenExpiration gets a reference to the given int64 and assigns it to the AmTokenExpiration field.
func (o *GatewayGetK8SAuthConfigOutput) SetAmTokenExpiration(v int64) {
	o.AmTokenExpiration = &v
}

// GetAuthMethodAccessId returns the AuthMethodAccessId field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetAuthMethodAccessId() string {
	if o == nil || o.AuthMethodAccessId == nil {
		var ret string
		return ret
	}
	return *o.AuthMethodAccessId
}

// GetAuthMethodAccessIdOk returns a tuple with the AuthMethodAccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetAuthMethodAccessIdOk() (*string, bool) {
	if o == nil || o.AuthMethodAccessId == nil {
		return nil, false
	}
	return o.AuthMethodAccessId, true
}

// HasAuthMethodAccessId returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasAuthMethodAccessId() bool {
	if o != nil && o.AuthMethodAccessId != nil {
		return true
	}

	return false
}

// SetAuthMethodAccessId gets a reference to the given string and assigns it to the AuthMethodAccessId field.
func (o *GatewayGetK8SAuthConfigOutput) SetAuthMethodAccessId(v string) {
	o.AuthMethodAccessId = &v
}

// GetAuthMethodPrvKeyPem returns the AuthMethodPrvKeyPem field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetAuthMethodPrvKeyPem() string {
	if o == nil || o.AuthMethodPrvKeyPem == nil {
		var ret string
		return ret
	}
	return *o.AuthMethodPrvKeyPem
}

// GetAuthMethodPrvKeyPemOk returns a tuple with the AuthMethodPrvKeyPem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetAuthMethodPrvKeyPemOk() (*string, bool) {
	if o == nil || o.AuthMethodPrvKeyPem == nil {
		return nil, false
	}
	return o.AuthMethodPrvKeyPem, true
}

// HasAuthMethodPrvKeyPem returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasAuthMethodPrvKeyPem() bool {
	if o != nil && o.AuthMethodPrvKeyPem != nil {
		return true
	}

	return false
}

// SetAuthMethodPrvKeyPem gets a reference to the given string and assigns it to the AuthMethodPrvKeyPem field.
func (o *GatewayGetK8SAuthConfigOutput) SetAuthMethodPrvKeyPem(v string) {
	o.AuthMethodPrvKeyPem = &v
}

// GetClusterApiType returns the ClusterApiType field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetClusterApiType() string {
	if o == nil || o.ClusterApiType == nil {
		var ret string
		return ret
	}
	return *o.ClusterApiType
}

// GetClusterApiTypeOk returns a tuple with the ClusterApiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetClusterApiTypeOk() (*string, bool) {
	if o == nil || o.ClusterApiType == nil {
		return nil, false
	}
	return o.ClusterApiType, true
}

// HasClusterApiType returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasClusterApiType() bool {
	if o != nil && o.ClusterApiType != nil {
		return true
	}

	return false
}

// SetClusterApiType gets a reference to the given string and assigns it to the ClusterApiType field.
func (o *GatewayGetK8SAuthConfigOutput) SetClusterApiType(v string) {
	o.ClusterApiType = &v
}

// GetDisableIssValidation returns the DisableIssValidation field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetDisableIssValidation() bool {
	if o == nil || o.DisableIssValidation == nil {
		var ret bool
		return ret
	}
	return *o.DisableIssValidation
}

// GetDisableIssValidationOk returns a tuple with the DisableIssValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetDisableIssValidationOk() (*bool, bool) {
	if o == nil || o.DisableIssValidation == nil {
		return nil, false
	}
	return o.DisableIssValidation, true
}

// HasDisableIssValidation returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasDisableIssValidation() bool {
	if o != nil && o.DisableIssValidation != nil {
		return true
	}

	return false
}

// SetDisableIssValidation gets a reference to the given bool and assigns it to the DisableIssValidation field.
func (o *GatewayGetK8SAuthConfigOutput) SetDisableIssValidation(v bool) {
	o.DisableIssValidation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GatewayGetK8SAuthConfigOutput) SetId(v string) {
	o.Id = &v
}

// GetK8sCaCert returns the K8sCaCert field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sCaCert() string {
	if o == nil || o.K8sCaCert == nil {
		var ret string
		return ret
	}
	return *o.K8sCaCert
}

// GetK8sCaCertOk returns a tuple with the K8sCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sCaCertOk() (*string, bool) {
	if o == nil || o.K8sCaCert == nil {
		return nil, false
	}
	return o.K8sCaCert, true
}

// HasK8sCaCert returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasK8sCaCert() bool {
	if o != nil && o.K8sCaCert != nil {
		return true
	}

	return false
}

// SetK8sCaCert gets a reference to the given string and assigns it to the K8sCaCert field.
func (o *GatewayGetK8SAuthConfigOutput) SetK8sCaCert(v string) {
	o.K8sCaCert = &v
}

// GetK8sHost returns the K8sHost field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sHost() string {
	if o == nil || o.K8sHost == nil {
		var ret string
		return ret
	}
	return *o.K8sHost
}

// GetK8sHostOk returns a tuple with the K8sHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sHostOk() (*string, bool) {
	if o == nil || o.K8sHost == nil {
		return nil, false
	}
	return o.K8sHost, true
}

// HasK8sHost returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasK8sHost() bool {
	if o != nil && o.K8sHost != nil {
		return true
	}

	return false
}

// SetK8sHost gets a reference to the given string and assigns it to the K8sHost field.
func (o *GatewayGetK8SAuthConfigOutput) SetK8sHost(v string) {
	o.K8sHost = &v
}

// GetK8sIssuer returns the K8sIssuer field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sIssuer() string {
	if o == nil || o.K8sIssuer == nil {
		var ret string
		return ret
	}
	return *o.K8sIssuer
}

// GetK8sIssuerOk returns a tuple with the K8sIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sIssuerOk() (*string, bool) {
	if o == nil || o.K8sIssuer == nil {
		return nil, false
	}
	return o.K8sIssuer, true
}

// HasK8sIssuer returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasK8sIssuer() bool {
	if o != nil && o.K8sIssuer != nil {
		return true
	}

	return false
}

// SetK8sIssuer gets a reference to the given string and assigns it to the K8sIssuer field.
func (o *GatewayGetK8SAuthConfigOutput) SetK8sIssuer(v string) {
	o.K8sIssuer = &v
}

// GetK8sPubKeysPem returns the K8sPubKeysPem field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sPubKeysPem() []string {
	if o == nil || o.K8sPubKeysPem == nil {
		var ret []string
		return ret
	}
	return o.K8sPubKeysPem
}

// GetK8sPubKeysPemOk returns a tuple with the K8sPubKeysPem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sPubKeysPemOk() ([]string, bool) {
	if o == nil || o.K8sPubKeysPem == nil {
		return nil, false
	}
	return o.K8sPubKeysPem, true
}

// HasK8sPubKeysPem returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasK8sPubKeysPem() bool {
	if o != nil && o.K8sPubKeysPem != nil {
		return true
	}

	return false
}

// SetK8sPubKeysPem gets a reference to the given []string and assigns it to the K8sPubKeysPem field.
func (o *GatewayGetK8SAuthConfigOutput) SetK8sPubKeysPem(v []string) {
	o.K8sPubKeysPem = v
}

// GetK8sTokenReviewerJwt returns the K8sTokenReviewerJwt field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sTokenReviewerJwt() string {
	if o == nil || o.K8sTokenReviewerJwt == nil {
		var ret string
		return ret
	}
	return *o.K8sTokenReviewerJwt
}

// GetK8sTokenReviewerJwtOk returns a tuple with the K8sTokenReviewerJwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetK8sTokenReviewerJwtOk() (*string, bool) {
	if o == nil || o.K8sTokenReviewerJwt == nil {
		return nil, false
	}
	return o.K8sTokenReviewerJwt, true
}

// HasK8sTokenReviewerJwt returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasK8sTokenReviewerJwt() bool {
	if o != nil && o.K8sTokenReviewerJwt != nil {
		return true
	}

	return false
}

// SetK8sTokenReviewerJwt gets a reference to the given string and assigns it to the K8sTokenReviewerJwt field.
func (o *GatewayGetK8SAuthConfigOutput) SetK8sTokenReviewerJwt(v string) {
	o.K8sTokenReviewerJwt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GatewayGetK8SAuthConfigOutput) SetName(v string) {
	o.Name = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetProtectionKey() string {
	if o == nil || o.ProtectionKey == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetProtectionKeyOk() (*string, bool) {
	if o == nil || o.ProtectionKey == nil {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasProtectionKey() bool {
	if o != nil && o.ProtectionKey != nil {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *GatewayGetK8SAuthConfigOutput) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetRancherApiKey returns the RancherApiKey field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetRancherApiKey() string {
	if o == nil || o.RancherApiKey == nil {
		var ret string
		return ret
	}
	return *o.RancherApiKey
}

// GetRancherApiKeyOk returns a tuple with the RancherApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetRancherApiKeyOk() (*string, bool) {
	if o == nil || o.RancherApiKey == nil {
		return nil, false
	}
	return o.RancherApiKey, true
}

// HasRancherApiKey returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasRancherApiKey() bool {
	if o != nil && o.RancherApiKey != nil {
		return true
	}

	return false
}

// SetRancherApiKey gets a reference to the given string and assigns it to the RancherApiKey field.
func (o *GatewayGetK8SAuthConfigOutput) SetRancherApiKey(v string) {
	o.RancherApiKey = &v
}

// GetRancherClusterId returns the RancherClusterId field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetRancherClusterId() string {
	if o == nil || o.RancherClusterId == nil {
		var ret string
		return ret
	}
	return *o.RancherClusterId
}

// GetRancherClusterIdOk returns a tuple with the RancherClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetRancherClusterIdOk() (*string, bool) {
	if o == nil || o.RancherClusterId == nil {
		return nil, false
	}
	return o.RancherClusterId, true
}

// HasRancherClusterId returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasRancherClusterId() bool {
	if o != nil && o.RancherClusterId != nil {
		return true
	}

	return false
}

// SetRancherClusterId gets a reference to the given string and assigns it to the RancherClusterId field.
func (o *GatewayGetK8SAuthConfigOutput) SetRancherClusterId(v string) {
	o.RancherClusterId = &v
}

// GetUseLocalCaJwt returns the UseLocalCaJwt field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfigOutput) GetUseLocalCaJwt() bool {
	if o == nil || o.UseLocalCaJwt == nil {
		var ret bool
		return ret
	}
	return *o.UseLocalCaJwt
}

// GetUseLocalCaJwtOk returns a tuple with the UseLocalCaJwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfigOutput) GetUseLocalCaJwtOk() (*bool, bool) {
	if o == nil || o.UseLocalCaJwt == nil {
		return nil, false
	}
	return o.UseLocalCaJwt, true
}

// HasUseLocalCaJwt returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfigOutput) HasUseLocalCaJwt() bool {
	if o != nil && o.UseLocalCaJwt != nil {
		return true
	}

	return false
}

// SetUseLocalCaJwt gets a reference to the given bool and assigns it to the UseLocalCaJwt field.
func (o *GatewayGetK8SAuthConfigOutput) SetUseLocalCaJwt(v bool) {
	o.UseLocalCaJwt = &v
}

func (o GatewayGetK8SAuthConfigOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmTokenExpiration != nil {
		toSerialize["am_token_expiration"] = o.AmTokenExpiration
	}
	if o.AuthMethodAccessId != nil {
		toSerialize["auth_method_access_id"] = o.AuthMethodAccessId
	}
	if o.AuthMethodPrvKeyPem != nil {
		toSerialize["auth_method_prv_key_pem"] = o.AuthMethodPrvKeyPem
	}
	if o.ClusterApiType != nil {
		toSerialize["cluster_api_type"] = o.ClusterApiType
	}
	if o.DisableIssValidation != nil {
		toSerialize["disable_iss_validation"] = o.DisableIssValidation
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.K8sCaCert != nil {
		toSerialize["k8s_ca_cert"] = o.K8sCaCert
	}
	if o.K8sHost != nil {
		toSerialize["k8s_host"] = o.K8sHost
	}
	if o.K8sIssuer != nil {
		toSerialize["k8s_issuer"] = o.K8sIssuer
	}
	if o.K8sPubKeysPem != nil {
		toSerialize["k8s_pub_keys_pem"] = o.K8sPubKeysPem
	}
	if o.K8sTokenReviewerJwt != nil {
		toSerialize["k8s_token_reviewer_jwt"] = o.K8sTokenReviewerJwt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProtectionKey != nil {
		toSerialize["protection_key"] = o.ProtectionKey
	}
	if o.RancherApiKey != nil {
		toSerialize["rancher_api_key"] = o.RancherApiKey
	}
	if o.RancherClusterId != nil {
		toSerialize["rancher_cluster_id"] = o.RancherClusterId
	}
	if o.UseLocalCaJwt != nil {
		toSerialize["use_local_ca_jwt"] = o.UseLocalCaJwt
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayGetK8SAuthConfigOutput struct {
	value *GatewayGetK8SAuthConfigOutput
	isSet bool
}

func (v NullableGatewayGetK8SAuthConfigOutput) Get() *GatewayGetK8SAuthConfigOutput {
	return v.value
}

func (v *NullableGatewayGetK8SAuthConfigOutput) Set(val *GatewayGetK8SAuthConfigOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayGetK8SAuthConfigOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayGetK8SAuthConfigOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayGetK8SAuthConfigOutput(val *GatewayGetK8SAuthConfigOutput) *NullableGatewayGetK8SAuthConfigOutput {
	return &NullableGatewayGetK8SAuthConfigOutput{value: val, isSet: true}
}

func (v NullableGatewayGetK8SAuthConfigOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayGetK8SAuthConfigOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


