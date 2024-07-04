# TargetCreateGlobalSignAtlas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | API Key of the GlobalSign Atlas account | 
**ApiSecret** | **string** | API Secret of the GlobalSign Atlas account | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**MtlsCertDataBase64** | Pointer to **string** | Mutual TLS Certificate contents of the GlobalSign Atlas account encoded in base64, either mtls-cert-file-path or mtls-cert-data-base64 must be supplied | [optional] 
**MtlsKeyDataBase64** | Pointer to **string** | Mutual TLS Key contents of the GlobalSign Atlas account encoded in base64, either mtls-key-file-path or mtls-data-base64 must be supplied | [optional] 
**Name** | **string** | Target name | 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateGlobalSignAtlas

`func NewTargetCreateGlobalSignAtlas(apiKey string, apiSecret string, name string, ) *TargetCreateGlobalSignAtlas`

NewTargetCreateGlobalSignAtlas instantiates a new TargetCreateGlobalSignAtlas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateGlobalSignAtlasWithDefaults

`func NewTargetCreateGlobalSignAtlasWithDefaults() *TargetCreateGlobalSignAtlas`

NewTargetCreateGlobalSignAtlasWithDefaults instantiates a new TargetCreateGlobalSignAtlas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TargetCreateGlobalSignAtlas) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TargetCreateGlobalSignAtlas) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TargetCreateGlobalSignAtlas) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiSecret

`func (o *TargetCreateGlobalSignAtlas) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *TargetCreateGlobalSignAtlas) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *TargetCreateGlobalSignAtlas) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.


### GetDescription

`func (o *TargetCreateGlobalSignAtlas) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateGlobalSignAtlas) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateGlobalSignAtlas) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateGlobalSignAtlas) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateGlobalSignAtlas) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateGlobalSignAtlas) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateGlobalSignAtlas) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateGlobalSignAtlas) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateGlobalSignAtlas) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateGlobalSignAtlas) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateGlobalSignAtlas) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateGlobalSignAtlas) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateGlobalSignAtlas) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateGlobalSignAtlas) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateGlobalSignAtlas) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateGlobalSignAtlas) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetMtlsCertDataBase64

`func (o *TargetCreateGlobalSignAtlas) GetMtlsCertDataBase64() string`

GetMtlsCertDataBase64 returns the MtlsCertDataBase64 field if non-nil, zero value otherwise.

### GetMtlsCertDataBase64Ok

`func (o *TargetCreateGlobalSignAtlas) GetMtlsCertDataBase64Ok() (*string, bool)`

GetMtlsCertDataBase64Ok returns a tuple with the MtlsCertDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsCertDataBase64

`func (o *TargetCreateGlobalSignAtlas) SetMtlsCertDataBase64(v string)`

SetMtlsCertDataBase64 sets MtlsCertDataBase64 field to given value.

### HasMtlsCertDataBase64

`func (o *TargetCreateGlobalSignAtlas) HasMtlsCertDataBase64() bool`

HasMtlsCertDataBase64 returns a boolean if a field has been set.

### GetMtlsKeyDataBase64

`func (o *TargetCreateGlobalSignAtlas) GetMtlsKeyDataBase64() string`

GetMtlsKeyDataBase64 returns the MtlsKeyDataBase64 field if non-nil, zero value otherwise.

### GetMtlsKeyDataBase64Ok

`func (o *TargetCreateGlobalSignAtlas) GetMtlsKeyDataBase64Ok() (*string, bool)`

GetMtlsKeyDataBase64Ok returns a tuple with the MtlsKeyDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsKeyDataBase64

`func (o *TargetCreateGlobalSignAtlas) SetMtlsKeyDataBase64(v string)`

SetMtlsKeyDataBase64 sets MtlsKeyDataBase64 field to given value.

### HasMtlsKeyDataBase64

`func (o *TargetCreateGlobalSignAtlas) HasMtlsKeyDataBase64() bool`

HasMtlsKeyDataBase64 returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateGlobalSignAtlas) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateGlobalSignAtlas) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateGlobalSignAtlas) SetName(v string)`

SetName sets Name field to given value.


### GetTimeout

`func (o *TargetCreateGlobalSignAtlas) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetCreateGlobalSignAtlas) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetCreateGlobalSignAtlas) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetCreateGlobalSignAtlas) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateGlobalSignAtlas) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateGlobalSignAtlas) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateGlobalSignAtlas) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateGlobalSignAtlas) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateGlobalSignAtlas) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateGlobalSignAtlas) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateGlobalSignAtlas) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateGlobalSignAtlas) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


