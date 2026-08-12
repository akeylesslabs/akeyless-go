# OpenAITargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**ApiKeyId** | Pointer to **string** |  | [optional] 
**AuthMode** | Pointer to **string** | AuthMode selects how this target authenticates. Empty (default) uses ApiKey as a static bearer token against BaseURL, matching all pre-existing behavior. OpenAIAuthModeChatGPTOAuth instead uses the OAuth* fields below. | [optional] 
**OauthAccessToken** | Pointer to **string** | OAuthAccessToken is the current ChatGPT-issued access token (the &#x60;tokens.access_token&#x60; field of the customer&#39;s local auth.json). Akeyless refreshes this automatically; do not treat it as long-lived. | [optional] 
**OauthAccountId** | Pointer to **string** | OAuthAccountID is the ChatGPT workspace/account id (&#x60;tokens.account_id&#x60; in auth.json), required on every request to the ChatGPT backend. | [optional] 
**OauthLastRefresh** | Pointer to **string** | OAuthLastRefresh is the RFC3339 timestamp of the last successful Akeyless-performed refresh; used as a fallback expiry heuristic when the access token&#39;s JWT exp claim can&#39;t be parsed. | [optional] 
**OauthRefreshToken** | Pointer to **string** | OAuthRefreshToken mints new access tokens. It rotates on every refresh - Akeyless persists the new value after each successful refresh, so the previous value becomes invalid. | [optional] 
**OpenaiUrl** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenAITargetDetails

`func NewOpenAITargetDetails() *OpenAITargetDetails`

NewOpenAITargetDetails instantiates a new OpenAITargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenAITargetDetailsWithDefaults

`func NewOpenAITargetDetailsWithDefaults() *OpenAITargetDetails`

NewOpenAITargetDetailsWithDefaults instantiates a new OpenAITargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *OpenAITargetDetails) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OpenAITargetDetails) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OpenAITargetDetails) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *OpenAITargetDetails) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiKeyId

`func (o *OpenAITargetDetails) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *OpenAITargetDetails) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *OpenAITargetDetails) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *OpenAITargetDetails) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetAuthMode

`func (o *OpenAITargetDetails) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *OpenAITargetDetails) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *OpenAITargetDetails) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *OpenAITargetDetails) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetOauthAccessToken

`func (o *OpenAITargetDetails) GetOauthAccessToken() string`

GetOauthAccessToken returns the OauthAccessToken field if non-nil, zero value otherwise.

### GetOauthAccessTokenOk

`func (o *OpenAITargetDetails) GetOauthAccessTokenOk() (*string, bool)`

GetOauthAccessTokenOk returns a tuple with the OauthAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAccessToken

`func (o *OpenAITargetDetails) SetOauthAccessToken(v string)`

SetOauthAccessToken sets OauthAccessToken field to given value.

### HasOauthAccessToken

`func (o *OpenAITargetDetails) HasOauthAccessToken() bool`

HasOauthAccessToken returns a boolean if a field has been set.

### GetOauthAccountId

`func (o *OpenAITargetDetails) GetOauthAccountId() string`

GetOauthAccountId returns the OauthAccountId field if non-nil, zero value otherwise.

### GetOauthAccountIdOk

`func (o *OpenAITargetDetails) GetOauthAccountIdOk() (*string, bool)`

GetOauthAccountIdOk returns a tuple with the OauthAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAccountId

`func (o *OpenAITargetDetails) SetOauthAccountId(v string)`

SetOauthAccountId sets OauthAccountId field to given value.

### HasOauthAccountId

`func (o *OpenAITargetDetails) HasOauthAccountId() bool`

HasOauthAccountId returns a boolean if a field has been set.

### GetOauthLastRefresh

`func (o *OpenAITargetDetails) GetOauthLastRefresh() string`

GetOauthLastRefresh returns the OauthLastRefresh field if non-nil, zero value otherwise.

### GetOauthLastRefreshOk

`func (o *OpenAITargetDetails) GetOauthLastRefreshOk() (*string, bool)`

GetOauthLastRefreshOk returns a tuple with the OauthLastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthLastRefresh

`func (o *OpenAITargetDetails) SetOauthLastRefresh(v string)`

SetOauthLastRefresh sets OauthLastRefresh field to given value.

### HasOauthLastRefresh

`func (o *OpenAITargetDetails) HasOauthLastRefresh() bool`

HasOauthLastRefresh returns a boolean if a field has been set.

### GetOauthRefreshToken

`func (o *OpenAITargetDetails) GetOauthRefreshToken() string`

GetOauthRefreshToken returns the OauthRefreshToken field if non-nil, zero value otherwise.

### GetOauthRefreshTokenOk

`func (o *OpenAITargetDetails) GetOauthRefreshTokenOk() (*string, bool)`

GetOauthRefreshTokenOk returns a tuple with the OauthRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthRefreshToken

`func (o *OpenAITargetDetails) SetOauthRefreshToken(v string)`

SetOauthRefreshToken sets OauthRefreshToken field to given value.

### HasOauthRefreshToken

`func (o *OpenAITargetDetails) HasOauthRefreshToken() bool`

HasOauthRefreshToken returns a boolean if a field has been set.

### GetOpenaiUrl

`func (o *OpenAITargetDetails) GetOpenaiUrl() string`

GetOpenaiUrl returns the OpenaiUrl field if non-nil, zero value otherwise.

### GetOpenaiUrlOk

`func (o *OpenAITargetDetails) GetOpenaiUrlOk() (*string, bool)`

GetOpenaiUrlOk returns a tuple with the OpenaiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenaiUrl

`func (o *OpenAITargetDetails) SetOpenaiUrl(v string)`

SetOpenaiUrl sets OpenaiUrl field to given value.

### HasOpenaiUrl

`func (o *OpenAITargetDetails) HasOpenaiUrl() bool`

HasOpenaiUrl returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OpenAITargetDetails) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OpenAITargetDetails) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OpenAITargetDetails) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OpenAITargetDetails) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectId

`func (o *OpenAITargetDetails) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OpenAITargetDetails) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OpenAITargetDetails) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *OpenAITargetDetails) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


