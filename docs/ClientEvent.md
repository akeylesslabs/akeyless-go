# ClientEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Client event kind. One of: injector-cert | [optional] 
**InjectorCertificate** | Pointer to [**InjectorCertificateEvent**](InjectorCertificateEvent.md) |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewClientEvent

`func NewClientEvent() *ClientEvent`

NewClientEvent instantiates a new ClientEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientEventWithDefaults

`func NewClientEventWithDefaults() *ClientEvent`

NewClientEventWithDefaults instantiates a new ClientEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ClientEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ClientEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ClientEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ClientEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetInjectorCertificate

`func (o *ClientEvent) GetInjectorCertificate() InjectorCertificateEvent`

GetInjectorCertificate returns the InjectorCertificate field if non-nil, zero value otherwise.

### GetInjectorCertificateOk

`func (o *ClientEvent) GetInjectorCertificateOk() (*InjectorCertificateEvent, bool)`

GetInjectorCertificateOk returns a tuple with the InjectorCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectorCertificate

`func (o *ClientEvent) SetInjectorCertificate(v InjectorCertificateEvent)`

SetInjectorCertificate sets InjectorCertificate field to given value.

### HasInjectorCertificate

`func (o *ClientEvent) HasInjectorCertificate() bool`

HasInjectorCertificate returns a boolean if a field has been set.

### GetJson

`func (o *ClientEvent) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ClientEvent) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ClientEvent) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ClientEvent) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *ClientEvent) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientEvent) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientEvent) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ClientEvent) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ClientEvent) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ClientEvent) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ClientEvent) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ClientEvent) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


