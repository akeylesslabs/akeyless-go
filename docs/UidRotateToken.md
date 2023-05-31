# UidRotateToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fork** | Pointer to **bool** | Create a new child token with default parameters | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**SendManualAckToken** | Pointer to **string** | The new rotated token to send manual ack for (with uid-token&#x3D;the-orig-token) | [optional] 
**UidToken** | Pointer to **string** | The Universal identity token | [optional] 
**WithManualAck** | Pointer to **bool** | Disable automatic ack | [optional] 

## Methods

### NewUidRotateToken

`func NewUidRotateToken() *UidRotateToken`

NewUidRotateToken instantiates a new UidRotateToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUidRotateTokenWithDefaults

`func NewUidRotateTokenWithDefaults() *UidRotateToken`

NewUidRotateTokenWithDefaults instantiates a new UidRotateToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFork

`func (o *UidRotateToken) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *UidRotateToken) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *UidRotateToken) SetFork(v bool)`

SetFork sets Fork field to given value.

### HasFork

`func (o *UidRotateToken) HasFork() bool`

HasFork returns a boolean if a field has been set.

### GetJson

`func (o *UidRotateToken) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UidRotateToken) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UidRotateToken) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UidRotateToken) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSendManualAckToken

`func (o *UidRotateToken) GetSendManualAckToken() string`

GetSendManualAckToken returns the SendManualAckToken field if non-nil, zero value otherwise.

### GetSendManualAckTokenOk

`func (o *UidRotateToken) GetSendManualAckTokenOk() (*string, bool)`

GetSendManualAckTokenOk returns a tuple with the SendManualAckToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendManualAckToken

`func (o *UidRotateToken) SetSendManualAckToken(v string)`

SetSendManualAckToken sets SendManualAckToken field to given value.

### HasSendManualAckToken

`func (o *UidRotateToken) HasSendManualAckToken() bool`

HasSendManualAckToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UidRotateToken) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UidRotateToken) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UidRotateToken) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UidRotateToken) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetWithManualAck

`func (o *UidRotateToken) GetWithManualAck() bool`

GetWithManualAck returns the WithManualAck field if non-nil, zero value otherwise.

### GetWithManualAckOk

`func (o *UidRotateToken) GetWithManualAckOk() (*bool, bool)`

GetWithManualAckOk returns a tuple with the WithManualAck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithManualAck

`func (o *UidRotateToken) SetWithManualAck(v bool)`

SetWithManualAck sets WithManualAck field to given value.

### HasWithManualAck

`func (o *UidRotateToken) HasWithManualAck() bool`

HasWithManualAck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


