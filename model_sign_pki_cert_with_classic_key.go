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

// SignPKICertWithClassicKey struct for SignPKICertWithClassicKey
type SignPKICertWithClassicKey struct {
	// The common name to be included in the PKI certificate
	CommonName *string `json:"common-name,omitempty"`
	// A comma-separated list of the country that will be set in the issued certificate
	Country *string `json:"country,omitempty"`
	// The name of the key to use in the sign PKI Cert process
	DisplayId string `json:"display-id"`
	// DNS Names to be included in the PKI certificate (in a comma-delimited list)
	DnsNames *string `json:"dns-names,omitempty"`
	// key-usage
	KeyUsage *string `json:"key-usage,omitempty"`
	// A comma-separated list of the locality that will be set in the issued certificate
	Locality *string `json:"locality,omitempty"`
	// A comma-separated list of organizational units (OU) that will be set in the issued certificate
	OrganizationalUnits *string `json:"organizational-units,omitempty"`
	// A comma-separated list of organizations (O) that will be set in the issued certificate
	Organizations *string `json:"organizations,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// A comma-separated list of the postal code that will be set in the issued certificate
	PostalCode *string `json:"postal-code,omitempty"`
	// A comma-separated list of the province that will be set in the issued certificate
	Province *string `json:"province,omitempty"`
	// PublicKey using for signing in a PEM format.
	PublicKeyPemData *string `json:"public-key-pem-data,omitempty"`
	// SigningMethod
	SigningMethod string `json:"signing-method"`
	// A comma-separated list of the street address that will be set in the issued certificate
	StreetAddress *string `json:"street-address,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The requested Time To Live for the certificate, use second units
	Ttl int64 `json:"ttl"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The URI Subject Alternative Names to be included in the PKI certificate (in a comma-delimited list)
	UriSans *string `json:"uri-sans,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
	// classic key version
	Version int32 `json:"version"`
}

// NewSignPKICertWithClassicKey instantiates a new SignPKICertWithClassicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignPKICertWithClassicKey(displayId string, signingMethod string, ttl int64, version int32, ) *SignPKICertWithClassicKey {
	this := SignPKICertWithClassicKey{}
	this.DisplayId = displayId
	var keyUsage string = "DigitalSignature,KeyAgreement,KeyEncipherment"
	this.KeyUsage = &keyUsage
	this.SigningMethod = signingMethod
	this.Ttl = ttl
	this.Version = version
	return &this
}

// NewSignPKICertWithClassicKeyWithDefaults instantiates a new SignPKICertWithClassicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignPKICertWithClassicKeyWithDefaults() *SignPKICertWithClassicKey {
	this := SignPKICertWithClassicKey{}
	var keyUsage string = "DigitalSignature,KeyAgreement,KeyEncipherment"
	this.KeyUsage = &keyUsage
	return &this
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *SignPKICertWithClassicKey) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *SignPKICertWithClassicKey) SetCountry(v string) {
	o.Country = &v
}

// GetDisplayId returns the DisplayId field value
func (o *SignPKICertWithClassicKey) GetDisplayId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetDisplayIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayId, true
}

// SetDisplayId sets field value
func (o *SignPKICertWithClassicKey) SetDisplayId(v string) {
	o.DisplayId = v
}

// GetDnsNames returns the DnsNames field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetDnsNames() string {
	if o == nil || o.DnsNames == nil {
		var ret string
		return ret
	}
	return *o.DnsNames
}

// GetDnsNamesOk returns a tuple with the DnsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetDnsNamesOk() (*string, bool) {
	if o == nil || o.DnsNames == nil {
		return nil, false
	}
	return o.DnsNames, true
}

// HasDnsNames returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasDnsNames() bool {
	if o != nil && o.DnsNames != nil {
		return true
	}

	return false
}

// SetDnsNames gets a reference to the given string and assigns it to the DnsNames field.
func (o *SignPKICertWithClassicKey) SetDnsNames(v string) {
	o.DnsNames = &v
}

// GetKeyUsage returns the KeyUsage field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetKeyUsage() string {
	if o == nil || o.KeyUsage == nil {
		var ret string
		return ret
	}
	return *o.KeyUsage
}

// GetKeyUsageOk returns a tuple with the KeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetKeyUsageOk() (*string, bool) {
	if o == nil || o.KeyUsage == nil {
		return nil, false
	}
	return o.KeyUsage, true
}

// HasKeyUsage returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasKeyUsage() bool {
	if o != nil && o.KeyUsage != nil {
		return true
	}

	return false
}

// SetKeyUsage gets a reference to the given string and assigns it to the KeyUsage field.
func (o *SignPKICertWithClassicKey) SetKeyUsage(v string) {
	o.KeyUsage = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetLocality() string {
	if o == nil || o.Locality == nil {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetLocalityOk() (*string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *SignPKICertWithClassicKey) SetLocality(v string) {
	o.Locality = &v
}

// GetOrganizationalUnits returns the OrganizationalUnits field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetOrganizationalUnits() string {
	if o == nil || o.OrganizationalUnits == nil {
		var ret string
		return ret
	}
	return *o.OrganizationalUnits
}

// GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetOrganizationalUnitsOk() (*string, bool) {
	if o == nil || o.OrganizationalUnits == nil {
		return nil, false
	}
	return o.OrganizationalUnits, true
}

// HasOrganizationalUnits returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasOrganizationalUnits() bool {
	if o != nil && o.OrganizationalUnits != nil {
		return true
	}

	return false
}

