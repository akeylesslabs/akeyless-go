/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// SAMLAccessRules struct for SAMLAccessRules
type SAMLAccessRules struct {
	// The attributes that login is restricted to.
	BoundAttributes *[]SAMLAttribute `json:"bound_attributes,omitempty"`
	// IDP metadata url
	IdpMetadataUrl *string `json:"idp_metadata_url,omitempty"`
	// IDP metadata XML
	IdpMetadataXml *string `json:"idp_metadata_xml,omitempty"`
}

// NewSAMLAccessRules instantiates a new SAMLAccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLAccessRules() *SAMLAccessRules {
	this := SAMLAccessRules{}
	return &this
}

// NewSAMLAccessRulesWithDefaults instantiates a new SAMLAccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLAccessRulesWithDefaults() *SAMLAccessRules {
	this := SAMLAccessRules{}
	return &this
}

// GetBoundAttributes returns the BoundAttributes field value if set, zero value otherwise.
func (o *SAMLAccessRules) GetBoundAttributes() []SAMLAttribute {
	if o == nil || o.BoundAttributes == nil {
		var ret []SAMLAttribute
		return ret
	}
	return *o.BoundAttributes
}

// GetBoundAttributesOk returns a tuple with the BoundAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAccessRules) GetBoundAttributesOk() (*[]SAMLAttribute, bool) {
	if o == nil || o.BoundAttributes == nil {
		return nil, false
	}
	return o.BoundAttributes, true
}

// HasBoundAttributes returns a boolean if a field has been set.
func (o *SAMLAccessRules) HasBoundAttributes() bool {
	if o != nil && o.BoundAttributes != nil {
		return true
	}

	return false
}

// SetBoundAttributes gets a reference to the given []SAMLAttribute and assigns it to the BoundAttributes field.
func (o *SAMLAccessRules) SetBoundAttributes(v []SAMLAttribute) {
	o.BoundAttributes = &v
}

// GetIdpMetadataUrl returns the IdpMetadataUrl field value if set, zero value otherwise.
func (o *SAMLAccessRules) GetIdpMetadataUrl() string {
	if o == nil || o.IdpMetadataUrl == nil {
		var ret string
		return ret
	}
	return *o.IdpMetadataUrl
}

// GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAccessRules) GetIdpMetadataUrlOk() (*string, bool) {
	if o == nil || o.IdpMetadataUrl == nil {
		return nil, false
	}
	return o.IdpMetadataUrl, true
}

// HasIdpMetadataUrl returns a boolean if a field has been set.
func (o *SAMLAccessRules) HasIdpMetadataUrl() bool {
	if o != nil && o.IdpMetadataUrl != nil {
		return true
	}

	return false
}

// SetIdpMetadataUrl gets a reference to the given string and assigns it to the IdpMetadataUrl field.
func (o *SAMLAccessRules) SetIdpMetadataUrl(v string) {
	o.IdpMetadataUrl = &v
}

// GetIdpMetadataXml returns the IdpMetadataXml field value if set, zero value otherwise.
func (o *SAMLAccessRules) GetIdpMetadataXml() string {
	if o == nil || o.IdpMetadataXml == nil {
		var ret string
		return ret
	}
	return *o.IdpMetadataXml
}

// GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAccessRules) GetIdpMetadataXmlOk() (*string, bool) {
	if o == nil || o.IdpMetadataXml == nil {
		return nil, false
	}
	return o.IdpMetadataXml, true
}

// HasIdpMetadataXml returns a boolean if a field has been set.
func (o *SAMLAccessRules) HasIdpMetadataXml() bool {
	if o != nil && o.IdpMetadataXml != nil {
		return true
	}

	return false
}

// SetIdpMetadataXml gets a reference to the given string and assigns it to the IdpMetadataXml field.
func (o *SAMLAccessRules) SetIdpMetadataXml(v string) {
	o.IdpMetadataXml = &v
}

func (o SAMLAccessRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BoundAttributes != nil {
		toSerialize["bound_attributes"] = o.BoundAttributes
	}
	if o.IdpMetadataUrl != nil {
		toSerialize["idp_metadata_url"] = o.IdpMetadataUrl
	}
	if o.IdpMetadataXml != nil {
		toSerialize["idp_metadata_xml"] = o.IdpMetadataXml
	}
	return json.Marshal(toSerialize)
}

type NullableSAMLAccessRules struct {
	value *SAMLAccessRules
	isSet bool
}

func (v NullableSAMLAccessRules) Get() *SAMLAccessRules {
	return v.value
}

func (v *NullableSAMLAccessRules) Set(val *SAMLAccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLAccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLAccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLAccessRules(val *SAMLAccessRules) *NullableSAMLAccessRules {
	return &NullableSAMLAccessRules{value: val, isSet: true}
}

func (v NullableSAMLAccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLAccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


