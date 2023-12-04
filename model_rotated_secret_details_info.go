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

// RotatedSecretDetailsInfo RotatedSecretDetailsInfo The rotated secret rotator info
type RotatedSecretDetailsInfo struct {
	DeletePreviousVersionInDays *int32 `json:"delete_previous_version_in_days,omitempty"`
	GwClusterId *int64 `json:"gw_cluster_id,omitempty"`
	LastRotationError *string `json:"last_rotation_error,omitempty"`
	NumberOfVersionsToSave *int32 `json:"number_of_versions_to_save,omitempty"`
	RotationHour *int32 `json:"rotation_hour,omitempty"`
	RotationIntervalMin *bool `json:"rotation_interval_min,omitempty"`
	RotationStatement *string `json:"rotation_statement,omitempty"`
	RotatorCredsType *string `json:"rotator_creds_type,omitempty"`
	// RotationStatus defines types of rotation Status
	RotatorStatus *string `json:"rotator_status,omitempty"`
	RotatorType *string `json:"rotator_type,omitempty"`
	SamePassword *bool `json:"same_password,omitempty"`
	ServicesDetails *[]WindowsService `json:"services_details,omitempty"`
}

// NewRotatedSecretDetailsInfo instantiates a new RotatedSecretDetailsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotatedSecretDetailsInfo() *RotatedSecretDetailsInfo {
	this := RotatedSecretDetailsInfo{}
	return &this
}

// NewRotatedSecretDetailsInfoWithDefaults instantiates a new RotatedSecretDetailsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotatedSecretDetailsInfoWithDefaults() *RotatedSecretDetailsInfo {
	this := RotatedSecretDetailsInfo{}
	return &this
}

// GetDeletePreviousVersionInDays returns the DeletePreviousVersionInDays field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetDeletePreviousVersionInDays() int32 {
	if o == nil || o.DeletePreviousVersionInDays == nil {
		var ret int32
		return ret
	}
	return *o.DeletePreviousVersionInDays
}

// GetDeletePreviousVersionInDaysOk returns a tuple with the DeletePreviousVersionInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetDeletePreviousVersionInDaysOk() (*int32, bool) {
	if o == nil || o.DeletePreviousVersionInDays == nil {
		return nil, false
	}
	return o.DeletePreviousVersionInDays, true
}

// HasDeletePreviousVersionInDays returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasDeletePreviousVersionInDays() bool {
	if o != nil && o.DeletePreviousVersionInDays != nil {
		return true
	}

	return false
}

// SetDeletePreviousVersionInDays gets a reference to the given int32 and assigns it to the DeletePreviousVersionInDays field.
func (o *RotatedSecretDetailsInfo) SetDeletePreviousVersionInDays(v int32) {
	o.DeletePreviousVersionInDays = &v
}

// GetGwClusterId returns the GwClusterId field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetGwClusterId() int64 {
	if o == nil || o.GwClusterId == nil {
		var ret int64
		return ret
	}
	return *o.GwClusterId
}

// GetGwClusterIdOk returns a tuple with the GwClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetGwClusterIdOk() (*int64, bool) {
	if o == nil || o.GwClusterId == nil {
		return nil, false
	}
	return o.GwClusterId, true
}

// HasGwClusterId returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasGwClusterId() bool {
	if o != nil && o.GwClusterId != nil {
		return true
	}

	return false
}

// SetGwClusterId gets a reference to the given int64 and assigns it to the GwClusterId field.
func (o *RotatedSecretDetailsInfo) SetGwClusterId(v int64) {
	o.GwClusterId = &v
}

// GetLastRotationError returns the LastRotationError field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetLastRotationError() string {
	if o == nil || o.LastRotationError == nil {
		var ret string
		return ret
	}
	return *o.LastRotationError
}

// GetLastRotationErrorOk returns a tuple with the LastRotationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetLastRotationErrorOk() (*string, bool) {
	if o == nil || o.LastRotationError == nil {
		return nil, false
	}
	return o.LastRotationError, true
}

// HasLastRotationError returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasLastRotationError() bool {
	if o != nil && o.LastRotationError != nil {
		return true
	}

	return false
}

// SetLastRotationError gets a reference to the given string and assigns it to the LastRotationError field.
func (o *RotatedSecretDetailsInfo) SetLastRotationError(v string) {
	o.LastRotationError = &v
}

// GetNumberOfVersionsToSave returns the NumberOfVersionsToSave field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetNumberOfVersionsToSave() int32 {
	if o == nil || o.NumberOfVersionsToSave == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfVersionsToSave
}

