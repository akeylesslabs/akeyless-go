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

// NotiForwarder struct for NotiForwarder
type NotiForwarder struct {
	ClientPermissions *[]string `json:"client_permissions,omitempty"`
	Comment *string `json:"comment,omitempty"`
	CreationDate *time.Time `json:"creation_date,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	EventTypes *[]string `json:"event_types,omitempty"`
	GatewayClusterId *int64 `json:"gateway_cluster_id,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	LastVersion *int32 `json:"last_version,omitempty"`
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	NotiForwarderId *int64 `json:"noti_forwarder_id,omitempty"`
	NotiForwarderName *string `json:"noti_forwarder_name,omitempty"`
	NotiForwarderType *string `json:"noti_forwarder_type,omitempty"`
	NotiForwarderVersions *[]ItemVersion `json:"noti_forwarder_versions,omitempty"`
	Paths *[]string `json:"paths,omitempty"`
	ProtectionKey *string `json:"protection_key,omitempty"`
	RunnerType *string `json:"runner_type,omitempty"`
	TimespanInSeconds *int64 `json:"timespan_in_seconds,omitempty"`
	ToEmails *[]EmailEntry `json:"to_emails,omitempty"`
	Username *string `json:"username,omitempty"`
	WithCustomerFragment *bool `json:"with_customer_fragment,omitempty"`
}

// NewNotiForwarder instantiates a new NotiForwarder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotiForwarder() *NotiForwarder {
	this := NotiForwarder{}
	return &this
}

// NewNotiForwarderWithDefaults instantiates a new NotiForwarder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotiForwarderWithDefaults() *NotiForwarder {
	this := NotiForwarder{}
	return &this
}

// GetClientPermissions returns the ClientPermissions field value if set, zero value otherwise.
func (o *NotiForwarder) GetClientPermissions() []string {
	if o == nil || o.ClientPermissions == nil {
		var ret []string
		return ret
	}
	return *o.ClientPermissions
}

// GetClientPermissionsOk returns a tuple with the ClientPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetClientPermissionsOk() (*[]string, bool) {
	if o == nil || o.ClientPermissions == nil {
		return nil, false
	}
	return o.ClientPermissions, true
}

// HasClientPermissions returns a boolean if a field has been set.
func (o *NotiForwarder) HasClientPermissions() bool {
	if o != nil && o.ClientPermissions != nil {
		return true
	}

	return false
}

// SetClientPermissions gets a reference to the given []string and assigns it to the ClientPermissions field.
func (o *NotiForwarder) SetClientPermissions(v []string) {
	o.ClientPermissions = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NotiForwarder) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NotiForwarder) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NotiForwarder) SetComment(v string) {
	o.Comment = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *NotiForwarder) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *NotiForwarder) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *NotiForwarder) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *NotiForwarder) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *NotiForwarder) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *NotiForwarder) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *NotiForwarder) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetEventTypesOk() (*[]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *NotiForwarder) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *NotiForwarder) SetEventTypes(v []string) {
	o.EventTypes = &v
}

// GetGatewayClusterId returns the GatewayClusterId field value if set, zero value otherwise.
func (o *NotiForwarder) GetGatewayClusterId() int64 {
	if o == nil || o.GatewayClusterId == nil {
		var ret int64
		return ret
	}
	return *o.GatewayClusterId
}

// GetGatewayClusterIdOk returns a tuple with the GatewayClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetGatewayClusterIdOk() (*int64, bool) {
	if o == nil || o.GatewayClusterId == nil {
		return nil, false
	}
	return o.GatewayClusterId, true
}

// HasGatewayClusterId returns a boolean if a field has been set.
func (o *NotiForwarder) HasGatewayClusterId() bool {
	if o != nil && o.GatewayClusterId != nil {
		return true
	}

	return false
}

// SetGatewayClusterId gets a reference to the given int64 and assigns it to the GatewayClusterId field.
func (o *NotiForwarder) SetGatewayClusterId(v int64) {
	o.GatewayClusterId = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *NotiForwarder) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *NotiForwarder) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *NotiForwarder) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetLastVersion returns the LastVersion field value if set, zero value otherwise.
