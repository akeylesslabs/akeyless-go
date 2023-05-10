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

// SetRoleRule struct for SetRoleRule
type SetRoleRule struct {
	// List of the approved/denied capabilities in the path options: [read, create, update, delete, list, deny]
	Capability []string `json:"capability"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The path the rule refers to
	Path string `json:"path"`
	// The role name to be updated
	RoleName string `json:"role-name"`
	// item-rule, target-rule, role-rule, auth-method-rule, search-rule, reports-rule, gw-reports-rule or sra-reports-rule
	RuleType *string `json:"rule-type,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// RoleRule ttl
	Ttl *int32 `json:"ttl,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewSetRoleRule instantiates a new SetRoleRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetRoleRule(capability []string, path string, roleName string, ) *SetRoleRule {
	this := SetRoleRule{}
	this.Capability = capability
	var json bool = false
	this.Json = &json
	this.Path = path
	this.RoleName = roleName
	var ruleType string = "item-rule"
	this.RuleType = &ruleType
	return &this
}

// NewSetRoleRuleWithDefaults instantiates a new SetRoleRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetRoleRuleWithDefaults() *SetRoleRule {
	this := SetRoleRule{}
	var json bool = false
	this.Json = &json
	var ruleType string = "item-rule"
	this.RuleType = &ruleType
	return &this
}

// GetCapability returns the Capability field value
func (o *SetRoleRule) GetCapability() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *SetRoleRule) GetCapabilityOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *SetRoleRule) SetCapability(v []string) {
	o.Capability = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *SetRoleRule) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRoleRule) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *SetRoleRule) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *SetRoleRule) SetJson(v bool) {
	o.Json = &v
}

// GetPath returns the Path field value
func (o *SetRoleRule) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *SetRoleRule) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *SetRoleRule) SetPath(v string) {
	o.Path = v
}

// GetRoleName returns the RoleName field value
func (o *SetRoleRule) GetRoleName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *SetRoleRule) GetRoleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *SetRoleRule) SetRoleName(v string) {
	o.RoleName = v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *SetRoleRule) GetRuleType() string {
	if o == nil || o.RuleType == nil {
		var ret string
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRoleRule) GetRuleTypeOk() (*string, bool) {
	if o == nil || o.RuleType == nil {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *SetRoleRule) HasRuleType() bool {
	if o != nil && o.RuleType != nil {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *SetRoleRule) SetRuleType(v string) {
	o.RuleType = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SetRoleRule) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRoleRule) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SetRoleRule) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SetRoleRule) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *SetRoleRule) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRoleRule) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *SetRoleRule) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *SetRoleRule) SetTtl(v int32) {
	o.Ttl = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *SetRoleRule) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRoleRule) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *SetRoleRule) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *SetRoleRule) SetUidToken(v string) {
	o.UidToken = &v
}

func (o SetRoleRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["capability"] = o.Capability
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["role-name"] = o.RoleName
	}
	if o.RuleType != nil {
		toSerialize["rule-type"] = o.RuleType
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableSetRoleRule struct {
	value *SetRoleRule
	isSet bool
}

func (v NullableSetRoleRule) Get() *SetRoleRule {
	return v.value
}

func (v *NullableSetRoleRule) Set(val *SetRoleRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSetRoleRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSetRoleRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetRoleRule(val *SetRoleRule) *NullableSetRoleRule {
	return &NullableSetRoleRule{value: val, isSet: true}
}

func (v NullableSetRoleRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetRoleRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


