# WebBastionRdpRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**AwsStorage**](AwsStorage.md) |  | [optional] 
**Azure** | Pointer to [**AzureStorage**](AzureStorage.md) |  | [optional] 
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