func (o *NotiForwarder) GetLastVersion() int32 {
	if o == nil || o.LastVersion == nil {
		var ret int32
		return ret
	}
	return *o.LastVersion
}

// GetLastVersionOk returns a tuple with the LastVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetLastVersionOk() (*int32, bool) {
	if o == nil || o.LastVersion == nil {
		return nil, false
	}
	return o.LastVersion, true
}

// HasLastVersion returns a boolean if a field has been set.
func (o *NotiForwarder) HasLastVersion() bool {
	if o != nil && o.LastVersion != nil {
		return true
	}

	return false
}

// SetLastVersion gets a reference to the given int32 and assigns it to the LastVersion field.
func (o *NotiForwarder) SetLastVersion(v int32) {
	o.LastVersion = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *NotiForwarder) GetModificationDate() time.Time {
	if o == nil || o.ModificationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || o.ModificationDate == nil {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *NotiForwarder) HasModificationDate() bool {
	if o != nil && o.ModificationDate != nil {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *NotiForwarder) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

// GetNotiForwarderId returns the NotiForwarderId field value if set, zero value otherwise.
func (o *NotiForwarder) GetNotiForwarderId() int64 {
	if o == nil || o.NotiForwarderId == nil {
		var ret int64
		return ret
	}
	return *o.NotiForwarderId
}

// GetNotiForwarderIdOk returns a tuple with the NotiForwarderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetNotiForwarderIdOk() (*int64, bool) {
	if o == nil || o.NotiForwarderId == nil {
		return nil, false
	}
	return o.NotiForwarderId, true
}

// HasNotiForwarderId returns a boolean if a field has been set.
func (o *NotiForwarder) HasNotiForwarderId() bool {
	if o != nil && o.NotiForwarderId != nil {
		return true
	}

	return false
}

// SetNotiForwarderId gets a reference to the given int64 and assigns it to the NotiForwarderId field.
func (o *NotiForwarder) SetNotiForwarderId(v int64) {
	o.NotiForwarderId = &v
}

// GetNotiForwarderName returns the NotiForwarderName field value if set, zero value otherwise.
func (o *NotiForwarder) GetNotiForwarderName() string {
	if o == nil || o.NotiForwarderName == nil {
		var ret string
		return ret
	}
	return *o.NotiForwarderName
}

// GetNotiForwarderNameOk returns a tuple with the NotiForwarderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetNotiForwarderNameOk() (*string, bool) {
	if o == nil || o.NotiForwarderName == nil {
		return nil, false
	}
	return o.NotiForwarderName, true
}

// HasNotiForwarderName returns a boolean if a field has been set.
func (o *NotiForwarder) HasNotiForwarderName() bool {
	if o != nil && o.NotiForwarderName != nil {
		return true
	}

	return false
}

// SetNotiForwarderName gets a reference to the given string and assigns it to the NotiForwarderName field.
func (o *NotiForwarder) SetNotiForwarderName(v string) {
	o.NotiForwarderName = &v
}

// GetNotiForwarderType returns the NotiForwarderType field value if set, zero value otherwise.
func (o *NotiForwarder) GetNotiForwarderType() string {
	if o == nil || o.NotiForwarderType == nil {
		var ret string
		return ret
	}
	return *o.NotiForwarderType
}

// GetNotiForwarderTypeOk returns a tuple with the NotiForwarderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetNotiForwarderTypeOk() (*string, bool) {
	if o == nil || o.NotiForwarderType == nil {
		return nil, false
	}
	return o.NotiForwarderType, true
}

// HasNotiForwarderType returns a boolean if a field has been set.
func (o *NotiForwarder) HasNotiForwarderType() bool {
	if o != nil && o.NotiForwarderType != nil {
		return true
	}

	return false
}

