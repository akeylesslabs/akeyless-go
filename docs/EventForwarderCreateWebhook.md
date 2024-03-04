# EventForwarderCreateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodsEventSourceLocations** | Pointer to **[]string** | Auth Method Event sources | [optional] 
**AuthToken** | Pointer to **string** | Base64 encoded Token string for authentication type Token | [optional] 
**AuthType** | Pointer to **string** | The Webhook authentication type [user-pass, token, certificate] | [optional] [default to "user-pass"]
**ClientCertData** | Pointer to **string** | Base64 encoded PEM certificate, relevant for certificate auth-type | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EventTypes** | Pointer to **[]string** | List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated] | [optional] 
**Every** | Pointer to **string** | Rate of periodic runner repetition in hours | [optional] 
**GatewaysEventSourceLocations** | **[]string** | Event sources | 
**ItemsEventSourceLocations** | Pointer to **[]string** | Items Event sources | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | EventForwarder name | 
**Password** | Pointer to **string** | Password for authentication relevant for user-pass auth-type | [optional] 
**PrivateKeyData** | Pointer to **string** | Base64 encoded PEM RSA Private Key, relevant for certificate auth-type | [optional] 
**RunnerType** | **string** |  | 
**ServerCertificates** | Pointer to **string** | Base64 encoded PEM certificate of the Webhook | [optional] 
**TargetsEventSourceLocations** | Pointer to **[]string** | Targets Event sources | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | Pointer to **string** | Webhook URL | [optional] 
**Username** | Pointer to **string** | Username for authentication relevant for user-pass auth-type | [optional] 

## Methods

### NewEventForwarderCreateWebhook

`func NewEventForwarderCreateWebhook(gatewaysEventSourceLocations []string, name string, runnerType string, ) *EventForwarderCreateWebhook`

NewEventForwarderCreateWebhook instantiates a new EventForwarderCreateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventForwarderCreateWebhookWithDefaults

`func NewEventForwarderCreateWebhookWithDefaults() *EventForwarderCreateWebhook`

NewEventForwarderCreateWebhookWithDefaults instantiates a new EventForwarderCreateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodsEventSourceLocations

`func (o *EventForwarderCreateWebhook) GetAuthMethodsEventSourceLocations() []string`

GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field if non-nil, zero value otherwise.

### GetAuthMethodsEventSourceLocationsOk

`func (o *EventForwarderCreateWebhook) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool)`

GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodsEventSourceLocations

`func (o *EventForwarderCreateWebhook) SetAuthMethodsEventSourceLocations(v []string)`

SetAuthMethodsEventSourceLocations sets AuthMethodsEventSourceLocations field to given value.

### HasAuthMethodsEventSourceLocations

`func (o *EventForwarderCreateWebhook) HasAuthMethodsEventSourceLocations() bool`

HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.

### GetAuthToken

`func (o *EventForwarderCreateWebhook) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *EventForwarderCreateWebhook) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *EventForwarderCreateWebhook) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *EventForwarderCreateWebhook) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetAuthType

`func (o *EventForwarderCreateWebhook) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *EventForwarderCreateWebhook) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *EventForwarderCreateWebhook) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *EventForwarderCreateWebhook) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientCertData

`func (o *EventForwarderCreateWebhook) GetClientCertData() string`

GetClientCertData returns the ClientCertData field if non-nil, zero value otherwise.

### GetClientCertDataOk

`func (o *EventForwarderCreateWebhook) GetClientCertDataOk() (*string, bool)`

GetClientCertDataOk returns a tuple with the ClientCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertData

`func (o *EventForwarderCreateWebhook) SetClientCertData(v string)`

SetClientCertData sets ClientCertData field to given value.

### HasClientCertData

`func (o *EventForwarderCreateWebhook) HasClientCertData() bool`

HasClientCertData returns a boolean if a field has been set.

### GetDescription

