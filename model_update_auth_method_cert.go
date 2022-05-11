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

// UpdateAuthMethodCert updateAuthMethodCert is a command that updates a new auth method that will be able to authenticate using a client certificae
type UpdateAuthMethodCert struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// A list of names. At least one must exist in the Common Name. Supports globbing.
	BoundCommonNames *[]string `json:"bound-common-names,omitempty"`
	// A list of DNS names. At least one must exist in the SANs. Supports globbing.
	BoundDnsSans *[]string `json:"bound-dns-sans,omitempty"`
	// A list of Email Addresses. At least one must exist in the SANs. Supports globbing.
	BoundEmailSans *[]string `json:"bound-email-sans,omitempty"`
	// A list of extensions formatted as \"oid:value\". Expects the extension value to be some type of ASN1 encoded string. All values much match. Supports globbing on \"value\".
	BoundExtensions *[]string `json:"bound-extensions,omitempty"`
	// A CIDR whitelist with the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// A list of Organizational Units names. At least one must exist in the OU field.
	BoundOrganizationalUnits *[]string `json:"bound-organizational-units,omitempty"`
	// A list of URIs. At least one must exist in the SANs. Supports globbing.
	BoundUriSans *[]string `json:"bound-uri-sans,omitempty"`
	// The certificate data in base64, if no file was provided
	CertificateData *string `json:"certificate-data,omitempty"`
	// if true: enforce role-association must include sub claims
	ForceSubClaims *bool `json:"force-sub-claims,omitempty"`
	// Jwt TTL
	JwtTtl *int64 `json:"jwt-ttl,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// Auth Method new name
	NewName *string `json:"new-name,omitempty"`
	// A list of revoked cert ids
	RevokedCertIds *[]string `json:"revoked-cert-ids,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// A unique identifier (ID) value should be configured, such as common_name or organizational_unit Whenever a user logs in with a token, these authentication types issue a \"sub claim\" that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization.
	UniqueIdentifier string `json:"unique-identifier"`
}

// NewUpdateAuthMethodCert instantiates a new UpdateAuthMethodCert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthMethodCert(name string, uniqueIdentifier string, ) *UpdateAuthMethodCert {
	this := UpdateAuthMethodCert{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	this.Name = name
	this.UniqueIdentifier = uniqueIdentifier
	return &this
}

// NewUpdateAuthMethodCertWithDefaults instantiates a new UpdateAuthMethodCert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthMethodCertWithDefaults() *UpdateAuthMethodCert {
	this := UpdateAuthMethodCert{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *UpdateAuthMethodCert) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetBoundCommonNames returns the BoundCommonNames field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetBoundCommonNames() []string {
	if o == nil || o.BoundCommonNames == nil {
		var ret []string
		return ret
	}
	return *o.BoundCommonNames
}

// GetBoundCommonNamesOk returns a tuple with the BoundCommonNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetBoundCommonNamesOk() (*[]string, bool) {
	if o == nil || o.BoundCommonNames == nil {
		return nil, false
	}
	return o.BoundCommonNames, true
}

// HasBoundCommonNames returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasBoundCommonNames() bool {
	if o != nil && o.BoundCommonNames != nil {
		return true
	}

	return false
}

// SetBoundCommonNames gets a reference to the given []string and assigns it to the BoundCommonNames field.
func (o *UpdateAuthMethodCert) SetBoundCommonNames(v []string) {
	o.BoundCommonNames = &v
}

// GetBoundDnsSans returns the BoundDnsSans field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetBoundDnsSans() []string {
	if o == nil || o.BoundDnsSans == nil {
		var ret []string
		return ret
	}
	return *o.BoundDnsSans
}

// GetBoundDnsSansOk returns a tuple with the BoundDnsSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetBoundDnsSansOk() (*[]string, bool) {
	if o == nil || o.BoundDnsSans == nil {
		return nil, false
	}
	return o.BoundDnsSans, true
}

