# DynamicSecretUpdateRabbitMq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
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

### NewDynamicSecretUpdateRabbitMq

`func NewDynamicSecretUpdateRabbitMq(name string, ) *DynamicSecretUpdateRabbitMq`

NewDynamicSecretUpdateRabbitMq instantiates a new DynamicSecretUpdateRabbitMq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateRabbitMqWithDefaults

`func NewDynamicSecretUpdateRabbitMqWithDefaults() *DynamicSecretUpdateRabbitMq`

NewDynamicSecretUpdateRabbitMqWithDefaults instantiates a new DynamicSecretUpdateRabbitMq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretUpdateRabbitMq) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateRabbitMq) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateRabbitMq) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateRabbitMq) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateRabbitMq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateRabbitMq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateRabbitMq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateRabbitMq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateRabbitMq) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateRabbitMq) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateRabbitMq) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateRabbitMq) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateRabbitMq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateRabbitMq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateRabbitMq) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateRabbitMq) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateRabbitMq) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateRabbitMq) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateRabbitMq) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretUpdateRabbitMq) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretUpdateRabbitMq) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretUpdateRabbitMq) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretUpdateRabbitMq) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateRabbitMq) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateRabbitMq) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateRabbitMq) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateRabbitMq) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRabbitmqAdminPwd

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqAdminPwd() string`

GetRabbitmqAdminPwd returns the RabbitmqAdminPwd field if non-nil, zero value otherwise.

### GetRabbitmqAdminPwdOk

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqAdminPwdOk() (*string, bool)`

GetRabbitmqAdminPwdOk returns a tuple with the RabbitmqAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqAdminPwd

`func (o *DynamicSecretUpdateRabbitMq) SetRabbitmqAdminPwd(v string)`

SetRabbitmqAdminPwd sets RabbitmqAdminPwd field to given value.

### HasRabbitmqAdminPwd

`func (o *DynamicSecretUpdateRabbitMq) HasRabbitmqAdminPwd() bool`

HasRabbitmqAdminPwd returns a boolean if a field has been set.

### GetRabbitmqAdminUser

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqAdminUser() string`

GetRabbitmqAdminUser returns the RabbitmqAdminUser field if non-nil, zero value otherwise.

### GetRabbitmqAdminUserOk

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqAdminUserOk() (*string, bool)`

GetRabbitmqAdminUserOk returns a tuple with the RabbitmqAdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqAdminUser

`func (o *DynamicSecretUpdateRabbitMq) SetRabbitmqAdminUser(v string)`

SetRabbitmqAdminUser sets RabbitmqAdminUser field to given value.

### HasRabbitmqAdminUser

`func (o *DynamicSecretUpdateRabbitMq) HasRabbitmqAdminUser() bool`

HasRabbitmqAdminUser returns a boolean if a field has been set.

### GetRabbitmqServerUri

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqServerUri() string`

GetRabbitmqServerUri returns the RabbitmqServerUri field if non-nil, zero value otherwise.

### GetRabbitmqServerUriOk

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqServerUriOk() (*string, bool)`

GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUri

`func (o *DynamicSecretUpdateRabbitMq) SetRabbitmqServerUri(v string)`

SetRabbitmqServerUri sets RabbitmqServerUri field to given value.

### HasRabbitmqServerUri

`func (o *DynamicSecretUpdateRabbitMq) HasRabbitmqServerUri() bool`

HasRabbitmqServerUri returns a boolean if a field has been set.

### GetRabbitmqUserConfPermission

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserConfPermission() string`

GetRabbitmqUserConfPermission returns the RabbitmqUserConfPermission field if non-nil, zero value otherwise.

### GetRabbitmqUserConfPermissionOk

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserConfPermissionOk() (*string, bool)`

GetRabbitmqUserConfPermissionOk returns a tuple with the RabbitmqUserConfPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserConfPermission

`func (o *DynamicSecretUpdateRabbitMq) SetRabbitmqUserConfPermission(v string)`

SetRabbitmqUserConfPermission sets RabbitmqUserConfPermission field to given value.

### HasRabbitmqUserConfPermission

`func (o *DynamicSecretUpdateRabbitMq) HasRabbitmqUserConfPermission() bool`

HasRabbitmqUserConfPermission returns a boolean if a field has been set.

