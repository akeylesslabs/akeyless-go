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

// CacheConfigPart struct for CacheConfigPart
type CacheConfigPart struct {
	CacheEnable *bool `json:"cache_enable,omitempty"`
	CacheTtl *string `json:"cache_ttl,omitempty"`
	NewProactiveCacheEnable *bool `json:"new_proactive_cache_enable,omitempty"`
	ProactiveCacheDumpInterval *string `json:"proactive_cache_dump_interval,omitempty"`
	ProactiveCacheEnable *bool `json:"proactive_cache_enable,omitempty"`
	ProactiveCacheMinimumFetchingTime *string `json:"proactive_cache_minimum_fetching_time,omitempty"`
}

// NewCacheConfigPart instantiates a new CacheConfigPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheConfigPart() *CacheConfigPart {
	this := CacheConfigPart{}
	return &this
}

// NewCacheConfigPartWithDefaults instantiates a new CacheConfigPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheConfigPartWithDefaults() *CacheConfigPart {
	this := CacheConfigPart{}
	return &this
}

// GetCacheEnable returns the CacheEnable field value if set, zero value otherwise.
func (o *CacheConfigPart) GetCacheEnable() bool {
	if o == nil || o.CacheEnable == nil {
		var ret bool
		return ret
	}
	return *o.CacheEnable
}

// GetCacheEnableOk returns a tuple with the CacheEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheConfigPart) GetCacheEnableOk() (*bool, bool) {
	if o == nil || o.CacheEnable == nil {
		return nil, false
	}
	return o.CacheEnable, true
}

// HasCacheEnable returns a boolean if a field has been set.
func (o *CacheConfigPart) HasCacheEnable() bool {
	if o != nil && o.CacheEnable != nil {
		return true
	}

	return false
}

// SetCacheEnable gets a reference to the given bool and assigns it to the CacheEnable field.
func (o *CacheConfigPart) SetCacheEnable(v bool) {
	o.CacheEnable = &v
}

// GetCacheTtl returns the CacheTtl field value if set, zero value otherwise.
func (o *CacheConfigPart) GetCacheTtl() string {
	if o == nil || o.CacheTtl == nil {
		var ret string
		return ret
	}
	return *o.CacheTtl
}

// GetCacheTtlOk returns a tuple with the CacheTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheConfigPart) GetCacheTtlOk() (*string, bool) {
	if o == nil || o.CacheTtl == nil {
		return nil, false
	}
	return o.CacheTtl, true
}

// HasCacheTtl returns a boolean if a field has been set.
func (o *CacheConfigPart) HasCacheTtl() bool {
	if o != nil && o.CacheTtl != nil {
		return true
	}

	return false
}

// SetCacheTtl gets a reference to the given string and assigns it to the CacheTtl field.
func (o *CacheConfigPart) SetCacheTtl(v string) {
	o.CacheTtl = &v
}

// GetNewProactiveCacheEnable returns the NewProactiveCacheEnable field value if set, zero value otherwise.
func (o *CacheConfigPart) GetNewProactiveCacheEnable() bool {
	if o == nil || o.NewProactiveCacheEnable == nil {
		var ret bool
		return ret
	}
	return *o.NewProactiveCacheEnable
}

// GetNewProactiveCacheEnableOk returns a tuple with the NewProactiveCacheEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheConfigPart) GetNewProactiveCacheEnableOk() (*bool, bool) {
	if o == nil || o.NewProactiveCacheEnable == nil {
		return nil, false
	}
	return o.NewProactiveCacheEnable, true
}

// HasNewProactiveCacheEnable returns a boolean if a field has been set.
func (o *CacheConfigPart) HasNewProactiveCacheEnable() bool {
	if o != nil && o.NewProactiveCacheEnable != nil {
		return true
	}

	return false
}

// SetNewProactiveCacheEnable gets a reference to the given bool and assigns it to the NewProactiveCacheEnable field.
func (o *CacheConfigPart) SetNewProactiveCacheEnable(v bool) {
	o.NewProactiveCacheEnable = &v
}

// GetProactiveCacheDumpInterval returns the ProactiveCacheDumpInterval field value if set, zero value otherwise.
func (o *CacheConfigPart) GetProactiveCacheDumpInterval() string {
	if o == nil || o.ProactiveCacheDumpInterval == nil {
		var ret string
		return ret
	}
	return *o.ProactiveCacheDumpInterval
}

// GetProactiveCacheDumpIntervalOk returns a tuple with the ProactiveCacheDumpInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheConfigPart) GetProactiveCacheDumpIntervalOk() (*string, bool) {
	if o == nil || o.ProactiveCacheDumpInterval == nil {
		return nil, false
	}
	return o.ProactiveCacheDumpInterval, true
}

