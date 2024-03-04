# RotatedSecretUpdateDockerhub

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
**RotationHour** | Pointer to **int32** | The Hour of the rotation in UTC | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (1-365) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotatedSecretUpdateDockerhub

`func NewRotatedSecretUpdateDockerhub(name string, ) *RotatedSecretUpdateDockerhub`

NewRotatedSecretUpdateDockerhub instantiates a new RotatedSecretUpdateDockerhub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretUpdateDockerhubWithDefaults

`func NewRotatedSecretUpdateDockerhubWithDefaults() *RotatedSecretUpdateDockerhub`

NewRotatedSecretUpdateDockerhubWithDefaults instantiates a new RotatedSecretUpdateDockerhub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *RotatedSecretUpdateDockerhub) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *RotatedSecretUpdateDockerhub) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *RotatedSecretUpdateDockerhub) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *RotatedSecretUpdateDockerhub) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAuthenticationCredentials

`func (o *RotatedSecretUpdateDockerhub) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *RotatedSecretUpdateDockerhub) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *RotatedSecretUpdateDockerhub) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *RotatedSecretUpdateDockerhub) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretUpdateDockerhub) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretUpdateDockerhub) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretUpdateDockerhub) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretUpdateDockerhub) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretUpdateDockerhub) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretUpdateDockerhub) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretUpdateDockerhub) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretUpdateDockerhub) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretUpdateDockerhub) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretUpdateDockerhub) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretUpdateDockerhub) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretUpdateDockerhub) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretUpdateDockerhub) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretUpdateDockerhub) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretUpdateDockerhub) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretUpdateDockerhub) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *RotatedSecretUpdateDockerhub) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *RotatedSecretUpdateDockerhub) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *RotatedSecretUpdateDockerhub) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *RotatedSecretUpdateDockerhub) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretUpdateDockerhub) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretUpdateDockerhub) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretUpdateDockerhub) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretUpdateDockerhub) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretUpdateDockerhub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretUpdateDockerhub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretUpdateDockerhub) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *RotatedSecretUpdateDockerhub) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *RotatedSecretUpdateDockerhub) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *RotatedSecretUpdateDockerhub) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *RotatedSecretUpdateDockerhub) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *RotatedSecretUpdateDockerhub) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretUpdateDockerhub) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretUpdateDockerhub) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretUpdateDockerhub) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRmTag

`func (o *RotatedSecretUpdateDockerhub) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *RotatedSecretUpdateDockerhub) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *RotatedSecretUpdateDockerhub) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *RotatedSecretUpdateDockerhub) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretUpdateDockerhub) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretUpdateDockerhub) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretUpdateDockerhub) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretUpdateDockerhub) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretUpdateDockerhub) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretUpdateDockerhub) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretUpdateDockerhub) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretUpdateDockerhub) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetToken

`func (o *RotatedSecretUpdateDockerhub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretUpdateDockerhub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretUpdateDockerhub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretUpdateDockerhub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretUpdateDockerhub) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretUpdateDockerhub) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretUpdateDockerhub) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretUpdateDockerhub) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


