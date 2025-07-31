# GatewayUpdateRemoteAccessRdpRecordings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsStorageAccessKeyId** | Pointer to **string** | AWS access key id. For more information refer to https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html | [optional] 
**AwsStorageBucketName** | Pointer to **string** | The AWS bucket name. For more information refer to https://docs.aws.amazon.com/s3/ | [optional] 
**AwsStorageBucketPrefix** | Pointer to **string** | The folder name in S3 bucket. For more information refer to https://docs.aws.amazon.com/s3/ | [optional] 
**AwsStorageRegion** | Pointer to **string** | The region where the storage is located | [optional] 
**AwsStorageSecretAccessKey** | Pointer to **string** | AWS secret access key. For more information refer to https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html | [optional] 
**AzureStorageAccountName** | Pointer to **string** | Azure account name. For more information refer to https://learn.microsoft.com/en-us/azure/storage/common/storage-account-overview | [optional] 
**AzureStorageClientId** | Pointer to **string** | Azure client id. For more information refer to https://learn.microsoft.com/en-us/azure/storage/common/storage-account-get-info?tabs&#x3D;portal | [optional] 
**AzureStorageClientSecret** | Pointer to **string** | Azure client secret. For more information refer to https://learn.microsoft.com/en-us/azure/storage/common/storage-account-get-info?tabs&#x3D;portal | [optional] 
**AzureStorageContainerName** | Pointer to **string** | Azure container name. For more information refer to https://learn.microsoft.com/en-us/rest/api/storageservices/naming-and-referencing-containers--blobs--and-metadata | [optional] 
**AzureStorageTenantId** | Pointer to **string** | Azure tenant id. For more information refer to https://learn.microsoft.com/en-us/entra/fundamentals/how-to-find-tenant | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**RdpSessionRecording** | Pointer to **string** | Enable recording of rdp session [true/false] | [optional] 
**RdpSessionRecordingCompress** | Pointer to **bool** | Whether to compress recording files before upload | [optional] 
**RdpSessionRecordingEncryptionKey** | Pointer to **string** | If provided, this key will be used to encrypt uploaded recordings. | [optional] 
**RdpSessionRecordingQuality** | Pointer to **string** | RDP session recording quality [low/medium/high] | [optional] 
**RdpSessionStorage** | Pointer to **string** | Rdp session recording storage destination [local/aws/azure] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateRemoteAccessRdpRecordings

`func NewGatewayUpdateRemoteAccessRdpRecordings() *GatewayUpdateRemoteAccessRdpRecordings`

NewGatewayUpdateRemoteAccessRdpRecordings instantiates a new GatewayUpdateRemoteAccessRdpRecordings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateRemoteAccessRdpRecordingsWithDefaults

`func NewGatewayUpdateRemoteAccessRdpRecordingsWithDefaults() *GatewayUpdateRemoteAccessRdpRecordings`

NewGatewayUpdateRemoteAccessRdpRecordingsWithDefaults instantiates a new GatewayUpdateRemoteAccessRdpRecordings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsStorageAccessKeyId

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageAccessKeyId() string`

GetAwsStorageAccessKeyId returns the AwsStorageAccessKeyId field if non-nil, zero value otherwise.

### GetAwsStorageAccessKeyIdOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageAccessKeyIdOk() (*string, bool)`

GetAwsStorageAccessKeyIdOk returns a tuple with the AwsStorageAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsStorageAccessKeyId

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAwsStorageAccessKeyId(v string)`

SetAwsStorageAccessKeyId sets AwsStorageAccessKeyId field to given value.

### HasAwsStorageAccessKeyId

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAwsStorageAccessKeyId() bool`

HasAwsStorageAccessKeyId returns a boolean if a field has been set.

### GetAwsStorageBucketName

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageBucketName() string`

GetAwsStorageBucketName returns the AwsStorageBucketName field if non-nil, zero value otherwise.

### GetAwsStorageBucketNameOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageBucketNameOk() (*string, bool)`

GetAwsStorageBucketNameOk returns a tuple with the AwsStorageBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsStorageBucketName

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAwsStorageBucketName(v string)`

SetAwsStorageBucketName sets AwsStorageBucketName field to given value.

### HasAwsStorageBucketName

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAwsStorageBucketName() bool`

HasAwsStorageBucketName returns a boolean if a field has been set.

