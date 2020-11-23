/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// DeleteRoleAssociation deleteRoleAssociation is a command that deletes an association between role and auth method.
type DeleteRoleAssociation struct {
	// The association id to be deleted
	AssocId string `json:"assoc-id"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDeleteRoleAssociation instantiates a new DeleteRoleAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRoleAssociation(assocId string, ) *DeleteRoleAssociation {
	this := DeleteRoleAssociation{}
	this.AssocId = assocId
	return &this
}

// NewDeleteRoleAssociationWithDefaults instantiates a new DeleteRoleAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRoleAssociationWithDefaults() *DeleteRoleAssociation {
	this := DeleteRoleAssociation{}
	return &this
}

// GetAssocId returns the AssocId field value
func (o *DeleteRoleAssociation) GetAssocId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AssocId
}

// GetAssocIdOk returns a tuple with the AssocId field value
// and a boolean to check if the value has been set.
func (o *DeleteRoleAssociation) GetAssocIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssocId, true
}

// SetAssocId sets field value
func (o *DeleteRoleAssociation) SetAssocId(v string) {
	o.AssocId = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DeleteRoleAssociation) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRoleAssociation) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DeleteRoleAssociation) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DeleteRoleAssociation) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DeleteRoleAssociation) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRoleAssociation) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DeleteRoleAssociation) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DeleteRoleAssociation) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DeleteRoleAssociation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assoc-id"] = o.AssocId
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteRoleAssociation struct {
	value *DeleteRoleAssociation
	isSet bool
}

func (v NullableDeleteRoleAssociation) Get() *DeleteRoleAssociation {
	return v.value
}

func (v *NullableDeleteRoleAssociation) Set(val *DeleteRoleAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRoleAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRoleAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRoleAssociation(val *DeleteRoleAssociation) *NullableDeleteRoleAssociation {
	return &NullableDeleteRoleAssociation{value: val, isSet: true}
}

func (v NullableDeleteRoleAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRoleAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


