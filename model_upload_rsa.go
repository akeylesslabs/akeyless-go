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

// UploadRSA struct for UploadRSA
type UploadRSA struct {
	// Key type. options: [RSA1024, RSA2048, RSA3072, RSA4096]
	Alg string `json:"alg"`
	// Path to a file that contain the certificate in a PEM format.
	Cert *string `json:"cert,omitempty"`
	// Certificate in a PEM format.
	CertFileData *string `json:"cert-file-data,omitempty"`
	// The customer fragment ID that will be used to split the key (if empty, the key will be created independently of a customer fragment)
	CustomerFrgId *string `json:"customer-frg-id,omitempty"`
	// A metadata about the key
	Metadata *string `json:"metadata,omitempty"`
	// Name of key to be created
	Name string `json:"name"`
	// RSA private key data, base64 encoded
	RsaFileData *string `json:"rsa-file-data,omitempty"`
	// RSA private key file path
	RsaKeyFilePath *string `json:"rsa-key-file-path,omitempty"`
	// The number of fragments that the item will be split into
	SplitLevel *int64 `json:"split-level,omitempty"`
	// List of the tags attached to this key
	Tag *[]string `json:"tag,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUploadRSA instantiates a new UploadRSA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadRSA(alg string, name string, ) *UploadRSA {
	this := UploadRSA{}
	this.Alg = alg
	this.Name = name
	var splitLevel int64 = 2
	this.SplitLevel = &splitLevel
	return &this
}

// NewUploadRSAWithDefaults instantiates a new UploadRSA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadRSAWithDefaults() *UploadRSA {
	this := UploadRSA{}
	var splitLevel int64 = 2
	this.SplitLevel = &splitLevel
	return &this
}

// GetAlg returns the Alg field value
func (o *UploadRSA) GetAlg() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetAlgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *UploadRSA) SetAlg(v string) {
	o.Alg = v
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *UploadRSA) GetCert() string {
	if o == nil || o.Cert == nil {
		var ret string
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetCertOk() (*string, bool) {
	if o == nil || o.Cert == nil {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *UploadRSA) HasCert() bool {
	if o != nil && o.Cert != nil {
		return true
	}

	return false
}

// SetCert gets a reference to the given string and assigns it to the Cert field.
func (o *UploadRSA) SetCert(v string) {
	o.Cert = &v
}

// GetCertFileData returns the CertFileData field value if set, zero value otherwise.
func (o *UploadRSA) GetCertFileData() string {
	if o == nil || o.CertFileData == nil {
		var ret string
		return ret
	}
	return *o.CertFileData
}

// GetCertFileDataOk returns a tuple with the CertFileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetCertFileDataOk() (*string, bool) {
	if o == nil || o.CertFileData == nil {
		return nil, false
	}
	return o.CertFileData, true
}

// HasCertFileData returns a boolean if a field has been set.
func (o *UploadRSA) HasCertFileData() bool {
	if o != nil && o.CertFileData != nil {
		return true
	}

	return false
}

// SetCertFileData gets a reference to the given string and assigns it to the CertFileData field.
func (o *UploadRSA) SetCertFileData(v string) {
	o.CertFileData = &v
}

// GetCustomerFrgId returns the CustomerFrgId field value if set, zero value otherwise.
func (o *UploadRSA) GetCustomerFrgId() string {
	if o == nil || o.CustomerFrgId == nil {
		var ret string
		return ret
	}
	return *o.CustomerFrgId
}

// GetCustomerFrgIdOk returns a tuple with the CustomerFrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetCustomerFrgIdOk() (*string, bool) {
	if o == nil || o.CustomerFrgId == nil {
		return nil, false
	}
	return o.CustomerFrgId, true
}

// HasCustomerFrgId returns a boolean if a field has been set.
func (o *UploadRSA) HasCustomerFrgId() bool {
	if o != nil && o.CustomerFrgId != nil {
		return true
	}

	return false
}

// SetCustomerFrgId gets a reference to the given string and assigns it to the CustomerFrgId field.
func (o *UploadRSA) SetCustomerFrgId(v string) {
	o.CustomerFrgId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UploadRSA) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UploadRSA) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *UploadRSA) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *UploadRSA) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UploadRSA) SetName(v string) {
	o.Name = v
}

// GetRsaFileData returns the RsaFileData field value if set, zero value otherwise.
func (o *UploadRSA) GetRsaFileData() string {
	if o == nil || o.RsaFileData == nil {
		var ret string
		return ret
	}
	return *o.RsaFileData
}

// GetRsaFileDataOk returns a tuple with the RsaFileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetRsaFileDataOk() (*string, bool) {
	if o == nil || o.RsaFileData == nil {
		return nil, false
	}
	return o.RsaFileData, true
}

// HasRsaFileData returns a boolean if a field has been set.
func (o *UploadRSA) HasRsaFileData() bool {
	if o != nil && o.RsaFileData != nil {
		return true
	}

	return false
}

// SetRsaFileData gets a reference to the given string and assigns it to the RsaFileData field.
func (o *UploadRSA) SetRsaFileData(v string) {
	o.RsaFileData = &v
}

// GetRsaKeyFilePath returns the RsaKeyFilePath field value if set, zero value otherwise.
func (o *UploadRSA) GetRsaKeyFilePath() string {
	if o == nil || o.RsaKeyFilePath == nil {
		var ret string
		return ret
	}
	return *o.RsaKeyFilePath
}

// GetRsaKeyFilePathOk returns a tuple with the RsaKeyFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetRsaKeyFilePathOk() (*string, bool) {
	if o == nil || o.RsaKeyFilePath == nil {
		return nil, false
	}
	return o.RsaKeyFilePath, true
}

// HasRsaKeyFilePath returns a boolean if a field has been set.
func (o *UploadRSA) HasRsaKeyFilePath() bool {
	if o != nil && o.RsaKeyFilePath != nil {
		return true
	}

	return false
}

// SetRsaKeyFilePath gets a reference to the given string and assigns it to the RsaKeyFilePath field.
func (o *UploadRSA) SetRsaKeyFilePath(v string) {
	o.RsaKeyFilePath = &v
}

// GetSplitLevel returns the SplitLevel field value if set, zero value otherwise.
func (o *UploadRSA) GetSplitLevel() int64 {
	if o == nil || o.SplitLevel == nil {
		var ret int64
		return ret
	}
	return *o.SplitLevel
}

// GetSplitLevelOk returns a tuple with the SplitLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetSplitLevelOk() (*int64, bool) {
	if o == nil || o.SplitLevel == nil {
		return nil, false
	}
	return o.SplitLevel, true
}

// HasSplitLevel returns a boolean if a field has been set.
func (o *UploadRSA) HasSplitLevel() bool {
	if o != nil && o.SplitLevel != nil {
		return true
	}

	return false
}

// SetSplitLevel gets a reference to the given int64 and assigns it to the SplitLevel field.
func (o *UploadRSA) SetSplitLevel(v int64) {
	o.SplitLevel = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *UploadRSA) GetTag() []string {
	if o == nil || o.Tag == nil {
		var ret []string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetTagOk() (*[]string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *UploadRSA) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *UploadRSA) SetTag(v []string) {
	o.Tag = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UploadRSA) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UploadRSA) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UploadRSA) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UploadRSA) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UploadRSA) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UploadRSA) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UploadRSA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if o.Cert != nil {
		toSerialize["cert"] = o.Cert
	}
	if o.CertFileData != nil {
		toSerialize["cert-file-data"] = o.CertFileData
	}
	if o.CustomerFrgId != nil {
		toSerialize["customer-frg-id"] = o.CustomerFrgId
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.RsaFileData != nil {
		toSerialize["rsa-file-data"] = o.RsaFileData
	}
	if o.RsaKeyFilePath != nil {
		toSerialize["rsa-key-file-path"] = o.RsaKeyFilePath
	}
	if o.SplitLevel != nil {
		toSerialize["split-level"] = o.SplitLevel
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableUploadRSA struct {
	value *UploadRSA
	isSet bool
}

func (v NullableUploadRSA) Get() *UploadRSA {
	return v.value
}

func (v *NullableUploadRSA) Set(val *UploadRSA) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadRSA) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadRSA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadRSA(val *UploadRSA) *NullableUploadRSA {
	return &NullableUploadRSA{value: val, isSet: true}
}

func (v NullableUploadRSA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadRSA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


