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

// GatewayCreateProducerCustom gatewayCreateProducerCustom is a command that creates a custom producer. [Deprecated: Use dynamic-secret-create-custom command]
type GatewayCreateProducerCustom struct {
	// Define rotation interval in days
	AdminRotationIntervalDays *int64 `json:"admin_rotation_interval_days,omitempty"`
	// URL of an endpoint that implements /sync/create method, for example https://webhook.example.com/sync/create
	CreateSyncUrl string `json:"create-sync-url"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Should admin credentials be rotated
	EnableAdminRotation *bool `json:"enable_admin_rotation,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Secret payload to be sent with each create/revoke webhook request
	Payload *string `json:"payload,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// URL of an endpoint that implements /sync/revoke method, for example https://webhook.example.com/sync/revoke
	RevokeSyncUrl string `json:"revoke-sync-url"`
	// URL of an endpoint that implements /sync/rotate method, for example https://webhook.example.com/sync/rotate
	RotateSyncUrl *string `json:"rotate-sync-url,omitempty"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Maximum allowed time in seconds for the webhook to return the results
	TimeoutSec *int64 `json:"timeout-sec,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayCreateProducerCustom instantiates a new GatewayCreateProducerCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerCustom(createSyncUrl string, name string, revokeSyncUrl string, ) *GatewayCreateProducerCustom {
	this := GatewayCreateProducerCustom{}
	this.CreateSyncUrl = createSyncUrl
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	var json bool = false
	this.Json = &json
	this.Name = name
	this.RevokeSyncUrl = revokeSyncUrl
	var timeoutSec int64 = 60
	this.TimeoutSec = &timeoutSec
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerCustomWithDefaults instantiates a new GatewayCreateProducerCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerCustomWithDefaults() *GatewayCreateProducerCustom {
	this := GatewayCreateProducerCustom{}
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	var json bool = false
	this.Json = &json
	var timeoutSec int64 = 60
	this.TimeoutSec = &timeoutSec
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetAdminRotationIntervalDays() int64 {
	if o == nil || o.AdminRotationIntervalDays == nil {
		var ret int64
		return ret
	}
	return *o.AdminRotationIntervalDays
}

// GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetAdminRotationIntervalDaysOk() (*int64, bool) {
	if o == nil || o.AdminRotationIntervalDays == nil {
		return nil, false
	}
	return o.AdminRotationIntervalDays, true
}

// HasAdminRotationIntervalDays returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasAdminRotationIntervalDays() bool {
	if o != nil && o.AdminRotationIntervalDays != nil {
		return true
	}

	return false
}

// SetAdminRotationIntervalDays gets a reference to the given int64 and assigns it to the AdminRotationIntervalDays field.
func (o *GatewayCreateProducerCustom) SetAdminRotationIntervalDays(v int64) {
	o.AdminRotationIntervalDays = &v
}

// GetCreateSyncUrl returns the CreateSyncUrl field value
func (o *GatewayCreateProducerCustom) GetCreateSyncUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CreateSyncUrl
}

// GetCreateSyncUrlOk returns a tuple with the CreateSyncUrl field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetCreateSyncUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreateSyncUrl, true
}

// SetCreateSyncUrl sets field value
func (o *GatewayCreateProducerCustom) SetCreateSyncUrl(v string) {
	o.CreateSyncUrl = v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerCustom) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetEnableAdminRotation returns the EnableAdminRotation field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetEnableAdminRotation() bool {
	if o == nil || o.EnableAdminRotation == nil {
		var ret bool
		return ret
	}
	return *o.EnableAdminRotation
}

// GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetEnableAdminRotationOk() (*bool, bool) {
	if o == nil || o.EnableAdminRotation == nil {
		return nil, false
	}
	return o.EnableAdminRotation, true
}

// HasEnableAdminRotation returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasEnableAdminRotation() bool {
	if o != nil && o.EnableAdminRotation != nil {
		return true
	}

	return false
}

// SetEnableAdminRotation gets a reference to the given bool and assigns it to the EnableAdminRotation field.
func (o *GatewayCreateProducerCustom) SetEnableAdminRotation(v bool) {
	o.EnableAdminRotation = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerCustom) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerCustom) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerCustom) SetName(v string) {
	o.Name = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *GatewayCreateProducerCustom) SetPayload(v string) {
	o.Payload = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerCustom) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRevokeSyncUrl returns the RevokeSyncUrl field value
func (o *GatewayCreateProducerCustom) GetRevokeSyncUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RevokeSyncUrl
}

// GetRevokeSyncUrlOk returns a tuple with the RevokeSyncUrl field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetRevokeSyncUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RevokeSyncUrl, true
}

// SetRevokeSyncUrl sets field value
func (o *GatewayCreateProducerCustom) SetRevokeSyncUrl(v string) {
	o.RevokeSyncUrl = v
}

// GetRotateSyncUrl returns the RotateSyncUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetRotateSyncUrl() string {
	if o == nil || o.RotateSyncUrl == nil {
		var ret string
		return ret
	}
	return *o.RotateSyncUrl
}

// GetRotateSyncUrlOk returns a tuple with the RotateSyncUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetRotateSyncUrlOk() (*string, bool) {
	if o == nil || o.RotateSyncUrl == nil {
		return nil, false
	}
	return o.RotateSyncUrl, true
}

// HasRotateSyncUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasRotateSyncUrl() bool {
	if o != nil && o.RotateSyncUrl != nil {
		return true
	}

	return false
}

// SetRotateSyncUrl gets a reference to the given string and assigns it to the RotateSyncUrl field.
func (o *GatewayCreateProducerCustom) SetRotateSyncUrl(v string) {
	o.RotateSyncUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerCustom) SetTags(v []string) {
	o.Tags = &v
}

// GetTimeoutSec returns the TimeoutSec field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetTimeoutSec() int64 {
	if o == nil || o.TimeoutSec == nil {
		var ret int64
		return ret
	}
	return *o.TimeoutSec
}

// GetTimeoutSecOk returns a tuple with the TimeoutSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetTimeoutSecOk() (*int64, bool) {
	if o == nil || o.TimeoutSec == nil {
		return nil, false
	}
	return o.TimeoutSec, true
}

// HasTimeoutSec returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasTimeoutSec() bool {
	if o != nil && o.TimeoutSec != nil {
		return true
	}

	return false
}

// SetTimeoutSec gets a reference to the given int64 and assigns it to the TimeoutSec field.
func (o *GatewayCreateProducerCustom) SetTimeoutSec(v int64) {
	o.TimeoutSec = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerCustom) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerCustom) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerCustom) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminRotationIntervalDays != nil {
		toSerialize["admin_rotation_interval_days"] = o.AdminRotationIntervalDays
	}
	if true {
		toSerialize["create-sync-url"] = o.CreateSyncUrl
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.EnableAdminRotation != nil {
		toSerialize["enable_admin_rotation"] = o.EnableAdminRotation
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if true {
		toSerialize["revoke-sync-url"] = o.RevokeSyncUrl
	}
	if o.RotateSyncUrl != nil {
		toSerialize["rotate-sync-url"] = o.RotateSyncUrl
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TimeoutSec != nil {
		toSerialize["timeout-sec"] = o.TimeoutSec
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerCustom struct {
	value *GatewayCreateProducerCustom
	isSet bool
}

func (v NullableGatewayCreateProducerCustom) Get() *GatewayCreateProducerCustom {
	return v.value
}

func (v *NullableGatewayCreateProducerCustom) Set(val *GatewayCreateProducerCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerCustom(val *GatewayCreateProducerCustom) *NullableGatewayCreateProducerCustom {
	return &NullableGatewayCreateProducerCustom{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


