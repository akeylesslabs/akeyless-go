# TargetUpdateGlobalSignAtlas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | API Key of the GlobalSign Atlas account | 
**ApiSecret** | **string** | API Secret of the GlobalSign Atlas account | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**MtlsCertDataBase64** | Pointer to **string** | Mutual TLS Certificate contents of the GlobalSign Atlas account encoded in base64, either mtls-cert-file-path or mtls-cert-data-base64 must be supplied | [optional] 
**MtlsKeyDataBase64** | Pointer to **string** | Mutual TLS Key contents of the GlobalSign Atlas account encoded in base64, either mtls-key-file-path or mtls-data-base64 must be supplied | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateGlobalSignAtlas

`func NewTargetUpdateGlobalSignAtlas(apiKey string, apiSecret string, name string, ) *TargetUpdateGlobalSignAtlas`

NewTargetUpdateGlobalSignAtlas instantiates a new TargetUpdateGlobalSignAtlas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateGlobalSignAtlasWithDefaults

`func NewTargetUpdateGlobalSignAtlasWithDefaults() *TargetUpdateGlobalSignAtlas`

NewTargetUpdateGlobalSignAtlasWithDefaults instantiates a new TargetUpdateGlobalSignAtlas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TargetUpdateGlobalSignAtlas) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TargetUpdateGlobalSignAtlas) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TargetUpdateGlobalSignAtlas) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiSecret

`func (o *TargetUpdateGlobalSignAtlas) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *TargetUpdateGlobalSignAtlas) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *TargetUpdateGlobalSignAtlas) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.


### GetDescription

`func (o *TargetUpdateGlobalSignAtlas) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateGlobalSignAtlas) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateGlobalSignAtlas) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateGlobalSignAtlas) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateGlobalSignAtlas) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateGlobalSignAtlas) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateGlobalSignAtlas) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateGlobalSignAtlas) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateGlobalSignAtlas) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateGlobalSignAtlas) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateGlobalSignAtlas) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateGlobalSignAtlas) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateGlobalSignAtlas) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateGlobalSignAtlas) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateGlobalSignAtlas) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateGlobalSignAtlas) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateGlobalSignAtlas) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateGlobalSignAtlas) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateGlobalSignAtlas) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateGlobalSignAtlas) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetMtlsCertDataBase64

`func (o *TargetUpdateGlobalSignAtlas) GetMtlsCertDataBase64() string`

GetMtlsCertDataBase64 returns the MtlsCertDataBase64 field if non-nil, zero value otherwise.

### GetMtlsCertDataBase64Ok

`func (o *TargetUpdateGlobalSignAtlas) GetMtlsCertDataBase64Ok() (*string, bool)`

GetMtlsCertDataBase64Ok returns a tuple with the MtlsCertDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsCertDataBase64

`func (o *TargetUpdateGlobalSignAtlas) SetMtlsCertDataBase64(v string)`

SetMtlsCertDataBase64 sets MtlsCertDataBase64 field to given value.

### HasMtlsCertDataBase64

`func (o *TargetUpdateGlobalSignAtlas) HasMtlsCertDataBase64() bool`

HasMtlsCertDataBase64 returns a boolean if a field has been set.

### GetMtlsKeyDataBase64

`func (o *TargetUpdateGlobalSignAtlas) GetMtlsKeyDataBase64() string`

GetMtlsKeyDataBase64 returns the MtlsKeyDataBase64 field if non-nil, zero value otherwise.

### GetMtlsKeyDataBase64Ok

`func (o *TargetUpdateGlobalSignAtlas) GetMtlsKeyDataBase64Ok() (*string, bool)`

GetMtlsKeyDataBase64Ok returns a tuple with the MtlsKeyDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsKeyDataBase64

`func (o *TargetUpdateGlobalSignAtlas) SetMtlsKeyDataBase64(v string)`

SetMtlsKeyDataBase64 sets MtlsKeyDataBase64 field to given value.

### HasMtlsKeyDataBase64

`func (o *TargetUpdateGlobalSignAtlas) HasMtlsKeyDataBase64() bool`

HasMtlsKeyDataBase64 returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateGlobalSignAtlas) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateGlobalSignAtlas) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateGlobalSignAtlas) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateGlobalSignAtlas) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateGlobalSignAtlas) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateGlobalSignAtlas) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateGlobalSignAtlas) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetTimeout

`func (o *TargetUpdateGlobalSignAtlas) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetUpdateGlobalSignAtlas) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetUpdateGlobalSignAtlas) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetUpdateGlobalSignAtlas) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateGlobalSignAtlas) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateGlobalSignAtlas) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateGlobalSignAtlas) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateGlobalSignAtlas) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateGlobalSignAtlas) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateGlobalSignAtlas) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateGlobalSignAtlas) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateGlobalSignAtlas) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


