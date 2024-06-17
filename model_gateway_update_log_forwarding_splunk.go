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

// GatewayUpdateLogForwardingSplunk gatewayUpdateLogForwardingSplunk is a command that updates log forwarding config (splunk target)
type GatewayUpdateLogForwardingSplunk struct {
	// Enable Log Forwarding [true/false]
	Enable *string `json:"enable,omitempty"`
	// Enable tls
	EnableTls *bool `json:"enable-tls,omitempty"`
	// Splunk index
	Index *string `json:"index,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Logs format [text/json]
	OutputFormat *string `json:"output-format,omitempty"`
	// Pull interval in seconds
	PullInterval *string `json:"pull-interval,omitempty"`
	// Splunk source
	Source *string `json:"source,omitempty"`
	// Splunk source type
	SourceType *string `json:"source-type,omitempty"`
	// Splunk token
	SplunkToken *string `json:"splunk-token,omitempty"`
	// Splunk server URL
	SplunkUrl *string `json:"splunk-url,omitempty"`
	// Splunk tls certificate
	TlsCertificate *string `json:"tls-certificate,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayUpdateLogForwardingSplunk instantiates a new GatewayUpdateLogForwardingSplunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateLogForwardingSplunk() *GatewayUpdateLogForwardingSplunk {
	this := GatewayUpdateLogForwardingSplunk{}
	var enable string = "true"
	this.Enable = &enable
	var json bool = false
	this.Json = &json
	var outputFormat string = "text"
	this.OutputFormat = &outputFormat
	var pullInterval string = "10"
	this.PullInterval = &pullInterval
	var source string = "use-existing"
	this.Source = &source
	var sourceType string = "use-existing"
	this.SourceType = &sourceType
	var tlsCertificate string = "use-existing"
	this.TlsCertificate = &tlsCertificate
	return &this
}

// NewGatewayUpdateLogForwardingSplunkWithDefaults instantiates a new GatewayUpdateLogForwardingSplunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateLogForwardingSplunkWithDefaults() *GatewayUpdateLogForwardingSplunk {
	this := GatewayUpdateLogForwardingSplunk{}
	var enable string = "true"
	this.Enable = &enable
	var json bool = false
	this.Json = &json
	var outputFormat string = "text"
	this.OutputFormat = &outputFormat
	var pullInterval string = "10"
	this.PullInterval = &pullInterval
	var source string = "use-existing"
	this.Source = &source
	var sourceType string = "use-existing"
	this.SourceType = &sourceType
	var tlsCertificate string = "use-existing"
	this.TlsCertificate = &tlsCertificate
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetEnable() string {
	if o == nil || o.Enable == nil {
		var ret string
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetEnableOk() (*string, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given string and assigns it to the Enable field.
func (o *GatewayUpdateLogForwardingSplunk) SetEnable(v string) {
	o.Enable = &v
}

// GetEnableTls returns the EnableTls field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetEnableTls() bool {
	if o == nil || o.EnableTls == nil {
		var ret bool
		return ret
	}
	return *o.EnableTls
}

// GetEnableTlsOk returns a tuple with the EnableTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetEnableTlsOk() (*bool, bool) {
	if o == nil || o.EnableTls == nil {
		return nil, false
	}
	return o.EnableTls, true
}

// HasEnableTls returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasEnableTls() bool {
	if o != nil && o.EnableTls != nil {
		return true
	}

	return false
}

// SetEnableTls gets a reference to the given bool and assigns it to the EnableTls field.
func (o *GatewayUpdateLogForwardingSplunk) SetEnableTls(v bool) {
	o.EnableTls = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *GatewayUpdateLogForwardingSplunk) SetIndex(v string) {
	o.Index = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateLogForwardingSplunk) SetJson(v bool) {
	o.Json = &v
}

// GetOutputFormat returns the OutputFormat field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetOutputFormat() string {
	if o == nil || o.OutputFormat == nil {
		var ret string
		return ret
	}
	return *o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetOutputFormatOk() (*string, bool) {
	if o == nil || o.OutputFormat == nil {
		return nil, false
	}
	return o.OutputFormat, true
}

// HasOutputFormat returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasOutputFormat() bool {
	if o != nil && o.OutputFormat != nil {
		return true
	}

	return false
}

// SetOutputFormat gets a reference to the given string and assigns it to the OutputFormat field.
func (o *GatewayUpdateLogForwardingSplunk) SetOutputFormat(v string) {
	o.OutputFormat = &v
}

// GetPullInterval returns the PullInterval field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetPullInterval() string {
	if o == nil || o.PullInterval == nil {
		var ret string
		return ret
	}
	return *o.PullInterval
}

// GetPullIntervalOk returns a tuple with the PullInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetPullIntervalOk() (*string, bool) {
	if o == nil || o.PullInterval == nil {
		return nil, false
	}
	return o.PullInterval, true
}

