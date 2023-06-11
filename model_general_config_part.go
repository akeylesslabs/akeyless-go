/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// checks if the GeneralConfigPart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneralConfigPart{}

// GeneralConfigPart struct for GeneralConfigPart
type GeneralConfigPart struct {
	// AkeylessUrl is here for BC only. Gator will still return it if it exists in the configuration, but new clients (>=2.34.0) will ignore it and override it with what exists in their local file. It will no longer be sent to Gator for update, so new clusters will only have the default value saved in the DB.
	AkeylessUrl *string `json:"akeyless_url,omitempty"`
	ApiTokenTtl *string `json:"api_token_ttl,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	EnableSniProxy *bool `json:"enable_sni_proxy,omitempty"`
	EnableTls *bool `json:"enable_tls,omitempty"`
	EnableTlsConfigure *bool `json:"enable_tls_configure,omitempty"`
	EnableTlsCurl *bool `json:"enable_tls_curl,omitempty"`
	EnableTlsHvp *bool `json:"enable_tls_hvp,omitempty"`
	GwClusterUrl *string `json:"gw_cluster_url,omitempty"`
	TcpPort *string `json:"tcp_port,omitempty"`
	TlsCert *string `json:"tls_cert,omitempty"`
	TlsKey *string `json:"tls_key,omitempty"`
}

// NewGeneralConfigPart instantiates a new GeneralConfigPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralConfigPart() *GeneralConfigPart {
	this := GeneralConfigPart{}
	return &this
}

// NewGeneralConfigPartWithDefaults instantiates a new GeneralConfigPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralConfigPartWithDefaults() *GeneralConfigPart {
	this := GeneralConfigPart{}
	return &this
}

// GetAkeylessUrl returns the AkeylessUrl field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetAkeylessUrl() string {
	if o == nil || IsNil(o.AkeylessUrl) {
		var ret string
		return ret
	}
	return *o.AkeylessUrl
}

// GetAkeylessUrlOk returns a tuple with the AkeylessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetAkeylessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AkeylessUrl) {
		return nil, false
	}
	return o.AkeylessUrl, true
}

// HasAkeylessUrl returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasAkeylessUrl() bool {
	if o != nil && !IsNil(o.AkeylessUrl) {
		return true
	}

	return false
}

// SetAkeylessUrl gets a reference to the given string and assigns it to the AkeylessUrl field.
func (o *GeneralConfigPart) SetAkeylessUrl(v string) {
	o.AkeylessUrl = &v
}

// GetApiTokenTtl returns the ApiTokenTtl field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetApiTokenTtl() string {
	if o == nil || IsNil(o.ApiTokenTtl) {
		var ret string
		return ret
	}
	return *o.ApiTokenTtl
}

// GetApiTokenTtlOk returns a tuple with the ApiTokenTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetApiTokenTtlOk() (*string, bool) {
	if o == nil || IsNil(o.ApiTokenTtl) {
		return nil, false
	}
	return o.ApiTokenTtl, true
}

// HasApiTokenTtl returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasApiTokenTtl() bool {
	if o != nil && !IsNil(o.ApiTokenTtl) {
		return true
	}

	return false
}

// SetApiTokenTtl gets a reference to the given string and assigns it to the ApiTokenTtl field.
func (o *GeneralConfigPart) SetApiTokenTtl(v string) {
	o.ApiTokenTtl = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GeneralConfigPart) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnableSniProxy returns the EnableSniProxy field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetEnableSniProxy() bool {
	if o == nil || IsNil(o.EnableSniProxy) {
		var ret bool
		return ret
	}
	return *o.EnableSniProxy
}

// GetEnableSniProxyOk returns a tuple with the EnableSniProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetEnableSniProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSniProxy) {
		return nil, false
	}
	return o.EnableSniProxy, true
}

// HasEnableSniProxy returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasEnableSniProxy() bool {
	if o != nil && !IsNil(o.EnableSniProxy) {
		return true
	}

	return false
}

// SetEnableSniProxy gets a reference to the given bool and assigns it to the EnableSniProxy field.
func (o *GeneralConfigPart) SetEnableSniProxy(v bool) {
	o.EnableSniProxy = &v
}

// GetEnableTls returns the EnableTls field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetEnableTls() bool {
	if o == nil || IsNil(o.EnableTls) {
		var ret bool
		return ret
	}
	return *o.EnableTls
}

// GetEnableTlsOk returns a tuple with the EnableTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetEnableTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTls) {
		return nil, false
	}
	return o.EnableTls, true
}

// HasEnableTls returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasEnableTls() bool {
	if o != nil && !IsNil(o.EnableTls) {
		return true
	}

	return false
}

// SetEnableTls gets a reference to the given bool and assigns it to the EnableTls field.
func (o *GeneralConfigPart) SetEnableTls(v bool) {
	o.EnableTls = &v
}

// GetEnableTlsConfigure returns the EnableTlsConfigure field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetEnableTlsConfigure() bool {
	if o == nil || IsNil(o.EnableTlsConfigure) {
		var ret bool
		return ret
	}
	return *o.EnableTlsConfigure
}

// GetEnableTlsConfigureOk returns a tuple with the EnableTlsConfigure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetEnableTlsConfigureOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTlsConfigure) {
		return nil, false
	}
	return o.EnableTlsConfigure, true
}

// HasEnableTlsConfigure returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasEnableTlsConfigure() bool {
	if o != nil && !IsNil(o.EnableTlsConfigure) {
		return true
	}

	return false
}

// SetEnableTlsConfigure gets a reference to the given bool and assigns it to the EnableTlsConfigure field.
func (o *GeneralConfigPart) SetEnableTlsConfigure(v bool) {
	o.EnableTlsConfigure = &v
}

// GetEnableTlsCurl returns the EnableTlsCurl field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetEnableTlsCurl() bool {
	if o == nil || IsNil(o.EnableTlsCurl) {
		var ret bool
		return ret
	}
	return *o.EnableTlsCurl
}

// GetEnableTlsCurlOk returns a tuple with the EnableTlsCurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetEnableTlsCurlOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTlsCurl) {
		return nil, false
	}
	return o.EnableTlsCurl, true
}

// HasEnableTlsCurl returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasEnableTlsCurl() bool {
	if o != nil && !IsNil(o.EnableTlsCurl) {
		return true
	}

	return false
}

// SetEnableTlsCurl gets a reference to the given bool and assigns it to the EnableTlsCurl field.
func (o *GeneralConfigPart) SetEnableTlsCurl(v bool) {
	o.EnableTlsCurl = &v
}

// GetEnableTlsHvp returns the EnableTlsHvp field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetEnableTlsHvp() bool {
	if o == nil || IsNil(o.EnableTlsHvp) {
		var ret bool
		return ret
	}
	return *o.EnableTlsHvp
}

// GetEnableTlsHvpOk returns a tuple with the EnableTlsHvp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetEnableTlsHvpOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTlsHvp) {
		return nil, false
	}
	return o.EnableTlsHvp, true
}

// HasEnableTlsHvp returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasEnableTlsHvp() bool {
	if o != nil && !IsNil(o.EnableTlsHvp) {
		return true
	}

	return false
}

// SetEnableTlsHvp gets a reference to the given bool and assigns it to the EnableTlsHvp field.
func (o *GeneralConfigPart) SetEnableTlsHvp(v bool) {
	o.EnableTlsHvp = &v
}

// GetGwClusterUrl returns the GwClusterUrl field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetGwClusterUrl() string {
	if o == nil || IsNil(o.GwClusterUrl) {
		var ret string
		return ret
	}
	return *o.GwClusterUrl
}

// GetGwClusterUrlOk returns a tuple with the GwClusterUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetGwClusterUrlOk() (*string, bool) {
	if o == nil || IsNil(o.GwClusterUrl) {
		return nil, false
	}
	return o.GwClusterUrl, true
}

// HasGwClusterUrl returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasGwClusterUrl() bool {
	if o != nil && !IsNil(o.GwClusterUrl) {
		return true
	}

	return false
}

// SetGwClusterUrl gets a reference to the given string and assigns it to the GwClusterUrl field.
func (o *GeneralConfigPart) SetGwClusterUrl(v string) {
	o.GwClusterUrl = &v
}

// GetTcpPort returns the TcpPort field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetTcpPort() string {
	if o == nil || IsNil(o.TcpPort) {
		var ret string
		return ret
	}
	return *o.TcpPort
}

// GetTcpPortOk returns a tuple with the TcpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetTcpPortOk() (*string, bool) {
	if o == nil || IsNil(o.TcpPort) {
		return nil, false
	}
	return o.TcpPort, true
}

// HasTcpPort returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasTcpPort() bool {
	if o != nil && !IsNil(o.TcpPort) {
		return true
	}

	return false
}

// SetTcpPort gets a reference to the given string and assigns it to the TcpPort field.
func (o *GeneralConfigPart) SetTcpPort(v string) {
	o.TcpPort = &v
}

// GetTlsCert returns the TlsCert field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetTlsCert() string {
	if o == nil || IsNil(o.TlsCert) {
		var ret string
		return ret
	}
	return *o.TlsCert
}

// GetTlsCertOk returns a tuple with the TlsCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetTlsCertOk() (*string, bool) {
	if o == nil || IsNil(o.TlsCert) {
		return nil, false
	}
	return o.TlsCert, true
}

// HasTlsCert returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasTlsCert() bool {
	if o != nil && !IsNil(o.TlsCert) {
		return true
	}

	return false
}

// SetTlsCert gets a reference to the given string and assigns it to the TlsCert field.
func (o *GeneralConfigPart) SetTlsCert(v string) {
	o.TlsCert = &v
}

// GetTlsKey returns the TlsKey field value if set, zero value otherwise.
func (o *GeneralConfigPart) GetTlsKey() string {
	if o == nil || IsNil(o.TlsKey) {
		var ret string
		return ret
	}
	return *o.TlsKey
}

// GetTlsKeyOk returns a tuple with the TlsKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralConfigPart) GetTlsKeyOk() (*string, bool) {
	if o == nil || IsNil(o.TlsKey) {
		return nil, false
	}
	return o.TlsKey, true
}

// HasTlsKey returns a boolean if a field has been set.
func (o *GeneralConfigPart) HasTlsKey() bool {
	if o != nil && !IsNil(o.TlsKey) {
		return true
	}

	return false
}

// SetTlsKey gets a reference to the given string and assigns it to the TlsKey field.
func (o *GeneralConfigPart) SetTlsKey(v string) {
	o.TlsKey = &v
}

func (o GeneralConfigPart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneralConfigPart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AkeylessUrl) {
		toSerialize["akeyless_url"] = o.AkeylessUrl
	}
	if !IsNil(o.ApiTokenTtl) {
		toSerialize["api_token_ttl"] = o.ApiTokenTtl
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.EnableSniProxy) {
		toSerialize["enable_sni_proxy"] = o.EnableSniProxy
	}
	if !IsNil(o.EnableTls) {
		toSerialize["enable_tls"] = o.EnableTls
	}
	if !IsNil(o.EnableTlsConfigure) {
		toSerialize["enable_tls_configure"] = o.EnableTlsConfigure
	}
	if !IsNil(o.EnableTlsCurl) {
		toSerialize["enable_tls_curl"] = o.EnableTlsCurl
	}
	if !IsNil(o.EnableTlsHvp) {
		toSerialize["enable_tls_hvp"] = o.EnableTlsHvp
	}
	if !IsNil(o.GwClusterUrl) {
		toSerialize["gw_cluster_url"] = o.GwClusterUrl
	}
	if !IsNil(o.TcpPort) {
		toSerialize["tcp_port"] = o.TcpPort
	}
	if !IsNil(o.TlsCert) {
		toSerialize["tls_cert"] = o.TlsCert
	}
	if !IsNil(o.TlsKey) {
		toSerialize["tls_key"] = o.TlsKey
	}
	return toSerialize, nil
}

type NullableGeneralConfigPart struct {
	value *GeneralConfigPart
	isSet bool
}

func (v NullableGeneralConfigPart) Get() *GeneralConfigPart {
	return v.value
}

func (v *NullableGeneralConfigPart) Set(val *GeneralConfigPart) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralConfigPart) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralConfigPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralConfigPart(val *GeneralConfigPart) *NullableGeneralConfigPart {
	return &NullableGeneralConfigPart{value: val, isSet: true}
}

func (v NullableGeneralConfigPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralConfigPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


