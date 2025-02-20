# SraSessionEntryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** |  | [optional] 
**ClientType** | Pointer to **string** |  | [optional] 
**ClusterUniqueId** | Pointer to **int64** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**ErrorMsg** | Pointer to **string** |  | [optional] 
**GatewayInfo** | Pointer to [**GatewayNameInfo**](GatewayNameInfo.md) |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**SecretName** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TargetHost** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **string** |  | [optional] 
**UserIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewSraSessionEntryOut

`func NewSraSessionEntryOut() *SraSessionEntryOut`

NewSraSessionEntryOut instantiates a new SraSessionEntryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSraSessionEntryOutWithDefaults

`func NewSraSessionEntryOutWithDefaults() *SraSessionEntryOut`

NewSraSessionEntryOutWithDefaults instantiates a new SraSessionEntryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *SraSessionEntryOut) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *SraSessionEntryOut) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *SraSessionEntryOut) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *SraSessionEntryOut) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetClientType

`func (o *SraSessionEntryOut) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *SraSessionEntryOut) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *SraSessionEntryOut) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *SraSessionEntryOut) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetClusterUniqueId

`func (o *SraSessionEntryOut) GetClusterUniqueId() int64`

GetClusterUniqueId returns the ClusterUniqueId field if non-nil, zero value otherwise.

### GetClusterUniqueIdOk

`func (o *SraSessionEntryOut) GetClusterUniqueIdOk() (*int64, bool)`

GetClusterUniqueIdOk returns a tuple with the ClusterUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUniqueId

`func (o *SraSessionEntryOut) SetClusterUniqueId(v int64)`

SetClusterUniqueId sets ClusterUniqueId field to given value.

### HasClusterUniqueId

`func (o *SraSessionEntryOut) HasClusterUniqueId() bool`

HasClusterUniqueId returns a boolean if a field has been set.

### GetConnectionType

`func (o *SraSessionEntryOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SraSessionEntryOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SraSessionEntryOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *SraSessionEntryOut) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetEndTime

`func (o *SraSessionEntryOut) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SraSessionEntryOut) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SraSessionEntryOut) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SraSessionEntryOut) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetErrorMsg

`func (o *SraSessionEntryOut) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *SraSessionEntryOut) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *SraSessionEntryOut) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *SraSessionEntryOut) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetGatewayInfo

`func (o *SraSessionEntryOut) GetGatewayInfo() GatewayNameInfo`

GetGatewayInfo returns the GatewayInfo field if non-nil, zero value otherwise.

### GetGatewayInfoOk

`func (o *SraSessionEntryOut) GetGatewayInfoOk() (*GatewayNameInfo, bool)`

GetGatewayInfoOk returns a tuple with the GatewayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInfo

`func (o *SraSessionEntryOut) SetGatewayInfo(v GatewayNameInfo)`

SetGatewayInfo sets GatewayInfo field to given value.

### HasGatewayInfo

`func (o *SraSessionEntryOut) HasGatewayInfo() bool`

HasGatewayInfo returns a boolean if a field has been set.

### GetInstanceId

`func (o *SraSessionEntryOut) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *SraSessionEntryOut) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *SraSessionEntryOut) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *SraSessionEntryOut) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetSecretName

`func (o *SraSessionEntryOut) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *SraSessionEntryOut) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *SraSessionEntryOut) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.

### HasSecretName

`func (o *SraSessionEntryOut) HasSecretName() bool`

HasSecretName returns a boolean if a field has been set.

### GetSessionId

`func (o *SraSessionEntryOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SraSessionEntryOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SraSessionEntryOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *SraSessionEntryOut) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStartTime

`func (o *SraSessionEntryOut) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SraSessionEntryOut) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SraSessionEntryOut) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SraSessionEntryOut) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *SraSessionEntryOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SraSessionEntryOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SraSessionEntryOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SraSessionEntryOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetHost

`func (o *SraSessionEntryOut) GetTargetHost() string`

GetTargetHost returns the TargetHost field if non-nil, zero value otherwise.

### GetTargetHostOk

`func (o *SraSessionEntryOut) GetTargetHostOk() (*string, bool)`

GetTargetHostOk returns a tuple with the TargetHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHost

`func (o *SraSessionEntryOut) SetTargetHost(v string)`

SetTargetHost sets TargetHost field to given value.

### HasTargetHost

`func (o *SraSessionEntryOut) HasTargetHost() bool`

HasTargetHost returns a boolean if a field has been set.

### GetTtl

`func (o *SraSessionEntryOut) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SraSessionEntryOut) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SraSessionEntryOut) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SraSessionEntryOut) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUserIdentifier

`func (o *SraSessionEntryOut) GetUserIdentifier() string`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *SraSessionEntryOut) GetUserIdentifierOk() (*string, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *SraSessionEntryOut) SetUserIdentifier(v string)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *SraSessionEntryOut) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