### GetRabbitmqUserReadPermission

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserReadPermission() string`

GetRabbitmqUserReadPermission returns the RabbitmqUserReadPermission field if non-nil, zero value otherwise.

### GetRabbitmqUserReadPermissionOk

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserReadPermissionOk() (*string, bool)`

GetRabbitmqUserReadPermissionOk returns a tuple with the RabbitmqUserReadPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserReadPermission

`func (o *DynamicSecretUpdateRabbitMq) SetRabbitmqUserReadPermission(v string)`

SetRabbitmqUserReadPermission sets RabbitmqUserReadPermission field to given value.

### HasRabbitmqUserReadPermission

`func (o *DynamicSecretUpdateRabbitMq) HasRabbitmqUserReadPermission() bool`

HasRabbitmqUserReadPermission returns a boolean if a field has been set.

### GetRabbitmqUserTags

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserTags() string`

GetRabbitmqUserTags returns the RabbitmqUserTags field if non-nil, zero value otherwise.

### GetRabbitmqUserTagsOk

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserTagsOk() (*string, bool)`

GetRabbitmqUserTagsOk returns a tuple with the RabbitmqUserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserTags

`func (o *DynamicSecretUpdateRabbitMq) SetRabbitmqUserTags(v string)`

SetRabbitmqUserTags sets RabbitmqUserTags field to given value.

### HasRabbitmqUserTags

`func (o *DynamicSecretUpdateRabbitMq) HasRabbitmqUserTags() bool`

HasRabbitmqUserTags returns a boolean if a field has been set.

### GetRabbitmqUserVhost

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserVhost() string`

GetRabbitmqUserVhost returns the RabbitmqUserVhost field if non-nil, zero value otherwise.

### GetRabbitmqUserVhostOk

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserVhostOk() (*string, bool)`

GetRabbitmqUserVhostOk returns a tuple with the RabbitmqUserVhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserVhost

`func (o *DynamicSecretUpdateRabbitMq) SetRabbitmqUserVhost(v string)`

SetRabbitmqUserVhost sets RabbitmqUserVhost field to given value.

### HasRabbitmqUserVhost

`func (o *DynamicSecretUpdateRabbitMq) HasRabbitmqUserVhost() bool`

HasRabbitmqUserVhost returns a boolean if a field has been set.

### GetRabbitmqUserWritePermission

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserWritePermission() string`

GetRabbitmqUserWritePermission returns the RabbitmqUserWritePermission field if non-nil, zero value otherwise.

### GetRabbitmqUserWritePermissionOk

`func (o *DynamicSecretUpdateRabbitMq) GetRabbitmqUserWritePermissionOk() (*string, bool)`

GetRabbitmqUserWritePermissionOk returns a tuple with the RabbitmqUserWritePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserWritePermission

`func (o *DynamicSecretUpdateRabbitMq) SetRabbitmqUserWritePermission(v string)`

SetRabbitmqUserWritePermission sets RabbitmqUserWritePermission field to given value.

### HasRabbitmqUserWritePermission

`func (o *DynamicSecretUpdateRabbitMq) HasRabbitmqUserWritePermission() bool`

HasRabbitmqUserWritePermission returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateRabbitMq) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateRabbitMq) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessUrl

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessUrl() string`

GetSecureAccessUrl returns the SecureAccessUrl field if non-nil, zero value otherwise.

### GetSecureAccessUrlOk

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessUrlOk() (*string, bool)`

GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUrl

`func (o *DynamicSecretUpdateRabbitMq) SetSecureAccessUrl(v string)`

SetSecureAccessUrl sets SecureAccessUrl field to given value.

### HasSecureAccessUrl

`func (o *DynamicSecretUpdateRabbitMq) HasSecureAccessUrl() bool`

HasSecureAccessUrl returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretUpdateRabbitMq) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretUpdateRabbitMq) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *DynamicSecretUpdateRabbitMq) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *DynamicSecretUpdateRabbitMq) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *DynamicSecretUpdateRabbitMq) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *DynamicSecretUpdateRabbitMq) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *DynamicSecretUpdateRabbitMq) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateRabbitMq) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateRabbitMq) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateRabbitMq) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateRabbitMq) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateRabbitMq) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateRabbitMq) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateRabbitMq) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateRabbitMq) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateRabbitMq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateRabbitMq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateRabbitMq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateRabbitMq) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateRabbitMq) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateRabbitMq) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateRabbitMq) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateRabbitMq) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateRabbitMq) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateRabbitMq) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateRabbitMq) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateRabbitMq) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


