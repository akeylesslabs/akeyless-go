# DynamicSecretCreateRabbitMq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RabbitmqAdminPwd** | Pointer to **string** | RabbitMQ Admin password | [optional] 
**RabbitmqAdminUser** | Pointer to **string** | RabbitMQ Admin User | [optional] 
**RabbitmqServerUri** | Pointer to **string** | Server URI | [optional] 
**RabbitmqUserConfPermission** | Pointer to **string** | User configuration permission | [optional] 
**RabbitmqUserReadPermission** | Pointer to **string** | User read permission | [optional] 
**RabbitmqUserTags** | Pointer to **string** | User Tags | [optional] 
**RabbitmqUserVhost** | Pointer to **string** | User Virtual Host | [optional] 
**RabbitmqUserWritePermission** | Pointer to **string** | User write permission | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessUrl** | Pointer to **string** | Destination URL to inject secrets | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to true]
**SecureAccessWebBrowsing** | Pointer to **bool** | Secure browser via Akeyless Web Access Bastion | [optional] [default to false]
**SecureAccessWebProxy** | Pointer to **bool** | Web-Proxy via Akeyless Web Access Bastion | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateRabbitMq

`func NewDynamicSecretCreateRabbitMq(name string, ) *DynamicSecretCreateRabbitMq`

NewDynamicSecretCreateRabbitMq instantiates a new DynamicSecretCreateRabbitMq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateRabbitMqWithDefaults

`func NewDynamicSecretCreateRabbitMqWithDefaults() *DynamicSecretCreateRabbitMq`

NewDynamicSecretCreateRabbitMqWithDefaults instantiates a new DynamicSecretCreateRabbitMq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretCreateRabbitMq) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateRabbitMq) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateRabbitMq) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateRabbitMq) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateRabbitMq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateRabbitMq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateRabbitMq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateRabbitMq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateRabbitMq) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateRabbitMq) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateRabbitMq) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateRabbitMq) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateRabbitMq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateRabbitMq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateRabbitMq) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *DynamicSecretCreateRabbitMq) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretCreateRabbitMq) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretCreateRabbitMq) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretCreateRabbitMq) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateRabbitMq) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateRabbitMq) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateRabbitMq) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateRabbitMq) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRabbitmqAdminPwd

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqAdminPwd() string`

GetRabbitmqAdminPwd returns the RabbitmqAdminPwd field if non-nil, zero value otherwise.

### GetRabbitmqAdminPwdOk

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqAdminPwdOk() (*string, bool)`

GetRabbitmqAdminPwdOk returns a tuple with the RabbitmqAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqAdminPwd

`func (o *DynamicSecretCreateRabbitMq) SetRabbitmqAdminPwd(v string)`

SetRabbitmqAdminPwd sets RabbitmqAdminPwd field to given value.

### HasRabbitmqAdminPwd

`func (o *DynamicSecretCreateRabbitMq) HasRabbitmqAdminPwd() bool`

HasRabbitmqAdminPwd returns a boolean if a field has been set.

### GetRabbitmqAdminUser

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqAdminUser() string`

GetRabbitmqAdminUser returns the RabbitmqAdminUser field if non-nil, zero value otherwise.

### GetRabbitmqAdminUserOk

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqAdminUserOk() (*string, bool)`

GetRabbitmqAdminUserOk returns a tuple with the RabbitmqAdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqAdminUser

`func (o *DynamicSecretCreateRabbitMq) SetRabbitmqAdminUser(v string)`

SetRabbitmqAdminUser sets RabbitmqAdminUser field to given value.

### HasRabbitmqAdminUser

`func (o *DynamicSecretCreateRabbitMq) HasRabbitmqAdminUser() bool`

HasRabbitmqAdminUser returns a boolean if a field has been set.

### GetRabbitmqServerUri

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqServerUri() string`

GetRabbitmqServerUri returns the RabbitmqServerUri field if non-nil, zero value otherwise.

### GetRabbitmqServerUriOk

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqServerUriOk() (*string, bool)`

GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUri

`func (o *DynamicSecretCreateRabbitMq) SetRabbitmqServerUri(v string)`

SetRabbitmqServerUri sets RabbitmqServerUri field to given value.

### HasRabbitmqServerUri

`func (o *DynamicSecretCreateRabbitMq) HasRabbitmqServerUri() bool`

HasRabbitmqServerUri returns a boolean if a field has been set.

