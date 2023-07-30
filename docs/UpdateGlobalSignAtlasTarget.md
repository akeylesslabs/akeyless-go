# UpdateGlobalSignAtlasTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | API Key of the GlobalSign Atlas account | 
**ApiSecret** | **string** | API Secret of the GlobalSign Atlas account | 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MtlsCertDataBase64** | Pointer to **string** | Mutual TLS Certificate contents of the GlobalSign Atlas account encoded in base64, either mtls-cert-file-path or mtls-cert-data-base64 must be supplied | [optional] 
**MtlsKeyDataBase64** | Pointer to **string** | Mutual TLS Key contents of the GlobalSign Atlas account encoded in base64, either mtls-key-file-path or mtls-data-base64 must be supplied | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 

## Methods

### NewUpdateGlobalSignAtlasTarget

`func NewUpdateGlobalSignAtlasTarget(apiKey string, apiSecret string, name string, ) *UpdateGlobalSignAtlasTarget`

NewUpdateGlobalSignAtlasTarget instantiates a new UpdateGlobalSignAtlasTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGlobalSignAtlasTargetWithDefaults

`func NewUpdateGlobalSignAtlasTargetWithDefaults() *UpdateGlobalSignAtlasTarget`

NewUpdateGlobalSignAtlasTargetWithDefaults instantiates a new UpdateGlobalSignAtlasTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *UpdateGlobalSignAtlasTarget) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateGlobalSignAtlasTarget) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateGlobalSignAtlasTarget) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiSecret

`func (o *UpdateGlobalSignAtlasTarget) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *UpdateGlobalSignAtlasTarget) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *UpdateGlobalSignAtlasTarget) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.


### GetComment

`func (o *UpdateGlobalSignAtlasTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateGlobalSignAtlasTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateGlobalSignAtlasTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateGlobalSignAtlasTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGlobalSignAtlasTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGlobalSignAtlasTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGlobalSignAtlasTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGlobalSignAtlasTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *UpdateGlobalSignAtlasTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateGlobalSignAtlasTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateGlobalSignAtlasTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateGlobalSignAtlasTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateGlobalSignAtlasTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateGlobalSignAtlasTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateGlobalSignAtlasTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateGlobalSignAtlasTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateGlobalSignAtlasTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateGlobalSignAtlasTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateGlobalSignAtlasTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateGlobalSignAtlasTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMtlsCertDataBase64

`func (o *UpdateGlobalSignAtlasTarget) GetMtlsCertDataBase64() string`

GetMtlsCertDataBase64 returns the MtlsCertDataBase64 field if non-nil, zero value otherwise.

### GetMtlsCertDataBase64Ok

`func (o *UpdateGlobalSignAtlasTarget) GetMtlsCertDataBase64Ok() (*string, bool)`

GetMtlsCertDataBase64Ok returns a tuple with the MtlsCertDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsCertDataBase64

`func (o *UpdateGlobalSignAtlasTarget) SetMtlsCertDataBase64(v string)`

SetMtlsCertDataBase64 sets MtlsCertDataBase64 field to given value.

### HasMtlsCertDataBase64

`func (o *UpdateGlobalSignAtlasTarget) HasMtlsCertDataBase64() bool`

HasMtlsCertDataBase64 returns a boolean if a field has been set.

### GetMtlsKeyDataBase64

`func (o *UpdateGlobalSignAtlasTarget) GetMtlsKeyDataBase64() string`

GetMtlsKeyDataBase64 returns the MtlsKeyDataBase64 field if non-nil, zero value otherwise.

### GetMtlsKeyDataBase64Ok

`func (o *UpdateGlobalSignAtlasTarget) GetMtlsKeyDataBase64Ok() (*string, bool)`

GetMtlsKeyDataBase64Ok returns a tuple with the MtlsKeyDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsKeyDataBase64

`func (o *UpdateGlobalSignAtlasTarget) SetMtlsKeyDataBase64(v string)`

SetMtlsKeyDataBase64 sets MtlsKeyDataBase64 field to given value.

### HasMtlsKeyDataBase64

`func (o *UpdateGlobalSignAtlasTarget) HasMtlsKeyDataBase64() bool`

HasMtlsKeyDataBase64 returns a boolean if a field has been set.

### GetName

`func (o *UpdateGlobalSignAtlasTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGlobalSignAtlasTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGlobalSignAtlasTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateGlobalSignAtlasTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateGlobalSignAtlasTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateGlobalSignAtlasTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateGlobalSignAtlasTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetTimeout

`func (o *UpdateGlobalSignAtlasTarget) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UpdateGlobalSignAtlasTarget) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UpdateGlobalSignAtlasTarget) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UpdateGlobalSignAtlasTarget) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *UpdateGlobalSignAtlasTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateGlobalSignAtlasTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateGlobalSignAtlasTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateGlobalSignAtlasTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateGlobalSignAtlasTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateGlobalSignAtlasTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateGlobalSignAtlasTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateGlobalSignAtlasTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateGlobalSignAtlasTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateGlobalSignAtlasTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateGlobalSignAtlasTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateGlobalSignAtlasTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


