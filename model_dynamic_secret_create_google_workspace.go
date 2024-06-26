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

// DynamicSecretCreateGoogleWorkspace dynamicSecretCreateGoogleWorkspace is a command that creates Google Workspace dynamic secret
type DynamicSecretCreateGoogleWorkspace struct {
	AccessMode string `json:"access-mode"`
	// Admin user email
	AdminEmail string `json:"admin-email"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Base64-encoded service account private key text
	GcpKey *string `json:"gcp-key,omitempty"`
	// A group email, relevant only for group access-mode
	GroupEmail *string `json:"group-email,omitempty"`
	GroupRole *string `json:"group-role,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Name of the admin role to assign to the user, relevant only for role access-mode
	RoleName *string `json:"role-name,omitempty"`
	RoleScope *string `json:"role-scope,omitempty"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Name of existing target to use in dynamic secret creation
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewDynamicSecretCreateGoogleWorkspace instantiates a new DynamicSecretCreateGoogleWorkspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretCreateGoogleWorkspace(accessMode string, adminEmail string, name string, ) *DynamicSecretCreateGoogleWorkspace {
	this := DynamicSecretCreateGoogleWorkspace{}
	this.AccessMode = accessMode
	this.AdminEmail = adminEmail
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretCreateGoogleWorkspaceWithDefaults instantiates a new DynamicSecretCreateGoogleWorkspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretCreateGoogleWorkspaceWithDefaults() *DynamicSecretCreateGoogleWorkspace {
	this := DynamicSecretCreateGoogleWorkspace{}
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetAccessMode returns the AccessMode field value
func (o *DynamicSecretCreateGoogleWorkspace) GetAccessMode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetAccessModeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessMode, true
}

// SetAccessMode sets field value
func (o *DynamicSecretCreateGoogleWorkspace) SetAccessMode(v string) {
	o.AccessMode = v
}

// GetAdminEmail returns the AdminEmail field value
func (o *DynamicSecretCreateGoogleWorkspace) GetAdminEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AdminEmail
}

// GetAdminEmailOk returns a tuple with the AdminEmail field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetAdminEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AdminEmail, true
}

// SetAdminEmail sets field value
func (o *DynamicSecretCreateGoogleWorkspace) SetAdminEmail(v string) {
	o.AdminEmail = v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretCreateGoogleWorkspace) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretCreateGoogleWorkspace) SetDescription(v string) {
	o.Description = &v
}

// GetGcpKey returns the GcpKey field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetGcpKey() string {
	if o == nil || o.GcpKey == nil {
		var ret string
		return ret
	}
	return *o.GcpKey
}

// GetGcpKeyOk returns a tuple with the GcpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetGcpKeyOk() (*string, bool) {
	if o == nil || o.GcpKey == nil {
		return nil, false
	}
	return o.GcpKey, true
}

// HasGcpKey returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasGcpKey() bool {
	if o != nil && o.GcpKey != nil {
		return true
	}

	return false
}

// SetGcpKey gets a reference to the given string and assigns it to the GcpKey field.
func (o *DynamicSecretCreateGoogleWorkspace) SetGcpKey(v string) {
	o.GcpKey = &v
}

// GetGroupEmail returns the GroupEmail field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetGroupEmail() string {
	if o == nil || o.GroupEmail == nil {
		var ret string
		return ret
	}
	return *o.GroupEmail
}

// GetGroupEmailOk returns a tuple with the GroupEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetGroupEmailOk() (*string, bool) {
	if o == nil || o.GroupEmail == nil {
		return nil, false
	}
	return o.GroupEmail, true
}

// HasGroupEmail returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasGroupEmail() bool {
	if o != nil && o.GroupEmail != nil {
		return true
	}

	return false
}

// SetGroupEmail gets a reference to the given string and assigns it to the GroupEmail field.
func (o *DynamicSecretCreateGoogleWorkspace) SetGroupEmail(v string) {
	o.GroupEmail = &v
}

// GetGroupRole returns the GroupRole field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetGroupRole() string {
	if o == nil || o.GroupRole == nil {
		var ret string
		return ret
	}
	return *o.GroupRole
}

// GetGroupRoleOk returns a tuple with the GroupRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetGroupRoleOk() (*string, bool) {
	if o == nil || o.GroupRole == nil {
		return nil, false
	}
	return o.GroupRole, true
}

// HasGroupRole returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasGroupRole() bool {
	if o != nil && o.GroupRole != nil {
		return true
	}

	return false
}

// SetGroupRole gets a reference to the given string and assigns it to the GroupRole field.
func (o *DynamicSecretCreateGoogleWorkspace) SetGroupRole(v string) {
	o.GroupRole = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretCreateGoogleWorkspace) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretCreateGoogleWorkspace) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretCreateGoogleWorkspace) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretCreateGoogleWorkspace) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *DynamicSecretCreateGoogleWorkspace) SetRoleName(v string) {
	o.RoleName = &v
}

// GetRoleScope returns the RoleScope field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetRoleScope() string {
	if o == nil || o.RoleScope == nil {
		var ret string
		return ret
	}
	return *o.RoleScope
}

// GetRoleScopeOk returns a tuple with the RoleScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetRoleScopeOk() (*string, bool) {
	if o == nil || o.RoleScope == nil {
		return nil, false
	}
	return o.RoleScope, true
}

// HasRoleScope returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasRoleScope() bool {
	if o != nil && o.RoleScope != nil {
		return true
	}

	return false
}

// SetRoleScope gets a reference to the given string and assigns it to the RoleScope field.
func (o *DynamicSecretCreateGoogleWorkspace) SetRoleScope(v string) {
	o.RoleScope = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretCreateGoogleWorkspace) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretCreateGoogleWorkspace) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretCreateGoogleWorkspace) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretCreateGoogleWorkspace) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretCreateGoogleWorkspace) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGoogleWorkspace) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretCreateGoogleWorkspace) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretCreateGoogleWorkspace) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretCreateGoogleWorkspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access-mode"] = o.AccessMode
	}
	if true {
		toSerialize["admin-email"] = o.AdminEmail
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GcpKey != nil {
		toSerialize["gcp-key"] = o.GcpKey
	}
	if o.GroupEmail != nil {
		toSerialize["group-email"] = o.GroupEmail
	}
	if o.GroupRole != nil {
		toSerialize["group-role"] = o.GroupRole
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.RoleName != nil {
		toSerialize["role-name"] = o.RoleName
	}
	if o.RoleScope != nil {
		toSerialize["role-scope"] = o.RoleScope
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetName != nil {
		toSerialize["target-name"] = o.TargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretCreateGoogleWorkspace struct {
	value *DynamicSecretCreateGoogleWorkspace
	isSet bool
}

func (v NullableDynamicSecretCreateGoogleWorkspace) Get() *DynamicSecretCreateGoogleWorkspace {
	return v.value
}

func (v *NullableDynamicSecretCreateGoogleWorkspace) Set(val *DynamicSecretCreateGoogleWorkspace) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretCreateGoogleWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretCreateGoogleWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretCreateGoogleWorkspace(val *DynamicSecretCreateGoogleWorkspace) *NullableDynamicSecretCreateGoogleWorkspace {
	return &NullableDynamicSecretCreateGoogleWorkspace{value: val, isSet: true}
}

func (v NullableDynamicSecretCreateGoogleWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretCreateGoogleWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