// HasBoundDnsSans returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasBoundDnsSans() bool {
	if o != nil && o.BoundDnsSans != nil {
		return true
	}

	return false
}

// SetBoundDnsSans gets a reference to the given []string and assigns it to the BoundDnsSans field.
func (o *UpdateAuthMethodCert) SetBoundDnsSans(v []string) {
	o.BoundDnsSans = &v
}

// GetBoundEmailSans returns the BoundEmailSans field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetBoundEmailSans() []string {
	if o == nil || o.BoundEmailSans == nil {
		var ret []string
		return ret
	}
	return *o.BoundEmailSans
}

// GetBoundEmailSansOk returns a tuple with the BoundEmailSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetBoundEmailSansOk() (*[]string, bool) {
	if o == nil || o.BoundEmailSans == nil {
		return nil, false
	}
	return o.BoundEmailSans, true
}

// HasBoundEmailSans returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasBoundEmailSans() bool {
	if o != nil && o.BoundEmailSans != nil {
		return true
	}

	return false
}

// SetBoundEmailSans gets a reference to the given []string and assigns it to the BoundEmailSans field.
func (o *UpdateAuthMethodCert) SetBoundEmailSans(v []string) {
	o.BoundEmailSans = &v
}

// GetBoundExtensions returns the BoundExtensions field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetBoundExtensions() []string {
	if o == nil || o.BoundExtensions == nil {
		var ret []string
		return ret
	}
	return *o.BoundExtensions
}

// GetBoundExtensionsOk returns a tuple with the BoundExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetBoundExtensionsOk() (*[]string, bool) {
	if o == nil || o.BoundExtensions == nil {
		return nil, false
	}
	return o.BoundExtensions, true
}

// HasBoundExtensions returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasBoundExtensions() bool {
	if o != nil && o.BoundExtensions != nil {
		return true
	}

	return false
}

// SetBoundExtensions gets a reference to the given []string and assigns it to the BoundExtensions field.
func (o *UpdateAuthMethodCert) SetBoundExtensions(v []string) {
	o.BoundExtensions = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *UpdateAuthMethodCert) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetBoundOrganizationalUnits returns the BoundOrganizationalUnits field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetBoundOrganizationalUnits() []string {
	if o == nil || o.BoundOrganizationalUnits == nil {
		var ret []string
		return ret
	}
	return *o.BoundOrganizationalUnits
}

// GetBoundOrganizationalUnitsOk returns a tuple with the BoundOrganizationalUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetBoundOrganizationalUnitsOk() (*[]string, bool) {
	if o == nil || o.BoundOrganizationalUnits == nil {
		return nil, false
	}
	return o.BoundOrganizationalUnits, true
}

// HasBoundOrganizationalUnits returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasBoundOrganizationalUnits() bool {
	if o != nil && o.BoundOrganizationalUnits != nil {
		return true
	}

	return false
}

// SetBoundOrganizationalUnits gets a reference to the given []string and assigns it to the BoundOrganizationalUnits field.
func (o *UpdateAuthMethodCert) SetBoundOrganizationalUnits(v []string) {
	o.BoundOrganizationalUnits = &v
}

// GetBoundUriSans returns the BoundUriSans field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetBoundUriSans() []string {
	if o == nil || o.BoundUriSans == nil {
		var ret []string
		return ret
	}
	return *o.BoundUriSans
}

// GetBoundUriSansOk returns a tuple with the BoundUriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetBoundUriSansOk() (*[]string, bool) {
	if o == nil || o.BoundUriSans == nil {
		return nil, false
	}
	return o.BoundUriSans, true
}

// HasBoundUriSans returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasBoundUriSans() bool {
	if o != nil && o.BoundUriSans != nil {
		return true
	}

	return false
}

// SetBoundUriSans gets a reference to the given []string and assigns it to the BoundUriSans field.
func (o *UpdateAuthMethodCert) SetBoundUriSans(v []string) {
	o.BoundUriSans = &v
}

