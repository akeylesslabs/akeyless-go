# CreateRotatedSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiId** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** | ApplicationId (used in azure) | [optional] 
**AuthenticationCredentials** | Pointer to **string** |  | [optional] 
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation | [optional] 
**CustomPayload** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Metadata** | Pointer to **string** | Metadata about the secret | [optional] 
**Name** | **string** | Secret name | 
**RotatedPassword** | Pointer to **string** |  | [optional] 
**RotatedUsername** | Pointer to **string** |  | [optional] 
**RotationHour** | Pointer to **int32** |  | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (1-365) | [optional] 
**RotatorCredsType** | Pointer to **string** |  | [optional] 
**RotatorCustomCmd** | Pointer to **string** |  | [optional] 
**RotatorType** | **string** | Rotator Type | 
**SshPassword** | Pointer to **string** | Deprecated: use RotatedPassword | [optional] 
**SshUsername** | Pointer to **string** | Deprecated: use RotatedUser | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserAttribute** | Pointer to **string** | User Attribute | [optional] 
**UserDn** | Pointer to **string** | User DN | [optional] 

## Methods

### NewCreateRotatedSecret

`func NewCreateRotatedSecret(name string, rotatorType string, targetName string, ) *CreateRotatedSecret`

NewCreateRotatedSecret instantiates a new CreateRotatedSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRotatedSecretWithDefaults

`func NewCreateRotatedSecretWithDefaults() *CreateRotatedSecret`

NewCreateRotatedSecretWithDefaults instantiates a new CreateRotatedSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiId

`func (o *CreateRotatedSecret) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *CreateRotatedSecret) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *CreateRotatedSecret) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *CreateRotatedSecret) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApiKey

`func (o *CreateRotatedSecret) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateRotatedSecret) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateRotatedSecret) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateRotatedSecret) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApplicationId

`func (o *CreateRotatedSecret) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CreateRotatedSecret) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CreateRotatedSecret) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CreateRotatedSecret) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAuthenticationCredentials

`func (o *CreateRotatedSecret) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *CreateRotatedSecret) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *CreateRotatedSecret) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *CreateRotatedSecret) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *CreateRotatedSecret) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *CreateRotatedSecret) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *CreateRotatedSecret) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *CreateRotatedSecret) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetCustomPayload

`func (o *CreateRotatedSecret) GetCustomPayload() string`

GetCustomPayload returns the CustomPayload field if non-nil, zero value otherwise.

### GetCustomPayloadOk

`func (o *CreateRotatedSecret) GetCustomPayloadOk() (*string, bool)`

GetCustomPayloadOk returns a tuple with the CustomPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayload

`func (o *CreateRotatedSecret) SetCustomPayload(v string)`

SetCustomPayload sets CustomPayload field to given value.

### HasCustomPayload

`func (o *CreateRotatedSecret) HasCustomPayload() bool`

HasCustomPayload returns a boolean if a field has been set.

### GetKey

`func (o *CreateRotatedSecret) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateRotatedSecret) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateRotatedSecret) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateRotatedSecret) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateRotatedSecret) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateRotatedSecret) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateRotatedSecret) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateRotatedSecret) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateRotatedSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRotatedSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRotatedSecret) SetName(v string)`

SetName sets Name field to given value.


### GetRotatedPassword

`func (o *CreateRotatedSecret) GetRotatedPassword() string`

GetRotatedPassword returns the RotatedPassword field if non-nil, zero value otherwise.

### GetRotatedPasswordOk

`func (o *CreateRotatedSecret) GetRotatedPasswordOk() (*string, bool)`

GetRotatedPasswordOk returns a tuple with the RotatedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedPassword

`func (o *CreateRotatedSecret) SetRotatedPassword(v string)`

SetRotatedPassword sets RotatedPassword field to given value.

### HasRotatedPassword

`func (o *CreateRotatedSecret) HasRotatedPassword() bool`

HasRotatedPassword returns a boolean if a field has been set.

### GetRotatedUsername

`func (o *CreateRotatedSecret) GetRotatedUsername() string`

GetRotatedUsername returns the RotatedUsername field if non-nil, zero value otherwise.

### GetRotatedUsernameOk

