# GatewayCreateProducerRabbitMQ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**Name** | **string** | Producer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RabbitmqAdminPwd** | **string** | RabbitMQ Admin password | 
**RabbitmqAdminUser** | **string** | RabbitMQ Admin User | 
**RabbitmqServerUri** | **string** | Server URI | 
**RabbitmqUserConfPermission** | **string** | User configuration permission | 
**RabbitmqUserReadPermission** | **string** | User read permission | 
**RabbitmqUserTags** | Pointer to **string** | User Tags | [optional] 
**RabbitmqUserVhost** | Pointer to **string** | User Virtual Host | [optional] 
**RabbitmqUserWritePermission** | **string** | User write permission | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayCreateProducerRabbitMQ

`func NewGatewayCreateProducerRabbitMQ(name string, rabbitmqAdminPwd string, rabbitmqAdminUser string, rabbitmqServerUri string, rabbitmqUserConfPermission string, rabbitmqUserReadPermission string, rabbitmqUserWritePermission string, ) *GatewayCreateProducerRabbitMQ`

NewGatewayCreateProducerRabbitMQ instantiates a new GatewayCreateProducerRabbitMQ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerRabbitMQWithDefaults

`func NewGatewayCreateProducerRabbitMQWithDefaults() *GatewayCreateProducerRabbitMQ`

NewGatewayCreateProducerRabbitMQWithDefaults instantiates a new GatewayCreateProducerRabbitMQ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayUrl

`func (o *GatewayCreateProducerRabbitMQ) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GatewayCreateProducerRabbitMQ) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GatewayCreateProducerRabbitMQ) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *GatewayCreateProducerRabbitMQ) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerRabbitMQ) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerRabbitMQ) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerRabbitMQ) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *GatewayCreateProducerRabbitMQ) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateProducerRabbitMQ) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateProducerRabbitMQ) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateProducerRabbitMQ) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerRabbitMQ) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerRabbitMQ) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerRabbitMQ) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerRabbitMQ) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRabbitmqAdminPwd

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqAdminPwd() string`

GetRabbitmqAdminPwd returns the RabbitmqAdminPwd field if non-nil, zero value otherwise.

### GetRabbitmqAdminPwdOk

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqAdminPwdOk() (*string, bool)`

GetRabbitmqAdminPwdOk returns a tuple with the RabbitmqAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqAdminPwd

`func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqAdminPwd(v string)`

SetRabbitmqAdminPwd sets RabbitmqAdminPwd field to given value.


### GetRabbitmqAdminUser

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqAdminUser() string`

GetRabbitmqAdminUser returns the RabbitmqAdminUser field if non-nil, zero value otherwise.

### GetRabbitmqAdminUserOk

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqAdminUserOk() (*string, bool)`

GetRabbitmqAdminUserOk returns a tuple with the RabbitmqAdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqAdminUser

`func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqAdminUser(v string)`

SetRabbitmqAdminUser sets RabbitmqAdminUser field to given value.


### GetRabbitmqServerUri

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqServerUri() string`

GetRabbitmqServerUri returns the RabbitmqServerUri field if non-nil, zero value otherwise.

### GetRabbitmqServerUriOk

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqServerUriOk() (*string, bool)`

GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUri

`func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqServerUri(v string)`

SetRabbitmqServerUri sets RabbitmqServerUri field to given value.


### GetRabbitmqUserConfPermission

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserConfPermission() string`

GetRabbitmqUserConfPermission returns the RabbitmqUserConfPermission field if non-nil, zero value otherwise.

### GetRabbitmqUserConfPermissionOk

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserConfPermissionOk() (*string, bool)`

GetRabbitmqUserConfPermissionOk returns a tuple with the RabbitmqUserConfPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserConfPermission

`func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserConfPermission(v string)`

SetRabbitmqUserConfPermission sets RabbitmqUserConfPermission field to given value.


### GetRabbitmqUserReadPermission

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserReadPermission() string`

GetRabbitmqUserReadPermission returns the RabbitmqUserReadPermission field if non-nil, zero value otherwise.

### GetRabbitmqUserReadPermissionOk

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserReadPermissionOk() (*string, bool)`

GetRabbitmqUserReadPermissionOk returns a tuple with the RabbitmqUserReadPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserReadPermission

`func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserReadPermission(v string)`

SetRabbitmqUserReadPermission sets RabbitmqUserReadPermission field to given value.


### GetRabbitmqUserTags

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserTags() string`

GetRabbitmqUserTags returns the RabbitmqUserTags field if non-nil, zero value otherwise.

### GetRabbitmqUserTagsOk

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserTagsOk() (*string, bool)`

GetRabbitmqUserTagsOk returns a tuple with the RabbitmqUserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserTags

`func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserTags(v string)`

SetRabbitmqUserTags sets RabbitmqUserTags field to given value.

### HasRabbitmqUserTags

`func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqUserTags() bool`

HasRabbitmqUserTags returns a boolean if a field has been set.

### GetRabbitmqUserVhost

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserVhost() string`

GetRabbitmqUserVhost returns the RabbitmqUserVhost field if non-nil, zero value otherwise.

### GetRabbitmqUserVhostOk

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserVhostOk() (*string, bool)`

GetRabbitmqUserVhostOk returns a tuple with the RabbitmqUserVhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserVhost

`func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserVhost(v string)`

SetRabbitmqUserVhost sets RabbitmqUserVhost field to given value.

### HasRabbitmqUserVhost

`func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqUserVhost() bool`

HasRabbitmqUserVhost returns a boolean if a field has been set.

### GetRabbitmqUserWritePermission

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserWritePermission() string`

GetRabbitmqUserWritePermission returns the RabbitmqUserWritePermission field if non-nil, zero value otherwise.

### GetRabbitmqUserWritePermissionOk

`func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserWritePermissionOk() (*string, bool)`

GetRabbitmqUserWritePermissionOk returns a tuple with the RabbitmqUserWritePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserWritePermission

`func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserWritePermission(v string)`

SetRabbitmqUserWritePermission sets RabbitmqUserWritePermission field to given value.


### GetToken

`func (o *GatewayCreateProducerRabbitMQ) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerRabbitMQ) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerRabbitMQ) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerRabbitMQ) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerRabbitMQ) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerRabbitMQ) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerRabbitMQ) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerRabbitMQ) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerRabbitMQ) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerRabbitMQ) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerRabbitMQ) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerRabbitMQ) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayCreateProducerRabbitMQ) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateProducerRabbitMQ) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateProducerRabbitMQ) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateProducerRabbitMQ) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


