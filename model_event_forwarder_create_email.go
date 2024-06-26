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

// EventForwarderCreateEmail eventForwarderCreateEmail is a command that creates email event forwarder
type EventForwarderCreateEmail struct {
	// Auth Method Event sources
	AuthMethodsEventSourceLocations *[]string `json:"auth-methods-event-source-locations,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// A comma seperated list of email addresses to send event to
	EmailTo *string `json:"email-to,omitempty"`
	// List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated]
	EventTypes *[]string `json:"event-types,omitempty"`
	// Rate of periodic runner repetition in hours
	Every *string `json:"every,omitempty"`
	// Event sources
	GatewaysEventSourceLocations []string `json:"gateways-event-source-locations"`
	// Items Event sources
	ItemsEventSourceLocations *[]string `json:"items-event-source-locations,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// EventForwarder name
	Name string `json:"name"`
	// Override Akeyless default URL with your Gateway url (port 18888)
	OverrideUrl *string `json:"override-url,omitempty"`
	RunnerType string `json:"runner-type"`
	// Targets Event sources
	TargetsEventSourceLocations *[]string `json:"targets-event-source-locations,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewEventForwarderCreateEmail instantiates a new EventForwarderCreateEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventForwarderCreateEmail(gatewaysEventSourceLocations []string, name string, runnerType string, ) *EventForwarderCreateEmail {
	this := EventForwarderCreateEmail{}
	this.GatewaysEventSourceLocations = gatewaysEventSourceLocations
	var json bool = false
	this.Json = &json
	this.Name = name
	this.RunnerType = runnerType
	return &this
}

// NewEventForwarderCreateEmailWithDefaults instantiates a new EventForwarderCreateEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventForwarderCreateEmailWithDefaults() *EventForwarderCreateEmail {
	this := EventForwarderCreateEmail{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetAuthMethodsEventSourceLocations() []string {
	if o == nil || o.AuthMethodsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.AuthMethodsEventSourceLocations
}

// GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.AuthMethodsEventSourceLocations == nil {
		return nil, false
	}
	return o.AuthMethodsEventSourceLocations, true
}

// HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasAuthMethodsEventSourceLocations() bool {
	if o != nil && o.AuthMethodsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetAuthMethodsEventSourceLocations gets a reference to the given []string and assigns it to the AuthMethodsEventSourceLocations field.
func (o *EventForwarderCreateEmail) SetAuthMethodsEventSourceLocations(v []string) {
	o.AuthMethodsEventSourceLocations = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventForwarderCreateEmail) SetDescription(v string) {
	o.Description = &v
}

// GetEmailTo returns the EmailTo field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetEmailTo() string {
	if o == nil || o.EmailTo == nil {
		var ret string
		return ret
	}
	return *o.EmailTo
}

// GetEmailToOk returns a tuple with the EmailTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetEmailToOk() (*string, bool) {
	if o == nil || o.EmailTo == nil {
		return nil, false
	}
	return o.EmailTo, true
}

// HasEmailTo returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasEmailTo() bool {
	if o != nil && o.EmailTo != nil {
		return true
	}

	return false
}

// SetEmailTo gets a reference to the given string and assigns it to the EmailTo field.
func (o *EventForwarderCreateEmail) SetEmailTo(v string) {
	o.EmailTo = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetEventTypesOk() (*[]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *EventForwarderCreateEmail) SetEventTypes(v []string) {
	o.EventTypes = &v
}

// GetEvery returns the Every field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetEvery() string {
	if o == nil || o.Every == nil {
		var ret string
		return ret
	}
	return *o.Every
}

// GetEveryOk returns a tuple with the Every field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetEveryOk() (*string, bool) {
	if o == nil || o.Every == nil {
		return nil, false
	}
	return o.Every, true
}

// HasEvery returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasEvery() bool {
	if o != nil && o.Every != nil {
		return true
	}

	return false
}

// SetEvery gets a reference to the given string and assigns it to the Every field.
func (o *EventForwarderCreateEmail) SetEvery(v string) {
	o.Every = &v
}

// GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field value
func (o *EventForwarderCreateEmail) GetGatewaysEventSourceLocations() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.GatewaysEventSourceLocations
}

// GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetGatewaysEventSourceLocationsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GatewaysEventSourceLocations, true
}

// SetGatewaysEventSourceLocations sets field value
func (o *EventForwarderCreateEmail) SetGatewaysEventSourceLocations(v []string) {
	o.GatewaysEventSourceLocations = v
}

// GetItemsEventSourceLocations returns the ItemsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetItemsEventSourceLocations() []string {
	if o == nil || o.ItemsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.ItemsEventSourceLocations
}

// GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetItemsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.ItemsEventSourceLocations == nil {
		return nil, false
	}
	return o.ItemsEventSourceLocations, true
}

// HasItemsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasItemsEventSourceLocations() bool {
	if o != nil && o.ItemsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetItemsEventSourceLocations gets a reference to the given []string and assigns it to the ItemsEventSourceLocations field.
func (o *EventForwarderCreateEmail) SetItemsEventSourceLocations(v []string) {
	o.ItemsEventSourceLocations = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *EventForwarderCreateEmail) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EventForwarderCreateEmail) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *EventForwarderCreateEmail) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventForwarderCreateEmail) SetName(v string) {
	o.Name = v
}

// GetOverrideUrl returns the OverrideUrl field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetOverrideUrl() string {
	if o == nil || o.OverrideUrl == nil {
		var ret string
		return ret
	}
	return *o.OverrideUrl
}

// GetOverrideUrlOk returns a tuple with the OverrideUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetOverrideUrlOk() (*string, bool) {
	if o == nil || o.OverrideUrl == nil {
		return nil, false
	}
	return o.OverrideUrl, true
}

// HasOverrideUrl returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasOverrideUrl() bool {
	if o != nil && o.OverrideUrl != nil {
		return true
	}

	return false
}

// SetOverrideUrl gets a reference to the given string and assigns it to the OverrideUrl field.
func (o *EventForwarderCreateEmail) SetOverrideUrl(v string) {
	o.OverrideUrl = &v
}

// GetRunnerType returns the RunnerType field value
func (o *EventForwarderCreateEmail) GetRunnerType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RunnerType
}

// GetRunnerTypeOk returns a tuple with the RunnerType field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetRunnerTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RunnerType, true
}

// SetRunnerType sets field value
func (o *EventForwarderCreateEmail) SetRunnerType(v string) {
	o.RunnerType = v
}

// GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetTargetsEventSourceLocations() []string {
	if o == nil || o.TargetsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.TargetsEventSourceLocations
}

// GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetTargetsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.TargetsEventSourceLocations == nil {
		return nil, false
	}
	return o.TargetsEventSourceLocations, true
}

// HasTargetsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasTargetsEventSourceLocations() bool {
	if o != nil && o.TargetsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetTargetsEventSourceLocations gets a reference to the given []string and assigns it to the TargetsEventSourceLocations field.
func (o *EventForwarderCreateEmail) SetTargetsEventSourceLocations(v []string) {
	o.TargetsEventSourceLocations = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EventForwarderCreateEmail) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EventForwarderCreateEmail) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateEmail) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EventForwarderCreateEmail) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EventForwarderCreateEmail) SetUidToken(v string) {
	o.UidToken = &v
}

func (o EventForwarderCreateEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthMethodsEventSourceLocations != nil {
		toSerialize["auth-methods-event-source-locations"] = o.AuthMethodsEventSourceLocations
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EmailTo != nil {
		toSerialize["email-to"] = o.EmailTo
	}
	if o.EventTypes != nil {
		toSerialize["event-types"] = o.EventTypes
	}
	if o.Every != nil {
		toSerialize["every"] = o.Every
	}
	if true {
		toSerialize["gateways-event-source-locations"] = o.GatewaysEventSourceLocations
	}
	if o.ItemsEventSourceLocations != nil {
		toSerialize["items-event-source-locations"] = o.ItemsEventSourceLocations
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OverrideUrl != nil {
		toSerialize["override-url"] = o.OverrideUrl
	}
	if true {
		toSerialize["runner-type"] = o.RunnerType
	}
	if o.TargetsEventSourceLocations != nil {
		toSerialize["targets-event-source-locations"] = o.TargetsEventSourceLocations
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableEventForwarderCreateEmail struct {
	value *EventForwarderCreateEmail
	isSet bool
}

func (v NullableEventForwarderCreateEmail) Get() *EventForwarderCreateEmail {
	return v.value
}

func (v *NullableEventForwarderCreateEmail) Set(val *EventForwarderCreateEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableEventForwarderCreateEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableEventForwarderCreateEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventForwarderCreateEmail(val *EventForwarderCreateEmail) *NullableEventForwarderCreateEmail {
	return &NullableEventForwarderCreateEmail{value: val, isSet: true}
}

func (v NullableEventForwarderCreateEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventForwarderCreateEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


