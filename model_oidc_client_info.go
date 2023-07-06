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

// OidcClientInfo struct for OidcClientInfo
type OidcClientInfo struct {
	AccessPermissionAssignment *[]AccessPermissionAssignment `json:"access_permission_assignment,omitempty"`
	Audience *[]string `json:"audience,omitempty"`
	ClientId *string `json:"client-id,omitempty"`
	GrantTypes *[]string `json:"grant_types,omitempty"`
	LogoutUris *[]string `json:"logout_uris,omitempty"`
	Public *bool `json:"public,omitempty"`
	RedirectUris *[]string `json:"redirect_uris,omitempty"`
	ResponseTypes *[]string `json:"response_types,omitempty"`
	Scopes *[]string `json:"scopes,omitempty"`
}

// NewOidcClientInfo instantiates a new OidcClientInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcClientInfo() *OidcClientInfo {
	this := OidcClientInfo{}
	return &this
}

// NewOidcClientInfoWithDefaults instantiates a new OidcClientInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcClientInfoWithDefaults() *OidcClientInfo {
	this := OidcClientInfo{}
	return &this
}

// GetAccessPermissionAssignment returns the AccessPermissionAssignment field value if set, zero value otherwise.
func (o *OidcClientInfo) GetAccessPermissionAssignment() []AccessPermissionAssignment {
	if o == nil || o.AccessPermissionAssignment == nil {
		var ret []AccessPermissionAssignment
		return ret
	}
	return *o.AccessPermissionAssignment
}

// GetAccessPermissionAssignmentOk returns a tuple with the AccessPermissionAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcClientInfo) GetAccessPermissionAssignmentOk() (*[]AccessPermissionAssignment, bool) {
	if o == nil || o.AccessPermissionAssignment == nil {
		return nil, false
	}
	return o.AccessPermissionAssignment, true
}

// HasAccessPermissionAssignment returns a boolean if a field has been set.
func (o *OidcClientInfo) HasAccessPermissionAssignment() bool {
	if o != nil && o.AccessPermissionAssignment != nil {
		return true
	}

	return false
}

// SetAccessPermissionAssignment gets a reference to the given []AccessPermissionAssignment and assigns it to the AccessPermissionAssignment field.
func (o *OidcClientInfo) SetAccessPermissionAssignment(v []AccessPermissionAssignment) {
	o.AccessPermissionAssignment = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *OidcClientInfo) GetAudience() []string {
	if o == nil || o.Audience == nil {
		var ret []string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcClientInfo) GetAudienceOk() (*[]string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *OidcClientInfo) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given []string and assigns it to the Audience field.
func (o *OidcClientInfo) SetAudience(v []string) {
	o.Audience = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OidcClientInfo) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcClientInfo) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OidcClientInfo) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OidcClientInfo) SetClientId(v string) {
	o.ClientId = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *OidcClientInfo) GetGrantTypes() []string {
	if o == nil || o.GrantTypes == nil {
		var ret []string
		return ret
	}
	return *o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcClientInfo) GetGrantTypesOk() (*[]string, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *OidcClientInfo) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *OidcClientInfo) SetGrantTypes(v []string) {
	o.GrantTypes = &v
}

// GetLogoutUris returns the LogoutUris field value if set, zero value otherwise.
func (o *OidcClientInfo) GetLogoutUris() []string {
	if o == nil || o.LogoutUris == nil {
		var ret []string
		return ret
	}
	return *o.LogoutUris
}

// GetLogoutUrisOk returns a tuple with the LogoutUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcClientInfo) GetLogoutUrisOk() (*[]string, bool) {
	if o == nil || o.LogoutUris == nil {
		return nil, false
	}
	return o.LogoutUris, true
}

// HasLogoutUris returns a boolean if a field has been set.
func (o *OidcClientInfo) HasLogoutUris() bool {
	if o != nil && o.LogoutUris != nil {
		return true
	}

	return false
}

// SetLogoutUris gets a reference to the given []string and assigns it to the LogoutUris field.
func (o *OidcClientInfo) SetLogoutUris(v []string) {
	o.LogoutUris = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *OidcClientInfo) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcClientInfo) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *OidcClientInfo) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *OidcClientInfo) SetPublic(v bool) {
	o.Public = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *OidcClientInfo) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcClientInfo) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *OidcClientInfo) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *OidcClientInfo) SetRedirectUris(v []string) {
	o.RedirectUris = &v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *OidcClientInfo) GetResponseTypes() []string {
	if o == nil || o.ResponseTypes == nil {
		var ret []string
		return ret
	}
	return *o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcClientInfo) GetResponseTypesOk() (*[]string, bool) {
	if o == nil || o.ResponseTypes == nil {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *OidcClientInfo) HasResponseTypes() bool {
	if o != nil && o.ResponseTypes != nil {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []string and assigns it to the ResponseTypes field.
func (o *OidcClientInfo) SetResponseTypes(v []string) {
	o.ResponseTypes = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OidcClientInfo) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcClientInfo) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OidcClientInfo) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *OidcClientInfo) SetScopes(v []string) {
	o.Scopes = &v
}

func (o OidcClientInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessPermissionAssignment != nil {
		toSerialize["access_permission_assignment"] = o.AccessPermissionAssignment
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.ClientId != nil {
		toSerialize["client-id"] = o.ClientId
	}
	if o.GrantTypes != nil {
		toSerialize["grant_types"] = o.GrantTypes
	}
	if o.LogoutUris != nil {
		toSerialize["logout_uris"] = o.LogoutUris
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	if o.RedirectUris != nil {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.ResponseTypes != nil {
		toSerialize["response_types"] = o.ResponseTypes
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return json.Marshal(toSerialize)
}

type NullableOidcClientInfo struct {
	value *OidcClientInfo
	isSet bool
}

func (v NullableOidcClientInfo) Get() *OidcClientInfo {
	return v.value
}

func (v *NullableOidcClientInfo) Set(val *OidcClientInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcClientInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcClientInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcClientInfo(val *OidcClientInfo) *NullableOidcClientInfo {
	return &NullableOidcClientInfo{value: val, isSet: true}
}

func (v NullableOidcClientInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcClientInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


