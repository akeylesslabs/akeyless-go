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

// GatewayCreateProducerCustom struct for GatewayCreateProducerCustom
type GatewayCreateProducerCustom struct {
	// URL of an endpoint that implements /sync/create method, for example https://webhook.example.com/sync/create
	CreateSyncUrl string `json:"create-sync-url"`
	// Gateway url
	GatewayUrl *string `json:"gateway-url,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Secret payload to be sent with each create/revoke webhook request
	Payload *string `json:"payload,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// URL of an endpoint that implements /sync/revoke method, for example https://webhook.example.com/sync/revoke
	RevokeSyncUrl string `json:"revoke-sync-url"`
	// URL of an endpoint that implements /sync/rotate method, for example https://webhook.example.com/sync/rotate
	RotateSyncUrl *string `json:"rotate-sync-url,omitempty"`
	// Maximum allowed time in seconds for the webhook to return the results
	TimeoutSec *int64 `json:"timeout-sec,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewGatewayCreateProducerCustom instantiates a new GatewayCreateProducerCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerCustom(createSyncUrl string, name string, revokeSyncUrl string, ) *GatewayCreateProducerCustom {
	this := GatewayCreateProducerCustom{}
	this.CreateSyncUrl = createSyncUrl
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
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
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	var timeoutSec int64 = 60
	this.TimeoutSec = &timeoutSec
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
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

// GetGatewayUrl returns the GatewayUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetGatewayUrl() string {
	if o == nil || o.GatewayUrl == nil {
		var ret string
		return ret
	}
	return *o.GatewayUrl
}

// GetGatewayUrlOk returns a tuple with the GatewayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetGatewayUrlOk() (*string, bool) {
	if o == nil || o.GatewayUrl == nil {
		return nil, false
	}
	return o.GatewayUrl, true
}

// HasGatewayUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasGatewayUrl() bool {
	if o != nil && o.GatewayUrl != nil {
		return true
	}

	return false
}

// SetGatewayUrl gets a reference to the given string and assigns it to the GatewayUrl field.
func (o *GatewayCreateProducerCustom) SetGatewayUrl(v string) {
	o.GatewayUrl = &v
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

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayCreateProducerCustom) SetPassword(v string) {
	o.Password = &v
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

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayCreateProducerCustom) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCustom) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerCustom) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayCreateProducerCustom) SetUsername(v string) {
	o.Username = &v
}

func (o GatewayCreateProducerCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["create-sync-url"] = o.CreateSyncUrl
	}
	if o.GatewayUrl != nil {
		toSerialize["gateway-url"] = o.GatewayUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
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


