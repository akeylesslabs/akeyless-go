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
	"time"
)

// AllowedAccess struct for AllowedAccess
type AllowedAccess struct {
	AccessId *string `json:"access_id,omitempty"`
	AccessType *string `json:"access_type,omitempty"`
	ClusterId *int64 `json:"cluster_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Editable *bool `json:"editable,omitempty"`
	Error *string `json:"error,omitempty"`
	Id *int64 `json:"id,omitempty"`
	IsValid *bool `json:"is_valid,omitempty"`
	Name *string `json:"name,omitempty"`
	Permissions *[]string `json:"permissions,omitempty"`
	SubClaims *map[string][]string `json:"sub_claims,omitempty"`
	SubClaimsCaseInsensitive *bool `json:"sub_claims_case_insensitive,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewAllowedAccess instantiates a new AllowedAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedAccess() *AllowedAccess {
	this := AllowedAccess{}
	return &this
}

// NewAllowedAccessWithDefaults instantiates a new AllowedAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedAccessWithDefaults() *AllowedAccess {
	this := AllowedAccess{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *AllowedAccess) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *AllowedAccess) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *AllowedAccess) SetAccessId(v string) {
	o.AccessId = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *AllowedAccess) GetAccessType() string {
	if o == nil || o.AccessType == nil {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetAccessTypeOk() (*string, bool) {
	if o == nil || o.AccessType == nil {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *AllowedAccess) HasAccessType() bool {
	if o != nil && o.AccessType != nil {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *AllowedAccess) SetAccessType(v string) {
	o.AccessType = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *AllowedAccess) GetClusterId() int64 {
	if o == nil || o.ClusterId == nil {
		var ret int64
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetClusterIdOk() (*int64, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *AllowedAccess) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given int64 and assigns it to the ClusterId field.
func (o *AllowedAccess) SetClusterId(v int64) {
	o.ClusterId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AllowedAccess) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AllowedAccess) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AllowedAccess) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AllowedAccess) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AllowedAccess) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AllowedAccess) SetDescription(v string) {
	o.Description = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *AllowedAccess) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *AllowedAccess) HasEditable() bool {
	if o != nil && o.Editable != nil {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *AllowedAccess) SetEditable(v bool) {
	o.Editable = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AllowedAccess) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AllowedAccess) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *AllowedAccess) SetError(v string) {
	o.Error = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AllowedAccess) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AllowedAccess) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AllowedAccess) SetId(v int64) {
	o.Id = &v
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *AllowedAccess) GetIsValid() bool {
	if o == nil || o.IsValid == nil {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetIsValidOk() (*bool, bool) {
	if o == nil || o.IsValid == nil {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *AllowedAccess) HasIsValid() bool {
	if o != nil && o.IsValid != nil {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *AllowedAccess) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AllowedAccess) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AllowedAccess) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AllowedAccess) SetName(v string) {
	o.Name = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *AllowedAccess) GetPermissions() []string {
	if o == nil || o.Permissions == nil {
		var ret []string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetPermissionsOk() (*[]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *AllowedAccess) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *AllowedAccess) SetPermissions(v []string) {
	o.Permissions = &v
}

// GetSubClaims returns the SubClaims field value if set, zero value otherwise.
func (o *AllowedAccess) GetSubClaims() map[string][]string {
	if o == nil || o.SubClaims == nil {
		var ret map[string][]string
		return ret
	}
	return *o.SubClaims
}

// GetSubClaimsOk returns a tuple with the SubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetSubClaimsOk() (*map[string][]string, bool) {
	if o == nil || o.SubClaims == nil {
		return nil, false
	}
	return o.SubClaims, true
}

// HasSubClaims returns a boolean if a field has been set.
func (o *AllowedAccess) HasSubClaims() bool {
	if o != nil && o.SubClaims != nil {
		return true
	}

	return false
}

// SetSubClaims gets a reference to the given map[string][]string and assigns it to the SubClaims field.
func (o *AllowedAccess) SetSubClaims(v map[string][]string) {
	o.SubClaims = &v
}

// GetSubClaimsCaseInsensitive returns the SubClaimsCaseInsensitive field value if set, zero value otherwise.
func (o *AllowedAccess) GetSubClaimsCaseInsensitive() bool {
	if o == nil || o.SubClaimsCaseInsensitive == nil {
		var ret bool
		return ret
	}
	return *o.SubClaimsCaseInsensitive
}

// GetSubClaimsCaseInsensitiveOk returns a tuple with the SubClaimsCaseInsensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetSubClaimsCaseInsensitiveOk() (*bool, bool) {
	if o == nil || o.SubClaimsCaseInsensitive == nil {
		return nil, false
	}
	return o.SubClaimsCaseInsensitive, true
}

// HasSubClaimsCaseInsensitive returns a boolean if a field has been set.
func (o *AllowedAccess) HasSubClaimsCaseInsensitive() bool {
	if o != nil && o.SubClaimsCaseInsensitive != nil {
		return true
	}

	return false
}

// SetSubClaimsCaseInsensitive gets a reference to the given bool and assigns it to the SubClaimsCaseInsensitive field.
func (o *AllowedAccess) SetSubClaimsCaseInsensitive(v bool) {
	o.SubClaimsCaseInsensitive = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AllowedAccess) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AllowedAccess) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AllowedAccess) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o AllowedAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access_id"] = o.AccessId
	}
	if o.AccessType != nil {
		toSerialize["access_type"] = o.AccessType
	}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Editable != nil {
		toSerialize["editable"] = o.Editable
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsValid != nil {
		toSerialize["is_valid"] = o.IsValid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.SubClaims != nil {
		toSerialize["sub_claims"] = o.SubClaims
	}
	if o.SubClaimsCaseInsensitive != nil {
		toSerialize["sub_claims_case_insensitive"] = o.SubClaimsCaseInsensitive
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAllowedAccess struct {
	value *AllowedAccess
	isSet bool
}

func (v NullableAllowedAccess) Get() *AllowedAccess {
	return v.value
}

func (v *NullableAllowedAccess) Set(val *AllowedAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedAccess(val *AllowedAccess) *NullableAllowedAccess {
	return &NullableAllowedAccess{value: val, isSet: true}
}

func (v NullableAllowedAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


