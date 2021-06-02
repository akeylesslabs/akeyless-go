# UpdateRabbitMQTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Target name | 
**NewVersion** | Pointer to **bool** | Whether to create a new version of not | [optional] [default to false]
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**RabbitmqServerPassword** | Pointer to **string** |  | [optional] 
**RabbitmqServerUri** | Pointer to **string** |  | [optional] 
**RabbitmqServerUser** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewUpdateRabbitMQTargetDetails

`func NewUpdateRabbitMQTargetDetails(name string, ) *UpdateRabbitMQTargetDetails`

NewUpdateRabbitMQTargetDetails instantiates a new UpdateRabbitMQTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRabbitMQTargetDetailsWithDefaults

`func NewUpdateRabbitMQTargetDetailsWithDefaults() *UpdateRabbitMQTargetDetails`

NewUpdateRabbitMQTargetDetailsWithDefaults instantiates a new UpdateRabbitMQTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRabbitMQTargetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRabbitMQTargetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRabbitMQTargetDetails) SetName(v string)`

SetName sets Name field to given value.


### GetNewVersion

`func (o *UpdateRabbitMQTargetDetails) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *UpdateRabbitMQTargetDetails) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *UpdateRabbitMQTargetDetails) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *UpdateRabbitMQTargetDetails) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateRabbitMQTargetDetails) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateRabbitMQTargetDetails) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateRabbitMQTargetDetails) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateRabbitMQTargetDetails) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProtectionKey

`func (o *UpdateRabbitMQTargetDetails) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *UpdateRabbitMQTargetDetails) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *UpdateRabbitMQTargetDetails) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *UpdateRabbitMQTargetDetails) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetRabbitmqServerPassword

`func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerPassword() string`

GetRabbitmqServerPassword returns the RabbitmqServerPassword field if non-nil, zero value otherwise.

### GetRabbitmqServerPasswordOk

`func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerPasswordOk() (*string, bool)`

GetRabbitmqServerPasswordOk returns a tuple with the RabbitmqServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerPassword

`func (o *UpdateRabbitMQTargetDetails) SetRabbitmqServerPassword(v string)`

SetRabbitmqServerPassword sets RabbitmqServerPassword field to given value.

### HasRabbitmqServerPassword

`func (o *UpdateRabbitMQTargetDetails) HasRabbitmqServerPassword() bool`

HasRabbitmqServerPassword returns a boolean if a field has been set.

### GetRabbitmqServerUri

`func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerUri() string`

GetRabbitmqServerUri returns the RabbitmqServerUri field if non-nil, zero value otherwise.

### GetRabbitmqServerUriOk

`func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerUriOk() (*string, bool)`

GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUri

`func (o *UpdateRabbitMQTargetDetails) SetRabbitmqServerUri(v string)`

SetRabbitmqServerUri sets RabbitmqServerUri field to given value.

### HasRabbitmqServerUri

`func (o *UpdateRabbitMQTargetDetails) HasRabbitmqServerUri() bool`

HasRabbitmqServerUri returns a boolean if a field has been set.

### GetRabbitmqServerUser

`func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerUser() string`

GetRabbitmqServerUser returns the RabbitmqServerUser field if non-nil, zero value otherwise.

### GetRabbitmqServerUserOk

`func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerUserOk() (*string, bool)`

GetRabbitmqServerUserOk returns a tuple with the RabbitmqServerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUser

`func (o *UpdateRabbitMQTargetDetails) SetRabbitmqServerUser(v string)`

SetRabbitmqServerUser sets RabbitmqServerUser field to given value.

### HasRabbitmqServerUser

`func (o *UpdateRabbitMQTargetDetails) HasRabbitmqServerUser() bool`

HasRabbitmqServerUser returns a boolean if a field has been set.

### GetToken

`func (o *UpdateRabbitMQTargetDetails) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateRabbitMQTargetDetails) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateRabbitMQTargetDetails) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateRabbitMQTargetDetails) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateRabbitMQTargetDetails) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateRabbitMQTargetDetails) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateRabbitMQTargetDetails) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateRabbitMQTargetDetails) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateRabbitMQTargetDetails) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateRabbitMQTargetDetails) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateRabbitMQTargetDetails) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateRabbitMQTargetDetails) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


