# CreateEventForwarder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminName** | Pointer to **string** | Workstation Admin Name | [optional] 
**AdminPwd** | Pointer to **string** | Workstation Admin password | [optional] 
**AppPrivateKeyBase64** | Pointer to **string** | The RSA Private Key PEM formatted in base64 to use when connecting to ServiceNow with jwt authentication | [optional] 
**AuthType** | Pointer to **string** | The authentication type to use when connecting to ServiceNow (user-pass / jwt) | [optional] [default to "user-pass"]
**ClientId** | Pointer to **string** | The client ID to use when connecting to ServiceNow with jwt authentication | [optional] 
**ClientSecret** | Pointer to **string** | The client secret to use when connecting to ServiceNow with jwt authentication | [optional] 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EmailTo** | Pointer to **string** | A comma seperated list of email addresses to send event to (relevant only for \&quot;email\&quot; Event Forwarder) | [optional] 
**EventSourceLocations** | **[]string** | Event sources | 
**EventSourceType** | Pointer to **string** | Event Source type [item, target, auth_method] | [optional] [default to "item"]
**EventTypes** | Pointer to **[]string** | List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, auth-method-pending-expiration, auth-method-expired, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure] | [optional] 
**Every** | Pointer to **string** | Rate of periodic runner repetition in hours | [optional] 
**ForwarderType** | **string** |  | 
**Host** | Pointer to **string** | Workstation Host | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | EventForwarder name | 
**RunnerType** | **string** |  | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserEmail** | Pointer to **string** | The user email to use when connecting to ServiceNow with jwt authentication | [optional] 

## Methods

### NewCreateEventForwarder

`func NewCreateEventForwarder(eventSourceLocations []string, forwarderType string, name string, runnerType string, ) *CreateEventForwarder`

NewCreateEventForwarder instantiates a new CreateEventForwarder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEventForwarderWithDefaults

`func NewCreateEventForwarderWithDefaults() *CreateEventForwarder`

NewCreateEventForwarderWithDefaults instantiates a new CreateEventForwarder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminName

`func (o *CreateEventForwarder) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *CreateEventForwarder) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *CreateEventForwarder) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *CreateEventForwarder) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAdminPwd

`func (o *CreateEventForwarder) GetAdminPwd() string`

GetAdminPwd returns the AdminPwd field if non-nil, zero value otherwise.

### GetAdminPwdOk

`func (o *CreateEventForwarder) GetAdminPwdOk() (*string, bool)`

GetAdminPwdOk returns a tuple with the AdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPwd

`func (o *CreateEventForwarder) SetAdminPwd(v string)`

SetAdminPwd sets AdminPwd field to given value.

### HasAdminPwd

`func (o *CreateEventForwarder) HasAdminPwd() bool`

HasAdminPwd returns a boolean if a field has been set.

### GetAppPrivateKeyBase64

`func (o *CreateEventForwarder) GetAppPrivateKeyBase64() string`

GetAppPrivateKeyBase64 returns the AppPrivateKeyBase64 field if non-nil, zero value otherwise.

### GetAppPrivateKeyBase64Ok

`func (o *CreateEventForwarder) GetAppPrivateKeyBase64Ok() (*string, bool)`

GetAppPrivateKeyBase64Ok returns a tuple with the AppPrivateKeyBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPrivateKeyBase64

`func (o *CreateEventForwarder) SetAppPrivateKeyBase64(v string)`

SetAppPrivateKeyBase64 sets AppPrivateKeyBase64 field to given value.

### HasAppPrivateKeyBase64

`func (o *CreateEventForwarder) HasAppPrivateKeyBase64() bool`

HasAppPrivateKeyBase64 returns a boolean if a field has been set.

### GetAuthType

`func (o *CreateEventForwarder) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateEventForwarder) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateEventForwarder) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *CreateEventForwarder) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *CreateEventForwarder) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateEventForwarder) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateEventForwarder) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateEventForwarder) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *CreateEventForwarder) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateEventForwarder) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateEventForwarder) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CreateEventForwarder) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetComment

