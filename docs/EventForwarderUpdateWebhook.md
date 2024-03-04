# EventForwarderUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodsEventSourceLocations** | Pointer to **[]string** | Auth Method Event sources | [optional] 
**AuthToken** | Pointer to **string** | Base64 encoded Token string for authentication type Token | [optional] 
**AuthType** | Pointer to **string** | The Webhook authentication type [user-pass, token, certificate] | [optional] [default to "user-pass"]
**ClientCertData** | Pointer to **string** | Base64 encoded PEM certificate, relevant for certificate auth-type | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Enable** | Pointer to **string** | Enable/Disable Event Forwarder [true/false] | [optional] [default to "true"]
**EventTypes** | Pointer to **[]string** | List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated] | [optional] 
**GatewaysEventSourceLocations** | **[]string** | Event sources | 
**ItemsEventSourceLocations** | Pointer to **[]string** | Items Event sources | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | EventForwarder name | 
**NewName** | Pointer to **string** | New EventForwarder name | [optional] 
**Password** | Pointer to **string** | Password for authentication relevant for user-pass auth-type | [optional] 
**PrivateKeyData** | Pointer to **string** | Base64 encoded PEM RSA Private Key, relevant for certificate auth-type | [optional] 
**ServerCertificates** | Pointer to **string** | Base64 encoded PEM certificate of the Webhook | [optional] 
**TargetsEventSourceLocations** | Pointer to **[]string** | Targets Event sources | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | Pointer to **string** | Webhook URL | [optional] 
**Username** | Pointer to **string** | Username for authentication relevant for user-pass auth-type | [optional] 

## Methods

### NewEventForwarderUpdateWebhook

`func NewEventForwarderUpdateWebhook(gatewaysEventSourceLocations []string, name string, ) *EventForwarderUpdateWebhook`

NewEventForwarderUpdateWebhook instantiates a new EventForwarderUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventForwarderUpdateWebhookWithDefaults

`func NewEventForwarderUpdateWebhookWithDefaults() *EventForwarderUpdateWebhook`

NewEventForwarderUpdateWebhookWithDefaults instantiates a new EventForwarderUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodsEventSourceLocations

`func (o *EventForwarderUpdateWebhook) GetAuthMethodsEventSourceLocations() []string`

GetAuthMethodsEventSourceLocations returns the AuthMethodsEventSourceLocations field if non-nil, zero value otherwise.

### GetAuthMethodsEventSourceLocationsOk

`func (o *EventForwarderUpdateWebhook) GetAuthMethodsEventSourceLocationsOk() (*[]string, bool)`

GetAuthMethodsEventSourceLocationsOk returns a tuple with the AuthMethodsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodsEventSourceLocations

`func (o *EventForwarderUpdateWebhook) SetAuthMethodsEventSourceLocations(v []string)`

SetAuthMethodsEventSourceLocations sets AuthMethodsEventSourceLocations field to given value.

### HasAuthMethodsEventSourceLocations

`func (o *EventForwarderUpdateWebhook) HasAuthMethodsEventSourceLocations() bool`

HasAuthMethodsEventSourceLocations returns a boolean if a field has been set.

### GetAuthToken

`func (o *EventForwarderUpdateWebhook) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *EventForwarderUpdateWebhook) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *EventForwarderUpdateWebhook) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *EventForwarderUpdateWebhook) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetAuthType

`func (o *EventForwarderUpdateWebhook) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *EventForwarderUpdateWebhook) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *EventForwarderUpdateWebhook) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *EventForwarderUpdateWebhook) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientCertData

`func (o *EventForwarderUpdateWebhook) GetClientCertData() string`

GetClientCertData returns the ClientCertData field if non-nil, zero value otherwise.

### GetClientCertDataOk

`func (o *EventForwarderUpdateWebhook) GetClientCertDataOk() (*string, bool)`

GetClientCertDataOk returns a tuple with the ClientCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertData

`func (o *EventForwarderUpdateWebhook) SetClientCertData(v string)`

SetClientCertData sets ClientCertData field to given value.

### HasClientCertData

`func (o *EventForwarderUpdateWebhook) HasClientCertData() bool`

HasClientCertData returns a boolean if a field has been set.

### GetDescription

`func (o *EventForwarderUpdateWebhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventForwarderUpdateWebhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventForwarderUpdateWebhook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventForwarderUpdateWebhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnable

`func (o *EventForwarderUpdateWebhook) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *EventForwarderUpdateWebhook) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *EventForwarderUpdateWebhook) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *EventForwarderUpdateWebhook) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventForwarderUpdateWebhook) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventForwarderUpdateWebhook) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventForwarderUpdateWebhook) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *EventForwarderUpdateWebhook) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetGatewaysEventSourceLocations

`func (o *EventForwarderUpdateWebhook) GetGatewaysEventSourceLocations() []string`

GetGatewaysEventSourceLocations returns the GatewaysEventSourceLocations field if non-nil, zero value otherwise.

### GetGatewaysEventSourceLocationsOk

`func (o *EventForwarderUpdateWebhook) GetGatewaysEventSourceLocationsOk() (*[]string, bool)`

