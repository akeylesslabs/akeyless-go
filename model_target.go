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

// Target struct for Target
type Target struct {
	AccessDate *time.Time `json:"access_date,omitempty"`
	AccessDateDisplay *string `json:"access_date_display,omitempty"`
	AccessRequestStatus *string `json:"access_request_status,omitempty"`
	// this is not \"omitempty\" since an empty value causes no update while an empty map will clear the attributes
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	ClientPermissions *[]string `json:"client_permissions,omitempty"`
	Comment *string `json:"comment,omitempty"`
	CreationDate *time.Time `json:"creation_date,omitempty"`
	CredentialsLess *bool `json:"credentials_less,omitempty"`
	IsAccessRequestEnabled *bool `json:"is_access_request_enabled,omitempty"`
	LastVersion *int32 `json:"last_version,omitempty"`
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	ProtectionKeyName *string `json:"protection_key_name,omitempty"`
	TargetDetails *string `json:"target_details,omitempty"`
	TargetId *int64 `json:"target_id,omitempty"`
	TargetItemsAssoc *[]TargetItemAssociation `json:"target_items_assoc,omitempty"`
	TargetName *string `json:"target_name,omitempty"`
	TargetSubType *string `json:"target_sub_type,omitempty"`
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

// GetAccessDate returns the AccessDate field value if set, zero value otherwise.
func (o *Target) GetAccessDate() time.Time {
	if o == nil || o.AccessDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AccessDate
}

// GetAccessDateOk returns a tuple with the AccessDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetAccessDateOk() (*time.Time, bool) {
	if o == nil || o.AccessDate == nil {
		return nil, false
	}
	return o.AccessDate, true
}

// HasAccessDate returns a boolean if a field has been set.
func (o *Target) HasAccessDate() bool {
	if o != nil && o.AccessDate != nil {
		return true
	}

	return false
}

// SetAccessDate gets a reference to the given time.Time and assigns it to the AccessDate field.
func (o *Target) SetAccessDate(v time.Time) {
	o.AccessDate = &v
}

// GetAccessDateDisplay returns the AccessDateDisplay field value if set, zero value otherwise.
func (o *Target) GetAccessDateDisplay() string {
	if o == nil || o.AccessDateDisplay == nil {
		var ret string
		return ret
	}
	return *o.AccessDateDisplay
}

// GetAccessDateDisplayOk returns a tuple with the AccessDateDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetAccessDateDisplayOk() (*string, bool) {
	if o == nil || o.AccessDateDisplay == nil {
		return nil, false
	}
	return o.AccessDateDisplay, true
}

// HasAccessDateDisplay returns a boolean if a field has been set.
func (o *Target) HasAccessDateDisplay() bool {
	if o != nil && o.AccessDateDisplay != nil {
		return true
	}

	return false
}

// SetAccessDateDisplay gets a reference to the given string and assigns it to the AccessDateDisplay field.
func (o *Target) SetAccessDateDisplay(v string) {
	o.AccessDateDisplay = &v
}

// GetAccessRequestStatus returns the AccessRequestStatus field value if set, zero value otherwise.
func (o *Target) GetAccessRequestStatus() string {
	if o == nil || o.AccessRequestStatus == nil {
		var ret string
		return ret
	}
	return *o.AccessRequestStatus
}

// GetAccessRequestStatusOk returns a tuple with the AccessRequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetAccessRequestStatusOk() (*string, bool) {
	if o == nil || o.AccessRequestStatus == nil {
		return nil, false
	}
	return o.AccessRequestStatus, true
}

// HasAccessRequestStatus returns a boolean if a field has been set.
func (o *Target) HasAccessRequestStatus() bool {
	if o != nil && o.AccessRequestStatus != nil {
		return true
	}

	return false
}

// SetAccessRequestStatus gets a reference to the given string and assigns it to the AccessRequestStatus field.
func (o *Target) SetAccessRequestStatus(v string) {
	o.AccessRequestStatus = &v
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

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Target) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Target) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *Target) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetCredentialsLess returns the CredentialsLess field value if set, zero value otherwise.
func (o *Target) GetCredentialsLess() bool {
	if o == nil || o.CredentialsLess == nil {
		var ret bool
		return ret
	}
	return *o.CredentialsLess
}

// GetCredentialsLessOk returns a tuple with the CredentialsLess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetCredentialsLessOk() (*bool, bool) {
	if o == nil || o.CredentialsLess == nil {
		return nil, false
	}
	return o.CredentialsLess, true
}

