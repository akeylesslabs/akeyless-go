# ScimStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveIdentities** | Pointer to **int64** |  | [optional] 
**BaseUrl** | Pointer to **string** |  | [optional] 
**BearerAccessId** | Pointer to **string** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Degraded** | Pointer to **bool** | Degraded is true when live-token revocation is not active, so deactivation only denies new logins. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**GatewayUrl** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **int64** |  | [optional] 
**InactiveIdentities** | Pointer to **int64** |  | [optional] 
**LastWriteAction** | Pointer to **string** |  | [optional] 
**LastWriteStatus** | Pointer to **string** |  | [optional] 
**LastWriteTs** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewScimStatus

`func NewScimStatus() *ScimStatus`

NewScimStatus instantiates a new ScimStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimStatusWithDefaults

`func NewScimStatusWithDefaults() *ScimStatus`

NewScimStatusWithDefaults instantiates a new ScimStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveIdentities

`func (o *ScimStatus) GetActiveIdentities() int64`

GetActiveIdentities returns the ActiveIdentities field if non-nil, zero value otherwise.

### GetActiveIdentitiesOk

`func (o *ScimStatus) GetActiveIdentitiesOk() (*int64, bool)`

GetActiveIdentitiesOk returns a tuple with the ActiveIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveIdentities

`func (o *ScimStatus) SetActiveIdentities(v int64)`

SetActiveIdentities sets ActiveIdentities field to given value.

### HasActiveIdentities

`func (o *ScimStatus) HasActiveIdentities() bool`

HasActiveIdentities returns a boolean if a field has been set.

### GetBaseUrl

`func (o *ScimStatus) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ScimStatus) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ScimStatus) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ScimStatus) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetBearerAccessId

`func (o *ScimStatus) GetBearerAccessId() string`

GetBearerAccessId returns the BearerAccessId field if non-nil, zero value otherwise.

### GetBearerAccessIdOk

`func (o *ScimStatus) GetBearerAccessIdOk() (*string, bool)`

GetBearerAccessIdOk returns a tuple with the BearerAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerAccessId

`func (o *ScimStatus) SetBearerAccessId(v string)`

SetBearerAccessId sets BearerAccessId field to given value.

### HasBearerAccessId

`func (o *ScimStatus) HasBearerAccessId() bool`

HasBearerAccessId returns a boolean if a field has been set.

### GetConnected

`func (o *ScimStatus) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ScimStatus) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ScimStatus) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ScimStatus) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetDegraded

`func (o *ScimStatus) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *ScimStatus) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *ScimStatus) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *ScimStatus) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetEnabled

`func (o *ScimStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScimStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScimStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ScimStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEndpoint

`func (o *ScimStatus) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ScimStatus) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ScimStatus) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ScimStatus) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetGatewayUrl

`func (o *ScimStatus) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *ScimStatus) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *ScimStatus) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *ScimStatus) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetGroups

`func (o *ScimStatus) GetGroups() int64`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ScimStatus) GetGroupsOk() (*int64, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ScimStatus) SetGroups(v int64)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ScimStatus) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetInactiveIdentities

`func (o *ScimStatus) GetInactiveIdentities() int64`

GetInactiveIdentities returns the InactiveIdentities field if non-nil, zero value otherwise.

### GetInactiveIdentitiesOk

`func (o *ScimStatus) GetInactiveIdentitiesOk() (*int64, bool)`

GetInactiveIdentitiesOk returns a tuple with the InactiveIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveIdentities

`func (o *ScimStatus) SetInactiveIdentities(v int64)`

SetInactiveIdentities sets InactiveIdentities field to given value.

### HasInactiveIdentities

`func (o *ScimStatus) HasInactiveIdentities() bool`

HasInactiveIdentities returns a boolean if a field has been set.

### GetLastWriteAction

`func (o *ScimStatus) GetLastWriteAction() string`

GetLastWriteAction returns the LastWriteAction field if non-nil, zero value otherwise.

### GetLastWriteActionOk

`func (o *ScimStatus) GetLastWriteActionOk() (*string, bool)`

GetLastWriteActionOk returns a tuple with the LastWriteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWriteAction

`func (o *ScimStatus) SetLastWriteAction(v string)`

SetLastWriteAction sets LastWriteAction field to given value.

### HasLastWriteAction

`func (o *ScimStatus) HasLastWriteAction() bool`

HasLastWriteAction returns a boolean if a field has been set.

### GetLastWriteStatus

`func (o *ScimStatus) GetLastWriteStatus() string`

GetLastWriteStatus returns the LastWriteStatus field if non-nil, zero value otherwise.

### GetLastWriteStatusOk

`func (o *ScimStatus) GetLastWriteStatusOk() (*string, bool)`

GetLastWriteStatusOk returns a tuple with the LastWriteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWriteStatus

`func (o *ScimStatus) SetLastWriteStatus(v string)`

SetLastWriteStatus sets LastWriteStatus field to given value.

### HasLastWriteStatus

`func (o *ScimStatus) HasLastWriteStatus() bool`

HasLastWriteStatus returns a boolean if a field has been set.

### GetLastWriteTs

`func (o *ScimStatus) GetLastWriteTs() time.Time`

GetLastWriteTs returns the LastWriteTs field if non-nil, zero value otherwise.

### GetLastWriteTsOk

`func (o *ScimStatus) GetLastWriteTsOk() (*time.Time, bool)`

GetLastWriteTsOk returns a tuple with the LastWriteTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWriteTs

`func (o *ScimStatus) SetLastWriteTs(v time.Time)`

SetLastWriteTs sets LastWriteTs field to given value.

### HasLastWriteTs

`func (o *ScimStatus) HasLastWriteTs() bool`

HasLastWriteTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


