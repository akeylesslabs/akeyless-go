# RotatedSecretCreateCassandra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationCredentials** | Pointer to **string** | The credentials to connect with use-user-creds/use-target-creds | [optional] [default to "use-user-creds"]
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation [true/false] | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Rotated secret name | 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**RotatedPassword** | Pointer to **string** | rotated-username password (relevant only for rotator-type&#x3D;password) | [optional] 
**RotatedUsername** | Pointer to **string** | username to be rotated, if selected use-self-creds at rotator-creds-type, this username will try to rotate it&#39;s own password, if use-target-creds is selected, target credentials will be use to rotate the rotated-password (relevant only for rotator-type&#x3D;password) | [optional] 
**RotationHour** | Pointer to **int32** | The Hour of the rotation in UTC | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (1-365) | [optional] 
**RotatorType** | **string** | The rotator type. options: [target/password] | 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotatedSecretCreateCassandra

`func NewRotatedSecretCreateCassandra(name string, rotatorType string, targetName string, ) *RotatedSecretCreateCassandra`

NewRotatedSecretCreateCassandra instantiates a new RotatedSecretCreateCassandra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretCreateCassandraWithDefaults

`func NewRotatedSecretCreateCassandraWithDefaults() *RotatedSecretCreateCassandra`

NewRotatedSecretCreateCassandraWithDefaults instantiates a new RotatedSecretCreateCassandra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationCredentials

`func (o *RotatedSecretCreateCassandra) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *RotatedSecretCreateCassandra) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *RotatedSecretCreateCassandra) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *RotatedSecretCreateCassandra) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretCreateCassandra) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretCreateCassandra) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretCreateCassandra) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretCreateCassandra) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretCreateCassandra) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretCreateCassandra) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretCreateCassandra) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretCreateCassandra) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretCreateCassandra) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretCreateCassandra) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretCreateCassandra) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretCreateCassandra) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretCreateCassandra) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretCreateCassandra) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretCreateCassandra) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretCreateCassandra) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretCreateCassandra) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretCreateCassandra) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretCreateCassandra) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretCreateCassandra) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretCreateCassandra) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretCreateCassandra) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretCreateCassandra) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *RotatedSecretCreateCassandra) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretCreateCassandra) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretCreateCassandra) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretCreateCassandra) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRotatedPassword

`func (o *RotatedSecretCreateCassandra) GetRotatedPassword() string`

GetRotatedPassword returns the RotatedPassword field if non-nil, zero value otherwise.

### GetRotatedPasswordOk

`func (o *RotatedSecretCreateCassandra) GetRotatedPasswordOk() (*string, bool)`

GetRotatedPasswordOk returns a tuple with the RotatedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedPassword

`func (o *RotatedSecretCreateCassandra) SetRotatedPassword(v string)`

SetRotatedPassword sets RotatedPassword field to given value.

### HasRotatedPassword

`func (o *RotatedSecretCreateCassandra) HasRotatedPassword() bool`

HasRotatedPassword returns a boolean if a field has been set.

### GetRotatedUsername

`func (o *RotatedSecretCreateCassandra) GetRotatedUsername() string`

GetRotatedUsername returns the RotatedUsername field if non-nil, zero value otherwise.

### GetRotatedUsernameOk

`func (o *RotatedSecretCreateCassandra) GetRotatedUsernameOk() (*string, bool)`

GetRotatedUsernameOk returns a tuple with the RotatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedUsername

`func (o *RotatedSecretCreateCassandra) SetRotatedUsername(v string)`

SetRotatedUsername sets RotatedUsername field to given value.

### HasRotatedUsername

`func (o *RotatedSecretCreateCassandra) HasRotatedUsername() bool`

HasRotatedUsername returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretCreateCassandra) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretCreateCassandra) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretCreateCassandra) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretCreateCassandra) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretCreateCassandra) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretCreateCassandra) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretCreateCassandra) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretCreateCassandra) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotatorType

`func (o *RotatedSecretCreateCassandra) GetRotatorType() string`

GetRotatorType returns the RotatorType field if non-nil, zero value otherwise.

### GetRotatorTypeOk

`func (o *RotatedSecretCreateCassandra) GetRotatorTypeOk() (*string, bool)`

GetRotatorTypeOk returns a tuple with the RotatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorType

`func (o *RotatedSecretCreateCassandra) SetRotatorType(v string)`

SetRotatorType sets RotatorType field to given value.


### GetTags

`func (o *RotatedSecretCreateCassandra) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RotatedSecretCreateCassandra) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RotatedSecretCreateCassandra) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RotatedSecretCreateCassandra) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *RotatedSecretCreateCassandra) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *RotatedSecretCreateCassandra) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *RotatedSecretCreateCassandra) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetToken

`func (o *RotatedSecretCreateCassandra) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretCreateCassandra) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretCreateCassandra) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretCreateCassandra) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretCreateCassandra) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretCreateCassandra) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretCreateCassandra) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretCreateCassandra) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


