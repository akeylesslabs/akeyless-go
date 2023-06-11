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

// CreateTokenizer createTokenizer is a command that creates a tokenizer item
type CreateTokenizer struct {
	// Alphabet to use in regexp vaultless tokenization
	Alphabet *string `json:"alphabet,omitempty"`
	// The Decoding output template to use in regexp vaultless tokenization
	DecodingTemplate *string `json:"decoding-template,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// The Encoding output template to use in regexp vaultless tokenization
	EncodingTemplate *string `json:"encoding-template,omitempty"`
	// AES key name to use in vaultless tokenization
	EncryptionKeyName *string `json:"encryption-key-name,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Deprecated - use description
	Metadata *string `json:"metadata,omitempty"`
	// Tokenizer name
	Name string `json:"name"`
	// Pattern to use in regexp vaultless tokenization
	Pattern *string `json:"pattern,omitempty"`
	// List of the tags attached to this key
	Tag []string `json:"tag,omitempty"`
	// Which template type this tokenizer is used for [SSN,CreditCard,USPhoneNumber,Email,Regexp]
	TemplateType string `json:"template-type"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Tokenizer type
	TokenizerType string `json:"tokenizer-type"`
	// The tweak type to use in vaultless tokenization [Supplied, Generated, Internal, Masking]
	TweakType *string `json:"tweak-type,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateTokenizer instantiates a new CreateTokenizer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokenizer(name string, templateType string, tokenizerType string) *CreateTokenizer {
	this := CreateTokenizer{}
	var json bool = false
	this.Json = &json
	this.Name = name
	this.TemplateType = templateType
	this.TokenizerType = tokenizerType
	return &this
}

// NewCreateTokenizerWithDefaults instantiates a new CreateTokenizer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokenizerWithDefaults() *CreateTokenizer {
	this := CreateTokenizer{}
	var json bool = false
	this.Json = &json
	var tokenizerType string = "vaultless"
	this.TokenizerType = tokenizerType
	return &this
}

// GetAlphabet returns the Alphabet field value if set, zero value otherwise.
func (o *CreateTokenizer) GetAlphabet() string {
	if o == nil || o.Alphabet == nil {
		var ret string
		return ret
	}
	return *o.Alphabet
}

// GetAlphabetOk returns a tuple with the Alphabet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetAlphabetOk() (*string, bool) {
	if o == nil || o.Alphabet == nil {
		return nil, false
	}
	return o.Alphabet, true
}

// HasAlphabet returns a boolean if a field has been set.
func (o *CreateTokenizer) HasAlphabet() bool {
	if o != nil && o.Alphabet != nil {
		return true
	}

	return false
}

// SetAlphabet gets a reference to the given string and assigns it to the Alphabet field.
func (o *CreateTokenizer) SetAlphabet(v string) {
	o.Alphabet = &v
}

// GetDecodingTemplate returns the DecodingTemplate field value if set, zero value otherwise.
func (o *CreateTokenizer) GetDecodingTemplate() string {
	if o == nil || o.DecodingTemplate == nil {
		var ret string
		return ret
	}
	return *o.DecodingTemplate
}

// GetDecodingTemplateOk returns a tuple with the DecodingTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetDecodingTemplateOk() (*string, bool) {
	if o == nil || o.DecodingTemplate == nil {
		return nil, false
	}
	return o.DecodingTemplate, true
}

// HasDecodingTemplate returns a boolean if a field has been set.
func (o *CreateTokenizer) HasDecodingTemplate() bool {
	if o != nil && o.DecodingTemplate != nil {
		return true
	}

	return false
}

// SetDecodingTemplate gets a reference to the given string and assigns it to the DecodingTemplate field.
func (o *CreateTokenizer) SetDecodingTemplate(v string) {
	o.DecodingTemplate = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *CreateTokenizer) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *CreateTokenizer) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *CreateTokenizer) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateTokenizer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateTokenizer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateTokenizer) SetDescription(v string) {
	o.Description = &v
}

// GetEncodingTemplate returns the EncodingTemplate field value if set, zero value otherwise.
func (o *CreateTokenizer) GetEncodingTemplate() string {
	if o == nil || o.EncodingTemplate == nil {
		var ret string
		return ret
	}
	return *o.EncodingTemplate
}

// GetEncodingTemplateOk returns a tuple with the EncodingTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetEncodingTemplateOk() (*string, bool) {
	if o == nil || o.EncodingTemplate == nil {
		return nil, false
	}
	return o.EncodingTemplate, true
}