// HasCredentialsLess returns a boolean if a field has been set.
func (o *Target) HasCredentialsLess() bool {
	if o != nil && o.CredentialsLess != nil {
		return true
	}

	return false
}

// SetCredentialsLess gets a reference to the given bool and assigns it to the CredentialsLess field.
func (o *Target) SetCredentialsLess(v bool) {
	o.CredentialsLess = &v
}

// GetIsAccessRequestEnabled returns the IsAccessRequestEnabled field value if set, zero value otherwise.
func (o *Target) GetIsAccessRequestEnabled() bool {
	if o == nil || o.IsAccessRequestEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsAccessRequestEnabled
}

// GetIsAccessRequestEnabledOk returns a tuple with the IsAccessRequestEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetIsAccessRequestEnabledOk() (*bool, bool) {
	if o == nil || o.IsAccessRequestEnabled == nil {
		return nil, false
	}
	return o.IsAccessRequestEnabled, true
}

// HasIsAccessRequestEnabled returns a boolean if a field has been set.
func (o *Target) HasIsAccessRequestEnabled() bool {
	if o != nil && o.IsAccessRequestEnabled != nil {
		return true
	}

	return false
}

// SetIsAccessRequestEnabled gets a reference to the given bool and assigns it to the IsAccessRequestEnabled field.
func (o *Target) SetIsAccessRequestEnabled(v bool) {
	o.IsAccessRequestEnabled = &v
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

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *Target) GetModificationDate() time.Time {
	if o == nil || o.ModificationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || o.ModificationDate == nil {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *Target) HasModificationDate() bool {
	if o != nil && o.ModificationDate != nil {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *Target) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
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

// GetTargetDetails returns the TargetDetails field value if set, zero value otherwise.
func (o *Target) GetTargetDetails() string {
	if o == nil || o.TargetDetails == nil {
		var ret string
		return ret
	}
	return *o.TargetDetails
}

// GetTargetDetailsOk returns a tuple with the TargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetTargetDetailsOk() (*string, bool) {
	if o == nil || o.TargetDetails == nil {
		return nil, false
	}
	return o.TargetDetails, true
}

// HasTargetDetails returns a boolean if a field has been set.
func (o *Target) HasTargetDetails() bool {
	if o != nil && o.TargetDetails != nil {
		return true
	}

	return false
}

// SetTargetDetails gets a reference to the given string and assigns it to the TargetDetails field.
func (o *Target) SetTargetDetails(v string) {
	o.TargetDetails = &v
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

// GetTargetSubType returns the TargetSubType field value if set, zero value otherwise.
func (o *Target) GetTargetSubType() string {
	if o == nil || o.TargetSubType == nil {
		var ret string
		return ret
	}
	return *o.TargetSubType
}

// GetTargetSubTypeOk returns a tuple with the TargetSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetTargetSubTypeOk() (*string, bool) {
	if o == nil || o.TargetSubType == nil {
		return nil, false
	}
	return o.TargetSubType, true
}

// HasTargetSubType returns a boolean if a field has been set.
func (o *Target) HasTargetSubType() bool {
	if o != nil && o.TargetSubType != nil {
		return true
	}

	return false
}

// SetTargetSubType gets a reference to the given string and assigns it to the TargetSubType field.
func (o *Target) SetTargetSubType(v string) {
	o.TargetSubType = &v
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
	if o.AccessDate != nil {
		toSerialize["access_date"] = o.AccessDate
	}
	if o.AccessDateDisplay != nil {
		toSerialize["access_date_display"] = o.AccessDateDisplay
	}
	if o.AccessRequestStatus != nil {
		toSerialize["access_request_status"] = o.AccessRequestStatus
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ClientPermissions != nil {
		toSerialize["client_permissions"] = o.ClientPermissions
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.CredentialsLess != nil {
		toSerialize["credentials_less"] = o.CredentialsLess
	}
	if o.IsAccessRequestEnabled != nil {
		toSerialize["is_access_request_enabled"] = o.IsAccessRequestEnabled
	}
	if o.LastVersion != nil {
		toSerialize["last_version"] = o.LastVersion
	}
	if o.ModificationDate != nil {
		toSerialize["modification_date"] = o.ModificationDate
	}
	if o.ProtectionKeyName != nil {
		toSerialize["protection_key_name"] = o.ProtectionKeyName
	}
	if o.TargetDetails != nil {
		toSerialize["target_details"] = o.TargetDetails
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
	if o.TargetSubType != nil {
		toSerialize["target_sub_type"] = o.TargetSubType
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


