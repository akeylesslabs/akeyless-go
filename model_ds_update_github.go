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

// DsUpdateGithub dsUpdateGithub is a command that updates github dynamic secret
type DsUpdateGithub struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Github app id
	GithubAppId *int64 `json:"github-app-id,omitempty"`
	// App private key
	GithubAppPrivateKey *string `json:"github-app-private-key,omitempty"`
	// Base URL
	GithubBaseUrl *string `json:"github-base-url,omitempty"`
	// Github app installation id
	InstallationId *int64 `json:"installation-id,omitempty"`
	// Repository that the app installation has access to
	InstallationRepository *string `json:"installation-repository,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret new name
	NewName *string `json:"new-name,omitempty"`
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

// NewDsUpdateGithub instantiates a new DsUpdateGithub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDsUpdateGithub(name string, ) *DsUpdateGithub {
	this := DsUpdateGithub{}
	var githubBaseUrl string = "https://api.github.com/"
	this.GithubBaseUrl = &githubBaseUrl
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewDsUpdateGithubWithDefaults instantiates a new DsUpdateGithub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDsUpdateGithubWithDefaults() *DsUpdateGithub {
	this := DsUpdateGithub{}
	var githubBaseUrl string = "https://api.github.com/"
	this.GithubBaseUrl = &githubBaseUrl
	var json bool = false
	this.Json = &json
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DsUpdateGithub) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetGithubAppId returns the GithubAppId field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetGithubAppId() int64 {
	if o == nil || o.GithubAppId == nil {
		var ret int64
		return ret
	}
	return *o.GithubAppId
}

// GetGithubAppIdOk returns a tuple with the GithubAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetGithubAppIdOk() (*int64, bool) {
	if o == nil || o.GithubAppId == nil {
		return nil, false
	}
	return o.GithubAppId, true
}

// HasGithubAppId returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasGithubAppId() bool {
	if o != nil && o.GithubAppId != nil {
		return true
	}

	return false
}

// SetGithubAppId gets a reference to the given int64 and assigns it to the GithubAppId field.
func (o *DsUpdateGithub) SetGithubAppId(v int64) {
	o.GithubAppId = &v
}

// GetGithubAppPrivateKey returns the GithubAppPrivateKey field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetGithubAppPrivateKey() string {
	if o == nil || o.GithubAppPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.GithubAppPrivateKey
}

// GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetGithubAppPrivateKeyOk() (*string, bool) {
	if o == nil || o.GithubAppPrivateKey == nil {
		return nil, false
	}
	return o.GithubAppPrivateKey, true
}

// HasGithubAppPrivateKey returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasGithubAppPrivateKey() bool {
	if o != nil && o.GithubAppPrivateKey != nil {
		return true
	}

	return false
}

// SetGithubAppPrivateKey gets a reference to the given string and assigns it to the GithubAppPrivateKey field.
func (o *DsUpdateGithub) SetGithubAppPrivateKey(v string) {
	o.GithubAppPrivateKey = &v
}

// GetGithubBaseUrl returns the GithubBaseUrl field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetGithubBaseUrl() string {
	if o == nil || o.GithubBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.GithubBaseUrl
}

// GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetGithubBaseUrlOk() (*string, bool) {
	if o == nil || o.GithubBaseUrl == nil {
		return nil, false
	}
	return o.GithubBaseUrl, true
}

// HasGithubBaseUrl returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasGithubBaseUrl() bool {
	if o != nil && o.GithubBaseUrl != nil {
		return true
	}

	return false
}

// SetGithubBaseUrl gets a reference to the given string and assigns it to the GithubBaseUrl field.
func (o *DsUpdateGithub) SetGithubBaseUrl(v string) {
	o.GithubBaseUrl = &v
}

// GetInstallationId returns the InstallationId field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetInstallationId() int64 {
	if o == nil || o.InstallationId == nil {
		var ret int64
		return ret
	}
	return *o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetInstallationIdOk() (*int64, bool) {
	if o == nil || o.InstallationId == nil {
		return nil, false
	}
	return o.InstallationId, true
}

// HasInstallationId returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasInstallationId() bool {
	if o != nil && o.InstallationId != nil {
		return true
	}

	return false
}

// SetInstallationId gets a reference to the given int64 and assigns it to the InstallationId field.
func (o *DsUpdateGithub) SetInstallationId(v int64) {
	o.InstallationId = &v
}

// GetInstallationRepository returns the InstallationRepository field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetInstallationRepository() string {
	if o == nil || o.InstallationRepository == nil {
		var ret string
		return ret
	}
	return *o.InstallationRepository
}

// GetInstallationRepositoryOk returns a tuple with the InstallationRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetInstallationRepositoryOk() (*string, bool) {
	if o == nil || o.InstallationRepository == nil {
		return nil, false
	}
	return o.InstallationRepository, true
}

// HasInstallationRepository returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasInstallationRepository() bool {
	if o != nil && o.InstallationRepository != nil {
		return true
	}

	return false
}

// SetInstallationRepository gets a reference to the given string and assigns it to the InstallationRepository field.
func (o *DsUpdateGithub) SetInstallationRepository(v string) {
	o.InstallationRepository = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DsUpdateGithub) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DsUpdateGithub) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DsUpdateGithub) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DsUpdateGithub) SetNewName(v string) {
	o.NewName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DsUpdateGithub) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DsUpdateGithub) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DsUpdateGithub) SetToken(v string) {
	o.Token = &v
}

// GetTokenPermissions returns the TokenPermissions field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetTokenPermissions() []string {
	if o == nil || o.TokenPermissions == nil {
		var ret []string
		return ret
	}
	return *o.TokenPermissions
}

// GetTokenPermissionsOk returns a tuple with the TokenPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetTokenPermissionsOk() (*[]string, bool) {
	if o == nil || o.TokenPermissions == nil {
		return nil, false
	}
	return o.TokenPermissions, true
}

// HasTokenPermissions returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasTokenPermissions() bool {
	if o != nil && o.TokenPermissions != nil {
		return true
	}

	return false
}

// SetTokenPermissions gets a reference to the given []string and assigns it to the TokenPermissions field.
func (o *DsUpdateGithub) SetTokenPermissions(v []string) {
	o.TokenPermissions = &v
}

// GetTokenRepositories returns the TokenRepositories field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetTokenRepositories() []string {
	if o == nil || o.TokenRepositories == nil {
		var ret []string
		return ret
	}
	return *o.TokenRepositories
}

// GetTokenRepositoriesOk returns a tuple with the TokenRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetTokenRepositoriesOk() (*[]string, bool) {
	if o == nil || o.TokenRepositories == nil {
		return nil, false
	}
	return o.TokenRepositories, true
}

// HasTokenRepositories returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasTokenRepositories() bool {
	if o != nil && o.TokenRepositories != nil {
		return true
	}

	return false
}

// SetTokenRepositories gets a reference to the given []string and assigns it to the TokenRepositories field.
func (o *DsUpdateGithub) SetTokenRepositories(v []string) {
	o.TokenRepositories = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DsUpdateGithub) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateGithub) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DsUpdateGithub) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DsUpdateGithub) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DsUpdateGithub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
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
	if o.InstallationRepository != nil {
		toSerialize["installation-repository"] = o.InstallationRepository
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
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

type NullableDsUpdateGithub struct {
	value *DsUpdateGithub
	isSet bool
}

func (v NullableDsUpdateGithub) Get() *DsUpdateGithub {
	return v.value
}

func (v *NullableDsUpdateGithub) Set(val *DsUpdateGithub) {
	v.value = val
	v.isSet = true
}

func (v NullableDsUpdateGithub) IsSet() bool {
	return v.isSet
}

func (v *NullableDsUpdateGithub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDsUpdateGithub(val *DsUpdateGithub) *NullableDsUpdateGithub {
	return &NullableDsUpdateGithub{value: val, isSet: true}
}

func (v NullableDsUpdateGithub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDsUpdateGithub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


