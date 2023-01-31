# KMIPClientListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**[]KMIPClient**](KMIPClient.md) |  | [optional] 

## Methods

### NewKMIPClientListResponse

`func NewKMIPClientListResponse() *KMIPClientListResponse`

NewKMIPClientListResponse instantiates a new KMIPClientListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMIPClientListResponseWithDefaults

`func NewKMIPClientListResponseWithDefaults() *KMIPClientListResponse`

NewKMIPClientListResponseWithDefaults instantiates a new KMIPClientListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *KMIPClientListResponse) GetClients() []KMIPClient`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *KMIPClientListResponse) GetClientsOk() (*[]KMIPClient, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *KMIPClientListResponse) SetClients(v []KMIPClient)`

SetClients sets Clients field to given value.

### HasClients

`func (o *KMIPClientListResponse) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


