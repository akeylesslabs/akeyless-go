# UpdateSecretVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Multiline** | Pointer to **bool** | The provided value is a multiline value (separated by &#39;\\n&#39;) | [optional] 
**Name** | **string** | Secret name | 
**NewVersion** | Pointer to **bool** | Deprecated | [optional] 
**PasswordManagerCustomField** | Pointer to **map[string]string** | For Password Management use, additional fields | [optional] 
**PasswordManagerInjectUrl** | Pointer to **[]string** | For Password Management use, reflect the website context | [optional] 
**PasswordManagerPassword** | Pointer to **string** | For Password Management use, additional fields | [optional] 
**PasswordManagerUsername** | Pointer to **string** | For Password Management use | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Value** | **string** | The new secret value | 

## Methods

### NewUpdateSecretVal

`func NewUpdateSecretVal(name string, value string, ) *UpdateSecretVal`

NewUpdateSecretVal instantiates a new UpdateSecretVal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecretValWithDefaults

`func NewUpdateSecretValWithDefaults() *UpdateSecretVal`

NewUpdateSecretValWithDefaults instantiates a new UpdateSecretVal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *UpdateSecretVal) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *UpdateSecretVal) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *UpdateSecretVal) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *UpdateSecretVal) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetJson

`func (o *UpdateSecretVal) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateSecretVal) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateSecretVal) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateSecretVal) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateSecretVal) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateSecretVal) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateSecretVal) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateSecretVal) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateSecretVal) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateSecretVal) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateSecretVal) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateSecretVal) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMultiline

`func (o *UpdateSecretVal) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *UpdateSecretVal) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *UpdateSecretVal) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.

### HasMultiline

`func (o *UpdateSecretVal) HasMultiline() bool`

HasMultiline returns a boolean if a field has been set.

### GetName

`func (o *UpdateSecretVal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSecretVal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSecretVal) SetName(v string)`

SetName sets Name field to given value.


### GetNewVersion

`func (o *UpdateSecretVal) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *UpdateSecretVal) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *UpdateSecretVal) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *UpdateSecretVal) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetPasswordManagerCustomField

`func (o *UpdateSecretVal) GetPasswordManagerCustomField() map[string]string`

GetPasswordManagerCustomField returns the PasswordManagerCustomField field if non-nil, zero value otherwise.

### GetPasswordManagerCustomFieldOk

`func (o *UpdateSecretVal) GetPasswordManagerCustomFieldOk() (*map[string]string, bool)`

GetPasswordManagerCustomFieldOk returns a tuple with the PasswordManagerCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManagerCustomField

`func (o *UpdateSecretVal) SetPasswordManagerCustomField(v map[string]string)`

SetPasswordManagerCustomField sets PasswordManagerCustomField field to given value.

### HasPasswordManagerCustomField

`func (o *UpdateSecretVal) HasPasswordManagerCustomField() bool`

HasPasswordManagerCustomField returns a boolean if a field has been set.

### GetPasswordManagerInjectUrl

`func (o *UpdateSecretVal) GetPasswordManagerInjectUrl() []string`

GetPasswordManagerInjectUrl returns the PasswordManagerInjectUrl field if non-nil, zero value otherwise.

### GetPasswordManagerInjectUrlOk

`func (o *UpdateSecretVal) GetPasswordManagerInjectUrlOk() (*[]string, bool)`

GetPasswordManagerInjectUrlOk returns a tuple with the PasswordManagerInjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManagerInjectUrl

`func (o *UpdateSecretVal) SetPasswordManagerInjectUrl(v []string)`

SetPasswordManagerInjectUrl sets PasswordManagerInjectUrl field to given value.

### HasPasswordManagerInjectUrl

`func (o *UpdateSecretVal) HasPasswordManagerInjectUrl() bool`

HasPasswordManagerInjectUrl returns a boolean if a field has been set.

### GetPasswordManagerPassword

`func (o *UpdateSecretVal) GetPasswordManagerPassword() string`

GetPasswordManagerPassword returns the PasswordManagerPassword field if non-nil, zero value otherwise.

### GetPasswordManagerPasswordOk

`func (o *UpdateSecretVal) GetPasswordManagerPasswordOk() (*string, bool)`

GetPasswordManagerPasswordOk returns a tuple with the PasswordManagerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManagerPassword

`func (o *UpdateSecretVal) SetPasswordManagerPassword(v string)`

SetPasswordManagerPassword sets PasswordManagerPassword field to given value.

### HasPasswordManagerPassword

`func (o *UpdateSecretVal) HasPasswordManagerPassword() bool`

HasPasswordManagerPassword returns a boolean if a field has been set.

### GetPasswordManagerUsername

`func (o *UpdateSecretVal) GetPasswordManagerUsername() string`

GetPasswordManagerUsername returns the PasswordManagerUsername field if non-nil, zero value otherwise.

### GetPasswordManagerUsernameOk

`func (o *UpdateSecretVal) GetPasswordManagerUsernameOk() (*string, bool)`

GetPasswordManagerUsernameOk returns a tuple with the PasswordManagerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManagerUsername

`func (o *UpdateSecretVal) SetPasswordManagerUsername(v string)`

SetPasswordManagerUsername sets PasswordManagerUsername field to given value.

### HasPasswordManagerUsername

`func (o *UpdateSecretVal) HasPasswordManagerUsername() bool`

HasPasswordManagerUsername returns a boolean if a field has been set.

### GetToken

`func (o *UpdateSecretVal) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateSecretVal) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateSecretVal) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateSecretVal) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateSecretVal) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateSecretVal) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateSecretVal) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateSecretVal) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetValue

`func (o *UpdateSecretVal) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateSecretVal) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateSecretVal) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


