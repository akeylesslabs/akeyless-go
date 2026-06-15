# RotatedSecretUpdateHashiVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**AutoRotate** | Pointer to **string** |  | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_metadata"]
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Rotated secret name | 
**NewName** | Pointer to **string** | New item name | [optional] 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**RotationEventIn** | Pointer to **[]string** | How many days before the rotation of the item would you like to be notified | [optional] 
**RotationHour** | Pointer to **int32** |  | [optional] 
**RotationInterval** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseCapitalLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one uppercase character from the ISO basic Latin alphabet (A to Z). [true/false] | [optional] 
**UseLowerLetters** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one lowercase character from the ISO basic Latin alphabet (a to z). [true/false] | [optional] 
**UseNumbers** | Pointer to **string** | Specifies whether the generated temporary password must contain at least one numeric character (0 to 9). [true/false] | [optional] 
**UseSpecialCharacters** | Pointer to **string** |  | [optional] 

## Methods

### NewRotatedSecretUpdateHashiVault

`func NewRotatedSecretUpdateHashiVault(name string, ) *RotatedSecretUpdateHashiVault`

NewRotatedSecretUpdateHashiVault instantiates a new RotatedSecretUpdateHashiVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretUpdateHashiVaultWithDefaults

`func NewRotatedSecretUpdateHashiVaultWithDefaults() *RotatedSecretUpdateHashiVault`

NewRotatedSecretUpdateHashiVaultWithDefaults instantiates a new RotatedSecretUpdateHashiVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *RotatedSecretUpdateHashiVault) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *RotatedSecretUpdateHashiVault) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *RotatedSecretUpdateHashiVault) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *RotatedSecretUpdateHashiVault) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAutoRotate

`func (o *RotatedSecretUpdateHashiVault) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *RotatedSecretUpdateHashiVault) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *RotatedSecretUpdateHashiVault) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *RotatedSecretUpdateHashiVault) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *RotatedSecretUpdateHashiVault) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *RotatedSecretUpdateHashiVault) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *RotatedSecretUpdateHashiVault) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *RotatedSecretUpdateHashiVault) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *RotatedSecretUpdateHashiVault) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RotatedSecretUpdateHashiVault) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RotatedSecretUpdateHashiVault) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RotatedSecretUpdateHashiVault) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputRule

`func (o *RotatedSecretUpdateHashiVault) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *RotatedSecretUpdateHashiVault) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *RotatedSecretUpdateHashiVault) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *RotatedSecretUpdateHashiVault) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *RotatedSecretUpdateHashiVault) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *RotatedSecretUpdateHashiVault) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *RotatedSecretUpdateHashiVault) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *RotatedSecretUpdateHashiVault) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretUpdateHashiVault) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretUpdateHashiVault) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretUpdateHashiVault) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretUpdateHashiVault) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *RotatedSecretUpdateHashiVault) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *RotatedSecretUpdateHashiVault) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *RotatedSecretUpdateHashiVault) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *RotatedSecretUpdateHashiVault) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *RotatedSecretUpdateHashiVault) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RotatedSecretUpdateHashiVault) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RotatedSecretUpdateHashiVault) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RotatedSecretUpdateHashiVault) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *RotatedSecretUpdateHashiVault) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *RotatedSecretUpdateHashiVault) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *RotatedSecretUpdateHashiVault) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *RotatedSecretUpdateHashiVault) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretUpdateHashiVault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretUpdateHashiVault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretUpdateHashiVault) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *RotatedSecretUpdateHashiVault) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *RotatedSecretUpdateHashiVault) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *RotatedSecretUpdateHashiVault) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *RotatedSecretUpdateHashiVault) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetOutputRule

`func (o *RotatedSecretUpdateHashiVault) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *RotatedSecretUpdateHashiVault) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *RotatedSecretUpdateHashiVault) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *RotatedSecretUpdateHashiVault) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetPasswordLength

`func (o *RotatedSecretUpdateHashiVault) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *RotatedSecretUpdateHashiVault) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *RotatedSecretUpdateHashiVault) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *RotatedSecretUpdateHashiVault) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetRmTag

