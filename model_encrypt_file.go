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

// checks if the EncryptFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptFile{}

// EncryptFile struct for EncryptFile
type EncryptFile struct {
	// The display id of the key to use in the encryption process
	DisplayId *string `json:"display-id,omitempty"`
	// Path to the file to be encrypted. If not provided, the content will be taken from stdin
	In string `json:"in"`
	// The item id of the key to use in the encryption process
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of the key to use in the encryption process
	KeyName string `json:"key-name"`
	// Path to the output file. If not provided, the output will be sent to stdout
	Out *string `json:"out,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewEncryptFile instantiates a new EncryptFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptFile(in string, keyName string) *EncryptFile {
	this := EncryptFile{}
	this.In = in
	var json bool = false
	this.Json = &json
	this.KeyName = keyName
	return &this
}

// NewEncryptFileWithDefaults instantiates a new EncryptFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptFileWithDefaults() *EncryptFile {
	this := EncryptFile{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *EncryptFile) GetDisplayId() string {
	if o == nil || IsNil(o.DisplayId) {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetDisplayIdOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayId) {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *EncryptFile) HasDisplayId() bool {
	if o != nil && !IsNil(o.DisplayId) {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *EncryptFile) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetIn returns the In field value
func (o *EncryptFile) GetIn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.In
}

// GetInOk returns a tuple with the In field value
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetInOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.In, true
}

// SetIn sets field value
func (o *EncryptFile) SetIn(v string) {
	o.In = v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *EncryptFile) GetItemId() int64 {
	if o == nil || IsNil(o.ItemId) {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetItemIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *EncryptFile) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *EncryptFile) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *EncryptFile) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *EncryptFile) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *EncryptFile) SetJson(v bool) {
	o.Json = &v
}

// GetKeyName returns the KeyName field value
func (o *EncryptFile) GetKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *EncryptFile) SetKeyName(v string) {
	o.KeyName = v
}

// GetOut returns the Out field value if set, zero value otherwise.
func (o *EncryptFile) GetOut() string {
	if o == nil || IsNil(o.Out) {
		var ret string
		return ret
	}
	return *o.Out
}

// GetOutOk returns a tuple with the Out field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetOutOk() (*string, bool) {
	if o == nil || IsNil(o.Out) {
		return nil, false
	}
	return o.Out, true
}

// HasOut returns a boolean if a field has been set.
func (o *EncryptFile) HasOut() bool {
	if o != nil && !IsNil(o.Out) {
		return true
	}

	return false
}

// SetOut gets a reference to the given string and assigns it to the Out field.
func (o *EncryptFile) SetOut(v string) {
	o.Out = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EncryptFile) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EncryptFile) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EncryptFile) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EncryptFile) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EncryptFile) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EncryptFile) SetUidToken(v string) {
	o.UidToken = &v
}

func (o EncryptFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayId) {
		toSerialize["display-id"] = o.DisplayId
	}
	toSerialize["in"] = o.In
	if !IsNil(o.ItemId) {
		toSerialize["item-id"] = o.ItemId
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["key-name"] = o.KeyName
	if !IsNil(o.Out) {
		toSerialize["out"] = o.Out
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableEncryptFile struct {
	value *EncryptFile
	isSet bool
}

func (v NullableEncryptFile) Get() *EncryptFile {
	return v.value
}

func (v *NullableEncryptFile) Set(val *EncryptFile) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptFile) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptFile(val *EncryptFile) *NullableEncryptFile {
	return &NullableEncryptFile{value: val, isSet: true}
}

func (v NullableEncryptFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