### GetAwsStorageBucketPrefix

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageBucketPrefix() string`

GetAwsStorageBucketPrefix returns the AwsStorageBucketPrefix field if non-nil, zero value otherwise.

### GetAwsStorageBucketPrefixOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageBucketPrefixOk() (*string, bool)`

GetAwsStorageBucketPrefixOk returns a tuple with the AwsStorageBucketPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsStorageBucketPrefix

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAwsStorageBucketPrefix(v string)`

SetAwsStorageBucketPrefix sets AwsStorageBucketPrefix field to given value.

### HasAwsStorageBucketPrefix

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAwsStorageBucketPrefix() bool`

HasAwsStorageBucketPrefix returns a boolean if a field has been set.

### GetAwsStorageRegion

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageRegion() string`

GetAwsStorageRegion returns the AwsStorageRegion field if non-nil, zero value otherwise.

### GetAwsStorageRegionOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageRegionOk() (*string, bool)`

GetAwsStorageRegionOk returns a tuple with the AwsStorageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsStorageRegion

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAwsStorageRegion(v string)`

SetAwsStorageRegion sets AwsStorageRegion field to given value.

### HasAwsStorageRegion

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAwsStorageRegion() bool`

HasAwsStorageRegion returns a boolean if a field has been set.

### GetAwsStorageSecretAccessKey

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageSecretAccessKey() string`

GetAwsStorageSecretAccessKey returns the AwsStorageSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsStorageSecretAccessKeyOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAwsStorageSecretAccessKeyOk() (*string, bool)`

GetAwsStorageSecretAccessKeyOk returns a tuple with the AwsStorageSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsStorageSecretAccessKey

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAwsStorageSecretAccessKey(v string)`

SetAwsStorageSecretAccessKey sets AwsStorageSecretAccessKey field to given value.

### HasAwsStorageSecretAccessKey

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAwsStorageSecretAccessKey() bool`

HasAwsStorageSecretAccessKey returns a boolean if a field has been set.

### GetAzureStorageAccountName

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageAccountName() string`

GetAzureStorageAccountName returns the AzureStorageAccountName field if non-nil, zero value otherwise.

### GetAzureStorageAccountNameOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageAccountNameOk() (*string, bool)`

GetAzureStorageAccountNameOk returns a tuple with the AzureStorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorageAccountName

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAzureStorageAccountName(v string)`

SetAzureStorageAccountName sets AzureStorageAccountName field to given value.

### HasAzureStorageAccountName

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAzureStorageAccountName() bool`

HasAzureStorageAccountName returns a boolean if a field has been set.

### GetAzureStorageClientId

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageClientId() string`

GetAzureStorageClientId returns the AzureStorageClientId field if non-nil, zero value otherwise.

### GetAzureStorageClientIdOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageClientIdOk() (*string, bool)`

GetAzureStorageClientIdOk returns a tuple with the AzureStorageClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorageClientId

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAzureStorageClientId(v string)`

SetAzureStorageClientId sets AzureStorageClientId field to given value.

### HasAzureStorageClientId

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAzureStorageClientId() bool`

HasAzureStorageClientId returns a boolean if a field has been set.

### GetAzureStorageClientSecret

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageClientSecret() string`

GetAzureStorageClientSecret returns the AzureStorageClientSecret field if non-nil, zero value otherwise.

### GetAzureStorageClientSecretOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageClientSecretOk() (*string, bool)`

GetAzureStorageClientSecretOk returns a tuple with the AzureStorageClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorageClientSecret

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAzureStorageClientSecret(v string)`

SetAzureStorageClientSecret sets AzureStorageClientSecret field to given value.

### HasAzureStorageClientSecret

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAzureStorageClientSecret() bool`

HasAzureStorageClientSecret returns a boolean if a field has been set.

### GetAzureStorageContainerName

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageContainerName() string`

GetAzureStorageContainerName returns the AzureStorageContainerName field if non-nil, zero value otherwise.

### GetAzureStorageContainerNameOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageContainerNameOk() (*string, bool)`

GetAzureStorageContainerNameOk returns a tuple with the AzureStorageContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorageContainerName

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAzureStorageContainerName(v string)`

SetAzureStorageContainerName sets AzureStorageContainerName field to given value.

### HasAzureStorageContainerName

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAzureStorageContainerName() bool`

