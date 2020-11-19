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

// Role struct for Role
type Role struct {
	ClientPermissions *[]string `json:"client_permissions,omitempty"`
	Comment *string `json:"comment,omitempty"`
	RoleAuthMethodsAssoc *[]RoleAuthMethodAssociation `json:"role_auth_methods_assoc,omitempty"`
	RoleName *string `json:"role_name,omitempty"`
	Rules *Rules `json:"rules,omitempty"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole() *Role {
	this := Role{}
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetClientPermissions returns the ClientPermissions field value if set, zero value otherwise.
func (o *Role) GetClientPermissions() []string {
	if o == nil || o.ClientPermissions == nil {
		var ret []string
		return ret
	}
	return *o.ClientPermissions
}

// GetClientPermissionsOk returns a tuple with the ClientPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetClientPermissionsOk() (*[]string, bool) {
	if o == nil || o.ClientPermissions == nil {
		return nil, false
	}
	return o.ClientPermissions, true
}

// HasClientPermissions returns a boolean if a field has been set.
func (o *Role) HasClientPermissions() bool {
	if o != nil && o.ClientPermissions != nil {
		return true
	}

	return false
}

// SetClientPermissions gets a reference to the given []string and assigns it to the ClientPermissions field.
func (o *Role) SetClientPermissions(v []string) {
	o.ClientPermissions = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Role) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Role) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Role) SetComment(v string) {
	o.Comment = &v
}

// GetRoleAuthMethodsAssoc returns the RoleAuthMethodsAssoc field value if set, zero value otherwise.
func (o *Role) GetRoleAuthMethodsAssoc() []RoleAuthMethodAssociation {
	if o == nil || o.RoleAuthMethodsAssoc == nil {
		var ret []RoleAuthMethodAssociation
		return ret
	}
	return *o.RoleAuthMethodsAssoc
}

// GetRoleAuthMethodsAssocOk returns a tuple with the RoleAuthMethodsAssoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRoleAuthMethodsAssocOk() (*[]RoleAuthMethodAssociation, bool) {
	if o == nil || o.RoleAuthMethodsAssoc == nil {
		return nil, false
	}
	return o.RoleAuthMethodsAssoc, true
}

// HasRoleAuthMethodsAssoc returns a boolean if a field has been set.
func (o *Role) HasRoleAuthMethodsAssoc() bool {
	if o != nil && o.RoleAuthMethodsAssoc != nil {
		return true
	}

	return false
}

// SetRoleAuthMethodsAssoc gets a reference to the given []RoleAuthMethodAssociation and assigns it to the RoleAuthMethodsAssoc field.
func (o *Role) SetRoleAuthMethodsAssoc(v []RoleAuthMethodAssociation) {
	o.RoleAuthMethodsAssoc = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *Role) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *Role) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *Role) SetRoleName(v string) {
	o.RoleName = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *Role) GetRules() Rules {
	if o == nil || o.Rules == nil {
		var ret Rules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRulesOk() (*Rules, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *Role) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given Rules and assigns it to the Rules field.
func (o *Role) SetRules(v Rules) {
	o.Rules = &v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientPermissions != nil {
		toSerialize["client_permissions"] = o.ClientPermissions
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.RoleAuthMethodsAssoc != nil {
		toSerialize["role_auth_methods_assoc"] = o.RoleAuthMethodsAssoc
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


