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

// ListItemsOutput struct for ListItemsOutput
type ListItemsOutput struct {
	Items []Item `json:"items,omitempty"`
	NextPage *string `json:"next_page,omitempty"`
}

// NewListItemsOutput instantiates a new ListItemsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListItemsOutput() *ListItemsOutput {
	this := ListItemsOutput{}
	return &this
}

// NewListItemsOutputWithDefaults instantiates a new ListItemsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListItemsOutputWithDefaults() *ListItemsOutput {
	this := ListItemsOutput{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListItemsOutput) GetItems() []Item {
	if o == nil || o.Items == nil {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItemsOutput) GetItemsOk() ([]Item, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListItemsOutput) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *ListItemsOutput) SetItems(v []Item) {
	o.Items = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *ListItemsOutput) GetNextPage() string {
	if o == nil || o.NextPage == nil {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItemsOutput) GetNextPageOk() (*string, bool) {
	if o == nil || o.NextPage == nil {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *ListItemsOutput) HasNextPage() bool {
	if o != nil && o.NextPage != nil {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *ListItemsOutput) SetNextPage(v string) {
	o.NextPage = &v
}

func (o ListItemsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextPage != nil {
		toSerialize["next_page"] = o.NextPage
	}
	return json.Marshal(toSerialize)
}

type NullableListItemsOutput struct {
	value *ListItemsOutput
	isSet bool
}

func (v NullableListItemsOutput) Get() *ListItemsOutput {
	return v.value
}

func (v *NullableListItemsOutput) Set(val *ListItemsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListItemsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListItemsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListItemsOutput(val *ListItemsOutput) *NullableListItemsOutput {
	return &NullableListItemsOutput{value: val, isSet: true}
}

func (v NullableListItemsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListItemsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


