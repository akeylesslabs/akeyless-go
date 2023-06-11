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

// checks if the GetKubeExecCreds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetKubeExecCreds{}

// GetKubeExecCreds getKubeExecCreds is a command that gets credentials for authentication with Kubernetes cluster based on a PKI cert issuer.
type GetKubeExecCreds struct {
	// The Subject Alternative Names to be included in the PKI certificate (in a comma-separated list) (if CSR is supplied this flag is ignored and any DNS.* names are taken from it)
	AltNames *string `json:"alt-names,omitempty"`
	// Client authentication API version
	ApiVersion *string `json:"api-version,omitempty"`
	// The name of the PKI certificate issuer
	CertIssuerName string `json:"cert-issuer-name"`
	// The common name to be included in the PKI certificate (if CSR is supplied this flag is ignored and the CSR subject CN is taken)
	CommonName *string `json:"common-name,omitempty"`
	// Certificate Signing Request contents encoded in base64 to generate the certificate with
	CsrDataBase64 *string `json:"csr-data-base64,omitempty"`
	// A comma-separated list of extended key usage requests which will be used for certificate issuance. Supported values: 'clientauth', 'serverauth'.
	ExtendedKeyUsage *string `json:"extended-key-usage,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// PKI key file contents. If this option is used, the certificate will be printed to stdout
	KeyDataBase64 *string `json:"key-data-base64,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Updated certificate lifetime in seconds (must be less than the Certificate Issuer default TTL)
	Ttl *int64 `json:"ttl,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The URI Subject Alternative Names to be included in the PKI certificate (in a comma-separated list) (if CSR is supplied this flag is ignored and any URI.* names are taken from it)
	UriSans *string `json:"uri-sans,omitempty"`
}

// NewGetKubeExecCreds instantiates a new GetKubeExecCreds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKubeExecCreds(certIssuerName string) *GetKubeExecCreds {
	this := GetKubeExecCreds{}
	var apiVersion string = "v1"
	this.ApiVersion = &apiVersion
	this.CertIssuerName = certIssuerName
	var json bool = false
	this.Json = &json
	return &this
}

// NewGetKubeExecCredsWithDefaults instantiates a new GetKubeExecCreds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKubeExecCredsWithDefaults() *GetKubeExecCreds {
	this := GetKubeExecCreds{}
	var apiVersion string = "v1"
	this.ApiVersion = &apiVersion
	var json bool = false
	this.Json = &json
	return &this
}

// GetAltNames returns the AltNames field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetAltNames() string {
	if o == nil || IsNil(o.AltNames) {
		var ret string
		return ret
	}
	return *o.AltNames
}

// GetAltNamesOk returns a tuple with the AltNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetAltNamesOk() (*string, bool) {
	if o == nil || IsNil(o.AltNames) {
		return nil, false
	}
	return o.AltNames, true
}

// HasAltNames returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasAltNames() bool {
	if o != nil && !IsNil(o.AltNames) {
		return true
	}

	return false
}

// SetAltNames gets a reference to the given string and assigns it to the AltNames field.
func (o *GetKubeExecCreds) SetAltNames(v string) {
	o.AltNames = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *GetKubeExecCreds) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetCertIssuerName returns the CertIssuerName field value
func (o *GetKubeExecCreds) GetCertIssuerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertIssuerName
}

// GetCertIssuerNameOk returns a tuple with the CertIssuerName field value
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetCertIssuerNameOk() (*string, bool) {
	if o == nil {
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
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *GetKubeExecCreds) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCsrDataBase64 returns the CsrDataBase64 field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetCsrDataBase64() string {
	if o == nil || IsNil(o.CsrDataBase64) {
		var ret string
		return ret
	}
	return *o.CsrDataBase64
}

// GetCsrDataBase64Ok returns a tuple with the CsrDataBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetCsrDataBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.CsrDataBase64) {
		return nil, false
	}
	return o.CsrDataBase64, true
}

// HasCsrDataBase64 returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasCsrDataBase64() bool {
	if o != nil && !IsNil(o.CsrDataBase64) {
		return true
	}

	return false
}

// SetCsrDataBase64 gets a reference to the given string and assigns it to the CsrDataBase64 field.
func (o *GetKubeExecCreds) SetCsrDataBase64(v string) {
	o.CsrDataBase64 = &v
}

// GetExtendedKeyUsage returns the ExtendedKeyUsage field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetExtendedKeyUsage() string {
	if o == nil || IsNil(o.ExtendedKeyUsage) {
		var ret string
		return ret
	}
	return *o.ExtendedKeyUsage
}

// GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetExtendedKeyUsageOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedKeyUsage) {
		return nil, false
	}
	return o.ExtendedKeyUsage, true
}

// HasExtendedKeyUsage returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasExtendedKeyUsage() bool {
	if o != nil && !IsNil(o.ExtendedKeyUsage) {
		return true
	}

	return false
}

// SetExtendedKeyUsage gets a reference to the given string and assigns it to the ExtendedKeyUsage field.
func (o *GetKubeExecCreds) SetExtendedKeyUsage(v string) {
	o.ExtendedKeyUsage = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GetKubeExecCreds) SetJson(v bool) {
	o.Json = &v
}

// GetKeyDataBase64 returns the KeyDataBase64 field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetKeyDataBase64() string {
	if o == nil || IsNil(o.KeyDataBase64) {
		var ret string
		return ret
	}
	return *o.KeyDataBase64
}

// GetKeyDataBase64Ok returns a tuple with the KeyDataBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetKeyDataBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.KeyDataBase64) {
		return nil, false
	}
	return o.KeyDataBase64, true
}

// HasKeyDataBase64 returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasKeyDataBase64() bool {
	if o != nil && !IsNil(o.KeyDataBase64) {
		return true
	}

	return false
}

// SetKeyDataBase64 gets a reference to the given string and assigns it to the KeyDataBase64 field.
func (o *GetKubeExecCreds) SetKeyDataBase64(v string) {
	o.KeyDataBase64 = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetKubeExecCreds) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetTtl() int64 {
	if o == nil || IsNil(o.Ttl) {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *GetKubeExecCreds) SetTtl(v int64) {
	o.Ttl = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GetKubeExecCreds) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
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
	if o == nil || IsNil(o.UriSans) {
		var ret string
		return ret
	}
	return *o.UriSans
}

// GetUriSansOk returns a tuple with the UriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCreds) GetUriSansOk() (*string, bool) {
	if o == nil || IsNil(o.UriSans) {
		return nil, false
	}
	return o.UriSans, true
}

// HasUriSans returns a boolean if a field has been set.
func (o *GetKubeExecCreds) HasUriSans() bool {
	if o != nil && !IsNil(o.UriSans) {
		return true
	}

	return false
}

// SetUriSans gets a reference to the given string and assigns it to the UriSans field.
func (o *GetKubeExecCreds) SetUriSans(v string) {
	o.UriSans = &v
}

func (o GetKubeExecCreds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetKubeExecCreds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AltNames) {
		toSerialize["alt-names"] = o.AltNames
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["api-version"] = o.ApiVersion
	}
	toSerialize["cert-issuer-name"] = o.CertIssuerName
	if !IsNil(o.CommonName) {
		toSerialize["common-name"] = o.CommonName
	}
	if !IsNil(o.CsrDataBase64) {
		toSerialize["csr-data-base64"] = o.CsrDataBase64
	}
	if !IsNil(o.ExtendedKeyUsage) {
		toSerialize["extended-key-usage"] = o.ExtendedKeyUsage
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.KeyDataBase64) {
		toSerialize["key-data-base64"] = o.KeyDataBase64
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	if !IsNil(o.UriSans) {
		toSerialize["uri-sans"] = o.UriSans
	}
	return toSerialize, nil
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


