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

// CreateClassicKey CreateClassicKey is a command that creates classic key
type CreateClassicKey struct {
	// Classic Key type; options: [AES128GCM, AES256GCM, AES128SIV, AES256SIV, RSA1024, RSA2048, RSA3072, RSA4096, EC256, EC384, GPG]
	Alg string `json:"alg"`
	// Certificate in a PEM format.
	CertFileData *string `json:"cert-file-data,omitempty"`
	// Common name for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateCommonName *string `json:"certificate-common-name,omitempty"`
	// Country name for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateCountry *string `json:"certificate-country,omitempty"`
	// Digest algorithm to be used for the certificate key signing. Currently, we support only \"sha256\" so we hide this option for CLI.
	CertificateDigestAlgo *string `json:"certificate-digest-algo,omitempty"`
	// Locality for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateLocality *string `json:"certificate-locality,omitempty"`
	// Organization name for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateOrganization *string `json:"certificate-organization,omitempty"`
	// Province name for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateProvince *string `json:"certificate-province,omitempty"`
	// TTL in days for the generated certificate. Required only for generate-self-signed-certificate.
	CertificateTtl *int64 `json:"certificate-ttl,omitempty"`
	// The csr config data in base64 encoding
	ConfFileData *string `json:"conf-file-data,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Whether to generate a self signed certificate with the key. If set, --certificate-ttl must be provided.
	GenerateSelfSignedCertificate *bool `json:"generate-self-signed-certificate,omitempty"`
	// gpg alg: Relevant only if GPG key type selected; options: [RSA1024, RSA2048, RSA3072, RSA4096, Ed25519]
	GpgAlg *string `json:"gpg-alg,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Base64-encoded classic key value
	KeyData *string `json:"key-data,omitempty"`
	// Deprecated - use description
	Metadata *string `json:"metadata,omitempty"`
	// ClassicKey name
	Name string `json:"name"`
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	ProtectionKeyName *string `json:"protection-key-name,omitempty"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateClassicKey instantiates a new CreateClassicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClassicKey(alg string, name string, ) *CreateClassicKey {
	this := CreateClassicKey{}
	this.Alg = alg
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewCreateClassicKeyWithDefaults instantiates a new CreateClassicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClassicKeyWithDefaults() *CreateClassicKey {
	this := CreateClassicKey{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAlg returns the Alg field value
func (o *CreateClassicKey) GetAlg() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetAlgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *CreateClassicKey) SetAlg(v string) {
	o.Alg = v
}

// GetCertFileData returns the CertFileData field value if set, zero value otherwise.
func (o *CreateClassicKey) GetCertFileData() string {
	if o == nil || o.CertFileData == nil {
		var ret string
		return ret
	}
	return *o.CertFileData
}

// GetCertFileDataOk returns a tuple with the CertFileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetCertFileDataOk() (*string, bool) {
	if o == nil || o.CertFileData == nil {
		return nil, false
	}
	return o.CertFileData, true
}

// HasCertFileData returns a boolean if a field has been set.
func (o *CreateClassicKey) HasCertFileData() bool {
	if o != nil && o.CertFileData != nil {
		return true
	}

	return false
}

// SetCertFileData gets a reference to the given string and assigns it to the CertFileData field.
func (o *CreateClassicKey) SetCertFileData(v string) {
	o.CertFileData = &v
}

// GetCertificateCommonName returns the CertificateCommonName field value if set, zero value otherwise.
func (o *CreateClassicKey) GetCertificateCommonName() string {
	if o == nil || o.CertificateCommonName == nil {
		var ret string
		return ret
	}
	return *o.CertificateCommonName
}

// GetCertificateCommonNameOk returns a tuple with the CertificateCommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetCertificateCommonNameOk() (*string, bool) {
	if o == nil || o.CertificateCommonName == nil {
		return nil, false
	}
	return o.CertificateCommonName, true
}

// HasCertificateCommonName returns a boolean if a field has been set.
func (o *CreateClassicKey) HasCertificateCommonName() bool {
	if o != nil && o.CertificateCommonName != nil {
		return true
	}

	return false
}

// SetCertificateCommonName gets a reference to the given string and assigns it to the CertificateCommonName field.
func (o *CreateClassicKey) SetCertificateCommonName(v string) {
	o.CertificateCommonName = &v
}

// GetCertificateCountry returns the CertificateCountry field value if set, zero value otherwise.
func (o *CreateClassicKey) GetCertificateCountry() string {
	if o == nil || o.CertificateCountry == nil {
		var ret string
		return ret
	}
	return *o.CertificateCountry
}

// GetCertificateCountryOk returns a tuple with the CertificateCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetCertificateCountryOk() (*string, bool) {
	if o == nil || o.CertificateCountry == nil {
		return nil, false
	}
	return o.CertificateCountry, true
}

// HasCertificateCountry returns a boolean if a field has been set.
func (o *CreateClassicKey) HasCertificateCountry() bool {
	if o != nil && o.CertificateCountry != nil {
		return true
	}

	return false
}

// SetCertificateCountry gets a reference to the given string and assigns it to the CertificateCountry field.
func (o *CreateClassicKey) SetCertificateCountry(v string) {
	o.CertificateCountry = &v
}

// GetCertificateDigestAlgo returns the CertificateDigestAlgo field value if set, zero value otherwise.
func (o *CreateClassicKey) GetCertificateDigestAlgo() string {
	if o == nil || o.CertificateDigestAlgo == nil {
		var ret string
		return ret
	}
	return *o.CertificateDigestAlgo
}

// GetCertificateDigestAlgoOk returns a tuple with the CertificateDigestAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetCertificateDigestAlgoOk() (*string, bool) {
	if o == nil || o.CertificateDigestAlgo == nil {
		return nil, false
	}
	return o.CertificateDigestAlgo, true
}

// HasCertificateDigestAlgo returns a boolean if a field has been set.
func (o *CreateClassicKey) HasCertificateDigestAlgo() bool {
	if o != nil && o.CertificateDigestAlgo != nil {
		return true
	}

	return false
}

// SetCertificateDigestAlgo gets a reference to the given string and assigns it to the CertificateDigestAlgo field.
func (o *CreateClassicKey) SetCertificateDigestAlgo(v string) {
	o.CertificateDigestAlgo = &v
}

// GetCertificateLocality returns the CertificateLocality field value if set, zero value otherwise.
func (o *CreateClassicKey) GetCertificateLocality() string {
	if o == nil || o.CertificateLocality == nil {
		var ret string
		return ret
	}
	return *o.CertificateLocality
}

// GetCertificateLocalityOk returns a tuple with the CertificateLocality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetCertificateLocalityOk() (*string, bool) {
	if o == nil || o.CertificateLocality == nil {
		return nil, false
	}
	return o.CertificateLocality, true
}

// HasCertificateLocality returns a boolean if a field has been set.
func (o *CreateClassicKey) HasCertificateLocality() bool {
	if o != nil && o.CertificateLocality != nil {
		return true
	}

	return false
}

// SetCertificateLocality gets a reference to the given string and assigns it to the CertificateLocality field.
func (o *CreateClassicKey) SetCertificateLocality(v string) {
	o.CertificateLocality = &v
}

// GetCertificateOrganization returns the CertificateOrganization field value if set, zero value otherwise.
func (o *CreateClassicKey) GetCertificateOrganization() string {
	if o == nil || o.CertificateOrganization == nil {
		var ret string
		return ret
	}
	return *o.CertificateOrganization
}

// GetCertificateOrganizationOk returns a tuple with the CertificateOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetCertificateOrganizationOk() (*string, bool) {
	if o == nil || o.CertificateOrganization == nil {
		return nil, false
	}
	return o.CertificateOrganization, true
}

// HasCertificateOrganization returns a boolean if a field has been set.
func (o *CreateClassicKey) HasCertificateOrganization() bool {
	if o != nil && o.CertificateOrganization != nil {
		return true
	}

	return false
}

// SetCertificateOrganization gets a reference to the given string and assigns it to the CertificateOrganization field.
func (o *CreateClassicKey) SetCertificateOrganization(v string) {
	o.CertificateOrganization = &v
}

// GetCertificateProvince returns the CertificateProvince field value if set, zero value otherwise.
func (o *CreateClassicKey) GetCertificateProvince() string {
	if o == nil || o.CertificateProvince == nil {
		var ret string
		return ret
	}
	return *o.CertificateProvince
}

// GetCertificateProvinceOk returns a tuple with the CertificateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetCertificateProvinceOk() (*string, bool) {
	if o == nil || o.CertificateProvince == nil {
		return nil, false
	}
	return o.CertificateProvince, true
}

// HasCertificateProvince returns a boolean if a field has been set.
func (o *CreateClassicKey) HasCertificateProvince() bool {
	if o != nil && o.CertificateProvince != nil {
		return true
	}

	return false
}

// SetCertificateProvince gets a reference to the given string and assigns it to the CertificateProvince field.
func (o *CreateClassicKey) SetCertificateProvince(v string) {
	o.CertificateProvince = &v
}

// GetCertificateTtl returns the CertificateTtl field value if set, zero value otherwise.
func (o *CreateClassicKey) GetCertificateTtl() int64 {
	if o == nil || o.CertificateTtl == nil {
		var ret int64
		return ret
	}
	return *o.CertificateTtl
}

// GetCertificateTtlOk returns a tuple with the CertificateTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetCertificateTtlOk() (*int64, bool) {
	if o == nil || o.CertificateTtl == nil {
		return nil, false
	}
	return o.CertificateTtl, true
}

// HasCertificateTtl returns a boolean if a field has been set.
func (o *CreateClassicKey) HasCertificateTtl() bool {
	if o != nil && o.CertificateTtl != nil {
		return true
	}

	return false
}

// SetCertificateTtl gets a reference to the given int64 and assigns it to the CertificateTtl field.
func (o *CreateClassicKey) SetCertificateTtl(v int64) {
	o.CertificateTtl = &v
}

// GetConfFileData returns the ConfFileData field value if set, zero value otherwise.
func (o *CreateClassicKey) GetConfFileData() string {
	if o == nil || o.ConfFileData == nil {
		var ret string
		return ret
	}
	return *o.ConfFileData
}

// GetConfFileDataOk returns a tuple with the ConfFileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetConfFileDataOk() (*string, bool) {
	if o == nil || o.ConfFileData == nil {
		return nil, false
	}
	return o.ConfFileData, true
}

// HasConfFileData returns a boolean if a field has been set.
func (o *CreateClassicKey) HasConfFileData() bool {
	if o != nil && o.ConfFileData != nil {
		return true
	}

	return false
}

// SetConfFileData gets a reference to the given string and assigns it to the ConfFileData field.
func (o *CreateClassicKey) SetConfFileData(v string) {
	o.ConfFileData = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *CreateClassicKey) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *CreateClassicKey) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *CreateClassicKey) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateClassicKey) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateClassicKey) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateClassicKey) SetDescription(v string) {
	o.Description = &v
}

// GetGenerateSelfSignedCertificate returns the GenerateSelfSignedCertificate field value if set, zero value otherwise.
func (o *CreateClassicKey) GetGenerateSelfSignedCertificate() bool {
	if o == nil || o.GenerateSelfSignedCertificate == nil {
		var ret bool
		return ret
	}
	return *o.GenerateSelfSignedCertificate
}

// GetGenerateSelfSignedCertificateOk returns a tuple with the GenerateSelfSignedCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetGenerateSelfSignedCertificateOk() (*bool, bool) {
	if o == nil || o.GenerateSelfSignedCertificate == nil {
		return nil, false
	}
	return o.GenerateSelfSignedCertificate, true
}

// HasGenerateSelfSignedCertificate returns a boolean if a field has been set.
func (o *CreateClassicKey) HasGenerateSelfSignedCertificate() bool {
	if o != nil && o.GenerateSelfSignedCertificate != nil {
		return true
	}

	return false
}

// SetGenerateSelfSignedCertificate gets a reference to the given bool and assigns it to the GenerateSelfSignedCertificate field.
func (o *CreateClassicKey) SetGenerateSelfSignedCertificate(v bool) {
	o.GenerateSelfSignedCertificate = &v
}

// GetGpgAlg returns the GpgAlg field value if set, zero value otherwise.
func (o *CreateClassicKey) GetGpgAlg() string {
	if o == nil || o.GpgAlg == nil {
		var ret string
		return ret
	}
	return *o.GpgAlg
}

// GetGpgAlgOk returns a tuple with the GpgAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetGpgAlgOk() (*string, bool) {
	if o == nil || o.GpgAlg == nil {
		return nil, false
	}
	return o.GpgAlg, true
}

// HasGpgAlg returns a boolean if a field has been set.
func (o *CreateClassicKey) HasGpgAlg() bool {
	if o != nil && o.GpgAlg != nil {
		return true
	}

	return false
}

// SetGpgAlg gets a reference to the given string and assigns it to the GpgAlg field.
func (o *CreateClassicKey) SetGpgAlg(v string) {
	o.GpgAlg = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateClassicKey) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateClassicKey) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateClassicKey) SetJson(v bool) {
	o.Json = &v
}

// GetKeyData returns the KeyData field value if set, zero value otherwise.
func (o *CreateClassicKey) GetKeyData() string {
	if o == nil || o.KeyData == nil {
		var ret string
		return ret
	}
	return *o.KeyData
}

// GetKeyDataOk returns a tuple with the KeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetKeyDataOk() (*string, bool) {
	if o == nil || o.KeyData == nil {
		return nil, false
	}
	return o.KeyData, true
}

// HasKeyData returns a boolean if a field has been set.
func (o *CreateClassicKey) HasKeyData() bool {
	if o != nil && o.KeyData != nil {
		return true
	}

	return false
}

// SetKeyData gets a reference to the given string and assigns it to the KeyData field.
func (o *CreateClassicKey) SetKeyData(v string) {
	o.KeyData = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateClassicKey) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateClassicKey) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreateClassicKey) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *CreateClassicKey) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateClassicKey) SetName(v string) {
	o.Name = v
}

// GetProtectionKeyName returns the ProtectionKeyName field value if set, zero value otherwise.
func (o *CreateClassicKey) GetProtectionKeyName() string {
	if o == nil || o.ProtectionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKeyName
}

// GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetProtectionKeyNameOk() (*string, bool) {
	if o == nil || o.ProtectionKeyName == nil {
		return nil, false
	}
	return o.ProtectionKeyName, true
}

// HasProtectionKeyName returns a boolean if a field has been set.
func (o *CreateClassicKey) HasProtectionKeyName() bool {
	if o != nil && o.ProtectionKeyName != nil {
		return true
	}

	return false
}

// SetProtectionKeyName gets a reference to the given string and assigns it to the ProtectionKeyName field.
func (o *CreateClassicKey) SetProtectionKeyName(v string) {
	o.ProtectionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateClassicKey) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateClassicKey) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateClassicKey) SetTags(v []string) {
	o.Tags = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateClassicKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateClassicKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateClassicKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateClassicKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateClassicKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateClassicKey) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateClassicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if o.CertFileData != nil {
		toSerialize["cert-file-data"] = o.CertFileData
	}
	if o.CertificateCommonName != nil {
		toSerialize["certificate-common-name"] = o.CertificateCommonName
	}
	if o.CertificateCountry != nil {
		toSerialize["certificate-country"] = o.CertificateCountry
	}
	if o.CertificateDigestAlgo != nil {
		toSerialize["certificate-digest-algo"] = o.CertificateDigestAlgo
	}
	if o.CertificateLocality != nil {
		toSerialize["certificate-locality"] = o.CertificateLocality
	}
	if o.CertificateOrganization != nil {
		toSerialize["certificate-organization"] = o.CertificateOrganization
	}
	if o.CertificateProvince != nil {
		toSerialize["certificate-province"] = o.CertificateProvince
	}
	if o.CertificateTtl != nil {
		toSerialize["certificate-ttl"] = o.CertificateTtl
	}
	if o.ConfFileData != nil {
		toSerialize["conf-file-data"] = o.ConfFileData
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GenerateSelfSignedCertificate != nil {
		toSerialize["generate-self-signed-certificate"] = o.GenerateSelfSignedCertificate
	}
	if o.GpgAlg != nil {
		toSerialize["gpg-alg"] = o.GpgAlg
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.KeyData != nil {
		toSerialize["key-data"] = o.KeyData
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProtectionKeyName != nil {
		toSerialize["protection-key-name"] = o.ProtectionKeyName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClassicKey struct {
	value *CreateClassicKey
	isSet bool
}

func (v NullableCreateClassicKey) Get() *CreateClassicKey {
	return v.value
}

func (v *NullableCreateClassicKey) Set(val *CreateClassicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClassicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClassicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClassicKey(val *CreateClassicKey) *NullableCreateClassicKey {
	return &NullableCreateClassicKey{value: val, isSet: true}
}

func (v NullableCreateClassicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClassicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


