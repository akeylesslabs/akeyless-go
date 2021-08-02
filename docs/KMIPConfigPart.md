# KMIPConfigPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**map[string]KMIPClient**](KMIPClient.md) |  | [optional] 
**ServerEnc** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewKMIPConfigPart

`func NewKMIPConfigPart() *KMIPConfigPart`

NewKMIPConfigPart instantiates a new KMIPConfigPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMIPConfigPartWithDefaults

`func NewKMIPConfigPartWithDefaults() *KMIPConfigPart`

NewKMIPConfigPartWithDefaults instantiates a new KMIPConfigPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *KMIPConfigPart) GetClients() map[string]KMIPClient`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *KMIPConfigPart) GetClientsOk() (*map[string]KMIPClient, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *KMIPConfigPart) SetClients(v map[string]KMIPClient)`

SetClients sets Clients field to given value.

### HasClients

`func (o *KMIPConfigPart) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetServerEnc

`func (o *KMIPConfigPart) GetServerEnc() []int32`

GetServerEnc returns the ServerEnc field if non-nil, zero value otherwise.

### GetServerEncOk

`func (o *KMIPConfigPart) GetServerEncOk() (*[]int32, bool)`

GetServerEncOk returns a tuple with the ServerEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerEnc

`func (o *KMIPConfigPart) SetServerEnc(v []int32)`

SetServerEnc sets ServerEnc field to given value.

### HasServerEnc

`func (o *KMIPConfigPart) HasServerEnc() bool`

HasServerEnc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


