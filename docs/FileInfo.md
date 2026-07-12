# FileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobStorageKey** | Pointer to **string** |  | [optional] 
**ClientFileSha256** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**DerivationData** | Pointer to **string** |  | [optional] 
**EncryptedBlobSha256** | Pointer to **string** |  | [optional] 
**EncryptedSizeBytes** | Pointer to **int64** |  | [optional] 
**EtagAtComplete** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **string** |  | [optional] 
**FileId** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**PlainSizeBytes** | Pointer to **int64** |  | [optional] 
**PreviousActiveFileInfo** | Pointer to [**FileInfo**](FileInfo.md) |  | [optional] 
**PreviousBlobStorageKey** | Pointer to **string** |  | [optional] 
**ProtectionKeyName** | Pointer to **string** |  | [optional] 
**ProtectionKeyVersion** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TransactionOwner** | Pointer to **string** |  | [optional] 

## Methods

### NewFileInfo

`func NewFileInfo() *FileInfo`

NewFileInfo instantiates a new FileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInfoWithDefaults

`func NewFileInfoWithDefaults() *FileInfo`

NewFileInfoWithDefaults instantiates a new FileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobStorageKey

`func (o *FileInfo) GetBlobStorageKey() string`

GetBlobStorageKey returns the BlobStorageKey field if non-nil, zero value otherwise.

### GetBlobStorageKeyOk

`func (o *FileInfo) GetBlobStorageKeyOk() (*string, bool)`

GetBlobStorageKeyOk returns a tuple with the BlobStorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobStorageKey

`func (o *FileInfo) SetBlobStorageKey(v string)`

SetBlobStorageKey sets BlobStorageKey field to given value.

### HasBlobStorageKey

`func (o *FileInfo) HasBlobStorageKey() bool`

HasBlobStorageKey returns a boolean if a field has been set.

### GetClientFileSha256

`func (o *FileInfo) GetClientFileSha256() string`

GetClientFileSha256 returns the ClientFileSha256 field if non-nil, zero value otherwise.

### GetClientFileSha256Ok

`func (o *FileInfo) GetClientFileSha256Ok() (*string, bool)`

GetClientFileSha256Ok returns a tuple with the ClientFileSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFileSha256

`func (o *FileInfo) SetClientFileSha256(v string)`

SetClientFileSha256 sets ClientFileSha256 field to given value.

### HasClientFileSha256

`func (o *FileInfo) HasClientFileSha256() bool`

HasClientFileSha256 returns a boolean if a field has been set.

### GetContentType

`func (o *FileInfo) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FileInfo) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FileInfo) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FileInfo) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDerivationData

`func (o *FileInfo) GetDerivationData() string`

GetDerivationData returns the DerivationData field if non-nil, zero value otherwise.

### GetDerivationDataOk

`func (o *FileInfo) GetDerivationDataOk() (*string, bool)`

GetDerivationDataOk returns a tuple with the DerivationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationData

`func (o *FileInfo) SetDerivationData(v string)`

SetDerivationData sets DerivationData field to given value.

### HasDerivationData

`func (o *FileInfo) HasDerivationData() bool`

HasDerivationData returns a boolean if a field has been set.

### GetEncryptedBlobSha256

`func (o *FileInfo) GetEncryptedBlobSha256() string`

GetEncryptedBlobSha256 returns the EncryptedBlobSha256 field if non-nil, zero value otherwise.

### GetEncryptedBlobSha256Ok

`func (o *FileInfo) GetEncryptedBlobSha256Ok() (*string, bool)`

GetEncryptedBlobSha256Ok returns a tuple with the EncryptedBlobSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedBlobSha256

`func (o *FileInfo) SetEncryptedBlobSha256(v string)`

SetEncryptedBlobSha256 sets EncryptedBlobSha256 field to given value.

### HasEncryptedBlobSha256

`func (o *FileInfo) HasEncryptedBlobSha256() bool`

HasEncryptedBlobSha256 returns a boolean if a field has been set.

### GetEncryptedSizeBytes

`func (o *FileInfo) GetEncryptedSizeBytes() int64`

GetEncryptedSizeBytes returns the EncryptedSizeBytes field if non-nil, zero value otherwise.

### GetEncryptedSizeBytesOk

`func (o *FileInfo) GetEncryptedSizeBytesOk() (*int64, bool)`

GetEncryptedSizeBytesOk returns a tuple with the EncryptedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSizeBytes

`func (o *FileInfo) SetEncryptedSizeBytes(v int64)`

SetEncryptedSizeBytes sets EncryptedSizeBytes field to given value.

### HasEncryptedSizeBytes

`func (o *FileInfo) HasEncryptedSizeBytes() bool`

HasEncryptedSizeBytes returns a boolean if a field has been set.

### GetEtagAtComplete

`func (o *FileInfo) GetEtagAtComplete() string`

GetEtagAtComplete returns the EtagAtComplete field if non-nil, zero value otherwise.

### GetEtagAtCompleteOk

`func (o *FileInfo) GetEtagAtCompleteOk() (*string, bool)`

GetEtagAtCompleteOk returns a tuple with the EtagAtComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtagAtComplete

`func (o *FileInfo) SetEtagAtComplete(v string)`

SetEtagAtComplete sets EtagAtComplete field to given value.

### HasEtagAtComplete

`func (o *FileInfo) HasEtagAtComplete() bool`

HasEtagAtComplete returns a boolean if a field has been set.

### GetExtension

`func (o *FileInfo) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *FileInfo) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *FileInfo) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *FileInfo) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFileId

