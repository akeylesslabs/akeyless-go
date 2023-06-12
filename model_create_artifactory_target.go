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

// checks if the CreateArtifactoryTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateArtifactoryTarget{}

// CreateArtifactoryTarget struct for CreateArtifactoryTarget
type CreateArtifactoryTarget struct {
	// Artifactory Admin Name
	ArtifactoryAdminName string `json:"artifactory-admin-name"`
	// Artifactory Admin password
	ArtifactoryAdminPwd string `json:"artifactory-admin-pwd"`
	// Base URL
	BaseUrl string `json:"base-url"`
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of a key used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateArtifactoryTarget instantiates a new CreateArtifactoryTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateArtifactoryTarget(artifactoryAdminName string, artifactoryAdminPwd string, baseUrl string, name string) *CreateArtifactoryTarget {
	this := CreateArtifactoryTarget{}
	this.ArtifactoryAdminName = artifactoryAdminName
	this.ArtifactoryAdminPwd = artifactoryAdminPwd
	this.BaseUrl = baseUrl
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewCreateArtifactoryTargetWithDefaults instantiates a new CreateArtifactoryTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateArtifactoryTargetWithDefaults() *CreateArtifactoryTarget {
	this := CreateArtifactoryTarget{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetArtifactoryAdminName returns the ArtifactoryAdminName field value
func (o *CreateArtifactoryTarget) GetArtifactoryAdminName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactoryAdminName
}

// GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field value
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetArtifactoryAdminNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactoryAdminName, true
}

// SetArtifactoryAdminName sets field value
func (o *CreateArtifactoryTarget) SetArtifactoryAdminName(v string) {
	o.ArtifactoryAdminName = v
}

// GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field value
func (o *CreateArtifactoryTarget) GetArtifactoryAdminPwd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactoryAdminPwd
}

// GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field value
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetArtifactoryAdminPwdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactoryAdminPwd, true
}

// SetArtifactoryAdminPwd sets field value
func (o *CreateArtifactoryTarget) SetArtifactoryAdminPwd(v string) {
	o.ArtifactoryAdminPwd = v
}

// GetBaseUrl returns the BaseUrl field value
func (o *CreateArtifactoryTarget) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *CreateArtifactoryTarget) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateArtifactoryTarget) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateArtifactoryTarget) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateArtifactoryTarget) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateArtifactoryTarget) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateArtifactoryTarget) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateArtifactoryTarget) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateArtifactoryTarget) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateArtifactoryTarget) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateArtifactoryTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateArtifactoryTarget) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateArtifactoryTarget) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateArtifactoryTarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *CreateArtifactoryTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateArtifactoryTarget) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateArtifactoryTarget) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateArtifactoryTarget) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateArtifactoryTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateArtifactoryTarget) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactoryTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateArtifactoryTarget) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateArtifactoryTarget) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateArtifactoryTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateArtifactoryTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artifactory-admin-name"] = o.ArtifactoryAdminName
	toSerialize["artifactory-admin-pwd"] = o.ArtifactoryAdminPwd
	toSerialize["base-url"] = o.BaseUrl
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableCreateArtifactoryTarget struct {
	value *CreateArtifactoryTarget
	isSet bool
}

func (v NullableCreateArtifactoryTarget) Get() *CreateArtifactoryTarget {
	return v.value
}

func (v *NullableCreateArtifactoryTarget) Set(val *CreateArtifactoryTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateArtifactoryTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateArtifactoryTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateArtifactoryTarget(val *CreateArtifactoryTarget) *NullableCreateArtifactoryTarget {
	return &NullableCreateArtifactoryTarget{value: val, isSet: true}
}

func (v NullableCreateArtifactoryTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateArtifactoryTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


