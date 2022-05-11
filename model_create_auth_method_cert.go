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

// CreateAuthMethodCert createAuthMethodCert is a command that creates a new auth method that will be able to authenticate using a client certificae
type CreateAuthMethodCert struct {
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
	// A list of revoked cert ids
	RevokedCertIds *[]string `json:"revoked-cert-ids,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// A unique identifier (ID) value should be configured, such as common_name or organizational_unit Whenever a user logs in with a token, these authentication types issue a \"sub claim\" that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization.
	UniqueIdentifier string `json:"unique-identifier"`
}

// NewCreateAuthMethodCert instantiates a new CreateAuthMethodCert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodCert(name string, uniqueIdentifier string, ) *CreateAuthMethodCert {
	this := CreateAuthMethodCert{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	this.Name = name
	this.UniqueIdentifier = uniqueIdentifier
	return &this
}

// NewCreateAuthMethodCertWithDefaults instantiates a new CreateAuthMethodCert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodCertWithDefaults() *CreateAuthMethodCert {
	this := CreateAuthMethodCert{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *CreateAuthMethodCert) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetBoundCommonNames returns the BoundCommonNames field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetBoundCommonNames() []string {
	if o == nil || o.BoundCommonNames == nil {
		var ret []string
		return ret
	}
	return *o.BoundCommonNames
}

// GetBoundCommonNamesOk returns a tuple with the BoundCommonNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetBoundCommonNamesOk() (*[]string, bool) {
	if o == nil || o.BoundCommonNames == nil {
		return nil, false
	}
	return o.BoundCommonNames, true
}

// HasBoundCommonNames returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasBoundCommonNames() bool {
	if o != nil && o.BoundCommonNames != nil {
		return true
	}

	return false
}

// SetBoundCommonNames gets a reference to the given []string and assigns it to the BoundCommonNames field.
func (o *CreateAuthMethodCert) SetBoundCommonNames(v []string) {
	o.BoundCommonNames = &v
}

// GetBoundDnsSans returns the BoundDnsSans field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetBoundDnsSans() []string {
	if o == nil || o.BoundDnsSans == nil {
		var ret []string
		return ret
	}
	return *o.BoundDnsSans
}

// GetBoundDnsSansOk returns a tuple with the BoundDnsSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetBoundDnsSansOk() (*[]string, bool) {
	if o == nil || o.BoundDnsSans == nil {
		return nil, false
	}
	return o.BoundDnsSans, true
}

// HasBoundDnsSans returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasBoundDnsSans() bool {
	if o != nil && o.BoundDnsSans != nil {
		return true
	}

	return false
}

// SetBoundDnsSans gets a reference to the given []string and assigns it to the BoundDnsSans field.
func (o *CreateAuthMethodCert) SetBoundDnsSans(v []string) {
	o.BoundDnsSans = &v
}

// GetBoundEmailSans returns the BoundEmailSans field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetBoundEmailSans() []string {
	if o == nil || o.BoundEmailSans == nil {
		var ret []string
		return ret
	}
	return *o.BoundEmailSans
}

// GetBoundEmailSansOk returns a tuple with the BoundEmailSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetBoundEmailSansOk() (*[]string, bool) {
	if o == nil || o.BoundEmailSans == nil {
		return nil, false
	}
	return o.BoundEmailSans, true
}

// HasBoundEmailSans returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasBoundEmailSans() bool {
	if o != nil && o.BoundEmailSans != nil {
		return true
	}

	return false
}

// SetBoundEmailSans gets a reference to the given []string and assigns it to the BoundEmailSans field.
func (o *CreateAuthMethodCert) SetBoundEmailSans(v []string) {
	o.BoundEmailSans = &v
}

// GetBoundExtensions returns the BoundExtensions field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetBoundExtensions() []string {
	if o == nil || o.BoundExtensions == nil {
		var ret []string
		return ret
	}
	return *o.BoundExtensions
}

// GetBoundExtensionsOk returns a tuple with the BoundExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetBoundExtensionsOk() (*[]string, bool) {
	if o == nil || o.BoundExtensions == nil {
		return nil, false
	}
	return o.BoundExtensions, true
}

// HasBoundExtensions returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasBoundExtensions() bool {
	if o != nil && o.BoundExtensions != nil {
		return true
	}

	return false
}

// SetBoundExtensions gets a reference to the given []string and assigns it to the BoundExtensions field.
func (o *CreateAuthMethodCert) SetBoundExtensions(v []string) {
	o.BoundExtensions = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *CreateAuthMethodCert) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetBoundOrganizationalUnits returns the BoundOrganizationalUnits field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetBoundOrganizationalUnits() []string {
	if o == nil || o.BoundOrganizationalUnits == nil {
		var ret []string
		return ret
	}
	return *o.BoundOrganizationalUnits
}

// GetBoundOrganizationalUnitsOk returns a tuple with the BoundOrganizationalUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetBoundOrganizationalUnitsOk() (*[]string, bool) {
	if o == nil || o.BoundOrganizationalUnits == nil {
		return nil, false
	}
	return o.BoundOrganizationalUnits, true
}

// HasBoundOrganizationalUnits returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasBoundOrganizationalUnits() bool {
	if o != nil && o.BoundOrganizationalUnits != nil {
		return true
	}

	return false
}

// SetBoundOrganizationalUnits gets a reference to the given []string and assigns it to the BoundOrganizationalUnits field.
func (o *CreateAuthMethodCert) SetBoundOrganizationalUnits(v []string) {
	o.BoundOrganizationalUnits = &v
}

// GetBoundUriSans returns the BoundUriSans field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetBoundUriSans() []string {
	if o == nil || o.BoundUriSans == nil {
		var ret []string
		return ret
	}
	return *o.BoundUriSans
}

// GetBoundUriSansOk returns a tuple with the BoundUriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetBoundUriSansOk() (*[]string, bool) {
	if o == nil || o.BoundUriSans == nil {
		return nil, false
	}
	return o.BoundUriSans, true
}

// HasBoundUriSans returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasBoundUriSans() bool {
	if o != nil && o.BoundUriSans != nil {
		return true
	}

	return false
}

// SetBoundUriSans gets a reference to the given []string and assigns it to the BoundUriSans field.
func (o *CreateAuthMethodCert) SetBoundUriSans(v []string) {
	o.BoundUriSans = &v
}

// GetCertificateData returns the CertificateData field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetCertificateData() string {
	if o == nil || o.CertificateData == nil {
		var ret string
		return ret
	}
	return *o.CertificateData
}

// GetCertificateDataOk returns a tuple with the CertificateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetCertificateDataOk() (*string, bool) {
	if o == nil || o.CertificateData == nil {
		return nil, false
	}
	return o.CertificateData, true
}

// HasCertificateData returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasCertificateData() bool {
	if o != nil && o.CertificateData != nil {
		return true
	}

	return false
}

// SetCertificateData gets a reference to the given string and assigns it to the CertificateData field.
func (o *CreateAuthMethodCert) SetCertificateData(v string) {
	o.CertificateData = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *CreateAuthMethodCert) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *CreateAuthMethodCert) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *CreateAuthMethodCert) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAuthMethodCert) SetName(v string) {
	o.Name = v
}

// GetRevokedCertIds returns the RevokedCertIds field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetRevokedCertIds() []string {
	if o == nil || o.RevokedCertIds == nil {
		var ret []string
		return ret
	}
	return *o.RevokedCertIds
}

// GetRevokedCertIdsOk returns a tuple with the RevokedCertIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetRevokedCertIdsOk() (*[]string, bool) {
	if o == nil || o.RevokedCertIds == nil {
		return nil, false
	}
	return o.RevokedCertIds, true
}

// HasRevokedCertIds returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasRevokedCertIds() bool {
	if o != nil && o.RevokedCertIds != nil {
		return true
	}

	return false
}

// SetRevokedCertIds gets a reference to the given []string and assigns it to the RevokedCertIds field.
func (o *CreateAuthMethodCert) SetRevokedCertIds(v []string) {
	o.RevokedCertIds = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateAuthMethodCert) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateAuthMethodCert) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateAuthMethodCert) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateAuthMethodCert) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUniqueIdentifier returns the UniqueIdentifier field value
func (o *CreateAuthMethodCert) GetUniqueIdentifier() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniqueIdentifier
}

// GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodCert) GetUniqueIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniqueIdentifier, true
}

// SetUniqueIdentifier sets field value
func (o *CreateAuthMethodCert) SetUniqueIdentifier(v string) {
	o.UniqueIdentifier = v
}

func (o CreateAuthMethodCert) MarshalJSON() ([]byte, error) {
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

type NullableCreateAuthMethodCert struct {
	value *CreateAuthMethodCert
	isSet bool
}

func (v NullableCreateAuthMethodCert) Get() *CreateAuthMethodCert {
	return v.value
}

func (v *NullableCreateAuthMethodCert) Set(val *CreateAuthMethodCert) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodCert) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodCert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodCert(val *CreateAuthMethodCert) *NullableCreateAuthMethodCert {
	return &NullableCreateAuthMethodCert{value: val, isSet: true}
}

func (v NullableCreateAuthMethodCert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodCert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


