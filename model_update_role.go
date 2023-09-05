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

// UpdateRole struct for UpdateRole
type UpdateRole struct {
	// Allow this role to view analytics. Currently only 'none', 'own', 'all' values are supported, allowing associated auth methods to view reports produced by the same auth methods.
	AnalyticsAccess *string `json:"analytics-access,omitempty"`
	// Allow this role to view audit logs. Currently only 'none', 'own' and 'all' values are supported, allowing associated auth methods to view audit logs produced by the same auth methods.
	AuditAccess *string `json:"audit-access,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Allow this role to view gw analytics. Currently only 'none', 'own', 'all' values are supported, allowing associated auth methods to view reports produced by the same auth methods.
	GwAnalyticsAccess *string `json:"gw-analytics-access,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Role name
	Name string `json:"name"`
	// Deprecated - use description
	NewComment *string `json:"new-comment,omitempty"`
	// New Role name
	NewName *string `json:"new-name,omitempty"`
	// Allow this role to view SRA Clusters. Currently only 'none', 'own', 'all' values are supported.
	SraReportsAccess *string `json:"sra-reports-access,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Allow this role to view Usage Report. Currently only 'none' and 'all' values are supported.
	UsageReportsAccess *string `json:"usage-reports-access,omitempty"`
}

// NewUpdateRole instantiates a new UpdateRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRole(name string, ) *UpdateRole {
	this := UpdateRole{}
	var description string = "default_comment"
	this.Description = &description
	var json bool = false
	this.Json = &json
	this.Name = name
	var newComment string = "default_comment"
	this.NewComment = &newComment
	return &this
}

// NewUpdateRoleWithDefaults instantiates a new UpdateRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleWithDefaults() *UpdateRole {
	this := UpdateRole{}
	var description string = "default_comment"
	this.Description = &description
	var json bool = false
	this.Json = &json
	var newComment string = "default_comment"
	this.NewComment = &newComment
	return &this
}

// GetAnalyticsAccess returns the AnalyticsAccess field value if set, zero value otherwise.
func (o *UpdateRole) GetAnalyticsAccess() string {
	if o == nil || o.AnalyticsAccess == nil {
		var ret string
		return ret
	}
	return *o.AnalyticsAccess
}

// GetAnalyticsAccessOk returns a tuple with the AnalyticsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetAnalyticsAccessOk() (*string, bool) {
	if o == nil || o.AnalyticsAccess == nil {
		return nil, false
	}
	return o.AnalyticsAccess, true
}

// HasAnalyticsAccess returns a boolean if a field has been set.
func (o *UpdateRole) HasAnalyticsAccess() bool {
	if o != nil && o.AnalyticsAccess != nil {
		return true
	}

	return false
}

// SetAnalyticsAccess gets a reference to the given string and assigns it to the AnalyticsAccess field.
func (o *UpdateRole) SetAnalyticsAccess(v string) {
	o.AnalyticsAccess = &v
}

// GetAuditAccess returns the AuditAccess field value if set, zero value otherwise.
func (o *UpdateRole) GetAuditAccess() string {
	if o == nil || o.AuditAccess == nil {
		var ret string
		return ret
	}
	return *o.AuditAccess
}

// GetAuditAccessOk returns a tuple with the AuditAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetAuditAccessOk() (*string, bool) {
	if o == nil || o.AuditAccess == nil {
		return nil, false
	}
	return o.AuditAccess, true
}

// HasAuditAccess returns a boolean if a field has been set.
func (o *UpdateRole) HasAuditAccess() bool {
	if o != nil && o.AuditAccess != nil {
		return true
	}

	return false
}

// SetAuditAccess gets a reference to the given string and assigns it to the AuditAccess field.
func (o *UpdateRole) SetAuditAccess(v string) {
	o.AuditAccess = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateRole) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateRole) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateRole) SetDescription(v string) {
	o.Description = &v
}

// GetGwAnalyticsAccess returns the GwAnalyticsAccess field value if set, zero value otherwise.
func (o *UpdateRole) GetGwAnalyticsAccess() string {
	if o == nil || o.GwAnalyticsAccess == nil {
		var ret string
		return ret
	}
	return *o.GwAnalyticsAccess
}

// GetGwAnalyticsAccessOk returns a tuple with the GwAnalyticsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetGwAnalyticsAccessOk() (*string, bool) {
	if o == nil || o.GwAnalyticsAccess == nil {
		return nil, false
	}
	return o.GwAnalyticsAccess, true
}

// HasGwAnalyticsAccess returns a boolean if a field has been set.
func (o *UpdateRole) HasGwAnalyticsAccess() bool {
	if o != nil && o.GwAnalyticsAccess != nil {
		return true
	}

	return false
}

// SetGwAnalyticsAccess gets a reference to the given string and assigns it to the GwAnalyticsAccess field.
func (o *UpdateRole) SetGwAnalyticsAccess(v string) {
	o.GwAnalyticsAccess = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateRole) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateRole) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateRole) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *UpdateRole) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateRole) SetName(v string) {
	o.Name = v
}

// GetNewComment returns the NewComment field value if set, zero value otherwise.
func (o *UpdateRole) GetNewComment() string {
	if o == nil || o.NewComment == nil {
		var ret string
		return ret
	}
	return *o.NewComment
}

// GetNewCommentOk returns a tuple with the NewComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetNewCommentOk() (*string, bool) {
	if o == nil || o.NewComment == nil {
		return nil, false
	}
	return o.NewComment, true
}

// HasNewComment returns a boolean if a field has been set.
func (o *UpdateRole) HasNewComment() bool {
	if o != nil && o.NewComment != nil {
		return true
	}

	return false
}

// SetNewComment gets a reference to the given string and assigns it to the NewComment field.
func (o *UpdateRole) SetNewComment(v string) {
	o.NewComment = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateRole) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateRole) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateRole) SetNewName(v string) {
	o.NewName = &v
}

// GetSraReportsAccess returns the SraReportsAccess field value if set, zero value otherwise.
func (o *UpdateRole) GetSraReportsAccess() string {
	if o == nil || o.SraReportsAccess == nil {
		var ret string
		return ret
	}
	return *o.SraReportsAccess
}

// GetSraReportsAccessOk returns a tuple with the SraReportsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetSraReportsAccessOk() (*string, bool) {
	if o == nil || o.SraReportsAccess == nil {
		return nil, false
	}
	return o.SraReportsAccess, true
}

// HasSraReportsAccess returns a boolean if a field has been set.
func (o *UpdateRole) HasSraReportsAccess() bool {
	if o != nil && o.SraReportsAccess != nil {
		return true
	}

	return false
}

// SetSraReportsAccess gets a reference to the given string and assigns it to the SraReportsAccess field.
func (o *UpdateRole) SetSraReportsAccess(v string) {
	o.SraReportsAccess = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateRole) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateRole) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateRole) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateRole) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateRole) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateRole) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsageReportsAccess returns the UsageReportsAccess field value if set, zero value otherwise.
func (o *UpdateRole) GetUsageReportsAccess() string {
	if o == nil || o.UsageReportsAccess == nil {
		var ret string
		return ret
	}
	return *o.UsageReportsAccess
}

// GetUsageReportsAccessOk returns a tuple with the UsageReportsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetUsageReportsAccessOk() (*string, bool) {
	if o == nil || o.UsageReportsAccess == nil {
		return nil, false
	}
	return o.UsageReportsAccess, true
}

// HasUsageReportsAccess returns a boolean if a field has been set.
func (o *UpdateRole) HasUsageReportsAccess() bool {
	if o != nil && o.UsageReportsAccess != nil {
		return true
	}

	return false
}

// SetUsageReportsAccess gets a reference to the given string and assigns it to the UsageReportsAccess field.
func (o *UpdateRole) SetUsageReportsAccess(v string) {
	o.UsageReportsAccess = &v
}

func (o UpdateRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnalyticsAccess != nil {
		toSerialize["analytics-access"] = o.AnalyticsAccess
	}
	if o.AuditAccess != nil {
		toSerialize["audit-access"] = o.AuditAccess
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GwAnalyticsAccess != nil {
		toSerialize["gw-analytics-access"] = o.GwAnalyticsAccess
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewComment != nil {
		toSerialize["new-comment"] = o.NewComment
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.SraReportsAccess != nil {
		toSerialize["sra-reports-access"] = o.SraReportsAccess
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UsageReportsAccess != nil {
		toSerialize["usage-reports-access"] = o.UsageReportsAccess
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRole struct {
	value *UpdateRole
	isSet bool
}

func (v NullableUpdateRole) Get() *UpdateRole {
	return v.value
}

func (v *NullableUpdateRole) Set(val *UpdateRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRole(val *UpdateRole) *NullableUpdateRole {
	return &NullableUpdateRole{value: val, isSet: true}
}

func (v NullableUpdateRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