`func (o *RotatedSecretUpdateHashiVault) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *RotatedSecretUpdateHashiVault) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *RotatedSecretUpdateHashiVault) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *RotatedSecretUpdateHashiVault) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotationEventIn

`func (o *RotatedSecretUpdateHashiVault) GetRotationEventIn() []string`

GetRotationEventIn returns the RotationEventIn field if non-nil, zero value otherwise.

### GetRotationEventInOk

`func (o *RotatedSecretUpdateHashiVault) GetRotationEventInOk() (*[]string, bool)`

GetRotationEventInOk returns a tuple with the RotationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationEventIn

`func (o *RotatedSecretUpdateHashiVault) SetRotationEventIn(v []string)`

SetRotationEventIn sets RotationEventIn field to given value.

### HasRotationEventIn

`func (o *RotatedSecretUpdateHashiVault) HasRotationEventIn() bool`

HasRotationEventIn returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretUpdateHashiVault) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretUpdateHashiVault) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretUpdateHashiVault) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretUpdateHashiVault) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *RotatedSecretUpdateHashiVault) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *RotatedSecretUpdateHashiVault) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *RotatedSecretUpdateHashiVault) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *RotatedSecretUpdateHashiVault) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetToken

`func (o *RotatedSecretUpdateHashiVault) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretUpdateHashiVault) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretUpdateHashiVault) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretUpdateHashiVault) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretUpdateHashiVault) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretUpdateHashiVault) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretUpdateHashiVault) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretUpdateHashiVault) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseCapitalLetters

`func (o *RotatedSecretUpdateHashiVault) GetUseCapitalLetters() string`

GetUseCapitalLetters returns the UseCapitalLetters field if non-nil, zero value otherwise.

### GetUseCapitalLettersOk

`func (o *RotatedSecretUpdateHashiVault) GetUseCapitalLettersOk() (*string, bool)`

GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCapitalLetters

`func (o *RotatedSecretUpdateHashiVault) SetUseCapitalLetters(v string)`

SetUseCapitalLetters sets UseCapitalLetters field to given value.

### HasUseCapitalLetters

`func (o *RotatedSecretUpdateHashiVault) HasUseCapitalLetters() bool`

HasUseCapitalLetters returns a boolean if a field has been set.

### GetUseLowerLetters

`func (o *RotatedSecretUpdateHashiVault) GetUseLowerLetters() string`

GetUseLowerLetters returns the UseLowerLetters field if non-nil, zero value otherwise.

### GetUseLowerLettersOk

`func (o *RotatedSecretUpdateHashiVault) GetUseLowerLettersOk() (*string, bool)`

GetUseLowerLettersOk returns a tuple with the UseLowerLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLowerLetters

`func (o *RotatedSecretUpdateHashiVault) SetUseLowerLetters(v string)`

SetUseLowerLetters sets UseLowerLetters field to given value.

### HasUseLowerLetters

`func (o *RotatedSecretUpdateHashiVault) HasUseLowerLetters() bool`

HasUseLowerLetters returns a boolean if a field has been set.

### GetUseNumbers

`func (o *RotatedSecretUpdateHashiVault) GetUseNumbers() string`

GetUseNumbers returns the UseNumbers field if non-nil, zero value otherwise.

### GetUseNumbersOk

`func (o *RotatedSecretUpdateHashiVault) GetUseNumbersOk() (*string, bool)`

GetUseNumbersOk returns a tuple with the UseNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumbers

`func (o *RotatedSecretUpdateHashiVault) SetUseNumbers(v string)`

SetUseNumbers sets UseNumbers field to given value.

### HasUseNumbers

`func (o *RotatedSecretUpdateHashiVault) HasUseNumbers() bool`

HasUseNumbers returns a boolean if a field has been set.

### GetUseSpecialCharacters

`func (o *RotatedSecretUpdateHashiVault) GetUseSpecialCharacters() string`

GetUseSpecialCharacters returns the UseSpecialCharacters field if non-nil, zero value otherwise.

### GetUseSpecialCharactersOk

`func (o *RotatedSecretUpdateHashiVault) GetUseSpecialCharactersOk() (*string, bool)`

GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpecialCharacters

`func (o *RotatedSecretUpdateHashiVault) SetUseSpecialCharacters(v string)`

SetUseSpecialCharacters sets UseSpecialCharacters field to given value.

### HasUseSpecialCharacters

`func (o *RotatedSecretUpdateHashiVault) HasUseSpecialCharacters() bool`

HasUseSpecialCharacters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


