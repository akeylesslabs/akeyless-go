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

// CreateLinkedTarget struct for CreateLinkedTarget
type CreateLinkedTarget struct {
	// Description of the object
	Description *string `json:"description,omitempty"`
	// A comma seperated list of server hosts and server descriptions joined by semicolon ';' (i.e. 'server-dev.com;My Dev server,server-prod.com;My Prod server description')
	Hosts *string `json:"hosts,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Target name
	Name string `json:"name"`
	// The parent Target name
	ParentTargetName *string `json:"parent-target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Specifies the hosts type, relevant only when working without parent target
	Type *string `json:"type,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateLinkedTarget instantiates a new CreateLinkedTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLinkedTarget(name string, ) *CreateLinkedTarget {
	this := CreateLinkedTarget{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewCreateLinkedTargetWithDefaults instantiates a new CreateLinkedTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLinkedTargetWithDefaults() *CreateLinkedTarget {
	this := CreateLinkedTarget{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateLinkedTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkedTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateLinkedTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateLinkedTarget) SetDescription(v string) {
	o.Description = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *CreateLinkedTarget) GetHosts() string {
	if o == nil || o.Hosts == nil {
		var ret string
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkedTarget) GetHostsOk() (*string, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *CreateLinkedTarget) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given string and assigns it to the Hosts field.
func (o *CreateLinkedTarget) SetHosts(v string) {
	o.Hosts = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateLinkedTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkedTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateLinkedTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateLinkedTarget) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *CreateLinkedTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateLinkedTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateLinkedTarget) SetName(v string) {
	o.Name = v
}

// GetParentTargetName returns the ParentTargetName field value if set, zero value otherwise.
func (o *CreateLinkedTarget) GetParentTargetName() string {
	if o == nil || o.ParentTargetName == nil {
		var ret string
		return ret
	}
	return *o.ParentTargetName
}

// GetParentTargetNameOk returns a tuple with the ParentTargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkedTarget) GetParentTargetNameOk() (*string, bool) {
	if o == nil || o.ParentTargetName == nil {
		return nil, false
	}
	return o.ParentTargetName, true
}

// HasParentTargetName returns a boolean if a field has been set.
func (o *CreateLinkedTarget) HasParentTargetName() bool {
	if o != nil && o.ParentTargetName != nil {
		return true
	}

	return false
}

// SetParentTargetName gets a reference to the given string and assigns it to the ParentTargetName field.
func (o *CreateLinkedTarget) SetParentTargetName(v string) {
	o.ParentTargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateLinkedTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkedTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateLinkedTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateLinkedTarget) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateLinkedTarget) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkedTarget) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateLinkedTarget) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateLinkedTarget) SetType(v string) {
	o.Type = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateLinkedTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkedTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateLinkedTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateLinkedTarget) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateLinkedTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ParentTargetName != nil {
		toSerialize["parent-target-name"] = o.ParentTargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLinkedTarget struct {
	value *CreateLinkedTarget
	isSet bool
}

func (v NullableCreateLinkedTarget) Get() *CreateLinkedTarget {
	return v.value
}

func (v *NullableCreateLinkedTarget) Set(val *CreateLinkedTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLinkedTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLinkedTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLinkedTarget(val *CreateLinkedTarget) *NullableCreateLinkedTarget {
	return &NullableCreateLinkedTarget{value: val, isSet: true}
}

func (v NullableCreateLinkedTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLinkedTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


