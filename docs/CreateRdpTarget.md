# CreateRdpTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Name** | **string** | Target name | 
**Pass** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRdpTarget

`func NewCreateRdpTarget(name string, ) *CreateRdpTarget`

NewCreateRdpTarget instantiates a new CreateRdpTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRdpTargetWithDefaults

`func NewCreateRdpTargetWithDefaults() *CreateRdpTarget`

NewCreateRdpTargetWithDefaults instantiates a new CreateRdpTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateRdpTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateRdpTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateRdpTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateRdpTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetHost

`func (o *CreateRdpTarget) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateRdpTarget) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateRdpTarget) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateRdpTarget) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *CreateRdpTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRdpTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRdpTarget) SetName(v string)`

SetName sets Name field to given value.


### GetPass

`func (o *CreateRdpTarget) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *CreateRdpTarget) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *CreateRdpTarget) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *CreateRdpTarget) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetPort

`func (o *CreateRdpTarget) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateRdpTarget) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateRdpTarget) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateRdpTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtectionKey

`func (o *CreateRdpTarget) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateRdpTarget) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateRdpTarget) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateRdpTarget) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetToken

`func (o *CreateRdpTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateRdpTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateRdpTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateRdpTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateRdpTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateRdpTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateRdpTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateRdpTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUser

`func (o *CreateRdpTarget) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateRdpTarget) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateRdpTarget) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateRdpTarget) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


