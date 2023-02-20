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

// GatewayUpdateProducerCustom struct for GatewayUpdateProducerCustom
type GatewayUpdateProducerCustom struct {
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
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// Secret payload to be sent with each create/revoke webhook request
	Payload *string `json:"payload,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// URL of an endpoint that implements /sync/revoke method, for example https://webhook.example.com/sync/revoke
	RevokeSyncUrl string `json:"revoke-sync-url"`
	// URL of an endpoint that implements /sync/rotate method, for example https://webhook.example.com/sync/rotate
	RotateSyncUrl *string `json:"rotate-sync-url,omitempty"`
	// List of the tags attached to this secret
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

// NewGatewayUpdateProducerCustom instantiates a new GatewayUpdateProducerCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerCustom(createSyncUrl string, name string, revokeSyncUrl string, ) *GatewayUpdateProducerCustom {
	this := GatewayUpdateProducerCustom{}
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

// NewGatewayUpdateProducerCustomWithDefaults instantiates a new GatewayUpdateProducerCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerCustomWithDefaults() *GatewayUpdateProducerCustom {
	this := GatewayUpdateProducerCustom{}
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
func (o *GatewayUpdateProducerCustom) GetAdminRotationIntervalDays() int64 {
	if o == nil || o.AdminRotationIntervalDays == nil {
		var ret int64
		return ret
	}
	return *o.AdminRotationIntervalDays
}

// GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetAdminRotationIntervalDaysOk() (*int64, bool) {
	if o == nil || o.AdminRotationIntervalDays == nil {
		return nil, false
	}
	return o.AdminRotationIntervalDays, true
}

// HasAdminRotationIntervalDays returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasAdminRotationIntervalDays() bool {
	if o != nil && o.AdminRotationIntervalDays != nil {
		return true
	}

	return false
}

// SetAdminRotationIntervalDays gets a reference to the given int64 and assigns it to the AdminRotationIntervalDays field.
func (o *GatewayUpdateProducerCustom) SetAdminRotationIntervalDays(v int64) {
	o.AdminRotationIntervalDays = &v
}

// GetCreateSyncUrl returns the CreateSyncUrl field value
func (o *GatewayUpdateProducerCustom) GetCreateSyncUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CreateSyncUrl
}

// GetCreateSyncUrlOk returns a tuple with the CreateSyncUrl field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetCreateSyncUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreateSyncUrl, true
}

// SetCreateSyncUrl sets field value
func (o *GatewayUpdateProducerCustom) SetCreateSyncUrl(v string) {
	o.CreateSyncUrl = v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerCustom) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetEnableAdminRotation returns the EnableAdminRotation field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetEnableAdminRotation() bool {
	if o == nil || o.EnableAdminRotation == nil {
		var ret bool
		return ret
	}
	return *o.EnableAdminRotation
}

// GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetEnableAdminRotationOk() (*bool, bool) {
	if o == nil || o.EnableAdminRotation == nil {
		return nil, false
	}
	return o.EnableAdminRotation, true
}

// HasEnableAdminRotation returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasEnableAdminRotation() bool {
	if o != nil && o.EnableAdminRotation != nil {
		return true
	}

	return false
}

// SetEnableAdminRotation gets a reference to the given bool and assigns it to the EnableAdminRotation field.
func (o *GatewayUpdateProducerCustom) SetEnableAdminRotation(v bool) {
	o.EnableAdminRotation = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerCustom) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerCustom) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerCustom) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerCustom) SetNewName(v string) {
	o.NewName = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *GatewayUpdateProducerCustom) SetPayload(v string) {
	o.Payload = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerCustom) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRevokeSyncUrl returns the RevokeSyncUrl field value
func (o *GatewayUpdateProducerCustom) GetRevokeSyncUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RevokeSyncUrl
}

// GetRevokeSyncUrlOk returns a tuple with the RevokeSyncUrl field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetRevokeSyncUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RevokeSyncUrl, true
}

// SetRevokeSyncUrl sets field value
func (o *GatewayUpdateProducerCustom) SetRevokeSyncUrl(v string) {
	o.RevokeSyncUrl = v
}

// GetRotateSyncUrl returns the RotateSyncUrl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetRotateSyncUrl() string {
	if o == nil || o.RotateSyncUrl == nil {
		var ret string
		return ret
	}
	return *o.RotateSyncUrl
}

// GetRotateSyncUrlOk returns a tuple with the RotateSyncUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetRotateSyncUrlOk() (*string, bool) {
	if o == nil || o.RotateSyncUrl == nil {
		return nil, false
	}
	return o.RotateSyncUrl, true
}

// HasRotateSyncUrl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasRotateSyncUrl() bool {
	if o != nil && o.RotateSyncUrl != nil {
		return true
	}

	return false
}

// SetRotateSyncUrl gets a reference to the given string and assigns it to the RotateSyncUrl field.
func (o *GatewayUpdateProducerCustom) SetRotateSyncUrl(v string) {
	o.RotateSyncUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerCustom) SetTags(v []string) {
	o.Tags = &v
}

// GetTimeoutSec returns the TimeoutSec field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetTimeoutSec() int64 {
	if o == nil || o.TimeoutSec == nil {
		var ret int64
		return ret
	}
	return *o.TimeoutSec
}

// GetTimeoutSecOk returns a tuple with the TimeoutSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetTimeoutSecOk() (*int64, bool) {
	if o == nil || o.TimeoutSec == nil {
		return nil, false
	}
	return o.TimeoutSec, true
}

// HasTimeoutSec returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasTimeoutSec() bool {
	if o != nil && o.TimeoutSec != nil {
		return true
	}

	return false
}

// SetTimeoutSec gets a reference to the given int64 and assigns it to the TimeoutSec field.
func (o *GatewayUpdateProducerCustom) SetTimeoutSec(v int64) {
	o.TimeoutSec = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerCustom) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerCustom) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCustom) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCustom) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCustom) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerCustom) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerCustom) MarshalJSON() ([]byte, error) {
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
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
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

type NullableGatewayUpdateProducerCustom struct {
	value *GatewayUpdateProducerCustom
	isSet bool
}

func (v NullableGatewayUpdateProducerCustom) Get() *GatewayUpdateProducerCustom {
	return v.value
}

func (v *NullableGatewayUpdateProducerCustom) Set(val *GatewayUpdateProducerCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerCustom(val *GatewayUpdateProducerCustom) *NullableGatewayUpdateProducerCustom {
	return &NullableGatewayUpdateProducerCustom{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


