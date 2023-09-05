# ImportPasswords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "personal"]
**Format** | Pointer to **string** | Password format type [LastPass/Chrome/Firefox] | [optional] [default to "LastPass"]
**ImportPath** | **string** | File path | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**TargetFolder** | Pointer to **string** | Target folder for imported passwords | [optional] [default to "/"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewImportPasswords

`func NewImportPasswords(importPath string, ) *ImportPasswords`

NewImportPasswords instantiates a new ImportPasswords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportPasswordsWithDefaults

`func NewImportPasswordsWithDefaults() *ImportPasswords`

NewImportPasswordsWithDefaults instantiates a new ImportPasswords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *ImportPasswords) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *ImportPasswords) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *ImportPasswords) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *ImportPasswords) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetFormat

`func (o *ImportPasswords) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ImportPasswords) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ImportPasswords) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ImportPasswords) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetImportPath

`func (o *ImportPasswords) GetImportPath() string`

GetImportPath returns the ImportPath field if non-nil, zero value otherwise.

### GetImportPathOk

`func (o *ImportPasswords) GetImportPathOk() (*string, bool)`

GetImportPathOk returns a tuple with the ImportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPath

`func (o *ImportPasswords) SetImportPath(v string)`

SetImportPath sets ImportPath field to given value.


### GetJson

`func (o *ImportPasswords) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ImportPasswords) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ImportPasswords) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ImportPasswords) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetProtectionKey

`func (o *ImportPasswords) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *ImportPasswords) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *ImportPasswords) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *ImportPasswords) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTargetFolder

`func (o *ImportPasswords) GetTargetFolder() string`

GetTargetFolder returns the TargetFolder field if non-nil, zero value otherwise.

### GetTargetFolderOk

`func (o *ImportPasswords) GetTargetFolderOk() (*string, bool)`

GetTargetFolderOk returns a tuple with the TargetFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolder

`func (o *ImportPasswords) SetTargetFolder(v string)`

SetTargetFolder sets TargetFolder field to given value.

### HasTargetFolder

`func (o *ImportPasswords) HasTargetFolder() bool`

HasTargetFolder returns a boolean if a field has been set.

### GetToken

`func (o *ImportPasswords) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ImportPasswords) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ImportPasswords) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ImportPasswords) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ImportPasswords) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ImportPasswords) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ImportPasswords) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ImportPasswords) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


