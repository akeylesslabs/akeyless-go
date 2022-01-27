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

// UpdateAuthMethodAWSIAM updateAuthMethodAWSIAM is a command that updates a new Auth Method that will be able to authenticate using AWS IAM credentials.
type UpdateAuthMethodAWSIAM struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// A list of full arns that the access is restricted to
	BoundArn *[]string `json:"bound-arn,omitempty"`
	// A list of AWS account-IDs that the access is restricted to
	BoundAwsAccountId []string `json:"bound-aws-account-id"`
	// A CIDR whitelist with the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// A list of full resource ids that the access is restricted to
	BoundResourceId *[]string `json:"bound-resource-id,omitempty"`
	// A list of full role ids that the access is restricted to
	BoundRoleId *[]string `json:"bound-role-id,omitempty"`
	// A list of full role-name that the access is restricted to
	BoundRoleName *[]string `json:"bound-role-name,omitempty"`
	// A list of full user ids that the access is restricted to
	BoundUserId *[]string `json:"bound-user-id,omitempty"`
	// A list of full user-name that the access is restricted to
	BoundUserName *[]string `json:"bound-user-name,omitempty"`
	// if true: enforce role-association must include sub claims
	ForceSubClaims *bool `json:"force-sub-claims,omitempty"`
	// Jwt TTL
	JwtTtl *int64 `json:"jwt-ttl,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// Auth Method new name
	NewName *string `json:"new-name,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// sts URL
	StsUrl *string `json:"sts-url,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewUpdateAuthMethodAWSIAM instantiates a new UpdateAuthMethodAWSIAM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthMethodAWSIAM(boundAwsAccountId []string, name string, ) *UpdateAuthMethodAWSIAM {
	this := UpdateAuthMethodAWSIAM{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	this.BoundAwsAccountId = boundAwsAccountId
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	this.Name = name
	var stsUrl string = "https://sts.amazonaws.com"
	this.StsUrl = &stsUrl
	return &this
}

// NewUpdateAuthMethodAWSIAMWithDefaults instantiates a new UpdateAuthMethodAWSIAM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthMethodAWSIAMWithDefaults() *UpdateAuthMethodAWSIAM {
	this := UpdateAuthMethodAWSIAM{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	var stsUrl string = "https://sts.amazonaws.com"
	this.StsUrl = &stsUrl
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *UpdateAuthMethodAWSIAM) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetBoundArn returns the BoundArn field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetBoundArn() []string {
	if o == nil || o.BoundArn == nil {
		var ret []string
		return ret
	}
	return *o.BoundArn
}

// GetBoundArnOk returns a tuple with the BoundArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetBoundArnOk() (*[]string, bool) {
	if o == nil || o.BoundArn == nil {
		return nil, false
	}
	return o.BoundArn, true
}

// HasBoundArn returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasBoundArn() bool {
	if o != nil && o.BoundArn != nil {
		return true
	}

	return false
}

// SetBoundArn gets a reference to the given []string and assigns it to the BoundArn field.
func (o *UpdateAuthMethodAWSIAM) SetBoundArn(v []string) {
	o.BoundArn = &v
}

// GetBoundAwsAccountId returns the BoundAwsAccountId field value
func (o *UpdateAuthMethodAWSIAM) GetBoundAwsAccountId() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.BoundAwsAccountId
}

// GetBoundAwsAccountIdOk returns a tuple with the BoundAwsAccountId field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetBoundAwsAccountIdOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BoundAwsAccountId, true
}

// SetBoundAwsAccountId sets field value
func (o *UpdateAuthMethodAWSIAM) SetBoundAwsAccountId(v []string) {
	o.BoundAwsAccountId = v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *UpdateAuthMethodAWSIAM) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetBoundResourceId returns the BoundResourceId field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetBoundResourceId() []string {
	if o == nil || o.BoundResourceId == nil {
		var ret []string
		return ret
	}
	return *o.BoundResourceId
}

// GetBoundResourceIdOk returns a tuple with the BoundResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetBoundResourceIdOk() (*[]string, bool) {
	if o == nil || o.BoundResourceId == nil {
		return nil, false
	}
	return o.BoundResourceId, true
}

// HasBoundResourceId returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasBoundResourceId() bool {
	if o != nil && o.BoundResourceId != nil {
		return true
	}

	return false
}

// SetBoundResourceId gets a reference to the given []string and assigns it to the BoundResourceId field.
func (o *UpdateAuthMethodAWSIAM) SetBoundResourceId(v []string) {
	o.BoundResourceId = &v
}

// GetBoundRoleId returns the BoundRoleId field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetBoundRoleId() []string {
	if o == nil || o.BoundRoleId == nil {
		var ret []string
		return ret
	}
	return *o.BoundRoleId
}

// GetBoundRoleIdOk returns a tuple with the BoundRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetBoundRoleIdOk() (*[]string, bool) {
	if o == nil || o.BoundRoleId == nil {
		return nil, false
	}
	return o.BoundRoleId, true
}

// HasBoundRoleId returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasBoundRoleId() bool {
	if o != nil && o.BoundRoleId != nil {
		return true
	}

	return false
}

// SetBoundRoleId gets a reference to the given []string and assigns it to the BoundRoleId field.
func (o *UpdateAuthMethodAWSIAM) SetBoundRoleId(v []string) {
	o.BoundRoleId = &v
}

// GetBoundRoleName returns the BoundRoleName field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetBoundRoleName() []string {
	if o == nil || o.BoundRoleName == nil {
		var ret []string
		return ret
	}
	return *o.BoundRoleName
}

// GetBoundRoleNameOk returns a tuple with the BoundRoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetBoundRoleNameOk() (*[]string, bool) {
	if o == nil || o.BoundRoleName == nil {
		return nil, false
	}
	return o.BoundRoleName, true
}

// HasBoundRoleName returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasBoundRoleName() bool {
	if o != nil && o.BoundRoleName != nil {
		return true
	}

	return false
}

// SetBoundRoleName gets a reference to the given []string and assigns it to the BoundRoleName field.
func (o *UpdateAuthMethodAWSIAM) SetBoundRoleName(v []string) {
	o.BoundRoleName = &v
}

// GetBoundUserId returns the BoundUserId field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetBoundUserId() []string {
	if o == nil || o.BoundUserId == nil {
		var ret []string
		return ret
	}
	return *o.BoundUserId
}

// GetBoundUserIdOk returns a tuple with the BoundUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetBoundUserIdOk() (*[]string, bool) {
	if o == nil || o.BoundUserId == nil {
		return nil, false
	}
	return o.BoundUserId, true
}

// HasBoundUserId returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasBoundUserId() bool {
	if o != nil && o.BoundUserId != nil {
		return true
	}

	return false
}

// SetBoundUserId gets a reference to the given []string and assigns it to the BoundUserId field.
func (o *UpdateAuthMethodAWSIAM) SetBoundUserId(v []string) {
	o.BoundUserId = &v
}

// GetBoundUserName returns the BoundUserName field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetBoundUserName() []string {
	if o == nil || o.BoundUserName == nil {
		var ret []string
		return ret
	}
	return *o.BoundUserName
}

// GetBoundUserNameOk returns a tuple with the BoundUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetBoundUserNameOk() (*[]string, bool) {
	if o == nil || o.BoundUserName == nil {
		return nil, false
	}
	return o.BoundUserName, true
}

// HasBoundUserName returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasBoundUserName() bool {
	if o != nil && o.BoundUserName != nil {
		return true
	}

	return false
}

// SetBoundUserName gets a reference to the given []string and assigns it to the BoundUserName field.
func (o *UpdateAuthMethodAWSIAM) SetBoundUserName(v []string) {
	o.BoundUserName = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *UpdateAuthMethodAWSIAM) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *UpdateAuthMethodAWSIAM) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *UpdateAuthMethodAWSIAM) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateAuthMethodAWSIAM) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateAuthMethodAWSIAM) SetNewName(v string) {
	o.NewName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateAuthMethodAWSIAM) SetPassword(v string) {
	o.Password = &v
}

// GetStsUrl returns the StsUrl field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetStsUrl() string {
	if o == nil || o.StsUrl == nil {
		var ret string
		return ret
	}
	return *o.StsUrl
}

// GetStsUrlOk returns a tuple with the StsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetStsUrlOk() (*string, bool) {
	if o == nil || o.StsUrl == nil {
		return nil, false
	}
	return o.StsUrl, true
}

// HasStsUrl returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasStsUrl() bool {
	if o != nil && o.StsUrl != nil {
		return true
	}

	return false
}

// SetStsUrl gets a reference to the given string and assigns it to the StsUrl field.
func (o *UpdateAuthMethodAWSIAM) SetStsUrl(v string) {
	o.StsUrl = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateAuthMethodAWSIAM) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateAuthMethodAWSIAM) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateAuthMethodAWSIAM) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodAWSIAM) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateAuthMethodAWSIAM) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateAuthMethodAWSIAM) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateAuthMethodAWSIAM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if o.BoundArn != nil {
		toSerialize["bound-arn"] = o.BoundArn
	}
	if true {
		toSerialize["bound-aws-account-id"] = o.BoundAwsAccountId
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if o.BoundResourceId != nil {
		toSerialize["bound-resource-id"] = o.BoundResourceId
	}
	if o.BoundRoleId != nil {
		toSerialize["bound-role-id"] = o.BoundRoleId
	}
	if o.BoundRoleName != nil {
		toSerialize["bound-role-name"] = o.BoundRoleName
	}
	if o.BoundUserId != nil {
		toSerialize["bound-user-id"] = o.BoundUserId
	}
	if o.BoundUserName != nil {
		toSerialize["bound-user-name"] = o.BoundUserName
	}
	if o.ForceSubClaims != nil {
		toSerialize["force-sub-claims"] = o.ForceSubClaims
	}
	if o.JwtTtl != nil {
		toSerialize["jwt-ttl"] = o.JwtTtl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.StsUrl != nil {
		toSerialize["sts-url"] = o.StsUrl
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAuthMethodAWSIAM struct {
	value *UpdateAuthMethodAWSIAM
	isSet bool
}

func (v NullableUpdateAuthMethodAWSIAM) Get() *UpdateAuthMethodAWSIAM {
	return v.value
}

func (v *NullableUpdateAuthMethodAWSIAM) Set(val *UpdateAuthMethodAWSIAM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthMethodAWSIAM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthMethodAWSIAM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthMethodAWSIAM(val *UpdateAuthMethodAWSIAM) *NullableUpdateAuthMethodAWSIAM {
	return &NullableUpdateAuthMethodAWSIAM{value: val, isSet: true}
}

func (v NullableUpdateAuthMethodAWSIAM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthMethodAWSIAM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


