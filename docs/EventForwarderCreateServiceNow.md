# EventForwarderCreateServiceNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminName** | Pointer to **string** | Workstation Admin Name | [optional] 
**AdminPwd** | Pointer to **string** | Workstation Admin Password | [optional] 
**AppPrivateKeyBase64** | Pointer to **string** | The RSA Private Key to use when connecting with jwt authentication | [optional] 
**AuthMethodsEventSourceLocations** | Pointer to **[]string** | Auth Method Event sources | [optional] 
**AuthType** | Pointer to **string** | The authentication type to use [user-pass/jwt] | [optional] [default to "user-pass"]
**ClientId** | Pointer to **string** | The client ID to use when connecting with jwt authentication | [optional] 
**ClientSecret** | Pointer to **string** | The client secret to use when connecting with jwt authentication | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EventTypes** | Pointer to **[]string** | List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated] | [optional] 
**Every** | Pointer to **string** | Rate of periodic runner repetition in hours | [optional] 
**GatewaysEventSourceLocations** | **[]string** | Event sources | 
**Host** | Pointer to **string** | Workstation Host | [optional] 
**ItemsEventSourceLocations** | Pointer to **[]string** | Items Event sources | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | EventForwarder name | 
**RunnerType** | **string** |  | 
**TargetsEventSourceLocations** | Pointer to **[]string** | Targets Event sources | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserEmail** | Pointer to **string** | The user email to identify with when connecting with jwt authentication | [optional] 

## Methods

### NewEventForwarderCreateServiceNow

`func NewEventForwarderCreateServiceNow(gatewaysEventSourceLocations []string, name string, runnerType string, ) *EventForwarderCreateServiceNow`

NewEventForwarderCreateServiceNow instantiates a new EventForwarderCreateServiceNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventForwarderCreateServiceNowWithDefaults

`func NewEventForwarderCreateServiceNowWithDefaults() *EventForwarderCreateServiceNow`

NewEventForwarderCreateServiceNowWithDefaults instantiates a new EventForwarderCreateServiceNow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminName

`func (o *EventForwarderCreateServiceNow) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *EventForwarderCreateServiceNow) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *EventForwarderCreateServiceNow) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *EventForwarderCreateServiceNow) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAdminPwd

`func (o *EventForwarderCreateServiceNow) GetAdminPwd() string`

GetAdminPwd returns the AdminPwd field if non-nil, zero value otherwise.

### GetAdminPwdOk

`func (o *EventForwarderCreateServiceNow) GetAdminPwdOk() (*string, bool)`

GetAdminPwdOk returns a tuple with the AdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPwd

`func (o *EventForwarderCreateServiceNow) SetAdminPwd(v string)`

SetAdminPwd sets AdminPwd field to given value.

### HasAdminPwd

`func (o *EventForwarderCreateServiceNow) HasAdminPwd() bool`

HasAdminPwd returns a boolean if a field has been set.

### GetAppPrivateKeyBase64

`func (o *EventForwarderCreateServiceNow) GetAppPrivateKeyBase64() string`

GetAppPrivateKeyBase64 returns the AppPrivateKeyBase64 field if non-nil, zero value otherwise.

### GetAppPrivateKeyBase64Ok

`func (o *EventForwarderCreateServiceNow) GetAppPrivateKeyBase64Ok() (*string, bool)`

GetAppPrivateKeyBase64Ok returns a tuple with the AppPrivateKeyBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPrivateKeyBase64

`func (o *EventForwarderCreateServiceNow) SetAppPrivateKeyBase64(v string)`

SetAppPrivateKeyBase64 sets AppPrivateKeyBase64 field to given value.

### HasAppPrivateKeyBase64

`func (o *EventForwarderCreateServiceNow) HasAppPrivateKeyBase64() bool`

HasAppPrivateKeyBase64 returns a boolean if a field has been set.

### GetAuthMethodsEventSourceLocations

`func (o *EventForwarderCreateServiceNow) GetAuthMethodsEventSourceLocations() []string`

GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field if non-nil, zero value otherwise.

### GetAuthMethodsEventSourceLocationsOk

`func (o *EventForwarderCreateServiceNow) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool)`

GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodsEventSourceLocations

`func (o *EventForwarderCreateServiceNow) SetAuthMethodsEventSourceLocations(v []string)`

SetAuthMethodsEventSourceLocations sets AuthMethodsEventSourceLocations field to given value.

### HasAuthMethodsEventSourceLocations

`func (o *EventForwarderCreateServiceNow) HasAuthMethodsEventSourceLocations() bool`

HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.

### GetAuthType

`func (o *EventForwarderCreateServiceNow) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *EventForwarderCreateServiceNow) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *EventForwarderCreateServiceNow) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *EventForwarderCreateServiceNow) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *EventForwarderCreateServiceNow) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EventForwarderCreateServiceNow) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EventForwarderCreateServiceNow) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *EventForwarderCreateServiceNow) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *EventForwarderCreateServiceNow) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *EventForwarderCreateServiceNow) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *EventForwarderCreateServiceNow) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *EventForwarderCreateServiceNow) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDescription

