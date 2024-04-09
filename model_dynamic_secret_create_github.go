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

// DynamicSecretCreateGithub dynamicSecretCreateGithub is a command that creates github dynamic secret
type DynamicSecretCreateGithub struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Github app id
	GithubAppId *int64 `json:"github-app-id,omitempty"`
	// App private key
	GithubAppPrivateKey *string `json:"github-app-private-key,omitempty"`
	// Base URL
	GithubBaseUrl *string `json:"github-base-url,omitempty"`
	// GitHub application installation id
	InstallationId *int64 `json:"installation-id,omitempty"`
	// Optional, mutually exclusive with installation id, GitHub organization name
	InstallationOrganization *string `json:"installation-organization,omitempty"`
	// Optional, mutually exclusive with installation id, GitHub repository '<owner>/<repo-name>'
	InstallationRepository *string `json:"installation-repository,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Optional - installation token's allowed permissions
	TokenPermissions *[]string `json:"token-permissions,omitempty"`
	// Optional - installation token's allowed repositories
	TokenRepositories *[]string `json:"token-repositories,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDynamicSecretCreateGithub instantiates a new DynamicSecretCreateGithub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretCreateGithub(name string, ) *DynamicSecretCreateGithub {
	this := DynamicSecretCreateGithub{}
	var githubBaseUrl string = "https://api.github.com/"
	this.GithubBaseUrl = &githubBaseUrl
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewDynamicSecretCreateGithubWithDefaults instantiates a new DynamicSecretCreateGithub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretCreateGithubWithDefaults() *DynamicSecretCreateGithub {
	this := DynamicSecretCreateGithub{}
	var githubBaseUrl string = "https://api.github.com/"
	this.GithubBaseUrl = &githubBaseUrl
	var json bool = false
	this.Json = &json
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretCreateGithub) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretCreateGithub) SetDescription(v string) {
	o.Description = &v
}

// GetGithubAppId returns the GithubAppId field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetGithubAppId() int64 {
	if o == nil || o.GithubAppId == nil {
		var ret int64
		return ret
	}
	return *o.GithubAppId
}

// GetGithubAppIdOk returns a tuple with the GithubAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetGithubAppIdOk() (*int64, bool) {
	if o == nil || o.GithubAppId == nil {
		return nil, false
	}
	return o.GithubAppId, true
}

// HasGithubAppId returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasGithubAppId() bool {
	if o != nil && o.GithubAppId != nil {
		return true
	}

	return false
}

// SetGithubAppId gets a reference to the given int64 and assigns it to the GithubAppId field.
func (o *DynamicSecretCreateGithub) SetGithubAppId(v int64) {
	o.GithubAppId = &v
}

// GetGithubAppPrivateKey returns the GithubAppPrivateKey field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetGithubAppPrivateKey() string {
	if o == nil || o.GithubAppPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.GithubAppPrivateKey
}

// GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetGithubAppPrivateKeyOk() (*string, bool) {
	if o == nil || o.GithubAppPrivateKey == nil {
		return nil, false
	}
	return o.GithubAppPrivateKey, true
}

// HasGithubAppPrivateKey returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasGithubAppPrivateKey() bool {
	if o != nil && o.GithubAppPrivateKey != nil {
		return true
	}

	return false
}

// SetGithubAppPrivateKey gets a reference to the given string and assigns it to the GithubAppPrivateKey field.
func (o *DynamicSecretCreateGithub) SetGithubAppPrivateKey(v string) {
	o.GithubAppPrivateKey = &v
}

// GetGithubBaseUrl returns the GithubBaseUrl field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetGithubBaseUrl() string {
	if o == nil || o.GithubBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.GithubBaseUrl
}

// GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetGithubBaseUrlOk() (*string, bool) {
	if o == nil || o.GithubBaseUrl == nil {
		return nil, false
	}
	return o.GithubBaseUrl, true
}

// HasGithubBaseUrl returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasGithubBaseUrl() bool {
	if o != nil && o.GithubBaseUrl != nil {
		return true
	}

	return false
}

// SetGithubBaseUrl gets a reference to the given string and assigns it to the GithubBaseUrl field.
func (o *DynamicSecretCreateGithub) SetGithubBaseUrl(v string) {
	o.GithubBaseUrl = &v
}

// GetInstallationId returns the InstallationId field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetInstallationId() int64 {
	if o == nil || o.InstallationId == nil {
		var ret int64
		return ret
	}
	return *o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetInstallationIdOk() (*int64, bool) {
	if o == nil || o.InstallationId == nil {
		return nil, false
	}
	return o.InstallationId, true
}

// HasInstallationId returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasInstallationId() bool {
	if o != nil && o.InstallationId != nil {
		return true
	}

	return false
}

// SetInstallationId gets a reference to the given int64 and assigns it to the InstallationId field.
func (o *DynamicSecretCreateGithub) SetInstallationId(v int64) {
	o.InstallationId = &v
}

// GetInstallationOrganization returns the InstallationOrganization field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetInstallationOrganization() string {
	if o == nil || o.InstallationOrganization == nil {
		var ret string
		return ret
	}
	return *o.InstallationOrganization
}

