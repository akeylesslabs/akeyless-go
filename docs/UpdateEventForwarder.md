# UpdateEventForwarder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminName** | Pointer to **string** | Workstation Admin Name | [optional] 
**AuthType** | Pointer to **string** | The authentication type to use when connecting to ServiceNow (user-pass / jwt) | [optional] [default to "user-pass"]
**ClientId** | Pointer to **string** | The client ID to use when connecting to ServiceNow with jwt authentication | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_comment"]
**EmailTo** | Pointer to **string** | A comma seperated list of email addresses to send event to (relevant only for \&quot;email\&quot; Event Forwarder) | [optional] 
**Enable** | Pointer to **string** | Enable/Disable Event Forwarder [true/false] | [optional] [default to "true"]
**EventSourceLocations** | Pointer to **[]string** | Event sources | [optional] 
**EventTypes** | Pointer to **[]string** | Event types | [optional] 
**Host** | Pointer to **string** | Workstation Host | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | EventForwarder name | 
**NewComment** | Pointer to **string** | Deprecated - use description | [optional] [default to "default_comment"]
**NewName** | Pointer to **string** | New EventForwarder name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserEmail** | Pointer to **string** | The user email to use when connecting to ServiceNow with jwt authentication | [optional] 

## Methods

### NewUpdateEventForwarder

`func NewUpdateEventForwarder(name string, ) *UpdateEventForwarder`

NewUpdateEventForwarder instantiates a new UpdateEventForwarder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEventForwarderWithDefaults

`func NewUpdateEventForwarderWithDefaults() *UpdateEventForwarder`

NewUpdateEventForwarderWithDefaults instantiates a new UpdateEventForwarder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminName

`func (o *UpdateEventForwarder) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *UpdateEventForwarder) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *UpdateEventForwarder) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *UpdateEventForwarder) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAuthType

`func (o *UpdateEventForwarder) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UpdateEventForwarder) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UpdateEventForwarder) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *UpdateEventForwarder) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *UpdateEventForwarder) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UpdateEventForwarder) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UpdateEventForwarder) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UpdateEventForwarder) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateEventForwarder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateEventForwarder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateEventForwarder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateEventForwarder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmailTo

`func (o *UpdateEventForwarder) GetEmailTo() string`

GetEmailTo returns the EmailTo field if non-nil, zero value otherwise.

### GetEmailToOk

`func (o *UpdateEventForwarder) GetEmailToOk() (*string, bool)`

GetEmailToOk returns a tuple with the EmailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTo

`func (o *UpdateEventForwarder) SetEmailTo(v string)`

SetEmailTo sets EmailTo field to given value.

### HasEmailTo

`func (o *UpdateEventForwarder) HasEmailTo() bool`

HasEmailTo returns a boolean if a field has been set.

### GetEnable

`func (o *UpdateEventForwarder) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UpdateEventForwarder) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UpdateEventForwarder) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *UpdateEventForwarder) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEventSourceLocations

`func (o *UpdateEventForwarder) GetEventSourceLocations() []string`

GetEventSourceLocations returns the EventSourceLocations field if non-nil, zero value otherwise.

### GetEventSourceLocationsOk

`func (o *UpdateEventForwarder) GetEventSourceLocationsOk() (*[]string, bool)`

GetEventSourceLocationsOk returns a tuple with the EventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSourceLocations

`func (o *UpdateEventForwarder) SetEventSourceLocations(v []string)`

SetEventSourceLocations sets EventSourceLocations field to given value.

### HasEventSourceLocations

`func (o *UpdateEventForwarder) HasEventSourceLocations() bool`

HasEventSourceLocations returns a boolean if a field has been set.

### GetEventTypes

`func (o *UpdateEventForwarder) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *UpdateEventForwarder) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *UpdateEventForwarder) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *UpdateEventForwarder) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetHost

`func (o *UpdateEventForwarder) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateEventForwarder) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateEventForwarder) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateEventForwarder) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *UpdateEventForwarder) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateEventForwarder) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateEventForwarder) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateEventForwarder) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *UpdateEventForwarder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEventForwarder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEventForwarder) SetName(v string)`

SetName sets Name field to given value.


### GetNewComment

`func (o *UpdateEventForwarder) GetNewComment() string`

GetNewComment returns the NewComment field if non-nil, zero value otherwise.

### GetNewCommentOk

`func (o *UpdateEventForwarder) GetNewCommentOk() (*string, bool)`

GetNewCommentOk returns a tuple with the NewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewComment

`func (o *UpdateEventForwarder) SetNewComment(v string)`

SetNewComment sets NewComment field to given value.

### HasNewComment

`func (o *UpdateEventForwarder) HasNewComment() bool`

HasNewComment returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateEventForwarder) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateEventForwarder) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateEventForwarder) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateEventForwarder) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateEventForwarder) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateEventForwarder) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateEventForwarder) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateEventForwarder) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateEventForwarder) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateEventForwarder) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateEventForwarder) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateEventForwarder) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserEmail

`func (o *UpdateEventForwarder) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *UpdateEventForwarder) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *UpdateEventForwarder) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *UpdateEventForwarder) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


