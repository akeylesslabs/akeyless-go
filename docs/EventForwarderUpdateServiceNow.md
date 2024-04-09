# EventForwarderUpdateServiceNow

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
**Enable** | Pointer to **string** | Enable/Disable Event Forwarder [true/false] | [optional] [default to "true"]
**EventTypes** | Pointer to **[]string** | List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated] | [optional] 
**GatewaysEventSourceLocations** | **[]string** | Event sources | 
**Host** | Pointer to **string** | Workstation Host | [optional] 
**ItemsEventSourceLocations** | Pointer to **[]string** | Items Event sources | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | EventForwarder name | 
**NewName** | Pointer to **string** | New EventForwarder name | [optional] 
**TargetsEventSourceLocations** | Pointer to **[]string** | Targets Event sources | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserEmail** | Pointer to **string** | The user email to identify with when connecting with jwt authentication | [optional] 

## Methods

### NewEventForwarderUpdateServiceNow

`func NewEventForwarderUpdateServiceNow(gatewaysEventSourceLocations []string, name string, ) *EventForwarderUpdateServiceNow`

NewEventForwarderUpdateServiceNow instantiates a new EventForwarderUpdateServiceNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventForwarderUpdateServiceNowWithDefaults

`func NewEventForwarderUpdateServiceNowWithDefaults() *EventForwarderUpdateServiceNow`

NewEventForwarderUpdateServiceNowWithDefaults instantiates a new EventForwarderUpdateServiceNow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminName

`func (o *EventForwarderUpdateServiceNow) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *EventForwarderUpdateServiceNow) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *EventForwarderUpdateServiceNow) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *EventForwarderUpdateServiceNow) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAdminPwd

`func (o *EventForwarderUpdateServiceNow) GetAdminPwd() string`

GetAdminPwd returns the AdminPwd field if non-nil, zero value otherwise.

### GetAdminPwdOk

`func (o *EventForwarderUpdateServiceNow) GetAdminPwdOk() (*string, bool)`

GetAdminPwdOk returns a tuple with the AdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPwd

`func (o *EventForwarderUpdateServiceNow) SetAdminPwd(v string)`

SetAdminPwd sets AdminPwd field to given value.

### HasAdminPwd

`func (o *EventForwarderUpdateServiceNow) HasAdminPwd() bool`

HasAdminPwd returns a boolean if a field has been set.

### GetAppPrivateKeyBase64

`func (o *EventForwarderUpdateServiceNow) GetAppPrivateKeyBase64() string`

GetAppPrivateKeyBase64 returns the AppPrivateKeyBase64 field if non-nil, zero value otherwise.

### GetAppPrivateKeyBase64Ok

`func (o *EventForwarderUpdateServiceNow) GetAppPrivateKeyBase64Ok() (*string, bool)`

GetAppPrivateKeyBase64Ok returns a tuple with the AppPrivateKeyBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPrivateKeyBase64

`func (o *EventForwarderUpdateServiceNow) SetAppPrivateKeyBase64(v string)`

SetAppPrivateKeyBase64 sets AppPrivateKeyBase64 field to given value.

### HasAppPrivateKeyBase64

`func (o *EventForwarderUpdateServiceNow) HasAppPrivateKeyBase64() bool`

HasAppPrivateKeyBase64 returns a boolean if a field has been set.

### GetAuthMethodsEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) GetAuthMethodsEventSourceLocations() []string`

GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field if non-nil, zero value otherwise.

### GetAuthMethodsEventSourceLocationsOk

`func (o *EventForwarderUpdateServiceNow) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool)`

GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodsEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) SetAuthMethodsEventSourceLocations(v []string)`

SetAuthMethodsEventSourceLocations sets AuthMethodsEventSourceLocations field to given value.

### HasAuthMethodsEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) HasAuthMethodsEventSourceLocations() bool`

HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.

### GetAuthType

