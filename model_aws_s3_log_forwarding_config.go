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

// AwsS3LogForwardingConfig struct for AwsS3LogForwardingConfig
type AwsS3LogForwardingConfig struct {
	AwsAccessId *string `json:"aws_access_id,omitempty"`
	AwsAccessKey *string `json:"aws_access_key,omitempty"`
	AwsRegion *string `json:"aws_region,omitempty"`
	BucketName *string `json:"bucket_name,omitempty"`
	LogFolder *string `json:"log_folder,omitempty"`
}

// NewAwsS3LogForwardingConfig instantiates a new AwsS3LogForwardingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsS3LogForwardingConfig() *AwsS3LogForwardingConfig {
	this := AwsS3LogForwardingConfig{}
	return &this
}

// NewAwsS3LogForwardingConfigWithDefaults instantiates a new AwsS3LogForwardingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsS3LogForwardingConfigWithDefaults() *AwsS3LogForwardingConfig {
	this := AwsS3LogForwardingConfig{}
	return &this
}

// GetAwsAccessId returns the AwsAccessId field value if set, zero value otherwise.
func (o *AwsS3LogForwardingConfig) GetAwsAccessId() string {
	if o == nil || o.AwsAccessId == nil {
		var ret string
		return ret
	}
	return *o.AwsAccessId
}

// GetAwsAccessIdOk returns a tuple with the AwsAccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3LogForwardingConfig) GetAwsAccessIdOk() (*string, bool) {
	if o == nil || o.AwsAccessId == nil {
		return nil, false
	}
	return o.AwsAccessId, true
}

// HasAwsAccessId returns a boolean if a field has been set.
func (o *AwsS3LogForwardingConfig) HasAwsAccessId() bool {
	if o != nil && o.AwsAccessId != nil {
		return true
	}

	return false
}

// SetAwsAccessId gets a reference to the given string and assigns it to the AwsAccessId field.
func (o *AwsS3LogForwardingConfig) SetAwsAccessId(v string) {
	o.AwsAccessId = &v
}

// GetAwsAccessKey returns the AwsAccessKey field value if set, zero value otherwise.
func (o *AwsS3LogForwardingConfig) GetAwsAccessKey() string {
	if o == nil || o.AwsAccessKey == nil {
		var ret string
		return ret
	}
	return *o.AwsAccessKey
}

// GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3LogForwardingConfig) GetAwsAccessKeyOk() (*string, bool) {
	if o == nil || o.AwsAccessKey == nil {
		return nil, false
	}
	return o.AwsAccessKey, true
}

// HasAwsAccessKey returns a boolean if a field has been set.
func (o *AwsS3LogForwardingConfig) HasAwsAccessKey() bool {
	if o != nil && o.AwsAccessKey != nil {
		return true
	}

	return false
}

// SetAwsAccessKey gets a reference to the given string and assigns it to the AwsAccessKey field.
func (o *AwsS3LogForwardingConfig) SetAwsAccessKey(v string) {
	o.AwsAccessKey = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *AwsS3LogForwardingConfig) GetAwsRegion() string {
	if o == nil || o.AwsRegion == nil {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3LogForwardingConfig) GetAwsRegionOk() (*string, bool) {
	if o == nil || o.AwsRegion == nil {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *AwsS3LogForwardingConfig) HasAwsRegion() bool {
	if o != nil && o.AwsRegion != nil {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *AwsS3LogForwardingConfig) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *AwsS3LogForwardingConfig) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3LogForwardingConfig) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *AwsS3LogForwardingConfig) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *AwsS3LogForwardingConfig) SetBucketName(v string) {
	o.BucketName = &v
}

// GetLogFolder returns the LogFolder field value if set, zero value otherwise.
func (o *AwsS3LogForwardingConfig) GetLogFolder() string {
	if o == nil || o.LogFolder == nil {
		var ret string
		return ret
	}
	return *o.LogFolder
}

// GetLogFolderOk returns a tuple with the LogFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3LogForwardingConfig) GetLogFolderOk() (*string, bool) {
	if o == nil || o.LogFolder == nil {
		return nil, false
	}
	return o.LogFolder, true
}

// HasLogFolder returns a boolean if a field has been set.
func (o *AwsS3LogForwardingConfig) HasLogFolder() bool {
	if o != nil && o.LogFolder != nil {
		return true
	}

	return false
}

// SetLogFolder gets a reference to the given string and assigns it to the LogFolder field.
func (o *AwsS3LogForwardingConfig) SetLogFolder(v string) {
	o.LogFolder = &v
}

func (o AwsS3LogForwardingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsAccessId != nil {
		toSerialize["aws_access_id"] = o.AwsAccessId
	}
	if o.AwsAccessKey != nil {
		toSerialize["aws_access_key"] = o.AwsAccessKey
	}
	if o.AwsRegion != nil {
		toSerialize["aws_region"] = o.AwsRegion
	}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.LogFolder != nil {
		toSerialize["log_folder"] = o.LogFolder
	}
	return json.Marshal(toSerialize)
}

type NullableAwsS3LogForwardingConfig struct {
	value *AwsS3LogForwardingConfig
	isSet bool
}

func (v NullableAwsS3LogForwardingConfig) Get() *AwsS3LogForwardingConfig {
	return v.value
}

func (v *NullableAwsS3LogForwardingConfig) Set(val *AwsS3LogForwardingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsS3LogForwardingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsS3LogForwardingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsS3LogForwardingConfig(val *AwsS3LogForwardingConfig) *NullableAwsS3LogForwardingConfig {
	return &NullableAwsS3LogForwardingConfig{value: val, isSet: true}
}

func (v NullableAwsS3LogForwardingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsS3LogForwardingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


