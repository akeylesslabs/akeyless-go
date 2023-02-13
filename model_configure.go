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

// Configure struct for Configure
type Configure struct {
	// Access ID
	AccessId *string `json:"access-id,omitempty"`
	// Access Key
	AccessKey *string `json:"access-key,omitempty"`
	// Access Type (access_key/password/azure_ad/saml/oidc/aws_iam/gcp/k8s)
	AccessType *string `json:"access-type,omitempty"`
	// Email (relevant only for access-type=password)
	AdminEmail *string `json:"admin-email,omitempty"`
	// Password (relevant only for access-type=password)
	AdminPassword *string `json:"admin-password,omitempty"`
	// Azure Active Directory ObjectId (relevant only for access-type=azure_ad)
	AzureAdObjectId *string `json:"azure_ad_object_id,omitempty"`
	// Certificate data encoded in base64. Used if file was not provided. (relevant only for access-type=cert in Curl Context)
	CertData *string `json:"cert-data,omitempty"`
	// GCP JWT audience
	GcpAudience *string `json:"gcp-audience,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The K8S Auth config name (relevant only for access-type=k8s)
	K8sAuthConfigName *string `json:"k8s-auth-config-name,omitempty"`
	// Private key data encoded in base64. Used if file was not provided.(relevant only for access-type=cert in Curl Context)
	KeyData *string `json:"key-data,omitempty"`
}

// NewConfigure instantiates a new Configure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigure() *Configure {
	this := Configure{}
	var accessType string = "access_key"
	this.AccessType = &accessType
	var gcpAudience string = "akeyless.io"
	this.GcpAudience = &gcpAudience
	var json bool = false
	this.Json = &json
	return &this
}

// NewConfigureWithDefaults instantiates a new Configure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigureWithDefaults() *Configure {
	this := Configure{}
	var accessType string = "access_key"
	this.AccessType = &accessType
	var gcpAudience string = "akeyless.io"
	this.GcpAudience = &gcpAudience
	var json bool = false
	this.Json = &json
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *Configure) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *Configure) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *Configure) SetAccessId(v string) {
	o.AccessId = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *Configure) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *Configure) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *Configure) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *Configure) GetAccessType() string {
	if o == nil || o.AccessType == nil {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetAccessTypeOk() (*string, bool) {
	if o == nil || o.AccessType == nil {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *Configure) HasAccessType() bool {
	if o != nil && o.AccessType != nil {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *Configure) SetAccessType(v string) {
	o.AccessType = &v
}

// GetAdminEmail returns the AdminEmail field value if set, zero value otherwise.
func (o *Configure) GetAdminEmail() string {
	if o == nil || o.AdminEmail == nil {
		var ret string
		return ret
	}
	return *o.AdminEmail
}

// GetAdminEmailOk returns a tuple with the AdminEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetAdminEmailOk() (*string, bool) {
	if o == nil || o.AdminEmail == nil {
		return nil, false
	}
	return o.AdminEmail, true
}

// HasAdminEmail returns a boolean if a field has been set.
func (o *Configure) HasAdminEmail() bool {
	if o != nil && o.AdminEmail != nil {
		return true
	}

	return false
}

// SetAdminEmail gets a reference to the given string and assigns it to the AdminEmail field.
func (o *Configure) SetAdminEmail(v string) {
	o.AdminEmail = &v
}

// GetAdminPassword returns the AdminPassword field value if set, zero value otherwise.
func (o *Configure) GetAdminPassword() string {
	if o == nil || o.AdminPassword == nil {
		var ret string
		return ret
	}
	return *o.AdminPassword
}

// GetAdminPasswordOk returns a tuple with the AdminPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetAdminPasswordOk() (*string, bool) {
	if o == nil || o.AdminPassword == nil {
		return nil, false
	}
	return o.AdminPassword, true
}

// HasAdminPassword returns a boolean if a field has been set.
func (o *Configure) HasAdminPassword() bool {
	if o != nil && o.AdminPassword != nil {
		return true
	}

	return false
}

// SetAdminPassword gets a reference to the given string and assigns it to the AdminPassword field.
func (o *Configure) SetAdminPassword(v string) {
	o.AdminPassword = &v
}

// GetAzureAdObjectId returns the AzureAdObjectId field value if set, zero value otherwise.
func (o *Configure) GetAzureAdObjectId() string {
	if o == nil || o.AzureAdObjectId == nil {
		var ret string
		return ret
	}
	return *o.AzureAdObjectId
}

// GetAzureAdObjectIdOk returns a tuple with the AzureAdObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetAzureAdObjectIdOk() (*string, bool) {
	if o == nil || o.AzureAdObjectId == nil {
		return nil, false
	}
	return o.AzureAdObjectId, true
}

// HasAzureAdObjectId returns a boolean if a field has been set.
func (o *Configure) HasAzureAdObjectId() bool {
	if o != nil && o.AzureAdObjectId != nil {
		return true
	}

	return false
}

// SetAzureAdObjectId gets a reference to the given string and assigns it to the AzureAdObjectId field.
func (o *Configure) SetAzureAdObjectId(v string) {
	o.AzureAdObjectId = &v
}

// GetCertData returns the CertData field value if set, zero value otherwise.
func (o *Configure) GetCertData() string {
	if o == nil || o.CertData == nil {
		var ret string
		return ret
	}
	return *o.CertData
}

// GetCertDataOk returns a tuple with the CertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetCertDataOk() (*string, bool) {
	if o == nil || o.CertData == nil {
		return nil, false
	}
	return o.CertData, true
}

// HasCertData returns a boolean if a field has been set.
func (o *Configure) HasCertData() bool {
	if o != nil && o.CertData != nil {
		return true
	}

	return false
}

// SetCertData gets a reference to the given string and assigns it to the CertData field.
func (o *Configure) SetCertData(v string) {
	o.CertData = &v
}

// GetGcpAudience returns the GcpAudience field value if set, zero value otherwise.
func (o *Configure) GetGcpAudience() string {
	if o == nil || o.GcpAudience == nil {
		var ret string
		return ret
	}
	return *o.GcpAudience
}

// GetGcpAudienceOk returns a tuple with the GcpAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetGcpAudienceOk() (*string, bool) {
	if o == nil || o.GcpAudience == nil {
		return nil, false
	}
	return o.GcpAudience, true
}

// HasGcpAudience returns a boolean if a field has been set.
func (o *Configure) HasGcpAudience() bool {
	if o != nil && o.GcpAudience != nil {
		return true
	}

	return false
}

// SetGcpAudience gets a reference to the given string and assigns it to the GcpAudience field.
func (o *Configure) SetGcpAudience(v string) {
	o.GcpAudience = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *Configure) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *Configure) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *Configure) SetJson(v bool) {
	o.Json = &v
}

// GetK8sAuthConfigName returns the K8sAuthConfigName field value if set, zero value otherwise.
func (o *Configure) GetK8sAuthConfigName() string {
	if o == nil || o.K8sAuthConfigName == nil {
		var ret string
		return ret
	}
	return *o.K8sAuthConfigName
}

// GetK8sAuthConfigNameOk returns a tuple with the K8sAuthConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetK8sAuthConfigNameOk() (*string, bool) {
	if o == nil || o.K8sAuthConfigName == nil {
		return nil, false
	}
	return o.K8sAuthConfigName, true
}

// HasK8sAuthConfigName returns a boolean if a field has been set.
func (o *Configure) HasK8sAuthConfigName() bool {
	if o != nil && o.K8sAuthConfigName != nil {
		return true
	}

	return false
}

// SetK8sAuthConfigName gets a reference to the given string and assigns it to the K8sAuthConfigName field.
func (o *Configure) SetK8sAuthConfigName(v string) {
	o.K8sAuthConfigName = &v
}

// GetKeyData returns the KeyData field value if set, zero value otherwise.
func (o *Configure) GetKeyData() string {
	if o == nil || o.KeyData == nil {
		var ret string
		return ret
	}
	return *o.KeyData
}

// GetKeyDataOk returns a tuple with the KeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configure) GetKeyDataOk() (*string, bool) {
	if o == nil || o.KeyData == nil {
		return nil, false
	}
	return o.KeyData, true
}

// HasKeyData returns a boolean if a field has been set.
func (o *Configure) HasKeyData() bool {
	if o != nil && o.KeyData != nil {
		return true
	}

	return false
}

// SetKeyData gets a reference to the given string and assigns it to the KeyData field.
func (o *Configure) SetKeyData(v string) {
	o.KeyData = &v
}

func (o Configure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access-id"] = o.AccessId
	}
	if o.AccessKey != nil {
		toSerialize["access-key"] = o.AccessKey
	}
	if o.AccessType != nil {
		toSerialize["access-type"] = o.AccessType
	}
	if o.AdminEmail != nil {
		toSerialize["admin-email"] = o.AdminEmail
	}
	if o.AdminPassword != nil {
		toSerialize["admin-password"] = o.AdminPassword
	}
	if o.AzureAdObjectId != nil {
		toSerialize["azure_ad_object_id"] = o.AzureAdObjectId
	}
	if o.CertData != nil {
		toSerialize["cert-data"] = o.CertData
	}
	if o.GcpAudience != nil {
		toSerialize["gcp-audience"] = o.GcpAudience
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.K8sAuthConfigName != nil {
		toSerialize["k8s-auth-config-name"] = o.K8sAuthConfigName
	}
	if o.KeyData != nil {
		toSerialize["key-data"] = o.KeyData
	}
	return json.Marshal(toSerialize)
}

type NullableConfigure struct {
	value *Configure
	isSet bool
}

func (v NullableConfigure) Get() *Configure {
	return v.value
}

func (v *NullableConfigure) Set(val *Configure) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigure) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigure(val *Configure) *NullableConfigure {
	return &NullableConfigure{value: val, isSet: true}
}

func (v NullableConfigure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


