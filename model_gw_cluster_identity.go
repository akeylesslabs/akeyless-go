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

// GwClusterIdentity struct for GwClusterIdentity
type GwClusterIdentity struct {
	Allowed *bool `json:"allowed,omitempty"`
	AllowedAccessIds *[]string `json:"allowed_access_ids,omitempty"`
	ClusterName *string `json:"cluster_name,omitempty"`
	ClusterUrl *string `json:"cluster_url,omitempty"`
	CurrentGw *bool `json:"current_gw,omitempty"`
	// Deprecated - use CustomerFragments instead
	CustomerFragmentIds *[]string `json:"customer_fragment_ids,omitempty"`
	CustomerFragments *[]CfInfo `json:"customer_fragments,omitempty"`
	DefaultProtectionKeyId *int64 `json:"default_protection_key_id,omitempty"`
	DefaultSecretLocation *string `json:"default_secret_location,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusDescription *string `json:"status_description,omitempty"`
}

// NewGwClusterIdentity instantiates a new GwClusterIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGwClusterIdentity() *GwClusterIdentity {
	this := GwClusterIdentity{}
	return &this
}

// NewGwClusterIdentityWithDefaults instantiates a new GwClusterIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGwClusterIdentityWithDefaults() *GwClusterIdentity {
	this := GwClusterIdentity{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetAllowed() bool {
	if o == nil || o.Allowed == nil {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetAllowedOk() (*bool, bool) {
	if o == nil || o.Allowed == nil {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasAllowed() bool {
	if o != nil && o.Allowed != nil {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *GwClusterIdentity) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetAllowedAccessIds returns the AllowedAccessIds field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetAllowedAccessIds() []string {
	if o == nil || o.AllowedAccessIds == nil {
		var ret []string
		return ret
	}
	return *o.AllowedAccessIds
}

// GetAllowedAccessIdsOk returns a tuple with the AllowedAccessIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetAllowedAccessIdsOk() (*[]string, bool) {
	if o == nil || o.AllowedAccessIds == nil {
		return nil, false
	}
	return o.AllowedAccessIds, true
}

// HasAllowedAccessIds returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasAllowedAccessIds() bool {
	if o != nil && o.AllowedAccessIds != nil {
		return true
	}

	return false
}

// SetAllowedAccessIds gets a reference to the given []string and assigns it to the AllowedAccessIds field.
func (o *GwClusterIdentity) SetAllowedAccessIds(v []string) {
	o.AllowedAccessIds = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *GwClusterIdentity) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterUrl returns the ClusterUrl field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetClusterUrl() string {
	if o == nil || o.ClusterUrl == nil {
		var ret string
		return ret
	}
	return *o.ClusterUrl
}

// GetClusterUrlOk returns a tuple with the ClusterUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetClusterUrlOk() (*string, bool) {
	if o == nil || o.ClusterUrl == nil {
		return nil, false
	}
	return o.ClusterUrl, true
}

// HasClusterUrl returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasClusterUrl() bool {
	if o != nil && o.ClusterUrl != nil {
		return true
	}

	return false
}

// SetClusterUrl gets a reference to the given string and assigns it to the ClusterUrl field.
func (o *GwClusterIdentity) SetClusterUrl(v string) {
	o.ClusterUrl = &v
}

// GetCurrentGw returns the CurrentGw field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetCurrentGw() bool {
	if o == nil || o.CurrentGw == nil {
		var ret bool
		return ret
	}
	return *o.CurrentGw
}

// GetCurrentGwOk returns a tuple with the CurrentGw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetCurrentGwOk() (*bool, bool) {
	if o == nil || o.CurrentGw == nil {
		return nil, false
	}
	return o.CurrentGw, true
}

// HasCurrentGw returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasCurrentGw() bool {
	if o != nil && o.CurrentGw != nil {
		return true
	}

	return false
}

// SetCurrentGw gets a reference to the given bool and assigns it to the CurrentGw field.
func (o *GwClusterIdentity) SetCurrentGw(v bool) {
	o.CurrentGw = &v
}

// GetCustomerFragmentIds returns the CustomerFragmentIds field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetCustomerFragmentIds() []string {
	if o == nil || o.CustomerFragmentIds == nil {
		var ret []string
		return ret
	}
	return *o.CustomerFragmentIds
}

// GetCustomerFragmentIdsOk returns a tuple with the CustomerFragmentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetCustomerFragmentIdsOk() (*[]string, bool) {
	if o == nil || o.CustomerFragmentIds == nil {
		return nil, false
	}
	return o.CustomerFragmentIds, true
}

// HasCustomerFragmentIds returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasCustomerFragmentIds() bool {
	if o != nil && o.CustomerFragmentIds != nil {
		return true
	}

	return false
}

// SetCustomerFragmentIds gets a reference to the given []string and assigns it to the CustomerFragmentIds field.
func (o *GwClusterIdentity) SetCustomerFragmentIds(v []string) {
	o.CustomerFragmentIds = &v
}

// GetCustomerFragments returns the CustomerFragments field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetCustomerFragments() []CfInfo {
	if o == nil || o.CustomerFragments == nil {
		var ret []CfInfo
		return ret
	}
	return *o.CustomerFragments
}

// GetCustomerFragmentsOk returns a tuple with the CustomerFragments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetCustomerFragmentsOk() (*[]CfInfo, bool) {
	if o == nil || o.CustomerFragments == nil {
		return nil, false
	}
	return o.CustomerFragments, true
}

// HasCustomerFragments returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasCustomerFragments() bool {
	if o != nil && o.CustomerFragments != nil {
		return true
	}

	return false
}

// SetCustomerFragments gets a reference to the given []CfInfo and assigns it to the CustomerFragments field.
func (o *GwClusterIdentity) SetCustomerFragments(v []CfInfo) {
	o.CustomerFragments = &v
}

// GetDefaultProtectionKeyId returns the DefaultProtectionKeyId field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetDefaultProtectionKeyId() int64 {
	if o == nil || o.DefaultProtectionKeyId == nil {
		var ret int64
		return ret
	}
	return *o.DefaultProtectionKeyId
}

// GetDefaultProtectionKeyIdOk returns a tuple with the DefaultProtectionKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetDefaultProtectionKeyIdOk() (*int64, bool) {
	if o == nil || o.DefaultProtectionKeyId == nil {
		return nil, false
	}
	return o.DefaultProtectionKeyId, true
}

// HasDefaultProtectionKeyId returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasDefaultProtectionKeyId() bool {
	if o != nil && o.DefaultProtectionKeyId != nil {
		return true
	}

	return false
}

// SetDefaultProtectionKeyId gets a reference to the given int64 and assigns it to the DefaultProtectionKeyId field.
func (o *GwClusterIdentity) SetDefaultProtectionKeyId(v int64) {
	o.DefaultProtectionKeyId = &v
}

// GetDefaultSecretLocation returns the DefaultSecretLocation field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetDefaultSecretLocation() string {
	if o == nil || o.DefaultSecretLocation == nil {
		var ret string
		return ret
	}
	return *o.DefaultSecretLocation
}

// GetDefaultSecretLocationOk returns a tuple with the DefaultSecretLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetDefaultSecretLocationOk() (*string, bool) {
	if o == nil || o.DefaultSecretLocation == nil {
		return nil, false
	}
	return o.DefaultSecretLocation, true
}

// HasDefaultSecretLocation returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasDefaultSecretLocation() bool {
	if o != nil && o.DefaultSecretLocation != nil {
		return true
	}

	return false
}

// SetDefaultSecretLocation gets a reference to the given string and assigns it to the DefaultSecretLocation field.
func (o *GwClusterIdentity) SetDefaultSecretLocation(v string) {
	o.DefaultSecretLocation = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GwClusterIdentity) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GwClusterIdentity) SetId(v int64) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GwClusterIdentity) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *GwClusterIdentity) GetStatusDescription() string {
	if o == nil || o.StatusDescription == nil {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GwClusterIdentity) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || o.StatusDescription == nil {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *GwClusterIdentity) HasStatusDescription() bool {
	if o != nil && o.StatusDescription != nil {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *GwClusterIdentity) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

func (o GwClusterIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allowed != nil {
		toSerialize["allowed"] = o.Allowed
	}
	if o.AllowedAccessIds != nil {
		toSerialize["allowed_access_ids"] = o.AllowedAccessIds
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.ClusterUrl != nil {
		toSerialize["cluster_url"] = o.ClusterUrl
	}
	if o.CurrentGw != nil {
		toSerialize["current_gw"] = o.CurrentGw
	}
	if o.CustomerFragmentIds != nil {
		toSerialize["customer_fragment_ids"] = o.CustomerFragmentIds
	}
	if o.CustomerFragments != nil {
		toSerialize["customer_fragments"] = o.CustomerFragments
	}
	if o.DefaultProtectionKeyId != nil {
		toSerialize["default_protection_key_id"] = o.DefaultProtectionKeyId
	}
	if o.DefaultSecretLocation != nil {
		toSerialize["default_secret_location"] = o.DefaultSecretLocation
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDescription != nil {
		toSerialize["status_description"] = o.StatusDescription
	}
	return json.Marshal(toSerialize)
}

type NullableGwClusterIdentity struct {
	value *GwClusterIdentity
	isSet bool
}

func (v NullableGwClusterIdentity) Get() *GwClusterIdentity {
	return v.value
}

func (v *NullableGwClusterIdentity) Set(val *GwClusterIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableGwClusterIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableGwClusterIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGwClusterIdentity(val *GwClusterIdentity) *NullableGwClusterIdentity {
	return &NullableGwClusterIdentity{value: val, isSet: true}
}

func (v NullableGwClusterIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGwClusterIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


