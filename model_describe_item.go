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

// DescribeItem struct for DescribeItem
type DescribeItem struct {
	// for personal password manager
	Accessibility *string `json:"accessibility,omitempty"`
	// Indicate if the item should return with ztb cluster details (url, etc)
	BastionDetails *bool `json:"bastion-details,omitempty"`
	// The display id of the item
	DisplayId *string `json:"display-id,omitempty"`
	// Indicate if the item should return with clusters details (url, etc)
	GatewayDetails *bool `json:"gateway-details,omitempty"`
	// Item id of the item
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Item name
	Name string `json:"name"`
	// Include all associated services details
	ServicesDetails *bool `json:"services-details,omitempty"`
	// Include all item versions in reply
	ShowVersions *bool `json:"show-versions,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDescribeItem instantiates a new DescribeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeItem(name string, ) *DescribeItem {
	this := DescribeItem{}
	var accessibility string = "regular"
	this.Accessibility = &accessibility
	var bastionDetails bool = false
	this.BastionDetails = &bastionDetails
	var gatewayDetails bool = false
	this.GatewayDetails = &gatewayDetails
	var json bool = false
	this.Json = &json
	this.Name = name
	var servicesDetails bool = false
	this.ServicesDetails = &servicesDetails
	var showVersions bool = false
	this.ShowVersions = &showVersions
	return &this
}

// NewDescribeItemWithDefaults instantiates a new DescribeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeItemWithDefaults() *DescribeItem {
	this := DescribeItem{}
	var accessibility string = "regular"
	this.Accessibility = &accessibility
	var bastionDetails bool = false
	this.BastionDetails = &bastionDetails
	var gatewayDetails bool = false
	this.GatewayDetails = &gatewayDetails
	var json bool = false
	this.Json = &json
	var servicesDetails bool = false
	this.ServicesDetails = &servicesDetails
	var showVersions bool = false
	this.ShowVersions = &showVersions
	return &this
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *DescribeItem) GetAccessibility() string {
	if o == nil || o.Accessibility == nil {
		var ret string
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetAccessibilityOk() (*string, bool) {
	if o == nil || o.Accessibility == nil {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *DescribeItem) HasAccessibility() bool {
	if o != nil && o.Accessibility != nil {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given string and assigns it to the Accessibility field.
func (o *DescribeItem) SetAccessibility(v string) {
	o.Accessibility = &v
}

// GetBastionDetails returns the BastionDetails field value if set, zero value otherwise.
func (o *DescribeItem) GetBastionDetails() bool {
	if o == nil || o.BastionDetails == nil {
		var ret bool
		return ret
	}
	return *o.BastionDetails
}

// GetBastionDetailsOk returns a tuple with the BastionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetBastionDetailsOk() (*bool, bool) {
	if o == nil || o.BastionDetails == nil {
		return nil, false
	}
	return o.BastionDetails, true
}

// HasBastionDetails returns a boolean if a field has been set.
func (o *DescribeItem) HasBastionDetails() bool {
	if o != nil && o.BastionDetails != nil {
		return true
	}

	return false
}

// SetBastionDetails gets a reference to the given bool and assigns it to the BastionDetails field.
func (o *DescribeItem) SetBastionDetails(v bool) {
	o.BastionDetails = &v
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *DescribeItem) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *DescribeItem) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *DescribeItem) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetGatewayDetails returns the GatewayDetails field value if set, zero value otherwise.
func (o *DescribeItem) GetGatewayDetails() bool {
	if o == nil || o.GatewayDetails == nil {
		var ret bool
		return ret
	}
	return *o.GatewayDetails
}

// GetGatewayDetailsOk returns a tuple with the GatewayDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetGatewayDetailsOk() (*bool, bool) {
	if o == nil || o.GatewayDetails == nil {
		return nil, false
	}
	return o.GatewayDetails, true
}

// HasGatewayDetails returns a boolean if a field has been set.
func (o *DescribeItem) HasGatewayDetails() bool {
	if o != nil && o.GatewayDetails != nil {
		return true
	}

	return false
}

// SetGatewayDetails gets a reference to the given bool and assigns it to the GatewayDetails field.
func (o *DescribeItem) SetGatewayDetails(v bool) {
	o.GatewayDetails = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *DescribeItem) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *DescribeItem) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *DescribeItem) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DescribeItem) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DescribeItem) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DescribeItem) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DescribeItem) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeItem) SetName(v string) {
	o.Name = v
}

// GetServicesDetails returns the ServicesDetails field value if set, zero value otherwise.
func (o *DescribeItem) GetServicesDetails() bool {
	if o == nil || o.ServicesDetails == nil {
		var ret bool
		return ret
	}
	return *o.ServicesDetails
}

// GetServicesDetailsOk returns a tuple with the ServicesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetServicesDetailsOk() (*bool, bool) {
	if o == nil || o.ServicesDetails == nil {
		return nil, false
	}
	return o.ServicesDetails, true
}

// HasServicesDetails returns a boolean if a field has been set.
func (o *DescribeItem) HasServicesDetails() bool {
	if o != nil && o.ServicesDetails != nil {
		return true
	}

	return false
}

// SetServicesDetails gets a reference to the given bool and assigns it to the ServicesDetails field.
func (o *DescribeItem) SetServicesDetails(v bool) {
	o.ServicesDetails = &v
}

// GetShowVersions returns the ShowVersions field value if set, zero value otherwise.
func (o *DescribeItem) GetShowVersions() bool {
	if o == nil || o.ShowVersions == nil {
		var ret bool
		return ret
	}
	return *o.ShowVersions
}

// GetShowVersionsOk returns a tuple with the ShowVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetShowVersionsOk() (*bool, bool) {
	if o == nil || o.ShowVersions == nil {
		return nil, false
	}
	return o.ShowVersions, true
}

// HasShowVersions returns a boolean if a field has been set.
func (o *DescribeItem) HasShowVersions() bool {
	if o != nil && o.ShowVersions != nil {
		return true
	}

	return false
}

// SetShowVersions gets a reference to the given bool and assigns it to the ShowVersions field.
func (o *DescribeItem) SetShowVersions(v bool) {
	o.ShowVersions = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DescribeItem) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DescribeItem) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DescribeItem) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DescribeItem) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeItem) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DescribeItem) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DescribeItem) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DescribeItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessibility != nil {
		toSerialize["accessibility"] = o.Accessibility
	}
	if o.BastionDetails != nil {
		toSerialize["bastion-details"] = o.BastionDetails
	}
	if o.DisplayId != nil {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.GatewayDetails != nil {
		toSerialize["gateway-details"] = o.GatewayDetails
	}
	if o.ItemId != nil {
		toSerialize["item-id"] = o.ItemId
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ServicesDetails != nil {
		toSerialize["services-details"] = o.ServicesDetails
	}
	if o.ShowVersions != nil {
		toSerialize["show-versions"] = o.ShowVersions
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDescribeItem struct {
	value *DescribeItem
	isSet bool
}

func (v NullableDescribeItem) Get() *DescribeItem {
	return v.value
}

func (v *NullableDescribeItem) Set(val *DescribeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeItem(val *DescribeItem) *NullableDescribeItem {
	return &NullableDescribeItem{value: val, isSet: true}
}

func (v NullableDescribeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