// GetCertificateData returns the CertificateData field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetCertificateData() string {
	if o == nil || o.CertificateData == nil {
		var ret string
		return ret
	}
	return *o.CertificateData
}

// GetCertificateDataOk returns a tuple with the CertificateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetCertificateDataOk() (*string, bool) {
	if o == nil || o.CertificateData == nil {
		return nil, false
	}
	return o.CertificateData, true
}

// HasCertificateData returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasCertificateData() bool {
	if o != nil && o.CertificateData != nil {
		return true
	}

	return false
}

// SetCertificateData gets a reference to the given string and assigns it to the CertificateData field.
func (o *UpdateAuthMethodCert) SetCertificateData(v string) {
	o.CertificateData = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *UpdateAuthMethodCert) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *UpdateAuthMethodCert) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *UpdateAuthMethodCert) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateAuthMethodCert) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateAuthMethodCert) SetNewName(v string) {
	o.NewName = &v
}

// GetRevokedCertIds returns the RevokedCertIds field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetRevokedCertIds() []string {
	if o == nil || o.RevokedCertIds == nil {
		var ret []string
		return ret
	}
	return *o.RevokedCertIds
}

// GetRevokedCertIdsOk returns a tuple with the RevokedCertIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetRevokedCertIdsOk() (*[]string, bool) {
	if o == nil || o.RevokedCertIds == nil {
		return nil, false
	}
	return o.RevokedCertIds, true
}

// HasRevokedCertIds returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasRevokedCertIds() bool {
	if o != nil && o.RevokedCertIds != nil {
		return true
	}

	return false
}

// SetRevokedCertIds gets a reference to the given []string and assigns it to the RevokedCertIds field.
func (o *UpdateAuthMethodCert) SetRevokedCertIds(v []string) {
	o.RevokedCertIds = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateAuthMethodCert) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateAuthMethodCert) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodCert) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateAuthMethodCert) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUniqueIdentifier returns the UniqueIdentifier field value
func (o *UpdateAuthMethodCert) GetUniqueIdentifier() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniqueIdentifier
}

// GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodCert) GetUniqueIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniqueIdentifier, true
}

// SetUniqueIdentifier sets field value
func (o *UpdateAuthMethodCert) SetUniqueIdentifier(v string) {
	o.UniqueIdentifier = v
}

func (o UpdateAuthMethodCert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if o.BoundCommonNames != nil {
		toSerialize["bound-common-names"] = o.BoundCommonNames
	}
	if o.BoundDnsSans != nil {
		toSerialize["bound-dns-sans"] = o.BoundDnsSans
	}
	if o.BoundEmailSans != nil {
		toSerialize["bound-email-sans"] = o.BoundEmailSans
	}
	if o.BoundExtensions != nil {
		toSerialize["bound-extensions"] = o.BoundExtensions
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if o.BoundOrganizationalUnits != nil {
		toSerialize["bound-organizational-units"] = o.BoundOrganizationalUnits
	}
	if o.BoundUriSans != nil {
		toSerialize["bound-uri-sans"] = o.BoundUriSans
	}
	if o.CertificateData != nil {
		toSerialize["certificate-data"] = o.CertificateData
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
	if o.RevokedCertIds != nil {
		toSerialize["revoked-cert-ids"] = o.RevokedCertIds
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if true {
		toSerialize["unique-identifier"] = o.UniqueIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAuthMethodCert struct {
	value *UpdateAuthMethodCert
	isSet bool
}

func (v NullableUpdateAuthMethodCert) Get() *UpdateAuthMethodCert {
	return v.value
}

func (v *NullableUpdateAuthMethodCert) Set(val *UpdateAuthMethodCert) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthMethodCert) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthMethodCert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthMethodCert(val *UpdateAuthMethodCert) *NullableUpdateAuthMethodCert {
	return &NullableUpdateAuthMethodCert{value: val, isSet: true}
}

func (v NullableUpdateAuthMethodCert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthMethodCert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