// SetNotiForwarderType gets a reference to the given string and assigns it to the NotiForwarderType field.
func (o *NotiForwarder) SetNotiForwarderType(v string) {
	o.NotiForwarderType = &v
}

// GetNotiForwarderVersions returns the NotiForwarderVersions field value if set, zero value otherwise.
func (o *NotiForwarder) GetNotiForwarderVersions() []ItemVersion {
	if o == nil || o.NotiForwarderVersions == nil {
		var ret []ItemVersion
		return ret
	}
	return *o.NotiForwarderVersions
}

// GetNotiForwarderVersionsOk returns a tuple with the NotiForwarderVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetNotiForwarderVersionsOk() (*[]ItemVersion, bool) {
	if o == nil || o.NotiForwarderVersions == nil {
		return nil, false
	}
	return o.NotiForwarderVersions, true
}

// HasNotiForwarderVersions returns a boolean if a field has been set.
func (o *NotiForwarder) HasNotiForwarderVersions() bool {
	if o != nil && o.NotiForwarderVersions != nil {
		return true
	}

	return false
}

// SetNotiForwarderVersions gets a reference to the given []ItemVersion and assigns it to the NotiForwarderVersions field.
func (o *NotiForwarder) SetNotiForwarderVersions(v []ItemVersion) {
	o.NotiForwarderVersions = &v
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *NotiForwarder) GetPaths() []string {
	if o == nil || o.Paths == nil {
		var ret []string
		return ret
	}
	return *o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetPathsOk() (*[]string, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *NotiForwarder) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *NotiForwarder) SetPaths(v []string) {
	o.Paths = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *NotiForwarder) GetProtectionKey() string {
	if o == nil || o.ProtectionKey == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetProtectionKeyOk() (*string, bool) {
	if o == nil || o.ProtectionKey == nil {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *NotiForwarder) HasProtectionKey() bool {
	if o != nil && o.ProtectionKey != nil {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *NotiForwarder) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetRunnerType returns the RunnerType field value if set, zero value otherwise.
func (o *NotiForwarder) GetRunnerType() string {
	if o == nil || o.RunnerType == nil {
		var ret string
		return ret
	}
	return *o.RunnerType
}

// GetRunnerTypeOk returns a tuple with the RunnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetRunnerTypeOk() (*string, bool) {
	if o == nil || o.RunnerType == nil {
		return nil, false
	}
	return o.RunnerType, true
}

// HasRunnerType returns a boolean if a field has been set.
func (o *NotiForwarder) HasRunnerType() bool {
	if o != nil && o.RunnerType != nil {
		return true
	}

	return false
}

// SetRunnerType gets a reference to the given string and assigns it to the RunnerType field.
func (o *NotiForwarder) SetRunnerType(v string) {
	o.RunnerType = &v
}

// GetTimespanInSeconds returns the TimespanInSeconds field value if set, zero value otherwise.
func (o *NotiForwarder) GetTimespanInSeconds() int64 {
	if o == nil || o.TimespanInSeconds == nil {
		var ret int64
		return ret
	}
	return *o.TimespanInSeconds
}

// GetTimespanInSecondsOk returns a tuple with the TimespanInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetTimespanInSecondsOk() (*int64, bool) {
	if o == nil || o.TimespanInSeconds == nil {
		return nil, false
	}
	return o.TimespanInSeconds, true
}

// HasTimespanInSeconds returns a boolean if a field has been set.
func (o *NotiForwarder) HasTimespanInSeconds() bool {
	if o != nil && o.TimespanInSeconds != nil {
		return true
	}

	return false
}

// SetTimespanInSeconds gets a reference to the given int64 and assigns it to the TimespanInSeconds field.
func (o *NotiForwarder) SetTimespanInSeconds(v int64) {
	o.TimespanInSeconds = &v
}

// GetToEmails returns the ToEmails field value if set, zero value otherwise.
func (o *NotiForwarder) GetToEmails() []EmailEntry {
	if o == nil || o.ToEmails == nil {
		var ret []EmailEntry
		return ret
	}
	return *o.ToEmails
}

// GetToEmailsOk returns a tuple with the ToEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetToEmailsOk() (*[]EmailEntry, bool) {
	if o == nil || o.ToEmails == nil {
		return nil, false
	}
	return o.ToEmails, true
}

