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

// GatewayUpdateAllowedAccess gatewayUpdateAllowedAccess is a command that update allowed access in Gator
type GatewayUpdateAllowedAccess struct {
	SubClaimsCaseInsensitive *bool `json:"SubClaimsCaseInsensitive,omitempty"`
	// Access ID The access id to be attached to this allowed access. Auth method with this access id should already exist.
	AccessId string `json:"access-id"`
	// Treat sub claims as case-sensitive [true/false]
	CaseSensitive *string `json:"case-sensitive,omitempty"`
	// Allowed access description
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Allowed access name
	Name string `json:"name"`
	// New allowed access name
	NewName *string `json:"new-name,omitempty"`
	// Permissions  Comma-seperated list of permissions for this allowed access. Available permissions: [defaults,targets,classic_keys,automatic_migration,ldap_auth,dynamic_secret,k8s_auth,log_forwarding,zero_knowledge_encryption,rotated_secret,caching,event_forwarding,admin,kmip,general]
	Permissions *string `json:"permissions,omitempty"`
	// Sub claims key/val of sub claims, e.g group=admins,developers
	SubClaims *map[string]string `json:"sub-claims,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayUpdateAllowedAccess instantiates a new GatewayUpdateAllowedAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateAllowedAccess(accessId string, name string, ) *GatewayUpdateAllowedAccess {
	this := GatewayUpdateAllowedAccess{}
	this.AccessId = accessId
	var caseSensitive string = "true"
	this.CaseSensitive = &caseSensitive
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewGatewayUpdateAllowedAccessWithDefaults instantiates a new GatewayUpdateAllowedAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateAllowedAccessWithDefaults() *GatewayUpdateAllowedAccess {
	this := GatewayUpdateAllowedAccess{}
	var caseSensitive string = "true"
	this.CaseSensitive = &caseSensitive
	var json bool = false
	this.Json = &json
	return &this
}

// GetSubClaimsCaseInsensitive returns the SubClaimsCaseInsensitive field value if set, zero value otherwise.
func (o *GatewayUpdateAllowedAccess) GetSubClaimsCaseInsensitive() bool {
	if o == nil || o.SubClaimsCaseInsensitive == nil {
		var ret bool
		return ret
	}
	return *o.SubClaimsCaseInsensitive
}

// GetSubClaimsCaseInsensitiveOk returns a tuple with the SubClaimsCaseInsensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetSubClaimsCaseInsensitiveOk() (*bool, bool) {
	if o == nil || o.SubClaimsCaseInsensitive == nil {
		return nil, false
	}
	return o.SubClaimsCaseInsensitive, true
}

// HasSubClaimsCaseInsensitive returns a boolean if a field has been set.
func (o *GatewayUpdateAllowedAccess) HasSubClaimsCaseInsensitive() bool {
	if o != nil && o.SubClaimsCaseInsensitive != nil {
		return true
	}

	return false
}

// SetSubClaimsCaseInsensitive gets a reference to the given bool and assigns it to the SubClaimsCaseInsensitive field.
func (o *GatewayUpdateAllowedAccess) SetSubClaimsCaseInsensitive(v bool) {
	o.SubClaimsCaseInsensitive = &v
}

// GetAccessId returns the AccessId field value
func (o *GatewayUpdateAllowedAccess) GetAccessId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetAccessIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessId, true
}

// SetAccessId sets field value
func (o *GatewayUpdateAllowedAccess) SetAccessId(v string) {
	o.AccessId = v
}

// GetCaseSensitive returns the CaseSensitive field value if set, zero value otherwise.
func (o *GatewayUpdateAllowedAccess) GetCaseSensitive() string {
	if o == nil || o.CaseSensitive == nil {
		var ret string
		return ret
	}
	return *o.CaseSensitive
}

// GetCaseSensitiveOk returns a tuple with the CaseSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetCaseSensitiveOk() (*string, bool) {
	if o == nil || o.CaseSensitive == nil {
		return nil, false
	}
	return o.CaseSensitive, true
}

// HasCaseSensitive returns a boolean if a field has been set.
func (o *GatewayUpdateAllowedAccess) HasCaseSensitive() bool {
	if o != nil && o.CaseSensitive != nil {
		return true
	}

	return false
}

// SetCaseSensitive gets a reference to the given string and assigns it to the CaseSensitive field.
func (o *GatewayUpdateAllowedAccess) SetCaseSensitive(v string) {
	o.CaseSensitive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GatewayUpdateAllowedAccess) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GatewayUpdateAllowedAccess) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GatewayUpdateAllowedAccess) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateAllowedAccess) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateAllowedAccess) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateAllowedAccess) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateAllowedAccess) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateAllowedAccess) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateAllowedAccess) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateAllowedAccess) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateAllowedAccess) SetNewName(v string) {
	o.NewName = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GatewayUpdateAllowedAccess) GetPermissions() string {
	if o == nil || o.Permissions == nil {
		var ret string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetPermissionsOk() (*string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GatewayUpdateAllowedAccess) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given string and assigns it to the Permissions field.
func (o *GatewayUpdateAllowedAccess) SetPermissions(v string) {
	o.Permissions = &v
}

// GetSubClaims returns the SubClaims field value if set, zero value otherwise.
func (o *GatewayUpdateAllowedAccess) GetSubClaims() map[string]string {
	if o == nil || o.SubClaims == nil {
		var ret map[string]string
		return ret
	}
	return *o.SubClaims
}

// GetSubClaimsOk returns a tuple with the SubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetSubClaimsOk() (*map[string]string, bool) {
	if o == nil || o.SubClaims == nil {
		return nil, false
	}
	return o.SubClaims, true
}

// HasSubClaims returns a boolean if a field has been set.
func (o *GatewayUpdateAllowedAccess) HasSubClaims() bool {
	if o != nil && o.SubClaims != nil {
		return true
	}

	return false
}

// SetSubClaims gets a reference to the given map[string]string and assigns it to the SubClaims field.
func (o *GatewayUpdateAllowedAccess) SetSubClaims(v map[string]string) {
	o.SubClaims = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateAllowedAccess) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateAllowedAccess) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateAllowedAccess) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateAllowedAccess) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateAllowedAccess) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateAllowedAccess) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateAllowedAccess) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayUpdateAllowedAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubClaimsCaseInsensitive != nil {
		toSerialize["SubClaimsCaseInsensitive"] = o.SubClaimsCaseInsensitive
	}
	if true {
		toSerialize["access-id"] = o.AccessId
	}
	if o.CaseSensitive != nil {
		toSerialize["case-sensitive"] = o.CaseSensitive
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.SubClaims != nil {
		toSerialize["sub-claims"] = o.SubClaims
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateAllowedAccess struct {
	value *GatewayUpdateAllowedAccess
	isSet bool
}

func (v NullableGatewayUpdateAllowedAccess) Get() *GatewayUpdateAllowedAccess {
	return v.value
}

func (v *NullableGatewayUpdateAllowedAccess) Set(val *GatewayUpdateAllowedAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateAllowedAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateAllowedAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateAllowedAccess(val *GatewayUpdateAllowedAccess) *NullableGatewayUpdateAllowedAccess {
	return &NullableGatewayUpdateAllowedAccess{value: val, isSet: true}
}

func (v NullableGatewayUpdateAllowedAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateAllowedAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


