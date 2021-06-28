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

// GatewayCreateProducerAzure gatewayCreateProducerAzure is a command that creates azure producer
type GatewayCreateProducerAzure struct {
	// Azure App Object Id
	AppObjId *string `json:"app-obj-id,omitempty"`
	// Azure Client ID
	ClientId string `json:"client-id"`
	// Azure Client Secret
	ClientSecret string `json:"client-secret"`
	// Producer name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Azure Tenant ID
	TenantId string `json:"tenant-id"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User Group Object Id
	UserGroupObjId *string `json:"user-group-obj-id,omitempty"`
	// Azure User portal access
	UserPortalAccess *bool `json:"user-portal-access,omitempty"`
	// User Principal Name
	UserPrincipalName *string `json:"user-principal-name,omitempty"`
	// Azure User programmatic access
	UserProgrammaticAccess *bool `json:"user-programmatic-access,omitempty"`
	// User Role Template Id
	UserRoleTemplateId *string `json:"user-role-template-id,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewGatewayCreateProducerAzure instantiates a new GatewayCreateProducerAzure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerAzure(clientId string, clientSecret string, name string, tenantId string, ) *GatewayCreateProducerAzure {
	this := GatewayCreateProducerAzure{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.Name = name
	this.TenantId = tenantId
	var userPortalAccess bool = false
	this.UserPortalAccess = &userPortalAccess
	var userProgrammaticAccess bool = true
	this.UserProgrammaticAccess = &userProgrammaticAccess
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerAzureWithDefaults instantiates a new GatewayCreateProducerAzure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerAzureWithDefaults() *GatewayCreateProducerAzure {
	this := GatewayCreateProducerAzure{}
	var userPortalAccess bool = false
	this.UserPortalAccess = &userPortalAccess
	var userProgrammaticAccess bool = true
	this.UserProgrammaticAccess = &userProgrammaticAccess
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetAppObjId returns the AppObjId field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetAppObjId() string {
	if o == nil || o.AppObjId == nil {
		var ret string
		return ret
	}
	return *o.AppObjId
}

// GetAppObjIdOk returns a tuple with the AppObjId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetAppObjIdOk() (*string, bool) {
	if o == nil || o.AppObjId == nil {
		return nil, false
	}
	return o.AppObjId, true
}

// HasAppObjId returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasAppObjId() bool {
	if o != nil && o.AppObjId != nil {
		return true
	}

	return false
}

// SetAppObjId gets a reference to the given string and assigns it to the AppObjId field.
func (o *GatewayCreateProducerAzure) SetAppObjId(v string) {
	o.AppObjId = &v
}

// GetClientId returns the ClientId field value
func (o *GatewayCreateProducerAzure) GetClientId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *GatewayCreateProducerAzure) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *GatewayCreateProducerAzure) GetClientSecret() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetClientSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *GatewayCreateProducerAzure) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerAzure) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerAzure) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayCreateProducerAzure) SetPassword(v string) {
	o.Password = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerAzure) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTenantId returns the TenantId field value
func (o *GatewayCreateProducerAzure) GetTenantId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *GatewayCreateProducerAzure) SetTenantId(v string) {
	o.TenantId = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerAzure) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerAzure) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserGroupObjId returns the UserGroupObjId field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserGroupObjId() string {
	if o == nil || o.UserGroupObjId == nil {
		var ret string
		return ret
	}
	return *o.UserGroupObjId
}

// GetUserGroupObjIdOk returns a tuple with the UserGroupObjId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserGroupObjIdOk() (*string, bool) {
	if o == nil || o.UserGroupObjId == nil {
		return nil, false
	}
	return o.UserGroupObjId, true
}

// HasUserGroupObjId returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserGroupObjId() bool {
	if o != nil && o.UserGroupObjId != nil {
		return true
	}

	return false
}

// SetUserGroupObjId gets a reference to the given string and assigns it to the UserGroupObjId field.
func (o *GatewayCreateProducerAzure) SetUserGroupObjId(v string) {
	o.UserGroupObjId = &v
}

// GetUserPortalAccess returns the UserPortalAccess field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserPortalAccess() bool {
	if o == nil || o.UserPortalAccess == nil {
		var ret bool
		return ret
	}
	return *o.UserPortalAccess
}

