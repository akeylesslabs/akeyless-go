# CreateSSHTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Ip** | Pointer to **[]string** |  | [optional] 
**Name** | **string** | Target name | 
**Port** | Pointer to **string** |  | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateSSHTarget

`func NewCreateSSHTarget(name string, ) *CreateSSHTarget`

NewCreateSSHTarget instantiates a new CreateSSHTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSSHTargetWithDefaults

`func NewCreateSSHTargetWithDefaults() *CreateSSHTarget`

NewCreateSSHTargetWithDefaults instantiates a new CreateSSHTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateSSHTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateSSHTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateSSHTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateSSHTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetIp

`func (o *CreateSSHTarget) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateSSHTarget) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateSSHTarget) SetIp(v []string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreateSSHTarget) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *CreateSSHTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSSHTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSSHTarget) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *CreateSSHTarget) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateSSHTarget) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateSSHTarget) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateSSHTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtectionKey

`func (o *CreateSSHTarget) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateSSHTarget) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateSSHTarget) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateSSHTarget) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetToken

`func (o *CreateSSHTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateSSHTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateSSHTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateSSHTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateSSHTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateSSHTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateSSHTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateSSHTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


