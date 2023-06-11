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

// checks if the CreateRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRole{}

// CreateRole struct for CreateRole
type CreateRole struct {
	// Allow this role to view analytics. Currently only 'none', 'own', 'all' values are supported, allowing associated auth methods to view reports produced by the same auth methods.
	AnalyticsAccess *string `json:"analytics-access,omitempty"`
	// Allow this role to view audit logs. Currently only 'none', 'own' and 'all' values are supported, allowing associated auth methods to view audit logs produced by the same auth methods.
	AuditAccess *string `json:"audit-access,omitempty"`
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Allow this role to view gw analytics. Currently only 'none', 'own', 'all' values are supported, allowing associated auth methods to view reports produced by the same auth methods.
	GwAnalyticsAccess *string `json:"gw-analytics-access,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Role name
	Name string `json:"name"`
	// Allow this role to view SRA Clusters. Currently only 'none', 'own', 'all' values are supported.
	SraReportsAccess *string `json:"sra-reports-access,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateRole instantiates a new CreateRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRole(name string) *CreateRole {
	this := CreateRole{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewCreateRoleWithDefaults instantiates a new CreateRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleWithDefaults() *CreateRole {
	this := CreateRole{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAnalyticsAccess returns the AnalyticsAccess field value if set, zero value otherwise.
func (o *CreateRole) GetAnalyticsAccess() string {
	if o == nil || IsNil(o.AnalyticsAccess) {
		var ret string
		return ret
	}
	return *o.AnalyticsAccess
}

// GetAnalyticsAccessOk returns a tuple with the AnalyticsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetAnalyticsAccessOk() (*string, bool) {
	if o == nil || IsNil(o.AnalyticsAccess) {
		return nil, false
	}
	return o.AnalyticsAccess, true
}

// HasAnalyticsAccess returns a boolean if a field has been set.
func (o *CreateRole) HasAnalyticsAccess() bool {
	if o != nil && !IsNil(o.AnalyticsAccess) {
		return true
	}

	return false
}

// SetAnalyticsAccess gets a reference to the given string and assigns it to the AnalyticsAccess field.
func (o *CreateRole) SetAnalyticsAccess(v string) {
	o.AnalyticsAccess = &v
}

// GetAuditAccess returns the AuditAccess field value if set, zero value otherwise.
func (o *CreateRole) GetAuditAccess() string {
	if o == nil || IsNil(o.AuditAccess) {
		var ret string
		return ret
	}
	return *o.AuditAccess
}

// GetAuditAccessOk returns a tuple with the AuditAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetAuditAccessOk() (*string, bool) {
	if o == nil || IsNil(o.AuditAccess) {
		return nil, false
	}
	return o.AuditAccess, true
}

// HasAuditAccess returns a boolean if a field has been set.
func (o *CreateRole) HasAuditAccess() bool {
	if o != nil && !IsNil(o.AuditAccess) {
		return true
	}

	return false
}

// SetAuditAccess gets a reference to the given string and assigns it to the AuditAccess field.
func (o *CreateRole) SetAuditAccess(v string) {
	o.AuditAccess = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateRole) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateRole) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateRole) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateRole) SetDescription(v string) {
	o.Description = &v
}

// GetGwAnalyticsAccess returns the GwAnalyticsAccess field value if set, zero value otherwise.
func (o *CreateRole) GetGwAnalyticsAccess() string {
	if o == nil || IsNil(o.GwAnalyticsAccess) {
		var ret string
		return ret
	}
	return *o.GwAnalyticsAccess
}

// GetGwAnalyticsAccessOk returns a tuple with the GwAnalyticsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetGwAnalyticsAccessOk() (*string, bool) {
	if o == nil || IsNil(o.GwAnalyticsAccess) {
		return nil, false
	}
	return o.GwAnalyticsAccess, true
}

// HasGwAnalyticsAccess returns a boolean if a field has been set.
func (o *CreateRole) HasGwAnalyticsAccess() bool {
	if o != nil && !IsNil(o.GwAnalyticsAccess) {
		return true
	}

	return false
}

// SetGwAnalyticsAccess gets a reference to the given string and assigns it to the GwAnalyticsAccess field.
func (o *CreateRole) SetGwAnalyticsAccess(v string) {
	o.GwAnalyticsAccess = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateRole) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateRole) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateRole) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *CreateRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRole) SetName(v string) {
	o.Name = v
}

// GetSraReportsAccess returns the SraReportsAccess field value if set, zero value otherwise.
func (o *CreateRole) GetSraReportsAccess() string {
	if o == nil || IsNil(o.SraReportsAccess) {
		var ret string
		return ret
	}
	return *o.SraReportsAccess
}

// GetSraReportsAccessOk returns a tuple with the SraReportsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetSraReportsAccessOk() (*string, bool) {
	if o == nil || IsNil(o.SraReportsAccess) {
		return nil, false
	}
	return o.SraReportsAccess, true
}

// HasSraReportsAccess returns a boolean if a field has been set.
func (o *CreateRole) HasSraReportsAccess() bool {
	if o != nil && !IsNil(o.SraReportsAccess) {
		return true
	}

	return false
}

// SetSraReportsAccess gets a reference to the given string and assigns it to the SraReportsAccess field.
func (o *CreateRole) SetSraReportsAccess(v string) {
	o.SraReportsAccess = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateRole) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateRole) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateRole) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateRole) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateRole) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateRole) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnalyticsAccess) {
		toSerialize["analytics-access"] = o.AnalyticsAccess
	}
	if !IsNil(o.AuditAccess) {
		toSerialize["audit-access"] = o.AuditAccess
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GwAnalyticsAccess) {
		toSerialize["gw-analytics-access"] = o.GwAnalyticsAccess
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.SraReportsAccess) {
		toSerialize["sra-reports-access"] = o.SraReportsAccess
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableCreateRole struct {
	value *CreateRole
	isSet bool
}

func (v NullableCreateRole) Get() *CreateRole {
	return v.value
}

func (v *NullableCreateRole) Set(val *CreateRole) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRole) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRole(val *CreateRole) *NullableCreateRole {
	return &NullableCreateRole{value: val, isSet: true}
}

func (v NullableCreateRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


