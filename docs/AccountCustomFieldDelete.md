# AccountCustomFieldDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The custom field id | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewAccountCustomFieldDelete

`func NewAccountCustomFieldDelete(id int64, ) *AccountCustomFieldDelete`

NewAccountCustomFieldDelete instantiates a new AccountCustomFieldDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCustomFieldDeleteWithDefaults

`func NewAccountCustomFieldDeleteWithDefaults() *AccountCustomFieldDelete`

NewAccountCustomFieldDeleteWithDefaults instantiates a new AccountCustomFieldDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountCustomFieldDelete) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountCustomFieldDelete) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountCustomFieldDelete) SetId(v int64)`

SetId sets Id field to given value.


### GetJson

`func (o *AccountCustomFieldDelete) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AccountCustomFieldDelete) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AccountCustomFieldDelete) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AccountCustomFieldDelete) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *AccountCustomFieldDelete) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountCustomFieldDelete) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountCustomFieldDelete) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountCustomFieldDelete) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AccountCustomFieldDelete) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AccountCustomFieldDelete) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AccountCustomFieldDelete) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AccountCustomFieldDelete) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


