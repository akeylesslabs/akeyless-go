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

// UpdateTokenizer updateTokenizer is a command that updates a tokenizer item
type UpdateTokenizer struct {
	// List of the new tags that will be attached to this item
	AddTag *[]string `json:"add-tag,omitempty"`
	// Alphabet to use in regexp vaultless tokenization
	Alphabet *string `json:"alphabet,omitempty"`
	// The Decryption output template to use in regexp vaultless tokenization
	DecryptionTemplate *string `json:"decryption-template,omitempty"`
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// AES key name to use in vaultless tokenization
	EncryptionKeyName *string `json:"encryption-key-name,omitempty"`
	// The Encryption output template to use in regexp vaultless tokenization
	EncryptionTemplate *string `json:"encryption-template,omitempty"`
	// Current item name
	Name string `json:"name"`
	// New item metadata
	NewMetadata *string `json:"new-metadata,omitempty"`
	// New item name
	NewName *string `json:"new-name,omitempty"`
	// Pattern to use in regexp vaultless tokenization
	Pattern *string `json:"pattern,omitempty"`
	// List of the existent tags that will be removed from this item
	RmTag *[]string `json:"rm-tag,omitempty"`
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

// NewUpdateTokenizer instantiates a new UpdateTokenizer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTokenizer(name string, templateType string, tokenizerType string, ) *UpdateTokenizer {
	this := UpdateTokenizer{}
	this.Name = name
	var newMetadata string = "default_metadata"
	this.NewMetadata = &newMetadata
	this.TemplateType = templateType
	this.TokenizerType = tokenizerType
	return &this
}

// NewUpdateTokenizerWithDefaults instantiates a new UpdateTokenizer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTokenizerWithDefaults() *UpdateTokenizer {
	this := UpdateTokenizer{}
	var newMetadata string = "default_metadata"
	this.NewMetadata = &newMetadata
	return &this
}

// GetAddTag returns the AddTag field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetAddTag() []string {
	if o == nil || o.AddTag == nil {
		var ret []string
		return ret
	}
	return *o.AddTag
}

// GetAddTagOk returns a tuple with the AddTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetAddTagOk() (*[]string, bool) {
	if o == nil || o.AddTag == nil {
		return nil, false
	}
	return o.AddTag, true
}

// HasAddTag returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasAddTag() bool {
	if o != nil && o.AddTag != nil {
		return true
	}

	return false
}

// SetAddTag gets a reference to the given []string and assigns it to the AddTag field.
func (o *UpdateTokenizer) SetAddTag(v []string) {
	o.AddTag = &v
}

// GetAlphabet returns the Alphabet field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetAlphabet() string {
	if o == nil || o.Alphabet == nil {
		var ret string
		return ret
	}
	return *o.Alphabet
}

// GetAlphabetOk returns a tuple with the Alphabet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetAlphabetOk() (*string, bool) {
	if o == nil || o.Alphabet == nil {
		return nil, false
	}
	return o.Alphabet, true
}

// HasAlphabet returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasAlphabet() bool {
	if o != nil && o.Alphabet != nil {
		return true
	}

	return false
}

// SetAlphabet gets a reference to the given string and assigns it to the Alphabet field.
func (o *UpdateTokenizer) SetAlphabet(v string) {
	o.Alphabet = &v
}

// GetDecryptionTemplate returns the DecryptionTemplate field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetDecryptionTemplate() string {
	if o == nil || o.DecryptionTemplate == nil {
		var ret string
		return ret
	}
	return *o.DecryptionTemplate
}

// GetDecryptionTemplateOk returns a tuple with the DecryptionTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetDecryptionTemplateOk() (*string, bool) {
	if o == nil || o.DecryptionTemplate == nil {
		return nil, false
	}
	return o.DecryptionTemplate, true
}

// HasDecryptionTemplate returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasDecryptionTemplate() bool {
	if o != nil && o.DecryptionTemplate != nil {
		return true
	}

	return false
}

// SetDecryptionTemplate gets a reference to the given string and assigns it to the DecryptionTemplate field.
func (o *UpdateTokenizer) SetDecryptionTemplate(v string) {
	o.DecryptionTemplate = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *UpdateTokenizer) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetEncryptionKeyName returns the EncryptionKeyName field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetEncryptionKeyName() string {
	if o == nil || o.EncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKeyName
}

// GetEncryptionKeyNameOk returns a tuple with the EncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.EncryptionKeyName == nil {
		return nil, false
	}
	return o.EncryptionKeyName, true
}

// HasEncryptionKeyName returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasEncryptionKeyName() bool {
	if o != nil && o.EncryptionKeyName != nil {
		return true
	}

	return false
}

// SetEncryptionKeyName gets a reference to the given string and assigns it to the EncryptionKeyName field.
func (o *UpdateTokenizer) SetEncryptionKeyName(v string) {
	o.EncryptionKeyName = &v
}

// GetEncryptionTemplate returns the EncryptionTemplate field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetEncryptionTemplate() string {
	if o == nil || o.EncryptionTemplate == nil {
		var ret string
		return ret
	}
	return *o.EncryptionTemplate
}

