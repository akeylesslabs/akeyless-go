# CreateRabbitMQTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Name** | **string** | Target name | 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**RabbitmqServerPassword** | Pointer to **string** |  | [optional] 
**RabbitmqServerUri** | Pointer to **string** |  | [optional] 
**RabbitmqServerUser** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateRabbitMQTarget

`func NewCreateRabbitMQTarget(name string, ) *CreateRabbitMQTarget`

NewCreateRabbitMQTarget instantiates a new CreateRabbitMQTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRabbitMQTargetWithDefaults

`func NewCreateRabbitMQTargetWithDefaults() *CreateRabbitMQTarget`

NewCreateRabbitMQTargetWithDefaults instantiates a new CreateRabbitMQTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateRabbitMQTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateRabbitMQTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateRabbitMQTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateRabbitMQTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *CreateRabbitMQTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRabbitMQTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRabbitMQTarget) SetName(v string)`

SetName sets Name field to given value.


### GetProtectionKey

`func (o *CreateRabbitMQTarget) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateRabbitMQTarget) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateRabbitMQTarget) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateRabbitMQTarget) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetRabbitmqServerPassword

`func (o *CreateRabbitMQTarget) GetRabbitmqServerPassword() string`

GetRabbitmqServerPassword returns the RabbitmqServerPassword field if non-nil, zero value otherwise.

### GetRabbitmqServerPasswordOk

`func (o *CreateRabbitMQTarget) GetRabbitmqServerPasswordOk() (*string, bool)`

GetRabbitmqServerPasswordOk returns a tuple with the RabbitmqServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerPassword

`func (o *CreateRabbitMQTarget) SetRabbitmqServerPassword(v string)`

SetRabbitmqServerPassword sets RabbitmqServerPassword field to given value.

### HasRabbitmqServerPassword

`func (o *CreateRabbitMQTarget) HasRabbitmqServerPassword() bool`

HasRabbitmqServerPassword returns a boolean if a field has been set.

### GetRabbitmqServerUri

`func (o *CreateRabbitMQTarget) GetRabbitmqServerUri() string`

GetRabbitmqServerUri returns the RabbitmqServerUri field if non-nil, zero value otherwise.

### GetRabbitmqServerUriOk

`func (o *CreateRabbitMQTarget) GetRabbitmqServerUriOk() (*string, bool)`

GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUri

`func (o *CreateRabbitMQTarget) SetRabbitmqServerUri(v string)`

SetRabbitmqServerUri sets RabbitmqServerUri field to given value.

### HasRabbitmqServerUri

`func (o *CreateRabbitMQTarget) HasRabbitmqServerUri() bool`

HasRabbitmqServerUri returns a boolean if a field has been set.

### GetRabbitmqServerUser

`func (o *CreateRabbitMQTarget) GetRabbitmqServerUser() string`

GetRabbitmqServerUser returns the RabbitmqServerUser field if non-nil, zero value otherwise.

### GetRabbitmqServerUserOk

`func (o *CreateRabbitMQTarget) GetRabbitmqServerUserOk() (*string, bool)`

GetRabbitmqServerUserOk returns a tuple with the RabbitmqServerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUser

`func (o *CreateRabbitMQTarget) SetRabbitmqServerUser(v string)`

SetRabbitmqServerUser sets RabbitmqServerUser field to given value.

### HasRabbitmqServerUser

`func (o *CreateRabbitMQTarget) HasRabbitmqServerUser() bool`

HasRabbitmqServerUser returns a boolean if a field has been set.

### GetToken

`func (o *CreateRabbitMQTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateRabbitMQTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateRabbitMQTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateRabbitMQTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateRabbitMQTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateRabbitMQTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateRabbitMQTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateRabbitMQTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


