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

// checks if the PathRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PathRule{}

// PathRule struct for PathRule
type PathRule struct {
	Assigners []RuleAssigner `json:"assigners,omitempty"`
	// The approved/denied capabilities in the path
	Capabilities []string `json:"capabilities,omitempty"`
	// flag that indicate that this rule is allowed to be access RemainingAccess of times.
	IsLimitAccess *bool `json:"is_limit_access,omitempty"`
	NumberOfAccessUsed *int64 `json:"number_of_access_used,omitempty"`
	NumberOfAllowedAccess *int64 `json:"number_of_allowed_access,omitempty"`
	// The path the rule refers to
	Path *string `json:"path,omitempty"`
	StartTime *int64 `json:"start_time,omitempty"`
	Ttl *int64 `json:"ttl,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewPathRule instantiates a new PathRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathRule() *PathRule {
	this := PathRule{}
	return &this
}

// NewPathRuleWithDefaults instantiates a new PathRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathRuleWithDefaults() *PathRule {
	this := PathRule{}
	return &this
}

// GetAssigners returns the Assigners field value if set, zero value otherwise.
func (o *PathRule) GetAssigners() []RuleAssigner {
	if o == nil || IsNil(o.Assigners) {
		var ret []RuleAssigner
		return ret
	}
	return o.Assigners
}

// GetAssignersOk returns a tuple with the Assigners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetAssignersOk() ([]RuleAssigner, bool) {
	if o == nil || IsNil(o.Assigners) {
		return nil, false
	}
	return o.Assigners, true
}

// HasAssigners returns a boolean if a field has been set.
func (o *PathRule) HasAssigners() bool {
	if o != nil && !IsNil(o.Assigners) {
		return true
	}

	return false
}

// SetAssigners gets a reference to the given []RuleAssigner and assigns it to the Assigners field.
func (o *PathRule) SetAssigners(v []RuleAssigner) {
	o.Assigners = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *PathRule) GetCapabilities() []string {
	if o == nil || IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *PathRule) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *PathRule) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetIsLimitAccess returns the IsLimitAccess field value if set, zero value otherwise.
func (o *PathRule) GetIsLimitAccess() bool {
	if o == nil || IsNil(o.IsLimitAccess) {
		var ret bool
		return ret
	}
	return *o.IsLimitAccess
}

// GetIsLimitAccessOk returns a tuple with the IsLimitAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetIsLimitAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLimitAccess) {
		return nil, false
	}
	return o.IsLimitAccess, true
}

// HasIsLimitAccess returns a boolean if a field has been set.
func (o *PathRule) HasIsLimitAccess() bool {
	if o != nil && !IsNil(o.IsLimitAccess) {
		return true
	}

	return false
}

// SetIsLimitAccess gets a reference to the given bool and assigns it to the IsLimitAccess field.
func (o *PathRule) SetIsLimitAccess(v bool) {
	o.IsLimitAccess = &v
}

// GetNumberOfAccessUsed returns the NumberOfAccessUsed field value if set, zero value otherwise.
func (o *PathRule) GetNumberOfAccessUsed() int64 {
	if o == nil || IsNil(o.NumberOfAccessUsed) {
		var ret int64
		return ret
	}
	return *o.NumberOfAccessUsed
}

// GetNumberOfAccessUsedOk returns a tuple with the NumberOfAccessUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetNumberOfAccessUsedOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfAccessUsed) {
		return nil, false
	}
	return o.NumberOfAccessUsed, true
}

// HasNumberOfAccessUsed returns a boolean if a field has been set.
func (o *PathRule) HasNumberOfAccessUsed() bool {
	if o != nil && !IsNil(o.NumberOfAccessUsed) {
		return true
	}

	return false
}

// SetNumberOfAccessUsed gets a reference to the given int64 and assigns it to the NumberOfAccessUsed field.
func (o *PathRule) SetNumberOfAccessUsed(v int64) {
	o.NumberOfAccessUsed = &v
}

// GetNumberOfAllowedAccess returns the NumberOfAllowedAccess field value if set, zero value otherwise.
func (o *PathRule) GetNumberOfAllowedAccess() int64 {
	if o == nil || IsNil(o.NumberOfAllowedAccess) {
		var ret int64
		return ret
	}
	return *o.NumberOfAllowedAccess
}

// GetNumberOfAllowedAccessOk returns a tuple with the NumberOfAllowedAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetNumberOfAllowedAccessOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfAllowedAccess) {
		return nil, false
	}
	return o.NumberOfAllowedAccess, true
}

// HasNumberOfAllowedAccess returns a boolean if a field has been set.
func (o *PathRule) HasNumberOfAllowedAccess() bool {
	if o != nil && !IsNil(o.NumberOfAllowedAccess) {
		return true
	}

	return false
}

// SetNumberOfAllowedAccess gets a reference to the given int64 and assigns it to the NumberOfAllowedAccess field.
func (o *PathRule) SetNumberOfAllowedAccess(v int64) {
	o.NumberOfAllowedAccess = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PathRule) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PathRule) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PathRule) SetPath(v string) {
	o.Path = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *PathRule) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *PathRule) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *PathRule) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *PathRule) GetTtl() int64 {
	if o == nil || IsNil(o.Ttl) {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *PathRule) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *PathRule) SetTtl(v int64) {
	o.Ttl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PathRule) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PathRule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PathRule) SetType(v string) {
	o.Type = &v
}

func (o PathRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assigners) {
		toSerialize["assigners"] = o.Assigners
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.IsLimitAccess) {
		toSerialize["is_limit_access"] = o.IsLimitAccess
	}
	if !IsNil(o.NumberOfAccessUsed) {
		toSerialize["number_of_access_used"] = o.NumberOfAccessUsed
	}
	if !IsNil(o.NumberOfAllowedAccess) {
		toSerialize["number_of_allowed_access"] = o.NumberOfAllowedAccess
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePathRule struct {
	value *PathRule
	isSet bool
}

func (v NullablePathRule) Get() *PathRule {
	return v.value
}

func (v *NullablePathRule) Set(val *PathRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePathRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePathRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathRule(val *PathRule) *NullablePathRule {
	return &NullablePathRule{value: val, isSet: true}
}

func (v NullablePathRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