// HasEncodingTemplate returns a boolean if a field has been set.
func (o *CreateTokenizer) HasEncodingTemplate() bool {
	if o != nil && o.EncodingTemplate != nil {
		return true
	}

	return false
}

// SetEncodingTemplate gets a reference to the given string and assigns it to the EncodingTemplate field.
func (o *CreateTokenizer) SetEncodingTemplate(v string) {
	o.EncodingTemplate = &v
}

// GetEncryptionKeyName returns the EncryptionKeyName field value if set, zero value otherwise.
func (o *CreateTokenizer) GetEncryptionKeyName() string {
	if o == nil || o.EncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKeyName
}

// GetEncryptionKeyNameOk returns a tuple with the EncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.EncryptionKeyName == nil {
		return nil, false
	}
	return o.EncryptionKeyName, true
}

// HasEncryptionKeyName returns a boolean if a field has been set.
func (o *CreateTokenizer) HasEncryptionKeyName() bool {
	if o != nil && o.EncryptionKeyName != nil {
		return true
	}

	return false
}

// SetEncryptionKeyName gets a reference to the given string and assigns it to the EncryptionKeyName field.
func (o *CreateTokenizer) SetEncryptionKeyName(v string) {
	o.EncryptionKeyName = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateTokenizer) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateTokenizer) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateTokenizer) SetJson(v bool) {
	o.Json = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateTokenizer) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateTokenizer) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreateTokenizer) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *CreateTokenizer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTokenizer) SetName(v string) {
	o.Name = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *CreateTokenizer) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *CreateTokenizer) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *CreateTokenizer) SetPattern(v string) {
	o.Pattern = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateTokenizer) GetTag() []string {
	if o == nil || o.Tag == nil {
		var ret []string
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetTagOk() ([]string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateTokenizer) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *CreateTokenizer) SetTag(v []string) {
	o.Tag = v
}

// GetTemplateType returns the TemplateType field value
func (o *CreateTokenizer) GetTemplateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetTemplateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateType, true
}

// SetTemplateType sets field value
func (o *CreateTokenizer) SetTemplateType(v string) {
	o.TemplateType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateTokenizer) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateTokenizer) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateTokenizer) SetToken(v string) {
	o.Token = &v
}

// GetTokenizerType returns the TokenizerType field value
func (o *CreateTokenizer) GetTokenizerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenizerType
}

// GetTokenizerTypeOk returns a tuple with the TokenizerType field value
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetTokenizerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenizerType, true
}

// SetTokenizerType sets field value
func (o *CreateTokenizer) SetTokenizerType(v string) {
	o.TokenizerType = v
}

// GetTweakType returns the TweakType field value if set, zero value otherwise.
func (o *CreateTokenizer) GetTweakType() string {
	if o == nil || o.TweakType == nil {
		var ret string
		return ret
	}
	return *o.TweakType
}

// GetTweakTypeOk returns a tuple with the TweakType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetTweakTypeOk() (*string, bool) {
	if o == nil || o.TweakType == nil {
		return nil, false
	}
	return o.TweakType, true
}

// HasTweakType returns a boolean if a field has been set.
func (o *CreateTokenizer) HasTweakType() bool {
	if o != nil && o.TweakType != nil {
		return true
	}

	return false
}

// SetTweakType gets a reference to the given string and assigns it to the TweakType field.
func (o *CreateTokenizer) SetTweakType(v string) {
	o.TweakType = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateTokenizer) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizer) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateTokenizer) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateTokenizer) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateTokenizer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alphabet != nil {
		toSerialize["alphabet"] = o.Alphabet
	}
	if o.DecodingTemplate != nil {
		toSerialize["decoding-template"] = o.DecodingTemplate
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EncodingTemplate != nil {
		toSerialize["encoding-template"] = o.EncodingTemplate
	}
	if o.EncryptionKeyName != nil {
		toSerialize["encryption-key-name"] = o.EncryptionKeyName
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["template-type"] = o.TemplateType
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["tokenizer-type"] = o.TokenizerType
	}
	if o.TweakType != nil {
		toSerialize["tweak-type"] = o.TweakType
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTokenizer struct {
	value *CreateTokenizer
	isSet bool
}

func (v NullableCreateTokenizer) Get() *CreateTokenizer {
	return v.value
}

func (v *NullableCreateTokenizer) Set(val *CreateTokenizer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokenizer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokenizer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokenizer(val *CreateTokenizer) *NullableCreateTokenizer {
	return &NullableCreateTokenizer{value: val, isSet: true}
}

func (v NullableCreateTokenizer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokenizer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