// HasProactiveCacheDumpInterval returns a boolean if a field has been set.
func (o *CacheConfigPart) HasProactiveCacheDumpInterval() bool {
	if o != nil && o.ProactiveCacheDumpInterval != nil {
		return true
	}

	return false
}

// SetProactiveCacheDumpInterval gets a reference to the given string and assigns it to the ProactiveCacheDumpInterval field.
func (o *CacheConfigPart) SetProactiveCacheDumpInterval(v string) {
	o.ProactiveCacheDumpInterval = &v
}

// GetProactiveCacheEnable returns the ProactiveCacheEnable field value if set, zero value otherwise.
func (o *CacheConfigPart) GetProactiveCacheEnable() bool {
	if o == nil || o.ProactiveCacheEnable == nil {
		var ret bool
		return ret
	}
	return *o.ProactiveCacheEnable
}

// GetProactiveCacheEnableOk returns a tuple with the ProactiveCacheEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheConfigPart) GetProactiveCacheEnableOk() (*bool, bool) {
	if o == nil || o.ProactiveCacheEnable == nil {
		return nil, false
	}
	return o.ProactiveCacheEnable, true
}

// HasProactiveCacheEnable returns a boolean if a field has been set.
func (o *CacheConfigPart) HasProactiveCacheEnable() bool {
	if o != nil && o.ProactiveCacheEnable != nil {
		return true
	}

	return false
}

// SetProactiveCacheEnable gets a reference to the given bool and assigns it to the ProactiveCacheEnable field.
func (o *CacheConfigPart) SetProactiveCacheEnable(v bool) {
	o.ProactiveCacheEnable = &v
}

// GetProactiveCacheMinimumFetchingTime returns the ProactiveCacheMinimumFetchingTime field value if set, zero value otherwise.
func (o *CacheConfigPart) GetProactiveCacheMinimumFetchingTime() string {
	if o == nil || o.ProactiveCacheMinimumFetchingTime == nil {
		var ret string
		return ret
	}
	return *o.ProactiveCacheMinimumFetchingTime
}

// GetProactiveCacheMinimumFetchingTimeOk returns a tuple with the ProactiveCacheMinimumFetchingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheConfigPart) GetProactiveCacheMinimumFetchingTimeOk() (*string, bool) {
	if o == nil || o.ProactiveCacheMinimumFetchingTime == nil {
		return nil, false
	}
	return o.ProactiveCacheMinimumFetchingTime, true
}

// HasProactiveCacheMinimumFetchingTime returns a boolean if a field has been set.
func (o *CacheConfigPart) HasProactiveCacheMinimumFetchingTime() bool {
	if o != nil && o.ProactiveCacheMinimumFetchingTime != nil {
		return true
	}

	return false
}

// SetProactiveCacheMinimumFetchingTime gets a reference to the given string and assigns it to the ProactiveCacheMinimumFetchingTime field.
func (o *CacheConfigPart) SetProactiveCacheMinimumFetchingTime(v string) {
	o.ProactiveCacheMinimumFetchingTime = &v
}

func (o CacheConfigPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CacheEnable != nil {
		toSerialize["cache_enable"] = o.CacheEnable
	}
	if o.CacheTtl != nil {
		toSerialize["cache_ttl"] = o.CacheTtl
	}
	if o.NewProactiveCacheEnable != nil {
		toSerialize["new_proactive_cache_enable"] = o.NewProactiveCacheEnable
	}
	if o.ProactiveCacheDumpInterval != nil {
		toSerialize["proactive_cache_dump_interval"] = o.ProactiveCacheDumpInterval
	}
	if o.ProactiveCacheEnable != nil {
		toSerialize["proactive_cache_enable"] = o.ProactiveCacheEnable
	}
	if o.ProactiveCacheMinimumFetchingTime != nil {
		toSerialize["proactive_cache_minimum_fetching_time"] = o.ProactiveCacheMinimumFetchingTime
	}
	return json.Marshal(toSerialize)
}

type NullableCacheConfigPart struct {
	value *CacheConfigPart
	isSet bool
}

func (v NullableCacheConfigPart) Get() *CacheConfigPart {
	return v.value
}

func (v *NullableCacheConfigPart) Set(val *CacheConfigPart) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheConfigPart) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheConfigPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheConfigPart(val *CacheConfigPart) *NullableCacheConfigPart {
	return &NullableCacheConfigPart{value: val, isSet: true}
}

func (v NullableCacheConfigPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheConfigPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


