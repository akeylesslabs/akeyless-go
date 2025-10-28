# TargetUpdateGemini

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | API key for Gemini | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_comment"]
**GeminiUrl** | Pointer to **string** | Base URL of the Gemini API | [optional] [default to "https://generativelanguage.googleapis.com"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewComment** | Pointer to **string** | Deprecated - use description | [optional] [default to "default_comment"]
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateGemini

`func NewTargetUpdateGemini(name string, ) *TargetUpdateGemini`

NewTargetUpdateGemini instantiates a new TargetUpdateGemini object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateGeminiWithDefaults

`func NewTargetUpdateGeminiWithDefaults() *TargetUpdateGemini`

NewTargetUpdateGeminiWithDefaults instantiates a new TargetUpdateGemini object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TargetUpdateGemini) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TargetUpdateGemini) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TargetUpdateGemini) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *TargetUpdateGemini) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateGemini) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateGemini) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateGemini) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateGemini) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGeminiUrl

`func (o *TargetUpdateGemini) GetGeminiUrl() string`

GetGeminiUrl returns the GeminiUrl field if non-nil, zero value otherwise.

### GetGeminiUrlOk

`func (o *TargetUpdateGemini) GetGeminiUrlOk() (*string, bool)`

GetGeminiUrlOk returns a tuple with the GeminiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeminiUrl

`func (o *TargetUpdateGemini) SetGeminiUrl(v string)`

SetGeminiUrl sets GeminiUrl field to given value.

### HasGeminiUrl

`func (o *TargetUpdateGemini) HasGeminiUrl() bool`

HasGeminiUrl returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateGemini) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateGemini) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateGemini) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateGemini) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateGemini) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateGemini) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateGemini) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateGemini) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateGemini) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateGemini) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateGemini) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateGemini) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateGemini) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateGemini) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateGemini) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateGemini) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateGemini) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateGemini) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateGemini) SetName(v string)`

SetName sets Name field to given value.


### GetNewComment

`func (o *TargetUpdateGemini) GetNewComment() string`

GetNewComment returns the NewComment field if non-nil, zero value otherwise.

### GetNewCommentOk

`func (o *TargetUpdateGemini) GetNewCommentOk() (*string, bool)`

GetNewCommentOk returns a tuple with the NewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewComment

`func (o *TargetUpdateGemini) SetNewComment(v string)`

SetNewComment sets NewComment field to given value.

### HasNewComment

`func (o *TargetUpdateGemini) HasNewComment() bool`

HasNewComment returns a boolean if a field has been set.

### GetNewName

`func (o *TargetUpdateGemini) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateGemini) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateGemini) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateGemini) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateGemini) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateGemini) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateGemini) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateGemini) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateGemini) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateGemini) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateGemini) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateGemini) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


