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

// checks if the PKICertificateIssueDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PKICertificateIssueDetails{}

// PKICertificateIssueDetails struct for PKICertificateIssueDetails
type PKICertificateIssueDetails struct {
	AllowAnyName *bool `json:"allow_any_name,omitempty"`
	AllowSubdomains *bool `json:"allow_subdomains,omitempty"`
	AllowedDomainsList []string `json:"allowed_domains_list,omitempty"`
	AllowedUriSans []string `json:"allowed_uri_sans,omitempty"`
	BasicConstraintsValidForNonCa *bool `json:"basic_constraints_valid_for_non_ca,omitempty"`
	CertificateAuthorityMode *string `json:"certificate_authority_mode,omitempty"`
	ClientFlag *bool `json:"client_flag,omitempty"`
	CodeSigningFlag *bool `json:"code_signing_flag,omitempty"`
	Country []string `json:"country,omitempty"`
	// DestinationPath is the destination to save generated certificates
	DestinationPath *string `json:"destination_path,omitempty"`
	EnforceHostnames *bool `json:"enforce_hostnames,omitempty"`
	// ExpirationNotification holds a list of expiration notices that should be sent in case a certificate is about to expire, this value is being propagated to the Certificate resources that are created
	ExpirationEvents []CertificateExpirationEvent `json:"expiration_events,omitempty"`
	// GWClusterURL is required when CAMode is \"public\" and it defines the cluster URL the PKI should be issued from. The GW cluster must have permissions to read associated target's details
	GwClusterUrl *string `json:"gw_cluster_url,omitempty"`
	IsCa *bool `json:"is_ca,omitempty"`
	KeyBits *int64 `json:"key_bits,omitempty"`
	KeyType *string `json:"key_type,omitempty"`
	KeyUsageList []string `json:"key_usage_list,omitempty"`
	Locality []string `json:"locality,omitempty"`
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
	NotBeforeDuration *int64 `json:"not_before_duration,omitempty"`
	OrganizationList []string `json:"organization_list,omitempty"`
	OrganizationUnitList []string `json:"organization_unit_list,omitempty"`
	PostalCode []string `json:"postal_code,omitempty"`
	// ProtectGeneratedCertificates dictates whether the created certificates should be protected from deletion
	ProtectGeneratedCertificates *bool `json:"protect_generated_certificates,omitempty"`
	Province []string `json:"province,omitempty"`
	RequireCn *bool `json:"require_cn,omitempty"`
	ServerFlag *bool `json:"server_flag,omitempty"`
	StreetAddress []string `json:"street_address,omitempty"`
}

// NewPKICertificateIssueDetails instantiates a new PKICertificateIssueDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPKICertificateIssueDetails() *PKICertificateIssueDetails {
	this := PKICertificateIssueDetails{}
	return &this
}

// NewPKICertificateIssueDetailsWithDefaults instantiates a new PKICertificateIssueDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKICertificateIssueDetailsWithDefaults() *PKICertificateIssueDetails {
	this := PKICertificateIssueDetails{}
	return &this
}

// GetAllowAnyName returns the AllowAnyName field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetAllowAnyName() bool {
	if o == nil || IsNil(o.AllowAnyName) {
		var ret bool
		return ret
	}
	return *o.AllowAnyName
}

// GetAllowAnyNameOk returns a tuple with the AllowAnyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetAllowAnyNameOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAnyName) {
		return nil, false
	}
	return o.AllowAnyName, true
}

// HasAllowAnyName returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasAllowAnyName() bool {
	if o != nil && !IsNil(o.AllowAnyName) {
		return true
	}

	return false
}

// SetAllowAnyName gets a reference to the given bool and assigns it to the AllowAnyName field.
func (o *PKICertificateIssueDetails) SetAllowAnyName(v bool) {
	o.AllowAnyName = &v
}

// GetAllowSubdomains returns the AllowSubdomains field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetAllowSubdomains() bool {
	if o == nil || IsNil(o.AllowSubdomains) {
		var ret bool
		return ret
	}
	return *o.AllowSubdomains
}

// GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetAllowSubdomainsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSubdomains) {
		return nil, false
	}
	return o.AllowSubdomains, true
}

// HasAllowSubdomains returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasAllowSubdomains() bool {
	if o != nil && !IsNil(o.AllowSubdomains) {
		return true
	}

	return false
}