// GetNumberOfVersionsToSaveOk returns a tuple with the NumberOfVersionsToSave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetNumberOfVersionsToSaveOk() (*int32, bool) {
	if o == nil || o.NumberOfVersionsToSave == nil {
		return nil, false
	}
	return o.NumberOfVersionsToSave, true
}

// HasNumberOfVersionsToSave returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasNumberOfVersionsToSave() bool {
	if o != nil && o.NumberOfVersionsToSave != nil {
		return true
	}

	return false
}

// SetNumberOfVersionsToSave gets a reference to the given int32 and assigns it to the NumberOfVersionsToSave field.
func (o *RotatedSecretDetailsInfo) SetNumberOfVersionsToSave(v int32) {
	o.NumberOfVersionsToSave = &v
}

// GetRotationHour returns the RotationHour field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetRotationHour() int32 {
	if o == nil || o.RotationHour == nil {
		var ret int32
		return ret
	}
	return *o.RotationHour
}

// GetRotationHourOk returns a tuple with the RotationHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetRotationHourOk() (*int32, bool) {
	if o == nil || o.RotationHour == nil {
		return nil, false
	}
	return o.RotationHour, true
}

// HasRotationHour returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasRotationHour() bool {
	if o != nil && o.RotationHour != nil {
		return true
	}

	return false
}

// SetRotationHour gets a reference to the given int32 and assigns it to the RotationHour field.
func (o *RotatedSecretDetailsInfo) SetRotationHour(v int32) {
	o.RotationHour = &v
}

// GetRotationIntervalMin returns the RotationIntervalMin field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetRotationIntervalMin() bool {
	if o == nil || o.RotationIntervalMin == nil {
		var ret bool
		return ret
	}
	return *o.RotationIntervalMin
}

// GetRotationIntervalMinOk returns a tuple with the RotationIntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetRotationIntervalMinOk() (*bool, bool) {
	if o == nil || o.RotationIntervalMin == nil {
		return nil, false
	}
	return o.RotationIntervalMin, true
}

// HasRotationIntervalMin returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasRotationIntervalMin() bool {
	if o != nil && o.RotationIntervalMin != nil {
		return true
	}

	return false
}

// SetRotationIntervalMin gets a reference to the given bool and assigns it to the RotationIntervalMin field.
func (o *RotatedSecretDetailsInfo) SetRotationIntervalMin(v bool) {
	o.RotationIntervalMin = &v
}

// GetRotationStatement returns the RotationStatement field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetRotationStatement() string {
	if o == nil || o.RotationStatement == nil {
		var ret string
		return ret
	}
	return *o.RotationStatement
}

// GetRotationStatementOk returns a tuple with the RotationStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetRotationStatementOk() (*string, bool) {
	if o == nil || o.RotationStatement == nil {
		return nil, false
	}
	return o.RotationStatement, true
}

// HasRotationStatement returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasRotationStatement() bool {
	if o != nil && o.RotationStatement != nil {
		return true
	}

	return false
}

// SetRotationStatement gets a reference to the given string and assigns it to the RotationStatement field.
func (o *RotatedSecretDetailsInfo) SetRotationStatement(v string) {
	o.RotationStatement = &v
}

// GetRotatorCredsType returns the RotatorCredsType field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetRotatorCredsType() string {
	if o == nil || o.RotatorCredsType == nil {
		var ret string
		return ret
	}
	return *o.RotatorCredsType
}

// GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetRotatorCredsTypeOk() (*string, bool) {
	if o == nil || o.RotatorCredsType == nil {
		return nil, false
	}
	return o.RotatorCredsType, true
}

// HasRotatorCredsType returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasRotatorCredsType() bool {
	if o != nil && o.RotatorCredsType != nil {
		return true
	}

	return false
}

// SetRotatorCredsType gets a reference to the given string and assigns it to the RotatorCredsType field.
func (o *RotatedSecretDetailsInfo) SetRotatorCredsType(v string) {
	o.RotatorCredsType = &v
}

// GetRotatorStatus returns the RotatorStatus field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetRotatorStatus() string {
	if o == nil || o.RotatorStatus == nil {
		var ret string
		return ret
	}
	return *o.RotatorStatus
}

// GetRotatorStatusOk returns a tuple with the RotatorStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetRotatorStatusOk() (*string, bool) {
	if o == nil || o.RotatorStatus == nil {
		return nil, false
	}
	return o.RotatorStatus, true
}

// HasRotatorStatus returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasRotatorStatus() bool {
	if o != nil && o.RotatorStatus != nil {
		return true
	}

	return false
}

// SetRotatorStatus gets a reference to the given string and assigns it to the RotatorStatus field.
func (o *RotatedSecretDetailsInfo) SetRotatorStatus(v string) {
	o.RotatorStatus = &v
}