`func (o *CreateRotatedSecret) GetRotatedUsernameOk() (*string, bool)`

GetRotatedUsernameOk returns a tuple with the RotatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedUsername

`func (o *CreateRotatedSecret) SetRotatedUsername(v string)`

SetRotatedUsername sets RotatedUsername field to given value.

### HasRotatedUsername

`func (o *CreateRotatedSecret) HasRotatedUsername() bool`

HasRotatedUsername returns a boolean if a field has been set.

### GetRotationHour

`func (o *CreateRotatedSecret) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *CreateRotatedSecret) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *CreateRotatedSecret) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *CreateRotatedSecret) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *CreateRotatedSecret) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *CreateRotatedSecret) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *CreateRotatedSecret) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *CreateRotatedSecret) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotatorCredsType

`func (o *CreateRotatedSecret) GetRotatorCredsType() string`

GetRotatorCredsType returns the RotatorCredsType field if non-nil, zero value otherwise.

### GetRotatorCredsTypeOk

`func (o *CreateRotatedSecret) GetRotatorCredsTypeOk() (*string, bool)`

GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorCredsType

`func (o *CreateRotatedSecret) SetRotatorCredsType(v string)`

SetRotatorCredsType sets RotatorCredsType field to given value.

### HasRotatorCredsType

`func (o *CreateRotatedSecret) HasRotatorCredsType() bool`

HasRotatorCredsType returns a boolean if a field has been set.

### GetRotatorCustomCmd

`func (o *CreateRotatedSecret) GetRotatorCustomCmd() string`

GetRotatorCustomCmd returns the RotatorCustomCmd field if non-nil, zero value otherwise.

### GetRotatorCustomCmdOk

`func (o *CreateRotatedSecret) GetRotatorCustomCmdOk() (*string, bool)`

GetRotatorCustomCmdOk returns a tuple with the RotatorCustomCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorCustomCmd

`func (o *CreateRotatedSecret) SetRotatorCustomCmd(v string)`

SetRotatorCustomCmd sets RotatorCustomCmd field to given value.

### HasRotatorCustomCmd

`func (o *CreateRotatedSecret) HasRotatorCustomCmd() bool`

HasRotatorCustomCmd returns a boolean if a field has been set.

### GetRotatorType

`func (o *CreateRotatedSecret) GetRotatorType() string`

GetRotatorType returns the RotatorType field if non-nil, zero value otherwise.

### GetRotatorTypeOk

`func (o *CreateRotatedSecret) GetRotatorTypeOk() (*string, bool)`

GetRotatorTypeOk returns a tuple with the RotatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorType

`func (o *CreateRotatedSecret) SetRotatorType(v string)`

SetRotatorType sets RotatorType field to given value.


### GetSshPassword

`func (o *CreateRotatedSecret) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *CreateRotatedSecret) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *CreateRotatedSecret) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *CreateRotatedSecret) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshUsername

`func (o *CreateRotatedSecret) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *CreateRotatedSecret) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *CreateRotatedSecret) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *CreateRotatedSecret) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetTags

`func (o *CreateRotatedSecret) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRotatedSecret) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRotatedSecret) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRotatedSecret) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *CreateRotatedSecret) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *CreateRotatedSecret) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *CreateRotatedSecret) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetToken

`func (o *CreateRotatedSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateRotatedSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateRotatedSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateRotatedSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateRotatedSecret) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateRotatedSecret) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateRotatedSecret) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateRotatedSecret) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserAttribute

`func (o *CreateRotatedSecret) GetUserAttribute() string`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *CreateRotatedSecret) GetUserAttributeOk() (*string, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *CreateRotatedSecret) SetUserAttribute(v string)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *CreateRotatedSecret) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### GetUserDn

`func (o *CreateRotatedSecret) GetUserDn() string`

GetUserDn returns the UserDn field if non-nil, zero value otherwise.

### GetUserDnOk

`func (o *CreateRotatedSecret) GetUserDnOk() (*string, bool)`

GetUserDnOk returns a tuple with the UserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDn

`func (o *CreateRotatedSecret) SetUserDn(v string)`

SetUserDn sets UserDn field to given value.

### HasUserDn

`func (o *CreateRotatedSecret) HasUserDn() bool`

HasUserDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


