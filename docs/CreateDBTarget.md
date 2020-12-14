# CreateDBTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**DbType** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**MongoDbName** | Pointer to **string** |  | [optional] 
**MongoUri** | Pointer to **string** |  | [optional] 
**Name** | **string** | Target name | 
**Port** | Pointer to **string** |  | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Pwd** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDBTarget

`func NewCreateDBTarget(name string, ) *CreateDBTarget`

NewCreateDBTarget instantiates a new CreateDBTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDBTargetWithDefaults

`func NewCreateDBTargetWithDefaults() *CreateDBTarget`

NewCreateDBTargetWithDefaults instantiates a new CreateDBTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateDBTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateDBTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateDBTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateDBTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDbType

`func (o *CreateDBTarget) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *CreateDBTarget) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *CreateDBTarget) SetDbType(v string)`

SetDbType sets DbType field to given value.

### HasDbType

`func (o *CreateDBTarget) HasDbType() bool`

HasDbType returns a boolean if a field has been set.

### GetHostName

`func (o *CreateDBTarget) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CreateDBTarget) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CreateDBTarget) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CreateDBTarget) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetMongoDbName

`func (o *CreateDBTarget) GetMongoDbName() string`

GetMongoDbName returns the MongoDbName field if non-nil, zero value otherwise.

### GetMongoDbNameOk

`func (o *CreateDBTarget) GetMongoDbNameOk() (*string, bool)`

GetMongoDbNameOk returns a tuple with the MongoDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDbName

`func (o *CreateDBTarget) SetMongoDbName(v string)`

SetMongoDbName sets MongoDbName field to given value.

### HasMongoDbName

`func (o *CreateDBTarget) HasMongoDbName() bool`

HasMongoDbName returns a boolean if a field has been set.

### GetMongoUri

`func (o *CreateDBTarget) GetMongoUri() string`

GetMongoUri returns the MongoUri field if non-nil, zero value otherwise.

### GetMongoUriOk

`func (o *CreateDBTarget) GetMongoUriOk() (*string, bool)`

GetMongoUriOk returns a tuple with the MongoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoUri

`func (o *CreateDBTarget) SetMongoUri(v string)`

SetMongoUri sets MongoUri field to given value.

### HasMongoUri

`func (o *CreateDBTarget) HasMongoUri() bool`

HasMongoUri returns a boolean if a field has been set.

### GetName

`func (o *CreateDBTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDBTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDBTarget) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *CreateDBTarget) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateDBTarget) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateDBTarget) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateDBTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtectionKey

`func (o *CreateDBTarget) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateDBTarget) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateDBTarget) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateDBTarget) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetPwd

`func (o *CreateDBTarget) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *CreateDBTarget) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *CreateDBTarget) SetPwd(v string)`

SetPwd sets Pwd field to given value.

### HasPwd

`func (o *CreateDBTarget) HasPwd() bool`

HasPwd returns a boolean if a field has been set.

### GetToken

`func (o *CreateDBTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateDBTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateDBTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateDBTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateDBTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateDBTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateDBTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateDBTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserName

`func (o *CreateDBTarget) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreateDBTarget) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreateDBTarget) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *CreateDBTarget) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


