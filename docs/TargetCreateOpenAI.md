# TargetCreateOpenAI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | API key for OpenAI | [optional] 
**ApiKeyId** | Pointer to **string** | API key ID | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Model** | Pointer to **string** | Default model to use with OpenAI | [optional] 
**Name** | **string** | Target name | 
**OpenaiUrl** | Pointer to **string** | Base URL of the OpenAI API | [optional] [default to "https://api.openai.com/v1"]
**OrganizationId** | Pointer to **string** | Organization ID | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateOpenAI

`func NewTargetCreateOpenAI(name string, ) *TargetCreateOpenAI`

NewTargetCreateOpenAI instantiates a new TargetCreateOpenAI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateOpenAIWithDefaults

`func NewTargetCreateOpenAIWithDefaults() *TargetCreateOpenAI`

NewTargetCreateOpenAIWithDefaults instantiates a new TargetCreateOpenAI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TargetCreateOpenAI) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TargetCreateOpenAI) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TargetCreateOpenAI) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *TargetCreateOpenAI) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiKeyId

`func (o *TargetCreateOpenAI) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *TargetCreateOpenAI) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *TargetCreateOpenAI) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *TargetCreateOpenAI) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateOpenAI) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateOpenAI) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateOpenAI) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateOpenAI) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateOpenAI) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateOpenAI) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateOpenAI) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateOpenAI) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateOpenAI) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateOpenAI) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateOpenAI) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateOpenAI) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateOpenAI) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateOpenAI) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateOpenAI) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateOpenAI) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetModel

`func (o *TargetCreateOpenAI) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TargetCreateOpenAI) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TargetCreateOpenAI) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TargetCreateOpenAI) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateOpenAI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateOpenAI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateOpenAI) SetName(v string)`

SetName sets Name field to given value.


### GetOpenaiUrl

`func (o *TargetCreateOpenAI) GetOpenaiUrl() string`

GetOpenaiUrl returns the OpenaiUrl field if non-nil, zero value otherwise.

### GetOpenaiUrlOk

`func (o *TargetCreateOpenAI) GetOpenaiUrlOk() (*string, bool)`

GetOpenaiUrlOk returns a tuple with the OpenaiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenaiUrl

`func (o *TargetCreateOpenAI) SetOpenaiUrl(v string)`

SetOpenaiUrl sets OpenaiUrl field to given value.

### HasOpenaiUrl

`func (o *TargetCreateOpenAI) HasOpenaiUrl() bool`

HasOpenaiUrl returns a boolean if a field has been set.

### GetOrganizationId

`func (o *TargetCreateOpenAI) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *TargetCreateOpenAI) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *TargetCreateOpenAI) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *TargetCreateOpenAI) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateOpenAI) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateOpenAI) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateOpenAI) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateOpenAI) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateOpenAI) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateOpenAI) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateOpenAI) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateOpenAI) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


