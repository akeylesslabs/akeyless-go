# TargetCreateGemini

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | API key for Gemini | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GeminiUrl** | Pointer to **string** | Base URL of the Gemini API | [optional] [default to "https://generativelanguage.googleapis.com"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateGemini

`func NewTargetCreateGemini(name string, ) *TargetCreateGemini`

NewTargetCreateGemini instantiates a new TargetCreateGemini object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateGeminiWithDefaults

`func NewTargetCreateGeminiWithDefaults() *TargetCreateGemini`

NewTargetCreateGeminiWithDefaults instantiates a new TargetCreateGemini object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TargetCreateGemini) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TargetCreateGemini) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TargetCreateGemini) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *TargetCreateGemini) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateGemini) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateGemini) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateGemini) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateGemini) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGeminiUrl

`func (o *TargetCreateGemini) GetGeminiUrl() string`

GetGeminiUrl returns the GeminiUrl field if non-nil, zero value otherwise.

### GetGeminiUrlOk

`func (o *TargetCreateGemini) GetGeminiUrlOk() (*string, bool)`

GetGeminiUrlOk returns a tuple with the GeminiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeminiUrl

`func (o *TargetCreateGemini) SetGeminiUrl(v string)`

SetGeminiUrl sets GeminiUrl field to given value.

### HasGeminiUrl

`func (o *TargetCreateGemini) HasGeminiUrl() bool`

HasGeminiUrl returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateGemini) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateGemini) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateGemini) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateGemini) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateGemini) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateGemini) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateGemini) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateGemini) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateGemini) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateGemini) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateGemini) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateGemini) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateGemini) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateGemini) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateGemini) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *TargetCreateGemini) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateGemini) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateGemini) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateGemini) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateGemini) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateGemini) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateGemini) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateGemini) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