`func (o *EventForwarderCreateServiceNow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventForwarderCreateServiceNow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventForwarderCreateServiceNow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventForwarderCreateServiceNow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventForwarderCreateServiceNow) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventForwarderCreateServiceNow) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventForwarderCreateServiceNow) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *EventForwarderCreateServiceNow) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetEvery

`func (o *EventForwarderCreateServiceNow) GetEvery() string`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *EventForwarderCreateServiceNow) GetEveryOk() (*string, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *EventForwarderCreateServiceNow) SetEvery(v string)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *EventForwarderCreateServiceNow) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### GetGatewaysEventSourceLocations

`func (o *EventForwarderCreateServiceNow) GetGatewaysEventSourceLocations() []string`

GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field if non-nil, zero value otherwise.

### GetGatewaysEventSourceLocationsOk

`func (o *EventForwarderCreateServiceNow) GetGatewaysEventSourceLocationsOk() (*[]string, bool)`

GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysEventSourceLocations

`func (o *EventForwarderCreateServiceNow) SetGatewaysEventSourceLocations(v []string)`

SetGatewaysEventSourceLocations sets GatewaysEventSourceLocations field to given value.


### GetHost

`func (o *EventForwarderCreateServiceNow) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EventForwarderCreateServiceNow) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EventForwarderCreateServiceNow) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EventForwarderCreateServiceNow) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetItemsEventSourceLocations

`func (o *EventForwarderCreateServiceNow) GetItemsEventSourceLocations() []string`

GetItemsEventSourceLocations returns the ItemsEventSourceLocations field if non-nil, zero value otherwise.

### GetItemsEventSourceLocationsOk

`func (o *EventForwarderCreateServiceNow) GetItemsEventSourceLocationsOk() (*[]string, bool)`

GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsEventSourceLocations

`func (o *EventForwarderCreateServiceNow) SetItemsEventSourceLocations(v []string)`

SetItemsEventSourceLocations sets ItemsEventSourceLocations field to given value.

### HasItemsEventSourceLocations

`func (o *EventForwarderCreateServiceNow) HasItemsEventSourceLocations() bool`

HasItemsEventSourceLocations returns a boolean if a field has been set.

### GetJson

`func (o *EventForwarderCreateServiceNow) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EventForwarderCreateServiceNow) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EventForwarderCreateServiceNow) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EventForwarderCreateServiceNow) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *EventForwarderCreateServiceNow) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventForwarderCreateServiceNow) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventForwarderCreateServiceNow) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EventForwarderCreateServiceNow) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *EventForwarderCreateServiceNow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventForwarderCreateServiceNow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventForwarderCreateServiceNow) SetName(v string)`

SetName sets Name field to given value.


### GetRunnerType

`func (o *EventForwarderCreateServiceNow) GetRunnerType() string`

GetRunnerType returns the RunnerType field if non-nil, zero value otherwise.

### GetRunnerTypeOk

`func (o *EventForwarderCreateServiceNow) GetRunnerTypeOk() (*string, bool)`

GetRunnerTypeOk returns a tuple with the RunnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerType

`func (o *EventForwarderCreateServiceNow) SetRunnerType(v string)`

SetRunnerType sets RunnerType field to given value.


### GetTargetsEventSourceLocations

`func (o *EventForwarderCreateServiceNow) GetTargetsEventSourceLocations() []string`

GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field if non-nil, zero value otherwise.

### GetTargetsEventSourceLocationsOk

`func (o *EventForwarderCreateServiceNow) GetTargetsEventSourceLocationsOk() (*[]string, bool)`

GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsEventSourceLocations

`func (o *EventForwarderCreateServiceNow) SetTargetsEventSourceLocations(v []string)`

SetTargetsEventSourceLocations sets TargetsEventSourceLocations field to given value.

### HasTargetsEventSourceLocations

`func (o *EventForwarderCreateServiceNow) HasTargetsEventSourceLocations() bool`

HasTargetsEventSourceLocations returns a boolean if a field has been set.

### GetToken

`func (o *EventForwarderCreateServiceNow) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EventForwarderCreateServiceNow) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EventForwarderCreateServiceNow) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EventForwarderCreateServiceNow) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EventForwarderCreateServiceNow) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EventForwarderCreateServiceNow) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EventForwarderCreateServiceNow) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EventForwarderCreateServiceNow) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserEmail

`func (o *EventForwarderCreateServiceNow) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *EventForwarderCreateServiceNow) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *EventForwarderCreateServiceNow) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *EventForwarderCreateServiceNow) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


