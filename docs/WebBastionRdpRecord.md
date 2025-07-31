# WebBastionRdpRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**AwsStorage**](AwsStorage.md) |  | [optional] 
**Azure** | Pointer to [**AzureStorage**](AzureStorage.md) |  | [optional] 
**Compress** | Pointer to **bool** |  | [optional] 
**EncryptionKey** | Pointer to **string** |  | [optional] 
**RecordingQuality** | Pointer to **string** |  | [optional] 
**StorageType** | Pointer to **string** |  | [optional] 

## Methods

### NewWebBastionRdpRecord

`func NewWebBastionRdpRecord() *WebBastionRdpRecord`

NewWebBastionRdpRecord instantiates a new WebBastionRdpRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebBastionRdpRecordWithDefaults

`func NewWebBastionRdpRecordWithDefaults() *WebBastionRdpRecord`

NewWebBastionRdpRecordWithDefaults instantiates a new WebBastionRdpRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *WebBastionRdpRecord) GetAws() AwsStorage`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *WebBastionRdpRecord) GetAwsOk() (*AwsStorage, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *WebBastionRdpRecord) SetAws(v AwsStorage)`

SetAws sets Aws field to given value.

### HasAws

`func (o *WebBastionRdpRecord) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *WebBastionRdpRecord) GetAzure() AzureStorage`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *WebBastionRdpRecord) GetAzureOk() (*AzureStorage, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *WebBastionRdpRecord) SetAzure(v AzureStorage)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *WebBastionRdpRecord) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetCompress

`func (o *WebBastionRdpRecord) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *WebBastionRdpRecord) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *WebBastionRdpRecord) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *WebBastionRdpRecord) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *WebBastionRdpRecord) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *WebBastionRdpRecord) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *WebBastionRdpRecord) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *WebBastionRdpRecord) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetRecordingQuality

`func (o *WebBastionRdpRecord) GetRecordingQuality() string`

GetRecordingQuality returns the RecordingQuality field if non-nil, zero value otherwise.

### GetRecordingQualityOk

`func (o *WebBastionRdpRecord) GetRecordingQualityOk() (*string, bool)`

GetRecordingQualityOk returns a tuple with the RecordingQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingQuality

`func (o *WebBastionRdpRecord) SetRecordingQuality(v string)`

SetRecordingQuality sets RecordingQuality field to given value.

### HasRecordingQuality

`func (o *WebBastionRdpRecord) HasRecordingQuality() bool`

HasRecordingQuality returns a boolean if a field has been set.

### GetStorageType

`func (o *WebBastionRdpRecord) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *WebBastionRdpRecord) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *WebBastionRdpRecord) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *WebBastionRdpRecord) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


