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

// GetSSHCertificate struct for GetSSHCertificate
type GetSSHCertificate struct {
	// The name of the SSH certificate issuer
	CertIssuerName string `json:"cert-issuer-name"`
	// The username to sign in the SSH certificate
	CertUsername string `json:"cert-username"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Set this option to output legacy ('ssh-rsa-cert-v01@openssh.com') signing algorithm name in the certificate.
	LegacySigningAlgName *bool `json:"legacy-signing-alg-name,omitempty"`
	// SSH public key file contents. If this option is used, the certificate will be printed to stdout
	PublicKeyData *string `json:"public-key-data,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Updated certificate lifetime in seconds (must be less than the Certificate Issuer default TTL)
	Ttl *int64 `json:"ttl,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGetSSHCertificate instantiates a new GetSSHCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSSHCertificate(certIssuerName string, certUsername string, ) *GetSSHCertificate {
	this := GetSSHCertificate{}
	this.CertIssuerName = certIssuerName
	this.CertUsername = certUsername
	var json bool = false
	this.Json = &json
	var legacySigningAlgName bool = false
	this.LegacySigningAlgName = &legacySigningAlgName
	return &this
}

// NewGetSSHCertificateWithDefaults instantiates a new GetSSHCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSSHCertificateWithDefaults() *GetSSHCertificate {
	this := GetSSHCertificate{}
	var json bool = false
	this.Json = &json
	var legacySigningAlgName bool = false
	this.LegacySigningAlgName = &legacySigningAlgName
	return &this
}

// GetCertIssuerName returns the CertIssuerName field value
func (o *GetSSHCertificate) GetCertIssuerName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CertIssuerName
}

// GetCertIssuerNameOk returns a tuple with the CertIssuerName field value
// and a boolean to check if the value has been set.
func (o *GetSSHCertificate) GetCertIssuerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertIssuerName, true
}

// SetCertIssuerName sets field value
func (o *GetSSHCertificate) SetCertIssuerName(v string) {
	o.CertIssuerName = v
}

// GetCertUsername returns the CertUsername field value
func (o *GetSSHCertificate) GetCertUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CertUsername
}

// GetCertUsernameOk returns a tuple with the CertUsername field value
// and a boolean to check if the value has been set.
func (o *GetSSHCertificate) GetCertUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertUsername, true
}

// SetCertUsername sets field value
func (o *GetSSHCertificate) SetCertUsername(v string) {
	o.CertUsername = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GetSSHCertificate) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSSHCertificate) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GetSSHCertificate) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GetSSHCertificate) SetJson(v bool) {
	o.Json = &v
}

// GetLegacySigningAlgName returns the LegacySigningAlgName field value if set, zero value otherwise.
func (o *GetSSHCertificate) GetLegacySigningAlgName() bool {
	if o == nil || o.LegacySigningAlgName == nil {
		var ret bool
		return ret
	}
	return *o.LegacySigningAlgName
}

// GetLegacySigningAlgNameOk returns a tuple with the LegacySigningAlgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSSHCertificate) GetLegacySigningAlgNameOk() (*bool, bool) {
	if o == nil || o.LegacySigningAlgName == nil {
		return nil, false
	}
	return o.LegacySigningAlgName, true
}

// HasLegacySigningAlgName returns a boolean if a field has been set.
func (o *GetSSHCertificate) HasLegacySigningAlgName() bool {
	if o != nil && o.LegacySigningAlgName != nil {
		return true
	}

	return false
}

// SetLegacySigningAlgName gets a reference to the given bool and assigns it to the LegacySigningAlgName field.
func (o *GetSSHCertificate) SetLegacySigningAlgName(v bool) {
	o.LegacySigningAlgName = &v
}

// GetPublicKeyData returns the PublicKeyData field value if set, zero value otherwise.
func (o *GetSSHCertificate) GetPublicKeyData() string {
	if o == nil || o.PublicKeyData == nil {
		var ret string
		return ret
	}
	return *o.PublicKeyData
}

// GetPublicKeyDataOk returns a tuple with the PublicKeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSSHCertificate) GetPublicKeyDataOk() (*string, bool) {
	if o == nil || o.PublicKeyData == nil {
		return nil, false
	}
	return o.PublicKeyData, true
}

// HasPublicKeyData returns a boolean if a field has been set.
func (o *GetSSHCertificate) HasPublicKeyData() bool {
	if o != nil && o.PublicKeyData != nil {
		return true
	}

	return false
}

// SetPublicKeyData gets a reference to the given string and assigns it to the PublicKeyData field.
func (o *GetSSHCertificate) SetPublicKeyData(v string) {
	o.PublicKeyData = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetSSHCertificate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSSHCertificate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetSSHCertificate) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetSSHCertificate) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *GetSSHCertificate) GetTtl() int64 {
	if o == nil || o.Ttl == nil {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSSHCertificate) GetTtlOk() (*int64, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *GetSSHCertificate) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *GetSSHCertificate) SetTtl(v int64) {
	o.Ttl = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GetSSHCertificate) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSSHCertificate) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GetSSHCertificate) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GetSSHCertificate) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GetSSHCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cert-issuer-name"] = o.CertIssuerName
	}
	if true {
		toSerialize["cert-username"] = o.CertUsername
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.LegacySigningAlgName != nil {
		toSerialize["legacy-signing-alg-name"] = o.LegacySigningAlgName
	}
	if o.PublicKeyData != nil {
		toSerialize["public-key-data"] = o.PublicKeyData
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGetSSHCertificate struct {
	value *GetSSHCertificate
	isSet bool
}

func (v NullableGetSSHCertificate) Get() *GetSSHCertificate {
	return v.value
}

func (v *NullableGetSSHCertificate) Set(val *GetSSHCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSSHCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSSHCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSSHCertificate(val *GetSSHCertificate) *NullableGetSSHCertificate {
	return &NullableGetSSHCertificate{value: val, isSet: true}
}

func (v NullableGetSSHCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSSHCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


