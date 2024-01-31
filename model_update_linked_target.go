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

// UpdateLinkedTarget struct for UpdateLinkedTarget
type UpdateLinkedTarget struct {
	// A comma seperated list of new server hosts and server descriptions joined by semicolon ';' that will be added to the Linked Target hosts.
	AddHosts *string `json:"add-hosts,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// A comma seperated list of server hosts and server descriptions joined by semicolon ';' (i.e. 'server-dev.com;My Dev server,server-prod.com;My Prod server description')
	Hosts *string `json:"hosts,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Whether to keep previous version [true/false]. If not set, use default according to account settings
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// Linked Target name
	Name string `json:"name"`
	// New Linked Target name
	NewName *string `json:"new-name,omitempty"`
	// The parent Target name
	ParentTargetName *string `json:"parent-target-name,omitempty"`
	// Comma separated list of existing hosts that will be removed from Linked Target hosts.
	RmHosts *string `json:"rm-hosts,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Specifies the hosts type, relevant only when working without parent target
	Type *string `json:"type,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateLinkedTarget instantiates a new UpdateLinkedTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLinkedTarget(name string, ) *UpdateLinkedTarget {
	this := UpdateLinkedTarget{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewUpdateLinkedTargetWithDefaults instantiates a new UpdateLinkedTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLinkedTargetWithDefaults() *UpdateLinkedTarget {
	this := UpdateLinkedTarget{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAddHosts returns the AddHosts field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetAddHosts() string {
	if o == nil || o.AddHosts == nil {
		var ret string
		return ret
	}
	return *o.AddHosts
}

// GetAddHostsOk returns a tuple with the AddHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetAddHostsOk() (*string, bool) {
	if o == nil || o.AddHosts == nil {
		return nil, false
	}
	return o.AddHosts, true
}

// HasAddHosts returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasAddHosts() bool {
	if o != nil && o.AddHosts != nil {
		return true
	}

	return false
}

// SetAddHosts gets a reference to the given string and assigns it to the AddHosts field.
func (o *UpdateLinkedTarget) SetAddHosts(v string) {
	o.AddHosts = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateLinkedTarget) SetDescription(v string) {
	o.Description = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetHosts() string {
	if o == nil || o.Hosts == nil {
		var ret string
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetHostsOk() (*string, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given string and assigns it to the Hosts field.
func (o *UpdateLinkedTarget) SetHosts(v string) {
	o.Hosts = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateLinkedTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateLinkedTarget) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetName returns the Name field value
func (o *UpdateLinkedTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateLinkedTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateLinkedTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetParentTargetName returns the ParentTargetName field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetParentTargetName() string {
	if o == nil || o.ParentTargetName == nil {
		var ret string
		return ret
	}
	return *o.ParentTargetName
}

// GetParentTargetNameOk returns a tuple with the ParentTargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetParentTargetNameOk() (*string, bool) {
	if o == nil || o.ParentTargetName == nil {
		return nil, false
	}
	return o.ParentTargetName, true
}

// HasParentTargetName returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasParentTargetName() bool {
	if o != nil && o.ParentTargetName != nil {
		return true
	}

	return false
}

// SetParentTargetName gets a reference to the given string and assigns it to the ParentTargetName field.
func (o *UpdateLinkedTarget) SetParentTargetName(v string) {
	o.ParentTargetName = &v
}

// GetRmHosts returns the RmHosts field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetRmHosts() string {
	if o == nil || o.RmHosts == nil {
		var ret string
		return ret
	}
	return *o.RmHosts
}

// GetRmHostsOk returns a tuple with the RmHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetRmHostsOk() (*string, bool) {
	if o == nil || o.RmHosts == nil {
		return nil, false
	}
	return o.RmHosts, true
}

// HasRmHosts returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasRmHosts() bool {
	if o != nil && o.RmHosts != nil {
		return true
	}

	return false
}

// SetRmHosts gets a reference to the given string and assigns it to the RmHosts field.
func (o *UpdateLinkedTarget) SetRmHosts(v string) {
	o.RmHosts = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateLinkedTarget) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateLinkedTarget) SetType(v string) {
	o.Type = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateLinkedTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLinkedTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateLinkedTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateLinkedTarget) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateLinkedTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddHosts != nil {
		toSerialize["add-hosts"] = o.AddHosts
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.ParentTargetName != nil {
		toSerialize["parent-target-name"] = o.ParentTargetName
	}
	if o.RmHosts != nil {
		toSerialize["rm-hosts"] = o.RmHosts
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

type NullableUpdateLinkedTarget struct {
	value *UpdateLinkedTarget
	isSet bool
}

func (v NullableUpdateLinkedTarget) Get() *UpdateLinkedTarget {
	return v.value
}

func (v *NullableUpdateLinkedTarget) Set(val *UpdateLinkedTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLinkedTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLinkedTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLinkedTarget(val *UpdateLinkedTarget) *NullableUpdateLinkedTarget {
	return &NullableUpdateLinkedTarget{value: val, isSet: true}
}

func (v NullableUpdateLinkedTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLinkedTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


