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

// SalesforceTargetDetails SalesforceTargetDetails
type SalesforceTargetDetails struct {
	// params needed for jwt auth AppPrivateKey is the rsa private key in PEM format
	AppPrivateKey *[]int32 `json:"app_private_key,omitempty"`
	AuthFlow *string `json:"auth_flow,omitempty"`
	// CACertData is the rsa 4096 certificate data in PEM format
	CaCertData *[]int32 `json:"ca_cert_data,omitempty"`
	// CACertName is the name of the certificate in SalesForce tenant
	CaCertName *string `json:"ca_cert_name,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	// params needed for password auth
	ClientSecret *string `json:"client_secret,omitempty"`
	Password *string `json:"password,omitempty"`
	SecurityToken *string `json:"security_token,omitempty"`
	TenantUrl *string `json:"tenant_url,omitempty"`
	UserName *string `json:"user_name,omitempty"`
}

// NewSalesforceTargetDetails instantiates a new SalesforceTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesforceTargetDetails() *SalesforceTargetDetails {
	this := SalesforceTargetDetails{}
	return &this
}

// NewSalesforceTargetDetailsWithDefaults instantiates a new SalesforceTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesforceTargetDetailsWithDefaults() *SalesforceTargetDetails {
	this := SalesforceTargetDetails{}
	return &this
}

// GetAppPrivateKey returns the AppPrivateKey field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetAppPrivateKey() []int32 {
	if o == nil || o.AppPrivateKey == nil {
		var ret []int32
		return ret
	}
	return *o.AppPrivateKey
}

// GetAppPrivateKeyOk returns a tuple with the AppPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetAppPrivateKeyOk() (*[]int32, bool) {
	if o == nil || o.AppPrivateKey == nil {
		return nil, false
	}
	return o.AppPrivateKey, true
}

// HasAppPrivateKey returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasAppPrivateKey() bool {
	if o != nil && o.AppPrivateKey != nil {
		return true
	}

	return false
}

// SetAppPrivateKey gets a reference to the given []int32 and assigns it to the AppPrivateKey field.
func (o *SalesforceTargetDetails) SetAppPrivateKey(v []int32) {
	o.AppPrivateKey = &v
}

// GetAuthFlow returns the AuthFlow field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetAuthFlow() string {
	if o == nil || o.AuthFlow == nil {
		var ret string
		return ret
	}
	return *o.AuthFlow
}

// GetAuthFlowOk returns a tuple with the AuthFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetAuthFlowOk() (*string, bool) {
	if o == nil || o.AuthFlow == nil {
		return nil, false
	}
	return o.AuthFlow, true
}

// HasAuthFlow returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasAuthFlow() bool {
	if o != nil && o.AuthFlow != nil {
		return true
	}

	return false
}

// SetAuthFlow gets a reference to the given string and assigns it to the AuthFlow field.
func (o *SalesforceTargetDetails) SetAuthFlow(v string) {
	o.AuthFlow = &v
}

// GetCaCertData returns the CaCertData field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetCaCertData() []int32 {
	if o == nil || o.CaCertData == nil {
		var ret []int32
		return ret
	}
	return *o.CaCertData
}

// GetCaCertDataOk returns a tuple with the CaCertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetCaCertDataOk() (*[]int32, bool) {
	if o == nil || o.CaCertData == nil {
		return nil, false
	}
	return o.CaCertData, true
}

// HasCaCertData returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasCaCertData() bool {
	if o != nil && o.CaCertData != nil {
		return true
	}

	return false
}

// SetCaCertData gets a reference to the given []int32 and assigns it to the CaCertData field.
func (o *SalesforceTargetDetails) SetCaCertData(v []int32) {
	o.CaCertData = &v
}

// GetCaCertName returns the CaCertName field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetCaCertName() string {
	if o == nil || o.CaCertName == nil {
		var ret string
		return ret
	}
	return *o.CaCertName
}

// GetCaCertNameOk returns a tuple with the CaCertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetCaCertNameOk() (*string, bool) {
	if o == nil || o.CaCertName == nil {
		return nil, false
	}
	return o.CaCertName, true
}

// HasCaCertName returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasCaCertName() bool {
	if o != nil && o.CaCertName != nil {
		return true
	}

	return false
}

// SetCaCertName gets a reference to the given string and assigns it to the CaCertName field.
func (o *SalesforceTargetDetails) SetCaCertName(v string) {
	o.CaCertName = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SalesforceTargetDetails) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *SalesforceTargetDetails) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SalesforceTargetDetails) SetPassword(v string) {
	o.Password = &v
}

// GetSecurityToken returns the SecurityToken field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetSecurityToken() string {
	if o == nil || o.SecurityToken == nil {
		var ret string
		return ret
	}
	return *o.SecurityToken
}

// GetSecurityTokenOk returns a tuple with the SecurityToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetSecurityTokenOk() (*string, bool) {
	if o == nil || o.SecurityToken == nil {
		return nil, false
	}
	return o.SecurityToken, true
}

// HasSecurityToken returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasSecurityToken() bool {
	if o != nil && o.SecurityToken != nil {
		return true
	}

	return false
}

// SetSecurityToken gets a reference to the given string and assigns it to the SecurityToken field.
func (o *SalesforceTargetDetails) SetSecurityToken(v string) {
	o.SecurityToken = &v
}

// GetTenantUrl returns the TenantUrl field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetTenantUrl() string {
	if o == nil || o.TenantUrl == nil {
		var ret string
		return ret
	}
	return *o.TenantUrl
}

// GetTenantUrlOk returns a tuple with the TenantUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetTenantUrlOk() (*string, bool) {
	if o == nil || o.TenantUrl == nil {
		return nil, false
	}
	return o.TenantUrl, true
}

// HasTenantUrl returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasTenantUrl() bool {
	if o != nil && o.TenantUrl != nil {
		return true
	}

	return false
}

// SetTenantUrl gets a reference to the given string and assigns it to the TenantUrl field.
func (o *SalesforceTargetDetails) SetTenantUrl(v string) {
	o.TenantUrl = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SalesforceTargetDetails) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceTargetDetails) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SalesforceTargetDetails) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SalesforceTargetDetails) SetUserName(v string) {
	o.UserName = &v
}

func (o SalesforceTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppPrivateKey != nil {
		toSerialize["app_private_key"] = o.AppPrivateKey
	}
	if o.AuthFlow != nil {
		toSerialize["auth_flow"] = o.AuthFlow
	}
	if o.CaCertData != nil {
		toSerialize["ca_cert_data"] = o.CaCertData
	}
	if o.CaCertName != nil {
		toSerialize["ca_cert_name"] = o.CaCertName
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.SecurityToken != nil {
		toSerialize["security_token"] = o.SecurityToken
	}
	if o.TenantUrl != nil {
		toSerialize["tenant_url"] = o.TenantUrl
	}
	if o.UserName != nil {
		toSerialize["user_name"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableSalesforceTargetDetails struct {
	value *SalesforceTargetDetails
	isSet bool
}

func (v NullableSalesforceTargetDetails) Get() *SalesforceTargetDetails {
	return v.value
}

func (v *NullableSalesforceTargetDetails) Set(val *SalesforceTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesforceTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesforceTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesforceTargetDetails(val *SalesforceTargetDetails) *NullableSalesforceTargetDetails {
	return &NullableSalesforceTargetDetails{value: val, isSet: true}
}

func (v NullableSalesforceTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesforceTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


