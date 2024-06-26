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

// AddGatewayAllowedAccessId Responses:  default: errorResponse 200: addGatewayAllowedAccessIdResponse
type AddGatewayAllowedAccessId struct {
	// The access id that will be able to access to gateway
	AccessId string `json:"access-id"`
	// The name of the updated cluster, e.g. acc-abcd12345678/p-123456789012/defaultCluster
	ClusterName string `json:"cluster-name"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// key/val of sub claims, e.g group=admins,developers
	SubClaims *map[string]string `json:"sub-claims,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewAddGatewayAllowedAccessId instantiates a new AddGatewayAllowedAccessId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGatewayAllowedAccessId(accessId string, clusterName string, ) *AddGatewayAllowedAccessId {
	this := AddGatewayAllowedAccessId{}
	this.AccessId = accessId
	this.ClusterName = clusterName
	var json bool = false
	this.Json = &json
	return &this
}

// NewAddGatewayAllowedAccessIdWithDefaults instantiates a new AddGatewayAllowedAccessId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGatewayAllowedAccessIdWithDefaults() *AddGatewayAllowedAccessId {
	this := AddGatewayAllowedAccessId{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAccessId returns the AccessId field value
func (o *AddGatewayAllowedAccessId) GetAccessId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value
// and a boolean to check if the value has been set.
func (o *AddGatewayAllowedAccessId) GetAccessIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessId, true
}

// SetAccessId sets field value
func (o *AddGatewayAllowedAccessId) SetAccessId(v string) {
	o.AccessId = v
}

// GetClusterName returns the ClusterName field value
func (o *AddGatewayAllowedAccessId) GetClusterName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *AddGatewayAllowedAccessId) GetClusterNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *AddGatewayAllowedAccessId) SetClusterName(v string) {
	o.ClusterName = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *AddGatewayAllowedAccessId) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGatewayAllowedAccessId) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *AddGatewayAllowedAccessId) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *AddGatewayAllowedAccessId) SetJson(v bool) {
	o.Json = &v
}

// GetSubClaims returns the SubClaims field value if set, zero value otherwise.
func (o *AddGatewayAllowedAccessId) GetSubClaims() map[string]string {
	if o == nil || o.SubClaims == nil {
		var ret map[string]string
		return ret
	}
	return *o.SubClaims
}

// GetSubClaimsOk returns a tuple with the SubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGatewayAllowedAccessId) GetSubClaimsOk() (*map[string]string, bool) {
	if o == nil || o.SubClaims == nil {
		return nil, false
	}
	return o.SubClaims, true
}

// HasSubClaims returns a boolean if a field has been set.
func (o *AddGatewayAllowedAccessId) HasSubClaims() bool {
	if o != nil && o.SubClaims != nil {
		return true
	}

	return false
}

// SetSubClaims gets a reference to the given map[string]string and assigns it to the SubClaims field.
func (o *AddGatewayAllowedAccessId) SetSubClaims(v map[string]string) {
	o.SubClaims = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AddGatewayAllowedAccessId) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGatewayAllowedAccessId) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AddGatewayAllowedAccessId) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AddGatewayAllowedAccessId) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *AddGatewayAllowedAccessId) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGatewayAllowedAccessId) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *AddGatewayAllowedAccessId) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *AddGatewayAllowedAccessId) SetUidToken(v string) {
	o.UidToken = &v
}

func (o AddGatewayAllowedAccessId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access-id"] = o.AccessId
	}
	if true {
		toSerialize["cluster-name"] = o.ClusterName
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.SubClaims != nil {
		toSerialize["sub-claims"] = o.SubClaims
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableAddGatewayAllowedAccessId struct {
	value *AddGatewayAllowedAccessId
	isSet bool
}

func (v NullableAddGatewayAllowedAccessId) Get() *AddGatewayAllowedAccessId {
	return v.value
}

func (v *NullableAddGatewayAllowedAccessId) Set(val *AddGatewayAllowedAccessId) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGatewayAllowedAccessId) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGatewayAllowedAccessId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGatewayAllowedAccessId(val *AddGatewayAllowedAccessId) *NullableAddGatewayAllowedAccessId {
	return &NullableAddGatewayAllowedAccessId{value: val, isSet: true}
}

func (v NullableAddGatewayAllowedAccessId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGatewayAllowedAccessId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


