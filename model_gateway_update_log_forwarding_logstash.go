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

// GatewayUpdateLogForwardingLogstash gatewayUpdateLogForwardingLogstash is a command that updates log forwarding config (logstash target)
type GatewayUpdateLogForwardingLogstash struct {
	// Logstash dns
	Dns *string `json:"dns,omitempty"`
	// Enable Log Forwarding [true/false]
	Enable *string `json:"enable,omitempty"`
	// Enable tls
	EnableTls *bool `json:"enable-tls,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Logs format [text/json]
	OutputFormat *string `json:"output-format,omitempty"`
	// Logstash protocol [tcp/udp]
	Protocol *string `json:"protocol,omitempty"`
	// Pull interval in seconds
	PullInterval *string `json:"pull-interval,omitempty"`
	// Logstash tls certificate
	TlsCertificate *string `json:"tls-certificate,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayUpdateLogForwardingLogstash instantiates a new GatewayUpdateLogForwardingLogstash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateLogForwardingLogstash() *GatewayUpdateLogForwardingLogstash {
	this := GatewayUpdateLogForwardingLogstash{}
	var enable string = "true"
	this.Enable = &enable
	var json bool = false
	this.Json = &json
	var outputFormat string = "text"
	this.OutputFormat = &outputFormat
	var pullInterval string = "10"
	this.PullInterval = &pullInterval
	var tlsCertificate string = "use-existing"
	this.TlsCertificate = &tlsCertificate
	return &this
}

// NewGatewayUpdateLogForwardingLogstashWithDefaults instantiates a new GatewayUpdateLogForwardingLogstash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateLogForwardingLogstashWithDefaults() *GatewayUpdateLogForwardingLogstash {
	this := GatewayUpdateLogForwardingLogstash{}
	var enable string = "true"
	this.Enable = &enable
	var json bool = false
	this.Json = &json
	var outputFormat string = "text"
	this.OutputFormat = &outputFormat
	var pullInterval string = "10"
	this.PullInterval = &pullInterval
	var tlsCertificate string = "use-existing"
	this.TlsCertificate = &tlsCertificate
	return &this
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetDns() string {
	if o == nil || o.Dns == nil {
		var ret string
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetDnsOk() (*string, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given string and assigns it to the Dns field.
func (o *GatewayUpdateLogForwardingLogstash) SetDns(v string) {
	o.Dns = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetEnable() string {
	if o == nil || o.Enable == nil {
		var ret string
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetEnableOk() (*string, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given string and assigns it to the Enable field.
func (o *GatewayUpdateLogForwardingLogstash) SetEnable(v string) {
	o.Enable = &v
}

// GetEnableTls returns the EnableTls field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetEnableTls() bool {
	if o == nil || o.EnableTls == nil {
		var ret bool
		return ret
	}
	return *o.EnableTls
}

// GetEnableTlsOk returns a tuple with the EnableTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetEnableTlsOk() (*bool, bool) {
	if o == nil || o.EnableTls == nil {
		return nil, false
	}
	return o.EnableTls, true
}

// HasEnableTls returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasEnableTls() bool {
	if o != nil && o.EnableTls != nil {
		return true
	}

	return false
}

// SetEnableTls gets a reference to the given bool and assigns it to the EnableTls field.
func (o *GatewayUpdateLogForwardingLogstash) SetEnableTls(v bool) {
	o.EnableTls = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateLogForwardingLogstash) SetJson(v bool) {
	o.Json = &v
}

// GetOutputFormat returns the OutputFormat field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetOutputFormat() string {
	if o == nil || o.OutputFormat == nil {
		var ret string
		return ret
	}
	return *o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetOutputFormatOk() (*string, bool) {
	if o == nil || o.OutputFormat == nil {
		return nil, false
	}
	return o.OutputFormat, true
}

// HasOutputFormat returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasOutputFormat() bool {
	if o != nil && o.OutputFormat != nil {
		return true
	}

	return false
}

// SetOutputFormat gets a reference to the given string and assigns it to the OutputFormat field.
func (o *GatewayUpdateLogForwardingLogstash) SetOutputFormat(v string) {
	o.OutputFormat = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GatewayUpdateLogForwardingLogstash) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPullInterval returns the PullInterval field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetPullInterval() string {
	if o == nil || o.PullInterval == nil {
		var ret string
		return ret
	}
	return *o.PullInterval
}

// GetPullIntervalOk returns a tuple with the PullInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetPullIntervalOk() (*string, bool) {
	if o == nil || o.PullInterval == nil {
		return nil, false
	}
	return o.PullInterval, true
}

// HasPullInterval returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasPullInterval() bool {
	if o != nil && o.PullInterval != nil {
		return true
	}

	return false
}

// SetPullInterval gets a reference to the given string and assigns it to the PullInterval field.
func (o *GatewayUpdateLogForwardingLogstash) SetPullInterval(v string) {
	o.PullInterval = &v
}

// GetTlsCertificate returns the TlsCertificate field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetTlsCertificate() string {
	if o == nil || o.TlsCertificate == nil {
		var ret string
		return ret
	}
	return *o.TlsCertificate
}

// GetTlsCertificateOk returns a tuple with the TlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetTlsCertificateOk() (*string, bool) {
	if o == nil || o.TlsCertificate == nil {
		return nil, false
	}
	return o.TlsCertificate, true
}

// HasTlsCertificate returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasTlsCertificate() bool {
	if o != nil && o.TlsCertificate != nil {
		return true
	}

	return false
}

// SetTlsCertificate gets a reference to the given string and assigns it to the TlsCertificate field.
func (o *GatewayUpdateLogForwardingLogstash) SetTlsCertificate(v string) {
	o.TlsCertificate = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateLogForwardingLogstash) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingLogstash) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingLogstash) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingLogstash) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateLogForwardingLogstash) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayUpdateLogForwardingLogstash) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.EnableTls != nil {
		toSerialize["enable-tls"] = o.EnableTls
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
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
	if o.TlsCertificate != nil {
		toSerialize["tls-certificate"] = o.TlsCertificate
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateLogForwardingLogstash struct {
	value *GatewayUpdateLogForwardingLogstash
	isSet bool
}

func (v NullableGatewayUpdateLogForwardingLogstash) Get() *GatewayUpdateLogForwardingLogstash {
	return v.value
}

func (v *NullableGatewayUpdateLogForwardingLogstash) Set(val *GatewayUpdateLogForwardingLogstash) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateLogForwardingLogstash) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateLogForwardingLogstash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateLogForwardingLogstash(val *GatewayUpdateLogForwardingLogstash) *NullableGatewayUpdateLogForwardingLogstash {
	return &NullableGatewayUpdateLogForwardingLogstash{value: val, isSet: true}
}

func (v NullableGatewayUpdateLogForwardingLogstash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateLogForwardingLogstash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