`func (o *EventForwarderCreateWebhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventForwarderCreateWebhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventForwarderCreateWebhook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventForwarderCreateWebhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventForwarderCreateWebhook) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventForwarderCreateWebhook) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventForwarderCreateWebhook) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *EventForwarderCreateWebhook) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetEvery

`func (o *EventForwarderCreateWebhook) GetEvery() string`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *EventForwarderCreateWebhook) GetEveryOk() (*string, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *EventForwarderCreateWebhook) SetEvery(v string)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *EventForwarderCreateWebhook) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### GetGatewaysEventSourceLocations

`func (o *EventForwarderCreateWebhook) GetGatewaysEventSourceLocations() []string`

GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field if non-nil, zero value otherwise.

### GetGatewaysEventSourceLocationsOk

`func (o *EventForwarderCreateWebhook) GetGatewaysEventSourceLocationsOk() (*[]string, bool)`

GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysEventSourceLocations

`func (o *EventForwarderCreateWebhook) SetGatewaysEventSourceLocations(v []string)`

SetGatewaysEventSourceLocations sets GatewaysEventSourceLocations field to given value.


### GetItemsEventSourceLocations

`func (o *EventForwarderCreateWebhook) GetItemsEventSourceLocations() []string`

GetItemsEventSourceLocations returns the ItemsEventSourceLocations field if non-nil, zero value otherwise.

### GetItemsEventSourceLocationsOk

`func (o *EventForwarderCreateWebhook) GetItemsEventSourceLocationsOk() (*[]string, bool)`

GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsEventSourceLocations

`func (o *EventForwarderCreateWebhook) SetItemsEventSourceLocations(v []string)`

SetItemsEventSourceLocations sets ItemsEventSourceLocations field to given value.

### HasItemsEventSourceLocations

`func (o *EventForwarderCreateWebhook) HasItemsEventSourceLocations() bool`

HasItemsEventSourceLocations returns a boolean if a field has been set.

### GetJson

`func (o *EventForwarderCreateWebhook) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EventForwarderCreateWebhook) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EventForwarderCreateWebhook) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EventForwarderCreateWebhook) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *EventForwarderCreateWebhook) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventForwarderCreateWebhook) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventForwarderCreateWebhook) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EventForwarderCreateWebhook) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *EventForwarderCreateWebhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventForwarderCreateWebhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventForwarderCreateWebhook) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *EventForwarderCreateWebhook) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EventForwarderCreateWebhook) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EventForwarderCreateWebhook) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EventForwarderCreateWebhook) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrivateKeyData

`func (o *EventForwarderCreateWebhook) GetPrivateKeyData() string`

GetPrivateKeyData returns the PrivateKeyData field if non-nil, zero value otherwise.

### GetPrivateKeyDataOk

`func (o *EventForwarderCreateWebhook) GetPrivateKeyDataOk() (*string, bool)`

GetPrivateKeyDataOk returns a tuple with the PrivateKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyData

`func (o *EventForwarderCreateWebhook) SetPrivateKeyData(v string)`

SetPrivateKeyData sets PrivateKeyData field to given value.

### HasPrivateKeyData

`func (o *EventForwarderCreateWebhook) HasPrivateKeyData() bool`

HasPrivateKeyData returns a boolean if a field has been set.

### GetRunnerType

`func (o *EventForwarderCreateWebhook) GetRunnerType() string`

GetRunnerType returns the RunnerType field if non-nil, zero value otherwise.

### GetRunnerTypeOk

`func (o *EventForwarderCreateWebhook) GetRunnerTypeOk() (*string, bool)`

GetRunnerTypeOk returns a tuple with the RunnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerType

`func (o *EventForwarderCreateWebhook) SetRunnerType(v string)`

SetRunnerType sets RunnerType field to given value.


### GetServerCertificates

`func (o *EventForwarderCreateWebhook) GetServerCertificates() string`

GetServerCertificates returns the ServerCertificates field if non-nil, zero value otherwise.

### GetServerCertificatesOk

`func (o *EventForwarderCreateWebhook) GetServerCertificatesOk() (*string, bool)`

GetServerCertificatesOk returns a tuple with the ServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificates

`func (o *EventForwarderCreateWebhook) SetServerCertificates(v string)`

SetServerCertificates sets ServerCertificates field to given value.

### HasServerCertificates

`func (o *EventForwarderCreateWebhook) HasServerCertificates() bool`

HasServerCertificates returns a boolean if a field has been set.

### GetTargetsEventSourceLocations

`func (o *EventForwarderCreateWebhook) GetTargetsEventSourceLocations() []string`

GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field if non-nil, zero value otherwise.

### GetTargetsEventSourceLocationsOk

`func (o *EventForwarderCreateWebhook) GetTargetsEventSourceLocationsOk() (*[]string, bool)`

GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsEventSourceLocations

`func (o *EventForwarderCreateWebhook) SetTargetsEventSourceLocations(v []string)`

SetTargetsEventSourceLocations sets TargetsEventSourceLocations field to given value.

### HasTargetsEventSourceLocations

`func (o *EventForwarderCreateWebhook) HasTargetsEventSourceLocations() bool`

HasTargetsEventSourceLocations returns a boolean if a field has been set.

### GetToken

`func (o *EventForwarderCreateWebhook) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EventForwarderCreateWebhook) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EventForwarderCreateWebhook) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EventForwarderCreateWebhook) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EventForwarderCreateWebhook) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EventForwarderCreateWebhook) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EventForwarderCreateWebhook) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EventForwarderCreateWebhook) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *EventForwarderCreateWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventForwarderCreateWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventForwarderCreateWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventForwarderCreateWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *EventForwarderCreateWebhook) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EventForwarderCreateWebhook) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EventForwarderCreateWebhook) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EventForwarderCreateWebhook) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


