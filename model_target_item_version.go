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
	"time"
)

// TargetItemVersion struct for TargetItemVersion
type TargetItemVersion struct {
	AccessDate *time.Time `json:"access_date,omitempty"`
	AccessDateDisplay *string `json:"access_date_display,omitempty"`
	CreationDate *time.Time `json:"creation_date,omitempty"`
	CustomerFragmentId *string `json:"customer_fragment_id,omitempty"`
	DeletionDate *time.Time `json:"deletion_date,omitempty"`
	// ItemState defines the different states an Item can be in
	ItemVersionState *string `json:"item_version_state,omitempty"`
	LatestVersion *bool `json:"latest_version,omitempty"`
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	ProtectionKeyName *string `json:"protection_key_name,omitempty"`
	TargetName *string `json:"target_name,omitempty"`
	Version *int32 `json:"version,omitempty"`
	WithCustomerFragment *bool `json:"with_customer_fragment,omitempty"`
}

// NewTargetItemVersion instantiates a new TargetItemVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetItemVersion() *TargetItemVersion {
	this := TargetItemVersion{}
	return &this
}

// NewTargetItemVersionWithDefaults instantiates a new TargetItemVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetItemVersionWithDefaults() *TargetItemVersion {
	this := TargetItemVersion{}
	return &this
}

// GetAccessDate returns the AccessDate field value if set, zero value otherwise.
func (o *TargetItemVersion) GetAccessDate() time.Time {
	if o == nil || o.AccessDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AccessDate
}

// GetAccessDateOk returns a tuple with the AccessDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetAccessDateOk() (*time.Time, bool) {
	if o == nil || o.AccessDate == nil {
		return nil, false
	}
	return o.AccessDate, true
}

// HasAccessDate returns a boolean if a field has been set.
func (o *TargetItemVersion) HasAccessDate() bool {
	if o != nil && o.AccessDate != nil {
		return true
	}

	return false
}

// SetAccessDate gets a reference to the given time.Time and assigns it to the AccessDate field.
func (o *TargetItemVersion) SetAccessDate(v time.Time) {
	o.AccessDate = &v
}

// GetAccessDateDisplay returns the AccessDateDisplay field value if set, zero value otherwise.
func (o *TargetItemVersion) GetAccessDateDisplay() string {
	if o == nil || o.AccessDateDisplay == nil {
		var ret string
		return ret
	}
	return *o.AccessDateDisplay
}

// GetAccessDateDisplayOk returns a tuple with the AccessDateDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetAccessDateDisplayOk() (*string, bool) {
	if o == nil || o.AccessDateDisplay == nil {
		return nil, false
	}
	return o.AccessDateDisplay, true
}

// HasAccessDateDisplay returns a boolean if a field has been set.
func (o *TargetItemVersion) HasAccessDateDisplay() bool {
	if o != nil && o.AccessDateDisplay != nil {
		return true
	}

	return false
}

// SetAccessDateDisplay gets a reference to the given string and assigns it to the AccessDateDisplay field.
func (o *TargetItemVersion) SetAccessDateDisplay(v string) {
	o.AccessDateDisplay = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *TargetItemVersion) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *TargetItemVersion) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *TargetItemVersion) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetCustomerFragmentId returns the CustomerFragmentId field value if set, zero value otherwise.
func (o *TargetItemVersion) GetCustomerFragmentId() string {
	if o == nil || o.CustomerFragmentId == nil {
		var ret string
		return ret
	}
	return *o.CustomerFragmentId
}

// GetCustomerFragmentIdOk returns a tuple with the CustomerFragmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetCustomerFragmentIdOk() (*string, bool) {
	if o == nil || o.CustomerFragmentId == nil {
		return nil, false
	}
	return o.CustomerFragmentId, true
}

// HasCustomerFragmentId returns a boolean if a field has been set.
func (o *TargetItemVersion) HasCustomerFragmentId() bool {
	if o != nil && o.CustomerFragmentId != nil {
		return true
	}

	return false
}

// SetCustomerFragmentId gets a reference to the given string and assigns it to the CustomerFragmentId field.
func (o *TargetItemVersion) SetCustomerFragmentId(v string) {
	o.CustomerFragmentId = &v
}

// GetDeletionDate returns the DeletionDate field value if set, zero value otherwise.
func (o *TargetItemVersion) GetDeletionDate() time.Time {
	if o == nil || o.DeletionDate == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletionDate
}

// GetDeletionDateOk returns a tuple with the DeletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetDeletionDateOk() (*time.Time, bool) {
	if o == nil || o.DeletionDate == nil {
		return nil, false
	}
	return o.DeletionDate, true
}

// HasDeletionDate returns a boolean if a field has been set.
func (o *TargetItemVersion) HasDeletionDate() bool {
	if o != nil && o.DeletionDate != nil {
		return true
	}

	return false
}

// SetDeletionDate gets a reference to the given time.Time and assigns it to the DeletionDate field.
func (o *TargetItemVersion) SetDeletionDate(v time.Time) {
	o.DeletionDate = &v
}

// GetItemVersionState returns the ItemVersionState field value if set, zero value otherwise.
func (o *TargetItemVersion) GetItemVersionState() string {
	if o == nil || o.ItemVersionState == nil {
		var ret string
		return ret
	}
	return *o.ItemVersionState
}

