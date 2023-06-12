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

// checks if the KmipDeleteClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KmipDeleteClient{}

// KmipDeleteClient struct for KmipDeleteClient
type KmipDeleteClient struct {
	ClientId *string `json:"client-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	Name *string `json:"name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewKmipDeleteClient instantiates a new KmipDeleteClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipDeleteClient() *KmipDeleteClient {
	this := KmipDeleteClient{}
	var json bool = false
	this.Json = &json
	return &this
}

// NewKmipDeleteClientWithDefaults instantiates a new KmipDeleteClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipDeleteClientWithDefaults() *KmipDeleteClient {
	this := KmipDeleteClient{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *KmipDeleteClient) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDeleteClient) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *KmipDeleteClient) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *KmipDeleteClient) SetClientId(v string) {
	o.ClientId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *KmipDeleteClient) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDeleteClient) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *KmipDeleteClient) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *KmipDeleteClient) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KmipDeleteClient) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDeleteClient) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KmipDeleteClient) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KmipDeleteClient) SetName(v string) {
	o.Name = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KmipDeleteClient) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDeleteClient) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KmipDeleteClient) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KmipDeleteClient) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *KmipDeleteClient) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDeleteClient) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *KmipDeleteClient) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *KmipDeleteClient) SetUidToken(v string) {
	o.UidToken = &v
}

func (o KmipDeleteClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KmipDeleteClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client-id"] = o.ClientId
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableKmipDeleteClient struct {
	value *KmipDeleteClient
	isSet bool
}

func (v NullableKmipDeleteClient) Get() *KmipDeleteClient {
	return v.value
}

func (v *NullableKmipDeleteClient) Set(val *KmipDeleteClient) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipDeleteClient) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipDeleteClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipDeleteClient(val *KmipDeleteClient) *NullableKmipDeleteClient {
	return &NullableKmipDeleteClient{value: val, isSet: true}
}

func (v NullableKmipDeleteClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipDeleteClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


