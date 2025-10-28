# AccountCustomFieldUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The custom field id | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | Pointer to **string** | The new name of the custom field | [optional] 
**Required** | Pointer to **bool** | Specify whether the custom field is mandatory | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewAccountCustomFieldUpdate

`func NewAccountCustomFieldUpdate(id int64, ) *AccountCustomFieldUpdate`

NewAccountCustomFieldUpdate instantiates a new AccountCustomFieldUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCustomFieldUpdateWithDefaults

`func NewAccountCustomFieldUpdateWithDefaults() *AccountCustomFieldUpdate`

NewAccountCustomFieldUpdateWithDefaults instantiates a new AccountCustomFieldUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountCustomFieldUpdate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountCustomFieldUpdate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountCustomFieldUpdate) SetId(v int64)`

SetId sets Id field to given value.


### GetJson

`func (o *AccountCustomFieldUpdate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AccountCustomFieldUpdate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AccountCustomFieldUpdate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AccountCustomFieldUpdate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *AccountCustomFieldUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCustomFieldUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCustomFieldUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountCustomFieldUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *AccountCustomFieldUpdate) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AccountCustomFieldUpdate) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AccountCustomFieldUpdate) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AccountCustomFieldUpdate) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetToken

`func (o *AccountCustomFieldUpdate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountCustomFieldUpdate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountCustomFieldUpdate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountCustomFieldUpdate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AccountCustomFieldUpdate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AccountCustomFieldUpdate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AccountCustomFieldUpdate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AccountCustomFieldUpdate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


