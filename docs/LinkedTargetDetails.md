# LinkedTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **map[string]string** | key hostname, value description | [optional] 

## Methods

### NewLinkedTargetDetails

`func NewLinkedTargetDetails() *LinkedTargetDetails`

NewLinkedTargetDetails instantiates a new LinkedTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedTargetDetailsWithDefaults

`func NewLinkedTargetDetailsWithDefaults() *LinkedTargetDetails`

NewLinkedTargetDetailsWithDefaults instantiates a new LinkedTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *LinkedTargetDetails) GetHosts() map[string]string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *LinkedTargetDetails) GetHostsOk() (*map[string]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *LinkedTargetDetails) SetHosts(v map[string]string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *LinkedTargetDetails) HasHosts() bool`

HasHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


