# EventForwarderDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | EventForwarder name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewEventForwarderDelete

`func NewEventForwarderDelete(name string, ) *EventForwarderDelete`

NewEventForwarderDelete instantiates a new EventForwarderDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventForwarderDeleteWithDefaults

`func NewEventForwarderDeleteWithDefaults() *EventForwarderDelete`

NewEventForwarderDeleteWithDefaults instantiates a new EventForwarderDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *EventForwarderDelete) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EventForwarderDelete) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EventForwarderDelete) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EventForwarderDelete) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *EventForwarderDelete) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventForwarderDelete) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventForwarderDelete) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *EventForwarderDelete) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EventForwarderDelete) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EventForwarderDelete) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EventForwarderDelete) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EventForwarderDelete) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EventForwarderDelete) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EventForwarderDelete) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EventForwarderDelete) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