`func (o *CreateEventForwarder) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateEventForwarder) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateEventForwarder) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateEventForwarder) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *CreateEventForwarder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEventForwarder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEventForwarder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEventForwarder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmailTo

`func (o *CreateEventForwarder) GetEmailTo() string`

GetEmailTo returns the EmailTo field if non-nil, zero value otherwise.

### GetEmailToOk

`func (o *CreateEventForwarder) GetEmailToOk() (*string, bool)`

GetEmailToOk returns a tuple with the EmailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTo

`func (o *CreateEventForwarder) SetEmailTo(v string)`

SetEmailTo sets EmailTo field to given value.

### HasEmailTo

`func (o *CreateEventForwarder) HasEmailTo() bool`

HasEmailTo returns a boolean if a field has been set.

### GetEventSourceLocations

`func (o *CreateEventForwarder) GetEventSourceLocations() []string`

GetEventSourceLocations returns the EventSourceLocations field if non-nil, zero value otherwise.

### GetEventSourceLocationsOk

`func (o *CreateEventForwarder) GetEventSourceLocationsOk() (*[]string, bool)`

GetEventSourceLocationsOk returns a tuple with the EventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSourceLocations

`func (o *CreateEventForwarder) SetEventSourceLocations(v []string)`

SetEventSourceLocations sets EventSourceLocations field to given value.


### GetEventSourceType

`func (o *CreateEventForwarder) GetEventSourceType() string`

GetEventSourceType returns the EventSourceType field if non-nil, zero value otherwise.

### GetEventSourceTypeOk

`func (o *CreateEventForwarder) GetEventSourceTypeOk() (*string, bool)`

GetEventSourceTypeOk returns a tuple with the EventSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSourceType

`func (o *CreateEventForwarder) SetEventSourceType(v string)`

SetEventSourceType sets EventSourceType field to given value.

### HasEventSourceType

`func (o *CreateEventForwarder) HasEventSourceType() bool`

HasEventSourceType returns a boolean if a field has been set.

### GetEventTypes

`func (o *CreateEventForwarder) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *CreateEventForwarder) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *CreateEventForwarder) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *CreateEventForwarder) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetEvery

`func (o *CreateEventForwarder) GetEvery() string`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *CreateEventForwarder) GetEveryOk() (*string, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *CreateEventForwarder) SetEvery(v string)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *CreateEventForwarder) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### GetForwarderType

`func (o *CreateEventForwarder) GetForwarderType() string`

GetForwarderType returns the ForwarderType field if non-nil, zero value otherwise.

### GetForwarderTypeOk

`func (o *CreateEventForwarder) GetForwarderTypeOk() (*string, bool)`

GetForwarderTypeOk returns a tuple with the ForwarderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarderType

`func (o *CreateEventForwarder) SetForwarderType(v string)`

SetForwarderType sets ForwarderType field to given value.


### GetHost

`func (o *CreateEventForwarder) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateEventForwarder) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateEventForwarder) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateEventForwarder) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *CreateEventForwarder) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateEventForwarder) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateEventForwarder) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateEventForwarder) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *CreateEventForwarder) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateEventForwarder) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateEventForwarder) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateEventForwarder) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateEventForwarder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEventForwarder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEventForwarder) SetName(v string)`

SetName sets Name field to given value.


### GetRunnerType

`func (o *CreateEventForwarder) GetRunnerType() string`

GetRunnerType returns the RunnerType field if non-nil, zero value otherwise.

### GetRunnerTypeOk

`func (o *CreateEventForwarder) GetRunnerTypeOk() (*string, bool)`

GetRunnerTypeOk returns a tuple with the RunnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerType

`func (o *CreateEventForwarder) SetRunnerType(v string)`

SetRunnerType sets RunnerType field to given value.


### GetToken

`func (o *CreateEventForwarder) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateEventForwarder) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateEventForwarder) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateEventForwarder) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateEventForwarder) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateEventForwarder) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateEventForwarder) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateEventForwarder) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserEmail

`func (o *CreateEventForwarder) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *CreateEventForwarder) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *CreateEventForwarder) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *CreateEventForwarder) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