`func (o *EventForwarderUpdateServiceNow) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *EventForwarderUpdateServiceNow) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *EventForwarderUpdateServiceNow) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *EventForwarderUpdateServiceNow) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *EventForwarderUpdateServiceNow) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EventForwarderUpdateServiceNow) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EventForwarderUpdateServiceNow) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *EventForwarderUpdateServiceNow) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *EventForwarderUpdateServiceNow) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *EventForwarderUpdateServiceNow) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *EventForwarderUpdateServiceNow) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *EventForwarderUpdateServiceNow) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDescription

`func (o *EventForwarderUpdateServiceNow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventForwarderUpdateServiceNow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventForwarderUpdateServiceNow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventForwarderUpdateServiceNow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnable

`func (o *EventForwarderUpdateServiceNow) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *EventForwarderUpdateServiceNow) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *EventForwarderUpdateServiceNow) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *EventForwarderUpdateServiceNow) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventForwarderUpdateServiceNow) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventForwarderUpdateServiceNow) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventForwarderUpdateServiceNow) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *EventForwarderUpdateServiceNow) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetGatewaysEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) GetGatewaysEventSourceLocations() []string`

GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field if non-nil, zero value otherwise.

### GetGatewaysEventSourceLocationsOk

`func (o *EventForwarderUpdateServiceNow) GetGatewaysEventSourceLocationsOk() (*[]string, bool)`

GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) SetGatewaysEventSourceLocations(v []string)`

SetGatewaysEventSourceLocations sets GatewaysEventSourceLocations field to given value.


### GetHost

`func (o *EventForwarderUpdateServiceNow) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EventForwarderUpdateServiceNow) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EventForwarderUpdateServiceNow) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EventForwarderUpdateServiceNow) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetItemsEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) GetItemsEventSourceLocations() []string`

GetItemsEventSourceLocations returns the ItemsEventSourceLocations field if non-nil, zero value otherwise.

### GetItemsEventSourceLocationsOk

`func (o *EventForwarderUpdateServiceNow) GetItemsEventSourceLocationsOk() (*[]string, bool)`

GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) SetItemsEventSourceLocations(v []string)`

SetItemsEventSourceLocations sets ItemsEventSourceLocations field to given value.

### HasItemsEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) HasItemsEventSourceLocations() bool`

HasItemsEventSourceLocations returns a boolean if a field has been set.

### GetJson

`func (o *EventForwarderUpdateServiceNow) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EventForwarderUpdateServiceNow) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EventForwarderUpdateServiceNow) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EventForwarderUpdateServiceNow) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *EventForwarderUpdateServiceNow) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *EventForwarderUpdateServiceNow) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *EventForwarderUpdateServiceNow) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *EventForwarderUpdateServiceNow) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *EventForwarderUpdateServiceNow) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventForwarderUpdateServiceNow) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventForwarderUpdateServiceNow) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EventForwarderUpdateServiceNow) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *EventForwarderUpdateServiceNow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventForwarderUpdateServiceNow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventForwarderUpdateServiceNow) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *EventForwarderUpdateServiceNow) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *EventForwarderUpdateServiceNow) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *EventForwarderUpdateServiceNow) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *EventForwarderUpdateServiceNow) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetTargetsEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) GetTargetsEventSourceLocations() []string`

GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field if non-nil, zero value otherwise.

### GetTargetsEventSourceLocationsOk

`func (o *EventForwarderUpdateServiceNow) GetTargetsEventSourceLocationsOk() (*[]string, bool)`

GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) SetTargetsEventSourceLocations(v []string)`

SetTargetsEventSourceLocations sets TargetsEventSourceLocations field to given value.

### HasTargetsEventSourceLocations

`func (o *EventForwarderUpdateServiceNow) HasTargetsEventSourceLocations() bool`

HasTargetsEventSourceLocations returns a boolean if a field has been set.

### GetToken

`func (o *EventForwarderUpdateServiceNow) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EventForwarderUpdateServiceNow) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EventForwarderUpdateServiceNow) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EventForwarderUpdateServiceNow) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EventForwarderUpdateServiceNow) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EventForwarderUpdateServiceNow) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EventForwarderUpdateServiceNow) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EventForwarderUpdateServiceNow) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserEmail

`func (o *EventForwarderUpdateServiceNow) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *EventForwarderUpdateServiceNow) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *EventForwarderUpdateServiceNow) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *EventForwarderUpdateServiceNow) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