// HasToEmails returns a boolean if a field has been set.
func (o *NotiForwarder) HasToEmails() bool {
	if o != nil && o.ToEmails != nil {
		return true
	}

	return false
}

// SetToEmails gets a reference to the given []EmailEntry and assigns it to the ToEmails field.
func (o *NotiForwarder) SetToEmails(v []EmailEntry) {
	o.ToEmails = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotiForwarder) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotiForwarder) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotiForwarder) SetUsername(v string) {
	o.Username = &v
}

// GetWithCustomerFragment returns the WithCustomerFragment field value if set, zero value otherwise.
func (o *NotiForwarder) GetWithCustomerFragment() bool {
	if o == nil || o.WithCustomerFragment == nil {
		var ret bool
		return ret
	}
	return *o.WithCustomerFragment
}

// GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotiForwarder) GetWithCustomerFragmentOk() (*bool, bool) {
	if o == nil || o.WithCustomerFragment == nil {
		return nil, false
	}
	return o.WithCustomerFragment, true
}

// HasWithCustomerFragment returns a boolean if a field has been set.
func (o *NotiForwarder) HasWithCustomerFragment() bool {
	if o != nil && o.WithCustomerFragment != nil {
		return true
	}

	return false
}

// SetWithCustomerFragment gets a reference to the given bool and assigns it to the WithCustomerFragment field.
func (o *NotiForwarder) SetWithCustomerFragment(v bool) {
	o.WithCustomerFragment = &v
}

func (o NotiForwarder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientPermissions != nil {
		toSerialize["client_permissions"] = o.ClientPermissions
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.EventTypes != nil {
		toSerialize["event_types"] = o.EventTypes
	}
	if o.GatewayClusterId != nil {
		toSerialize["gateway_cluster_id"] = o.GatewayClusterId
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.LastVersion != nil {
		toSerialize["last_version"] = o.LastVersion
	}
	if o.ModificationDate != nil {
		toSerialize["modification_date"] = o.ModificationDate
	}
	if o.NotiForwarderId != nil {
		toSerialize["noti_forwarder_id"] = o.NotiForwarderId
	}
	if o.NotiForwarderName != nil {
		toSerialize["noti_forwarder_name"] = o.NotiForwarderName
	}
	if o.NotiForwarderType != nil {
		toSerialize["noti_forwarder_type"] = o.NotiForwarderType
	}
	if o.NotiForwarderVersions != nil {
		toSerialize["noti_forwarder_versions"] = o.NotiForwarderVersions
	}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	if o.ProtectionKey != nil {
		toSerialize["protection_key"] = o.ProtectionKey
	}
	if o.RunnerType != nil {
		toSerialize["runner_type"] = o.RunnerType
	}
	if o.TimespanInSeconds != nil {
		toSerialize["timespan_in_seconds"] = o.TimespanInSeconds
	}
	if o.ToEmails != nil {
		toSerialize["to_emails"] = o.ToEmails
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.WithCustomerFragment != nil {
		toSerialize["with_customer_fragment"] = o.WithCustomerFragment
	}
	return json.Marshal(toSerialize)
}

type NullableNotiForwarder struct {
	value *NotiForwarder
	isSet bool
}

func (v NullableNotiForwarder) Get() *NotiForwarder {
	return v.value
}

func (v *NullableNotiForwarder) Set(val *NotiForwarder) {
	v.value = val
	v.isSet = true
}

func (v NullableNotiForwarder) IsSet() bool {
	return v.isSet
}

func (v *NullableNotiForwarder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotiForwarder(val *NotiForwarder) *NullableNotiForwarder {
	return &NullableNotiForwarder{value: val, isSet: true}
}

func (v NullableNotiForwarder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotiForwarder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

