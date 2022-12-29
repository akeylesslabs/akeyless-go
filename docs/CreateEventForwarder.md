# CreateEventForwarder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminName** | Pointer to **string** | Workstation Admin Name | [optional] 
**AdminPwd** | Pointer to **string** | Workstation Admin password | [optional] 
**Comment** | Pointer to **string** | Comment about the EventForwarder | [optional] 
**EmailTo** | Pointer to **string** | A comma seperated list of email addresses to send event to (relevant only for \\\&quot;email\\\&quot; Event Forwarder) | [optional] 
**EventSourceLocations** | **[]string** | Event sources | 
**EventSourceType** | Pointer to **string** | Event Source type [item, target] | [optional] [default to "item"]
**EventTypes** | Pointer to **[]string** | Event types | [optional] 
**Every** | Pointer to **string** | Rate of periodic runner repetition in hours | [optional] 
**ForwarderType** | **string** |  | 
**Host** | Pointer to **string** | Workstation Host | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | EventForwarder name | 
**RunnerType** | **string** |  | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


