/*
 * includes all operations supported by Akeyless.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GetRSAPublic getRSAPublic is a command that obtains the public key from a specific RSA private key.
type GetRSAPublic struct {
	// Name of RSA key to extract the public key from
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGetRSAPublic instantiates a new GetRSAPublic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRSAPublic(name string, ) *GetRSAPublic {
	this := GetRSAPublic{}
	this.Name = name
	return &this
}

// NewGetRSAPublicWithDefaults instantiates a new GetRSAPublic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRSAPublicWithDefaults() *GetRSAPublic {
	this := GetRSAPublic{}
	return &this
}

// GetName returns the Name field value
func (o *GetRSAPublic) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetRSAPublic) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetRSAPublic) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetRSAPublic) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRSAPublic) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetRSAPublic) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetRSAPublic) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GetRSAPublic) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRSAPublic) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GetRSAPublic) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GetRSAPublic) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GetRSAPublic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGetRSAPublic struct {
	value *GetRSAPublic
	isSet bool
}

func (v NullableGetRSAPublic) Get() *GetRSAPublic {
	return v.value
}

func (v *NullableGetRSAPublic) Set(val *GetRSAPublic) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRSAPublic) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRSAPublic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRSAPublic(val *GetRSAPublic) *NullableGetRSAPublic {
	return &NullableGetRSAPublic{value: val, isSet: true}
}

func (v NullableGetRSAPublic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRSAPublic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


