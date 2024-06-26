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
	// Certificate in a PEM format.
	CertFileData *string `json:"cert-file-data,omitempty"`
	// The customer fragment ID that will be used to split the key (if empty, the key will be created independently of a customer fragment)
	CustomerFrgId *string `json:"customer-frg-id,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Deprecated - use description
	Metadata *string `json:"metadata,omitempty"`
	// Name of key to be created
	Name string `json:"name"`
	// When the overwrite flag is set, this command will only update an existing key [true/false]
	Overwrite *string `json:"overwrite,omitempty"`
	// RSA private key data, base64 encoded
	RsaFileData *string `json:"rsa-file-data,omitempty"`
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
	var json bool = false
	this.Json = &json
	this.Name = name
	var overwrite string = "false"
	this.Overwrite = &overwrite
	var splitLevel int64 = 3
	this.SplitLevel = &splitLevel
	return &this
}

// NewUploadRSAWithDefaults instantiates a new UploadRSA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadRSAWithDefaults() *UploadRSA {
	this := UploadRSA{}
	var json bool = false
	this.Json = &json
	var overwrite string = "false"
	this.Overwrite = &overwrite
	var splitLevel int64 = 3
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

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *UploadRSA) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *UploadRSA) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *UploadRSA) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UploadRSA) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UploadRSA) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UploadRSA) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UploadRSA) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UploadRSA) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UploadRSA) SetJson(v bool) {
	o.Json = &v
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

// GetOverwrite returns the Overwrite field value if set, zero value otherwise.
func (o *UploadRSA) GetOverwrite() string {
	if o == nil || o.Overwrite == nil {
		var ret string
		return ret
	}
	return *o.Overwrite
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRSA) GetOverwriteOk() (*string, bool) {
	if o == nil || o.Overwrite == nil {
		return nil, false
	}
	return o.Overwrite, true
}

// HasOverwrite returns a boolean if a field has been set.
func (o *UploadRSA) HasOverwrite() bool {
	if o != nil && o.Overwrite != nil {
		return true
	}

	return false
}

// SetOverwrite gets a reference to the given string and assigns it to the Overwrite field.
func (o *UploadRSA) SetOverwrite(v string) {
	o.Overwrite = &v
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
	if o.CertFileData != nil {
		toSerialize["cert-file-data"] = o.CertFileData
	}
	if o.CustomerFrgId != nil {
		toSerialize["customer-frg-id"] = o.CustomerFrgId
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Overwrite != nil {
		toSerialize["overwrite"] = o.Overwrite
	}
	if o.RsaFileData != nil {
		toSerialize["rsa-file-data"] = o.RsaFileData
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


