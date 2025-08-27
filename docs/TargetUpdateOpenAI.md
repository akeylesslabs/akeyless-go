# TargetUpdateOpenAI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | API key for OpenAI | [optional] 
**ApiKeyId** | Pointer to **string** | API key ID | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_comment"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Model** | Pointer to **string** | Default model to use with OpenAI | [optional] 
**Name** | **string** | Target name | 
**NewComment** | Pointer to **string** | Deprecated - use description | [optional] [default to "default_comment"]
**NewName** | Pointer to **string** | New target name | [optional] 
**OpenaiUrl** | Pointer to **string** | Base URL of the OpenAI API | [optional] [default to "https://api.openai.com/v1"]
**OrganizationId** | Pointer to **string** | Organization ID | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateOpenAI

`func NewTargetUpdateOpenAI(name string, ) *TargetUpdateOpenAI`

NewTargetUpdateOpenAI instantiates a new TargetUpdateOpenAI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateOpenAIWithDefaults

`func NewTargetUpdateOpenAIWithDefaults() *TargetUpdateOpenAI`

NewTargetUpdateOpenAIWithDefaults instantiates a new TargetUpdateOpenAI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TargetUpdateOpenAI) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TargetUpdateOpenAI) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TargetUpdateOpenAI) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *TargetUpdateOpenAI) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiKeyId

`func (o *TargetUpdateOpenAI) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *TargetUpdateOpenAI) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *TargetUpdateOpenAI) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *TargetUpdateOpenAI) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateOpenAI) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateOpenAI) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateOpenAI) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateOpenAI) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateOpenAI) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateOpenAI) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateOpenAI) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateOpenAI) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateOpenAI) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateOpenAI) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateOpenAI) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateOpenAI) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateOpenAI) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateOpenAI) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateOpenAI) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateOpenAI) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateOpenAI) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateOpenAI) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateOpenAI) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateOpenAI) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetModel

`func (o *TargetUpdateOpenAI) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TargetUpdateOpenAI) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TargetUpdateOpenAI) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TargetUpdateOpenAI) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateOpenAI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateOpenAI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateOpenAI) SetName(v string)`

SetName sets Name field to given value.


### GetNewComment

`func (o *TargetUpdateOpenAI) GetNewComment() string`

GetNewComment returns the NewComment field if non-nil, zero value otherwise.

### GetNewCommentOk

`func (o *TargetUpdateOpenAI) GetNewCommentOk() (*string, bool)`

GetNewCommentOk returns a tuple with the NewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewComment

`func (o *TargetUpdateOpenAI) SetNewComment(v string)`

SetNewComment sets NewComment field to given value.

### HasNewComment

`func (o *TargetUpdateOpenAI) HasNewComment() bool`

HasNewComment returns a boolean if a field has been set.

### GetNewName

`func (o *TargetUpdateOpenAI) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateOpenAI) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateOpenAI) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateOpenAI) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetOpenaiUrl

`func (o *TargetUpdateOpenAI) GetOpenaiUrl() string`

GetOpenaiUrl returns the OpenaiUrl field if non-nil, zero value otherwise.

### GetOpenaiUrlOk

`func (o *TargetUpdateOpenAI) GetOpenaiUrlOk() (*string, bool)`

GetOpenaiUrlOk returns a tuple with the OpenaiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenaiUrl

`func (o *TargetUpdateOpenAI) SetOpenaiUrl(v string)`

SetOpenaiUrl sets OpenaiUrl field to given value.

### HasOpenaiUrl

`func (o *TargetUpdateOpenAI) HasOpenaiUrl() bool`

HasOpenaiUrl returns a boolean if a field has been set.

### GetOrganizationId

`func (o *TargetUpdateOpenAI) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *TargetUpdateOpenAI) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *TargetUpdateOpenAI) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *TargetUpdateOpenAI) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateOpenAI) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateOpenAI) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateOpenAI) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateOpenAI) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateOpenAI) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateOpenAI) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateOpenAI) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateOpenAI) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


