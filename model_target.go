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

// Target struct for Target
type Target struct {
	// this is not \"omitempty\" since an empty value causes no update while an empty map will clear the attributes
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	ClientPermissions *[]string `json:"client_permissions,omitempty"`
	Comment *string `json:"comment,omitempty"`
	LastVersion *int32 `json:"last_version,omitempty"`
	ProtectionKeyName *string `json:"protection_key_name,omitempty"`
	TargetId *int64 `json:"target_id,omitempty"`
	TargetItemsAssoc *[]TargetItemAssociation `json:"target_items_assoc,omitempty"`
	TargetName *string `json:"target_name,omitempty"`
	TargetType *string `json:"target_type,omitempty"`
	TargetVersions *[]ItemVersion `json:"target_versions,omitempty"`
	WithCustomerFragment *bool `json:"with_customer_fragment,omitempty"`
}

// NewTarget instantiates a new Target object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTarget() *Target {
	this := Target{}
	return &this
}

// NewTargetWithDefaults instantiates a new Target object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetWithDefaults() *Target {
	this := Target{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Target) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetAttributesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Target) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *Target) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

// GetClientPermissions returns the ClientPermissions field value if set, zero value otherwise.
func (o *Target) GetClientPermissions() []string {
	if o == nil || o.ClientPermissions == nil {
		var ret []string
		return ret
	}
	return *o.ClientPermissions
}

// GetClientPermissionsOk returns a tuple with the ClientPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetClientPermissionsOk() (*[]string, bool) {
	if o == nil || o.ClientPermissions == nil {
		return nil, false
	}
	return o.ClientPermissions, true
}

// HasClientPermissions returns a boolean if a field has been set.
func (o *Target) HasClientPermissions() bool {
	if o != nil && o.ClientPermissions != nil {
		return true
	}

	return false
}

// SetClientPermissions gets a reference to the given []string and assigns it to the ClientPermissions field.
func (o *Target) SetClientPermissions(v []string) {
	o.ClientPermissions = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Target) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Target) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Target) SetComment(v string) {
	o.Comment = &v
}

// GetLastVersion returns the LastVersion field value if set, zero value otherwise.
func (o *Target) GetLastVersion() int32 {
	if o == nil || o.LastVersion == nil {
		var ret int32
		return ret
	}
	return *o.LastVersion
}

// GetLastVersionOk returns a tuple with the LastVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetLastVersionOk() (*int32, bool) {
	if o == nil || o.LastVersion == nil {
		return nil, false
	}
	return o.LastVersion, true
}

// HasLastVersion returns a boolean if a field has been set.
func (o *Target) HasLastVersion() bool {
	if o != nil && o.LastVersion != nil {
		return true
	}

	return false
}

// SetLastVersion gets a reference to the given int32 and assigns it to the LastVersion field.
func (o *Target) SetLastVersion(v int32) {
	o.LastVersion = &v
}

// GetProtectionKeyName returns the ProtectionKeyName field value if set, zero value otherwise.
func (o *Target) GetProtectionKeyName() string {
	if o == nil || o.ProtectionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKeyName
}

// GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetProtectionKeyNameOk() (*string, bool) {
	if o == nil || o.ProtectionKeyName == nil {
		return nil, false
	}
	return o.ProtectionKeyName, true
}

// HasProtectionKeyName returns a boolean if a field has been set.
func (o *Target) HasProtectionKeyName() bool {
	if o != nil && o.ProtectionKeyName != nil {
		return true
	}

	return false
}

// SetProtectionKeyName gets a reference to the given string and assigns it to the ProtectionKeyName field.
func (o *Target) SetProtectionKeyName(v string) {
	o.ProtectionKeyName = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *Target) GetTargetId() int64 {
	if o == nil || o.TargetId == nil {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetTargetIdOk() (*int64, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *Target) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *Target) SetTargetId(v int64) {
	o.TargetId = &v
}

// GetTargetItemsAssoc returns the TargetItemsAssoc field value if set, zero value otherwise.
func (o *Target) GetTargetItemsAssoc() []TargetItemAssociation {
	if o == nil || o.TargetItemsAssoc == nil {
		var ret []TargetItemAssociation
		return ret
	}
	return *o.TargetItemsAssoc
}

// GetTargetItemsAssocOk returns a tuple with the TargetItemsAssoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetTargetItemsAssocOk() (*[]TargetItemAssociation, bool) {
	if o == nil || o.TargetItemsAssoc == nil {
		return nil, false
	}
	return o.TargetItemsAssoc, true
}