// GetInstallationOrganizationOk returns a tuple with the InstallationOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetInstallationOrganizationOk() (*string, bool) {
	if o == nil || o.InstallationOrganization == nil {
		return nil, false
	}
	return o.InstallationOrganization, true
}

// HasInstallationOrganization returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasInstallationOrganization() bool {
	if o != nil && o.InstallationOrganization != nil {
		return true
	}

	return false
}

// SetInstallationOrganization gets a reference to the given string and assigns it to the InstallationOrganization field.
func (o *DynamicSecretCreateGithub) SetInstallationOrganization(v string) {
	o.InstallationOrganization = &v
}

// GetInstallationRepository returns the InstallationRepository field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetInstallationRepository() string {
	if o == nil || o.InstallationRepository == nil {
		var ret string
		return ret
	}
	return *o.InstallationRepository
}

// GetInstallationRepositoryOk returns a tuple with the InstallationRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetInstallationRepositoryOk() (*string, bool) {
	if o == nil || o.InstallationRepository == nil {
		return nil, false
	}
	return o.InstallationRepository, true
}

// HasInstallationRepository returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasInstallationRepository() bool {
	if o != nil && o.InstallationRepository != nil {
		return true
	}

	return false
}

// SetInstallationRepository gets a reference to the given string and assigns it to the InstallationRepository field.
func (o *DynamicSecretCreateGithub) SetInstallationRepository(v string) {
	o.InstallationRepository = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretCreateGithub) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretCreateGithub) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretCreateGithub) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretCreateGithub) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretCreateGithub) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretCreateGithub) SetToken(v string) {
	o.Token = &v
}

// GetTokenPermissions returns the TokenPermissions field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetTokenPermissions() []string {
	if o == nil || o.TokenPermissions == nil {
		var ret []string
		return ret
	}
	return *o.TokenPermissions
}

// GetTokenPermissionsOk returns a tuple with the TokenPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetTokenPermissionsOk() (*[]string, bool) {
	if o == nil || o.TokenPermissions == nil {
		return nil, false
	}
	return o.TokenPermissions, true
}

// HasTokenPermissions returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasTokenPermissions() bool {
	if o != nil && o.TokenPermissions != nil {
		return true
	}

	return false
}

// SetTokenPermissions gets a reference to the given []string and assigns it to the TokenPermissions field.
func (o *DynamicSecretCreateGithub) SetTokenPermissions(v []string) {
	o.TokenPermissions = &v
}

// GetTokenRepositories returns the TokenRepositories field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetTokenRepositories() []string {
	if o == nil || o.TokenRepositories == nil {
		var ret []string
		return ret
	}
	return *o.TokenRepositories
}

// GetTokenRepositoriesOk returns a tuple with the TokenRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetTokenRepositoriesOk() (*[]string, bool) {
	if o == nil || o.TokenRepositories == nil {
		return nil, false
	}
	return o.TokenRepositories, true
}

// HasTokenRepositories returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasTokenRepositories() bool {
	if o != nil && o.TokenRepositories != nil {
		return true
	}

	return false
}

// SetTokenRepositories gets a reference to the given []string and assigns it to the TokenRepositories field.
func (o *DynamicSecretCreateGithub) SetTokenRepositories(v []string) {
	o.TokenRepositories = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretCreateGithub) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateGithub) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateGithub) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretCreateGithub) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DynamicSecretCreateGithub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GithubAppId != nil {
		toSerialize["github-app-id"] = o.GithubAppId
	}
	if o.GithubAppPrivateKey != nil {
		toSerialize["github-app-private-key"] = o.GithubAppPrivateKey
	}
	if o.GithubBaseUrl != nil {
		toSerialize["github-base-url"] = o.GithubBaseUrl
	}
	if o.InstallationId != nil {
		toSerialize["installation-id"] = o.InstallationId
	}
	if o.InstallationOrganization != nil {
		toSerialize["installation-organization"] = o.InstallationOrganization
	}
	if o.InstallationRepository != nil {
		toSerialize["installation-repository"] = o.InstallationRepository
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetName != nil {
		toSerialize["target-name"] = o.TargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.TokenPermissions != nil {
		toSerialize["token-permissions"] = o.TokenPermissions
	}
	if o.TokenRepositories != nil {
		toSerialize["token-repositories"] = o.TokenRepositories
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretCreateGithub struct {
	value *DynamicSecretCreateGithub
	isSet bool
}

func (v NullableDynamicSecretCreateGithub) Get() *DynamicSecretCreateGithub {
	return v.value
}

func (v *NullableDynamicSecretCreateGithub) Set(val *DynamicSecretCreateGithub) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretCreateGithub) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretCreateGithub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretCreateGithub(val *DynamicSecretCreateGithub) *NullableDynamicSecretCreateGithub {
	return &NullableDynamicSecretCreateGithub{value: val, isSet: true}
}

func (v NullableDynamicSecretCreateGithub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretCreateGithub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