// SetOrganizationalUnits gets a reference to the given string and assigns it to the OrganizationalUnits field.
func (o *SignPKICertWithClassicKey) SetOrganizationalUnits(v string) {
	o.OrganizationalUnits = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetOrganizations() string {
	if o == nil || o.Organizations == nil {
		var ret string
		return ret
	}
	return *o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetOrganizationsOk() (*string, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given string and assigns it to the Organizations field.
func (o *SignPKICertWithClassicKey) SetOrganizations(v string) {
	o.Organizations = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SignPKICertWithClassicKey) SetPassword(v string) {
	o.Password = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *SignPKICertWithClassicKey) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetProvince() string {
	if o == nil || o.Province == nil {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetProvinceOk() (*string, bool) {
	if o == nil || o.Province == nil {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasProvince() bool {
	if o != nil && o.Province != nil {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *SignPKICertWithClassicKey) SetProvince(v string) {
	o.Province = &v
}

// GetPublicKeyPemData returns the PublicKeyPemData field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetPublicKeyPemData() string {
	if o == nil || o.PublicKeyPemData == nil {
		var ret string
		return ret
	}
	return *o.PublicKeyPemData
}

// GetPublicKeyPemDataOk returns a tuple with the PublicKeyPemData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetPublicKeyPemDataOk() (*string, bool) {
	if o == nil || o.PublicKeyPemData == nil {
		return nil, false
	}
	return o.PublicKeyPemData, true
}

// HasPublicKeyPemData returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasPublicKeyPemData() bool {
	if o != nil && o.PublicKeyPemData != nil {
		return true
	}

	return false
}

// SetPublicKeyPemData gets a reference to the given string and assigns it to the PublicKeyPemData field.
func (o *SignPKICertWithClassicKey) SetPublicKeyPemData(v string) {
	o.PublicKeyPemData = &v
}

// GetSigningMethod returns the SigningMethod field value
func (o *SignPKICertWithClassicKey) GetSigningMethod() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SigningMethod
}

// GetSigningMethodOk returns a tuple with the SigningMethod field value
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetSigningMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SigningMethod, true
}

// SetSigningMethod sets field value
func (o *SignPKICertWithClassicKey) SetSigningMethod(v string) {
	o.SigningMethod = v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *SignPKICertWithClassicKey) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SignPKICertWithClassicKey) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value
func (o *SignPKICertWithClassicKey) GetTtl() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetTtlOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *SignPKICertWithClassicKey) SetTtl(v int64) {
	o.Ttl = v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *SignPKICertWithClassicKey) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUriSans returns the UriSans field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetUriSans() string {
	if o == nil || o.UriSans == nil {
		var ret string
		return ret
	}
	return *o.UriSans
}

// GetUriSansOk returns a tuple with the UriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetUriSansOk() (*string, bool) {
	if o == nil || o.UriSans == nil {
		return nil, false
	}
	return o.UriSans, true
}

// HasUriSans returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasUriSans() bool {
	if o != nil && o.UriSans != nil {
		return true
	}

	return false
}

// SetUriSans gets a reference to the given string and assigns it to the UriSans field.
func (o *SignPKICertWithClassicKey) SetUriSans(v string) {
	o.UriSans = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SignPKICertWithClassicKey) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SignPKICertWithClassicKey) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SignPKICertWithClassicKey) SetUsername(v string) {
	o.Username = &v
}

// GetVersion returns the Version field value
func (o *SignPKICertWithClassicKey) GetVersion() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SignPKICertWithClassicKey) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SignPKICertWithClassicKey) SetVersion(v int32) {
	o.Version = v
}

func (o SignPKICertWithClassicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommonName != nil {
		toSerialize["common-name"] = o.CommonName
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if true {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.DnsNames != nil {
		toSerialize["dns-names"] = o.DnsNames
	}
	if o.KeyUsage != nil {
		toSerialize["key-usage"] = o.KeyUsage
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.OrganizationalUnits != nil {
		toSerialize["organizational-units"] = o.OrganizationalUnits
	}
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PostalCode != nil {
		toSerialize["postal-code"] = o.PostalCode
	}
	if o.Province != nil {
		toSerialize["province"] = o.Province
	}
	if o.PublicKeyPemData != nil {
		toSerialize["public-key-pem-data"] = o.PublicKeyPemData
	}
	if true {
		toSerialize["signing-method"] = o.SigningMethod
	}
	if o.StreetAddress != nil {
		toSerialize["street-address"] = o.StreetAddress
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["ttl"] = o.Ttl
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UriSans != nil {
		toSerialize["uri-sans"] = o.UriSans
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSignPKICertWithClassicKey struct {
	value *SignPKICertWithClassicKey
	isSet bool
}

func (v NullableSignPKICertWithClassicKey) Get() *SignPKICertWithClassicKey {
	return v.value
}

func (v *NullableSignPKICertWithClassicKey) Set(val *SignPKICertWithClassicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSignPKICertWithClassicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSignPKICertWithClassicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignPKICertWithClassicKey(val *SignPKICertWithClassicKey) *NullableSignPKICertWithClassicKey {
	return &NullableSignPKICertWithClassicKey{value: val, isSet: true}
}

func (v NullableSignPKICertWithClassicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignPKICertWithClassicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