// GetItemVersionStateOk returns a tuple with the ItemVersionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetItemVersionStateOk() (*string, bool) {
	if o == nil || o.ItemVersionState == nil {
		return nil, false
	}
	return o.ItemVersionState, true
}

// HasItemVersionState returns a boolean if a field has been set.
func (o *TargetItemVersion) HasItemVersionState() bool {
	if o != nil && o.ItemVersionState != nil {
		return true
	}

	return false
}

// SetItemVersionState gets a reference to the given string and assigns it to the ItemVersionState field.
func (o *TargetItemVersion) SetItemVersionState(v string) {
	o.ItemVersionState = &v
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *TargetItemVersion) GetLatestVersion() bool {
	if o == nil || o.LatestVersion == nil {
		var ret bool
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetLatestVersionOk() (*bool, bool) {
	if o == nil || o.LatestVersion == nil {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *TargetItemVersion) HasLatestVersion() bool {
	if o != nil && o.LatestVersion != nil {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given bool and assigns it to the LatestVersion field.
func (o *TargetItemVersion) SetLatestVersion(v bool) {
	o.LatestVersion = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *TargetItemVersion) GetModificationDate() time.Time {
	if o == nil || o.ModificationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || o.ModificationDate == nil {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *TargetItemVersion) HasModificationDate() bool {
	if o != nil && o.ModificationDate != nil {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *TargetItemVersion) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

// GetProtectionKeyName returns the ProtectionKeyName field value if set, zero value otherwise.
func (o *TargetItemVersion) GetProtectionKeyName() string {
	if o == nil || o.ProtectionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKeyName
}

// GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetProtectionKeyNameOk() (*string, bool) {
	if o == nil || o.ProtectionKeyName == nil {
		return nil, false
	}
	return o.ProtectionKeyName, true
}

// HasProtectionKeyName returns a boolean if a field has been set.
func (o *TargetItemVersion) HasProtectionKeyName() bool {
	if o != nil && o.ProtectionKeyName != nil {
		return true
	}

	return false
}

// SetProtectionKeyName gets a reference to the given string and assigns it to the ProtectionKeyName field.
func (o *TargetItemVersion) SetProtectionKeyName(v string) {
	o.ProtectionKeyName = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *TargetItemVersion) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *TargetItemVersion) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *TargetItemVersion) SetTargetName(v string) {
	o.TargetName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TargetItemVersion) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TargetItemVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *TargetItemVersion) SetVersion(v int32) {
	o.Version = &v
}

// GetWithCustomerFragment returns the WithCustomerFragment field value if set, zero value otherwise.
func (o *TargetItemVersion) GetWithCustomerFragment() bool {
	if o == nil || o.WithCustomerFragment == nil {
		var ret bool
		return ret
	}
	return *o.WithCustomerFragment
}

// GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemVersion) GetWithCustomerFragmentOk() (*bool, bool) {
	if o == nil || o.WithCustomerFragment == nil {
		return nil, false
	}
	return o.WithCustomerFragment, true
}

// HasWithCustomerFragment returns a boolean if a field has been set.
func (o *TargetItemVersion) HasWithCustomerFragment() bool {
	if o != nil && o.WithCustomerFragment != nil {
		return true
	}

	return false
}

// SetWithCustomerFragment gets a reference to the given bool and assigns it to the WithCustomerFragment field.
func (o *TargetItemVersion) SetWithCustomerFragment(v bool) {
	o.WithCustomerFragment = &v
}

func (o TargetItemVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessDate != nil {
		toSerialize["access_date"] = o.AccessDate
	}
	if o.AccessDateDisplay != nil {
		toSerialize["access_date_display"] = o.AccessDateDisplay
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.CustomerFragmentId != nil {
		toSerialize["customer_fragment_id"] = o.CustomerFragmentId
	}
	if o.DeletionDate != nil {
		toSerialize["deletion_date"] = o.DeletionDate
	}
	if o.ItemVersionState != nil {
		toSerialize["item_version_state"] = o.ItemVersionState
	}
	if o.LatestVersion != nil {
		toSerialize["latest_version"] = o.LatestVersion
	}
	if o.ModificationDate != nil {
		toSerialize["modification_date"] = o.ModificationDate
	}
	if o.ProtectionKeyName != nil {
		toSerialize["protection_key_name"] = o.ProtectionKeyName
	}
	if o.TargetName != nil {
		toSerialize["target_name"] = o.TargetName
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.WithCustomerFragment != nil {
		toSerialize["with_customer_fragment"] = o.WithCustomerFragment
	}
	return json.Marshal(toSerialize)
}

type NullableTargetItemVersion struct {
	value *TargetItemVersion
	isSet bool
}

func (v NullableTargetItemVersion) Get() *TargetItemVersion {
	return v.value
}

func (v *NullableTargetItemVersion) Set(val *TargetItemVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetItemVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetItemVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetItemVersion(val *TargetItemVersion) *NullableTargetItemVersion {
	return &NullableTargetItemVersion{value: val, isSet: true}
}

func (v NullableTargetItemVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetItemVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


