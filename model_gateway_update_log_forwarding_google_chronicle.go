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

// GatewayUpdateLogForwardingGoogleChronicle gatewayUpdateLogForwardingGoogleChronicle is a command that updates log forwarding config (google-chronicle target)
type GatewayUpdateLogForwardingGoogleChronicle struct {
	// Google chronicle customer id
	CustomerId *string `json:"customer-id,omitempty"`
	// Enable Log Forwarding [true/false]
	Enable *string `json:"enable,omitempty"`
	// Base64-encoded service account private key text
	GcpKey *string `json:"gcp-key,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Google chronicle log type
	LogType *string `json:"log-type,omitempty"`
	// Logs format [text/json]
	OutputFormat *string `json:"output-format,omitempty"`
	// Pull interval in seconds
	PullInterval *string `json:"pull-interval,omitempty"`
	// Google chronicle region [eu_multi_region/london/us_multi_region/singapore/tel_aviv]
	Region *string `json:"region,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayUpdateLogForwardingGoogleChronicle instantiates a new GatewayUpdateLogForwardingGoogleChronicle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateLogForwardingGoogleChronicle() *GatewayUpdateLogForwardingGoogleChronicle {
	this := GatewayUpdateLogForwardingGoogleChronicle{}
	var enable string = "true"
	this.Enable = &enable
	var json bool = false
	this.Json = &json
	var outputFormat string = "text"
	this.OutputFormat = &outputFormat
	var pullInterval string = "10"
	this.PullInterval = &pullInterval
	return &this
}

// NewGatewayUpdateLogForwardingGoogleChronicleWithDefaults instantiates a new GatewayUpdateLogForwardingGoogleChronicle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateLogForwardingGoogleChronicleWithDefaults() *GatewayUpdateLogForwardingGoogleChronicle {
	this := GatewayUpdateLogForwardingGoogleChronicle{}
	var enable string = "true"
	this.Enable = &enable
	var json bool = false
	this.Json = &json
	var outputFormat string = "text"
	this.OutputFormat = &outputFormat
	var pullInterval string = "10"
	this.PullInterval = &pullInterval
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetEnable() string {
	if o == nil || o.Enable == nil {
		var ret string
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetEnableOk() (*string, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given string and assigns it to the Enable field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetEnable(v string) {
	o.Enable = &v
}

// GetGcpKey returns the GcpKey field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetGcpKey() string {
	if o == nil || o.GcpKey == nil {
		var ret string
		return ret
	}
	return *o.GcpKey
}

// GetGcpKeyOk returns a tuple with the GcpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetGcpKeyOk() (*string, bool) {
	if o == nil || o.GcpKey == nil {
		return nil, false
	}
	return o.GcpKey, true
}

// HasGcpKey returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasGcpKey() bool {
	if o != nil && o.GcpKey != nil {
		return true
	}

	return false
}

// SetGcpKey gets a reference to the given string and assigns it to the GcpKey field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetGcpKey(v string) {
	o.GcpKey = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetJson(v bool) {
	o.Json = &v
}

// GetLogType returns the LogType field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetLogType() string {
	if o == nil || o.LogType == nil {
		var ret string
		return ret
	}
	return *o.LogType
}

// GetLogTypeOk returns a tuple with the LogType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetLogTypeOk() (*string, bool) {
	if o == nil || o.LogType == nil {
		return nil, false
	}
	return o.LogType, true
}

// HasLogType returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasLogType() bool {
	if o != nil && o.LogType != nil {
		return true
	}

	return false
}

// SetLogType gets a reference to the given string and assigns it to the LogType field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetLogType(v string) {
	o.LogType = &v
}

// GetOutputFormat returns the OutputFormat field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetOutputFormat() string {
	if o == nil || o.OutputFormat == nil {
		var ret string
		return ret
	}
	return *o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetOutputFormatOk() (*string, bool) {
	if o == nil || o.OutputFormat == nil {
		return nil, false
	}
	return o.OutputFormat, true
}

// HasOutputFormat returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasOutputFormat() bool {
	if o != nil && o.OutputFormat != nil {
		return true
	}

	return false
}

// SetOutputFormat gets a reference to the given string and assigns it to the OutputFormat field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetOutputFormat(v string) {
	o.OutputFormat = &v
}

// GetPullInterval returns the PullInterval field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetPullInterval() string {
	if o == nil || o.PullInterval == nil {
		var ret string
		return ret
	}
	return *o.PullInterval
}

// GetPullIntervalOk returns a tuple with the PullInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetPullIntervalOk() (*string, bool) {
	if o == nil || o.PullInterval == nil {
		return nil, false
	}
	return o.PullInterval, true
}

// HasPullInterval returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasPullInterval() bool {
	if o != nil && o.PullInterval != nil {
		return true
	}

	return false
}

// SetPullInterval gets a reference to the given string and assigns it to the PullInterval field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetPullInterval(v string) {
	o.PullInterval = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetRegion(v string) {
	o.Region = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingGoogleChronicle) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateLogForwardingGoogleChronicle) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayUpdateLogForwardingGoogleChronicle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerId != nil {
		toSerialize["customer-id"] = o.CustomerId
	}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.GcpKey != nil {
		toSerialize["gcp-key"] = o.GcpKey
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.LogType != nil {
		toSerialize["log-type"] = o.LogType
	}
	if o.OutputFormat != nil {
		toSerialize["output-format"] = o.OutputFormat
	}
	if o.PullInterval != nil {
		toSerialize["pull-interval"] = o.PullInterval
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateLogForwardingGoogleChronicle struct {
	value *GatewayUpdateLogForwardingGoogleChronicle
	isSet bool
}

func (v NullableGatewayUpdateLogForwardingGoogleChronicle) Get() *GatewayUpdateLogForwardingGoogleChronicle {
	return v.value
}

func (v *NullableGatewayUpdateLogForwardingGoogleChronicle) Set(val *GatewayUpdateLogForwardingGoogleChronicle) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateLogForwardingGoogleChronicle) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateLogForwardingGoogleChronicle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateLogForwardingGoogleChronicle(val *GatewayUpdateLogForwardingGoogleChronicle) *NullableGatewayUpdateLogForwardingGoogleChronicle {
	return &NullableGatewayUpdateLogForwardingGoogleChronicle{value: val, isSet: true}
}

func (v NullableGatewayUpdateLogForwardingGoogleChronicle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateLogForwardingGoogleChronicle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


