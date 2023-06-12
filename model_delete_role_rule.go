/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// checks if the DeleteRoleRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRoleRule{}

// DeleteRoleRule struct for DeleteRoleRule
type DeleteRoleRule struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The path the rule refers to
	Path string `json:"path"`
	// The role name to be updated
	RoleName string `json:"role-name"`
	// item-rule, role-rule, auth-method-rule, search-rule, reports-rule, gw-reports-rule or sra-reports-rule
	RuleType *string `json:"rule-type,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDeleteRoleRule instantiates a new DeleteRoleRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRoleRule(path string, roleName string) *DeleteRoleRule {
	this := DeleteRoleRule{}
	var json bool = false
	this.Json = &json
	this.Path = path
	this.RoleName = roleName
	var ruleType string = "item-rule"
	this.RuleType = &ruleType
	return &this
}

// NewDeleteRoleRuleWithDefaults instantiates a new DeleteRoleRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRoleRuleWithDefaults() *DeleteRoleRule {
	this := DeleteRoleRule{}
	var json bool = false
	this.Json = &json
	var ruleType string = "item-rule"
	this.RuleType = &ruleType
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DeleteRoleRule) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRoleRule) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DeleteRoleRule) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DeleteRoleRule) SetJson(v bool) {
	o.Json = &v
}

// GetPath returns the Path field value
func (o *DeleteRoleRule) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DeleteRoleRule) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DeleteRoleRule) SetPath(v string) {
	o.Path = v
}

// GetRoleName returns the RoleName field value
func (o *DeleteRoleRule) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *DeleteRoleRule) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *DeleteRoleRule) SetRoleName(v string) {
	o.RoleName = v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *DeleteRoleRule) GetRuleType() string {
	if o == nil || IsNil(o.RuleType) {
		var ret string
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRoleRule) GetRuleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *DeleteRoleRule) HasRuleType() bool {
	if o != nil && !IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *DeleteRoleRule) SetRuleType(v string) {
	o.RuleType = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DeleteRoleRule) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRoleRule) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DeleteRoleRule) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DeleteRoleRule) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DeleteRoleRule) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRoleRule) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DeleteRoleRule) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DeleteRoleRule) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DeleteRoleRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRoleRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["path"] = o.Path
	toSerialize["role-name"] = o.RoleName
	if !IsNil(o.RuleType) {
		toSerialize["rule-type"] = o.RuleType
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableDeleteRoleRule struct {
	value *DeleteRoleRule
	isSet bool
}

func (v NullableDeleteRoleRule) Get() *DeleteRoleRule {
	return v.value
}

func (v *NullableDeleteRoleRule) Set(val *DeleteRoleRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRoleRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRoleRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRoleRule(val *DeleteRoleRule) *NullableDeleteRoleRule {
	return &NullableDeleteRoleRule{value: val, isSet: true}
}

func (v NullableDeleteRoleRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRoleRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


