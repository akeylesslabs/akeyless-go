# KMIPClientsConfigPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**map[string]KMIPClient**](KMIPClient.md) |  | [optional] 
**Environments** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKMIPClientsConfigPart

`func NewKMIPClientsConfigPart() *KMIPClientsConfigPart`

NewKMIPClientsConfigPart instantiates a new KMIPClientsConfigPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMIPClientsConfigPartWithDefaults

`func NewKMIPClientsConfigPartWithDefaults() *KMIPClientsConfigPart`

NewKMIPClientsConfigPartWithDefaults instantiates a new KMIPClientsConfigPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *KMIPClientsConfigPart) GetClients() map[string]KMIPClient`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *KMIPClientsConfigPart) GetClientsOk() (*map[string]KMIPClient, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *KMIPClientsConfigPart) SetClients(v map[string]KMIPClient)`

SetClients sets Clients field to given value.

### HasClients

`func (o *KMIPClientsConfigPart) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetEnvironments

`func (o *KMIPClientsConfigPart) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *KMIPClientsConfigPart) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *KMIPClientsConfigPart) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *KMIPClientsConfigPart) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


