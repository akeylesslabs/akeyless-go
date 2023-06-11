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

// checks if the Update type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Update{}

// Update struct for Update
type Update struct {
	// Alternative CLI repository url. e.g. https://artifacts.site2.akeyless.io
	ArtifactRepository *string `json:"artifact-repository,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Show the changelog between the current version and the latest one and exit (update will not be performed)
	ShowChangelog *bool `json:"show-changelog,omitempty"`
	// The CLI version
	Version *string `json:"version,omitempty"`
}

// NewUpdate instantiates a new Update object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdate() *Update {
	this := Update{}
	var json bool = false
	this.Json = &json
	var version string = "latest"
	this.Version = &version
	return &this
}

// NewUpdateWithDefaults instantiates a new Update object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWithDefaults() *Update {
	this := Update{}
	var json bool = false
	this.Json = &json
	var version string = "latest"
	this.Version = &version
	return &this
}

// GetArtifactRepository returns the ArtifactRepository field value if set, zero value otherwise.
func (o *Update) GetArtifactRepository() string {
	if o == nil || IsNil(o.ArtifactRepository) {
		var ret string
		return ret
	}
	return *o.ArtifactRepository
}

// GetArtifactRepositoryOk returns a tuple with the ArtifactRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetArtifactRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactRepository) {
		return nil, false
	}
	return o.ArtifactRepository, true
}

// HasArtifactRepository returns a boolean if a field has been set.
func (o *Update) HasArtifactRepository() bool {
	if o != nil && !IsNil(o.ArtifactRepository) {
		return true
	}

	return false
}

// SetArtifactRepository gets a reference to the given string and assigns it to the ArtifactRepository field.
func (o *Update) SetArtifactRepository(v string) {
	o.ArtifactRepository = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *Update) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *Update) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *Update) SetJson(v bool) {
	o.Json = &v
}

// GetShowChangelog returns the ShowChangelog field value if set, zero value otherwise.
func (o *Update) GetShowChangelog() bool {
	if o == nil || IsNil(o.ShowChangelog) {
		var ret bool
		return ret
	}
	return *o.ShowChangelog
}

// GetShowChangelogOk returns a tuple with the ShowChangelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetShowChangelogOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowChangelog) {
		return nil, false
	}
	return o.ShowChangelog, true
}

// HasShowChangelog returns a boolean if a field has been set.
func (o *Update) HasShowChangelog() bool {
	if o != nil && !IsNil(o.ShowChangelog) {
		return true
	}

	return false
}

// SetShowChangelog gets a reference to the given bool and assigns it to the ShowChangelog field.
func (o *Update) SetShowChangelog(v bool) {
	o.ShowChangelog = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Update) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Update) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Update) SetVersion(v string) {
	o.Version = &v
}

func (o Update) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Update) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArtifactRepository) {
		toSerialize["artifact-repository"] = o.ArtifactRepository
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.ShowChangelog) {
		toSerialize["show-changelog"] = o.ShowChangelog
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableUpdate struct {
	value *Update
	isSet bool
}

func (v NullableUpdate) Get() *Update {
	return v.value
}

func (v *NullableUpdate) Set(val *Update) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdate(val *Update) *NullableUpdate {
	return &NullableUpdate{value: val, isSet: true}
}

func (v NullableUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