// SetAllowSubdomains gets a reference to the given bool and assigns it to the AllowSubdomains field.
func (o *PKICertificateIssueDetails) SetAllowSubdomains(v bool) {
	o.AllowSubdomains = &v
}

// GetAllowedDomainsList returns the AllowedDomainsList field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetAllowedDomainsList() []string {
	if o == nil || IsNil(o.AllowedDomainsList) {
		var ret []string
		return ret
	}
	return o.AllowedDomainsList
}

// GetAllowedDomainsListOk returns a tuple with the AllowedDomainsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetAllowedDomainsListOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedDomainsList) {
		return nil, false
	}
	return o.AllowedDomainsList, true
}

// HasAllowedDomainsList returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasAllowedDomainsList() bool {
	if o != nil && !IsNil(o.AllowedDomainsList) {
		return true
	}

	return false
}

// SetAllowedDomainsList gets a reference to the given []string and assigns it to the AllowedDomainsList field.
func (o *PKICertificateIssueDetails) SetAllowedDomainsList(v []string) {
	o.AllowedDomainsList = v
}

// GetAllowedUriSans returns the AllowedUriSans field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetAllowedUriSans() []string {
	if o == nil || IsNil(o.AllowedUriSans) {
		var ret []string
		return ret
	}
	return o.AllowedUriSans
}

// GetAllowedUriSansOk returns a tuple with the AllowedUriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetAllowedUriSansOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedUriSans) {
		return nil, false
	}
	return o.AllowedUriSans, true
}

// HasAllowedUriSans returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasAllowedUriSans() bool {
	if o != nil && !IsNil(o.AllowedUriSans) {
		return true
	}

	return false
}

// SetAllowedUriSans gets a reference to the given []string and assigns it to the AllowedUriSans field.
func (o *PKICertificateIssueDetails) SetAllowedUriSans(v []string) {
	o.AllowedUriSans = v
}

// GetBasicConstraintsValidForNonCa returns the BasicConstraintsValidForNonCa field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetBasicConstraintsValidForNonCa() bool {
	if o == nil || IsNil(o.BasicConstraintsValidForNonCa) {
		var ret bool
		return ret
	}
	return *o.BasicConstraintsValidForNonCa
}

// GetBasicConstraintsValidForNonCaOk returns a tuple with the BasicConstraintsValidForNonCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetBasicConstraintsValidForNonCaOk() (*bool, bool) {
	if o == nil || IsNil(o.BasicConstraintsValidForNonCa) {
		return nil, false
	}
	return o.BasicConstraintsValidForNonCa, true
}

// HasBasicConstraintsValidForNonCa returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasBasicConstraintsValidForNonCa() bool {
	if o != nil && !IsNil(o.BasicConstraintsValidForNonCa) {
		return true
	}

	return false
}

// SetBasicConstraintsValidForNonCa gets a reference to the given bool and assigns it to the BasicConstraintsValidForNonCa field.
func (o *PKICertificateIssueDetails) SetBasicConstraintsValidForNonCa(v bool) {
	o.BasicConstraintsValidForNonCa = &v
}

// GetCertificateAuthorityMode returns the CertificateAuthorityMode field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetCertificateAuthorityMode() string {
	if o == nil || IsNil(o.CertificateAuthorityMode) {
		var ret string
		return ret
	}
	return *o.CertificateAuthorityMode
}

// GetCertificateAuthorityModeOk returns a tuple with the CertificateAuthorityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetCertificateAuthorityModeOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateAuthorityMode) {
		return nil, false
	}
	return o.CertificateAuthorityMode, true
}

// HasCertificateAuthorityMode returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasCertificateAuthorityMode() bool {
	if o != nil && !IsNil(o.CertificateAuthorityMode) {
		return true
	}

	return false
}

// SetCertificateAuthorityMode gets a reference to the given string and assigns it to the CertificateAuthorityMode field.
func (o *PKICertificateIssueDetails) SetCertificateAuthorityMode(v string) {
	o.CertificateAuthorityMode = &v
}

// GetClientFlag returns the ClientFlag field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetClientFlag() bool {
	if o == nil || IsNil(o.ClientFlag) {
		var ret bool
		return ret
	}
	return *o.ClientFlag
}

// GetClientFlagOk returns a tuple with the ClientFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetClientFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientFlag) {
		return nil, false
	}
	return o.ClientFlag, true
}

