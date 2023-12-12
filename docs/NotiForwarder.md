# NotiForwarder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** | Auth - JWT | [optional] 
**ClientPermissions** | Pointer to **[]string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**EventTypes** | Pointer to **[]string** |  | [optional] 
**GatewayClusterId** | Pointer to **int64** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**LastVersion** | Pointer to **int32** |  | [optional] 
**ModificationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**NotiForwarderId** | Pointer to **int64** |  | [optional] 
**NotiForwarderName** | Pointer to **string** |  | [optional] 
**NotiForwarderType** | Pointer to **string** |  | [optional] 
**NotiForwarderVersions** | Pointer to [**[]ItemVersion**](ItemVersion.md) |  | [optional] 
**Paths** | Pointer to **[]string** |  | [optional] 
**ProtectionKey** | Pointer to **string** |  | [optional] 
**RunnerType** | Pointer to **string** |  | [optional] 
**TimespanInSeconds** | Pointer to **int64** |  | [optional] 
**ToEmails** | Pointer to [**[]EmailEntry**](EmailEntry.md) |  | [optional] 
**UserEmail** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** | Auth - User Password | [optional] 
**WithCustomerFragment** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotiForwarder

`func NewNotiForwarder() *NotiForwarder`

NewNotiForwarder instantiates a new NotiForwarder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotiForwarderWithDefaults

`func NewNotiForwarderWithDefaults() *NotiForwarder`

NewNotiForwarderWithDefaults instantiates a new NotiForwarder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *NotiForwarder) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *NotiForwarder) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *NotiForwarder) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *NotiForwarder) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *NotiForwarder) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NotiForwarder) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NotiForwarder) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *NotiForwarder) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientPermissions

`func (o *NotiForwarder) GetClientPermissions() []string`

GetClientPermissions returns the ClientPermissions field if non-nil, zero value otherwise.

### GetClientPermissionsOk

`func (o *NotiForwarder) GetClientPermissionsOk() (*[]string, bool)`

GetClientPermissionsOk returns a tuple with the ClientPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPermissions

`func (o *NotiForwarder) SetClientPermissions(v []string)`

SetClientPermissions sets ClientPermissions field to given value.

### HasClientPermissions

`func (o *NotiForwarder) HasClientPermissions() bool`

HasClientPermissions returns a boolean if a field has been set.

### GetComment

`func (o *NotiForwarder) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NotiForwarder) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NotiForwarder) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NotiForwarder) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationDate

`func (o *NotiForwarder) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *NotiForwarder) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *NotiForwarder) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *NotiForwarder) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetEndpoint

`func (o *NotiForwarder) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *NotiForwarder) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *NotiForwarder) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *NotiForwarder) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEventTypes

`func (o *NotiForwarder) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *NotiForwarder) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *NotiForwarder) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *NotiForwarder) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetGatewayClusterId

`func (o *NotiForwarder) GetGatewayClusterId() int64`

GetGatewayClusterId returns the GatewayClusterId field if non-nil, zero value otherwise.

### GetGatewayClusterIdOk

`func (o *NotiForwarder) GetGatewayClusterIdOk() (*int64, bool)`

GetGatewayClusterIdOk returns a tuple with the GatewayClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayClusterId

`func (o *NotiForwarder) SetGatewayClusterId(v int64)`

SetGatewayClusterId sets GatewayClusterId field to given value.

### HasGatewayClusterId

`func (o *NotiForwarder) HasGatewayClusterId() bool`

HasGatewayClusterId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *NotiForwarder) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotiForwarder) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotiForwarder) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *NotiForwarder) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetLastVersion

`func (o *NotiForwarder) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *NotiForwarder) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *NotiForwarder) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *NotiForwarder) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### GetModificationDate

`func (o *NotiForwarder) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *NotiForwarder) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *NotiForwarder) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *NotiForwarder) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetNotiForwarderId

`func (o *NotiForwarder) GetNotiForwarderId() int64`

GetNotiForwarderId returns the NotiForwarderId field if non-nil, zero value otherwise.

### GetNotiForwarderIdOk

`func (o *NotiForwarder) GetNotiForwarderIdOk() (*int64, bool)`

GetNotiForwarderIdOk returns a tuple with the NotiForwarderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotiForwarderId

`func (o *NotiForwarder) SetNotiForwarderId(v int64)`

SetNotiForwarderId sets NotiForwarderId field to given value.

### HasNotiForwarderId

`func (o *NotiForwarder) HasNotiForwarderId() bool`

HasNotiForwarderId returns a boolean if a field has been set.

### GetNotiForwarderName

`func (o *NotiForwarder) GetNotiForwarderName() string`

GetNotiForwarderName returns the NotiForwarderName field if non-nil, zero value otherwise.

### GetNotiForwarderNameOk

`func (o *NotiForwarder) GetNotiForwarderNameOk() (*string, bool)`

GetNotiForwarderNameOk returns a tuple with the NotiForwarderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotiForwarderName

`func (o *NotiForwarder) SetNotiForwarderName(v string)`

SetNotiForwarderName sets NotiForwarderName field to given value.

### HasNotiForwarderName

`func (o *NotiForwarder) HasNotiForwarderName() bool`

HasNotiForwarderName returns a boolean if a field has been set.

### GetNotiForwarderType

`func (o *NotiForwarder) GetNotiForwarderType() string`

GetNotiForwarderType returns the NotiForwarderType field if non-nil, zero value otherwise.

### GetNotiForwarderTypeOk

`func (o *NotiForwarder) GetNotiForwarderTypeOk() (*string, bool)`

GetNotiForwarderTypeOk returns a tuple with the NotiForwarderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotiForwarderType

`func (o *NotiForwarder) SetNotiForwarderType(v string)`

SetNotiForwarderType sets NotiForwarderType field to given value.

### HasNotiForwarderType

`func (o *NotiForwarder) HasNotiForwarderType() bool`

HasNotiForwarderType returns a boolean if a field has been set.

### GetNotiForwarderVersions

`func (o *NotiForwarder) GetNotiForwarderVersions() []ItemVersion`

GetNotiForwarderVersions returns the NotiForwarderVersions field if non-nil, zero value otherwise.

### GetNotiForwarderVersionsOk

`func (o *NotiForwarder) GetNotiForwarderVersionsOk() (*[]ItemVersion, bool)`

GetNotiForwarderVersionsOk returns a tuple with the NotiForwarderVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotiForwarderVersions

`func (o *NotiForwarder) SetNotiForwarderVersions(v []ItemVersion)`

SetNotiForwarderVersions sets NotiForwarderVersions field to given value.

### HasNotiForwarderVersions

`func (o *NotiForwarder) HasNotiForwarderVersions() bool`

HasNotiForwarderVersions returns a boolean if a field has been set.

### GetPaths

`func (o *NotiForwarder) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *NotiForwarder) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *NotiForwarder) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *NotiForwarder) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetProtectionKey

`func (o *NotiForwarder) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *NotiForwarder) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *NotiForwarder) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *NotiForwarder) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetRunnerType

`func (o *NotiForwarder) GetRunnerType() string`

GetRunnerType returns the RunnerType field if non-nil, zero value otherwise.

### GetRunnerTypeOk

`func (o *NotiForwarder) GetRunnerTypeOk() (*string, bool)`

GetRunnerTypeOk returns a tuple with the RunnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerType

`func (o *NotiForwarder) SetRunnerType(v string)`

SetRunnerType sets RunnerType field to given value.

### HasRunnerType

`func (o *NotiForwarder) HasRunnerType() bool`

HasRunnerType returns a boolean if a field has been set.

### GetTimespanInSeconds

`func (o *NotiForwarder) GetTimespanInSeconds() int64`

GetTimespanInSeconds returns the TimespanInSeconds field if non-nil, zero value otherwise.

### GetTimespanInSecondsOk

`func (o *NotiForwarder) GetTimespanInSecondsOk() (*int64, bool)`

GetTimespanInSecondsOk returns a tuple with the TimespanInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimespanInSeconds

`func (o *NotiForwarder) SetTimespanInSeconds(v int64)`

SetTimespanInSeconds sets TimespanInSeconds field to given value.

### HasTimespanInSeconds

`func (o *NotiForwarder) HasTimespanInSeconds() bool`

HasTimespanInSeconds returns a boolean if a field has been set.

### GetToEmails

`func (o *NotiForwarder) GetToEmails() []EmailEntry`

GetToEmails returns the ToEmails field if non-nil, zero value otherwise.

### GetToEmailsOk

`func (o *NotiForwarder) GetToEmailsOk() (*[]EmailEntry, bool)`

GetToEmailsOk returns a tuple with the ToEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEmails

`func (o *NotiForwarder) SetToEmails(v []EmailEntry)`

SetToEmails sets ToEmails field to given value.

### HasToEmails

`func (o *NotiForwarder) HasToEmails() bool`

HasToEmails returns a boolean if a field has been set.

### GetUserEmail

`func (o *NotiForwarder) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *NotiForwarder) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *NotiForwarder) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *NotiForwarder) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUsername

`func (o *NotiForwarder) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotiForwarder) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotiForwarder) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotiForwarder) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWithCustomerFragment

`func (o *NotiForwarder) GetWithCustomerFragment() bool`

GetWithCustomerFragment returns the WithCustomerFragment field if non-nil, zero value otherwise.

### GetWithCustomerFragmentOk

`func (o *NotiForwarder) GetWithCustomerFragmentOk() (*bool, bool)`

GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCustomerFragment

`func (o *NotiForwarder) SetWithCustomerFragment(v bool)`

SetWithCustomerFragment sets WithCustomerFragment field to given value.

### HasWithCustomerFragment

`func (o *NotiForwarder) HasWithCustomerFragment() bool`

HasWithCustomerFragment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