// GetRotatorType returns the RotatorType field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetRotatorType() string {
	if o == nil || o.RotatorType == nil {
		var ret string
		return ret
	}
	return *o.RotatorType
}

// GetRotatorTypeOk returns a tuple with the RotatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetRotatorTypeOk() (*string, bool) {
	if o == nil || o.RotatorType == nil {
		return nil, false
	}
	return o.RotatorType, true
}

// HasRotatorType returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasRotatorType() bool {
	if o != nil && o.RotatorType != nil {
		return true
	}

	return false
}

// SetRotatorType gets a reference to the given string and assigns it to the RotatorType field.
func (o *RotatedSecretDetailsInfo) SetRotatorType(v string) {
	o.RotatorType = &v
}

// GetSamePassword returns the SamePassword field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetSamePassword() bool {
	if o == nil || o.SamePassword == nil {
		var ret bool
		return ret
	}
	return *o.SamePassword
}

// GetSamePasswordOk returns a tuple with the SamePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetSamePasswordOk() (*bool, bool) {
	if o == nil || o.SamePassword == nil {
		return nil, false
	}
	return o.SamePassword, true
}

// HasSamePassword returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasSamePassword() bool {
	if o != nil && o.SamePassword != nil {
		return true
	}

	return false
}

// SetSamePassword gets a reference to the given bool and assigns it to the SamePassword field.
func (o *RotatedSecretDetailsInfo) SetSamePassword(v bool) {
	o.SamePassword = &v
}

// GetServicesDetails returns the ServicesDetails field value if set, zero value otherwise.
func (o *RotatedSecretDetailsInfo) GetServicesDetails() []WindowsService {
	if o == nil || o.ServicesDetails == nil {
		var ret []WindowsService
		return ret
	}
	return *o.ServicesDetails
}

// GetServicesDetailsOk returns a tuple with the ServicesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretDetailsInfo) GetServicesDetailsOk() (*[]WindowsService, bool) {
	if o == nil || o.ServicesDetails == nil {
		return nil, false
	}
	return o.ServicesDetails, true
}

// HasServicesDetails returns a boolean if a field has been set.
func (o *RotatedSecretDetailsInfo) HasServicesDetails() bool {
	if o != nil && o.ServicesDetails != nil {
		return true
	}

	return false
}

// SetServicesDetails gets a reference to the given []WindowsService and assigns it to the ServicesDetails field.
func (o *RotatedSecretDetailsInfo) SetServicesDetails(v []WindowsService) {
	o.ServicesDetails = &v
}

func (o RotatedSecretDetailsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeletePreviousVersionInDays != nil {
		toSerialize["delete_previous_version_in_days"] = o.DeletePreviousVersionInDays
	}
	if o.GwClusterId != nil {
		toSerialize["gw_cluster_id"] = o.GwClusterId
	}
	if o.LastRotationError != nil {
		toSerialize["last_rotation_error"] = o.LastRotationError
	}
	if o.NumberOfVersionsToSave != nil {
		toSerialize["number_of_versions_to_save"] = o.NumberOfVersionsToSave
	}
	if o.RotationHour != nil {
		toSerialize["rotation_hour"] = o.RotationHour
	}
	if o.RotationIntervalMin != nil {
		toSerialize["rotation_interval_min"] = o.RotationIntervalMin
	}
	if o.RotationStatement != nil {
		toSerialize["rotation_statement"] = o.RotationStatement
	}
	if o.RotatorCredsType != nil {
		toSerialize["rotator_creds_type"] = o.RotatorCredsType
	}
	if o.RotatorStatus != nil {
		toSerialize["rotator_status"] = o.RotatorStatus
	}
	if o.RotatorType != nil {
		toSerialize["rotator_type"] = o.RotatorType
	}
	if o.SamePassword != nil {
		toSerialize["same_password"] = o.SamePassword
	}
	if o.ServicesDetails != nil {
		toSerialize["services_details"] = o.ServicesDetails
	}
	return json.Marshal(toSerialize)
}

type NullableRotatedSecretDetailsInfo struct {
	value *RotatedSecretDetailsInfo
	isSet bool
}

func (v NullableRotatedSecretDetailsInfo) Get() *RotatedSecretDetailsInfo {
	return v.value
}

func (v *NullableRotatedSecretDetailsInfo) Set(val *RotatedSecretDetailsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRotatedSecretDetailsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRotatedSecretDetailsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotatedSecretDetailsInfo(val *RotatedSecretDetailsInfo) *NullableRotatedSecretDetailsInfo {
	return &NullableRotatedSecretDetailsInfo{value: val, isSet: true}
}

func (v NullableRotatedSecretDetailsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotatedSecretDetailsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


