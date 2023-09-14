# CreateGlobalSignAtlasTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | API Key of the GlobalSign Atlas account | 
**ApiSecret** | **string** | API Secret of the GlobalSign Atlas account | 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MtlsCertDataBase64** | Pointer to **string** | Mutual TLS Certificate contents of the GlobalSign Atlas account encoded in base64, either mtls-cert-file-path or mtls-cert-data-base64 must be supplied | [optional] 
**MtlsKeyDataBase64** | Pointer to **string** | Mutual TLS Key contents of the GlobalSign Atlas account encoded in base64, either mtls-key-file-path or mtls-data-base64 must be supplied | [optional] 
**Name** | **string** | Target name | 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateGlobalSignAtlasTarget

`func NewCreateGlobalSignAtlasTarget(apiKey string, apiSecret string, name string, ) *CreateGlobalSignAtlasTarget`

NewCreateGlobalSignAtlasTarget instantiates a new CreateGlobalSignAtlasTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGlobalSignAtlasTargetWithDefaults

`func NewCreateGlobalSignAtlasTargetWithDefaults() *CreateGlobalSignAtlasTarget`

NewCreateGlobalSignAtlasTargetWithDefaults instantiates a new CreateGlobalSignAtlasTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *CreateGlobalSignAtlasTarget) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateGlobalSignAtlasTarget) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateGlobalSignAtlasTarget) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiSecret

`func (o *CreateGlobalSignAtlasTarget) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *CreateGlobalSignAtlasTarget) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *CreateGlobalSignAtlasTarget) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.


### GetComment

`func (o *CreateGlobalSignAtlasTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateGlobalSignAtlasTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateGlobalSignAtlasTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateGlobalSignAtlasTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *CreateGlobalSignAtlasTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGlobalSignAtlasTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGlobalSignAtlasTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGlobalSignAtlasTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *CreateGlobalSignAtlasTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateGlobalSignAtlasTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateGlobalSignAtlasTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateGlobalSignAtlasTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *CreateGlobalSignAtlasTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateGlobalSignAtlasTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateGlobalSignAtlasTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateGlobalSignAtlasTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMtlsCertDataBase64

`func (o *CreateGlobalSignAtlasTarget) GetMtlsCertDataBase64() string`

GetMtlsCertDataBase64 returns the MtlsCertDataBase64 field if non-nil, zero value otherwise.

### GetMtlsCertDataBase64Ok

`func (o *CreateGlobalSignAtlasTarget) GetMtlsCertDataBase64Ok() (*string, bool)`

GetMtlsCertDataBase64Ok returns a tuple with the MtlsCertDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsCertDataBase64

`func (o *CreateGlobalSignAtlasTarget) SetMtlsCertDataBase64(v string)`

SetMtlsCertDataBase64 sets MtlsCertDataBase64 field to given value.

### HasMtlsCertDataBase64

`func (o *CreateGlobalSignAtlasTarget) HasMtlsCertDataBase64() bool`

HasMtlsCertDataBase64 returns a boolean if a field has been set.

### GetMtlsKeyDataBase64

`func (o *CreateGlobalSignAtlasTarget) GetMtlsKeyDataBase64() string`

GetMtlsKeyDataBase64 returns the MtlsKeyDataBase64 field if non-nil, zero value otherwise.

### GetMtlsKeyDataBase64Ok

`func (o *CreateGlobalSignAtlasTarget) GetMtlsKeyDataBase64Ok() (*string, bool)`

GetMtlsKeyDataBase64Ok returns a tuple with the MtlsKeyDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsKeyDataBase64

`func (o *CreateGlobalSignAtlasTarget) SetMtlsKeyDataBase64(v string)`

SetMtlsKeyDataBase64 sets MtlsKeyDataBase64 field to given value.

### HasMtlsKeyDataBase64

`func (o *CreateGlobalSignAtlasTarget) HasMtlsKeyDataBase64() bool`

HasMtlsKeyDataBase64 returns a boolean if a field has been set.

### GetName

`func (o *CreateGlobalSignAtlasTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGlobalSignAtlasTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGlobalSignAtlasTarget) SetName(v string)`

SetName sets Name field to given value.


### GetTimeout

`func (o *CreateGlobalSignAtlasTarget) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CreateGlobalSignAtlasTarget) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CreateGlobalSignAtlasTarget) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CreateGlobalSignAtlasTarget) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *CreateGlobalSignAtlasTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateGlobalSignAtlasTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateGlobalSignAtlasTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateGlobalSignAtlasTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateGlobalSignAtlasTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateGlobalSignAtlasTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateGlobalSignAtlasTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateGlobalSignAtlasTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


