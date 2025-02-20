# ListSraSessionsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedGateways** | Pointer to [**[]GatewayNameInfo**](GatewayNameInfo.md) |  | [optional] 
**NextPage** | Pointer to **string** |  | [optional] 
**Sessions** | Pointer to [**[]SraSessionEntryOut**](SraSessionEntryOut.md) |  | [optional] 

## Methods

### NewListSraSessionsOutput

`func NewListSraSessionsOutput() *ListSraSessionsOutput`

NewListSraSessionsOutput instantiates a new ListSraSessionsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSraSessionsOutputWithDefaults

`func NewListSraSessionsOutputWithDefaults() *ListSraSessionsOutput`

NewListSraSessionsOutputWithDefaults instantiates a new ListSraSessionsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedGateways

`func (o *ListSraSessionsOutput) GetAllowedGateways() []GatewayNameInfo`

GetAllowedGateways returns the AllowedGateways field if non-nil, zero value otherwise.

### GetAllowedGatewaysOk

`func (o *ListSraSessionsOutput) GetAllowedGatewaysOk() (*[]GatewayNameInfo, bool)`

GetAllowedGatewaysOk returns a tuple with the AllowedGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGateways

`func (o *ListSraSessionsOutput) SetAllowedGateways(v []GatewayNameInfo)`

SetAllowedGateways sets AllowedGateways field to given value.

### HasAllowedGateways

`func (o *ListSraSessionsOutput) HasAllowedGateways() bool`

HasAllowedGateways returns a boolean if a field has been set.

### GetNextPage

`func (o *ListSraSessionsOutput) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ListSraSessionsOutput) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ListSraSessionsOutput) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ListSraSessionsOutput) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetSessions

`func (o *ListSraSessionsOutput) GetSessions() []SraSessionEntryOut`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ListSraSessionsOutput) GetSessionsOk() (*[]SraSessionEntryOut, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ListSraSessionsOutput) SetSessions(v []SraSessionEntryOut)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ListSraSessionsOutput) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