// HasClientFlag returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasClientFlag() bool {
	if o != nil && !IsNil(o.ClientFlag) {
		return true
	}

	return false
}

// SetClientFlag gets a reference to the given bool and assigns it to the ClientFlag field.
func (o *PKICertificateIssueDetails) SetClientFlag(v bool) {
	o.ClientFlag = &v
}

// GetCodeSigningFlag returns the CodeSigningFlag field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetCodeSigningFlag() bool {
	if o == nil || IsNil(o.CodeSigningFlag) {
		var ret bool
		return ret
	}
	return *o.CodeSigningFlag
}

// GetCodeSigningFlagOk returns a tuple with the CodeSigningFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetCodeSigningFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.CodeSigningFlag) {
		return nil, false
	}
	return o.CodeSigningFlag, true
}

// HasCodeSigningFlag returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasCodeSigningFlag() bool {
	if o != nil && !IsNil(o.CodeSigningFlag) {
		return true
	}

	return false
}

// SetCodeSigningFlag gets a reference to the given bool and assigns it to the CodeSigningFlag field.
func (o *PKICertificateIssueDetails) SetCodeSigningFlag(v bool) {
	o.CodeSigningFlag = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetCountry() []string {
	if o == nil || IsNil(o.Country) {
		var ret []string
		return ret
	}
	return o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetCountryOk() ([]string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given []string and assigns it to the Country field.
func (o *PKICertificateIssueDetails) SetCountry(v []string) {
	o.Country = v
}

// GetDestinationPath returns the DestinationPath field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetDestinationPath() string {
	if o == nil || IsNil(o.DestinationPath) {
		var ret string
		return ret
	}
	return *o.DestinationPath
}

// GetDestinationPathOk returns a tuple with the DestinationPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetDestinationPathOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationPath) {
		return nil, false
	}
	return o.DestinationPath, true
}

// HasDestinationPath returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasDestinationPath() bool {
	if o != nil && !IsNil(o.DestinationPath) {
		return true
	}

	return false
}

// SetDestinationPath gets a reference to the given string and assigns it to the DestinationPath field.
func (o *PKICertificateIssueDetails) SetDestinationPath(v string) {
	o.DestinationPath = &v
}

// GetEnforceHostnames returns the EnforceHostnames field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetEnforceHostnames() bool {
	if o == nil || IsNil(o.EnforceHostnames) {
		var ret bool
		return ret
	}
	return *o.EnforceHostnames
}

// GetEnforceHostnamesOk returns a tuple with the EnforceHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetEnforceHostnamesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceHostnames) {
		return nil, false
	}
	return o.EnforceHostnames, true
}

// HasEnforceHostnames returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasEnforceHostnames() bool {
	if o != nil && !IsNil(o.EnforceHostnames) {
		return true
	}

	return false
}

// SetEnforceHostnames gets a reference to the given bool and assigns it to the EnforceHostnames field.
func (o *PKICertificateIssueDetails) SetEnforceHostnames(v bool) {
	o.EnforceHostnames = &v
}

// GetExpirationEvents returns the ExpirationEvents field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetExpirationEvents() []CertificateExpirationEvent {
	if o == nil || IsNil(o.ExpirationEvents) {
		var ret []CertificateExpirationEvent
		return ret
	}
	return o.ExpirationEvents
}

// GetExpirationEventsOk returns a tuple with the ExpirationEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetExpirationEventsOk() ([]CertificateExpirationEvent, bool) {
	if o == nil || IsNil(o.ExpirationEvents) {
		return nil, false
	}
	return o.ExpirationEvents, true
}

// HasExpirationEvents returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasExpirationEvents() bool {
	if o != nil && !IsNil(o.ExpirationEvents) {
		return true
	}

	return false
}

// SetExpirationEvents gets a reference to the given []CertificateExpirationEvent and assigns it to the ExpirationEvents field.
func (o *PKICertificateIssueDetails) SetExpirationEvents(v []CertificateExpirationEvent) {
	o.ExpirationEvents = v
}

// GetGwClusterUrl returns the GwClusterUrl field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetGwClusterUrl() string {
	if o == nil || IsNil(o.GwClusterUrl) {
		var ret string
		return ret
	}
	return *o.GwClusterUrl
}

