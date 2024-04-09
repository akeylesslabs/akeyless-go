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

// EventForwarderCreateSlack eventForwarderCreateSlack is a command that creates slack event forwarder
type EventForwarderCreateSlack struct {
	// Auth Method Event sources
	AuthMethodsEventSourceLocations *[]string `json:"auth-methods-event-source-locations,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
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
	RunnerType string `json:"runner-type"`
	// Targets Event sources
	TargetsEventSourceLocations *[]string `json:"targets-event-source-locations,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Slack Webhook URL
	Url string `json:"url"`
}

// NewEventForwarderCreateSlack instantiates a new EventForwarderCreateSlack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventForwarderCreateSlack(gatewaysEventSourceLocations []string, name string, runnerType string, url string, ) *EventForwarderCreateSlack {
	this := EventForwarderCreateSlack{}
	this.GatewaysEventSourceLocations = gatewaysEventSourceLocations
	var json bool = false
	this.Json = &json
	this.Name = name
	this.RunnerType = runnerType
	this.Url = url
	return &this
}

// NewEventForwarderCreateSlackWithDefaults instantiates a new EventForwarderCreateSlack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventForwarderCreateSlackWithDefaults() *EventForwarderCreateSlack {
	this := EventForwarderCreateSlack{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetAuthMethodsEventSourceLocations() []string {
	if o == nil || o.AuthMethodsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.AuthMethodsEventSourceLocations
}

// GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.AuthMethodsEventSourceLocations == nil {
		return nil, false
	}
	return o.AuthMethodsEventSourceLocations, true
}

// HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasAuthMethodsEventSourceLocations() bool {
	if o != nil && o.AuthMethodsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetAuthMethodsEventSourceLocations gets a reference to the given []string and assigns it to the AuthMethodsEventSourceLocations field.
func (o *EventForwarderCreateSlack) SetAuthMethodsEventSourceLocations(v []string) {
	o.AuthMethodsEventSourceLocations = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventForwarderCreateSlack) SetDescription(v string) {
	o.Description = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetEventTypesOk() (*[]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *EventForwarderCreateSlack) SetEventTypes(v []string) {
	o.EventTypes = &v
}

// GetEvery returns the Every field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetEvery() string {
	if o == nil || o.Every == nil {
		var ret string
		return ret
	}
	return *o.Every
}

// GetEveryOk returns a tuple with the Every field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetEveryOk() (*string, bool) {
	if o == nil || o.Every == nil {
		return nil, false
	}
	return o.Every, true
}

// HasEvery returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasEvery() bool {
	if o != nil && o.Every != nil {
		return true
	}

	return false
}

// SetEvery gets a reference to the given string and assigns it to the Every field.
func (o *EventForwarderCreateSlack) SetEvery(v string) {
	o.Every = &v
}

// GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field value
func (o *EventForwarderCreateSlack) GetGatewaysEventSourceLocations() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.GatewaysEventSourceLocations
}

// GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetGatewaysEventSourceLocationsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GatewaysEventSourceLocations, true
}

// SetGatewaysEventSourceLocations sets field value
func (o *EventForwarderCreateSlack) SetGatewaysEventSourceLocations(v []string) {
	o.GatewaysEventSourceLocations = v
}

// GetItemsEventSourceLocations returns the ItemsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetItemsEventSourceLocations() []string {
	if o == nil || o.ItemsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.ItemsEventSourceLocations
}

// GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetItemsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.ItemsEventSourceLocations == nil {
		return nil, false
	}
	return o.ItemsEventSourceLocations, true
}

// HasItemsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasItemsEventSourceLocations() bool {
	if o != nil && o.ItemsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetItemsEventSourceLocations gets a reference to the given []string and assigns it to the ItemsEventSourceLocations field.
func (o *EventForwarderCreateSlack) SetItemsEventSourceLocations(v []string) {
	o.ItemsEventSourceLocations = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *EventForwarderCreateSlack) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EventForwarderCreateSlack) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *EventForwarderCreateSlack) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventForwarderCreateSlack) SetName(v string) {
	o.Name = v
}

// GetRunnerType returns the RunnerType field value
func (o *EventForwarderCreateSlack) GetRunnerType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RunnerType
}

// GetRunnerTypeOk returns a tuple with the RunnerType field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetRunnerTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RunnerType, true
}

// SetRunnerType sets field value
func (o *EventForwarderCreateSlack) SetRunnerType(v string) {
	o.RunnerType = v
}

// GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetTargetsEventSourceLocations() []string {
	if o == nil || o.TargetsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.TargetsEventSourceLocations
}

// GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetTargetsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.TargetsEventSourceLocations == nil {
		return nil, false
	}
	return o.TargetsEventSourceLocations, true
}

// HasTargetsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasTargetsEventSourceLocations() bool {
	if o != nil && o.TargetsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetTargetsEventSourceLocations gets a reference to the given []string and assigns it to the TargetsEventSourceLocations field.
func (o *EventForwarderCreateSlack) SetTargetsEventSourceLocations(v []string) {
	o.TargetsEventSourceLocations = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EventForwarderCreateSlack) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EventForwarderCreateSlack) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EventForwarderCreateSlack) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EventForwarderCreateSlack) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUrl returns the Url field value
func (o *EventForwarderCreateSlack) GetUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *EventForwarderCreateSlack) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *EventForwarderCreateSlack) SetUrl(v string) {
	o.Url = v
}

func (o EventForwarderCreateSlack) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthMethodsEventSourceLocations != nil {
		toSerialize["auth-methods-event-source-locations"] = o.AuthMethodsEventSourceLocations
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableEventForwarderCreateSlack struct {
	value *EventForwarderCreateSlack
	isSet bool
}

func (v NullableEventForwarderCreateSlack) Get() *EventForwarderCreateSlack {
	return v.value
}

func (v *NullableEventForwarderCreateSlack) Set(val *EventForwarderCreateSlack) {
	v.value = val
	v.isSet = true
}

func (v NullableEventForwarderCreateSlack) IsSet() bool {
	return v.isSet
}

func (v *NullableEventForwarderCreateSlack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventForwarderCreateSlack(val *EventForwarderCreateSlack) *NullableEventForwarderCreateSlack {
	return &NullableEventForwarderCreateSlack{value: val, isSet: true}
}

func (v NullableEventForwarderCreateSlack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventForwarderCreateSlack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