// HasPullInterval returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasPullInterval() bool {
	if o != nil && o.PullInterval != nil {
		return true
	}

	return false
}

// SetPullInterval gets a reference to the given string and assigns it to the PullInterval field.
func (o *GatewayUpdateLogForwardingSplunk) SetPullInterval(v string) {
	o.PullInterval = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *GatewayUpdateLogForwardingSplunk) SetSource(v string) {
	o.Source = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetSourceType() string {
	if o == nil || o.SourceType == nil {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetSourceTypeOk() (*string, bool) {
	if o == nil || o.SourceType == nil {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasSourceType() bool {
	if o != nil && o.SourceType != nil {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *GatewayUpdateLogForwardingSplunk) SetSourceType(v string) {
	o.SourceType = &v
}

// GetSplunkToken returns the SplunkToken field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetSplunkToken() string {
	if o == nil || o.SplunkToken == nil {
		var ret string
		return ret
	}
	return *o.SplunkToken
}

// GetSplunkTokenOk returns a tuple with the SplunkToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetSplunkTokenOk() (*string, bool) {
	if o == nil || o.SplunkToken == nil {
		return nil, false
	}
	return o.SplunkToken, true
}

// HasSplunkToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasSplunkToken() bool {
	if o != nil && o.SplunkToken != nil {
		return true
	}

	return false
}

// SetSplunkToken gets a reference to the given string and assigns it to the SplunkToken field.
func (o *GatewayUpdateLogForwardingSplunk) SetSplunkToken(v string) {
	o.SplunkToken = &v
}

// GetSplunkUrl returns the SplunkUrl field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetSplunkUrl() string {
	if o == nil || o.SplunkUrl == nil {
		var ret string
		return ret
	}
	return *o.SplunkUrl
}

// GetSplunkUrlOk returns a tuple with the SplunkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetSplunkUrlOk() (*string, bool) {
	if o == nil || o.SplunkUrl == nil {
		return nil, false
	}
	return o.SplunkUrl, true
}

// HasSplunkUrl returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasSplunkUrl() bool {
	if o != nil && o.SplunkUrl != nil {
		return true
	}

	return false
}

// SetSplunkUrl gets a reference to the given string and assigns it to the SplunkUrl field.
func (o *GatewayUpdateLogForwardingSplunk) SetSplunkUrl(v string) {
	o.SplunkUrl = &v
}

// GetTlsCertificate returns the TlsCertificate field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetTlsCertificate() string {
	if o == nil || o.TlsCertificate == nil {
		var ret string
		return ret
	}
	return *o.TlsCertificate
}

// GetTlsCertificateOk returns a tuple with the TlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetTlsCertificateOk() (*string, bool) {
	if o == nil || o.TlsCertificate == nil {
		return nil, false
	}
	return o.TlsCertificate, true
}

// HasTlsCertificate returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasTlsCertificate() bool {
	if o != nil && o.TlsCertificate != nil {
		return true
	}

	return false
}

// SetTlsCertificate gets a reference to the given string and assigns it to the TlsCertificate field.
func (o *GatewayUpdateLogForwardingSplunk) SetTlsCertificate(v string) {
	o.TlsCertificate = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateLogForwardingSplunk) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateLogForwardingSplunk) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLogForwardingSplunk) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateLogForwardingSplunk) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateLogForwardingSplunk) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayUpdateLogForwardingSplunk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.EnableTls != nil {
		toSerialize["enable-tls"] = o.EnableTls
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.OutputFormat != nil {
		toSerialize["output-format"] = o.OutputFormat
	}
	if o.PullInterval != nil {
		toSerialize["pull-interval"] = o.PullInterval
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.SourceType != nil {
		toSerialize["source-type"] = o.SourceType
	}
	if o.SplunkToken != nil {
		toSerialize["splunk-token"] = o.SplunkToken
	}
	if o.SplunkUrl != nil {
		toSerialize["splunk-url"] = o.SplunkUrl
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

type NullableGatewayUpdateLogForwardingSplunk struct {
	value *GatewayUpdateLogForwardingSplunk
	isSet bool
}

func (v NullableGatewayUpdateLogForwardingSplunk) Get() *GatewayUpdateLogForwardingSplunk {
	return v.value
}

func (v *NullableGatewayUpdateLogForwardingSplunk) Set(val *GatewayUpdateLogForwardingSplunk) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateLogForwardingSplunk) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateLogForwardingSplunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateLogForwardingSplunk(val *GatewayUpdateLogForwardingSplunk) *NullableGatewayUpdateLogForwardingSplunk {
	return &NullableGatewayUpdateLogForwardingSplunk{value: val, isSet: true}
}

func (v NullableGatewayUpdateLogForwardingSplunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateLogForwardingSplunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


