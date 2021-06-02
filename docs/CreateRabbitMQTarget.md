# CreateRabbitMQTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**RabbitmqServerPassword** | Pointer to **string** |  | [optional] 
**RabbitmqServerUri** | Pointer to **string** |  | [optional] 
**RabbitmqServerUser** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

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

### GetKey

`func (o *CreateRabbitMQTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateRabbitMQTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateRabbitMQTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateRabbitMQTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

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


### GetPassword

`func (o *CreateRabbitMQTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateRabbitMQTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateRabbitMQTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateRabbitMQTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

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

### GetUsername

`func (o *CreateRabbitMQTarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateRabbitMQTarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateRabbitMQTarget) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateRabbitMQTarget) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