// GetGwClusterUrlOk returns a tuple with the GwClusterUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetGwClusterUrlOk() (*string, bool) {
	if o == nil || IsNil(o.GwClusterUrl) {
		return nil, false
	}
	return o.GwClusterUrl, true
}

// HasGwClusterUrl returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasGwClusterUrl() bool {
	if o != nil && !IsNil(o.GwClusterUrl) {
		return true
	}

	return false
}

// SetGwClusterUrl gets a reference to the given string and assigns it to the GwClusterUrl field.
func (o *PKICertificateIssueDetails) SetGwClusterUrl(v string) {
	o.GwClusterUrl = &v
}

// GetIsCa returns the IsCa field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetIsCa() bool {
	if o == nil || IsNil(o.IsCa) {
		var ret bool
		return ret
	}
	return *o.IsCa
}

// GetIsCaOk returns a tuple with the IsCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetIsCaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCa) {
		return nil, false
	}
	return o.IsCa, true
}

// HasIsCa returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasIsCa() bool {
	if o != nil && !IsNil(o.IsCa) {
		return true
	}

	return false
}

// SetIsCa gets a reference to the given bool and assigns it to the IsCa field.
func (o *PKICertificateIssueDetails) SetIsCa(v bool) {
	o.IsCa = &v
}

// GetKeyBits returns the KeyBits field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetKeyBits() int64 {
	if o == nil || IsNil(o.KeyBits) {
		var ret int64
		return ret
	}
	return *o.KeyBits
}

// GetKeyBitsOk returns a tuple with the KeyBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetKeyBitsOk() (*int64, bool) {
	if o == nil || IsNil(o.KeyBits) {
		return nil, false
	}
	return o.KeyBits, true
}

// HasKeyBits returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasKeyBits() bool {
	if o != nil && !IsNil(o.KeyBits) {
		return true
	}

	return false
}

// SetKeyBits gets a reference to the given int64 and assigns it to the KeyBits field.
func (o *PKICertificateIssueDetails) SetKeyBits(v int64) {
	o.KeyBits = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetKeyType() string {
	if o == nil || IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasKeyType() bool {
	if o != nil && !IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *PKICertificateIssueDetails) SetKeyType(v string) {
	o.KeyType = &v
}

// GetKeyUsageList returns the KeyUsageList field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetKeyUsageList() []string {
	if o == nil || IsNil(o.KeyUsageList) {
		var ret []string
		return ret
	}
	return o.KeyUsageList
}

// GetKeyUsageListOk returns a tuple with the KeyUsageList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetKeyUsageListOk() ([]string, bool) {
	if o == nil || IsNil(o.KeyUsageList) {
		return nil, false
	}
	return o.KeyUsageList, true
}

// HasKeyUsageList returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasKeyUsageList() bool {
	if o != nil && !IsNil(o.KeyUsageList) {
		return true
	}

	return false
}

// SetKeyUsageList gets a reference to the given []string and assigns it to the KeyUsageList field.
func (o *PKICertificateIssueDetails) SetKeyUsageList(v []string) {
	o.KeyUsageList = v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetLocality() []string {
	if o == nil || IsNil(o.Locality) {
		var ret []string
		return ret
	}
	return o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetLocalityOk() ([]string, bool) {
	if o == nil || IsNil(o.Locality) {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasLocality() bool {
	if o != nil && !IsNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given []string and assigns it to the Locality field.
func (o *PKICertificateIssueDetails) SetLocality(v []string) {
	o.Locality = v
}

// GetNotBeforeDuration returns the NotBeforeDuration field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetNotBeforeDuration() int64 {
	if o == nil || IsNil(o.NotBeforeDuration) {
		var ret int64
		return ret
	}
	return *o.NotBeforeDuration
}

// GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetNotBeforeDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.NotBeforeDuration) {
		return nil, false
	}
	return o.NotBeforeDuration, true
}

// HasNotBeforeDuration returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasNotBeforeDuration() bool {
	if o != nil && !IsNil(o.NotBeforeDuration) {
		return true
	}

	return false
}

// SetNotBeforeDuration gets a reference to the given int64 and assigns it to the NotBeforeDuration field.
func (o *PKICertificateIssueDetails) SetNotBeforeDuration(v int64) {
	o.NotBeforeDuration = &v
}

// GetOrganizationList returns the OrganizationList field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetOrganizationList() []string {
	if o == nil || IsNil(o.OrganizationList) {
		var ret []string
		return ret
	}
	return o.OrganizationList
}