// HasTargetItemsAssoc returns a boolean if a field has been set.
func (o *Target) HasTargetItemsAssoc() bool {
	if o != nil && o.TargetItemsAssoc != nil {
		return true
	}

	return false
}

// SetTargetItemsAssoc gets a reference to the given []TargetItemAssociation and assigns it to the TargetItemsAssoc field.
func (o *Target) SetTargetItemsAssoc(v []TargetItemAssociation) {
	o.TargetItemsAssoc = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *Target) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *Target) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *Target) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *Target) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *Target) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *Target) SetTargetType(v string) {
	o.TargetType = &v
}

// GetTargetVersions returns the TargetVersions field value if set, zero value otherwise.
func (o *Target) GetTargetVersions() []ItemVersion {
	if o == nil || o.TargetVersions == nil {
		var ret []ItemVersion
		return ret
	}
	return *o.TargetVersions
}

// GetTargetVersionsOk returns a tuple with the TargetVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetTargetVersionsOk() (*[]ItemVersion, bool) {
	if o == nil || o.TargetVersions == nil {
		return nil, false
	}
	return o.TargetVersions, true
}

// HasTargetVersions returns a boolean if a field has been set.
func (o *Target) HasTargetVersions() bool {
	if o != nil && o.TargetVersions != nil {
		return true
	}

	return false
}

// SetTargetVersions gets a reference to the given []ItemVersion and assigns it to the TargetVersions field.
func (o *Target) SetTargetVersions(v []ItemVersion) {
	o.TargetVersions = &v
}

// GetWithCustomerFragment returns the WithCustomerFragment field value if set, zero value otherwise.
func (o *Target) GetWithCustomerFragment() bool {
	if o == nil || o.WithCustomerFragment == nil {
		var ret bool
		return ret
	}
	return *o.WithCustomerFragment
}

// GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetWithCustomerFragmentOk() (*bool, bool) {
	if o == nil || o.WithCustomerFragment == nil {
		return nil, false
	}
	return o.WithCustomerFragment, true
}

// HasWithCustomerFragment returns a boolean if a field has been set.
func (o *Target) HasWithCustomerFragment() bool {
	if o != nil && o.WithCustomerFragment != nil {
		return true
	}

	return false
}

// SetWithCustomerFragment gets a reference to the given bool and assigns it to the WithCustomerFragment field.
func (o *Target) SetWithCustomerFragment(v bool) {
	o.WithCustomerFragment = &v
}

func (o Target) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ClientPermissions != nil {
		toSerialize["client_permissions"] = o.ClientPermissions
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.LastVersion != nil {
		toSerialize["last_version"] = o.LastVersion
	}
	if o.ProtectionKeyName != nil {
		toSerialize["protection_key_name"] = o.ProtectionKeyName
	}
	if o.TargetId != nil {
		toSerialize["target_id"] = o.TargetId
	}
	if o.TargetItemsAssoc != nil {
		toSerialize["target_items_assoc"] = o.TargetItemsAssoc
	}
	if o.TargetName != nil {
		toSerialize["target_name"] = o.TargetName
	}
	if o.TargetType != nil {
		toSerialize["target_type"] = o.TargetType
	}
	if o.TargetVersions != nil {
		toSerialize["target_versions"] = o.TargetVersions
	}
	if o.WithCustomerFragment != nil {
		toSerialize["with_customer_fragment"] = o.WithCustomerFragment
	}
	return json.Marshal(toSerialize)
}

type NullableTarget struct {
	value *Target
	isSet bool
}

func (v NullableTarget) Get() *Target {
	return v.value
}

func (v *NullableTarget) Set(val *Target) {
	v.value = val
	v.isSet = true
}

func (v NullableTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTarget(val *Target) *NullableTarget {
	return &NullableTarget{value: val, isSet: true}
}

func (v NullableTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