HasAzureStorageContainerName returns a boolean if a field has been set.

### GetAzureStorageTenantId

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageTenantId() string`

GetAzureStorageTenantId returns the AzureStorageTenantId field if non-nil, zero value otherwise.

### GetAzureStorageTenantIdOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetAzureStorageTenantIdOk() (*string, bool)`

GetAzureStorageTenantIdOk returns a tuple with the AzureStorageTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorageTenantId

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetAzureStorageTenantId(v string)`

SetAzureStorageTenantId sets AzureStorageTenantId field to given value.

### HasAzureStorageTenantId

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasAzureStorageTenantId() bool`

HasAzureStorageTenantId returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetRdpSessionRecording

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionRecording() string`

GetRdpSessionRecording returns the RdpSessionRecording field if non-nil, zero value otherwise.

### GetRdpSessionRecordingOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionRecordingOk() (*string, bool)`

GetRdpSessionRecordingOk returns a tuple with the RdpSessionRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpSessionRecording

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetRdpSessionRecording(v string)`

SetRdpSessionRecording sets RdpSessionRecording field to given value.

### HasRdpSessionRecording

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasRdpSessionRecording() bool`

HasRdpSessionRecording returns a boolean if a field has been set.

### GetRdpSessionRecordingCompress

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionRecordingCompress() bool`

GetRdpSessionRecordingCompress returns the RdpSessionRecordingCompress field if non-nil, zero value otherwise.

### GetRdpSessionRecordingCompressOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionRecordingCompressOk() (*bool, bool)`

GetRdpSessionRecordingCompressOk returns a tuple with the RdpSessionRecordingCompress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpSessionRecordingCompress

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetRdpSessionRecordingCompress(v bool)`

SetRdpSessionRecordingCompress sets RdpSessionRecordingCompress field to given value.

### HasRdpSessionRecordingCompress

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasRdpSessionRecordingCompress() bool`

HasRdpSessionRecordingCompress returns a boolean if a field has been set.

### GetRdpSessionRecordingEncryptionKey

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionRecordingEncryptionKey() string`

GetRdpSessionRecordingEncryptionKey returns the RdpSessionRecordingEncryptionKey field if non-nil, zero value otherwise.

### GetRdpSessionRecordingEncryptionKeyOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionRecordingEncryptionKeyOk() (*string, bool)`

GetRdpSessionRecordingEncryptionKeyOk returns a tuple with the RdpSessionRecordingEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpSessionRecordingEncryptionKey

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetRdpSessionRecordingEncryptionKey(v string)`

SetRdpSessionRecordingEncryptionKey sets RdpSessionRecordingEncryptionKey field to given value.

### HasRdpSessionRecordingEncryptionKey

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasRdpSessionRecordingEncryptionKey() bool`

HasRdpSessionRecordingEncryptionKey returns a boolean if a field has been set.

### GetRdpSessionRecordingQuality

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionRecordingQuality() string`

GetRdpSessionRecordingQuality returns the RdpSessionRecordingQuality field if non-nil, zero value otherwise.

### GetRdpSessionRecordingQualityOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionRecordingQualityOk() (*string, bool)`

GetRdpSessionRecordingQualityOk returns a tuple with the RdpSessionRecordingQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpSessionRecordingQuality

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetRdpSessionRecordingQuality(v string)`

SetRdpSessionRecordingQuality sets RdpSessionRecordingQuality field to given value.

### HasRdpSessionRecordingQuality

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasRdpSessionRecordingQuality() bool`

HasRdpSessionRecordingQuality returns a boolean if a field has been set.

### GetRdpSessionStorage

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionStorage() string`

GetRdpSessionStorage returns the RdpSessionStorage field if non-nil, zero value otherwise.

### GetRdpSessionStorageOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetRdpSessionStorageOk() (*string, bool)`

GetRdpSessionStorageOk returns a tuple with the RdpSessionStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpSessionStorage

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetRdpSessionStorage(v string)`

SetRdpSessionStorage sets RdpSessionStorage field to given value.

### HasRdpSessionStorage

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasRdpSessionStorage() bool`

HasRdpSessionStorage returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateRemoteAccessRdpRecordings) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateRemoteAccessRdpRecordings) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateRemoteAccessRdpRecordings) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


