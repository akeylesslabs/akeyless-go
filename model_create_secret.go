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

// CreateSecret struct for CreateSecret
type CreateSecret struct {
	// Protection from accidental deletion of this item
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Metadata about the secret
	Metadata *string `json:"metadata,omitempty"`
	// The provided value is a multiline value (separated by '\\n')
	MultilineValue *bool `json:"multiline_value,omitempty"`
	// Secret name
	Name string `json:"name"`
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	ProtectionKey *string `json:"protection_key,omitempty"`
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	SecureAccessSshCreds *string `json:"secure-access-ssh-creds,omitempty"`
	SecureAccessSshUser *string `json:"secure-access-ssh-user,omitempty"`
	SecureAccessUrl *string `json:"secure-access-url,omitempty"`
	SecureAccessWebBrowsing *bool `json:"secure-access-web-browsing,omitempty"`
	SecureAccessWebProxy *bool `json:"secure-access-web-proxy,omitempty"`
	// List of the tags attached to this secret
	Tags *[]string `json:"tags,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The secret value
	Value string `json:"value"`
}

// NewCreateSecret instantiates a new CreateSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecret(name string, value string, ) *CreateSecret {
	this := CreateSecret{}
	this.Name = name
	this.Value = value
	return &this
}

// NewCreateSecretWithDefaults instantiates a new CreateSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecretWithDefaults() *CreateSecret {
	this := CreateSecret{}
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *CreateSecret) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *CreateSecret) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *CreateSecret) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateSecret) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateSecret) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreateSecret) SetMetadata(v string) {
	o.Metadata = &v
}

// GetMultilineValue returns the MultilineValue field value if set, zero value otherwise.
func (o *CreateSecret) GetMultilineValue() bool {
	if o == nil || o.MultilineValue == nil {
		var ret bool
		return ret
	}
	return *o.MultilineValue
}

// GetMultilineValueOk returns a tuple with the MultilineValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetMultilineValueOk() (*bool, bool) {
	if o == nil || o.MultilineValue == nil {
		return nil, false
	}
	return o.MultilineValue, true
}

// HasMultilineValue returns a boolean if a field has been set.
func (o *CreateSecret) HasMultilineValue() bool {
	if o != nil && o.MultilineValue != nil {
		return true
	}

	return false
}

// SetMultilineValue gets a reference to the given bool and assigns it to the MultilineValue field.
func (o *CreateSecret) SetMultilineValue(v bool) {
	o.MultilineValue = &v
}

// GetName returns the Name field value
func (o *CreateSecret) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSecret) SetName(v string) {
	o.Name = v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *CreateSecret) GetProtectionKey() string {
	if o == nil || o.ProtectionKey == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetProtectionKeyOk() (*string, bool) {
	if o == nil || o.ProtectionKey == nil {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *CreateSecret) HasProtectionKey() bool {
	if o != nil && o.ProtectionKey != nil {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *CreateSecret) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *CreateSecret) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *CreateSecret) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *CreateSecret) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *CreateSecret) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *CreateSecret) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *CreateSecret) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *CreateSecret) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *CreateSecret) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *CreateSecret) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessSshCreds returns the SecureAccessSshCreds field value if set, zero value otherwise.
func (o *CreateSecret) GetSecureAccessSshCreds() string {
	if o == nil || o.SecureAccessSshCreds == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessSshCreds
}

// GetSecureAccessSshCredsOk returns a tuple with the SecureAccessSshCreds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetSecureAccessSshCredsOk() (*string, bool) {
	if o == nil || o.SecureAccessSshCreds == nil {
		return nil, false
	}
	return o.SecureAccessSshCreds, true
}

// HasSecureAccessSshCreds returns a boolean if a field has been set.
func (o *CreateSecret) HasSecureAccessSshCreds() bool {
	if o != nil && o.SecureAccessSshCreds != nil {
		return true
	}

	return false
}

// SetSecureAccessSshCreds gets a reference to the given string and assigns it to the SecureAccessSshCreds field.
func (o *CreateSecret) SetSecureAccessSshCreds(v string) {
	o.SecureAccessSshCreds = &v
}

// GetSecureAccessSshUser returns the SecureAccessSshUser field value if set, zero value otherwise.
func (o *CreateSecret) GetSecureAccessSshUser() string {
	if o == nil || o.SecureAccessSshUser == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessSshUser
}

// GetSecureAccessSshUserOk returns a tuple with the SecureAccessSshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetSecureAccessSshUserOk() (*string, bool) {
	if o == nil || o.SecureAccessSshUser == nil {
		return nil, false
	}
	return o.SecureAccessSshUser, true
}

// HasSecureAccessSshUser returns a boolean if a field has been set.
func (o *CreateSecret) HasSecureAccessSshUser() bool {
	if o != nil && o.SecureAccessSshUser != nil {
		return true
	}

	return false
}

// SetSecureAccessSshUser gets a reference to the given string and assigns it to the SecureAccessSshUser field.
func (o *CreateSecret) SetSecureAccessSshUser(v string) {
	o.SecureAccessSshUser = &v
}

// GetSecureAccessUrl returns the SecureAccessUrl field value if set, zero value otherwise.
func (o *CreateSecret) GetSecureAccessUrl() string {
	if o == nil || o.SecureAccessUrl == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessUrl
}

// GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetSecureAccessUrlOk() (*string, bool) {
	if o == nil || o.SecureAccessUrl == nil {
		return nil, false
	}
	return o.SecureAccessUrl, true
}

// HasSecureAccessUrl returns a boolean if a field has been set.
func (o *CreateSecret) HasSecureAccessUrl() bool {
	if o != nil && o.SecureAccessUrl != nil {
		return true
	}

	return false
}

// SetSecureAccessUrl gets a reference to the given string and assigns it to the SecureAccessUrl field.
func (o *CreateSecret) SetSecureAccessUrl(v string) {
	o.SecureAccessUrl = &v
}

// GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field value if set, zero value otherwise.
func (o *CreateSecret) GetSecureAccessWebBrowsing() bool {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebBrowsing
}

// GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetSecureAccessWebBrowsingOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		return nil, false
	}
	return o.SecureAccessWebBrowsing, true
}

// HasSecureAccessWebBrowsing returns a boolean if a field has been set.
func (o *CreateSecret) HasSecureAccessWebBrowsing() bool {
	if o != nil && o.SecureAccessWebBrowsing != nil {
		return true
	}

	return false
}

// SetSecureAccessWebBrowsing gets a reference to the given bool and assigns it to the SecureAccessWebBrowsing field.
func (o *CreateSecret) SetSecureAccessWebBrowsing(v bool) {
	o.SecureAccessWebBrowsing = &v
}

// GetSecureAccessWebProxy returns the SecureAccessWebProxy field value if set, zero value otherwise.
func (o *CreateSecret) GetSecureAccessWebProxy() bool {
	if o == nil || o.SecureAccessWebProxy == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebProxy
}

// GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetSecureAccessWebProxyOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebProxy == nil {
		return nil, false
	}
	return o.SecureAccessWebProxy, true
}

// HasSecureAccessWebProxy returns a boolean if a field has been set.
func (o *CreateSecret) HasSecureAccessWebProxy() bool {
	if o != nil && o.SecureAccessWebProxy != nil {
		return true
	}

	return false
}

// SetSecureAccessWebProxy gets a reference to the given bool and assigns it to the SecureAccessWebProxy field.
func (o *CreateSecret) SetSecureAccessWebProxy(v bool) {
	o.SecureAccessWebProxy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateSecret) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateSecret) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateSecret) SetTags(v []string) {
	o.Tags = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateSecret) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateSecret) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateSecret) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateSecret) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateSecret) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateSecret) SetUidToken(v string) {
	o.UidToken = &v
}

// GetValue returns the Value field value
func (o *CreateSecret) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreateSecret) SetValue(v string) {
	o.Value = v
}

func (o CreateSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MultilineValue != nil {
		toSerialize["multiline_value"] = o.MultilineValue
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProtectionKey != nil {
		toSerialize["protection_key"] = o.ProtectionKey
	}
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessHost != nil {
		toSerialize["secure-access-host"] = o.SecureAccessHost
	}
	if o.SecureAccessSshCreds != nil {
		toSerialize["secure-access-ssh-creds"] = o.SecureAccessSshCreds
	}
	if o.SecureAccessSshUser != nil {
		toSerialize["secure-access-ssh-user"] = o.SecureAccessSshUser
	}
	if o.SecureAccessUrl != nil {
		toSerialize["secure-access-url"] = o.SecureAccessUrl
	}
	if o.SecureAccessWebBrowsing != nil {
		toSerialize["secure-access-web-browsing"] = o.SecureAccessWebBrowsing
	}
	if o.SecureAccessWebProxy != nil {
		toSerialize["secure-access-web-proxy"] = o.SecureAccessWebProxy
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
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSecret struct {
	value *CreateSecret
	isSet bool
}

func (v NullableCreateSecret) Get() *CreateSecret {
	return v.value
}

func (v *NullableCreateSecret) Set(val *CreateSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecret(val *CreateSecret) *NullableCreateSecret {
	return &NullableCreateSecret{value: val, isSet: true}
}

func (v NullableCreateSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


