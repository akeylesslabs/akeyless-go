# AzureStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**StorageAccount** | Pointer to **string** |  | [optional] 
**StorageContainerName** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** | creds | [optional] 

## Methods

### NewAzureStorage

`func NewAzureStorage() *AzureStorage`

NewAzureStorage instantiates a new AzureStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageWithDefaults

`func NewAzureStorageWithDefaults() *AzureStorage`

NewAzureStorageWithDefaults instantiates a new AzureStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *AzureStorage) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AzureStorage) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AzureStorage) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AzureStorage) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *AzureStorage) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureStorage) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureStorage) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AzureStorage) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AzureStorage) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureStorage) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureStorage) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AzureStorage) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetStorageAccount

`func (o *AzureStorage) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AzureStorage) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AzureStorage) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *AzureStorage) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetStorageContainerName

`func (o *AzureStorage) GetStorageContainerName() string`

GetStorageContainerName returns the StorageContainerName field if non-nil, zero value otherwise.

### GetStorageContainerNameOk

`func (o *AzureStorage) GetStorageContainerNameOk() (*string, bool)`

GetStorageContainerNameOk returns a tuple with the StorageContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerName

`func (o *AzureStorage) SetStorageContainerName(v string)`

SetStorageContainerName sets StorageContainerName field to given value.

### HasStorageContainerName

`func (o *AzureStorage) HasStorageContainerName() bool`

HasStorageContainerName returns a boolean if a field has been set.

### GetTenantId

`func (o *AzureStorage) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureStorage) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureStorage) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureStorage) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