// GetOrganizationListOk returns a tuple with the OrganizationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetOrganizationListOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationList) {
		return nil, false
	}
	return o.OrganizationList, true
}

// HasOrganizationList returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasOrganizationList() bool {
	if o != nil && !IsNil(o.OrganizationList) {
		return true
	}

	return false
}

// SetOrganizationList gets a reference to the given []string and assigns it to the OrganizationList field.
func (o *PKICertificateIssueDetails) SetOrganizationList(v []string) {
	o.OrganizationList = v
}

// GetOrganizationUnitList returns the OrganizationUnitList field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetOrganizationUnitList() []string {
	if o == nil || IsNil(o.OrganizationUnitList) {
		var ret []string
		return ret
	}
	return o.OrganizationUnitList
}

// GetOrganizationUnitListOk returns a tuple with the OrganizationUnitList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetOrganizationUnitListOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationUnitList) {
		return nil, false
	}
	return o.OrganizationUnitList, true
}

// HasOrganizationUnitList returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasOrganizationUnitList() bool {
	if o != nil && !IsNil(o.OrganizationUnitList) {
		return true
	}

	return false
}

// SetOrganizationUnitList gets a reference to the given []string and assigns it to the OrganizationUnitList field.
func (o *PKICertificateIssueDetails) SetOrganizationUnitList(v []string) {
	o.OrganizationUnitList = v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetPostalCode() []string {
	if o == nil || IsNil(o.PostalCode) {
		var ret []string
		return ret
	}
	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetPostalCodeOk() ([]string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given []string and assigns it to the PostalCode field.
func (o *PKICertificateIssueDetails) SetPostalCode(v []string) {
	o.PostalCode = v
}

// GetProtectGeneratedCertificates returns the ProtectGeneratedCertificates field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetProtectGeneratedCertificates() bool {
	if o == nil || IsNil(o.ProtectGeneratedCertificates) {
		var ret bool
		return ret
	}
	return *o.ProtectGeneratedCertificates
}

// GetProtectGeneratedCertificatesOk returns a tuple with the ProtectGeneratedCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetProtectGeneratedCertificatesOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectGeneratedCertificates) {
		return nil, false
	}
	return o.ProtectGeneratedCertificates, true
}

// HasProtectGeneratedCertificates returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasProtectGeneratedCertificates() bool {
	if o != nil && !IsNil(o.ProtectGeneratedCertificates) {
		return true
	}

	return false
}

// SetProtectGeneratedCertificates gets a reference to the given bool and assigns it to the ProtectGeneratedCertificates field.
func (o *PKICertificateIssueDetails) SetProtectGeneratedCertificates(v bool) {
	o.ProtectGeneratedCertificates = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetProvince() []string {
	if o == nil || IsNil(o.Province) {
		var ret []string
		return ret
	}
	return o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetProvinceOk() ([]string, bool) {
	if o == nil || IsNil(o.Province) {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasProvince() bool {
	if o != nil && !IsNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given []string and assigns it to the Province field.
func (o *PKICertificateIssueDetails) SetProvince(v []string) {
	o.Province = v
}

// GetRequireCn returns the RequireCn field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetRequireCn() bool {
	if o == nil || IsNil(o.RequireCn) {
		var ret bool
		return ret
	}
	return *o.RequireCn
}

// GetRequireCnOk returns a tuple with the RequireCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetRequireCnOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireCn) {
		return nil, false
	}
	return o.RequireCn, true
}

// HasRequireCn returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasRequireCn() bool {
	if o != nil && !IsNil(o.RequireCn) {
		return true
	}

	return false
}

// SetRequireCn gets a reference to the given bool and assigns it to the RequireCn field.
func (o *PKICertificateIssueDetails) SetRequireCn(v bool) {
	o.RequireCn = &v
}

// GetServerFlag returns the ServerFlag field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetServerFlag() bool {
	if o == nil || IsNil(o.ServerFlag) {
		var ret bool
		return ret
	}
	return *o.ServerFlag
}

// GetServerFlagOk returns a tuple with the ServerFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetServerFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerFlag) {
		return nil, false
	}
	return o.ServerFlag, true
}

