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

// UpdateAccountSettings struct for UpdateAccountSettings
type UpdateAccountSettings struct {
	// Address
	Address *string `json:"address,omitempty"`
	// City
	City *string `json:"city,omitempty"`
	// Company name
	CompanyName *string `json:"company-name,omitempty"`
	// Country
	Country *string `json:"country,omitempty"`
	// Default ttl
	JwtTtlDefault *int64 `json:"jwt-ttl-default,omitempty"`
	// Maximum ttl
	JwtTtlMax *int64 `json:"jwt-ttl-max,omitempty"`
	// Minimum ttl
	JwtTtlMin *int64 `json:"jwt-ttl-min,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Phone number
	Phone *string `json:"phone,omitempty"`
	// Postal code
	PostalCode *string `json:"postal-code,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewUpdateAccountSettings instantiates a new UpdateAccountSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountSettings() *UpdateAccountSettings {
	this := UpdateAccountSettings{}
	return &this
}

// NewUpdateAccountSettingsWithDefaults instantiates a new UpdateAccountSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountSettingsWithDefaults() *UpdateAccountSettings {
	this := UpdateAccountSettings{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UpdateAccountSettings) SetAddress(v string) {
	o.Address = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *UpdateAccountSettings) SetCity(v string) {
	o.City = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UpdateAccountSettings) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *UpdateAccountSettings) SetCountry(v string) {
	o.Country = &v
}

// GetJwtTtlDefault returns the JwtTtlDefault field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetJwtTtlDefault() int64 {
	if o == nil || o.JwtTtlDefault == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtlDefault
}

// GetJwtTtlDefaultOk returns a tuple with the JwtTtlDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetJwtTtlDefaultOk() (*int64, bool) {
	if o == nil || o.JwtTtlDefault == nil {
		return nil, false
	}
	return o.JwtTtlDefault, true
}

// HasJwtTtlDefault returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasJwtTtlDefault() bool {
	if o != nil && o.JwtTtlDefault != nil {
		return true
	}

	return false
}

// SetJwtTtlDefault gets a reference to the given int64 and assigns it to the JwtTtlDefault field.
func (o *UpdateAccountSettings) SetJwtTtlDefault(v int64) {
	o.JwtTtlDefault = &v
}

// GetJwtTtlMax returns the JwtTtlMax field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetJwtTtlMax() int64 {
	if o == nil || o.JwtTtlMax == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtlMax
}

// GetJwtTtlMaxOk returns a tuple with the JwtTtlMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetJwtTtlMaxOk() (*int64, bool) {
	if o == nil || o.JwtTtlMax == nil {
		return nil, false
	}
	return o.JwtTtlMax, true
}

// HasJwtTtlMax returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasJwtTtlMax() bool {
	if o != nil && o.JwtTtlMax != nil {
		return true
	}

	return false
}

// SetJwtTtlMax gets a reference to the given int64 and assigns it to the JwtTtlMax field.
func (o *UpdateAccountSettings) SetJwtTtlMax(v int64) {
	o.JwtTtlMax = &v
}

// GetJwtTtlMin returns the JwtTtlMin field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetJwtTtlMin() int64 {
	if o == nil || o.JwtTtlMin == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtlMin
}

// GetJwtTtlMinOk returns a tuple with the JwtTtlMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetJwtTtlMinOk() (*int64, bool) {
	if o == nil || o.JwtTtlMin == nil {
		return nil, false
	}
	return o.JwtTtlMin, true
}

// HasJwtTtlMin returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasJwtTtlMin() bool {
	if o != nil && o.JwtTtlMin != nil {
		return true
	}

	return false
}

// SetJwtTtlMin gets a reference to the given int64 and assigns it to the JwtTtlMin field.
func (o *UpdateAccountSettings) SetJwtTtlMin(v int64) {
	o.JwtTtlMin = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateAccountSettings) SetPassword(v string) {
	o.Password = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UpdateAccountSettings) SetPhone(v string) {
	o.Phone = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *UpdateAccountSettings) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateAccountSettings) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateAccountSettings) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateAccountSettings) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettings) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateAccountSettings) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateAccountSettings) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateAccountSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CompanyName != nil {
		toSerialize["company-name"] = o.CompanyName
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.JwtTtlDefault != nil {
		toSerialize["jwt-ttl-default"] = o.JwtTtlDefault
	}
	if o.JwtTtlMax != nil {
		toSerialize["jwt-ttl-max"] = o.JwtTtlMax
	}
	if o.JwtTtlMin != nil {
		toSerialize["jwt-ttl-min"] = o.JwtTtlMin
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.PostalCode != nil {
		toSerialize["postal-code"] = o.PostalCode
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

type NullableUpdateAccountSettings struct {
	value *UpdateAccountSettings
	isSet bool
}

func (v NullableUpdateAccountSettings) Get() *UpdateAccountSettings {
	return v.value
}

func (v *NullableUpdateAccountSettings) Set(val *UpdateAccountSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountSettings(val *UpdateAccountSettings) *NullableUpdateAccountSettings {
	return &NullableUpdateAccountSettings{value: val, isSet: true}
}

func (v NullableUpdateAccountSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