`func (o *FileInfo) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileInfo) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileInfo) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *FileInfo) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetFilename

`func (o *FileInfo) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FileInfo) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FileInfo) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FileInfo) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetPlainSizeBytes

`func (o *FileInfo) GetPlainSizeBytes() int64`

GetPlainSizeBytes returns the PlainSizeBytes field if non-nil, zero value otherwise.

### GetPlainSizeBytesOk

`func (o *FileInfo) GetPlainSizeBytesOk() (*int64, bool)`

GetPlainSizeBytesOk returns a tuple with the PlainSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlainSizeBytes

`func (o *FileInfo) SetPlainSizeBytes(v int64)`

SetPlainSizeBytes sets PlainSizeBytes field to given value.

### HasPlainSizeBytes

`func (o *FileInfo) HasPlainSizeBytes() bool`

HasPlainSizeBytes returns a boolean if a field has been set.

### GetPreviousActiveFileInfo

`func (o *FileInfo) GetPreviousActiveFileInfo() FileInfo`

GetPreviousActiveFileInfo returns the PreviousActiveFileInfo field if non-nil, zero value otherwise.

### GetPreviousActiveFileInfoOk

`func (o *FileInfo) GetPreviousActiveFileInfoOk() (*FileInfo, bool)`

GetPreviousActiveFileInfoOk returns a tuple with the PreviousActiveFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousActiveFileInfo

`func (o *FileInfo) SetPreviousActiveFileInfo(v FileInfo)`

SetPreviousActiveFileInfo sets PreviousActiveFileInfo field to given value.

### HasPreviousActiveFileInfo

`func (o *FileInfo) HasPreviousActiveFileInfo() bool`

HasPreviousActiveFileInfo returns a boolean if a field has been set.

### GetPreviousBlobStorageKey

`func (o *FileInfo) GetPreviousBlobStorageKey() string`

GetPreviousBlobStorageKey returns the PreviousBlobStorageKey field if non-nil, zero value otherwise.

### GetPreviousBlobStorageKeyOk

`func (o *FileInfo) GetPreviousBlobStorageKeyOk() (*string, bool)`

GetPreviousBlobStorageKeyOk returns a tuple with the PreviousBlobStorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlobStorageKey

`func (o *FileInfo) SetPreviousBlobStorageKey(v string)`

SetPreviousBlobStorageKey sets PreviousBlobStorageKey field to given value.

### HasPreviousBlobStorageKey

`func (o *FileInfo) HasPreviousBlobStorageKey() bool`

HasPreviousBlobStorageKey returns a boolean if a field has been set.

### GetProtectionKeyName

`func (o *FileInfo) GetProtectionKeyName() string`

GetProtectionKeyName returns the ProtectionKeyName field if non-nil, zero value otherwise.

### GetProtectionKeyNameOk

`func (o *FileInfo) GetProtectionKeyNameOk() (*string, bool)`

GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyName

`func (o *FileInfo) SetProtectionKeyName(v string)`

SetProtectionKeyName sets ProtectionKeyName field to given value.

### HasProtectionKeyName

`func (o *FileInfo) HasProtectionKeyName() bool`

HasProtectionKeyName returns a boolean if a field has been set.

### GetProtectionKeyVersion

`func (o *FileInfo) GetProtectionKeyVersion() int32`

GetProtectionKeyVersion returns the ProtectionKeyVersion field if non-nil, zero value otherwise.

### GetProtectionKeyVersionOk

`func (o *FileInfo) GetProtectionKeyVersionOk() (*int32, bool)`

GetProtectionKeyVersionOk returns a tuple with the ProtectionKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyVersion

`func (o *FileInfo) SetProtectionKeyVersion(v int32)`

SetProtectionKeyVersion sets ProtectionKeyVersion field to given value.

### HasProtectionKeyVersion

`func (o *FileInfo) HasProtectionKeyVersion() bool`

HasProtectionKeyVersion returns a boolean if a field has been set.

### GetStatus

`func (o *FileInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FileInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionOwner

`func (o *FileInfo) GetTransactionOwner() string`

GetTransactionOwner returns the TransactionOwner field if non-nil, zero value otherwise.

### GetTransactionOwnerOk

`func (o *FileInfo) GetTransactionOwnerOk() (*string, bool)`

GetTransactionOwnerOk returns a tuple with the TransactionOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionOwner

`func (o *FileInfo) SetTransactionOwner(v string)`

SetTransactionOwner sets TransactionOwner field to given value.

### HasTransactionOwner

`func (o *FileInfo) HasTransactionOwner() bool`

HasTransactionOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


