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

// checks if the ShareItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShareItem{}

// ShareItem struct for ShareItem
type ShareItem struct {
	// for personal password manager
	Accessibility *string `json:"accessibility,omitempty"`
	// Action to be performed on the item [start/stop/describe]
	Action string `json:"action"`
	// For Password Management use, reflect the website context
	Emails []string `json:"emails,omitempty"`
	// Item name
	ItemName string `json:"item-name"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// TTL of the Availability of the shared secret in seconds
	Ttl *int32 `json:"ttl,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// ViewOnlyOnce Shared secrets can only be viewed once [true/false]
	ViewOnce *bool `json:"view-once,omitempty"`
}

// NewShareItem instantiates a new ShareItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareItem(action string, itemName string) *ShareItem {
	this := ShareItem{}
	var accessibility string = "regular"
	this.Accessibility = &accessibility
	this.Action = action
	this.ItemName = itemName
	var json bool = false
	this.Json = &json
	var viewOnce bool = false
	this.ViewOnce = &viewOnce
	return &this
}

// NewShareItemWithDefaults instantiates a new ShareItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareItemWithDefaults() *ShareItem {
	this := ShareItem{}
	var accessibility string = "regular"
	this.Accessibility = &accessibility
	var json bool = false
	this.Json = &json
	var viewOnce bool = false
	this.ViewOnce = &viewOnce
	return &this
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *ShareItem) GetAccessibility() string {
	if o == nil || IsNil(o.Accessibility) {
		var ret string
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareItem) GetAccessibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Accessibility) {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *ShareItem) HasAccessibility() bool {
	if o != nil && !IsNil(o.Accessibility) {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given string and assigns it to the Accessibility field.
func (o *ShareItem) SetAccessibility(v string) {
	o.Accessibility = &v
}

// GetAction returns the Action field value
func (o *ShareItem) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ShareItem) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ShareItem) SetAction(v string) {
	o.Action = v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *ShareItem) GetEmails() []string {
	if o == nil || IsNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareItem) GetEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *ShareItem) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *ShareItem) SetEmails(v []string) {
	o.Emails = v
}

// GetItemName returns the ItemName field value
func (o *ShareItem) GetItemName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemName
}

// GetItemNameOk returns a tuple with the ItemName field value
// and a boolean to check if the value has been set.
func (o *ShareItem) GetItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemName, true
}

// SetItemName sets field value
func (o *ShareItem) SetItemName(v string) {
	o.ItemName = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ShareItem) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareItem) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ShareItem) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *ShareItem) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ShareItem) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareItem) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ShareItem) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ShareItem) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ShareItem) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareItem) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ShareItem) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *ShareItem) SetTtl(v int32) {
	o.Ttl = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *ShareItem) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareItem) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *ShareItem) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *ShareItem) SetUidToken(v string) {
	o.UidToken = &v
}

// GetViewOnce returns the ViewOnce field value if set, zero value otherwise.
func (o *ShareItem) GetViewOnce() bool {
	if o == nil || IsNil(o.ViewOnce) {
		var ret bool
		return ret
	}
	return *o.ViewOnce
}

// GetViewOnceOk returns a tuple with the ViewOnce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareItem) GetViewOnceOk() (*bool, bool) {
	if o == nil || IsNil(o.ViewOnce) {
		return nil, false
	}
	return o.ViewOnce, true
}

// HasViewOnce returns a boolean if a field has been set.
func (o *ShareItem) HasViewOnce() bool {
	if o != nil && !IsNil(o.ViewOnce) {
		return true
	}

	return false
}

// SetViewOnce gets a reference to the given bool and assigns it to the ViewOnce field.
func (o *ShareItem) SetViewOnce(v bool) {
	o.ViewOnce = &v
}

func (o ShareItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShareItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accessibility) {
		toSerialize["accessibility"] = o.Accessibility
	}
	toSerialize["action"] = o.Action
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	toSerialize["item-name"] = o.ItemName
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	if !IsNil(o.ViewOnce) {
		toSerialize["view-once"] = o.ViewOnce
	}
	return toSerialize, nil
}

type NullableShareItem struct {
	value *ShareItem
	isSet bool
}

func (v NullableShareItem) Get() *ShareItem {
	return v.value
}

func (v *NullableShareItem) Set(val *ShareItem) {
	v.value = val
	v.isSet = true
}

func (v NullableShareItem) IsSet() bool {
	return v.isSet
}

func (v *NullableShareItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareItem(val *ShareItem) *NullableShareItem {
	return &NullableShareItem{value: val, isSet: true}
}

func (v NullableShareItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


