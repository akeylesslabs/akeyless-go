# CreateAwsTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Name** | **string** | Target name | 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SessionToken** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateAwsTarget

`func NewCreateAwsTarget(name string, ) *CreateAwsTarget`

NewCreateAwsTarget instantiates a new CreateAwsTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAwsTargetWithDefaults

`func NewCreateAwsTargetWithDefaults() *CreateAwsTarget`

NewCreateAwsTargetWithDefaults instantiates a new CreateAwsTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *CreateAwsTarget) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CreateAwsTarget) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CreateAwsTarget) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *CreateAwsTarget) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *CreateAwsTarget) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *CreateAwsTarget) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *CreateAwsTarget) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *CreateAwsTarget) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetComment

`func (o *CreateAwsTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateAwsTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateAwsTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateAwsTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *CreateAwsTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAwsTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAwsTarget) SetName(v string)`

SetName sets Name field to given value.


### GetProtectionKey

`func (o *CreateAwsTarget) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateAwsTarget) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateAwsTarget) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateAwsTarget) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetRegion

`func (o *CreateAwsTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateAwsTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateAwsTarget) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateAwsTarget) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSessionToken

`func (o *CreateAwsTarget) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *CreateAwsTarget) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *CreateAwsTarget) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *CreateAwsTarget) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetToken

`func (o *CreateAwsTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAwsTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAwsTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAwsTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAwsTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAwsTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAwsTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAwsTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


