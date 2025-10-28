# AccountCustomFieldList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Object** | Pointer to **string** | The object to filter by the custom fields in: body | [optional] 
**ObjectType** | Pointer to **string** | The object type to filter by the custom fields in: body | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewAccountCustomFieldList

`func NewAccountCustomFieldList() *AccountCustomFieldList`

NewAccountCustomFieldList instantiates a new AccountCustomFieldList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCustomFieldListWithDefaults

`func NewAccountCustomFieldListWithDefaults() *AccountCustomFieldList`

NewAccountCustomFieldListWithDefaults instantiates a new AccountCustomFieldList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *AccountCustomFieldList) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AccountCustomFieldList) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AccountCustomFieldList) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AccountCustomFieldList) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetObject

`func (o *AccountCustomFieldList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AccountCustomFieldList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AccountCustomFieldList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AccountCustomFieldList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectType

`func (o *AccountCustomFieldList) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccountCustomFieldList) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccountCustomFieldList) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AccountCustomFieldList) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetToken

`func (o *AccountCustomFieldList) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountCustomFieldList) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountCustomFieldList) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountCustomFieldList) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AccountCustomFieldList) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AccountCustomFieldList) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AccountCustomFieldList) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AccountCustomFieldList) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


