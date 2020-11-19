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

// UpdateRole struct for UpdateRole
type UpdateRole struct {
	// Role name
	Name string `json:"name"`
	// New comment about the role
	NewComment *string `json:"new-comment,omitempty"`
	// New Role name
	NewName *string `json:"new-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateRole instantiates a new UpdateRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRole(name string, ) *UpdateRole {
	this := UpdateRole{}
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
	var newComment string = "default_comment"
	this.NewComment = &newComment
	return &this
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

func (o UpdateRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewComment != nil {
		toSerialize["new-comment"] = o.NewComment
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
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