GetGatewaysEventSourceLocationsOk returns a tuple with the GatewaysEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysEventSourceLocations

`func (o *EventForwarderUpdateWebhook) SetGatewaysEventSourceLocations(v []string)`

SetGatewaysEventSourceLocations sets GatewaysEventSourceLocations field to given value.


### GetItemsEventSourceLocations

`func (o *EventForwarderUpdateWebhook) GetItemsEventSourceLocations() []string`

GetItemsEventSourceLocations returns the ItemsEventSourceLocations field if non-nil, zero value otherwise.

### GetItemsEventSourceLocationsOk

`func (o *EventForwarderUpdateWebhook) GetItemsEventSourceLocationsOk() (*[]string, bool)`

GetItemsEventSourceLocationsOk returns a tuple with the ItemsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsEventSourceLocations

`func (o *EventForwarderUpdateWebhook) SetItemsEventSourceLocations(v []string)`

SetItemsEventSourceLocations sets ItemsEventSourceLocations field to given value.

### HasItemsEventSourceLocations

`func (o *EventForwarderUpdateWebhook) HasItemsEventSourceLocations() bool`

HasItemsEventSourceLocations returns a boolean if a field has been set.

### GetJson

`func (o *EventForwarderUpdateWebhook) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EventForwarderUpdateWebhook) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EventForwarderUpdateWebhook) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EventForwarderUpdateWebhook) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *EventForwarderUpdateWebhook) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *EventForwarderUpdateWebhook) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *EventForwarderUpdateWebhook) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *EventForwarderUpdateWebhook) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *EventForwarderUpdateWebhook) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventForwarderUpdateWebhook) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventForwarderUpdateWebhook) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EventForwarderUpdateWebhook) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *EventForwarderUpdateWebhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventForwarderUpdateWebhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventForwarderUpdateWebhook) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *EventForwarderUpdateWebhook) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *EventForwarderUpdateWebhook) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *EventForwarderUpdateWebhook) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *EventForwarderUpdateWebhook) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *EventForwarderUpdateWebhook) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EventForwarderUpdateWebhook) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EventForwarderUpdateWebhook) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EventForwarderUpdateWebhook) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrivateKeyData

`func (o *EventForwarderUpdateWebhook) GetPrivateKeyData() string`

GetPrivateKeyData returns the PrivateKeyData field if non-nil, zero value otherwise.

### GetPrivateKeyDataOk

`func (o *EventForwarderUpdateWebhook) GetPrivateKeyDataOk() (*string, bool)`

GetPrivateKeyDataOk returns a tuple with the PrivateKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyData

`func (o *EventForwarderUpdateWebhook) SetPrivateKeyData(v string)`

SetPrivateKeyData sets PrivateKeyData field to given value.

### HasPrivateKeyData

`func (o *EventForwarderUpdateWebhook) HasPrivateKeyData() bool`

HasPrivateKeyData returns a boolean if a field has been set.

### GetServerCertificates

`func (o *EventForwarderUpdateWebhook) GetServerCertificates() string`

GetServerCertificates returns the ServerCertificates field if non-nil, zero value otherwise.

### GetServerCertificatesOk

`func (o *EventForwarderUpdateWebhook) GetServerCertificatesOk() (*string, bool)`

GetServerCertificatesOk returns a tuple with the ServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificates

`func (o *EventForwarderUpdateWebhook) SetServerCertificates(v string)`

SetServerCertificates sets ServerCertificates field to given value.

### HasServerCertificates

`func (o *EventForwarderUpdateWebhook) HasServerCertificates() bool`

HasServerCertificates returns a boolean if a field has been set.

### GetTargetsEventSourceLocations

`func (o *EventForwarderUpdateWebhook) GetTargetsEventSourceLocations() []string`

GetTargetsEventSourceLocations returns the TargetsEventSourceLocations field if non-nil, zero value otherwise.

### GetTargetsEventSourceLocationsOk

`func (o *EventForwarderUpdateWebhook) GetTargetsEventSourceLocationsOk() (*[]string, bool)`

GetTargetsEventSourceLocationsOk returns a tuple with the TargetsEventSourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsEventSourceLocations

`func (o *EventForwarderUpdateWebhook) SetTargetsEventSourceLocations(v []string)`

SetTargetsEventSourceLocations sets TargetsEventSourceLocations field to given value.

### HasTargetsEventSourceLocations

`func (o *EventForwarderUpdateWebhook) HasTargetsEventSourceLocations() bool`

HasTargetsEventSourceLocations returns a boolean if a field has been set.

### GetToken

`func (o *EventForwarderUpdateWebhook) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EventForwarderUpdateWebhook) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EventForwarderUpdateWebhook) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EventForwarderUpdateWebhook) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EventForwarderUpdateWebhook) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EventForwarderUpdateWebhook) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EventForwarderUpdateWebhook) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EventForwarderUpdateWebhook) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *EventForwarderUpdateWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventForwarderUpdateWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventForwarderUpdateWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventForwarderUpdateWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *EventForwarderUpdateWebhook) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EventForwarderUpdateWebhook) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EventForwarderUpdateWebhook) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EventForwarderUpdateWebhook) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


