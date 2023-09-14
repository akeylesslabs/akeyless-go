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

// AccessOrGroupPermissionAssignment struct for AccessOrGroupPermissionAssignment
type AccessOrGroupPermissionAssignment struct {
	AccessId *string `json:"access_id,omitempty"`
	AssignmentType *string `json:"assignment_type,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	SubClaims *map[string][]string `json:"sub_claims,omitempty"`
}

// NewAccessOrGroupPermissionAssignment instantiates a new AccessOrGroupPermissionAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessOrGroupPermissionAssignment() *AccessOrGroupPermissionAssignment {
	this := AccessOrGroupPermissionAssignment{}
	return &this
}

// NewAccessOrGroupPermissionAssignmentWithDefaults instantiates a new AccessOrGroupPermissionAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessOrGroupPermissionAssignmentWithDefaults() *AccessOrGroupPermissionAssignment {
	this := AccessOrGroupPermissionAssignment{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *AccessOrGroupPermissionAssignment) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessOrGroupPermissionAssignment) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *AccessOrGroupPermissionAssignment) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *AccessOrGroupPermissionAssignment) SetAccessId(v string) {
	o.AccessId = &v
}

// GetAssignmentType returns the AssignmentType field value if set, zero value otherwise.
func (o *AccessOrGroupPermissionAssignment) GetAssignmentType() string {
	if o == nil || o.AssignmentType == nil {
		var ret string
		return ret
	}
	return *o.AssignmentType
}

// GetAssignmentTypeOk returns a tuple with the AssignmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessOrGroupPermissionAssignment) GetAssignmentTypeOk() (*string, bool) {
	if o == nil || o.AssignmentType == nil {
		return nil, false
	}
	return o.AssignmentType, true
}

// HasAssignmentType returns a boolean if a field has been set.
func (o *AccessOrGroupPermissionAssignment) HasAssignmentType() bool {
	if o != nil && o.AssignmentType != nil {
		return true
	}

	return false
}

// SetAssignmentType gets a reference to the given string and assigns it to the AssignmentType field.
func (o *AccessOrGroupPermissionAssignment) SetAssignmentType(v string) {
	o.AssignmentType = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AccessOrGroupPermissionAssignment) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessOrGroupPermissionAssignment) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AccessOrGroupPermissionAssignment) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AccessOrGroupPermissionAssignment) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSubClaims returns the SubClaims field value if set, zero value otherwise.
func (o *AccessOrGroupPermissionAssignment) GetSubClaims() map[string][]string {
	if o == nil || o.SubClaims == nil {
		var ret map[string][]string
		return ret
	}
	return *o.SubClaims
}

// GetSubClaimsOk returns a tuple with the SubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessOrGroupPermissionAssignment) GetSubClaimsOk() (*map[string][]string, bool) {
	if o == nil || o.SubClaims == nil {
		return nil, false
	}
	return o.SubClaims, true
}

// HasSubClaims returns a boolean if a field has been set.
func (o *AccessOrGroupPermissionAssignment) HasSubClaims() bool {
	if o != nil && o.SubClaims != nil {
		return true
	}

	return false
}

// SetSubClaims gets a reference to the given map[string][]string and assigns it to the SubClaims field.
func (o *AccessOrGroupPermissionAssignment) SetSubClaims(v map[string][]string) {
	o.SubClaims = &v
}

func (o AccessOrGroupPermissionAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access_id"] = o.AccessId
	}
	if o.AssignmentType != nil {
		toSerialize["assignment_type"] = o.AssignmentType
	}
	if o.GroupId != nil {
		toSerialize["group_id"] = o.GroupId
	}
	if o.SubClaims != nil {
		toSerialize["sub_claims"] = o.SubClaims
	}
	return json.Marshal(toSerialize)
}

type NullableAccessOrGroupPermissionAssignment struct {
	value *AccessOrGroupPermissionAssignment
	isSet bool
}

func (v NullableAccessOrGroupPermissionAssignment) Get() *AccessOrGroupPermissionAssignment {
	return v.value
}

func (v *NullableAccessOrGroupPermissionAssignment) Set(val *AccessOrGroupPermissionAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessOrGroupPermissionAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessOrGroupPermissionAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessOrGroupPermissionAssignment(val *AccessOrGroupPermissionAssignment) *NullableAccessOrGroupPermissionAssignment {
	return &NullableAccessOrGroupPermissionAssignment{value: val, isSet: true}
}

func (v NullableAccessOrGroupPermissionAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessOrGroupPermissionAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


