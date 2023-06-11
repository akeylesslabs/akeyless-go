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

// checks if the TargetItemAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetItemAssociation{}

// TargetItemAssociation TargetItemAssociation includes details of an association between a target and an item.
type TargetItemAssociation struct {
	AssocId *string `json:"assoc_id,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty"`
	ClusterId *int64 `json:"cluster_id,omitempty"`
	ItemName *string `json:"item_name,omitempty"`
	ItemType *string `json:"item_type,omitempty"`
	Relationship *string `json:"relationship,omitempty"`
}

// NewTargetItemAssociation instantiates a new TargetItemAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetItemAssociation() *TargetItemAssociation {
	this := TargetItemAssociation{}
	return &this
}

// NewTargetItemAssociationWithDefaults instantiates a new TargetItemAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetItemAssociationWithDefaults() *TargetItemAssociation {
	this := TargetItemAssociation{}
	return &this
}

// GetAssocId returns the AssocId field value if set, zero value otherwise.
func (o *TargetItemAssociation) GetAssocId() string {
	if o == nil || IsNil(o.AssocId) {
		var ret string
		return ret
	}
	return *o.AssocId
}

// GetAssocIdOk returns a tuple with the AssocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemAssociation) GetAssocIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssocId) {
		return nil, false
	}
	return o.AssocId, true
}

// HasAssocId returns a boolean if a field has been set.
func (o *TargetItemAssociation) HasAssocId() bool {
	if o != nil && !IsNil(o.AssocId) {
		return true
	}

	return false
}

// SetAssocId gets a reference to the given string and assigns it to the AssocId field.
func (o *TargetItemAssociation) SetAssocId(v string) {
	o.AssocId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TargetItemAssociation) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemAssociation) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TargetItemAssociation) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *TargetItemAssociation) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *TargetItemAssociation) GetClusterId() int64 {
	if o == nil || IsNil(o.ClusterId) {
		var ret int64
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemAssociation) GetClusterIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *TargetItemAssociation) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given int64 and assigns it to the ClusterId field.
func (o *TargetItemAssociation) SetClusterId(v int64) {
	o.ClusterId = &v
}

// GetItemName returns the ItemName field value if set, zero value otherwise.
func (o *TargetItemAssociation) GetItemName() string {
	if o == nil || IsNil(o.ItemName) {
		var ret string
		return ret
	}
	return *o.ItemName
}

// GetItemNameOk returns a tuple with the ItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemAssociation) GetItemNameOk() (*string, bool) {
	if o == nil || IsNil(o.ItemName) {
		return nil, false
	}
	return o.ItemName, true
}

// HasItemName returns a boolean if a field has been set.
func (o *TargetItemAssociation) HasItemName() bool {
	if o != nil && !IsNil(o.ItemName) {
		return true
	}

	return false
}

// SetItemName gets a reference to the given string and assigns it to the ItemName field.
func (o *TargetItemAssociation) SetItemName(v string) {
	o.ItemName = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *TargetItemAssociation) GetItemType() string {
	if o == nil || IsNil(o.ItemType) {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemAssociation) GetItemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemType) {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *TargetItemAssociation) HasItemType() bool {
	if o != nil && !IsNil(o.ItemType) {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *TargetItemAssociation) SetItemType(v string) {
	o.ItemType = &v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *TargetItemAssociation) GetRelationship() string {
	if o == nil || IsNil(o.Relationship) {
		var ret string
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetItemAssociation) GetRelationshipOk() (*string, bool) {
	if o == nil || IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *TargetItemAssociation) HasRelationship() bool {
	if o != nil && !IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given string and assigns it to the Relationship field.
func (o *TargetItemAssociation) SetRelationship(v string) {
	o.Relationship = &v
}

func (o TargetItemAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetItemAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssocId) {
		toSerialize["assoc_id"] = o.AssocId
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.ClusterId) {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if !IsNil(o.ItemName) {
		toSerialize["item_name"] = o.ItemName
	}
	if !IsNil(o.ItemType) {
		toSerialize["item_type"] = o.ItemType
	}
	if !IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	return toSerialize, nil
}

type NullableTargetItemAssociation struct {
	value *TargetItemAssociation
	isSet bool
}

func (v NullableTargetItemAssociation) Get() *TargetItemAssociation {
	return v.value
}

func (v *NullableTargetItemAssociation) Set(val *TargetItemAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetItemAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetItemAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetItemAssociation(val *TargetItemAssociation) *NullableTargetItemAssociation {
	return &NullableTargetItemAssociation{value: val, isSet: true}
}

func (v NullableTargetItemAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetItemAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


