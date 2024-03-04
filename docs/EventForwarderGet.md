# EventForwarderGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | EventForwarder name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewEventForwarderGet

`func NewEventForwarderGet(name string, ) *EventForwarderGet`

NewEventForwarderGet instantiates a new EventForwarderGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventForwarderGetWithDefaults

`func NewEventForwarderGetWithDefaults() *EventForwarderGet`

NewEventForwarderGetWithDefaults instantiates a new EventForwarderGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *EventForwarderGet) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EventForwarderGet) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EventForwarderGet) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EventForwarderGet) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *EventForwarderGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventForwarderGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventForwarderGet) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *EventForwarderGet) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EventForwarderGet) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EventForwarderGet) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EventForwarderGet) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EventForwarderGet) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EventForwarderGet) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EventForwarderGet) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EventForwarderGet) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


