# AccountCustomFieldCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | The name of the custom field | 
**Object** | **string** | The object to create the custom field | [default to "ITEM"]
**ObjectType** | **string** | The object type to create the custom field [e.g. STATIC_SECRET,DYNAMIC_SECRET,ROTATED_SECRET] | 
**Required** | Pointer to **bool** | Specify whether the custom field is mandatory | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewAccountCustomFieldCreate

`func NewAccountCustomFieldCreate(name string, object string, objectType string, ) *AccountCustomFieldCreate`

NewAccountCustomFieldCreate instantiates a new AccountCustomFieldCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCustomFieldCreateWithDefaults

`func NewAccountCustomFieldCreateWithDefaults() *AccountCustomFieldCreate`

NewAccountCustomFieldCreateWithDefaults instantiates a new AccountCustomFieldCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *AccountCustomFieldCreate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AccountCustomFieldCreate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AccountCustomFieldCreate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AccountCustomFieldCreate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *AccountCustomFieldCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCustomFieldCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCustomFieldCreate) SetName(v string)`

SetName sets Name field to given value.


### GetObject

`func (o *AccountCustomFieldCreate) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AccountCustomFieldCreate) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AccountCustomFieldCreate) SetObject(v string)`

SetObject sets Object field to given value.


### GetObjectType

`func (o *AccountCustomFieldCreate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccountCustomFieldCreate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccountCustomFieldCreate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRequired

`func (o *AccountCustomFieldCreate) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AccountCustomFieldCreate) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AccountCustomFieldCreate) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AccountCustomFieldCreate) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetToken

`func (o *AccountCustomFieldCreate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountCustomFieldCreate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountCustomFieldCreate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountCustomFieldCreate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AccountCustomFieldCreate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AccountCustomFieldCreate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AccountCustomFieldCreate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AccountCustomFieldCreate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