### GetRabbitmqUserConfPermission

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserConfPermission() string`

GetRabbitmqUserConfPermission returns the RabbitmqUserConfPermission field if non-nil, zero value otherwise.

### GetRabbitmqUserConfPermissionOk

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserConfPermissionOk() (*string, bool)`

GetRabbitmqUserConfPermissionOk returns a tuple with the RabbitmqUserConfPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserConfPermission

`func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserConfPermission(v string)`

SetRabbitmqUserConfPermission sets RabbitmqUserConfPermission field to given value.

### HasRabbitmqUserConfPermission

`func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserConfPermission() bool`

HasRabbitmqUserConfPermission returns a boolean if a field has been set.

### GetRabbitmqUserReadPermission

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserReadPermission() string`

GetRabbitmqUserReadPermission returns the RabbitmqUserReadPermission field if non-nil, zero value otherwise.

### GetRabbitmqUserReadPermissionOk

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserReadPermissionOk() (*string, bool)`

GetRabbitmqUserReadPermissionOk returns a tuple with the RabbitmqUserReadPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserReadPermission

`func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserReadPermission(v string)`

SetRabbitmqUserReadPermission sets RabbitmqUserReadPermission field to given value.

### HasRabbitmqUserReadPermission

`func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserReadPermission() bool`

HasRabbitmqUserReadPermission returns a boolean if a field has been set.

### GetRabbitmqUserTags

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserTags() string`

GetRabbitmqUserTags returns the RabbitmqUserTags field if non-nil, zero value otherwise.

### GetRabbitmqUserTagsOk

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserTagsOk() (*string, bool)`

GetRabbitmqUserTagsOk returns a tuple with the RabbitmqUserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserTags

`func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserTags(v string)`

SetRabbitmqUserTags sets RabbitmqUserTags field to given value.

### HasRabbitmqUserTags

`func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserTags() bool`

HasRabbitmqUserTags returns a boolean if a field has been set.

### GetRabbitmqUserVhost

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserVhost() string`

GetRabbitmqUserVhost returns the RabbitmqUserVhost field if non-nil, zero value otherwise.

### GetRabbitmqUserVhostOk

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserVhostOk() (*string, bool)`

GetRabbitmqUserVhostOk returns a tuple with the RabbitmqUserVhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserVhost

`func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserVhost(v string)`

SetRabbitmqUserVhost sets RabbitmqUserVhost field to given value.

### HasRabbitmqUserVhost

`func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserVhost() bool`

HasRabbitmqUserVhost returns a boolean if a field has been set.

### GetRabbitmqUserWritePermission

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserWritePermission() string`

GetRabbitmqUserWritePermission returns the RabbitmqUserWritePermission field if non-nil, zero value otherwise.

### GetRabbitmqUserWritePermissionOk

`func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserWritePermissionOk() (*string, bool)`

GetRabbitmqUserWritePermissionOk returns a tuple with the RabbitmqUserWritePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserWritePermission

`func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserWritePermission(v string)`

SetRabbitmqUserWritePermission sets RabbitmqUserWritePermission field to given value.

### HasRabbitmqUserWritePermission

`func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserWritePermission() bool`

HasRabbitmqUserWritePermission returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateRabbitMq) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateRabbitMq) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessUrl

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessUrl() string`

GetSecureAccessUrl returns the SecureAccessUrl field if non-nil, zero value otherwise.

### GetSecureAccessUrlOk

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessUrlOk() (*string, bool)`

GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUrl

`func (o *DynamicSecretCreateRabbitMq) SetSecureAccessUrl(v string)`

SetSecureAccessUrl sets SecureAccessUrl field to given value.

### HasSecureAccessUrl

`func (o *DynamicSecretCreateRabbitMq) HasSecureAccessUrl() bool`

HasSecureAccessUrl returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretCreateRabbitMq) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretCreateRabbitMq) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *DynamicSecretCreateRabbitMq) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *DynamicSecretCreateRabbitMq) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *DynamicSecretCreateRabbitMq) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *DynamicSecretCreateRabbitMq) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateRabbitMq) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateRabbitMq) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateRabbitMq) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateRabbitMq) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateRabbitMq) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateRabbitMq) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateRabbitMq) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateRabbitMq) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateRabbitMq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateRabbitMq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateRabbitMq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateRabbitMq) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateRabbitMq) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateRabbitMq) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateRabbitMq) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateRabbitMq) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateRabbitMq) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateRabbitMq) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateRabbitMq) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateRabbitMq) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


