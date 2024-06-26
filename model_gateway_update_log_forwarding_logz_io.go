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

// GatewayUpdateLogForwardingLogzIo gatewayUpdateLogForwardingLogzIo is a command that updates log forwarding config (logz-io target)
type GatewayUpdateLogForwardingLogzIo struct {
	// Enable Log Forwarding [true/false]
	Enable *string `json:"enable,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Logz-io token
	LogzIoToken *string `json:"logz-io-token,omitempty"`
	// Logs format [text/json]
	OutputFormat *string `json:"output-format,omitempty"`
	// LogzIo protocol [tcp/https]
	Protocol *string `json:"protocol,omitempty"`
	// Pull interval in seconds
	PullInterval *string `json:"pull-interval,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayUpdateLogForwardingLogzIo instantiates a new GatewayUpdateLogForwardingLogzIo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateLogForwardingLogzIo() *GatewayUpdateLogForwardingLogzIo {
	this := GatewayUpdateLogForwardingLogzIo{}
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

// NewGatewayUpdateLogForwardingLogzIoWithDefaults instantiates a new GatewayUpdateLogForwardingLogzIo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateLogForwardingLogzIoWithDefaults() *GatewayUpdateLogForwardingLogzIo {
	this := GatewayUpdateLogForwardingLogzIo{}
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

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogzIo) GetEnable() string {
	if o == nil || o.Enable == nil {
		var ret string
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogzIo) GetEnableOk() (*string, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogzIo) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given string and assigns it to the Enable field.
func (o *GatewayUpdateLogForwardingLogzIo) SetEnable(v string) {
	o.Enable = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogzIo) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogzIo) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogzIo) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateLogForwardingLogzIo) SetJson(v bool) {
	o.Json = &v
}

// GetLogzIoToken returns the LogzIoToken field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogzIo) GetLogzIoToken() string {
	if o == nil || o.LogzIoToken == nil {
		var ret string
		return ret
	}
	return *o.LogzIoToken
}

// GetLogzIoTokenOk returns a tuple with the LogzIoToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogzIo) GetLogzIoTokenOk() (*string, bool) {
	if o == nil || o.LogzIoToken == nil {
		return nil, false
	}
	return o.LogzIoToken, true
}

// HasLogzIoToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogzIo) HasLogzIoToken() bool {
	if o != nil && o.LogzIoToken != nil {
		return true
	}

	return false
}

// SetLogzIoToken gets a reference to the given string and assigns it to the LogzIoToken field.
func (o *GatewayUpdateLogForwardingLogzIo) SetLogzIoToken(v string) {
	o.LogzIoToken = &v
}

// GetOutputFormat returns the OutputFormat field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogzIo) GetOutputFormat() string {
	if o == nil || o.OutputFormat == nil {
		var ret string
		return ret
	}
	return *o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogzIo) GetOutputFormatOk() (*string, bool) {
	if o == nil || o.OutputFormat == nil {
		return nil, false
	}
	return o.OutputFormat, true
}

// HasOutputFormat returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogzIo) HasOutputFormat() bool {
	if o != nil && o.OutputFormat != nil {
		return true
	}

	return false
}

// SetOutputFormat gets a reference to the given string and assigns it to the OutputFormat field.
func (o *GatewayUpdateLogForwardingLogzIo) SetOutputFormat(v string) {
	o.OutputFormat = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogzIo) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogzIo) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogzIo) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GatewayUpdateLogForwardingLogzIo) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPullInterval returns the PullInterval field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogzIo) GetPullInterval() string {
	if o == nil || o.PullInterval == nil {
		var ret string
		return ret
	}
	return *o.PullInterval
}

// GetPullIntervalOk returns a tuple with the PullInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogzIo) GetPullIntervalOk() (*string, bool) {
	if o == nil || o.PullInterval == nil {
		return nil, false
	}
	return o.PullInterval, true
}

// HasPullInterval returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogzIo) HasPullInterval() bool {
	if o != nil && o.PullInterval != nil {
		return true
	}

	return false
}

// SetPullInterval gets a reference to the given string and assigns it to the PullInterval field.
func (o *GatewayUpdateLogForwardingLogzIo) SetPullInterval(v string) {
	o.PullInterval = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogzIo) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogzIo) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogzIo) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateLogForwardingLogzIo) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogzIo) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogzIo) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogzIo) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateLogForwardingLogzIo) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayUpdateLogForwardingLogzIo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.LogzIoToken != nil {
		toSerialize["logz-io-token"] = o.LogzIoToken
	}
	if o.OutputFormat != nil {
		toSerialize["output-format"] = o.OutputFormat
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.PullInterval != nil {
		toSerialize["pull-interval"] = o.PullInterval
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateLogForwardingLogzIo struct {
	value *GatewayUpdateLogForwardingLogzIo
	isSet bool
}

func (v NullableGatewayUpdateLogForwardingLogzIo) Get() *GatewayUpdateLogForwardingLogzIo {
	return v.value
}

func (v *NullableGatewayUpdateLogForwardingLogzIo) Set(val *GatewayUpdateLogForwardingLogzIo) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateLogForwardingLogzIo) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateLogForwardingLogzIo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateLogForwardingLogzIo(val *GatewayUpdateLogForwardingLogzIo) *NullableGatewayUpdateLogForwardingLogzIo {
	return &NullableGatewayUpdateLogForwardingLogzIo{value: val, isSet: true}
}

func (v NullableGatewayUpdateLogForwardingLogzIo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateLogForwardingLogzIo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


