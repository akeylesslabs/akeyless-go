/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 0.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GetKubeExecCreds getKubeExecCreds is a command that gets credentials for authentication with Kubernetes cluster based on a PKI cert issuer.
type GetKubeExecCreds struct {
	// The Subject Alternative Names to be included in the PKI certificate (in a comma-delimited list)
	AltNames *string `json:"alt-names,omitempty"`
	// The name of the PKI certificate issuer
	CertIssuerName string `json:"cert-issuer-name"`
	// The common name to be included in the PKI certificate
	CommonName *string `json:"common-name,omitempty"`
	// The client public or private key file path (in case of a private key, it will be use to extract the public key)
	KeyFilePath string `json:"key-file-path"`
	// Output file path with the certificate. If not provided, the file with the certificate will be created in the same location of the provided public key with the -cert extension
	Outfile *string `json:"outfile,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The URI Subject Alternative Names to be included in the PKI certificate (in a comma-delimited list)
	UriSans *string `json:"uri-sans,omitempty"`
}

// NewGetKubeExecCreds instantiates a new GetKubeExecCreds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKubeExecCreds(certIssuerName string, keyFilePath string, ) *GetKubeExecCreds {
	this := GetKubeExecCreds{}
	this.CertIssuerName = certIssuerName
	this.KeyFilePath = keyFilePath
	return &this
}

// NewGetKubeExecCredsWithDefaults instantiates a new GetKubeExecCreds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKubeExecCredsWithDefaults() *GetKubeExecCreds {
	this := GetKubeExecCreds{}
	return &this
}

// GetAltNames returns the AltNames field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetAltNames() string {
	if o == nil || o.AltNames == nil {
		var ret string
		return ret
	}
	return *o.AltNames
}

// GetAltNamesOk returns a tuple with the AltNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetAltNamesOk() (*string, bool) {
	if o == nil || o.AltNames == nil {
		return nil, false
	}
	return o.AltNames, true
}

// HasAltNames returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasAltNames() bool {
	if o != nil && o.AltNames != nil {
		return true
	}

	return false
}

// SetAltNames gets a reference to the given string and assigns it to the AltNames field.
func (o *GetKubeExecCreds) SetAltNames(v string) {
	o.AltNames = &v
}

// GetCertIssuerName returns the CertIssuerName field value
func (o *GetKubeExecCreds) GetCertIssuerName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CertIssuerName
}

// GetCertIssuerNameOk returns a tuple with the CertIssuerName field value
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetCertIssuerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertIssuerName, true
}

// SetCertIssuerName sets field value
func (o *GetKubeExecCreds) SetCertIssuerName(v string) {
	o.CertIssuerName = v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *GetKubeExecCreds) SetCommonName(v string) {
	o.CommonName = &v
}

// GetKeyFilePath returns the KeyFilePath field value
func (o *GetKubeExecCreds) GetKeyFilePath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.KeyFilePath
}

// GetKeyFilePathOk returns a tuple with the KeyFilePath field value
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetKeyFilePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyFilePath, true
}

// SetKeyFilePath sets field value
func (o *GetKubeExecCreds) SetKeyFilePath(v string) {
	o.KeyFilePath = v
}

// GetOutfile returns the Outfile field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetOutfile() string {
	if o == nil || o.Outfile == nil {
		var ret string
		return ret
	}
	return *o.Outfile
}

// GetOutfileOk returns a tuple with the Outfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetOutfileOk() (*string, bool) {
	if o == nil || o.Outfile == nil {
		return nil, false
	}
	return o.Outfile, true
}

// HasOutfile returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasOutfile() bool {
	if o != nil && o.Outfile != nil {
		return true
	}

	return false
}

// SetOutfile gets a reference to the given string and assigns it to the Outfile field.
func (o *GetKubeExecCreds) SetOutfile(v string) {
	o.Outfile = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetKubeExecCreds) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GetKubeExecCreds) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUriSans returns the UriSans field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetUriSans() string {
	if o == nil || o.UriSans == nil {
		var ret string
		return ret
	}
	return *o.UriSans
}

// GetUriSansOk returns a tuple with the UriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetUriSansOk() (*string, bool) {
	if o == nil || o.UriSans == nil {
		return nil, false
	}
	return o.UriSans, true
}

// HasUriSans returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasUriSans() bool {
	if o != nil && o.UriSans != nil {
		return true
	}

	return false
}

// SetUriSans gets a reference to the given string and assigns it to the UriSans field.
func (o *GetKubeExecCreds) SetUriSans(v string) {
	o.UriSans = &v
}

func (o GetKubeExecCreds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AltNames != nil {
		toSerialize["alt-names"] = o.AltNames
	}
	if true {
		toSerialize["cert-issuer-name"] = o.CertIssuerName
	}
	if o.CommonName != nil {
		toSerialize["common-name"] = o.CommonName
	}
	if true {
		toSerialize["key-file-path"] = o.KeyFilePath
	}
	if o.Outfile != nil {
		toSerialize["outfile"] = o.Outfile
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UriSans != nil {
		toSerialize["uri-sans"] = o.UriSans
	}
	return json.Marshal(toSerialize)
}

type NullableGetKubeExecCreds struct {
	value *GetKubeExecCreds
	isSet bool
}

func (v NullableGetKubeExecCreds) Get() *GetKubeExecCreds {
	return v.value
}

func (v *NullableGetKubeExecCreds) Set(val *GetKubeExecCreds) {
	v.value = val
	v.isSet = true
}

func (v NullableGetKubeExecCreds) IsSet() bool {
	return v.isSet
}

func (v *NullableGetKubeExecCreds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetKubeExecCreds(val *GetKubeExecCreds) *NullableGetKubeExecCreds {
	return &NullableGetKubeExecCreds{value: val, isSet: true}
}

func (v NullableGetKubeExecCreds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetKubeExecCreds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


