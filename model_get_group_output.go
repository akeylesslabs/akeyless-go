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

// GetGroupOutput struct for GetGroupOutput
type GetGroupOutput struct {
	AccountId *string `json:"account_id,omitempty"`
	CreationDate *time.Time `json:"creation_date,omitempty"`
	Description *string `json:"description,omitempty"`
	GroupAlias *string `json:"group_alias,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	GroupName *string `json:"group_name,omitempty"`
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	UserAssignments *[]AccessPermissionAssignment `json:"user_assignments,omitempty"`
}

// NewGetGroupOutput instantiates a new GetGroupOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGroupOutput() *GetGroupOutput {
	this := GetGroupOutput{}
	return &this
}

// NewGetGroupOutputWithDefaults instantiates a new GetGroupOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGroupOutputWithDefaults() *GetGroupOutput {
	this := GetGroupOutput{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *GetGroupOutput) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupOutput) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *GetGroupOutput) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *GetGroupOutput) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *GetGroupOutput) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupOutput) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *GetGroupOutput) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *GetGroupOutput) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetGroupOutput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupOutput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetGroupOutput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetGroupOutput) SetDescription(v string) {
	o.Description = &v
}

// GetGroupAlias returns the GroupAlias field value if set, zero value otherwise.
func (o *GetGroupOutput) GetGroupAlias() string {
	if o == nil || o.GroupAlias == nil {
		var ret string
		return ret
	}
	return *o.GroupAlias
}

// GetGroupAliasOk returns a tuple with the GroupAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupOutput) GetGroupAliasOk() (*string, bool) {
	if o == nil || o.GroupAlias == nil {
		return nil, false
	}
	return o.GroupAlias, true
}

// HasGroupAlias returns a boolean if a field has been set.
func (o *GetGroupOutput) HasGroupAlias() bool {
	if o != nil && o.GroupAlias != nil {
		return true
	}

	return false
}

// SetGroupAlias gets a reference to the given string and assigns it to the GroupAlias field.
func (o *GetGroupOutput) SetGroupAlias(v string) {
	o.GroupAlias = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GetGroupOutput) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupOutput) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GetGroupOutput) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GetGroupOutput) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *GetGroupOutput) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupOutput) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *GetGroupOutput) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *GetGroupOutput) SetGroupName(v string) {
	o.GroupName = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *GetGroupOutput) GetModificationDate() time.Time {
	if o == nil || o.ModificationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupOutput) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || o.ModificationDate == nil {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *GetGroupOutput) HasModificationDate() bool {
	if o != nil && o.ModificationDate != nil {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *GetGroupOutput) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

// GetUserAssignments returns the UserAssignments field value if set, zero value otherwise.
func (o *GetGroupOutput) GetUserAssignments() []AccessPermissionAssignment {
	if o == nil || o.UserAssignments == nil {
		var ret []AccessPermissionAssignment
		return ret
	}
	return *o.UserAssignments
}

// GetUserAssignmentsOk returns a tuple with the UserAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupOutput) GetUserAssignmentsOk() (*[]AccessPermissionAssignment, bool) {
	if o == nil || o.UserAssignments == nil {
		return nil, false
	}
	return o.UserAssignments, true
}

// HasUserAssignments returns a boolean if a field has been set.
func (o *GetGroupOutput) HasUserAssignments() bool {
	if o != nil && o.UserAssignments != nil {
		return true
	}

	return false
}

// SetUserAssignments gets a reference to the given []AccessPermissionAssignment and assigns it to the UserAssignments field.
func (o *GetGroupOutput) SetUserAssignments(v []AccessPermissionAssignment) {
	o.UserAssignments = &v
}

func (o GetGroupOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GroupAlias != nil {
		toSerialize["group_alias"] = o.GroupAlias
	}
	if o.GroupId != nil {
		toSerialize["group_id"] = o.GroupId
	}
	if o.GroupName != nil {
		toSerialize["group_name"] = o.GroupName
	}
	if o.ModificationDate != nil {
		toSerialize["modification_date"] = o.ModificationDate
	}
	if o.UserAssignments != nil {
		toSerialize["user_assignments"] = o.UserAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableGetGroupOutput struct {
	value *GetGroupOutput
	isSet bool
}

func (v NullableGetGroupOutput) Get() *GetGroupOutput {
	return v.value
}

func (v *NullableGetGroupOutput) Set(val *GetGroupOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGroupOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGroupOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGroupOutput(val *GetGroupOutput) *NullableGetGroupOutput {
	return &NullableGetGroupOutput{value: val, isSet: true}
}

func (v NullableGetGroupOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGroupOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


