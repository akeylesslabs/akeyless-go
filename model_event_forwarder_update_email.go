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

// EventForwarderUpdateEmail eventForwarderUpdateEmail is a command that updates email event forwarder
type EventForwarderUpdateEmail struct {
	// Auth Method Event sources
	AuthMethodsEventSourceLocations *[]string `json:"auth-methods-event-source-locations,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// A comma seperated list of email addresses to send event to
	EmailTo *string `json:"email-to,omitempty"`
	// Enable/Disable Event Forwarder [true/false]
	Enable *string `json:"enable,omitempty"`
	// List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated]
	EventTypes *[]string `json:"event-types,omitempty"`
	// Event sources
	GatewaysEventSourceLocations []string `json:"gateways-event-source-locations"`
	// Items Event sources
	ItemsEventSourceLocations *[]string `json:"items-event-source-locations,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Whether to keep previous version [true/false]. If not set, use default according to account settings
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// EventForwarder name
	Name string `json:"name"`
	// New EventForwarder name
	NewName *string `json:"new-name,omitempty"`
	// Override Akeyless default URL with your Gateway url (port 18888)
	OverrideUrl *string `json:"override-url,omitempty"`
	// Targets Event sources
	TargetsEventSourceLocations *[]string `json:"targets-event-source-locations,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewEventForwarderUpdateEmail instantiates a new EventForwarderUpdateEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventForwarderUpdateEmail(gatewaysEventSourceLocations []string, name string, ) *EventForwarderUpdateEmail {
	this := EventForwarderUpdateEmail{}
	var enable string = "true"
	this.Enable = &enable
	this.GatewaysEventSourceLocations = gatewaysEventSourceLocations
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewEventForwarderUpdateEmailWithDefaults instantiates a new EventForwarderUpdateEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventForwarderUpdateEmailWithDefaults() *EventForwarderUpdateEmail {
	this := EventForwarderUpdateEmail{}
	var enable string = "true"
	this.Enable = &enable
	var json bool = false
	this.Json = &json
	return &this
}

// GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetAuthMethodsEventSourceLocations() []string {
	if o == nil || o.AuthMethodsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.AuthMethodsEventSourceLocations
}

// GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.AuthMethodsEventSourceLocations == nil {
		return nil, false
	}
	return o.AuthMethodsEventSourceLocations, true
}

// HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasAuthMethodsEventSourceLocations() bool {
	if o != nil && o.AuthMethodsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetAuthMethodsEventSourceLocations gets a reference to the given []string and assigns it to the AuthMethodsEventSourceLocations field.
func (o *EventForwarderUpdateEmail) SetAuthMethodsEventSourceLocations(v []string) {
	o.AuthMethodsEventSourceLocations = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventForwarderUpdateEmail) SetDescription(v string) {
	o.Description = &v
}

// GetEmailTo returns the EmailTo field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetEmailTo() string {
	if o == nil || o.EmailTo == nil {
		var ret string
		return ret
	}
	return *o.EmailTo
}

// GetEmailToOk returns a tuple with the EmailTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetEmailToOk() (*string, bool) {
	if o == nil || o.EmailTo == nil {
		return nil, false
	}
	return o.EmailTo, true
}

// HasEmailTo returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasEmailTo() bool {
	if o != nil && o.EmailTo != nil {
		return true
	}

	return false
}

// SetEmailTo gets a reference to the given string and assigns it to the EmailTo field.
func (o *EventForwarderUpdateEmail) SetEmailTo(v string) {
	o.EmailTo = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetEnable() string {
	if o == nil || o.Enable == nil {
		var ret string
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetEnableOk() (*string, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given string and assigns it to the Enable field.
func (o *EventForwarderUpdateEmail) SetEnable(v string) {
	o.Enable = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetEventTypesOk() (*[]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *EventForwarderUpdateEmail) SetEventTypes(v []string) {
	o.EventTypes = &v
}

// GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field value
func (o *EventForwarderUpdateEmail) GetGatewaysEventSourceLocations() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.GatewaysEventSourceLocations
}

// GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field value
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetGatewaysEventSourceLocationsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GatewaysEventSourceLocations, true
}

// SetGatewaysEventSourceLocations sets field value
func (o *EventForwarderUpdateEmail) SetGatewaysEventSourceLocations(v []string) {
	o.GatewaysEventSourceLocations = v
}

// GetItemsEventSourceLocations returns the ItemsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetItemsEventSourceLocations() []string {
	if o == nil || o.ItemsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.ItemsEventSourceLocations
}

// GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetItemsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.ItemsEventSourceLocations == nil {
		return nil, false
	}
	return o.ItemsEventSourceLocations, true
}

// HasItemsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasItemsEventSourceLocations() bool {
	if o != nil && o.ItemsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetItemsEventSourceLocations gets a reference to the given []string and assigns it to the ItemsEventSourceLocations field.
func (o *EventForwarderUpdateEmail) SetItemsEventSourceLocations(v []string) {
	o.ItemsEventSourceLocations = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *EventForwarderUpdateEmail) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *EventForwarderUpdateEmail) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EventForwarderUpdateEmail) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *EventForwarderUpdateEmail) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventForwarderUpdateEmail) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *EventForwarderUpdateEmail) SetNewName(v string) {
	o.NewName = &v
}

// GetOverrideUrl returns the OverrideUrl field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetOverrideUrl() string {
	if o == nil || o.OverrideUrl == nil {
		var ret string
		return ret
	}
	return *o.OverrideUrl
}

// GetOverrideUrlOk returns a tuple with the OverrideUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetOverrideUrlOk() (*string, bool) {
	if o == nil || o.OverrideUrl == nil {
		return nil, false
	}
	return o.OverrideUrl, true
}

// HasOverrideUrl returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasOverrideUrl() bool {
	if o != nil && o.OverrideUrl != nil {
		return true
	}

	return false
}

// SetOverrideUrl gets a reference to the given string and assigns it to the OverrideUrl field.
func (o *EventForwarderUpdateEmail) SetOverrideUrl(v string) {
	o.OverrideUrl = &v
}

// GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetTargetsEventSourceLocations() []string {
	if o == nil || o.TargetsEventSourceLocations == nil {
		var ret []string
		return ret
	}
	return *o.TargetsEventSourceLocations
}

// GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetTargetsEventSourceLocationsOk() (*[]string, bool) {
	if o == nil || o.TargetsEventSourceLocations == nil {
		return nil, false
	}
	return o.TargetsEventSourceLocations, true
}

// HasTargetsEventSourceLocations returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasTargetsEventSourceLocations() bool {
	if o != nil && o.TargetsEventSourceLocations != nil {
		return true
	}

	return false
}

// SetTargetsEventSourceLocations gets a reference to the given []string and assigns it to the TargetsEventSourceLocations field.
func (o *EventForwarderUpdateEmail) SetTargetsEventSourceLocations(v []string) {
	o.TargetsEventSourceLocations = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EventForwarderUpdateEmail) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EventForwarderUpdateEmail) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderUpdateEmail) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EventForwarderUpdateEmail) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EventForwarderUpdateEmail) SetUidToken(v string) {
	o.UidToken = &v
}

func (o EventForwarderUpdateEmail) MarshalJSON() ([]byte, error) {
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
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.EventTypes != nil {
		toSerialize["event-types"] = o.EventTypes
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
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.OverrideUrl != nil {
		toSerialize["override-url"] = o.OverrideUrl
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

type NullableEventForwarderUpdateEmail struct {
	value *EventForwarderUpdateEmail
	isSet bool
}

func (v NullableEventForwarderUpdateEmail) Get() *EventForwarderUpdateEmail {
	return v.value
}

func (v *NullableEventForwarderUpdateEmail) Set(val *EventForwarderUpdateEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableEventForwarderUpdateEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableEventForwarderUpdateEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventForwarderUpdateEmail(val *EventForwarderUpdateEmail) *NullableEventForwarderUpdateEmail {
	return &NullableEventForwarderUpdateEmail{value: val, isSet: true}
}

func (v NullableEventForwarderUpdateEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventForwarderUpdateEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


