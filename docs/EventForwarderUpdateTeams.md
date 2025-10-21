# EventForwarderUpdateTeams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodsEventSourceLocations** | Pointer to **[]string** | Auth Method Event sources | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Enable** | Pointer to **string** | Enable/Disable Event Forwarder [true/false] | [optional] [default to "true"]
**EventTypes** | Pointer to **[]string** | List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, next-automatic-rotation, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated, rate-limiting, usage-report, secret-sync] | [optional] 
**GatewaysEventSourceLocations** | **[]string** | Event sources | 
**ItemsEventSourceLocations** | Pointer to **[]string** | Items Event sources | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | EventForwarder name | 
**NewName** | Pointer to **string** | New EventForwarder name | [optional] 
**TargetsEventSourceLocations** | Pointer to **[]string** | Targets Event sources | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | **string** | Teams Webhook URL | 

## Methods

### NewEventForwarderUpdateTeams

`func NewEventForwarderUpdateTeams(gatewaysEventSourceLocations []string, name string, url string, ) *EventForwarderUpdateTeams`

NewEventForwarderUpdateTeams instantiates a new EventForwarderUpdateTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventForwarderUpdateTeamsWithDefaults

`func NewEventForwarderUpdateTeamsWithDefaults() *EventForwarderUpdateTeams`

NewEventForwarderUpdateTeamsWithDefaults instantiates a new EventForwarderUpdateTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodsEventSourceLocations

`func (o *EventForwarderUpdateTeams) GetAuthMethodsEventSourceLocations() []string`

GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field if non-nil, zero value otherwise.

### GetAuthMethodsEventSourceLocationsOk

`func (o *EventForwarderUpdateTeams) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool)`

GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodsEventSourceLocations

`func (o *EventForwarderUpdateTeams) SetAuthMethodsEventSourceLocations(v []string)`

SetAuthMethodsEventSourceLocations sets AuthMethodsEventSourceLocations field to given value.

### HasAuthMethodsEventSourceLocations

`func (o *EventForwarderUpdateTeams) HasAuthMethodsEventSourceLocations() bool`

HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.

### GetDescription

`func (o *EventForwarderUpdateTeams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventForwarderUpdateTeams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventForwarderUpdateTeams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventForwarderUpdateTeams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnable

`func (o *EventForwarderUpdateTeams) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *EventForwarderUpdateTeams) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *EventForwarderUpdateTeams) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *EventForwarderUpdateTeams) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventForwarderUpdateTeams) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventForwarderUpdateTeams) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventForwarderUpdateTeams) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *EventForwarderUpdateTeams) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetGatewaysEventSourceLocations

`func (o *EventForwarderUpdateTeams) GetGatewaysEventSourceLocations() []string`

GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field if non-nil, zero value otherwise.

### GetGatewaysEventSourceLocationsOk

`func (o *EventForwarderUpdateTeams) GetGatewaysEventSourceLocationsOk() (*[]string, bool)`

GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysEventSourceLocations

`func (o *EventForwarderUpdateTeams) SetGatewaysEventSourceLocations(v []string)`

SetGatewaysEventSourceLocations sets GatewaysEventSourceLocations field to given value.


### GetItemsEventSourceLocations

`func (o *EventForwarderUpdateTeams) GetItemsEventSourceLocations() []string`

GetItemsEventSourceLocations returns the ItemsEventSourceLocations field if non-nil, zero value otherwise.

### GetItemsEventSourceLocationsOk

`func (o *EventForwarderUpdateTeams) GetItemsEventSourceLocationsOk() (*[]string, bool)`

GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsEventSourceLocations

`func (o *EventForwarderUpdateTeams) SetItemsEventSourceLocations(v []string)`

SetItemsEventSourceLocations sets ItemsEventSourceLocations field to given value.

### HasItemsEventSourceLocations

`func (o *EventForwarderUpdateTeams) HasItemsEventSourceLocations() bool`

HasItemsEventSourceLocations returns a boolean if a field has been set.

### GetJson

`func (o *EventForwarderUpdateTeams) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EventForwarderUpdateTeams) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EventForwarderUpdateTeams) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EventForwarderUpdateTeams) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *EventForwarderUpdateTeams) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *EventForwarderUpdateTeams) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *EventForwarderUpdateTeams) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *EventForwarderUpdateTeams) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *EventForwarderUpdateTeams) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventForwarderUpdateTeams) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventForwarderUpdateTeams) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EventForwarderUpdateTeams) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *EventForwarderUpdateTeams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventForwarderUpdateTeams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventForwarderUpdateTeams) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *EventForwarderUpdateTeams) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *EventForwarderUpdateTeams) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *EventForwarderUpdateTeams) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *EventForwarderUpdateTeams) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetTargetsEventSourceLocations

`func (o *EventForwarderUpdateTeams) GetTargetsEventSourceLocations() []string`

GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field if non-nil, zero value otherwise.

### GetTargetsEventSourceLocationsOk

`func (o *EventForwarderUpdateTeams) GetTargetsEventSourceLocationsOk() (*[]string, bool)`

GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsEventSourceLocations

`func (o *EventForwarderUpdateTeams) SetTargetsEventSourceLocations(v []string)`

SetTargetsEventSourceLocations sets TargetsEventSourceLocations field to given value.

### HasTargetsEventSourceLocations

`func (o *EventForwarderUpdateTeams) HasTargetsEventSourceLocations() bool`

HasTargetsEventSourceLocations returns a boolean if a field has been set.

### GetToken

`func (o *EventForwarderUpdateTeams) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EventForwarderUpdateTeams) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EventForwarderUpdateTeams) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EventForwarderUpdateTeams) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EventForwarderUpdateTeams) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EventForwarderUpdateTeams) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EventForwarderUpdateTeams) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EventForwarderUpdateTeams) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *EventForwarderUpdateTeams) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventForwarderUpdateTeams) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventForwarderUpdateTeams) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