// GetUserPortalAccessOk returns a tuple with the UserPortalAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserPortalAccessOk() (*bool, bool) {
	if o == nil || o.UserPortalAccess == nil {
		return nil, false
	}
	return o.UserPortalAccess, true
}

// HasUserPortalAccess returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserPortalAccess() bool {
	if o != nil && o.UserPortalAccess != nil {
		return true
	}

	return false
}

// SetUserPortalAccess gets a reference to the given bool and assigns it to the UserPortalAccess field.
func (o *GatewayCreateProducerAzure) SetUserPortalAccess(v bool) {
	o.UserPortalAccess = &v
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil || o.UserPrincipalName == nil {
		return nil, false
	}
	return o.UserPrincipalName, true
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName != nil {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.
func (o *GatewayCreateProducerAzure) SetUserPrincipalName(v string) {
	o.UserPrincipalName = &v
}

// GetUserProgrammaticAccess returns the UserProgrammaticAccess field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserProgrammaticAccess() bool {
	if o == nil || o.UserProgrammaticAccess == nil {
		var ret bool
		return ret
	}
	return *o.UserProgrammaticAccess
}

// GetUserProgrammaticAccessOk returns a tuple with the UserProgrammaticAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserProgrammaticAccessOk() (*bool, bool) {
	if o == nil || o.UserProgrammaticAccess == nil {
		return nil, false
	}
	return o.UserProgrammaticAccess, true
}

// HasUserProgrammaticAccess returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserProgrammaticAccess() bool {
	if o != nil && o.UserProgrammaticAccess != nil {
		return true
	}

	return false
}

// SetUserProgrammaticAccess gets a reference to the given bool and assigns it to the UserProgrammaticAccess field.
func (o *GatewayCreateProducerAzure) SetUserProgrammaticAccess(v bool) {
	o.UserProgrammaticAccess = &v
}

// GetUserRoleTemplateId returns the UserRoleTemplateId field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserRoleTemplateId() string {
	if o == nil || o.UserRoleTemplateId == nil {
		var ret string
		return ret
	}
	return *o.UserRoleTemplateId
}

// GetUserRoleTemplateIdOk returns a tuple with the UserRoleTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserRoleTemplateIdOk() (*string, bool) {
	if o == nil || o.UserRoleTemplateId == nil {
		return nil, false
	}
	return o.UserRoleTemplateId, true
}

// HasUserRoleTemplateId returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserRoleTemplateId() bool {
	if o != nil && o.UserRoleTemplateId != nil {
		return true
	}

	return false
}

// SetUserRoleTemplateId gets a reference to the given string and assigns it to the UserRoleTemplateId field.
func (o *GatewayCreateProducerAzure) SetUserRoleTemplateId(v string) {
	o.UserRoleTemplateId = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerAzure) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayCreateProducerAzure) SetUsername(v string) {
	o.Username = &v
}

func (o GatewayCreateProducerAzure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppObjId != nil {
		toSerialize["app-obj-id"] = o.AppObjId
	}
	if true {
		toSerialize["client-id"] = o.ClientId
	}
	if true {
		toSerialize["client-secret"] = o.ClientSecret
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if true {
		toSerialize["tenant-id"] = o.TenantId
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UserGroupObjId != nil {
		toSerialize["user-group-obj-id"] = o.UserGroupObjId
	}
	if o.UserPortalAccess != nil {
		toSerialize["user-portal-access"] = o.UserPortalAccess
	}
	if o.UserPrincipalName != nil {
		toSerialize["user-principal-name"] = o.UserPrincipalName
	}
	if o.UserProgrammaticAccess != nil {
		toSerialize["user-programmatic-access"] = o.UserProgrammaticAccess
	}
	if o.UserRoleTemplateId != nil {
		toSerialize["user-role-template-id"] = o.UserRoleTemplateId
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerAzure struct {
	value *GatewayCreateProducerAzure
	isSet bool
}

func (v NullableGatewayCreateProducerAzure) Get() *GatewayCreateProducerAzure {
	return v.value
}

func (v *NullableGatewayCreateProducerAzure) Set(val *GatewayCreateProducerAzure) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerAzure) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerAzure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerAzure(val *GatewayCreateProducerAzure) *NullableGatewayCreateProducerAzure {
	return &NullableGatewayCreateProducerAzure{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerAzure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerAzure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


