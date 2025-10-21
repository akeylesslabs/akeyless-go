# EventForwarderCreateTeams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodsEventSourceLocations** | Pointer to **[]string** | Auth Method Event sources | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EventTypes** | Pointer to **[]string** | List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, next-automatic-rotation, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated, rate-limiting, usage-report, secret-sync] | [optional] 
**Every** | Pointer to **string** | Rate of periodic runner repetition in hours | [optional] 
**GatewaysEventSourceLocations** | **[]string** | Event sources | 
**ItemsEventSourceLocations** | Pointer to **[]string** | Items Event sources | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | EventForwarder name | 
**RunnerType** | **string** |  | 
**TargetsEventSourceLocations** | Pointer to **[]string** | Targets Event sources | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | **string** | Teams Webhook URL | 

## Methods

### NewEventForwarderCreateTeams

`func NewEventForwarderCreateTeams(gatewaysEventSourceLocations []string, name string, runnerType string, url string, ) *EventForwarderCreateTeams`

NewEventForwarderCreateTeams instantiates a new EventForwarderCreateTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventForwarderCreateTeamsWithDefaults

`func NewEventForwarderCreateTeamsWithDefaults() *EventForwarderCreateTeams`

NewEventForwarderCreateTeamsWithDefaults instantiates a new EventForwarderCreateTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodsEventSourceLocations

`func (o *EventForwarderCreateTeams) GetAuthMethodsEventSourceLocations() []string`

GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field if non-nil, zero value otherwise.

### GetAuthMethodsEventSourceLocationsOk

`func (o *EventForwarderCreateTeams) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool)`

GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodsEventSourceLocations

`func (o *EventForwarderCreateTeams) SetAuthMethodsEventSourceLocations(v []string)`

SetAuthMethodsEventSourceLocations sets AuthMethodsEventSourceLocations field to given value.

### HasAuthMethodsEventSourceLocations

`func (o *EventForwarderCreateTeams) HasAuthMethodsEventSourceLocations() bool`

HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.

### GetDescription

`func (o *EventForwarderCreateTeams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventForwarderCreateTeams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventForwarderCreateTeams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventForwarderCreateTeams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventForwarderCreateTeams) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventForwarderCreateTeams) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventForwarderCreateTeams) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *EventForwarderCreateTeams) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetEvery

`func (o *EventForwarderCreateTeams) GetEvery() string`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *EventForwarderCreateTeams) GetEveryOk() (*string, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *EventForwarderCreateTeams) SetEvery(v string)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *EventForwarderCreateTeams) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### GetGatewaysEventSourceLocations

`func (o *EventForwarderCreateTeams) GetGatewaysEventSourceLocations() []string`

GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field if non-nil, zero value otherwise.

### GetGatewaysEventSourceLocationsOk

`func (o *EventForwarderCreateTeams) GetGatewaysEventSourceLocationsOk() (*[]string, bool)`

GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysEventSourceLocations

`func (o *EventForwarderCreateTeams) SetGatewaysEventSourceLocations(v []string)`

SetGatewaysEventSourceLocations sets GatewaysEventSourceLocations field to given value.


### GetItemsEventSourceLocations

`func (o *EventForwarderCreateTeams) GetItemsEventSourceLocations() []string`

GetItemsEventSourceLocations returns the ItemsEventSourceLocations field if non-nil, zero value otherwise.

### GetItemsEventSourceLocationsOk

`func (o *EventForwarderCreateTeams) GetItemsEventSourceLocationsOk() (*[]string, bool)`

GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsEventSourceLocations

`func (o *EventForwarderCreateTeams) SetItemsEventSourceLocations(v []string)`

SetItemsEventSourceLocations sets ItemsEventSourceLocations field to given value.

### HasItemsEventSourceLocations

`func (o *EventForwarderCreateTeams) HasItemsEventSourceLocations() bool`

HasItemsEventSourceLocations returns a boolean if a field has been set.

### GetJson

`func (o *EventForwarderCreateTeams) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EventForwarderCreateTeams) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EventForwarderCreateTeams) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EventForwarderCreateTeams) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *EventForwarderCreateTeams) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventForwarderCreateTeams) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventForwarderCreateTeams) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EventForwarderCreateTeams) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *EventForwarderCreateTeams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventForwarderCreateTeams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventForwarderCreateTeams) SetName(v string)`

SetName sets Name field to given value.


### GetRunnerType

`func (o *EventForwarderCreateTeams) GetRunnerType() string`

GetRunnerType returns the RunnerType field if non-nil, zero value otherwise.

### GetRunnerTypeOk

`func (o *EventForwarderCreateTeams) GetRunnerTypeOk() (*string, bool)`

GetRunnerTypeOk returns a tuple with the RunnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerType

`func (o *EventForwarderCreateTeams) SetRunnerType(v string)`

SetRunnerType sets RunnerType field to given value.


### GetTargetsEventSourceLocations

`func (o *EventForwarderCreateTeams) GetTargetsEventSourceLocations() []string`

GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field if non-nil, zero value otherwise.

### GetTargetsEventSourceLocationsOk

`func (o *EventForwarderCreateTeams) GetTargetsEventSourceLocationsOk() (*[]string, bool)`

GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsEventSourceLocations

`func (o *EventForwarderCreateTeams) SetTargetsEventSourceLocations(v []string)`

SetTargetsEventSourceLocations sets TargetsEventSourceLocations field to given value.

### HasTargetsEventSourceLocations

`func (o *EventForwarderCreateTeams) HasTargetsEventSourceLocations() bool`

HasTargetsEventSourceLocations returns a boolean if a field has been set.

### GetToken

`func (o *EventForwarderCreateTeams) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EventForwarderCreateTeams) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EventForwarderCreateTeams) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EventForwarderCreateTeams) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EventForwarderCreateTeams) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EventForwarderCreateTeams) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EventForwarderCreateTeams) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EventForwarderCreateTeams) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *EventForwarderCreateTeams) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventForwarderCreateTeams) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventForwarderCreateTeams) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


