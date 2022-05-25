# GatewayCreateProducerRabbitMQ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Producer name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RabbitmqAdminPwd** | Pointer to **string** | RabbitMQ Admin password | [optional] 
**RabbitmqAdminUser** | Pointer to **string** | RabbitMQ Admin User | [optional] 
**RabbitmqServerUri** | Pointer to **string** | Server URI | [optional] 
**RabbitmqUserConfPermission** | Pointer to **string** | User configuration permission | [optional] 
**RabbitmqUserReadPermission** | Pointer to **string** | User read permission | [optional] 
**RabbitmqUserTags** | Pointer to **string** | User Tags | [optional] 
**RabbitmqUserVhost** | Pointer to **string** | User Virtual Host | [optional] 
**RabbitmqUserWritePermission** | Pointer to **string** | User write permission | [optional] 
**SecureAccessEnable** | Pointer to **string** |  | [optional] 
**SecureAccessUrl** | Pointer to **string** |  | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Secure Access Web Category | [optional] [default to true]
**SecureAccessWebBrowsing** | Pointer to **bool** |  | [optional] 
**SecureAccessWebProxy** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerRabbitMQ

`func NewGatewayCreateProducerRabbitMQ(name string, ) *GatewayCreateProducerRabbitMQ`

NewGatewayCreateProducerRabbitMQ instantiates a new GatewayCreateProducerRabbitMQ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerRabbitMQWithDefaults

`func NewGatewayCreateProducerRabbitMQWithDefaults() *GatewayCreateProducerRabbitMQ`

NewGatewayCreateProducerRabbitMQWithDefaults instantiates a new GatewayCreateProducerRabbitMQ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasRabbitmqAdminPwd

`func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqAdminPwd() bool`

HasRabbitmqAdminPwd returns a boolean if a field has been set.

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

### HasRabbitmqAdminUser

`func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqAdminUser() bool`

HasRabbitmqAdminUser returns a boolean if a field has been set.

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

### HasRabbitmqServerUri

`func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqServerUri() bool`

HasRabbitmqServerUri returns a boolean if a field has been set.

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

### HasRabbitmqUserConfPermission

`func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqUserConfPermission() bool`

HasRabbitmqUserConfPermission returns a boolean if a field has been set.

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

### HasRabbitmqUserReadPermission

`func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqUserReadPermission() bool`

HasRabbitmqUserReadPermission returns a boolean if a field has been set.

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

### HasRabbitmqUserWritePermission

`func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqUserWritePermission() bool`

HasRabbitmqUserWritePermission returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayCreateProducerRabbitMQ) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayCreateProducerRabbitMQ) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessUrl

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessUrl() string`

GetSecureAccessUrl returns the SecureAccessUrl field if non-nil, zero value otherwise.

### GetSecureAccessUrlOk

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessUrlOk() (*string, bool)`

GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUrl

`func (o *GatewayCreateProducerRabbitMQ) SetSecureAccessUrl(v string)`

SetSecureAccessUrl sets SecureAccessUrl field to given value.

### HasSecureAccessUrl

`func (o *GatewayCreateProducerRabbitMQ) HasSecureAccessUrl() bool`

HasSecureAccessUrl returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayCreateProducerRabbitMQ) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayCreateProducerRabbitMQ) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *GatewayCreateProducerRabbitMQ) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *GatewayCreateProducerRabbitMQ) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *GatewayCreateProducerRabbitMQ) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *GatewayCreateProducerRabbitMQ) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *GatewayCreateProducerRabbitMQ) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerRabbitMQ) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerRabbitMQ) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerRabbitMQ) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerRabbitMQ) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerRabbitMQ) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerRabbitMQ) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerRabbitMQ) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerRabbitMQ) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


