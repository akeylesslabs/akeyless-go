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

// AkeylessGatewayConfig struct for AkeylessGatewayConfig
type AkeylessGatewayConfig struct {
	Admins *AdminsConfigPart `json:"admins,omitempty"`
	Cache *CacheConfigPart `json:"cache,omitempty"`
	Cf *CFConfigPart `json:"cf,omitempty"`
	ConfigProtectionKeyName *string `json:"config_protection_key_name,omitempty"`
	General *GeneralConfigPart `json:"general,omitempty"`
	Ldap *LdapConfigPart `json:"ldap,omitempty"`
	Leadership *LeadershipConfigPart `json:"leadership,omitempty"`
	LogForwarding *LogForwardingConfigPart `json:"log_forwarding,omitempty"`
	Migrations *MigrationsConfigPart `json:"migrations,omitempty"`
	Producers *ProducersConfigPart `json:"producers,omitempty"`
	Rotators *RotatorsConfigPart `json:"rotators,omitempty"`
	Saml *DefaultConfigPart `json:"saml,omitempty"`
	Uidentity *UIdentityConfigPart `json:"uidentity,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewAkeylessGatewayConfig instantiates a new AkeylessGatewayConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAkeylessGatewayConfig() *AkeylessGatewayConfig {
	this := AkeylessGatewayConfig{}
	return &this
}

// NewAkeylessGatewayConfigWithDefaults instantiates a new AkeylessGatewayConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAkeylessGatewayConfigWithDefaults() *AkeylessGatewayConfig {
	this := AkeylessGatewayConfig{}
	return &this
}

// GetAdmins returns the Admins field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetAdmins() AdminsConfigPart {
	if o == nil || o.Admins == nil {
		var ret AdminsConfigPart
		return ret
	}
	return *o.Admins
}

// GetAdminsOk returns a tuple with the Admins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetAdminsOk() (*AdminsConfigPart, bool) {
	if o == nil || o.Admins == nil {
		return nil, false
	}
	return o.Admins, true
}

// HasAdmins returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasAdmins() bool {
	if o != nil && o.Admins != nil {
		return true
	}

	return false
}

// SetAdmins gets a reference to the given AdminsConfigPart and assigns it to the Admins field.
func (o *AkeylessGatewayConfig) SetAdmins(v AdminsConfigPart) {
	o.Admins = &v
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetCache() CacheConfigPart {
	if o == nil || o.Cache == nil {
		var ret CacheConfigPart
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetCacheOk() (*CacheConfigPart, bool) {
	if o == nil || o.Cache == nil {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasCache() bool {
	if o != nil && o.Cache != nil {
		return true
	}

	return false
}

// SetCache gets a reference to the given CacheConfigPart and assigns it to the Cache field.
func (o *AkeylessGatewayConfig) SetCache(v CacheConfigPart) {
	o.Cache = &v
}

// GetCf returns the Cf field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetCf() CFConfigPart {
	if o == nil || o.Cf == nil {
		var ret CFConfigPart
		return ret
	}
	return *o.Cf
}

// GetCfOk returns a tuple with the Cf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetCfOk() (*CFConfigPart, bool) {
	if o == nil || o.Cf == nil {
		return nil, false
	}
	return o.Cf, true
}

// HasCf returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasCf() bool {
	if o != nil && o.Cf != nil {
		return true
	}

	return false
}

// SetCf gets a reference to the given CFConfigPart and assigns it to the Cf field.
func (o *AkeylessGatewayConfig) SetCf(v CFConfigPart) {
	o.Cf = &v
}

// GetConfigProtectionKeyName returns the ConfigProtectionKeyName field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetConfigProtectionKeyName() string {
	if o == nil || o.ConfigProtectionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ConfigProtectionKeyName
}

// GetConfigProtectionKeyNameOk returns a tuple with the ConfigProtectionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetConfigProtectionKeyNameOk() (*string, bool) {
	if o == nil || o.ConfigProtectionKeyName == nil {
		return nil, false
	}
	return o.ConfigProtectionKeyName, true
}

// HasConfigProtectionKeyName returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasConfigProtectionKeyName() bool {
	if o != nil && o.ConfigProtectionKeyName != nil {
		return true
	}

	return false
}

// SetConfigProtectionKeyName gets a reference to the given string and assigns it to the ConfigProtectionKeyName field.
func (o *AkeylessGatewayConfig) SetConfigProtectionKeyName(v string) {
	o.ConfigProtectionKeyName = &v
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetGeneral() GeneralConfigPart {
	if o == nil || o.General == nil {
		var ret GeneralConfigPart
		return ret
	}
	return *o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetGeneralOk() (*GeneralConfigPart, bool) {
	if o == nil || o.General == nil {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasGeneral() bool {
	if o != nil && o.General != nil {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given GeneralConfigPart and assigns it to the General field.
func (o *AkeylessGatewayConfig) SetGeneral(v GeneralConfigPart) {
	o.General = &v
}

// GetLdap returns the Ldap field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetLdap() LdapConfigPart {
	if o == nil || o.Ldap == nil {
		var ret LdapConfigPart
		return ret
	}
	return *o.Ldap
}

// GetLdapOk returns a tuple with the Ldap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetLdapOk() (*LdapConfigPart, bool) {
	if o == nil || o.Ldap == nil {
		return nil, false
	}
	return o.Ldap, true
}

// HasLdap returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasLdap() bool {
	if o != nil && o.Ldap != nil {
		return true
	}

	return false
}

// SetLdap gets a reference to the given LdapConfigPart and assigns it to the Ldap field.
func (o *AkeylessGatewayConfig) SetLdap(v LdapConfigPart) {
	o.Ldap = &v
}

// GetLeadership returns the Leadership field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetLeadership() LeadershipConfigPart {
	if o == nil || o.Leadership == nil {
		var ret LeadershipConfigPart
		return ret
	}
	return *o.Leadership
}

// GetLeadershipOk returns a tuple with the Leadership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetLeadershipOk() (*LeadershipConfigPart, bool) {
	if o == nil || o.Leadership == nil {
		return nil, false
	}
	return o.Leadership, true
}

// HasLeadership returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasLeadership() bool {
	if o != nil && o.Leadership != nil {
		return true
	}

	return false
}

// SetLeadership gets a reference to the given LeadershipConfigPart and assigns it to the Leadership field.
func (o *AkeylessGatewayConfig) SetLeadership(v LeadershipConfigPart) {
	o.Leadership = &v
}

// GetLogForwarding returns the LogForwarding field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetLogForwarding() LogForwardingConfigPart {
	if o == nil || o.LogForwarding == nil {
		var ret LogForwardingConfigPart
		return ret
	}
	return *o.LogForwarding
}

// GetLogForwardingOk returns a tuple with the LogForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetLogForwardingOk() (*LogForwardingConfigPart, bool) {
	if o == nil || o.LogForwarding == nil {
		return nil, false
	}
	return o.LogForwarding, true
}

// HasLogForwarding returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasLogForwarding() bool {
	if o != nil && o.LogForwarding != nil {
		return true
	}

	return false
}

// SetLogForwarding gets a reference to the given LogForwardingConfigPart and assigns it to the LogForwarding field.
func (o *AkeylessGatewayConfig) SetLogForwarding(v LogForwardingConfigPart) {
	o.LogForwarding = &v
}

// GetMigrations returns the Migrations field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetMigrations() MigrationsConfigPart {
	if o == nil || o.Migrations == nil {
		var ret MigrationsConfigPart
		return ret
	}
	return *o.Migrations
}

// GetMigrationsOk returns a tuple with the Migrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetMigrationsOk() (*MigrationsConfigPart, bool) {
	if o == nil || o.Migrations == nil {
		return nil, false
	}
	return o.Migrations, true
}

// HasMigrations returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasMigrations() bool {
	if o != nil && o.Migrations != nil {
		return true
	}

	return false
}

// SetMigrations gets a reference to the given MigrationsConfigPart and assigns it to the Migrations field.
func (o *AkeylessGatewayConfig) SetMigrations(v MigrationsConfigPart) {
	o.Migrations = &v
}

// GetProducers returns the Producers field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetProducers() ProducersConfigPart {
	if o == nil || o.Producers == nil {
		var ret ProducersConfigPart
		return ret
	}
	return *o.Producers
}

// GetProducersOk returns a tuple with the Producers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetProducersOk() (*ProducersConfigPart, bool) {
	if o == nil || o.Producers == nil {
		return nil, false
	}
	return o.Producers, true
}

// HasProducers returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasProducers() bool {
	if o != nil && o.Producers != nil {
		return true
	}

	return false
}

// SetProducers gets a reference to the given ProducersConfigPart and assigns it to the Producers field.
func (o *AkeylessGatewayConfig) SetProducers(v ProducersConfigPart) {
	o.Producers = &v
}

// GetRotators returns the Rotators field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetRotators() RotatorsConfigPart {
	if o == nil || o.Rotators == nil {
		var ret RotatorsConfigPart
		return ret
	}
	return *o.Rotators
}

// GetRotatorsOk returns a tuple with the Rotators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetRotatorsOk() (*RotatorsConfigPart, bool) {
	if o == nil || o.Rotators == nil {
		return nil, false
	}
	return o.Rotators, true
}

// HasRotators returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasRotators() bool {
	if o != nil && o.Rotators != nil {
		return true
	}

	return false
}

// SetRotators gets a reference to the given RotatorsConfigPart and assigns it to the Rotators field.
func (o *AkeylessGatewayConfig) SetRotators(v RotatorsConfigPart) {
	o.Rotators = &v
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetSaml() DefaultConfigPart {
	if o == nil || o.Saml == nil {
		var ret DefaultConfigPart
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetSamlOk() (*DefaultConfigPart, bool) {
	if o == nil || o.Saml == nil {
		return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasSaml() bool {
	if o != nil && o.Saml != nil {
		return true
	}

	return false
}

// SetSaml gets a reference to the given DefaultConfigPart and assigns it to the Saml field.
func (o *AkeylessGatewayConfig) SetSaml(v DefaultConfigPart) {
	o.Saml = &v
}

// GetUidentity returns the Uidentity field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetUidentity() UIdentityConfigPart {
	if o == nil || o.Uidentity == nil {
		var ret UIdentityConfigPart
		return ret
	}
	return *o.Uidentity
}

// GetUidentityOk returns a tuple with the Uidentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetUidentityOk() (*UIdentityConfigPart, bool) {
	if o == nil || o.Uidentity == nil {
		return nil, false
	}
	return o.Uidentity, true
}

// HasUidentity returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasUidentity() bool {
	if o != nil && o.Uidentity != nil {
		return true
	}

	return false
}

// SetUidentity gets a reference to the given UIdentityConfigPart and assigns it to the Uidentity field.
func (o *AkeylessGatewayConfig) SetUidentity(v UIdentityConfigPart) {
	o.Uidentity = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AkeylessGatewayConfig) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkeylessGatewayConfig) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AkeylessGatewayConfig) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *AkeylessGatewayConfig) SetVersion(v int32) {
	o.Version = &v
}

func (o AkeylessGatewayConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admins != nil {
		toSerialize["admins"] = o.Admins
	}
	if o.Cache != nil {
		toSerialize["cache"] = o.Cache
	}
	if o.Cf != nil {
		toSerialize["cf"] = o.Cf
	}
	if o.ConfigProtectionKeyName != nil {
		toSerialize["config_protection_key_name"] = o.ConfigProtectionKeyName
	}
	if o.General != nil {
		toSerialize["general"] = o.General
	}
	if o.Ldap != nil {
		toSerialize["ldap"] = o.Ldap
	}
	if o.Leadership != nil {
		toSerialize["leadership"] = o.Leadership
	}
	if o.LogForwarding != nil {
		toSerialize["log_forwarding"] = o.LogForwarding
	}
	if o.Migrations != nil {
		toSerialize["migrations"] = o.Migrations
	}
	if o.Producers != nil {
		toSerialize["producers"] = o.Producers
	}
	if o.Rotators != nil {
		toSerialize["rotators"] = o.Rotators
	}
	if o.Saml != nil {
		toSerialize["saml"] = o.Saml
	}
	if o.Uidentity != nil {
		toSerialize["uidentity"] = o.Uidentity
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableAkeylessGatewayConfig struct {
	value *AkeylessGatewayConfig
	isSet bool
}

func (v NullableAkeylessGatewayConfig) Get() *AkeylessGatewayConfig {
	return v.value
}

func (v *NullableAkeylessGatewayConfig) Set(val *AkeylessGatewayConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAkeylessGatewayConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAkeylessGatewayConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAkeylessGatewayConfig(val *AkeylessGatewayConfig) *NullableAkeylessGatewayConfig {
	return &NullableAkeylessGatewayConfig{value: val, isSet: true}
}

func (v NullableAkeylessGatewayConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAkeylessGatewayConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


