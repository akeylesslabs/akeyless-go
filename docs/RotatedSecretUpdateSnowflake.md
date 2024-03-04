# RotatedSecretUpdateSnowflake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**AuthenticationCredentials** | Pointer to **string** | The credentials to connect with use-user-creds/use-target-creds | [optional] [default to "use-user-creds"]
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation [true/false] | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_metadata"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Rotated secret name | 
**NewName** | Pointer to **string** | New item name | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**RotatedPassword** | Pointer to **string** | rotated-username password (relevant only for rotator-type&#x3D;password) | [optional] 
**RotatedUsername** | Pointer to **string** | username to be rotated, if selected use-self-creds at rotator-creds-type, this username will try to rotate it&#39;s own password, if use-target-creds is selected, target credentials will be use to rotate the rotated-password (relevant only for rotator-type&#x3D;password) | [optional] 
**RotationHour** | Pointer to **int32** | The Hour of the rotation in UTC | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (1-365) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotatedSecretUpdateSnowflake

`func NewRotatedSecretUpdateSnowflake(name string, ) *RotatedSecretUpdateSnowflake`

NewRotatedSecretUpdateSnowflake instantiates a new RotatedSecretUpdateSnowflake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretUpdateSnowflakeWithDefaults

`func NewRotatedSecretUpdateSnowflakeWithDefaults() *RotatedSecretUpdateSnowflake`

NewRotatedSecretUpdateSnowflakeWithDefaults instantiates a new RotatedSecretUpdateSnowflake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *RotatedSecretUpdateSnowflake) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *RotatedSecretUpdateSnowflake) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *RotatedSecretUpdateSnowflake) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *RotatedSecretUpdateSnowflake) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAuthenticationCredentials

`func (o *RotatedSecretUpdateSnowflake) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *RotatedSecretUpdateSnowflake) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *RotatedSecretUpdateSnowflake) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *RotatedSecretUpdateSnowflake) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretUpdateSnowflake) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretUpdateSnowflake) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretUpdateSnowflake) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretUpdateSnowflake) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretUpdateSnowflake) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretUpdateSnowflake) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretUpdateSnowflake) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretUpdateSnowflake) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretUpdateSnowflake) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretUpdateSnowflake) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretUpdateSnowflake) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretUpdateSnowflake) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretUpdateSnowflake) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretUpdateSnowflake) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretUpdateSnowflake) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretUpdateSnowflake) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *RotatedSecretUpdateSnowflake) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *RotatedSecretUpdateSnowflake) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *RotatedSecretUpdateSnowflake) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *RotatedSecretUpdateSnowflake) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretUpdateSnowflake) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretUpdateSnowflake) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretUpdateSnowflake) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretUpdateSnowflake) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretUpdateSnowflake) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretUpdateSnowflake) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretUpdateSnowflake) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *RotatedSecretUpdateSnowflake) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *RotatedSecretUpdateSnowflake) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *RotatedSecretUpdateSnowflake) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *RotatedSecretUpdateSnowflake) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *RotatedSecretUpdateSnowflake) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretUpdateSnowflake) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretUpdateSnowflake) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretUpdateSnowflake) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRmTag

`func (o *RotatedSecretUpdateSnowflake) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *RotatedSecretUpdateSnowflake) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *RotatedSecretUpdateSnowflake) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *RotatedSecretUpdateSnowflake) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotatedPassword

`func (o *RotatedSecretUpdateSnowflake) GetRotatedPassword() string`

GetRotatedPassword returns the RotatedPassword field if non-nil, zero value otherwise.

### GetRotatedPasswordOk

`func (o *RotatedSecretUpdateSnowflake) GetRotatedPasswordOk() (*string, bool)`

GetRotatedPasswordOk returns a tuple with the RotatedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedPassword

`func (o *RotatedSecretUpdateSnowflake) SetRotatedPassword(v string)`

SetRotatedPassword sets RotatedPassword field to given value.

### HasRotatedPassword

`func (o *RotatedSecretUpdateSnowflake) HasRotatedPassword() bool`

HasRotatedPassword returns a boolean if a field has been set.

### GetRotatedUsername

`func (o *RotatedSecretUpdateSnowflake) GetRotatedUsername() string`

GetRotatedUsername returns the RotatedUsername field if non-nil, zero value otherwise.

### GetRotatedUsernameOk

`func (o *RotatedSecretUpdateSnowflake) GetRotatedUsernameOk() (*string, bool)`

GetRotatedUsernameOk returns a tuple with the RotatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedUsername

`func (o *RotatedSecretUpdateSnowflake) SetRotatedUsername(v string)`

SetRotatedUsername sets RotatedUsername field to given value.

### HasRotatedUsername

`func (o *RotatedSecretUpdateSnowflake) HasRotatedUsername() bool`

HasRotatedUsername returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretUpdateSnowflake) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretUpdateSnowflake) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretUpdateSnowflake) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretUpdateSnowflake) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretUpdateSnowflake) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretUpdateSnowflake) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretUpdateSnowflake) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretUpdateSnowflake) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetToken

`func (o *RotatedSecretUpdateSnowflake) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretUpdateSnowflake) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretUpdateSnowflake) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretUpdateSnowflake) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretUpdateSnowflake) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretUpdateSnowflake) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretUpdateSnowflake) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretUpdateSnowflake) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


