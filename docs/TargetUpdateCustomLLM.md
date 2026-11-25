# TargetUpdateCustomLLM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | API key for the provider | [optional] 
**CustomLlmUrl** | Pointer to **string** |  | [optional] [default to "the whole"]
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_comment"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LockOnRead** | Pointer to **string** | Lock this secret after each successful value read | [optional] 
**LockTtl** | Pointer to **string** | Lock TTL in minutes | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewComment** | Pointer to **string** | Deprecated - use description | [optional] [default to "default_comment"]
**NewName** | Pointer to **string** | New target name | [optional] 
**RotateOnUnlock** | Pointer to **string** | Rotate this secret after it is unlocked | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateCustomLLM

`func NewTargetUpdateCustomLLM(name string, ) *TargetUpdateCustomLLM`

NewTargetUpdateCustomLLM instantiates a new TargetUpdateCustomLLM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateCustomLLMWithDefaults

`func NewTargetUpdateCustomLLMWithDefaults() *TargetUpdateCustomLLM`

NewTargetUpdateCustomLLMWithDefaults instantiates a new TargetUpdateCustomLLM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TargetUpdateCustomLLM) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TargetUpdateCustomLLM) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TargetUpdateCustomLLM) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *TargetUpdateCustomLLM) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetCustomLlmUrl

`func (o *TargetUpdateCustomLLM) GetCustomLlmUrl() string`

GetCustomLlmUrl returns the CustomLlmUrl field if non-nil, zero value otherwise.

### GetCustomLlmUrlOk

`func (o *TargetUpdateCustomLLM) GetCustomLlmUrlOk() (*string, bool)`

GetCustomLlmUrlOk returns a tuple with the CustomLlmUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLlmUrl

`func (o *TargetUpdateCustomLLM) SetCustomLlmUrl(v string)`

SetCustomLlmUrl sets CustomLlmUrl field to given value.

### HasCustomLlmUrl

`func (o *TargetUpdateCustomLLM) HasCustomLlmUrl() bool`

HasCustomLlmUrl returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *TargetUpdateCustomLLM) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *TargetUpdateCustomLLM) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *TargetUpdateCustomLLM) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *TargetUpdateCustomLLM) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateCustomLLM) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateCustomLLM) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateCustomLLM) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateCustomLLM) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateCustomLLM) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateCustomLLM) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateCustomLLM) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateCustomLLM) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateCustomLLM) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateCustomLLM) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateCustomLLM) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateCustomLLM) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateCustomLLM) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateCustomLLM) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateCustomLLM) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateCustomLLM) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLockOnRead

`func (o *TargetUpdateCustomLLM) GetLockOnRead() string`

GetLockOnRead returns the LockOnRead field if non-nil, zero value otherwise.

### GetLockOnReadOk

`func (o *TargetUpdateCustomLLM) GetLockOnReadOk() (*string, bool)`

GetLockOnReadOk returns a tuple with the LockOnRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOnRead

`func (o *TargetUpdateCustomLLM) SetLockOnRead(v string)`

SetLockOnRead sets LockOnRead field to given value.

### HasLockOnRead

`func (o *TargetUpdateCustomLLM) HasLockOnRead() bool`

HasLockOnRead returns a boolean if a field has been set.

### GetLockTtl

`func (o *TargetUpdateCustomLLM) GetLockTtl() string`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *TargetUpdateCustomLLM) GetLockTtlOk() (*string, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *TargetUpdateCustomLLM) SetLockTtl(v string)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *TargetUpdateCustomLLM) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateCustomLLM) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateCustomLLM) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateCustomLLM) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateCustomLLM) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateCustomLLM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateCustomLLM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateCustomLLM) SetName(v string)`

SetName sets Name field to given value.


### GetNewComment

`func (o *TargetUpdateCustomLLM) GetNewComment() string`

GetNewComment returns the NewComment field if non-nil, zero value otherwise.

### GetNewCommentOk

`func (o *TargetUpdateCustomLLM) GetNewCommentOk() (*string, bool)`

GetNewCommentOk returns a tuple with the NewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewComment

`func (o *TargetUpdateCustomLLM) SetNewComment(v string)`

SetNewComment sets NewComment field to given value.

### HasNewComment

`func (o *TargetUpdateCustomLLM) HasNewComment() bool`

HasNewComment returns a boolean if a field has been set.

### GetNewName

`func (o *TargetUpdateCustomLLM) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateCustomLLM) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateCustomLLM) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateCustomLLM) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRotateOnUnlock

`func (o *TargetUpdateCustomLLM) GetRotateOnUnlock() string`

GetRotateOnUnlock returns the RotateOnUnlock field if non-nil, zero value otherwise.

### GetRotateOnUnlockOk

`func (o *TargetUpdateCustomLLM) GetRotateOnUnlockOk() (*string, bool)`

GetRotateOnUnlockOk returns a tuple with the RotateOnUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateOnUnlock

`func (o *TargetUpdateCustomLLM) SetRotateOnUnlock(v string)`

SetRotateOnUnlock sets RotateOnUnlock field to given value.

### HasRotateOnUnlock

`func (o *TargetUpdateCustomLLM) HasRotateOnUnlock() bool`

HasRotateOnUnlock returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateCustomLLM) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateCustomLLM) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateCustomLLM) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateCustomLLM) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateCustomLLM) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateCustomLLM) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateCustomLLM) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateCustomLLM) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