// GetEncryptionTemplateOk returns a tuple with the EncryptionTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetEncryptionTemplateOk() (*string, bool) {
	if o == nil || o.EncryptionTemplate == nil {
		return nil, false
	}
	return o.EncryptionTemplate, true
}

// HasEncryptionTemplate returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasEncryptionTemplate() bool {
	if o != nil && o.EncryptionTemplate != nil {
		return true
	}

	return false
}

// SetEncryptionTemplate gets a reference to the given string and assigns it to the EncryptionTemplate field.
func (o *UpdateTokenizer) SetEncryptionTemplate(v string) {
	o.EncryptionTemplate = &v
}

// GetName returns the Name field value
func (o *UpdateTokenizer) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateTokenizer) SetName(v string) {
	o.Name = v
}

// GetNewMetadata returns the NewMetadata field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetNewMetadata() string {
	if o == nil || o.NewMetadata == nil {
		var ret string
		return ret
	}
	return *o.NewMetadata
}

// GetNewMetadataOk returns a tuple with the NewMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetNewMetadataOk() (*string, bool) {
	if o == nil || o.NewMetadata == nil {
		return nil, false
	}
	return o.NewMetadata, true
}

// HasNewMetadata returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasNewMetadata() bool {
	if o != nil && o.NewMetadata != nil {
		return true
	}

	return false
}

// SetNewMetadata gets a reference to the given string and assigns it to the NewMetadata field.
func (o *UpdateTokenizer) SetNewMetadata(v string) {
	o.NewMetadata = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateTokenizer) SetNewName(v string) {
	o.NewName = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *UpdateTokenizer) SetPattern(v string) {
	o.Pattern = &v
}

// GetRmTag returns the RmTag field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetRmTag() []string {
	if o == nil || o.RmTag == nil {
		var ret []string
		return ret
	}
	return *o.RmTag
}

// GetRmTagOk returns a tuple with the RmTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetRmTagOk() (*[]string, bool) {
	if o == nil || o.RmTag == nil {
		return nil, false
	}
	return o.RmTag, true
}

// HasRmTag returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasRmTag() bool {
	if o != nil && o.RmTag != nil {
		return true
	}

	return false
}

// SetRmTag gets a reference to the given []string and assigns it to the RmTag field.
func (o *UpdateTokenizer) SetRmTag(v []string) {
	o.RmTag = &v
}

// GetTemplateType returns the TemplateType field value
func (o *UpdateTokenizer) GetTemplateType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetTemplateTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateType, true
}

// SetTemplateType sets field value
func (o *UpdateTokenizer) SetTemplateType(v string) {
	o.TemplateType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateTokenizer) SetToken(v string) {
	o.Token = &v
}

// GetTokenizerType returns the TokenizerType field value
func (o *UpdateTokenizer) GetTokenizerType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TokenizerType
}

// GetTokenizerTypeOk returns a tuple with the TokenizerType field value
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetTokenizerTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenizerType, true
}

// SetTokenizerType sets field value
func (o *UpdateTokenizer) SetTokenizerType(v string) {
	o.TokenizerType = v
}

// GetTweakType returns the TweakType field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetTweakType() string {
	if o == nil || o.TweakType == nil {
		var ret string
		return ret
	}
	return *o.TweakType
}

// GetTweakTypeOk returns a tuple with the TweakType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetTweakTypeOk() (*string, bool) {
	if o == nil || o.TweakType == nil {
		return nil, false
	}
	return o.TweakType, true
}

// HasTweakType returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasTweakType() bool {
	if o != nil && o.TweakType != nil {
		return true
	}

	return false
}

// SetTweakType gets a reference to the given string and assigns it to the TweakType field.
func (o *UpdateTokenizer) SetTweakType(v string) {
	o.TweakType = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateTokenizer) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTokenizer) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateTokenizer) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateTokenizer) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateTokenizer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddTag != nil {
		toSerialize["add-tag"] = o.AddTag
	}
	if o.Alphabet != nil {
		toSerialize["alphabet"] = o.Alphabet
	}
	if o.DecryptionTemplate != nil {
		toSerialize["decryption-template"] = o.DecryptionTemplate
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.EncryptionKeyName != nil {
		toSerialize["encryption-key-name"] = o.EncryptionKeyName
	}
	if o.EncryptionTemplate != nil {
		toSerialize["encryption-template"] = o.EncryptionTemplate
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewMetadata != nil {
		toSerialize["new-metadata"] = o.NewMetadata
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.RmTag != nil {
		toSerialize["rm-tag"] = o.RmTag
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

type NullableUpdateTokenizer struct {
	value *UpdateTokenizer
	isSet bool
}

func (v NullableUpdateTokenizer) Get() *UpdateTokenizer {
	return v.value
}

func (v *NullableUpdateTokenizer) Set(val *UpdateTokenizer) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTokenizer) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTokenizer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTokenizer(val *UpdateTokenizer) *NullableUpdateTokenizer {
	return &NullableUpdateTokenizer{value: val, isSet: true}
}

func (v NullableUpdateTokenizer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTokenizer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


