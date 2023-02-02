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

// UpdateEventForwarder struct for UpdateEventForwarder
type UpdateEventForwarder struct {
	// Workstation Admin Name
	AdminName *string `json:"admin-name,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// A comma seperated list of email addresses to send event to (relevant only for \\\"email\\\" Event Forwarder)
	EmailTo *string `json:"email-to,omitempty"`
	// Enable
	Enable *string `json:"enable,omitempty"`
	// Event sources
	EventSourceLocations *[]string `json:"event-source-locations,omitempty"`
	// Event types
	EventTypes *[]string `json:"event-types,omitempty"`
	// Workstation Host
	Host *string `json:"host,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// EventForwarder name
	Name string `json:"name"`
	// Deprecated - use description
	NewComment *string `json:"new-comment,omitempty"`
	// New EventForwarder name
	NewName *string `json:"new-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateEventForwarder instantiates a new UpdateEventForwarder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEventForwarder(name string, ) *UpdateEventForwarder {
	this := UpdateEventForwarder{}
	var description string = "default_comment"
	this.Description = &description
	this.Name = name
	var newComment string = "default_comment"
	this.NewComment = &newComment
	return &this
}

// NewUpdateEventForwarderWithDefaults instantiates a new UpdateEventForwarder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEventForwarderWithDefaults() *UpdateEventForwarder {
	this := UpdateEventForwarder{}
	var description string = "default_comment"
	this.Description = &description
	var newComment string = "default_comment"
	this.NewComment = &newComment
	return &this
}

// GetAdminName returns the AdminName field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetAdminName() string {
	if o == nil || o.AdminName == nil {
		var ret string
		return ret
	}
	return *o.AdminName
}

// GetAdminNameOk returns a tuple with the AdminName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetAdminNameOk() (*string, bool) {
	if o == nil || o.AdminName == nil {
		return nil, false
	}
	return o.AdminName, true
}

// HasAdminName returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasAdminName() bool {
	if o != nil && o.AdminName != nil {
		return true
	}

	return false
}

// SetAdminName gets a reference to the given string and assigns it to the AdminName field.
func (o *UpdateEventForwarder) SetAdminName(v string) {
	o.AdminName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateEventForwarder) SetDescription(v string) {
	o.Description = &v
}

// GetEmailTo returns the EmailTo field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetEmailTo() string {
	if o == nil || o.EmailTo == nil {
		var ret string
		return ret
	}
	return *o.EmailTo
}

// GetEmailToOk returns a tuple with the EmailTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetEmailToOk() (*string, bool) {
	if o == nil || o.EmailTo == nil {
		return nil, false
	}
	return o.EmailTo, true
}

// HasEmailTo returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasEmailTo() bool {
	if o != nil && o.EmailTo != nil {
		return true
	}

	return false
}

// SetEmailTo gets a reference to the given string and assigns it to the EmailTo field.
func (o *UpdateEventForwarder) SetEmailTo(v string) {
	o.EmailTo = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetEnable() string {
	if o == nil || o.Enable == nil {
		var ret string
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetEnableOk() (*string, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given string and assigns it to the Enable field.
func (o *UpdateEventForwarder) SetEnable(v string) {
	o.Enable = &v
}

// GetEventSourceLocations returns the EventSourceLocations field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetEventSourceLocations() []string {
	if o == nil || o.EventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.EventSourceLocations
}

// GetEventSourceLocationsOk returns a tuple with the EventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.EventSourceLocations == nil {
		return nil, false
	}
	return o.EventSourceLocations, true
}

// HasEventSourceLocations returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasEventSourceLocations() bool {
	if o != nil && o.EventSourceLocations != nil {
		return true
	}

	return false
}

// SetEventSourceLocations gets a reference to the given []string and assigns it to the EventSourceLocations field.
func (o *UpdateEventForwarder) SetEventSourceLocations(v []string) {
	o.EventSourceLocations = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetEventTypesOk() (*[]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *UpdateEventForwarder) SetEventTypes(v []string) {
	o.EventTypes = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UpdateEventForwarder) SetHost(v string) {
	o.Host = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateEventForwarder) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *UpdateEventForwarder) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateEventForwarder) SetName(v string) {
	o.Name = v
}

// GetNewComment returns the NewComment field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetNewComment() string {
	if o == nil || o.NewComment == nil {
		var ret string
		return ret
	}
	return *o.NewComment
}

// GetNewCommentOk returns a tuple with the NewComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetNewCommentOk() (*string, bool) {
	if o == nil || o.NewComment == nil {
		return nil, false
	}
	return o.NewComment, true
}

// HasNewComment returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasNewComment() bool {
	if o != nil && o.NewComment != nil {
		return true
	}

	return false
}

// SetNewComment gets a reference to the given string and assigns it to the NewComment field.
func (o *UpdateEventForwarder) SetNewComment(v string) {
	o.NewComment = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateEventForwarder) SetNewName(v string) {
	o.NewName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateEventForwarder) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateEventForwarder) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEventForwarder) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateEventForwarder) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateEventForwarder) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateEventForwarder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminName != nil {
		toSerialize["admin-name"] = o.AdminName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EmailTo != nil {
		toSerialize["email-to"] = o.EmailTo
	}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.EventSourceLocations != nil {
		toSerialize["event-source-locations"] = o.EventSourceLocations
	}
	if o.EventTypes != nil {
		toSerialize["event-types"] = o.EventTypes
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
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
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateEventForwarder struct {
	value *UpdateEventForwarder
	isSet bool
}

func (v NullableUpdateEventForwarder) Get() *UpdateEventForwarder {
	return v.value
}

func (v *NullableUpdateEventForwarder) Set(val *UpdateEventForwarder) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEventForwarder) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEventForwarder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEventForwarder(val *UpdateEventForwarder) *NullableUpdateEventForwarder {
	return &NullableUpdateEventForwarder{value: val, isSet: true}
}

func (v NullableUpdateEventForwarder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEventForwarder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


