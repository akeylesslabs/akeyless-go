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

// CreateAuthMethodHuawei createAuthMethodHuawei is a command that creates a new auth method that will be able to authenticate using Huawei credentials.
type CreateAuthMethodHuawei struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// sts URL
	AuthUrl *string `json:"auth-url,omitempty"`
	// A list of domain IDs that the access is restricted to
	BoundDomainId *[]string `json:"bound-domain-id,omitempty"`
	// A list of domain names that the access is restricted to
	BoundDomainName *[]string `json:"bound-domain-name,omitempty"`
	// A CIDR whitelist of the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// A list of full tenant ids that the access is restricted to
	BoundTenantId *[]string `json:"bound-tenant-id,omitempty"`
	// A list of full tenant names that the access is restricted to
	BoundTenantName *[]string `json:"bound-tenant-name,omitempty"`
	// A list of full user ids that the access is restricted to
	BoundUserId *[]string `json:"bound-user-id,omitempty"`
	// A list of full user-name that the access is restricted to
	BoundUserName *[]string `json:"bound-user-name,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateAuthMethodHuawei instantiates a new CreateAuthMethodHuawei object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodHuawei(name string, ) *CreateAuthMethodHuawei {
	this := CreateAuthMethodHuawei{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var authUrl string = "https://iam.myhwclouds.com:443/v3"
	this.AuthUrl = &authUrl
	this.Name = name
	return &this
}

// NewCreateAuthMethodHuaweiWithDefaults instantiates a new CreateAuthMethodHuawei object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodHuaweiWithDefaults() *CreateAuthMethodHuawei {
	this := CreateAuthMethodHuawei{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var authUrl string = "https://iam.myhwclouds.com:443/v3"
	this.AuthUrl = &authUrl
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *CreateAuthMethodHuawei) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetAuthUrl() string {
	if o == nil || o.AuthUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetAuthUrlOk() (*string, bool) {
	if o == nil || o.AuthUrl == nil {
		return nil, false
	}
	return o.AuthUrl, true
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasAuthUrl() bool {
	if o != nil && o.AuthUrl != nil {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given string and assigns it to the AuthUrl field.
func (o *CreateAuthMethodHuawei) SetAuthUrl(v string) {
	o.AuthUrl = &v
}

// GetBoundDomainId returns the BoundDomainId field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetBoundDomainId() []string {
	if o == nil || o.BoundDomainId == nil {
		var ret []string
		return ret
	}
	return *o.BoundDomainId
}

// GetBoundDomainIdOk returns a tuple with the BoundDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetBoundDomainIdOk() (*[]string, bool) {
	if o == nil || o.BoundDomainId == nil {
		return nil, false
	}
	return o.BoundDomainId, true
}

// HasBoundDomainId returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasBoundDomainId() bool {
	if o != nil && o.BoundDomainId != nil {
		return true
	}

	return false
}

// SetBoundDomainId gets a reference to the given []string and assigns it to the BoundDomainId field.
func (o *CreateAuthMethodHuawei) SetBoundDomainId(v []string) {
	o.BoundDomainId = &v
}

// GetBoundDomainName returns the BoundDomainName field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetBoundDomainName() []string {
	if o == nil || o.BoundDomainName == nil {
		var ret []string
		return ret
	}
	return *o.BoundDomainName
}

// GetBoundDomainNameOk returns a tuple with the BoundDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetBoundDomainNameOk() (*[]string, bool) {
	if o == nil || o.BoundDomainName == nil {
		return nil, false
	}
	return o.BoundDomainName, true
}

// HasBoundDomainName returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasBoundDomainName() bool {
	if o != nil && o.BoundDomainName != nil {
		return true
	}

	return false
}

// SetBoundDomainName gets a reference to the given []string and assigns it to the BoundDomainName field.
func (o *CreateAuthMethodHuawei) SetBoundDomainName(v []string) {
	o.BoundDomainName = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *CreateAuthMethodHuawei) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetBoundTenantId returns the BoundTenantId field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetBoundTenantId() []string {
	if o == nil || o.BoundTenantId == nil {
		var ret []string
		return ret
	}
	return *o.BoundTenantId
}

// GetBoundTenantIdOk returns a tuple with the BoundTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetBoundTenantIdOk() (*[]string, bool) {
	if o == nil || o.BoundTenantId == nil {
		return nil, false
	}
	return o.BoundTenantId, true
}

// HasBoundTenantId returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasBoundTenantId() bool {
	if o != nil && o.BoundTenantId != nil {
		return true
	}

	return false
}

// SetBoundTenantId gets a reference to the given []string and assigns it to the BoundTenantId field.
func (o *CreateAuthMethodHuawei) SetBoundTenantId(v []string) {
	o.BoundTenantId = &v
}

// GetBoundTenantName returns the BoundTenantName field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetBoundTenantName() []string {
	if o == nil || o.BoundTenantName == nil {
		var ret []string
		return ret
	}
	return *o.BoundTenantName
}

// GetBoundTenantNameOk returns a tuple with the BoundTenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetBoundTenantNameOk() (*[]string, bool) {
	if o == nil || o.BoundTenantName == nil {
		return nil, false
	}
	return o.BoundTenantName, true
}

// HasBoundTenantName returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasBoundTenantName() bool {
	if o != nil && o.BoundTenantName != nil {
		return true
	}

	return false
}

// SetBoundTenantName gets a reference to the given []string and assigns it to the BoundTenantName field.
func (o *CreateAuthMethodHuawei) SetBoundTenantName(v []string) {
	o.BoundTenantName = &v
}

// GetBoundUserId returns the BoundUserId field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetBoundUserId() []string {
	if o == nil || o.BoundUserId == nil {
		var ret []string
		return ret
	}
	return *o.BoundUserId
}

// GetBoundUserIdOk returns a tuple with the BoundUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetBoundUserIdOk() (*[]string, bool) {
	if o == nil || o.BoundUserId == nil {
		return nil, false
	}
	return o.BoundUserId, true
}

// HasBoundUserId returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasBoundUserId() bool {
	if o != nil && o.BoundUserId != nil {
		return true
	}

	return false
}

// SetBoundUserId gets a reference to the given []string and assigns it to the BoundUserId field.
func (o *CreateAuthMethodHuawei) SetBoundUserId(v []string) {
	o.BoundUserId = &v
}

// GetBoundUserName returns the BoundUserName field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetBoundUserName() []string {
	if o == nil || o.BoundUserName == nil {
		var ret []string
		return ret
	}
	return *o.BoundUserName
}

// GetBoundUserNameOk returns a tuple with the BoundUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetBoundUserNameOk() (*[]string, bool) {
	if o == nil || o.BoundUserName == nil {
		return nil, false
	}
	return o.BoundUserName, true
}

// HasBoundUserName returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasBoundUserName() bool {
	if o != nil && o.BoundUserName != nil {
		return true
	}

	return false
}

// SetBoundUserName gets a reference to the given []string and assigns it to the BoundUserName field.
func (o *CreateAuthMethodHuawei) SetBoundUserName(v []string) {
	o.BoundUserName = &v
}

// GetName returns the Name field value
func (o *CreateAuthMethodHuawei) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAuthMethodHuawei) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateAuthMethodHuawei) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateAuthMethodHuawei) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodHuawei) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateAuthMethodHuawei) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateAuthMethodHuawei) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateAuthMethodHuawei) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if o.AuthUrl != nil {
		toSerialize["auth-url"] = o.AuthUrl
	}
	if o.BoundDomainId != nil {
		toSerialize["bound-domain-id"] = o.BoundDomainId
	}
	if o.BoundDomainName != nil {
		toSerialize["bound-domain-name"] = o.BoundDomainName
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if o.BoundTenantId != nil {
		toSerialize["bound-tenant-id"] = o.BoundTenantId
	}
	if o.BoundTenantName != nil {
		toSerialize["bound-tenant-name"] = o.BoundTenantName
	}
	if o.BoundUserId != nil {
		toSerialize["bound-user-id"] = o.BoundUserId
	}
	if o.BoundUserName != nil {
		toSerialize["bound-user-name"] = o.BoundUserName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAuthMethodHuawei struct {
	value *CreateAuthMethodHuawei
	isSet bool
}

func (v NullableCreateAuthMethodHuawei) Get() *CreateAuthMethodHuawei {
	return v.value
}

func (v *NullableCreateAuthMethodHuawei) Set(val *CreateAuthMethodHuawei) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodHuawei) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodHuawei) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodHuawei(val *CreateAuthMethodHuawei) *NullableCreateAuthMethodHuawei {
	return &NullableCreateAuthMethodHuawei{value: val, isSet: true}
}

func (v NullableCreateAuthMethodHuawei) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodHuawei) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