// HasServerFlag returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasServerFlag() bool {
	if o != nil && !IsNil(o.ServerFlag) {
		return true
	}

	return false
}

// SetServerFlag gets a reference to the given bool and assigns it to the ServerFlag field.
func (o *PKICertificateIssueDetails) SetServerFlag(v bool) {
	o.ServerFlag = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetStreetAddress() []string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret []string
		return ret
	}
	return o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetStreetAddressOk() ([]string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given []string and assigns it to the StreetAddress field.
func (o *PKICertificateIssueDetails) SetStreetAddress(v []string) {
	o.StreetAddress = v
}

func (o PKICertificateIssueDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PKICertificateIssueDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowAnyName) {
		toSerialize["allow_any_name"] = o.AllowAnyName
	}
	if !IsNil(o.AllowSubdomains) {
		toSerialize["allow_subdomains"] = o.AllowSubdomains
	}
	if !IsNil(o.AllowedDomainsList) {
		toSerialize["allowed_domains_list"] = o.AllowedDomainsList
	}
	if !IsNil(o.AllowedUriSans) {
		toSerialize["allowed_uri_sans"] = o.AllowedUriSans
	}
	if !IsNil(o.BasicConstraintsValidForNonCa) {
		toSerialize["basic_constraints_valid_for_non_ca"] = o.BasicConstraintsValidForNonCa
	}
	if !IsNil(o.CertificateAuthorityMode) {
		toSerialize["certificate_authority_mode"] = o.CertificateAuthorityMode
	}
	if !IsNil(o.ClientFlag) {
		toSerialize["client_flag"] = o.ClientFlag
	}
	if !IsNil(o.CodeSigningFlag) {
		toSerialize["code_signing_flag"] = o.CodeSigningFlag
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.DestinationPath) {
		toSerialize["destination_path"] = o.DestinationPath
	}
	if !IsNil(o.EnforceHostnames) {
		toSerialize["enforce_hostnames"] = o.EnforceHostnames
	}
	if !IsNil(o.ExpirationEvents) {
		toSerialize["expiration_events"] = o.ExpirationEvents
	}
	if !IsNil(o.GwClusterUrl) {
		toSerialize["gw_cluster_url"] = o.GwClusterUrl
	}
	if !IsNil(o.IsCa) {
		toSerialize["is_ca"] = o.IsCa
	}
	if !IsNil(o.KeyBits) {
		toSerialize["key_bits"] = o.KeyBits
	}
	if !IsNil(o.KeyType) {
		toSerialize["key_type"] = o.KeyType
	}
	if !IsNil(o.KeyUsageList) {
		toSerialize["key_usage_list"] = o.KeyUsageList
	}
	if !IsNil(o.Locality) {
		toSerialize["locality"] = o.Locality
	}
	if !IsNil(o.NotBeforeDuration) {
		toSerialize["not_before_duration"] = o.NotBeforeDuration
	}
	if !IsNil(o.OrganizationList) {
		toSerialize["organization_list"] = o.OrganizationList
	}
	if !IsNil(o.OrganizationUnitList) {
		toSerialize["organization_unit_list"] = o.OrganizationUnitList
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.ProtectGeneratedCertificates) {
		toSerialize["protect_generated_certificates"] = o.ProtectGeneratedCertificates
	}
	if !IsNil(o.Province) {
		toSerialize["province"] = o.Province
	}
	if !IsNil(o.RequireCn) {
		toSerialize["require_cn"] = o.RequireCn
	}
	if !IsNil(o.ServerFlag) {
		toSerialize["server_flag"] = o.ServerFlag
	}
	if !IsNil(o.StreetAddress) {
		toSerialize["street_address"] = o.StreetAddress
	}
	return toSerialize, nil
}

type NullablePKICertificateIssueDetails struct {
	value *PKICertificateIssueDetails
	isSet bool
}

func (v NullablePKICertificateIssueDetails) Get() *PKICertificateIssueDetails {
	return v.value
}

func (v *NullablePKICertificateIssueDetails) Set(val *PKICertificateIssueDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePKICertificateIssueDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePKICertificateIssueDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePKICertificateIssueDetails(val *PKICertificateIssueDetails) *NullablePKICertificateIssueDetails {
	return &NullablePKICertificateIssueDetails{value: val, isSet: true}
}

func (v NullablePKICertificateIssueDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePKICertificateIssueDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


