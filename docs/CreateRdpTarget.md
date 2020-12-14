# CreateRdpTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminName** | Pointer to **string** |  | [optional] 
**AdminPwd** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**HostPort** | Pointer to **string** |  | [optional] 
**Name** | **string** | Target name | 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

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

### GetAdminName

`func (o *CreateRdpTarget) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *CreateRdpTarget) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *CreateRdpTarget) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *CreateRdpTarget) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAdminPwd

`func (o *CreateRdpTarget) GetAdminPwd() string`

GetAdminPwd returns the AdminPwd field if non-nil, zero value otherwise.

### GetAdminPwdOk

`func (o *CreateRdpTarget) GetAdminPwdOk() (*string, bool)`

GetAdminPwdOk returns a tuple with the AdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPwd

`func (o *CreateRdpTarget) SetAdminPwd(v string)`

SetAdminPwd sets AdminPwd field to given value.

### HasAdminPwd

`func (o *CreateRdpTarget) HasAdminPwd() bool`

HasAdminPwd returns a boolean if a field has been set.

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

### GetHostName

`func (o *CreateRdpTarget) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CreateRdpTarget) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CreateRdpTarget) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CreateRdpTarget) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHostPort

`func (o *CreateRdpTarget) GetHostPort() string`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *CreateRdpTarget) GetHostPortOk() (*string, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *CreateRdpTarget) SetHostPort(v string)`

SetHostPort sets HostPort field to given value.

### HasHostPort

`func (o *CreateRdpTarget) HasHostPort() bool`

HasHostPort returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


