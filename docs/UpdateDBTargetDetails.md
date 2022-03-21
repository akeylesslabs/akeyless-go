# UpdateDBTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbType** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**KeepPrevVersion** | Pointer to **string** |  | [optional] 
**MongoDbName** | Pointer to **string** |  | [optional] 
**MongoUri** | Pointer to **string** |  | [optional] 
**Name** | **string** | Target name | 
**NewVersion** | Pointer to **bool** | Deprecated | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Pwd** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateDBTargetDetails

`func NewUpdateDBTargetDetails(name string, ) *UpdateDBTargetDetails`

NewUpdateDBTargetDetails instantiates a new UpdateDBTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDBTargetDetailsWithDefaults

`func NewUpdateDBTargetDetailsWithDefaults() *UpdateDBTargetDetails`

NewUpdateDBTargetDetailsWithDefaults instantiates a new UpdateDBTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbType

`func (o *UpdateDBTargetDetails) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *UpdateDBTargetDetails) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *UpdateDBTargetDetails) SetDbType(v string)`

SetDbType sets DbType field to given value.

### HasDbType

`func (o *UpdateDBTargetDetails) HasDbType() bool`

HasDbType returns a boolean if a field has been set.

### GetHostName

`func (o *UpdateDBTargetDetails) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *UpdateDBTargetDetails) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *UpdateDBTargetDetails) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *UpdateDBTargetDetails) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateDBTargetDetails) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateDBTargetDetails) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateDBTargetDetails) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateDBTargetDetails) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetMongoDbName

`func (o *UpdateDBTargetDetails) GetMongoDbName() string`

GetMongoDbName returns the MongoDbName field if non-nil, zero value otherwise.

### GetMongoDbNameOk

`func (o *UpdateDBTargetDetails) GetMongoDbNameOk() (*string, bool)`

GetMongoDbNameOk returns a tuple with the MongoDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDbName

`func (o *UpdateDBTargetDetails) SetMongoDbName(v string)`

SetMongoDbName sets MongoDbName field to given value.

### HasMongoDbName

`func (o *UpdateDBTargetDetails) HasMongoDbName() bool`

HasMongoDbName returns a boolean if a field has been set.

### GetMongoUri

`func (o *UpdateDBTargetDetails) GetMongoUri() string`

GetMongoUri returns the MongoUri field if non-nil, zero value otherwise.

### GetMongoUriOk

`func (o *UpdateDBTargetDetails) GetMongoUriOk() (*string, bool)`

GetMongoUriOk returns a tuple with the MongoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoUri

`func (o *UpdateDBTargetDetails) SetMongoUri(v string)`

SetMongoUri sets MongoUri field to given value.

### HasMongoUri

`func (o *UpdateDBTargetDetails) HasMongoUri() bool`

HasMongoUri returns a boolean if a field has been set.

### GetName

`func (o *UpdateDBTargetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDBTargetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDBTargetDetails) SetName(v string)`

SetName sets Name field to given value.


### GetNewVersion

`func (o *UpdateDBTargetDetails) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *UpdateDBTargetDetails) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *UpdateDBTargetDetails) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *UpdateDBTargetDetails) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetPort

`func (o *UpdateDBTargetDetails) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateDBTargetDetails) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateDBTargetDetails) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateDBTargetDetails) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtectionKey

`func (o *UpdateDBTargetDetails) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *UpdateDBTargetDetails) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *UpdateDBTargetDetails) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *UpdateDBTargetDetails) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetPwd

`func (o *UpdateDBTargetDetails) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *UpdateDBTargetDetails) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *UpdateDBTargetDetails) SetPwd(v string)`

SetPwd sets Pwd field to given value.

### HasPwd

`func (o *UpdateDBTargetDetails) HasPwd() bool`

HasPwd returns a boolean if a field has been set.

### GetToken

`func (o *UpdateDBTargetDetails) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateDBTargetDetails) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateDBTargetDetails) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateDBTargetDetails) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateDBTargetDetails) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateDBTargetDetails) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateDBTargetDetails) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateDBTargetDetails) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserName

`func (o *UpdateDBTargetDetails) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UpdateDBTargetDetails) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UpdateDBTargetDetails) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UpdateDBTargetDetails) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


