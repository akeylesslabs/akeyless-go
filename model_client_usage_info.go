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

// ClientUsageInfo struct for ClientUsageInfo
type ClientUsageInfo struct {
	AccessId *string `json:"access_id,omitempty"`
	AccessType *string `json:"access_type,omitempty"`
	AuthMethodName *string `json:"auth_method_name,omitempty"`
	ClientUniqueId *string `json:"client_unique_id,omitempty"`
	ExceededClients *int64 `json:"exceeded_clients,omitempty"`
}

// NewClientUsageInfo instantiates a new ClientUsageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientUsageInfo() *ClientUsageInfo {
	this := ClientUsageInfo{}
	return &this
}

// NewClientUsageInfoWithDefaults instantiates a new ClientUsageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientUsageInfoWithDefaults() *ClientUsageInfo {
	this := ClientUsageInfo{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *ClientUsageInfo) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientUsageInfo) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *ClientUsageInfo) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *ClientUsageInfo) SetAccessId(v string) {
	o.AccessId = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *ClientUsageInfo) GetAccessType() string {
	if o == nil || o.AccessType == nil {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientUsageInfo) GetAccessTypeOk() (*string, bool) {
	if o == nil || o.AccessType == nil {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *ClientUsageInfo) HasAccessType() bool {
	if o != nil && o.AccessType != nil {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *ClientUsageInfo) SetAccessType(v string) {
	o.AccessType = &v
}

// GetAuthMethodName returns the AuthMethodName field value if set, zero value otherwise.
func (o *ClientUsageInfo) GetAuthMethodName() string {
	if o == nil || o.AuthMethodName == nil {
		var ret string
		return ret
	}
	return *o.AuthMethodName
}

// GetAuthMethodNameOk returns a tuple with the AuthMethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientUsageInfo) GetAuthMethodNameOk() (*string, bool) {
	if o == nil || o.AuthMethodName == nil {
		return nil, false
	}
	return o.AuthMethodName, true
}

// HasAuthMethodName returns a boolean if a field has been set.
func (o *ClientUsageInfo) HasAuthMethodName() bool {
	if o != nil && o.AuthMethodName != nil {
		return true
	}

	return false
}

// SetAuthMethodName gets a reference to the given string and assigns it to the AuthMethodName field.
func (o *ClientUsageInfo) SetAuthMethodName(v string) {
	o.AuthMethodName = &v
}

// GetClientUniqueId returns the ClientUniqueId field value if set, zero value otherwise.
func (o *ClientUsageInfo) GetClientUniqueId() string {
	if o == nil || o.ClientUniqueId == nil {
		var ret string
		return ret
	}
	return *o.ClientUniqueId
}

// GetClientUniqueIdOk returns a tuple with the ClientUniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientUsageInfo) GetClientUniqueIdOk() (*string, bool) {
	if o == nil || o.ClientUniqueId == nil {
		return nil, false
	}
	return o.ClientUniqueId, true
}

// HasClientUniqueId returns a boolean if a field has been set.
func (o *ClientUsageInfo) HasClientUniqueId() bool {
	if o != nil && o.ClientUniqueId != nil {
		return true
	}

	return false
}

// SetClientUniqueId gets a reference to the given string and assigns it to the ClientUniqueId field.
func (o *ClientUsageInfo) SetClientUniqueId(v string) {
	o.ClientUniqueId = &v
}

// GetExceededClients returns the ExceededClients field value if set, zero value otherwise.
func (o *ClientUsageInfo) GetExceededClients() int64 {
	if o == nil || o.ExceededClients == nil {
		var ret int64
		return ret
	}
	return *o.ExceededClients
}

// GetExceededClientsOk returns a tuple with the ExceededClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientUsageInfo) GetExceededClientsOk() (*int64, bool) {
	if o == nil || o.ExceededClients == nil {
		return nil, false
	}
	return o.ExceededClients, true
}

// HasExceededClients returns a boolean if a field has been set.
func (o *ClientUsageInfo) HasExceededClients() bool {
	if o != nil && o.ExceededClients != nil {
		return true
	}

	return false
}

// SetExceededClients gets a reference to the given int64 and assigns it to the ExceededClients field.
func (o *ClientUsageInfo) SetExceededClients(v int64) {
	o.ExceededClients = &v
}

func (o ClientUsageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access_id"] = o.AccessId
	}
	if o.AccessType != nil {
		toSerialize["access_type"] = o.AccessType
	}
	if o.AuthMethodName != nil {
		toSerialize["auth_method_name"] = o.AuthMethodName
	}
	if o.ClientUniqueId != nil {
		toSerialize["client_unique_id"] = o.ClientUniqueId
	}
	if o.ExceededClients != nil {
		toSerialize["exceeded_clients"] = o.ExceededClients
	}
	return json.Marshal(toSerialize)
}

type NullableClientUsageInfo struct {
	value *ClientUsageInfo
	isSet bool
}

func (v NullableClientUsageInfo) Get() *ClientUsageInfo {
	return v.value
}

func (v *NullableClientUsageInfo) Set(val *ClientUsageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableClientUsageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableClientUsageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientUsageInfo(val *ClientUsageInfo) *NullableClientUsageInfo {
	return &NullableClientUsageInfo{value: val, isSet: true}
}

func (v NullableClientUsageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientUsageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

