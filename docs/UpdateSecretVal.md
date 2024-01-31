# UpdateSecretVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**CustomField** | Pointer to **map[string]string** | For Password Management use, additional fields | [optional] 
**InjectUrl** | Pointer to **[]string** | For Password Management use, reflect the website context | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LastVersion** | Pointer to **int32** | The last version number before the update | [optional] 
**Multiline** | Pointer to **bool** | The provided value is a multiline value (separated by &#39;\\n&#39;) | [optional] 
**Name** | **string** | Secret name | 
**NewVersion** | Pointer to **bool** | Deprecated | [optional] 
**Password** | Pointer to **string** | For Password Management use, additional fields | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | For Password Management use | [optional] 
**Value** | **string** | The secret value (only relevant for type &#39;generic&#39;) | 

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

### GetCustomField

`func (o *UpdateSecretVal) GetCustomField() map[string]string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *UpdateSecretVal) GetCustomFieldOk() (*map[string]string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *UpdateSecretVal) SetCustomField(v map[string]string)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *UpdateSecretVal) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.

### GetInjectUrl

`func (o *UpdateSecretVal) GetInjectUrl() []string`

GetInjectUrl returns the InjectUrl field if non-nil, zero value otherwise.

### GetInjectUrlOk

`func (o *UpdateSecretVal) GetInjectUrlOk() (*[]string, bool)`

GetInjectUrlOk returns a tuple with the InjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectUrl

`func (o *UpdateSecretVal) SetInjectUrl(v []string)`

SetInjectUrl sets InjectUrl field to given value.

### HasInjectUrl

`func (o *UpdateSecretVal) HasInjectUrl() bool`

HasInjectUrl returns a boolean if a field has been set.

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

### GetLastVersion

`func (o *UpdateSecretVal) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *UpdateSecretVal) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *UpdateSecretVal) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *UpdateSecretVal) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

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

### GetPassword

`func (o *UpdateSecretVal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateSecretVal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateSecretVal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateSecretVal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

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

### GetUsername

`func (o *UpdateSecretVal) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateSecretVal) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateSecretVal) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateSecretVal) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

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


