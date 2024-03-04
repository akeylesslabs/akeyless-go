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

// UpdateAuthMethodOCI updateAuthMethodOCI is a command that updates an auth method that will be used in the account using OCI principle and groups.
type UpdateAuthMethodOCI struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// A CIDR whitelist with the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// Auth Method description
	Description *string `json:"description,omitempty"`
	// if true: enforce role-association must include sub claims
	ForceSubClaims *bool `json:"force-sub-claims,omitempty"`
	// A list of required groups ocids
	GroupOcid []string `json:"group-ocid"`
	// A CIDR whitelist with the GW IPs that the access is restricted to
	GwBoundIps *[]string `json:"gw-bound-ips,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Jwt TTL
	JwtTtl *int64 `json:"jwt-ttl,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// Auth Method new name
	NewName *string `json:"new-name,omitempty"`
	// Choose the relevant product type for the auth method [sm, sra, pm, dp, ca]
	ProductType *[]string `json:"product-type,omitempty"`
	// The Oracle Cloud tenant ID
	TenantOcid string `json:"tenant-ocid"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateAuthMethodOCI instantiates a new UpdateAuthMethodOCI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthMethodOCI(groupOcid []string, name string, tenantOcid string, ) *UpdateAuthMethodOCI {
	this := UpdateAuthMethodOCI{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	this.GroupOcid = groupOcid
	var json bool = false
	this.Json = &json
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	this.Name = name
	this.TenantOcid = tenantOcid
	return &this
}

// NewUpdateAuthMethodOCIWithDefaults instantiates a new UpdateAuthMethodOCI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthMethodOCIWithDefaults() *UpdateAuthMethodOCI {
	this := UpdateAuthMethodOCI{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var json bool = false
	this.Json = &json
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *UpdateAuthMethodOCI) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *UpdateAuthMethodOCI) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateAuthMethodOCI) SetDescription(v string) {
	o.Description = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *UpdateAuthMethodOCI) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetGroupOcid returns the GroupOcid field value
func (o *UpdateAuthMethodOCI) GetGroupOcid() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.GroupOcid
}

// GetGroupOcidOk returns a tuple with the GroupOcid field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetGroupOcidOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GroupOcid, true
}

// SetGroupOcid sets field value
func (o *UpdateAuthMethodOCI) SetGroupOcid(v []string) {
	o.GroupOcid = v
}

// GetGwBoundIps returns the GwBoundIps field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetGwBoundIps() []string {
	if o == nil || o.GwBoundIps == nil {
		var ret []string
		return ret
	}
	return *o.GwBoundIps
}

// GetGwBoundIpsOk returns a tuple with the GwBoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetGwBoundIpsOk() (*[]string, bool) {
	if o == nil || o.GwBoundIps == nil {
		return nil, false
	}
	return o.GwBoundIps, true
}

// HasGwBoundIps returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasGwBoundIps() bool {
	if o != nil && o.GwBoundIps != nil {
		return true
	}

	return false
}

// SetGwBoundIps gets a reference to the given []string and assigns it to the GwBoundIps field.
func (o *UpdateAuthMethodOCI) SetGwBoundIps(v []string) {
	o.GwBoundIps = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateAuthMethodOCI) SetJson(v bool) {
	o.Json = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *UpdateAuthMethodOCI) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *UpdateAuthMethodOCI) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateAuthMethodOCI) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateAuthMethodOCI) SetNewName(v string) {
	o.NewName = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetProductType() []string {
	if o == nil || o.ProductType == nil {
		var ret []string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetProductTypeOk() (*[]string, bool) {
	if o == nil || o.ProductType == nil {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasProductType() bool {
	if o != nil && o.ProductType != nil {
		return true
	}

	return false
}

// SetProductType gets a reference to the given []string and assigns it to the ProductType field.
func (o *UpdateAuthMethodOCI) SetProductType(v []string) {
	o.ProductType = &v
}

// GetTenantOcid returns the TenantOcid field value
func (o *UpdateAuthMethodOCI) GetTenantOcid() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TenantOcid
}

// GetTenantOcidOk returns a tuple with the TenantOcid field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetTenantOcidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantOcid, true
}

// SetTenantOcid sets field value
func (o *UpdateAuthMethodOCI) SetTenantOcid(v string) {
	o.TenantOcid = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateAuthMethodOCI) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateAuthMethodOCI) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOCI) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodOCI) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateAuthMethodOCI) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateAuthMethodOCI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ForceSubClaims != nil {
		toSerialize["force-sub-claims"] = o.ForceSubClaims
	}
	if true {
		toSerialize["group-ocid"] = o.GroupOcid
	}
	if o.GwBoundIps != nil {
		toSerialize["gw-bound-ips"] = o.GwBoundIps
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.JwtTtl != nil {
		toSerialize["jwt-ttl"] = o.JwtTtl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.ProductType != nil {
		toSerialize["product-type"] = o.ProductType
	}
	if true {
		toSerialize["tenant-ocid"] = o.TenantOcid
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAuthMethodOCI struct {
	value *UpdateAuthMethodOCI
	isSet bool
}

func (v NullableUpdateAuthMethodOCI) Get() *UpdateAuthMethodOCI {
	return v.value
}

func (v *NullableUpdateAuthMethodOCI) Set(val *UpdateAuthMethodOCI) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthMethodOCI) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthMethodOCI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthMethodOCI(val *UpdateAuthMethodOCI) *NullableUpdateAuthMethodOCI {
	return &NullableUpdateAuthMethodOCI{value: val, isSet: true}
}

func (v NullableUpdateAuthMethodOCI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthMethodOCI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


